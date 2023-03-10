// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow/core/data/service/common.proto

package service

import (
	framework "github.com/yangchnet/tf-serving/tensorflow/core/framework"
	protobuf "github.com/yangchnet/tf-serving/tensorflow/core/protobuf"
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

// Specifies which tf.data service workers to read from.
type TargetWorkers int32

const (
	TargetWorkers_TARGET_WORKERS_UNSPECIFIED TargetWorkers = 0
	// tf.data service runtime decides which workers to read from.
	TargetWorkers_TARGET_WORKERS_AUTO TargetWorkers = 1
	// Reads from any available worker.
	TargetWorkers_TARGET_WORKERS_ANY TargetWorkers = 2
	// Only reads from local workers. If no local worker is found, it is an error.
	TargetWorkers_TARGET_WORKERS_LOCAL TargetWorkers = 3
)

// Enum value maps for TargetWorkers.
var (
	TargetWorkers_name = map[int32]string{
		0: "TARGET_WORKERS_UNSPECIFIED",
		1: "TARGET_WORKERS_AUTO",
		2: "TARGET_WORKERS_ANY",
		3: "TARGET_WORKERS_LOCAL",
	}
	TargetWorkers_value = map[string]int32{
		"TARGET_WORKERS_UNSPECIFIED": 0,
		"TARGET_WORKERS_AUTO":        1,
		"TARGET_WORKERS_ANY":         2,
		"TARGET_WORKERS_LOCAL":       3,
	}
)

func (x TargetWorkers) Enum() *TargetWorkers {
	p := new(TargetWorkers)
	*p = x
	return p
}

func (x TargetWorkers) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TargetWorkers) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_data_service_common_proto_enumTypes[0].Descriptor()
}

func (TargetWorkers) Type() protoreflect.EnumType {
	return &file_tensorflow_core_data_service_common_proto_enumTypes[0]
}

func (x TargetWorkers) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TargetWorkers.Descriptor instead.
func (TargetWorkers) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_data_service_common_proto_rawDescGZIP(), []int{0}
}

// Next tag: 2
type DatasetDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// We represent datasets as tensorflow GraphDefs which define the operations
	// needed to create a tf.data dataset.
	Graph *framework.GraphDef `protobuf:"bytes,1,opt,name=graph,proto3" json:"graph,omitempty"`
}

func (x *DatasetDef) Reset() {
	*x = DatasetDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_data_service_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetDef) ProtoMessage() {}

func (x *DatasetDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_data_service_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetDef.ProtoReflect.Descriptor instead.
func (*DatasetDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_data_service_common_proto_rawDescGZIP(), []int{0}
}

func (x *DatasetDef) GetGraph() *framework.GraphDef {
	if x != nil {
		return x.Graph
	}
	return nil
}

// Next tag: 3
type IterationKeyDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Iteration int64  `protobuf:"varint,2,opt,name=iteration,proto3" json:"iteration,omitempty"`
}

func (x *IterationKeyDef) Reset() {
	*x = IterationKeyDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_data_service_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IterationKeyDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IterationKeyDef) ProtoMessage() {}

func (x *IterationKeyDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_data_service_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IterationKeyDef.ProtoReflect.Descriptor instead.
func (*IterationKeyDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_data_service_common_proto_rawDescGZIP(), []int{1}
}

func (x *IterationKeyDef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IterationKeyDef) GetIteration() int64 {
	if x != nil {
		return x.Iteration
	}
	return 0
}

// Next tag: 14
type TaskDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The dataset to iterate over.
	//
	// Types that are assignable to Dataset:
	//
	//	*TaskDef_DatasetDef
	//	*TaskDef_Path
	Dataset     isTaskDef_Dataset `protobuf_oneof:"dataset"`
	DatasetId   string            `protobuf:"bytes,3,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	TaskId      int64             `protobuf:"varint,4,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	IterationId int64             `protobuf:"varint,5,opt,name=iteration_id,json=iterationId,proto3" json:"iteration_id,omitempty"`
	// In distributed epoch processing mode, we use one split provider for each
	// source that feeds into the dataset. In parallel_epochs mode,
	// `num_split_providers` is always zero.
	NumSplitProviders int64 `protobuf:"varint,9,opt,name=num_split_providers,json=numSplitProviders,proto3" json:"num_split_providers,omitempty"`
	// Address of the worker that the task is assigned to.
	WorkerAddress     string                      `protobuf:"bytes,8,opt,name=worker_address,json=workerAddress,proto3" json:"worker_address,omitempty"`
	ProcessingModeDef *protobuf.ProcessingModeDef `protobuf:"bytes,10,opt,name=processing_mode_def,json=processingModeDef,proto3" json:"processing_mode_def,omitempty"`
	// Optional number of consumers. If set, the results of the task will be
	// provided to consumers round-robin.
	//
	// Types that are assignable to OptionalNumConsumers:
	//
	//	*TaskDef_NumConsumers
	OptionalNumConsumers isTaskDef_OptionalNumConsumers `protobuf_oneof:"optional_num_consumers"`
	// Number of workers and the worker index. These are only populated when the
	// `processing_mode_def` specifies a static sharding policy.
	NumWorkers  int64 `protobuf:"varint,11,opt,name=num_workers,json=numWorkers,proto3" json:"num_workers,omitempty"`
	WorkerIndex int64 `protobuf:"varint,12,opt,name=worker_index,json=workerIndex,proto3" json:"worker_index,omitempty"`
	// True if cross-trainer cache is enabled.
	UseCrossTrainerCache bool `protobuf:"varint,13,opt,name=use_cross_trainer_cache,json=useCrossTrainerCache,proto3" json:"use_cross_trainer_cache,omitempty"`
}

func (x *TaskDef) Reset() {
	*x = TaskDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_data_service_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDef) ProtoMessage() {}

func (x *TaskDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_data_service_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDef.ProtoReflect.Descriptor instead.
func (*TaskDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_data_service_common_proto_rawDescGZIP(), []int{2}
}

func (m *TaskDef) GetDataset() isTaskDef_Dataset {
	if m != nil {
		return m.Dataset
	}
	return nil
}

func (x *TaskDef) GetDatasetDef() *DatasetDef {
	if x, ok := x.GetDataset().(*TaskDef_DatasetDef); ok {
		return x.DatasetDef
	}
	return nil
}

func (x *TaskDef) GetPath() string {
	if x, ok := x.GetDataset().(*TaskDef_Path); ok {
		return x.Path
	}
	return ""
}

func (x *TaskDef) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

func (x *TaskDef) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskDef) GetIterationId() int64 {
	if x != nil {
		return x.IterationId
	}
	return 0
}

func (x *TaskDef) GetNumSplitProviders() int64 {
	if x != nil {
		return x.NumSplitProviders
	}
	return 0
}

func (x *TaskDef) GetWorkerAddress() string {
	if x != nil {
		return x.WorkerAddress
	}
	return ""
}

func (x *TaskDef) GetProcessingModeDef() *protobuf.ProcessingModeDef {
	if x != nil {
		return x.ProcessingModeDef
	}
	return nil
}

func (m *TaskDef) GetOptionalNumConsumers() isTaskDef_OptionalNumConsumers {
	if m != nil {
		return m.OptionalNumConsumers
	}
	return nil
}

func (x *TaskDef) GetNumConsumers() int64 {
	if x, ok := x.GetOptionalNumConsumers().(*TaskDef_NumConsumers); ok {
		return x.NumConsumers
	}
	return 0
}

func (x *TaskDef) GetNumWorkers() int64 {
	if x != nil {
		return x.NumWorkers
	}
	return 0
}

func (x *TaskDef) GetWorkerIndex() int64 {
	if x != nil {
		return x.WorkerIndex
	}
	return 0
}

func (x *TaskDef) GetUseCrossTrainerCache() bool {
	if x != nil {
		return x.UseCrossTrainerCache
	}
	return false
}

type isTaskDef_Dataset interface {
	isTaskDef_Dataset()
}

type TaskDef_DatasetDef struct {
	DatasetDef *DatasetDef `protobuf:"bytes,1,opt,name=dataset_def,json=datasetDef,proto3,oneof"`
}

type TaskDef_Path struct {
	Path string `protobuf:"bytes,2,opt,name=path,proto3,oneof"`
}

func (*TaskDef_DatasetDef) isTaskDef_Dataset() {}

func (*TaskDef_Path) isTaskDef_Dataset() {}

type isTaskDef_OptionalNumConsumers interface {
	isTaskDef_OptionalNumConsumers()
}

type TaskDef_NumConsumers struct {
	NumConsumers int64 `protobuf:"varint,7,opt,name=num_consumers,json=numConsumers,proto3,oneof"`
}

func (*TaskDef_NumConsumers) isTaskDef_OptionalNumConsumers() {}

// Next tag: 8
type TaskInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the worker processing the task.
	WorkerAddress string `protobuf:"bytes,1,opt,name=worker_address,json=workerAddress,proto3" json:"worker_address,omitempty"`
	// The transfer address of the worker processing the task.
	TransferAddress string `protobuf:"bytes,4,opt,name=transfer_address,json=transferAddress,proto3" json:"transfer_address,omitempty"`
	// Tags attached to the worker. This allows reading from selected workers.
	// For example, by applying a "COLOCATED" tag, tf.data service is able to read
	// from the local tf.data worker if one exists, then from off-TF-host workers,
	// to avoid cross-TF-host reads.
	WorkerTags []string `protobuf:"bytes,6,rep,name=worker_tags,json=workerTags,proto3" json:"worker_tags,omitempty"`
	// The task id.
	TaskId int64 `protobuf:"varint,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// The id of the iteration that the task is part of.
	IterationId int64 `protobuf:"varint,3,opt,name=iteration_id,json=iterationId,proto3" json:"iteration_id,omitempty"`
	// The UID of the worker Borg job, used for telemetry.
	WorkerUid int64 `protobuf:"varint,7,opt,name=worker_uid,json=workerUid,proto3" json:"worker_uid,omitempty"`
	// The round to start reading from the task in. For non-round-robin reads,
	// this is always 0.
	StartingRound int64 `protobuf:"varint,5,opt,name=starting_round,json=startingRound,proto3" json:"starting_round,omitempty"`
}

func (x *TaskInfo) Reset() {
	*x = TaskInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_data_service_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInfo) ProtoMessage() {}

func (x *TaskInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_data_service_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInfo.ProtoReflect.Descriptor instead.
func (*TaskInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_data_service_common_proto_rawDescGZIP(), []int{3}
}

func (x *TaskInfo) GetWorkerAddress() string {
	if x != nil {
		return x.WorkerAddress
	}
	return ""
}

func (x *TaskInfo) GetTransferAddress() string {
	if x != nil {
		return x.TransferAddress
	}
	return ""
}

func (x *TaskInfo) GetWorkerTags() []string {
	if x != nil {
		return x.WorkerTags
	}
	return nil
}

func (x *TaskInfo) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskInfo) GetIterationId() int64 {
	if x != nil {
		return x.IterationId
	}
	return 0
}

func (x *TaskInfo) GetWorkerUid() int64 {
	if x != nil {
		return x.WorkerUid
	}
	return 0
}

func (x *TaskInfo) GetStartingRound() int64 {
	if x != nil {
		return x.StartingRound
	}
	return 0
}

var File_tensorflow_core_data_service_common_proto protoreflect.FileDescriptor

var file_tensorflow_core_data_service_common_proto_rawDesc = []byte{
	0x0a, 0x29, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x25, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x38, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x44, 0x65, 0x66, 0x12, 0x2a,
	0x0a, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x44, 0x65, 0x66, 0x52, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x22, 0x43, 0x0a, 0x0f, 0x49, 0x74,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x66, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xb2, 0x04, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x66, 0x12, 0x3e, 0x0a, 0x0b, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x44, 0x65, 0x66, 0x48, 0x00, 0x52,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x44, 0x65, 0x66, 0x12, 0x14, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x74, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13,
	0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6e, 0x75, 0x6d, 0x53, 0x70,
	0x6c, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x52, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x44, 0x65, 0x66, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x4d, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66, 0x12, 0x25, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x5f, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01,
	0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x35, 0x0a, 0x17, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x5f,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x14, 0x75, 0x73, 0x65, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x63, 0x68, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x42, 0x18, 0x0a, 0x16, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x4a, 0x04,
	0x08, 0x06, 0x10, 0x07, 0x22, 0xff, 0x01, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x54, 0x61, 0x67, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x2a, 0x7a, 0x0a, 0x0d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x41, 0x52, 0x47, 0x45,
	0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x41, 0x52, 0x47, 0x45,
	0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x53, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x10, 0x01,
	0x12, 0x16, 0x0a, 0x12, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x45,
	0x52, 0x53, 0x5f, 0x41, 0x4e, 0x59, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x41, 0x52, 0x47,
	0x45, 0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x53, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c,
	0x10, 0x03, 0x42, 0xbe, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70,
	0x2f, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x03, 0x54, 0x44, 0x58, 0xaa, 0x02,
	0x0f, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0xca, 0x02, 0x0f, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x44, 0x61,
	0x74, 0x61, 0xe2, 0x02, 0x1b, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c,
	0x44, 0x61, 0x74, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x10, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_data_service_common_proto_rawDescOnce sync.Once
	file_tensorflow_core_data_service_common_proto_rawDescData = file_tensorflow_core_data_service_common_proto_rawDesc
)

func file_tensorflow_core_data_service_common_proto_rawDescGZIP() []byte {
	file_tensorflow_core_data_service_common_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_data_service_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_data_service_common_proto_rawDescData)
	})
	return file_tensorflow_core_data_service_common_proto_rawDescData
}

var file_tensorflow_core_data_service_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_data_service_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_core_data_service_common_proto_goTypes = []interface{}{
	(TargetWorkers)(0),                 // 0: tensorflow.data.TargetWorkers
	(*DatasetDef)(nil),                 // 1: tensorflow.data.DatasetDef
	(*IterationKeyDef)(nil),            // 2: tensorflow.data.IterationKeyDef
	(*TaskDef)(nil),                    // 3: tensorflow.data.TaskDef
	(*TaskInfo)(nil),                   // 4: tensorflow.data.TaskInfo
	(*framework.GraphDef)(nil),         // 5: tensorflow.GraphDef
	(*protobuf.ProcessingModeDef)(nil), // 6: tensorflow.data.ProcessingModeDef
}
var file_tensorflow_core_data_service_common_proto_depIdxs = []int32{
	5, // 0: tensorflow.data.DatasetDef.graph:type_name -> tensorflow.GraphDef
	1, // 1: tensorflow.data.TaskDef.dataset_def:type_name -> tensorflow.data.DatasetDef
	6, // 2: tensorflow.data.TaskDef.processing_mode_def:type_name -> tensorflow.data.ProcessingModeDef
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_core_data_service_common_proto_init() }
func file_tensorflow_core_data_service_common_proto_init() {
	if File_tensorflow_core_data_service_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_data_service_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetDef); i {
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
		file_tensorflow_core_data_service_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IterationKeyDef); i {
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
		file_tensorflow_core_data_service_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDef); i {
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
		file_tensorflow_core_data_service_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskInfo); i {
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
	file_tensorflow_core_data_service_common_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*TaskDef_DatasetDef)(nil),
		(*TaskDef_Path)(nil),
		(*TaskDef_NumConsumers)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_data_service_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_data_service_common_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_data_service_common_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_data_service_common_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_data_service_common_proto_msgTypes,
	}.Build()
	File_tensorflow_core_data_service_common_proto = out.File
	file_tensorflow_core_data_service_common_proto_rawDesc = nil
	file_tensorflow_core_data_service_common_proto_goTypes = nil
	file_tensorflow_core_data_service_common_proto_depIdxs = nil
}
