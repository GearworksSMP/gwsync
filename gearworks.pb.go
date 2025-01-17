// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: gearworks.proto

package gwsync

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

type SaveStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *SaveStateRequest) Reset() {
	*x = SaveStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gearworks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveStateRequest) ProtoMessage() {}

func (x *SaveStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gearworks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveStateRequest.ProtoReflect.Descriptor instead.
func (*SaveStateRequest) Descriptor() ([]byte, []int) {
	return file_gearworks_proto_rawDescGZIP(), []int{0}
}

func (x *SaveStateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type SaveStateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *SaveStateReply) Reset() {
	*x = SaveStateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gearworks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveStateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveStateReply) ProtoMessage() {}

func (x *SaveStateReply) ProtoReflect() protoreflect.Message {
	mi := &file_gearworks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveStateReply.ProtoReflect.Descriptor instead.
func (*SaveStateReply) Descriptor() ([]byte, []int) {
	return file_gearworks_proto_rawDescGZIP(), []int{1}
}

func (x *SaveStateReply) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type LoadStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *LoadStateRequest) Reset() {
	*x = LoadStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gearworks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadStateRequest) ProtoMessage() {}

func (x *LoadStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gearworks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadStateRequest.ProtoReflect.Descriptor instead.
func (*LoadStateRequest) Descriptor() ([]byte, []int) {
	return file_gearworks_proto_rawDescGZIP(), []int{2}
}

func (x *LoadStateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type LoadStateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *LoadStateReply) Reset() {
	*x = LoadStateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gearworks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadStateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadStateReply) ProtoMessage() {}

func (x *LoadStateReply) ProtoReflect() protoreflect.Message {
	mi := &file_gearworks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadStateReply.ProtoReflect.Descriptor instead.
func (*LoadStateReply) Descriptor() ([]byte, []int) {
	return file_gearworks_proto_rawDescGZIP(), []int{3}
}

func (x *LoadStateReply) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Message  string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gearworks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gearworks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_gearworks_proto_rawDescGZIP(), []int{4}
}

func (x *ChatMessage) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ChatMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ChatMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_gearworks_proto protoreflect.FileDescriptor

var file_gearworks_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x65, 0x61, 0x72, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x67, 0x65, 0x61, 0x72, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x22, 0x2b, 0x0a, 0x10,
	0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0e, 0x53, 0x61, 0x76,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x10, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x29, 0x0a, 0x0e, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5c, 0x0a, 0x0b,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xe4, 0x01, 0x0a, 0x04, 0x53,
	0x79, 0x6e, 0x63, 0x12, 0x4b, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x65, 0x61, 0x72, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x65, 0x61, 0x72, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x0f, 0x4c, 0x6f, 0x61, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x65, 0x61, 0x72, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e,
	0x4c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x67, 0x65, 0x61, 0x72, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x4c, 0x6f, 0x61,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x0a, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x65,
	0x61, 0x72, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x65, 0x61, 0x72, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x4d, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x61, 0x72, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x73, 0x6d, 0x70, 0x2e, 0x67, 0x65, 0x61, 0x72, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e,
	0x67, 0x77, 0x73, 0x79, 0x6e, 0x63, 0x42, 0x06, 0x47, 0x57, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x01,
	0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x61,
	0x72, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x73, 0x6d, 0x70, 0x2f, 0x67, 0x77, 0x73, 0x79, 0x6e, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gearworks_proto_rawDescOnce sync.Once
	file_gearworks_proto_rawDescData = file_gearworks_proto_rawDesc
)

func file_gearworks_proto_rawDescGZIP() []byte {
	file_gearworks_proto_rawDescOnce.Do(func() {
		file_gearworks_proto_rawDescData = protoimpl.X.CompressGZIP(file_gearworks_proto_rawDescData)
	})
	return file_gearworks_proto_rawDescData
}

var file_gearworks_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gearworks_proto_goTypes = []any{
	(*SaveStateRequest)(nil), // 0: gearworks.SaveStateRequest
	(*SaveStateReply)(nil),   // 1: gearworks.SaveStateReply
	(*LoadStateRequest)(nil), // 2: gearworks.LoadStateRequest
	(*LoadStateReply)(nil),   // 3: gearworks.LoadStateReply
	(*ChatMessage)(nil),      // 4: gearworks.ChatMessage
}
var file_gearworks_proto_depIdxs = []int32{
	0, // 0: gearworks.Sync.SavePlayerState:input_type -> gearworks.SaveStateRequest
	2, // 1: gearworks.Sync.LoadPlayerState:input_type -> gearworks.LoadStateRequest
	4, // 2: gearworks.Sync.HandleChat:input_type -> gearworks.ChatMessage
	1, // 3: gearworks.Sync.SavePlayerState:output_type -> gearworks.SaveStateReply
	3, // 4: gearworks.Sync.LoadPlayerState:output_type -> gearworks.LoadStateReply
	4, // 5: gearworks.Sync.HandleChat:output_type -> gearworks.ChatMessage
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gearworks_proto_init() }
func file_gearworks_proto_init() {
	if File_gearworks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gearworks_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SaveStateRequest); i {
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
		file_gearworks_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SaveStateReply); i {
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
		file_gearworks_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LoadStateRequest); i {
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
		file_gearworks_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LoadStateReply); i {
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
		file_gearworks_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ChatMessage); i {
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
			RawDescriptor: file_gearworks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gearworks_proto_goTypes,
		DependencyIndexes: file_gearworks_proto_depIdxs,
		MessageInfos:      file_gearworks_proto_msgTypes,
	}.Build()
	File_gearworks_proto = out.File
	file_gearworks_proto_rawDesc = nil
	file_gearworks_proto_goTypes = nil
	file_gearworks_proto_depIdxs = nil
}
