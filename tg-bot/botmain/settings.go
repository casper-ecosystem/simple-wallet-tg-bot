package botmain

import (
	"context"
	"log"
	"strconv"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	tggateway "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func (b *BotMain) HandleExportPKState(msg *pb.TgTextMessage) error {
	log.Println("received exportpk state message", msg)
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 || state[0] != "ExportPK" {
		err = b.State.DeleteUserState(msg.GetFrom().GetId())
		if err != nil {
			return errors.Wrap(err, "failed delete user state")
		}
		return errors.Wrap(err, "bad get user state")
	}

	switch state[1] {
	case "AskPass":
		log.Println("received askPass state message", msg)
		err := b.ExportPKAskPass(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle transfer amount")
		}

	}
	return nil
}

func (b *BotMain) HandleSettings(msg tggateway.TgMessageMsg) error {
	out := pb.TgSettingsButton{}
	err := proto.Unmarshal(msg.Data, &out)
	if err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	loggedin, err := b.CheckLogin(out.GetUser())
	if err != nil {
		return err
	}
	if !loggedin {
		return nil
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}

	lockTimeout := u.LockTimeout
	if lockTimeout == 0 {
		lockTimeout = DEFAULT_LOGOUT_TIMEOUT
	}
	data := pb.SettingsResponse{
		User:          out.GetUser(),
		PublicKey:     u.PublicKey,
		Notifications: u.Notify,
		NotifyTime:    int32(u.NotifyTime),
		LockTimeout:   lockTimeout,
		MsgId:         out.GetMsgId(),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "SettingsMsg",
		Data: dataBytes,
	}
	return nil

}

func (b *BotMain) HandleLogoutState(msg *pb.TgTextMessage) error {
	// out := pb.TgLogoutButton{}
	// err := proto.Unmarshal(msg.Data, &out)
	// if err != nil {
	// 	return errors.Wrap(err, "failed unmarshal")
	// }
	u, err := b.DB.User.Query().Where(user.ID(msg.From.GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	if msg.Text == "CONFIRM" {
		err = b.DB.User.DeleteOne(u).Exec(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed update user")
		}
		data := pb.LogoutResponse{
			User: msg.GetFrom(),
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "LogoutMsg",
			Data: dataBytes,
		}
	} else {
		err = b.State.DeleteUserState(msg.GetFrom().GetId())
		if err != nil {
			return errors.Wrap(err, "failed update user state")
		}
		data := pb.RejectLogout{
			User: msg.GetFrom(),
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "RejectLogout",
			Data: dataBytes,
		}
		startMsgdata := pb.TgCommandStart{
			From:  msg.GetFrom(),
			MsgId: msg.GetMsgId() - 1,
		}
		dataBytes, err = proto.Marshal(&startMsgdata)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		startmsg := tggateway.TgMessageMsg{
			Name: "/start",
			Data: dataBytes,
		}
		err = b.HandleStart(startmsg)
		if err != nil {
			return err
		}

	}
	return nil

}

func (b *BotMain) CancelLogout(msg tggateway.TgMessageMsg) error {
	out := pb.TgLogoutButton{}
	err := proto.Unmarshal(msg.Data, &out)
	if err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	// u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	// if err != nil {
	// 	return errors.Wrap(err, "failed get user")
	// }

	err = b.State.DeleteUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed update user state")
	}
	data := pb.RejectLogout{
		User: out.GetUser(),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "RejectLogout",
		Data: dataBytes,
	}
	startMsgdata := pb.TgCommandStart{
		From:  out.GetUser(),
		MsgId: out.GetMsgId(),
	}
	dataBytes, err = proto.Marshal(&startMsgdata)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	startmsg := tggateway.TgMessageMsg{
		Name: "/start",
		Data: dataBytes,
	}
	err = b.HandleStart(startmsg)
	if err != nil {
		return err
	}

	return nil

}

func (b *BotMain) CancelChangeTimeout(msg tggateway.TgMessageMsg) error {
	out := pb.CancelChangeTimeoutButton{}
	err := proto.Unmarshal(msg.Data, &out)
	if err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	// u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	// if err != nil {
	// 	return errors.Wrap(err, "failed get user")
	// }

	err = b.State.DeleteUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed update user state")
	}

	Settings := pb.TgSettingsButton{
		User:  out.GetUser(),
		MsgId: out.GetMsgId(),
	}
	dataBytes, err := proto.Marshal(&Settings)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	startmsg := tggateway.TgMessageMsg{
		Name: "/start",
		Data: dataBytes,
	}
	err = b.HandleStart(startmsg)
	if err != nil {
		return err
	}

	return nil

}

func (b *BotMain) HandleLogout(msg tggateway.TgMessageMsg) error {
	out := pb.TgLogoutButton{}
	err := proto.Unmarshal(msg.Data, &out)
	if err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	err = b.State.SetUserState(out.GetUser().GetId(), []string{"logout"})
	if err != nil {
		return errors.Wrap(err, "failed update user state")
	}
	data := pb.LogoutConfirmation{
		User:  out.GetUser(),
		MsgId: out.GetMsgId(),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "LogoutConfirmationMsg",
		Data: dataBytes,
	}
	return nil

}

func (b *BotMain) ChangeLockTimeoutButtonHandler(msg tggateway.TgMessageMsg) error {
	out := pb.TgChangeLockTimeoutButton{}
	err := proto.Unmarshal(msg.Data, &out)
	if err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}

	//_, err = u.Update().SetState("ChangeLockTime/" + strconv.Itoa(int(out.GetMsgId()))).Save(context.Background())
	err = b.State.SetUserState(out.GetUser().GetId(), []string{"ChangeLockTime", strconv.Itoa(int(out.GetMsgId()))})
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	currentTime := u.LockTimeout
	if currentTime == 0 {
		currentTime = DEFAULT_LOGOUT_TIMEOUT
	}
	data := pb.ChangeLogoutAskTime{
		User:        out.GetUser(),
		MsgId:       out.GetMsgId(),
		CurrentTime: currentTime,
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "ChangeLogoutAskTime",
		Data: dataBytes,
	}
	return nil
}

func (b *BotMain) ChangeLockTimeout(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	timeout, err := strconv.Atoi(msg.GetText())
	if err != nil {
		return err
	}
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get state")
	}
	if len(state) != 2 {
		return errors.Wrap(err, "failed get state")
	}
	msgid, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get state")
	}
	err = b.State.DeleteUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get state")
	}
	_, err = u.Update().SetLockTimeout(int64(timeout)).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	// current_time := u.LockTimeout
	// if current_time == 0 {
	// 	current_time = DEFAULT_LOGOUT_TIMEOUT
	// }
	data := pb.ChangeLogoutAskTimeResponse{
		User:        msg.GetFrom(),
		MsgId:       int64(msgid),
		CurrentTume: int64(timeout),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "ChangeLogoutAskTimeResponse",
		Data: dataBytes,
	}
	return nil

}

func (b *BotMain) SendYieldLoadingResponse(user *pb.User, msgid int64, msg string) error {
	loadresp := pb.YieldLoadingResponse{
		User:  user,
		MsgId: msgid,
	}
	loadrespBytes, err := proto.Marshal(&loadresp)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "YieldLoadResponse",
		Data: loadrespBytes,
	}
	return nil
}

func (b *BotMain) HandleOnOffNotifications(msg tggateway.TgMessageMsg) error {
	out := pb.TgOnOffNotifications{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get userInst")
	}
	err = u.Update().SetNotify(!u.Notify).Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update userInst")
	}
	userInst := out.GetUser()
	msgSett := pb.TgNotifySettingsButton{
		User:  userInst,
		MsgId: out.GetMsgId(),
	}
	outSett, err := proto.Marshal(&msgSett)
	if err != nil {
		return err
	}
	msgtosett := types.TgMessageMsg{
		Name: "NotifySettings",
		Data: outSett,
	}
	err = b.HandleNotifySettings(msgtosett)
	return err

}

func (b *BotMain) HandleChangeRewardsNotifyTime(msg tggateway.TgMessageMsg) error {
	out := pb.ChangeRewardsNotifyTime{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get userInst")
	}
	var newTime int8
	switch {
	case u.NotifyTime >= 0 && u.NotifyTime <= 1:
		newTime = 2
	case u.NotifyTime >= 2 && u.NotifyTime <= 3:
		newTime = 4
	case u.NotifyTime >= 4 && u.NotifyTime <= 7:
		newTime = 8
	case u.NotifyTime >= 8 && u.NotifyTime <= 15:
		newTime = 16
	case u.NotifyTime >= 16 && u.NotifyTime <= 23:
		newTime = 24
	default:
		newTime = 0
	}

	log.Println("new_time", newTime)

	err = u.Update().SetNotifyTime(newTime).Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update userInst")
	}
	userInst := out.GetUser()
	msgSett := pb.TgNotifySettingsButton{
		User:  userInst,
		MsgId: out.GetMsgId(),
	}
	outSett, err := proto.Marshal(&msgSett)
	if err != nil {
		return err
	}
	msgtosett := types.TgMessageMsg{
		Name: "NotifySettings",
		Data: outSett,
	}
	err = b.HandleNotifySettings(msgtosett)
	return err

}

func (b *BotMain) HandleNotifySettings(msg tggateway.TgMessageMsg) error {
	out := pb.TgSettingsButton{}
	err := proto.Unmarshal(msg.Data, &out)
	if err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	loggedin, err := b.CheckLogin(out.GetUser())
	if err != nil {
		return err
	}
	if !loggedin {
		return nil
	}

	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}

	data := pb.NotifySettingsResponse{
		User:          out.GetUser(),
		Notifications: u.Notify,
		MsgId:         out.GetMsgId(),
		NotifyTime:    int32(u.NotifyTime),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "NotifySettingsMsg",
		Data: dataBytes,
	}
	return nil

}

func (b *BotMain) HandleExportPrivateKey(msg tggateway.TgMessageMsg) error {
	out := pb.ExportPrivateKeyButton{}
	err := proto.Unmarshal(msg.Data, &out)
	if err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	loggedin, err := b.CheckLogin(out.GetUser())
	if err != nil {
		return err
	}
	if !loggedin {
		return nil
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	if !u.StorePrivatKey {
		data := pb.ErrorExportPKNotStore{
			User:  out.GetUser(),
			MsgId: out.GetMsgId(),
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "ErrorExportPKNotStore",
			Data: dataBytes,
		}
		return nil

	}
	// set user state ask password export private key
	err = b.State.SetUserState(out.GetUser().GetId(), []string{"ExportPK", "AskPass", strconv.Itoa(int(out.GetMsgId()))})
	if err != nil {
		return errors.Wrap(err, "failed update user state")
	}
	data := pb.ExportAskPassword{
		User:  out.GetUser(),
		MsgId: out.GetMsgId(),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "ExportAskPassword",
		Data: dataBytes,
	}
	return nil

}

func (b *BotMain) ExportPKAskPass(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) != 3 {
		return errors.Wrap(err, "failed get user state")
	}
	msgid, err := strconv.Atoi(state[2])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	pk, err := b.Crypto.PemFromDB(u.ID, msg.GetText())
	if err != nil {
		//bad pass
		data := pb.ExportIncorrectPassword{
			User:  msg.GetFrom(),
			MsgId: int64(msgid),
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "ExportIncorrectPassword",
			Data: dataBytes,
		}
		return nil
	}
	data := pb.AuthResponse{
		User:   msg.GetFrom(),
		Stage:  "SendPrivatKey",
		Data:   pk,
		Pubkey: u.PublicKey,
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AuthMsg",
		Data: dataBytes,
	}
	err = b.State.DeleteUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed delete user state")
	}

	return nil
}

func (b *BotMain) PrivacySettings(msg tggateway.TgMessageMsg) error {
	out := pb.TgPrivacySettingsButton{}
	err := proto.Unmarshal(msg.Data, &out)
	if err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	loggedin, err := b.CheckLogin(out.GetUser())
	if err != nil {
		return err
	}
	if !loggedin {
		return nil
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}

	data := pb.PrivacyMenu{
		User:      out.GetUser(),
		MsgId:     out.GetMsgId(),
		LogStatus: u.EnableLogging,
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "PrivacySettingsResponse",
		Data: dataBytes,
	}
	return nil

}

func (b *BotMain) ToggleLogging(msg tggateway.TgMessageMsg) error {
	out := pb.ToggleLogging{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	u, err = u.Update().SetEnableLogging(!u.EnableLogging).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	data := pb.PrivacyMenu{
		User:      out.GetUser(),
		MsgId:     out.GetMsgId(),
		LogStatus: u.EnableLogging,
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "PrivacySettingsResponse",
		Data: dataBytes,
	}
	return nil

}
