package ossUtils

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"os"
)

func OssUtils(file *multipart.FileHeader, username string) string {
	client, err := oss.New("cn-beijing.oss.aliyuncs.com", "key", "key secret")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	fmt.Println("user:", username)
	// yourBucketName填写存储空间名称。
	bucketName := "bucketName-c"
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

	return "https://bucketName-c.oss-cn-beijing.aliyuncs.com/" + objectName
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
