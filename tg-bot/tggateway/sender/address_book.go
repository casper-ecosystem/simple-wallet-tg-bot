package sender

import (
	"log"
	"strconv"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/messages"
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	"google.golang.org/protobuf/proto"
	tele "gopkg.in/telebot.v3"
)

func (S *Sender) SendAddressBookMsg(msg types.TgResponseMsg) {
	out := pb.AddressResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received address book message")
	data := types.AddressBookResponse{
		Offset: out.GetOffset(),
		Total:  out.GetTotal(),
		Data:   out.GetAddresses(),
	}
	what, opts, err := messages.GetAddressBookMsg("eng", data)
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}

	if out.GetMsgId() != -1 {
		smsg := tele.StoredMessage{
			MessageID: strconv.Itoa(int(out.GetMsgId())),
			ChatID:    out.GetUser().GetId(),
		}
		_, err = S.bot.Edit(&smsg, what, opts...)
		if err != nil {
			S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
		}
	} else {
		recipient := &recipient{Id: out.User.GetId()}
		_, err = S.bot.Send(recipient, what, opts...)
		if err != nil {
			return
		}
	}
	//S.bot.Send(&recipient{Id: out.GetUser().GetId()}, what, opts...)

}

func (S *Sender) AskNameAddressBookMsg(msg types.TgResponseMsg) {
	out := pb.AskNameAddressBook{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetCreateEntryAddressBookNameMsg("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	smsg := tele.StoredMessage{
		MessageID: strconv.Itoa(int(out.GetMsgId())),
		ChatID:    out.GetUser().GetId(),
	}
	_, err = S.bot.Edit(&smsg, what, opts...)
	if err != nil {
		S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
	}
}

func (S *Sender) AskAddressAddressBook(msg types.TgResponseMsg) {
	out := pb.AskAddressAddressBook{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetAskAddresAdressBookMsg("eng", out.GetName())
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	smsg := tele.StoredMessage{
		MessageID: strconv.Itoa(int(out.GetMsgId())),
		ChatID:    out.GetUser().GetId(),
	}
	_, err = S.bot.Edit(&smsg, what, opts...)
	if err != nil {
		return
	}
}
func (S *Sender) AskAddressInvalidAdress(msg types.TgResponseMsg) {
	out := pb.AskAddressInvalidResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetAskAddresInvalidAdress("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	smsg := tele.StoredMessage{
		MessageID: strconv.Itoa(int(out.GetMsgId())),
		ChatID:    out.GetUser().GetId(),
	}
	_, err = S.bot.Edit(&smsg, what, opts...)
	if err != nil {
		return
	}
}

func (S *Sender) SendAddressBookDetailed(msg types.TgResponseMsg) {
	out := pb.AddressBookDetailed{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	data := types.AddressBookDetailed{
		Name:      out.GetName(),
		Address:   out.GetAddress(),
		CreatedAt: out.Created.AsTime(),
		Id:        out.GetId(),
	}
	what, opts, err := messages.GetAddressDetailedMsg("eng", data)
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	smsg := tele.StoredMessage{
		MessageID: strconv.Itoa(int(out.GetMsgId())),
		ChatID:    out.GetUser().GetId(),
	}
	_, err = S.bot.Edit(&smsg, what, opts...)
	if err != nil {
		S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
	}
}

func (S *Sender) DeleteAdressBookConfirm(msg types.TgResponseMsg) {
	out := pb.DeleteEntryAddressBookConfirmationMessage{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetDeleteEntryAddressBookConfirmationMessage("eng", out.GetName(), out.GetAddress())
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	smsg := tele.StoredMessage{
		MessageID: strconv.Itoa(int(out.GetMsgId())),
		ChatID:    out.GetUser().GetId(),
	}
	_, err = S.bot.Edit(&smsg, what, opts...)
	if err != nil {
		S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
	}
}
