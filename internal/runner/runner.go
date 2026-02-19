package runner

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"syscall"

	"github.com/runex/runex/internal/detector"
)

type Runner struct {
	detector  *detector.Detector
	output    io.Writer
	errOutput io.Writer
}

func New(d *detector.Detector, out io.Writer, errOut io.Writer) *Runner {
	return &Runner{
		detector:  d,
		output:    out,
		errOutput: errOut,
	}
}

type RunResult struct {
	ExitCode int
	HasError bool
	Error    *detector.DetectedError
}

func (r *Runner) Run(ctx context.Context, cmd string, args []string) (*RunResult, error) {
	command := exec.CommandContext(ctx, cmd, args...)
	command.Stdout = r.output
	command.Stderr = r.errOutput
	command.Stdin = os.Stdin

	var wg sync.WaitGroup
	var mu sync.Mutex
	var detectedError *detector.DetectedError

	pr, pw, err := os.Pipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create pipe: %w", err)
	}

	command.Stderr = pw

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer pr.Close()

		buf := make([]byte, 1024)
		for {
			n, readErr := pr.Read(buf)
			if n > 0 {
				if _, err := r.errOutput.Write(buf[:n]); err != nil {
					// ignore error
				}

				mu.Lock()
				if detectedError == nil {
					if err := r.detector.Detect(string(buf[:n])); err != nil {
						detectedError = err
					}
				}
				mu.Unlock()
			}
			if readErr != nil {
				break
			}
		}
	}()

	err = command.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start command: %w", err)
	}

	errChan := make(chan error, 1)
	go func() {
		errChan <- command.Wait()
	}()

	select {
	case err := <-errChan:
		pw.Close()
		wg.Wait()

		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
					return &RunResult{
						ExitCode: status.ExitStatus(),
						HasError: detectedError != nil,
						Error:    detectedError,
					}, nil
				}
			}
			return nil, err
		}

		return &RunResult{
			ExitCode: 0,
			HasError: detectedError != nil,
			Error:    detectedError,
		}, nil

	case <-ctx.Done():
		pw.Close()
		wg.Wait()
		if err := command.Process.Kill(); err != nil {
			// ignore error
		}
		return nil, ctx.Err()
	}
}
