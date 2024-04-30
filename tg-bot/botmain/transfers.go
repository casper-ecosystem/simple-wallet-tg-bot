package botmain

import (
	"context"
	"log"
	"math/big"
	"strconv"

	"github.com/Simplewallethq/tg-bot/botmain/crypto"
	"github.com/Simplewallethq/tg-bot/ent/adressbook"
	"github.com/Simplewallethq/tg-bot/ent/recentinvoices"
	"github.com/Simplewallethq/tg-bot/ent/user"
	pb "github.com/Simplewallethq/tg-bot/tggateway/proto"
	tggateway "github.com/Simplewallethq/tg-bot/tggateway/types"
	"github.com/make-software/casper-go-sdk/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func (b *BotMain) HandleTransferState(msg *pb.TgTextMessage) error {
	log.Println("received transfers state message", msg)
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 || state[0] != "Transfer" {
		err = b.State.DeleteUserState(msg.GetFrom().GetId())
		if err != nil {
			return errors.Wrap(err, "failed delete user state")
		}
		return errors.Wrap(err, "bad get user state")
	}

	switch state[2] {
	case "askCustomAddress":
		err := b.SetCustomTransferAddress(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle transfer custom address")
		}
	case "askAmount":
		err := b.SetTransferAmount(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle transfer amount")
		}
	case "askMemo":
		err := b.SetTransferMemo(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle transfer amount")
		}
	case "sign":
		err := b.SignAndPutTransfer(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle transfer amount")
		}
	}
	return nil
}

func (b *BotMain) HandleNewTransferButton(msg tggateway.TgMessageMsg) error {
	out := pb.NewTransferButton{}
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

	transferDB, err := b.DB.Transfers.Create().SetFromPubkey(u.PublicKey).SetSenderBalance(balanceRes).SetOwner(u).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed create transfer")
	}
	log.Println(transferDB.ID)

	err = b.State.SetUserState(out.GetUser().GetId(), []string{"Transfer", strconv.Itoa(transferDB.ID)})
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}

	data := pb.NewTransferResponseStage1{
		User:    out.GetUser(),
		MsgId:   out.MsgId,
		Balance: balanceRes,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "NewTransferResponseStage1",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) HandleTransferCustomAddress(msg tggateway.TgMessageMsg) error {
	out := pb.TransferCustomAddressButton{}
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
	_, err = b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}

	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	// if len(state) != 2 && state[0] != "Transfer" {
	// 	b.State.DeleteUserState(out.GetUser().GetId())
	// 	return errors.Wrap(err, "bad get user state")
	// }

	// trID, err := strconv.Atoi(state[1])
	// if err != nil {
	// 	return errors.Wrap(err, "failed get user state")
	// }

	state = append(state, "askCustomAddress")
	state = append(state, strconv.Itoa(int(out.MsgId)))
	err = b.State.SetUserState(out.GetUser().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}

	data := pb.TransferAskCustomAddress{
		User:  out.GetUser(),
		MsgId: out.MsgId,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "TransferAskAdress",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) SetCustomTransferAddress(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 && state[0] != "Transfer" {
		err = b.State.DeleteUserState(msg.GetFrom().GetId())
		if err != nil {
			return errors.Wrap(err, "failed delete user state")
		}
		return errors.Wrap(err, "bad get user state")
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
	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	transferDB, err := b.DB.Transfers.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}

	valid, err := b.Restclient.IsAddress(b.RPCNode, msg.GetText())
	if err != nil {
		return errors.Wrap(err, "failed check address")
	}
	if !valid {
		data := pb.TransferAddressIsNotValid{
			User:  msg.GetFrom(),
			MsgId: msgid,
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}

		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "TransferAddressNotValid",
			Data: dataBytes,
		}
		return nil

	}

	transferDB, err = transferDB.Update().SetToPubkey(msg.Text).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update transfer")
	}

	uBalance, err := strconv.ParseFloat(transferDB.SenderBalance, 64)
	if err != nil {
		return errors.Wrap(err, "failed parse float")
	}

	state[2] = "askAmount"
	err = b.State.SetUserState(msg.GetFrom().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}

	data := pb.TransferAskAmount{
		User:        msg.GetFrom(),
		MsgId:       msgid,
		Recommended: uBalance - 10,
		ToPubkey:    transferDB.ToPubkey,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "TransferAskAmount",
		Data: dataBytes,
	}

	return nil
}

func (b *BotMain) SetTransferAmount(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 && state[0] != "Transfer" {
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
	var msgid int64
	if len(state) < 4 {
		msgid = -1
	} else {
		msgid, err = strconv.ParseInt(state[3], 10, 64)
		if err != nil {
			msgid = -1
		}
	}
	log.Println(state)
	transferDB, err := b.DB.Transfers.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}

	//try msg.GetText() to big float
	biAmount, ok := new(big.Float).SetString(msg.GetText())
	if !ok || (biAmount.Cmp(big.NewFloat(0))) < 0 {
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

	transferDB, err = transferDB.Update().SetAmount(msg.GetText()).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update transfer")
	}

	state[2] = "askMemo"
	err = b.State.SetUserState(msg.GetFrom().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}

	data := pb.TransferAskMemo{
		User:   msg.GetFrom(),
		MsgId:  msgid,
		Amount: transferDB.Amount,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "TransferAskMemo",
		Data: dataBytes,
	}

	return nil
}

func (b *BotMain) SetTransferMemo(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 && state[0] != "Transfer" {
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
	var msgid int64
	if len(state) < 4 {
		msgid = -1
	} else {
		msgid, err = strconv.ParseInt(state[3], 10, 64)
		if err != nil {
			msgid = -1
		}
	}

	transferDB, err := b.DB.Transfers.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}

	//check we can convert msg.GetText() to uint64

	memo, err := strconv.ParseUint(msg.GetText(), 10, 64)
	if err != nil {
		data := pb.TransferMemoIsNotValid{
			User:  msg.GetFrom(),
			MsgId: msgid,
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}

		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "TransferMemoNotValid",
			Data: dataBytes,
		}
		return nil
	}
	transferDB, err = transferDB.Update().SetMemoID(memo).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update transfer")
	}

	state[2] = "confirmation"
	err = b.State.SetUserState(msg.GetFrom().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}

	data := pb.TransferAskConfirmation{
		User:     msg.GetFrom(),
		MsgId:    msgid,
		Amount:   transferDB.Amount,
		ToPubkey: transferDB.ToPubkey,
		Name:     transferDB.Name,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "TransferAskConfirmation",
		Data: dataBytes,
	}

	return nil
}

func (b *BotMain) HandleTransferConfirmButton(msg tggateway.TgMessageMsg) error {
	out := pb.TransferConfirmButton{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 && state[0] != "Transfer" {
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

func (b *BotMain) SignAndPutTransfer(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 && state[0] != "Transfer" {
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

	transferDB, err := b.DB.Transfers.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}

	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}

	userPass := msg.GetText()
	//string amount to big int
	amount, ok := new(big.Float).SetString(transferDB.Amount)
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
	transfer := crypto.Transfer{
		ToPubkey: transferDB.ToPubkey,
		Amount:   motesAmountInt.String(),
		Memo:     transferDB.MemoID,
	}
	var dep *types.Deploy
	if u.StorePrivatKey {
		dep, err = b.Crypto.SignTransferWithPassword(transfer, msg.GetFrom().GetId(), userPass)
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
		dep, err = b.Crypto.SignTransferWithPK(transfer, msg.GetFrom().GetId(), userPass)
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

	data := pb.TransferSuccesResponse{
		User:     msg.GetFrom(),
		MsgId:    msgid,
		Amount:   transferDB.Amount,
		Name:     transferDB.Name,
		ToPubkey: transferDB.ToPubkey,
		Memo:     transferDB.MemoID,
		Hash:     res,
	}

	if transferDB.AdditionalType == "invoice" && u.EnableLogging {
		rec, err := u.QueryRecentInvoices().Where(recentinvoices.
			InvoiceID(transferDB.InvoiceID)).
			Only(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed get recent invoice")
		}
		err = rec.Update().SetStatus("paidByTransfer").Exec(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed set status for recent invoice")
		}
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "TransferSuccessResponse",
		Data: dataBytes,
	}

	if !u.EnableLogging {
		err = b.DB.Transfers.DeleteOneID(trID).Exec(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed del transfer")
		}
	}

	err = b.State.DeleteUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed delete user state")
	}

	return nil
}

func (b *BotMain) HandleTransferSetMaximumAmount(msg tggateway.TgMessageMsg) error {
	out := pb.TransferMaximumButton{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 && state[0] != "Transfer" {
		err = b.State.DeleteUserState(out.GetUser().GetId())
		if err != nil {
			return errors.Wrap(err, "failed delete user state")
		}
		return errors.Wrap(err, "bad get user state")
	}
	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	transferDB, err := b.DB.Transfers.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}

	//transferDB.Amount to bigint and minus 10 cspr
	amount, ok := new(big.Float).SetString(transferDB.SenderBalance)
	if !ok {
		return errors.Wrap(err, "failed parse amount")
	}
	amount.Sub(amount, big.NewFloat(10))

	transferDB, err = transferDB.Update().SetAmount(amount.String()).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update transfer")
	}

	state[2] = "askMemo"
	err = b.State.SetUserState(out.GetUser().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed set user state")
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

	data := pb.TransferAskMemo{
		User:   out.GetUser(),
		MsgId:  msgid,
		Amount: transferDB.Amount,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "TransferAskMemo",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) HandleTransferAddressBookButton(msg tggateway.TgMessageMsg) error {
	out := pb.TgAddressButton{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	loggedin, err := b.CheckLogin(out.GetFrom())
	if err != nil {
		return err
	}
	if !loggedin {
		return nil
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}

	addressesDB, err := u.QueryAddressBook().Offset(int(out.GetOffset())).Limit(5).All(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get address book")
	}

	count, err := u.QueryAddressBook().Count(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get address book")
	}

	var addresses []*pb.AddressRow

	for _, address := range addressesDB {
		addresses = append(addresses, &pb.AddressRow{
			Address: address.Address,
			Name:    address.Name,
			Id:      uint64(address.ID),
		})
	}

	data := pb.AddressResponse{
		User:      out.GetFrom(),
		MsgId:     out.GetMsgId(),
		Offset:    out.GetOffset(),
		Total:     int64(count),
		Addresses: addresses,
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "TransferAddressBookMsg",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) PickAddressFromAddressBook(msg tggateway.TgMessageMsg) error {
	out := pb.PickTransferAddress{}

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

	addressesDB, err := u.QueryAddressBook().Where(adressbook.ID(int(out.GetId()))).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get address book")
	}
	if addressesDB == nil {
		return errors.Wrap(err, "failed get address book")
	}

	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	transferDB, err := b.DB.Transfers.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}

	transferDB, err = transferDB.Update().SetToPubkey(addressesDB.Address).SetName(addressesDB.Name).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update transfer")
	}

	uBalance, err := strconv.ParseFloat(transferDB.SenderBalance, 64)
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

	data := pb.TransferAskAmount{
		User:        out.GetUser(),
		MsgId:       out.GetMsgId(),
		Recommended: uBalance - 10,
		ToPubkey:    transferDB.ToPubkey,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "TransferAskAmount",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) HandleTransferWithoutMemo(msg tggateway.TgMessageMsg) error {
	out := pb.TransferWithoutMemo{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 && state[0] != "Transfer" {
		err = b.State.DeleteUserState(out.GetUser().GetId())
		if err != nil {
			return errors.Wrap(err, "failed delete user state")
		}
		return errors.Wrap(err, "bad get user state")
	}
	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	transferDB, err := b.DB.Transfers.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}

	state[2] = "confirmation"
	err = b.State.SetUserState(out.GetUser().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}

	data := pb.TransferAskConfirmation{
		User:     out.GetUser(),
		MsgId:    out.GetMsgId(),
		Amount:   transferDB.Amount,
		ToPubkey: transferDB.ToPubkey,
		Name:     transferDB.Name,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "TransferAskConfirmation",
		Data: dataBytes,
	}

	return nil

}
