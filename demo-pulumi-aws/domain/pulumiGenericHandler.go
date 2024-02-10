package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optdestroy"
	"net/http"
	"os"
)

type PulumiHandler struct {
}

func NewPulumiHandler() *PulumiHandler {
	return &PulumiHandler{}
}

var project = "demo-pulumi-aws"

func (ph PulumiHandler) DeleteStack(ctx *gin.Context) {
	stackName := ctx.Param("stack")
	s, err := auto.SelectStackInlineSource(ctx, stackName, project, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_ = s.SetConfig(ctx, "aws:region", auto.ConfigValue{Value: "us-west-2"})

	_, err = s.Destroy(ctx, optdestroy.ProgressStreams(os.Stdout))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// delete the stack and all associated history and config
	err = s.Workspace().RemoveStack(ctx, stackName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"operation_message": "Successfully Deleted Stack",
	})
}
