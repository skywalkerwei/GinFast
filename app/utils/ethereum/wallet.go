package ethereum

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct{}

// Create @Title CreateWallet
// @Description Create a new wallet
func (*Wallet) Create() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", errors.New("privateKey Failed, err:" + err.Error())
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyStr := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return privateKeyStr, address, nil
}
