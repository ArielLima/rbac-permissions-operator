module github.com/openshift/rbac-permissions-operator

require (
	cloud.google.com/go v0.62.0 // indirect
	github.com/Azure/go-autorest/autorest v0.11.2 // indirect
	github.com/appscode/jsonpatch v0.0.0-20190108182946-7c0e3b262f30 // indirect
	github.com/coreos/prometheus-operator v0.41.0
	github.com/emicklei/go-restful v2.13.0+incompatible // indirect
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/go-openapi/spec v0.19.9
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gophercloud/gophercloud v0.12.0 // indirect
	github.com/imdario/mergo v0.3.10 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mailru/easyjson v0.7.2 // indirect
	github.com/openshift/api v3.9.1-0.20190424152011-77b8897ec79a+incompatible
	github.com/openshift/client-go v0.0.0-20200116152001-92a2713fa240
	github.com/openshift/operator-custom-metrics v0.3.0
	github.com/operator-framework/operator-sdk v0.12.0
	github.com/prometheus/client_golang v1.1.0
	github.com/spf13/pflag v1.0.5
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/sys v0.0.0-20200803210538-64077c9b5642 // indirect
	golang.org/x/tools v0.0.0-20200804011535-6c149bb5ef0d // indirect
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
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191008115822-1210218b4a26
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20181117043124-c2090bec4d9b
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20180711000925-0cf8f7e6ed1d
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.1.10
	sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.1.11-0.20190411181648-9d55346c2bde
)

go 1.13
