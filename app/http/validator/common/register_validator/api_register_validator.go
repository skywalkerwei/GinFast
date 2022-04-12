package register_validator

import (
	"ginfast/app/core/container"
	"ginfast/app/global/consts"
	"ginfast/app/http/validator/api"
)

// ApiRegisterValidator
//各个业务模块验证器必须进行注册（初始化），程序启动时会自动加载到容器
func ApiRegisterValidator() {
	//创建容器
	containers := container.CreateContainersFactory()

	//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
	var key string

	// 注册门户类表单参数验证器
	key = consts.ValidatorPrefix + "SendCode"
	containers.Set(key, validates_api.SendCode{})

	key = consts.ValidatorPrefix + "Login"
	containers.Set(key, validates_api.Login{})

	key = consts.ValidatorPrefix + "EthCreate"
	containers.Set(key, validates_api.EthCreate{})

	key = consts.ValidatorPrefix + "EthTokenURI"
	containers.Set(key, validates_api.EthTokenURI{})

	key = consts.ValidatorPrefix + "EthSetTokenURI"
	containers.Set(key, validates_api.EthSetTokenURI{})

	key = consts.ValidatorPrefix + "EthCreateTokenByUser"
	containers.Set(key, validates_api.EthCreateTokenByUser{})

	key = consts.ValidatorPrefix + "EthTransferFrom"
	containers.Set(key, validates_api.EthTransferFrom{})

	key = consts.ValidatorPrefix + "EthTransferFromByUser"
	containers.Set(key, validates_api.EthTransferFromByUser{})

	key = consts.ValidatorPrefix + "EthTransferEth"
	containers.Set(key, validates_api.EthTransferEth{})

	//key = consts.ValidatorPrefix + "MRefreshToken"
	//containers.Set(key, validates.MRefreshToken{})

	//key = consts.ValidatorPrefix + "WxLogin"
	//containers.Set(key, validates.WxLogin{})

	//key = consts.ValidatorPrefix + "WxBind"
	//containers.Set(key, validates.WxBind{})

}
