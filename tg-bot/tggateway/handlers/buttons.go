package handlers

import (
	"log"
	"strconv"
	"strings"

	"github.com/Simplewallethq/tg-bot/tggateway/messages"
	pb "github.com/Simplewallethq/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/tg-bot/tggateway/types"
	"google.golang.org/protobuf/proto"
	tele "gopkg.in/telebot.v3"
)

func (h *Handler) ButtonHandler(c tele.Context) error {
	err := c.Respond()
	if err != nil {
		return err
	}
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	log.Println(data)

	switch data[0] {
	case "Lock":
		return h.LockHandler(c)

	case "balance":
		return h.BalanceHandler(c)
	case "mainMenu":
		return h.MainMenuHandler(c)
	case "logout":
		return h.LogoutHandler(c)
	case "history":
		return h.HistoryHandler(c)
	case "balanceHistory":
		return h.BalanceHistoryHandler(c)
	case "rewardsHistory":
		return h.RewardsHistoryHandler(c)
	case "moveHistory":
		return h.HistoryMoveHandler(c)
	case "moveRewardsHistory":
		return h.RewardsHistoryMoveHandler(c)
	case "ChangeLockTimeout":
		return h.ChangeLockTimeoutHandler(c)
	case "settings":
		return h.ShowSettings(c)
	case "yield":
		return h.YieldHandler(c)
	case "OnOffNotify":
		return h.OnOffNotifications(c)
	case "NotifySettings":
		return h.ShowNotifySettings(c)
	case "ChangeRewardsNotifyTime":
		return h.ChangeRewardsNotifyTime(c)
	case "addressBook":
		return h.AddressBookHandler(c)
	case "createEntryAdressBook":
		return h.CreateEntryAddressBookHandler(c)
	case "moveAddressBook":
		return h.MoveAddressBookHandler(c)
	case "moveTransferAddressBook":
		return h.MoveTransferAddressBookHandler(c)
	case "showAddress":
		return h.ShowAddressHandler(c)
	case "changeNameAddressBook":
		return h.ChangeNameAddressBookHandler(c)
	case "changeAddressAddressBook":
		return h.ChangeAddressAddressBookHandler(c)
	case "deleteAddressBook":
		return h.DeleteAddressBookHandler(c)
	case "ConfirmDeleteAdressBook":
		return h.ConfirmDeleteAdressBook(c)
	case "cancelLogout":
		return h.CancelLogout(c)
	case "cancelAddressBook":
		return h.CancelAddressBook(c)
	case "cancelChangeTimeout":
		return h.CancelChangeTimeout(c)
	case "addExistingWallet":
		return h.AddExistingWallet(c)
	case "NotStoreKey":
		return h.NotStoreKeyButtonHandler(c)
	case "StoreKey":
		return h.StoreKeyButtonHandler(c)
	case "createWallet":
		return h.createNewWallet(c)
	case "newTransfer":
		return h.newTransfer(c)
	case "TransferCustomAddress":
		return h.TransferCustomAddress(c)
	case "transferConfirm":
		return h.TransferConfirmButton(c)
	case "TransferMaximum":
		return h.TransferMaximum(c)
	case "TransferAddressBook":
		return h.TransferAddressBook(c)
	case "pickTransferAddress":
		return h.PickTransferAddress(c)
	case "transferWithoutMemo":
		return h.TransferWithoutMemo(c)
	case "newDelegate":
		return h.DelegateButton(c)
	case "newUndelegate":
		return h.UndelegateButton(c)
	case "pickDelegateValidator":
		return h.PickDelegateValidator(c)
	case "DelegateSelectAmount":
		return h.PickDelegateAmount(c)
	case "delegateConfirm":
		return h.DelegateConfirmButton(c)
	case "moveDelegateValidators":
		return h.MoveDelegateValidators(c)
	case "pickUndelegateValidator":
		return h.pickUndelegateValidator(c)
	case "UndelegateSelectAmount":
		return h.UndelegateSelectAmount(c)
	case "undelegateConfirm":
		return h.UndelegateConfirmButton(c)
	case "newDeposit":
		return h.NewDepositButton(c)
	case "depositBySwap":
		return h.SwapBySwapButton(c)
	case "moveSwapPairs":
		return h.MoveSwapPairs(c)
	case "pickSwapPair":
		return h.pickSwapPair(c)
	case "swapConfirmAmount":
		return h.SwapConfirmAmount(c)
	case "pickSwapChain":
		return h.pickSwapChain(c)
	case "withdrawBySwap":
		return h.WithdrawBySwapButton(c)
	case "ExportPrivat":
		return h.ExportPrivateKey(c)
	case "PrivacyOptions":
		return h.PrivacyOptions(c)
	case "toggleLogging":
		return h.ToggleLogging(c)
	case "invoices":
		return h.InvoicesHandler(c)
	case "createNewInvoice":
		return h.NewInvoiceHandler(c)
	case "moveInvoicesList":
		return h.MoveInvoicesHandler(c)
	case "showInvoice":
		return h.ShowInvoiceHandler(c)
	case "deleteInvoice":
		return h.DeleteInvoice(c)
	case "payInvoice":
		return h.PayInvoice(c)
	case "payInvoiceTransfer":
		return h.PayInvoiceTransfer(c)
	case "payInvoiceSwap":
		return h.PayInvoiceSwap(c)
	case "showInvoicePayments":
		return h.ShowInvoicePayments(c)
	case "movePaymentsList":
		return h.MovePaymentsHandler(c)
	case "recentInvoices":
		return h.ShowRecentInvoices(c)
	case "moveRecentInvoicesList":
		return h.MoveRecentInvoicesList(c)
	case "exportPaymentsInvoice":
		return h.ExportPaymentsInvoice(c)
	case "deleteInvoiceConfirm":
		return h.DeleteInvoiceConfirm(c)

	}
	return nil
}

func (h *Handler) LockHandler(c tele.Context) error {
	log.Printf("user %d press Lock", c.Sender().ID)
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
	msg := pb.TgLockButton{
		User: user,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "Lock",
		Data: out,
	}
	return nil
}

func (h *Handler) BalanceHandler(c tele.Context) error {
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
	msg := pb.TgBalanceButton{
		From:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "Balance",
		Data: out,
	}
	log.Printf("user %d press balance", c.Sender().ID)

	return nil
}

func (h *Handler) MainMenuHandler(c tele.Context) error {
	// what, opts, err := messages.GetWelcomeMsg("eng")
	// if err != nil {
	// 	what, opts, _ = messages.GetErrorMessage()
	// }
	// _, err = c.Bot().Edit(c.Callback().Message, what, opts...)
	// return err
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
	startMsgdata := pb.TgCommandStart{
		From:  user,
		MsgId: int64(c.Callback().Message.ID),
	}

	out, err := proto.Marshal(&startMsgdata)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "/start",
		Data: out,
	}
	return nil
}

func (h *Handler) LogoutHandler(c tele.Context) error {
	log.Printf("user %d press Logout", c.Sender().ID)
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
	msg := pb.TgLogoutButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "Logout",
		Data: out,
	}
	return nil
}

func (h *Handler) CancelLogout(c tele.Context) error {
	log.Printf("user %d press Logout", c.Sender().ID)
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
	msg := pb.CancelLogoutButton{
		User: user,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "CancelLogout",
		Data: out,
	}
	return nil
}

func (h *Handler) HistoryHandler(c tele.Context) error {
	what, opts, err := messages.GetHistoryMenu("eng")
	if err != nil {
		return err
	}
	_, err = c.Bot().Edit(c.Message(), what, opts...)
	return err
}

func (h *Handler) BalanceHistoryHandler(c tele.Context) error {
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
	msg := pb.TgBalanceButtonHistory{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Start: -1,
		End:   -1,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "BalanceHistory",
		Data: out,
	}
	return nil
}

func (h *Handler) HistoryMoveHandler(c tele.Context) error {
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 2 {
		start, err := strconv.Atoi(data[1])
		if err != nil {
			return err
		}
		end := start + 500
		user := &pb.User{
			Id:       c.Sender().ID,
			Username: c.Sender().Username,
		}
		msg := pb.TgBalanceButtonHistory{
			User:  user,
			MsgId: int64(c.Callback().Message.ID),
			Start: int64(start),
			End:   int64(end),
		}
		out, err := proto.Marshal(&msg)
		if err != nil {
			return err
		}
		h.MessagesMsg <- types.TgMessageMsg{
			Name: "BalanceHistory",
			Data: out,
		}
	}
	return nil
}

func (h *Handler) RewardsHistoryHandler(c tele.Context) error {
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
	msg := pb.TgButtonRewardsHistory{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
		Start: -1,
		End:   -1,
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "RewardsHistory",
		Data: out,
	}
	return nil
}

func (h *Handler) RewardsHistoryMoveHandler(c tele.Context) error {
	data := strings.Split(strings.TrimSpace(c.Callback().Data), "|")
	if len(data) == 2 {
		start, err := strconv.Atoi(data[1])
		if err != nil {
			return err
		}
		end := start + 10
		user := &pb.User{
			Id:       c.Sender().ID,
			Username: c.Sender().Username,
		}
		msg := pb.TgButtonRewardsHistory{
			User:  user,
			MsgId: int64(c.Callback().Message.ID),
			Start: int64(start),
			End:   int64(end),
		}
		out, err := proto.Marshal(&msg)
		if err != nil {
			return err
		}
		h.MessagesMsg <- types.TgMessageMsg{
			Name: "RewardsHistory",
			Data: out,
		}
	}
	return nil
}

func (h *Handler) ChangeLockTimeoutHandler(c tele.Context) error {
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
	msg := pb.TgChangeLockTimeoutButton{
		User:  user,
		MsgId: int64(c.Callback().Message.ID),
	}
	out, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	h.MessagesMsg <- types.TgMessageMsg{
		Name: "ChangeLockTimeout",
		Data: out,
	}
	return nil
}
