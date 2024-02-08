package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"log"
	"net/http"
	"os"
)

type BucketCreateRequest struct {
	Name       string `json:"name"`
	Versioning bool   `json:"enable_versioning"`
}

var project = "demo-pulumi-aws"

func main() {
	ensurePlugins()
	r := gin.Default()

	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal("could not set trusted proxies: ", err)
		return
	}

	r.POST("/:stack/bucket", func(ginCtx *gin.Context) {
		req := &BucketCreateRequest{}

		if err := json.NewDecoder(ginCtx.Request.Body).Decode(&req); err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "couldn't parse request body"})
			return
		}

		ctx := context.Background()
		stackName := ginCtx.Param("stack")
		program := initPulumiProgram(req)

		// Create or select the stack
		s, err := auto.UpsertStackInlineSource(ctx, stackName, project, program)
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		_ = s.SetConfig(ctx, "aws:region", auto.ConfigValue{Value: "us-west-2"})

		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ginCtx.JSON(http.StatusOK, gin.H{
			"message": "Resource created successfully",
		})
	})

	_ = r.Run("127.0.0.1:8083")
}

func initPulumiProgram(req *BucketCreateRequest) pulumi.RunFunc {
	return func(ctx *pulumi.Context) error {
		_, err := s3.NewBucket(ctx, req.Name, &s3.BucketArgs{
			Versioning: s3.BucketVersioningArgs{
				Enabled: pulumi.Bool(req.Versioning),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("bucketName", pulumi.String(req.Name))
		return nil
	}
}

func ensurePlugins() {
	ctx := context.Background()
	w, err := auto.NewLocalWorkspace(ctx)
	if err != nil {
		fmt.Printf("Failed to setup and run http server: %v\n", err)
		os.Exit(1)
	}
	err = w.InstallPlugin(ctx, "aws", "v6.21.0")
	if err != nil {
		fmt.Printf("Failed to install program plugins: %v\n", err)
		os.Exit(1)
	}
}
