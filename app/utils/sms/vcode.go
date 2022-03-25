package sms

import (
	"bytes"
	"errors"
	"ginfast/app/utils/redis_factory"
	"ginfast/app/utils/yml_config"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type VCode struct {
	Template string
	Mobile   string
}

func New(Template string, Mobile string) *VCode {
	return &VCode{Template, Mobile}
}

// Send 发送
func (v *VCode) Send() error {
	//debug 状态不发送不校验
	debug := yml_config.CreateYamlFactory().GetBool("VCode.debug")
	if debug {
		return nil
	}
	//testUsers 不发送不校验
	testUsers := yml_config.CreateYamlFactory().GetString("VCode.testUsers")
	//testMobiles := strings.Split(testUsers, "|")
	//for _, testMobile := range testMobiles {
	//	if testMobile == v.Mobile {
	//		return true, nil
	//	}
	//}
	index := strings.Contains(testUsers, v.Mobile)
	if index {
		return nil
	}

	redisClient := redis_factory.GetOneRedisClient()
	defer redisClient.ReleaseOneRedisClient()
	
	key := v.getKey()
	isExists, _ := redisClient.Exists(key)
	if isExists {
		return errors.New("验证码已发送，请勿重复请求")
	}
	//缓存
	code := RandCode(yml_config.CreateYamlFactory().GetInt("VCode.length"))
	life := yml_config.CreateYamlFactory().GetInt64("VCode.life")
	err := redisClient.Set(key, code, life)
	if err != nil {
		return err
	}
	_ = redisClient.Set(key+"_checkTimes", 0, life)

	//发送短信
	aliSms := AliSms{}
	return aliSms.SendCode(v.Template, v.Mobile, code)
}

// Check 验证
func (v *VCode) Check(code string) error {
	//debug 状态不发送不校验
	debug := yml_config.CreateYamlFactory().GetBool("VCode.debug")
	if debug {
		return nil
	}
	//testUsers 不发送不校验
	testUsers := yml_config.CreateYamlFactory().GetString("VCode.testUsers")
	index := strings.Contains(testUsers, v.Mobile)
	if index {
		return nil
	}
	// 魔法密码直接放过
	magicCode := yml_config.CreateYamlFactory().GetString("VCode.magicCode")
	if magicCode == code {
		return nil
	}
	// 正常校验
	key := v.getKey()

	redisClient := redis_factory.GetOneRedisClient()
	defer redisClient.ReleaseOneRedisClient()

	vCode, err := redisClient.GetString(key)
	if err != nil {
		return errors.New("验证码有误")
	}
	if vCode != code {
		_, _ = redisClient.Incr(key + "_checkTimes")
		maxCheckTimes := yml_config.CreateYamlFactory().GetInt64("VCode.maxCheckTimes")
		checkTimes, _ := redisClient.GetInt64(key + "_checkTimes")
		if maxCheckTimes < checkTimes {
			_ = redisClient.Del(key)
			_ = redisClient.Del(key + "_checkTimes")
			return errors.New("校验次数已超过最大校验次数")
		}
		return errors.New("验证码有误")
	}
	//
	_ = redisClient.Del(key)
	_ = redisClient.Del(key + "_checkTimes")

	return nil
}

func (v *VCode) getKey() string {
	return v.Template + "_validate_code_" + v.Mobile
}

//RandCode 生成随机数
func RandCode(length int) string {
	randNum := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(math.Pow10(length)))
	s := bytes.Buffer{}
	s.WriteString(strconv.Itoa(int(randNum)))
	return s.String()
}
