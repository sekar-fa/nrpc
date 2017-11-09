// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nrpc.proto

/*
Package nrpc is a generated protocol buffer package.

It is generated from these files:
	nrpc.proto

It has these top-level messages:
	Error
	Void
	NoReply
*/
package nrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SubjectRule int32

const (
	SubjectRule_COPY    SubjectRule = 0
	SubjectRule_TOLOWER SubjectRule = 1
)

var SubjectRule_name = map[int32]string{
	0: "COPY",
	1: "TOLOWER",
}
var SubjectRule_value = map[string]int32{
	"COPY":    0,
	"TOLOWER": 1,
}

func (x SubjectRule) String() string {
	return proto.EnumName(SubjectRule_name, int32(x))
}
func (SubjectRule) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Error_Type int32

const (
	Error_CLIENT Error_Type = 0
	Error_SERVER Error_Type = 1
)

var Error_Type_name = map[int32]string{
	0: "CLIENT",
	1: "SERVER",
}
var Error_Type_value = map[string]int32{
	"CLIENT": 0,
	"SERVER": 1,
}

func (x Error_Type) String() string {
	return proto.EnumName(Error_Type_name, int32(x))
}
func (Error_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Error struct {
	Type    Error_Type `protobuf:"varint,1,opt,name=type,enum=nrpc.Error_Type" json:"type,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Error) GetType() Error_Type {
	if m != nil {
		return m.Type
	}
	return Error_CLIENT
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type NoReply struct {
}

func (m *NoReply) Reset()                    { *m = NoReply{} }
func (m *NoReply) String() string            { return proto.CompactTextString(m) }
func (*NoReply) ProtoMessage()               {}
func (*NoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

var E_PackageSubject = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "nrpc.packageSubject",
	Tag:           "bytes,50000,opt,name=packageSubject",
	Filename:      "nrpc.proto",
}

var E_PackageSubjectParams = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: ([]string)(nil),
	Field:         50001,
	Name:          "nrpc.packageSubjectParams",
	Tag:           "bytes,50001,rep,name=packageSubjectParams",
	Filename:      "nrpc.proto",
}

var E_ServiceSubjectRule = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*SubjectRule)(nil),
	Field:         50002,
	Name:          "nrpc.serviceSubjectRule",
	Tag:           "varint,50002,opt,name=serviceSubjectRule,enum=nrpc.SubjectRule",
	Filename:      "nrpc.proto",
}

var E_MethodSubjectRule = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*SubjectRule)(nil),
	Field:         50003,
	Name:          "nrpc.methodSubjectRule",
	Tag:           "varint,50003,opt,name=methodSubjectRule,enum=nrpc.SubjectRule",
	Filename:      "nrpc.proto",
}

var E_ServiceSubject = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51000,
	Name:          "nrpc.serviceSubject",
	Tag:           "bytes,51000,opt,name=serviceSubject",
	Filename:      "nrpc.proto",
}

var E_ServiceSubjectParams = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: ([]string)(nil),
	Field:         51001,
	Name:          "nrpc.serviceSubjectParams",
	Tag:           "bytes,51001,rep,name=serviceSubjectParams",
	Filename:      "nrpc.proto",
}

var E_MethodSubject = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         52000,
	Name:          "nrpc.methodSubject",
	Tag:           "bytes,52000,opt,name=methodSubject",
	Filename:      "nrpc.proto",
}

var E_MethodSubjectParams = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: ([]string)(nil),
	Field:         52001,
	Name:          "nrpc.methodSubjectParams",
	Tag:           "bytes,52001,rep,name=methodSubjectParams",
	Filename:      "nrpc.proto",
}

func init() {
	proto.RegisterType((*Error)(nil), "nrpc.Error")
	proto.RegisterType((*Void)(nil), "nrpc.Void")
	proto.RegisterType((*NoReply)(nil), "nrpc.NoReply")
	proto.RegisterEnum("nrpc.SubjectRule", SubjectRule_name, SubjectRule_value)
	proto.RegisterEnum("nrpc.Error_Type", Error_Type_name, Error_Type_value)
	proto.RegisterExtension(E_PackageSubject)
	proto.RegisterExtension(E_PackageSubjectParams)
	proto.RegisterExtension(E_ServiceSubjectRule)
	proto.RegisterExtension(E_MethodSubjectRule)
	proto.RegisterExtension(E_ServiceSubject)
	proto.RegisterExtension(E_ServiceSubjectParams)
	proto.RegisterExtension(E_MethodSubject)
	proto.RegisterExtension(E_MethodSubjectParams)
}

func init() { proto.RegisterFile("nrpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x6b, 0x62, 0x12, 0x72, 0x2a, 0x2c, 0x77, 0xe8, 0xc2, 0x5c, 0x54, 0x2c, 0xab, 0x8b,
	0x88, 0x85, 0x2d, 0x95, 0xdd, 0x2c, 0xa9, 0x1c, 0xa9, 0x52, 0xa9, 0xab, 0x49, 0x28, 0x82, 0x0d,
	0xf2, 0x65, 0x70, 0x0d, 0x76, 0x66, 0x34, 0xb6, 0x91, 0xf2, 0x02, 0x59, 0x66, 0x99, 0x35, 0xbc,
	0x05, 0xbc, 0x05, 0x97, 0x17, 0x42, 0x1e, 0xdb, 0x52, 0x4c, 0x0c, 0x61, 0x37, 0x73, 0xce, 0xf9,
	0xbf, 0xf3, 0x9f, 0x1f, 0x60, 0x21, 0x78, 0x68, 0x73, 0xc1, 0x0a, 0x86, 0xd4, 0xea, 0xfd, 0xc8,
	0x8c, 0x19, 0x8b, 0x53, 0xea, 0xc8, 0x5a, 0x50, 0xbe, 0x77, 0x22, 0x9a, 0x87, 0x22, 0xe1, 0x05,
	0x13, 0xf5, 0x9c, 0x15, 0xc3, 0x5d, 0x57, 0x08, 0x26, 0xd0, 0x29, 0xa8, 0xc5, 0x92, 0x53, 0x43,
	0x31, 0x95, 0x89, 0x76, 0xa6, 0xdb, 0x92, 0x25, 0x5b, 0xf6, 0x7c, 0xc9, 0x29, 0x91, 0x5d, 0x64,
	0xc0, 0x28, 0xa3, 0x79, 0xee, 0xc7, 0xd4, 0xb8, 0x63, 0x2a, 0x93, 0x31, 0x69, 0xbf, 0xd6, 0x09,
	0xa8, 0xd5, 0x1c, 0x02, 0x18, 0x9e, 0x5f, 0x5e, 0xb8, 0x57, 0x73, 0xfd, 0xa0, 0x7a, 0xcf, 0x5c,
	0x72, 0xe3, 0x12, 0x5d, 0xb1, 0x86, 0xa0, 0xde, 0xb0, 0x24, 0xb2, 0xc6, 0x30, 0xba, 0x62, 0x84,
	0xf2, 0x74, 0xf9, 0xec, 0x14, 0x0e, 0x67, 0x65, 0xf0, 0x81, 0x86, 0x05, 0x29, 0x53, 0x8a, 0xee,
	0x81, 0x7a, 0xee, 0x5d, 0xbf, 0xd1, 0x0f, 0xd0, 0x21, 0x8c, 0xe6, 0xde, 0xa5, 0xf7, 0xba, 0x12,
	0xe2, 0x29, 0x68, 0xdc, 0x0f, 0x3f, 0xfa, 0x31, 0x6d, 0x86, 0xd1, 0x13, 0xbb, 0x3e, 0xcb, 0x6e,
	0xcf, 0xb2, 0xa7, 0x49, 0x4a, 0x3d, 0x5e, 0x24, 0x6c, 0x91, 0x1b, 0xdf, 0x57, 0x03, 0xe9, 0xec,
	0x0f, 0x15, 0x26, 0x70, 0xdc, 0xad, 0x5c, 0xfb, 0xc2, 0xcf, 0xf2, 0x3d, 0xb4, 0x1f, 0xab, 0x81,
	0x39, 0x98, 0x8c, 0x49, 0xaf, 0x16, 0xfb, 0x80, 0x72, 0x2a, 0x3e, 0x25, 0x21, 0xdd, 0x3e, 0xe4,
	0xdf, 0xc4, 0x9f, 0xd2, 0x9f, 0x76, 0x76, 0x54, 0x47, 0xbc, 0x25, 0x24, 0x3d, 0x30, 0xfc, 0x0e,
	0x8e, 0x32, 0x5a, 0xdc, 0xb2, 0xe8, 0xff, 0x37, 0xfc, 0xfa, 0xfb, 0x86, 0x5d, 0x16, 0xbe, 0x00,
	0xad, 0xbb, 0x16, 0x3d, 0xdd, 0xa1, 0xcf, 0xea, 0x81, 0x76, 0xc1, 0xd7, 0x75, 0x13, 0x71, 0x57,
	0x88, 0x5f, 0xc1, 0x71, 0xb7, 0xd2, 0x44, 0xbc, 0x17, 0xf8, 0x6d, 0xdd, 0xa4, 0xdc, 0x27, 0xc7,
	0x53, 0xb8, 0xdf, 0xb1, 0x8d, 0x4e, 0x76, 0x78, 0x2f, 0x65, 0xbf, 0xc5, 0x7d, 0xde, 0xd4, 0xfe,
	0xba, 0x32, 0x4c, 0xe0, 0x41, 0xa7, 0xd0, 0xb8, 0xdb, 0x47, 0xfb, 0xb2, 0xa9, 0xcd, 0xf5, 0x89,
	0x5f, 0x3c, 0x7e, 0xfb, 0x30, 0x4e, 0x8a, 0xdb, 0x32, 0xb0, 0x43, 0x96, 0x39, 0xc2, 0xe7, 0x49,
	0x94, 0x32, 0xc6, 0x9d, 0x2a, 0xf9, 0x60, 0x28, 0x89, 0xcf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x03, 0x85, 0xae, 0x67, 0x99, 0x03, 0x00, 0x00,
}
