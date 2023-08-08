package api

import (
	"io/ioutil"
	"strconv"
	"vbbs/model/constants"
	"vbbs/pkg/uploader"

	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"vbbs/pkg/simple/web"

	"vbbs/services"
)

type UploadController struct {
	Ctx iris.Context
}

func (c *UploadController) Post() *web.JsonResult {
	user := services.UserTokenService.GetCurrent(c.Ctx)
	if err := services.UserService.CheckPostStatus(user); err != nil {
		return web.JsonError(err)
	}

	file, header, err := c.Ctx.FormFile("image")
	if err != nil {
		return web.JsonError(err)
	}
	defer file.Close()

	if header.Size > constants.UploadMaxBytes {
		return web.JsonErrorMsg("图片不能超过" + strconv.Itoa(constants.UploadMaxM) + "M")
	}

	contentType := header.Header.Get("Content-Type")
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return web.JsonError(err)
	}

	logrus.Info("上传文件：", header.Filename, " size:", header.Size)

	url, err := uploader.PutImage(fileBytes, contentType, header.Size)
	if err != nil {
		return web.JsonError(err)
	}
	return web.NewEmptyRspBuilder().Put("url", url).JsonResult()
}
