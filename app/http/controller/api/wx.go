package api

import (
	"fmt"
	"ginfast/app/global/consts"
	"ginfast/app/utils/response"
	"ginfast/app/utils/wechat"
	"github.com/gin-gonic/gin"
)

type Wx struct {
}

func (u *Wx) Wx(context *gin.Context) {

	code := context.GetString(consts.ValidatorPrefix + "code")

	session, err := wechat.Code2Session(code)

	if err != nil {
		fmt.Println("error", err, code)
		return
	}
	response.Success(context, "ok", session)
	//
}

func (u *Wx) Bind(context *gin.Context) {

	encryptedData := context.GetString(consts.ValidatorPrefix + "encryptedData")
	iv := context.GetString(consts.ValidatorPrefix + "iv")
	openID := context.GetString(consts.ValidatorPrefix + "openid")

	decrypt, err := wechat.Decrypt(openID, encryptedData, iv)
	if err != nil {
		return
	}
	response.Success(context, "ok", decrypt)

}
