// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/jwt_authn/v3/config.proto

package envoy_extensions_filters_http_jwt_authn_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type JwtProvider struct {
	Issuer    string   `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Audiences []string `protobuf:"bytes,2,rep,name=audiences,proto3" json:"audiences,omitempty"`
	// Types that are valid to be assigned to JwksSourceSpecifier:
	//	*JwtProvider_RemoteJwks
	//	*JwtProvider_LocalJwks
	JwksSourceSpecifier  isJwtProvider_JwksSourceSpecifier `protobuf_oneof:"jwks_source_specifier"`
	Forward              bool                              `protobuf:"varint,5,opt,name=forward,proto3" json:"forward,omitempty"`
	FromHeaders          []*JwtHeader                      `protobuf:"bytes,6,rep,name=from_headers,json=fromHeaders,proto3" json:"from_headers,omitempty"`
	FromParams           []string                          `protobuf:"bytes,7,rep,name=from_params,json=fromParams,proto3" json:"from_params,omitempty"`
	ForwardPayloadHeader string                            `protobuf:"bytes,8,opt,name=forward_payload_header,json=forwardPayloadHeader,proto3" json:"forward_payload_header,omitempty"`
	PayloadInMetadata    string                            `protobuf:"bytes,9,opt,name=payload_in_metadata,json=payloadInMetadata,proto3" json:"payload_in_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *JwtProvider) Reset()         { *m = JwtProvider{} }
func (m *JwtProvider) String() string { return proto.CompactTextString(m) }
func (*JwtProvider) ProtoMessage()    {}
func (*JwtProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_733511ccb445d825, []int{0}
}

func (m *JwtProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtProvider.Unmarshal(m, b)
}
func (m *JwtProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtProvider.Marshal(b, m, deterministic)
}
func (m *JwtProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtProvider.Merge(m, src)
}
func (m *JwtProvider) XXX_Size() int {
	return xxx_messageInfo_JwtProvider.Size(m)
}
func (m *JwtProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtProvider.DiscardUnknown(m)
}

var xxx_messageInfo_JwtProvider proto.InternalMessageInfo

func (m *JwtProvider) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *JwtProvider) GetAudiences() []string {
	if m != nil {
		return m.Audiences
	}
	return nil
}

type isJwtProvider_JwksSourceSpecifier interface {
	isJwtProvider_JwksSourceSpecifier()
}

type JwtProvider_RemoteJwks struct {
	RemoteJwks *RemoteJwks `protobuf:"bytes,3,opt,name=remote_jwks,json=remoteJwks,proto3,oneof"`
}

type JwtProvider_LocalJwks struct {
	LocalJwks *v3.DataSource `protobuf:"bytes,4,opt,name=local_jwks,json=localJwks,proto3,oneof"`
}

func (*JwtProvider_RemoteJwks) isJwtProvider_JwksSourceSpecifier() {}

func (*JwtProvider_LocalJwks) isJwtProvider_JwksSourceSpecifier() {}

func (m *JwtProvider) GetJwksSourceSpecifier() isJwtProvider_JwksSourceSpecifier {
	if m != nil {
		return m.JwksSourceSpecifier
	}
	return nil
}

func (m *JwtProvider) GetRemoteJwks() *RemoteJwks {
	if x, ok := m.GetJwksSourceSpecifier().(*JwtProvider_RemoteJwks); ok {
		return x.RemoteJwks
	}
	return nil
}

func (m *JwtProvider) GetLocalJwks() *v3.DataSource {
	if x, ok := m.GetJwksSourceSpecifier().(*JwtProvider_LocalJwks); ok {
		return x.LocalJwks
	}
	return nil
}

func (m *JwtProvider) GetForward() bool {
	if m != nil {
		return m.Forward
	}
	return false
}

func (m *JwtProvider) GetFromHeaders() []*JwtHeader {
	if m != nil {
		return m.FromHeaders
	}
	return nil
}

func (m *JwtProvider) GetFromParams() []string {
	if m != nil {
		return m.FromParams
	}
	return nil
}

func (m *JwtProvider) GetForwardPayloadHeader() string {
	if m != nil {
		return m.ForwardPayloadHeader
	}
	return ""
}

func (m *JwtProvider) GetPayloadInMetadata() string {
	if m != nil {
		return m.PayloadInMetadata
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*JwtProvider) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*JwtProvider_RemoteJwks)(nil),
		(*JwtProvider_LocalJwks)(nil),
	}
}

type RemoteJwks struct {
	HttpUri              *v3.HttpUri        `protobuf:"bytes,1,opt,name=http_uri,json=httpUri,proto3" json:"http_uri,omitempty"`
	CacheDuration        *duration.Duration `protobuf:"bytes,2,opt,name=cache_duration,json=cacheDuration,proto3" json:"cache_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RemoteJwks) Reset()         { *m = RemoteJwks{} }
func (m *RemoteJwks) String() string { return proto.CompactTextString(m) }
func (*RemoteJwks) ProtoMessage()    {}
func (*RemoteJwks) Descriptor() ([]byte, []int) {
	return fileDescriptor_733511ccb445d825, []int{1}
}

func (m *RemoteJwks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteJwks.Unmarshal(m, b)
}
func (m *RemoteJwks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteJwks.Marshal(b, m, deterministic)
}
func (m *RemoteJwks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteJwks.Merge(m, src)
}
func (m *RemoteJwks) XXX_Size() int {
	return xxx_messageInfo_RemoteJwks.Size(m)
}
func (m *RemoteJwks) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteJwks.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteJwks proto.InternalMessageInfo

func (m *RemoteJwks) GetHttpUri() *v3.HttpUri {
	if m != nil {
		return m.HttpUri
	}
	return nil
}

func (m *RemoteJwks) GetCacheDuration() *duration.Duration {
	if m != nil {
		return m.CacheDuration
	}
	return nil
}

type JwtHeader struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ValuePrefix          string   `protobuf:"bytes,2,opt,name=value_prefix,json=valuePrefix,proto3" json:"value_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JwtHeader) Reset()         { *m = JwtHeader{} }
func (m *JwtHeader) String() string { return proto.CompactTextString(m) }
func (*JwtHeader) ProtoMessage()    {}
func (*JwtHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_733511ccb445d825, []int{2}
}

func (m *JwtHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtHeader.Unmarshal(m, b)
}
func (m *JwtHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtHeader.Marshal(b, m, deterministic)
}
func (m *JwtHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtHeader.Merge(m, src)
}
func (m *JwtHeader) XXX_Size() int {
	return xxx_messageInfo_JwtHeader.Size(m)
}
func (m *JwtHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtHeader.DiscardUnknown(m)
}

var xxx_messageInfo_JwtHeader proto.InternalMessageInfo

func (m *JwtHeader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JwtHeader) GetValuePrefix() string {
	if m != nil {
		return m.ValuePrefix
	}
	return ""
}

type ProviderWithAudiences struct {
	ProviderName         string   `protobuf:"bytes,1,opt,name=provider_name,json=providerName,proto3" json:"provider_name,omitempty"`
	Audiences            []string `protobuf:"bytes,2,rep,name=audiences,proto3" json:"audiences,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderWithAudiences) Reset()         { *m = ProviderWithAudiences{} }
func (m *ProviderWithAudiences) String() string { return proto.CompactTextString(m) }
func (*ProviderWithAudiences) ProtoMessage()    {}
func (*ProviderWithAudiences) Descriptor() ([]byte, []int) {
	return fileDescriptor_733511ccb445d825, []int{3}
}

func (m *ProviderWithAudiences) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderWithAudiences.Unmarshal(m, b)
}
func (m *ProviderWithAudiences) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderWithAudiences.Marshal(b, m, deterministic)
}
func (m *ProviderWithAudiences) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderWithAudiences.Merge(m, src)
}
func (m *ProviderWithAudiences) XXX_Size() int {
	return xxx_messageInfo_ProviderWithAudiences.Size(m)
}
func (m *ProviderWithAudiences) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderWithAudiences.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderWithAudiences proto.InternalMessageInfo

func (m *ProviderWithAudiences) GetProviderName() string {
	if m != nil {
		return m.ProviderName
	}
	return ""
}

func (m *ProviderWithAudiences) GetAudiences() []string {
	if m != nil {
		return m.Audiences
	}
	return nil
}

type JwtRequirement struct {
	// Types that are valid to be assigned to RequiresType:
	//	*JwtRequirement_ProviderName
	//	*JwtRequirement_ProviderAndAudiences
	//	*JwtRequirement_RequiresAny
	//	*JwtRequirement_RequiresAll
	//	*JwtRequirement_AllowMissingOrFailed
	//	*JwtRequirement_AllowMissing
	RequiresType         isJwtRequirement_RequiresType `protobuf_oneof:"requires_type"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *JwtRequirement) Reset()         { *m = JwtRequirement{} }
func (m *JwtRequirement) String() string { return proto.CompactTextString(m) }
func (*JwtRequirement) ProtoMessage()    {}
func (*JwtRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_733511ccb445d825, []int{4}
}

func (m *JwtRequirement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtRequirement.Unmarshal(m, b)
}
func (m *JwtRequirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtRequirement.Marshal(b, m, deterministic)
}
func (m *JwtRequirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtRequirement.Merge(m, src)
}
func (m *JwtRequirement) XXX_Size() int {
	return xxx_messageInfo_JwtRequirement.Size(m)
}
func (m *JwtRequirement) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtRequirement.DiscardUnknown(m)
}

var xxx_messageInfo_JwtRequirement proto.InternalMessageInfo

type isJwtRequirement_RequiresType interface {
	isJwtRequirement_RequiresType()
}

type JwtRequirement_ProviderName struct {
	ProviderName string `protobuf:"bytes,1,opt,name=provider_name,json=providerName,proto3,oneof"`
}

type JwtRequirement_ProviderAndAudiences struct {
	ProviderAndAudiences *ProviderWithAudiences `protobuf:"bytes,2,opt,name=provider_and_audiences,json=providerAndAudiences,proto3,oneof"`
}

type JwtRequirement_RequiresAny struct {
	RequiresAny *JwtRequirementOrList `protobuf:"bytes,3,opt,name=requires_any,json=requiresAny,proto3,oneof"`
}

type JwtRequirement_RequiresAll struct {
	RequiresAll *JwtRequirementAndList `protobuf:"bytes,4,opt,name=requires_all,json=requiresAll,proto3,oneof"`
}

type JwtRequirement_AllowMissingOrFailed struct {
	AllowMissingOrFailed *empty.Empty `protobuf:"bytes,5,opt,name=allow_missing_or_failed,json=allowMissingOrFailed,proto3,oneof"`
}

type JwtRequirement_AllowMissing struct {
	AllowMissing *empty.Empty `protobuf:"bytes,6,opt,name=allow_missing,json=allowMissing,proto3,oneof"`
}

func (*JwtRequirement_ProviderName) isJwtRequirement_RequiresType() {}

func (*JwtRequirement_ProviderAndAudiences) isJwtRequirement_RequiresType() {}

func (*JwtRequirement_RequiresAny) isJwtRequirement_RequiresType() {}

func (*JwtRequirement_RequiresAll) isJwtRequirement_RequiresType() {}

func (*JwtRequirement_AllowMissingOrFailed) isJwtRequirement_RequiresType() {}

func (*JwtRequirement_AllowMissing) isJwtRequirement_RequiresType() {}

func (m *JwtRequirement) GetRequiresType() isJwtRequirement_RequiresType {
	if m != nil {
		return m.RequiresType
	}
	return nil
}

func (m *JwtRequirement) GetProviderName() string {
	if x, ok := m.GetRequiresType().(*JwtRequirement_ProviderName); ok {
		return x.ProviderName
	}
	return ""
}

func (m *JwtRequirement) GetProviderAndAudiences() *ProviderWithAudiences {
	if x, ok := m.GetRequiresType().(*JwtRequirement_ProviderAndAudiences); ok {
		return x.ProviderAndAudiences
	}
	return nil
}

func (m *JwtRequirement) GetRequiresAny() *JwtRequirementOrList {
	if x, ok := m.GetRequiresType().(*JwtRequirement_RequiresAny); ok {
		return x.RequiresAny
	}
	return nil
}

func (m *JwtRequirement) GetRequiresAll() *JwtRequirementAndList {
	if x, ok := m.GetRequiresType().(*JwtRequirement_RequiresAll); ok {
		return x.RequiresAll
	}
	return nil
}

func (m *JwtRequirement) GetAllowMissingOrFailed() *empty.Empty {
	if x, ok := m.GetRequiresType().(*JwtRequirement_AllowMissingOrFailed); ok {
		return x.AllowMissingOrFailed
	}
	return nil
}

func (m *JwtRequirement) GetAllowMissing() *empty.Empty {
	if x, ok := m.GetRequiresType().(*JwtRequirement_AllowMissing); ok {
		return x.AllowMissing
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*JwtRequirement) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*JwtRequirement_ProviderName)(nil),
		(*JwtRequirement_ProviderAndAudiences)(nil),
		(*JwtRequirement_RequiresAny)(nil),
		(*JwtRequirement_RequiresAll)(nil),
		(*JwtRequirement_AllowMissingOrFailed)(nil),
		(*JwtRequirement_AllowMissing)(nil),
	}
}

type JwtRequirementOrList struct {
	Requirements         []*JwtRequirement `protobuf:"bytes,1,rep,name=requirements,proto3" json:"requirements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JwtRequirementOrList) Reset()         { *m = JwtRequirementOrList{} }
func (m *JwtRequirementOrList) String() string { return proto.CompactTextString(m) }
func (*JwtRequirementOrList) ProtoMessage()    {}
func (*JwtRequirementOrList) Descriptor() ([]byte, []int) {
	return fileDescriptor_733511ccb445d825, []int{5}
}

func (m *JwtRequirementOrList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtRequirementOrList.Unmarshal(m, b)
}
func (m *JwtRequirementOrList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtRequirementOrList.Marshal(b, m, deterministic)
}
func (m *JwtRequirementOrList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtRequirementOrList.Merge(m, src)
}
func (m *JwtRequirementOrList) XXX_Size() int {
	return xxx_messageInfo_JwtRequirementOrList.Size(m)
}
func (m *JwtRequirementOrList) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtRequirementOrList.DiscardUnknown(m)
}

var xxx_messageInfo_JwtRequirementOrList proto.InternalMessageInfo

func (m *JwtRequirementOrList) GetRequirements() []*JwtRequirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

type JwtRequirementAndList struct {
	Requirements         []*JwtRequirement `protobuf:"bytes,1,rep,name=requirements,proto3" json:"requirements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JwtRequirementAndList) Reset()         { *m = JwtRequirementAndList{} }
func (m *JwtRequirementAndList) String() string { return proto.CompactTextString(m) }
func (*JwtRequirementAndList) ProtoMessage()    {}
func (*JwtRequirementAndList) Descriptor() ([]byte, []int) {
	return fileDescriptor_733511ccb445d825, []int{6}
}

func (m *JwtRequirementAndList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtRequirementAndList.Unmarshal(m, b)
}
func (m *JwtRequirementAndList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtRequirementAndList.Marshal(b, m, deterministic)
}
func (m *JwtRequirementAndList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtRequirementAndList.Merge(m, src)
}
func (m *JwtRequirementAndList) XXX_Size() int {
	return xxx_messageInfo_JwtRequirementAndList.Size(m)
}
func (m *JwtRequirementAndList) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtRequirementAndList.DiscardUnknown(m)
}

var xxx_messageInfo_JwtRequirementAndList proto.InternalMessageInfo

func (m *JwtRequirementAndList) GetRequirements() []*JwtRequirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

type RequirementRule struct {
	Match                *v31.RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	Requires             *JwtRequirement `protobuf:"bytes,2,opt,name=requires,proto3" json:"requires,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RequirementRule) Reset()         { *m = RequirementRule{} }
func (m *RequirementRule) String() string { return proto.CompactTextString(m) }
func (*RequirementRule) ProtoMessage()    {}
func (*RequirementRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_733511ccb445d825, []int{7}
}

func (m *RequirementRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequirementRule.Unmarshal(m, b)
}
func (m *RequirementRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequirementRule.Marshal(b, m, deterministic)
}
func (m *RequirementRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequirementRule.Merge(m, src)
}
func (m *RequirementRule) XXX_Size() int {
	return xxx_messageInfo_RequirementRule.Size(m)
}
func (m *RequirementRule) XXX_DiscardUnknown() {
	xxx_messageInfo_RequirementRule.DiscardUnknown(m)
}

var xxx_messageInfo_RequirementRule proto.InternalMessageInfo

func (m *RequirementRule) GetMatch() *v31.RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *RequirementRule) GetRequires() *JwtRequirement {
	if m != nil {
		return m.Requires
	}
	return nil
}

type FilterStateRule struct {
	Name                 string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Requires             map[string]*JwtRequirement `protobuf:"bytes,3,rep,name=requires,proto3" json:"requires,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FilterStateRule) Reset()         { *m = FilterStateRule{} }
func (m *FilterStateRule) String() string { return proto.CompactTextString(m) }
func (*FilterStateRule) ProtoMessage()    {}
func (*FilterStateRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_733511ccb445d825, []int{8}
}

func (m *FilterStateRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterStateRule.Unmarshal(m, b)
}
func (m *FilterStateRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterStateRule.Marshal(b, m, deterministic)
}
func (m *FilterStateRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterStateRule.Merge(m, src)
}
func (m *FilterStateRule) XXX_Size() int {
	return xxx_messageInfo_FilterStateRule.Size(m)
}
func (m *FilterStateRule) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterStateRule.DiscardUnknown(m)
}

var xxx_messageInfo_FilterStateRule proto.InternalMessageInfo

func (m *FilterStateRule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FilterStateRule) GetRequires() map[string]*JwtRequirement {
	if m != nil {
		return m.Requires
	}
	return nil
}

type JwtAuthentication struct {
	Providers            map[string]*JwtProvider `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Rules                []*RequirementRule      `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	FilterStateRules     *FilterStateRule        `protobuf:"bytes,3,opt,name=filter_state_rules,json=filterStateRules,proto3" json:"filter_state_rules,omitempty"`
	BypassCorsPreflight  bool                    `protobuf:"varint,4,opt,name=bypass_cors_preflight,json=bypassCorsPreflight,proto3" json:"bypass_cors_preflight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *JwtAuthentication) Reset()         { *m = JwtAuthentication{} }
func (m *JwtAuthentication) String() string { return proto.CompactTextString(m) }
func (*JwtAuthentication) ProtoMessage()    {}
func (*JwtAuthentication) Descriptor() ([]byte, []int) {
	return fileDescriptor_733511ccb445d825, []int{9}
}

func (m *JwtAuthentication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtAuthentication.Unmarshal(m, b)
}
func (m *JwtAuthentication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtAuthentication.Marshal(b, m, deterministic)
}
func (m *JwtAuthentication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtAuthentication.Merge(m, src)
}
func (m *JwtAuthentication) XXX_Size() int {
	return xxx_messageInfo_JwtAuthentication.Size(m)
}
func (m *JwtAuthentication) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtAuthentication.DiscardUnknown(m)
}

var xxx_messageInfo_JwtAuthentication proto.InternalMessageInfo

func (m *JwtAuthentication) GetProviders() map[string]*JwtProvider {
	if m != nil {
		return m.Providers
	}
	return nil
}

func (m *JwtAuthentication) GetRules() []*RequirementRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *JwtAuthentication) GetFilterStateRules() *FilterStateRule {
	if m != nil {
		return m.FilterStateRules
	}
	return nil
}

func (m *JwtAuthentication) GetBypassCorsPreflight() bool {
	if m != nil {
		return m.BypassCorsPreflight
	}
	return false
}

func init() {
	proto.RegisterType((*JwtProvider)(nil), "envoy.extensions.filters.http.jwt_authn.v3.JwtProvider")
	proto.RegisterType((*RemoteJwks)(nil), "envoy.extensions.filters.http.jwt_authn.v3.RemoteJwks")
	proto.RegisterType((*JwtHeader)(nil), "envoy.extensions.filters.http.jwt_authn.v3.JwtHeader")
	proto.RegisterType((*ProviderWithAudiences)(nil), "envoy.extensions.filters.http.jwt_authn.v3.ProviderWithAudiences")
	proto.RegisterType((*JwtRequirement)(nil), "envoy.extensions.filters.http.jwt_authn.v3.JwtRequirement")
	proto.RegisterType((*JwtRequirementOrList)(nil), "envoy.extensions.filters.http.jwt_authn.v3.JwtRequirementOrList")
	proto.RegisterType((*JwtRequirementAndList)(nil), "envoy.extensions.filters.http.jwt_authn.v3.JwtRequirementAndList")
	proto.RegisterType((*RequirementRule)(nil), "envoy.extensions.filters.http.jwt_authn.v3.RequirementRule")
	proto.RegisterType((*FilterStateRule)(nil), "envoy.extensions.filters.http.jwt_authn.v3.FilterStateRule")
	proto.RegisterMapType((map[string]*JwtRequirement)(nil), "envoy.extensions.filters.http.jwt_authn.v3.FilterStateRule.RequiresEntry")
	proto.RegisterType((*JwtAuthentication)(nil), "envoy.extensions.filters.http.jwt_authn.v3.JwtAuthentication")
	proto.RegisterMapType((map[string]*JwtProvider)(nil), "envoy.extensions.filters.http.jwt_authn.v3.JwtAuthentication.ProvidersEntry")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/jwt_authn/v3/config.proto", fileDescriptor_733511ccb445d825)
}

var fileDescriptor_733511ccb445d825 = []byte{
	// 1234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4b, 0x8f, 0x1b, 0x45,
	0x10, 0xde, 0xf1, 0xbe, 0xec, 0xf2, 0x6e, 0x1e, 0x9d, 0x6c, 0x62, 0x36, 0x21, 0x71, 0x1c, 0x21,
	0xad, 0x10, 0x1a, 0x4b, 0xce, 0x6b, 0xf1, 0x26, 0xb0, 0x76, 0x1e, 0x38, 0xab, 0x2c, 0x31, 0x13,
	0xf1, 0x3a, 0x8d, 0x7a, 0xed, 0xf6, 0xba, 0x93, 0xf1, 0xf4, 0xd0, 0xdd, 0x63, 0xc7, 0xff, 0x20,
	0xe2, 0xc8, 0x91, 0x2b, 0x67, 0xfe, 0x04, 0x12, 0x08, 0x71, 0xe3, 0x3f, 0x70, 0xe7, 0x8a, 0x72,
	0x42, 0xfd, 0x18, 0xdb, 0xb3, 0x6b, 0x05, 0x66, 0x83, 0xb8, 0xf5, 0x54, 0x75, 0x7d, 0xfd, 0xd5,
	0x57, 0xd5, 0x5d, 0x03, 0x77, 0x48, 0x38, 0x64, 0xe3, 0x2a, 0x79, 0x29, 0x49, 0x28, 0x28, 0x0b,
	0x45, 0xb5, 0x47, 0x03, 0x49, 0xb8, 0xa8, 0xf6, 0xa5, 0x8c, 0xaa, 0xcf, 0x47, 0xd2, 0xc7, 0xb1,
	0xec, 0x87, 0xd5, 0xe1, 0x8d, 0x6a, 0x87, 0x85, 0x3d, 0x7a, 0xe8, 0x46, 0x9c, 0x49, 0x86, 0xde,
	0xd7, 0x81, 0xee, 0x34, 0xd0, 0xb5, 0x81, 0xae, 0x0a, 0x74, 0x27, 0x81, 0xee, 0xf0, 0xc6, 0xe6,
	0x55, 0x73, 0x88, 0x89, 0xaf, 0x76, 0x18, 0x27, 0x0a, 0xee, 0x00, 0x0b, 0x62, 0xc0, 0x36, 0xaf,
	0xcf, 0xdd, 0xa0, 0x80, 0xfc, 0x98, 0x53, 0xbb, 0xe9, 0x83, 0xd4, 0x26, 0xce, 0x62, 0xa9, 0x77,
	0xe9, 0x85, 0xdf, 0x61, 0x83, 0x88, 0x85, 0x24, 0x94, 0xc2, 0xee, 0xbe, 0x72, 0xc8, 0xd8, 0x61,
	0x40, 0xaa, 0xfa, 0xeb, 0x20, 0xee, 0x55, 0xbb, 0x31, 0xc7, 0x92, 0xb2, 0xd0, 0xfa, 0x2f, 0x1d,
	0xf5, 0x93, 0x41, 0x24, 0xc7, 0xd6, 0x79, 0x2d, 0xee, 0x46, 0xb8, 0x8a, 0xc3, 0x90, 0x49, 0x1d,
	0x23, 0xaa, 0x43, 0xc2, 0x55, 0x96, 0x34, 0xb4, 0xf9, 0x6f, 0x5e, 0x1c, 0xe2, 0x80, 0x76, 0xb1,
	0xe2, 0x60, 0x17, 0xc6, 0x51, 0xf9, 0x69, 0x09, 0x8a, 0x7b, 0x23, 0xd9, 0xe6, 0x6c, 0x48, 0xbb,
	0x84, 0xa3, 0xab, 0xb0, 0x42, 0x85, 0x88, 0x09, 0x2f, 0x39, 0x65, 0x67, 0xab, 0xd0, 0x5c, 0x7d,
	0xdd, 0x5c, 0xe2, 0xb9, 0xb2, 0xe3, 0x59, 0x33, 0xba, 0x0c, 0x05, 0x1c, 0x77, 0x29, 0x09, 0x3b,
	0x44, 0x94, 0x72, 0xe5, 0xc5, 0xad, 0x82, 0x37, 0x35, 0xa0, 0xaf, 0xa1, 0xc8, 0xc9, 0x80, 0x49,
	0xe2, 0x3f, 0x1f, 0xbd, 0x10, 0xa5, 0xc5, 0xb2, 0xb3, 0x55, 0xac, 0xdd, 0x76, 0xff, 0xbd, 0xfa,
	0xae, 0xa7, 0xc3, 0xf7, 0x46, 0x2f, 0x44, 0x6b, 0xc1, 0x03, 0x3e, 0xf9, 0x42, 0x0d, 0x80, 0x80,
	0x75, 0x70, 0x60, 0x90, 0x97, 0x34, 0x72, 0xd9, 0x22, 0xdb, 0x5a, 0xab, 0x52, 0x28, 0x8c, 0x07,
	0x58, 0xe2, 0x67, 0x2c, 0xe6, 0x1d, 0xd2, 0x5a, 0xf0, 0x0a, 0x3a, 0x4a, 0x43, 0x94, 0x60, 0xb5,
	0xc7, 0xf8, 0x08, 0xf3, 0x6e, 0x69, 0xb9, 0xec, 0x6c, 0xe5, 0xbd, 0xe4, 0x13, 0x7d, 0x05, 0x6b,
	0x3d, 0xce, 0x06, 0x7e, 0x9f, 0xe0, 0x2e, 0xe1, 0xa2, 0xb4, 0x52, 0x5e, 0xdc, 0x2a, 0xd6, 0x6e,
	0x65, 0x21, 0xbe, 0x37, 0x92, 0x2d, 0x1d, 0xed, 0x15, 0x15, 0x94, 0x59, 0x0b, 0x74, 0x15, 0xf4,
	0xa7, 0x1f, 0x61, 0x8e, 0x07, 0xa2, 0xb4, 0xaa, 0x15, 0x03, 0x65, 0x6a, 0x6b, 0x0b, 0xba, 0x09,
	0x17, 0x2c, 0x0b, 0x3f, 0xc2, 0xe3, 0x80, 0xe1, 0xae, 0x65, 0x51, 0xca, 0xab, 0x0a, 0x78, 0xe7,
	0xad, 0xb7, 0x6d, 0x9c, 0x06, 0x17, 0xb9, 0x70, 0x2e, 0xd9, 0x4d, 0x43, 0x7f, 0x40, 0x24, 0xee,
	0x62, 0x89, 0x4b, 0x05, 0x1d, 0x72, 0xd6, 0xba, 0x1e, 0x87, 0xfb, 0xd6, 0x51, 0xbf, 0xf7, 0xfd,
	0xcf, 0xaf, 0xae, 0x6c, 0xc3, 0xed, 0x94, 0x5e, 0x26, 0x99, 0x63, 0xb9, 0xd4, 0x70, 0x10, 0xf5,
	0xb1, 0x3b, 0xd3, 0x16, 0xcd, 0xcb, 0xb0, 0xa1, 0x64, 0xf7, 0x85, 0x56, 0xd5, 0x17, 0x11, 0xe9,
	0xd0, 0x1e, 0x25, 0x1c, 0x2d, 0xfe, 0xd5, 0x74, 0x2a, 0xbf, 0x38, 0x00, 0xd3, 0xba, 0xa1, 0x6d,
	0xc8, 0x27, 0x97, 0x41, 0x77, 0x51, 0xb1, 0xf6, 0xee, 0xfc, 0x3a, 0xb5, 0xa4, 0x8c, 0x3e, 0xe7,
	0xd4, 0x5b, 0xed, 0x9b, 0x05, 0xda, 0x85, 0x53, 0x1d, 0xdc, 0xe9, 0x13, 0x3f, 0x69, 0xff, 0x52,
	0x4e, 0xc7, 0xbf, 0xe3, 0x9a, 0xfe, 0x77, 0x93, 0xfe, 0x77, 0x1f, 0xd8, 0x0d, 0xde, 0xba, 0x0e,
	0x48, 0x3e, 0xeb, 0x77, 0x55, 0x9e, 0x77, 0xe0, 0x56, 0x86, 0x3c, 0xa7, 0xcc, 0x2b, 0xaf, 0x1c,
	0x28, 0x4c, 0xea, 0x88, 0x2e, 0xc1, 0x52, 0x88, 0x07, 0xe4, 0xe8, 0x4d, 0xd0, 0x46, 0x74, 0x0d,
	0xd6, 0x86, 0x38, 0x88, 0x89, 0x1f, 0x71, 0xd2, 0xa3, 0x2f, 0x35, 0xd1, 0x82, 0x57, 0xd4, 0xb6,
	0xb6, 0x36, 0xd5, 0x77, 0x14, 0x97, 0xdb, 0x70, 0x33, 0x9b, 0xe6, 0xe6, 0xf0, 0xca, 0x0f, 0x0e,
	0x6c, 0x24, 0xf2, 0x7f, 0x49, 0x65, 0xbf, 0x31, 0xb9, 0x63, 0xd7, 0x61, 0x3d, 0xb2, 0x0e, 0x7f,
	0xca, 0xcf, 0x5b, 0x4b, 0x8c, 0x9f, 0x2a, 0x7a, 0x6f, 0xbc, 0xa6, 0xf5, 0x4f, 0x14, 0xb3, 0x26,
	0xec, 0x66, 0x60, 0x36, 0x97, 0x4b, 0xe5, 0x8f, 0x25, 0x38, 0xb5, 0x37, 0x92, 0x1e, 0xf9, 0x26,
	0xa6, 0x9c, 0x0c, 0x48, 0x28, 0xd1, 0x7b, 0x73, 0xe9, 0xb5, 0x16, 0x8e, 0x10, 0x1c, 0xc3, 0x85,
	0xc9, 0x36, 0x1c, 0x76, 0xfd, 0x59, 0xb6, 0xaa, 0xe4, 0x8d, 0x2c, 0x77, 0x6f, 0x2e, 0xb9, 0xd6,
	0x82, 0x77, 0x3e, 0x39, 0xa2, 0x11, 0x76, 0xa7, 0x02, 0x12, 0x58, 0xe3, 0x86, 0xb0, 0xf0, 0x71,
	0x38, 0xb6, 0xaf, 0xd4, 0x6e, 0xc6, 0xcb, 0x3e, 0x93, 0xf3, 0x53, 0xfe, 0x84, 0x0a, 0xd9, 0x5a,
	0xf0, 0x8a, 0x09, 0x6e, 0x23, 0x1c, 0xa3, 0xde, 0xec, 0x31, 0x41, 0x60, 0x9f, 0xac, 0xc6, 0xc9,
	0x8f, 0x69, 0x84, 0xdd, 0x63, 0xe7, 0x04, 0x01, 0x7a, 0x0a, 0x17, 0x71, 0x10, 0xb0, 0x91, 0x3f,
	0xa0, 0x42, 0xd0, 0xf0, 0xd0, 0x67, 0xdc, 0xef, 0x61, 0x1a, 0x10, 0xf3, 0xca, 0x15, 0x6b, 0x17,
	0x8e, 0xdd, 0x9e, 0x87, 0x6a, 0x7a, 0x28, 0x7d, 0x74, 0xe0, 0xbe, 0x89, 0x7b, 0xca, 0x1f, 0xe9,
	0x28, 0x74, 0x0f, 0xd6, 0x53, 0x80, 0xa5, 0x95, 0x7f, 0x80, 0x59, 0x9b, 0x85, 0xa9, 0xef, 0xaa,
	0xe6, 0xda, 0x81, 0x0f, 0xb3, 0xb5, 0xfd, 0x4c, 0x9e, 0xcd, 0xd3, 0xb0, 0x3e, 0x51, 0x4e, 0x8e,
	0x23, 0x52, 0xf9, 0xd5, 0x81, 0xf3, 0xf3, 0x24, 0x47, 0xfd, 0x89, 0xc6, 0xca, 0x28, 0x4a, 0x8e,
	0x7e, 0xb7, 0xeb, 0x27, 0xd7, 0xb8, 0x99, 0x7f, 0xdd, 0x5c, 0xfe, 0xce, 0xc9, 0xe5, 0x73, 0x5e,
	0x0a, 0xb9, 0xfe, 0x48, 0x65, 0xd5, 0x80, 0x8f, 0x4f, 0x9c, 0x95, 0x61, 0x5c, 0xf9, 0xcd, 0x81,
	0x8d, 0xb9, 0x65, 0xfd, 0x1f, 0x73, 0x39, 0xc1, 0xf5, 0x9f, 0x4b, 0xb9, 0xf2, 0xa7, 0x03, 0xa7,
	0x67, 0xcc, 0x5e, 0x1c, 0x10, 0xd4, 0x80, 0xe5, 0x01, 0x96, 0x9d, 0xbe, 0x7d, 0xfa, 0xaf, 0xa5,
	0x9f, 0x7e, 0xfd, 0xff, 0xa3, 0xe7, 0xbc, 0x5a, 0xec, 0xab, 0x8d, 0x9a, 0xe6, 0xb7, 0x4e, 0xee,
	0x8c, 0xe3, 0x99, 0x48, 0xf4, 0x05, 0xe4, 0x93, 0xfa, 0xdb, 0xd7, 0xe0, 0x2d, 0x54, 0xf0, 0x26,
	0x58, 0xf5, 0x86, 0xca, 0xfb, 0x2e, 0xd4, 0x33, 0x0d, 0x87, 0x54, 0x76, 0x95, 0xdf, 0x73, 0x70,
	0xfa, 0x91, 0x8e, 0x78, 0x26, 0xb1, 0x24, 0x3a, 0xe3, 0x37, 0xce, 0x09, 0x32, 0x93, 0xcb, 0xa2,
	0xae, 0xe8, 0xe3, 0x2c, 0xb9, 0x1c, 0x39, 0x2b, 0xe1, 0x23, 0x1e, 0x86, 0x92, 0x8f, 0xa7, 0xa9,
	0x6d, 0x8e, 0x60, 0x3d, 0xe5, 0x42, 0x67, 0x60, 0xf1, 0x05, 0x19, 0xdb, 0xd9, 0xa0, 0x96, 0xa8,
	0x0d, 0xcb, 0x7a, 0x3a, 0xfd, 0x07, 0x92, 0x1a, 0xa0, 0x7a, 0x6e, 0xdb, 0x39, 0x89, 0xa6, 0x47,
	0x72, 0xaa, 0xfc, 0xb8, 0x04, 0x67, 0xf7, 0x46, 0xb2, 0x11, 0xcb, 0x3e, 0x09, 0x25, 0xed, 0xe8,
	0x49, 0x8e, 0x9e, 0x43, 0x21, 0x79, 0xbd, 0x93, 0xbb, 0xf0, 0x24, 0x23, 0xe5, 0x34, 0xe2, 0x64,
	0x4a, 0x58, 0xf1, 0xa6, 0xf0, 0xe8, 0x33, 0x58, 0xe6, 0x71, 0x60, 0x27, 0x65, 0xb1, 0xb6, 0x93,
	0xed, 0x87, 0x35, 0xd5, 0x21, 0x9e, 0x41, 0x42, 0x14, 0x90, 0x89, 0xf1, 0x85, 0x4a, 0xd4, 0x37,
	0xf8, 0x66, 0xd4, 0xec, 0xbc, 0x45, 0x07, 0x78, 0x67, 0x7a, 0x69, 0x83, 0x40, 0x35, 0xd8, 0x38,
	0x18, 0x47, 0x58, 0x08, 0xbf, 0xc3, 0xb8, 0xd0, 0x3f, 0x24, 0x01, 0x3d, 0xec, 0x4b, 0x3d, 0x71,
	0xf2, 0xde, 0x39, 0xe3, 0xbc, 0xcf, 0xb8, 0x68, 0x27, 0xae, 0xcd, 0x18, 0x4e, 0xa5, 0xe5, 0x98,
	0xd3, 0x30, 0xfb, 0xe9, 0x86, 0xb9, 0x93, 0x51, 0xfd, 0x04, 0x7f, 0xb6, 0x5b, 0xee, 0xab, 0x6e,
	0xf9, 0x08, 0xee, 0x66, 0x7b, 0x79, 0xd2, 0x75, 0x6c, 0x3e, 0x86, 0x6d, 0xca, 0x0c, 0x99, 0x88,
	0xb3, 0x97, 0xe3, 0x0c, 0xbc, 0x9a, 0xc5, 0xfb, 0xfa, 0xcc, 0xb6, 0x1a, 0x61, 0x6d, 0xe7, 0x60,
	0x45, 0xcf, 0xb2, 0x1b, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xfc, 0x5a, 0x94, 0x59, 0x0e,
	0x00, 0x00,
}
