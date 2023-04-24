package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func OssUtils(filename string, t string) string {
	// 设置连接数为10，每个主机的最大闲置连接数为20，每个主机的最大连接数为20。
	conn := oss.MaxConns(10, 20, 20)
	// 设置HTTP连接超时时间为20秒，HTTP读取或写入超时时间为60秒。
	time := oss.Timeout(20, 60)
	// 设置是否支持将自定义域名作为Endpoint，默认不支持。
	cname := oss.UseCname(true)
	// 设置HTTP的User-Agent头，默认为aliyun-sdk-go。
	userAgent := oss.UserAgent("aliyun-sdk-go")
	// 设置是否开启HTTP重定向，默认开启。
	redirect := oss.RedirectEnabled(true)
	// 设置是否开启SSL证书校验，默认关闭。
	verifySsl := oss.InsecureSkipVerify(false)
	// 设置代理服务器地址和端口。
	//proxy := oss.Proxy("yourProxyHost")
	// 设置代理服务器的主机地址和端口，代理服务器验证的用户名和密码。
	authProxy := oss.AuthProxy("yourProxyHost", "yourProxyUserName", "yourProxyPassword")
	// 开启CRC加密。
	crc := oss.EnableCRC(true)
	// 设置日志模式。
	logLevel := oss.SetLogLevel(oss.LogOff)

	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5tKuBewz8ZWTtamtJQYy", "mkhYdOS2WyFGqvkWvFmdJ2ZWAa1ggd", conn, time, cname, userAgent, authProxy, verifySsl, redirect, crc, logLevel)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	fmt.Printf("%#v\n", client)

	// yourBucketName填写存储空间名称。
	bucketName := "recommendation-c"
	// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。
	objectName := "yourObjectName"
	// yourLocalFileName填写本地文件的完整路径。
	localFileName := "yourLocalFileName"
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 上传文件
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		handleError(err)
	}

	return "https://recommendation-c.oss-cn-beijing.aliyuncs.com/" + t + filename
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
