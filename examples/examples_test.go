// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		Config: map[string]string{
			"postgresql:host": "127.0.0.1",
            "postgresql:port": "5432",
			"postgresql:username": "postgres",
			"postgresql:password": "password",
            "postgresql:sslmode": "disable",
		},
	}
}
