// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow/core/function/polymorphism/function_type.proto

package polymorphism

import (
	trace_type "github.com/yangchnet/tf-serving/tensorflow/core/function/trace_type"
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

type Parameter_Kind int32

const (
	Parameter_UNDEFINED             Parameter_Kind = 0
	Parameter_POSITIONAL_ONLY       Parameter_Kind = 1
	Parameter_POSITIONAL_OR_KEYWORD Parameter_Kind = 2
	Parameter_VAR_POSITIONAL        Parameter_Kind = 3
	Parameter_KEYWORD_ONLY          Parameter_Kind = 4
	Parameter_VAR_KEYWORD           Parameter_Kind = 5
)

// Enum value maps for Parameter_Kind.
var (
	Parameter_Kind_name = map[int32]string{
		0: "UNDEFINED",
		1: "POSITIONAL_ONLY",
		2: "POSITIONAL_OR_KEYWORD",
		3: "VAR_POSITIONAL",
		4: "KEYWORD_ONLY",
		5: "VAR_KEYWORD",
	}
	Parameter_Kind_value = map[string]int32{
		"UNDEFINED":             0,
		"POSITIONAL_ONLY":       1,
		"POSITIONAL_OR_KEYWORD": 2,
		"VAR_POSITIONAL":        3,
		"KEYWORD_ONLY":          4,
		"VAR_KEYWORD":           5,
	}
)

func (x Parameter_Kind) Enum() *Parameter_Kind {
	p := new(Parameter_Kind)
	*p = x
	return p
}

func (x Parameter_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Parameter_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_function_polymorphism_function_type_proto_enumTypes[0].Descriptor()
}

func (Parameter_Kind) Type() protoreflect.EnumType {
	return &file_tensorflow_core_function_polymorphism_function_type_proto_enumTypes[0]
}

func (x Parameter_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Parameter_Kind) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Parameter_Kind(num)
	return nil
}

// Deprecated: Use Parameter_Kind.Descriptor instead.
func (Parameter_Kind) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_function_polymorphism_function_type_proto_rawDescGZIP(), []int{0, 0}
}

// Represents a serialized Parameter type.
type Parameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           *string                         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Kind           *Parameter_Kind                 `protobuf:"varint,2,opt,name=kind,enum=tensorflow.core.function.polymorphism.function_type.Parameter_Kind" json:"kind,omitempty"`
	IsOptional     *bool                           `protobuf:"varint,3,opt,name=is_optional,json=isOptional" json:"is_optional,omitempty"`
	TypeConstraint *trace_type.SerializedTraceType `protobuf:"bytes,4,opt,name=type_constraint,json=typeConstraint" json:"type_constraint,omitempty"`
}

func (x *Parameter) Reset() {
	*x = Parameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_function_polymorphism_function_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parameter) ProtoMessage() {}

func (x *Parameter) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_function_polymorphism_function_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parameter.ProtoReflect.Descriptor instead.
func (*Parameter) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_function_polymorphism_function_type_proto_rawDescGZIP(), []int{0}
}

func (x *Parameter) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Parameter) GetKind() Parameter_Kind {
	if x != nil && x.Kind != nil {
		return *x.Kind
	}
	return Parameter_UNDEFINED
}

func (x *Parameter) GetIsOptional() bool {
	if x != nil && x.IsOptional != nil {
		return *x.IsOptional
	}
	return false
}

func (x *Parameter) GetTypeConstraint() *trace_type.SerializedTraceType {
	if x != nil {
		return x.TypeConstraint
	}
	return nil
}

// Represents a serialized Capture type.
type Capture struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           *string                         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	TypeConstraint *trace_type.SerializedTraceType `protobuf:"bytes,2,opt,name=type_constraint,json=typeConstraint" json:"type_constraint,omitempty"`
}

func (x *Capture) Reset() {
	*x = Capture{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_function_polymorphism_function_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capture) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capture) ProtoMessage() {}

func (x *Capture) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_function_polymorphism_function_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capture.ProtoReflect.Descriptor instead.
func (*Capture) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_function_polymorphism_function_type_proto_rawDescGZIP(), []int{1}
}

func (x *Capture) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Capture) GetTypeConstraint() *trace_type.SerializedTraceType {
	if x != nil {
		return x.TypeConstraint
	}
	return nil
}

// Represents a serialized FunctionType.
type FunctionType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameters []*Parameter `protobuf:"bytes,1,rep,name=parameters" json:"parameters,omitempty"`
	Captures   []*Capture   `protobuf:"bytes,2,rep,name=captures" json:"captures,omitempty"` // TODO(fmuham): Add support for return type.
}

func (x *FunctionType) Reset() {
	*x = FunctionType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_function_polymorphism_function_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionType) ProtoMessage() {}

func (x *FunctionType) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_function_polymorphism_function_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionType.ProtoReflect.Descriptor instead.
func (*FunctionType) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_function_polymorphism_function_type_proto_rawDescGZIP(), []int{2}
}

func (x *FunctionType) GetParameters() []*Parameter {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *FunctionType) GetCaptures() []*Capture {
	if x != nil {
		return x.Captures
	}
	return nil
}

var File_tensorflow_core_function_polymorphism_function_type_proto protoreflect.FileDescriptor

var file_tensorflow_core_function_polymorphism_function_type_proto_rawDesc = []byte{
	0x0a, 0x39, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x6d,
	0x6f, 0x72, 0x70, 0x68, 0x69, 0x73, 0x6d, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x33, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x69,
	0x73, 0x6d, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x1a, 0x37, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x03, 0x0a, 0x09, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x57, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x43, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x69, 0x73,
	0x6d, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x6f, 0x0a, 0x0f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x22, 0x7c, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0d,
	0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x4f, 0x4e, 0x4c, 0x59,
	0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c,
	0x5f, 0x4f, 0x52, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x02, 0x12, 0x12, 0x0a,
	0x0e, 0x56, 0x41, 0x52, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10,
	0x03, 0x12, 0x10, 0x0a, 0x0c, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x4f, 0x4e, 0x4c,
	0x59, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x41, 0x52, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f,
	0x52, 0x44, 0x10, 0x05, 0x22, 0x8e, 0x01, 0x0a, 0x07, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x6f, 0x0a, 0x0f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x22, 0xc8, 0x01, 0x0a, 0x0c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x69,
	0x73, 0x6d, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x58, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x69, 0x73, 0x6d,
	0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43,
	0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x52, 0x08, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x42, 0x84, 0x03, 0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x69, 0x73, 0x6d, 0x2e, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x11, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48,
	0x02, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71,
	0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x6c,
	0x79, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x69, 0x73, 0x6d, 0xa2, 0x02, 0x05, 0x54, 0x43, 0x46, 0x50,
	0x46, 0xaa, 0x02, 0x32, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43,
	0x6f, 0x72, 0x65, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6c,
	0x79, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x69, 0x73, 0x6d, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0xca, 0x02, 0x33, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x5c, 0x50, 0x6f, 0x6c, 0x79, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x69, 0x73, 0x6d, 0x5c,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0xe2, 0x02, 0x3f, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x5c, 0x50, 0x6f, 0x6c, 0x79, 0x6d, 0x6f, 0x72,
	0x70, 0x68, 0x69, 0x73, 0x6d, 0x5c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x36, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x43, 0x6f, 0x72,
	0x65, 0x3a, 0x3a, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x6f, 0x6c,
	0x79, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x69, 0x73, 0x6d, 0x3a, 0x3a, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
}

var (
	file_tensorflow_core_function_polymorphism_function_type_proto_rawDescOnce sync.Once
	file_tensorflow_core_function_polymorphism_function_type_proto_rawDescData = file_tensorflow_core_function_polymorphism_function_type_proto_rawDesc
)

func file_tensorflow_core_function_polymorphism_function_type_proto_rawDescGZIP() []byte {
	file_tensorflow_core_function_polymorphism_function_type_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_function_polymorphism_function_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_function_polymorphism_function_type_proto_rawDescData)
	})
	return file_tensorflow_core_function_polymorphism_function_type_proto_rawDescData
}

var file_tensorflow_core_function_polymorphism_function_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_function_polymorphism_function_type_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_core_function_polymorphism_function_type_proto_goTypes = []interface{}{
	(Parameter_Kind)(0),                    // 0: tensorflow.core.function.polymorphism.function_type.Parameter.Kind
	(*Parameter)(nil),                      // 1: tensorflow.core.function.polymorphism.function_type.Parameter
	(*Capture)(nil),                        // 2: tensorflow.core.function.polymorphism.function_type.Capture
	(*FunctionType)(nil),                   // 3: tensorflow.core.function.polymorphism.function_type.FunctionType
	(*trace_type.SerializedTraceType)(nil), // 4: tensorflow.core.function.trace_type.serialization.SerializedTraceType
}
var file_tensorflow_core_function_polymorphism_function_type_proto_depIdxs = []int32{
	0, // 0: tensorflow.core.function.polymorphism.function_type.Parameter.kind:type_name -> tensorflow.core.function.polymorphism.function_type.Parameter.Kind
	4, // 1: tensorflow.core.function.polymorphism.function_type.Parameter.type_constraint:type_name -> tensorflow.core.function.trace_type.serialization.SerializedTraceType
	4, // 2: tensorflow.core.function.polymorphism.function_type.Capture.type_constraint:type_name -> tensorflow.core.function.trace_type.serialization.SerializedTraceType
	1, // 3: tensorflow.core.function.polymorphism.function_type.FunctionType.parameters:type_name -> tensorflow.core.function.polymorphism.function_type.Parameter
	2, // 4: tensorflow.core.function.polymorphism.function_type.FunctionType.captures:type_name -> tensorflow.core.function.polymorphism.function_type.Capture
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tensorflow_core_function_polymorphism_function_type_proto_init() }
func file_tensorflow_core_function_polymorphism_function_type_proto_init() {
	if File_tensorflow_core_function_polymorphism_function_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_function_polymorphism_function_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parameter); i {
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
		file_tensorflow_core_function_polymorphism_function_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capture); i {
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
		file_tensorflow_core_function_polymorphism_function_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionType); i {
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
			RawDescriptor: file_tensorflow_core_function_polymorphism_function_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_function_polymorphism_function_type_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_function_polymorphism_function_type_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_function_polymorphism_function_type_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_function_polymorphism_function_type_proto_msgTypes,
	}.Build()
	File_tensorflow_core_function_polymorphism_function_type_proto = out.File
	file_tensorflow_core_function_polymorphism_function_type_proto_rawDesc = nil
	file_tensorflow_core_function_polymorphism_function_type_proto_goTypes = nil
	file_tensorflow_core_function_polymorphism_function_type_proto_depIdxs = nil
}
