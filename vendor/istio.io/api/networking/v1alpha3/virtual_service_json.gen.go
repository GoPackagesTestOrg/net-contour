// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/virtual_service.proto

// Configuration affecting traffic routing. Here are a few terms useful to define
// in the context of traffic routing.
//
// `Service` a unit of application behavior bound to a unique name in a
// service registry. Services consist of multiple network *endpoints*
// implemented by workload instances running on pods, containers, VMs etc.
//
// `Service versions (a.k.a. subsets)` - In a continuous deployment
// scenario, for a given service, there can be distinct subsets of
// instances running different variants of the application binary. These
// variants are not necessarily different API versions. They could be
// iterative changes to the same service, deployed in different
// environments (prod, staging, dev, etc.). Common scenarios where this
// occurs include A/B testing, canary rollouts, etc. The choice of a
// particular version can be decided based on various criterion (headers,
// url, etc.) and/or by weights assigned to each version. Each service has
// a default version consisting of all its instances.
//
// `Source` - A downstream client calling a service.
//
// `Host` - The address used by a client when attempting to connect to a
// service.
//
// `Access model` - Applications address only the destination service
// (Host) without knowledge of individual service versions (subsets). The
// actual choice of the version is determined by the proxy/sidecar, enabling the
// application code to decouple itself from the evolution of dependent
// services.
//
// A `VirtualService` defines a set of traffic routing rules to apply when a host is
// addressed. Each routing rule defines matching criteria for traffic of a specific
// protocol. If the traffic is matched, then it is sent to a named destination service
// (or subset/version of it) defined in the registry.
//
// The source of traffic can also be matched in a routing rule. This allows routing
// to be customized for specific client contexts.
//
// The following example on Kubernetes, routes all HTTP traffic by default to
// pods of the reviews service with label "version: v1". In addition,
// HTTP requests with path starting with /wpcatalog/ or /consumercatalog/ will
// be rewritten to /newcatalog and sent to pods with label "version: v2".
//
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: reviews-route
// spec:
//   hosts:
//   - reviews.prod.svc.cluster.local
//   http:
//   - name: "reviews-v2-routes"
//     match:
//     - uri:
//         prefix: "/wpcatalog"
//     - uri:
//         prefix: "/consumercatalog"
//     rewrite:
//       uri: "/newcatalog"
//     route:
//     - destination:
//         host: reviews.prod.svc.cluster.local
//         subset: v2
//   - name: "reviews-v1-route"
//     route:
//     - destination:
//         host: reviews.prod.svc.cluster.local
//         subset: v1
// ```
//
// A subset/version of a route destination is identified with a reference
// to a named service subset which must be declared in a corresponding
// `DestinationRule`.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: reviews-destination
// spec:
//   host: reviews.prod.svc.cluster.local
//   subsets:
//   - name: v1
//     labels:
//       version: v1
//   - name: v2
//     labels:
//       version: v2
// ```
//

package v1alpha3

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for VirtualService
func (this *VirtualService) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualService
func (this *VirtualService) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Destination
func (this *Destination) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Destination
func (this *Destination) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPRoute
func (this *HTTPRoute) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRoute
func (this *HTTPRoute) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Headers
func (this *Headers) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Headers
func (this *Headers) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Headers_HeaderOperations
func (this *Headers_HeaderOperations) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Headers_HeaderOperations
func (this *Headers_HeaderOperations) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TLSRoute
func (this *TLSRoute) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TLSRoute
func (this *TLSRoute) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TCPRoute
func (this *TCPRoute) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TCPRoute
func (this *TCPRoute) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPMatchRequest
func (this *HTTPMatchRequest) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPMatchRequest
func (this *HTTPMatchRequest) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPRouteDestination
func (this *HTTPRouteDestination) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRouteDestination
func (this *HTTPRouteDestination) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RouteDestination
func (this *RouteDestination) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteDestination
func (this *RouteDestination) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for L4MatchAttributes
func (this *L4MatchAttributes) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for L4MatchAttributes
func (this *L4MatchAttributes) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TLSMatchAttributes
func (this *TLSMatchAttributes) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TLSMatchAttributes
func (this *TLSMatchAttributes) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPRedirect
func (this *HTTPRedirect) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRedirect
func (this *HTTPRedirect) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPRewrite
func (this *HTTPRewrite) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRewrite
func (this *HTTPRewrite) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for StringMatch
func (this *StringMatch) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for StringMatch
func (this *StringMatch) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPRetry
func (this *HTTPRetry) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRetry
func (this *HTTPRetry) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CorsPolicy
func (this *CorsPolicy) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CorsPolicy
func (this *CorsPolicy) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPFaultInjection
func (this *HTTPFaultInjection) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPFaultInjection
func (this *HTTPFaultInjection) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPFaultInjection_Delay
func (this *HTTPFaultInjection_Delay) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPFaultInjection_Delay
func (this *HTTPFaultInjection_Delay) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPFaultInjection_Abort
func (this *HTTPFaultInjection_Abort) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPFaultInjection_Abort
func (this *HTTPFaultInjection_Abort) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for PortSelector
func (this *PortSelector) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PortSelector
func (this *PortSelector) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Percent
func (this *Percent) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Percent
func (this *Percent) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	VirtualServiceMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	VirtualServiceUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
