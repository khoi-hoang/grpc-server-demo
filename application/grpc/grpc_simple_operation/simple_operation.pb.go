// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.1
// source: simple_operation.proto

package grpc_simple_operation

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

type SimpleOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First  int64 `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	Second int64 `protobuf:"varint,2,opt,name=second,proto3" json:"second,omitempty"`
}

func (x *SimpleOperationRequest) Reset() {
	*x = SimpleOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_operation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleOperationRequest) ProtoMessage() {}

func (x *SimpleOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simple_operation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleOperationRequest.ProtoReflect.Descriptor instead.
func (*SimpleOperationRequest) Descriptor() ([]byte, []int) {
	return file_simple_operation_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleOperationRequest) GetFirst() int64 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *SimpleOperationRequest) GetSecond() int64 {
	if x != nil {
		return x.Second
	}
	return 0
}

type SimpleOperationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SimpleOperationResponse) Reset() {
	*x = SimpleOperationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_operation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleOperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleOperationResponse) ProtoMessage() {}

func (x *SimpleOperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_simple_operation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleOperationResponse.ProtoReflect.Descriptor instead.
func (*SimpleOperationResponse) Descriptor() ([]byte, []int) {
	return file_simple_operation_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleOperationResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_simple_operation_proto protoreflect.FileDescriptor

var file_simple_operation_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x16, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x22, 0x31, 0x0a, 0x17, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0x4b, 0x0a, 0x0f, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x17, 0x2e,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_simple_operation_proto_rawDescOnce sync.Once
	file_simple_operation_proto_rawDescData = file_simple_operation_proto_rawDesc
)

func file_simple_operation_proto_rawDescGZIP() []byte {
	file_simple_operation_proto_rawDescOnce.Do(func() {
		file_simple_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_simple_operation_proto_rawDescData)
	})
	return file_simple_operation_proto_rawDescData
}

var file_simple_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_simple_operation_proto_goTypes = []interface{}{
	(*SimpleOperationRequest)(nil),  // 0: SimpleOperationRequest
	(*SimpleOperationResponse)(nil), // 1: SimpleOperationResponse
}
var file_simple_operation_proto_depIdxs = []int32{
	0, // 0: SimpleOperation.Sum:input_type -> SimpleOperationRequest
	1, // 1: SimpleOperation.Sum:output_type -> SimpleOperationResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_simple_operation_proto_init() }
func file_simple_operation_proto_init() {
	if File_simple_operation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_simple_operation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleOperationRequest); i {
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
		file_simple_operation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleOperationResponse); i {
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
			RawDescriptor: file_simple_operation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_simple_operation_proto_goTypes,
		DependencyIndexes: file_simple_operation_proto_depIdxs,
		MessageInfos:      file_simple_operation_proto_msgTypes,
	}.Build()
	File_simple_operation_proto = out.File
	file_simple_operation_proto_rawDesc = nil
	file_simple_operation_proto_goTypes = nil
	file_simple_operation_proto_depIdxs = nil
}