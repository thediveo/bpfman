// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: bpfd.proto

package v1

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

type ProgramType int32

const (
	ProgramType_XDP        ProgramType = 0
	ProgramType_TC_INGRESS ProgramType = 1
	ProgramType_TC_EGRESS  ProgramType = 2
)

// Enum value maps for ProgramType.
var (
	ProgramType_name = map[int32]string{
		0: "XDP",
		1: "TC_INGRESS",
		2: "TC_EGRESS",
	}
	ProgramType_value = map[string]int32{
		"XDP":        0,
		"TC_INGRESS": 1,
		"TC_EGRESS":  2,
	}
)

func (x ProgramType) Enum() *ProgramType {
	p := new(ProgramType)
	*p = x
	return p
}

func (x ProgramType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProgramType) Descriptor() protoreflect.EnumDescriptor {
	return file_bpfd_proto_enumTypes[0].Descriptor()
}

func (ProgramType) Type() protoreflect.EnumType {
	return &file_bpfd_proto_enumTypes[0]
}

func (x ProgramType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProgramType.Descriptor instead.
func (ProgramType) EnumDescriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{0}
}

type ProceedOn int32

const (
	ProceedOn_ABORTED           ProceedOn = 0
	ProceedOn_DROP              ProceedOn = 1
	ProceedOn_PASS              ProceedOn = 2
	ProceedOn_TX                ProceedOn = 3
	ProceedOn_REDIRECT          ProceedOn = 4
	ProceedOn_DISPATCHER_RETURN ProceedOn = 31
)

// Enum value maps for ProceedOn.
var (
	ProceedOn_name = map[int32]string{
		0:  "ABORTED",
		1:  "DROP",
		2:  "PASS",
		3:  "TX",
		4:  "REDIRECT",
		31: "DISPATCHER_RETURN",
	}
	ProceedOn_value = map[string]int32{
		"ABORTED":           0,
		"DROP":              1,
		"PASS":              2,
		"TX":                3,
		"REDIRECT":          4,
		"DISPATCHER_RETURN": 31,
	}
)

func (x ProceedOn) Enum() *ProceedOn {
	p := new(ProceedOn)
	*p = x
	return p
}

func (x ProceedOn) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProceedOn) Descriptor() protoreflect.EnumDescriptor {
	return file_bpfd_proto_enumTypes[1].Descriptor()
}

func (ProceedOn) Type() protoreflect.EnumType {
	return &file_bpfd_proto_enumTypes[1]
}

func (x ProceedOn) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProceedOn.Descriptor instead.
func (ProceedOn) EnumDescriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{1}
}

type LoadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path        string      `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	SectionName string      `protobuf:"bytes,2,opt,name=section_name,json=sectionName,proto3" json:"section_name,omitempty"`
	ProgramType ProgramType `protobuf:"varint,3,opt,name=program_type,json=programType,proto3,enum=bpfd.v1.ProgramType" json:"program_type,omitempty"`
	Priority    int32       `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	Iface       string      `protobuf:"bytes,5,opt,name=iface,proto3" json:"iface,omitempty"`
	ProceedOn   []ProceedOn `protobuf:"varint,6,rep,packed,name=proceed_on,json=proceedOn,proto3,enum=bpfd.v1.ProceedOn" json:"proceed_on,omitempty"`
}

func (x *LoadRequest) Reset() {
	*x = LoadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadRequest) ProtoMessage() {}

func (x *LoadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadRequest.ProtoReflect.Descriptor instead.
func (*LoadRequest) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{0}
}

func (x *LoadRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *LoadRequest) GetSectionName() string {
	if x != nil {
		return x.SectionName
	}
	return ""
}

func (x *LoadRequest) GetProgramType() ProgramType {
	if x != nil {
		return x.ProgramType
	}
	return ProgramType_XDP
}

func (x *LoadRequest) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *LoadRequest) GetIface() string {
	if x != nil {
		return x.Iface
	}
	return ""
}

func (x *LoadRequest) GetProceedOn() []ProceedOn {
	if x != nil {
		return x.ProceedOn
	}
	return nil
}

type LoadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LoadResponse) Reset() {
	*x = LoadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadResponse) ProtoMessage() {}

func (x *LoadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadResponse.ProtoReflect.Descriptor instead.
func (*LoadResponse) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{1}
}

func (x *LoadResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UnloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iface string `protobuf:"bytes,1,opt,name=iface,proto3" json:"iface,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UnloadRequest) Reset() {
	*x = UnloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnloadRequest) ProtoMessage() {}

func (x *UnloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnloadRequest.ProtoReflect.Descriptor instead.
func (*UnloadRequest) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{2}
}

func (x *UnloadRequest) GetIface() string {
	if x != nil {
		return x.Iface
	}
	return ""
}

func (x *UnloadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UnloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnloadResponse) Reset() {
	*x = UnloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnloadResponse) ProtoMessage() {}

func (x *UnloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnloadResponse.ProtoReflect.Descriptor instead.
func (*UnloadResponse) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{3}
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iface string `protobuf:"bytes,1,opt,name=iface,proto3" json:"iface,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{4}
}

func (x *ListRequest) GetIface() string {
	if x != nil {
		return x.Iface
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XdpMode string                     `protobuf:"bytes,1,opt,name=xdp_mode,json=xdpMode,proto3" json:"xdp_mode,omitempty"`
	Results []*ListResponse_ListResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{5}
}

func (x *ListResponse) GetXdpMode() string {
	if x != nil {
		return x.XdpMode
	}
	return ""
}

func (x *ListResponse) GetResults() []*ListResponse_ListResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type GetMapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iface      string `protobuf:"bytes,1,opt,name=iface,proto3" json:"iface,omitempty"`
	Id         string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	MapName    string `protobuf:"bytes,3,opt,name=map_name,json=mapName,proto3" json:"map_name,omitempty"`
	SocketPath string `protobuf:"bytes,4,opt,name=socket_path,json=socketPath,proto3" json:"socket_path,omitempty"`
}

func (x *GetMapRequest) Reset() {
	*x = GetMapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMapRequest) ProtoMessage() {}

func (x *GetMapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMapRequest.ProtoReflect.Descriptor instead.
func (*GetMapRequest) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{6}
}

func (x *GetMapRequest) GetIface() string {
	if x != nil {
		return x.Iface
	}
	return ""
}

func (x *GetMapRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetMapRequest) GetMapName() string {
	if x != nil {
		return x.MapName
	}
	return ""
}

func (x *GetMapRequest) GetSocketPath() string {
	if x != nil {
		return x.SocketPath
	}
	return ""
}

type GetMapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMapResponse) Reset() {
	*x = GetMapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMapResponse) ProtoMessage() {}

func (x *GetMapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMapResponse.ProtoReflect.Descriptor instead.
func (*GetMapResponse) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{7}
}

type ListResponse_ListResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Position  uint32      `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
	Priority  int32       `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	Path      string      `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	ProceedOn []ProceedOn `protobuf:"varint,6,rep,packed,name=proceed_on,json=proceedOn,proto3,enum=bpfd.v1.ProceedOn" json:"proceed_on,omitempty"`
}

func (x *ListResponse_ListResult) Reset() {
	*x = ListResponse_ListResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse_ListResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse_ListResult) ProtoMessage() {}

func (x *ListResponse_ListResult) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse_ListResult.ProtoReflect.Descriptor instead.
func (*ListResponse_ListResult) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ListResponse_ListResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListResponse_ListResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListResponse_ListResult) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *ListResponse_ListResult) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *ListResponse_ListResult) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ListResponse_ListResult) GetProceedOn() []ProceedOn {
	if x != nil {
		return x.ProceedOn
	}
	return nil
}

var File_bpfd_proto protoreflect.FileDescriptor

var file_bpfd_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x70,
	0x66, 0x64, 0x2e, 0x76, 0x31, 0x22, 0xe2, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x66, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x66, 0x61, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x62, 0x70,
	0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x65, 0x64, 0x4f, 0x6e, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0x1e, 0x0a, 0x0c, 0x4c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0d, 0x55, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x66, 0x61, 0x63, 0x65, 0x22, 0x97, 0x02, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x78, 0x64, 0x70,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x78, 0x64, 0x70,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x1a, 0xaf, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x31, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x65, 0x64, 0x4f, 0x6e, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x65, 0x64,
	0x4f, 0x6e, 0x22, 0x71, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x66, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x70,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x35, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x44, 0x50, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x54, 0x43, 0x5f, 0x49, 0x4e, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x54, 0x43, 0x5f, 0x45, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x2a, 0x59,
	0x0a, 0x09, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x41,
	0x42, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x52, 0x4f, 0x50,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x53, 0x53, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02,
	0x54, 0x58, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54,
	0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x49, 0x53, 0x50, 0x41, 0x54, 0x43, 0x48, 0x45, 0x52,
	0x5f, 0x52, 0x45, 0x54, 0x55, 0x52, 0x4e, 0x10, 0x1f, 0x32, 0xe8, 0x01, 0x0a, 0x06, 0x4c, 0x6f,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x2e, 0x62,
	0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x55, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x70,
	0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x62,
	0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x70, 0x12, 0x16, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x70,
	0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x68, 0x61, 0x74, 0x2d, 0x65, 0x74, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x62, 0x70, 0x66, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bpfd_proto_rawDescOnce sync.Once
	file_bpfd_proto_rawDescData = file_bpfd_proto_rawDesc
)

func file_bpfd_proto_rawDescGZIP() []byte {
	file_bpfd_proto_rawDescOnce.Do(func() {
		file_bpfd_proto_rawDescData = protoimpl.X.CompressGZIP(file_bpfd_proto_rawDescData)
	})
	return file_bpfd_proto_rawDescData
}

var file_bpfd_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bpfd_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_bpfd_proto_goTypes = []interface{}{
	(ProgramType)(0),                // 0: bpfd.v1.ProgramType
	(ProceedOn)(0),                  // 1: bpfd.v1.ProceedOn
	(*LoadRequest)(nil),             // 2: bpfd.v1.LoadRequest
	(*LoadResponse)(nil),            // 3: bpfd.v1.LoadResponse
	(*UnloadRequest)(nil),           // 4: bpfd.v1.UnloadRequest
	(*UnloadResponse)(nil),          // 5: bpfd.v1.UnloadResponse
	(*ListRequest)(nil),             // 6: bpfd.v1.ListRequest
	(*ListResponse)(nil),            // 7: bpfd.v1.ListResponse
	(*GetMapRequest)(nil),           // 8: bpfd.v1.GetMapRequest
	(*GetMapResponse)(nil),          // 9: bpfd.v1.GetMapResponse
	(*ListResponse_ListResult)(nil), // 10: bpfd.v1.ListResponse.ListResult
}
var file_bpfd_proto_depIdxs = []int32{
	0,  // 0: bpfd.v1.LoadRequest.program_type:type_name -> bpfd.v1.ProgramType
	1,  // 1: bpfd.v1.LoadRequest.proceed_on:type_name -> bpfd.v1.ProceedOn
	10, // 2: bpfd.v1.ListResponse.results:type_name -> bpfd.v1.ListResponse.ListResult
	1,  // 3: bpfd.v1.ListResponse.ListResult.proceed_on:type_name -> bpfd.v1.ProceedOn
	2,  // 4: bpfd.v1.Loader.Load:input_type -> bpfd.v1.LoadRequest
	4,  // 5: bpfd.v1.Loader.Unload:input_type -> bpfd.v1.UnloadRequest
	6,  // 6: bpfd.v1.Loader.List:input_type -> bpfd.v1.ListRequest
	8,  // 7: bpfd.v1.Loader.GetMap:input_type -> bpfd.v1.GetMapRequest
	3,  // 8: bpfd.v1.Loader.Load:output_type -> bpfd.v1.LoadResponse
	5,  // 9: bpfd.v1.Loader.Unload:output_type -> bpfd.v1.UnloadResponse
	7,  // 10: bpfd.v1.Loader.List:output_type -> bpfd.v1.ListResponse
	9,  // 11: bpfd.v1.Loader.GetMap:output_type -> bpfd.v1.GetMapResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_bpfd_proto_init() }
func file_bpfd_proto_init() {
	if File_bpfd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bpfd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadRequest); i {
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
		file_bpfd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadResponse); i {
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
		file_bpfd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnloadRequest); i {
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
		file_bpfd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnloadResponse); i {
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
		file_bpfd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_bpfd_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_bpfd_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMapRequest); i {
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
		file_bpfd_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMapResponse); i {
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
		file_bpfd_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse_ListResult); i {
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
			RawDescriptor: file_bpfd_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bpfd_proto_goTypes,
		DependencyIndexes: file_bpfd_proto_depIdxs,
		EnumInfos:         file_bpfd_proto_enumTypes,
		MessageInfos:      file_bpfd_proto_msgTypes,
	}.Build()
	File_bpfd_proto = out.File
	file_bpfd_proto_rawDesc = nil
	file_bpfd_proto_goTypes = nil
	file_bpfd_proto_depIdxs = nil
}