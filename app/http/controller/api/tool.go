package api

import (
	"ginfast/app/global/consts"
	"ginfast/app/utils/response"
	"ginfast/app/utils/sms"
	"github.com/gin-gonic/gin"
)

type Tool struct {
}

// SendCode
// @Tags 工具
// @Summary 发送验证码
// @Produce json
// @Param mobile formData string true "mobile"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /tool/send_code [post]
func (t *Tool) SendCode(c *gin.Context) {
	mobile := c.PostForm("mobile")
	err := sms.New("SMS_219510061", mobile).Send()
	if err != nil {
		response.Fail(c, consts.CurdCreatFailCode, err.Error(), "")
	} else {
		response.Success(c, consts.CurdStatusOkMsg, consts.CurdStatusOkMsg)
	}

}
