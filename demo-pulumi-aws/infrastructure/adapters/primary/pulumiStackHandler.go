package primary

import (
	"demo-pulumi-aws/infrastructure/adapters/secondary"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PulumiStackHandler struct {
	pulumiStackService *secondary.PulumiStackService
}

func NewPulumiHandler(pulumiStackService *secondary.PulumiStackService) *PulumiStackHandler {
	return &PulumiStackHandler{pulumiStackService: pulumiStackService}
}

var project = "demo-pulumi-aws"

func (ph PulumiStackHandler) DeleteStack(ctx *gin.Context) {
	stackName := ctx.Param("stack")
	err := ph.pulumiStackService.DeleteStack(ctx, stackName, project)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"operation_message": "Successfully Deleted Stack",
	})
}
