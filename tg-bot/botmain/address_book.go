package botmain

import (
	"context"
	"strconv"
	"time"

	"github.com/Simplewallethq/tg-bot/ent/adressbook"
	"github.com/Simplewallethq/tg-bot/ent/user"
	pb "github.com/Simplewallethq/tg-bot/tggateway/proto"

	//"github.com/Simplewallethq/tg-bot/tggateway/types"
	tggateway "github.com/Simplewallethq/tg-bot/tggateway/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (b *BotMain) HandleAddressBook(msg tggateway.TgMessageMsg) error {
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

	// addressesDB, err := u.QueryAddressBook().All(context.Background())
	// if err != nil {
	// 	return errors.Wrap(err, "failed get address book")
	// }
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
		Name: "AddressBookMsg",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) CreateEntryAddressBook(msg tggateway.TgMessageMsg) error {
	out := pb.CreateEntryAddressBookButton{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	data := pb.AskNameAddressBook{
		User:  out.GetUser(),
		MsgId: out.GetMsgId(),
	}
	_, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}

	//_, err = u.Update().SetState("addressBook/askName").Save(context.Background())
	err = b.State.SetUserState(out.GetUser().GetId(), []string{"addressBook", "askName", strconv.Itoa(int(out.GetMsgId()))})
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
		Name: "AskNameAddressBookMsg",
		Data: dataBytes,
	}
	return nil
}

func (b *BotMain) HandleAddressBookState(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())

	if err != nil {
		return errors.Wrap(err, "failed get from")
	}
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get from state")
	}
	ustate, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get from state")
	}
	var editableMsgId int64
	if len(ustate) >= 3 {
		id, err := strconv.Atoi(ustate[2])
		if err != nil {
			return errors.Wrap(err, "failed get id from from state")
		}
		editableMsgId = int64(id)
	} else {
		editableMsgId = -1
	}
	switch state[1] {
	case "askName":
		ab, err := b.DB.AdressBook.Create().SetName(msg.GetText()).SetAddress("").SetCreatedAt(time.Now()).SetInUpdate(true).SetOwner(u).Save(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed create address book entry")
		}
		//_, err = u.Update().SetState("addressBook/askAddress").Save(context.Background())
		// if err != nil {
		// 	return errors.Wrap(err, "failed update from")
		// }
		err = b.State.SetUserState(msg.GetFrom().GetId(), []string{"addressBook", "askAddress", strconv.Itoa(int(editableMsgId))})
		if err != nil {
			return errors.Wrap(err, "failed set from state")
		}
		data := pb.AskAddressAddressBook{
			User:  msg.GetFrom(),
			MsgId: editableMsgId,
			Name:  ab.Name,
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}

		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "AskAddressAddressBookMsg",
			Data: dataBytes,
		}
		if err != nil {
			return errors.Wrap(err, "failed update from")
		}
	case "askAddress":
		from := msg.GetFrom()
		valid, err := b.Restclient.IsAddress(b.RPCNode, msg.GetText())
		if err != nil {
			return errors.Wrap(err, "failed check address")
		}
		if !valid {
			data := pb.AskAddressInvalidResponse{
				User:  from,
				MsgId: editableMsgId,
			}
			dataBytes, err := proto.Marshal(&data)
			if err != nil {
				return errors.Wrap(err, "failed marshal")
			}
			b.ResponseChan <- tggateway.TgResponseMsg{
				Name: "AskAdressInvalidResponse",
				Data: dataBytes,
			}

			return nil
		}
		_, err = b.DB.AdressBook.Update().Where(adressbook.InUpdate(true)).SetAddress(msg.GetText()).SetCreatedAt(time.Now()).SetInUpdate(false).Save(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed update address book entry")
		}
		// _, err = b.DB.User.Update().SetState("").Save(context.Background())
		// if err != nil {
		// 	return errors.Wrap(err, "failed update from")
		// }
		err = b.State.DeleteUserState(msg.GetFrom().GetId())
		if err != nil {
			return errors.Wrap(err, "failed update from")
		}

		msgSett := pb.TgAddressButton{
			From:  from,
			MsgId: editableMsgId,
		}
		outSett, err := proto.Marshal(&msgSett)
		if err != nil {
			return err
		}
		msgtosett := tggateway.TgMessageMsg{
			Name: "AddressBook",
			Data: outSett,
		}
		err = b.HandleAddressBook(msgtosett)
		return err
	case "changeName":
		id, err := strconv.Atoi(state[3])
		from := msg.GetFrom()
		if err != nil {
			return errors.Wrap(err, "failed convert id")
		}
		_, err = b.DB.AdressBook.Update().Where(adressbook.ID(id)).SetName(msg.GetText()).Save(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed update address book entry")
		}
		err = b.State.DeleteUserState(msg.GetFrom().GetId())
		if err != nil {
			return errors.Wrap(err, "failed update from")
		}
		msgSett := pb.TgAddressButton{
			From:  from,
			MsgId: editableMsgId,
		}
		outSett, err := proto.Marshal(&msgSett)
		if err != nil {
			return err
		}
		msgtosett := tggateway.TgMessageMsg{
			Name: "AddressBook",
			Data: outSett,
		}
		err = b.HandleAddressBook(msgtosett)
		return err
	case "changeAddress":
		valid, err := b.Restclient.IsAddress(b.RPCNode, msg.GetText())
		if err != nil {
			return errors.Wrap(err, "failed check address")
		}
		if !valid {
			data := pb.AskAddressInvalidResponse{
				User:  msg.GetFrom(),
				MsgId: editableMsgId,
			}
			dataBytes, err := proto.Marshal(&data)
			if err != nil {
				return errors.Wrap(err, "failed marshal")
			}
			b.ResponseChan <- tggateway.TgResponseMsg{
				Name: "AskAdressInvalidResponse",
				Data: dataBytes,
			}

			return nil
		}
		id, err := strconv.Atoi(state[3])
		from := msg.GetFrom()
		if err != nil {
			return errors.Wrap(err, "failed convert id")
		}
		_, err = b.DB.AdressBook.Update().Where(adressbook.ID(id)).SetAddress(msg.GetText()).Save(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed update address book entry")
		}
		err = b.State.DeleteUserState(msg.GetFrom().GetId())
		if err != nil {
			return errors.Wrap(err, "failed update from")
		}
		msgSett := pb.TgAddressButton{
			From:  from,
			MsgId: editableMsgId,
		}
		outSett, err := proto.Marshal(&msgSett)
		if err != nil {
			return err
		}
		msgtosett := tggateway.TgMessageMsg{
			Name: "AddressBook",
			Data: outSett,
		}
		err = b.HandleAddressBook(msgtosett)
		return err
	}

	return nil
}

func (b *BotMain) AskAddressBookDetailed(msg tggateway.TgMessageMsg) error {
	out := pb.AskAddressBookDetailed{}
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

	addressesDB, err := u.QueryAddressBook().Where(adressbook.ID(int(out.GetId()))).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get address book")
	}
	if addressesDB == nil {
		return errors.Wrap(err, "failed get address book")
	}
	data := pb.AddressBookDetailed{
		User:    out.GetUser(),
		MsgId:   out.GetMsgId(),
		Address: addressesDB.Address,
		Name:    addressesDB.Name,
		Id:      uint64(addressesDB.ID),
		Created: timestamppb.New(addressesDB.CreatedAt),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AddressBookDetailed",
		Data: dataBytes,
	}

	return nil

}
func (b *BotMain) ChangeNameAddressBook(msg tggateway.TgMessageMsg) error {
	out := pb.ChangeNameAddressBook{}
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
	//_, err = u.Update().SetState("addressBook/changeName/" + strconv.Itoa(int(out.GetId()))).Save(context.Background())
	err = b.State.SetUserState(out.GetUser().GetId(), []string{"addressBook", "changeName", strconv.Itoa(int(out.GetMsgId())), strconv.Itoa(int(out.GetId()))})
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}

	data := pb.AskNameAddressBook{
		User:  out.GetUser(),
		MsgId: out.GetMsgId(),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AskNameAddressBookMsg",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) ChangeAddressAddressBook(msg tggateway.TgMessageMsg) error {
	out := pb.ChangeAddressAddressBook{}
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
	//_, err = u.Update().SetState("addressBook/changeAddress/" + strconv.Itoa(int(out.GetId()))).Save(context.Background())
	err = b.State.SetUserState(out.GetUser().GetId(), []string{"addressBook", "changeAddress", strconv.Itoa(int(out.GetMsgId())), strconv.Itoa(int(out.GetId()))})
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	ab, err := b.DB.AdressBook.Query().Where(adressbook.ID(int(out.GetId()))).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get address book name")
	}

	data := pb.AskAddressAddressBook{
		User:  out.GetUser(),
		MsgId: out.GetMsgId(),
		Name:  ab.Name,
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AskAddressAddressBookMsg",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) DeleteEntryAddressBook(msg tggateway.TgMessageMsg) error {
	out := pb.DeleteEntryAddressBook{}
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
	addressesDB, err := u.QueryAddressBook().Where(adressbook.ID(int(out.GetId()))).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get address book")
	}
	if addressesDB == nil {
		return errors.Wrap(err, "failed get address book")
	}
	err = b.State.SetUserState(out.GetUser().GetId(), []string{"DeleteAdressBook", strconv.Itoa(int(out.GetId()))})
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}
	data := pb.DeleteEntryAddressBookConfirmationMessage{
		User:    out.GetUser(),
		MsgId:   out.MsgId,
		Name:    addressesDB.Name,
		Address: addressesDB.Address,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "DeleteAdressBookConfirm",
		Data: dataBytes,
	}
	return nil
}

func (b *BotMain) DeleteEntryAddressBookConfirm(msg tggateway.TgMessageMsg) error {
	out := pb.DeleteEntryAddressBookConfirm{}
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

	state, err := b.State.GetUserState(u.ID)
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	var idToDel int
	if len(state) == 2 && state[0] == "DeleteAdressBook" {
		idToDel, err = strconv.Atoi(state[1])
		if err != nil {
			return errors.Wrap(err, "bad address book field id")
		}
	}
	addressesDB, err := u.QueryAddressBook().Where(adressbook.ID(idToDel)).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get address book")
	}
	if addressesDB == nil {
		return errors.Wrap(err, "failed get address book")
	}
	err = b.DB.AdressBook.DeleteOne(addressesDB).Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed delete address book entry")
	}

	msgSett := pb.TgAddressButton{
		From:  out.GetUser(),
		MsgId: -1,
	}
	outSett, err := proto.Marshal(&msgSett)
	if err != nil {
		return err
	}
	msgtosett := tggateway.TgMessageMsg{
		Name: "AddressBook",
		Data: outSett,
	}
	err = b.HandleAddressBook(msgtosett)
	return err

}

func (b *BotMain) CancelAddressBook(msg tggateway.TgMessageMsg) error {
	out := pb.CancelAddressBook{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	addressesDB, err := u.QueryAddressBook().Where(adressbook.Address("")).All(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get address book")
	}
	if addressesDB == nil {
		return errors.Wrap(err, "failed get address book")
	}
	for _, addr := range addressesDB {
		err = b.DB.AdressBook.DeleteOne(addr).Exec(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed delete address book entry")
		}
	}
	err = b.State.DeleteUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}

	msgSett := pb.TgAddressButton{
		From:  out.GetUser(),
		MsgId: -1,
	}
	outSett, err := proto.Marshal(&msgSett)
	if err != nil {
		return err
	}
	msgtosett := tggateway.TgMessageMsg{
		Name: "AddressBook",
		Data: outSett,
	}
	err = b.HandleAddressBook(msgtosett)
	return err

}
