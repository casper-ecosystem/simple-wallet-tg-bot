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

func (h *Handler) UndelegateButton(c tele.Context) error {
	log.Printf("user %d press UndelegateButton", c.Sender().ID)
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
	msg := pb.NewUndelegateButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "NewUndelegateButton",
		Data: out,
	}
	return nil
}

func (h *Handler) pickUndelegateValidator(c tele.Context) error {
	var id uint64
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 2 {
		idt, err := strconv.Atoi(data[1])
		if err != nil {
			return err
		}
		id = uint64(idt)
	} else {
		return errors.New("address is empty")
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
	msg := pb.PickUndelegateValidator{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Id:    id,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "pickUndelegateValidator",
		Data: out,
	}
	log.Printf("user %d press ask  pickUndelegateValidator", c.Sender().ID)

	return nil
}

func (h *Handler) UndelegateSelectAmount(c tele.Context) error {
	var amount float64
	var err error
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 2 {
		amount, err = strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}
	} else {
		return errors.New("amount is empty")
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
	msg := pb.UndelegatePickAmount{
		User:   user,
		MsgId:  int64(c.Callback().Message.ID),
		Amount: amount,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "pickUndelegateAmount",
		Data: out,
	}
	log.Printf("user %d press ask  pickUndelegateAmount", c.Sender().ID)

	return nil
}

func (h *Handler) UndelegateConfirmButton(c tele.Context) error {
	log.Printf("user %d press UndelegateConfirmButton", c.Sender().ID)
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
	msg := pb.UndelegateConfirmButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "UndelegateConfirmButton",
		Data: out,
	}
	return nil
}
