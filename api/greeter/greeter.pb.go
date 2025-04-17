// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.4
// source: api/greeter/greeter.proto

package greeter

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorReason int32

const (
	// 为某个枚举单独设置错误码
	ErrorReason_USER_NOT_FOUND  ErrorReason = 0
	ErrorReason_CONTENT_MISSING ErrorReason = 1
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "USER_NOT_FOUND",
		1: "CONTENT_MISSING",
	}
	ErrorReason_value = map[string]int32{
		"USER_NOT_FOUND":  0,
		"CONTENT_MISSING": 1,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_greeter_greeter_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_api_greeter_greeter_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_greeter_greeter_proto_rawDescGZIP(), []int{0}
}

// The request message containing the user's name.
type PingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	mi := &file_api_greeter_greeter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeter_greeter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_api_greeter_greeter_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type PingReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PingReply) Reset() {
	*x = PingReply{}
	mi := &file_api_greeter_greeter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReply) ProtoMessage() {}

func (x *PingReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_greeter_greeter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReply.ProtoReflect.Descriptor instead.
func (*PingReply) Descriptor() ([]byte, []int) {
	return file_api_greeter_greeter_proto_rawDescGZIP(), []int{1}
}

func (x *PingReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_greeter_greeter_proto protoreflect.FileDescriptor

const file_api_greeter_greeter_proto_rawDesc = "" +
	"\n" +
	"\x19api/greeter/greeter.proto\x12\agreeter\x1a\x1cgoogle/api/annotations.proto\x1a\x13errors/errors.proto\"!\n" +
	"\vPingRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"%\n" +
	"\tPingReply\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage*H\n" +
	"\vErrorReason\x12\x18\n" +
	"\x0eUSER_NOT_FOUND\x10\x00\x1a\x04\xa8E\x94\x03\x12\x19\n" +
	"\x0fCONTENT_MISSING\x10\x01\x1a\x04\xa8E\x90\x03\x1a\x04\xa0E\xf4\x032Q\n" +
	"\aGreeter\x12F\n" +
	"\x04Ping\x12\x14.greeter.PingRequest\x1a\x12.greeter.PingReply\"\x14\x82\xd3\xe4\x93\x02\x0e\x12\f/ping/{name}B`\n" +
	"\x16dev.matina.api.greeterB\fGreeterProtoP\x01Z6github.com/go-matina/kratos-layout/api/greeter;greeterb\x06proto3"

var (
	file_api_greeter_greeter_proto_rawDescOnce sync.Once
	file_api_greeter_greeter_proto_rawDescData []byte
)

func file_api_greeter_greeter_proto_rawDescGZIP() []byte {
	file_api_greeter_greeter_proto_rawDescOnce.Do(func() {
		file_api_greeter_greeter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_greeter_greeter_proto_rawDesc), len(file_api_greeter_greeter_proto_rawDesc)))
	})
	return file_api_greeter_greeter_proto_rawDescData
}

var file_api_greeter_greeter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_greeter_greeter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_greeter_greeter_proto_goTypes = []any{
	(ErrorReason)(0),    // 0: greeter.ErrorReason
	(*PingRequest)(nil), // 1: greeter.PingRequest
	(*PingReply)(nil),   // 2: greeter.PingReply
}
var file_api_greeter_greeter_proto_depIdxs = []int32{
	1, // 0: greeter.Greeter.Ping:input_type -> greeter.PingRequest
	2, // 1: greeter.Greeter.Ping:output_type -> greeter.PingReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_greeter_greeter_proto_init() }
func file_api_greeter_greeter_proto_init() {
	if File_api_greeter_greeter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_greeter_greeter_proto_rawDesc), len(file_api_greeter_greeter_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_greeter_greeter_proto_goTypes,
		DependencyIndexes: file_api_greeter_greeter_proto_depIdxs,
		EnumInfos:         file_api_greeter_greeter_proto_enumTypes,
		MessageInfos:      file_api_greeter_greeter_proto_msgTypes,
	}.Build()
	File_api_greeter_greeter_proto = out.File
	file_api_greeter_greeter_proto_goTypes = nil
	file_api_greeter_greeter_proto_depIdxs = nil
}
