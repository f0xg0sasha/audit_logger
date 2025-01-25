// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.28.3
// source: audit.proto

package audit

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type LogRequest_Actions int32

const (
	LogRequest_REGISTER LogRequest_Actions = 0
	LogRequest_LOGIN    LogRequest_Actions = 1
	LogRequest_CREATE   LogRequest_Actions = 2
	LogRequest_UPDATE   LogRequest_Actions = 3
	LogRequest_GET      LogRequest_Actions = 4
	LogRequest_DELETE   LogRequest_Actions = 5
)

// Enum value maps for LogRequest_Actions.
var (
	LogRequest_Actions_name = map[int32]string{
		0: "REGISTER",
		1: "LOGIN",
		2: "CREATE",
		3: "UPDATE",
		4: "GET",
		5: "DELETE",
	}
	LogRequest_Actions_value = map[string]int32{
		"REGISTER": 0,
		"LOGIN":    1,
		"CREATE":   2,
		"UPDATE":   3,
		"GET":      4,
		"DELETE":   5,
	}
)

func (x LogRequest_Actions) Enum() *LogRequest_Actions {
	p := new(LogRequest_Actions)
	*p = x
	return p
}

func (x LogRequest_Actions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogRequest_Actions) Descriptor() protoreflect.EnumDescriptor {
	return file_audit_proto_enumTypes[0].Descriptor()
}

func (LogRequest_Actions) Type() protoreflect.EnumType {
	return &file_audit_proto_enumTypes[0]
}

func (x LogRequest_Actions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogRequest_Actions.Descriptor instead.
func (LogRequest_Actions) EnumDescriptor() ([]byte, []int) {
	return file_audit_proto_rawDescGZIP(), []int{0, 0}
}

type LogRequest_Entities int32

const (
	LogRequest_USER LogRequest_Entities = 0
	LogRequest_BOOK LogRequest_Entities = 1
)

// Enum value maps for LogRequest_Entities.
var (
	LogRequest_Entities_name = map[int32]string{
		0: "USER",
		1: "BOOK",
	}
	LogRequest_Entities_value = map[string]int32{
		"USER": 0,
		"BOOK": 1,
	}
)

func (x LogRequest_Entities) Enum() *LogRequest_Entities {
	p := new(LogRequest_Entities)
	*p = x
	return p
}

func (x LogRequest_Entities) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogRequest_Entities) Descriptor() protoreflect.EnumDescriptor {
	return file_audit_proto_enumTypes[1].Descriptor()
}

func (LogRequest_Entities) Type() protoreflect.EnumType {
	return &file_audit_proto_enumTypes[1]
}

func (x LogRequest_Entities) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogRequest_Entities.Descriptor instead.
func (LogRequest_Entities) EnumDescriptor() ([]byte, []int) {
	return file_audit_proto_rawDescGZIP(), []int{0, 1}
}

type LogRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Action        LogRequest_Actions     `protobuf:"varint,1,opt,name=action,proto3,enum=audit.LogRequest_Actions" json:"action,omitempty"`
	Entity        LogRequest_Entities    `protobuf:"varint,2,opt,name=entity,proto3,enum=audit.LogRequest_Entities" json:"entity,omitempty"`
	EntityId      int64                  `protobuf:"varint,3,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	mi := &file_audit_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audit_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_audit_proto_rawDescGZIP(), []int{0}
}

func (x *LogRequest) GetAction() LogRequest_Actions {
	if x != nil {
		return x.Action
	}
	return LogRequest_REGISTER
}

func (x *LogRequest) GetEntity() LogRequest_Entities {
	if x != nil {
		return x.Entity
	}
	return LogRequest_USER
}

func (x *LogRequest) GetEntityId() int64 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *LogRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_audit_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_audit_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_audit_proto_rawDescGZIP(), []int{1}
}

var File_audit_proto protoreflect.FileDescriptor

var file_audit_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x4f, 0x0a, 0x07, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0c, 0x0a,
	0x08, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c,
	0x4f, 0x47, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x07,
	0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x45, 0x10, 0x05, 0x22, 0x1e, 0x0a, 0x08, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f,
	0x4b, 0x10, 0x01, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x38, 0x0a, 0x0c,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x03,
	0x4c, 0x6f, 0x67, 0x12, 0x11, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_audit_proto_rawDescOnce sync.Once
	file_audit_proto_rawDescData []byte
)

func file_audit_proto_rawDescGZIP() []byte {
	file_audit_proto_rawDescOnce.Do(func() {
		file_audit_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_audit_proto_rawDesc), len(file_audit_proto_rawDesc)))
	})
	return file_audit_proto_rawDescData
}

var file_audit_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_audit_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_audit_proto_goTypes = []any{
	(LogRequest_Actions)(0),       // 0: audit.LogRequest.Actions
	(LogRequest_Entities)(0),      // 1: audit.LogRequest.Entities
	(*LogRequest)(nil),            // 2: audit.LogRequest
	(*Empty)(nil),                 // 3: audit.Empty
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_audit_proto_depIdxs = []int32{
	0, // 0: audit.LogRequest.action:type_name -> audit.LogRequest.Actions
	1, // 1: audit.LogRequest.entity:type_name -> audit.LogRequest.Entities
	4, // 2: audit.LogRequest.timestamp:type_name -> google.protobuf.Timestamp
	2, // 3: audit.AuditService.Log:input_type -> audit.LogRequest
	3, // 4: audit.AuditService.Log:output_type -> audit.Empty
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_audit_proto_init() }
func file_audit_proto_init() {
	if File_audit_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_audit_proto_rawDesc), len(file_audit_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_audit_proto_goTypes,
		DependencyIndexes: file_audit_proto_depIdxs,
		EnumInfos:         file_audit_proto_enumTypes,
		MessageInfos:      file_audit_proto_msgTypes,
	}.Build()
	File_audit_proto = out.File
	file_audit_proto_goTypes = nil
	file_audit_proto_depIdxs = nil
}
