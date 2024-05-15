package tggateway

import (
	"strconv"
	"time"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/handlers"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/messages"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

type TGbot struct {
	logger *logrus.Logger
	Bot    *tele.Bot
}

func NewTGbot(token string, logger *logrus.Logger) (*TGbot, error) {

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, ctx tele.Context) {
			logger.Error(err)
		},
	}
	b, err := tele.NewBot(pref)

	if err != nil {
		return nil, err
	}
	return &TGbot{
		Bot:    b,
		logger: logger,
	}, nil
}

func (b *TGbot) Run(h *handlers.Handler) {
	//b.Bot.Group().Handle()
	b.Bot.Handle("/start", h.Start, b.BlockNotPrivate)
	b.Bot.Handle("/lock", h.LockHandler, b.BlockNotPrivate)
	b.Bot.Handle("/settings", h.ShowSettings, b.BlockNotPrivate)
	b.Bot.Handle(tele.OnCallback, h.ButtonHandler)
	b.Bot.Handle(tele.OnText, h.TextHandler)
	b.Bot.Handle(tele.OnDocument, h.DocumentHandler)
	//b.Bot.Handle(tele.OnQuery, b.QueryHandler)
	b.Bot.Start()
}

func (b *TGbot) QueryHandler(c tele.Context) error {
	urls := []string{
		"http://photo.jpg",
		"http://photo2.jpg",
	}

	results := make(tele.Results, len(urls)) // []tele.Result
	for i, url := range urls {
		result := &tele.PhotoResult{
			URL:      url,
			ThumbURL: url, // required for photos
		}

		results[i] = result
		// needed to set a unique string ID for each result
		results[i].SetResultID(strconv.Itoa(i))
	}

	return c.Answer(&tele.QueryResponse{
		Results:   results,
		CacheTime: 60, // a minute
	})
}

type recipient struct {
	Id int64
}

func (r *recipient) Recipient() string {
	return strconv.FormatInt(r.Id, 10)
}

func (b *TGbot) BlockNotPrivate(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		if c.Chat().Type == tele.ChatGroup {
			what, opts, err := messages.BlockForGroup("eng")
			if err != nil {
				what, opts, _ = messages.GetErrorMessage()
			}

			recipient := &recipient{Id: c.Chat().ID}
			_, err = c.Bot().Send(recipient, what, opts...)
			if err != nil {
				b.logger.Error("error send message: ", err, "uid: ", c.Chat().ID)

			}
		}
		return next(c) // continue execution chain
	}
}
