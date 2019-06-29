// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nest_formula.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type NestFormula_FormulaType int32

const (
	NestFormula_MONO_FORMULA NestFormula_FormulaType = 0
	NestFormula_POLY_FORMULA NestFormula_FormulaType = 1
)

var NestFormula_FormulaType_name = map[int32]string{
	0: "MONO_FORMULA",
	1: "POLY_FORMULA",
}

var NestFormula_FormulaType_value = map[string]int32{
	"MONO_FORMULA": 0,
	"POLY_FORMULA": 1,
}

func (x NestFormula_FormulaType) String() string {
	return proto.EnumName(NestFormula_FormulaType_name, int32(x))
}

func (NestFormula_FormulaType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_897a04f0fec97625, []int{0, 0}
}

type NestFormula struct {
	OpString             string                  `protobuf:"bytes,1,opt,name=op_string,json=opString,proto3" json:"op_string,omitempty"`
	Poly                 []*NestFormula          `protobuf:"bytes,2,rep,name=poly,proto3" json:"poly,omitempty"`
	Mono                 int32                   `protobuf:"varint,3,opt,name=mono,proto3" json:"mono,omitempty"`
	Type                 NestFormula_FormulaType `protobuf:"varint,4,opt,name=type,proto3,enum=api.NestFormula_FormulaType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *NestFormula) Reset()         { *m = NestFormula{} }
func (m *NestFormula) String() string { return proto.CompactTextString(m) }
func (*NestFormula) ProtoMessage()    {}
func (*NestFormula) Descriptor() ([]byte, []int) {
	return fileDescriptor_897a04f0fec97625, []int{0}
}

func (m *NestFormula) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NestFormula.Unmarshal(m, b)
}
func (m *NestFormula) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NestFormula.Marshal(b, m, deterministic)
}
func (m *NestFormula) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NestFormula.Merge(m, src)
}
func (m *NestFormula) XXX_Size() int {
	return xxx_messageInfo_NestFormula.Size(m)
}
func (m *NestFormula) XXX_DiscardUnknown() {
	xxx_messageInfo_NestFormula.DiscardUnknown(m)
}

var xxx_messageInfo_NestFormula proto.InternalMessageInfo

func (m *NestFormula) GetOpString() string {
	if m != nil {
		return m.OpString
	}
	return ""
}

func (m *NestFormula) GetPoly() []*NestFormula {
	if m != nil {
		return m.Poly
	}
	return nil
}

func (m *NestFormula) GetMono() int32 {
	if m != nil {
		return m.Mono
	}
	return 0
}

func (m *NestFormula) GetType() NestFormula_FormulaType {
	if m != nil {
		return m.Type
	}
	return NestFormula_MONO_FORMULA
}

func init() {
	proto.RegisterEnum("api.NestFormula_FormulaType", NestFormula_FormulaType_name, NestFormula_FormulaType_value)
	proto.RegisterType((*NestFormula)(nil), "api.NestFormula")
}

func init() { proto.RegisterFile("nest_formula.proto", fileDescriptor_897a04f0fec97625) }

var fileDescriptor_897a04f0fec97625 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0x4b, 0x2d, 0x2e,
	0x89, 0x4f, 0xcb, 0x2f, 0xca, 0x2d, 0xcd, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4e, 0x2c, 0xc8, 0x54, 0x3a, 0xc9, 0xc8, 0xc5, 0xed, 0x97, 0x5a, 0x5c, 0xe2, 0x06, 0x91, 0x12,
	0x92, 0xe6, 0xe2, 0xcc, 0x2f, 0x88, 0x2f, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x97, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0xe2, 0xc8, 0x2f, 0x08, 0x06, 0xf3, 0x85, 0x54, 0xb8, 0x58, 0x0a, 0xf2, 0x73,
	0x2a, 0x25, 0x98, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0x04, 0xf4, 0x12, 0x0b, 0x32, 0xf5, 0x90, 0x34,
	0x07, 0x81, 0x65, 0x85, 0x84, 0xb8, 0x58, 0x72, 0xf3, 0xf3, 0xf2, 0x25, 0x98, 0x15, 0x18, 0x35,
	0x58, 0x83, 0xc0, 0x6c, 0x21, 0x03, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0x09, 0x16, 0x05, 0x46,
	0x0d, 0x3e, 0x23, 0x19, 0x74, 0x9d, 0x7a, 0x50, 0x3a, 0xa4, 0xb2, 0x20, 0x35, 0x08, 0xac, 0x52,
	0xc9, 0x90, 0x8b, 0x1b, 0x49, 0x50, 0x48, 0x80, 0x8b, 0xc7, 0xd7, 0xdf, 0xcf, 0x3f, 0xde, 0xcd,
	0x3f, 0xc8, 0x37, 0xd4, 0xc7, 0x51, 0x80, 0x01, 0x24, 0x12, 0xe0, 0xef, 0x13, 0x09, 0x17, 0x61,
	0x4c, 0x62, 0x03, 0xfb, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x76, 0xa1, 0xc1, 0xed,
	0x00, 0x00, 0x00,
}
