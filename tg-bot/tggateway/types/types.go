package types

import (
	"time"

	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
)

type TgMessageMsg struct {
	Name string
	Data []byte
}

type TgResponseMsg struct {
	Name string
	Data []byte
}

type BalanceResponse struct {
	Liquid           float64
	Price            float64
	TotalUSD         float64
	Total            float64
	Delegated        *pb.DelegatedBalance
	BeingDelegated   *pb.BeingDelegatedBalance
	BeingUndelegated *pb.BeingUndelegatedBalance
}

type HistoryResponse struct {
	Start             int64
	StartDate         string
	End               int64
	EndDate           string
	Transfers         []*pb.Transfer
	DelegateHistory   []*pb.DelegateHistory
	UndelegateHistory []*pb.UndelegateHistory
	IsFirst           bool
}

type RewardsHistoryResponse struct {
	Start     int64
	End       int64
	StartDate string
	EndDate   string
	Rewards   []*pb.Reward
	IsFirst   bool
}

type YieldResponse struct {
	TotalRewards   float64
	RewardsUSD     float64
	NetworkApy     float64
	Estimates      []*pb.YieldEstimate
	Validators     []*pb.YieldValidatorData
	TotalDelegated float64
	Proj365Days    float64
	Proj30Days     float64
	Proj365DaysUSD float64
	Proj30DaysUSD  float64
	ProjApy        float64
}

type SettingsResponse struct {
	// bool notifications = 3;
	// int32 notifyTime = 4;
	// int64 LockTimeout=5;
	NotificationsEnabled bool
	NotifyTime           int32
	LockTimeout          int64
	PublicKey            string
}

type NotifySettingsResponse struct {
	NotificationsEnabled bool
	NotyfyTime           int32
}

type Reward struct {
	Validator      string
	Amount         float64
	FirstEra       uint64
	LastEra        uint64
	LastRewardTime string
}

type NotifyNewRewards struct {
	Rewards           []Reward
	Delegated         float64
	FirstEra          uint64
	LastEra           uint64
	FirstEraTimestamp string
	LastEraTimestamp  string
}

type AddressBookResponse struct {
	Offset int64
	Total  int64
	Data   []*pb.AddressRow
}

type DelegateValidatorsResponse struct {
	Offset      int64
	Total       int64
	UserBalance string
	Data        []*pb.ValidatorsRow
}

type AddressBookDetailed struct {
	Name      string
	Address   string
	Id        uint64
	CreatedAt time.Time
}

type UndelegateDelegatesList struct {
	Offset int64
	Total  int64
	Data   []*pb.DelegatesRow
}
