package logging

import (
	"fmt"
	"io"
	"os"
	"time"
)

// MultiWriter encapsulates writing logs to both a file and stdout, tailored for Pulumi deployment logging.
// @see also: [io.Writer]
type MultiWriter struct {
	File   *os.File
	Stdout io.Writer
}

// NewMultiWriter initializes a MultiWriter for a given project and stack, creating a log file in ./logs.
func NewMultiWriter(project, stackName string) (*MultiWriter, error) {
	logsDir := "./logs"
	if err := os.MkdirAll(logsDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %v", err)
	}

	logFileName := fmt.Sprintf("%s/pulumi-deployments-%s-%s-%d.log", logsDir, project, stackName, time.Now().Unix())
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	return &MultiWriter{
		File:   logFile,
		Stdout: os.Stdout,
	}, nil
}

func (mw *MultiWriter) Write(p []byte) (n int, err error) {
	if _, err := mw.File.Write(p); err != nil {
		return 0, err
	}
	return mw.Stdout.Write(p)
}

func (mw *MultiWriter) Close() error {
	return mw.File.Close()
}
