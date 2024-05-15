package eng

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	tele "gopkg.in/telebot.v3"
)

func GetWelcomeMsg() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		btnBalance               = selector.Data("💵 Balances", "balance")
		btnHistory               = selector.Data("📜 Show history", "history")
		btnYield                 = selector.Data("📈 Yield analyzer", "yield")
		btnNewDeposit            = selector.Data("💰 Deposit", "newDeposit")
		btnWithdrawBySwapService = selector.Data("🔄 Withdraw", "withdrawBySwap")
		btnNewTransfer           = selector.Data("💸 Transfer", "newTransfer")
		btnNewDelegate           = selector.Data("✅ Delegate", "newDelegate")
		btnNewUndelegate         = selector.Data("❎ Undelegate", "newUndelegate")
		btnAddressBook           = selector.Data("📖 Address book", "addressBook")
		btnInvoices              = selector.Data("🧾 Invoices", "invoices")
		btnLock                  = selector.Data("🔒 Lock the wallet", "Lock")
		btnSettings              = selector.Data("⚙ Settings", "settings")
	)
	selector.Inline(
		selector.Row(btnBalance, btnNewTransfer),
		selector.Row(btnNewDeposit, btnWithdrawBySwapService),
		selector.Row(btnNewDelegate, btnNewUndelegate),
		selector.Row(btnHistory),
		selector.Row(btnYield),
		selector.Row(btnAddressBook),
		selector.Row(btnInvoices),
		selector.Row(btnLock),
		selector.Row(btnSettings),
	)

	return `⭐ <b>Welcome to Casper Blockchain Telegram Wallet</b> ⭐

Hello and welcome! You've just unlocked a world of possibilities with our wallet bot. Here's what you can do:
	
💰 <b>Check Your Balances:</b> Instantly view all your wallet balances, so you're always in the know.
	
⭐ <b>Reward and transfer notifications:</b> Receive updates on rewards and incoming transfers as soon as they happen.
	
📧 <b>Send, Delegate, and Undelegate CSPR:</b> Seamlessly manage your CSPR. Send it to friends, delegate it to validators, or undelegate whenever you want.
	
📈 <b>Yield Analysis:</b> Dive deep into your yield and explore the potential of your assets with our analytical tools.
	
Your journey into the Casper Blockchain begins here. Feel free to explore and manage your assets with ease. If you have any questions or need assistance, don't hesitate to ask. Happy wallet management! 🚀💼🔒
	
Main menu
`, []interface{}{tele.ModeHTML, selector}
}

func GetMenuMsg() (interface{}, []interface{}) {
	var (
		menu        = &tele.ReplyMarkup{ResizeKeyboard: true}
		btnLockMenu = menu.Text("🔒 Lock")
		btnSettings = menu.Text("⚙ Settings")
	)
	menu.Reply(
		menu.Row(btnSettings),
		menu.Row(btnLockMenu),
	)

	return "", []interface{}{menu}
}

func GetWelcomeAuthMessage() (interface{}, []interface{}) {
	return "Welcome to the Casper Telegram Wallet\nTo start please enter your public key", []interface{}{tele.ModeMarkdownV2}
}

func GetAuthPasswordMessage() (interface{}, []interface{}) {
	return `🔒 <b>Set Encryption Password</b> 🔒

Your new password is crucial for encrypting your private key and safeguarding your wallet. Remember, without this password or your private key, there's no way to restore access. Keep it secure! 🔒✨`, []interface{}{tele.ModeHTML}
}

func GetAuthInvalidPubKeyMessage() (interface{}, []interface{}) {
	return `The public key you intered is invalid. Please enter a correct public key`, []interface{}{}
}

func GetAuthInvalidPrivateMessage() (interface{}, []interface{}) {
	return `The private key you intered is invalid. Please enter a correct private key`, []interface{}{}
}

func GetAuthRepeatPasswordMessage() (interface{}, []interface{}) {
	return `🔁 <b>Password Confirmation</b> 🔁
	Please re-enter the password you just provided. If it doesn't match, you'll need to start over with setting a new password. Double-check for accuracy! 🔐🔁`, []interface{}{tele.ModeHTML}
}

func GetAuthRepeatPasswordInvalidMessage() (interface{}, []interface{}) {
	return `Passwords you enter did not match. Please enter your new password`, []interface{}{}
}

func GetRegisterSuccessMessage(pubkey string) (interface{}, []interface{}) {
	return fmt.Sprintf(`🚀 <b>Registration Successful</b> 🚀

Your registration is a success! You're all set to use your wallet. Your unique wallet address is: <code>%s</code>. Enjoy secure transactions and asset management. ⭐💼🔒`, pubkey), []interface{}{tele.ModeHTML}
}

func GetSendPrivatKey(data []byte, pubkey string) (interface{}, []interface{}) {
	reader := bytes.NewReader(data)
	a := &tele.Document{File: tele.FromReader(reader), FileName: pubkey[:6] + "_pkey.pem",
		Caption: `🔐 IMPORTANT: Private Key Alert 🔐
		
The file attached to this message is your private key, a critical element in securing your wallet. Please take these steps seriously:
1️⃣ Save it Securely: Store this file in a highly secure location, away from prying eyes and potential threats.
2️⃣ Never Share it: Do not share this private key with anyone, under any circumstances. It's your key to safeguarding your assets.
3️⃣ Loss = No Access: Understand that if you lose this key, you'll lose access to your wallet permanently. There's no way for us to recover it.
4️⃣ Automatic Deletion: Be aware that this message will be automatically deleted in 24 hours with no way to recover it. So, please ensure you've saved your private key securely within this time frame.
You can use this private key to securely sign transactions by sending it to our bot whenever necessary. Your security is our top priority. Keep your key safe, and your assets will stay secure. 💼🔒
`}
	return a, []interface{}{tele.ModeHTML}
}

func GetLoadBalanceMessage() (interface{}, []interface{}) {
	return `⭐ Balances Loading... Please Hang Tight ⭐

We're working diligently to fetch your balances. Just a moment longer, and you'll have all your balances at your fingertips! ⏳💼✨`, []interface{}{tele.ModeHTML}
}

func GetBalanceMsg(balance types.BalanceResponse) (interface{}, []interface{}) {
	var (
		selector    = &tele.ReplyMarkup{}
		btnMain     = selector.Data("🏠 Back to main menu", "mainMenu")
		SwapService = selector.Data("🔄 Withdraw by SimpleSwap", "withdrawBySwap")
	)

	selector.Inline(selector.Row(SwapService), selector.Row(btnMain))
	header := `⭐ <b>Wallet Balances Info</b> ⭐`
	footer := `Your assets are looking fantastic! Keep managing them wisely. ⭐💼💸`

	// var DelegatedData string

	// for _, d := range balance.Delegated.GetData() {
	// 	DelegatedData += fmt.Sprintf("<b>%.6f CSPR</b> to <b>%s</b>\n\n", d.Amount, d.Validator)
	// }
	var totalDelegated float64
	for _, d := range balance.Delegated.GetData() {
		totalDelegated += d.Amount
	}

	var BeingDelegatedData string
	for _, d := range balance.BeingDelegated.GetData() {
		//log.Println(d.GetValidator())
		BeingDelegatedData += fmt.Sprintf("<b>%.6f CSPR</b> to <b>%s</b> Era when finished: %d ", d.Amount, d.Validator, d.EraDelegationFinished)
	}
	if len(balance.BeingDelegated.GetData()) == 0 {
		BeingDelegatedData += "<b>0.000000 CSPR</b>"
	}

	var BeingUndelegatedData string
	for _, d := range balance.BeingUndelegated.GetData() {
		BeingUndelegatedData += fmt.Sprintf("<b>%.6f CSPR</b> to <b>%s</b> Era when finished: %d ", d.Amount, d.Validator, d.EraUndelegationFinished)
	}
	if len(balance.BeingUndelegated.GetData()) == 0 {
		BeingUndelegatedData += "<b>0.000000 CSPR</b>"
	}

	balanceMsg := fmt.Sprintf("💰<b>Your liquid balance:</b> \n<b>%.6f CSPR</b>\n", balance.Liquid)
	priceMsg := fmt.Sprintf("📈<b>CSPR Token Price:</b> \n<b>%.6f USD</b>\n", balance.Price)
	totalMsg := fmt.Sprintf("🌐<b>Your total balance:</b> \n<b>%.6f CSPR (%.6f USD)</b>\n", balance.Total, balance.TotalUSD)
	delegatedMsg := fmt.Sprintf("🤝<b>Your delegated balance:</b>\n<b>%.6f CSPR</b>\n", totalDelegated)
	beingDelegatedMsg := fmt.Sprintf("💼<b>Your being delegated balance:</b> \n%s\n", BeingDelegatedData)
	beingUndelegatedMsg := fmt.Sprintf("🚀<b>Your being undelegated balance:</b> \n%s\n", BeingUndelegatedData)

	return fmt.Sprintf("%s\n\n%s\n%s\n%s\n%s\n%s\n%s\n%s", header, priceMsg, balanceMsg, delegatedMsg, beingDelegatedMsg, beingUndelegatedMsg, totalMsg, footer),
		[]interface{}{tele.ModeHTML, selector}
}

func GetErrorMessage() (interface{}, []interface{}) {
	return `*Error*
	Please contact support`, []interface{}{tele.ModeMarkdownV2}
}

func GetLoginMessage(LogoutManual bool) (interface{}, []interface{}) {
	var (
		selector  = &tele.ReplyMarkup{}
		btnLogout = selector.Data("⛔️ Reset the wallet", "logout")
	)

	selector.Inline(
		selector.Row(btnLogout),
	)
	var resp string
	ManualLogoutText := `🔒 <b>Logged Out</b>

You have been logged out for your security.

🔑 <b>Log In to Your Casper Telegram Wallet </b>

To access your wallet, please enter your password now.`
	AutoLogoutText := `🔒 Safeguarding Your Access 🔒

You've been automatically logged out due to the inactivity period you configured. To regain entry to your Casper Telegram Wallet, please input your password. We're here to keep your access secure!`
	if LogoutManual {
		resp = ManualLogoutText
	} else {
		resp = AutoLogoutText
	}
	return resp, []interface{}{tele.ModeHTML, selector}
}

func GetLoginSuccessMessage() (interface{}, []interface{}) {
	return `<b>Access Granted!</b> 😊

You have successfully logged in. 🎉`, []interface{}{tele.ModeHTML}
}

func GetLoginPassInvalidMessage() (interface{}, []interface{}) {
	var (
		selector  = &tele.ReplyMarkup{}
		btnLogout = selector.Data("⛔️ Reset the wallet", "logout")
	)

	selector.Inline(
		selector.Row(btnLogout),
	)
	return `❌ <b>Incorrect Password Entered</b>

The password you provided does not match our records.

🔑 <b>Please Try Again </b>

To access your wallet, re-enter your password correctly.`, []interface{}{tele.ModeHTML, selector}
}

func GetLockMessage() (interface{}, []interface{}) {
	return `🔐 <b>Your Wallet Is Now Securely Locked</b> 🔐

For access, please send the command /start to unlock your wallet.`, []interface{}{tele.ModeHTML}
}

func GetLogoutConfirmationMessage() (interface{}, []interface{}) {
	var (
		selector  = &tele.ReplyMarkup{}
		btnCancel = selector.Data("Cancel", "mainMenu")
	)

	selector.Inline(
		selector.Row(btnCancel),
	)
	return `🔒 <b>Avoid Accidental Reset</b> 🔒

To ensure you don't accidentally reset your Casper Telegram Wallet, please type "CONFIRM" to proceed or press "CANCEL" to keep your settings intact. We've got your back!`, []interface{}{tele.ModeHTML, selector}
}

func GetLogoutMessage() (interface{}, []interface{}) {
	return `✨ <b>Reset Complete!</b> ✨
	
Come back anytime, and use /start to create a new wallet when you're ready to begin your journey. We'll be here waiting for you! 🚀🔒`, []interface{}{tele.ModeHTML}
}

func GetSettingsMessage(settings types.SettingsResponse) (interface{}, []interface{}) {
	var (
		selector             = &tele.ReplyMarkup{}
		btnMain              = selector.Data("🏠 Back to main menu", "mainMenu")
		btnLogout            = selector.Data("⛔️ Reset the wallet", "logout")
		btnChangeLockTimeout = selector.Data("⌛ Set automatic lock timeout", "ChangeLockTimeout")
		//	btnOnOffNotify       = selector.Data("🔔 On/Off notify", "OnOffNotify")
		btnNotifySettings = selector.Data("🔔 Notification settings", "NotifySettings")
		btnSos            = selector.URL("🆘 Contact Support", "https://t.me/simplewallet_cspr")
		btnExportPrivat   = selector.Data("🔐 Export your private key", "ExportPrivat")
		//:ninja: Privacy options
		btnPrivacy = selector.Data("🕵️ Privacy options", "PrivacyOptions")
	)

	selector.Inline(
		selector.Row(btnNotifySettings),    //1
		selector.Row(btnPrivacy),           //2
		selector.Row(btnChangeLockTimeout), //3
		selector.Row(btnExportPrivat),      //4
		selector.Row(btnSos),               //5
		selector.Row(btnLogout),            //6
		selector.Row(btnMain),              //7

		//selector.Row(btnOnOffNotify),
	)

	var notificationSet string
	if settings.NotificationsEnabled {
		notificationSet = fmt.Sprintf("<b>Enabled (rewards every %d hours</b>)", settings.NotifyTime)
	} else {
		notificationSet = "<b>🚫 Currently Disabled</b>"
	}
	LockTimeoutInfo := fmt.Sprintf("%d", settings.LockTimeout)

	//return fmt.Sprintf("<b>Welcome to your Casper Telegram Wallet \n Settings</b>\n\nYour public key:\n<b>%s</b>\n\nNotifications:\n%s\n\nAutomatic lock timeout:\n%s", settings.PublicKey, notificationSet, LockTimeoutInfo), []interface{}{tele.ModeHTML, selector}
	return fmt.Sprintf(`⭐ <b>Welcome to Your Casper Telegram Wallet Settings</b> ⭐

🔑 <b>Your Public Key:</b> <code>%s</code>

🔔 <b>Notifications:</b> %s

⏲️ <b>Automatic Lock Timeout:</b> 🔐 Your wallet is set to automatically lock after %s minutes of inactivity for enhanced security.`, settings.PublicKey, notificationSet, LockTimeoutInfo), []interface{}{tele.ModeHTML, selector}
}

func GetNotifySettingsMessage(settings types.NotifySettingsResponse) (interface{}, []interface{}) {
	var notifyStatus string
	var btnOnOffText string
	if settings.NotificationsEnabled {
		notifyStatus = "✅ Enabled"
		btnOnOffText = "🔔 Turn operations notification OFF"
	} else {
		btnOnOffText = "🔔 Turn operations notification ON"
		notifyStatus = "🚫 Currently Disabled"
	}
	var (
		selector                   = &tele.ReplyMarkup{}
		btnMain                    = selector.Data("🏠 Back to main menu", "mainMenu")
		btnOnOffNotify             = selector.Data(btnOnOffText, "OnOffNotify")
		btnChangeRewardsNotifyTime = selector.Data("⌛ Change rewards notification period", "ChangeRewardsNotifyTime")
	)

	selector.Inline(
		selector.Row(btnOnOffNotify),
		selector.Row(btnChangeRewardsNotifyTime),
		selector.Row(btnMain),
	)

	var NotifyTime string

	if settings.NotyfyTime == 0 {
		NotifyTime = "🚫 Currently Disabled"
	} else {
		NotifyTime = "⏰ Every " + strconv.Itoa(int(settings.NotyfyTime)) + " hours"
	}

	return fmt.Sprintf(`🔔 <b>Notification Settings Overview</b> 🔔

<b>Operations Notification Status:</b> %s

<b>Rewards Notification Period:</b> %s`, notifyStatus, NotifyTime), []interface{}{tele.ModeHTML, selector}

}

func GetChangeLockTimeoutMessage(currentTime int64) (interface{}, []interface{}) {
	var (
		selector  = &tele.ReplyMarkup{}
		btnCancel = selector.Data("Cancel", "cancelChangeTimeout")
	)

	selector.Inline(
		selector.Row(btnCancel),
	)
	text := fmt.Sprintf(`🔐 <b>Enhance Your Wallet's Security</b>

Please configure the automatic lock timeout for your wallet. This feature is designed to boost your security by automatically securing your wallet after a period of inactivity.

⏱️ <b>Current Setting:</b> Your wallet is currently set to automatically lock after %d minutes of inactivity.

🔒💼 By customizing this setting, you ensure enhanced protection for your assets, maintaining peace of mind and control over your wallet's security. Set your desired inactivity period in minutes now.`, currentTime)
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetChangeLockTimeoutMessageSuccess(current int64) (interface{}, []interface{}) {
	var selector = &tele.ReplyMarkup{}
	var btnMain = selector.Data("🏠 Back to main menu", "mainMenu")
	selector.Inline(
		selector.Row(btnMain),
	)
	return fmt.Sprintf("You have successfully set automatic timeout time to <b>%d minutes</b>", current), []interface{}{tele.ModeHTML, selector}
}

func GetHistoryMenu() (interface{}, []interface{}) {
	var (
		selector          = &tele.ReplyMarkup{}
		btnBalanceHistory = selector.Data("Show balance history", "balanceHistory")
		btnRewardsHisory  = selector.Data("Show rewards history", "rewardsHistory")
		btnMain           = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(
		selector.Row(btnBalanceHistory),
		selector.Row(btnRewardsHisory),
		selector.Row(btnMain),
	)

	return `<b>⏳ Operations History Menu ⏳</b>

Welcome to the Operations History section of your wallet. This feature allows for a comprehensive review of your financial transactions. Please select from the following options to access detailed historical data:

📊 <b>Balance History:</b> View a chronological record of your wallet's balance fluctuations. This feature provides insights into your transactional history, enabling you to trace the evolution of your funds over time.

🎁 <b>Rewards History:</b> Explore a detailed account of the rewards accrued. This section offers a thorough overview of your earnings, illustrating the growth of your assets.

Each selection offers a detailed perspective on different aspects of your wallet's history. Utilize these tools to gain a clearer understanding of your financial trajectory. 🔍📜
`, []interface{}{tele.ModeHTML, selector}
}
func GetLoadHistoryMessage() (interface{}, []interface{}) {
	return `⌛ <b>Loading History... Patience Is Key</b> ⌛
	
🛠️ We're meticulously gathering your historical data. Hold on just a little longer, and soon you'll have a complete overview of your operations at your command! 📊🔍✨`, []interface{}{tele.ModeHTML}
}

// func GetHistoryMsg(history types.HistoryResponse) (interface{}, []interface{}) {
// 	var (
// 		selector = &tele.ReplyMarkup{}
// 		btnNext  = selector.Data("Next", "moveHistory", strconv.Itoa(int(history.Start+500)))
// 		btnPrev  = selector.Data("Prev", "moveHistory", strconv.Itoa(int(history.Start-500)))
// 		btnMain  = selector.Data("🏠 Main Menu", "mainMenu")
// 	)

// 	if history.IsFirst {
// 		selector.Inline(selector.Row(btnPrev), selector.Row(btnMain))
// 	} else {
// 		selector.Inline(selector.Row(btnPrev, btnNext), selector.Row(btnMain))
// 	}

// 	var Transfers string
// 	Transfers += "💱 Transfer history: \n"
// 	for _, t := range history.Transfers {
// 		Transfers += fmt.Sprintf("Transfer \n<b>%.6f CSPR</b> from <b>%s</b> to <b>%s</b> \n block height= %d \n timestamp = %s \n\n", t.GetAmount(), t.GetFrom(), t.GetTo(), t.GetHeight(), t.GetDate())
// 	}
// 	var Delegates string
// 	Delegates += "✅ Delegate history: \n"
// 	for _, d := range history.DelegateHistory {
// 		Delegates += fmt.Sprintf("Delegate \n<b>%.6f CSPR</b> to validator <b>%s</b> in era <b>%d</b>. Finished? <b>%v</b>\n block height= %d \n timestamp = %s \n\n", d.GetAmount(), d.GetValidator(), d.GetEra(), d.GetIsFinished(), d.GetHeight(), d.GetDate())
// 	}
// 	var Undelegates string
// 	Undelegates += "❎ Undelegate history: \n"
// 	for _, u := range history.UndelegateHistory {
// 		Undelegates += fmt.Sprintf("Undelegate \n<b>%.6f CSPR</b> from validator <b>%s</b> in era <b>%d</b>. Finished? <b>%v</b> \n block height= %d \n timestamp = %s \n\n", u.GetAmount(), u.GetValidator(), u.GetEra(), u.GetIsFinished(), u.GetHeight(), u.GetDate())
// 	}

// 	debugInfo := fmt.Sprintf("<b>DEBUG INFO</b> \nBlock Start: %d \nBlock End: %d \n", history.Start, history.End)
// 	resp := fmt.Sprintf("%s%s%s%s", Transfers, Delegates, Undelegates, debugInfo)
// 	return resp, []interface{}{tele.ModeHTML, selector}
// }

func GetHistoryMsg(history types.HistoryResponse) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnNext  = selector.Data("Next page", "moveHistory", strconv.Itoa(int(history.Start+500)))
		btnPrev  = selector.Data("Previous page", "moveHistory", strconv.Itoa(int(history.Start-500)))
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)

	if history.IsFirst {
		selector.Inline(selector.Row(btnPrev), selector.Row(btnMain))
	} else {
		selector.Inline(selector.Row(btnPrev, btnNext), selector.Row(btnMain))
	}
	events := make(map[int64]string)
	for _, t := range history.Transfers {
		if t.Outward {
			events[t.Height] += fmt.Sprintf("<b>%s</b>\nOutward transfer: \n%.6f CSPR to %s \n\n", t.GetDate(), t.GetAmount(), t.GetTo())
		} else {
			events[t.Height] += fmt.Sprintf("<b>%s</b>\nInward transfer: \n%.6f CSPR from %s \n\n", t.GetDate(), t.GetAmount(), t.GetFrom())
		}

	}
	for _, d := range history.DelegateHistory {
		//events[d.Height] += fmt.Sprintf("<b>%s</b>\nDelegate \n<b>%.6f CSPR</b> to validator <b>%s</b> in era <b>%d</b>. Finished? <b>%v</b>\n\n", d.GetDate(), d.GetAmount(), d.GetValidator(), d.GetEra(), d.GetIsFinished())
		events[d.Height] += fmt.Sprintf("<b>%s</b>\nDelegation: \n%.6f CSPR to validator %s \n\n", d.GetDate(), d.GetAmount(), d.GetValidator())

	}
	for _, u := range history.UndelegateHistory {
		events[u.Height] += fmt.Sprintf("<b>%s</b>\nUndelegate \n%.6f CSPR from validator %s\n\n", u.GetDate(), u.GetAmount(), u.GetValidator())
	}
	var resp string
	debugInfo := fmt.Sprintf("🕒 Operations History \n\n<b>Period:</b> \n<b>⏰ From:</b> %s \n<b>⏰ To:</b> %s \n\n<b>🔍Result:</b> \n", history.StartDate, history.EndDate)
	resp += debugInfo + "\n"
	for _, text := range events {
		//block_string := strconv.Itoa(int(block))
		//resp += "<b>block " + block_string + "</b> \n\n"
		resp += text
	}
	if len(events) == 0 {
		resp += "🚫 No operations found"
	}
	return resp, []interface{}{tele.ModeHTML, selector}
}

func GetRewardsHistoryMsg(history types.RewardsHistoryResponse) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnNext  = selector.Data("Next page", "moveRewardsHistory", strconv.Itoa(int(history.Start+10)))
		btnPrev  = selector.Data("Previous page", "moveRewardsHistory", strconv.Itoa(int(history.Start-10)))
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	if history.IsFirst {
		selector.Inline(selector.Row(btnPrev), selector.Row(btnMain))
	} else {
		selector.Inline(selector.Row(btnPrev, btnNext), selector.Row(btnMain))
	}

	var Rewards string
	Rewards += fmt.Sprintf("🏆 Rewards History \n\n<b>Period:</b> \n<b>⏰ From:</b> %s \n<b>⏰ To:</b> %s \n\n<b>🔍Result:</b>\n", history.StartDate, history.EndDate)
	for _, r := range history.Rewards {
		Rewards += fmt.Sprintf("<b>%s</b>\nReward %.6f CSPR from validator %s (era %d)\n\n", r.GetTimestamp(), r.GetAmount(), "*"+r.GetValidator()[:6], r.GetEra())
	}
	if len(history.Rewards) == 0 {
		Rewards += "🚫 No operations found"
	}

	//debugInfo := fmt.Sprintf("<b>DEBUG INFO</b> \nStart Era: %d \nEnd Era: %d \n", history.Start, history.End)
	resp := Rewards
	return resp, []interface{}{tele.ModeHTML, selector}
}

func GetYieldMsg(history types.YieldResponse) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(selector.Row(btnMain))

	rewards := fmt.Sprintf("🏆<b>Past rewards:</b> Total rewards for the last 7 days: <b>%.6f CSPR (%.6f USD)</b>\n\n", history.TotalRewards, history.RewardsUSD)

	// var sb strings.Builder
	// sb.WriteString("💰 Estimated Yield : \n")
	// for _, estim := range history.Estimates {
	// 	sb.WriteString(fmt.Sprintf("Validator <b>%s</b>: \n<b>%.6f CSPR</b>\n\n", estim.Validator, estim.Amount))
	// }

	prev := "⭐ <b>Casper Network Earnings Overview </b>⭐"

	var validators strings.Builder
	if len(history.Validators) != 0 {
		validators.WriteString("💵 <b>Your delegated balance:</b> \n\n")
		for _, estim := range history.Validators {
			validators.WriteString(fmt.Sprintf("Validator <b>%s</b>: \n<b>%.6f CSPR</b> (%d %s commission, Validator APY: %.6f %%)\n", "*"+estim.Address[:5], estim.Amount, int(estim.Fee), "%", estim.Apy))
		}
		validators.WriteString(fmt.Sprintf("\nTotal delegated balance: %.6f CSPR \n\n", history.TotalDelegated))
	} else {
		validators.WriteString("💵 <b>Your delegated balance:</b> ")
		validators.WriteString(fmt.Sprintf("\nTotal delegated balance: %.6f CSPR \n\n", history.TotalDelegated))
	}

	var Projected string

	Projected = "💰 <b>Projected rewards</b> \n\n"
	Projected += fmt.Sprintf("-- 30 days rewards: %.6f CSPR (%.6f USD)\n", history.Proj30Days, history.Proj30DaysUSD)
	Projected += fmt.Sprintf("-- 365 days rewards: %.6f CSPR (%.6f USD)\n\n", history.Proj365Days, history.Proj365DaysUSD)
	if math.IsNaN(history.ProjApy) {
		Projected += "📊 <b>Your projected APY:</b> No rewards data available \n"
	} else {
		Projected += fmt.Sprintf("📊 <b>Your projected APY:</b> %.6f %% \n", history.ProjApy)
	}

	apy := fmt.Sprintf("📈 <b>Current Casper network APY:</b> %.6f %%\n\n", history.NetworkApy)

	resp := prev + "\n\n" + apy + validators.String() + rewards + Projected
	//log.Println("YIELD MESSAGE:", resp)
	return resp, []interface{}{tele.ModeHTML, selector}
}

func GetCustomYieldMsg() (interface{}, []interface{}) {
	return `⏳ <b>Gathering Yield Data... A Moment of Patience, Please</b> ⏳

⭐ We're currently processing your yield data. Hang tight just a bit longer, and you'll soon have all the insights you need at your fingertips! 📈 💫`, []interface{}{tele.ModeHTML}
}

func GetTooManyTasksMsg() (interface{}, []interface{}) {
	return "*Too many tasks, please wait*", []interface{}{tele.ModeMarkdownV2}
}

//GetNotifyNewUndelegateMessage
//GetNotifyNewDelegateMessage
//GetNotifyNewTransferMessage

func GetNotifyNewTransferMessage(amount float64, from string, to string, balance float64) (interface{}, []interface{}) {
	return fmt.Sprintf("📬 You have received <b>%.6f CSPR</b> from <b>%s</b> to <b>%s</b> \n New balance: <b>%.6f CSPR</b>", amount, from, to, balance), []interface{}{tele.ModeHTML}
}

func GetNotifyNewDelegateMessage(amount float64, validator string, era int64, balance float64) (interface{}, []interface{}) {
	return fmt.Sprintf("📬 You have delegated <b>%.6f CSPR</b> to validator <b>%s</b> in era <b>%d</b>  \n New balance: <b>%.6f CSPR</b>", amount, validator, era, balance), []interface{}{tele.ModeHTML}
}

func GetNotifyNewUndelegateMessage(amount float64, validator string, era int64, balance float64) (interface{}, []interface{}) {
	return fmt.Sprintf("📬 You have undelegated <b>%.6f CSPR</b> from validator <b>%s</b> in era <b>%d</b>  \n New balance: <b>%.6f CSPR</b>", amount, validator, era, balance), []interface{}{tele.ModeHTML}
}

func GetNotifyNewRewards(rews types.NotifyNewRewards) (interface{}, []interface{}) {
	var rewards string
	for _, r := range rews.Rewards {
		rewards += fmt.Sprintf("<b>%.6f CSPR</b> from validator <b>*%s</b>\n\n", r.Amount, r.Validator[:6])
	}
	text := ""
	if rews.LastEra == rews.FirstEra {
		text = fmt.Sprintf(`🏆 <b>Casper Rewards Notification</b>
<b>Period:</b> %s
<b>Era:</b> %d
✨ <b>Rewards Earned: </b>
%s 
💰 <b>Current Delegated Balance:</b> %.6f CSPR`, rews.FirstEraTimestamp, rews.FirstEra, rewards, rews.Delegated)
	} else {
		text = fmt.Sprintf(`🏆 <b>Casper Rewards Notification</b>
<b>Period:</b> %s to %s 
<b>Era range:</b> %d-%d
✨ <b>Rewards Earned: </b>
%s 
💰 <b>Current Delegated Balance:</b> %.6f CSPR`, rews.FirstEraTimestamp, rews.LastEraTimestamp, rews.FirstEra, rews.LastEra, rewards, rews.Delegated)
	}
	return text, []interface{}{tele.ModeHTML}
}
func GetNotifyNewBalance(amount float64, old float64) (interface{}, []interface{}) {
	text := fmt.Sprintf(`
✅ <b>Liquid Balance Updated</b>
	
🔄 <b>New Balance:</b> %f CSPR
🔙 <b>Previous Balance:</b> %f CSPR
`, amount, old)

	if amount > old {
		text += fmt.Sprintf(`
		📈 An increase of %f CSPR reflects recent activity in your wallet`, amount-old)
	} else {
		text += fmt.Sprintf(`
		📉 An decrease of %f CSPR reflects recent activity in your wallet`, old-amount)
	}
	return text, []interface{}{tele.ModeHTML}
}

func GetAddressBookMsg(book types.AddressBookResponse) (interface{}, []interface{}) {
	var (
		selector       = &tele.ReplyMarkup{}
		btnMain        = selector.Data("🏠 Back to main menu", "mainMenu")
		btnCreateEntry = selector.Data("📝 Create new entry", "createEntryAdressBook")
		btnNext        = selector.Data("Next page", "moveAddressBook", strconv.Itoa(int(book.Offset+5)))
		btnPrev        = selector.Data("Previous page", "moveAddressBook", strconv.Itoa(int(book.Offset-5)))
	)

	var namedButtons []tele.Row

	for _, boobook := range book.Data {
		namedButtons = append(namedButtons, selector.Row(selector.Data(boobook.GetName(), "showAddress", strconv.Itoa(int(boobook.GetId())))))
	}
	//namedButtons = append(namedButtons, selector.Row(btnMain))
	if book.Total > 5 {
		tempSelector := selector.Row()
		if book.Offset > 0 {
			tempSelector = append(tempSelector, btnPrev)
		}
		if book.Offset+5 < book.Total {
			tempSelector = append(tempSelector, btnNext)
		}
		namedButtons = append(namedButtons, tempSelector)
	}
	namedButtons = append(namedButtons, selector.Row(btnCreateEntry))
	namedButtons = append(namedButtons, selector.Row(btnMain))

	selector.Inline(namedButtons...)

	text := "<b>Welcome to your Casper Telegram Wallet</b>\n\n📖 <b>Your Address Book</b> \n"
	if len(book.Data) == 0 {
		text = "<b>Welcome to your Casper Telegram Wallet</b>\n\n📖 <b>Your Address Book</b> \n\n Currently, your address book is empty. Let's begin populating it with addresses for easier access and convenience in the future."
	}

	return text, []interface{}{tele.ModeHTML, selector}

}

func GetCreateEntryAddressBookNameMsg() (interface{}, []interface{}) {
	var (
		selector  = &tele.ReplyMarkup{}
		btnCancel = selector.Data("Cancel", "cancelAddressBook")
	)
	selector.Inline(
		selector.Row(btnCancel),
	)
	return `📓 Create New Address Book Entry
	
	Please enter a name for this new addition to your address book.
	`, []interface{}{tele.ModeHTML, selector}
}

func GetAskAddresAdressBookMsg(name string) (interface{}, []interface{}) {
	var (
		selector  = &tele.ReplyMarkup{}
		btnCancel = selector.Data("Cancel", "cancelAddressBook")
	)
	selector.Inline(
		selector.Row(btnCancel),
	)
	return fmt.Sprintf(`🌐 Address for Entry [%s]
	Please enter the Casper blockchain address you wish to add to your address book.`, name), []interface{}{tele.ModeHTML, selector}
}

func GetAskAddresInvalidAdress() (interface{}, []interface{}) {
	var (
		selector  = &tele.ReplyMarkup{}
		btnCancel = selector.Data("Cancel", "cancelAddressBook")
	)
	selector.Inline(
		selector.Row(btnCancel),
	)
	return `❗ Invalid Address Detected
	
	Please check the address and try entering it again.`, []interface{}{tele.ModeHTML, selector}
}

func GetAddressDetailedMsg(data types.AddressBookDetailed) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}

		btnChangeName    = selector.Data("📝 Change name", "changeNameAddressBook", strconv.Itoa(int(data.Id)))
		btnChangeAddress = selector.Data("📝 Change address", "changeAddressAddressBook", strconv.Itoa(int(data.Id)))
		btnDelete        = selector.Data("🗑 Delete this entry", "deleteAddressBook", strconv.Itoa(int(data.Id)))
		btnAddressBook   = selector.Data("📖 Back", "addressBook")
	)

	selector.Inline(
		selector.Row(btnChangeName),
		selector.Row(btnChangeAddress),
		selector.Row(btnDelete),
		selector.Row(btnAddressBook))
	msg := fmt.Sprintf("👤<b>Name:</b> %s \n\n📍<b>Address:</b> <code>%s</code> \n\n⏰<b>Created at:</b> %s", data.Name, data.Address, data.CreatedAt.Format("2006.01.02 15:04:05"))
	return msg, []interface{}{tele.ModeHTML, selector}
}

func GetDeleteEntryAddressBookConfirmationMessage(name, address string) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}

		btnConfirm     = selector.Data("📝 Confirm", "ConfirmDeleteAdressBook")
		btnAddressBook = selector.Data("✅ Back", "addressBook")
	)

	selector.Inline(
		selector.Row(btnConfirm),
		selector.Row(btnAddressBook))
	msg := fmt.Sprintf(`
❓ <b>Confirm Deletion</b>

Are you absolutely certain you want to delete this address book entry:

👤 <b>Name:</b> %s

📍 <b>Address:</b> <code>%s</code>`, name, address)
	return msg, []interface{}{tele.ModeHTML, selector}
}

func GetChangeAuthTypeMessage() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		NewWallet      = selector.Data("Create new wallet", "createWallet")
		ExistingWallet = selector.Data("Add existing wallet", "addExistingWallet")
	)
	selector.Inline(
		selector.Row(NewWallet),
		selector.Row(ExistingWallet),
	)
	return `⭐ <b>Welcome to Casper Blockchain Telegram Wallet</b> ⭐

Hello and welcome! You've just unlocked a world of possibilities with our wallet bot. Here's what you can do:

💰 <b>Check Your Balances:</b> Instantly view all your wallet balances, so you're always in the know.

⭐ <b>Reward and transfer notifications:</b> Receive updates on rewards and incoming transfers as soon as they happen.

📧 <b>Send, Delegate, and Undelegate CSPR:</b> Seamlessly manage your CSPR. Send it to friends, delegate it to validators, or undelegate whenever you want.

📈 <b>Yield Analysis:</b> Dive deep into your yield and explore the potential of your assets with our analytical tools.

Your journey into the Casper Blockchain begins here. Feel free to explore and manage your assets with ease. If you have any questions or need assistance, don't hesitate to ask. Happy wallet management! 🚀💼🔒`, []interface{}{tele.ModeHTML, selector}
}

func GetAskStoreTheKeyMessage() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		Store    = selector.Data("Store the key", "StoreKey")
		NotStore = selector.Data("Do not store the key", "NotStoreKey")
	)
	selector.Inline(
		selector.Row(Store),
		selector.Row(NotStore),
	)
	return `We've got some exciting options for you:

1️⃣Store your Private Key 🛡️: Opt for convenience! We'll securely encrypt your private key, keeping it safe and accessible only to you. No worries, we won't have access to it.

2️⃣ Don't Store Your Private Key 🚫: Prefer an extra layer of security? Choose this option, but be prepared to provide your private key each time you want to make a transaction. It's less convenient but offers peace of mind.

The choice is yours! Please select the option that suits your needs best. Your privacy and security are our priorities. 🔒✨`, []interface{}{tele.ModeHTML, selector}
}

func GetAskPrivatKeyMessage() (interface{}, []interface{}) {
	return `🔑 <b>Share Your Private Key</b> 🔑


Feel free to send us either the complete .pem file or just the private key inside it. We're all about making it convenient for you! 📤🔐`, []interface{}{tele.ModeHTML}
}

func GetSendTransferStage1Message(balance string) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		fromAddressBook = selector.Data("Select from address book", "TransferAddressBook")
		enterCustom     = selector.Data("Enter address manually", "TransferCustomAddress")
		btnMain         = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(
		selector.Row(fromAddressBook),
		selector.Row(enterCustom),
		selector.Row(btnMain),
	)
	text := fmt.Sprintf(`💧 <b>Your Liquid Balance:</b> %s CSPR
	
🚀 <b>Send CSPR</b>:
Would you like to send CSPR to an address from your address book, or would you prefer to enter the address manually? Please select your preferred method.`, balance)
	return text, []interface{}{tele.ModeHTML, selector}
}

func SendTransferAskAdressMessage() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		btnMain = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(
		selector.Row(btnMain),
	)
	text := `📬 <b>Select CSPR Transfer Destination</b>
	
Please enter the destination address for your CSPR transfer`
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetTransferAskAmountMessage(balance float64, toPubkey string) (interface{}, []interface{}) {
	text := fmt.Sprintf(`
💸 <b>Set Your Transfer Amount</b>

Ready to transfer CSPR? Go ahead and select how much you'd like to send to the address: <code>%s</code>

💼 <b>Your Liquid Balance:</b> %f CSPR

💡 <b>Tip:</b> It's wise to keep at least 10 CSPR in your balance for gas fees.

🔢 Use the button to pick a preset amount or enter a custom value. The choice is yours!`, toPubkey, balance+10)

	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		sendMaximum = selector.Data(fmt.Sprintf("Send maximum (%.6f CSPR)", balance), "TransferMaximum")
		btnMain     = selector.Data("🏠 Cancel", "mainMenu")
	)

	if balance > 0 {
		selector.Inline(
			selector.Row(sendMaximum),
			selector.Row(btnMain),
		)
	} else {
		selector.Inline(
			selector.Row(btnMain),
		)
	}
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetTransferAskMemo(amount string) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		withoutMemo = selector.Data("Continue without memo", "transferWithoutMemo")
		btnMain     = selector.Data("🏠 Cancel", "mainMenu")
	)
	selector.Inline(
		selector.Row(withoutMemo),
		selector.Row(btnMain),
	)
	text := "You are transferring <b>" + amount + " CSPR</b> \n\nPlease input the memo / id / tag for this transfer or press Continue without memo "
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetTransferAskConfirmation(amount, topubkey, name string) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		//withoutMemo = selector.Data("Continue without memo", "transferWithoutMemo")
		confirm = selector.Data("Confirm", "transferConfirm")
		cancel  = selector.Data("Cancel", "mainMenu")
	)
	selector.Inline(
		selector.Row(confirm),
		selector.Row(cancel),
	)
	text := fmt.Sprintf("✅ Final Confirmation ✅\n\n You are transfering <b>%s CSPR</b> to %s (...%s) \n\nPlease confirm the transfer or press Cancel. 🚀😳", amount, name, topubkey[len(topubkey)-5:])
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetSignDeployAskPasswordMessage() (interface{}, []interface{}) {
	return "Please enter your password to sign deploy", []interface{}{tele.ModeMarkdownV2}
}

func GetExportAskPasswordMessage() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		cancel   = selector.Data("Cancel", "mainMenu")
	)
	selector.Inline(
		selector.Row(cancel),
	)
	return `🔑 Export Private Key - Password Required 🔑

To proceed with the export of your private key, your password is needed for verification. This step ensures the security of your key and wallet.

🔒 Please Enter Your Password

Enter your password carefully to successfully export your private key. We prioritize the safety and confidentiality of your wallet details!`, []interface{}{tele.ModeHTML}
}

func GetExportIncorrectPasswordMessage() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		cancel   = selector.Data("Cancel", "mainMenu")
	)
	selector.Inline(
		selector.Row(cancel),
	)
	return `❌ Incorrect Password Entered

The password you provided does not match our records.

🔑 Please Try Again

To export the private key, re-enter your password correctly.`, []interface{}{tele.ModeHTML}
}

func GetSignDeployAskPK() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		cancel   = selector.Data("Cancel", "mainMenu")
	)
	selector.Inline(
		selector.Row(cancel),
	)
	return `🔐 <b>Operation Confirmation Required</b>

Since you opted not to store your private key with us, it's necessary to provide it now to confirm your operation.
🔑 Please securely send your private key either as a .pem file or in text format.
`, []interface{}{tele.ModeHTML, selector}
}

func GetSuccessTransferMessage(amount string, toPubkey string, toName string, memo uint64, hash string) (interface{}, []interface{}) {
	text := fmt.Sprintf(`You transaction has been succesfull submitted!
	
	Transfering %s CSPR to %s (...%s) with tag %d.
	
	Transaction link https://testnet.cspr.live/deploy/%s`, amount, toName, toPubkey[len(toPubkey)-5:], memo, hash)
	return text, []interface{}{tele.ModeHTML}
}

func GetTransferAddressIsNotValidMsg() (interface{}, []interface{}) {
	return `❗ <b>Invalid Address Detected</b>

	The address you entered does not seem to be valid.

	🔄 Please double-check it and try entering it again.
	`, []interface{}{tele.ModeHTML}
}

func GetTransferAmountIsNotValidMsg() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Row(btnMain)
	return `⚠️ Invalid Amount Entered ⚠️

The amount entered must be a number greater than 0. Please ensure you input a valid amount.

<b>💡 Enter Valid Amount</b>

Enter a positive numerical value to accurately represent the transaction amount for your invoice.
We appreciate your attention to detail and accuracy. Thank you for using Casper Telegram Wallet for your invoicing needs`, []interface{}{tele.ModeHTML, selector}
}

func GetTransferMemoIsNotValidMsg() (interface{}, []interface{}) {
	return "The memo you entered is not valid, please enter number 0-18446744073709551615", []interface{}{tele.ModeHTML}
}

func GetTransferAddressBookMsg(book types.AddressBookResponse) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
		btnNext  = selector.Data("Next page", "moveTransferAddressBook", strconv.Itoa(int(book.Offset+5)))
		btnPrev  = selector.Data("Previous page", "moveTransferAddressBook", strconv.Itoa(int(book.Offset-5)))
	)

	var namedButtons []tele.Row

	for _, boobook := range book.Data {
		//namedButtons = append(namedButtons, selector.Row(selector.Data(boobook.GetName(), "showAddress", strconv.Itoa(int(boobook.GetId())))))
		namedButtons = append(namedButtons, selector.Row(selector.Data(boobook.GetName(), "pickTransferAddress", strconv.Itoa(int(boobook.GetId())))))
	}
	//namedButtons = append(namedButtons, selector.Row(btnMain))
	if book.Total > 5 {
		tempSelector := selector.Row()
		if book.Offset > 0 {
			tempSelector = append(tempSelector, btnPrev)
		}
		if book.Offset+5 < book.Total {
			tempSelector = append(tempSelector, btnNext)
		}
		namedButtons = append(namedButtons, tempSelector)
	}
	namedButtons = append(namedButtons, selector.Row(btnMain))

	selector.Inline(namedButtons...)

	return "*Welcome to your Casper Telegram Wallet*\n\n📖 Address book \n", []interface{}{tele.ModeHTML, selector}

}

func GetTransferUnknownError() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		cancel   = selector.Data("🏠 Return to main menu", "mainMenu")
	)
	selector.Inline(
		selector.Row(cancel),
	)
	return `⚠️ <b>Unknown Error Encountered During Transfer</b>

An unexpected issue has occurred.
🔄 Please attempt the transfer again. If the problem persists, contact support for assistance.
`, []interface{}{tele.ModeHTML}
}

func GetTransferBadPassword() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		cancel   = selector.Data("Cancel", "mainMenu")
	)
	selector.Inline(
		selector.Row(cancel),
	)
	return "incorrect password, try again", []interface{}{tele.ModeHTML, selector}
}

func GetTransferBadPK() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		cancel   = selector.Data("Cancel", "mainMenu")
	)
	selector.Inline(
		selector.Row(cancel),
	)
	return `❌ <b>Incorrect Private Key Detected</b>

The key you provided does not match our records.
🔄 Please try again. You can resend the private key either as a .pem file or in text format.
`, []interface{}{tele.ModeHTML, selector}
}

func GetDelegatorValidators(book types.DelegateValidatorsResponse) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
		btnNext  = selector.Data("Next page ➡️", "moveDelegateValidators", strconv.Itoa(int(book.Offset+5)))
		btnPrev  = selector.Data("⬅️ Previous page", "moveDelegateValidators", strconv.Itoa(int(book.Offset-5)))
	)

	var namedButtons []tele.Row

	for _, boobook := range book.Data {
		//namedButtons = append(namedButtons, selector.Row(selector.Data(boobook.GetName(), "showAddress", strconv.Itoa(int(boobook.GetId())))))
		namedButtons = append(namedButtons, selector.Row(selector.Data(fmt.Sprintf("Validator: *%s, Fee: %d %%", boobook.GetAddress()[len(boobook.GetAddress())-5:], int(boobook.GetFee())),
			"pickDelegateValidator", strconv.Itoa(int(boobook.GetId())))))
	}
	//namedButtons = append(namedButtons, selector.Row(btnMain))
	if book.Total > 5 {
		tempSelector := selector.Row()
		if book.Offset > 0 {
			tempSelector = append(tempSelector, btnPrev)
		}
		if book.Offset+5 < book.Total {
			tempSelector = append(tempSelector, btnNext)
		}
		namedButtons = append(namedButtons, tempSelector)
	}
	namedButtons = append(namedButtons, selector.Row(btnMain))

	selector.Inline(namedButtons...)

	text := fmt.Sprintf(`<b>💧 Your Liquid Balance:</b> %s
	
⭐ <b>Choose a Validator to Stake With:</b>
Explore and select from the network of validators to delegate your CSPR.

🔢 <b>Total Validators:</b> %d 

📖 <b>Currently Viewing Validators:</b> %d-%d

⬅️➡️ Use the buttons to navigate through the validator list.
	`, book.UserBalance, book.Total, book.Offset, book.Offset+5)

	return text, []interface{}{tele.ModeHTML, selector}
}

func GetDelegateAskAmountMessage(balance float64, validator string) (interface{}, []interface{}) {
	text := fmt.Sprintf(`<b>Please select amount CSPR to stake to %s </b> 
	
	recomended maximum leaves 10 CSPR in your liquid balance for the gas 
	
	
	You can choose amount by button or enter custom amount`, validator)

	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		balance25  = balance / 100 * 25
		balance50  = balance / 100 * 55
		balance75  = balance / 100 * 75
		balanceRec = balance - 10
		perc25     = selector.Data(fmt.Sprintf("Send 25%% (%.6f CSPR)", balance25), "DelegateSelectAmount", fmt.Sprintf("%f", balance25))
		perc50     = selector.Data(fmt.Sprintf("Send 55%% (%.6f CSPR)", balance50), "DelegateSelectAmount", fmt.Sprintf("%f", balance50))
		perc75     = selector.Data(fmt.Sprintf("Send 75%% (%.6f CSPR)", balance75), "DelegateSelectAmount", fmt.Sprintf("%f", balance75))
		percRec    = selector.Data(fmt.Sprintf("Send reccomended maximum (%.6f CSPR)", balanceRec), "DelegateSelectAmount", fmt.Sprintf("%f", balanceRec))
		mainMenu   = selector.Data("Main menu", "mainMenu")
	)

	if balance25 >= 500 {
		selector.Inline(
			selector.Row(perc25),
			selector.Row(perc50),
			selector.Row(perc75),
			selector.Row(percRec),
		)
	} else if balance50 >= 500 {
		selector.Inline(
			selector.Row(perc50),
			selector.Row(perc75),
			selector.Row(percRec),
		)
	} else if balance75 >= 500 {
		selector.Inline(
			selector.Row(perc75),
			selector.Row(percRec),
		)
	} else if balanceRec >= 500 {
		selector.Inline(
			selector.Row(percRec),
		)
	} else {
		text = fmt.Sprintf(`🚫 <b>Insufficient Balance for Delegation</b>
Your current balance: %f

💡 <b>Minimum Required to Delegate:</b> 500 CSPR
Please top up your balance and try again.`, balance)
		selector.Inline(
			selector.Row(mainMenu),
		)
	}

	return text, []interface{}{tele.ModeHTML, selector}
}

func GetDelegateAskConfirmation(amount float64, validator string) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		//withoutMemo = selector.Data("Continue without memo", "transferWithoutMemo")
		confirm = selector.Data("Confirm", "delegateConfirm")
		cancel  = selector.Data("Cancel", "mainMenu")
	)
	selector.Inline(
		selector.Row(confirm),
		selector.Row(cancel),
	)
	text := fmt.Sprintf("You are staking <b>%f CSPR</b> to %s  \n\nPlease confirm the delegation or press Cancel", amount, validator)
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetSuccessDelegateMessage(amount string, validator string, hash string) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		cancel   = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(
		selector.Row(cancel),
	)
	text := fmt.Sprintf(`You delegate has been succesfull submitted!
	
	Delegating %s CSPR to %s .
	
	Transaction link https://testnet.cspr.live/deploy/%s`, amount, validator[len(validator)-5:], hash)
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetSuccessUndelegateMessage(amount string, validator string, hash string) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		cancel   = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(
		selector.Row(cancel),
	)
	text := fmt.Sprintf(`🎉  <b>You delegate has been succesfull submitted! </b>
	
	Delegating %s CSPR to %s .

	🔗 <b>Transaction Details:</b>

	View your transaction here: https://testnet.cspr.live/deploy/%s`, amount, validator[len(validator)-5:], hash)
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetUndelegateDelegates(book types.UndelegateDelegatesList) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
		btnNext  = selector.Data("Next page ➡️", "moveDelegateValidators", strconv.Itoa(int(book.Offset+5)))
		btnPrev  = selector.Data("⬅️ Previous page", "moveDelegateValidators", strconv.Itoa(int(book.Offset-5)))
	)

	var namedButtons []tele.Row

	for _, boobook := range book.Data {
		//namedButtons = append(namedButtons, selector.Row(selector.Data(boobook.GetName(), "showAddress", strconv.Itoa(int(boobook.GetId())))))
		namedButtons = append(namedButtons, selector.Row(selector.Data(fmt.Sprintf("Validator: *%s, Delegated: %s cspr", boobook.GetAddress()[len(boobook.GetAddress())-5:], boobook.GetAmount()),
			"pickUndelegateValidator", strconv.Itoa(int(boobook.GetId())))))
	}
	//namedButtons = append(namedButtons, selector.Row(btnMain))
	if book.Total > 5 {
		tempSelector := selector.Row()
		if book.Offset > 0 {
			tempSelector = append(tempSelector, btnPrev)
		}
		if book.Offset+5 < book.Total {
			tempSelector = append(tempSelector, btnNext)
		}
		namedButtons = append(namedButtons, tempSelector)
	}
	namedButtons = append(namedButtons, selector.Row(btnMain))

	selector.Inline(namedButtons...)

	text := fmt.Sprintf(`🔓 <b>Select Validator for Undelegating</b>

Choose which validator you'd like to unedelegate from.

🔢 <b>Your Total Delegations: </b>%d

📑 <b>Viewing Delegations: </b>%d-%d

⬅️➡️ Use the buttons to navigate through the delegations list.`, book.Total, book.Offset, book.Offset+5)

	return text, []interface{}{tele.ModeHTML, selector}
}

func GetUndelegateAskAmountMessage(balance string, validator string) (interface{}, []interface{}) {

	fb, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		log.Println(err)
	}

	text := fmt.Sprintf(`💸 <b>Choose Undelegation Amount</b>

Select the amount of CSPR to undelegate from validator <code>%s</code>.

🔢 You can choose a preset amount using the button or enter a custom amount. Your choice!`, validator)

	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		balance25  = fb / 100 * 25
		balance50  = fb / 100 * 55
		balance75  = fb / 100 * 75
		balanceRec = fb
		perc25     = selector.Data(fmt.Sprintf("Unstake 25%% (%.6f CSPR)", balance25), "UndelegateSelectAmount", fmt.Sprintf("%f", balance25))
		perc50     = selector.Data(fmt.Sprintf("Unstake 55%% (%.6f CSPR)", balance50), "UndelegateSelectAmount", fmt.Sprintf("%f", balance50))
		perc75     = selector.Data(fmt.Sprintf("Unstake 75%% (%.6f CSPR)", balance75), "UndelegateSelectAmount", fmt.Sprintf("%f", balance75))
		percRec    = selector.Data(fmt.Sprintf("Unstake maximum (%.6f CSPR)", balanceRec), "UndelegateSelectAmount", fmt.Sprintf("%f", balanceRec))
		mainMenu   = selector.Data("🏠 Back to main menu", "mainMenu")
	)

	selector.Inline(
		selector.Row(perc25),
		selector.Row(perc50),
		selector.Row(perc75),
		selector.Row(percRec),
		selector.Row(mainMenu),
	)

	return text, []interface{}{tele.ModeHTML, selector}
}

func GetUndelegateAskConfirmation(amount float64, validator string) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		//withoutMemo = selector.Data("Continue without memo", "transferWithoutMemo")
		confirm = selector.Data("✅ Confirm", "undelegateConfirm")
		cancel  = selector.Data("❌ Cancel", "mainMenu")
	)
	selector.Inline(
		selector.Row(confirm),
		selector.Row(cancel),
	)
	text := fmt.Sprintf("You are unstaking <b>%f CSPR</b> from %s  \n\nPlease confirm the delegation or press Cancel", amount, validator)
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetDepositMessage(address string) (interface{}, []interface{}) {
	var (
		selector    = &tele.ReplyMarkup{}
		SwapService = selector.Data("🔄 Deposit by SimpleSwap", "depositBySwap")
		btnMain     = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(selector.Row(SwapService), selector.Row(btnMain))

	text := fmt.Sprintf(`🔐<b> Your Casper Deposit Address:</b> <code>%s</code> 

📝 <b>Memo / ID / Tag Instructions:</b>
When transferring from an exchange or another wallet, if prompted for a memo, ID, or tag, you can enter "1", any other number, or simply leave it blank.`, address)
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetSwapPairs(offset, total int64, pairs []string, swapType string) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
		btnNext  = selector.Data("Next page ➡️", "moveSwapPairs", strconv.Itoa(int(offset+5)))
		btnPrev  = selector.Data("⬅️ Previous page", "moveSwapPairs", strconv.Itoa(int(offset-5)))
	)

	var namedButtons []tele.Row

	for _, boobook := range pairs {
		//namedButtons = append(namedButtons, selector.Row(selector.Data(boobook.GetName(), "showAddress", strconv.Itoa(int(boobook.GetId())))))
		namedButtons = append(namedButtons, selector.Row(selector.Data(strings.ToUpper(boobook),
			"pickSwapPair", boobook)))
	}
	//namedButtons = append(namedButtons, selector.Row(btnMain))
	if total > 5 {
		tempSelector := selector.Row()
		if offset > 0 {
			tempSelector = append(tempSelector, btnPrev)
		}
		if offset+5 < total {
			tempSelector = append(tempSelector, btnNext)
		}
		namedButtons = append(namedButtons, tempSelector)
	}
	namedButtons = append(namedButtons, selector.Row(btnMain))

	selector.Inline(namedButtons...)
	var text string
	if swapType == "deposit" {
		text = `⭐ Select a Coin for Deposit via SimpleSwap ⭐
		
		Welcome to the deposit gateway!
		
		Please select the coin you wish to deposit into your Casper Telegram Wallet. A wide range of cryptocurrencies are at your fingertips, ready for secure and efficient transactions.
		
		🔄 To begin, simply choose from the list of available coins. Your selection will initiate the deposit process, ensuring a smooth and secure transfer into your wallet.
		Your security and satisfaction are our top priorities. Happy transacting!`
	} else if swapType == "withdraw" {
		text = `🔄 SimpleSwap Withdrawal - Choose Your Coin 🔄
			
		Initiate your withdrawal with SimpleSwap! Easily convert and withdraw to a different cryptocurrency.
		
		🔹 Quick Coin Selection
		Select your desired withdrawal coin from our supported list. SimpleSwap ensures a smooth and secure transaction.`
	} else if swapType == "invoice" {
		text = `⭐ Select a Coin for pay invoice via SimpleSwap ⭐

Welcome to the invoice payment gateway!

To settle your invoice, please choose the cryptocurrency you'd like to use. We offer a broad spectrum of options for a secure and efficient payment experience.

🔄 Begin by selecting from our extensive list of available coins. This will kickstart the payment process, ensuring a seamless and secure transaction.

We prioritize your security and satisfaction above all. Here's to smooth and happy transacting!`
	}

	return text, []interface{}{tele.ModeHTML, selector}
}

func GetSwapAskAmount(fromCur string, toCur string, min, max float64) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(selector.Row(btnMain))

	var mins, maxs string
	//compare float with 0 if it is 0 then it is unlimited
	if min == 0 {
		mins = "unlimited"
	} else {
		mins = fmt.Sprintf("%f %s", min, strings.ToUpper(fromCur))
	}
	if max == 0 {
		maxs = "unlimited"
	} else {
		maxs = fmt.Sprintf("%f %s", max, strings.ToUpper(fromCur))
	}

	text := fmt.Sprintf(`🔹 Specify Your %s Amount 🔹

Ready to proceed with your transaction? Please enter the amount of %s you wish to use.

👉 Minimum Amount: %s
👉 Maximum Amount: %s

Ensure that your entry is within the specified limits for a smooth and secure transaction process.`, strings.ToUpper(fromCur), strings.ToUpper(fromCur), mins, maxs)
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetSwapShowEstimated(estim float64, curr string, amount float64, currFrom string) (interface{}, []interface{}) {
	var (
		selector   = &tele.ReplyMarkup{}
		btnMain    = selector.Data("🏠 Back to main menu", "mainMenu")
		btnConfirm = selector.Data("Confirm", "swapConfirmAmount")
	)
	selector.Inline(selector.Row(btnConfirm), selector.Row(btnMain))

	text := fmt.Sprintf(`⭐ Transaction Confirmation ⭐

You are set %f %s to receive %f %s.

✅ Please Confirm to proceed
🔄 Or, enter a different amount if you wish to adjust the transaction.

Your confirmation ensures accuracy and security in your transaction. Thank you for using Casper Telegram Wallet!`,
		amount, strings.ToUpper(currFrom), estim, strings.ToUpper(curr))
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetSwapAskRefund(curr string, chain string) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(selector.Row(btnMain))

	text := fmt.Sprintf(`🔁 %s %s Refund Address Needed 🔁
To process a refund in USDC, please enter your refund address below.

🔑 Enter Your %s %s Address

Ensure that the address you provide is accurate to facilitate a smooth and secure refund process.
Thank you for using the Casper Telegram Wallet. Your transaction security is our priority!`, strings.ToUpper(curr), strings.ToUpper(chain), strings.ToUpper(curr), strings.ToUpper(chain))
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetSwapSuccess(id, addr, fromCur string, toCur string, amount float64, estimated float64) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(selector.Row(btnMain))

	text := fmt.Sprintf(`🚀 Transaction Ready 🚀

To receive your %f %s, please send %f %s to the following address:

🔗 %s Transfer Address: %s

After completing the transfer, your CSPR will be processed. For seamless exchange options, you may also use the provided SimpleSwap link: https://simpleswap.io/exchange?id=%s
We appreciate your trust in Casper Telegram Wallet. Securely facilitating your transaction is our priority!`, estimated, strings.ToUpper(toCur), amount, strings.ToUpper(fromCur), strings.ToUpper(fromCur), addr, id)
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetSwapChains(data *pb.ShowSwapChains) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)

	var namedButtons []tele.Row

	for _, boobook := range data.GetChains() {
		//namedButtons = append(namedButtons, selector.Row(selector.Data(boobook.GetName(), "showAddress", strconv.Itoa(int(boobook.GetId())))))
		namedButtons = append(namedButtons, selector.Row(selector.Data(boobook.Name, "pickSwapChain", boobook.Short)))
	}
	//namedButtons = append(namedButtons, selector.Row(btnMain))
	namedButtons = append(namedButtons, selector.Row(btnMain))

	selector.Inline(namedButtons...)

	text := fmt.Sprintf(`🔗 Select a Blockchain for %s 🔗

Now, choose the blockchain network for your %s transaction.`, strings.ToUpper(data.GetCoin()), strings.ToUpper(data.GetCoin()))

	return text, []interface{}{tele.ModeHTML, selector}
}

func GetSwapAskAddress(curr string) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(selector.Row(btnMain))

	text := fmt.Sprintf(`🔑 %s Address Requested 🔑

To complete your withdrawal, we need your %s address.

📍 Please Enter Your %s Address

Kindly provide your %s address where the withdrawn funds will be sent. Ensure the address is correct to guarantee a safe and secure transfer of your coins.

Thank you for using Casper Telegram Wallet. We prioritize the safety and security of your transactions!`, strings.ToUpper(curr), strings.ToUpper(curr), strings.ToUpper(curr), strings.ToUpper(curr))
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetPrivacySettingsMessage(isLogging bool) (interface{}, []interface{}) {
	var (
		selector      = &tele.ReplyMarkup{}
		btnMain       = selector.Data("🏠 Back to main menu", "mainMenu")
		btnLogging    tele.Btn //
		loggingStatus string
	)

	if isLogging {
		btnLogging = selector.Data("📝 Turn logging OFF", "toggleLogging")
		loggingStatus = "✅ Enabled"
	} else {
		btnLogging = selector.Data("📝 Turn logging ON", "toggleLogging")
		loggingStatus = "🚫 Currently Disabled"
	}

	selector.Inline(selector.Row(btnLogging), selector.Row(btnMain))
	// :star:️ Casper Blockchain Telegram Wallet - Security Options :star:️
	// :book: Choose Your Logging Preferences
	// With this feature, you can opt to log the interactions you have with our wallet bot. This includes commands you use, balance inquiries, and general navigation through the bot. Please note that while you have the choice to log these interactions, all transactions made on the blockchain, including sends, delegates, and undelegates of CSPR, are inherently public and will always be recorded on the blockchain ledger.
	// :globe_with_meridians: Blockchain Transparency
	// Remember, the blockchain is an open ledger. This means all transactional activities, such as transfers, rewards, delegations, and undelegations, are transparent and visible to the public. These transactions are permanently recorded on the blockchain and are accessible to anyone.
	text := fmt.Sprintf(`
📝 <b>Logging Status & Options 📝</b>

Current Logging: %s

🔍 Understanding Your Logging Preferences
Your interactions are currently being recorded to enhance your experience, ensuring faster and more efficient service through the use of cached data.

🚫 Opt-Out Option
Prefer not to be logged? You can choose to opt out of logging. Please note, opting out may result in increased delays as we won't utilize cached data for your transactions.
🔄 To change your logging preference, please press the button in the menu. We're committed to balancing your security with efficient service!`,
		loggingStatus)
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetSwapLoadMsg() (interface{}, []interface{}) {
	return `⭐ Data Loading... Please Hang Tight ⭐`, []interface{}{tele.ModeHTML}
}

func GetErrorExportPKNotStore() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		cancel   = selector.Data("Back to main menu", "mainMenu")
	)
	selector.Inline(
		selector.Row(cancel),
	)
	return `🔒 Key Export Unavailable 🔒

We've noted your preference to not store the key, and as a result, we're unable to export it at this time.

🛑 Important Notice

This decision enhances your privacy but means that key retrieval or export is not possible through our service.
For any further assistance or inquiries, please don't hesitate to reach out. Your security and privacy are our top priorities at Casper Telegram Wallet.`, []interface{}{tele.ModeHTML, selector}
}

func GetInvoicesListMsg(invoices *pb.InvoicesListResponse) (interface{}, []interface{}) {
	var (
		selector          = &tele.ReplyMarkup{}
		btnMain           = selector.Data("🏠 Back to main menu", "mainMenu")
		btnCreateEntry    = selector.Data("📝 Create new invoice", "createNewInvoice")
		btnRecentInvoices = selector.Data("📝 Recent invoices", "recentInvoices")
		btnNext           = selector.Data("Next page", "moveInvoicesList", strconv.Itoa(int(invoices.GetOffset()+5)))
		btnPrev           = selector.Data("Previous page", "moveInvoicesList", strconv.Itoa(int(invoices.GetOffset()-5)))
	)

	var namedButtons []tele.Row

	for _, boobook := range invoices.GetInvoices() {
		namedButtons = append(namedButtons, selector.Row(selector.Data(boobook.GetName(), "showInvoice", strconv.Itoa(int(boobook.GetId())))))
	}
	//namedButtons = append(namedButtons, selector.Row(btnMain))
	if invoices.GetTotal() > 5 {
		tempSelector := selector.Row()
		if invoices.GetOffset() > 0 {
			tempSelector = append(tempSelector, btnPrev)
		}
		if invoices.GetOffset()+5 < invoices.GetTotal() {
			tempSelector = append(tempSelector, btnNext)
		}
		namedButtons = append(namedButtons, tempSelector)
	}
	namedButtons = append(namedButtons, selector.Row(btnRecentInvoices))
	namedButtons = append(namedButtons, selector.Row(btnCreateEntry))
	namedButtons = append(namedButtons, selector.Row(btnMain))

	selector.Inline(namedButtons...)

	text := `Welcome to your Casper Telegram Wallet

📄 Invoices Hub 📄

Welcome to your Invoices Hub, where managing your transactions is effortless and intuitive.

📋 <b>Your Invoice List</b>

Browse through your list of created invoices. Each entry is at your fingertips for review, and should you need to, you can easily delete any invoice that's no longer required.

👀 <b>View Recent Invoices</b>

Access and review your most recent invoices with a simple tap. Stay informed about your transaction history and financial movements.

🆕 <b>Create New Invoice</b>

Looking to issue a new invoice? Click here to start crafting a new invoice tailored to your specifications, ensuring your billing needs are met with precision.
Navigate below to manage your invoices efficiently. We're here to support your financial organization and security within the Casper Telegram Wallet.
`
	// if len(book.Data) == 0 {
	// 	text = ("<b>Welcome to your Casper Telegram Wallet</b>\n\n📖 <b>Your Address Book</b> \n\n Currently, your address book is empty. Let's begin populating it with addresses for easier access and convenience in the future.")
	// }

	return text, []interface{}{tele.ModeHTML, selector}

}

func AskInvoiceName() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(
		selector.Row(btnMain),
	)
	return `🆕 Create New Invoice 🆕

To begin crafting your new invoice, please start by giving it a unique name.

✏️ <b>Invoice Name:</b>

This will help identify your invoice for both you and the recipient. Choose a name that reflects the purpose or contents of the invoice.
Once named, we'll guide you through the next steps to complete your invoice details.`, []interface{}{tele.ModeHTML, selector}
}

func AskInvoiceAmount() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(
		selector.Row(btnMain),
	)
	return `🆕 Create New Invoice

Please specify the amount in CSPR for this invoice. This will determine the total chargeable amount to the recipient.

💰 <b>Enter Invoice Amount</b>

Enter the desired CSPR amount to accurately reflect the value of the transaction or service provided.
Your precision in this step ensures clarity and transparency for both parties involved in the transaction. Thank you for using Casper Telegram Wallet for your invoicing needs.`, []interface{}{tele.ModeHTML, selector}
}

func AskInvoiceRepeatability() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(
		selector.Row(btnMain),
	)
	return `🔄 Set Invoice Payment Frequency 🔄

Please specify the frequency with which this invoice can be paid. Enter '0' for unlimited payments, allowing the invoice to be paid multiple times as needed.

🔢 Enter Payment Frequency

Choose how many times the invoice should be settled. A value greater than '0' sets a specific limit, while '0' allows for flexible, repeated payments.

Adjust according to your billing requirements and preferences. Your flexibility in managing payments enhances your invoicing efficiency with Casper Telegram Wallet.
	`, []interface{}{tele.ModeHTML, selector}
}

func AskInvoiceComment() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(
		selector.Row(btnMain),
	)
	return `🖊️ Personalize Your Invoice 🖊️

We invite you to add a comment to your invoice. This could be anything from transaction details, a note of appreciation, or any specific instructions for the recipient.`, []interface{}{tele.ModeHTML, selector}
}

func InvoiceCreateSuccess(res *pb.InvoiceCreateSuccess, botname string) (interface{}, []interface{}) {
	var (
		selector    = &tele.ReplyMarkup{}
		btnMain     = selector.Data("🏠 Back to main menu", "mainMenu")
		btnInvoices = selector.Data("🧾 Invoices", "invoices")
	)
	selector.Inline(
		selector.Row(btnInvoices),
		selector.Row(btnMain),
	)

	text := fmt.Sprintf(`✅ Invoice Successfully Created ✅

Your invoice has been generated with the following details:

<b>- Name: %s</b>
<b>- Amount: %s</b>
<b>- Repeatability: %d</b>
<b>- Comment: %s</b>

🔗 https://t.me/%s?start=inv%s

You can now share this link directly with the intended recipient for payment. Thank you for using Casper Telegram Wallet for your invoicing needs!`, res.GetName(), res.GetAmount(), res.GetRepeatability(), res.GetComment(), botname, res.GetShort())
	return text, []interface{}{tele.ModeHTML, selector}
}

func InvoiceDetailed(res *pb.InvoiceDetailed, botname string) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
		//btnInvoices     = selector.Data("🧾 Show transfers(TODO)", "showInvoiceTransfers")
		btnDelete       = selector.Data("🗑 Delete this invoice", "deleteInvoice", strconv.Itoa(int(res.Id)))
		btnShowPayments = selector.Data("💵 Show payments", "showInvoicePayments", strconv.Itoa(int(res.Id)))
	)
	selector.Inline(
		selector.Row(btnShowPayments),
		//selector.Row(btnInvoices),
		selector.Row(btnDelete),
		selector.Row(btnMain),
	)

	text := fmt.Sprintf(`📄 Invoice Details 📄

<b>Your Invoice Summary:</b>

<b>- Name: %s</b>
<b>- Amount: %s</b>
<b>- Repeatability: %d</b>
<b>- Comment: %s</b>
<b>- Payments Received: %d</b>

🔗 Invoice Link: https://t.me/%s?start=inv%s

Feel free to copy this link and share it with the relevant person or group for payment.

<b>Invoice Management Options:</b>

<b>-View Payments:</b> Keep an eye on all received payments for this invoice.
<b>-Delete Invoice:</b> Remove this invoice if it's no longer necessary.

Utilize the buttons below to explore payment details or to delete the invoice. We're here to ensure your invoice management is smooth and efficient with Casper Telegram Wallet.`, res.GetName(), res.GetAmount(), res.GetRepeatability(), res.GetComment(), res.Paid, botname, res.GetShort())
	return text, []interface{}{tele.ModeHTML, selector}
}

func DeleteInvoiceConfirmation(id uint64) (interface{}, []interface{}) {
	var (
		selector    = &tele.ReplyMarkup{}
		btnDelete   = selector.Data("Confirm", "deleteInvoiceConfirm", strconv.Itoa(int(id)))
		btnInvoices = selector.Data("Cancel", "invoices")
	)
	selector.Inline(
		selector.Row(btnDelete),
		selector.Row(btnInvoices),
	)

	text := `🗑️ Confirm Invoice Deletion 🗑️

Are you sure you want to delete this invoice? This action cannot be undone.

<b>Confirm Deletion:</b> If you're certain, please confirm to permanently remove the invoice from your records.
<b>Cancel: </b>If you wish to keep the invoice, simply cancel this request.

Please choose wisely to ensure the integrity of your financial records. Your peace of mind is important to us at Casper Telegram Wallet.`
	return text, []interface{}{tele.ModeHTML, selector}
}

func InvoiceAskRegisterPM() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		NewWallet      = selector.Data("Create new wallet", "createWallet")
		ExistingWallet = selector.Data("Add existing wallet", "addExistingWallet")
	)
	selector.Inline(
		selector.Row(NewWallet),
		selector.Row(ExistingWallet),
	)
	return `⭐ <b>Welcome to Casper Blockchain Telegram Wallet</b> ⭐

To complete your invoice payment, a wallet is essential. You have the flexibility to either set up a brand new wallet or add an existing one.

🔐 <b>Set Up or Add Wallet</b>

<b>Create New Wallet:</b> Choose this option for a quick and secure setup of a new wallet.
<b>Add Existing Wallet:</b> Select this if you'd like to link an already existing wallet to this platform.

Follow the simple steps to get started and unlock the full potential of managing invoices with ease.

We're thrilled to guide you through this process and ensure a smooth and secure experience on the Casper Telegram Wallet!`, []interface{}{tele.ModeHTML, selector}
}

func PayInvoiceNotAviablePM() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		//btnLock = selector.Data("Lock", "Lock")
		btnMain = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Inline(
		selector.Row(btnMain),
	)
	return `🚫 Invoice Payment Unavailable 🚫
Unfortunately, this invoice cannot currently be processed for payment. We kindly ask that you reach out to the individual or entity that issued this invoice for further assistance.
Thank you for your understanding, and we apologize for any inconvenience this may cause.`, []interface{}{tele.ModeHTML, selector}
}

func PayInvoiceRegisteredResponse(res *pb.PayInvoiceRegisteredResponse) (interface{}, []interface{}) {
	log.Println("SHORT: ", res.GetShort())
	var (
		selector              = &tele.ReplyMarkup{}
		btnMain               = selector.Data("🏠 Back to main menu", "mainMenu")
		btnPayByTransfer      = selector.Data("Pay by transfer", "payInvoiceTransfer", res.GetShort())
		btnPayBySwap          = selector.Data("Pay by swap", "payInvoiceSwap", res.GetShort())
		btnPayByTransferBlock = selector.Data("Pay by transfer 🚫", "PASS")
	)
	if res.BalanceEnough {
		selector.Inline(
			selector.Row(btnPayByTransfer, btnPayBySwap),
			selector.Row(btnMain),
		)
	} else {
		selector.Inline(
			selector.Row(btnPayByTransferBlock, btnPayBySwap),
			selector.Row(btnMain),
		)
	}

	text := fmt.Sprintf(`
🔓 Invoice Payment Options 🔓 

You're all set to pay the following invoice:

<b>- Name: %s</b>
<b>- Amount: %s</b>
<b>- Comment: %s</b>

Choose how you'd like to proceed with the payment:
1. <b>*Pay with CSPR Balance:*</b>
Directly use your CSPR balance for a fast and secure transaction.
2. <b>*Opt for SimpleSwap:*</b>
Prefer a different cryptocurrency? Convert it to CSPR through SimpleSwap and complete your payment seamlessly.

Select the method that aligns with your preferences. We're dedicated to providing a straightforward and secure payment process.

Thank you for trusting Casper Telegram Wallet for your transactions. We're committed to making your experience smooth and efficient!"

This message includes detailed invoice information and clearly explains the payment options available to the user, ensuring they have all the necessary information to proceed.`, res.GetName(), res.GetAmount(), res.GetComment())
	return text, []interface{}{tele.ModeHTML, selector}
}

func BlockForGroup() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
	)

	return `⭐ <b>Welcome to Casper Blockchain Telegram Wallet</b> ⭐


You can create invoices by /invoice@botname`, []interface{}{tele.ModeHTML, selector}
}

func GetPaymentsListMsg(payments *pb.PaymentsListResponse) (interface{}, []interface{}) {
	var (
		selector  = &tele.ReplyMarkup{}
		btnMain   = selector.Data("🏠 Back to main menu", "mainMenu")
		btnExport = selector.Data("Export to .csv", "exportPaymentsInvoice", strconv.Itoa(int(payments.GetId())))
		//btnCreateEntry = selector.Data("📝 Create new invoice", "createNewInvoice")
		btnNext = selector.Data("Next page", "movePaymentsList",
			strconv.Itoa(int(payments.GetId())), strconv.Itoa(int(payments.GetOffset()+10)))
		btnPrev = selector.Data("Previous page", "movePaymentsList",
			strconv.Itoa(int(payments.GetId())), strconv.Itoa(int(payments.GetOffset()-10)))
	)

	var namedButtons []tele.Row
	if payments.Total == 0 {
		text := `
🚫 No Payments Received 🚫

Currently, there are no payments recorded for this invoice. This could be due to the invoice not being viewed or paid yet.Please check back later or follow up with the intended recipient to ensure they've received and reviewed the invoice details.

We're here to assist with any further actions you may need to take regarding this invoice through Casper Telegram Wallet.`
		namedButtons = append(namedButtons, selector.Row(btnMain))
		selector.Inline(namedButtons...)
		return text, []interface{}{tele.ModeHTML, selector}
	}

	namedButtons = append(namedButtons, selector.Row(btnExport))

	text := `📝 Invoice Payments Overview 📝

Here's a summary of the recent payments made towards your invoices. This overview includes details about each payer, the amount paid, and the payment status.

- Payer: The individual or entity that has made the payment.
- Amount: The total sum paid towards the invoice.
- Status: Indicates whether the invoice has been fully settled or if there are outstanding amounts.

Thank you for managing your invoicing with us. We're committed to ensuring a transparent and efficient tracking of your transactions!

`

	for _, boobook := range payments.GetPayments() {
		text += fmt.Sprintf("<b>From: %s \nAmount: %s \nCorrect: %t </b> \n\n", boobook.From, boobook.Amount, boobook.Success)
	}
	//namedButtons = append(namedButtons, selector.Row(btnMain))
	if payments.GetTotal() > 10 {
		tempSelector := selector.Row()
		if payments.GetOffset() > 0 {
			tempSelector = append(tempSelector, btnPrev)
		}
		if payments.GetOffset()+10 < payments.GetTotal() {
			tempSelector = append(tempSelector, btnNext)
		}
		namedButtons = append(namedButtons, tempSelector)
	}
	namedButtons = append(namedButtons, selector.Row(btnMain))
	selector.Inline(namedButtons...)
	return text, []interface{}{tele.ModeHTML, selector}
}

func GetRecentlyInvoices(invoices *pb.RecentInvoicesListResponse) (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
		btnNext  = selector.Data("Next page", "moveRecentInvoicesList", strconv.Itoa(int(invoices.GetOffset()+5)))
		btnPrev  = selector.Data("Previous page", "moveRecentInvoicesList", strconv.Itoa(int(invoices.GetOffset()-5)))
	)

	var namedButtons []tele.Row

	for _, boobook := range invoices.GetInvoices() {
		namedButtons = append(namedButtons, selector.Row(selector.Data(boobook.GetName()+" "+boobook.Status,
			"payInvoice", boobook.Short)))
	}
	//namedButtons = append(namedButtons, selector.Row(btnMain))
	if invoices.GetTotal() > 5 {
		tempSelector := selector.Row()
		if invoices.GetOffset() > 0 {
			tempSelector = append(tempSelector, btnPrev)
		}
		if invoices.GetOffset()+5 < invoices.GetTotal() {
			tempSelector = append(tempSelector, btnNext)
		}
		namedButtons = append(namedButtons, tempSelector)
	}

	namedButtons = append(namedButtons, selector.Row(btnMain))

	selector.Inline(namedButtons...)

	text := `📑 Quick Access to Your Recent Invoices 📑

In this section, you can effortlessly browse through the 20 invoices you've most recently opened. This feature is designed to give you quick and easy access to your latest invoicing activities, helping you stay on top of your financial management.

Thank you for choosing our services for your invoicing needs. Your convenience and efficiency are our top priorities!
`
	return text, []interface{}{tele.ModeHTML, selector}

}

func GetExportPaymentsInvoice(data []byte, short string) (interface{}, []interface{}) {
	reader := bytes.NewReader(data)
	a := &tele.Document{File: tele.FromReader(reader), FileName: short + "_payments.csv",
		Caption: `📊Invoice Payments Exported 📊

Your request to export the payment details for your invoices has been successfully processed. Attached, you'll find the .csv file containing all the relevant payment information.

📁 <b>File Details:</b>
- Format: .csv
- Contents: Payments for invoices, including dates, amounts, payers, and related comments or memos.

Thank you for using our services for your invoicing needs. We're here to support your financial management and reporting efforts!`}
	return a, []interface{}{tele.ModeHTML}
}

func GetInvoiceRepeatabilityIsNotValid() (interface{}, []interface{}) {
	var (
		selector = &tele.ReplyMarkup{}
		btnMain  = selector.Data("🏠 Back to main menu", "mainMenu")
	)
	selector.Row(btnMain)
	return `
⚠️ Invalid Frequency Entered ⚠️

The frequency value you entered is incorrect. Please ensure the number is an integer that is 0 or greater.

🔢 <b>Enter Valid Frequency</b>

Input a valid numerical value to set the number of times this invoice can be paid. Remember, '0' allows for unlimited payments.

Kindly re-enter a valid number to proceed with setting up your invoice payment terms. We're here to assist with your invoicing needs at Casper Telegram Wallet.
`, []interface{}{tele.ModeHTML, selector}
}
