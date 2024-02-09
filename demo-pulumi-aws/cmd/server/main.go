package main

import (
	"demo-pulumi-aws/adapters"
	"demo-pulumi-aws/domain"
	"demo-pulumi-aws/infrastructure"
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
	objectStorageService := adapters.NewPulumiObjectStorageService()
	objectStorageHandler := domain.NewObjectStorageHandler(project, objectStorageService)
	httpRouter := infrastructure.NewHttpRouter(objectStorageHandler)
	pulumiSetup := infrastructure.NewPulumiSetup()

	r := httpRouter.SetupHttpServer()

	pulumiSetup.EnsurePlugins()

	_ = r.Run("127.0.0.1:8083")
}
