package infrastructure

import (
	"demo-pulumi-aws/domain"
	"github.com/gin-gonic/gin"
	"log"
)

type HttpRouter struct {
	objectStorageHandler *domain.ObjectStorageHandler
	pulumiHandler        *domain.PulumiHandler
}

func NewHttpRouter(objectStorageHandler *domain.ObjectStorageHandler, pulumiHandler *domain.PulumiHandler) *HttpRouter {
	return &HttpRouter{
		objectStorageHandler: objectStorageHandler,
		pulumiHandler:        pulumiHandler,
	}
}

func (hr HttpRouter) SetupHttpServer() (r *gin.Engine) {
	r = gin.Default()

	//Stack operations
	r.DELETE("/:stack", hr.pulumiHandler.DeleteStack)

	//Storage Objects
	r.POST("/:stack/bucket", hr.objectStorageHandler.CreateBucket)

	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal("could not set trusted proxies: ", err)
	}
	return
}
