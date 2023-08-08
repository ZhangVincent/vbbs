package uploader

import (
	"strings"
	"sync"

	"vbbs/pkg/simple/common/strs"
	"vbbs/pkg/simple/common/urls"

	"vbbs/pkg/config"
)

// 这里的filesize是为了适配minio
type uploader interface {
	PutImage(data []byte, contentType string, filesize int64) (string, error)
	PutObject(key string, data []byte, contentType string, filesize int64) (string, error)
	CopyImage(originUrl string) (string, error)
}

var (
	aliyun = &aliyunOssUploader{
		once:   sync.Once{},
		bucket: nil,
	}
	local = &localUploader{}
)

func PutImage(data []byte, contentType string, filesize int64) (string, error) {
	return getUploader().PutImage(data, contentType, filesize)
}

func PutObject(key string, data []byte, contentType string, filesize int64) (string, error) {
	return getUploader().PutObject(key, data, contentType, filesize)
}

func CopyImage(url string) (string, error) {
	u1 := urls.ParseUrl(url).GetURL()
	u2 := urls.ParseUrl(config.Instance.BaseUrl).GetURL()
	// 本站host，不下载
	if u1.Host == u2.Host {
		return url, nil
	}
	return getUploader().CopyImage(url)
}

func getUploader() uploader {
	if IsEnabledOss() {
		return aliyun
	} else {
		return local
	}
}

// IsEnabledOss 是否启用阿里云oss
func IsEnabledOss() bool {
	enable := config.Instance.Uploader.Enable
	return strs.EqualsIgnoreCase(enable, "aliyun") || strs.EqualsIgnoreCase(enable, "oss") ||
		strs.EqualsIgnoreCase(enable, "aliyunOss")
}

// IsOssImageUrl 是否是存放在阿里云oss中的图片
func IsOssImageUrl(url string) bool {
	host := urls.ParseUrl(config.Instance.Uploader.AliyunOss.Host).GetURL().Host
	return strings.Contains(url, host)
}
