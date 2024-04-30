package blockchain

import (
	"math/big"
	"time"

	"github.com/make-software/casper-go-sdk/types"
)

type TransferResponse struct {
	DeployHash string `json:"deploy_hash"`
	From       string `json:"from"`
	FromPubKey string `json:"from_pub_key"`
	To         string `json:"to"`
	ToPubKey   string `json:"to_pub_key"`
	Amount     string `json:"amount"`
	Gas        string `json:"gas"`
	Height     uint64 `json:"height"`
	Memo       uint64 `json:"memo"`
}

type ValidatorsResponse struct {
	Validators []Validator
	Decimals   int
}
type Validator struct {
	Address string
	Fee     float32
	Active  bool
}

// Return dictionary: [integer: amount being staked in motes, string: validator which address is staking to, integer: era when finished, integer: amount of decimals in token]
type UnstakingBalance struct {
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator_pubkey"`
	Era             int    `json:"era"`
}

type HistroyUnstake struct {
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator_pubkey"`
	Era             int    `json:"era"`
	Finished        bool   `json:"is_finished"`
	Height          uint64 `json:"height"`
}

type HistroyStake struct {
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator_pubkey"`
	Era             int    `json:"block_height"`
	Finished        bool   `json:"is_finished"`
	Height          uint64 `json:"height"`
}

type BeingStaked struct {
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator_pubkey"`
	Era             int    `json:"era"`
}

type RewardsByBlock struct {
	Amount          *big.Int `json:"amount"`
	ValidatorPubkey string   `json:"validator_pubkey"`
	Block           uint64   `json:"block"`
}

type RewardsByEra struct {
	Amount          *big.Int  `json:"amount"`
	ValidatorPubkey string    `json:"validator_pubkey"`
	Era             uint32    `json:"era"`
	Date            time.Time `json:"timestamp"`
}

type BlockEvents struct {
	Transfers   []TransferResponse
	Delegates   []DelegateData
	Undelegates []UndelegateData
	Rewards     []RewardsData
	Era         uint32
	Decimals    int
	Date        time.Time `json:"timestamp"`
}

type DelegateData struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator_pubkey"`
	Era             int    `json:"era"`
	Finished        bool   `json:"is_finished"`
}

type UndelegateData struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ValidatorPubkey string `json:"validator_pubkey"`
	Era             int    `json:"block_height"`
	Finished        bool   `json:"is_finished"`
}

type RewardsData struct {
	Delagator string   `json:"delagator"`
	Validator string   `json:"validator"`
	Amount    *big.Int `json:"amount"`
}
type Client interface {
	// GetCurrentState returns the current block height and current era.
	GetCurrentState(rpc string) (int, int, error)

	// GetTimestampByEra returns the estimated time when the specified era starts (up to 100 eras in the future).
	GetTimestampByEra(rpc string, era int) (time.Time, bool, error)

	// GetTimestampByBlock returns the estimated time when the specified block starts (up to 1000 blocks in the future).
	GetTimestampByBlock(rpc string, block int) (time.Time, bool, error)

	// GetByTimestampEra returns the era that includes the specified time.
	GetEraByTimestamp(rpc string, timestamp time.Time) (int, error)

	// GetByTimestampBlock returns the block that includes the specified time.
	GetBlockByTimestamp(rpc string, timestamp time.Time) (int, error)

	// IsAddress returns true if the provided address is a valid chain address, otherwise false.
	IsAddress(address string) bool

	// GetBalanceMain returns the current wallet balance in motes (wei) and decimals of the token as an integer.
	GetBalanceBase(rpc string, address string) (big.Int, int, error)

	// GetBalanceStaked returns the amount staked in motes, the validator staked to, and the amount of decimals in the token.
	GetBalanceStaked(rpc string, address string) (map[string]big.Int, int, error)

	// GetBalanceBeingStaked returns the amount being staked in motes, the validator staked to, the era when finished, and the amount of decimals in the token.
	GetBalanceBeingStaked(rpc string, address string) ([]BeingStaked, int, error)

	// GetBalanceBeingUnstaked returns the amount being unstaked in motes, the validator staked to, the era when finished, and the amount of decimals in the token.
	GetBalanceBeingUnstaked(rpc string, address string) ([]UnstakingBalance, int, error)

	// GetPriceMainCoin returns the token price in USD at the specified time (using a separate data provider).
	GetPriceMainCoin() (float64, error)

	// GetEraBounds returns the block height of the first and last block of the specified era.
	GetEraBounds(rpc string, era int) (int64, int64, error)

	// CalculateCurrentChainAPY returns the current APY for the specified chain.
	CalculateCurrentChainAPY(rpc string, chain string) (float64, error)

	// GetHistoryDelegate returns the amount staked in motes, the validator staked to, the first era when accrual started, and the amount of decimals in the token.
	GetHistoryDelegate(rpc string, address string, blockStart int, blockEnd int) ([]HistroyStake, int, error)

	// GetHistoryUndelegate returns the amount staked in motes, the validator unstaked from, the first era when accrual started, and the amount of decimals in the token.
	GetHistoryUndelegate(rpc string, address string, blockStart int, blockEnd int) ([]HistroyUnstake, int, error)

	// GetHistoryTransfers returns the amount transferred in motes, the address to/from, and the amount of decimals in the token.
	GetHistoryTransfers(rpc string, address string, blockStart int, blockEnd int) ([]TransferResponse, int, error)

	// GetRewardsByBlock returns the amount of rewards during the specified period.
	GetRewardsByBlock(rpc string, address string, blockStart int64, blockEnd int64) ([]RewardsByBlock, int, error)
	GetRewardsByEra(rpc string, address string, era_start int, era_end int) ([]RewardsByEra, int, error)

	GetAPRByERA(rpc string, address string, era_start int, era_end int) (float64, error)

	//CheckChain returns true if the provided chain is the current node chain, otherwise false.
	CheckChain(rpc string, chain string) (bool, error)
	GetBlockAllEvents(rpc string, height uint64, transfers bool, delegates bool, undelegates bool, rewards bool) (BlockEvents, error)
	GetValidators(rpc string) (ValidatorsResponse, error)

	PutDeploy(rpc string, deploy types.Deploy) (string, error)
}
