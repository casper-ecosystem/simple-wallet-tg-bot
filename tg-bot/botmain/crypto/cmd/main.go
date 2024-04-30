package main

import (
	"encoding/json"
	"log"

	"github.com/Simplewallethq/tg-bot/botmain/crypto"
	"github.com/Simplewallethq/tg-bot/botmain/restclient"
	"github.com/Simplewallethq/tg-bot/ent"
	_ "github.com/lib/pq"
)

func main() {
	DB, err := ent.Open("postgres", "host="+"127.0.0.1"+" port="+"5432"+" user="+"postgres"+" dbname="+"tgbot"+" password="+"changeme"+" sslmode="+"disable")
	if err != nil {
		panic(err)
	}
	c := crypto.NewCrypto(DB, "TEST SALT", "casper-test")
	client := restclient.NewClient("http://65.108.2.174/rest/api/v1/cspr-testnet")
	tr := crypto.Transfer{
		ToPubkey: "020237037ff4845669e59d3e7698e7d58eb97ca378960ac57478a86a6a3535460292",
		Amount:   "2500000000",
		Memo:     123,
	}
	dep, err := c.SignTransferWithPassword(tr, 158287363, "123")
	if err != nil {
		panic(err)
	}
	log.Println(dep)

	// marshal dep to json
	jsres, err := json.Marshal(dep)
	if err != nil {
		panic(err)
	}
	log.Println(string(jsres))

	res, err := client.PutDeploy("65.21.238.180", *dep)
	if err != nil {
		panic(err)
	}
	log.Println(res)

	// keypair, content, err := crypto.GenerateEd25519Pair()
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(string(content))
	// println(keypair.PublicKey().String())

}
