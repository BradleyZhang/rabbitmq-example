package common

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitMQ() *amqp.Channel {
	// 连接 RabbitMQ
	conn, err := amqp.Dial(Url)
	FailOnError(err, "无法连接到 RabbitMQ")

	// 创建通道
	ch, err := conn.Channel()
	FailOnError(err, "无法打开通道")
	return ch
}
func InitMQS() *amqp.Channel {
	// 1. 加载客户端证书和密钥（双向认证时需要）
	cert, err := tls.LoadX509KeyPair(CertFile, KeyFile)
	if err != nil {
		log.Fatalf("加载客户端证书失败: %v", err)
	}

	// 2. 加载CA证书（验证服务器证书）
	caCert, err := os.ReadFile(CaCert)
	if err != nil {
		log.Fatalf("读取CA证书失败: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// 3. 配置TLS
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert}, // 客户端证书（双向认证时需要）
		RootCAs:            caCertPool,              // 信任的CA
		InsecureSkipVerify: false,                   // 必须验证服务器证书
	}

	// 连接 RabbitMQ
	conn, err := amqp.DialTLS(SecureUrl, tlsConfig)
	if err != nil {
		log.Fatalf("无法连接到 RabbitMQ: %v", err)
	}

	// 创建通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("无法打开通道: %v", err)
	}
	return ch
}
