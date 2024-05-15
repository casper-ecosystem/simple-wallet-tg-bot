package botmain

import (
	"context"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	tggateway "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func (b *BotMain) HandleStart(msg tggateway.TgMessageMsg) error {
	out := pb.TgCommandStart{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}

	_, err := b.DB.User.Query().Where(user.ID(out.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		var entErr *ent.NotFoundError
		if errors.As(err, &entErr) {
			err := b.HandleStartAuth(msg)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	//log.Println(u.PublicKey)
	authed, err := b.CheckLoginWithEditMsg(out.GetFrom(), out.GetMsgId())
	if err != nil {
		return err
	}
	if !authed {
		return nil
	}

	err = b.State.DeleteUserState(out.GetFrom().Id)
	if err != nil {
		b.logger.Error("failed to delete user state")
	}
	data := pb.WelcomeResponse{
		User:  out.GetFrom(),
		MsgId: out.GetMsgId(),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "WelcomeMsg",
		Data: dataBytes,
	}
	return nil
}

// func (b *BotMain) HandleStartAuth(msg tggateway.TgMessageMsg) error {
// 	out := pb.TgCommandStart{}
// 	if err := proto.Unmarshal(msg.Data, &out); err != nil {
// 		return errors.Wrap(err, "failed unmarshal")
// 	}
// 	_, err := b.DB.User.Create().SetID(out.GetFrom().GetId()).Save(context.Background())
// 	if err != nil {
// 		return errors.Wrap(err, "failed create user")
// 	}
// 	err = b.State.SetUserState(out.GetFrom().GetId(), []string{"Auth", "1"})
// 	if err != nil {
// 		return err
// 	}
// 	data := pb.AuthResponse{
// 		User:  out.GetFrom(),
// 		Stage: 1,
// 	}
// 	dataBytes, err := proto.Marshal(&data)
// 	if err != nil {
// 		return errors.Wrap(err, "failed marshal")
// 	}
// 	b.ResponseChan <- tggateway.TgResponseMsg{
// 		Name: "AuthMsg",
// 		Data: dataBytes,
// 	}

// 	return nil
// }

func (b *BotMain) HandleStartAuth(msg tggateway.TgMessageMsg) error {
	out := pb.TgCommandStart{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	_, err := b.DB.User.Create().SetID(out.GetFrom().GetId()).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed create user")
	}
	err = b.State.SetUserState(out.GetFrom().GetId(), []string{"Auth", "AskStoreKey"})
	if err != nil {
		return err
	}
	data := pb.AuthRegisterType{
		User:  out.GetFrom(),
		MsgId: out.GetFrom().Id,
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AuthRegisterMsg",
		Data: dataBytes,
	}

	return nil
}
