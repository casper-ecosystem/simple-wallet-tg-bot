package main

import (
	"context"
	"os"

	"github.com/Simplewallethq/tg-bot/botmain"
	"github.com/Simplewallethq/tg-bot/botmain/crypto"
	"github.com/Simplewallethq/tg-bot/botmain/swap"
	"github.com/Simplewallethq/tg-bot/ent/migrate"
	"github.com/Simplewallethq/tg-bot/tggateway"
	"github.com/Simplewallethq/tg-bot/tggateway/handlers"
	"github.com/Simplewallethq/tg-bot/tggateway/sender"
	"github.com/Simplewallethq/tg-bot/tggateway/types"
	"github.com/sirupsen/logrus"

	//pb "github.com/Simplewallethq/tg-bot/tggateway/proto"
	_ "github.com/lib/pq"
	//"google.golang.org/protobuf/proto"
)

type Config struct {
	TgToken    string
	Resthost   string
	RestToken  string
	RPCnode    string
	Chain      string
	DB         botmain.DBConfig
	PK_SALT    string
	SWAP_TOKEN string
}

func main() {
	logger := logrus.New()
	// dbconf := botmain.DBConfig{
	// 	User:     "postgres",
	// 	Password: "changeme",
	// 	Host:     "127.0.0.1",
	// 	Port:     "5432",
	// 	DBName:   "tgbot",
	// 	SSLMode:  "disable",
	// }
	// conf := Config{
	// 	TgToken: "6230100048:AAHOsK6anbCDVBTJ97gy7Lb20n3s3gN93q8",
	// 	DB:      dbconf,
	// }
	conf := ParseConfig()
	token := conf.TgToken
	bot, err := tggateway.NewTGbot(token, logger)
	if err != nil {
		panic(err)
	}
	MessageMsg := make(chan types.TgMessageMsg)
	ResponseMsg := make(chan types.TgResponseMsg)
	h := handlers.NewHandler(MessageMsg, ResponseMsg, logger)
	go bot.Run(h)

	client, err := botmain.NewDBClient(conf.DB)
	if err != nil {
		logger.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); err != nil {
		logger.Fatalf("failed creating schema resources: %v", err)
	}
	cryptoMod := crypto.NewCrypto(client, conf.PK_SALT, conf.Chain)
	SwapClient := swap.NewSwapClient(conf.SWAP_TOKEN)
	var botmainInst *botmain.BotMain
	if conf.RestToken != "" {
		botmainInst = botmain.NewBotMainWithToken(MessageMsg, ResponseMsg, client, conf.Resthost, conf.RPCnode, conf.RestToken, cryptoMod, logger, SwapClient)
		logrus.Info("CONFIG BOT WITH TOKEN")
	} else {
		logrus.Info("CONFIG BOT WITHOUT TOKEN")
		botmainInst = botmain.NewBotMain(MessageMsg, ResponseMsg, client, conf.Resthost, conf.RPCnode, cryptoMod, logger, SwapClient)

	}
	//botmain := botmain.NewBotMain(MessageMsg, ResponseMsg, client, conf.Resthost, conf.RPCnode, logger)
	go botmainInst.HandleIncome()
	go botmainInst.TaskRecoverer.RecoverOnStartup()
	go botmainInst.Notificator.Start()
	go botmainInst.ValidatorsCrawler.Start()

	senderInst := sender.NewSender(bot.Bot, MessageMsg, ResponseMsg, logger)
	senderInst.ListenIncome()
}

func ParseConfig() Config {
	config := Config{
		os.Getenv("TG_TOKEN"),
		os.Getenv("REST_HOST"),
		os.Getenv("REST_TOKEN"),
		os.Getenv("RPC_NODE"),
		os.Getenv("CHAIN"),
		botmain.DBConfig{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
		os.Getenv("PK_SALT"),
		os.Getenv("SWAP_TOKEN"),
	}
	return config
}
