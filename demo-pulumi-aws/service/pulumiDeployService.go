package service

import (
	"context"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"os"
)

type PulumiDeployService struct {
}

func NewPulumiDeployService() *PulumiDeployService {
	return &PulumiDeployService{}
}

func (service *PulumiDeployService) PrepareAndDeployResource(ctx context.Context, stackName, project string, programRun pulumi.RunFunc) (*auto.UpResult, error) {
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
