package service

import (
	"context"
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"io"
	"os"
	"time"
)

type PulumiDeployService struct {
}

// MultiWriter writes to both stdout and a log file TODO Move class
type MultiWriter struct {
	File   *os.File
	Stdout io.Writer
}

func (mw *MultiWriter) Write(p []byte) (n int, err error) {
	if _, err := mw.File.Write(p); err != nil {
		return 0, err
	}
	return mw.Stdout.Write(p)
}

func NewPulumiDeployService() *PulumiDeployService {
	return &PulumiDeployService{}
}

//end TODO Move class

func (service *PulumiDeployService) PrepareAndDeployResource(ctx context.Context, stackName, project string, programRun pulumi.RunFunc) (*auto.UpResult, error) {

	// Create the logs directory if it doesn't already exist
	logsDir := "./logs"
	if err := os.MkdirAll(logsDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %v", err)
	}

	// Create a dynamic log file name with the current date
	logFileName := fmt.Sprintf("%s/pulumi-deployments-%s-%s--%d.log", logsDir, project, stackName, time.Now().Unix())

	// Create a custom writer to log to a file
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}
	defer logFile.Close()

	mw := &MultiWriter{
		File:   logFile,
		Stdout: os.Stdout,
	}

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
