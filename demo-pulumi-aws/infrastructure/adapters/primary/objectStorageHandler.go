package primary

import (
	"demo-pulumi-aws/application/ports/driven"
	"demo-pulumi-aws/dto"
	"demo-pulumi-aws/infrastructure/adapters/secondary"
	"demo-pulumi-aws/infrastructure/setup"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ObjectStorageHandler struct {
	pulumiObjectStoragePort driven.PulumiObjectStoragePort
	pulumiStackService      *secondary.PulumiStackService
	dbService               *secondary.ResourceDBService
}

type ObjectStorageCreateResponse struct {
	Stack         string `json:"stack"`
	BucketName    string `json:"bucket_name"`
	QualifiedName string `json:"qualified_name"`
	Domain        string `json:"domain"`
}

func NewObjectStorageHandler(pulumiObjectStoragePort driven.PulumiObjectStoragePort, pulumiStackService *secondary.PulumiStackService, dbService *secondary.ResourceDBService) *ObjectStorageHandler {
	return &ObjectStorageHandler{
		pulumiObjectStoragePort: pulumiObjectStoragePort,
		pulumiStackService:      pulumiStackService,
		dbService:               dbService,
	}
}

func (objHandler *ObjectStorageHandler) CreateObjectStorage(ctx *gin.Context) {
	req := &dto.ObjectStorageCreateRequest{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "couldn't parse request body"})
		return
	}

	stackName := ctx.Param("stack")
	resources, err := objHandler.dbService.FetchAllResources(stackName)

	storageResource := objHandler.pulumiObjectStoragePort.CreateObjectStorageResource(req, resources)

	// Create or select the stack
	upRes, err := objHandler.pulumiStackService.PrepareAndDeployResource(ctx, stackName, setup.ProjectName, storageResource)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	qualifiedName := upRes.Outputs["bucketName"].Value.(string)

	err = objHandler.dbService.SaveResource(dto.ResourceDTO{
		ResourceName:   req.Name,
		QualifiedName:  qualifiedName,
		ResourceType:   "object-storage",
		StackName:      stackName,
		Status:         "created",
		Configurations: []dto.ConfigurationDTO{{ConfigKey: "Versioning", ConfigValue: strconv.FormatBool(req.Versioning)}},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &ObjectStorageCreateResponse{
		Stack:         stackName,
		BucketName:    req.Name,
		QualifiedName: qualifiedName,
		Domain:        upRes.Outputs["bucketDomain"].Value.(string),
	})
}

func (objHandler *ObjectStorageHandler) RefreshObjectStorage(ctx *gin.Context) {
	stackName := ctx.Param("stack")
	resources, err := objHandler.dbService.FetchAllResources(stackName)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	storageResource := objHandler.pulumiObjectStoragePort.RefreshObjectStorageResource(resources)

	// Create or select the stack
	_, err = objHandler.pulumiStackService.PrepareAndDeployResource(ctx, stackName, setup.ProjectName, storageResource)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
