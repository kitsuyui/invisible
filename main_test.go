package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/subcommands"
)

func TestAddNoiseNegativeFrequency(t *testing.T) {
	cmd := &addNoise{frequency: -0.5, maxSize: 1}
	status := cmd.Execute(context.Background(), flag.NewFlagSet("add-noise", flag.ContinueOnError))
	if status != subcommands.ExitUsageError {
		t.Errorf("Execute() = %v, want ExitUsageError", status)
	}
}

func TestAddNoiseNegativeMaxSize(t *testing.T) {
	cmd := &addNoise{frequency: 0.5, maxSize: -1}
	status := cmd.Execute(context.Background(), flag.NewFlagSet("add-noise", flag.ContinueOnError))
	if status != subcommands.ExitUsageError {
		t.Errorf("Execute() = %v, want ExitUsageError", status)
	}
}

func TestEncodeEmptyMessage(t *testing.T) {
	cmd := &encode{message: ""}
	status := cmd.Execute(context.Background(), flag.NewFlagSet("encode", flag.ContinueOnError))
	if status != subcommands.ExitUsageError {
		t.Errorf("Execute() = %v, want ExitUsageError", status)
	}
}

func TestVersionCmdReturnsSuccess(t *testing.T) {
	cmd := &versionCmd{}
	status := cmd.Execute(context.Background(), flag.NewFlagSet("version", flag.ContinueOnError))
	if status != subcommands.ExitSuccess {
		t.Errorf("Execute() = %v, want ExitSuccess", status)
	}
}

func TestPrintCommandErrorAddsCommandAndStage(t *testing.T) {
	stderr := captureStderr(t, func() {
		printCommandError("decode", "extract hidden message", errors.New("unexpected EOF"))
	})

	want := "error: decode: extract hidden message: unexpected EOF\n"
	if stderr != want {
		t.Errorf("stderr = %q, want %q", stderr, want)
	}
}

func TestPrintCommandErrorHandlesNilError(t *testing.T) {
	stderr := captureStderr(t, func() {
		printCommandError("encode", "embed message", nil)
	})

	if !strings.Contains(stderr, "error: encode: embed message: unknown error") {
		t.Errorf("stderr = %q, want command and stage context", stderr)
	}
}

func captureStderr(t *testing.T, fn func()) string {
	t.Helper()

	original := os.Stderr
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.Stderr = original
		_ = reader.Close()
	}()

	os.Stderr = writer
	fn()
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, reader); err != nil {
		t.Fatal(err)
	}
	return buf.String()
}
