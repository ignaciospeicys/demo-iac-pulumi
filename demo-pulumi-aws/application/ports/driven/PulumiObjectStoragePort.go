package driven

import (
	"demo-pulumi-aws/dto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PulumiObjectStoragePort interface {
	CreateObjectStorageResource(req *dto.ObjectStorageCreateRequest, resources []dto.ResourceDTO) pulumi.RunFunc
}
