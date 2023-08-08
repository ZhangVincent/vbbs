package render

import (
	"vbbs/model"

	"github.com/sirupsen/logrus"
	"vbbs/pkg/simple/common/jsons"
	"vbbs/pkg/simple/common/strs"
)

func buildImageList(imageListStr string) (imageList []model.ImageInfo) {
	if strs.IsNotBlank(imageListStr) {
		var images []model.ImageDTO
		if err := jsons.Parse(imageListStr, &images); err == nil {
			if len(images) > 0 {
				for _, image := range images {
					imageList = append(imageList, model.ImageInfo{
						Url:     HandleOssImageStyleDetail(image.Url),
						Preview: HandleOssImageStylePreview(image.Url),
					})
				}
			}
		} else {
			logrus.Error(err)
		}
	}
	return
}

func buildImage(imageStr string) *model.ImageInfo {
	if strs.IsBlank(imageStr) {
		return nil
	}
	var img *model.ImageDTO
	if err := jsons.Parse(imageStr, &img); err != nil {
		logrus.Error(err)
		return nil
	} else {
		return &model.ImageInfo{
			Url:     HandleOssImageStyleDetail(img.Url),
			Preview: HandleOssImageStylePreview(img.Url),
		}
	}
}
