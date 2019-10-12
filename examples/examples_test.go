// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/testing/integration"
)

func TestAccDatabase(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "database"),
		})

	integration.ProgramTest(t, &test)
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/postgresql",
		},
	})

	return baseJS
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		Tracing: "https://tracing.pulumi-engineering.com/collector/api/v1/spans",
		Config: map[string]string{
			"mysql:endpoint": "127.0.0.1:3306",
			"mysql:username": "root",
		},
	}
}
