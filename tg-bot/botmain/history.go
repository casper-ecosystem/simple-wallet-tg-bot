package botmain

import (
	"context"
	"strconv"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	tggateway "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

const defaultEraRange = 11
const defaultBlockRange = 500

func (b *BotMain) HandleBalanceHistory(msg tggateway.TgMessageMsg) error {
	out := pb.TgBalanceButtonHistory{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	loggedin, err := b.CheckLogin(out.GetUser())
	if err != nil {
		return err
	}
	if !loggedin {
		return nil
	}
	err = b.SendLoadingHistoryBalance(&out)
	if err != nil {
		return err
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	pubkey := u.PublicKey
	state, err := b.Restclient.GetState(b.RPCNode)
	if err != nil {
		return errors.Wrap(err, "failed get state")
	}
	var Start, End int64
	var isFirst = false
	if out.GetStart() == -1 || out.GetEnd() > state.BlockHeight {
		Start = state.BlockHeight - defaultBlockRange
		End = state.BlockHeight
		isFirst = true
	} else {

		Start = out.GetStart()
		End = out.GetEnd()
	}
	historyTransfers, err := b.Restclient.GetHistoryTransfers(b.RPCNode, pubkey, Start, End)
	if err != nil {
		return errors.Wrap(err, "failed get history transfers")
	}
	resp := pb.TransferHistoryResponse{}
	resp.IsFirst = isFirst
	for _, transfer := range historyTransfers.Data {
		var from, to string
		if transfer.FromPubkey != "" {
			from = transfer.FromPubkey
		} else {
			from = transfer.From
		}
		if transfer.ToPubkey != "" {
			to = transfer.ToPubkey
		} else {
			to = transfer.To
		}
		var outward bool
		if from == pubkey {
			outward = true
		} else {
			outward = false
		}
		date, err := b.Restclient.GetTimestampByBlock(b.RPCNode, transfer.Height)
		if err != nil {
			return errors.Wrap(err, "failed get timestamp by block")
		}
		resp.Transfers = append(resp.Transfers, &pb.Transfer{
			From:    from,
			To:      to,
			Amount:  transfer.Amount,
			Hash:    transfer.DeployHash,
			Height:  int64(transfer.Height),
			Date:    date,
			Outward: outward,
		})
	}

	historyDelegate, err := b.Restclient.GetHistoryDelegate(b.RPCNode, pubkey, Start, End)
	if err != nil {
		return errors.Wrap(err, "failed get history delegate")
	}
	for _, delegate := range historyDelegate.Data {
		date, err := b.Restclient.GetTimestampByBlock(b.RPCNode, delegate.Height)
		if err != nil {
			return errors.Wrap(err, "failed get timestamp by block")
		}
		resp.Delegates = append(resp.Delegates, &pb.DelegateHistory{
			Validator:  delegate.ValidatorPubkey,
			Amount:     delegate.Amount,
			Era:        int64(delegate.Era),
			IsFinished: delegate.IsFinished,
			Height:     int64(delegate.Height),
			Date:       date,
		})
	}
	historyUndelegate, err := b.Restclient.GetHistoryUndelegate(b.RPCNode, pubkey, Start, End)
	if err != nil {
		return errors.Wrap(err, "failed get history undelegate")
	}
	for _, undelegate := range historyUndelegate.Data {
		date, err := b.Restclient.GetTimestampByBlock(b.RPCNode, undelegate.Height)
		if err != nil {
			return errors.Wrap(err, "failed get timestamp by block")
		}
		resp.Undelegates = append(resp.Undelegates, &pb.UndelegateHistory{
			Validator:  undelegate.ValidatorPubkey,
			Amount:     undelegate.Amount,
			Era:        int64(undelegate.Era),
			IsFinished: undelegate.IsFinished,
			Height:     int64(undelegate.Height),
			Date:       date,
		})
	}
	resp.User = out.GetUser()
	resp.MsgId = out.GetMsgId()
	resp.Start = Start
	resp.End = End
	startDate, err := b.Restclient.GetTimestampByBlock(b.RPCNode, uint64(Start))
	if err != nil {
		return errors.Wrap(err, "failed get timestamp by block")
	}
	endDate, err := b.Restclient.GetTimestampByBlock(b.RPCNode, uint64(End))
	if err != nil {
		return errors.Wrap(err, "failed get timestamp by block")
	}
	resp.StartDate = startDate
	resp.EndDate = endDate

	respBytes, err := proto.Marshal(&resp)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "TransferHistoryMsg",
		Data: respBytes,
	}

	return nil

}

func (b *BotMain) SendLoadingHistoryBalance(out *pb.TgBalanceButtonHistory) error {
	loadresp := pb.LoadHistoryResponse{
		User:  out.GetUser(),
		MsgId: out.GetMsgId(),
	}
	loadrespBytes, err := proto.Marshal(&loadresp)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "LoadHistoryMsg",
		Data: loadrespBytes,
	}
	return nil
}

func (b *BotMain) SendLoadingRewards(out *pb.TgButtonRewardsHistory) error {
	loadresp := pb.LoadHistoryResponse{
		User:  out.GetUser(),
		MsgId: out.GetMsgId(),
	}
	loadrespBytes, err := proto.Marshal(&loadresp)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "LoadHistoryMsg",
		Data: loadrespBytes,
	}
	return nil
}

func (b *BotMain) HandleRewards(msg tggateway.TgMessageMsg) error {
	out := pb.TgButtonRewardsHistory{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	loggedin, err := b.CheckLogin(out.GetUser())
	if err != nil {
		return err
	}
	if !loggedin {
		return nil
	}
	err = b.SendLoadingRewards(&out)
	if err != nil {
		return err
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	pubkey := u.PublicKey
	state, err := b.Restclient.GetState(b.RPCNode)
	if err != nil {
		return errors.Wrap(err, "failed get state")
	}
	var Start, End int
	var isFirst = false
	if out.GetStart() == -1 || out.GetEnd() > int64(state.CurrentEra)-1 {
		Start = state.CurrentEra - defaultEraRange
		End = state.CurrentEra - 1
		isFirst = true
	} else {

		Start = int(out.GetStart())
		End = int(out.GetEnd())
	}
	historyRewards, err := b.Restclient.GetRewardsByEra(b.RPCNode, pubkey, Start, End)
	if err != nil {
		return errors.Wrap(err, "failed get history rewards")
	}
	// log.Println(historyRewards)
	resp := pb.RewardsHistoryResponse{}
	resp.IsFirst = isFirst
	for _, reward := range historyRewards.Rewards {
		amount, err := strconv.ParseFloat(reward.Amount, 64)
		if err != nil {
			return errors.Wrap(err, "failed parse amount of reward")
		}
		resp.Rewards = append(resp.Rewards, &pb.Reward{
			Validator: reward.Validator,
			Amount:    amount,
			Era:       int64(reward.Era),
			Timestamp: reward.Timestamp,
		})
	}
	resp.User = out.GetUser()
	resp.MsgId = out.GetMsgId()
	resp.Start = int64(Start)
	resp.End = int64(End)
	startDate, err := b.Restclient.GetTimestampByBlock(b.RPCNode, uint64(Start))
	if err != nil {
		return errors.Wrap(err, "failed get timestamp by block")
	}
	endDate, err := b.Restclient.GetTimestampByBlock(b.RPCNode, uint64(End))
	if err != nil {
		return errors.Wrap(err, "failed get timestamp by block")
	}
	resp.StartDate = startDate
	resp.EndDate = endDate

	respBytes, err := proto.Marshal(&resp)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "RewardsHistoryMsg",
		Data: respBytes,
	}

	return nil

}
