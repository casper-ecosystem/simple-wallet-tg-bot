package sender

import (
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/messages"
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	"google.golang.org/protobuf/proto"
)

func (S *Sender) SendNotifyNewTransfer(msg types.TgResponseMsg) {
	out := pb.NotificationNewTransfer{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		S.logger.Error("error unmarshalling message NotificationNewTransfer: ", err)
		return
	}
	what, opts, err := messages.GetNotifyNewTransferMessage("eng", out.GetAmount(), out.GetFrom(), out.GetTo(), out.GetBalance())
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	_, err = S.bot.Send(&recipient{Id: out.GetUser().GetId()}, what, opts...)
	if err != nil {
		S.logger.Error("error sending message NotificationNewTransfer: ", err, "uid: ", out.GetUser().GetId())
	}
}

func (S *Sender) SendNotifyNewDelegate(msg types.TgResponseMsg) {
	out := pb.NotificationNewDelegate{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		S.logger.Error("error unmarshalling message NotificationNewDelegate: ", err)
		return
	}
	what, opts, err := messages.GetNotifyNewDelegateMessage("eng", out.GetAmount(), out.GetValidator(), out.GetHeight(), out.GetBalance())
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	_, err = S.bot.Send(&recipient{Id: out.GetUser().GetId()}, what, opts...)
	if err != nil {
		S.logger.Error("error sending message NotificationNewDelegate: ", err, "uid: ", out.GetUser().GetId())
	}
}

func (S *Sender) SendNotifyNewUndelegate(msg types.TgResponseMsg) {
	out := pb.NotificationNewUndelegate{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		S.logger.Error("error unmarshalling message NotificationNewUndelegate: ", err)
		return
	}
	what, opts, err := messages.GetNotifyNewUndelegateMessage("eng", out.GetAmount(), out.GetValidator(), out.GetEra(), out.GetBalance())
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	_, err = S.bot.Send(&recipient{Id: out.GetUser().GetId()}, what, opts...)
	if err != nil {
		S.logger.Error("error sending message NotificationNewUndelegate: ", err, "uid: ", out.GetUser().GetId())
	}
}

func (S *Sender) SendNotifyNewRewards(msg types.TgResponseMsg) {
	//log.Println("SendNotifyNewRewards")
	out := pb.NotificationNewReward{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		S.logger.Error("error unmarshalling message NotificationNewRewards: ", err)
		return
	}
	var data types.NotifyNewRewards
	for _, val := range out.GetRewards() {
		data.Rewards = append(data.Rewards, types.Reward{
			Validator:      val.GetValidator(),
			Amount:         val.GetAmount(),
			FirstEra:       val.GetFirstEra(),
			LastEra:        val.GetLastEra(),
			LastRewardTime: val.GetLastRewardTime(),
		})
	}
	data.Delegated = out.GetDelegated()
	data.FirstEra = out.GetFirstEra()
	data.LastEra = out.GetLastEra()
	data.FirstEraTimestamp = out.GetFirstEraTimestamp()
	data.LastEraTimestamp = out.GetLastEraTimestamp()

	what, opts, err := messages.GetNotifyNewRewards("eng", data)
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	_, err = S.bot.Send(&recipient{Id: out.GetUser().GetId()}, what, opts...)
	if err != nil {
		S.logger.Error("error sending message NotificationNewRewards: ", err, "uid: ", out.GetUser().GetId())
	}
}

func (S *Sender) SendNotifyNewBalance(msg types.TgResponseMsg) {
	out := pb.NotificationNewBalance{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		S.logger.Error("error unmarshalling message NotificationNewBalance: ", err)
		return
	}
	what, opts, err := messages.GetNotifyNewBalance("eng", out.GetBalance(), out.GetOldBalance())
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	_, err = S.bot.Send(&recipient{Id: out.GetUser().GetId()}, what, opts...)
	if err != nil {
		S.logger.Error("error sending message NotificationNewBalance: ", err, "uid: ", out.GetUser().GetId())
	}
}
