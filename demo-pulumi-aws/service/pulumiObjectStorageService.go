package service

import (
	"demo-pulumi-aws/dto"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PulumiObjectStorageService struct {
	pulumiDeployService *PulumiDeployService
}

func NewPulumiObjectStorageService(pulumiDeployService *PulumiDeployService) *PulumiObjectStorageService {
	return &PulumiObjectStorageService{pulumiDeployService: pulumiDeployService}
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
