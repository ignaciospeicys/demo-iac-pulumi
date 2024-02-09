package main

import (
	"demo-pulumi-aws/domain"
	"demo-pulumi-aws/infrastructure"
	"demo-pulumi-aws/service"
)

type BucketCreateRequest struct {
	Name       string `json:"name"`
	Versioning bool   `json:"enable_versioning"`
}

type BucketCreateResponse struct {
	Stack string `json:"stack"`
	URL   string `json:"url"`
}

var project = "demo-pulumi-aws"

func main() {
	pulumiDeployService := service.NewPulumiDeployService()
	objectStorageService := service.NewPulumiObjectStorageService(pulumiDeployService)
	objectStorageHandler := domain.NewObjectStorageHandler(project, objectStorageService, pulumiDeployService)
	pulumiHandler := domain.NewPulumiHandler()
	httpRouter := infrastructure.NewHttpRouter(objectStorageHandler, pulumiHandler)
	pulumiSetup := infrastructure.NewPulumiSetup()

	r := httpRouter.SetupHttpServer()

	pulumiSetup.EnsurePlugins()

	_ = r.Run("127.0.0.1:8083")
}
