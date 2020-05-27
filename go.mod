module knative.dev/net-contour

go 1.14

require (
	github.com/google/go-cmp v0.4.0
	github.com/gorilla/websocket v1.4.0
	github.com/mikefarah/yq/v3 v3.0.0-20200501003153-6fc3566acd3a
	github.com/projectcontour/contour v1.4.1-0.20200507033955-65d52b253570
	google.golang.org/grpc v1.28.1
	gopkg.in/yaml.v2 v2.2.8
	istio.io/client-go v0.0.0-20200505182340-146ba01d5357 // indirect
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	knative.dev/pkg v0.0.0-20200527024749-495174c96651
	knative.dev/serving v0.15.1-0.20200527113659-fd494b3d37a7
	knative.dev/test-infra v0.0.0-20200527032149-1fa26194e82d
)

replace (
	k8s.io/api => k8s.io/api v0.16.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.4
	k8s.io/client-go => k8s.io/client-go v0.16.4
	k8s.io/code-generator => k8s.io/code-generator v0.16.4
)
