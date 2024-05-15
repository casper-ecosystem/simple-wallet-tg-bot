package botmain

import (
	"context"
	"log"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/botmain/crypto"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/botmain/notificator"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/botmain/restclient"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/botmain/swap"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/botmain/userstate"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/botmain/validators"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/botmain/taskrecover"
	tggateway "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
)

type BotMain struct {
	MessagesChan      chan tggateway.TgMessageMsg
	ResponseChan      chan tggateway.TgResponseMsg
	DB                *ent.Client
	TestUserState     map[int64]string
	Restclient        *restclient.Client
	TaskRecoverer     *taskrecover.TaskRecoverer
	Notificator       *notificator.Notificator
	ValidatorsCrawler *validators.Crawler
	logger            *logrus.Logger
	State             *userstate.State
	Crypto            *crypto.Crypto
	RPCNode           string
	SwapClient        *swap.Client
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
	Schema   string
}

func NewDBClient(cfg DBConfig) (*ent.Client, error) {
	//	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=tgbot password=changeme sslmode=disable")
	client, err := ent.Open("postgres", "host="+cfg.Host+" port="+cfg.Port+" user="+cfg.User+" dbname="+cfg.DBName+" password="+cfg.Password+" sslmode="+cfg.SSLMode)
	if err != nil {
		return nil, errors.Wrap(err, "failed opening connection to postgres")
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, errors.Wrap(err, "failed creating schema resources")
	}
	return client, nil

}
func NewBotMain(income chan tggateway.TgMessageMsg, outcome chan tggateway.TgResponseMsg, DB *ent.Client, resthost string, rpcnode string, crypto *crypto.Crypto, logger *logrus.Logger, swap *swap.Client) *BotMain {

	return &BotMain{
		Restclient:        restclient.NewClient(resthost),
		TestUserState:     make(map[int64]string),
		MessagesChan:      income,
		ResponseChan:      outcome,
		DB:                DB,
		TaskRecoverer:     taskrecover.NewTaskRecoverer(DB, income, outcome, logger),
		Notificator:       notificator.NewNotificator(DB, outcome, resthost, rpcnode, logger),
		ValidatorsCrawler: validators.NewValidatorsCrawler(DB, resthost, rpcnode, logger),
		logger:            logger,
		State:             userstate.NewState(DB, logger),
		RPCNode:           rpcnode,
		Crypto:            crypto,
		SwapClient:        swap,
	}
}

func NewBotMainWithToken(income chan tggateway.TgMessageMsg, outcome chan tggateway.TgResponseMsg, DB *ent.Client, resthost string, rpcnode string, RESTtoken string, crypto *crypto.Crypto, logger *logrus.Logger, swap *swap.Client) *BotMain {

	return &BotMain{
		Restclient:        restclient.NewClientWithToken(resthost, RESTtoken),
		TestUserState:     make(map[int64]string),
		MessagesChan:      income,
		ResponseChan:      outcome,
		DB:                DB,
		TaskRecoverer:     taskrecover.NewTaskRecoverer(DB, income, outcome, logger),
		Notificator:       notificator.NewNotificatorWithToken(DB, outcome, resthost, rpcnode, RESTtoken, logger),
		ValidatorsCrawler: validators.NewValidatorsCrawlerWithToken(DB, resthost, rpcnode, logger, RESTtoken),
		logger:            logger,
		State:             userstate.NewState(DB, logger),
		RPCNode:           rpcnode,
		Crypto:            crypto,
		SwapClient:        swap,
	}
}
func (b *BotMain) HandleIncome() {
	for {

		//select {
		msg := <-b.MessagesChan
		go func() {
			var err error
			log.Println("botmain recieve income message", msg.Name)
			switch msg.Name {
			case "/start":
				err = b.HandleStart(msg)
			case "OnText":
				err = b.HandleText(msg)
			case "Balance":
				err = b.HandleBalance(msg)
			case "Lock":
				err = b.HandleLock(msg)
			case "Settings":
				err = b.HandleSettings(msg)
			case "Logout":
				err = b.HandleLogout(msg)
			case "BalanceHistory":
				err = b.HandleBalanceHistory(msg)
			case "RewardsHistory":
				err = b.HandleRewards(msg)
			case "ChangeLockTimeout":
				err = b.ChangeLockTimeoutButtonHandler(msg)
			case "Yield":
				err = b.HandleYield(msg)
			case "OnOffNotifications":
				err = b.HandleOnOffNotifications(msg)
			case "NotifySettings":
				err = b.HandleNotifySettings(msg)
			case "ChangeRewardsNotifyTime":
				err = b.HandleChangeRewardsNotifyTime(msg)
			case "AddressBook":
				err = b.HandleAddressBook(msg)
			case "CreateEntryAddressBook":
				err = b.CreateEntryAddressBook(msg)
			case "AskAddressBookDetailed":
				err = b.AskAddressBookDetailed(msg)
			case "ChangeNameAddressBook":
				err = b.ChangeNameAddressBook(msg)
				if err != nil {
					log.Println("failed handle income message", err)
				}
			case "ChangeAddressAddressBook":
				err = b.ChangeAddressAddressBook(msg)
			case "DeleteEntryAddressBook":
				err = b.DeleteEntryAddressBook(msg)
			case "DeleteEntryAddressBookConfirm":
				err = b.DeleteEntryAddressBookConfirm(msg)
			case "CancelAddressBook":
				err = b.CancelAddressBook(msg)
			case "CancelLogout":
				err = b.CancelLogout(msg)
			case "CancelChangeTimeout":
				err = b.CancelChangeTimeout(msg)
			case "AddExistingWallet":
				err = b.AddExistingWallet(msg)
			case "StoreTheKeyMenu":
				err = b.HandleStoreTheKeyButton(msg)
			case "CreateNewWallet":
				err = b.CreateNewWallet(msg)
			case "NewTransferButton":
				err = b.HandleNewTransferButton(msg)
			case "TransferCustomAddress":
				err = b.HandleTransferCustomAddress(msg)
			case "TransferConfirmButton":
				err = b.HandleTransferConfirmButton(msg)
			case "TransferMaximumButton":
				err = b.HandleTransferSetMaximumAmount(msg)
			case "TransferAddressBookButton":
				err = b.HandleTransferAddressBookButton(msg)
			case "PickTransferAddress":
				err = b.PickAddressFromAddressBook(msg)
			case "TransferWithoutMemo":
				err = b.HandleTransferWithoutMemo(msg)
			case "NewDelegateButton":
				err = b.HandleNewDelegateButton(msg)
			case "pickDelegateValidator":
				err = b.PickDelegateValidator(msg)
			case "pickDelegateAmount":
				err = b.PickDelegateAmount(msg)
			case "DelegateConfirmButton":
				err = b.DelegateConfirmButton(msg)
			case "NewUndelegateButton":
				err = b.HandleNewUnelegateButton(msg)
			case "pickUndelegateValidator":
				err = b.PickUndelegateValidator(msg)
			case "pickUndelegateAmount":
				err = b.PickUndelegateAmount(msg)
			case "UndelegateConfirmButton":
				err = b.UndelegateConfirmButton(msg)
			case "NewDepositButton":
				err = b.NewDepositButton(msg)
			case "SwapBySwapButton":
				err = b.SwapBySwapButton(msg)
			case "AskSwapPairs":
				err = b.AskSwapPairs(msg)
			case "pickSwapPair":
				err = b.PickSwapPair(msg)
			case "SwapConfirmAmount":
				err = b.SwapConfirmAmount(msg)
			case "pickSwapChain":
				err = b.PickSwapChain(msg)
			case "ExportPrivateKeyButton":
				err = b.HandleExportPrivateKey(msg)
			case "PrivacySettings":
				err = b.PrivacySettings(msg)
			case "ToggleLogging":
				err = b.ToggleLogging(msg)
			case "Invoices":
				err = b.HandleInvoiceButton(msg)
			case "NewInvoiceButton":
				err = b.CreateNewInvoice(msg)
			case "AskInvoiceDetailed":
				err = b.AskInvoiceDetailed(msg)
			case "DeleteInvoice":
				err = b.DeleteInvoice(msg)
			case "DeleteInvoiceConfirm":
				err = b.DeleteInvoiceConfirm(msg)
			case "PayInvoiceHandler":
				err = b.HandlePayInvoice(msg)
			case "PayInvoiceTransfer":
				err = b.PayInvoiceTransfer(msg)
			case "PayInvoiceSwap":
				err = b.PayInvoiceSwap(msg)
			case "ShowInvoicePayments":
				err = b.ShowInvoicePayments(msg)
			case "ShowRecentInvoices":
				err = b.ShowRecentInvoices(msg)
			case "ExportPaymentsInvoice":
				err = b.ExportPaymentsInvoice(msg)

			}
			if err != nil {
				b.logger.Error(err)
			}
		}()

	}

}

func (b *BotMain) HandleText(msg tggateway.TgMessageMsg) error {
	out := pb.TgTextMessage{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	if out.GetFrom().GetGroup() {
		//todo group logic
		return nil
	}
	_, err := b.DB.User.Query().Where(user.ID(out.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		var entErr *ent.NotFoundError
		if errors.As(err, &entErr) {
			err := b.HandleStartAuth(msg)
			return err
		}
		return errors.Wrap(err, "failed get user")
	}
	state, err := b.State.GetUserState(out.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) == 0 {
		err := b.HandleStart(msg)
		return err
	}
	switch state[0] {
	case "Auth":
		err := b.HandleAuth(&out)
		return err
	case "CreateNewWallet":
		err := b.HandleAuth(&out)
		return err
	case "CreateNewWalletWithoutPK":
		err := b.HandleAuth(&out)
		return err
	case "AddWalletWithPK":
		err := b.HandleAuth(&out)
		return err
	case "login":
		err := b.HandleLogin(&out)
		return err
	case "ChangeLockTime":
		err := b.ChangeLockTimeout(&out)
		return err
	case "addressBook":
		err := b.HandleAddressBookState(&out)
		return err
	case "logout":
		err := b.HandleLogoutState(&out)
		return err
	case "Transfer":
		err := b.HandleTransferState(&out)
		return err
	case "Delegate":
		err := b.HandleDelegateState(&out)
		return err
	case "Undelegate":
		err := b.HandleUndelegateState(&out)
		return err
	case "Swap":
		err := b.HandleSwapState(&out)
		return err
	case "ExportPK":
		err := b.HandleExportPKState(&out)
		return err
	case "newInvoice":
		err := b.HandleNewInvoiceState(&out)
		return err
	}
	return nil
}
