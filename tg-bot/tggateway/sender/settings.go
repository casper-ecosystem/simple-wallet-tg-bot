package sender

import (
	"log"
	"strconv"

	"github.com/Simplewallethq/tg-bot/tggateway/messages"
	pb "github.com/Simplewallethq/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/tg-bot/tggateway/types"
	"google.golang.org/protobuf/proto"
	tele "gopkg.in/telebot.v3"
)

func (S *Sender) SendSettingsMsg(msg types.TgResponseMsg) {
	out := pb.SettingsResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		S.logger.Error("error unmarshalling message SettingsResponse: ", err)
		return
	}
	settings := types.SettingsResponse{
		PublicKey:            out.GetPublicKey(),
		NotificationsEnabled: out.GetNotifications(),
		LockTimeout:          out.GetLockTimeout(),
		NotifyTime:           out.GetNotifyTime(),
	}
	what, opts, err := messages.GetSettingsMessage("eng", settings)
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	if out.GetMsgId() == 0 {
		_, err = S.bot.Send(&recipient{Id: out.GetUser().GetId()}, what, opts...)
		if err != nil {
			S.logger.Error("error sending message SettingsResponse: ", err, "uid: ", out.GetUser().GetId())
		}
	} else {
		smsg := tele.StoredMessage{
			MessageID: strconv.Itoa(int(out.GetMsgId())),
			ChatID:    out.GetUser().GetId(),
		}
		_, err = S.bot.Edit(&smsg, what, opts...)
		if err != nil {
			S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
		}
	}

}

func (S *Sender) SendChangeLogoutAskTime(msg types.TgResponseMsg) {
	out := pb.ChangeLogoutAskTime{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		S.logger.Error("error unmarshalling message ChangeLogoutAskTime: ", err)
		return
	}
	what, opts, err := messages.GetChangeLockTimeoutMessage("eng", out.CurrentTime)
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	log.Println(out.GetMsgId())
	smsg := tele.StoredMessage{
		MessageID: strconv.Itoa(int(out.GetMsgId())),
		ChatID:    out.GetUser().GetId(),
	}
	_, err = S.bot.Edit(&smsg, what, opts...)
	if err != nil {
		S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
	}

}

func (S *Sender) SendChangeLogoutAskTimeResponse(msg types.TgResponseMsg) {
	out := pb.ChangeLogoutAskTimeResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		S.logger.Error("error unmarshalling message ChangeLogoutAskTimeResponse: ", err)
		return
	}
	what, opts, err := messages.GetChangeLockTimeoutMessageSuccess("eng", out.CurrentTume)
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

func (S *Sender) SendNotifySettingsMsg(msg types.TgResponseMsg) {
	out := pb.NotifySettingsResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		S.logger.Error("error unmarshalling message NotifySettingsResponse: ", err)
		return
	}
	settings := types.NotifySettingsResponse{
		NotificationsEnabled: out.GetNotifications(),
		NotyfyTime:           out.GetNotifyTime(),
	}
	what, opts, err := messages.GetNotifySettingsMessage("eng", settings)
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	log.Println(out.GetMsgId())
	smsg := tele.StoredMessage{
		MessageID: strconv.Itoa(int(out.GetMsgId())),
		ChatID:    out.GetUser().GetId(),
	}
	_, err = S.bot.Edit(&smsg, what, opts...)
	if err != nil {
		S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
	}

}

func (S *Sender) SendExportAskPassword(msg types.TgResponseMsg) {
	out := pb.ExportAskPassword{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received send transfer ask confirmation message")
	what, opts, err := messages.GetExportAskPasswordMessage("eng")
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

//SendPrivacySettingsResponse

func (S *Sender) SendPrivacySettingsResponse(msg types.TgResponseMsg) {
	out := pb.PrivacyMenu{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		S.logger.Error("error unmarshalling message PrivacySettingsResponse: ", err)
		return
	}
	what, opts, err := messages.GetPrivacySettingsMessage("eng", out.GetLogStatus())
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	log.Println(out.GetMsgId())
	smsg := tele.StoredMessage{
		MessageID: strconv.Itoa(int(out.GetMsgId())),
		ChatID:    out.GetUser().GetId(),
	}
	_, err = S.bot.Edit(&smsg, what, opts...)
	if err != nil {
		S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
	}

}

func (S *Sender) SendExportIncorrectPassword(msg types.TgResponseMsg) {
	out := pb.ExportAskPassword{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received send export incorrect password message")
	what, opts, err := messages.GetExportIncorrectPasswordMessage("eng")
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

func (S *Sender) SendErrorExportPKNotStore(msg types.TgResponseMsg) {
	out := pb.ErrorExportPKNotStore{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received ErrorExportPKNotStore message")
	what, opts, err := messages.GetErrorExportPKNotStore("eng")
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
