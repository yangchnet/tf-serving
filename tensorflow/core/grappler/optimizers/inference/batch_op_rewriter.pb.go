// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow/core/grappler/optimizers/inference/batch_op_rewriter.proto

package inference

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Config for the batch op rewriter. This should be serialized
// and set a param in RewriterConfig with key kBatchOpRewriteParamKey.
type BatchOpRewriteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnableAdaptiveSharedBatchingThreadPool bool `protobuf:"varint,4,opt,name=enable_adaptive_shared_batching_thread_pool,json=enableAdaptiveSharedBatchingThreadPool,proto3" json:"enable_adaptive_shared_batching_thread_pool,omitempty"`
	// Keyed by model name, meaning all batch-ops in one saved model would use the
	// same adaptive-batch-scheduler option.
	ModelSchedulerOptions map[string]*BatchOpRewriteConfig_AdaptiveBatchSchedulerOption `protobuf:"bytes,1,rep,name=model_scheduler_options,json=modelSchedulerOptions,proto3" json:"model_scheduler_options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BatchOpRewriteConfig) Reset() {
	*x = BatchOpRewriteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOpRewriteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOpRewriteConfig) ProtoMessage() {}

func (x *BatchOpRewriteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOpRewriteConfig.ProtoReflect.Descriptor instead.
func (*BatchOpRewriteConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_rawDescGZIP(), []int{0}
}

func (x *BatchOpRewriteConfig) GetEnableAdaptiveSharedBatchingThreadPool() bool {
	if x != nil {
		return x.EnableAdaptiveSharedBatchingThreadPool
	}
	return false
}

func (x *BatchOpRewriteConfig) GetModelSchedulerOptions() map[string]*BatchOpRewriteConfig_AdaptiveBatchSchedulerOption {
	if x != nil {
		return x.ModelSchedulerOptions
	}
	return nil
}

// The options for tensorflow::serving::AdaptiveSharedBatchScheduler.
// See AdaptiveSharedBatchScheduler::Options for meaning of each field.
//
// NOTE:
// Leave this unset to pick up default settings which should work for most
// scenarios.
//
// Example scenarios when tuning helps:
// * Latency sensitive
type BatchOpRewriteConfig_AdaptiveBatchSchedulerOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinInflightBatchesLimit     *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=min_inflight_batches_limit,json=minInflightBatchesLimit,proto3" json:"min_inflight_batches_limit,omitempty"`
	InitialInflightBatchesLimit *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=initial_inflight_batches_limit,json=initialInflightBatchesLimit,proto3" json:"initial_inflight_batches_limit,omitempty"`
	MaxInflightBatchesLimit     *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=max_inflight_batches_limit,json=maxInflightBatchesLimit,proto3" json:"max_inflight_batches_limit,omitempty"`
	// You can use QPS as a reference to decide how quickly to react to workload
	// changes.
	BatchesToAverageOver *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=batches_to_average_over,json=batchesToAverageOver,proto3" json:"batches_to_average_over,omitempty"`
}

func (x *BatchOpRewriteConfig_AdaptiveBatchSchedulerOption) Reset() {
	*x = BatchOpRewriteConfig_AdaptiveBatchSchedulerOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOpRewriteConfig_AdaptiveBatchSchedulerOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOpRewriteConfig_AdaptiveBatchSchedulerOption) ProtoMessage() {}

func (x *BatchOpRewriteConfig_AdaptiveBatchSchedulerOption) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOpRewriteConfig_AdaptiveBatchSchedulerOption.ProtoReflect.Descriptor instead.
func (*BatchOpRewriteConfig_AdaptiveBatchSchedulerOption) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BatchOpRewriteConfig_AdaptiveBatchSchedulerOption) GetMinInflightBatchesLimit() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MinInflightBatchesLimit
	}
	return nil
}

func (x *BatchOpRewriteConfig_AdaptiveBatchSchedulerOption) GetInitialInflightBatchesLimit() *wrapperspb.UInt32Value {
	if x != nil {
		return x.InitialInflightBatchesLimit
	}
	return nil
}

func (x *BatchOpRewriteConfig_AdaptiveBatchSchedulerOption) GetMaxInflightBatchesLimit() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxInflightBatchesLimit
	}
	return nil
}

func (x *BatchOpRewriteConfig_AdaptiveBatchSchedulerOption) GetBatchesToAverageOver() *wrapperspb.UInt32Value {
	if x != nil {
		return x.BatchesToAverageOver
	}
	return nil
}

var File_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto protoreflect.FileDescriptor

var file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_rawDesc = []byte{
	0x0a, 0x45, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6d,
	0x69, 0x7a, 0x65, 0x72, 0x73, 0x2f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2f,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6f, 0x70, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x06, 0x0a, 0x14,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x5b, 0x0a, 0x2b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61,
	0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x70,
	0x6f, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x26, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x41, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x6f,
	0x6c, 0x12, 0x7b, 0x0a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x43, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x52,
	0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x15, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x8c,
	0x03, 0x0a, 0x1c, 0x41, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x59, 0x0a, 0x1a, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x17, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x61, 0x0a, 0x1e, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x1b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x59, 0x0a,
	0x1a, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x17, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x53, 0x0a, 0x17, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6f,
	0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x54, 0x6f, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x1a, 0x8f, 0x01,
	0x0a, 0x1a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x5b,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0xe7, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x42, 0x14, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x4f, 0x70, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x48, 0x02, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x71, 0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x72, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x2f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0xa2, 0x02, 0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x12, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x12,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0xe2, 0x02, 0x1e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_rawDescOnce sync.Once
	file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_rawDescData = file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_rawDesc
)

func file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_rawDescGZIP() []byte {
	file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_rawDescData)
	})
	return file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_rawDescData
}

var file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_goTypes = []interface{}{
	(*BatchOpRewriteConfig)(nil),                              // 0: tensorflow.serving.BatchOpRewriteConfig
	(*BatchOpRewriteConfig_AdaptiveBatchSchedulerOption)(nil), // 1: tensorflow.serving.BatchOpRewriteConfig.AdaptiveBatchSchedulerOption
	nil,                            // 2: tensorflow.serving.BatchOpRewriteConfig.ModelSchedulerOptionsEntry
	(*wrapperspb.UInt32Value)(nil), // 3: google.protobuf.UInt32Value
}
var file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_depIdxs = []int32{
	2, // 0: tensorflow.serving.BatchOpRewriteConfig.model_scheduler_options:type_name -> tensorflow.serving.BatchOpRewriteConfig.ModelSchedulerOptionsEntry
	3, // 1: tensorflow.serving.BatchOpRewriteConfig.AdaptiveBatchSchedulerOption.min_inflight_batches_limit:type_name -> google.protobuf.UInt32Value
	3, // 2: tensorflow.serving.BatchOpRewriteConfig.AdaptiveBatchSchedulerOption.initial_inflight_batches_limit:type_name -> google.protobuf.UInt32Value
	3, // 3: tensorflow.serving.BatchOpRewriteConfig.AdaptiveBatchSchedulerOption.max_inflight_batches_limit:type_name -> google.protobuf.UInt32Value
	3, // 4: tensorflow.serving.BatchOpRewriteConfig.AdaptiveBatchSchedulerOption.batches_to_average_over:type_name -> google.protobuf.UInt32Value
	1, // 5: tensorflow.serving.BatchOpRewriteConfig.ModelSchedulerOptionsEntry.value:type_name -> tensorflow.serving.BatchOpRewriteConfig.AdaptiveBatchSchedulerOption
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_init() }
func file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_init() {
	if File_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOpRewriteConfig); i {
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
		file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOpRewriteConfig_AdaptiveBatchSchedulerOption); i {
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
			RawDescriptor: file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_msgTypes,
	}.Build()
	File_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto = out.File
	file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_rawDesc = nil
	file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_goTypes = nil
	file_tensorflow_core_grappler_optimizers_inference_batch_op_rewriter_proto_depIdxs = nil
}
