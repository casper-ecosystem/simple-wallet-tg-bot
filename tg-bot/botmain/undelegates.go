package botmain

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/botmain/crypto"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	tggateway "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	"github.com/make-software/casper-go-sdk/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func (b *BotMain) HandleUndelegateState(msg *pb.TgTextMessage) error {
	log.Println("received undelegate state message", msg)
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 || state[0] != "Undelegate" {
		err = b.State.DeleteUserState(msg.GetFrom().GetId())
		if err != nil {
			return errors.Wrap(err, "failed delete user state")
		}
		return errors.Wrap(err, "bad get user state")
	}

	switch state[2] {
	case "sign":
		err := b.SignAndPutUndelegate(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle transfer amount")
		}
	case "askAmount":
		err := b.PickUndelegateAmountCustom(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle transfer amount")
		}
	}
	return nil
}

func (b *BotMain) HandleNewUnelegateButton(msg tggateway.TgMessageMsg) error {
	out := pb.NewUndelegateButton{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	loggedin, err := b.CheckLogin(out.GetUser())
	if err != nil {
		return err
	}
	if !loggedin {
		return nil
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}

	staked, err := b.Restclient.GetBalanceDelegated(b.RPCNode, u.PublicKey)
	if err != nil {
		return errors.Wrap(err, "failed get staked balance")
	}

	delDB, err := b.DB.Undelegates.Create().SetDelegator(u.PublicKey).SetOwner(u).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed create transfer")
	}
	log.Println(delDB.ID)

	err = b.State.SetUserState(out.GetUser().GetId(), []string{"Undelegate", strconv.Itoa(delDB.ID)})
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}

	var addresses []*pb.DelegatesRow

	for i, val := range staked.Data {
		balanceRes := strconv.FormatFloat(val.Amount, 'f', 6, 64)
		addresses = append(addresses, &pb.DelegatesRow{
			Address: val.Validator,
			Amount:  balanceRes,
			Id:      uint64(i),
		})
	}
	data := pb.DelegatesList{
		User:      out.GetUser(),
		MsgId:     out.GetMsgId(),
		Offset:    out.GetOffset(),
		Delegates: addresses,
		Total:     int64(len(staked.Data)),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "UndelegateDelegatesList",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) PickUndelegateValidator(msg tggateway.TgMessageMsg) error {
	out := pb.PickUndelegateValidator{}

	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}

	staked, err := b.Restclient.GetBalanceDelegated(b.RPCNode, u.PublicKey)
	if err != nil {
		return errors.Wrap(err, "failed get staked balance")
	}
	var validator string
	var stakedBalance string

	if len(staked.Data) >= int(out.GetId()) {
		validator = staked.Data[out.GetId()].Validator
		stakedBalance = strconv.FormatFloat(staked.Data[out.GetId()].Amount, 'f', 6, 64)

	} else {
		return errors.Wrap(err, "failed get validator")
	}

	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	DelDB, err := b.DB.Undelegates.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}

	DelDB, err = DelDB.Update().SetValidator(validator).SetStakedBalance(stakedBalance).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update undelegate")
	}

	if len(state) < 3 {
		state = append(state, "askAmount")
	} else {
		state[2] = "askAmount"
	}
	err = b.State.SetUserState(out.GetUser().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}

	data := pb.UndelegateAskAmount{
		User:          out.GetUser(),
		MsgId:         out.GetMsgId(),
		StakedBalance: stakedBalance,
		Validator:     DelDB.Validator,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "UndelegateAskAmount",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) PickUndelegateAmount(msg tggateway.TgMessageMsg) error {
	out := pb.UndelegatePickAmount{}

	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	DelDB, err := b.DB.Undelegates.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get undelegate")
	}

	DelDB, err = DelDB.Update().SetAmount(fmt.Sprintf("%f", out.GetAmount())).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update undelegate")
	}

	if len(state) < 3 {
		state = append(state, "confirmation")
	} else {
		state[2] = "confirmation"
	}
	err = b.State.SetUserState(out.GetUser().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}

	amountF, err := strconv.ParseFloat(DelDB.Amount, 64)
	if err != nil {
		return errors.Wrap(err, "failed parse amount")
	}
	data := pb.UndelegateConfirmation{
		User:      out.GetUser(),
		MsgId:     out.GetMsgId(),
		Validator: DelDB.Validator,
		Amount:    amountF,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "UndelegateAskConfirmation",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) PickUndelegateAmountCustom(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	// u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	// if err != nil {
	// 	return errors.Wrap(err, "failed get user")
	// }

	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	DelDB, err := b.DB.Undelegates.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}
	_, ok := new(big.Float).SetString(msg.GetText())
	var msgid int64
	if len(state) < 4 {
		msgid = -1
	} else {
		msgid, err = strconv.ParseInt(state[3], 10, 64)
		if err != nil {
			msgid = -1
		}
	}
	if !ok {
		data := pb.TransferAmountIsNotValid{
			User:  msg.GetFrom(),
			MsgId: msgid,
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}

		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "TransferAmountNotValid",
			Data: dataBytes,
		}
		return nil
	}

	DelDB, err = DelDB.Update().SetAmount(msg.GetText()).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update transfer")
	}

	if len(state) < 3 {
		state = append(state, "confirmation")
	} else {
		state[2] = "confirmation"
	}
	err = b.State.SetUserState(msg.GetFrom().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}

	amountF, err := strconv.ParseFloat(DelDB.Amount, 64)
	if err != nil {
		return errors.Wrap(err, "failed parse amount")
	}
	data := pb.UndelegateConfirmation{
		User:      msg.GetFrom(),
		MsgId:     msgid,
		Validator: DelDB.Validator,
		Amount:    amountF,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "UndelegateAskConfirmation",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) UndelegateConfirmButton(msg tggateway.TgMessageMsg) error {
	out := pb.UndelegateConfirmButton{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 && state[0] != "Undelegate" {
		err = b.State.DeleteUserState(out.GetUser().GetId())
		if err != nil {
			return errors.Wrap(err, "failed delete user state")
		}
		return errors.Wrap(err, "bad get user state")
	}

	state[2] = "sign"
	err = b.State.SetUserState(out.GetUser().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}

	if u.StorePrivatKey {
		data := pb.SignDeployAskPassword{
			User:  out.GetUser(),
			MsgId: out.MsgId,
		}

		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}

		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "SignDeployAskPassword",
			Data: dataBytes,
		}
	} else {
		data := pb.SignDeployAskPK{
			User:  out.GetUser(),
			MsgId: out.MsgId,
		}

		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}

		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "SignDeployAskPK",
			Data: dataBytes,
		}
	}

	return nil

}

func (b *BotMain) SignAndPutUndelegate(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 && state[0] != "Undelegate" {
		err = b.State.DeleteUserState(msg.GetFrom().GetId())
		if err != nil {
			return errors.Wrap(err, "failed delete user state")
		}
		return errors.Wrap(err, "bad get user state")
	}
	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	DelDB, err := b.DB.Undelegates.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}

	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}

	userPass := msg.GetText()
	//string amount to big int
	amount, ok := new(big.Float).SetString(DelDB.Amount)
	if !ok {
		log.Println("failed parse amount")
		return errors.Wrap(err, "failed parse amount")
	}
	var msgid int64
	if len(state) < 4 {
		msgid = -1
	} else {
		msgid, err = strconv.ParseInt(state[3], 10, 64)
		if err != nil {
			msgid = -1
		}
	}

	//big int balance * 100000000 (new var balance in motes)
	motesAmount := new(big.Float).Mul(amount, big.NewFloat(1000000000))
	//big float motes amount to big int
	motesAmountInt, _ := motesAmount.Int(nil)
	delegate := crypto.Delegate{
		Validator: DelDB.Validator,
		Amount:    motesAmountInt.String(),
		Delegator: DelDB.Delegator,
	}
	var dep *types.Deploy
	if u.StorePrivatKey {
		dep, err = b.Crypto.SignUndelegateWithPassword(delegate, msg.GetFrom().GetId(), userPass)
		if err != nil {
			data := pb.TransferBadPassword{
				User:  msg.GetFrom(),
				MsgId: msgid,
			}

			dataBytes, err := proto.Marshal(&data)
			if err != nil {
				return errors.Wrap(err, "failed marshal")
			}

			b.ResponseChan <- tggateway.TgResponseMsg{
				Name: "TransferBadPassword",
				Data: dataBytes,
			}
			return errors.Wrap(err, "failed sign transfer")
		}
	} else {
		dep, err = b.Crypto.SignUndelegateWithPK(delegate, msg.GetFrom().GetId(), userPass)
		if err != nil {
			data := pb.TransferBadPK{
				User:  msg.GetFrom(),
				MsgId: msgid,
			}

			dataBytes, err := proto.Marshal(&data)
			if err != nil {
				return errors.Wrap(err, "failed marshal")
			}

			b.ResponseChan <- tggateway.TgResponseMsg{
				Name: "TransferBadPK",
				Data: dataBytes,
			}
			return errors.Wrap(err, "failed sign transfer")
		}
	}

	res, err := b.Restclient.PutDeploy(b.RPCNode, *dep)
	if err != nil {
		data := pb.TransferUnknownError{
			User:  msg.GetFrom(),
			MsgId: msgid,
		}

		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}

		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "TransferUnknownError",
			Data: dataBytes,
		}
		err = b.State.DeleteUserState(msg.GetFrom().GetId())
		if err != nil {
			return errors.Wrap(err, "failed delete user state")
		}
		return errors.Wrap(err, "failed put deploy")
	}

	data := pb.UndelegateSuccesResponse{
		User:      msg.GetFrom(),
		MsgId:     msgid,
		Amount:    DelDB.Amount,
		Delegator: DelDB.Delegator,
		Validator: DelDB.Validator,
		Hash:      res,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "UndelegateSuccessResponse",
		Data: dataBytes,
	}
	if !u.EnableLogging {
		err = b.DB.Undelegates.DeleteOneID(trID).Exec(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed del undelegate")
		}
	}
	err = b.State.DeleteUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed delete user state")
	}

	return nil
}
