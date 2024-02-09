package infrastructure

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

func (ps *PulumiSetup) EnsurePlugins() {
	ctx := context.Background()
	w, err := auto.NewLocalWorkspace(ctx)
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
