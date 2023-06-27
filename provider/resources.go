// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package postgresql

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/cyrilgdn/terraform-provider-postgresql/postgresql"
	"github.com/pulumi/pulumi-postgresql/provider/v3/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/x"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "postgresql"
	// modules:
	mainMod = "index"
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(postgresql.Provider())

	prov := tfbridge.ProviderInfo{
		P:                p,
		Name:             "postgresql",
		Description:      "A Pulumi package for creating and managing postgresql cloud resources.",
		Keywords:         []string{"pulumi", "postgresql"},
		License:          "Apache-2.0",
		Homepage:         "https://pulumi.io",
		Repository:       "https://github.com/pulumi/pulumi-postgresql",
		GitHubOrg:        "cyrilgdn",
		UpstreamRepoPath: "./upstream",
		Config: map[string]*tfbridge.SchemaInfo{
			"sslmode": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PGSSLMODE"},
				},
			},
			"connect_timeout": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PGCONNECT_TIMEOUT"},
					Value:   180,
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"postgresql_grant_role": {
				Tok: makeResource(mainMod, "GrantRole"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"grant_role": {
						CSharpName: "GrantRoleName",
					},
				},
			},
			"postgresql_grant": {
				Tok:                 makeResource(mainMod, "Grant"),
				DeleteBeforeReplace: true,
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				mainPkg: "PostgreSql",
			},
		},
	}

	prov.RenameResourceWithAlias("postgresql_default_privileges", makeResource(mainMod, "DefaultPrivileg"),
		makeResource(mainMod, "DefaultPrivileges"), mainMod, mainMod, &tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{
				Source: "postgresql_default_privileges.html.markdown",
			},
		})

	err := x.ComputeDefaults(&prov, x.TokensSingleModule("postgresql_", mainMod, x.MakeStandardToken(mainPkg)))
	contract.AssertNoErrorf(err, "Failed to map tokens")

	prov.SetAutonaming(255, "-")

	return prov
}
