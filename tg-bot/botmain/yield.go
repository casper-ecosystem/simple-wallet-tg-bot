package botmain

import (
	"context"
	"log"
	"time"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/botmain/restclient"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/botmain/types"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	tggateway "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func (b *BotMain) HandleYield(msg tggateway.TgMessageMsg) error {
	out := pb.TgYieldButton{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	if !out.IsRestored {
		loggedin, err := b.CheckLogin(out.GetUser())
		if err != nil {
			return err
		}
		if !loggedin {
			return nil
		}
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	log.Println("user", u.PublicKey, "press yield")
	state, err := b.Restclient.GetState(b.RPCNode)
	if err != nil {
		return errors.Wrap(err, "failed get state")
	}
	err = b.SendYieldLoadingResponse(out.GetUser(), out.GetMsgId(), "")
	if err != nil {
		return errors.Wrap(err, "failed send loading balance response")
	}
	err = b.TaskRecoverer.SetYieldTask(&out)
	if err != nil {
		var errTooManyTasks *types.TooManyTasksError
		if errors.As(err, &errTooManyTasks) {
			return err // Specifically handle the TooManyTasksError
		} else {
			return err
		}
	}
	defer func() {
		err = b.TaskRecoverer.ClearUserTasks(out.GetUser().GetId())
		if err != nil {
			log.Println(errors.Wrap(err, "failed clear user tasks"))
		}
	}()
	if err != nil {
		log.Println("failed set yield task", err)
		return err
	}
	end := state.CurrentEra - 1
	//7 day befor now
	startEraByTime, err := b.Restclient.GetEraByTimestamp(b.RPCNode, time.Now().AddDate(0, 0, -7))
	if err != nil {
		return errors.Wrap(err, "failed get era by time")
	}
	rewards, err := b.Restclient.GetRewardsSummByEra(b.RPCNode, u.PublicKey, startEraByTime.EraID, end)
	if err != nil {
		return errors.Wrap(err, "failed get rewards")
	}
	log.Println("rewards:", rewards)
	apy, err := b.Restclient.CalculateCurrentChainAPY(b.RPCNode)
	if err != nil {
		return errors.Wrap(err, "failed get apy")
	}
	price, err := b.Restclient.GetPrice()
	if err != nil {
		return errors.Wrap(err, "failed to get price of cspr token")
	}
	rewardsusd := rewards * price

	delegated, err := b.Restclient.GetBalanceDelegated(b.RPCNode, u.PublicKey)
	if err != nil {
		return errors.Wrap(err, "failed to get delegated balance")
	}
	validators, err := b.Restclient.GetValidators(b.RPCNode)
	if err != nil {
		return errors.Wrap(err, "failed to get validators")
	}
	est, proj365, proj30, err := b.CalculateEstimatedYield(u.PublicKey, apy, delegated, validators)
	if err != nil {
		return errors.Wrap(err, "failed calculate estimated yield")
	}

	validatorsData, totalDelegated, err := b.GetValidatorsData(u.PublicKey, apy, delegated, validators)
	if err != nil {
		return errors.Wrap(err, "failed to get validators data")
	}
	data := pb.YieldResponse{
		User:       out.GetUser(),
		MsgId:      out.GetMsgId(),
		Rewards:    rewards,
		NetworkApy: apy,
		RewardsUSD: rewardsusd,
	}
	data.Estim365Days = proj365
	data.Estim30Days = proj30
	data.Estim365DaysUSD = proj365 * price
	data.Estim30DaysUSD = proj30 * price
	data.EstimApy = proj365 / totalDelegated * 100

	for _, v := range est {
		data.Estimates = append(data.Estimates, &pb.YieldEstimate{
			Validator: v.Validator,
			Amount:    v.EstimatedYield,
		})
	}
	for _, v := range validatorsData {
		data.ValidatorsData = append(data.ValidatorsData, &pb.YieldValidatorData{
			Address: v.Validator,
			Amount:  v.Amount,
			Fee:     v.Fee,
			Apy:     v.Apy,
		})
	}
	data.TotalDelegated = totalDelegated

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "YieldMsg",
		Data: dataBytes,
	}
	err = b.TaskRecoverer.ClearUserTasks(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed clear user tasks")
	}

	return nil
}

type EstimatedYield struct {
	Validator        string
	EstimatedYield   float64
	Estimated30Yield float64
}

func (b *BotMain) CalculateEstimatedYield(address string, apy float64, delegated restclient.DelegatedBalance, validators restclient.ValidatorsResponse) ([]EstimatedYield, float64, float64, error) {
	var result []EstimatedYield
	var Projected365 float64
	var Projected30 float64
	// delegated, err := b.Restclient.GetBalanceDelegated(b.RPCNode, address)
	// if err != nil {
	// 	return result, err
	// }
	// validators, err := b.Restclient.GetValidators(b.RPCNode)
	// if err != nil {
	// 	return result, err
	// }
	for _, del := range delegated.Data {
		var fee float32
		fee = 1
		for _, validator := range validators.Validators {
			if del.Validator == validator.Address {
				fee = validator.Fee

				break
			}
		}
		estimated365 := (del.Amount * (apy / 100)) * (1 - float64(fee)/100)
		estimated30 := (del.Amount * (apy / 100)) * (1 - float64(fee)/100) / 365 * 30
		Projected365 += estimated365
		Projected30 += estimated30
		result = append(result, EstimatedYield{
			Validator:      del.Validator,
			EstimatedYield: estimated365,
		})
	}
	return result, Projected365, Projected30, nil
}

type ValidatorsData struct {
	Validator string
	Amount    float64
	Fee       float64
	Apy       float64
}

func (b *BotMain) GetValidatorsData(address string, apy float64, delegated restclient.DelegatedBalance, validators restclient.ValidatorsResponse) ([]ValidatorsData, float64, error) {
	var result []ValidatorsData
	// delegated, err := b.Restclient.GetBalanceDelegated(b.RPCNode, address)
	// if err != nil {
	// 	return result, err
	// }
	// validators, err := b.Restclient.GetValidators(b.RPCNode)
	// if err != nil {
	// 	return result, err
	// }
	var total float64
	for _, del := range delegated.Data {
		var fee float64
		fee = 1
		for _, validator := range validators.Validators {
			if del.Validator == validator.Address {
				fee = float64(validator.Fee)

				break
			}
		}
		total += del.Amount
		result = append(result, ValidatorsData{
			Validator: del.Validator,
			Fee:       fee,
			Amount:    del.Amount,
			Apy:       apy * ((100 - fee) / 100),
		})
	}
	return result, total, nil
}
