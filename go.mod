module github.com/pulumi/pulumi-postgresql

go 1.12

replace (
	git.apache.org/thrift.git => github.com/apache/thrift v0.12.0
	github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v31.1.0+incompatible
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

require (
	cloud.google.com/go v0.40.0 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/hashicorp/terraform v0.12.6
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.0.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190828172748-3f206601e7a1
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.5 // indirect
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	github.com/terraform-providers/terraform-provider-postgresql v1.1.0
	go.opencensus.io v0.22.0 // indirect
	golang.org/x/sys v0.0.0-20190616124812-15dcb6c0061f // indirect
	google.golang.org/appengine v1.6.1 // indirect
	google.golang.org/genproto v0.0.0-20190611190212-a7e196e89fd3 // indirect
)
