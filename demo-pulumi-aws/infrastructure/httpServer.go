package infrastructure

import (
	"demo-pulumi-aws/domain"
	"github.com/gin-gonic/gin"
	"log"
)

type HttpRouter struct {
	objectStorageHandler *domain.ObjectStorageHandler
}

func NewHttpRouter(objectStorageHandler *domain.ObjectStorageHandler) *HttpRouter {
	return &HttpRouter{
		objectStorageHandler: objectStorageHandler,
	}
}

func (hr HttpRouter) SetupHttpServer() (r *gin.Engine) {
	r = gin.Default()

	r.POST("/:stack/bucket", hr.objectStorageHandler.CreateBucket)

	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal("could not set trusted proxies: ", err)
	}
	return
}
