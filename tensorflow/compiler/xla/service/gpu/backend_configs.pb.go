// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow/compiler/xla/service/gpu/backend_configs.proto

package gpu

import (
	xla "github.com/yangchnet/tf-serving/tensorflow/compiler/xla"
	_ "github.com/yangchnet/tf-serving/tensorflow/compiler/xla/stream_executor"
	protobuf "github.com/yangchnet/tf-serving/tensorflow/tsl/protobuf"
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

// cublasLt matmul epilogue.
type GemmBackendConfig_Epilogue int32

const (
	GemmBackendConfig_DEFAULT       GemmBackendConfig_Epilogue = 0
	GemmBackendConfig_BIAS          GemmBackendConfig_Epilogue = 1
	GemmBackendConfig_RELU          GemmBackendConfig_Epilogue = 2
	GemmBackendConfig_BIAS_RELU     GemmBackendConfig_Epilogue = 3
	GemmBackendConfig_GELU          GemmBackendConfig_Epilogue = 4
	GemmBackendConfig_GELU_AUX      GemmBackendConfig_Epilogue = 5
	GemmBackendConfig_BIAS_GELU     GemmBackendConfig_Epilogue = 6
	GemmBackendConfig_BIAS_GELU_AUX GemmBackendConfig_Epilogue = 7
)

// Enum value maps for GemmBackendConfig_Epilogue.
var (
	GemmBackendConfig_Epilogue_name = map[int32]string{
		0: "DEFAULT",
		1: "BIAS",
		2: "RELU",
		3: "BIAS_RELU",
		4: "GELU",
		5: "GELU_AUX",
		6: "BIAS_GELU",
		7: "BIAS_GELU_AUX",
	}
	GemmBackendConfig_Epilogue_value = map[string]int32{
		"DEFAULT":       0,
		"BIAS":          1,
		"RELU":          2,
		"BIAS_RELU":     3,
		"GELU":          4,
		"GELU_AUX":      5,
		"BIAS_GELU":     6,
		"BIAS_GELU_AUX": 7,
	}
)

func (x GemmBackendConfig_Epilogue) Enum() *GemmBackendConfig_Epilogue {
	p := new(GemmBackendConfig_Epilogue)
	*p = x
	return p
}

func (x GemmBackendConfig_Epilogue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GemmBackendConfig_Epilogue) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_enumTypes[0].Descriptor()
}

func (GemmBackendConfig_Epilogue) Type() protoreflect.EnumType {
	return &file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_enumTypes[0]
}

func (x GemmBackendConfig_Epilogue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GemmBackendConfig_Epilogue.Descriptor instead.
func (GemmBackendConfig_Epilogue) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDescGZIP(), []int{1, 0}
}

// Backend config for a convolution that runs through cudnn.
type CudnnConvBackendConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Opaque algorithm number and tuning knobs chosen for this conv.
	Algorithm *protobuf.AlgorithmProto `protobuf:"bytes,6,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	// The scaling factor multiplied with the convolution result.
	ConvResultScale float64 `protobuf:"fixed64,4,opt,name=conv_result_scale,json=convResultScale,proto3" json:"conv_result_scale,omitempty"`
	// The requested activation (e.g. relu) after the convolution. It is with type
	// stream_executor::dnn::ActivationMode.
	ActivationMode int64 `protobuf:"varint,3,opt,name=activation_mode,json=activationMode,proto3" json:"activation_mode,omitempty"`
	// The scaling factor multiplied with the side input. If no side input buffer
	// is provided, this field must be 0.
	SideInputScale float64 `protobuf:"fixed64,5,opt,name=side_input_scale,json=sideInputScale,proto3" json:"side_input_scale,omitempty"`
}

func (x *CudnnConvBackendConfig) Reset() {
	*x = CudnnConvBackendConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CudnnConvBackendConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CudnnConvBackendConfig) ProtoMessage() {}

func (x *CudnnConvBackendConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CudnnConvBackendConfig.ProtoReflect.Descriptor instead.
func (*CudnnConvBackendConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDescGZIP(), []int{0}
}

func (x *CudnnConvBackendConfig) GetAlgorithm() *protobuf.AlgorithmProto {
	if x != nil {
		return x.Algorithm
	}
	return nil
}

func (x *CudnnConvBackendConfig) GetConvResultScale() float64 {
	if x != nil {
		return x.ConvResultScale
	}
	return 0
}

func (x *CudnnConvBackendConfig) GetActivationMode() int64 {
	if x != nil {
		return x.ActivationMode
	}
	return 0
}

func (x *CudnnConvBackendConfig) GetSideInputScale() float64 {
	if x != nil {
		return x.SideInputScale
	}
	return 0
}

// Backend config for the GEMM operation running through cuBLAS.
type GemmBackendConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Opaque optional algorithm number. No chosen number indicates that a
	// different cuBLAS API will be used, which does not allow for choosing an
	// algorithm.
	//
	// Types that are assignable to Algorithm:
	//
	//	*GemmBackendConfig_SelectedAlgorithm
	Algorithm           isGemmBackendConfig_Algorithm `protobuf_oneof:"algorithm"`
	AlphaReal           float64                       `protobuf:"fixed64,2,opt,name=alpha_real,json=alphaReal,proto3" json:"alpha_real,omitempty"`
	AlphaImag           float64                       `protobuf:"fixed64,9,opt,name=alpha_imag,json=alphaImag,proto3" json:"alpha_imag,omitempty"`
	Beta                float64                       `protobuf:"fixed64,3,opt,name=beta,proto3" json:"beta,omitempty"`
	DotDimensionNumbers *xla.DotDimensionNumbers      `protobuf:"bytes,7,opt,name=dot_dimension_numbers,json=dotDimensionNumbers,proto3" json:"dot_dimension_numbers,omitempty"`
	PrecisionConfig     *xla.PrecisionConfig          `protobuf:"bytes,12,opt,name=precision_config,json=precisionConfig,proto3" json:"precision_config,omitempty"`
	Epilogue            GemmBackendConfig_Epilogue    `protobuf:"varint,13,opt,name=epilogue,proto3,enum=xla.gpu.GemmBackendConfig_Epilogue" json:"epilogue,omitempty"`
}

func (x *GemmBackendConfig) Reset() {
	*x = GemmBackendConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GemmBackendConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GemmBackendConfig) ProtoMessage() {}

func (x *GemmBackendConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GemmBackendConfig.ProtoReflect.Descriptor instead.
func (*GemmBackendConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDescGZIP(), []int{1}
}

func (m *GemmBackendConfig) GetAlgorithm() isGemmBackendConfig_Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return nil
}

func (x *GemmBackendConfig) GetSelectedAlgorithm() int64 {
	if x, ok := x.GetAlgorithm().(*GemmBackendConfig_SelectedAlgorithm); ok {
		return x.SelectedAlgorithm
	}
	return 0
}

func (x *GemmBackendConfig) GetAlphaReal() float64 {
	if x != nil {
		return x.AlphaReal
	}
	return 0
}

func (x *GemmBackendConfig) GetAlphaImag() float64 {
	if x != nil {
		return x.AlphaImag
	}
	return 0
}

func (x *GemmBackendConfig) GetBeta() float64 {
	if x != nil {
		return x.Beta
	}
	return 0
}

func (x *GemmBackendConfig) GetDotDimensionNumbers() *xla.DotDimensionNumbers {
	if x != nil {
		return x.DotDimensionNumbers
	}
	return nil
}

func (x *GemmBackendConfig) GetPrecisionConfig() *xla.PrecisionConfig {
	if x != nil {
		return x.PrecisionConfig
	}
	return nil
}

func (x *GemmBackendConfig) GetEpilogue() GemmBackendConfig_Epilogue {
	if x != nil {
		return x.Epilogue
	}
	return GemmBackendConfig_DEFAULT
}

type isGemmBackendConfig_Algorithm interface {
	isGemmBackendConfig_Algorithm()
}

type GemmBackendConfig_SelectedAlgorithm struct {
	SelectedAlgorithm int64 `protobuf:"varint,1,opt,name=selected_algorithm,json=selectedAlgorithm,proto3,oneof"`
}

func (*GemmBackendConfig_SelectedAlgorithm) isGemmBackendConfig_Algorithm() {}

// Backend config for bitcast operation generated from MLIR MHLO dialect.
type BitcastBackendConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceLayout *xla.LayoutProto `protobuf:"bytes,1,opt,name=source_layout,json=sourceLayout,proto3" json:"source_layout,omitempty"`
	ResultLayout *xla.LayoutProto `protobuf:"bytes,2,opt,name=result_layout,json=resultLayout,proto3" json:"result_layout,omitempty"`
}

func (x *BitcastBackendConfig) Reset() {
	*x = BitcastBackendConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BitcastBackendConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BitcastBackendConfig) ProtoMessage() {}

func (x *BitcastBackendConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BitcastBackendConfig.ProtoReflect.Descriptor instead.
func (*BitcastBackendConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDescGZIP(), []int{2}
}

func (x *BitcastBackendConfig) GetSourceLayout() *xla.LayoutProto {
	if x != nil {
		return x.SourceLayout
	}
	return nil
}

func (x *BitcastBackendConfig) GetResultLayout() *xla.LayoutProto {
	if x != nil {
		return x.ResultLayout
	}
	return nil
}

var File_tensorflow_compiler_xla_service_gpu_backend_configs_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDesc = []byte{
	0x0a, 0x39, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x67, 0x70, 0x75, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x78, 0x6c, 0x61,
	0x2e, 0x67, 0x70, 0x75, 0x1a, 0x31, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x64, 0x6e,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61,
	0x2f, 0x78, 0x6c, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe6, 0x01, 0x0a, 0x16, 0x43, 0x75, 0x64, 0x6e, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x42, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x09, 0x61, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e,
	0x64, 0x6e, 0x6e, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x2a, 0x0a,
	0x11, 0x63, 0x6f, 0x6e, 0x76, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x73, 0x69,
	0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x4a, 0x04, 0x08, 0x01,
	0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0xe9, 0x03, 0x0a, 0x11, 0x47, 0x65, 0x6d,
	0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2f,
	0x0a, 0x12, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x11, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x72, 0x65, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x52, 0x65, 0x61, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x09, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x49, 0x6d, 0x61, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x62, 0x65, 0x74,
	0x61, 0x12, 0x4c, 0x0a, 0x15, 0x64, 0x6f, 0x74, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x44, 0x6f, 0x74, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x13, 0x64, 0x6f, 0x74, 0x44,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12,
	0x3f, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x78, 0x6c, 0x61, 0x2e,
	0x50, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0f, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x3f, 0x0a, 0x08, 0x65, 0x70, 0x69, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x23, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x67, 0x70, 0x75, 0x2e, 0x47, 0x65, 0x6d,
	0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45,
	0x70, 0x69, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x08, 0x65, 0x70, 0x69, 0x6c, 0x6f, 0x67, 0x75,
	0x65, 0x22, 0x74, 0x0a, 0x08, 0x45, 0x70, 0x69, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x49,
	0x41, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x4c, 0x55, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x42, 0x49, 0x41, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x55, 0x10, 0x03, 0x12, 0x08, 0x0a,
	0x04, 0x47, 0x45, 0x4c, 0x55, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x45, 0x4c, 0x55, 0x5f,
	0x41, 0x55, 0x58, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x49, 0x41, 0x53, 0x5f, 0x47, 0x45,
	0x4c, 0x55, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x49, 0x41, 0x53, 0x5f, 0x47, 0x45, 0x4c,
	0x55, 0x5f, 0x41, 0x55, 0x58, 0x10, 0x07, 0x42, 0x0b, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x22, 0x84, 0x01, 0x0a, 0x14, 0x42, 0x69, 0x74, 0x63, 0x61, 0x73, 0x74,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x35, 0x0a,
	0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x4c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x61,
	0x79, 0x6f, 0x75, 0x74, 0x12, 0x35, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6c,
	0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x78, 0x6c,
	0x61, 0x2e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x42, 0xa5, 0x01, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x67, 0x70, 0x75, 0x42, 0x13, 0x42, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x48, 0x02, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x71, 0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x70, 0x75, 0xa2, 0x02, 0x03, 0x58, 0x47, 0x58, 0xaa, 0x02,
	0x07, 0x58, 0x6c, 0x61, 0x2e, 0x47, 0x70, 0x75, 0xca, 0x02, 0x07, 0x58, 0x6c, 0x61, 0x5c, 0x47,
	0x70, 0x75, 0xe2, 0x02, 0x13, 0x58, 0x6c, 0x61, 0x5c, 0x47, 0x70, 0x75, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x58, 0x6c, 0x61, 0x3a, 0x3a,
	0x47, 0x70, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDescData = file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDesc
)

func file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDescData)
	})
	return file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDescData
}

var file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_goTypes = []interface{}{
	(GemmBackendConfig_Epilogue)(0), // 0: xla.gpu.GemmBackendConfig.Epilogue
	(*CudnnConvBackendConfig)(nil),  // 1: xla.gpu.CudnnConvBackendConfig
	(*GemmBackendConfig)(nil),       // 2: xla.gpu.GemmBackendConfig
	(*BitcastBackendConfig)(nil),    // 3: xla.gpu.BitcastBackendConfig
	(*protobuf.AlgorithmProto)(nil), // 4: stream_executor.dnn.AlgorithmProto
	(*xla.DotDimensionNumbers)(nil), // 5: xla.DotDimensionNumbers
	(*xla.PrecisionConfig)(nil),     // 6: xla.PrecisionConfig
	(*xla.LayoutProto)(nil),         // 7: xla.LayoutProto
}
var file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_depIdxs = []int32{
	4, // 0: xla.gpu.CudnnConvBackendConfig.algorithm:type_name -> stream_executor.dnn.AlgorithmProto
	5, // 1: xla.gpu.GemmBackendConfig.dot_dimension_numbers:type_name -> xla.DotDimensionNumbers
	6, // 2: xla.gpu.GemmBackendConfig.precision_config:type_name -> xla.PrecisionConfig
	0, // 3: xla.gpu.GemmBackendConfig.epilogue:type_name -> xla.gpu.GemmBackendConfig.Epilogue
	7, // 4: xla.gpu.BitcastBackendConfig.source_layout:type_name -> xla.LayoutProto
	7, // 5: xla.gpu.BitcastBackendConfig.result_layout:type_name -> xla.LayoutProto
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_init() }
func file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_init() {
	if File_tensorflow_compiler_xla_service_gpu_backend_configs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CudnnConvBackendConfig); i {
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
		file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GemmBackendConfig); i {
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
		file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BitcastBackendConfig); i {
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
	file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GemmBackendConfig_SelectedAlgorithm)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_depIdxs,
		EnumInfos:         file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_enumTypes,
		MessageInfos:      file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_xla_service_gpu_backend_configs_proto = out.File
	file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_rawDesc = nil
	file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_goTypes = nil
	file_tensorflow_compiler_xla_service_gpu_backend_configs_proto_depIdxs = nil
}
