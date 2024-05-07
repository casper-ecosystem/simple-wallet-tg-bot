package main

import (
	"log"
	"time"

	"github.com/Simplewallethq/source-code/library/blockchain"
	csprclient "github.com/Simplewallethq/source-code/library/client"
)

var rpc_url string = "http://65.21.238.180:7777/rpc"

var client blockchain.Client = csprclient.NewClient()

func GetLatestBlock() {
	h, e, err := client.GetCurrentState(rpc_url)

	if err != nil {
		log.Println("can't get latest block")
	}
	log.Printf("height: %d, era: %d", h, e)
}

func GetBalanceBeingUnstaked() {
	res, _, err := client.GetBalanceBeingUnstaked(rpc_url, "01c4105e4152fb185b8a1fa041b022ca56eb6c8ba556dbf41030899ce9e250a6a9")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res)
}

func GetTimestampByBlock() {
	t, estimated, err := client.GetTimestampByBlock(rpc_url, 1653302)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(t, estimated)
}

func GetTimestampByEra() {
	t, estimated, err := client.GetTimestampByEra(rpc_url, 3700)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(t, estimated)
}

func GetByTimestampEra() {
	// time now - 7 days
	preptime := time.Date(2023, 4, 11, 11, 18, 04, 261201025, time.UTC)
	era, err := client.GetEraByTimestamp(rpc_url, preptime)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(era, preptime)
}

func GetByTimestampBlock() {
	// time now - 7 days
	preptime := time.Date(2023, 4, 11, 15, 18, 04, 261201025, time.UTC)
	block, err := client.GetBlockByTimestamp(rpc_url, preptime)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(block, preptime)
}

func GetRewardsByBlock() {
	res, _, err := client.GetRewardsByBlock(rpc_url, "011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635", 1639816, 1639836)
	if err != nil {
		log.Println(err)
		return
	}
	for _, rew := range res {
		log.Println(rew.Amount.String())
		log.Println(rew.ValidatorPubkey)
		log.Println(rew.Block)
	}
}

func CheckAddress() {
	is := client.IsAddress("020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292")

	log.Println(is)
}

func CheckChain() {
	is, err := client.CheckChain(rpc_url, "casper-test")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(is)
}

func GetAccountBalance() {
	balance, _, err := client.GetBalanceBase(rpc_url, "02038f2267e1f40294fc1e681f83732cb8bbe6bfbfbb887f8dcd84a76154f35a453d")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(balance.String())
}

func GetHistoryTransfers() {
	history, _, err := client.GetHistoryTransfers(rpc_url, "011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635", 2005519, 2006019)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(history)
}

func GetAccountStaked() {
	balance, _, err := client.GetBalanceStaked(rpc_url, "02038f2267e1f40294fc1e681f83732cb8bbe6bfbfbb887f8dcd84a76154f35a453d")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(balance)
}

func GetHistoryUndelegate() {
	history, _, err := client.GetHistoryUndelegate(rpc_url, "02038f2267e1f40294fc1e681f83732cb8bbe6bfbfbb887f8dcd84a76154f35a453d", 1840245, 1840276)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(history)
}

func GetPriceMainCoin() {
	price, err := client.GetPriceMainCoin()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(price)
}

func GetEraBounds() {
	b1, b2, err := client.GetEraBounds(rpc_url, 9459)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(b1, b2)
}

func GetRewardsByEra() {
	res, _, err := client.GetRewardsByEra(rpc_url, "020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292", 8902, 8904)
	if err != nil {
		log.Println(err)
		return
	}
	for _, rew := range res {
		log.Println(rew.Amount.String())
		log.Println(rew.ValidatorPubkey)
		log.Println(rew.Era)
	}
}

func GetHistoryDelegate() {
	history, _, err := client.GetHistoryDelegate(rpc_url, "020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292", 1674240, 1674249)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(history)
}

func GetAPYByERA() {
	apy, err := client.GetAPRByERA(rpc_url, "020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292", 8924, 8925)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(apy)
}
func GetBalanceBeingStaked() {
	staked, _, err := client.GetBalanceBeingStaked(rpc_url, "011b87a676e4ac0336f54cb40141a97600464cddab056e2664d5f76d77dbd94635")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(staked)
}

func CalculateCurrentChainAPY() {
	apy, err := client.CalculateCurrentChainAPY(rpc_url, "testnet")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(apy)

	apym, err := client.CalculateCurrentChainAPY("http://13.58.157.225:7777/rpc", "mainnet")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(apym)
}

func GetBlockEvents() {
	events, err := client.GetBlockAllEvents(rpc_url, 1890454, true, true, true, true)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(events)
	// events, err = client.GetBlockAllEvents(rpc_url, 1808739, true, true, true, true)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Println(events)
	// events, err = client.GetBlockAllEvents(rpc_url, 1827451, true, true, true, true)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Println(events)
}

func GetValidators() {
	validators, err := client.GetValidators(rpc_url)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(validators)
}

func main() {
	//GetLatestBlock() //work
	//GetAccountStaked() //work
	//GetRewardsByBlock() //work
	//GetTimestampByBlock() //work
	//GetTimestampByEra() //work
	//GetByTimestampEra()   //work
	//GetByTimestampBlock() //work
	//GetHistoryTransfers() //work
	//GetBalanceBeingUnstaked() //work
	//GetHistoryUndelegate() //work
	//GetHistoryDelegate() //work
	//GetAPYByERA() //work
	//GetBalanceBeingStaked() //work need test
	//CalculateCurrentChainAPY()
	GetPriceMainCoin() //work
	//GetAccountBalance() //work
	//GetEraBounds() //work
	//GetRewardsByEra() //work
	//GetBlockEvents()
	//CheckChain()
}
