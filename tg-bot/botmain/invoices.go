package botmain

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/Simplewallethq/tg-bot/ent"
	"github.com/Simplewallethq/tg-bot/ent/invoice"
	"github.com/Simplewallethq/tg-bot/ent/recentinvoices"
	"github.com/Simplewallethq/tg-bot/ent/user"
	pb "github.com/Simplewallethq/tg-bot/tggateway/proto"

	//"github.com/Simplewallethq/tg-bot/tggateway/types"
	tggateway "github.com/Simplewallethq/tg-bot/tggateway/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func (b *BotMain) HandleNewInvoiceState(msg *pb.TgTextMessage) error {
	log.Println("received new invoice state message", msg)
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 || state[0] != "newInvoice" {
		err = b.State.DeleteUserState(msg.GetFrom().GetId())
		if err != nil {
			return errors.Wrap(err, "failed delete user state")
		}
		return errors.Wrap(err, "bad get user state")
	}

	switch state[2] {
	case "askName":
		err := b.PickInvoiceName(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle invoice name in state")
		}
	case "askAmount":
		err := b.PickInvoiceAmount(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle invoice amount")
		}
	case "askRepeatability":
		err := b.PickInvoiceRepeatability(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle invoice Repeatability")
		}
	case "askComment":
		err := b.PickInvoiceComment(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle invoice Repeatability")
		}
	}
	return nil
}

func (b *BotMain) HandleInvoiceButton(msg tggateway.TgMessageMsg) error {
	out := pb.TgInvoicesButton{}
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

	// addressesDB, err := u.QueryAddressBook().All(context.Background())
	// if err != nil {
	// 	return errors.Wrap(err, "failed get address book")
	// }
	invoiceDB, err := u.QueryInvoices().Offset(int(out.GetOffset())).Limit(5).Where(invoice.Active(true)).All(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get invoices")
	}

	count, err := u.QueryInvoices().Where(invoice.Active(true)).Count(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get address book")
	}

	var invoices []*pb.InvoiceRow

	for _, inv := range invoiceDB {
		if !inv.Active {
			continue
		}
		invoices = append(invoices, &pb.InvoiceRow{
			Name:   inv.Name,
			Id:     uint64(inv.ID),
			Amount: inv.Amount,
		})
	}

	data := pb.InvoicesListResponse{
		User:     out.GetFrom(),
		MsgId:    out.GetMsgId(),
		Offset:   out.GetOffset(),
		Total:    int64(count),
		Invoices: invoices,
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "InvoicesListMsg",
		Data: dataBytes,
	}

	return nil

}
func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
func (b *BotMain) invoiceGenerateShort() (string, error) {
	for {
		random, err := generateRandomString(7)
		if err != nil {
			return "", err
		}
		exist, err := b.DB.Invoice.Query().Where(invoice.Short(random)).Exist(context.Background())
		if err != nil {
			return "", err
		}
		if !exist {
			return random, err
		}
	}
}
func (b *BotMain) invoiceGenerateMemo() (uint64, error) {
	for {
		// var by [8]byte
		// if _, err := rand.Read(by[:]); err != nil {
		// 	return 0, err
		// }
		// randmemo := binary.LittleEndian.Uint64(by[:])
		rint, err := rand.Int(rand.Reader, big.NewInt(99999999999))
		if err != nil {
			return 0, err
		}
		randmemo := 100000000000 + rint.Uint64()
		exist, err := b.DB.Invoice.Query().Where(invoice.Memo(randmemo)).Exist(context.Background())
		if err != nil {
			return 0, err
		}
		if !exist {
			return randmemo, err
		}
	}
}

func (b *BotMain) CreateNewInvoice(msg tggateway.TgMessageMsg) error {
	out := pb.TgNewInvoiceButton{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	data := pb.AskInvoiceName{
		User:  out.GetFrom(),
		MsgId: out.GetMsgId(),
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	short, err := b.invoiceGenerateShort()
	if err != nil {
		return errors.Wrap(err, "failed generate short link")
	}
	memo, err := b.invoiceGenerateMemo()
	if err != nil {
		return errors.Wrap(err, "failed generate random memo")
	}
	invDB, err := b.DB.Invoice.Create().
		SetOwner(u).
		SetAddress(u.PublicKey).
		SetActive(true).
		SetShort(short).
		SetMemo(memo).
		Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed create invoice record in database")
	}

	err = b.State.SetUserState(out.GetFrom().GetId(), []string{"newInvoice", strconv.Itoa(invDB.ID), "askName", strconv.Itoa(int(out.GetMsgId()))})
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AskNewInvoiceName",
		Data: dataBytes,
	}
	return nil
}

func (b *BotMain) PickInvoiceName(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	InvDB, err := b.DB.Invoice.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get invoice db")
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

	InvDB, err = InvDB.Update().SetName(msg.GetText()).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update invoice")
	}

	data := pb.AskInvoiceAmount{
		User:  msg.GetFrom(),
		MsgId: msgid,
		Name:  InvDB.Name,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AskInvoiceAmount",
		Data: dataBytes,
	}

	state[2] = "askAmount"
	err = b.State.SetUserState(msg.GetFrom().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed update user state")
	}

	return nil

}

func (b *BotMain) PickInvoiceAmount(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	InvDB, err := b.DB.Invoice.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get invoice db")
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
	bAmount, ok := new(big.Float).SetString(msg.GetText())
	if !ok || bAmount.Sign() != 1 {
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
	InvDB, err = InvDB.Update().SetAmount(msg.GetText()).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update invoice")
	}

	data := pb.AskInvoiceRepeatability{
		User:  msg.GetFrom(),
		MsgId: msgid,
		Name:  InvDB.Name,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AskInvoiceRepeatability",
		Data: dataBytes,
	}

	state[2] = "askRepeatability"
	err = b.State.SetUserState(msg.GetFrom().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed update user state")
	}

	return nil

}

func (b *BotMain) PickInvoiceRepeatability(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	InvDB, err := b.DB.Invoice.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get invoice db")
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
	rep, err := strconv.Atoi(msg.GetText())
	if err != nil || rep < 0 {
		data := pb.InvoiceRepeatabilityIsNotValid{
			User:  msg.GetFrom(),
			MsgId: msgid,
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}

		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "InvoiceRepeatabilityIsNotValid",
			Data: dataBytes,
		}
		return nil
	}
	InvDB, err = InvDB.Update().SetRepeatability(rep).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update invoice")
	}

	data := pb.AskInvoiceComment{
		User:  msg.GetFrom(),
		MsgId: msgid,
		Name:  InvDB.Name,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AskInvoiceComment",
		Data: dataBytes,
	}

	state[2] = "askComment"
	err = b.State.SetUserState(msg.GetFrom().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed update user state")
	}

	return nil

}

func (b *BotMain) PickInvoiceComment(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	InvDB, err := b.DB.Invoice.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get invoice db")
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
	InvDB, err = InvDB.Update().SetComment(msg.GetText()).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update invoice")
	}

	data := pb.InvoiceCreateSuccess{
		User:          msg.GetFrom(),
		MsgId:         msgid,
		Name:          InvDB.Name,
		Id:            int64(InvDB.ID),
		Amount:        InvDB.Amount,
		Comment:       InvDB.Comment,
		Repeatability: int64(InvDB.Repeatability),
		Short:         InvDB.Short,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "InvoiceCreateSuccess",
		Data: dataBytes,
	}

	err = b.State.DeleteUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed update user state")
	}

	return nil

}

func (b *BotMain) AskInvoiceDetailed(msg tggateway.TgMessageMsg) error {
	out := pb.AskInvoiceDetailed{}
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

	InvDB, err := u.QueryInvoices().Where(invoice.ID(int(out.GetId()))).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get invoice entry from db")
	}
	if InvDB == nil {
		return errors.Wrap(err, "failed get invoice entry from db")
	}
	paidCount := InvDB.Paid

	data := pb.InvoiceDetailed{
		User:          out.GetUser(),
		MsgId:         out.GetMsgId(),
		Name:          InvDB.Name,
		Amount:        InvDB.Amount,
		Repeatability: int64(InvDB.Repeatability),
		Comment:       InvDB.Comment,
		Paid:          int64(paidCount),
		Id:            int64(InvDB.ID),
		Short:         InvDB.Short,
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "InvoiceDetailed",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) DeleteInvoice(msg tggateway.TgMessageMsg) error {
	out := pb.DeleteInvoice{}
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
	data := pb.DeleteInvoiceConfirmationMessage{
		User:      out.GetUser(),
		MsgId:     out.GetMsgId(),
		InvoiceID: out.GetId(),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "DeleteInvoiceConfirmationMessage",
		Data: dataBytes,
	}
	return nil
}

func (b *BotMain) DeleteInvoiceConfirm(msg tggateway.TgMessageMsg) error {
	out := pb.DeleteInvoiceConfirm{}
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
	invDB, err := u.QueryInvoices().Where(invoice.ID(int(out.GetId()))).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get invoice entry from db")
	}
	if invDB == nil {
		return errors.Wrap(err, "failed get invoice entry from db")
	}
	//err = b.DB.Invoice.DeleteOne(invDB).Exec(context.Background())
	//if err != nil {
	//	return errors.Wrap(err, "failed delete invoice entry")
	//}
	err = invDB.Update().SetActive(false).Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed set invoice status incactivr")
	}

	msgSett := pb.TgInvoicesButton{
		From:   out.GetUser(),
		MsgId:  out.GetMsgId(),
		Offset: 0,
	}
	outSett, err := proto.Marshal(&msgSett)
	if err != nil {
		return err
	}
	msgtosett := tggateway.TgMessageMsg{
		Name: "Invoices",
		Data: outSett,
	}
	err = b.HandleInvoiceButton(msgtosett)
	return err

}

func (b *BotMain) CheckInvoiceAvailability(short string) (bool, error) {
	invexist, err := b.DB.Invoice.Query().Where(invoice.Short(short)).Exist(context.Background())
	if err != nil {
		return false, errors.Wrap(err, "failed get invoice entry from db")
	}
	if !invexist {
		//todo handle invoice not exist
		log.Println("INVOICE NOT EXIST")
		return false, nil
	}
	InvDB, err := b.DB.Invoice.Query().Where(invoice.Short(short)).Only(context.Background())
	if err != nil {
		return false, errors.Wrap(err, "failed get invoice entry from db")
	}
	if InvDB == nil {
		return false, errors.Wrap(err, "failed get invoice entry from db")
	}
	if InvDB.Active && InvDB.Paid <= InvDB.Repeatability {
		return true, nil
	}
	return false, nil
}

func (b *BotMain) HandlePayInvoice(msg tggateway.TgMessageMsg) error {
	out := pb.PayInvoiceHandler{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}

	userExist, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId()), user.Registered(true)).Exist(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	if userExist {
		//handle registered user way
		aviable, err := b.CheckInvoiceAvailability(out.GetShort())
		if err != nil {
			return errors.Wrap(err, "failed get invoice")
		}
		if !aviable {
			data := pb.PayInvoiceNotAviablePM{User: out.GetUser()}
			dataBytes, err := proto.Marshal(&data)
			if err != nil {
				return errors.Wrap(err, "failed marshal")
			}
			b.ResponseChan <- tggateway.TgResponseMsg{
				Name: "PayInvoiceNotAviablePM",
				Data: dataBytes,
			}
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
		InvDB, err := b.DB.Invoice.Query().Where(invoice.Short(out.GetShort())).Only(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed get invoice entry from db")
		}
		paidCount, err := InvDB.QueryPayments().Count(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed count invoice payments from db")
		}
		amountf64, err := strconv.ParseFloat(InvDB.Amount, 64)
		if err != nil {
			return errors.Wrap(err, "failed convert invoice amount")
		}
		data := pb.PayInvoiceRegisteredResponse{
			User:          out.GetUser(),
			MsgId:         -1,
			Name:          InvDB.Name,
			Amount:        InvDB.Amount,
			BalanceEnough: balance > amountf64,
			Repeatability: int64(InvDB.Repeatability),
			Comment:       InvDB.Comment,
			Paid:          int64(paidCount),
			Id:            int64(InvDB.ID),
			Short:         InvDB.Short,
			Memo:          InvDB.Memo,
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "PayInvoiceRegisteredResponse",
			Data: dataBytes,
		}
		if u.EnableLogging {
			err = b.DB.RecentInvoices.Create().SetInvoiceID(int64(InvDB.ID)).SetOwner(u).SetStatus("opened").OnConflictColumns(recentinvoices.FieldInvoiceID).Ignore().Exec(context.Background())
			if err != nil {
				return errors.Wrap(err, "failed set invoice to recent db")
			}
		}
	} else {
		//ask user to register
		_, err := b.DB.User.Create().SetID(out.GetUser().GetId()).Save(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed create user")
		}
		data := pb.PayInvoiceNotRegisteredPM{User: out.GetUser()}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "PayInvoiceNotRegisteredPM",
			Data: dataBytes,
		}
	}

	return nil

}

func (b *BotMain) PayInvoiceTransfer(msg tggateway.TgMessageMsg) error {
	out := pb.PayInvoiceTransfer{}
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

	invDB, err := u.QueryInvoices().Where(invoice.Short(out.GetShort())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get invoice entry from db")
	}
	if invDB == nil {
		return errors.Wrap(err, "failed get invoice entry from db")
	}
	balance, err := b.Restclient.GetBalance(b.RPCNode, u.PublicKey)
	if err != nil {
		return errors.Wrap(err, "failed get balance")
	}

	balanceRes := strconv.FormatFloat(balance, 'f', 6, 64)
	transferDB, err := b.DB.Transfers.Create().
		SetFromPubkey(u.PublicKey).
		SetSenderBalance(balanceRes).
		SetAmount(invDB.Amount).
		SetMemoID(invDB.Memo).
		SetToPubkey(invDB.Address).
		SetCreatedAt(time.Now()).
		SetAdditionalType("invoice").
		SetInvoiceID(int64(invDB.ID)).
		SetOwner(u).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed create transfer")
	}
	err = b.State.SetUserState(out.GetUser().GetId(), []string{"Transfer", strconv.Itoa(transferDB.ID), "confirmation", strconv.Itoa(int(out.MsgId))})
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

func (b *BotMain) PayInvoiceSwap(msg tggateway.TgMessageMsg) error {
	out := pb.PayInvoiceSwap{}
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

	invDB, err := u.QueryInvoices().Where(invoice.Short(out.GetShort())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get invoice entry from db")
	}
	if invDB == nil {
		return errors.Wrap(err, "failed get invoice entry from db")
	}
	swapsDB, err := b.DB.Swaps.Create().SetToAddress(invDB.Address).
		SetToCurrency("cspr").
		SetType("invoice").
		SetAmount(invDB.Amount).
		SetOwner(u).
		SetExtraID(strconv.Itoa(int(invDB.Memo))).
		SetInvoiceID(int64(invDB.ID)).
		Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed create swap")
	}
	err = b.State.SetUserState(out.GetUser().GetId(), []string{"Swap", strconv.Itoa(swapsDB.ID), "pickCurrency"})
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}

	data := pb.AskSwapPairs{
		User:   out.GetUser(),
		MsgId:  out.MsgId,
		Limit:  5,
		Offset: 0,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.MessagesChan <- tggateway.TgMessageMsg{
		Name: "AskSwapPairs",
		Data: dataBytes,
	}

	return nil
}

func (b *BotMain) ShowInvoicePayments(msg tggateway.TgMessageMsg) error {
	out := pb.ShowInvoicePayments{}
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

	InvDB, err := u.QueryInvoices().Where(invoice.ID(int(out.GetId()))).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get invoice entry from db")
	}
	if InvDB == nil {
		return errors.Wrap(err, "failed get invoice entry from db")
	}

	count, err := InvDB.QueryPayments().Count(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed count invoice payments from db")
	}
	payDB, err := InvDB.QueryPayments().Limit(10).Offset(int(out.GetOffset())).All(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get invoice payments from db")
	}
	for _, item := range payDB {
		log.Println(item.From, item.Amount, item.Correct)
	}

	var payments []*pb.PaymentRow

	for _, pay := range payDB {
		payments = append(payments, &pb.PaymentRow{
			From:    pay.From,
			Success: pay.Correct,
			Amount:  pay.Amount,
		})
	}

	data := pb.PaymentsListResponse{
		User:     out.GetUser(),
		MsgId:    out.GetMsgId(),
		Offset:   out.GetOffset(),
		Total:    int64(count),
		Payments: payments,
		Id:       int64(out.GetId()),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "PaymentsListResponse",
		Data: dataBytes,
	}
	return nil

}

func (b *BotMain) ShowRecentInvoices(msg tggateway.TgMessageMsg) error {
	out := pb.ShowRecentInvoices{}
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

	recentDB, err := u.QueryRecentInvoices().Order(ent.Asc(recentinvoices.FieldID)).
		Limit(5).Offset(int(out.GetOffset())).All(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get recent invoice entry from db")
	}
	if recentDB == nil {
		return errors.Wrap(err, "failed get invoice entry from db")
	}

	count, err := u.QueryRecentInvoices().Limit(20).Count(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed count invoice payments from db")
	}

	var invoices []*pb.RecentInvoiceRow

	for _, pay := range recentDB {
		invExist, err := b.DB.Invoice.Query().Where(invoice.ID(int(pay.InvoiceID))).Exist(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed get invoiceInst")
		}
		if !invExist {
			continue
		}
		invoiceInst, err := b.DB.Invoice.Query().Where(invoice.ID(int(pay.InvoiceID))).Only(context.Background())
		if err != nil {

			log.Println("bad invoice id: ", int(pay.InvoiceID))
			return errors.Wrap(err, "failed get invoiceInst")
		}
		invoices = append(invoices, &pb.RecentInvoiceRow{
			Name:   invoiceInst.Name,
			Status: pay.Status,
			Short:  invoiceInst.Short,
		})
	}

	data := pb.RecentInvoicesListResponse{
		User:     out.GetUser(),
		MsgId:    out.GetMsgId(),
		Offset:   out.GetOffset(),
		Total:    int64(count),
		Invoices: invoices,
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "RecentInvoiceResponse",
		Data: dataBytes,
	}
	return nil

}

func (b *BotMain) ExportPaymentsInvoice(msg tggateway.TgMessageMsg) error {
	out := pb.ExportPaymentsInvoice{}
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

	InvDB, err := u.QueryInvoices().Where(invoice.ID(int(out.GetId()))).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get invoice entry from db")
	}
	if InvDB == nil {
		return errors.Wrap(err, "failed get invoice entry from db")
	}

	payDB, err := InvDB.QueryPayments().All(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get invoice payments from db")
	}

	var bb bytes.Buffer
	w := csv.NewWriter(bufio.NewWriter(&bb))

	for _, item := range payDB {
		tempString := []string{item.From, item.Amount, fmt.Sprintf("%t", item.Correct)}
		if err := w.Write(tempString); err != nil {
			return errors.Wrap(err, "failed make csv")
		}
	}
	w.Flush()
	log.Println(bb.String())

	data := pb.ExportPaymentsInvoiceResponse{
		User:  out.GetUser(),
		MsgId: out.GetMsgId(),
		Data:  bb.Bytes(),
		Short: InvDB.Short,
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "ExportPaymentsInvoiceResponse",
		Data: dataBytes,
	}
	return nil

}
