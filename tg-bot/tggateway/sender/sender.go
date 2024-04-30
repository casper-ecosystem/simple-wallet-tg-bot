package sender

import (
	"fmt"
	"strconv"

	"github.com/Simplewallethq/tg-bot/tggateway/messages"
	pb "github.com/Simplewallethq/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/tg-bot/tggateway/types"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	tele "gopkg.in/telebot.v3"
)

type Sender struct {
	bot         *tele.Bot
	MessageMsg  chan types.TgMessageMsg
	ResponseMsg chan types.TgResponseMsg
	logger      *logrus.Logger
}

func NewSender(bot *tele.Bot, income chan types.TgMessageMsg, outcome chan types.TgResponseMsg, logger *logrus.Logger) *Sender {
	return &Sender{
		bot:         bot,
		MessageMsg:  income,
		ResponseMsg: outcome,
		logger:      logger,
	}
}

func (S *Sender) ListenIncome() {
	for {
		//select {
		msg := <-S.ResponseMsg
		fmt.Println("received message", msg.Name)
		switch msg.Name {
		case "WelcomeMsg":
			S.SendWelcomeMsg(msg)
		case "AuthMsg":
			S.SendAuthMsg(msg)
		case "AuthRegisterMsg":
			S.SendAuthTypeMessage(msg)
		case "BalanceMsg":
			S.SendBalanceMsg(msg)
		case "LoadBalanceMsg":
			S.LoadBalanceMsg(msg)
		case "AskLoginMsg":
			S.SendLoginMsg(msg)
		case "LoginPassInvalidMsg":
			S.SendLoginPassInvalidMsg(msg)
		case "LoginSuccessMsg":
			S.SendLoginSuccessMsg(msg)
		case "LockMsg":
			S.SendLockMsg(msg)
		case "SettingsMsg":
			S.SendSettingsMsg(msg)
		case "LogoutMsg":
			S.SendLogoutMsg(msg)
		case "LoadHistoryMsg":
			S.LoadHistoryMsg(msg)
		case "TransferHistoryMsg":
			S.SendHistory(msg)
		case "RewardsHistoryMsg":
			S.SendRewardsHistory(msg)
		case "ChangeLogoutAskTime":
			S.SendChangeLogoutAskTime(msg)
		case "ChangeLogoutAskTimeResponse":
			S.SendChangeLogoutAskTimeResponse(msg)
		case "YieldMsg":
			S.SendYieldMsg(msg)
		case "YieldLoadResponse":
			S.SendYieldLoadMsg(msg)
		case "NotifyManyTasks":
			S.SendTooManyTasks(msg)
		//NotifyNewTransfer
		//NotifyNewDelegate
		//NotifyNewUndelegate
		case "NotifyNewTransfer":
			S.SendNotifyNewTransfer(msg)
		case "NotifyNewDelegate":
			S.SendNotifyNewDelegate(msg)
		case "NotifyNewUndelegate":
			S.SendNotifyNewUndelegate(msg)
		case "NotifySettingsMsg":
			S.SendNotifySettingsMsg(msg)
		case "NotifyNewRewards":
			S.SendNotifyNewRewards(msg)
		case "NotifyNewBalance":
			S.SendNotifyNewBalance(msg)
		case "AddressBookMsg":
			S.SendAddressBookMsg(msg)
		case "AskNameAddressBookMsg":
			S.AskNameAddressBookMsg(msg)
		case "AskAddressAddressBookMsg":
			S.AskAddressAddressBook(msg)
		case "AskAdressInvalidResponse":
			S.AskAddressInvalidAdress(msg)
		case "AddressBookDetailed":
			S.SendAddressBookDetailed(msg)
		case "DeleteAdressBookConfirm":
			S.DeleteAdressBookConfirm(msg)
		case "LogoutConfirmationMsg":
			S.SendLogoutConfirmationMsg(msg)
		case "NewTransferResponseStage1":
			S.SendNewTransferResponseStage1(msg)
		case "TransferAskAdress":
			S.SendTransferAskAdress(msg)
		case "TransferAskAmount":
			S.SendTransferAskAmount(msg)
		case "TransferAskMemo":
			S.SendAskMemo(msg)
		case "TransferAskConfirmation":
			S.SendTransferAskConfirmation(msg)
		case "SignDeployAskPassword":
			S.SendSignDeployAskPassword(msg)
		case "TransferSuccessResponse":
			S.SendTransferSuccessResponse(msg)
		case "TransferAddressNotValid":
			S.SendTransferAddressNotValid(msg)
		case "TransferAmountNotValid":
			S.SendTransferAmountNotValid(msg)
		case "TransferMemoNotValid":
			S.SendTransferMemoNotValid(msg)
		case "TransferAddressBookMsg":
			S.SendTransferAddressBookMsg(msg)
		case "TransferUnknownError":
			S.SendTransferUnknownError(msg)
		case "TransferBadPassword":
			S.SendTransferBadPassword(msg)
		case "SignDeployAskPK":
			S.SendSignDeployAskPK(msg)
		case "TransferBadPK":
			S.SendTransferBadPK(msg)
		case "DelegateListValidators":
			S.SendDelegateValidatorList(msg)
		case "DelegateAskAmount":
			S.SendDelegateAskAmount(msg)
		case "DelegateAskConfirmation":
			S.SendDelegateAskConfirmation(msg)
		case "DelegateSuccessResponse":
			S.SendDelegatorSuccessResponse(msg)
		case "UndelegateSuccessResponse":
			S.SendUndelegatorSuccessResponse(msg)
		case "UndelegateDelegatesList":
			S.SendUndelegateDelegatesList(msg)
		case "UndelegateAskAmount":
			S.SendUndelegateAskAmount(msg)
		case "UndelegateAskConfirmation":
			S.SendUndelegateAskConfirmation(msg)
		case "DepositMessage":
			S.SendNewDepositMessage(msg)
		case "ShowSwapPairs":
			S.SendSwapPairs(msg)
		case "SwapAskAmount":
			S.SendSwapAskAmount(msg)
		case "SwapShowEstimatedAmount":
			S.SendEstimatedAmount(msg)
		case "SwapAskRefundAddress":
			S.SendSwapAskRefundAddress(msg)
		case "SwapSuccessResponse":
			S.SendSwapSuccessResponse(msg)
		case "SwapShowChains":
			S.SendSwapShowChains(msg)
		case "SwapAskToAddress":
			S.SendSwapAskToAddress(msg)
		case "ExportAskPassword":
			S.SendExportAskPassword(msg)
		case "PrivacySettingsResponse":
			S.SendPrivacySettingsResponse(msg)
		case "ExportIncorrectPassword":
			S.SendExportIncorrectPassword(msg)
		case "SwapLoadMsg":
			S.SendSwapLoadMsg(msg)
		case "ErrorExportPKNotStore":
			S.SendErrorExportPKNotStore(msg)
		case "InvoicesListMsg":
			S.SendInvoicesMsg(msg)
		case "AskNewInvoiceName":
			S.AskNewInvoiceName(msg)
		case "AskInvoiceAmount":
			S.AskInvoiceAmount(msg)
		case "AskInvoiceRepeatability":
			S.AskInvoiceRepeatability(msg)
		case "AskInvoiceComment":
			S.AskInvoiceComment(msg)
		case "InvoiceCreateSuccess":
			S.InvoiceCreateSuccess(msg)
		case "InvoiceDetailed":
			S.InvoiceDetailed(msg)
		case "PayInvoiceNotRegisteredPM":
			S.PayInvoiceNotRegisteredPM(msg)
		case "PayInvoiceNotAviablePM":
			S.PayInvoiceNotAviablePM(msg)
		case "PayInvoiceRegisteredResponse":
			S.PayInvoiceRegisteredResponse(msg)
		case "PaymentsListResponse":
			S.PaymentsListMsg(msg)
		case "RecentInvoiceResponse":
			S.RecentInvoiceResponse(msg)
		case "ExportPaymentsInvoiceResponse":
			S.ExportPaymentsInvoiceResponse(msg)
		case "InvoiceRepeatabilityIsNotValid":
			S.SendInvoiceRepeatabilityIsNotValid(msg)
		case "DeleteInvoiceConfirmationMessage":
			S.DeleteInvoiceConfirmation(msg)
		}

		//	}
	}

}

type recipient struct {
	Id int64
}

func (r *recipient) Recipient() string {
	return strconv.FormatInt(r.Id, 10)
}

func (S *Sender) SendWelcomeMsg(msg types.TgResponseMsg) {
	out := pb.WelcomeResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetWelcomeMsg("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	if out.GetMsgId() != 0 {
		smsg := tele.StoredMessage{
			MessageID: strconv.Itoa(int(out.GetMsgId())),
			ChatID:    out.GetUser().GetId(),
		}
		_, err = S.bot.Edit(&smsg, what, opts...)
		if err != nil {
			S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
		}
	} else {
		recipient := &recipient{Id: out.User.GetId()}
		_, err = S.bot.Send(recipient, what, opts...)
		if err != nil {
			return
		}
	}
}

func (S *Sender) SendYieldLoadMsg(msg types.TgResponseMsg) {
	out := pb.YieldLoadingResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetCustomYieldMsg("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	smsg := tele.StoredMessage{
		MessageID: strconv.Itoa(int(out.GetMsgId())),
		ChatID:    out.GetUser().GetId(),
	}
	_, err = S.bot.Edit(&smsg, what, opts...)
	if err != nil {
		S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
	}
}

func (S *Sender) SendTooManyTasks(msg types.TgResponseMsg) {
	out := pb.TooManyTasksResponse{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return
	}
	what, opts, err := messages.GetTooManyTasksMsg("eng")
	if err != nil {
		what, opts, _ = messages.GetErrorMessage()
	}
	smsg := tele.StoredMessage{
		MessageID: strconv.Itoa(int(out.GetMsgId())),
		ChatID:    out.GetUser().GetId(),
	}
	_, err = S.bot.Edit(&smsg, what, opts...)
	if err != nil {
		S.logger.Error("error editing message: ", err, "uid: ", out.GetUser().GetId(), "msgid: ", out.GetMsgId())
	}
}
