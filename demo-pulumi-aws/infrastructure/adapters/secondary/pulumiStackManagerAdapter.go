package secondary

import (
	"context"
	"demo-pulumi-aws/infrastructure/logging"
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optdestroy"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"os"
)

type PulumiStackService struct {
}

func NewPulumiStackService() *PulumiStackService {
	return &PulumiStackService{}
}

func (service *PulumiStackService) PrepareAndDeployResource(ctx context.Context, stackName, project string, programRun pulumi.RunFunc) (*auto.UpResult, error) {
	mw, err := logging.NewMultiWriter(project, stackName)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize multi-writer: %v", err)
	}
	defer mw.Close()

	// Create or select the stack
	s, err := auto.UpsertStackInlineSource(ctx, stackName, project, programRun)
	if err != nil {
		return nil, err
	}
	_ = s.SetConfig(ctx, "aws:region", auto.ConfigValue{Value: "us-west-2"})

	upRes, err := s.Up(ctx, optup.ProgressStreams(mw))
	if err != nil {
		return nil, err
	}

	return &upRes, nil
}

func (service *PulumiStackService) DeleteStack(ctx context.Context, project, stackName string) error {
	s, err := auto.SelectStackInlineSource(ctx, stackName, project, nil)
	if err != nil {
		return err
	}
	_ = s.SetConfig(ctx, "aws:region", auto.ConfigValue{Value: "us-west-2"})

	_, err = s.Destroy(ctx, optdestroy.ProgressStreams(os.Stdout))
	if err != nil {
		return err
	}
	// delete the stack and all associated history and config
	err = s.Workspace().RemoveStack(ctx, stackName)
	if err != nil {
		return err
	}
	return nil
}
