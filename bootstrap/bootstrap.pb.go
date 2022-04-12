// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: proto/bootstrap.proto

package bootstrap

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

type Error int32

const (
	Error_AUTHENTICATION Error = 0
)

// Enum value maps for Error.
var (
	Error_name = map[int32]string{
		0: "AUTHENTICATION",
	}
	Error_value = map[string]int32{
		"AUTHENTICATION": 0,
	}
)

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}

func (x Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_bootstrap_proto_enumTypes[0].Descriptor()
}

func (Error) Type() protoreflect.EnumType {
	return &file_proto_bootstrap_proto_enumTypes[0]
}

func (x Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error.Descriptor instead.
func (Error) EnumDescriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{0}
}

type CommandTarget int32

const (
	CommandTarget_FILE CommandTarget = 0
	CommandTarget_DIR  CommandTarget = 1
)

// Enum value maps for CommandTarget.
var (
	CommandTarget_name = map[int32]string{
		0: "FILE",
		1: "DIR",
	}
	CommandTarget_value = map[string]int32{
		"FILE": 0,
		"DIR":  1,
	}
)

func (x CommandTarget) Enum() *CommandTarget {
	p := new(CommandTarget)
	*p = x
	return p
}

func (x CommandTarget) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandTarget) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_bootstrap_proto_enumTypes[1].Descriptor()
}

func (CommandTarget) Type() protoreflect.EnumType {
	return &file_proto_bootstrap_proto_enumTypes[1]
}

func (x CommandTarget) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandTarget.Descriptor instead.
func (CommandTarget) EnumDescriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{1}
}

type LoginCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginCommand) Reset() {
	*x = LoginCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bootstrap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginCommand) ProtoMessage() {}

func (x *LoginCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bootstrap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginCommand.ProtoReflect.Descriptor instead.
func (*LoginCommand) Descriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{0}
}

func (x *LoginCommand) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginCommand) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session   string        `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Bootstrap *BootstrapDTO `protobuf:"bytes,2,opt,name=bootstrap,proto3" json:"bootstrap,omitempty"`
	Error     *ErrorDTO     `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *LoginDocument) Reset() {
	*x = LoginDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bootstrap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginDocument) ProtoMessage() {}

func (x *LoginDocument) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bootstrap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginDocument.ProtoReflect.Descriptor instead.
func (*LoginDocument) Descriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{1}
}

func (x *LoginDocument) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

func (x *LoginDocument) GetBootstrap() *BootstrapDTO {
	if x != nil {
		return x.Bootstrap
	}
	return nil
}

func (x *LoginDocument) GetError() *ErrorDTO {
	if x != nil {
		return x.Error
	}
	return nil
}

type SessionCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session string `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *SessionCommand) Reset() {
	*x = SessionCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bootstrap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionCommand) ProtoMessage() {}

func (x *SessionCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bootstrap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionCommand.ProtoReflect.Descriptor instead.
func (*SessionCommand) Descriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{2}
}

func (x *SessionCommand) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

type SessionDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session   string        `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Bootstrap *BootstrapDTO `protobuf:"bytes,2,opt,name=bootstrap,proto3" json:"bootstrap,omitempty"`
	Error     *ErrorDTO     `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SessionDocument) Reset() {
	*x = SessionDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bootstrap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionDocument) ProtoMessage() {}

func (x *SessionDocument) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bootstrap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionDocument.ProtoReflect.Descriptor instead.
func (*SessionDocument) Descriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{3}
}

func (x *SessionDocument) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

func (x *SessionDocument) GetBootstrap() *BootstrapDTO {
	if x != nil {
		return x.Bootstrap
	}
	return nil
}

func (x *SessionDocument) GetError() *ErrorDTO {
	if x != nil {
		return x.Error
	}
	return nil
}

type HelloQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HelloQuery) Reset() {
	*x = HelloQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bootstrap_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloQuery) ProtoMessage() {}

func (x *HelloQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bootstrap_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloQuery.ProtoReflect.Descriptor instead.
func (*HelloQuery) Descriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{4}
}

type HelloDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session   string        `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Bootstrap *BootstrapDTO `protobuf:"bytes,2,opt,name=bootstrap,proto3" json:"bootstrap,omitempty"`
	Error     *ErrorDTO     `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *HelloDocument) Reset() {
	*x = HelloDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bootstrap_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloDocument) ProtoMessage() {}

func (x *HelloDocument) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bootstrap_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloDocument.ProtoReflect.Descriptor instead.
func (*HelloDocument) Descriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{5}
}

func (x *HelloDocument) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

func (x *HelloDocument) GetBootstrap() *BootstrapDTO {
	if x != nil {
		return x.Bootstrap
	}
	return nil
}

func (x *HelloDocument) GetError() *ErrorDTO {
	if x != nil {
		return x.Error
	}
	return nil
}

type ErrorDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   Error  `protobuf:"varint,1,opt,name=error,proto3,enum=Error" json:"error,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorDTO) Reset() {
	*x = ErrorDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bootstrap_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDTO) ProtoMessage() {}

func (x *ErrorDTO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bootstrap_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDTO.ProtoReflect.Descriptor instead.
func (*ErrorDTO) Descriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{6}
}

func (x *ErrorDTO) GetError() Error {
	if x != nil {
		return x.Error
	}
	return Error_AUTHENTICATION
}

func (x *ErrorDTO) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BootstrapDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Modules []*ModuleDTO `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *BootstrapDTO) Reset() {
	*x = BootstrapDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bootstrap_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapDTO) ProtoMessage() {}

func (x *BootstrapDTO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bootstrap_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapDTO.ProtoReflect.Descriptor instead.
func (*BootstrapDTO) Descriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{7}
}

func (x *BootstrapDTO) GetModules() []*ModuleDTO {
	if x != nil {
		return x.Modules
	}
	return nil
}

type ModuleDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []*EndpointDTO `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Commands  []*CommandDTO  `protobuf:"bytes,2,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *ModuleDTO) Reset() {
	*x = ModuleDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bootstrap_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleDTO) ProtoMessage() {}

func (x *ModuleDTO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bootstrap_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleDTO.ProtoReflect.Descriptor instead.
func (*ModuleDTO) Descriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{8}
}

func (x *ModuleDTO) GetEndpoints() []*EndpointDTO {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *ModuleDTO) GetCommands() []*CommandDTO {
	if x != nil {
		return x.Commands
	}
	return nil
}

type EndpointDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method  string      `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path    string      `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Handler *HandlerDTO `protobuf:"bytes,3,opt,name=handler,proto3" json:"handler,omitempty"`
}

func (x *EndpointDTO) Reset() {
	*x = EndpointDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bootstrap_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointDTO) ProtoMessage() {}

func (x *EndpointDTO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bootstrap_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointDTO.ProtoReflect.Descriptor instead.
func (*EndpointDTO) Descriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{9}
}

func (x *EndpointDTO) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *EndpointDTO) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *EndpointDTO) GetHandler() *HandlerDTO {
	if x != nil {
		return x.Handler
	}
	return nil
}

type CommandDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Children []*CommandDTO `protobuf:"bytes,2,rep,name=children,proto3" json:"children,omitempty"`
	Handler  *HandlerDTO   `protobuf:"bytes,3,opt,name=handler,proto3" json:"handler,omitempty"`
	Target   CommandTarget `protobuf:"varint,4,opt,name=target,proto3,enum=CommandTarget" json:"target,omitempty"` // TODO is sort of decied by the HandlerDTO though?
}

func (x *CommandDTO) Reset() {
	*x = CommandDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bootstrap_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandDTO) ProtoMessage() {}

func (x *CommandDTO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bootstrap_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandDTO.ProtoReflect.Descriptor instead.
func (*CommandDTO) Descriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{10}
}

func (x *CommandDTO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommandDTO) GetChildren() []*CommandDTO {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *CommandDTO) GetHandler() *HandlerDTO {
	if x != nil {
		return x.Handler
	}
	return nil
}

func (x *CommandDTO) GetTarget() CommandTarget {
	if x != nil {
		return x.Target
	}
	return CommandTarget_FILE
}

//
// pwd - get current dir
// file_read - read current file
// file_write - write body to file
// proxy - forward the request
// db_write - store in internal db
// db_read - read from internal db
type HandlerDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HandlerDTO) Reset() {
	*x = HandlerDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bootstrap_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandlerDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandlerDTO) ProtoMessage() {}

func (x *HandlerDTO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bootstrap_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandlerDTO.ProtoReflect.Descriptor instead.
func (*HandlerDTO) Descriptor() ([]byte, []int) {
	return file_proto_bootstrap_proto_rawDescGZIP(), []int{11}
}

var File_proto_bootstrap_proto protoreflect.FileDescriptor

var file_proto_bootstrap_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x77, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x09, 0x62, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x44, 0x54, 0x4f, 0x52, 0x09, 0x62, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x54,
	0x4f, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x79, 0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x2b, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x44, 0x54, 0x4f, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x1f,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x54, 0x4f, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x0c, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x77, 0x0a,
	0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x44, 0x54, 0x4f, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x54, 0x4f, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x42, 0x0a, 0x08, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44,
	0x54, 0x4f, 0x12, 0x1c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x34, 0x0a, 0x0c, 0x42, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x44, 0x54, 0x4f, 0x12, 0x24, 0x0a, 0x07, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x44, 0x54, 0x4f, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x22, 0x60, 0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x54, 0x4f, 0x12, 0x2a, 0x0a,
	0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x54, 0x4f, 0x52, 0x09,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x44, 0x54, 0x4f, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x22, 0x60, 0x0a, 0x0b, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x54,
	0x4f, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a,
	0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x52, 0x07, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x22, 0x98, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x44, 0x54, 0x4f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64,
	0x72, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x44, 0x54, 0x4f, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x12, 0x25, 0x0a, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x52, 0x07,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22,
	0x0c, 0x0a, 0x0a, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x2a, 0x1b, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e,
	0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x2a, 0x22, 0x0a, 0x0d, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x46,
	0x49, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x49, 0x52, 0x10, 0x01, 0x32, 0x87,
	0x01, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x24, 0x0a, 0x05,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0b, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x1a, 0x0e, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x0e, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x10, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x62, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bootstrap_proto_rawDescOnce sync.Once
	file_proto_bootstrap_proto_rawDescData = file_proto_bootstrap_proto_rawDesc
)

func file_proto_bootstrap_proto_rawDescGZIP() []byte {
	file_proto_bootstrap_proto_rawDescOnce.Do(func() {
		file_proto_bootstrap_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bootstrap_proto_rawDescData)
	})
	return file_proto_bootstrap_proto_rawDescData
}

var file_proto_bootstrap_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_bootstrap_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_bootstrap_proto_goTypes = []interface{}{
	(Error)(0),              // 0: Error
	(CommandTarget)(0),      // 1: CommandTarget
	(*LoginCommand)(nil),    // 2: LoginCommand
	(*LoginDocument)(nil),   // 3: LoginDocument
	(*SessionCommand)(nil),  // 4: SessionCommand
	(*SessionDocument)(nil), // 5: SessionDocument
	(*HelloQuery)(nil),      // 6: HelloQuery
	(*HelloDocument)(nil),   // 7: HelloDocument
	(*ErrorDTO)(nil),        // 8: ErrorDTO
	(*BootstrapDTO)(nil),    // 9: BootstrapDTO
	(*ModuleDTO)(nil),       // 10: ModuleDTO
	(*EndpointDTO)(nil),     // 11: EndpointDTO
	(*CommandDTO)(nil),      // 12: CommandDTO
	(*HandlerDTO)(nil),      // 13: HandlerDTO
}
var file_proto_bootstrap_proto_depIdxs = []int32{
	9,  // 0: LoginDocument.bootstrap:type_name -> BootstrapDTO
	8,  // 1: LoginDocument.error:type_name -> ErrorDTO
	9,  // 2: SessionDocument.bootstrap:type_name -> BootstrapDTO
	8,  // 3: SessionDocument.error:type_name -> ErrorDTO
	9,  // 4: HelloDocument.bootstrap:type_name -> BootstrapDTO
	8,  // 5: HelloDocument.error:type_name -> ErrorDTO
	0,  // 6: ErrorDTO.error:type_name -> Error
	10, // 7: BootstrapDTO.modules:type_name -> ModuleDTO
	11, // 8: ModuleDTO.endpoints:type_name -> EndpointDTO
	12, // 9: ModuleDTO.commands:type_name -> CommandDTO
	13, // 10: EndpointDTO.handler:type_name -> HandlerDTO
	12, // 11: CommandDTO.children:type_name -> CommandDTO
	13, // 12: CommandDTO.handler:type_name -> HandlerDTO
	1,  // 13: CommandDTO.target:type_name -> CommandTarget
	6,  // 14: Bootstrap.hello:input_type -> HelloQuery
	2,  // 15: Bootstrap.login:input_type -> LoginCommand
	4,  // 16: Bootstrap.session:input_type -> SessionCommand
	7,  // 17: Bootstrap.hello:output_type -> HelloDocument
	3,  // 18: Bootstrap.login:output_type -> LoginDocument
	5,  // 19: Bootstrap.session:output_type -> SessionDocument
	17, // [17:20] is the sub-list for method output_type
	14, // [14:17] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_proto_bootstrap_proto_init() }
func file_proto_bootstrap_proto_init() {
	if File_proto_bootstrap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_bootstrap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginCommand); i {
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
		file_proto_bootstrap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginDocument); i {
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
		file_proto_bootstrap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionCommand); i {
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
		file_proto_bootstrap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionDocument); i {
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
		file_proto_bootstrap_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloQuery); i {
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
		file_proto_bootstrap_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloDocument); i {
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
		file_proto_bootstrap_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDTO); i {
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
		file_proto_bootstrap_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapDTO); i {
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
		file_proto_bootstrap_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleDTO); i {
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
		file_proto_bootstrap_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointDTO); i {
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
		file_proto_bootstrap_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandDTO); i {
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
		file_proto_bootstrap_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandlerDTO); i {
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
			RawDescriptor: file_proto_bootstrap_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_bootstrap_proto_goTypes,
		DependencyIndexes: file_proto_bootstrap_proto_depIdxs,
		EnumInfos:         file_proto_bootstrap_proto_enumTypes,
		MessageInfos:      file_proto_bootstrap_proto_msgTypes,
	}.Build()
	File_proto_bootstrap_proto = out.File
	file_proto_bootstrap_proto_rawDesc = nil
	file_proto_bootstrap_proto_goTypes = nil
	file_proto_bootstrap_proto_depIdxs = nil
}
