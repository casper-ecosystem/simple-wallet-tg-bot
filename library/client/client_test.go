package client

import (
	"context"
	"errors"
	"log"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/Simplewallethq/source-code/library/blockchain"
	"github.com/google/go-cmp/cmp"
	newsdk "github.com/make-software/casper-go-sdk/casper"
	"github.com/stretchr/testify/assert"
)

var rpc_url string = "http://3.136.227.9:7777/rpc"
var client = NewClient()

func TestRpcClient_GetLatestBlock(t *testing.T) {
	_, _, err := client.GetCurrentState(rpc_url)

	if err != nil {
		t.Errorf("can't get latest block")
	}
}

func TestRpcClient_IsAddress(t *testing.T) { //TODO
	// Define the necessary variables
	real_address := "020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292"
	fake_address := "fdsfsdfds"

	// Test real address
	valid := client.IsAddress(real_address)

	if !valid {
		t.Errorf("Expected valid address for real address, got invalid")
	}

	// Test fake address
	valid = client.IsAddress(fake_address)

	if valid {
		t.Errorf("Expected invalid address for fake address, got valid")
	}
}

func TestRpcClient_BalanceMain(t *testing.T) {
	// Define the necessary variables
	real_address := "017d96b9a63abcb61c870a4f55187a0a7ac24096bdb5fc585c12a686a4d892009e"
	fake_address := "fdsfsdfds"

	// Test real address
	_, _, err := client.GetBalanceBase(rpc_url, real_address)

	if err != nil {
		t.Errorf("Error while get balance: %s", err)
	}

	// Test fake address
	_, _, err = client.GetBalanceBase(rpc_url, fake_address)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	} else {
		var addressInvalidError *AdressInvalidError
		if !errors.As(err, &addressInvalidError) {
			t.Errorf("Expected AddressInvalidError, got: %T", err)
		}
	}

}

func TestRpcClient_BalanceStaked(t *testing.T) {
	// Define the necessary variables
	real_address := "017d96b9a63abcb61c870a4f55187a0a7ac24096bdb5fc585c12a686a4d892009e"
	fake_address := "fdsfsdfds"

	// Test real address
	_, _, err := client.GetBalanceStaked(rpc_url, real_address)

	if err != nil {
		t.Errorf("Error while get balance: %s", err)
	}

	// Test fake address
	_, _, err = client.GetBalanceStaked(rpc_url, fake_address)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	} else {
		var addressInvalidError *AdressInvalidError
		if !errors.As(err, &addressInvalidError) {
			t.Errorf("Expected AddressInvalidError, got: %T", err)
		}
	}

}

func TestRpcClient_getHistoryTransactions(t *testing.T) {
	real_address := "020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292"
	transactions, dec, err := client.GetHistoryTransfers(rpc_url, real_address, 1638650, 1638655)
	if err != nil {
		t.Errorf("Error while getting history transfers: %s", err)
	}
	if dec != CasperDecimals {
		t.Errorf("Decimals is not equal to CasperDecimals")
	}

	// Check the length of the transactions slice
	expectedLength := 1
	if len(transactions) != expectedLength {
		t.Errorf("Expected %d transactions, but got %d", expectedLength, len(transactions))
	}

	// Check properties of the first transaction
	if len(transactions) > 0 {
		firstTransaction := transactions[0]
		expectedDeployHash := "78a13e44a8d9a34ace15acaf00437da7e07fdbd9c629b2544c888873e4825d3a"
		expectedFrom := "cfc1dd3329c25ee0640881d844317dd7306da1cd0adbbdbc764fcebd7a887c0a"
		expectedFromPubkey := "020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292"
		expectedTo := "5a558b721112b6ec722dfb9f350ccd7651889ed19edbc24701eb8612f842fcc7"
		expectedToPubkey := "011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635"
		expectedAmount := "700000000000"
		expectedGas := "0"

		if firstTransaction.DeployHash != expectedDeployHash {
			t.Errorf("Expected DeployHash: %s, but got: %s", expectedDeployHash, firstTransaction.DeployHash)
		}
		if firstTransaction.From != expectedFrom {
			t.Errorf("Expected From: %s, but got: %s", expectedFrom, firstTransaction.From)
		}
		if firstTransaction.FromPubKey != expectedFromPubkey {
			t.Errorf("Expected FromPubkey: %s, but got: %s", expectedFromPubkey, firstTransaction.FromPubKey)
		}
		if firstTransaction.To != expectedTo {
			t.Errorf("Expected To: %s, but got: %s", expectedTo, firstTransaction.To)
		}
		if firstTransaction.ToPubKey != expectedToPubkey {
			t.Errorf("Expected ToPubkey: %s, but got: %s", expectedToPubkey, firstTransaction.ToPubKey)
		}
		if firstTransaction.Amount != expectedAmount {
			t.Errorf("Expected Amount: %s, but got: %s", expectedAmount, firstTransaction.Amount)
		}
		if firstTransaction.Gas != expectedGas {
			t.Errorf("Expected Gas: %s, but got: %s", expectedGas, firstTransaction.Gas)
		}
	}

}

func TestRpcClient_GetRewardsByBlock(t *testing.T) {
	testAddress := "011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635"
	expectedRewards := "44853156"

	rewards, _, err := client.GetRewardsByBlock(rpc_url, testAddress, 1639836, 1639836)

	// Check for errors
	assert.NoError(t, err)

	// Compare the result with the expected value
	expectedRewardsBigInt := new(big.Int)
	expectedRewardsBigInt.SetString(expectedRewards, 10)
	assert.Equal(t, *expectedRewardsBigInt, rewards)
}
func TestRpcClient_GetRewardsByEra(t *testing.T) {
	testAddress := "011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635"
	expectedRewards := "44853156"

	rewards, _, err := client.GetRewardsByEra(rpc_url, testAddress, 8759, 8759)
	if err != nil {
		t.Errorf("Error while getting rewards: %s", err)
	}

	// Compare the result with the expected value
	expectedRewardsBigInt := new(big.Int)
	expectedRewardsBigInt.SetString(expectedRewards, 10)
	assert.Equal(t, *expectedRewardsBigInt, rewards)
}

func TestAverageBlockTime(t *testing.T) {
	var block_start, block_end uint64
	block_start = 1653245
	block_end = 1653255

	avgBlockTime, err := client.averageBlockTime(rpc_url, block_start, block_end)
	if err != nil {
		t.Errorf("averageBlockTime() error: %v", err)
		return
	}

	if avgBlockTime < 0 {
		t.Errorf("averageBlockTime() returned negative duration: %v", avgBlockTime)
		return
	}

	t.Logf("Average block time between blocks %d and %d is %v", block_start, block_end, avgBlockTime)
}

func TestGetTimestampByBlock(t *testing.T) {
	currentBlock, _, err := client.GetCurrentState(rpc_url)
	if err != nil {
		t.Errorf("GetCurrentState() error: %v", err)
		return
	}

	// Test cases
	testCases := []struct {
		name          string
		block         int
		estimated     bool
		expectedError string
	}{
		{
			name:      "Current block",
			block:     currentBlock,
			estimated: false,
		},
		{
			name:      "Future block",
			block:     currentBlock + 500,
			estimated: true,
		},
		{
			name:      "Past block",
			block:     currentBlock - 500,
			estimated: false,
		},
		{
			name:          "Too large block",
			block:         currentBlock + 1500,
			expectedError: "block number is too big",
			estimated:     true,
		},
	}

	client := NewClient()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			timestamp, estimated, err := client.GetTimestampByBlock(rpc_url, tc.block)
			t.Logf("currentBlock: %d, tc.block: %d", currentBlock, tc.block)
			if tc.expectedError == "" {
				if err != nil {
					t.Errorf("GetTimestampByBlock() error: %v", err)
					return
				}
				if estimated != tc.estimated {
					t.Errorf("GetTimestampByBlock() estimated: %v, expected: %v", estimated, tc.estimated)
					return
				}
				t.Logf("Timestamp for block %d is %v", tc.block, timestamp)
			} else {
				if err == nil || err.Error() != tc.expectedError {
					t.Errorf("GetTimestampByBlock() expected error: %s, got: %v", tc.expectedError, err)
				}
			}
		})
	}
}

func TestGetTimestampByEra(t *testing.T) {
	// Replace the following value with a valid RPC URL

	// Test cases
	testCases := []struct {
		name          string
		era           int
		expectedError string
	}{
		{
			name: "Era 8820",
			era:  8820,
		},
		// Add more test cases if needed
	}

	client := NewClient()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			timestamp, _, err := client.GetTimestampByEra(rpc_url, tc.era)
			if tc.expectedError == "" {
				if err != nil {
					t.Errorf("GetTimestampByEra() error: %v", err)
					return
				}
				t.Logf("Timestamp for era %d is %v", tc.era, timestamp)
			} else {
				if err == nil || err.Error() != tc.expectedError {
					t.Errorf("GetTimestampByEra() expected error: %s, got: %v", tc.expectedError, err)
				}
			}
		})
	}
}

func TestGetByTimestampEra(t *testing.T) {
	client := NewClient()
	preptime := time.Date(2023, 4, 11, 11, 18, 04, 261201025, time.UTC)

	era, err := client.GetEraByTimestamp(rpc_url, preptime)
	if err != nil {
		t.Fatalf("Error getting era by timestamp: %s", err)
	}

	expectedEra := 8738
	if era != expectedEra {
		t.Fatalf("Expected era %d, but got %d", expectedEra, era)
	}

	log.Println("Era:", era, "Timestamp:", preptime)
}

func TestGetByTimestampBlock(t *testing.T) {
	client := NewClient()
	preptime := time.Date(2023, 4, 11, 11, 18, 04, 261201025, time.UTC)

	block, err := client.GetBlockByTimestamp(rpc_url, preptime)
	if err != nil {
		t.Fatalf("Error getting era by timestamp: %s", err)
	}

	expectedBlock := 1635233
	if block != expectedBlock {
		t.Fatalf("Expected era %d, but got %d", expectedBlock, block)
	}

	log.Println("Era:", block, "Timestamp:", preptime)
}

// func TestCollectUserWithdraws(t *testing.T) {
// 	client := NewClient()
// 	rpc := newsdk.NewRPCClient(newsdk.NewRPCHandler(rpc_url, http.DefaultClient))
// 	depolys := []string{"b65ddaac5f5a4f926ff4bfb5f8f82fd660d344ce02434c268ee26c8e11de4586"}
// 	address := "020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292"
// 	withdraw := "withdraw-a67b1aca3084c820fb46de3351f0810153f8ce8d903ce10bb1651db0162b58b1"
// 	withdraws, err := client.collectUserWithdrawsUndelegate(rpc, depolys, address)
// 	if err != nil {
// 		t.Fatalf("Error collecting user withdraws: %s", err)
// 	}
// 	if len(withdraws) != 1 {
// 		t.Fatalf("Expected 1 withdraw, but got %d", len(withdraws))
// 	}
// 	if withdraws[0].withdrawHash != withdraw {
// 		t.Fatalf("Expected withdraw %s, but got %s", withdraw, withdraws[0].withdrawHash)
// 	}
// 	//fmt.Println(withdraws)

// }

func TestProcessUserWithdraws(t *testing.T) {
	client := NewClient()
	rpc := newsdk.NewRPCClient(newsdk.NewRPCHandler(rpc_url, http.DefaultClient))

	withdraw := UncheckedWithdraws{
		withdrawHash: "withdraw-b98aeebe049967c8273088fac9301978ac1edbd8bff07c5256e0957e52a0ccac",
		amount:       "1000000000000000000",
	}
	nowblock, err := rpc.GetBlockLatest(context.Background())
	if err != nil {
		t.Fatalf("Error getting latest block: %s", err)
	}
	_, err = client.processUserWithdrawsUnstaking(rpc, []UncheckedWithdraws{withdraw}, nowblock.Block.Hash.String())
	if err != nil {
		t.Fatalf("Error processing user withdraws: %s", err)
	}

}

func TestGetPriceMainCoin(t *testing.T) {
	client := NewClient()

	price, err := client.GetPriceMainCoin()
	if err != nil {
		t.Errorf("Error should be nil, got: %v", err)
	}
	if price <= 0 {
		t.Errorf("Expected price to be greater than 0, got: %f", price)
	}
}

func TestGetEraBounds(t *testing.T) {
	client := NewClient()
	era := 8500
	firstBlock, lastBlock, err := client.GetEraBounds(rpc_url, era)
	if err != nil {
		t.Errorf("Error should be nil, got: %v", err)
	}
	if firstBlock != 1582782 {
		t.Errorf("Expected first block to be 1582782, got: %d", firstBlock)
	}
	if lastBlock != 1583001 {
		t.Errorf("Expected last block to be 1584101, got: %d", lastBlock)
	}
}

func TestGetHistoryDelegate(t *testing.T) {
	client := NewClient()
	want := blockchain.HistroyStake{
		ValidatorPubkey: "020377bc3ad54b5505971e001044ea822a3f6f307f8dc93fa45a05b7463c0a053bed",
		Amount:          "500000000000",
		Era:             8916,
	}
	history, _, err := client.GetHistoryDelegate(rpc_url, "020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292", 1674242, 1674246)
	if err != nil {
		t.Errorf("Error should be nil, got: %v", err)
	}
	if len(history) != 1 {
		t.Errorf("Expected history to be 1, got: %d", len(history))
	}
	if history[0].ValidatorPubkey != want.ValidatorPubkey {
		t.Errorf("Expected ValidatorPubkey to be %s, got: %s", want.ValidatorPubkey, history[0].ValidatorPubkey)
	}
	if history[0].Amount != want.Amount {
		t.Errorf("Expected Amount to be %s, got: %s", want.Amount, history[0].Amount)
	}
	if history[0].Era != want.Era {
		t.Errorf("Expected Era to be %d, got: %d", want.Era, history[0].Era)
	}

}

// test apy, err := client.GetAPYByERA(rpc_url, "020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292", 8924, 8925)
func TestGetAPYByERA(t *testing.T) {
	client := NewClient()
	apy, err := client.GetAPRByERA(rpc_url, "020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292", 8924, 8925)
	if err != nil {
		t.Errorf("Error should be nil, got: %v", err)
	}
	//apy must be like 21.59027172361426, compare float number is not accurate
	if apy < 21.5 || apy > 21.6 {
		t.Errorf("Expected apy to be 21.59027172361426, got: %f", apy)
	}

}

func TestGetBalanceBeingStaked(t *testing.T) {
	client := NewClient()
	_, _, err := client.GetBalanceBeingStaked(rpc_url, "020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292")
	if err != nil {
		t.Errorf("Error should be nil, got: %v", err)
	}
}

func TestCalculateApy(t *testing.T) {
	client := NewClient()
	apy, err := client.CalculateCurrentChainAPY(rpc_url, "testnet")
	if err != nil {
		t.Errorf("Error should be nil, got: %v", err)
		return
	}
	if apy <= 0 {
		t.Errorf("Expected apy to be greater than 0, got: %f", apy)
	}

}

func TestCheckChain(t *testing.T) {
	rpc_testnet := "http://65.21.238.180:7777/rpc"
	chain := "casper-test"
	res1, err := client.CheckChain(rpc_testnet, chain)
	if err != nil {
		t.Errorf("Error should be nil, got: %v", err)
		return
	}
	if res1 != true {
		t.Errorf("Expected res1 to be true, got: %v", res1)
	}
	rpc_mainnet := "http://65.109.54.159:7777/rpc"
	chain = "casper"
	res2, err := client.CheckChain(rpc_mainnet, chain)
	if err != nil {
		t.Errorf("Error should be nil, got: %v", err)
		return
	}
	if res2 != true {
		t.Errorf("Expected res2 to be true, got: %v", res2)
	}
}

func TestGetBlockAllEvents(t *testing.T) {

	tests := []struct {
		blockNum uint64
		expected blockchain.BlockEvents
	}{
		{
			blockNum: 1808765,
			expected: blockchain.BlockEvents{
				Transfers: []blockchain.TransferResponse{},
				Undelegates: []blockchain.UndelegateData{
					//{525843085970 020377bc3ad54b5505971e001044ea822a3f6f307f8dc93fa45a05b7463c0a053bed 9529 false
					{Address: "011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635",
						Amount:          "525843085970",
						ValidatorPubkey: "020377bc3ad54b5505971e001044ea822a3f6f307f8dc93fa45a05b7463c0a053bed",
						Era:             9529,
						Finished:        false},
				},
				Decimals: 9,
			},
		},
		{
			blockNum: 1808739,
			expected: blockchain.BlockEvents{
				Transfers: []blockchain.TransferResponse{
					{DeployHash: "84726d430635dacc7451af0e352a3d6116382c7882ae0405d4d769919ded9d40",
						From:       "5a558b721112b6ec722dfb9f350ccd7651889ed19edbc24701eb8612f842fcc7",
						FromPubKey: "011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635",
						To:         "6174cf2e6f8fed1715c9a3bace9c50bfe572eecb763b0ed3f644532616452008",
						ToPubKey:   "",
						Amount:     "500000000000",
						Gas:        "0",
					},
				},
				Delegates: []blockchain.DelegateData{
					//{525843085970 020377bc3ad54b5505971e001044ea822a3f6f307f8dc93fa45a05b7463c0a053bed 9529 false
					{Address: "011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635",
						Amount:          "500000000000",
						ValidatorPubkey: "01ace6578907bfe6eba3a618e863bbe7274284c88e405e2857be80dd094726a223",
						Era:             9529,
						Finished:        true},
				},
				Undelegates: []blockchain.UndelegateData{},
				Decimals:    9,
			},
		},
	}

	for _, tt := range tests {
		events, err := client.GetBlockAllEvents(rpc_url, tt.blockNum, true, true, true, true)
		if err != nil {
			log.Println(err)
			return
		}

		if !cmp.Equal(events, tt.expected) {
			t.Errorf("GetBlockAllEvents() diff = %v", cmp.Diff(events, tt.expected))
		}
	}
}
