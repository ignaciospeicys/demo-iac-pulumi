package driving

import "github.com/gin-gonic/gin"

type ObjectStorageHandler interface {
	CreateObjectStorage(ctx *gin.Context)
	RefreshObjectStorage(ctx *gin.Context)
}
