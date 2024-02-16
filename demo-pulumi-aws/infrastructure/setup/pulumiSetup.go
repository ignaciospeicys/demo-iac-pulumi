package setup

import (
	"context"
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
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

	// File backend
	localBackendURL := "file:///Users/ignaciospeicys/pulumi-backend/"
	envVars := auto.EnvVars(map[string]string{
		"PULUMI_BACKEND_URL": localBackendURL,
	})

	w, err := auto.NewLocalWorkspace(ctx, envVars)
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
