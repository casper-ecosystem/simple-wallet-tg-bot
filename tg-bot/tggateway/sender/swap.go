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

func (S *Sender) SendSwapPairs(msg types.TgResponseMsg) {
	out := pb.ShowSwapPairs{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received address book message")
	log.Println(out.GetPairs())
	what, opts, err := messages.GetSwapPairs("eng", out.GetOffset(), out.GetTotal(), out.GetPairs(), out.GetType())
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

func (S *Sender) SendSwapAskAmount(msg types.TgResponseMsg) {
	out := pb.SwapAskAmount{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received address book message")
	what, opts, err := messages.GetSwapAskAmount("eng", out.GetFromCurr(), out.GetToCurr(), out.GetMin(), out.GetMax())
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

func (S *Sender) SendEstimatedAmount(msg types.TgResponseMsg) {
	out := pb.SwapAmountEstimated{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received address book message")
	what, opts, err := messages.GetSwapShowEstimated("eng", out.GetEstimated(), out.GetCurr(), out.GetAmount(), out.GetCurrFrom())
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

func (S *Sender) SendSwapAskRefundAddress(msg types.TgResponseMsg) {
	out := pb.SwapAskRefundAddress{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received swap ask refund address")
	what, opts, err := messages.GetSwapAskRefund("eng", out.GetCurr(), out.GetChain())
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

func (S *Sender) SendSwapAskToAddress(msg types.TgResponseMsg) {
	out := pb.SwapAskToAddress{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received swap ask refund address")
	what, opts, err := messages.GetSwapAskAdress("eng", out.GetCurr())
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

func (S *Sender) SendSwapSuccessResponse(msg types.TgResponseMsg) {
	out := pb.SwapSuccessResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received swap ask refund address")
	what, opts, err := messages.GetSwapSuccess("eng", out.GetId(), out.GetDepAddress(), out.GetFromCur(), out.GetToCur(), out.GetAmount(), out.GetEstimated())
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

func (S *Sender) SendSwapShowChains(msg types.TgResponseMsg) {
	out := pb.ShowSwapChains{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received address book message")
	log.Println(out.GetChains())
	what, opts, err := messages.GetSwapChains("eng", &out)
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

func (S *Sender) SendSwapLoadMsg(msg types.TgResponseMsg) {
	out := pb.YieldLoadingResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetSwapLoadMsg("eng")
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
