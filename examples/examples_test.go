// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/testing/integration"
)

func TestExamples(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err, "expected a valid working directory: %v", err) {
		return
	}

	// base options shared amongst all tests.
	base := integration.ProgramTestOptions{
		Config: map[string]string{
			// Configuration map
		},
		Tracing: "https://tracing.pulumi-engineering.com/collector/api/v1/spans",
	}

	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			// JavaScript dependencies
		},
	})

	examples := []integration.ProgramTestOptions{
		// List each test
		baseJS.With(integration.ProgramTestOptions{
			Dir: path.Join(cwd, "database"),
			Config: map[string]string{
				"postgresql:host": "127.0.0.1",
				"postgresql:username": "postgres",
				"postgresql:sslmode": "disable",
			},
			Dependencies: []string{
				"@pulumi/postgresql",
			},
		}),
	}

	if !testing.Short() {
		// Append any longer running tests
	}

	for _, ex := range examples {
		example := ex
		t.Run(example.Dir, func(t *testing.T) {
			integration.ProgramTest(t, &example)
		})
	}
}
