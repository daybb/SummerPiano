module helloworld

go 1.16

replace google.golang.org/protobuf v1.27.1 => google.golang.org/protobuf v1.27.2-0.20211129185528-26e8bcb3c743

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.7
	github.com/go-kratos/gin v0.1.0
	github.com/go-kratos/kratos/v2 v2.1.2
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/google/wire v0.5.0
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.17.0 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/genproto v0.0.0-20211016002631-37fc39342514
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.2.1
	gorm.io/gorm v1.22.4
)
