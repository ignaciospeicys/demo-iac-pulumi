package main

import (
	"demo-pulumi-aws/infrastructure"
	"demo-pulumi-aws/infrastructure/adapters/primary"
	"demo-pulumi-aws/infrastructure/adapters/secondary"
	"demo-pulumi-aws/infrastructure/setup"
)

func main() {
	db := setup.CreateDBConnection()
	gormDB := setup.InitGormDB(db)
	resourceRepository := secondary.NewResourceRepository(gormDB)
	resourceConfigurationRepository := secondary.NewConfigurationRepository(gormDB)

	pulumiStackService := secondary.NewPulumiStackService()
	objectStorageService := secondary.NewPulumiObjectStorageService()
	objectStorageHandler := primary.NewObjectStorageHandler(objectStorageService, pulumiStackService)
	pulumiHandler := primary.NewPulumiHandler(pulumiStackService)
	httpRouter := infrastructure.NewHttpRouter(objectStorageHandler, pulumiHandler)
	pulumiSetup := setup.NewPulumiSetup()

	r := httpRouter.SetupRoutes()

	pulumiSetup.CreateWorkspace()

	_ = r.Run("127.0.0.1:8083")
}
