// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: core/services/synchronization/telem/telem_head_report.proto

package telem

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

type HeadReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainID   string `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	Latest    *Block `protobuf:"bytes,2,opt,name=latest,proto3" json:"latest,omitempty"`
	Finalized *Block `protobuf:"bytes,3,opt,name=finalized,proto3,oneof" json:"finalized,omitempty"`
}

func (x *HeadReportRequest) Reset() {
	*x = HeadReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_services_synchronization_telem_telem_head_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeadReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeadReportRequest) ProtoMessage() {}

func (x *HeadReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_services_synchronization_telem_telem_head_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeadReportRequest.ProtoReflect.Descriptor instead.
func (*HeadReportRequest) Descriptor() ([]byte, []int) {
	return file_core_services_synchronization_telem_telem_head_report_proto_rawDescGZIP(), []int{0}
}

func (x *HeadReportRequest) GetChainID() string {
	if x != nil {
		return x.ChainID
	}
	return ""
}

func (x *HeadReportRequest) GetLatest() *Block {
	if x != nil {
		return x.Latest
	}
	return nil
}

func (x *HeadReportRequest) GetFinalized() *Block {
	if x != nil {
		return x.Finalized
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Number    uint64 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Hash      string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_services_synchronization_telem_telem_head_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_core_services_synchronization_telem_telem_head_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_core_services_synchronization_telem_telem_head_report_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Block) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_core_services_synchronization_telem_telem_head_report_proto protoreflect.FileDescriptor

var file_core_services_synchronization_telem_telem_head_report_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x22, 0x92, 0x01, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x09, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x22, 0x51, 0x0a, 0x05, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x42, 0x4e, 0x5a, 0x4c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b, 0x69, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_services_synchronization_telem_telem_head_report_proto_rawDescOnce sync.Once
	file_core_services_synchronization_telem_telem_head_report_proto_rawDescData = file_core_services_synchronization_telem_telem_head_report_proto_rawDesc
)

func file_core_services_synchronization_telem_telem_head_report_proto_rawDescGZIP() []byte {
	file_core_services_synchronization_telem_telem_head_report_proto_rawDescOnce.Do(func() {
		file_core_services_synchronization_telem_telem_head_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_services_synchronization_telem_telem_head_report_proto_rawDescData)
	})
	return file_core_services_synchronization_telem_telem_head_report_proto_rawDescData
}

var file_core_services_synchronization_telem_telem_head_report_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_core_services_synchronization_telem_telem_head_report_proto_goTypes = []any{
	(*HeadReportRequest)(nil), // 0: telem.HeadReportRequest
	(*Block)(nil),             // 1: telem.Block
}
var file_core_services_synchronization_telem_telem_head_report_proto_depIdxs = []int32{
	1, // 0: telem.HeadReportRequest.latest:type_name -> telem.Block
	1, // 1: telem.HeadReportRequest.finalized:type_name -> telem.Block
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_core_services_synchronization_telem_telem_head_report_proto_init() }
func file_core_services_synchronization_telem_telem_head_report_proto_init() {
	if File_core_services_synchronization_telem_telem_head_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_services_synchronization_telem_telem_head_report_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HeadReportRequest); i {
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
		file_core_services_synchronization_telem_telem_head_report_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Block); i {
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
	file_core_services_synchronization_telem_telem_head_report_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_services_synchronization_telem_telem_head_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_services_synchronization_telem_telem_head_report_proto_goTypes,
		DependencyIndexes: file_core_services_synchronization_telem_telem_head_report_proto_depIdxs,
		MessageInfos:      file_core_services_synchronization_telem_telem_head_report_proto_msgTypes,
	}.Build()
	File_core_services_synchronization_telem_telem_head_report_proto = out.File
	file_core_services_synchronization_telem_telem_head_report_proto_rawDesc = nil
	file_core_services_synchronization_telem_telem_head_report_proto_goTypes = nil
	file_core_services_synchronization_telem_telem_head_report_proto_depIdxs = nil
}
