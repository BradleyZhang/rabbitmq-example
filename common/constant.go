package common

var (
	Username   = "admin"
	Password   = "password"
	IP         = "127.0.0.1"
	Port       = "5672"
	SecurePort = "5671"
	Url        = "amqp://" + Username + ":" + Password + "@" + IP + ":" + Port
	SecureUrl  = "amqps://" + Username + ":" + Password + "@" + IP + ":" + SecurePort
)

var (
	CertFile = "./ca/client_certificate.pem"
	KeyFile  = "./ca/client_key.pem"
	CaCert   = "./ca/ca_certificate.pem"
)
