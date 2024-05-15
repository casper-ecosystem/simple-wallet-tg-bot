package notificator

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/botmain/restclient"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/invoice"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/rewardsdata"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	csprsdk "github.com/make-software/casper-go-sdk/casper"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type Notificator struct {
	DB           *ent.Client
	MessagesChan chan types.TgResponseMsg
	Restclient   *restclient.Client
	logger       *logrus.Logger
	RPCNode      string
}

func NewNotificator(DB *ent.Client, MessagesChan chan types.TgResponseMsg, resthost string, rpcNode string, logger *logrus.Logger) *Notificator {
	return &Notificator{
		DB:           DB,
		MessagesChan: MessagesChan,
		Restclient:   restclient.NewClient(resthost),
		logger:       logger,
		RPCNode:      rpcNode,
	}
}
func NewNotificatorWithToken(DB *ent.Client, MessagesChan chan types.TgResponseMsg, resthost string, rpcNode string, RESTtoken string, logger *logrus.Logger) *Notificator {
	return &Notificator{
		DB:           DB,
		MessagesChan: MessagesChan,
		Restclient:   restclient.NewClientWithToken(resthost, RESTtoken),
		logger:       logger,
		RPCNode:      rpcNode,
	}
}

func (n *Notificator) Start() {
	// ticker 60 second
	err := n.ScanNetwork()
	if err != nil {
		n.logger.Error(err)
	}
	log.Println("start notificator")
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		err := n.ScanNetwork()
		if err != nil {
			n.logger.Error(err)
		}

		// select {
		// case <-ticker.C:
		// 	err := n.ScanNetwork()
		// 	if err != nil {
		// 		n.logger.Error(err)
		// 		return err
		// 	}
		// }
	}
}

func (n *Notificator) ScanNetwork() error {
	state, err := n.Restclient.GetState(n.RPCNode)
	if err != nil {
		return err
	}
	optcount, err := n.DB.Settings.Query().Count(context.Background())
	if err != nil {
		return err
	}
	if optcount == 0 {
		_, err = n.DB.Settings.Create().SetLastScannedBlockNotificator(state.BlockHeight).Save(context.Background())
		if err != nil {
			return err
		}
	}

	settings, err := n.DB.Settings.Query().Only(context.Background())
	if err != nil {
		return err
	}
	users, err := n.DB.User.Query().Where(user.Notify(true)).All(context.Background())
	if err != nil {
		return err
	}
	invUsers, err := n.DB.Invoice.Query().Where(invoice.Active(true)).All(context.Background())
	if err != nil {
		return err
	}
	// pubkeys_to_check := make([]string, 0, len(users))
	// for _, user := range users {
	// 	pubkeys_to_check = append(pubkeys_to_check, user.PublicKey)
	// }
	log.Println("scan network", settings.LastScannedBlockNotificator, state.BlockHeight)
	var sendedBalance map[int64]struct{}
	for settings.LastScannedBlockNotificator < state.BlockHeight {
		log.Println("scan network", settings.LastScannedBlockNotificator)
		sendedBalance, err = n.handleEvents(uint64(settings.LastScannedBlockNotificator), users, invUsers)
		if err != nil {
			return err
		}

		settings, err = settings.Update().SetLastScannedBlockNotificator(settings.LastScannedBlockNotificator + 1).Save(context.Background())
		if err != nil {
			return err
		}

	}
	err = n.handleBalance(uint64(state.BlockHeight), users, sendedBalance)
	if err != nil {
		n.logger.Error(err)
		return err
	}

	return nil
}

func (n *Notificator) handleEvents(height uint64, users []*ent.User, invUsers []*ent.Invoice) (map[int64]struct{}, error) {
	sendedBalance := make(map[int64]struct{})
	events, err := n.Restclient.GetBlockEvents(n.RPCNode, height, true, true, true, true)
	if err != nil {
		return sendedBalance, err
	}

	divisor := new(big.Float).SetFloat64(math.Pow10(events.Decimals))
	for _, transfer := range events.Transfers {
		// if slice_contains_users(users, transfer.FromPubkey) {
		// 	log.Println("new transaction", transfer.FromPubkey, transfer.ToPubkey, transfer.Amount)
		// }
		if uid, ok := sliceContainsUsers(users, transfer.ToPubkey); ok {
			log.Println("incoming transaction", transfer.FromPubkey, transfer.ToPubkey, transfer.Amount)
			bigAmount, success := new(big.Int).SetString(transfer.Amount, 10)
			if !success {
				return sendedBalance, errors.New("invalid amount format")
			}

			floatAmount := new(big.Float).SetInt(bigAmount)
			adjustedAmount := new(big.Float).Quo(floatAmount, divisor)
			f64bAmount, _ := adjustedAmount.Float64()
			//log.Println("incoming transaction", transfer.FromPubkey, transfer.ToPubkey, f64bAmount)
			var from string
			if transfer.FromPubkey != "" {
				from = transfer.FromPubkey
			} else {
				from = transfer.From
			}
			gas, err := strconv.ParseFloat(transfer.Gas, 64)
			if err != nil {
				return sendedBalance, err
			}
			actualBalance, err := n.Restclient.GetBalance(n.RPCNode, transfer.ToPubkey)
			if err != nil {
				return sendedBalance, err
			}
			sendedBalance[uid] = struct{}{}
			data := pb.NotificationNewTransfer{
				User: &pb.User{
					Id: uid,
				},
				From:    from,
				To:      transfer.ToPubkey,
				Amount:  f64bAmount,
				Hash:    transfer.DeployHash,
				Gas:     gas,
				Balance: actualBalance,
			}
			dataBytes, err := proto.Marshal(&data)
			if err != nil {
				return sendedBalance, err
			}
			n.MessagesChan <- types.TgResponseMsg{
				Name: "NotifyNewTransfer",
				Data: dataBytes,
			}

		}
		if invItem, ok := sliceContainsInvoices(invUsers, transfer); ok {
			if invItem.Memo == transfer.Memo {
				log.Println("new invoice pay", transfer.FromPubkey, transfer.ToPubkey, transfer.Amount)
				bigAmount, success := new(big.Float).SetString(transfer.Amount)
				if !success {
					return sendedBalance, errors.New("invalid amount format")
				}
				bigAmountInvoice, success := new(big.Float).SetString(invItem.Amount)
				if !success {
					return sendedBalance, errors.New("invalid amount format")
				}

				bigAmount = bigAmount.Quo(bigAmount, big.NewFloat(1000000000))
				// log.Println(bigAmount.String())
				// log.Println(bigAmountInvoice.String())
				var percDiff float64
				dv := new(big.Float).Quo(bigAmount, bigAmountInvoice)
				log.Println("dv after div two big", dv)
				dv = dv.Mul(dv, big.NewFloat(100))
				log.Println("dv after mul", dv)

				if dv.Cmp(big.NewFloat(100)) == 1 {
					fl, _ := dv.Float64()
					percDiff = fl - 100
				} else {

					percDiff, _ = dv.Float64()
				}
				var from string
				if transfer.FromPubkey != "" {
					from = transfer.FromPubkey
				} else {
					from = transfer.From
				}
				paymentsDB, err := n.DB.Invoices_payments.Create().
					SetInvoice(invItem).SetAmount(bigAmount.String()).SetFrom(from).Save(context.Background())
				if err != nil {
					return sendedBalance, err
				}

				log.Println(percDiff)
				if 100-percDiff < 10 {
					_, err := invItem.Update().SetPaid(invItem.Paid + 1).Save(context.Background())
					if err != nil {
						return nil, err
					}
					_, err = paymentsDB.Update().SetCorrect(true).Save(context.Background())
					if err != nil {
						return sendedBalance, err
					}
				}

			}

		}

	}

	for _, delegate := range events.Delegates {
		if uid, ok := sliceContainsUsers(users, delegate.Address); ok {
			bigAmount, success := new(big.Int).SetString(delegate.Amount, 10)
			if !success {
				return sendedBalance, errors.New("invalid amount format")
			}

			floatAmount := new(big.Float).SetInt(bigAmount)
			adjustedAmount := new(big.Float).Quo(floatAmount, divisor)
			f64bAmount, _ := adjustedAmount.Float64()

			log.Println("new delegate", delegate.Address, f64bAmount, uid)
			actualBalance, err := n.Restclient.GetBalance(n.RPCNode, delegate.Address)
			if err != nil {
				return sendedBalance, err
			}
			sendedBalance[uid] = struct{}{}
			data := pb.NotificationNewDelegate{
				User: &pb.User{
					Id: uid,
				},
				Validator: delegate.ValidatorPubkey,
				Amount:    f64bAmount,
				Balance:   actualBalance,
				Height:    int64(delegate.BlockHeight),
			}
			dataBytes, err := proto.Marshal(&data)
			if err != nil {
				return sendedBalance, err
			}
			n.MessagesChan <- types.TgResponseMsg{
				Name: "NotifyNewDelegate",
				Data: dataBytes,
			}
		}
	}

	for _, undelegate := range events.Undelegates {
		if uid, ok := sliceContainsUsers(users, undelegate.Address); ok {
			bigAmount, success := new(big.Int).SetString(undelegate.Amount, 10)
			if !success {
				return sendedBalance, errors.New("invalid amount format")
			}

			floatAmount := new(big.Float).SetInt(bigAmount)
			adjustedAmount := new(big.Float).Quo(floatAmount, divisor)
			f64bAmount, _ := adjustedAmount.Float64()

			log.Println("new undelegate", undelegate.Address, f64bAmount, uid)
			actualBalance, err := n.Restclient.GetBalance(n.RPCNode, undelegate.Address)
			if err != nil {
				return sendedBalance, err
			}
			sendedBalance[uid] = struct{}{}
			data := pb.NotificationNewUndelegate{
				User: &pb.User{
					Id: uid,
				},
				Validator: undelegate.ValidatorPubkey,
				Amount:    f64bAmount,
				Era:       int64(undelegate.Era),
				Balance:   actualBalance,
			}
			dataBytes, err := proto.Marshal(&data)
			if err != nil {
				return sendedBalance, err
			}
			n.MessagesChan <- types.TgResponseMsg{
				Name: "NotifyNewUndelegate",
				Data: dataBytes,
			}
		}
	}

	usersAll, err := n.DB.User.Query().All(context.Background())
	if err != nil {
		return sendedBalance, err
	}
	usersForRewards := make([]*ent.User, 0)
	for _, u := range usersAll {
		if u.NotifyLastTime.Before(time.Now().Add(-time.Duration(u.NotifyTime)*time.Hour)) && u.NotifyTime != 0 {
			usersForRewards = append(usersForRewards, u)
		}
	}

	usersForRewardsDb := make([]*ent.User, 0)
	for _, u := range usersAll {
		if u.NotifyTime != 0 {
			usersForRewardsDb = append(usersForRewardsDb, u)
		}
	}
	//log.Println("users for rewards", users_for_rewards)
	var mapRewards = make(map[int64]*pb.NotificationNewReward)
	for _, reward := range events.Rewards {
		if id, ok := sliceContainsUsers(usersForRewardsDb, reward.Delagator); ok {
			log.Println("new reward", reward.Delagator, reward.Validator, reward.Amount)
			bigAmount, success := new(big.Int).SetString(reward.Amount, 10)
			if !success {
				return sendedBalance, errors.New("invalid amount format")
			}

			floatAmount := new(big.Float).SetInt(bigAmount)
			adjustedAmount := new(big.Float).Quo(floatAmount, divisor)
			u, err := n.DB.User.Query().Where(user.ID(id)).Only(context.Background())
			if err != nil {
				return sendedBalance, errors.New("failed get u")
			}
			RewardsDB := u.QueryRewardsData()

			validatorDB, err := RewardsDB.Where(rewardsdata.ValidatorEQ(reward.Validator)).All(context.Background())
			if err != nil {
				log.Println(err)
			}
			if len(validatorDB) == 0 {
				log.Println("validatorDB = 0")
				err = n.DB.RewardsData.Create().SetValidator(reward.Validator).SetAmount(adjustedAmount.String()).SetLastReward(time.Now()).SetFirstEra(int64(events.EraID)).SetFirstEraTimestamp(events.Date).SetLastEra(int64(events.EraID)).SetLastEraTimestamp(events.Date).SetOwner(u).Exec(context.Background())
				if err != nil {
					//log.Println(err)
					return sendedBalance, errors.New("failed set reward")
				}
			} else {
				log.Println("validatorDB not 0")
				strAmount := validatorDB[0].Amount
				bigAmountOld, yes := big.NewFloat(0).SetString(strAmount)
				if !yes {
					return sendedBalance, fmt.Errorf("error parse amount")
				}
				// big_amount_new, yes := big.NewFloat(0).SetString(reward.Amount)
				// if yes != true {
				// 	return sendedBalance, fmt.Errorf("error parse amount")
				// }

				finalAmount := bigAmountOld.Add(bigAmountOld, adjustedAmount)
				log.Println("OLD AMOUNT:", bigAmountOld.String())
				//log.Println("NEW AMOUNT:", big_amount_new.String())
				log.Println("FINAL AMOUNT:", finalAmount.String())
				err = validatorDB[0].Update().SetAmount(finalAmount.String()).SetLastEra(int64(events.EraID)).SetLastEraTimestamp(events.Date).Exec(context.Background())
				if err != nil {
					return sendedBalance, err
				}
			}
		}
	}
	for _, u := range usersForRewards {
		RewardsDB, err := u.QueryRewardsData().All(context.Background())
		if err != nil {
			return sendedBalance, errors.New("failed get u")
		}
		for _, reward := range RewardsDB {
			amount, err := strconv.ParseFloat(reward.Amount, 64)
			if err != nil {
				return sendedBalance, errors.New("failed parse float")
			}

			if val, ok := mapRewards[u.ID]; ok {

				val.Rewards = append(val.Rewards, &pb.Reward{
					Validator:      reward.Validator,
					Amount:         amount,
					FirstEra:       uint64(reward.FirstEra),
					LastEra:        uint64(reward.LastEra),
					LastRewardTime: reward.LastReward.UTC().String(),
				})
				if uint64(reward.FirstEra) < val.FirstEra {
					val.FirstEra = uint64(reward.FirstEra)
					val.FirstEraTimestamp = reward.FirstEraTimestamp
				}
				if uint64(reward.LastEra) > val.LastEra {
					val.LastEra = uint64(reward.LastEra)
					val.LastEraTimestamp = reward.LastEraTimestamp
				}
			} else {
				mapRewards[u.ID] = &pb.NotificationNewReward{
					User: &pb.User{
						Id: u.ID,
					},
					FirstEra:          uint64(reward.FirstEra),
					FirstEraTimestamp: reward.FirstEraTimestamp,
					LastEraTimestamp:  reward.LastEraTimestamp,
					LastEra:           uint64(reward.LastEra),
					Rewards: []*pb.Reward{{Validator: reward.Validator,
						Amount:         amount,
						FirstEra:       uint64(reward.FirstEra),
						LastEra:        uint64(reward.LastEra),
						LastRewardTime: reward.LastReward.UTC().String()}},
				}
			}

		}

	}

	for _, val := range mapRewards {
		u, err := n.DB.User.Query().Where(user.IDEQ(val.User.Id)).Only(context.Background())
		if err != nil {
			n.logger.Error(err)
			return sendedBalance, err
		}
		u, err = u.Update().SetNotifyLastTime(time.Now()).Save(context.Background())
		if err != nil {
			n.logger.Error(err)

			return sendedBalance, err
		}
		delegated, err := n.Restclient.GetBalanceDelegated(n.RPCNode, u.PublicKey)
		if err != nil {
			n.logger.Error(err)
			return sendedBalance, err
		}
		var staked float64
		for _, delegation := range delegated.Data {
			staked += delegation.Amount
		}
		val.Delegated = staked

		dataBytes, err := proto.Marshal(val)
		if err != nil {
			return sendedBalance, err
		}
		n.MessagesChan <- types.TgResponseMsg{
			Name: "NotifyNewRewards",
			Data: dataBytes,
		}
		rewardsDB, err := u.QueryRewardsData().All(context.Background())
		if err != nil {
			return sendedBalance, errors.New("failed get REWARDS table")
		}
		for _, rew := range rewardsDB {
			err = n.DB.RewardsData.DeleteOne(rew).Exec(context.Background())
			if err != nil {
				return sendedBalance, errors.New("failed clean rewards table")
			}
		}
		//n.DB.RewardsData.Delete().Where(rewardsdata.Ow()

	}

	return sendedBalance, nil
}

func (n *Notificator) handleBalance(height uint64, users []*ent.User, sendedBalance map[int64]struct{}) error {
	for _, u := range users {
		actualBalance, err := n.Restclient.GetBalance(n.RPCNode, u.PublicKey)
		if err != nil {
			return err
		}
		balanceExist, err := u.QueryBalance().Exist(context.Background())
		if err != nil {
			return err
		}
		if !balanceExist {
			_, err := n.DB.Balances.Create().SetBalance(actualBalance).SetHeight(height).SetOwner(u).Save(context.Background())
			if err != nil {
				return err
			}
		} else {
			balance, err := u.QueryBalance().Only(context.Background())
			if err != nil {
				return err
			}
			if balance.Balance != actualBalance {
				log.Println("balance changed", balance.Balance, actualBalance)
				if _, ok := sendedBalance[u.ID]; !ok {
					data := pb.NotificationNewBalance{
						User: &pb.User{
							Id: u.ID,
						},
						Balance:    actualBalance,
						OldBalance: balance.Balance,
					}
					dataBytes, err := proto.Marshal(&data)
					if err != nil {
						return err
					}
					n.MessagesChan <- types.TgResponseMsg{
						Name: "NotifyNewBalance",
						Data: dataBytes,
					}
				}
				_, err = balance.Update().SetHeight(height).SetBalance(actualBalance).Save(context.Background())
				if err != nil {
					n.logger.Error(err)
					return err
				}
			}
		}
	}
	return nil
}

func sliceContainsUsers(slice []*ent.User, val string) (int64, bool) {
	for _, item := range slice {
		if item.PublicKey == val {
			return item.ID, true
		}
	}
	return 0, false
}

func sliceContainsInvoices(slice []*ent.Invoice, tr restclient.Transfer) (*ent.Invoice, bool) {
	for _, item := range slice {
		hash, err := csprsdk.NewPublicKey(item.Address)
		if err != nil {
			return nil, false
		}
		if item.Address == tr.ToPubkey || hash.AccountHash().String() == tr.To {
			return item, true
		}
	}
	return nil, false
}
