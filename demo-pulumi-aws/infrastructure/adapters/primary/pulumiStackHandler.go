package primary

import (
	"demo-pulumi-aws/infrastructure/adapters/secondary"
	"demo-pulumi-aws/infrastructure/setup"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PulumiStackHandler struct {
	pulumiStackService *secondary.PulumiStackService
	dbService          *secondary.ResourceDBService
}

func NewPulumiHandler(pulumiStackService *secondary.PulumiStackService, dbService *secondary.ResourceDBService) *PulumiStackHandler {
	return &PulumiStackHandler{
		pulumiStackService: pulumiStackService,
		dbService:          dbService,
	}
}

func (ph PulumiStackHandler) DeleteStack(ctx *gin.Context) {
	stackName := ctx.Param("stack")
	err := ph.pulumiStackService.DeleteStack(ctx, setup.ProjectName, stackName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = ph.dbService.DeleteResourcesByStackName(stackName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"operation_message": "Successfully Deleted Stack",
	})
}
