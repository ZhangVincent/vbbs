package api

import (
	"vbbs/model"
	"vbbs/services"

	"github.com/kataras/iris/v12"
	"vbbs/pkg/simple/common/dates"
	"vbbs/pkg/simple/web"
	"vbbs/pkg/simple/web/params"
)

type UserReportController struct {
	Ctx iris.Context
}

func (c *UserReportController) PostSubmit() *web.JsonResult {
	var (
		dataId, _ = params.FormValueInt64(c.Ctx, "dataId")
		dataType  = params.FormValue(c.Ctx, "dataId")
		reason    = params.FormValue(c.Ctx, "reason")
	)
	report := &model.UserReport{
		DataId:     dataId,
		DataType:   dataType,
		Reason:     reason,
		CreateTime: dates.NowTimestamp(),
	}

	if user := services.UserTokenService.GetCurrent(c.Ctx); user != nil {
		report.UserId = user.Id
	}
	services.UserReportService.Create(report)
	return web.JsonSuccess()
}
