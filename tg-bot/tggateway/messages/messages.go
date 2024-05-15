package messages

import (
	"errors"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/messages/eng"
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
)

func GetWelcomeMsg(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetWelcomeMsg()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}
func GetMenuMsg(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetMenuMsg()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetWelcomeAuthMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetWelcomeAuthMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetAuthPasswordMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetAuthPasswordMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetErrorMessage() (interface{}, []interface{}, error) {
	what, opts := eng.GetErrorMessage()
	return what, opts, nil
}

func GetAuthInvalidPubKeyMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetAuthInvalidPubKeyMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetAuthInvalidPrivateMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetAuthInvalidPrivateMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetAuthRepeatPasswordMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetAuthRepeatPasswordMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetAuthRepeatPasswordInvalidMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetAuthRepeatPasswordInvalidMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetRegisterSuccessMessage(lang string, pubkey string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetRegisterSuccessMessage(pubkey)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSendPrivatKey(lang string, data []byte, pubkey string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSendPrivatKey(data, pubkey)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetLoadBalanceMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetLoadBalanceMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")

}

func GetBalanceMsg(lang string, balance types.BalanceResponse) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetBalanceMsg(balance)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetLoginMessage(lang string, LogoutManual bool) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetLoginMessage(LogoutManual)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetLoginPassInvalidMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetLoginPassInvalidMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetLoginSuccessMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetLoginSuccessMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetLockMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetLockMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetLogoutConfirmationMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetLogoutConfirmationMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetLogoutMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetLogoutMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSettingsMessage(lang string, settings types.SettingsResponse) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSettingsMessage(settings)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}
func GetNotifySettingsMessage(lang string, settings types.NotifySettingsResponse) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetNotifySettingsMessage(settings)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetHistoryMenu(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetHistoryMenu()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetLoadHistoryMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetLoadHistoryMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")

}

func GetHistoryMsg(lang string, history types.HistoryResponse) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetHistoryMsg(history)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetRewardsHistoryMsg(lang string, history types.RewardsHistoryResponse) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetRewardsHistoryMsg(history)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetChangeLockTimeoutMessage(lang string, currentTime int64) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetChangeLockTimeoutMessage(currentTime)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetChangeLockTimeoutMessageSuccess(lang string, current int64) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetChangeLockTimeoutMessageSuccess(current)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetYieldMsg(lang string, yield types.YieldResponse) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetYieldMsg(yield)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetCustomYieldMsg(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetCustomYieldMsg()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetTooManyTasksMsg(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetTooManyTasksMsg()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

//GetNotifyNewUndelegateMessage
//GetNotifyNewDelegateMessage
//GetNotifyNewTransferMessage

func GetNotifyNewUndelegateMessage(lang string, amount float64, validator string, era int64, balance float64) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetNotifyNewUndelegateMessage(amount, validator, era, balance)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetNotifyNewDelegateMessage(lang string, amount float64, validator string, height int64, balance float64) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetNotifyNewDelegateMessage(amount, validator, height, balance)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetNotifyNewTransferMessage(lang string, amount float64, from string, to string, balance float64) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetNotifyNewTransferMessage(amount, from, to, balance)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetNotifyNewRewards(lang string, rews types.NotifyNewRewards) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetNotifyNewRewards(rews)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetNotifyNewBalance(lang string, amount float64, old float64) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetNotifyNewBalance(amount, old)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetAddressBookMsg(lang string, addressBook types.AddressBookResponse) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetAddressBookMsg(addressBook)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetCreateEntryAddressBookNameMsg(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetCreateEntryAddressBookNameMsg()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetAskAddresAdressBookMsg(lang string, name string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetAskAddresAdressBookMsg(name)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}
func GetAskAddresInvalidAdress(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetAskAddresInvalidAdress()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetAddressDetailedMsg(lang string, data types.AddressBookDetailed) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetAddressDetailedMsg(data)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetDeleteEntryAddressBookConfirmationMessage(lang string, name, address string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetDeleteEntryAddressBookConfirmationMessage(name, address)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetChangeAuthTypeMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetChangeAuthTypeMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetAskStoreTheKeyMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetAskStoreTheKeyMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetAskPrivatKeyMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetAskPrivatKeyMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSendTransferStage1Message(lang string, balance string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSendTransferStage1Message(balance)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetTransferAskAdressMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.SendTransferAskAdressMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetTransferAskAmountMessage(lang string, balance float64, toPubkey string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetTransferAskAmountMessage(balance, toPubkey)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetTransferAskMemo(lang string, amount string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetTransferAskMemo(amount)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetTransferAskConfirmation(lang string, amount, topubkey, name string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetTransferAskConfirmation(amount, topubkey, name)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSignDeployAskPasswordMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSignDeployAskPasswordMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetExportAskPasswordMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetExportAskPasswordMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetExportIncorrectPasswordMessage(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetExportIncorrectPasswordMessage()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSignDeployAskPK(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSignDeployAskPK()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSuccessTransferMessage(lang string, amount string, toPubkey string, toName string, memo uint64, hash string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSuccessTransferMessage(amount, toPubkey, toName, memo, hash)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetTransferAddressIsNotValidMsg(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetTransferAddressIsNotValidMsg()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetTransferAmountIsNotValidMsg(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetTransferAmountIsNotValidMsg()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetTransferMemoIsNotValidMsg(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetTransferMemoIsNotValidMsg()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetTransferBadPassword(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetTransferBadPassword()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetTransferBadPK(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetTransferBadPK()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetTransferUnknownError(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetTransferUnknownError()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetTransferAddressBookMsg(lang string, addressBook types.AddressBookResponse) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetTransferAddressBookMsg(addressBook)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetDelegatorValidators(lang string, data types.DelegateValidatorsResponse) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetDelegatorValidators(data)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetUndelegateDelegates(lang string, data types.UndelegateDelegatesList) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetUndelegateDelegates(data)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetDelegateAskAmountMessage(lang string, balance float64, validator string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetDelegateAskAmountMessage(balance, validator)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetDelegateAskConfirmation(lang string, amount float64, validator string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetDelegateAskConfirmation(amount, validator)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSuccessDelegateMessage(lang string, amount string, validator string, hash string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSuccessDelegateMessage(amount, validator, hash)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSuccessUndelegateMessage(lang string, amount string, validator string, hash string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSuccessUndelegateMessage(amount, validator, hash)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetUndelegateAskAmountMessage(lang string, balance string, validator string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetUndelegateAskAmountMessage(balance, validator)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetUndelegateAskConfirmation(lang string, amount float64, validator string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetUndelegateAskConfirmation(amount, validator)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetDepositMessage(lang string, address string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetDepositMessage(address)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSwapPairs(lang string, offset, total int64, pairs []string, swapType string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSwapPairs(offset, total, pairs, swapType)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSwapChains(lang string, data *pb.ShowSwapChains) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSwapChains(data)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSwapAskAmount(lang string, fromCur string, toCur string, min, max float64) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSwapAskAmount(fromCur, toCur, min, max)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSwapShowEstimated(lang string, estim float64, curr string, amount float64, currFrom string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSwapShowEstimated(estim, curr, amount, currFrom)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSwapAskRefund(lang string, curr string, chain string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSwapAskRefund(curr, chain)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSwapSuccess(lang string, id, addr, fromCur string, toCur string, amount float64, estimated float64) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSwapSuccess(id, addr, fromCur, toCur, amount, estimated)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSwapAskAdress(lang string, curr string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSwapAskAddress(curr)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

//GetPrivacySettingsMessage

func GetPrivacySettingsMessage(lang string, logging bool) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetPrivacySettingsMessage(logging)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetSwapLoadMsg(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetSwapLoadMsg()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetErrorExportPKNotStore(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetErrorExportPKNotStore()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetInvoicesListMsg(lang string, invoices *pb.InvoicesListResponse) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetInvoicesListMsg(invoices)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func AskInvoiceName(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.AskInvoiceName()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func AskInvoiceAmount(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.AskInvoiceAmount()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func AskInvoiceRepeatability(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.AskInvoiceRepeatability()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func AskInvoiceComment(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.AskInvoiceComment()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func InvoiceCreateSuccess(lang string, res *pb.InvoiceCreateSuccess, botname string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.InvoiceCreateSuccess(res, botname)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func InvoiceDetailed(lang string, res *pb.InvoiceDetailed, botname string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.InvoiceDetailed(res, botname)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func InvoiceAskRegisterPM(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.InvoiceAskRegisterPM()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func PayInvoiceNotAviablePM(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.PayInvoiceNotAviablePM()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func PayInvoiceRegisteredResponse(lang string, res *pb.PayInvoiceRegisteredResponse) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.PayInvoiceRegisteredResponse(res)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func BlockForGroup(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.BlockForGroup()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetPaymentsListMsg(lang string, payments *pb.PaymentsListResponse) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetPaymentsListMsg(payments)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetRecentlyInvoices(lang string, inv *pb.RecentInvoicesListResponse) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetRecentlyInvoices(inv)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func GetExportPaymentsInvoice(lang string, data []byte, short string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetExportPaymentsInvoice(data, short)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}

func DeleteInvoiceConfirmation(lang string, id uint64) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.DeleteInvoiceConfirmation(id)
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")

}

func GetInvoiceRepeatabilityIsNotValid(lang string) (interface{}, []interface{}, error) {
	switch lang {
	case "eng":
		what, opts := eng.GetInvoiceRepeatabilityIsNotValid()
		return what, opts, nil
	}
	return nil, nil, errors.New("language not found")
}
