package crypto

import (
	"fmt"
	"log"
	"math/big"

	"github.com/make-software/casper-go-sdk/casper"
	"github.com/make-software/casper-go-sdk/types"
	"github.com/make-software/casper-go-sdk/types/clvalue"
	"github.com/make-software/casper-go-sdk/types/key"
	"github.com/pkg/errors"
)

const AuctionContractTestnet string = "93d923e336b20a4c4ca14d592b60e5bd3fe330775618290104f9beb326db7ae2"
const AuctionContractMainnet string = "86f2d45f024d7bb7fb5266b2390d7c253b588a0a16ebd946a60cb4314600af74"

func (c *Crypto) GetAuctionContractHash() string {
	if c.Chain == "casper-test" {
		return AuctionContractTestnet
	}
	return AuctionContractMainnet
}

type Transfer struct {
	ToPubkey string
	Amount   string
	Memo     uint64
}

type Delegate struct {
	Validator string
	Delegator string
	Amount    string
}

func (c *Crypto) SignTransferWithPassword(trData Transfer, uid int64, password string) (*types.Deploy, error) {
	facetKeys, err := c.DecodePemFromDB(uid, password)
	if err != nil {
		return nil, err
	}
	accountPublicKey := facetKeys.PublicKey()
	log.Println(accountPublicKey.String())

	targetPubKey, err := casper.NewPublicKey(trData.ToPubkey)
	if err != nil {
		log.Println(err)
	}
	amount := new(big.Int)
	amount, ok := amount.SetString(trData.Amount, 10)
	if !ok {
		fmt.Println("SetString: error")
		return nil, errors.New("failed parse amount")
	}

	session := casper.ExecutableDeployItem{
		Transfer: &types.TransferDeployItem{
			Args: *(&casper.Args{}).AddArgument("target", clvalue.NewCLByteArray(targetPubKey.AccountHash().Bytes())).
				AddArgument("amount", *clvalue.NewCLUInt512(amount)).
				AddArgument("id", clvalue.NewCLOption(*clvalue.NewCLUInt64(trData.Memo))),
		},
	}

	payment := casper.StandardPayment(amount)

	deployHeader := casper.DefaultHeader()
	deployHeader.Account = accountPublicKey
	deployHeader.ChainName = c.Chain

	newDeploy, err := casper.MakeDeploy(deployHeader, payment, session)
	if err != nil {
		log.Println(err)
	}
	err = newDeploy.SignDeploy(*facetKeys)
	if err != nil {
		log.Println(err)
	}
	return newDeploy, err
}

func (c *Crypto) SignTransferWithPK(trData Transfer, uid int64, pk string) (*types.Deploy, error) {
	facetKeys, err := c.NewKeypairFromPEM(uid, []byte(pk))
	if err != nil {
		return nil, err
	}
	accountPublicKey := facetKeys.PublicKey()
	log.Println(accountPublicKey.String())

	targetPubKey, err := casper.NewPublicKey(trData.ToPubkey)
	if err != nil {
		log.Println(err)
	}
	amount := new(big.Int)
	amount, ok := amount.SetString(trData.Amount, 10)
	if !ok {
		fmt.Println("SetString: error")
		return nil, errors.New("failed parse amount")
	}

	session := casper.ExecutableDeployItem{
		Transfer: &types.TransferDeployItem{
			Args: *(&casper.Args{}).AddArgument("target", clvalue.NewCLByteArray(targetPubKey.AccountHash().Bytes())).
				AddArgument("amount", *clvalue.NewCLUInt512(amount)).
				AddArgument("id", clvalue.NewCLOption(*clvalue.NewCLUInt64(trData.Memo))),
		},
	}

	payment := casper.StandardPayment(amount)

	deployHeader := casper.DefaultHeader()
	deployHeader.Account = accountPublicKey
	deployHeader.ChainName = c.Chain

	newDeploy, err := casper.MakeDeploy(deployHeader, payment, session)
	if err != nil {
		log.Println(err)
	}
	err = newDeploy.SignDeploy(*facetKeys)
	if err != nil {
		log.Println(err)
	}
	return newDeploy, err
}

func (c *Crypto) SignDelegateWithPassword(trData Delegate, uid int64, password string) (*types.Deploy, error) {
	facetKeys, err := c.DecodePemFromDB(uid, password)
	if err != nil {
		return nil, err
	}
	accountPublicKey := facetKeys.PublicKey()
	log.Println(accountPublicKey.String())

	contractHash, err := key.NewContract(c.GetAuctionContractHash())
	if err != nil {
		return nil, err
	}
	validator, err := casper.NewPublicKey(trData.Validator)
	if err != nil {
		return nil, err
	}
	amount := new(big.Int)
	amount, ok := amount.SetString(trData.Amount, 10)
	if !ok {
		fmt.Println("SetString: error")
		return nil, errors.New("failed parse amount")
	}

	contract := casper.StoredContractByHash{
		Hash:       contractHash,
		EntryPoint: "delegate",
		Args: (&casper.Args{}).
			AddArgument("validator", clvalue.NewCLPublicKey(validator)).
			AddArgument("delegator", clvalue.NewCLPublicKey(accountPublicKey)).
			AddArgument("amount", *clvalue.NewCLUInt512(amount)),
	}
	session := casper.ExecutableDeployItem{
		StoredContractByHash: &contract,
	}

	payment := casper.StandardPayment(big.NewInt(2500000000))

	deployHeader := casper.DefaultHeader()
	deployHeader.Account = accountPublicKey
	deployHeader.ChainName = c.Chain

	newDeploy, err := casper.MakeDeploy(deployHeader, payment, session)
	if err != nil {
		log.Println(err)
	}
	err = newDeploy.SignDeploy(*facetKeys)
	if err != nil {
		log.Println(err)
	}
	return newDeploy, err
}

func (c *Crypto) SignDelegateWithPK(trData Delegate, uid int64, pk string) (*types.Deploy, error) {
	facetKeys, err := c.NewKeypairFromPEM(uid, []byte(pk))
	if err != nil {
		return nil, err
	}
	accountPublicKey := facetKeys.PublicKey()
	log.Println(accountPublicKey.String())

	contractHash, err := key.NewContract(c.GetAuctionContractHash())
	if err != nil {
		return nil, err
	}
	validator, err := casper.NewPublicKey(trData.Validator)
	if err != nil {
		return nil, err
	}
	amount := new(big.Int)
	amount, ok := amount.SetString(trData.Amount, 10)
	if !ok {
		fmt.Println("SetString: error")
		return nil, errors.New("failed parse amount")
	}

	contract := casper.StoredContractByHash{
		Hash:       contractHash,
		EntryPoint: "delegate",
		Args: (&casper.Args{}).
			AddArgument("validator", clvalue.NewCLPublicKey(validator)).
			AddArgument("delegator", clvalue.NewCLPublicKey(accountPublicKey)).
			AddArgument("amount", *clvalue.NewCLUInt512(amount)),
	}
	session := casper.ExecutableDeployItem{
		StoredContractByHash: &contract,
	}

	payment := casper.StandardPayment(big.NewInt(2500000000))

	deployHeader := casper.DefaultHeader()
	deployHeader.Account = accountPublicKey
	deployHeader.ChainName = c.Chain

	newDeploy, err := casper.MakeDeploy(deployHeader, payment, session)
	if err != nil {
		log.Println(err)
	}
	err = newDeploy.SignDeploy(*facetKeys)
	if err != nil {
		log.Println(err)
	}
	return newDeploy, err
}

func (c *Crypto) SignUndelegateWithPassword(trData Delegate, uid int64, password string) (*types.Deploy, error) {
	facetKeys, err := c.DecodePemFromDB(uid, password)
	if err != nil {
		return nil, err
	}
	accountPublicKey := facetKeys.PublicKey()
	log.Println(accountPublicKey.String())

	contractHash, err := key.NewContract(c.GetAuctionContractHash())
	if err != nil {
		return nil, err
	}
	validator, err := casper.NewPublicKey(trData.Validator)
	if err != nil {
		return nil, err
	}
	amount := new(big.Int)
	amount, ok := amount.SetString(trData.Amount, 10)
	if !ok {
		fmt.Println("SetString: error")
		return nil, errors.New("failed parse amount")
	}

	contract := casper.StoredContractByHash{
		Hash:       contractHash,
		EntryPoint: "undelegate",
		Args: (&casper.Args{}).
			AddArgument("validator", clvalue.NewCLPublicKey(validator)).
			AddArgument("delegator", clvalue.NewCLPublicKey(accountPublicKey)).
			AddArgument("amount", *clvalue.NewCLUInt512(amount)),
	}
	session := casper.ExecutableDeployItem{
		StoredContractByHash: &contract,
	}

	payment := casper.StandardPayment(big.NewInt(2500000000))

	deployHeader := casper.DefaultHeader()
	deployHeader.Account = accountPublicKey
	deployHeader.ChainName = c.Chain

	newDeploy, err := casper.MakeDeploy(deployHeader, payment, session)
	if err != nil {
		log.Println(err)
	}
	err = newDeploy.SignDeploy(*facetKeys)
	if err != nil {
		log.Println(err)
	}
	return newDeploy, err
}

func (c *Crypto) SignUndelegateWithPK(trData Delegate, uid int64, pk string) (*types.Deploy, error) {
	facetKeys, err := c.NewKeypairFromPEM(uid, []byte(pk))
	if err != nil {
		return nil, err
	}
	accountPublicKey := facetKeys.PublicKey()
	log.Println(accountPublicKey.String())

	contractHash, err := key.NewContract(c.GetAuctionContractHash())
	if err != nil {
		return nil, err
	}
	validator, err := casper.NewPublicKey(trData.Validator)
	if err != nil {
		return nil, err
	}
	amount := new(big.Int)
	amount, ok := amount.SetString(trData.Amount, 10)
	if !ok {
		fmt.Println("SetString: error")
		return nil, errors.New("failed parse amount")
	}

	contract := casper.StoredContractByHash{
		Hash:       contractHash,
		EntryPoint: "undelegate",
		Args: (&casper.Args{}).
			AddArgument("validator", clvalue.NewCLPublicKey(validator)).
			AddArgument("delegator", clvalue.NewCLPublicKey(accountPublicKey)).
			AddArgument("amount", *clvalue.NewCLUInt512(amount)),
	}
	session := casper.ExecutableDeployItem{
		StoredContractByHash: &contract,
	}

	payment := casper.StandardPayment(big.NewInt(2500000000))

	deployHeader := casper.DefaultHeader()
	deployHeader.Account = accountPublicKey
	deployHeader.ChainName = c.Chain

	newDeploy, err := casper.MakeDeploy(deployHeader, payment, session)
	if err != nil {
		log.Println(err)
	}
	err = newDeploy.SignDeploy(*facetKeys)
	if err != nil {
		log.Println(err)
	}
	return newDeploy, err
}
