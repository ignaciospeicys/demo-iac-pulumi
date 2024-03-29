package setup

import (
	"context"
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"
	"os"
)

// verify latest version here: https://github.com/pulumi/pulumi-aws
var awsPluginVersion = "v6.21.0"

type PulumiSetup struct {
}

func NewPulumiSetup() *PulumiSetup {
	return &PulumiSetup{}
}

const ProjectName = "demo-pulumi-aws"

func (ps *PulumiSetup) CreateWorkspace() {
	ctx := context.Background()

	proj := auto.Project(workspace.Project{
		Name:    ProjectName,
		Runtime: workspace.NewProjectRuntimeInfo("go", nil),
		Backend: &workspace.ProjectBackend{
			URL: "file:///Users/ignaciospeicys/pulumi-file-backend/",
		},
	})

	w, err := auto.NewLocalWorkspace(ctx, proj)
	if err != nil {
		fmt.Printf("Failed to initialize local workspace: %v\n", err)
		os.Exit(1)
	}

	err = w.InstallPlugin(ctx, "aws", awsPluginVersion)
	if err != nil {
		fmt.Printf("Failed to install program plugins: %v\n", err)
		os.Exit(1)
	}
}
