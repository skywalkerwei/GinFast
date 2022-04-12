package validates_api

import (
	"ginfast/app/global/consts"
	"ginfast/app/http/controller/api"
	"ginfast/app/http/validator/core/data_transfer"
	"ginfast/app/utils/response"
	"github.com/gin-gonic/gin"
)

type WxLogin struct {
	Code string `form:"code" json:"code"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
}

func (n WxLogin) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&n); err != nil {
		response.ValidatorError(context, err)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(n, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "表单验证器json化失败", "")
	} else {
		(&api.Wx{}).Wx(extraAddBindDataContext)
	}
}

type WxBind struct {
	EncryptedData string `form:"encryptedData" json:"encryptedData"  binding:"required,min=1"`
	Iv            string `form:"iv" json:"iv"  binding:"required,min=1"`
	Openid        string `form:"openid" json:"openid"  binding:"required,min=1"`
}

func (n WxBind) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&n); err != nil {
		response.ValidatorError(context, err)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(n, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "表单验证器json化失败", "")
	} else {
		(&api.Wx{}).Bind(extraAddBindDataContext)
	}

}
