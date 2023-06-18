package ossUtils

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"mime/multipart"
	"os"
)

func OssUtils(file *multipart.FileHeader, username string) string {
	config := viper.New()
	config.SetConfigName("application")
	config.AddConfigPath("./config")
	config.SetConfigType("yaml")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	client, err := oss.New(config.GetString("oss.endpoint"), config.GetString("oss.accessKeyID"), config.GetString("oss.accessKeySecret"))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// yourBucketName填写存储空间名称。
	bucketName := config.GetString("oss.bucketName")
	// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。
	objectName := username + "avatar.jpg"
	// yourLocalFileName填写本地文件的完整路径。
	tempFile, _ := file.Open()
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("err")
		handleError(err)
	}

	// 上传文件
	err = bucket.PutObject(objectName, tempFile)
	if err != nil {
		fmt.Println("err2")
		handleError(err)
	}

	return "https://recommendation-c.oss-cn-beijing.aliyuncs.com/" + objectName
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
