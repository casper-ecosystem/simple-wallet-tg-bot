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

func (S *Sender) SendInvoicesMsg(msg types.TgResponseMsg) {
	out := pb.InvoicesListResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received invoices list msg")
	what, opts, err := messages.GetInvoicesListMsg("eng", &out)
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

func (S *Sender) AskNewInvoiceName(msg types.TgResponseMsg) {
	out := pb.AskInvoiceName{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.AskInvoiceName("eng")
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

func (S *Sender) AskInvoiceAmount(msg types.TgResponseMsg) {
	out := pb.AskInvoiceAmount{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.AskInvoiceAmount("eng")
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

func (S *Sender) AskInvoiceRepeatability(msg types.TgResponseMsg) {
	out := pb.AskInvoiceRepeatability{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.AskInvoiceRepeatability("eng")
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

func (S *Sender) AskInvoiceComment(msg types.TgResponseMsg) {
	out := pb.AskInvoiceComment{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.AskInvoiceComment("eng")
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

func (S *Sender) InvoiceCreateSuccess(msg types.TgResponseMsg) {
	out := pb.InvoiceCreateSuccess{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	botname := S.bot.Me.Username
	what, opts, err := messages.InvoiceCreateSuccess("eng", &out, botname)
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

func (S *Sender) InvoiceDetailed(msg types.TgResponseMsg) {
	out := pb.InvoiceDetailed{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	botname := S.bot.Me.Username
	what, opts, err := messages.InvoiceDetailed("eng", &out, botname)
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

func (S *Sender) PayInvoiceNotRegisteredPM(msg types.TgResponseMsg) {
	out := pb.PayInvoiceNotRegisteredPM{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.InvoiceAskRegisterPM("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}

	recipient := &recipient{Id: out.User.GetId()}
	_, err = S.bot.Send(recipient, what, opts...)
	if err != nil {
		return
	}

}

func (S *Sender) PayInvoiceNotAviablePM(msg types.TgResponseMsg) {
	out := pb.PayInvoiceNotAviablePM{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.PayInvoiceNotAviablePM("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}

	recipient := &recipient{Id: out.User.GetId()}
	_, err = S.bot.Send(recipient, what, opts...)
	if err != nil {
		return
	}
}

func (S *Sender) PayInvoiceRegisteredResponse(msg types.TgResponseMsg) {
	out := pb.PayInvoiceRegisteredResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.PayInvoiceRegisteredResponse("eng", &out)
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}

	recipient := &recipient{Id: out.User.GetId()}
	_, err = S.bot.Send(recipient, what, opts...)
	if err != nil {
		S.logger.Error("error send message: ", err, "uid: ", out.GetUser().GetId())

	}
}

func (S *Sender) PaymentsListMsg(msg types.TgResponseMsg) {
	out := pb.PaymentsListResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received payments list msg")
	what, opts, err := messages.GetPaymentsListMsg("eng", &out)
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

func (S *Sender) RecentInvoiceResponse(msg types.TgResponseMsg) {
	out := pb.RecentInvoicesListResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received payments list msg")
	what, opts, err := messages.GetRecentlyInvoices("eng", &out)
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

func (S *Sender) ExportPaymentsInvoiceResponse(msg types.TgResponseMsg) {
	out := pb.ExportPaymentsInvoiceResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received payments list msg")
	what, opts, err := messages.GetExportPaymentsInvoice("eng", out.GetData(), out.GetShort())
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}

	// if out.GetMsgId() != -1 {
	// 	smsg := tele.StoredMessage{
	// 		MessageID: strconv.Itoa(int(out.GetMsgId())),
	// 		ChatID:    out.GetUser().GetId(),
	// 	}
	// 	_, err = S.bot.Edit(&smsg, what, opts...)
	// 	if err != nil {
	// 		S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
	// 	}
	// } else {
	recipient := &recipient{Id: out.User.GetId()}
	_, err = S.bot.Send(recipient, what, opts...)
	if err != nil {
		return
	}
	// }
	//S.bot.Send(&recipient{Id: out.GetUser().GetId()}, what, opts...)

}

func (S *Sender) DeleteInvoiceConfirmation(msg types.TgResponseMsg) {
	out := pb.DeleteInvoiceConfirmationMessage{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received payments list msg")
	what, opts, err := messages.DeleteInvoiceConfirmation("eng", out.GetInvoiceID())
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

func (S *Sender) SendInvoiceRepeatabilityIsNotValid(msg types.TgResponseMsg) {
	//log.Println("HERE")
	out := pb.InvoiceRepeatabilityIsNotValid{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received invoice repeatebility not valid message")
	what, opts, err := messages.GetInvoiceRepeatabilityIsNotValid("eng")
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
