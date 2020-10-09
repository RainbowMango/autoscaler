// This is a generated file. Do not edit directly.
// Ensure you've carefully read
// https://git.k8s.io/community/contributors/devel/sig-architecture/vendor.md
// Run hack/pin-dependency.sh to change pinned dependency versions.
// Run hack/update-vendor.sh to update go.mod files and the vendor directory.

module k8s.io/autoscaler/cluster-autoscaler

go 1.15

require (
	cloud.google.com/go v0.68.0
	github.com/Azure/azure-sdk-for-go v46.4.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.10
	github.com/Azure/go-autorest/autorest/adal v0.9.5
	github.com/Azure/go-autorest/autorest/date v0.3.0
	github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/aws/aws-sdk-go v1.35.6
	github.com/digitalocean/godo v1.46.0
	github.com/ghodss/yaml v1.0.0
	github.com/golang/mock v1.4.4
	github.com/jmespath/go-jmespath v0.4.0
	github.com/json-iterator/go v1.1.10
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	google.golang.org/api v0.32.0
	gopkg.in/gcfg.v1 v1.2.3
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.20.0-alpha.1
	k8s.io/apimachinery v0.20.0-alpha.1
	k8s.io/apiserver v0.20.0-alpha.1
	k8s.io/client-go v0.20.0-alpha.1
	k8s.io/cloud-provider v0.20.0-alpha.1
	k8s.io/component-base v0.20.0-alpha.1
	k8s.io/klog/v2 v2.3.0
	k8s.io/kubernetes v1.20.0-alpha.1
	k8s.io/legacy-cloud-providers v0.0.0
	k8s.io/utils v0.0.0-20201005171033-6301aaf42dc7
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
	k8s.io/api => k8s.io/api v0.20.0-alpha.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.0-alpha.1
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.0-alpha.1
	k8s.io/apiserver => k8s.io/apiserver v0.20.0-alpha.1
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.0-alpha.1
	k8s.io/client-go => k8s.io/client-go v0.20.0-alpha.0
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.20.0-alpha.1
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.20.0-alpha.1
	k8s.io/code-generator => k8s.io/code-generator v0.20.0-alpha.1
	k8s.io/component-base => k8s.io/component-base v0.20.0-alpha.1
	k8s.io/cri-api => k8s.io/cri-api v0.20.0-alpha.1
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.20.0-alpha.1
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.20.0-alpha.1
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.20.0-alpha.1
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.20.0-alpha.1
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.20.0-alpha.1
	k8s.io/kubectl => k8s.io/kubectl v0.20.0-alpha.1
	k8s.io/kubelet => k8s.io/kubelet v0.20.0-alpha.1
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.20.0-alpha.1
	k8s.io/metrics => k8s.io/metrics v0.20.0-alpha.1
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.20.0-alpha.1
)
