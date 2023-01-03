// This proto intends to match format expected by pprof tool.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow/tsl/profiler/protobuf/profile.proto

package protobuf

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

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SampleType        []*ValueType `protobuf:"bytes,1,rep,name=sample_type,json=sampleType,proto3" json:"sample_type,omitempty"`
	Sample            []*Sample    `protobuf:"bytes,2,rep,name=sample,proto3" json:"sample,omitempty"`
	Mapping           []*Mapping   `protobuf:"bytes,3,rep,name=mapping,proto3" json:"mapping,omitempty"`
	Location          []*Location  `protobuf:"bytes,4,rep,name=location,proto3" json:"location,omitempty"`
	Function          []*Function  `protobuf:"bytes,5,rep,name=function,proto3" json:"function,omitempty"`
	StringTable       []string     `protobuf:"bytes,6,rep,name=string_table,json=stringTable,proto3" json:"string_table,omitempty"`
	DropFrames        int64        `protobuf:"varint,7,opt,name=drop_frames,json=dropFrames,proto3" json:"drop_frames,omitempty"`
	KeepFrames        int64        `protobuf:"varint,8,opt,name=keep_frames,json=keepFrames,proto3" json:"keep_frames,omitempty"`
	TimeNanos         int64        `protobuf:"varint,9,opt,name=time_nanos,json=timeNanos,proto3" json:"time_nanos,omitempty"`
	DurationNanos     int64        `protobuf:"varint,10,opt,name=duration_nanos,json=durationNanos,proto3" json:"duration_nanos,omitempty"`
	PeriodType        *ValueType   `protobuf:"bytes,11,opt,name=period_type,json=periodType,proto3" json:"period_type,omitempty"`
	Period            int64        `protobuf:"varint,12,opt,name=period,proto3" json:"period,omitempty"`
	Comment           []int64      `protobuf:"varint,13,rep,packed,name=comment,proto3" json:"comment,omitempty"`
	DefaultSampleType int64        `protobuf:"varint,14,opt,name=default_sample_type,json=defaultSampleType,proto3" json:"default_sample_type,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetSampleType() []*ValueType {
	if x != nil {
		return x.SampleType
	}
	return nil
}

func (x *Profile) GetSample() []*Sample {
	if x != nil {
		return x.Sample
	}
	return nil
}

func (x *Profile) GetMapping() []*Mapping {
	if x != nil {
		return x.Mapping
	}
	return nil
}

func (x *Profile) GetLocation() []*Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Profile) GetFunction() []*Function {
	if x != nil {
		return x.Function
	}
	return nil
}

func (x *Profile) GetStringTable() []string {
	if x != nil {
		return x.StringTable
	}
	return nil
}

func (x *Profile) GetDropFrames() int64 {
	if x != nil {
		return x.DropFrames
	}
	return 0
}

func (x *Profile) GetKeepFrames() int64 {
	if x != nil {
		return x.KeepFrames
	}
	return 0
}

func (x *Profile) GetTimeNanos() int64 {
	if x != nil {
		return x.TimeNanos
	}
	return 0
}

func (x *Profile) GetDurationNanos() int64 {
	if x != nil {
		return x.DurationNanos
	}
	return 0
}

func (x *Profile) GetPeriodType() *ValueType {
	if x != nil {
		return x.PeriodType
	}
	return nil
}

func (x *Profile) GetPeriod() int64 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *Profile) GetComment() []int64 {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *Profile) GetDefaultSampleType() int64 {
	if x != nil {
		return x.DefaultSampleType
	}
	return 0
}

type ValueType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int64 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Unit int64 `protobuf:"varint,2,opt,name=unit,proto3" json:"unit,omitempty"`
}

func (x *ValueType) Reset() {
	*x = ValueType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueType) ProtoMessage() {}

func (x *ValueType) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueType.ProtoReflect.Descriptor instead.
func (*ValueType) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescGZIP(), []int{1}
}

func (x *ValueType) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ValueType) GetUnit() int64 {
	if x != nil {
		return x.Unit
	}
	return 0
}

type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationId []uint64 `protobuf:"varint,1,rep,packed,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Value      []int64  `protobuf:"varint,2,rep,packed,name=value,proto3" json:"value,omitempty"`
	Label      []*Label `protobuf:"bytes,3,rep,name=label,proto3" json:"label,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescGZIP(), []int{2}
}

func (x *Sample) GetLocationId() []uint64 {
	if x != nil {
		return x.LocationId
	}
	return nil
}

func (x *Sample) GetValue() []int64 {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Sample) GetLabel() []*Label {
	if x != nil {
		return x.Label
	}
	return nil
}

type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key int64 `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Str int64 `protobuf:"varint,2,opt,name=str,proto3" json:"str,omitempty"`
	Num int64 `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescGZIP(), []int{3}
}

func (x *Label) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *Label) GetStr() int64 {
	if x != nil {
		return x.Str
	}
	return 0
}

func (x *Label) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type Mapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MemoryStart     uint64 `protobuf:"varint,2,opt,name=memory_start,json=memoryStart,proto3" json:"memory_start,omitempty"`
	MemoryLimit     uint64 `protobuf:"varint,3,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	FileOffset      uint64 `protobuf:"varint,4,opt,name=file_offset,json=fileOffset,proto3" json:"file_offset,omitempty"`
	Filename        int64  `protobuf:"varint,5,opt,name=filename,proto3" json:"filename,omitempty"`
	BuildId         int64  `protobuf:"varint,6,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	HasFunctions    bool   `protobuf:"varint,7,opt,name=has_functions,json=hasFunctions,proto3" json:"has_functions,omitempty"`
	HasFilenames    bool   `protobuf:"varint,8,opt,name=has_filenames,json=hasFilenames,proto3" json:"has_filenames,omitempty"`
	HasLineNumbers  bool   `protobuf:"varint,9,opt,name=has_line_numbers,json=hasLineNumbers,proto3" json:"has_line_numbers,omitempty"`
	HasInlineFrames bool   `protobuf:"varint,10,opt,name=has_inline_frames,json=hasInlineFrames,proto3" json:"has_inline_frames,omitempty"`
}

func (x *Mapping) Reset() {
	*x = Mapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mapping) ProtoMessage() {}

func (x *Mapping) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mapping.ProtoReflect.Descriptor instead.
func (*Mapping) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescGZIP(), []int{4}
}

func (x *Mapping) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Mapping) GetMemoryStart() uint64 {
	if x != nil {
		return x.MemoryStart
	}
	return 0
}

func (x *Mapping) GetMemoryLimit() uint64 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *Mapping) GetFileOffset() uint64 {
	if x != nil {
		return x.FileOffset
	}
	return 0
}

func (x *Mapping) GetFilename() int64 {
	if x != nil {
		return x.Filename
	}
	return 0
}

func (x *Mapping) GetBuildId() int64 {
	if x != nil {
		return x.BuildId
	}
	return 0
}

func (x *Mapping) GetHasFunctions() bool {
	if x != nil {
		return x.HasFunctions
	}
	return false
}

func (x *Mapping) GetHasFilenames() bool {
	if x != nil {
		return x.HasFilenames
	}
	return false
}

func (x *Mapping) GetHasLineNumbers() bool {
	if x != nil {
		return x.HasLineNumbers
	}
	return false
}

func (x *Mapping) GetHasInlineFrames() bool {
	if x != nil {
		return x.HasInlineFrames
	}
	return false
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MappingId uint64  `protobuf:"varint,2,opt,name=mapping_id,json=mappingId,proto3" json:"mapping_id,omitempty"`
	Address   uint64  `protobuf:"varint,3,opt,name=address,proto3" json:"address,omitempty"`
	Line      []*Line `protobuf:"bytes,4,rep,name=line,proto3" json:"line,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescGZIP(), []int{5}
}

func (x *Location) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Location) GetMappingId() uint64 {
	if x != nil {
		return x.MappingId
	}
	return 0
}

func (x *Location) GetAddress() uint64 {
	if x != nil {
		return x.Address
	}
	return 0
}

func (x *Location) GetLine() []*Line {
	if x != nil {
		return x.Line
	}
	return nil
}

type Line struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FunctionId uint64 `protobuf:"varint,1,opt,name=function_id,json=functionId,proto3" json:"function_id,omitempty"`
	Line       int64  `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
}

func (x *Line) Reset() {
	*x = Line{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Line) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Line) ProtoMessage() {}

func (x *Line) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Line.ProtoReflect.Descriptor instead.
func (*Line) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescGZIP(), []int{6}
}

func (x *Line) GetFunctionId() uint64 {
	if x != nil {
		return x.FunctionId
	}
	return 0
}

func (x *Line) GetLine() int64 {
	if x != nil {
		return x.Line
	}
	return 0
}

type Function struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       int64  `protobuf:"varint,2,opt,name=name,proto3" json:"name,omitempty"`
	SystemName int64  `protobuf:"varint,3,opt,name=system_name,json=systemName,proto3" json:"system_name,omitempty"`
	Filename   int64  `protobuf:"varint,4,opt,name=filename,proto3" json:"filename,omitempty"`
	StartLine  int64  `protobuf:"varint,5,opt,name=start_line,json=startLine,proto3" json:"start_line,omitempty"`
}

func (x *Function) Reset() {
	*x = Function{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Function) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Function) ProtoMessage() {}

func (x *Function) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Function.ProtoReflect.Descriptor instead.
func (*Function) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescGZIP(), []int{7}
}

func (x *Function) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Function) GetName() int64 {
	if x != nil {
		return x.Name
	}
	return 0
}

func (x *Function) GetSystemName() int64 {
	if x != nil {
		return x.SystemName
	}
	return 0
}

func (x *Function) GetFilename() int64 {
	if x != nil {
		return x.Filename
	}
	return 0
}

func (x *Function) GetStartLine() int64 {
	if x != nil {
		return x.StartLine
	}
	return 0
}

var File_tensorflow_tsl_profiler_protobuf_profile_proto protoreflect.FileDescriptor

var file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x66, 0x70,
	0x72, 0x6f, 0x66, 0x2e, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x22, 0x93, 0x05, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x66, 0x70, 0x72, 0x6f, 0x66, 0x2e, 0x70,
	0x70, 0x72, 0x6f, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x66, 0x70, 0x72, 0x6f, 0x66, 0x2e, 0x70,
	0x70, 0x72, 0x6f, 0x66, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x06, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x74, 0x66, 0x70, 0x72, 0x6f, 0x66, 0x2e, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x2e, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12,
	0x3d, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74,
	0x66, 0x70, 0x72, 0x6f, 0x66, 0x2e, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d,
	0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x66,
	0x70, 0x72, 0x6f, 0x66, 0x2e, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x72, 0x6f, 0x70, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6b, 0x65, 0x65, 0x70, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4e, 0x61, 0x6e, 0x6f,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61,
	0x6e, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6e, 0x6f, 0x73, 0x12, 0x43, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x66, 0x70, 0x72, 0x6f,
	0x66, 0x2e, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x2e, 0x0a, 0x13, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x33, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x22, 0x75, 0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x74, 0x66, 0x70, 0x72, 0x6f, 0x66, 0x2e, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x3d, 0x0a, 0x05, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0xd7, 0x02, 0x0a, 0x07, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x5f, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x61, 0x73,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x61, 0x73,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x68, 0x61, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x28,
	0x0a, 0x10, 0x68, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x4c, 0x69, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x68, 0x61, 0x73, 0x5f,
	0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x68, 0x61, 0x73, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x66, 0x70, 0x72, 0x6f, 0x66, 0x2e, 0x70, 0x70, 0x72,
	0x6f, 0x66, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x3b, 0x0a,
	0x04, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x08, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x42, 0xec, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x66, 0x70, 0x72, 0x6f,
	0x66, 0x2e, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x65, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f, 0x74,
	0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xa2, 0x02, 0x03, 0x54, 0x54, 0x50,
	0xaa, 0x02, 0x17, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x66,
	0x70, 0x72, 0x6f, 0x66, 0x2e, 0x50, 0x70, 0x72, 0x6f, 0x66, 0xca, 0x02, 0x17, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x54, 0x66, 0x70, 0x72, 0x6f, 0x66, 0x5c, 0x50,
	0x70, 0x72, 0x6f, 0x66, 0xe2, 0x02, 0x23, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x5c, 0x54, 0x66, 0x70, 0x72, 0x6f, 0x66, 0x5c, 0x50, 0x70, 0x72, 0x6f, 0x66, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x54, 0x66, 0x70, 0x72, 0x6f, 0x66, 0x3a,
	0x3a, 0x50, 0x70, 0x72, 0x6f, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescOnce sync.Once
	file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescData = file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDesc
)

func file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescGZIP() []byte {
	file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescOnce.Do(func() {
		file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescData)
	})
	return file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDescData
}

var file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tensorflow_tsl_profiler_protobuf_profile_proto_goTypes = []interface{}{
	(*Profile)(nil),   // 0: tensorflow.tfprof.pprof.Profile
	(*ValueType)(nil), // 1: tensorflow.tfprof.pprof.ValueType
	(*Sample)(nil),    // 2: tensorflow.tfprof.pprof.Sample
	(*Label)(nil),     // 3: tensorflow.tfprof.pprof.Label
	(*Mapping)(nil),   // 4: tensorflow.tfprof.pprof.Mapping
	(*Location)(nil),  // 5: tensorflow.tfprof.pprof.Location
	(*Line)(nil),      // 6: tensorflow.tfprof.pprof.Line
	(*Function)(nil),  // 7: tensorflow.tfprof.pprof.Function
}
var file_tensorflow_tsl_profiler_protobuf_profile_proto_depIdxs = []int32{
	1, // 0: tensorflow.tfprof.pprof.Profile.sample_type:type_name -> tensorflow.tfprof.pprof.ValueType
	2, // 1: tensorflow.tfprof.pprof.Profile.sample:type_name -> tensorflow.tfprof.pprof.Sample
	4, // 2: tensorflow.tfprof.pprof.Profile.mapping:type_name -> tensorflow.tfprof.pprof.Mapping
	5, // 3: tensorflow.tfprof.pprof.Profile.location:type_name -> tensorflow.tfprof.pprof.Location
	7, // 4: tensorflow.tfprof.pprof.Profile.function:type_name -> tensorflow.tfprof.pprof.Function
	1, // 5: tensorflow.tfprof.pprof.Profile.period_type:type_name -> tensorflow.tfprof.pprof.ValueType
	3, // 6: tensorflow.tfprof.pprof.Sample.label:type_name -> tensorflow.tfprof.pprof.Label
	6, // 7: tensorflow.tfprof.pprof.Location.line:type_name -> tensorflow.tfprof.pprof.Line
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_tensorflow_tsl_profiler_protobuf_profile_proto_init() }
func file_tensorflow_tsl_profiler_protobuf_profile_proto_init() {
	if File_tensorflow_tsl_profiler_protobuf_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueType); i {
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
		file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
		file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
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
		file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mapping); i {
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
		file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Line); i {
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
		file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Function); i {
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
			RawDescriptor: file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_tsl_profiler_protobuf_profile_proto_goTypes,
		DependencyIndexes: file_tensorflow_tsl_profiler_protobuf_profile_proto_depIdxs,
		MessageInfos:      file_tensorflow_tsl_profiler_protobuf_profile_proto_msgTypes,
	}.Build()
	File_tensorflow_tsl_profiler_protobuf_profile_proto = out.File
	file_tensorflow_tsl_profiler_protobuf_profile_proto_rawDesc = nil
	file_tensorflow_tsl_profiler_protobuf_profile_proto_goTypes = nil
	file_tensorflow_tsl_profiler_protobuf_profile_proto_depIdxs = nil
}