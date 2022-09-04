// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: protocol/tri.proto

package protocol

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

// The request message containing the user's name.
type AddTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Priority    int32  `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *AddTaskRequest) Reset() {
	*x = AddTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_tri_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskRequest) ProtoMessage() {}

func (x *AddTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_tri_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskRequest.ProtoReflect.Descriptor instead.
func (*AddTaskRequest) Descriptor() ([]byte, []int) {
	return file_protocol_tri_proto_rawDescGZIP(), []int{0}
}

func (x *AddTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddTaskRequest) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

// The response message containing the greetings
type AddTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddTaskReply) Reset() {
	*x = AddTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_tri_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskReply) ProtoMessage() {}

func (x *AddTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_tri_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskReply.ProtoReflect.Descriptor instead.
func (*AddTaskReply) Descriptor() ([]byte, []int) {
	return file_protocol_tri_proto_rawDescGZIP(), []int{1}
}

func (x *AddTaskReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The request message containing the user's name.
type SetDoneTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *SetDoneTaskRequest) Reset() {
	*x = SetDoneTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_tri_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDoneTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDoneTaskRequest) ProtoMessage() {}

func (x *SetDoneTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_tri_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDoneTaskRequest.ProtoReflect.Descriptor instead.
func (*SetDoneTaskRequest) Descriptor() ([]byte, []int) {
	return file_protocol_tri_proto_rawDescGZIP(), []int{2}
}

func (x *SetDoneTaskRequest) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

// The response message containing the greetings
type SetDoneTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SetDoneTaskReply) Reset() {
	*x = SetDoneTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_tri_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDoneTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDoneTaskReply) ProtoMessage() {}

func (x *SetDoneTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_tri_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDoneTaskReply.ProtoReflect.Descriptor instead.
func (*SetDoneTaskReply) Descriptor() ([]byte, []int) {
	return file_protocol_tri_proto_rawDescGZIP(), []int{3}
}

func (x *SetDoneTaskReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top      int32 `protobuf:"varint,1,opt,name=top,proto3" json:"top,omitempty"`
	ShowDone bool  `protobuf:"varint,2,opt,name=showDone,proto3" json:"showDone,omitempty"`
	ShowAll  bool  `protobuf:"varint,3,opt,name=showAll,proto3" json:"showAll,omitempty"`
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_tri_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_tri_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_protocol_tri_proto_rawDescGZIP(), []int{4}
}

func (x *GetTaskRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *GetTaskRequest) GetShowDone() bool {
	if x != nil {
		return x.ShowDone
	}
	return false
}

func (x *GetTaskRequest) GetShowAll() bool {
	if x != nil {
		return x.ShowAll
	}
	return false
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Priority    int32  `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	Done        bool   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	Id          uint32 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_tri_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_tri_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_protocol_tri_proto_rawDescGZIP(), []int{5}
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Task) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *Task) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *GetTaskReply) Reset() {
	*x = GetTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_tri_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskReply) ProtoMessage() {}

func (x *GetTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_tri_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskReply.ProtoReflect.Descriptor instead.
func (*GetTaskReply) Descriptor() ([]byte, []int) {
	return file_protocol_tri_proto_rawDescGZIP(), []int{6}
}

func (x *GetTaskReply) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_protocol_tri_proto protoreflect.FileDescriptor

var file_protocol_tri_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x74, 0x72, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x22, 0x28, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26,
	0x0a, 0x12, 0x53, 0x65, 0x74, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x44, 0x6f, 0x6e,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x58, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x77,
	0x44, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x77,
	0x44, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x22, 0x68,
	0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x32, 0xa0, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x13, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x6f, 0x6e, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_tri_proto_rawDescOnce sync.Once
	file_protocol_tri_proto_rawDescData = file_protocol_tri_proto_rawDesc
)

func file_protocol_tri_proto_rawDescGZIP() []byte {
	file_protocol_tri_proto_rawDescOnce.Do(func() {
		file_protocol_tri_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_tri_proto_rawDescData)
	})
	return file_protocol_tri_proto_rawDescData
}

var file_protocol_tri_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protocol_tri_proto_goTypes = []interface{}{
	(*AddTaskRequest)(nil),     // 0: AddTaskRequest
	(*AddTaskReply)(nil),       // 1: AddTaskReply
	(*SetDoneTaskRequest)(nil), // 2: SetDoneTaskRequest
	(*SetDoneTaskReply)(nil),   // 3: SetDoneTaskReply
	(*GetTaskRequest)(nil),     // 4: GetTaskRequest
	(*Task)(nil),               // 5: Task
	(*GetTaskReply)(nil),       // 6: GetTaskReply
}
var file_protocol_tri_proto_depIdxs = []int32{
	5, // 0: GetTaskReply.tasks:type_name -> Task
	0, // 1: TaskHandler.AddTask:input_type -> AddTaskRequest
	2, // 2: TaskHandler.SetDoneTask:input_type -> SetDoneTaskRequest
	4, // 3: TaskHandler.GetTask:input_type -> GetTaskRequest
	1, // 4: TaskHandler.AddTask:output_type -> AddTaskReply
	3, // 5: TaskHandler.SetDoneTask:output_type -> SetDoneTaskReply
	6, // 6: TaskHandler.GetTask:output_type -> GetTaskReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protocol_tri_proto_init() }
func file_protocol_tri_proto_init() {
	if File_protocol_tri_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_tri_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskRequest); i {
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
		file_protocol_tri_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskReply); i {
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
		file_protocol_tri_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDoneTaskRequest); i {
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
		file_protocol_tri_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDoneTaskReply); i {
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
		file_protocol_tri_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskRequest); i {
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
		file_protocol_tri_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_protocol_tri_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskReply); i {
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
			RawDescriptor: file_protocol_tri_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_tri_proto_goTypes,
		DependencyIndexes: file_protocol_tri_proto_depIdxs,
		MessageInfos:      file_protocol_tri_proto_msgTypes,
	}.Build()
	File_protocol_tri_proto = out.File
	file_protocol_tri_proto_rawDesc = nil
	file_protocol_tri_proto_goTypes = nil
	file_protocol_tri_proto_depIdxs = nil
}
