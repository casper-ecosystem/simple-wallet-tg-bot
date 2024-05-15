package handlers

import (
	"log"
	"strconv"
	"strings"

	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	tele "gopkg.in/telebot.v3"
)

func (h *Handler) DelegateButton(c tele.Context) error {
	log.Printf("user %d press DelegateButton", c.Sender().ID)
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
	msg := pb.NewDelegateButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "NewDelegateButton",
		Data: out,
	}
	return nil
}

func (h *Handler) PickDelegateValidator(c tele.Context) error {
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
	msg := pb.PickDelegateValidator{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Id:    id,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "pickDelegateValidator",
		Data: out,
	}
	log.Printf("user %d press ask  pickDelegateValidator", c.Sender().ID)

	return nil
}

func (h *Handler) PickDelegateAmount(c tele.Context) error {
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
	msg := pb.DelegatePickAmount{
		User:   user,
		MsgId:  int64(c.Callback().Message.ID),
		Amount: amount,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "pickDelegateAmount",
		Data: out,
	}
	log.Printf("user %d press ask  pickDelegateAmount", c.Sender().ID)

	return nil
}

func (h *Handler) DelegateConfirmButton(c tele.Context) error {
	log.Printf("user %d press DelegateConfirmButton", c.Sender().ID)
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
	msg := pb.DelegateConfirmButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "DelegateConfirmButton",
		Data: out,
	}
	return nil
}

func (h *Handler) MoveDelegateValidators(c tele.Context) error {
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
	msg := pb.NewDelegateButton{
		User:   user,
		MsgId:  int64(c.Callback().Message.ID),
		Offset: int64(offset),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "NewDelegateButton",
		Data: out,
	}
	log.Printf("user %d press Move Address Book", c.Sender().ID)

	return nil
}

func (h *Handler) NewDepositButton(c tele.Context) error {
	log.Printf("user %d press NewDeposit by swap", c.Sender().ID)
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
	msg := pb.NewDepositButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "NewDepositButton",
		Data: out,
	}
	return nil
}
