package api

import (
	"github.com/kataras/iris/v12"
	"vbbs/pkg/simple/web"

	"vbbs/services"
)

type ConfigController struct {
	Ctx iris.Context
}

func (c *ConfigController) GetConfigs() *web.JsonResult {
	config := services.SysConfigService.GetConfig()
	return web.JsonData(config)
}
