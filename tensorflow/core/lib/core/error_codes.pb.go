// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow/core/lib/core/error_codes.proto

package core

import (
	protobuf "gitee.com/qciip-icp/tf-serving/tensorflow/tsl/protobuf"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of tensorflow/tsl/protobuf/error_codes.proto.

type Code = protobuf.Code

const Code_OK = protobuf.Code_OK
const Code_CANCELLED = protobuf.Code_CANCELLED
const Code_UNKNOWN = protobuf.Code_UNKNOWN
const Code_INVALID_ARGUMENT = protobuf.Code_INVALID_ARGUMENT
const Code_DEADLINE_EXCEEDED = protobuf.Code_DEADLINE_EXCEEDED
const Code_NOT_FOUND = protobuf.Code_NOT_FOUND
const Code_ALREADY_EXISTS = protobuf.Code_ALREADY_EXISTS
const Code_PERMISSION_DENIED = protobuf.Code_PERMISSION_DENIED
const Code_UNAUTHENTICATED = protobuf.Code_UNAUTHENTICATED
const Code_RESOURCE_EXHAUSTED = protobuf.Code_RESOURCE_EXHAUSTED
const Code_FAILED_PRECONDITION = protobuf.Code_FAILED_PRECONDITION
const Code_ABORTED = protobuf.Code_ABORTED
const Code_OUT_OF_RANGE = protobuf.Code_OUT_OF_RANGE
const Code_UNIMPLEMENTED = protobuf.Code_UNIMPLEMENTED
const Code_INTERNAL = protobuf.Code_INTERNAL
const Code_UNAVAILABLE = protobuf.Code_UNAVAILABLE
const Code_DATA_LOSS = protobuf.Code_DATA_LOSS
const Code_DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_ = protobuf.Code_DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_

var Code_name = protobuf.Code_name
var Code_value = protobuf.Code_value

var File_tensorflow_core_lib_core_error_codes_proto protoreflect.FileDescriptor

var file_tensorflow_core_lib_core_error_codes_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x4e, 0x42, 0x0f, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x69, 0x69, 0x70, 0x2d,
	0x69, 0x63, 0x70, 0x2f, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c,
	0x69, 0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_tensorflow_core_lib_core_error_codes_proto_goTypes = []interface{}{}
var file_tensorflow_core_lib_core_error_codes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_lib_core_error_codes_proto_init() }
func file_tensorflow_core_lib_core_error_codes_proto_init() {
	if File_tensorflow_core_lib_core_error_codes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_lib_core_error_codes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_lib_core_error_codes_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_lib_core_error_codes_proto_depIdxs,
	}.Build()
	File_tensorflow_core_lib_core_error_codes_proto = out.File
	file_tensorflow_core_lib_core_error_codes_proto_rawDesc = nil
	file_tensorflow_core_lib_core_error_codes_proto_goTypes = nil
	file_tensorflow_core_lib_core_error_codes_proto_depIdxs = nil
}