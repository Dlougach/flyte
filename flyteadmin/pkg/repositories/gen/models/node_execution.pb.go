// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.1
// source: node_execution.proto

package models

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

// Data about the node execution that is only referenced by FlyteAdmin and never by external callers.
type NodeExecutionInternalData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event version > 0 indicates that node execution events can now include populated IsParent and IsDynamic fields.
	EventVersion int32 `protobuf:"varint,1,opt,name=event_version,json=eventVersion,proto3" json:"event_version,omitempty"`
}

func (x *NodeExecutionInternalData) Reset() {
	*x = NodeExecutionInternalData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_execution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeExecutionInternalData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeExecutionInternalData) ProtoMessage() {}

func (x *NodeExecutionInternalData) ProtoReflect() protoreflect.Message {
	mi := &file_node_execution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeExecutionInternalData.ProtoReflect.Descriptor instead.
func (*NodeExecutionInternalData) Descriptor() ([]byte, []int) {
	return file_node_execution_proto_rawDescGZIP(), []int{0}
}

func (x *NodeExecutionInternalData) GetEventVersion() int32 {
	if x != nil {
		return x.EventVersion
	}
	return 0
}

var File_node_execution_proto protoreflect.FileDescriptor

var file_node_execution_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x40, 0x0a, 0x19, 0x4e, 0x6f, 0x64,
	0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x5a, 0x07, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_execution_proto_rawDescOnce sync.Once
	file_node_execution_proto_rawDescData = file_node_execution_proto_rawDesc
)

func file_node_execution_proto_rawDescGZIP() []byte {
	file_node_execution_proto_rawDescOnce.Do(func() {
		file_node_execution_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_execution_proto_rawDescData)
	})
	return file_node_execution_proto_rawDescData
}

var file_node_execution_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_node_execution_proto_goTypes = []interface{}{
	(*NodeExecutionInternalData)(nil), // 0: flyteadmin.models.NodeExecutionInternalData
}
var file_node_execution_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_node_execution_proto_init() }
func file_node_execution_proto_init() {
	if File_node_execution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_execution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeExecutionInternalData); i {
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
			RawDescriptor: file_node_execution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_node_execution_proto_goTypes,
		DependencyIndexes: file_node_execution_proto_depIdxs,
		MessageInfos:      file_node_execution_proto_msgTypes,
	}.Build()
	File_node_execution_proto = out.File
	file_node_execution_proto_rawDesc = nil
	file_node_execution_proto_goTypes = nil
	file_node_execution_proto_depIdxs = nil
}
