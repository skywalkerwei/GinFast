package user_service

import (
	"errors"
	"ginfast/app/global/variable"
	"ginfast/app/model"
	usersToken "ginfast/app/service/token"
	"github.com/gin-gonic/gin"
	"time"
)

type User struct{}

func (u User) Token(mobile string, clientIP string) (data gin.H, err error) {
	m := model.CreateUsersFactory()
	var userModel = model.UserModel{}
	err = m.Where("mobile = ?", mobile).First(&userModel).Error
	if err != nil {
		return nil, err
	}
	if userModel.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	role := "u"
	if userToken, err := usersToken.CreateUserFactory().GenerateToken(userModel.ID, userModel.Name, userModel.Mobile, role, 0, variable.ConfigYml.GetInt64("Token.JwtTokenCreatedExpireAt")); err == nil {
		data = gin.H{
			"uid":      userModel.ID,
			"nickname": userModel.Name,
			"mobile":   userModel.Mobile,
			"token":    "Bearer " + userToken,
		}
		go m.Model(m).Where("id = ?", userModel.ID).Updates(map[string]interface{}{
			"login_at": time.Now(),
			"login_ip": clientIP})
		return data, nil

	}
	return nil, errors.New("用户不存在")
}
