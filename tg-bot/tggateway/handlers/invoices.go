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

func (h *Handler) InvoicesHandler(c tele.Context) error {
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
	msg := pb.TgInvoicesButton{
		From:   user,
		MsgId:  int64(c.Callback().Message.ID),
		Offset: 0,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "Invoices",
		Data: out,
	}
	log.Printf("user %d press invoices", c.Sender().ID)

	return nil
}

func (h *Handler) MoveInvoicesHandler(c tele.Context) error {
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
	msg := pb.TgInvoicesButton{
		From:   user,
		MsgId:  int64(c.Callback().Message.ID),
		Offset: int64(offset),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "Invoices",
		Data: out,
	}
	log.Printf("user %d press Move Invoices", c.Sender().ID)

	return nil
}

func (h *Handler) NewInvoiceHandler(c tele.Context) error {
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
	msg := pb.TgNewInvoiceButton{
		From:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "NewInvoiceButton",
		Data: out,
	}
	log.Printf("user %d press NewInvoiceButton", c.Sender().ID)

	return nil
}

func (h *Handler) ShowInvoiceHandler(c tele.Context) error {
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
	msg := pb.AskInvoiceDetailed{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Id:    id,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "AskInvoiceDetailed",
		Data: out,
	}
	log.Printf("user %d press ask invoice detailed", c.Sender().ID)

	return nil
}

func (h *Handler) DeleteInvoice(c tele.Context) error {
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
	msg := pb.DeleteInvoice{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Id:    uint64(id),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "DeleteInvoice",
		Data: out,
	}
	log.Printf("user %d press DeleteInvoice", c.Sender().ID)

	return nil
}

func (h *Handler) DeleteInvoiceConfirm(c tele.Context) error {
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
	msg := pb.DeleteInvoiceConfirm{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Id:    uint64(id),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "DeleteInvoiceConfirm",
		Data: out,
	}
	log.Printf("user %d press DeleteInvoice", c.Sender().ID)

	return nil
}

func (h *Handler) PayInvoiceTransfer(c tele.Context) error {
	var short string
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 2 {
		short = data[1]
	} else {
		return errors.New("err parse button data")
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
	msg := pb.PayInvoiceTransfer{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Short: short,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "PayInvoiceTransfer",
		Data: out,
	}
	log.Printf("user %d press PayInvoiceTransfer", c.Sender().ID)

	return nil
}

func (h *Handler) PayInvoiceSwap(c tele.Context) error {
	var short string
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 2 {
		short = data[1]
	} else {
		return errors.New("err parse button data")
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
	msg := pb.PayInvoiceSwap{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Short: short,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "PayInvoiceSwap",
		Data: out,
	}
	log.Printf("user %d press PayInvoiceSwap", c.Sender().ID)

	return nil
}

func (h *Handler) ShowInvoicePayments(c tele.Context) error {
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
	msg := pb.ShowInvoicePayments{
		User:   user,
		MsgId:  int64(c.Callback().Message.ID),
		Id:     uint64(id),
		Offset: 0,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "ShowInvoicePayments",
		Data: out,
	}
	log.Printf("user %d press ShowInvoicePayments", c.Sender().ID)

	return nil
}

func (h *Handler) MovePaymentsHandler(c tele.Context) error {
	var offset int
	var invoice int
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 3 {
		var err error
		offset, err = strconv.Atoi(data[2])
		if err != nil {
			return err
		}
		invoice, err = strconv.Atoi(data[1])
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
	msg := pb.ShowInvoicePayments{
		User:   user,
		MsgId:  int64(c.Callback().Message.ID),
		Offset: int64(offset),
		Id:     uint64(invoice),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "ShowInvoicePayments",
		Data: out,
	}
	log.Printf("user %d press Move Invoices", c.Sender().ID)

	return nil
}

func (h *Handler) ShowRecentInvoices(c tele.Context) error {

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
	msg := pb.ShowRecentInvoices{
		User:   user,
		MsgId:  int64(c.Callback().Message.ID),
		Offset: 0,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "ShowRecentInvoices",
		Data: out,
	}
	log.Printf("user %d press ShowRecentInvoices", c.Sender().ID)

	return nil
}

func (h *Handler) PayInvoice(c tele.Context) error {

	var short string
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 2 {
		short = data[1]
	} else {
		return errors.New("err parse button data")
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
	msg := pb.PayInvoiceHandler{
		User:  user,
		Short: short,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "PayInvoiceHandler",
		Data: out,
	}
	log.Printf("user %d press PayInvoiceHandler ", c.Sender().ID)

	return nil
}

func (h *Handler) MoveRecentInvoicesList(c tele.Context) error {
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
	msg := pb.ShowRecentInvoices{
		User:   user,
		MsgId:  int64(c.Callback().Message.ID),
		Offset: int64(offset),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "ShowInvoicePayments",
		Data: out,
	}
	log.Printf("user %d press Move Invoices", c.Sender().ID)

	return nil
}

func (h *Handler) ExportPaymentsInvoice(c tele.Context) error {
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
	msg := pb.ExportPaymentsInvoice{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Id:    int64(id),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "ExportPaymentsInvoice",
		Data: out,
	}
	log.Printf("user %d press exportPaymentsInvoice", c.Sender().ID)

	return nil
}
