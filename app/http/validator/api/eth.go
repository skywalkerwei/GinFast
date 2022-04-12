package validates_api

import (
	"ginfast/app/global/consts"
	"ginfast/app/http/controller/api"
	"ginfast/app/http/validator/core/data_transfer"
	"ginfast/app/utils/response"
	"github.com/gin-gonic/gin"
)

type EthCreate struct {
	TokenUrl string `form:"tokenUrl" json:"tokenUrl"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
}

func (n EthCreate) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&n); err != nil {
		response.ValidatorError(context, err)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(n, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "表单验证器json化失败", "")
	} else {
		(&api.Eth{}).CreateNft(extraAddBindDataContext)
	}
}

type EthTokenURI struct {
	TokenId int `form:"tokenId" json:"tokenId"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
}

func (n EthTokenURI) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&n); err != nil {
		response.ValidatorError(context, err)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(n, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "表单验证器json化失败", "")
	} else {
		(&api.Eth{}).TokenURI(extraAddBindDataContext)
	}
}

type EthSetTokenURI struct {
	TokenId  int    `form:"tokenId" json:"tokenId"  binding:"required,min=1"`   // 必填、对于文本,表示它的长度>=1
	TokenUrl string `form:"tokenUrl" json:"tokenUrl"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
}

func (n EthSetTokenURI) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&n); err != nil {
		response.ValidatorError(context, err)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(n, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "表单验证器json化失败", "")
	} else {
		(&api.Eth{}).SetTokenURI(extraAddBindDataContext)
	}
}

type EthCreateTokenByUser struct {
	PrivateKey string `form:"privateKey" json:"privateKey"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
	TokenUrl   string `form:"tokenUrl" json:"tokenUrl"  binding:"required,min=1"`     // 必填、对于文本,表示它的长度>=1
}

func (n EthCreateTokenByUser) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&n); err != nil {
		response.ValidatorError(context, err)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(n, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "表单验证器json化失败", "")
	} else {
		(&api.Eth{}).CreateTokenByUser(extraAddBindDataContext)
	}
}

type EthTransferFrom struct {
	From    string `form:"from" json:"from"  binding:"required,min=1"`       // 必填、对于文本,表示它的长度>=1
	To      string `form:"to" json:"to"  binding:"required,min=1"`           // 必填、对于文本,表示它的长度>=1
	TokenId int    `form:"tokenId" json:"tokenId"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
}

func (n EthTransferFrom) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&n); err != nil {
		response.ValidatorError(context, err)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(n, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "表单验证器json化失败", "")
	} else {
		(&api.Eth{}).TransferFrom(extraAddBindDataContext)
	}
}

type EthTransferFromByUser struct {
	FromPrivateKey string `form:"fromPrivateKey" json:"fromPrivateKey"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
	To             string `form:"to" json:"to"  binding:"required,min=1"`                         // 必填、对于文本,表示它的长度>=1
	TokenId        int    `form:"tokenId" json:"tokenId"  binding:"required,min=1"`               // 必填、对于文本,表示它的长度>=1
}

func (n EthTransferFromByUser) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&n); err != nil {
		response.ValidatorError(context, err)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(n, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "表单验证器json化失败", "")
	} else {
		(&api.Eth{}).TransferFromByUser(extraAddBindDataContext)
	}
}

type EthTransferEth struct {
	ToAddress string `form:"toAddress" json:"toAddress"  binding:"required,min=1"`
	Amount    int    `form:"amount" json:"amount"  binding:"required,min=1"`
}

func (n EthTransferEth) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&n); err != nil {
		response.ValidatorError(context, err)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(n, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "表单验证器json化失败", "")
	} else {
		(&api.Eth{}).TransferEth(extraAddBindDataContext)
	}
}
