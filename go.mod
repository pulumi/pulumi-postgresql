module github.com/pulumi/pulumi-postgresql

go 1.12

replace (
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)

require (
	cloud.google.com/go v0.40.0 // indirect
	github.com/Nvveen/Gotty v0.0.0-20170406111628-a8b993ba6abd // indirect
	github.com/OpenPeeDeeP/depguard v0.0.0-20181229194401-1f388ab2d810 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/apache/thrift v0.12.0 // indirect
	github.com/chzyer/logex v1.1.11-0.20160617073814-96a4d311aa9b // indirect
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/etcd v3.3.13+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20190612170431-362f06ec6bc1 // indirect
	github.com/go-critic/go-critic v0.3.4 // indirect
	github.com/go-ini/ini v1.31.0 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/golangci/errcheck v0.0.0-20181223084120-ef45e06d44b6 // indirect
	github.com/golangci/go-tools v0.0.0-20190124090046-35a9f45a5db0 // indirect
	github.com/golangci/gocyclo v0.0.0-20180528144436-0a533e8fa43d // indirect
	github.com/golangci/gofmt v0.0.0-20181222123516-0b8337e80d98 // indirect
	github.com/golangci/golangci-lint v1.17.1 // indirect
	github.com/golangci/gosec v0.0.0-20180901114220-8afd9cbb6cfb // indirect
	github.com/golangci/lint-1 v0.0.0-20181222135242-d2cdd8c08219 // indirect
	github.com/golangci/revgrep v0.0.0-20180812185044-276a5c0a1039 // indirect
	github.com/google/pprof v0.0.0-20190515194954-54271f7e092f // indirect
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/gostaticanalysis/analysisutil v0.0.0-20190329151158-56bca42c7635 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.1 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-sockaddr v1.0.0 // indirect
	github.com/hashicorp/serf v0.8.2-0.20171022020050-c20a0b1b1ea9 // indirect
	github.com/hashicorp/terraform v0.12.6
	github.com/kisielk/errcheck v1.2.0 // indirect
	github.com/klauspost/compress v1.7.0 // indirect
	github.com/klauspost/cpuid v1.2.1 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pty v1.1.5 // indirect
	github.com/logrusorgru/aurora v0.0.0-20190428105938-cea283e61946 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/miekg/dns v1.0.14 // indirect
	github.com/nbutton23/zxcvbn-go v0.0.0-20180912185939-ae427f1e4c1d // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/openzipkin/zipkin-go v0.1.6 // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.0.0 // indirect
	github.com/pulumi/pulumi v0.17.28-0.20190731182900-6804d640fc7c
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190828172748-3f206601e7a1
	github.com/rogpeppe/fastuuid v1.1.0 // indirect
	github.com/russross/blackfriday v2.0.0+incompatible // indirect
	github.com/shirou/gopsutil v2.18.12+incompatible // indirect
	github.com/shurcooL/go v0.0.0-20190330031554-6713ea532688 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/smartystreets/assertions v0.0.0-20190215210624-980c5ac6f3ac // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.5 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.4.0 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	github.com/terraform-providers/terraform-provider-postgresql v1.1.0
	github.com/ugorji/go v1.1.5-pre // indirect
	github.com/valyala/fasthttp v1.3.0 // indirect
	go.etcd.io/bbolt v1.3.3 // indirect
	go.opencensus.io v0.22.0 // indirect
	golang.org/x/build v0.0.0-20190314133821-5284462c4bec // indirect
	golang.org/x/exp v0.0.0-20190510132918-efd6b22b2522 // indirect
	golang.org/x/image v0.0.0-20190616094056-33659d3de4f5 // indirect
	golang.org/x/mobile v0.0.0-20190607214518-6fa95d984e88 // indirect
	golang.org/x/mod v0.1.0 // indirect
	golang.org/x/net v0.0.0-20190613194153-d28f0bde5980 // indirect
	golang.org/x/sys v0.0.0-20190616124812-15dcb6c0061f // indirect
	golang.org/x/tools v0.0.0-20190614205625-5aca471b1d59 // indirect
	google.golang.org/appengine v1.6.1 // indirect
	google.golang.org/genproto v0.0.0-20190611190212-a7e196e89fd3 // indirect
	google.golang.org/grpc v1.21.1 // indirect
	gopkg.in/vmihailenco/msgpack.v2 v2.9.1 // indirect
	honnef.co/go/tools v0.0.0-20190614002413-cb51c254f01b // indirect
	mvdan.cc/unparam v0.0.0-20190310220240-1b9ccfa71afe // indirect
	sourcegraph.com/sqs/pbtypes v1.0.0 // indirect
)
