package botmain

import (
	"context"
	"log"
	"time"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	tggateway "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/proto"
)

const DEFAULT_LOGOUT_TIMEOUT = 15

func (b *BotMain) HandleAuth(msg *pb.TgTextMessage) error {
	_, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	//path := strings.Split(u.State, "/")
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	log.Println("state in auth botmain: ", state)
	switch state[0] {
	case "Auth":
		switch state[1] {
		case "1":
			err := b.HandleAuth1(msg)
			return err
		case "2":
			err := b.HandleAuth2(msg)
			return err
		case "3":
			err := b.HandleAuth3(msg)
			return err
		}
	case "CreateNewWallet":
		switch state[1] {
		case "2":
			err := b.HandleCreateNewWallet2(msg)
			return err
		case "3":
			err := b.HandleCreateNewWallet3(msg)
			return err
		}
	case "CreateNewWalletWithoutPK":
		switch state[1] {
		case "2":
			err := b.HandleCreateNewWalletWithoutPK2(msg)
			return err
		case "3":
			err := b.HandleCreateNewWalletWithoutPK3(msg)
			return err
		}
	case "AddWalletWithPK":
		switch state[1] {
		case "2":
			err := b.HandleAddWalletWithPK2(msg)
			return err
		case "3":
			err := b.HandleAddWalletWithPK3(msg)
			return err
		case "4":
			err := b.HandleAddWalletWithPK4(msg)
			return err
		}
	}
	return nil
}

func (b *BotMain) HandleAuth1(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	pubkey := msg.GetText()
	valid, err := b.Restclient.IsAddress(b.RPCNode, pubkey)
	if err != nil {
		return errors.Wrap(err, "failed check address")
	}
	if !valid {
		data := pb.AuthResponse{
			User:  msg.GetFrom(),
			Error: "Invalid",
			Stage: "1",
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "AuthMsg",
			Data: dataBytes,
		}
		return nil
	}
	//_, err = u.Update().SetState("auth/2").SetPublicKey(pubkey).Save(context.Background())
	_, err = u.Update().SetPublicKey(pubkey).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	//msgid := strconv.Itoa(int(msg.MsgId))
	err = b.State.SetUserState(msg.GetFrom().GetId(), []string{"Auth", "2"})
	if err != nil {
		return errors.Wrap(err, "failed update user state")
	}
	data := pb.AuthResponse{
		User:  msg.GetFrom(),
		Stage: "2",
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AuthMsg",
		Data: dataBytes,
	}
	return nil

}

func (b *BotMain) HandleAuth2(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	passw := msg.GetText()
	//log.Printf("user with id %d and pubkey %s send password %s", u.ID, u.PublicKey, passw)
	hash, err := bcrypt.GenerateFromPassword([]byte(passw), bcrypt.MinCost)
	if err != nil {
		return errors.Wrap(err, "failed generate hash")
	}
	_, err = u.Update().SetPassword(string(hash)).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	err = b.State.SetUserState(msg.GetFrom().GetId(), []string{"Auth", "3"})
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	data := pb.AuthResponse{
		User:  msg.GetFrom(),
		Stage: "3",
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AuthMsg",
		Data: dataBytes,
	}
	return nil
}

func (b *BotMain) HandleAuth3(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	passw := msg.GetText()
	//log.Printf("user with id %d and pubkey %s send password %s", u.ID, u.PublicKey, passw)
	matchErr := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passw))
	if matchErr != nil {
		err = b.State.SetUserState(msg.GetFrom().GetId(), []string{"Auth", "2"})
		if err != nil {
			return errors.Wrap(err, "failed update user")
		}
		data := pb.AuthResponse{
			User:  msg.GetFrom(),
			Error: "Invalid",
			Stage: "3",
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "AuthMsg",
			Data: dataBytes,
		}
		return nil
	}
	err = b.State.DeleteUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed delete user state")
	}
	u, err = u.Update().SetLoggedIn(true).SetLastAccess(time.Now()).SetRegistered(true).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	data := pb.AuthResponse{
		User:   msg.GetFrom(),
		Stage:  "4",
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
	return nil
}

func (b *BotMain) CheckLogin(pbuser *pb.User) (bool, error) {
	u, err := b.DB.User.Query().Where(user.ID(pbuser.GetId())).Only(context.Background())
	if err != nil {
		return false, errors.Wrap(err, "failed get user")
	}
	if !u.Registered {
		data := pb.AuthRegisterType{
			User:  pbuser,
			MsgId: -1,
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return false, errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "AuthRegisterMsg",
			Data: dataBytes,
		}
	}
	var timeout int64
	if u.LockTimeout > 0 {
		timeout = u.LockTimeout
	} else {
		timeout = DEFAULT_LOGOUT_TIMEOUT
	}
	if u.LastAccess.Add(time.Minute*time.Duration(timeout)).Before(time.Now()) && u.PublicKey != "" {
		_, err := u.Update().SetLoggedIn(false).SetLockedManual(false).Save(context.Background())
		if err != nil {
			return false, errors.Wrap(err, "failed update user")
		}
		u.LoggedIn = false
	}
	_, err = u.Update().SetLastAccess(time.Now()).Save(context.Background())
	if err != nil {
		return false, errors.Wrap(err, "failed update user")
	}
	if !u.LoggedIn {
		//_, err := u.Update().SetState("login").Save(context.Background())
		err = b.State.SetUserState(pbuser.GetId(), []string{"login"})
		if err != nil {
			return false, errors.Wrap(err, "failed update user")
		}
		data := pb.AskLoginResponse{
			User:         pbuser,
			ManualLogout: u.LockedManual,
			MsgId:        0,
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return false, errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "AskLoginMsg",
			Data: dataBytes,
		}
	}

	return u.LoggedIn, nil
}

func (b *BotMain) CheckLoginWithEditMsg(pbuser *pb.User, msgid int64) (bool, error) {
	u, err := b.DB.User.Query().Where(user.ID(pbuser.GetId())).Only(context.Background())
	if err != nil {
		return false, errors.Wrap(err, "failed get user")
	}
	var timeout int64
	if u.LockTimeout > 0 {
		timeout = u.LockTimeout
	} else {
		timeout = DEFAULT_LOGOUT_TIMEOUT
	}
	if u.LastAccess.Add(time.Minute*time.Duration(timeout)).Before(time.Now()) && u.PublicKey != "" {
		_, err := u.Update().SetLoggedIn(false).SetLockedManual(false).Save(context.Background())
		if err != nil {
			return false, errors.Wrap(err, "failed update user")
		}
		u.LoggedIn = false
	}
	_, err = u.Update().SetLastAccess(time.Now()).Save(context.Background())
	if err != nil {
		return false, errors.Wrap(err, "failed update user")
	}
	if !u.LoggedIn {
		//_, err := u.Update().SetState("login").Save(context.Background())
		err = b.State.SetUserState(pbuser.GetId(), []string{"login"})
		if err != nil {
			return false, errors.Wrap(err, "failed update user")
		}
		data := pb.AskLoginResponse{
			User:         pbuser,
			ManualLogout: u.LockedManual,
			MsgId:        msgid,
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return false, errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "AskLoginMsg",
			Data: dataBytes,
		}
	}

	return u.LoggedIn, nil
}

func (b *BotMain) HandleLogin(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	passw := msg.GetText()
	//log.Printf("user with id %d and pubkey %s send password %s", u.ID, u.PublicKey, passw)
	matchErr := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passw))
	if matchErr != nil {
		data := pb.LoginPassInvalidResponse{
			User: msg.GetFrom(),
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "LoginPassInvalidMsg",
			Data: dataBytes,
		}
		return nil
	}
	err = b.State.DeleteUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed delete user state")
	}
	_, err = u.Update().SetLoggedIn(true).SetLockedManual(false).SetRegistered(true).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	data := pb.LoginSuccessResponse{
		User: msg.GetFrom(),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "LoginSuccessMsg",
		Data: dataBytes,
	}
	return nil
}

func (b *BotMain) HandleLock(msg tggateway.TgMessageMsg) error {
	out := pb.TgLockButton{}
	err := proto.Unmarshal(msg.Data, &out)
	if err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	_, err = u.Update().SetLoggedIn(false).SetLockedManual(true).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	data := pb.LockResponse{
		User: out.GetUser(),
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "LockMsg",
		Data: dataBytes,
	}
	return nil

}

func (b *BotMain) AddExistingWallet(msg tggateway.TgMessageMsg) error {
	out := pb.AddExistingWalletButton{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	err := b.State.SetUserState(out.GetUser().GetId(), []string{"Auth", "1"})
	if err != nil {
		return err
	}
	data := pb.AuthResponse{
		User:  out.GetUser(),
		Stage: "AskStoreKey",
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AuthMsg",
		Data: dataBytes,
	}
	return nil
}

func (b *BotMain) HandleStoreTheKeyButton(msg tggateway.TgMessageMsg) error {
	//TODO logic
	//log.Println("HERE")
	out := pb.StoreTheKeyButton{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	switch state[0] {
	case "Auth":
		if !out.IsStore {
			err = b.State.SetUserState(out.GetUser().GetId(), []string{"Auth", "1"})
			if err != nil {
				return err
			}
			u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
			if err != nil {
				return errors.Wrap(err, "failed get user")
			}
			_, err = u.Update().SetStorePrivatKey(out.GetIsStore()).Save(context.Background())
			if err != nil {
				return errors.Wrap(err, "failed update user")
			}
			data := pb.AuthResponse{
				User:  out.GetUser(),
				Stage: "1",
			}
			dataBytes, err := proto.Marshal(&data)
			if err != nil {
				return errors.Wrap(err, "failed marshal")
			}
			b.ResponseChan <- tggateway.TgResponseMsg{
				Name: "AuthMsg",
				Data: dataBytes,
			}
		} else {
			err = b.State.SetUserState(out.GetUser().GetId(), []string{"AddWalletWithPK", "2"})
			if err != nil {
				return err
			}
			u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
			if err != nil {
				return errors.Wrap(err, "failed get user")
			}
			_, err = u.Update().SetStorePrivatKey(out.GetIsStore()).Save(context.Background())
			if err != nil {
				return errors.Wrap(err, "failed update user")
			}
			data := pb.AuthResponse{
				User:  out.GetUser(),
				Stage: "2",
			}
			dataBytes, err := proto.Marshal(&data)
			if err != nil {
				return errors.Wrap(err, "failed marshal")
			}
			b.ResponseChan <- tggateway.TgResponseMsg{
				Name: "AuthMsg",
				Data: dataBytes,
			}
		}

	case "CreateNewWallet":
		if out.IsStore {
			err = b.State.SetUserState(out.GetUser().GetId(), []string{"CreateNewWallet", "2"})
			if err != nil {
				return err
			}

			u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
			if err != nil {
				return errors.Wrap(err, "failed get user")
			}
			_, err = u.Update().SetStorePrivatKey(out.GetIsStore()).Save(context.Background())
			if err != nil {
				return errors.Wrap(err, "failed update user")
			}
			data := pb.AuthResponse{
				User:  out.GetUser(),
				Stage: "2",
			}
			dataBytes, err := proto.Marshal(&data)
			if err != nil {
				return errors.Wrap(err, "failed marshal")
			}
			b.ResponseChan <- tggateway.TgResponseMsg{
				Name: "AuthMsg",
				Data: dataBytes,
			}
		} else {
			err = b.State.SetUserState(out.GetUser().GetId(), []string{"CreateNewWalletWithoutPK", "2"})
			if err != nil {
				return err
			}

			u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
			if err != nil {
				return errors.Wrap(err, "failed get user")
			}
			_, err = u.Update().SetStorePrivatKey(out.GetIsStore()).Save(context.Background())
			if err != nil {
				return errors.Wrap(err, "failed update user")
			}
			data := pb.AuthResponse{
				User:  out.GetUser(),
				Stage: "2",
			}
			dataBytes, err := proto.Marshal(&data)
			if err != nil {
				return errors.Wrap(err, "failed marshal")
			}
			b.ResponseChan <- tggateway.TgResponseMsg{
				Name: "AuthMsg",
				Data: dataBytes,
			}
		}
	}

	return nil
}

func (b *BotMain) CreateNewWallet(msg tggateway.TgMessageMsg) error {
	out := pb.CreateNewWalletButton{}
	log.Println("botmain: create new wallet")
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	err := b.State.SetUserState(out.GetUser().GetId(), []string{"CreateNewWallet", "1"})
	if err != nil {
		return err
	}
	data := pb.AuthResponse{
		User:  out.GetUser(),
		Stage: "AskStoreKey",
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AuthMsg",
		Data: dataBytes,
	}
	return nil
}

func (b *BotMain) HandleCreateNewWallet2(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	passw := msg.GetText()
	//log.Printf("user with id %d and pubkey %s send password %s", u.ID, u.PublicKey, passw)
	hash, err := bcrypt.GenerateFromPassword([]byte(passw), bcrypt.MinCost)
	if err != nil {
		return errors.Wrap(err, "failed generate hash")
	}
	_, err = u.Update().SetPassword(string(hash)).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	err = b.State.SetUserState(msg.GetFrom().GetId(), []string{"CreateNewWallet", "3"})
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	data := pb.AuthResponse{
		User:  msg.GetFrom(),
		Stage: "3",
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AuthMsg",
		Data: dataBytes,
	}
	return nil
}

func (b *BotMain) HandleCreateNewWalletWithoutPK2(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	passw := msg.GetText()
	//log.Printf("user with id %d and pubkey %s send password %s", u.ID, u.PublicKey, passw)
	hash, err := bcrypt.GenerateFromPassword([]byte(passw), bcrypt.MinCost)
	if err != nil {
		return errors.Wrap(err, "failed generate hash")
	}
	_, err = u.Update().SetPassword(string(hash)).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	err = b.State.SetUserState(msg.GetFrom().GetId(), []string{"CreateNewWalletWithoutPK", "3"})
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	data := pb.AuthResponse{
		User:  msg.GetFrom(),
		Stage: "3",
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AuthMsg",
		Data: dataBytes,
	}
	return nil
}

func (b *BotMain) HandleCreateNewWallet3(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	passw := msg.GetText()
	//log.Printf("user with id %d and pubkey %s send password %s", u.ID, u.PublicKey, passw)
	matchErr := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passw))
	if matchErr != nil {
		err = b.State.SetUserState(msg.GetFrom().GetId(), []string{"CreateNewWallet", "2"})
		if err != nil {
			return errors.Wrap(err, "failed update user")
		}
		data := pb.AuthResponse{
			User:  msg.GetFrom(),
			Error: "Invalid",
			Stage: "3",
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "AuthMsg",
			Data: dataBytes,
		}
		return nil
	}
	pem, err := b.Crypto.GenerateNewUserWallet(u.ID, passw)
	if err != nil {
		return errors.Wrap(err, "failed generate new user wallet")
	}
	err = b.State.DeleteUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed delete user state")
	}
	u, err = u.Update().SetLoggedIn(true).SetLastAccess(time.Now()).SetRegistered(true).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	data := pb.AuthResponse{
		User:   msg.GetFrom(),
		Stage:  "SendPrivatKey",
		Data:   pem,
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
	return nil
}

func (b *BotMain) HandleCreateNewWalletWithoutPK3(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	passw := msg.GetText()
	//log.Printf("user with id %d and pubkey %s send password %s", u.ID, u.PublicKey, passw)
	matchErr := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passw))
	if matchErr != nil {
		err = b.State.SetUserState(msg.GetFrom().GetId(), []string{"CreateNewWalletWithoutPK", "2"})
		if err != nil {
			return errors.Wrap(err, "failed update user")
		}
		data := pb.AuthResponse{
			User:  msg.GetFrom(),
			Error: "Invalid",
			Stage: "3",
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "AuthMsg",
			Data: dataBytes,
		}
		return nil
	}
	pem, err := b.Crypto.GenerateNewUserWalletWithoutStorePK(u.ID, passw)
	if err != nil {
		return errors.Wrap(err, "failed generate new user wallet")
	}
	err = b.State.DeleteUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed delete user state")
	}
	u, err = u.Update().SetLoggedIn(true).SetLastAccess(time.Now()).SetRegistered(true).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	data := pb.AuthResponse{
		User:   msg.GetFrom(),
		Stage:  "SendPrivatKey",
		Data:   pem,
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
	return nil
}

func (b *BotMain) HandleAddWalletWithPK2(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	passw := msg.GetText()
	//log.Printf("user with id %d and pubkey %s send password %s", u.ID, u.PublicKey, passw)
	hash, err := bcrypt.GenerateFromPassword([]byte(passw), bcrypt.MinCost)
	if err != nil {
		return errors.Wrap(err, "failed generate hash")
	}
	_, err = u.Update().SetPassword(string(hash)).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	err = b.State.SetUserState(msg.GetFrom().GetId(), []string{"AddWalletWithPK", "3"})
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	data := pb.AuthResponse{
		User:  msg.GetFrom(),
		Stage: "3",
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AuthMsg",
		Data: dataBytes,
	}
	return nil
}

func (b *BotMain) HandleAddWalletWithPK3(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	passw := msg.GetText()
	//log.Printf("user with id %d and pubkey %s send password %s", u.ID, u.PublicKey, passw)
	matchErr := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passw))
	if matchErr != nil {
		err = b.State.SetUserState(msg.GetFrom().GetId(), []string{"AddWalletWithPK", "2"})
		if err != nil {
			return errors.Wrap(err, "failed update user")
		}
		data := pb.AuthResponse{
			User:  msg.GetFrom(),
			Error: "Invalid",
			Stage: "3",
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "AuthMsg",
			Data: dataBytes,
		}
		return nil
	}
	err = b.Crypto.SaveUserPassword(u.ID, passw)
	if err != nil {
		return errors.Wrap(err, "failed save user password")
	}
	err = b.State.SetUserState(msg.GetFrom().GetId(), []string{"AddWalletWithPK", "4"})
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}
	_, err = u.Update().SetLoggedIn(true).SetLastAccess(time.Now()).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	data := pb.AuthResponse{
		User:  msg.GetFrom(),
		Stage: "AskPrivatKey",
	}
	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}
	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "AuthMsg",
		Data: dataBytes,
	}
	return nil
}

func (b *BotMain) HandleAddWalletWithPK4(msg *pb.TgTextMessage) error {
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	pemText := msg.GetText()
	log.Printf("user with id %d and pubkey %s send pem %s", u.ID, u.PublicKey, pemText)
	err = b.Crypto.NewWalletFromPEM(u.ID, []byte(pemText))
	if err != nil {
		log.Println(err)
		data := pb.AuthResponse{
			User:  msg.GetFrom(),
			Error: "Invalid",
			Stage: "AskPrivatKey",
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "AuthMsg",
			Data: dataBytes,
		}
		return nil
	}
	err = b.State.DeleteUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed delete user state")
	}
	u, err = u.Update().SetLoggedIn(true).SetLastAccess(time.Now()).SetRegistered(true).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	data := pb.AuthResponse{
		User:   msg.GetFrom(),
		Stage:  "4",
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
	return nil

}
