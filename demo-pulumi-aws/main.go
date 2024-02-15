package main

import (
	"demo-pulumi-aws/infrastructure"
	"demo-pulumi-aws/infrastructure/adapters/primary"
	"demo-pulumi-aws/infrastructure/adapters/secondary"
	"demo-pulumi-aws/infrastructure/setup"
)

var project = "demo-pulumi-aws"

func main() {
	pulumiStackService := secondary.NewPulumiStackService()
	objectStorageService := secondary.NewPulumiObjectStorageService()
	objectStorageHandler := primary.NewObjectStorageHandler(project, objectStorageService, pulumiStackService)
	pulumiHandler := primary.NewPulumiHandler(pulumiStackService)
	httpRouter := infrastructure.NewHttpRouter(objectStorageHandler, pulumiHandler)
	pulumiSetup := setup.NewPulumiSetup()

	r := httpRouter.SetupRoutes()

	pulumiSetup.EnsurePlugins()

	_ = r.Run("127.0.0.1:8083")
}
