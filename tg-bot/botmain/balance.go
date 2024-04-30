package botmain

import (
	"context"

	"github.com/Simplewallethq/tg-bot/ent/user"
	pb "github.com/Simplewallethq/tg-bot/tggateway/proto"
	tggateway "github.com/Simplewallethq/tg-bot/tggateway/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func (b *BotMain) HandleBalance(msg tggateway.TgMessageMsg) error {
	out := pb.TgBalanceButton{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	loggedin, err := b.CheckLogin(out.GetFrom())
	if err != nil {
		return err
	}
	if !loggedin {
		return nil
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	err = b.SendLoadingBalanceResponse(&out)
	if err != nil {
		return errors.Wrap(err, "failed send loading balance response")
	}

	var total float64
	balance, err := b.Restclient.GetBalance(b.RPCNode, u.PublicKey)
	if err != nil {
		return errors.Wrap(err, "failed get balance")
	}
	total += balance

	delegatedBalance, totalDelegated, err := b.CalculateDelegated(u.PublicKey)
	if err != nil {
		return errors.Wrap(err, "failed get delegated balance")
	}
	total += totalDelegated

	price, err := b.Restclient.GetPrice()
	if err != nil {
		return errors.Wrap(err, "failed get price")
	}

	beingDelegatedBalance, _, err := b.CalculateBeingDelegated(u.PublicKey)
	if err != nil {
		return errors.Wrap(err, "failed get being delegated balance")
	}

	beingUndelegatedBalance, totalBeingUNDg, err := b.CalculateBeingUndelegated(u.PublicKey)
	if err != nil {
		return errors.Wrap(err, "failed get being undelegated balance")
	}
	total += totalBeingUNDg

	totalUSD := total * price

	data := pb.BalanceResponse{
		User:                    out.GetFrom(),
		MsgId:                   out.GetMsgId(),
		Balance:                 balance,
		DelegatedBalance:        delegatedBalance,
		BeingDelegatedBalance:   beingDelegatedBalance,
		BeingUndelegatedBalance: beingUndelegatedBalance,
		Total:                   total,
		TotalUSD:                totalUSD,
		Price:                   price,
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "BalanceMsg",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) CalculateDelegated(pubkey string) (*pb.DelegatedBalance, float64, error) {
	var total float64
	delegatedBalance, err := b.Restclient.GetBalanceDelegated(b.RPCNode, pubkey)
	if err != nil {
		return nil, 0.0, errors.Wrap(err, "failed get delegated balance")
	}
	var pbDelegatedBalance pb.DelegatedBalance
	for _, v := range delegatedBalance.Data {
		pbDelegatedBalance.Data = append(pbDelegatedBalance.Data, &pb.DelegatedBalanceData{
			Validator: v.Validator,
			Amount:    v.Amount,
		})
		total += v.Amount
	}
	return &pbDelegatedBalance, total, nil
}

func (b *BotMain) CalculateBeingDelegated(pubkey string) (*pb.BeingDelegatedBalance, float64, error) {
	var total float64
	beingDelegatedBalance, err := b.Restclient.GetBalanceBeingDelegated(b.RPCNode, pubkey)
	if err != nil {
		return nil, 0.0, errors.Wrap(err, "failed get being delegated balance")
	}
	var pbBeingDelegatedBalance pb.BeingDelegatedBalance
	for _, v := range beingDelegatedBalance.Data {
		pbBeingDelegatedBalance.Data = append(pbBeingDelegatedBalance.Data, &pb.BeingDelegatedBalanceData{
			Validator:             v.ValidatorPubkey,
			EraDelegationFinished: int64(v.EraDelegationFinished),
			Amount:                v.Amount,
		})
		total += v.Amount
	}

	return &pbBeingDelegatedBalance, total, nil
}

func (b *BotMain) CalculateBeingUndelegated(pubkey string) (*pb.BeingUndelegatedBalance, float64, error) {
	var total float64
	beingUndelegatedBalance, err := b.Restclient.GetBalanceBeingUndelegated(b.RPCNode, pubkey)
	if err != nil {
		return nil, 0.0, errors.Wrap(err, "failed get being undelegated balance")
	}
	var pbBeingUndelegatedBalance pb.BeingUndelegatedBalance
	for _, v := range beingUndelegatedBalance.Data {
		pbBeingUndelegatedBalance.Data = append(pbBeingUndelegatedBalance.Data, &pb.BeingUndelegatedBalanceData{
			Validator:               v.ValidatorPubkey,
			EraUndelegationFinished: int64(v.EraUndelegationFinished),
			Amount:                  v.Amount,
		})
		total += v.Amount
	}
	return &pbBeingUndelegatedBalance, total, nil
}

func (b *BotMain) SendLoadingBalanceResponse(out *pb.TgBalanceButton) error {
	loadresp := pb.LoadBalanceResponse{
		User:  out.GetFrom(),
		MsgId: out.GetMsgId(),
	}
	loadrespBytes, err := proto.Marshal(&loadresp)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "LoadBalanceMsg",
		Data: loadrespBytes,
	}
	return nil
}
