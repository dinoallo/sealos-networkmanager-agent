// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
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

type NewCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId uint32 `protobuf:"varint,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
}

func (x *NewCounter) Reset() {
	*x = NewCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_counter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCounter) ProtoMessage() {}

func (x *NewCounter) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use NewCounter.ProtoReflect.Descriptor instead.
func (*NewCounter) Descriptor() ([]byte, []int) {
	return file_proto_counter_proto_rawDescGZIP(), []int{0}
}

func (x *NewCounter) GetEndpointId() uint32 {
	if x != nil {
		return x.EndpointId
	}
	return 0
}

type Counter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId uint32 `protobuf:"varint,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
}

func (x *Counter) Reset() {
	*x = Counter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_counter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Counter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counter) ProtoMessage() {}

func (x *Counter) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Counter.ProtoReflect.Descriptor instead.
func (*Counter) Descriptor() ([]byte, []int) {
	return file_proto_counter_proto_rawDescGZIP(), []int{1}
}

func (x *Counter) GetEndpointId() uint32 {
	if x != nil {
		return x.EndpointId
	}
	return 0
}

type Dump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity uint32 `protobuf:"varint,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Bytes    uint64 `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Type     uint32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Dump) Reset() {
	*x = Dump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_counter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dump) ProtoMessage() {}

func (x *Dump) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Dump.ProtoReflect.Descriptor instead.
func (*Dump) Descriptor() ([]byte, []int) {
	return file_proto_counter_proto_rawDescGZIP(), []int{2}
}

func (x *Dump) GetIdentity() uint32 {
	if x != nil {
		return x.Identity
	}
	return 0
}

func (x *Dump) GetBytes() uint64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *Dump) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type CounterDumps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId uint32  `protobuf:"varint,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	Dumps      []*Dump `protobuf:"bytes,2,rep,name=dumps,proto3" json:"dumps,omitempty"`
}

func (x *CounterDumps) Reset() {
	*x = CounterDumps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_counter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CounterDumps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterDumps) ProtoMessage() {}

func (x *CounterDumps) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CounterDumps.ProtoReflect.Descriptor instead.
func (*CounterDumps) Descriptor() ([]byte, []int) {
	return file_proto_counter_proto_rawDescGZIP(), []int{3}
}

func (x *CounterDumps) GetEndpointId() uint32 {
	if x != nil {
		return x.EndpointId
	}
	return 0
}

func (x *CounterDumps) GetDumps() []*Dump {
	if x != nil {
		return x.Dumps
	}
	return nil
}

var File_proto_counter_proto protoreflect.FileDescriptor

var file_proto_counter_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x0a,
	0x4e, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x07, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x04, 0x44, 0x75, 0x6d, 0x70, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x44, 0x75, 0x6d, 0x70, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x05, 0x64, 0x75, 0x6d, 0x70, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75,
	0x6d, 0x70, 0x52, 0x05, 0x64, 0x75, 0x6d, 0x70, 0x73, 0x32, 0x7d, 0x0a, 0x0f, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b, 0x44, 0x75, 0x6d, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x44, 0x75, 0x6d, 0x70, 0x73, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x6e, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x2f,
	0x73, 0x65, 0x61, 0x6c, 0x6f, 0x73, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_proto_counter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_counter_proto_goTypes = []interface{}{
	(*NewCounter)(nil),   // 0: proto.NewCounter
	(*Counter)(nil),      // 1: proto.Counter
	(*Dump)(nil),         // 2: proto.Dump
	(*CounterDumps)(nil), // 3: proto.CounterDumps
}
var file_proto_counter_proto_depIdxs = []int32{
	2, // 0: proto.CounterDumps.dumps:type_name -> proto.Dump
	0, // 1: proto.CountingService.CreateCounter:input_type -> proto.NewCounter
	1, // 2: proto.CountingService.DumpCounter:input_type -> proto.Counter
	1, // 3: proto.CountingService.CreateCounter:output_type -> proto.Counter
	3, // 4: proto.CountingService.DumpCounter:output_type -> proto.CounterDumps
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_counter_proto_init() }
func file_proto_counter_proto_init() {
	if File_proto_counter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_counter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCounter); i {
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
		file_proto_counter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dump); i {
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
			switch v := v.(*CounterDumps); i {
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
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_counter_proto_goTypes,
		DependencyIndexes: file_proto_counter_proto_depIdxs,
		MessageInfos:      file_proto_counter_proto_msgTypes,
	}.Build()
	File_proto_counter_proto = out.File
	file_proto_counter_proto_rawDesc = nil
	file_proto_counter_proto_goTypes = nil
	file_proto_counter_proto_depIdxs = nil
}
