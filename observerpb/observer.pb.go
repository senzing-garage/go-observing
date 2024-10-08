// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: observer.proto

package observerpb

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

type UpdateObserverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateObserverRequest) Reset() {
	*x = UpdateObserverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateObserverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateObserverRequest) ProtoMessage() {}

func (x *UpdateObserverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_observer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateObserverRequest.ProtoReflect.Descriptor instead.
func (*UpdateObserverRequest) Descriptor() ([]byte, []int) {
	return file_observer_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateObserverRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateObserverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateObserverResponse) Reset() {
	*x = UpdateObserverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateObserverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateObserverResponse) ProtoMessage() {}

func (x *UpdateObserverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_observer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateObserverResponse.ProtoReflect.Descriptor instead.
func (*UpdateObserverResponse) Descriptor() ([]byte, []int) {
	return file_observer_proto_rawDescGZIP(), []int{1}
}

var File_observer_proto protoreflect.FileDescriptor

var file_observer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x22, 0x31, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x65, 0x0a, 0x08, 0x4f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x5d, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x6e, 0x7a, 0x69, 0x6e, 0x67, 0x2e,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x62, 0x42, 0x0d, 0x4f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x6e, 0x7a, 0x69, 0x6e, 0x67,
	0x2d, 0x67, 0x61, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_observer_proto_rawDescOnce sync.Once
	file_observer_proto_rawDescData = file_observer_proto_rawDesc
)

func file_observer_proto_rawDescGZIP() []byte {
	file_observer_proto_rawDescOnce.Do(func() {
		file_observer_proto_rawDescData = protoimpl.X.CompressGZIP(file_observer_proto_rawDescData)
	})
	return file_observer_proto_rawDescData
}

var file_observer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_observer_proto_goTypes = []interface{}{
	(*UpdateObserverRequest)(nil),  // 0: observerpb.UpdateObserverRequest
	(*UpdateObserverResponse)(nil), // 1: observerpb.UpdateObserverResponse
}
var file_observer_proto_depIdxs = []int32{
	0, // 0: observerpb.Observer.UpdateObserver:input_type -> observerpb.UpdateObserverRequest
	1, // 1: observerpb.Observer.UpdateObserver:output_type -> observerpb.UpdateObserverResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_observer_proto_init() }
func file_observer_proto_init() {
	if File_observer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_observer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateObserverRequest); i {
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
		file_observer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateObserverResponse); i {
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
			RawDescriptor: file_observer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_observer_proto_goTypes,
		DependencyIndexes: file_observer_proto_depIdxs,
		MessageInfos:      file_observer_proto_msgTypes,
	}.Build()
	File_observer_proto = out.File
	file_observer_proto_rawDesc = nil
	file_observer_proto_goTypes = nil
	file_observer_proto_depIdxs = nil
}
