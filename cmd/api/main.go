package main

import (
	"ginfast/app/global/variable"
	_ "ginfast/bootstrap"
	"ginfast/routers"
)

// @title Kyle API
// @version 0.0.1
// @description kyle api 文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	router := routers.InitApiRouter()
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Api.Port"))
}
