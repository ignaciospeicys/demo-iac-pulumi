package secondary

import (
	"demo-pulumi-aws/dto"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"log"
	"strconv"
)

type PulumiObjectStorageService struct {
}

func NewPulumiObjectStorageService() *PulumiObjectStorageService {
	return &PulumiObjectStorageService{}
}

func (service *PulumiObjectStorageService) CreateObjectStorageResource(req *dto.ObjectStorageCreateRequest, resources []dto.ResourceDTO) pulumi.RunFunc {
	return func(ctx *pulumi.Context) error {
		//re-declares existing resources so pulumi has them in the context
		for _, r := range resources {
			_, err := s3.NewBucket(ctx, r.ResourceName, &s3.BucketArgs{
				Versioning: s3.BucketVersioningArgs{
					Enabled: pulumi.Bool(findVersioningConfig(r.Configurations)),
				},
			})
			if err != nil {
				return err
			}
		}
		//this is the bucket we actually want to add onto the state
		bucket, err := s3.NewBucket(ctx, req.Name, &s3.BucketArgs{
			Versioning: s3.BucketVersioningArgs{
				Enabled: pulumi.Bool(req.Versioning),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("bucketName", bucket.Bucket)
		ctx.Export("bucketDomain", bucket.BucketDomainName)
		return nil
	}
}

func findVersioningConfig(configs []dto.ConfigurationDTO) bool {
	for _, config := range configs {
		if config.ConfigKey == "Versioning" {
			versioningEnabled, err := strconv.ParseBool(config.ConfigValue)
			if err != nil {
				log.Printf("Error parsing versioning config value: %v", err)
				return false
			}
			return versioningEnabled
		}
	}
	return false
}

func (service *PulumiObjectStorageService) RefreshObjectStorageResource(resources []dto.ResourceDTO) pulumi.RunFunc {
	return func(ctx *pulumi.Context) error {
		//re-declares existing resources so pulumi has them in the context
		for _, r := range resources {
			_, err := s3.NewBucket(ctx, r.ResourceName, &s3.BucketArgs{
				Versioning: s3.BucketVersioningArgs{
					Enabled: pulumi.Bool(findVersioningConfig(r.Configurations)),
				},
			})
			if err != nil {
				return err
			}
		}
		return nil
	}
}
