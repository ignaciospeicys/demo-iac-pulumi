package main

import (
	"demo-pulumi-aws/infrastructure"
	"demo-pulumi-aws/infrastructure/adapters/primary"
	"demo-pulumi-aws/infrastructure/adapters/secondary"
	"demo-pulumi-aws/infrastructure/setup"
)

func main() {
	gormDB := setup.CreateGormDBConnection()

	resourceRepository := secondary.NewResourceRepository(gormDB)
	resourceConfigurationRepository := secondary.NewConfigurationRepository(gormDB)

	dbService := secondary.NewResourceDBService(resourceRepository, resourceConfigurationRepository)
	pulumiStackService := secondary.NewPulumiStackService()
	objectStorageService := secondary.NewPulumiObjectStorageService()

	objectStorageHandler := primary.NewObjectStorageHandler(objectStorageService, pulumiStackService, dbService)
	pulumiHandler := primary.NewPulumiHandler(pulumiStackService, dbService)

	httpRouter := infrastructure.NewHttpRouter(objectStorageHandler, pulumiHandler)

	pulumiSetup := setup.NewPulumiSetup()

	r := httpRouter.SetupRoutes()

	pulumiSetup.CreateWorkspace()

	_ = r.Run("127.0.0.1:8083")
}
