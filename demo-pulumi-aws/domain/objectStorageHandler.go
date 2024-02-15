package domain

import (
	"demo-pulumi-aws/dto"
	"demo-pulumi-aws/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ObjectStorageHandler struct {
	project                    string
	pulumiObjectStorageService *service.PulumiObjectStorageService
	pulumiDeployService        *service.PulumiDeployService
}

type ObjectStorageCreateResponse struct {
	Stack      string `json:"stack"`
	BucketName string `json:"bucket_name"`
	Domain     string `json:"domain"`
}

func NewObjectStorageHandler(project string, pulumiObjectStorageService *service.PulumiObjectStorageService, pulumiDeployService *service.PulumiDeployService) *ObjectStorageHandler {
	return &ObjectStorageHandler{
		project:                    project,
		pulumiObjectStorageService: pulumiObjectStorageService,
		pulumiDeployService:        pulumiDeployService,
	}
}

func (objHandler *ObjectStorageHandler) CreateBucket(ctx *gin.Context) {
	req := &dto.ObjectStorageCreateRequest{}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "couldn't parse request body"})
		return
	}

	// Create the resource
	storageResource := objHandler.pulumiObjectStorageService.CreateObjectStorageResource(req)

	// Create or select the stack
	stackName := ctx.Param("stack")
	upRes, err := objHandler.pulumiDeployService.PrepareAndDeployResource(ctx, stackName, objHandler.project, storageResource)
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
