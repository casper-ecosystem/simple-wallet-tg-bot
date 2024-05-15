package handlers

import (
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	tele "gopkg.in/telebot.v3"
)

type Handler struct {
	TestUserState map[int64]string
	MessagesMsg   chan types.TgMessageMsg
	ResponseMsg   chan types.TgResponseMsg
	logger        *logrus.Logger
}

func NewHandler(MessageMsg chan types.TgMessageMsg, ResponseMsg chan types.TgResponseMsg, logger *logrus.Logger) *Handler {
	return &Handler{
		TestUserState: make(map[int64]string),
		MessagesMsg:   MessageMsg,
		ResponseMsg:   ResponseMsg,
		logger:        logger,
	}
}

func (h *Handler) Start(c tele.Context) error {
	if c.Data() != "" {
		// handle argument
		if len(c.Data()) >= 3 && c.Data()[:3] == "inv" {
			//handle invoice argument

			user := &pb.User{
				Id:       c.Sender().ID,
				Username: c.Sender().Username,
			}
			msg := pb.PayInvoiceHandler{
				User:  user,
				Short: c.Data()[3:],
			}
			out, err := proto.Marshal(&msg)
			if err != nil {
				return err
			}
			h.MessagesMsg <- types.TgMessageMsg{
				Name: "PayInvoiceHandler",
				Data: out,
			}
			err = c.Bot().Delete(c.Message())
			return err
		}
	}
	group := false
	if c.Chat().Type == tele.ChatGroup {
		group = true
	}
	user := &pb.User{
		Id:       c.Sender().ID,
		Username: c.Sender().Username,
		Group:    group,
		ChatId:   c.Chat().ID,
	}
	msg := pb.TgCommandStart{
		From: user,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "/start",
		Data: out,
	}
	err = c.Bot().Delete(c.Message())
	return err
}
