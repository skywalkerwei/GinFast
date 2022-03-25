package api

import (
	"ginfast/app/global/consts"
	"ginfast/app/global/variable"
	"ginfast/app/http/middleware/my_jwt"
	usersToken "ginfast/app/service/token"
	"ginfast/app/service/user"
	"ginfast/app/utils/response"
	"ginfast/app/utils/sms"
	"github.com/gin-gonic/gin"
	"strings"
)

type Auth struct {
}

// Login
// @Tags Auth
// @Summary 获取token
// @Produce json
// @Param mobile formData string true "手机号码"
// @Param code formData string true "验证码"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/login [post]
func (u *Auth) Login(context *gin.Context) {
	mobile := context.GetString(consts.ValidatorPrefix + "mobile")
	code := context.GetString(consts.ValidatorPrefix + "code")
	err := sms.New("SMS_219510061", mobile).Check(code)
	if err != nil {
		response.Fail(context, consts.CurdCreatFailCode, err.Error(), "")
		return
	}
	//
	userService := user_service.User{}
	data, err := userService.Token(mobile, context.ClientIP())
	if err != nil {
		response.Fail(context, consts.CurdCreatFailCode, err.Error(), "")
		return
	}
	response.Success(context, consts.CurdStatusOkMsg, data)
	return

}

// GetAuthInfo 获取用户-强制登录
func GetAuthInfo(context *gin.Context) my_jwt.CustomClaims {
	auth, isAuth := context.Get(variable.ConfigYml.GetString("Token.BindContextKeyName"))
	if isAuth {
		return auth.(my_jwt.CustomClaims)
	}
	return my_jwt.CustomClaims{}
}

// GetUserInfo 获取用户-非强制登录
func GetUserInfo(context *gin.Context) (bool, my_jwt.CustomClaims) {
	authorization := context.Request.Header.Get("Authorization")
	token := strings.Split(authorization, " ")
	if len(token) == 2 && len(token[1]) >= 20 {
		tokenIsEffective := usersToken.CreateUserFactory().IsEffective(token[1])
		if tokenIsEffective {
			if customToken, err := usersToken.CreateUserFactory().ParseToken(token[1]); err == nil {
				return true, customToken
			}
		}
	}
	return false, my_jwt.CustomClaims{}
}
