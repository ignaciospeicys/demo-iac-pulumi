package domain

import (
	"context"
	"demo-pulumi-aws/adapters"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ObjectStorageHandler struct {
	project                    string
	pulumiObjectStorageService *adapters.PulumiObjectStorageService
}

type ObjectStorageCreateRequest struct {
	Name       string `json:"name"`
	Versioning bool   `json:"enable_versioning"`
}

type ObjectStorageCreateResponse struct {
	Stack string `json:"stack"`
	URL   string `json:"url"`
}

func NewObjectStorageHandler(project string, pulumiObjectStorageService *adapters.PulumiObjectStorageService) *ObjectStorageHandler {
	return &ObjectStorageHandler{
		project:                    project,
		pulumiObjectStorageService: pulumiObjectStorageService,
	}
}

func (objHandler *ObjectStorageHandler) CreateBucket(ginCtx *gin.Context) {
	req := &ObjectStorageCreateRequest{}

	if err := json.NewDecoder(ginCtx.Request.Body).Decode(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "couldn't parse request body"})
		return
	}

	ctx := context.Background()

	// Create the resource
	storageResource := objHandler.pulumiObjectStorageService.CreateObjectStorageResource(req)

	// Create or select the stack
	stackName := ginCtx.Param("stack")
	upRes, err := objHandler.pulumiObjectStorageService.PrepareAndDeployResource(ctx, stackName, objHandler.project, storageResource)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, &ObjectStorageCreateResponse{
		Stack: stackName,
		URL:   upRes.Outputs["bucketName"].Value.(string),
	})
}
