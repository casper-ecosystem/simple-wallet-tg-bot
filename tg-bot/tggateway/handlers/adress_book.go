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

func (h *Handler) AddressBookHandler(c tele.Context) error {
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
	msg := pb.TgAddressButton{
		From:   user,
		MsgId:  int64(c.Callback().Message.ID),
		Offset: 0,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "AddressBook",
		Data: out,
	}
	log.Printf("user %d press address book", c.Sender().ID)

	return nil
}

func (h *Handler) MoveAddressBookHandler(c tele.Context) error {
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
	msg := pb.TgAddressButton{
		From:   user,
		MsgId:  int64(c.Callback().Message.ID),
		Offset: int64(offset),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "AddressBook",
		Data: out,
	}
	log.Printf("user %d press Move Address Book", c.Sender().ID)

	return nil
}

func (h *Handler) CreateEntryAddressBookHandler(c tele.Context) error {
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
	msg := pb.CreateEntryAddressBookButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "CreateEntryAddressBook",
		Data: out,
	}
	log.Printf("user %d press balance", c.Sender().ID)

	return nil
}

func (h *Handler) ShowAddressHandler(c tele.Context) error {
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
	msg := pb.AskAddressBookDetailed{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Id:    id,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "AskAddressBookDetailed",
		Data: out,
	}
	log.Printf("user %d press ask  address book detailed", c.Sender().ID)

	return nil
}

func (h *Handler) ChangeNameAddressBookHandler(c tele.Context) error {
	var id int
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 2 {
		var err error
		id, err = strconv.Atoi(data[1])
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
	msg := pb.ChangeNameAddressBook{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Id:    uint64(id),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "ChangeNameAddressBook",
		Data: out,
	}
	log.Printf("user %d press ChangeNameAddressBook", c.Sender().ID)

	return nil
}

func (h *Handler) ChangeAddressAddressBookHandler(c tele.Context) error {
	var id int
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 2 {
		var err error
		id, err = strconv.Atoi(data[1])
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
	msg := pb.ChangeAddressAddressBook{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Id:    uint64(id),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "ChangeAddressAddressBook",
		Data: out,
	}
	log.Printf("user %d press ChangeAddressAddressBook", c.Sender().ID)

	return nil
}

func (h *Handler) DeleteAddressBookHandler(c tele.Context) error {
	var id int
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 2 {
		var err error
		id, err = strconv.Atoi(data[1])
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
	msg := pb.DeleteEntryAddressBook{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Id:    uint64(id),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "DeleteEntryAddressBook",
		Data: out,
	}
	log.Printf("user %d press DeleteEntryAddressBook", c.Sender().ID)

	return nil
}

func (h *Handler) ConfirmDeleteAdressBook(c tele.Context) error {

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
	msg := pb.DeleteEntryAddressBookConfirm{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "DeleteEntryAddressBookConfirm",
		Data: out,
	}
	log.Printf("user %d press DeleteEntryAddressBook confirm", c.Sender().ID)

	return nil
}

func (h *Handler) CancelAddressBook(c tele.Context) error {

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
	msg := pb.CancelAddressBook{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "CancelAddressBook",
		Data: out,
	}

	return nil
}
