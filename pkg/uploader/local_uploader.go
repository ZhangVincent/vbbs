package uploader

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"log"
	"strings"
	"vbbs/pkg/bbsurls"
	"vbbs/pkg/config"
)

// 本地文件系统
type localUploader struct{}

func (local *localUploader) PutImage(data []byte, contentType string, filesize int64) (string, error) {
	key := generateImageKey(data, contentType)
	return local.PutObject(key, data, contentType, filesize)
}

func (local *localUploader) PutObject(key string, data []byte, contentType string, filesize int64) (string, error) {
	c := config.Instance.Uploader.Local
	ctx := context.Background()
	bucketName := strings.Trim(c.Path, "/")
	minioClient := GetMinioUploader()

	exists, err := minioClient.BucketExists(ctx, bucketName)
	if !exists {
		return "", err
	}

	objectName := key
	filereader := bytes.NewReader(data)
	info, err := minioClient.PutObject(ctx, bucketName, objectName, filereader, filesize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	return bbsurls.UrlJoin(c.Host, c.Path, key), nil
}

func (local *localUploader) CopyImage(originUrl string) (string, error) {
	data, contentType, err := download(originUrl)
	if err != nil {
		return "", err
	}
	return local.PutImage(data, contentType, 0)
}
