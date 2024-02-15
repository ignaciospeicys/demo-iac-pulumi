package driven

import (
	"context"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PulumiStackManagerPort interface {
	PrepareAndDeployResource(ctx context.Context, stackName, project string, programRun pulumi.RunFunc) (*auto.UpResult, error)

	DeleteStack(ctx context.Context, project, stackName string) error
}
