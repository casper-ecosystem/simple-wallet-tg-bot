package handlers

import (
	"log"

	pb "github.com/Simplewallethq/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/tg-bot/tggateway/types"
	"google.golang.org/protobuf/proto"
	tele "gopkg.in/telebot.v3"
)

func (h *Handler) AddExistingWallet(c tele.Context) error {
	log.Printf("user %d press AddExistingWallet", c.Sender().ID)
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
	msg := pb.AddExistingWalletButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "AddExistingWallet",
		Data: out,
	}
	return nil
}

func (h *Handler) createNewWallet(c tele.Context) error {
	log.Printf("user %d press createNewWallet", c.Sender().ID)
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
	msg := pb.CreateNewWalletButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "CreateNewWallet",
		Data: out,
	}
	return nil
}

func (h *Handler) NotStoreKeyButtonHandler(c tele.Context) error {
	log.Printf("user %d press NotStoreKeyButton", c.Sender().ID)
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
	msg := pb.StoreTheKeyButton{
		User:    user,
		MsgId:   int64(c.Callback().Message.ID),
		IsStore: false,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "StoreTheKeyMenu",
		Data: out,
	}
	return nil
}

func (h *Handler) StoreKeyButtonHandler(c tele.Context) error {
	log.Printf("user %d press StoreKeyButton", c.Sender().ID)
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
	msg := pb.StoreTheKeyButton{
		User:    user,
		MsgId:   int64(c.Callback().Message.ID),
		IsStore: true,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "StoreTheKeyMenu",
		Data: out,
	}
	return nil
}
