package infrastructure

import (
	"demo-pulumi-aws/application/ports/driving"
	"github.com/gin-gonic/gin"
	"log"
)

type HttpRouter struct {
	objectStorageHandler driving.ObjectStorageHandler
	stackManagerHandler  driving.StackManagerHandler
}

func NewHttpRouter(objectStorageHandler driving.ObjectStorageHandler, stackManagerHandler driving.StackManagerHandler) *HttpRouter {
	return &HttpRouter{
		objectStorageHandler: objectStorageHandler,
		stackManagerHandler:  stackManagerHandler,
	}
}

func (hr HttpRouter) SetupRoutes() (r *gin.Engine) {
	r = gin.Default()

	//Stack operations
	r.DELETE("/:stack", hr.stackManagerHandler.DeleteStack)

	//Storage Objects
	r.POST("/:stack/bucket", hr.objectStorageHandler.CreateObjectStorage)

	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal("could not set trusted proxies: ", err)
	}
	return
}
