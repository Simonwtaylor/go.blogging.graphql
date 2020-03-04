// Code generated by protoc-gen-go. DO NOT EDIT.
// source: servers/blogging/proto/blogging.proto

package go_blogging_graphql

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

type AddPostRequest struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPostRequest) Reset()         { *m = AddPostRequest{} }
func (m *AddPostRequest) String() string { return proto.CompactTextString(m) }
func (*AddPostRequest) ProtoMessage()    {}
func (*AddPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12e57f654310df34, []int{0}
}

func (m *AddPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPostRequest.Unmarshal(m, b)
}
func (m *AddPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPostRequest.Marshal(b, m, deterministic)
}
func (m *AddPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPostRequest.Merge(m, src)
}
func (m *AddPostRequest) XXX_Size() int {
	return xxx_messageInfo_AddPostRequest.Size(m)
}
func (m *AddPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddPostRequest proto.InternalMessageInfo

func (m *AddPostRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *AddPostRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type AddPostResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPostResponse) Reset()         { *m = AddPostResponse{} }
func (m *AddPostResponse) String() string { return proto.CompactTextString(m) }
func (*AddPostResponse) ProtoMessage()    {}
func (*AddPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12e57f654310df34, []int{1}
}

func (m *AddPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPostResponse.Unmarshal(m, b)
}
func (m *AddPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPostResponse.Marshal(b, m, deterministic)
}
func (m *AddPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPostResponse.Merge(m, src)
}
func (m *AddPostResponse) XXX_Size() int {
	return xxx_messageInfo_AddPostResponse.Size(m)
}
func (m *AddPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddPostResponse proto.InternalMessageInfo

func (m *AddPostResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AddPostResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*AddPostRequest)(nil), "go.blogging.graphql.AddPostRequest")
	proto.RegisterType((*AddPostResponse)(nil), "go.blogging.graphql.AddPostResponse")
}

func init() {
	proto.RegisterFile("servers/blogging/proto/blogging.proto", fileDescriptor_12e57f654310df34)
}

var fileDescriptor_12e57f654310df34 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0x2a, 0xd6, 0x4f, 0xca, 0xc9, 0x4f, 0x4f, 0xcf, 0xcc, 0x4b, 0xd7, 0x2f, 0x28, 0xca,
	0x2f, 0xc9, 0x87, 0x73, 0xf5, 0xc0, 0x5c, 0x21, 0xe1, 0xf4, 0x7c, 0x3d, 0xb8, 0x50, 0x7a, 0x51,
	0x62, 0x41, 0x46, 0x61, 0x8e, 0x92, 0x13, 0x17, 0x9f, 0x63, 0x4a, 0x4a, 0x40, 0x7e, 0x71, 0x49,
	0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x7e, 0x5e, 0x49, 0x6a,
	0x5e, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0x24, 0xc6, 0xc5, 0x56, 0x5a,
	0x9c, 0x5a, 0xe4, 0x99, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0xe5, 0x29, 0x59, 0x73,
	0xf1, 0xc3, 0xcd, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xe2, 0xe3, 0x62, 0xf2, 0x74, 0x81,
	0xea, 0x67, 0xf2, 0x74, 0x41, 0x36, 0x94, 0x09, 0xc5, 0x50, 0xa3, 0x24, 0x2e, 0x0e, 0x27, 0xa8,
	0xa3, 0x84, 0xc2, 0xb8, 0xd8, 0xa1, 0x06, 0x09, 0x29, 0xeb, 0x61, 0x71, 0xad, 0x1e, 0xaa, 0x53,
	0xa5, 0x54, 0xf0, 0x2b, 0x82, 0xb8, 0x45, 0x89, 0x21, 0x89, 0x0d, 0x1c, 0x00, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbe, 0xdd, 0xb0, 0xec, 0x29, 0x01, 0x00, 0x00,
}