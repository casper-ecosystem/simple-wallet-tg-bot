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

func (S *Sender) SendNewTransferResponseStage1(msg types.TgResponseMsg) {
	out := pb.NewTransferResponseStage1{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received address book message")
	what, opts, err := messages.GetSendTransferStage1Message("eng", out.GetBalance())
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

func (S *Sender) SendTransferAskAdress(msg types.TgResponseMsg) {
	out := pb.TransferAskCustomAddress{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received address book message")
	what, opts, err := messages.GetTransferAskAdressMessage("eng")
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

func (S *Sender) SendTransferAskAmount(msg types.TgResponseMsg) {
	out := pb.TransferAskAmount{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received address book message")
	what, opts, err := messages.GetTransferAskAmountMessage("eng", out.GetRecommended(), out.GetToPubkey())
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

func (S *Sender) SendAskMemo(msg types.TgResponseMsg) {
	out := pb.TransferAskMemo{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received address book message")
	what, opts, err := messages.GetTransferAskMemo("eng", out.GetAmount())
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

func (S *Sender) SendTransferAskConfirmation(msg types.TgResponseMsg) {
	out := pb.TransferAskConfirmation{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received address book message")
	what, opts, err := messages.GetTransferAskConfirmation("eng", out.GetAmount(), out.GetToPubkey(), out.GetName())
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

func (S *Sender) SendSignDeployAskPassword(msg types.TgResponseMsg) {
	out := pb.TransferAskConfirmation{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received send transfer ask confirmation message")
	what, opts, err := messages.GetSignDeployAskPasswordMessage("eng")
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

func (S *Sender) SendSignDeployAskPK(msg types.TgResponseMsg) {
	out := pb.SignDeployAskPK{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received send transfer ask confirmation message")
	what, opts, err := messages.GetSignDeployAskPK("eng")
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

func (S *Sender) SendTransferSuccessResponse(msg types.TgResponseMsg) {
	out := pb.TransferSuccesResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received address book message")
	what, opts, err := messages.GetSuccessTransferMessage("eng", out.GetAmount(), out.GetToPubkey(), out.GetName(), out.GetMemo(), out.GetHash())
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

func (S *Sender) SendTransferAddressNotValid(msg types.TgResponseMsg) {
	out := pb.TransferAddressIsNotValid{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received transfer address not valid message")
	what, opts, err := messages.GetTransferAddressIsNotValidMsg("eng")
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

func (S *Sender) SendTransferAmountNotValid(msg types.TgResponseMsg) {
	//log.Println("HERE")
	out := pb.TransferAddressIsNotValid{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received transfer address not valid message")
	what, opts, err := messages.GetTransferAmountIsNotValidMsg("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	log.Println(out.GetMsgId())
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

func (S *Sender) SendTransferMemoNotValid(msg types.TgResponseMsg) {
	out := pb.TransferAddressIsNotValid{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received transfer address not valid message")
	what, opts, err := messages.GetTransferMemoIsNotValidMsg("eng")
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

func (S *Sender) SendTransferAddressBookMsg(msg types.TgResponseMsg) {
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
	what, opts, err := messages.GetTransferAddressBookMsg("eng", data)
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

}

func (S *Sender) SendTransferUnknownError(msg types.TgResponseMsg) {
	out := pb.TransferUnknownError{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received transfer address not valid message")
	what, opts, err := messages.GetTransferUnknownError("eng")
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

func (S *Sender) SendTransferBadPassword(msg types.TgResponseMsg) {
	out := pb.TransferBadPassword{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received transfer address not valid message")
	what, opts, err := messages.GetTransferBadPassword("eng")
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

func (S *Sender) SendTransferBadPK(msg types.TgResponseMsg) {
	out := pb.TransferBadPassword{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received transfer address not valid message")
	what, opts, err := messages.GetTransferBadPK("eng")
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
