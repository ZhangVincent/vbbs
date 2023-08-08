// @description minio文件管理系统
// @author zkp15
// @date 2023/7/28 20:43
// @version 1.0

package uploader

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"vbbs/pkg/config"
)

//var minioClient *minio.Client = getMinioUploader()

func GetMinioUploader() *minio.Client {
	m := config.Instance.Uploader.Minio
	var (
		endpoint        = m.Endpoint
		accessKeyID     = m.AccessKeyID
		secretAccessKey = m.SecretAccessKey
		useSSL          = m.UseSSL
	)

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil
	}

	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, "mediafiles")
	fmt.Println(exists, err)

	return minioClient
}
