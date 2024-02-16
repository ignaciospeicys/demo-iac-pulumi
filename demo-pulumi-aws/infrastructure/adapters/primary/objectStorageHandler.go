package primary

import (
	"demo-pulumi-aws/application/ports/driven"
	"demo-pulumi-aws/dto"
	"demo-pulumi-aws/infrastructure/adapters/secondary"
	"demo-pulumi-aws/infrastructure/setup"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ObjectStorageHandler struct {
	pulumiObjectStoragePort driven.PulumiObjectStoragePort
	pulumiStackService      *secondary.PulumiStackService
}

type ObjectStorageCreateResponse struct {
	Stack      string `json:"stack"`
	BucketName string `json:"bucket_name"`
	Domain     string `json:"domain"`
}

func NewObjectStorageHandler(pulumiObjectStoragePort driven.PulumiObjectStoragePort, pulumiStackService *secondary.PulumiStackService) *ObjectStorageHandler {
	return &ObjectStorageHandler{
		pulumiObjectStoragePort: pulumiObjectStoragePort,
		pulumiStackService:      pulumiStackService,
	}
}

func (objHandler *ObjectStorageHandler) CreateObjectStorage(ctx *gin.Context) {
	req := &dto.ObjectStorageCreateRequest{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "couldn't parse request body"})
		return
	}
	// Create the resource
	storageResource := objHandler.pulumiObjectStoragePort.CreateObjectStorageResource(req)

	// Create or select the stack
	stackName := ctx.Param("stack")
	upRes, err := objHandler.pulumiStackService.PrepareAndDeployResource(ctx, stackName, setup.ProjectName, storageResource)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &ObjectStorageCreateResponse{
		Stack:      stackName,
		BucketName: upRes.Outputs["bucketName"].Value.(string),
		Domain:     upRes.Outputs["bucketDomain"].Value.(string),
	})
}
