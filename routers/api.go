package routers

import (
	"ginfast/app/global/consts"
	"ginfast/app/global/variable"
	"ginfast/app/http/controller/api"
	"ginfast/app/http/middleware/authorization"
	"ginfast/app/http/middleware/cors"
	validatorFactory "ginfast/app/http/validator/core/factory"
	"ginfast/app/utils/gin_release"
	_ "ginfast/cmd/api/docs"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
)

// 该路由主要设置门户类网站等前台路由

func InitApiRouter() *gin.Engine {
	var router *gin.Engine
	// 非调试模式（生产模式） 日志写到日志文件
	if variable.ConfigYml.GetBool("AppDebug") == false {
		//1.gin自行记录接口访问日志，不需要nginx，如果开启以下3行，那么请屏蔽第 34 行代码
		//gin.DisableConsoleColor()
		//f, _ := os.Create(variable.BasePath + variable.ConfigYml.GetString("Logs.GinLogName"))
		//gin.DefaultWriter = io.MultiWriter(f)

		//【生产模式】
		// 根据 gin 官方的说明：[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
		// 如果部署到生产环境，请使用以下模式：
		// 1.生产模式(release) 和开发模式的变化主要是禁用 gin 记录接口访问日志，
		// 2.go服务就必须使用nginx作为前置代理服务，这样也方便实现负载均衡
		// 3.如果程序发生 panic 等异常使用自定义的 panic 恢复中间件拦截、记录到日志
		router = gin_release.ReleaseRouter()
	} else {
		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		router = gin.Default()
		pprof.Register(router)
	}
	// 设置可信任的代理服务器列表,gin (2021-11-24发布的v1.7.7版本之后出的新功能)
	if variable.ConfigYml.GetInt("HttpServer.TrustProxies.IsOpen") == 1 {
		if err := router.SetTrustedProxies(variable.ConfigYml.GetStringSlice("HttpServer.TrustProxies.ProxyServerList")); err != nil {
			variable.ZapLog.Error(consts.GinSetTrustProxyError, zap.Error(err))
		}
	} else {
		_ = router.SetTrustedProxies(nil)
	}

	//根据配置进行设置跨域
	if variable.ConfigYml.GetBool("HttpServer.AllowCrossDomain") {
		router.Use(cors.Next())
	}
	//router.Use(gin_err.ErrHandler())
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Api 模块接口 hello word！")
	})

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//处理静态资源（不建议gin框架处理静态资源，参见 Public/readme.md 说明 ）
	router.Static("/public", "./public") //  		定义静态资源路由与实际目录映射关系
	//router.StaticFile("/readme", "./public/readme.md") // 可以根据文件名绑定需要返回的文件名

	router.GET("ws", validatorFactory.Create(consts.ValidatorPrefix+"WebsocketConnect"))
	//  创建一个门户类接口路由组
	vApi := router.Group("/api/")
	{
		//无需验证
		noAuth := vApi.Group("/")
		{

			noAuth.GET("tt", (&api.Home{}).Tt)
			noAuth.POST("tool/send_code", validatorFactory.Create(consts.ValidatorPrefix+"SendCode"))
			noAuth.POST("auth/login", validatorFactory.Create(consts.ValidatorPrefix+"Login"))

			//noAuth.POST("login", (&api.Home{}).Wx)
			//noAuth.POST("login", validatorFactory.Create(consts.ValidatorPrefix+"WxLogin"))
			//noAuth.POST("binding", validatorFactory.Create(consts.ValidatorPrefix+"WxBind"))
			//noAuth.GET("w", (&api.Test{}).WsSendMsg)
			//noAuth.POST("test", validatorFactory.Create(consts.ValidatorPrefix+"TestTs"))
			//noAuth.POST("refreshtoken", validatorFactory.Create(consts.ValidatorPrefix+"MRefreshToken"))

		}

		//vApi.Use(authorization.RefreshTokenConditionCheck()).POST("refreshtoken", validatorFactory.Create(consts.ValidatorPrefix+"MRefreshToken"))

		// 【需要token】中间件验证的路由
		needAuth := vApi.Use(authorization.CheckTokenAuth("u"))
		{
			needAuth.GET("me", (&api.Home{}).Tt)
		}

	}
	return router
}
