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

func (S *Sender) LoadHistoryMsg(msg types.TgResponseMsg) {
	out := pb.LoadHistoryResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetLoadHistoryMessage("eng")
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

func (S *Sender) SendHistory(msg types.TgResponseMsg) {
	out := pb.TransferHistoryResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received history message", out.GetTransfers())
	history := types.HistoryResponse{
		Start:             out.GetStart(),
		StartDate:         out.StartDate,
		End:               out.GetEnd(),
		EndDate:           out.EndDate,
		Transfers:         out.GetTransfers(),
		DelegateHistory:   out.GetDelegates(),
		UndelegateHistory: out.GetUndelegates(),
		IsFirst:           out.GetIsFirst(),
	}
	what, opts, err := messages.GetHistoryMsg("eng", history)
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	log.Println(out.GetUser().GetId())
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

func (S *Sender) SendRewardsHistory(msg types.TgResponseMsg) {
	out := pb.RewardsHistoryResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received history message", out.GetRewards())
	history := types.RewardsHistoryResponse{
		Start:     out.GetStart(),
		End:       out.GetEnd(),
		Rewards:   out.GetRewards(),
		IsFirst:   out.GetIsFirst(),
		StartDate: out.GetStartDate(),
		EndDate:   out.GetEndDate(),
	}
	what, opts, err := messages.GetRewardsHistoryMsg("eng", history)
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	log.Println(out.GetUser().GetId())
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
