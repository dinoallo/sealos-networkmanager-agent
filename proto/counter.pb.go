// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/counter.proto

package sealos_networkmanager_agent

import (
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

type Direction int32

const (
	Direction_Unspec    Direction = 0
	Direction_V4Ingress Direction = 1
	Direction_V4Egress  Direction = 2
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "Unspec",
		1: "V4Ingress",
		2: "V4Egress",
	}
	Direction_value = map[string]int32{
		"Unspec":    0,
		"V4Ingress": 1,
		"V4Egress":  2,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_counter_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_proto_counter_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_proto_counter_proto_rawDescGZIP(), []int{0}
}

type Counter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId int64 `protobuf:"varint,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
}

func (x *Counter) Reset() {
	*x = Counter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_counter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Counter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counter) ProtoMessage() {}

func (x *Counter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_counter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Counter.ProtoReflect.Descriptor instead.
func (*Counter) Descriptor() ([]byte, []int) {
	return file_proto_counter_proto_rawDescGZIP(), []int{0}
}

func (x *Counter) GetEndpointId() int64 {
	if x != nil {
		return x.EndpointId
	}
	return 0
}

type CreateCounterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId int64     `protobuf:"varint,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	Direction  Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=proto.Direction" json:"direction,omitempty"`
}

func (x *CreateCounterRequest) Reset() {
	*x = CreateCounterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_counter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCounterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCounterRequest) ProtoMessage() {}

func (x *CreateCounterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_counter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCounterRequest.ProtoReflect.Descriptor instead.
func (*CreateCounterRequest) Descriptor() ([]byte, []int) {
	return file_proto_counter_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCounterRequest) GetEndpointId() int64 {
	if x != nil {
		return x.EndpointId
	}
	return 0
}

func (x *CreateCounterRequest) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_Unspec
}

type RemoveCounterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId int64     `protobuf:"varint,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	Direction  Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=proto.Direction" json:"direction,omitempty"`
}

func (x *RemoveCounterRequest) Reset() {
	*x = RemoveCounterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_counter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCounterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCounterRequest) ProtoMessage() {}

func (x *RemoveCounterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_counter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCounterRequest.ProtoReflect.Descriptor instead.
func (*RemoveCounterRequest) Descriptor() ([]byte, []int) {
	return file_proto_counter_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveCounterRequest) GetEndpointId() int64 {
	if x != nil {
		return x.EndpointId
	}
	return 0
}

func (x *RemoveCounterRequest) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_Unspec
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_counter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_counter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_counter_proto_rawDescGZIP(), []int{3}
}

var File_proto_counter_proto protoreflect.FileDescriptor

var file_proto_counter_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x07,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x67, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x2a, 0x34, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0a, 0x0a, 0x06, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x56, 0x34, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x56,
	0x34, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x02, 0x32, 0x8d, 0x01, 0x0a, 0x0f, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x6e, 0x6f, 0x61, 0x6c, 0x6c, 0x6f,
	0x2f, 0x73, 0x65, 0x61, 0x6c, 0x6f, 0x73, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_counter_proto_rawDescOnce sync.Once
	file_proto_counter_proto_rawDescData = file_proto_counter_proto_rawDesc
)

func file_proto_counter_proto_rawDescGZIP() []byte {
	file_proto_counter_proto_rawDescOnce.Do(func() {
		file_proto_counter_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_counter_proto_rawDescData)
	})
	return file_proto_counter_proto_rawDescData
}

var file_proto_counter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_counter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_counter_proto_goTypes = []interface{}{
	(Direction)(0),               // 0: proto.Direction
	(*Counter)(nil),              // 1: proto.Counter
	(*CreateCounterRequest)(nil), // 2: proto.CreateCounterRequest
	(*RemoveCounterRequest)(nil), // 3: proto.RemoveCounterRequest
	(*Empty)(nil),                // 4: proto.Empty
}
var file_proto_counter_proto_depIdxs = []int32{
	0, // 0: proto.CreateCounterRequest.direction:type_name -> proto.Direction
	0, // 1: proto.RemoveCounterRequest.direction:type_name -> proto.Direction
	2, // 2: proto.CountingService.CreateCounter:input_type -> proto.CreateCounterRequest
	3, // 3: proto.CountingService.RemoveCounter:input_type -> proto.RemoveCounterRequest
	4, // 4: proto.CountingService.CreateCounter:output_type -> proto.Empty
	4, // 5: proto.CountingService.RemoveCounter:output_type -> proto.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_counter_proto_init() }
func file_proto_counter_proto_init() {
	if File_proto_counter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_counter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Counter); i {
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
		file_proto_counter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCounterRequest); i {
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
		file_proto_counter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCounterRequest); i {
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
		file_proto_counter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_proto_counter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_counter_proto_goTypes,
		DependencyIndexes: file_proto_counter_proto_depIdxs,
		EnumInfos:         file_proto_counter_proto_enumTypes,
		MessageInfos:      file_proto_counter_proto_msgTypes,
	}.Build()
	File_proto_counter_proto = out.File
	file_proto_counter_proto_rawDesc = nil
	file_proto_counter_proto_goTypes = nil
	file_proto_counter_proto_depIdxs = nil
}
