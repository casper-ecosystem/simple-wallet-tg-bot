package handlers

import (
	"log"

	pb "github.com/Simplewallethq/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/tg-bot/tggateway/types"
	"google.golang.org/protobuf/proto"
	tele "gopkg.in/telebot.v3"
)

func (h *Handler) ShowSettings(c tele.Context) error {
	log.Printf("user %d press settings", c.Sender().ID)
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
	var msgid int64
	if c.Callback() != nil {
		msgid = int64(c.Callback().Message.ID)
	} else {
		msgid = 0
	}
	msg := pb.TgSettingsButton{
		User:  user,
		MsgId: msgid,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "Settings",
		Data: out,
	}
	return nil
}

func (h *Handler) OnOffNotifications(c tele.Context) error {
	log.Printf("user %d press notifications", c.Sender().ID)
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
	msg := pb.TgOnOffNotifications{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "OnOffNotifications",
		Data: out,
	}
	return nil
}

func (h *Handler) ChangeRewardsNotifyTime(c tele.Context) error {
	log.Printf("user %d press notifications", c.Sender().ID)
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
	msg := pb.ChangeRewardsNotifyTime{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "ChangeRewardsNotifyTime",
		Data: out,
	}
	return nil
}

func (h *Handler) ShowNotifySettings(c tele.Context) error {
	log.Printf("user %d press notify settings", c.Sender().ID)
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
	var msgid int64
	if c.Callback() != nil {
		msgid = int64(c.Callback().Message.ID)
	} else {
		msgid = 0
	}
	msg := pb.TgNotifySettingsButton{
		User:  user,
		MsgId: msgid,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "NotifySettings",
		Data: out,
	}
	return nil
}

func (h *Handler) CancelChangeTimeout(c tele.Context) error {

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
	msg := pb.CancelChangeTimeoutButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "Settings",
		Data: out,
	}

	return nil
}

func (h *Handler) ExportPrivateKey(c tele.Context) error {

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
	msg := pb.ExportPrivateKeyButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "ExportPrivateKeyButton",
		Data: out,
	}

	return nil
}

func (h *Handler) PrivacyOptions(c tele.Context) error {
	log.Printf("user %d press privacy settings", c.Sender().ID)
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
	var msgid int64
	if c.Callback() != nil {
		msgid = int64(c.Callback().Message.ID)
	} else {
		msgid = 0
	}
	msg := pb.TgPrivacySettingsButton{
		User:  user,
		MsgId: msgid,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "PrivacySettings",
		Data: out,
	}
	return nil
}

func (h *Handler) ToggleLogging(c tele.Context) error {
	log.Printf("user %d press ToggleLogging", c.Sender().ID)
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
	msg := pb.ToggleLogging{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "ToggleLogging",
		Data: out,
	}
	return nil
}
