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

func (ps *PulumiSetup) CreateWorkspace() {
	ctx := context.Background()

	proj := auto.Project(workspace.Project{
		Name:    "demo-aws-pulumi",
		Runtime: workspace.NewProjectRuntimeInfo("go", nil),
		Backend: &workspace.ProjectBackend{
			URL: "file:///Users/ignaciospeicys/pulumi-file-backend/",
		},
	})

	// File backend
	localBackendURL := "file:///Users/ignaciospeicys/pulumi-file-backend/"
	envVars := auto.EnvVars(map[string]string{
		"PULUMI_BACKEND_URL":       localBackendURL,
		"PULUMI_CONFIG_PASSPHRASE": "n4ch0Pu1um1",
	})

	// Creating a new local workspace with these environment variables
	workspaceOptions := []auto.LocalWorkspaceOption{
		proj,
		envVars,
	}

	w, err := auto.NewLocalWorkspace(ctx, workspaceOptions...)
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
