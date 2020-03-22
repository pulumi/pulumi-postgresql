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
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/pulumi/pulumi-terraform-bridge/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/terraform-providers/terraform-provider-postgresql/postgresql"
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
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	p := postgresql.Provider().(*schema.Provider)

	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "postgresql",
		Description: "A Pulumi package for creating and managing postgresql cloud resources.",
		Keywords:    []string{"pulumi", "postgresql"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-postgresql",
		Config: map[string]*tfbridge.SchemaInfo{
			"host": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PGHOST"},
				},
			},
			"port": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PGPORT"},
					Value:   5432,
				},
			},
			"database": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PGDATABASE"},
					Value:   "postgres",
				},
			},
			"username": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PGUSER"},
					Value:   "postgres",
				},
			},
			"password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PGPASSWORD"},
				},
			},
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
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^1.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=1.0.0,<2.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "1.12.1-preview",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				mainPkg: "PostgreSql",
			},
		},
	}

	prov.RenameResourceWithAlias("postgresql_default_privileges", makeResource(mainMod, "DefaultPrivileg"),
		makeResource(mainMod, "DefaultPrivileges"), mainMod, mainMod, nil)

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
