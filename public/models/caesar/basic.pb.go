// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: protobuf/caesar/basic.proto

package caesar

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Code for response status.
type ResponseCode int32

const (
	// Default value.
	ResponseCode_UNSPECIFIED ResponseCode = 0
	// Successful.
	ResponseCode_SUCCESSFUL ResponseCode = 100
	// Failed
	ResponseCode_FAILED ResponseCode = 200
)

// Enum value maps for ResponseCode.
var (
	ResponseCode_name = map[int32]string{
		0:   "UNSPECIFIED",
		100: "SUCCESSFUL",
		200: "FAILED",
	}
	ResponseCode_value = map[string]int32{
		"UNSPECIFIED": 0,
		"SUCCESSFUL":  100,
		"FAILED":      200,
	}
)

func (x ResponseCode) Enum() *ResponseCode {
	p := new(ResponseCode)
	*p = x
	return p
}

func (x ResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_caesar_basic_proto_enumTypes[0].Descriptor()
}

func (ResponseCode) Type() protoreflect.EnumType {
	return &file_protobuf_caesar_basic_proto_enumTypes[0]
}

func (x ResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseCode.Descriptor instead.
func (ResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_caesar_basic_proto_rawDescGZIP(), []int{0}
}

// Basic response for all request.
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status code.
	Code ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=caesar.caesar.ResponseCode" json:"code,omitempty"`
	// Status message.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The time  when response created.
	Time *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	// Response body.
	Result *any.Any `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_caesar_basic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_caesar_basic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_protobuf_caesar_basic_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetCode() ResponseCode {
	if x != nil {
		return x.Code
	}
	return ResponseCode_UNSPECIFIED
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Response) GetResult() *any.Any {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_protobuf_caesar_basic_proto protoreflect.FileDescriptor

var file_protobuf_caesar_basic_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x61, 0x65, 0x73, 0x61,
	0x72, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63,
	0x61, 0x65, 0x73, 0x61, 0x72, 0x2e, 0x63, 0x61, 0x65, 0x73, 0x61, 0x72, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x61, 0x65, 0x73, 0x61, 0x72, 0x2e, 0x63, 0x61, 0x65,
	0x73, 0x61, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x3c,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x46, 0x55, 0x4c, 0x10, 0x64, 0x12,
	0x0b, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xc8, 0x01, 0x42, 0x16, 0x5a, 0x14,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x61,
	0x65, 0x73, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_caesar_basic_proto_rawDescOnce sync.Once
	file_protobuf_caesar_basic_proto_rawDescData = file_protobuf_caesar_basic_proto_rawDesc
)

func file_protobuf_caesar_basic_proto_rawDescGZIP() []byte {
	file_protobuf_caesar_basic_proto_rawDescOnce.Do(func() {
		file_protobuf_caesar_basic_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_caesar_basic_proto_rawDescData)
	})
	return file_protobuf_caesar_basic_proto_rawDescData
}

var file_protobuf_caesar_basic_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_caesar_basic_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_caesar_basic_proto_goTypes = []interface{}{
	(ResponseCode)(0),           // 0: caesar.caesar.ResponseCode
	(*Response)(nil),            // 1: caesar.caesar.Response
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*any.Any)(nil),             // 3: google.protobuf.Any
}
var file_protobuf_caesar_basic_proto_depIdxs = []int32{
	0, // 0: caesar.caesar.Response.code:type_name -> caesar.caesar.ResponseCode
	2, // 1: caesar.caesar.Response.time:type_name -> google.protobuf.Timestamp
	3, // 2: caesar.caesar.Response.result:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protobuf_caesar_basic_proto_init() }
func file_protobuf_caesar_basic_proto_init() {
	if File_protobuf_caesar_basic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_caesar_basic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_protobuf_caesar_basic_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_caesar_basic_proto_goTypes,
		DependencyIndexes: file_protobuf_caesar_basic_proto_depIdxs,
		EnumInfos:         file_protobuf_caesar_basic_proto_enumTypes,
		MessageInfos:      file_protobuf_caesar_basic_proto_msgTypes,
	}.Build()
	File_protobuf_caesar_basic_proto = out.File
	file_protobuf_caesar_basic_proto_rawDesc = nil
	file_protobuf_caesar_basic_proto_goTypes = nil
	file_protobuf_caesar_basic_proto_depIdxs = nil
}
