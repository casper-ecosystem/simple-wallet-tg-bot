package sender

import (
	"log"
	"strconv"
	"strings"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/messages"
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	"google.golang.org/protobuf/proto"
	tele "gopkg.in/telebot.v3"
)

func (S *Sender) SendYieldMsg(msg types.TgResponseMsg) {
	out := pb.YieldResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	log.Println("received yield message", out.GetRewards())
	yield := types.YieldResponse{
		TotalRewards:   out.GetRewards(),
		NetworkApy:     out.GetNetworkApy(),
		RewardsUSD:     out.GetRewardsUSD(),
		Estimates:      out.GetEstimates(),
		Validators:     out.GetValidatorsData(),
		TotalDelegated: out.GetTotalDelegated(),
		Proj365Days:    out.GetEstim365Days(),
		Proj30Days:     out.GetEstim30Days(),
		Proj365DaysUSD: out.GetEstim365DaysUSD(),
		Proj30DaysUSD:  out.GetEstim30DaysUSD(),
		ProjApy:        out.GetEstimApy(),
	}

	what, opts, err := messages.GetYieldMsg("eng", yield)
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	if len(what.(string)) >= 1024 {
		//messagesData := []string{"test", "this", "algho"}
		messagesData := splitMessage(what.(string), 1024)
		smsg := tele.StoredMessage{
			MessageID: strconv.Itoa(int(out.GetMsgId())),
			ChatID:    out.GetUser().GetId(),
		}

		log.Println(smsg.ChatID, smsg.MessageID)
		log.Println(out.GetUser())
		//log.Println(messagesData[0], opts)
		_, err = S.bot.Edit(&smsg, messagesData[0], opts[:1]...)
		if err != nil {
			S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
		}
		for i := 1; i < len(messagesData); i++ {
			var newopts []interface{}
			if i != len(messagesData)-1 {
				newopts = opts[:1]
			} else {
				newopts = opts
			}
			//log.Println(messagesData[i], opts)
			_, err = S.bot.Send(&recipient{Id: out.GetUser().GetId()}, messagesData[i], newopts...)
			if err != nil {
				S.logger.Error("error sending message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
			}
		}

	} else {
		smsg := tele.StoredMessage{
			MessageID: strconv.Itoa(int(out.GetMsgId())),
			ChatID:    out.GetUser().GetId(),
		}

		log.Println(smsg.ChatID, smsg.MessageID)
		log.Println(out.GetUser())
		_, err = S.bot.Edit(&smsg, what, opts...)
		if err != nil {
			S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
		}
	}
}

func splitMessage(text string, max int) []string {
	var result []string
	var buf string
	splitted := strings.Split(text, "\n")
	for _, str := range splitted {
		if len(buf)+len(str) > max {
			result = append(result, buf)
			buf = ""
		}
		buf += str + "\n"

	}
	if len(buf) != 0 {
		result = append(result, buf)
	}
	return result

}
