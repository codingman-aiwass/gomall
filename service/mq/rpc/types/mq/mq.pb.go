// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: mq/rpc/mq.proto

package mq

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

type SendDelayMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic        string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`                // 消息主题
	Payload      []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`            // 消息内容
	DelaySeconds int64  `protobuf:"varint,3,opt,name=delaySeconds,proto3" json:"delaySeconds,omitempty"` // 延时秒数
}

func (x *SendDelayMessageReq) Reset() {
	*x = SendDelayMessageReq{}
	mi := &file_mq_rpc_mq_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendDelayMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDelayMessageReq) ProtoMessage() {}

func (x *SendDelayMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_mq_rpc_mq_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDelayMessageReq.ProtoReflect.Descriptor instead.
func (*SendDelayMessageReq) Descriptor() ([]byte, []int) {
	return file_mq_rpc_mq_proto_rawDescGZIP(), []int{0}
}

func (x *SendDelayMessageReq) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *SendDelayMessageReq) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SendDelayMessageReq) GetDelaySeconds() int64 {
	if x != nil {
		return x.DelaySeconds
	}
	return 0
}

type SendDelayMessageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=messageId,proto3" json:"messageId,omitempty"` // 消息ID
}

func (x *SendDelayMessageResp) Reset() {
	*x = SendDelayMessageResp{}
	mi := &file_mq_rpc_mq_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendDelayMessageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDelayMessageResp) ProtoMessage() {}

func (x *SendDelayMessageResp) ProtoReflect() protoreflect.Message {
	mi := &file_mq_rpc_mq_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDelayMessageResp.ProtoReflect.Descriptor instead.
func (*SendDelayMessageResp) Descriptor() ([]byte, []int) {
	return file_mq_rpc_mq_proto_rawDescGZIP(), []int{1}
}

func (x *SendDelayMessageResp) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

type SendMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic      string            `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`                                                                                                   // 消息主题
	Payload    []byte            `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`                                                                                               // 消息内容
	Properties map[string]string `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 消息属性（可选，用于附加键值对）
}

func (x *SendMessageReq) Reset() {
	*x = SendMessageReq{}
	mi := &file_mq_rpc_mq_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReq) ProtoMessage() {}

func (x *SendMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_mq_rpc_mq_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReq.ProtoReflect.Descriptor instead.
func (*SendMessageReq) Descriptor() ([]byte, []int) {
	return file_mq_rpc_mq_proto_rawDescGZIP(), []int{2}
}

func (x *SendMessageReq) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *SendMessageReq) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SendMessageReq) GetProperties() map[string]string {
	if x != nil {
		return x.Properties
	}
	return nil
}

type SendMessageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=messageId,proto3" json:"messageId,omitempty"` // 消息ID
}

func (x *SendMessageResp) Reset() {
	*x = SendMessageResp{}
	mi := &file_mq_rpc_mq_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResp) ProtoMessage() {}

func (x *SendMessageResp) ProtoReflect() protoreflect.Message {
	mi := &file_mq_rpc_mq_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResp.ProtoReflect.Descriptor instead.
func (*SendMessageResp) Descriptor() ([]byte, []int) {
	return file_mq_rpc_mq_proto_rawDescGZIP(), []int{3}
}

func (x *SendMessageResp) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

var File_mq_rpc_mq_proto protoreflect.FileDescriptor

var file_mq_rpc_mq_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x71, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x6d, 0x71, 0x22, 0x69, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x6c,
	0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x64, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x22, 0x34, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0xc3, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x6d, 0x71, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x3d, 0x0a,
	0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2f, 0x0a, 0x0f,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x32, 0x83, 0x01,
	0x0a, 0x02, 0x4d, 0x71, 0x12, 0x45, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x61,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x71, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x18, 0x2e, 0x6d, 0x71, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x61, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x0b, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x2e, 0x6d, 0x71, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x6d, 0x71, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x6d, 0x71, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mq_rpc_mq_proto_rawDescOnce sync.Once
	file_mq_rpc_mq_proto_rawDescData = file_mq_rpc_mq_proto_rawDesc
)

func file_mq_rpc_mq_proto_rawDescGZIP() []byte {
	file_mq_rpc_mq_proto_rawDescOnce.Do(func() {
		file_mq_rpc_mq_proto_rawDescData = protoimpl.X.CompressGZIP(file_mq_rpc_mq_proto_rawDescData)
	})
	return file_mq_rpc_mq_proto_rawDescData
}

var file_mq_rpc_mq_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mq_rpc_mq_proto_goTypes = []any{
	(*SendDelayMessageReq)(nil),  // 0: mq.SendDelayMessageReq
	(*SendDelayMessageResp)(nil), // 1: mq.SendDelayMessageResp
	(*SendMessageReq)(nil),       // 2: mq.SendMessageReq
	(*SendMessageResp)(nil),      // 3: mq.SendMessageResp
	nil,                          // 4: mq.SendMessageReq.PropertiesEntry
}
var file_mq_rpc_mq_proto_depIdxs = []int32{
	4, // 0: mq.SendMessageReq.properties:type_name -> mq.SendMessageReq.PropertiesEntry
	0, // 1: mq.Mq.SendDelayMessage:input_type -> mq.SendDelayMessageReq
	2, // 2: mq.Mq.SendMessage:input_type -> mq.SendMessageReq
	1, // 3: mq.Mq.SendDelayMessage:output_type -> mq.SendDelayMessageResp
	3, // 4: mq.Mq.SendMessage:output_type -> mq.SendMessageResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mq_rpc_mq_proto_init() }
func file_mq_rpc_mq_proto_init() {
	if File_mq_rpc_mq_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mq_rpc_mq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mq_rpc_mq_proto_goTypes,
		DependencyIndexes: file_mq_rpc_mq_proto_depIdxs,
		MessageInfos:      file_mq_rpc_mq_proto_msgTypes,
	}.Build()
	File_mq_rpc_mq_proto = out.File
	file_mq_rpc_mq_proto_rawDesc = nil
	file_mq_rpc_mq_proto_goTypes = nil
	file_mq_rpc_mq_proto_depIdxs = nil
}
