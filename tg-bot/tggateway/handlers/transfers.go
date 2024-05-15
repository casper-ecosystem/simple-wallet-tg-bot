package handlers

import (
	"errors"
	"log"
	"strconv"
	"strings"

	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	"google.golang.org/protobuf/proto"
	tele "gopkg.in/telebot.v3"
)

func (h *Handler) newTransfer(c tele.Context) error {
	log.Printf("user %d press New transfer", c.Sender().ID)
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
	msg := pb.NewTransferButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "NewTransferButton",
		Data: out,
	}
	return nil
}

func (h *Handler) TransferCustomAddress(c tele.Context) error {
	log.Printf("user %d press TransferCustomAddress", c.Sender().ID)
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
	msg := pb.TransferCustomAddressButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "TransferCustomAddress",
		Data: out,
	}
	return nil
}

func (h *Handler) TransferConfirmButton(c tele.Context) error {
	log.Printf("user %d press TransferConfirmButton", c.Sender().ID)
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
	msg := pb.TransferConfirmButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "TransferConfirmButton",
		Data: out,
	}
	return nil
}

func (h *Handler) TransferMaximum(c tele.Context) error {
	log.Printf("user %d press TransferMaximum button", c.Sender().ID)
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
	msg := pb.TransferMaximumButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "TransferMaximumButton",
		Data: out,
	}
	return nil
}

func (h *Handler) TransferAddressBook(c tele.Context) error {
	log.Printf("user %d press TransferAddressBook button", c.Sender().ID)
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
	msg := pb.TransferAddressBookButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "TransferAddressBookButton",
		Data: out,
	}
	return nil
}

func (h *Handler) PickTransferAddress(c tele.Context) error {
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
	msg := pb.PickTransferAddress{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Id:    id,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "PickTransferAddress",
		Data: out,
	}
	log.Printf("user %d press ask  PickTransferAddress", c.Sender().ID)

	return nil
}

func (h *Handler) MoveTransferAddressBookHandler(c tele.Context) error {
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
	msg := pb.TransferAddressBookButton{
		User:   user,
		MsgId:  int64(c.Callback().Message.ID),
		Offset: int64(offset),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "TransferAddressBookButton",
		Data: out,
	}
	log.Printf("user %d press Move Address Book", c.Sender().ID)

	return nil
}

func (h *Handler) TransferWithoutMemo(c tele.Context) error {
	log.Printf("user %d press TransferWithoutMemo button", c.Sender().ID)
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
	msg := pb.TransferWithoutMemo{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "TransferWithoutMemo",
		Data: out,
	}
	return nil
}
