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

func (S *Sender) SendAuthMsg(msg types.TgResponseMsg) {
	out := pb.AuthResponse{}
	var RangeToDel int
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	var what interface{}
	var opts []interface{}
	var err error
	log.Println("stage in sender", out.Stage)
	switch out.Stage {
	case "AskStoreKey":
		what, opts, err = messages.GetAskStoreTheKeyMessage("eng")
		if err != nil {
			what, opts, _ = messages.GetErrorMessage()
		}
		RangeToDel = 1
	case "AskPrivatKey":
		if out.Error != "" {
			if out.Error == "Invalid" {
				what, opts, err = messages.GetAuthInvalidPrivateMessage("eng")
				if err != nil {
					what, opts, _ = messages.GetErrorMessage()
				}
			}
		} else {
			what, opts, err = messages.GetAskPrivatKeyMessage("eng")
			if err != nil {
				what, opts, _ = messages.GetErrorMessage()
			}
		}
		RangeToDel = 2
	case "1":
		if out.Error != "" {
			if out.Error == "Invalid" {
				what, opts, err = messages.GetAuthInvalidPubKeyMessage("eng")
				if err != nil {
					what, opts, _ = messages.GetErrorMessage()
				}
			}
		} else {
			what, opts, err = messages.GetWelcomeAuthMessage("eng")
			if err != nil {
				what, opts, _ = messages.GetErrorMessage()
			}
		}
		RangeToDel = 1
	case "2":
		what, opts, err = messages.GetAuthPasswordMessage("eng")
		if err != nil {
			what, opts, _ = messages.GetErrorMessage()
		}
		RangeToDel = 1
	case "3":
		//log.Println(out.Error)
		if out.Error != "" {
			if out.Error == "Invalid" {
				what, opts, err = messages.GetAuthRepeatPasswordInvalidMessage("eng")
				if err != nil {
					what, opts, _ = messages.GetErrorMessage()
				}
			}
		} else {
			what, opts, err = messages.GetAuthRepeatPasswordMessage("eng")
			if err != nil {
				what, opts, _ = messages.GetErrorMessage()
			}

		}
		RangeToDel = 2
	case "4":
		what, opts, err = messages.GetRegisterSuccessMessage("eng", out.Pubkey)
		if err != nil {
			what, opts, _ = messages.GetErrorMessage()
		}
		recipient := &recipient{Id: out.User.GetId()}
		//fmt.Println((opts))
		_, err = S.bot.Send(recipient, what, opts...)
		if err != nil {
			return
		}
		what, opts, err = messages.GetWelcomeMsg("eng")
		if err != nil {
			what, opts, _ = messages.GetErrorMessage()
		}
		RangeToDel = 3
	case "SendPrivatKey":
		what, opts, err = messages.GetSendPrivatKey("eng", out.GetData(), out.GetPubkey())
		if err != nil {
			what, opts, _ = messages.GetErrorMessage()
		}
		recipient := &recipient{Id: out.User.GetId()}
		//fmt.Println((opts))
		pkMsg, err := S.bot.Send(recipient, what, opts...)
		if err != nil {
			return
		}
		log.Println("Send pk, msg ID:", pkMsg.ID)
		what, opts, err = messages.GetWelcomeMsg("eng")
		if err != nil {
			what, opts, _ = messages.GetErrorMessage()
		}
		RangeToDel = 3
	}

	recipient := &recipient{Id: out.User.GetId()}
	//fmt.Println((opts))
	res, err := S.bot.Send(recipient, what, opts...)
	if err != nil {
		return
	}
	//clean temp messages
	// _ = res
	// _ = RangeToDel
	msgid := strconv.Itoa(res.ID - RangeToDel)
	toDelete := tele.StoredMessage{
		MessageID: msgid,
		ChatID:    res.Chat.ID,
	}
	err = S.bot.Delete(toDelete)
	if err != nil {
		log.Println(err)
		return
	}

}

func (S *Sender) SendAuthTypeMessage(msg types.TgResponseMsg) {
	out := pb.AuthRegisterType{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetChangeAuthTypeMessage("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	recipient := &recipient{Id: out.User.GetId()}
	_, err = S.bot.Send(recipient, what, opts...)
	if err != nil {
		return
	}

}

func (S *Sender) SendLoginMsg(msg types.TgResponseMsg) {
	out := pb.AskLoginResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetLoginMessage("eng", out.ManualLogout)
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	if out.GetMsgId() != 0 {
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

func (S *Sender) SendLoginPassInvalidMsg(msg types.TgResponseMsg) {
	out := pb.LoginPassInvalidResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetLoginPassInvalidMessage("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	recipient := &recipient{Id: out.User.GetId()}
	_, err = S.bot.Send(recipient, what, opts...)
	if err != nil {
		S.logger.Error(err)
		return
	}
}

func (S *Sender) SendLoginSuccessMsg(msg types.TgResponseMsg) {
	out := pb.LoginSuccessResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	recipient := &recipient{Id: out.User.GetId()}
	what, opts, err := messages.GetLoginSuccessMessage("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	_, err = S.bot.Send(recipient, what, opts...)
	if err != nil {
		S.logger.Error(err)
		return
	}
	what, opts, err = messages.GetWelcomeMsg("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	_, err = S.bot.Send(recipient, what, opts...)
	if err != nil {
		S.logger.Error(err)
		return
	}
}

func (S *Sender) SendLockMsg(msg types.TgResponseMsg) {
	out := pb.LockResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetLockMessage("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	recipient := &recipient{Id: out.User.GetId()}
	_, err = S.bot.Send(recipient, what, opts...)
	if err != nil {
		return
	}
}

func (S *Sender) SendLogoutConfirmationMsg(msg types.TgResponseMsg) {
	out := pb.LogoutConfirmation{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetLogoutConfirmationMessage("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	if out.GetMsgId() != -0 {
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

func (S *Sender) SendLogoutMsg(msg types.TgResponseMsg) {
	out := pb.LogoutResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetLogoutMessage("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	recipient := &recipient{Id: out.User.GetId()}
	res, err := S.bot.Send(recipient, what, opts...)
	if err != nil {
		return
	}
	msgid := strconv.Itoa(res.ID - 2)
	toDelete := tele.StoredMessage{
		MessageID: msgid,
		ChatID:    res.Chat.ID,
	}
	err = S.bot.Delete(toDelete)
	if err != nil {
		log.Println(err)
		return
	}
}
