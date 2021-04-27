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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pulumi/pulumi-postgresql/provider/v3/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
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

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	p := shimv1.NewProvider(postgresql.Provider().(*schema.Provider))

	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "postgresql",
		Description: "A Pulumi package for creating and managing postgresql cloud resources.",
		Keywords:    []string{"pulumi", "postgresql"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-postgresql",
		GitHubOrg:   "cyrilgdn",
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
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"postgresql_database":  {Tok: makeResource(mainMod, "Database")},
			"postgresql_extension": {Tok: makeResource(mainMod, "Extension")},
			"postgresql_grant":     {Tok: makeResource(mainMod, "Grant")},
			"postgresql_role":      {Tok: makeResource(mainMod, "Role")},
			"postgresql_schema":    {Tok: makeResource(mainMod, "Schema")},
			"postgresql_grant_role": {
				Tok: makeResource(mainMod, "GrantRole"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"grant_role": {
						CSharpName: "GrantRoleName",
					},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{},
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
				"Pulumi":                       "3.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				mainPkg: "PostgreSql",
			},
		},
	}

	prov.RenameResourceWithAlias("postgresql_default_privileges", makeResource(mainMod, "DefaultPrivileg"),
		makeResource(mainMod, "DefaultPrivileges"), mainMod, mainMod, nil)

	prov.SetAutonaming(255, "-")

	return prov
}
