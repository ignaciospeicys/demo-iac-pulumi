package driving

import "github.com/gin-gonic/gin"

type StackManagerHandler interface {
	DeleteStack(ctx *gin.Context)
}
