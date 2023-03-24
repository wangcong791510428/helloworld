// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.13.0
// source: student/v1/student.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type StudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  uint32 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *StudentRequest) Reset() {
	*x = StudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentRequest) ProtoMessage() {}

func (x *StudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_student_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentRequest.ProtoReflect.Descriptor instead.
func (*StudentRequest) Descriptor() ([]byte, []int) {
	return file_student_v1_student_proto_rawDescGZIP(), []int{0}
}

func (x *StudentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StudentRequest) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type StudentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StudentReply) Reset() {
	*x = StudentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentReply) ProtoMessage() {}

func (x *StudentReply) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_student_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentReply.ProtoReflect.Descriptor instead.
func (*StudentReply) Descriptor() ([]byte, []int) {
	return file_student_v1_student_proto_rawDescGZIP(), []int{1}
}

func (x *StudentReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_student_v1_student_proto protoreflect.FileDescriptor

var file_student_v1_student_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x22, 0x28, 0x0a, 0x0c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x87, 0x01, 0x0a, 0x07, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x7c, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x15, 0x2f, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x7b, 0x61, 0x67,
	0x65, 0x7d, 0x5a, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x3a, 0x01, 0x2a, 0x42, 0x30, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1c, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_student_v1_student_proto_rawDescOnce sync.Once
	file_student_v1_student_proto_rawDescData = file_student_v1_student_proto_rawDesc
)

func file_student_v1_student_proto_rawDescGZIP() []byte {
	file_student_v1_student_proto_rawDescOnce.Do(func() {
		file_student_v1_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_student_v1_student_proto_rawDescData)
	})
	return file_student_v1_student_proto_rawDescData
}

var file_student_v1_student_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_student_v1_student_proto_goTypes = []interface{}{
	(*StudentRequest)(nil), // 0: api.student.v1.StudentRequest
	(*StudentReply)(nil),   // 1: api.student.v1.StudentReply
}
var file_student_v1_student_proto_depIdxs = []int32{
	0, // 0: api.student.v1.Student.CallStudent:input_type -> api.student.v1.StudentRequest
	1, // 1: api.student.v1.Student.CallStudent:output_type -> api.student.v1.StudentReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_student_v1_student_proto_init() }
func file_student_v1_student_proto_init() {
	if File_student_v1_student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_student_v1_student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentRequest); i {
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
		file_student_v1_student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentReply); i {
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
			RawDescriptor: file_student_v1_student_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_student_v1_student_proto_goTypes,
		DependencyIndexes: file_student_v1_student_proto_depIdxs,
		MessageInfos:      file_student_v1_student_proto_msgTypes,
	}.Build()
	File_student_v1_student_proto = out.File
	file_student_v1_student_proto_rawDesc = nil
	file_student_v1_student_proto_goTypes = nil
	file_student_v1_student_proto_depIdxs = nil
}
