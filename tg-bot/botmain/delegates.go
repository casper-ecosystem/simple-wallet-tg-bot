package botmain

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/Simplewallethq/tg-bot/botmain/crypto"
	"github.com/Simplewallethq/tg-bot/ent"
	"github.com/Simplewallethq/tg-bot/ent/user"
	entval "github.com/Simplewallethq/tg-bot/ent/validators"
	pb "github.com/Simplewallethq/tg-bot/tggateway/proto"
	tggateway "github.com/Simplewallethq/tg-bot/tggateway/types"
	"github.com/make-software/casper-go-sdk/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

const maxDelegators = 1200

func (b *BotMain) HandleDelegateState(msg *pb.TgTextMessage) error {
	log.Println("received delegate state message", msg)
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 || state[0] != "Delegate" {
		err = b.State.DeleteUserState(msg.GetFrom().GetId())
		if err != nil {
			return errors.Wrap(err, "failed delete user state")
		}
		return errors.Wrap(err, "bad get user state")
	}

	switch state[2] {
	case "sign":
		err := b.SignAndPutDelegate(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle transfer amount")
		}
	case "askAmount":
		err := b.PickDelegateAmountCustom(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle transfer amount")
		}
	}
	return nil
}

func (b *BotMain) HandleNewDelegateButton(msg tggateway.TgMessageMsg) error {
	out := pb.NewDelegateButton{}
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
	balance, err := b.Restclient.GetBalance(b.RPCNode, u.PublicKey)
	if err != nil {
		return errors.Wrap(err, "failed get balance")
	}

	balanceRes := strconv.FormatFloat(balance, 'f', 6, 64)
	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) == 0 || (len(state) > 0 && state[0] != "Delegate") {
		delDB, err := b.DB.Delegates.Create().SetDelegator(u.PublicKey).SetUserBalance(balanceRes).SetOwner(u).Save(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed create transfer")
		}
		log.Println(delDB.ID)
		err = b.State.SetUserState(out.GetUser().GetId(), []string{"Delegate", strconv.Itoa(delDB.ID)})
		if err != nil {
			return errors.Wrap(err, "failed set user state")
		}
	}

	ValDB, err := b.DB.Validators.Query().Where(entval.And(entval.Active(true),
		entval.DelegatorsLT(maxDelegators))).
		Order(ent.Asc(entval.FieldFee), ent.Desc(entval.FieldDelegators)).
		Offset(int(out.GetOffset())).
		Limit(5).
		All(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get validators")
	}

	count, err := b.DB.Validators.Query().Where(entval.And(entval.Active(true),
		entval.DelegatorsLT(maxDelegators))).Count(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get validators count")
	}

	var addresses []*pb.ValidatorsRow

	for _, val := range ValDB {
		addresses = append(addresses, &pb.ValidatorsRow{
			Address:    val.Address,
			Fee:        float64(val.Fee),
			Delegators: int32(val.Delegators),
			Id:         uint64(val.ID),
		})
	}
	data := pb.DelegateValidatorsList{
		User:        out.GetUser(),
		MsgId:       out.GetMsgId(),
		Offset:      out.GetOffset(),
		Validators:  addresses,
		Total:       int64(count),
		UserBalance: balanceRes,
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "DelegateListValidators",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) PickDelegateValidator(msg tggateway.TgMessageMsg) error {
	out := pb.PickDelegateValidator{}

	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	// u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	// if err != nil {
	// 	return errors.Wrap(err, "failed get user")
	// }

	Validator, err := b.DB.Validators.Get(context.Background(), int(out.GetId()))
	if err != nil {
		return errors.Wrap(err, "failed get address book")
	}

	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	DelDB, err := b.DB.Delegates.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}

	DelDB, err = DelDB.Update().SetValidator(Validator.Address).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update transfer")
	}

	uBalance, err := strconv.ParseFloat(DelDB.UserBalance, 64)
	if err != nil {
		return errors.Wrap(err, "failed parse float")
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

	data := pb.DelegateAskAmount{
		User:        out.GetUser(),
		MsgId:       out.GetMsgId(),
		UserBalance: uBalance,
		Validator:   DelDB.Validator,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "DelegateAskAmount",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) PickDelegateAmount(msg tggateway.TgMessageMsg) error {
	out := pb.DelegatePickAmount{}

	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	state, err := b.State.GetUserState(out.GetUser().GetId())
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

	DelDB, err := b.DB.Delegates.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}

	DelDB, err = DelDB.Update().SetAmount(fmt.Sprintf("%f", out.GetAmount())).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update transfer")
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
	data := pb.DelegateConfirmation{
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
		Name: "DelegateAskConfirmation",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) PickDelegateAmountCustom(msg *pb.TgTextMessage) error {
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

	DelDB, err := b.DB.Delegates.Get(context.Background(), trID)
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
	data := pb.DelegateConfirmation{
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
		Name: "DelegateAskConfirmation",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) DelegateConfirmButton(msg tggateway.TgMessageMsg) error {
	out := pb.DelegateConfirmButton{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 && state[0] != "Delegate" {
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

func (b *BotMain) SignAndPutDelegate(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 && state[0] != "Delegate" {
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

	DelDB, err := b.DB.Delegates.Get(context.Background(), trID)
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
		dep, err = b.Crypto.SignDelegateWithPassword(delegate, msg.GetFrom().GetId(), userPass)
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
		dep, err = b.Crypto.SignDelegateWithPK(delegate, msg.GetFrom().GetId(), userPass)
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

	data := pb.DelegateSuccesResponse{
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
		Name: "DelegateSuccessResponse",
		Data: dataBytes,
	}

	err = b.State.DeleteUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed delete user state")
	}

	if !u.EnableLogging {
		err := b.DB.Delegates.DeleteOneID(trID).Exec(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed delete delegate")
		}
	}

	return nil
}

func (b *BotMain) NewDepositButton(msg tggateway.TgMessageMsg) error {
	out := pb.NewDepositButton{}
	err := proto.Unmarshal(msg.Data, &out)
	if err != nil {
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

	data := pb.DepositResponse{
		User:    out.GetUser(),
		Address: u.PublicKey,
		MsgId:   out.GetMsgId(),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "DepositMessage",
		Data: dataBytes,
	}
	return nil

}
