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

func (S *Sender) SendBalanceMsg(msg types.TgResponseMsg) {
	out := pb.BalanceResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received balance message", out.GetBalance())
	balance := types.BalanceResponse{
		Liquid:           out.GetBalance(),
		Price:            out.GetPrice(),
		TotalUSD:         out.GetTotalUSD(),
		Total:            out.GetTotal(),
		Delegated:        out.GetDelegatedBalance(),
		BeingDelegated:   out.GetBeingDelegatedBalance(),
		BeingUndelegated: out.GetBeingUndelegatedBalance(),
	}
	what, opts, err := messages.GetBalanceMsg("eng", balance)
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
	//S.bot.Send(&recipient{Id: out.GetUser().GetId()}, what, opts...)

}

func (S *Sender) LoadBalanceMsg(msg types.TgResponseMsg) {
	out := pb.LoadBalanceResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetLoadBalanceMessage("eng")
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
