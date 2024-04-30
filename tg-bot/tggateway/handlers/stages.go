package handlers

import (
	"bytes"
	"log"
	"strings"

	pb "github.com/Simplewallethq/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/tg-bot/tggateway/types"
	"google.golang.org/protobuf/proto"
	tele "gopkg.in/telebot.v3"
)

func (h *Handler) GetUserState(id int64) (string, bool) {
	if state, ok := h.TestUserState[id]; ok {
		if state != "" {
			return state, true
		}
	}

	return "", false
}

// func (h *Handler) TextHandler(c tele.Context) error {
// 	//log.Println("id", c.Sender().ID)
// 	//log.Println(h.TestUserState)
// 	if state, ok := h.GetUserState(c.Sender().ID); ok {
// 		//log.Printf("id %d state %s", c.Sender().ID, state)
// 		path := strings.Split(state, "/")
// 		switch path[0] {
// 		case "Auth":
// 			return h.Auth(c)
// 			// case "auth1":
// 			// 	return h.Auth1(c)
// 			// case "auth2":
// 			// 	return h.Auth2(c)
// 		}
// 	}
// 	err := c.Send(fmt.Sprintf("Unknown command: %s", c.Text()))
// 	return err
// }

func (h *Handler) TextHandler(c tele.Context) error {

	group := false
	if c.Chat().Type == tele.ChatGroup {

		for _, entity := range c.Message().Entities {
			log.Println(entity)
			if entity.Type == tele.EntityMention {
				if strings.Contains(c.Message().Text, c.Bot().Me.Username) {
					log.Println("GOOD")

				}
			}
		}

		group = true

	}

	user := &pb.User{
		Id:       c.Sender().ID,
		Username: c.Sender().Username,
		Group:    group,
		ChatId:   c.Chat().ID,
	}
	msg := pb.TgTextMessage{
		From:  user,
		MsgId: int64(c.Message().ID),
		Text:  c.Text(),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "OnText",
		Data: out,
	}

	//cannot get permissons
	// if c.Chat().Permissions != nil && c.Chat().Permissions.CanDeleteMessages {
	// 	log.Println("can delete user message")
	// 	err = c.Bot().Delete(c.Message())
	// } else {
	// 	log.Println("can not delete user message")
	// }
	_ = c.Bot().Delete(c.Message())
	return err
}

func (h *Handler) DocumentHandler(c tele.Context) error {
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
	log.Printf("user with id %d send document", user.Id)
	log.Println(c.Message().Document)
	reader, err := c.Bot().File(&c.Message().Document.File)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)
	if err != nil {
		return err
	}
	log.Println(buf.Bytes())

	msg := pb.TgTextMessage{
		From:  user,
		MsgId: int64(c.Message().ID),
		Text:  buf.String(),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "OnText",
		Data: out,
	}
	err = c.Bot().Delete(c.Message())
	return err
}
