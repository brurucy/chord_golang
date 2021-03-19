// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: chord.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LookupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LookupRequest) Reset() {
	*x = LookupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupRequest) ProtoMessage() {}

func (x *LookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupRequest.ProtoReflect.Descriptor instead.
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{0}
}

func (x *LookupRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindSuccessorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindSuccessorRequest) Reset() {
	*x = FindSuccessorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindSuccessorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindSuccessorRequest) ProtoMessage() {}

func (x *FindSuccessorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindSuccessorRequest.ProtoReflect.Descriptor instead.
func (*FindSuccessorRequest) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{1}
}

func (x *FindSuccessorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindPredecessorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindPredecessorRequest) Reset() {
	*x = FindPredecessorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPredecessorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPredecessorRequest) ProtoMessage() {}

func (x *FindPredecessorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPredecessorRequest.ProtoReflect.Descriptor instead.
func (*FindPredecessorRequest) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{2}
}

func (x *FindPredecessorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ClosestNodeToRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ClosestNodeToRequest) Reset() {
	*x = ClosestNodeToRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClosestNodeToRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClosestNodeToRequest) ProtoMessage() {}

func (x *ClosestNodeToRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClosestNodeToRequest.ProtoReflect.Descriptor instead.
func (*ClosestNodeToRequest) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{3}
}

func (x *ClosestNodeToRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{4}
}

func (x *Node) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Node) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alive bool `protobuf:"varint,1,opt,name=alive,proto3" json:"alive,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chord_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_chord_proto_rawDescGZIP(), []int{5}
}

func (x *PingResponse) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

var File_chord_proto protoreflect.FileDescriptor

var file_chord_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f,
	0x0a, 0x0d, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x26, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x50,
	0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x54, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x04, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x24, 0x0a, 0x0c, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x6c, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x76,
	0x65, 0x32, 0x9c, 0x04, 0x0a, 0x05, 0x43, 0x68, 0x6f, 0x72, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x31,
	0x0a, 0x0b, 0x53, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x53, 0x75, 0x63, 0x63, 0x12, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64,
	0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x0c, 0x53, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x7a, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x54,
	0x6f, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x06, 0x4c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00,
	0x42, 0x12, 0x5a, 0x10, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chord_proto_rawDescOnce sync.Once
	file_chord_proto_rawDescData = file_chord_proto_rawDesc
)

func file_chord_proto_rawDescGZIP() []byte {
	file_chord_proto_rawDescOnce.Do(func() {
		file_chord_proto_rawDescData = protoimpl.X.CompressGZIP(file_chord_proto_rawDescData)
	})
	return file_chord_proto_rawDescData
}

var file_chord_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chord_proto_goTypes = []interface{}{
	(*LookupRequest)(nil),          // 0: pb.LookupRequest
	(*FindSuccessorRequest)(nil),   // 1: pb.FindSuccessorRequest
	(*FindPredecessorRequest)(nil), // 2: pb.FindPredecessorRequest
	(*ClosestNodeToRequest)(nil),   // 3: pb.ClosestNodeToRequest
	(*Node)(nil),                   // 4: pb.Node
	(*PingResponse)(nil),           // 5: pb.PingResponse
	(*empty.Empty)(nil),            // 6: google.protobuf.Empty
}
var file_chord_proto_depIdxs = []int32{
	6,  // 0: pb.Chord.Ping:input_type -> google.protobuf.Empty
	4,  // 1: pb.Chord.SetSucc:input_type -> pb.Node
	4,  // 2: pb.Chord.SetSuccSucc:input_type -> pb.Node
	1,  // 3: pb.Chord.FindSuccessor:input_type -> pb.FindSuccessorRequest
	2,  // 4: pb.Chord.FindPredecessor:input_type -> pb.FindPredecessorRequest
	4,  // 5: pb.Chord.Join:input_type -> pb.Node
	6,  // 6: pb.Chord.Stabilize:input_type -> google.protobuf.Empty
	6,  // 7: pb.Chord.StabilizeAll:input_type -> google.protobuf.Empty
	3,  // 8: pb.Chord.ClosestNodeTo:input_type -> pb.ClosestNodeToRequest
	0,  // 9: pb.Chord.Lookup:input_type -> pb.LookupRequest
	5,  // 10: pb.Chord.Ping:output_type -> pb.PingResponse
	6,  // 11: pb.Chord.SetSucc:output_type -> google.protobuf.Empty
	6,  // 12: pb.Chord.SetSuccSucc:output_type -> google.protobuf.Empty
	4,  // 13: pb.Chord.FindSuccessor:output_type -> pb.Node
	4,  // 14: pb.Chord.FindPredecessor:output_type -> pb.Node
	6,  // 15: pb.Chord.Join:output_type -> google.protobuf.Empty
	6,  // 16: pb.Chord.Stabilize:output_type -> google.protobuf.Empty
	6,  // 17: pb.Chord.StabilizeAll:output_type -> google.protobuf.Empty
	4,  // 18: pb.Chord.ClosestNodeTo:output_type -> pb.Node
	4,  // 19: pb.Chord.Lookup:output_type -> pb.Node
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_chord_proto_init() }
func file_chord_proto_init() {
	if File_chord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindSuccessorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPredecessorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClosestNodeToRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chord_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chord_proto_goTypes,
		DependencyIndexes: file_chord_proto_depIdxs,
		MessageInfos:      file_chord_proto_msgTypes,
	}.Build()
	File_chord_proto = out.File
	file_chord_proto_rawDesc = nil
	file_chord_proto_goTypes = nil
	file_chord_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChordClient is the client API for Chord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChordClient interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	SetSucc(ctx context.Context, in *Node, opts ...grpc.CallOption) (*empty.Empty, error)
	SetSuccSucc(ctx context.Context, in *Node, opts ...grpc.CallOption) (*empty.Empty, error)
	FindSuccessor(ctx context.Context, in *FindSuccessorRequest, opts ...grpc.CallOption) (*Node, error)
	FindPredecessor(ctx context.Context, in *FindPredecessorRequest, opts ...grpc.CallOption) (*Node, error)
	Join(ctx context.Context, in *Node, opts ...grpc.CallOption) (*empty.Empty, error)
	Stabilize(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	StabilizeAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	ClosestNodeTo(ctx context.Context, in *ClosestNodeToRequest, opts ...grpc.CallOption) (*Node, error)
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*Node, error)
}

type chordClient struct {
	cc grpc.ClientConnInterface
}

func NewChordClient(cc grpc.ClientConnInterface) ChordClient {
	return &chordClient{cc}
}

func (c *chordClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/pb.Chord/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) SetSucc(ctx context.Context, in *Node, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.Chord/SetSucc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) SetSuccSucc(ctx context.Context, in *Node, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.Chord/SetSuccSucc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) FindSuccessor(ctx context.Context, in *FindSuccessorRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/pb.Chord/FindSuccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) FindPredecessor(ctx context.Context, in *FindPredecessorRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/pb.Chord/FindPredecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) Join(ctx context.Context, in *Node, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.Chord/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) Stabilize(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.Chord/Stabilize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) StabilizeAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.Chord/StabilizeAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) ClosestNodeTo(ctx context.Context, in *ClosestNodeToRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/pb.Chord/ClosestNodeTo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/pb.Chord/Lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChordServer is the server API for Chord service.
type ChordServer interface {
	Ping(context.Context, *empty.Empty) (*PingResponse, error)
	SetSucc(context.Context, *Node) (*empty.Empty, error)
	SetSuccSucc(context.Context, *Node) (*empty.Empty, error)
	FindSuccessor(context.Context, *FindSuccessorRequest) (*Node, error)
	FindPredecessor(context.Context, *FindPredecessorRequest) (*Node, error)
	Join(context.Context, *Node) (*empty.Empty, error)
	Stabilize(context.Context, *empty.Empty) (*empty.Empty, error)
	StabilizeAll(context.Context, *empty.Empty) (*empty.Empty, error)
	ClosestNodeTo(context.Context, *ClosestNodeToRequest) (*Node, error)
	Lookup(context.Context, *LookupRequest) (*Node, error)
}

// UnimplementedChordServer can be embedded to have forward compatible implementations.
type UnimplementedChordServer struct {
}

func (*UnimplementedChordServer) Ping(context.Context, *empty.Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedChordServer) SetSucc(context.Context, *Node) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSucc not implemented")
}
func (*UnimplementedChordServer) SetSuccSucc(context.Context, *Node) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSuccSucc not implemented")
}
func (*UnimplementedChordServer) FindSuccessor(context.Context, *FindSuccessorRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessor not implemented")
}
func (*UnimplementedChordServer) FindPredecessor(context.Context, *FindPredecessorRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPredecessor not implemented")
}
func (*UnimplementedChordServer) Join(context.Context, *Node) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (*UnimplementedChordServer) Stabilize(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stabilize not implemented")
}
func (*UnimplementedChordServer) StabilizeAll(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StabilizeAll not implemented")
}
func (*UnimplementedChordServer) ClosestNodeTo(context.Context, *ClosestNodeToRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClosestNodeTo not implemented")
}
func (*UnimplementedChordServer) Lookup(context.Context, *LookupRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lookup not implemented")
}

func RegisterChordServer(s *grpc.Server, srv ChordServer) {
	s.RegisterService(&_Chord_serviceDesc, srv)
}

func _Chord_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_SetSucc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).SetSucc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/SetSucc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).SetSucc(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_SetSuccSucc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).SetSuccSucc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/SetSuccSucc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).SetSuccSucc(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_FindSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSuccessorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).FindSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/FindSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).FindSuccessor(ctx, req.(*FindSuccessorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_FindPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPredecessorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).FindPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/FindPredecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).FindPredecessor(ctx, req.(*FindPredecessorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).Join(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_Stabilize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).Stabilize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/Stabilize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).Stabilize(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_StabilizeAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).StabilizeAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/StabilizeAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).StabilizeAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_ClosestNodeTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClosestNodeToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).ClosestNodeTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/ClosestNodeTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).ClosestNodeTo(ctx, req.(*ClosestNodeToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chord/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chord_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Chord",
	HandlerType: (*ChordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Chord_Ping_Handler,
		},
		{
			MethodName: "SetSucc",
			Handler:    _Chord_SetSucc_Handler,
		},
		{
			MethodName: "SetSuccSucc",
			Handler:    _Chord_SetSuccSucc_Handler,
		},
		{
			MethodName: "FindSuccessor",
			Handler:    _Chord_FindSuccessor_Handler,
		},
		{
			MethodName: "FindPredecessor",
			Handler:    _Chord_FindPredecessor_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _Chord_Join_Handler,
		},
		{
			MethodName: "Stabilize",
			Handler:    _Chord_Stabilize_Handler,
		},
		{
			MethodName: "StabilizeAll",
			Handler:    _Chord_StabilizeAll_Handler,
		},
		{
			MethodName: "ClosestNodeTo",
			Handler:    _Chord_ClosestNodeTo_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _Chord_Lookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chord.proto",
}
