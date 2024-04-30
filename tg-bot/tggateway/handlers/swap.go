package handlers

import (
	"errors"
	"log"
	"strconv"
	"strings"

	pb "github.com/Simplewallethq/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/tg-bot/tggateway/types"
	"google.golang.org/protobuf/proto"
	tele "gopkg.in/telebot.v3"
)

func (h *Handler) SwapBySwapButton(c tele.Context) error {
	log.Printf("user %d press Deposit by swap button", c.Sender().ID)
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
	msg := pb.SwapBySwapButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "SwapBySwapButton",
		Data: out,
	}
	return nil
}

func (h *Handler) MoveSwapPairs(c tele.Context) error {
	var offset int
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 2 {
		var err error
		offset, err = strconv.Atoi(data[1])
		if err != nil {
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
	msg := pb.AskSwapPairs{
		User:   user,
		MsgId:  int64(c.Callback().Message.ID),
		Limit:  5,
		Offset: int64(offset),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "AskSwapPairs",
		Data: out,
	}
	log.Printf("user %d press Move Address Book", c.Sender().ID)

	return nil
}

func (h *Handler) pickSwapPair(c tele.Context) error {
	var cur string
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 2 {
		cur = data[1]
	} else {
		return errors.New("currency is empty")
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
	msg := pb.PickSwapPair{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Cur:   cur,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "pickSwapPair",
		Data: out,
	}
	log.Printf("user %d press ask  pick Swap Pair", c.Sender().ID)

	return nil
}

func (h *Handler) SwapConfirmAmount(c tele.Context) error {
	log.Printf("user %d press Swap confirm amount", c.Sender().ID)
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
	msg := pb.SwapConfirmAmount{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "SwapConfirmAmount",
		Data: out,
	}
	return nil
}

func (h *Handler) pickSwapChain(c tele.Context) error {
	var cur string
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 2 {
		cur = data[1]
	} else {
		cur = ""
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
	msg := pb.PickSwapChain{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Chain: cur,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "pickSwapChain",
		Data: out,
	}
	log.Printf("user %d press ask  pick Swap Chain", c.Sender().ID)

	return nil
}

func (h *Handler) WithdrawBySwapButton(c tele.Context) error {
	log.Printf("user %d press Withdraw by swap button", c.Sender().ID)
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
	msg := pb.SwapBySwapButton{
		User:     user,
		MsgId:    int64(c.Callback().Message.ID),
		Withdraw: true,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "SwapBySwapButton",
		Data: out,
	}
	return nil
}
