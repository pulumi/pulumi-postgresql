// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "postgresql:index/database:Database":
		r = &Database{}
	case "postgresql:index/defaultPrivileg:DefaultPrivileg":
		r = &DefaultPrivileg{}
	case "postgresql:index/defaultPrivileges:DefaultPrivileges":
		r = &DefaultPrivileges{}
	case "postgresql:index/extension:Extension":
		r = &Extension{}
	case "postgresql:index/grant:Grant":
		r = &Grant{}
	case "postgresql:index/grantRole:GrantRole":
		r = &GrantRole{}
	case "postgresql:index/physicalReplicationSlot:PhysicalReplicationSlot":
		r = &PhysicalReplicationSlot{}
	case "postgresql:index/replicationSlot:ReplicationSlot":
		r = &ReplicationSlot{}
	case "postgresql:index/role:Role":
		r = &Role{}
	case "postgresql:index/schema:Schema":
		r = &Schema{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:postgresql" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"postgresql",
		"index/database",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"postgresql",
		"index/defaultPrivileg",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"postgresql",
		"index/defaultPrivileges",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"postgresql",
		"index/extension",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"postgresql",
		"index/grant",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"postgresql",
		"index/grantRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"postgresql",
		"index/physicalReplicationSlot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"postgresql",
		"index/replicationSlot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"postgresql",
		"index/role",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"postgresql",
		"index/schema",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"postgresql",
		&pkg{version},
	)
}
