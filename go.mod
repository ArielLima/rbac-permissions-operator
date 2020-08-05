module github.com/openshift/rbac-permissions-operator

require (
	contrib.go.opencensus.io/exporter/ocagent v0.7.0 // indirect
	git.apache.org/thrift.git v0.12.0 // indirect
	github.com/Azure/go-autorest/autorest v0.11.2 // indirect
	github.com/apache/thrift v0.12.0 // indirect
	github.com/appscode/jsonpatch v0.0.0-20190108182946-7c0e3b262f30 // indirect
	github.com/census-instrumentation/opencensus-proto v0.3.0 // indirect
	github.com/coreos/prometheus-operator v0.41.0
	github.com/emicklei/go-restful v2.13.0+incompatible // indirect
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/go-openapi/spec v0.19.9
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/golang/lint v0.0.0-20180702182130-06c8688daad7 // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gophercloud/gophercloud v0.12.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/imdario/mergo v0.3.10 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mailru/easyjson v0.7.2 // indirect
	github.com/openshift/api v3.9.1-0.20190424152011-77b8897ec79a+incompatible
	github.com/openshift/client-go v0.0.0-20200116152001-92a2713fa240
	github.com/openshift/dedicated-admin-operator v0.0.0-20200610141258-cc7a464155de
	github.com/openshift/operator-custom-metrics v0.3.0
	github.com/openzipkin/zipkin-go v0.1.6 // indirect
	github.com/operator-framework/operator-sdk v0.12.0
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/prometheus/client_golang v1.1.0
	github.com/spf13/pflag v1.0.5
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	google.golang.org/api v0.30.0 // indirect
	google.golang.org/genproto v0.0.0-20200804151602-45615f50871c // indirect
	gopkg.in/resty.v1 v1.12.0 // indirect
	k8s.io/api v0.18.4
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/code-generator v0.18.6
	k8s.io/gengo v0.0.0-20200728071708-7794989d0000
	k8s.io/kube-openapi v0.0.0-20200727223308-4c7aaf529f79
	sigs.k8s.io/controller-runtime v0.3.0
	sigs.k8s.io/controller-tools v0.3.0
	sigs.k8s.io/testing_frameworks v0.1.1 // indirect
)

// Pinned to kubernetes-1.13.1
replace (
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20181213153335-0fe22c71c476
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.3
)

replace (
	github.com/coreos/prometheus-operator => github.com/coreos/prometheus-operator v0.29.0
	github.com/operator-framework/operator-sdk => github.com/operator-framework/operator-sdk v0.8.1
	k8s.io/api => k8s.io/api v0.18.3
	k8s.io/client-go => k8s.io/client-go v0.18.3
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20181117043124-c2090bec4d9b
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20180711000925-0cf8f7e6ed1d
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.1.10
	sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.1.11-0.20190411181648-9d55346c2bde
)

go 1.13
