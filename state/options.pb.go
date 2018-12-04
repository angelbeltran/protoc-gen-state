// Code generated by protoc-gen-go. DO NOT EDIT.
// source: state/options.proto

package state // import "github.com/tcncloud/protoc-gen-state/state"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StateMessageType int32

const (
	StateMessageType_REDUX_STATE   StateMessageType = 0
	StateMessageType_CUSTOM_ACTION StateMessageType = 1
	StateMessageType_EXTERNAL_LINK StateMessageType = 2
)

var StateMessageType_name = map[int32]string{
	0: "REDUX_STATE",
	1: "CUSTOM_ACTION",
	2: "EXTERNAL_LINK",
}
var StateMessageType_value = map[string]int32{
	"REDUX_STATE":   0,
	"CUSTOM_ACTION": 1,
	"EXTERNAL_LINK": 2,
}

func (x StateMessageType) Enum() *StateMessageType {
	p := new(StateMessageType)
	*p = x
	return p
}
func (x StateMessageType) String() string {
	return proto.EnumName(StateMessageType_name, int32(x))
}
func (x *StateMessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StateMessageType_value, data, "StateMessageType")
	if err != nil {
		return err
	}
	*x = StateMessageType(value)
	return nil
}
func (StateMessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_options_05d0aedec47de19f, []int{0}
}

type StringFieldOptions struct {
	Create               *string  `protobuf:"bytes,1,opt,name=create" json:"create,omitempty"`
	Get                  *string  `protobuf:"bytes,2,opt,name=get" json:"get,omitempty"`
	List                 *string  `protobuf:"bytes,3,opt,name=list" json:"list,omitempty"`
	Update               *string  `protobuf:"bytes,4,opt,name=update" json:"update,omitempty"`
	Delete               *string  `protobuf:"bytes,5,opt,name=delete" json:"delete,omitempty"`
	Custom               *string  `protobuf:"bytes,6,opt,name=custom" json:"custom,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringFieldOptions) Reset()         { *m = StringFieldOptions{} }
func (m *StringFieldOptions) String() string { return proto.CompactTextString(m) }
func (*StringFieldOptions) ProtoMessage()    {}
func (*StringFieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_options_05d0aedec47de19f, []int{0}
}
func (m *StringFieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringFieldOptions.Unmarshal(m, b)
}
func (m *StringFieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringFieldOptions.Marshal(b, m, deterministic)
}
func (dst *StringFieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringFieldOptions.Merge(dst, src)
}
func (m *StringFieldOptions) XXX_Size() int {
	return xxx_messageInfo_StringFieldOptions.Size(m)
}
func (m *StringFieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_StringFieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_StringFieldOptions proto.InternalMessageInfo

func (m *StringFieldOptions) GetCreate() string {
	if m != nil && m.Create != nil {
		return *m.Create
	}
	return ""
}

func (m *StringFieldOptions) GetGet() string {
	if m != nil && m.Get != nil {
		return *m.Get
	}
	return ""
}

func (m *StringFieldOptions) GetList() string {
	if m != nil && m.List != nil {
		return *m.List
	}
	return ""
}

func (m *StringFieldOptions) GetUpdate() string {
	if m != nil && m.Update != nil {
		return *m.Update
	}
	return ""
}

func (m *StringFieldOptions) GetDelete() string {
	if m != nil && m.Delete != nil {
		return *m.Delete
	}
	return ""
}

func (m *StringFieldOptions) GetCustom() string {
	if m != nil && m.Custom != nil {
		return *m.Custom
	}
	return ""
}

type IntFieldOptions struct {
	Create               *int64   `protobuf:"varint,1,opt,name=create" json:"create,omitempty"`
	Get                  *int64   `protobuf:"varint,2,opt,name=get" json:"get,omitempty"`
	List                 *int64   `protobuf:"varint,3,opt,name=list" json:"list,omitempty"`
	Update               *int64   `protobuf:"varint,4,opt,name=update" json:"update,omitempty"`
	Delete               *int64   `protobuf:"varint,5,opt,name=delete" json:"delete,omitempty"`
	Custom               *int64   `protobuf:"varint,6,opt,name=custom" json:"custom,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntFieldOptions) Reset()         { *m = IntFieldOptions{} }
func (m *IntFieldOptions) String() string { return proto.CompactTextString(m) }
func (*IntFieldOptions) ProtoMessage()    {}
func (*IntFieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_options_05d0aedec47de19f, []int{1}
}
func (m *IntFieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntFieldOptions.Unmarshal(m, b)
}
func (m *IntFieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntFieldOptions.Marshal(b, m, deterministic)
}
func (dst *IntFieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntFieldOptions.Merge(dst, src)
}
func (m *IntFieldOptions) XXX_Size() int {
	return xxx_messageInfo_IntFieldOptions.Size(m)
}
func (m *IntFieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_IntFieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_IntFieldOptions proto.InternalMessageInfo

func (m *IntFieldOptions) GetCreate() int64 {
	if m != nil && m.Create != nil {
		return *m.Create
	}
	return 0
}

func (m *IntFieldOptions) GetGet() int64 {
	if m != nil && m.Get != nil {
		return *m.Get
	}
	return 0
}

func (m *IntFieldOptions) GetList() int64 {
	if m != nil && m.List != nil {
		return *m.List
	}
	return 0
}

func (m *IntFieldOptions) GetUpdate() int64 {
	if m != nil && m.Update != nil {
		return *m.Update
	}
	return 0
}

func (m *IntFieldOptions) GetDelete() int64 {
	if m != nil && m.Delete != nil {
		return *m.Delete
	}
	return 0
}

func (m *IntFieldOptions) GetCustom() int64 {
	if m != nil && m.Custom != nil {
		return *m.Custom
	}
	return 0
}

type StateMessageOptions struct {
	Type                 *StateMessageType `protobuf:"varint,1,req,name=type,enum=state.StateMessageType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StateMessageOptions) Reset()         { *m = StateMessageOptions{} }
func (m *StateMessageOptions) String() string { return proto.CompactTextString(m) }
func (*StateMessageOptions) ProtoMessage()    {}
func (*StateMessageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_options_05d0aedec47de19f, []int{2}
}
func (m *StateMessageOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateMessageOptions.Unmarshal(m, b)
}
func (m *StateMessageOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateMessageOptions.Marshal(b, m, deterministic)
}
func (dst *StateMessageOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateMessageOptions.Merge(dst, src)
}
func (m *StateMessageOptions) XXX_Size() int {
	return xxx_messageInfo_StateMessageOptions.Size(m)
}
func (m *StateMessageOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_StateMessageOptions.DiscardUnknown(m)
}

var xxx_messageInfo_StateMessageOptions proto.InternalMessageInfo

func (m *StateMessageOptions) GetType() StateMessageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return StateMessageType_REDUX_STATE
}

var E_StateOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*StateMessageOptions)(nil),
	Field:         550020,
	Name:          "state.state_options",
	Tag:           "bytes,550020,opt,name=state_options,json=stateOptions",
	Filename:      "state/options.proto",
}

var E_Timeout = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         550001,
	Name:          "state.timeout",
	Tag:           "varint,550001,opt,name=timeout",
	Filename:      "state/options.proto",
}

var E_Retries = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         550002,
	Name:          "state.retries",
	Tag:           "varint,550002,opt,name=retries",
	Filename:      "state/options.proto",
}

var E_Method = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*StringFieldOptions)(nil),
	Field:         550003,
	Name:          "state.method",
	Tag:           "bytes,550003,opt,name=method",
	Filename:      "state/options.proto",
}

var E_MethodTimeout = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*IntFieldOptions)(nil),
	Field:         550004,
	Name:          "state.method_timeout",
	Tag:           "bytes,550004,opt,name=method_timeout,json=methodTimeout",
	Filename:      "state/options.proto",
}

var E_MethodRetries = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*IntFieldOptions)(nil),
	Field:         550005,
	Name:          "state.method_retries",
	Tag:           "bytes,550005,opt,name=method_retries,json=methodRetries",
	Filename:      "state/options.proto",
}

var E_DefaultTimeout = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         550006,
	Name:          "state.default_timeout",
	Tag:           "varint,550006,opt,name=default_timeout,json=defaultTimeout",
	Filename:      "state/options.proto",
}

var E_DefaultRetries = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         550007,
	Name:          "state.default_retries",
	Tag:           "varint,550007,opt,name=default_retries,json=defaultRetries",
	Filename:      "state/options.proto",
}

var E_Debug = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         550008,
	Name:          "state.debug",
	Tag:           "varint,550008,opt,name=debug",
	Filename:      "state/options.proto",
}

var E_Port = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         550010,
	Name:          "state.port",
	Tag:           "varint,550010,opt,name=port",
	Filename:      "state/options.proto",
}

var E_Debounce = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         550011,
	Name:          "state.debounce",
	Tag:           "varint,550011,opt,name=debounce",
	Filename:      "state/options.proto",
}

var E_ProtocTsPath = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         550012,
	Name:          "state.protoc_ts_path",
	Tag:           "bytes,550012,opt,name=protoc_ts_path,json=protocTsPath",
	Filename:      "state/options.proto",
}

var E_Hostname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         550013,
	Name:          "state.hostname",
	Tag:           "bytes,550013,opt,name=hostname",
	Filename:      "state/options.proto",
}

var E_HostnameLocation = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         550014,
	Name:          "state.hostname_location",
	Tag:           "bytes,550014,opt,name=hostname_location,json=hostnameLocation",
	Filename:      "state/options.proto",
}

var E_AuthTokenLocation = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         550015,
	Name:          "state.auth_token_location",
	Tag:           "bytes,550015,opt,name=auth_token_location,json=authTokenLocation",
	Filename:      "state/options.proto",
}

func init() {
	proto.RegisterType((*StringFieldOptions)(nil), "state.StringFieldOptions")
	proto.RegisterType((*IntFieldOptions)(nil), "state.IntFieldOptions")
	proto.RegisterType((*StateMessageOptions)(nil), "state.StateMessageOptions")
	proto.RegisterEnum("state.StateMessageType", StateMessageType_name, StateMessageType_value)
	proto.RegisterExtension(E_StateOptions)
	proto.RegisterExtension(E_Timeout)
	proto.RegisterExtension(E_Retries)
	proto.RegisterExtension(E_Method)
	proto.RegisterExtension(E_MethodTimeout)
	proto.RegisterExtension(E_MethodRetries)
	proto.RegisterExtension(E_DefaultTimeout)
	proto.RegisterExtension(E_DefaultRetries)
	proto.RegisterExtension(E_Debug)
	proto.RegisterExtension(E_Port)
	proto.RegisterExtension(E_Debounce)
	proto.RegisterExtension(E_ProtocTsPath)
	proto.RegisterExtension(E_Hostname)
	proto.RegisterExtension(E_HostnameLocation)
	proto.RegisterExtension(E_AuthTokenLocation)
}

func init() { proto.RegisterFile("state/options.proto", fileDescriptor_options_05d0aedec47de19f) }

var fileDescriptor_options_05d0aedec47de19f = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4b, 0x4f, 0xdb, 0x40,
	0x10, 0xae, 0x71, 0xa0, 0xb0, 0x40, 0x30, 0x46, 0xa2, 0x2e, 0x6a, 0x55, 0xe0, 0x84, 0x5a, 0xe1,
	0x54, 0x51, 0x2f, 0x75, 0x4f, 0x3c, 0xd2, 0x2a, 0x22, 0x24, 0x95, 0x63, 0x24, 0xd4, 0x43, 0x2d,
	0xc7, 0x5e, 0x6c, 0xab, 0x8e, 0xd7, 0xb2, 0xc7, 0x07, 0xee, 0xfd, 0x0f, 0x95, 0xfa, 0x8b, 0xe8,
	0xcf, 0xe9, 0xbb, 0xea, 0x5b, 0xfb, 0xb0, 0x9b, 0x86, 0x80, 0x73, 0xb1, 0x76, 0xc6, 0xdf, 0xf7,
	0xcd, 0x37, 0x33, 0x2b, 0x2d, 0x5a, 0xcb, 0xc0, 0x01, 0xdc, 0x20, 0x09, 0x84, 0x24, 0xce, 0xf4,
	0x24, 0x25, 0x40, 0xd4, 0x59, 0x96, 0xdc, 0xd8, 0xf4, 0x09, 0xf1, 0x23, 0xdc, 0x60, 0xc9, 0x41,
	0x7e, 0xd6, 0xf0, 0x70, 0xe6, 0xa6, 0x61, 0x02, 0x24, 0xe5, 0xc0, 0xed, 0xb7, 0x12, 0x52, 0xfb,
	0x90, 0x86, 0xb1, 0xff, 0x34, 0xc4, 0x91, 0xd7, 0xe3, 0x2a, 0xea, 0x3a, 0x9a, 0x73, 0x53, 0xec,
	0x00, 0xd6, 0xa4, 0x4d, 0x69, 0x67, 0xc1, 0x14, 0x91, 0xaa, 0x20, 0xd9, 0xc7, 0xa0, 0xcd, 0xb0,
	0x24, 0x3d, 0xaa, 0x2a, 0xaa, 0x45, 0x61, 0x06, 0x9a, 0xcc, 0x52, 0xec, 0x4c, 0xd9, 0x79, 0xe2,
	0x51, 0x76, 0x8d, 0xb3, 0x79, 0x44, 0xf3, 0x1e, 0x8e, 0x30, 0x60, 0x6d, 0x96, 0xe7, 0x79, 0xc4,
	0xaa, 0xe5, 0x19, 0x90, 0xa1, 0x36, 0x27, 0xaa, 0xb1, 0x68, 0xfb, 0x8d, 0x84, 0x56, 0xda, 0x31,
	0x5c, 0xe3, 0x4c, 0x9e, 0xe4, 0x4c, 0xbe, 0xec, 0x4c, 0x9e, 0xe8, 0x4c, 0xbe, 0xc2, 0x99, 0x7c,
	0x85, 0x33, 0xb9, 0x74, 0xb6, 0x8f, 0xd6, 0xfa, 0x74, 0xc2, 0xc7, 0x38, 0xcb, 0x1c, 0x1f, 0x17,
	0xe6, 0x1e, 0xa0, 0x1a, 0x9c, 0x27, 0xd4, 0xda, 0xcc, 0x4e, 0xbd, 0x79, 0x4b, 0x67, 0x5b, 0xd0,
	0x47, 0x91, 0xd6, 0x79, 0x82, 0x4d, 0x06, 0xba, 0xdf, 0x46, 0xca, 0xf8, 0x1f, 0x75, 0x05, 0x2d,
	0x9a, 0xad, 0xc3, 0x93, 0x53, 0xbb, 0x6f, 0xed, 0x59, 0x2d, 0xe5, 0x86, 0xba, 0x8a, 0x96, 0x0f,
	0x4e, 0xfa, 0x56, 0xef, 0xd8, 0xde, 0x3b, 0xb0, 0xda, 0xbd, 0xae, 0x22, 0xd1, 0x54, 0xeb, 0xd4,
	0x6a, 0x99, 0xdd, 0xbd, 0x8e, 0xdd, 0x69, 0x77, 0x8f, 0x94, 0x19, 0x63, 0x80, 0x96, 0x59, 0x29,
	0x5b, 0xdc, 0x02, 0xf5, 0x9e, 0xce, 0x37, 0xaf, 0x17, 0x9b, 0xd7, 0xff, 0x77, 0xaa, 0xbd, 0x7e,
	0xb7, 0xb5, 0x29, 0xed, 0x2c, 0x36, 0x37, 0x26, 0x78, 0x14, 0x18, 0x73, 0x89, 0xfd, 0x12, 0x91,
	0xf1, 0x18, 0xdd, 0x84, 0x70, 0x88, 0x49, 0x0e, 0xea, 0xdd, 0x4b, 0xea, 0xa3, 0x2b, 0xd2, 0xde,
	0x5f, 0x6c, 0xb1, 0x61, 0x15, 0x78, 0x4a, 0x4d, 0x31, 0xa4, 0x21, 0xce, 0xaa, 0xa8, 0x1f, 0x0a,
	0xaa, 0xc0, 0x1b, 0x26, 0x9a, 0x1b, 0x62, 0x08, 0x88, 0x57, 0xc5, 0xfc, 0x78, 0xc1, 0x1b, 0xba,
	0x5d, 0x36, 0x34, 0x7e, 0xa9, 0x4d, 0xa1, 0x64, 0xbc, 0x44, 0x75, 0x7e, 0xb2, 0xa7, 0x6c, 0xe8,
	0x93, 0xd0, 0x5e, 0x17, 0xda, 0x63, 0x77, 0xd2, 0x5c, 0xe6, 0x72, 0x96, 0x68, 0xf7, 0x9f, 0xfe,
	0x94, 0x5d, 0x7f, 0x9e, 0x4e, 0xdf, 0x14, 0x33, 0x79, 0x86, 0x56, 0x3c, 0x7c, 0xe6, 0xe4, 0x11,
	0x94, 0x0d, 0xdc, 0x99, 0x50, 0x20, 0x2a, 0x97, 0xfd, 0x45, 0x4c, 0xb5, 0x2e, 0x68, 0x85, 0xd1,
	0x11, 0xa1, 0xc2, 0xe9, 0xf5, 0x42, 0x5f, 0xc7, 0x84, 0x0a, 0x47, 0x8f, 0xd0, 0xac, 0x87, 0x07,
	0xb9, 0x5f, 0x41, 0xff, 0xc6, 0xe8, 0xf3, 0x26, 0x07, 0x1b, 0x4d, 0x54, 0x4b, 0x48, 0x5a, 0x65,
	0xfe, 0xbb, 0xa8, 0xc9, 0xb0, 0x86, 0x81, 0xe6, 0x3d, 0x3c, 0x20, 0x79, 0xec, 0xe2, 0x0a, 0xde,
	0x0f, 0xc1, 0x2b, 0xf1, 0xc6, 0x21, 0xaa, 0x33, 0xa4, 0x6b, 0x43, 0x66, 0x27, 0x0e, 0x04, 0x15,
	0x0a, 0x3f, 0x99, 0xc2, 0x82, 0xb9, 0xc4, 0x59, 0x56, 0xf6, 0xdc, 0x81, 0x80, 0x3a, 0x08, 0x48,
	0x06, 0xb1, 0x33, 0xac, 0x72, 0xf0, 0x4b, 0xf0, 0x4b, 0xbc, 0x71, 0x84, 0x56, 0x8b, 0xb3, 0x1d,
	0x11, 0xd7, 0xa1, 0xb8, 0x0a, 0x91, 0xdf, 0x42, 0x44, 0x29, 0x88, 0x1d, 0xc1, 0x33, 0xba, 0x68,
	0xcd, 0xc9, 0x21, 0xb0, 0x81, 0xbc, 0xc2, 0xf1, 0xb4, 0x72, 0x7f, 0x84, 0xdc, 0x2a, 0xa5, 0x5a,
	0x94, 0x59, 0xe8, 0xed, 0x37, 0x5f, 0x3c, 0xf4, 0x43, 0x08, 0xf2, 0x81, 0xee, 0x92, 0x61, 0x03,
	0xdc, 0xd8, 0x8d, 0x48, 0xee, 0xf1, 0xb7, 0xc3, 0xdd, 0xf5, 0x71, 0xbc, 0xcb, 0x5f, 0x1a, 0xf6,
	0x7d, 0xc2, 0xbe, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xb9, 0x13, 0x13, 0x7d, 0x06, 0x00,
	0x00,
}
