package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"log"
	"net/http"
	"os"
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

var VarBucketExport = "bucketName"

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
		programRun := runPulumi(req)

		// Create or select the stack
		s, err := auto.UpsertStackInlineSource(ctx, stackName, project, programRun)
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

		upRes, err := s.Up(ctx, optup.ProgressStreams(os.Stdout))
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response := &BucketCreateResponse{
			Stack: stackName,
			URL:   upRes.Outputs[VarBucketExport].Value.(string),
		}
		ginCtx.JSON(http.StatusOK, response)
	})

	_ = r.Run("127.0.0.1:8083")
}

func runPulumi(req *BucketCreateRequest) pulumi.RunFunc {
	return func(ctx *pulumi.Context) error {
		_, err := s3.NewBucket(ctx, req.Name, &s3.BucketArgs{
			Versioning: s3.BucketVersioningArgs{
				Enabled: pulumi.Bool(req.Versioning),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export(VarBucketExport, pulumi.String(req.Name))
		return nil
	}
}

func ensurePlugins() {
	ctx := context.Background()
	w, err := auto.NewLocalWorkspace(ctx)
	if err != nil {
		fmt.Printf("Failed to initialize local workspace: %v\n", err)
		os.Exit(1)
	}
	err = w.InstallPlugin(ctx, "aws", "v6.21.0") //verify latest version here: https://github.com/pulumi/pulumi-aws
	if err != nil {
		fmt.Printf("Failed to install program plugins: %v\n", err)
		os.Exit(1)
	}
}
