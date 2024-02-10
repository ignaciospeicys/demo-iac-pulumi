package main

import (
	"demo-pulumi-aws/domain"
	"demo-pulumi-aws/infrastructure"
	"demo-pulumi-aws/service"
)

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
