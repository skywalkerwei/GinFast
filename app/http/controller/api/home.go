package api

import (
	"github.com/gin-gonic/gin"
)

type Home struct {
}

func (u *Home) Tt(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (u *Home) Tt2(c *gin.Context) {

	//privateKey, err := crypto.GenerateKey()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//privateKeyBytes := crypto.FromECDSA(privateKey)
	//fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19
	//
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}
	//
	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05
	//
	//address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	//url := "https://testnets-api.opensea.io/api/v1/asset/0x6b5e78f92c4894f833a6150388869d6fa2f925d4/10"
	//req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Add("Accept", "application/json")
	//res, _ := http.DefaultClient.Do(req)
	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(res)
	//fmt.Println(string(body))

	//cli := goCurl.CreateHttpClient()
	//resp, err := cli.Post("https://testnets-api.opensea.io/api/v1/asset/0x6b5e78f92c4894f833a6150388869d6fa2f925d4/10",
	//	goCurl.Options{
	//		Headers: map[string]interface{}{
	//			"Content-Type": "application/json",
	//		},
	//		FormParams:    map[string]interface{}{},
	//		SetResCharset: "utf-8",
	//		Timeout:       10,
	//	})
	//fmt.Println(resp, err)
	//response.Success(c, "ok", resp)

	//if resp, err := cli.Get("https://testnets-api.opensea.io/api/v1/asset/0x6b5e78f92c4894f833a6150388869d6fa2f925d4/10"); err == nil {
	//	content, err := resp.GetContents()
	//	if err != nil {
	//		fmt.Println("error - cli", err)
	//	}
	//	response.Success(c, "ok", content)
	//}

	//response.Success(c, "ok", model.CreateUsersFactory().List(1, 10))
	//id := c.Query("id")
	//page := c.DefaultQuery("page", "0")
	//name := c.PostForm("name")
	//message := c.PostForm("message")
	//code := c.PostForm("code")
	//fmt.Printf("id: %s; page: %s; name: %s; message: %s code:%s", id, page, name, message, code)

	//c := cacher.New()
	//m := map[string]interface{}{
	//	"key":  "home",
	//	"age":  18,
	//	"name": "kyle",
	//}
	//_ = redis_factory.GetOneRedisClient().Set("test", m, 6*60*60)
	//var m2 = map[string]interface{}{}
	//err, data := c.GetMap("test2")
	//err := redis_factory.GetOneRedisClient().Exists("test")
	//
	//if err != nil {
	//	fmt.Println("error - redis_factory", err)
	//
	//}
	//send, err := sms.New("SMS_219510061", "18627111095").Check("232425")
	//if err != nil {
	//	response.Fail(context, consts.CurdCreatFailCode, err.Error(), "")
	//} else {
	//response.Success(c, "ok", yml_config.CreateYamlFactory().GetString("Redis.Host")+":"+yml_config.CreateYamlFactory().GetString("Redis.Port"))
	//}

	//  由于本项目骨架已经将表单验证器的字段(成员)绑定在上下文，因此可以按照 GetString()、GetInt64()、GetFloat64（）等快捷获取需要的数据类型
	// 当然也可以通过gin框架的上下文原原始方法获取，例如： context.PostForm("name") 获取，这样获取的数据格式为文本，需要自己继续转换
	//newsType := context.GetString(consts.ValidatorPrefix + "newsType")
	//page := context.GetFloat64(consts.ValidatorPrefix + "page")
	//limit := context.GetFloat64(consts.ValidatorPrefix + "limit")
	//userIp := context.ClientIP()
	//isLogin,user :=	GetUserInfo(context)
	//if isLogin {
	//	fmt.Println(user)
	//}

	//newsType := c.GetString(consts.ValidatorPrefix + "newsType")
	//page := c.GetFloat64(consts.ValidatorPrefix + "page")
	//limit := c.GetFloat64(consts.ValidatorPrefix + "limit")
	//userIp := c.ClientIP()
	//
	////// 这里随便模拟一条数据返回
	//response.Success(c, "ok", gin.H{
	//	"newsType": newsType,
	//	"page":     page,
	//	"limit":    limit,
	//	"userIp":   userIp,
	//	"title":    "门户首页公司新闻标题001",
	//	"content":  "门户新闻内容001",
	//})
}
