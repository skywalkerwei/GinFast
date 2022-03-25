package websocket

import (
	"ginfast/app/global/consts"
	"ginfast/app/global/variable"
	controllerWs "ginfast/app/http/controller/websocket"
	"ginfast/app/http/validator/core/data_transfer"
	userstoken "ginfast/app/service/token"
	"ginfast/app/utils/response"
	"github.com/gin-gonic/gin"
)

type Connect struct {
	Token string `form:"token" binding:"required,min=10"`
}

// 验证器语法，参见 Register.go文件，有详细说明

func (c Connect) CheckParams(context *gin.Context) {

	// 1. 首先检查是否开启websocket服务配置（在配置项中开启）
	if variable.ConfigYml.GetInt("Websocket.Start") != 1 {
		response.Fail(context, consts.WsServerNotStartCode, consts.WsServerNotStartMsg, "")
		return
	}
	//2.基本的验证规则没有通过
	if err := context.ShouldBind(&c); err != nil {
		errs := gin.H{
			"tips": "请在get参数中提交token信息,demo格式：ws://127.0.0.1:6001?ws?token=this is a series token...",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(c, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "websocket-Connect 表单验证器json化失败", "")
		context.Abort()
		return
	} else {
		token := extraAddBindDataContext.GetString(consts.ValidatorPrefix + "Token")
		uid := checkToken(token)
		if uid == 0 {
			response.ErrorSystem(context, "用户token不合法", "")
			context.Abort()
			return
		}
		extraAddBindDataContext.Set(consts.ValidatorPrefix+"uid", uid)
		if serviceWs, ok := (&controllerWs.Ws{}).OnOpen(extraAddBindDataContext); ok == false {
			response.Fail(context, consts.WsOpenFailCode, consts.WsOpenFailMsg, "")
		} else {
			(&controllerWs.Ws{}).OnMessage(serviceWs, extraAddBindDataContext) // 注意这里传递的service_ws必须是调用open返回的，必须保证的ws对象的一致性
		}
	}
}

func checkToken(token string) (uid int64) {

	if customToken, err := userstoken.CreateUserFactory().ParseToken(token); err == nil {
		uid = customToken.UserId
	}

	return uid
}
