package main

import (
	"context"
	"flag"
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
