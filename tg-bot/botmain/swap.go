package botmain

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"

	"github.com/Simplewallethq/tg-bot/ent"
	"github.com/Simplewallethq/tg-bot/ent/recentinvoices"
	"github.com/Simplewallethq/tg-bot/ent/user"
	pb "github.com/Simplewallethq/tg-bot/tggateway/proto"
	tggateway "github.com/Simplewallethq/tg-bot/tggateway/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func (b *BotMain) HandleSwapState(msg *pb.TgTextMessage) error {
	log.Println("received swap state message", msg)
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	if len(state) < 3 || state[0] != "Swap" {
		err = b.State.DeleteUserState(msg.GetFrom().GetId())
		if err != nil {
			return errors.Wrap(err, "failed delete user state")
		}
		return errors.Wrap(err, "bad get user state")
	}

	switch state[2] {
	case "askAmount":
		err := b.PickSwapAmountCustom(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle swap amount")
		}
	case "askRefund":
		err := b.SwapSetRefundAddress(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle swap refund ")
		}
	case "askAddress":
		err := b.SwapSetAddress(msg)
		if err != nil {
			return errors.Wrap(err, "failed handle swap refund ")
		}
	}
	return nil
}

func (b *BotMain) SwapBySwapButton(msg tggateway.TgMessageMsg) error {
	out := pb.SwapBySwapButton{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
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
	var swapsDB *ent.Swaps
	if out.GetWithdraw() {
		swapsDB, err = b.DB.Swaps.Create().SetRefundAddress(u.PublicKey).SetFromCurrency("cspr").SetType("withdraw").SetOwner(u).Save(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed create swap")
		}
	} else {
		swapsDB, err = b.DB.Swaps.Create().SetToAddress(u.PublicKey).SetToCurrency("cspr").SetType("deposit").SetOwner(u).Save(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed create swap")
		}
	}

	log.Println(swapsDB.ID)

	err = b.State.SetUserState(out.GetUser().GetId(), []string{"Swap", strconv.Itoa(swapsDB.ID), "pickCurrency"})
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}

	data := pb.AskSwapPairs{
		User:   out.GetUser(),
		MsgId:  out.MsgId,
		Limit:  5,
		Offset: 0,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.MessagesChan <- tggateway.TgMessageMsg{
		Name: "AskSwapPairs",
		Data: dataBytes,
	}

	return nil
}

func (b *BotMain) AskSwapPairs(msg tggateway.TgMessageMsg) error {
	out := pb.AskSwapPairs{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	loggedin, err := b.CheckLogin(out.GetUser())
	if err != nil {
		return err
	}
	if !loggedin {
		return nil
	}

	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed delete user state")
	}
	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	SwapDB, err := b.DB.Swaps.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get swapdb")
	}

	//GET PAIRS FROM SIMPLESWAP API
	// pairs, err := b.SwapClient.GetCSPRPairs()
	// if err != nil {
	// 	return errors.Wrap(err, "failed get swap pairs")
	// }

	//STATIC PAIRS FOR CSPR
	var pairs []string
	if SwapDB.Type == "withdraw" {
		pairs = []string{"btc"}
	} else if SwapDB.Type == "deposit" || SwapDB.Type == "invoice" {
		pairs = []string{"btc", "eth", "usdt", "usdc"}
	}

	var preparedPairs []string

	if len(pairs) >= int(out.GetOffset()) {
		if len(pairs) >= int(out.GetLimit())+int(out.GetOffset()) {
			preparedPairs = pairs[out.GetOffset() : out.GetLimit()+out.GetOffset()]
		} else {
			preparedPairs = pairs[out.GetOffset():]
		}

	}
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	data := pb.ShowSwapPairs{
		User:   out.GetUser(),
		MsgId:  out.MsgId,
		Offset: out.GetOffset(),
		Pairs:  preparedPairs,
		Type:   SwapDB.Type,
		Total:  int64(len(pairs)),
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "ShowSwapPairs",
		Data: dataBytes,
	}

	return nil
}

func (b *BotMain) PickSwapPair(msg tggateway.TgMessageMsg) error {
	out := pb.PickSwapPair{}

	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	// u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	// if err != nil {
	// 	return errors.Wrap(err, "failed get user")
	// }

	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	SwapDB, err := b.DB.Swaps.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get swapdb")
	}
	if SwapDB.Type == "withdraw" {
		SwapDB, err = SwapDB.Update().SetToCurrency(out.GetCur()).Save(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed update swaps")
		}
	} else if SwapDB.Type == "deposit" || SwapDB.Type == "invoice" {
		SwapDB, err = SwapDB.Update().SetFromCurrency(out.GetCur()).Save(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed update swaps")
		}
	} else {
		//return unknown swaptype error
		return errors.Wrap(err, "unknown swap type")
	}

	//map currency:slice of chains

	type chain struct {
		name string
		code string
	}
	chains := map[string][]chain{
		"eth":  {{"Ethereum", "base"}, {"Binance Smart Chain", "bsc"}, {"Polygon", "op"}, {"Zcash", "zc"}, {"Ethereum ERC20", "erc20"}},
		"usdt": {{"Ethereum", "erc20"}, {"Tron", "trc20"}, {"Binance Smart Chain", "bep20"}, {"Near", "near"}},
		"usdc": {{"Tron", "trc20"}, {"Binance Smart Chain", "bep20"}, {"Tron", "trc20"}},
	}

	if val, ok := chains[out.GetCur()]; ok {
		if len(state) < 3 {
			state = append(state, "askChain")
		} else {
			state[2] = "askChain"
		}
		state = append(state, strconv.Itoa(int(out.GetMsgId())))
		err = b.State.SetUserState(out.GetUser().GetId(), state)
		if err != nil {
			return errors.Wrap(err, "failed set user state")
		}
		var result []*pb.Chain
		for i := range val {
			result = append(result, &pb.Chain{Name: val[i].name, Short: val[i].code})
		}
		data := pb.ShowSwapChains{
			User:   out.GetUser(),
			MsgId:  out.MsgId,
			Coin:   out.GetCur(),
			Chains: result,
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}

		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "SwapShowChains",
			Data: dataBytes,
		}
	} else {
		if len(state) < 3 {
			state = append(state, "askAmount")
		} else {
			state[2] = "askAmount"
		}
		if SwapDB.Type == "invoice" {
			err = b.CalculateSwapAmountForInvoice(out.GetUser())
			return err
		}
		state = append(state, strconv.Itoa(int(out.GetMsgId())))
		err = b.State.SetUserState(out.GetUser().GetId(), state)
		if err != nil {
			return errors.Wrap(err, "failed set user state")
		}
		min, max, err := b.SwapClient.GetRanges(SwapDB.FromCurrency, SwapDB.ToCurrency)
		if err != nil {
			return errors.Wrap(err, "err get sawp ranges")
		}

		data := pb.SwapAskAmount{
			User:     out.GetUser(),
			MsgId:    out.GetMsgId(),
			FromCurr: SwapDB.FromCurrency,
			ToCurr:   SwapDB.ToCurrency,
			Min:      min,
			Max:      max,
		}

		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}

		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "SwapAskAmount",
			Data: dataBytes,
		}
	}

	return nil
}

func (b *BotMain) CalculateSwapAmountForInvoice(user *pb.User) error {
	state, err := b.State.GetUserState(user.GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	// u, err := b.DB.User.Query().Where(user.ID(User.GetId())).Only(context.Background())
	// if err != nil {
	// 	return errors.Wrap(err, "failed get user")
	// }

	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	SwapDB, err := b.DB.Swaps.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get swapdb")
	}
	amountCSPR, err := strconv.ParseFloat(SwapDB.Amount, 64)
	if err != nil {
		return errors.Wrap(err, "failed convert string to float")
	}
	try := 0
	var estimFrom float64
	var amountTo float64
	estimFrom, err = b.SwapClient.GetEstimated(SwapDB.ToCurrency+SwapDB.ToNetwork, SwapDB.FromCurrency+SwapDB.FromNetwork, amountCSPR)
	if err != nil {
		return errors.Wrap(err, "failed get estimated amount")
	}
	for try <= 5 {
		amountTo, err = b.SwapClient.GetEstimated(SwapDB.FromCurrency+SwapDB.FromNetwork, SwapDB.ToCurrency+SwapDB.ToNetwork, estimFrom)
		if err != nil {
			return errors.Wrap(err, "failed get estimated amount")
		}
		dif := ((amountTo - amountCSPR) / ((amountTo + amountCSPR) / 2)) * 100
		if math.Abs(dif) > 5 {
			if dif < 0 {
				estimFrom = estimFrom + ((estimFrom / 100) * math.Abs(dif) / 2)
			} else if dif > 0 {
				estimFrom = estimFrom - ((estimFrom / 100) * math.Abs(dif) / 2)
			}
		} else {
			break
		}
		try++
	}
	SwapDB, err = SwapDB.Update().SetAmount(fmt.Sprintf("%f", estimFrom)).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update transfer")
	}
	var msgid int64
	if len(state) < 4 {
		msgid = -1
	} else {
		msgid, err = strconv.ParseInt(state[3], 10, 64)
		if err != nil {
			msgid = -1
		}
	}
	data := pb.SwapAmountEstimated{
		User:      user,
		MsgId:     msgid,
		Estimated: amountTo,
		Curr:      SwapDB.ToCurrency,
		Amount:    estimFrom,
		CurrFrom:  SwapDB.FromCurrency,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "SwapShowEstimatedAmount",
		Data: dataBytes,
	}

	return nil

}

func (b *BotMain) PickSwapChain(msg tggateway.TgMessageMsg) error {
	out := pb.PickSwapChain{}

	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	// u, err := b.DB.User.Query().Where(user.ID(out.GetUser().GetId())).Only(context.Background())
	// if err != nil {
	// 	return errors.Wrap(err, "failed get user")
	// }

	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	SwapDB, err := b.DB.Swaps.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get swapdb")
	}

	if SwapDB.Type == "withdraw" {
		SwapDB, err = SwapDB.Update().SetToNetwork(out.GetChain()).Save(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed update swaps")
		}
	} else if SwapDB.Type == "deposit" || SwapDB.Type == "invoice" {
		SwapDB, err = SwapDB.Update().SetFromNetwork(out.GetChain()).Save(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed update swaps")
		}
	} else {
		//return unknown swaptype error
		return errors.Wrap(err, "unknown swap type")
	}

	data := pb.SwapLoadingResponse{
		User:  out.GetUser(),
		MsgId: out.GetMsgId(),
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "SwapLoadMsg",
		Data: dataBytes,
	}

	if len(state) < 3 {
		state = append(state, "askAmount")
	} else {
		state[2] = "askAmount"
	}
	state = append(state, strconv.Itoa(int(out.GetMsgId())))
	err = b.State.SetUserState(out.GetUser().GetId(), state)
	if err != nil {
		return errors.Wrap(err, "failed set user state")
	}
	log.Println("swap from", SwapDB.FromCurrency, "to", SwapDB.ToCurrency)
	if SwapDB.Type == "invoice" {
		err = b.CalculateSwapAmountForInvoice(out.GetUser())
		return err
	}

	min, max, err := b.SwapClient.GetRanges(SwapDB.FromCurrency+SwapDB.ToNetwork, SwapDB.ToCurrency)
	if err != nil {
		return errors.Wrap(err, "err get sawp ranges")
	}

	data2 := pb.SwapAskAmount{
		User:     out.GetUser(),
		MsgId:    out.GetMsgId(),
		FromCurr: SwapDB.FromCurrency,
		ToCurr:   SwapDB.ToCurrency,
		Min:      min,
		Max:      max,
	}

	dataBytes, err = proto.Marshal(&data2)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "SwapAskAmount",
		Data: dataBytes,
	}

	return nil
}

func (b *BotMain) PickSwapAmountCustom(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	SwapDB, err := b.DB.Swaps.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}
	_, ok := new(big.Float).SetString(msg.GetText())
	var msgid int64
	if len(state) < 4 {
		msgid = -1
	} else {
		msgid, err = strconv.ParseInt(state[3], 10, 64)
		if err != nil {
			msgid = -1
		}
	}
	if !ok {
		data := pb.TransferAmountIsNotValid{
			User:  msg.GetFrom(),
			MsgId: msgid,
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}

		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "TransferAmountNotValid",
			Data: dataBytes,
		}
		return nil
	}

	SwapDB, err = SwapDB.Update().SetAmount(msg.GetText()).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update transfer")
	}

	amountF, err := strconv.ParseFloat(SwapDB.Amount, 64)
	if err != nil {
		return errors.Wrap(err, "failed parse amount")
	}

	estim, err := b.SwapClient.GetEstimated(SwapDB.FromCurrency+SwapDB.ToNetwork, SwapDB.ToCurrency, amountF)
	if err != nil {
		return errors.Wrap(err, "failed get estimated amount")
	}

	data := pb.SwapAmountEstimated{
		User:      msg.GetFrom(),
		MsgId:     msgid,
		Estimated: estim,
		Curr:      SwapDB.ToCurrency,
		Amount:    amountF,
		CurrFrom:  SwapDB.FromCurrency,
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "SwapShowEstimatedAmount",
		Data: dataBytes,
	}

	_, err = SwapDB.Update().SetAmountRecive(fmt.Sprintf("%f", estim)).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update ReciveAmount")
	}

	return nil

}

func (b *BotMain) SwapConfirmAmount(msg tggateway.TgMessageMsg) error {
	out := pb.SwapConfirmAmount{}
	if err := proto.Unmarshal(msg.Data, &out); err != nil {
		return errors.Wrap(err, "failed unmarshal")
	}
	loggedin, err := b.CheckLogin(out.GetUser())
	if err != nil {
		return err
	}
	if !loggedin {
		return nil
	}

	state, err := b.State.GetUserState(out.GetUser().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	SwapDB, err := b.DB.Swaps.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get transfer")
	}
	if SwapDB.Type == "deposit" || SwapDB.Type == "invoice" {
		state[2] = "askRefund"
		err = b.State.SetUserState(out.GetUser().GetId(), state)
		if err != nil {
			return errors.Wrap(err, "failed set user state")
		}
		data := pb.SwapAskRefundAddress{
			User:  out.GetUser(),
			MsgId: out.MsgId,
			Curr:  SwapDB.FromCurrency,
			Chain: SwapDB.FromNetwork,
		}

		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}

		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "SwapAskRefundAddress",
			Data: dataBytes,
		}
	} else if SwapDB.Type == "withdraw" {
		state[2] = "askAddress"
		err = b.State.SetUserState(out.GetUser().GetId(), state)
		if err != nil {
			return errors.Wrap(err, "failed set user state")
		}
		data := pb.SwapAskToAddress{
			User:  out.GetUser(),
			MsgId: out.MsgId,
			Curr:  SwapDB.ToCurrency + " " + SwapDB.ToNetwork,
		}

		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}

		b.ResponseChan <- tggateway.TgResponseMsg{
			Name: "SwapAskToAddress",
			Data: dataBytes,
		}
	} else {
		return errors.Wrap(err, "unknown swap type")
	}

	return nil
}

func (b *BotMain) SwapSetRefundAddress(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	var msgid int64
	if len(state) < 4 {
		msgid = -1
	} else {
		msgid, err = strconv.ParseInt(state[3], 10, 64)
		if err != nil {
			msgid = -1
		}
	}
	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	SwapDB, err := b.DB.Swaps.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get swap")
	}

	SwapDB, err = SwapDB.Update().SetRefundAddress(msg.GetText()).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update swap")
	}

	amountF, err := strconv.ParseFloat(SwapDB.Amount, 64)
	if err != nil {
		return errors.Wrap(err, "failed parse amount")
	}

	swap, err := b.SwapClient.MakeExchange(SwapDB.FromCurrency+SwapDB.FromNetwork, SwapDB.ToCurrency+SwapDB.ToNetwork, SwapDB.RefundAddress, SwapDB.ToAddress, amountF, SwapDB.ExtraID)
	if err != nil {
		return errors.Wrap(err, "failed make exchange")
	}

	_, err = SwapDB.Update().SetSwapID(SwapDB.SwapID).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update swap")
	}

	err = b.State.DeleteUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed delete user state")
	}

	estimated, err := strconv.ParseFloat(SwapDB.AmountRecive, 64)
	if err != nil {
		return errors.Wrap(err, "failed calculate estimated")
	}

	data := pb.SwapSuccessResponse{
		User:       msg.GetFrom(),
		MsgId:      msgid,
		DepAddress: swap.AddressFrom,
		Id:         swap.ID,
		FromCur:    SwapDB.FromCurrency + " " + SwapDB.FromNetwork,
		ToCur:      SwapDB.ToCurrency + " " + SwapDB.ToNetwork,
		Amount:     amountF,
		Estimated:  estimated,
	}
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	if !u.EnableLogging {
		err = b.DB.Swaps.DeleteOneID(trID).Exec(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed del swap")
		}
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "SwapSuccessResponse",
		Data: dataBytes,
	}

	if SwapDB.Type == "invoice" && u.EnableLogging {
		rec, err := u.QueryRecentInvoices().Where(recentinvoices.
			InvoiceID(SwapDB.InvoiceID)).
			Only(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed get recent invoice")
		}
		err = rec.Update().SetStatus("CreateSwap").Exec(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed set status for recent invoice")
		}
	}
	return nil

}

func (b *BotMain) SwapSetAddress(msg *pb.TgTextMessage) error {
	state, err := b.State.GetUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}
	var msgid int64
	if len(state) < 4 {
		msgid = -1
	} else {
		msgid, err = strconv.ParseInt(state[3], 10, 64)
		if err != nil {
			msgid = -1
		}
	}
	trID, err := strconv.Atoi(state[1])
	if err != nil {
		return errors.Wrap(err, "failed get user state")
	}

	SwapDB, err := b.DB.Swaps.Get(context.Background(), trID)
	if err != nil {
		return errors.Wrap(err, "failed get swap")
	}

	SwapDB, err = SwapDB.Update().SetToAddress(msg.GetText()).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update swap")
	}

	amountF, err := strconv.ParseFloat(SwapDB.Amount, 64)
	if err != nil {
		return errors.Wrap(err, "failed parse amount")
	}

	log.Println(SwapDB.FromNetwork)

	swap, err := b.SwapClient.MakeExchange(SwapDB.FromCurrency+SwapDB.FromNetwork, SwapDB.ToCurrency+SwapDB.ToNetwork, SwapDB.RefundAddress, SwapDB.ToAddress, amountF, SwapDB.ExtraID)
	if err != nil {
		//handle Error
		b.logger.Error(err)
		return errors.Wrap(err, "failed make exchange")
	}

	_, err = SwapDB.Update().SetSwapID(SwapDB.SwapID).Save(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update swap")
	}

	err = b.State.DeleteUserState(msg.GetFrom().GetId())
	if err != nil {
		return errors.Wrap(err, "failed delete user state")
	}

	estimated, err := strconv.ParseFloat(SwapDB.AmountRecive, 64)
	if err != nil {
		return errors.Wrap(err, "failed calculate estimated")
	}
	data := pb.SwapSuccessResponse{
		User:       msg.GetFrom(),
		MsgId:      msgid,
		DepAddress: swap.AddressFrom,
		Id:         swap.ID,
		FromCur:    SwapDB.FromCurrency + " " + SwapDB.FromNetwork,
		ToCur:      SwapDB.ToCurrency + " " + SwapDB.ToNetwork,
		Amount:     amountF,
		Estimated:  estimated,
	}
	u, err := b.DB.User.Query().Where(user.ID(msg.GetFrom().GetId())).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	if !u.EnableLogging {
		err = b.DB.Swaps.DeleteOneID(trID).Exec(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed del swap")
		}
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed marshal")
	}

	b.ResponseChan <- tggateway.TgResponseMsg{
		Name: "SwapSuccessResponse",
		Data: dataBytes,
	}

	return nil

}
