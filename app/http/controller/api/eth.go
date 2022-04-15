package api

import (
	"ginfast/app/global/consts"
	"ginfast/app/utils/ethereum"
	"ginfast/app/utils/ethereum/nft"
	"ginfast/app/utils/response"
	"github.com/gin-gonic/gin"
)

type Eth struct {
}

// Wallet
// @Tags Eth
// @Summary 生成钱包
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /eth/wallet [Get]
func (e *Eth) Wallet(c *gin.Context) {

	wallet := ethereum.Wallet{}
	privateKeyStr, address, err := wallet.Create()
	if err != nil {
		response.Fail(c, consts.CurdCreatFailCode, err.Error(), "")
		return
	}
	response.Success(c, "ok", gin.H{
		"private": privateKeyStr,
		"address": address,
	})

}

// CreateNft
// @Tags Eth
// @Summary 生成NFT
// @Param tokenUrl formData string true "tokenUrl"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /eth/crateNft [Post]
func (e *Eth) CreateNft(c *gin.Context) {
	tokenUrl := c.GetString(consts.ValidatorPrefix + "tokenUrl")
	nftC := nft.CreateContractT()
	hashCode, tokenID := nftC.CreateToken(tokenUrl)
	response.Success(c, "ok", gin.H{
		"tokenID":  tokenID,
		"hashCode": hashCode,
	})

}

// TokenURI
// @Tags Eth
// @Summary meta信息
// @Param tokenId query string true "tokenId"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /eth/tokenURI [get]
func (e *Eth) TokenURI(c *gin.Context) {
	tokenId := c.GetFloat64(consts.ValidatorPrefix + "tokenId")
	nftC := nft.CreateContractT()
	response.Success(c, "ok", nftC.TokenURI(int64(tokenId)))
}

// TotalSupply
// @Tags Eth
// @Summary 总数
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /eth/totalSupply [get]
func (e *Eth) TotalSupply(c *gin.Context) {
	nftC := nft.CreateContractT()
	response.Success(c, "ok", nftC.TotalSupply())
}

// SetTokenURI
// @Tags Eth
// @Summary 设置meta信息
// @Param tokenId formData string true "tokenId"
// @Param tokenUrl formData string true "tokenUrl"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /eth/setTokenURI [Post]
func (e *Eth) SetTokenURI(c *gin.Context) {
	tokenId := c.GetFloat64(consts.ValidatorPrefix + "tokenId")
	tokenUrl := c.GetString(consts.ValidatorPrefix + "tokenUrl")
	nftC := nft.CreateContractT()
	response.Success(c, "ok", nftC.SetTokenURI(int64(tokenId), tokenUrl))
}

// CreateTokenByUser
// @Tags Eth
// @Summary 用户创建NFT
// @Param privateKey formData string true "privateKey"
// @Param tokenUrl formData string true "tokenUrl"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /eth/createTokenByUser [Post]
func (e *Eth) CreateTokenByUser(c *gin.Context) {
	privateKey := c.GetString(consts.ValidatorPrefix + "privateKey")
	tokenUrl := c.GetString(consts.ValidatorPrefix + "tokenUrl")
	nftC := nft.CreateContractT()
	response.Success(c, "ok", nftC.CreateTokenByUser(privateKey, tokenUrl))
}

// TransferFrom
// @Tags Eth
// @Summary 转移nft
// @Param from formData string true "from"
// @Param to formData string true "to"
// @Param tokenId formData string true "tokenId"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /eth/transferFrom [Post]
func (e *Eth) TransferFrom(c *gin.Context) {
	from := c.GetString(consts.ValidatorPrefix + "from")
	to := c.GetString(consts.ValidatorPrefix + "to")
	tokenId := c.GetFloat64(consts.ValidatorPrefix + "tokenId")
	nftC := nft.CreateContractT()
	response.Success(c, "ok", nftC.TransferFrom(from, to, int64(tokenId)))
}

// TransferFromByUser
// @Tags Eth
// @Summary 用户转移nft
// @Param fromPrivateKey formData string true "fromPrivateKey"
// @Param to formData string true "to"
// @Param tokenId formData string true "tokenId"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /eth/transferFromByUser [Post]
func (e *Eth) TransferFromByUser(c *gin.Context) {
	fromPrivateKey := c.GetString(consts.ValidatorPrefix + "fromPrivateKey")
	to := c.GetString(consts.ValidatorPrefix + "to")
	tokenId := c.GetFloat64(consts.ValidatorPrefix + "tokenId")
	nftC := nft.CreateContractT()
	response.Success(c, "ok", nftC.TransferFromByUser(fromPrivateKey, to, int64(tokenId)))
}

// TransferEth
// @Tags Eth
// @Summary 转移eth
// @Param toAddress formData string true "toAddress"
// @Param amount formData string true "amount"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /eth/transferEth [Post]
func (e *Eth) TransferEth(c *gin.Context) {
	toAddress := c.GetString(consts.ValidatorPrefix + "toAddress")
	amount := c.GetFloat64(consts.ValidatorPrefix + "amount")
	nftC := nft.CreateContractT()
	response.Success(c, "ok", nftC.TransferEth(toAddress, int64(amount)))
}
