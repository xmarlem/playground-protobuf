// Code generated by protoc-gen-go. DO NOT EDIT.
// source: employee.proto

package employee

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

type Employee struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Employee) Reset()         { *m = Employee{} }
func (m *Employee) String() string { return proto.CompactTextString(m) }
func (*Employee) ProtoMessage()    {}
func (*Employee) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{0}
}

func (m *Employee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Employee.Unmarshal(m, b)
}
func (m *Employee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Employee.Marshal(b, m, deterministic)
}
func (m *Employee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Employee.Merge(m, src)
}
func (m *Employee) XXX_Size() int {
	return xxx_messageInfo_Employee.Size(m)
}
func (m *Employee) XXX_DiscardUnknown() {
	xxx_messageInfo_Employee.DiscardUnknown(m)
}

var xxx_messageInfo_Employee proto.InternalMessageInfo

func (m *Employee) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Employee) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Employee) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GetAllResponse struct {
	Employee             []*Employee `protobuf:"bytes,1,rep,name=employee,proto3" json:"employee,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetAllResponse) Reset()         { *m = GetAllResponse{} }
func (m *GetAllResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllResponse) ProtoMessage()    {}
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{1}
}

func (m *GetAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllResponse.Unmarshal(m, b)
}
func (m *GetAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllResponse.Marshal(b, m, deterministic)
}
func (m *GetAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllResponse.Merge(m, src)
}
func (m *GetAllResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllResponse.Size(m)
}
func (m *GetAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllResponse proto.InternalMessageInfo

func (m *GetAllResponse) GetEmployee() []*Employee {
	if m != nil {
		return m.Employee
	}
	return nil
}

type GetAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllRequest) Reset()         { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()    {}
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{2}
}

func (m *GetAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllRequest.Unmarshal(m, b)
}
func (m *GetAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllRequest.Marshal(b, m, deterministic)
}
func (m *GetAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllRequest.Merge(m, src)
}
func (m *GetAllRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllRequest.Size(m)
}
func (m *GetAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Employee)(nil), "Employee")
	proto.RegisterType((*GetAllResponse)(nil), "GetAllResponse")
	proto.RegisterType((*GetAllRequest)(nil), "GetAllRequest")
}

func init() { proto.RegisterFile("employee.proto", fileDescriptor_eb50a19aa79a6eac) }

var fileDescriptor_eb50a19aa79a6eac = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcd, 0x2d, 0xc8,
	0xc9, 0xaf, 0x4c, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x0a, 0xe1, 0xe2, 0x70, 0x85,
	0x8a, 0x08, 0xf1, 0x71, 0x31, 0x79, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x31, 0x79,
	0xba, 0x08, 0xc9, 0x70, 0x71, 0xa6, 0x65, 0x16, 0x15, 0x97, 0xf8, 0x25, 0xe6, 0xa6, 0x4a, 0x30,
	0x29, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x04, 0x84, 0xa4, 0xb8, 0x38, 0x72, 0x12, 0xa1, 0x92, 0xcc,
	0x60, 0x49, 0x38, 0x5f, 0xc9, 0x9c, 0x8b, 0xcf, 0x3d, 0xb5, 0xc4, 0x31, 0x27, 0x27, 0x28, 0xb5,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x95, 0x8b, 0x03, 0x66, 0xb3, 0x04, 0xa3, 0x02, 0xb3,
	0x06, 0xb7, 0x11, 0xa7, 0x1e, 0xcc, 0xe2, 0x20, 0xb8, 0x94, 0x12, 0x3f, 0x17, 0x2f, 0x4c, 0x63,
	0x61, 0x69, 0x6a, 0x71, 0x89, 0x91, 0x0d, 0x17, 0x3f, 0x4c, 0x59, 0x70, 0x6a, 0x51, 0x59, 0x66,
	0x72, 0xaa, 0x90, 0x26, 0x17, 0x1b, 0x44, 0x8d, 0x10, 0x9f, 0x1e, 0x8a, 0x62, 0x29, 0x7e, 0x3d,
	0x54, 0x5b, 0x93, 0xd8, 0xc0, 0x9e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xea, 0xe5,
	0x97, 0xf6, 0x00, 0x00, 0x00,
}