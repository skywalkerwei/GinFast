package validates_api

import (
	"ginfast/app/global/consts"
	"ginfast/app/http/controller/api"
	"ginfast/app/http/validator/common/data_type"
	"ginfast/app/http/validator/core/data_transfer"
	"ginfast/app/utils/response"
	"github.com/gin-gonic/gin"
)

type Login struct {
	data_type.Mobile
	Code string `form:"code" json:"code"  binding:"required,min=1"`
}

func (n Login) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&n); err != nil {
		response.ValidatorError(context, err)
		return
	}
	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := data_transfer.DataAddContext(n, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&api.Auth{}).Login(extraAddBindDataContext)
	}

}

//
//type MRefreshToken struct {
//	Authorization string `json:"token" header:"Authorization" binding:"required,min=20"`
//}
//
//func (n MRefreshToken) CheckParams(context *gin.Context) {
//	if err := context.ShouldBind(&n); err != nil {
//		response.ValidatorError(context, err)
//		return
//	}
//	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
//	extraAddBindDataContext := data_transfer.DataAddContext(n, consts.ValidatorPrefix, context)
//	if extraAddBindDataContext == nil {
//		response.ErrorSystem(context, "表单验证器json化失败", "")
//	} else {
//		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
//		//(&api.Auth{}).MToken(extraAddBindDataContext)
//	}
//}
