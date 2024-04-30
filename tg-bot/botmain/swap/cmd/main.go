package main

import (
	"fmt"
	"log"

	"github.com/Simplewallethq/tg-bot/botmain/swap"
)

func main() {
	client := swap.NewSwapClient("71954344-35b9-4f3a-bf3b-811e6162e99e")
	//pairs, err := client.GetCSPRPairs()
	res, err := client.MakeExchange("btc",
		"cspr",
		"16ftSEQ4ctQFDtVZiUBusQUjRrGhM3JYwe",
		"0167d33cfa53a1498797136d6b95ed1d95f2c9d8d56f2053655920c58fc9f2eb84",
		0.001, "168567599160")
	if err != nil {
		panic(err)
	}
	log.Println(res.ID)
	log.Println(res.AddressFrom)
	fmt.Printf("%+v\n", res)

	// res, err := client.GetCSPRPairs()
	// if err != nil {
	// 	panic(err)
	// }
	log.Println(res)

}
