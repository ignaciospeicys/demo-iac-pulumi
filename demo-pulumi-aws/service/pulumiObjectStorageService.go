package service

import (
	"context"
	"demo-pulumi-aws/dto"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"os"
)

type PulumiObjectStorageService struct {
}

func NewPulumiObjectStorageService() *PulumiObjectStorageService {
	return &PulumiObjectStorageService{}
}

func (service *PulumiObjectStorageService) PrepareAndDeployResource(ctx context.Context, stackName, project string, programRun pulumi.RunFunc) (*auto.UpResult, error) {
	// Create or select the stack
	s, err := auto.UpsertStackInlineSource(ctx, stackName, project, programRun)
	if err != nil {
		return nil, err
	}
	_ = s.SetConfig(ctx, "aws:region", auto.ConfigValue{Value: "us-west-2"})
	upRes, err := s.Up(ctx, optup.ProgressStreams(os.Stdout))
	if err != nil {
		return nil, err
	}
	return &upRes, nil
}

func (service *PulumiObjectStorageService) CreateObjectStorageResource(req *dto.ObjectStorageCreateRequest) pulumi.RunFunc {
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
