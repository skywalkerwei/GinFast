package nft

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
)

func CreateContractT() *ContractT {
	//return &ContractT{
	//	DialURL:         "http://127.0.0.1:8545",
	//	ContractAddress: "0xbD8d0BaF93f62D9cC41DfdE374d245c62cB5f78B",
	//	MainPrivateKey:  "22884c29fc4fce16956cf0aac1e1b65c7dd45ff6c0f934e09bbb0a80f81a6e24",
	//	ChainID:         new(big.Int).SetInt64(1337),
	//}
	return &ContractT{
		DialURL:         "https://testnet.emerald.oasis.dev",
		ContractAddress: "0xB12845aD394832C01630960d3D9359c6ebd1e415",
		MainPrivateKey:  "5b504c58b38904ab0695ccf187fcffb30bc8f14dd5110c17934d32c1a396ecd5",
		ChainID:         new(big.Int).SetInt64(42261),
	}
}

type ContractT struct {
	DialURL         string
	ContractAddress string
	MainPrivateKey  string
	ChainID         *big.Int
}

func (c *ContractT) Client() *ethclient.Client {
	//http://127.0.0.1:8545
	client, err := ethclient.Dial(c.DialURL)
	if err != nil {
		fmt.Println(err)
	}
	return client
}

func (c *ContractT) ConnectContract() *NftAbi {
	client := c.Client()
	defer client.Close()
	//创建合约
	contract, err := NewNftAbi(common.HexToAddress(c.ContractAddress), client)
	if err != nil {
		fmt.Print(err)
	}
	return contract
}

func (c *ContractT) TotalSupply() *big.Int {
	contract := c.ConnectContract()
	totalSupply, err := contract.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return totalSupply
	}
	return totalSupply
}

func (c *ContractT) TokenURI(tokenId int64) string {
	contract := c.ConnectContract()
	tokenUrl, err := contract.TokenURI(&bind.CallOpts{}, big.NewInt(tokenId))
	fmt.Println(tokenId, tokenUrl, err)
	if err != nil {
		return tokenUrl
	}
	return tokenUrl
}

func (c *ContractT) Auth(privateKey string) *bind.TransactOpts {

	ecdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(ecdsa, c.ChainID)
	if err != nil {
		fmt.Println(err)
	}
	return auth
}

func (c *ContractT) CreateToken(tokenUrl string) (string, int) {
	client := c.Client()
	defer client.Close()
	//创建合约
	contract, err := NewNftAbi(common.HexToAddress(c.ContractAddress), client)
	if err != nil {
		fmt.Print(err)
	}
	auth := c.Auth(c.MainPrivateKey)
	tx, err := contract.CreateToken(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  big.NewInt(0),
	}, tokenUrl)
	if err != nil {
		fmt.Println("err", err)
	}
	ctx := context.Background()
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		fmt.Println("err", err)
	}
	json, _ := receipt.MarshalJSON()

	fmt.Println("receipt", receipt, string(json))
	tokenID := 0
	//for _, log := range receipt.Logs {
	//	fmt.Println("log", log)
	//	fmt.Println("log.Topics", log.Topics)
	for _, topic := range receipt.Logs[0].Topics {
		fmt.Printf("topic is :%s\n", topic.String())
		tokenID = int(topic.Big().Int64())
	}
	//}
	return tx.Hash().Hex(), tokenID
	//for i, b := range json {
	//	fmt.Printf("%d:%d\n", i, b)
	//}
	//return tx.Hash().String()
	//fmt.Printf("receipt is :%s\n", receipt.TxHash)
	//return tx.Hash().Hex()
}

func (c *ContractT) CreateTokenByUser(privateKey string, tokenUrl string) string {
	contract := c.ConnectContract()
	auth := c.Auth(privateKey)
	tx, err := contract.CreateToken(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  big.NewInt(0),
	}, tokenUrl)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("tx sent:", tx.Hash().Hex())
	return tx.Hash().Hex()
}

func (c *ContractT) SetTokenURI(tokenId int64, tokenUrl string) string {
	contract := c.ConnectContract()
	auth := c.Auth(c.MainPrivateKey)
	tx, err := contract.SetTokenURI(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  big.NewInt(0),
	}, big.NewInt(tokenId), tokenUrl)
	if err != nil {
		fmt.Println("err", err)
		return ""
	}
	return tx.Hash().Hex()
}

func (c *ContractT) TransferFrom(from string, to string, tokenId int64) string {
	contract := c.ConnectContract()
	auth := c.Auth(c.MainPrivateKey)
	tx, err := contract.TransferFrom(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  big.NewInt(0),
	}, common.HexToAddress(from), common.HexToAddress(to), big.NewInt(tokenId))
	if err != nil {
		fmt.Println("err", err)
		return ""
	}
	return tx.Hash().Hex()
}

func (c *ContractT) TransferEth(toAddress string, amount int64) string {
	contract := c.ConnectContract()
	auth := c.Auth(c.MainPrivateKey)
	tx, err := contract.Eth(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  big.NewInt(amount),
	}, common.HexToAddress(toAddress))
	if err != nil {
		fmt.Println("err", err)
		return ""
	}
	fmt.Println("tx sent:", tx.Hash().Hex())
	return tx.Hash().Hex()
}

func (c *ContractT) TransferFromByUser(fromPrivateKey string, to string, tokenId int64) string {
	contract := c.ConnectContract()
	auth := c.Auth(fromPrivateKey)
	tx, err := contract.TransferFrom(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  big.NewInt(0),
	}, auth.From, common.HexToAddress(to), big.NewInt(tokenId))
	if err != nil {
		fmt.Println("err", err)
		return ""
	}
	return tx.Hash().Hex()
}

func (c *ContractT) getBalance(address string) string {
	client := c.Client()
	defer client.Close()
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return ""
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue.String()
}

func (c *ContractT) TransferFormEth(toAddress string, amount int64) error {

	client := c.Client()
	defer client.Close()
	auth := c.Auth(c.MainPrivateKey)
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	toAddr := common.HexToAddress(toAddress)
	data := make([]byte, 0)
	gasLimit := uint64(21000)

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddr,
		Value:    big.NewInt(amount),
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	ecdsa, err := crypto.HexToECDSA(c.MainPrivateKey)
	if err != nil {
		fmt.Println(err)
	}

	signTx, err := types.SignTx(tx, types.NewEIP155Signer(c.ChainID), ecdsa)
	if err != nil {
		return err
	}

	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		return err
	}
	fmt.Println("signTx", err, signTx)

	return nil
}

//EthToWei
//https://stackoverrun.com/cn/q/13021596
func EthToWei(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	// Set precision if required.
	// bigval.SetPrec(64)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}
