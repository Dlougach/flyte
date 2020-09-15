// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/matchable_resource.proto

package admin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Defines a resource that can be configured by customizable Project-, ProjectDomain- or WorkflowAttributes
// based on matching tags.
type MatchableResource int32

const (
	// Applies to customizable task resource requests and limits.
	MatchableResource_TASK_RESOURCE MatchableResource = 0
	// Applies to configuring templated kubernetes cluster resources.
	MatchableResource_CLUSTER_RESOURCE MatchableResource = 1
	// Configures task and dynamic task execution queue assignment.
	MatchableResource_EXECUTION_QUEUE MatchableResource = 2
	// Configures the K8s cluster label to be used for execution to be run
	MatchableResource_EXECUTION_CLUSTER_LABEL MatchableResource = 3
	// Configures default quality of service when undefined in an execution spec.
	MatchableResource_QUALITY_OF_SERVICE_SPECIFICATION MatchableResource = 4
	// Selects configurable plugin implementation behavior for a given task type.
	MatchableResource_PLUGIN_OVERRIDE MatchableResource = 5
)

var MatchableResource_name = map[int32]string{
	0: "TASK_RESOURCE",
	1: "CLUSTER_RESOURCE",
	2: "EXECUTION_QUEUE",
	3: "EXECUTION_CLUSTER_LABEL",
	4: "QUALITY_OF_SERVICE_SPECIFICATION",
	5: "PLUGIN_OVERRIDE",
}

var MatchableResource_value = map[string]int32{
	"TASK_RESOURCE":                    0,
	"CLUSTER_RESOURCE":                 1,
	"EXECUTION_QUEUE":                  2,
	"EXECUTION_CLUSTER_LABEL":          3,
	"QUALITY_OF_SERVICE_SPECIFICATION": 4,
	"PLUGIN_OVERRIDE":                  5,
}

func (x MatchableResource) String() string {
	return proto.EnumName(MatchableResource_name, int32(x))
}

func (MatchableResource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{0}
}

type PluginOverride_MissingPluginBehavior int32

const (
	PluginOverride_FAIL PluginOverride_MissingPluginBehavior = 0
	// Uses the system-configured default implementation.
	PluginOverride_USE_DEFAULT PluginOverride_MissingPluginBehavior = 1
)

var PluginOverride_MissingPluginBehavior_name = map[int32]string{
	0: "FAIL",
	1: "USE_DEFAULT",
}

var PluginOverride_MissingPluginBehavior_value = map[string]int32{
	"FAIL":        0,
	"USE_DEFAULT": 1,
}

func (x PluginOverride_MissingPluginBehavior) String() string {
	return proto.EnumName(PluginOverride_MissingPluginBehavior_name, int32(x))
}

func (PluginOverride_MissingPluginBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{5, 0}
}

type TaskResourceSpec struct {
	Cpu                  string   `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Gpu                  string   `protobuf:"bytes,2,opt,name=gpu,proto3" json:"gpu,omitempty"`
	Memory               string   `protobuf:"bytes,3,opt,name=memory,proto3" json:"memory,omitempty"`
	Storage              string   `protobuf:"bytes,4,opt,name=storage,proto3" json:"storage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResourceSpec) Reset()         { *m = TaskResourceSpec{} }
func (m *TaskResourceSpec) String() string { return proto.CompactTextString(m) }
func (*TaskResourceSpec) ProtoMessage()    {}
func (*TaskResourceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{0}
}

func (m *TaskResourceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResourceSpec.Unmarshal(m, b)
}
func (m *TaskResourceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResourceSpec.Marshal(b, m, deterministic)
}
func (m *TaskResourceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResourceSpec.Merge(m, src)
}
func (m *TaskResourceSpec) XXX_Size() int {
	return xxx_messageInfo_TaskResourceSpec.Size(m)
}
func (m *TaskResourceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResourceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResourceSpec proto.InternalMessageInfo

func (m *TaskResourceSpec) GetCpu() string {
	if m != nil {
		return m.Cpu
	}
	return ""
}

func (m *TaskResourceSpec) GetGpu() string {
	if m != nil {
		return m.Gpu
	}
	return ""
}

func (m *TaskResourceSpec) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func (m *TaskResourceSpec) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

type TaskResourceAttributes struct {
	Defaults             *TaskResourceSpec `protobuf:"bytes,1,opt,name=defaults,proto3" json:"defaults,omitempty"`
	Limits               *TaskResourceSpec `protobuf:"bytes,2,opt,name=limits,proto3" json:"limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskResourceAttributes) Reset()         { *m = TaskResourceAttributes{} }
func (m *TaskResourceAttributes) String() string { return proto.CompactTextString(m) }
func (*TaskResourceAttributes) ProtoMessage()    {}
func (*TaskResourceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{1}
}

func (m *TaskResourceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResourceAttributes.Unmarshal(m, b)
}
func (m *TaskResourceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResourceAttributes.Marshal(b, m, deterministic)
}
func (m *TaskResourceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResourceAttributes.Merge(m, src)
}
func (m *TaskResourceAttributes) XXX_Size() int {
	return xxx_messageInfo_TaskResourceAttributes.Size(m)
}
func (m *TaskResourceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResourceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResourceAttributes proto.InternalMessageInfo

func (m *TaskResourceAttributes) GetDefaults() *TaskResourceSpec {
	if m != nil {
		return m.Defaults
	}
	return nil
}

func (m *TaskResourceAttributes) GetLimits() *TaskResourceSpec {
	if m != nil {
		return m.Limits
	}
	return nil
}

type ClusterResourceAttributes struct {
	// Custom resource attributes which will be applied in cluster resource creation (e.g. quotas).
	// Map keys are the *case-sensitive* names of variables in templatized resource files.
	// Map values should be the custom values which get substituted during resource creation.
	Attributes           map[string]string `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClusterResourceAttributes) Reset()         { *m = ClusterResourceAttributes{} }
func (m *ClusterResourceAttributes) String() string { return proto.CompactTextString(m) }
func (*ClusterResourceAttributes) ProtoMessage()    {}
func (*ClusterResourceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{2}
}

func (m *ClusterResourceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterResourceAttributes.Unmarshal(m, b)
}
func (m *ClusterResourceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterResourceAttributes.Marshal(b, m, deterministic)
}
func (m *ClusterResourceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterResourceAttributes.Merge(m, src)
}
func (m *ClusterResourceAttributes) XXX_Size() int {
	return xxx_messageInfo_ClusterResourceAttributes.Size(m)
}
func (m *ClusterResourceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterResourceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterResourceAttributes proto.InternalMessageInfo

func (m *ClusterResourceAttributes) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type ExecutionQueueAttributes struct {
	// Tags used for assigning execution queues for tasks defined within this project.
	Tags                 []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionQueueAttributes) Reset()         { *m = ExecutionQueueAttributes{} }
func (m *ExecutionQueueAttributes) String() string { return proto.CompactTextString(m) }
func (*ExecutionQueueAttributes) ProtoMessage()    {}
func (*ExecutionQueueAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{3}
}

func (m *ExecutionQueueAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionQueueAttributes.Unmarshal(m, b)
}
func (m *ExecutionQueueAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionQueueAttributes.Marshal(b, m, deterministic)
}
func (m *ExecutionQueueAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionQueueAttributes.Merge(m, src)
}
func (m *ExecutionQueueAttributes) XXX_Size() int {
	return xxx_messageInfo_ExecutionQueueAttributes.Size(m)
}
func (m *ExecutionQueueAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionQueueAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionQueueAttributes proto.InternalMessageInfo

func (m *ExecutionQueueAttributes) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ExecutionClusterLabel struct {
	// Label value to determine where the execution will be run
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionClusterLabel) Reset()         { *m = ExecutionClusterLabel{} }
func (m *ExecutionClusterLabel) String() string { return proto.CompactTextString(m) }
func (*ExecutionClusterLabel) ProtoMessage()    {}
func (*ExecutionClusterLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{4}
}

func (m *ExecutionClusterLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionClusterLabel.Unmarshal(m, b)
}
func (m *ExecutionClusterLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionClusterLabel.Marshal(b, m, deterministic)
}
func (m *ExecutionClusterLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionClusterLabel.Merge(m, src)
}
func (m *ExecutionClusterLabel) XXX_Size() int {
	return xxx_messageInfo_ExecutionClusterLabel.Size(m)
}
func (m *ExecutionClusterLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionClusterLabel.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionClusterLabel proto.InternalMessageInfo

func (m *ExecutionClusterLabel) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// This MatchableAttribute configures selecting alternate plugin implementations for a given task type.
// In addition to an override implementation a selection of fallbacks can be provided or other modes
// for handling cases where the desired plugin override is not enabled in a given Flyte deployment.
type PluginOverride struct {
	// A predefined yet extensible Task type identifier.
	TaskType string `protobuf:"bytes,1,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// A set of plugin ids which should handle tasks of this type instead of the default registered plugin. The list will be tried in order until a plugin is found with that id.
	PluginId []string `protobuf:"bytes,2,rep,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	// Defines the behavior when no plugin from the plugin_id list is not found.
	MissingPluginBehavior PluginOverride_MissingPluginBehavior `protobuf:"varint,4,opt,name=missing_plugin_behavior,json=missingPluginBehavior,proto3,enum=flyteidl.admin.PluginOverride_MissingPluginBehavior" json:"missing_plugin_behavior,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                             `json:"-"`
	XXX_unrecognized      []byte                               `json:"-"`
	XXX_sizecache         int32                                `json:"-"`
}

func (m *PluginOverride) Reset()         { *m = PluginOverride{} }
func (m *PluginOverride) String() string { return proto.CompactTextString(m) }
func (*PluginOverride) ProtoMessage()    {}
func (*PluginOverride) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{5}
}

func (m *PluginOverride) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginOverride.Unmarshal(m, b)
}
func (m *PluginOverride) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginOverride.Marshal(b, m, deterministic)
}
func (m *PluginOverride) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginOverride.Merge(m, src)
}
func (m *PluginOverride) XXX_Size() int {
	return xxx_messageInfo_PluginOverride.Size(m)
}
func (m *PluginOverride) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginOverride.DiscardUnknown(m)
}

var xxx_messageInfo_PluginOverride proto.InternalMessageInfo

func (m *PluginOverride) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *PluginOverride) GetPluginId() []string {
	if m != nil {
		return m.PluginId
	}
	return nil
}

func (m *PluginOverride) GetMissingPluginBehavior() PluginOverride_MissingPluginBehavior {
	if m != nil {
		return m.MissingPluginBehavior
	}
	return PluginOverride_FAIL
}

// Generic container for encapsulating all types of the above attributes messages.
type MatchingAttributes struct {
	// Types that are valid to be assigned to Target:
	//	*MatchingAttributes_TaskResourceAttributes
	//	*MatchingAttributes_ClusterResourceAttributes
	//	*MatchingAttributes_ExecutionQueueAttributes
	//	*MatchingAttributes_ExecutionClusterLabel
	//	*MatchingAttributes_QualityOfService
	//	*MatchingAttributes_PluginOverride
	Target               isMatchingAttributes_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MatchingAttributes) Reset()         { *m = MatchingAttributes{} }
func (m *MatchingAttributes) String() string { return proto.CompactTextString(m) }
func (*MatchingAttributes) ProtoMessage()    {}
func (*MatchingAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{6}
}

func (m *MatchingAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchingAttributes.Unmarshal(m, b)
}
func (m *MatchingAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchingAttributes.Marshal(b, m, deterministic)
}
func (m *MatchingAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchingAttributes.Merge(m, src)
}
func (m *MatchingAttributes) XXX_Size() int {
	return xxx_messageInfo_MatchingAttributes.Size(m)
}
func (m *MatchingAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchingAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_MatchingAttributes proto.InternalMessageInfo

type isMatchingAttributes_Target interface {
	isMatchingAttributes_Target()
}

type MatchingAttributes_TaskResourceAttributes struct {
	TaskResourceAttributes *TaskResourceAttributes `protobuf:"bytes,1,opt,name=task_resource_attributes,json=taskResourceAttributes,proto3,oneof"`
}

type MatchingAttributes_ClusterResourceAttributes struct {
	ClusterResourceAttributes *ClusterResourceAttributes `protobuf:"bytes,2,opt,name=cluster_resource_attributes,json=clusterResourceAttributes,proto3,oneof"`
}

type MatchingAttributes_ExecutionQueueAttributes struct {
	ExecutionQueueAttributes *ExecutionQueueAttributes `protobuf:"bytes,3,opt,name=execution_queue_attributes,json=executionQueueAttributes,proto3,oneof"`
}

type MatchingAttributes_ExecutionClusterLabel struct {
	ExecutionClusterLabel *ExecutionClusterLabel `protobuf:"bytes,4,opt,name=execution_cluster_label,json=executionClusterLabel,proto3,oneof"`
}

type MatchingAttributes_QualityOfService struct {
	QualityOfService *core.QualityOfService `protobuf:"bytes,5,opt,name=quality_of_service,json=qualityOfService,proto3,oneof"`
}

type MatchingAttributes_PluginOverride struct {
	PluginOverride *PluginOverride `protobuf:"bytes,6,opt,name=plugin_override,json=pluginOverride,proto3,oneof"`
}

func (*MatchingAttributes_TaskResourceAttributes) isMatchingAttributes_Target() {}

func (*MatchingAttributes_ClusterResourceAttributes) isMatchingAttributes_Target() {}

func (*MatchingAttributes_ExecutionQueueAttributes) isMatchingAttributes_Target() {}

func (*MatchingAttributes_ExecutionClusterLabel) isMatchingAttributes_Target() {}

func (*MatchingAttributes_QualityOfService) isMatchingAttributes_Target() {}

func (*MatchingAttributes_PluginOverride) isMatchingAttributes_Target() {}

func (m *MatchingAttributes) GetTarget() isMatchingAttributes_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *MatchingAttributes) GetTaskResourceAttributes() *TaskResourceAttributes {
	if x, ok := m.GetTarget().(*MatchingAttributes_TaskResourceAttributes); ok {
		return x.TaskResourceAttributes
	}
	return nil
}

func (m *MatchingAttributes) GetClusterResourceAttributes() *ClusterResourceAttributes {
	if x, ok := m.GetTarget().(*MatchingAttributes_ClusterResourceAttributes); ok {
		return x.ClusterResourceAttributes
	}
	return nil
}

func (m *MatchingAttributes) GetExecutionQueueAttributes() *ExecutionQueueAttributes {
	if x, ok := m.GetTarget().(*MatchingAttributes_ExecutionQueueAttributes); ok {
		return x.ExecutionQueueAttributes
	}
	return nil
}

func (m *MatchingAttributes) GetExecutionClusterLabel() *ExecutionClusterLabel {
	if x, ok := m.GetTarget().(*MatchingAttributes_ExecutionClusterLabel); ok {
		return x.ExecutionClusterLabel
	}
	return nil
}

func (m *MatchingAttributes) GetQualityOfService() *core.QualityOfService {
	if x, ok := m.GetTarget().(*MatchingAttributes_QualityOfService); ok {
		return x.QualityOfService
	}
	return nil
}

func (m *MatchingAttributes) GetPluginOverride() *PluginOverride {
	if x, ok := m.GetTarget().(*MatchingAttributes_PluginOverride); ok {
		return x.PluginOverride
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MatchingAttributes) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MatchingAttributes_TaskResourceAttributes)(nil),
		(*MatchingAttributes_ClusterResourceAttributes)(nil),
		(*MatchingAttributes_ExecutionQueueAttributes)(nil),
		(*MatchingAttributes_ExecutionClusterLabel)(nil),
		(*MatchingAttributes_QualityOfService)(nil),
		(*MatchingAttributes_PluginOverride)(nil),
	}
}

// Represents a custom set of attributes applied for either a domain; a domain and project; or
// domain, project and workflow name.
type MatchableAttributesConfiguration struct {
	Attributes           *MatchingAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	Domain               string              `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Project              string              `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Workflow             string              `protobuf:"bytes,4,opt,name=workflow,proto3" json:"workflow,omitempty"`
	LaunchPlan           string              `protobuf:"bytes,5,opt,name=launch_plan,json=launchPlan,proto3" json:"launch_plan,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MatchableAttributesConfiguration) Reset()         { *m = MatchableAttributesConfiguration{} }
func (m *MatchableAttributesConfiguration) String() string { return proto.CompactTextString(m) }
func (*MatchableAttributesConfiguration) ProtoMessage()    {}
func (*MatchableAttributesConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{7}
}

func (m *MatchableAttributesConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchableAttributesConfiguration.Unmarshal(m, b)
}
func (m *MatchableAttributesConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchableAttributesConfiguration.Marshal(b, m, deterministic)
}
func (m *MatchableAttributesConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchableAttributesConfiguration.Merge(m, src)
}
func (m *MatchableAttributesConfiguration) XXX_Size() int {
	return xxx_messageInfo_MatchableAttributesConfiguration.Size(m)
}
func (m *MatchableAttributesConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchableAttributesConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_MatchableAttributesConfiguration proto.InternalMessageInfo

func (m *MatchableAttributesConfiguration) GetAttributes() *MatchingAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *MatchableAttributesConfiguration) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *MatchableAttributesConfiguration) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *MatchableAttributesConfiguration) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

func (m *MatchableAttributesConfiguration) GetLaunchPlan() string {
	if m != nil {
		return m.LaunchPlan
	}
	return ""
}

// Request all matching resource attributes.
type ListMatchableAttributesRequest struct {
	ResourceType         MatchableResource `protobuf:"varint,1,opt,name=resource_type,json=resourceType,proto3,enum=flyteidl.admin.MatchableResource" json:"resource_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListMatchableAttributesRequest) Reset()         { *m = ListMatchableAttributesRequest{} }
func (m *ListMatchableAttributesRequest) String() string { return proto.CompactTextString(m) }
func (*ListMatchableAttributesRequest) ProtoMessage()    {}
func (*ListMatchableAttributesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{8}
}

func (m *ListMatchableAttributesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMatchableAttributesRequest.Unmarshal(m, b)
}
func (m *ListMatchableAttributesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMatchableAttributesRequest.Marshal(b, m, deterministic)
}
func (m *ListMatchableAttributesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMatchableAttributesRequest.Merge(m, src)
}
func (m *ListMatchableAttributesRequest) XXX_Size() int {
	return xxx_messageInfo_ListMatchableAttributesRequest.Size(m)
}
func (m *ListMatchableAttributesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMatchableAttributesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMatchableAttributesRequest proto.InternalMessageInfo

func (m *ListMatchableAttributesRequest) GetResourceType() MatchableResource {
	if m != nil {
		return m.ResourceType
	}
	return MatchableResource_TASK_RESOURCE
}

// Response for a request for all matching resource attributes.
type ListMatchableAttributesResponse struct {
	Configurations       []*MatchableAttributesConfiguration `protobuf:"bytes,1,rep,name=configurations,proto3" json:"configurations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ListMatchableAttributesResponse) Reset()         { *m = ListMatchableAttributesResponse{} }
func (m *ListMatchableAttributesResponse) String() string { return proto.CompactTextString(m) }
func (*ListMatchableAttributesResponse) ProtoMessage()    {}
func (*ListMatchableAttributesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{9}
}

func (m *ListMatchableAttributesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMatchableAttributesResponse.Unmarshal(m, b)
}
func (m *ListMatchableAttributesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMatchableAttributesResponse.Marshal(b, m, deterministic)
}
func (m *ListMatchableAttributesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMatchableAttributesResponse.Merge(m, src)
}
func (m *ListMatchableAttributesResponse) XXX_Size() int {
	return xxx_messageInfo_ListMatchableAttributesResponse.Size(m)
}
func (m *ListMatchableAttributesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMatchableAttributesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMatchableAttributesResponse proto.InternalMessageInfo

func (m *ListMatchableAttributesResponse) GetConfigurations() []*MatchableAttributesConfiguration {
	if m != nil {
		return m.Configurations
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.admin.MatchableResource", MatchableResource_name, MatchableResource_value)
	proto.RegisterEnum("flyteidl.admin.PluginOverride_MissingPluginBehavior", PluginOverride_MissingPluginBehavior_name, PluginOverride_MissingPluginBehavior_value)
	proto.RegisterType((*TaskResourceSpec)(nil), "flyteidl.admin.TaskResourceSpec")
	proto.RegisterType((*TaskResourceAttributes)(nil), "flyteidl.admin.TaskResourceAttributes")
	proto.RegisterType((*ClusterResourceAttributes)(nil), "flyteidl.admin.ClusterResourceAttributes")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.ClusterResourceAttributes.AttributesEntry")
	proto.RegisterType((*ExecutionQueueAttributes)(nil), "flyteidl.admin.ExecutionQueueAttributes")
	proto.RegisterType((*ExecutionClusterLabel)(nil), "flyteidl.admin.ExecutionClusterLabel")
	proto.RegisterType((*PluginOverride)(nil), "flyteidl.admin.PluginOverride")
	proto.RegisterType((*MatchingAttributes)(nil), "flyteidl.admin.MatchingAttributes")
	proto.RegisterType((*MatchableAttributesConfiguration)(nil), "flyteidl.admin.MatchableAttributesConfiguration")
	proto.RegisterType((*ListMatchableAttributesRequest)(nil), "flyteidl.admin.ListMatchableAttributesRequest")
	proto.RegisterType((*ListMatchableAttributesResponse)(nil), "flyteidl.admin.ListMatchableAttributesResponse")
}

func init() {
	proto.RegisterFile("flyteidl/admin/matchable_resource.proto", fileDescriptor_1d15bcabb02640f4)
}

var fileDescriptor_1d15bcabb02640f4 = []byte{
	// 958 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xdd, 0x72, 0xda, 0x46,
	0x14, 0x46, 0x06, 0x53, 0x7c, 0x68, 0xb0, 0xb2, 0x8d, 0x6d, 0x05, 0x4f, 0x63, 0xaa, 0xe9, 0x8f,
	0xdb, 0x99, 0x40, 0x4b, 0x7a, 0x91, 0x76, 0xda, 0x0b, 0x20, 0xa2, 0x68, 0x8a, 0x83, 0xbd, 0x40,
	0x26, 0xe9, 0x8d, 0x46, 0x88, 0x45, 0x6c, 0x91, 0xb4, 0xb2, 0xb4, 0x72, 0xca, 0xf4, 0x25, 0xfa,
	0x1a, 0x7d, 0x80, 0x3e, 0x4a, 0xfb, 0x14, 0x7d, 0x88, 0x8c, 0xfe, 0xf8, 0x51, 0x20, 0xe3, 0x3b,
	0x9d, 0xb3, 0xe7, 0x9c, 0xef, 0x9c, 0xb3, 0xdf, 0x7e, 0x00, 0x5f, 0xcd, 0xac, 0x25, 0x27, 0x74,
	0x6a, 0x35, 0xf4, 0xa9, 0x4d, 0x9d, 0x86, 0xad, 0x73, 0x63, 0xae, 0x4f, 0x2c, 0xa2, 0x79, 0xc4,
	0x67, 0x81, 0x67, 0x90, 0xba, 0xeb, 0x31, 0xce, 0x50, 0x25, 0x0d, 0xac, 0x47, 0x81, 0xd5, 0xf3,
	0x4c, 0xa2, 0xc1, 0x6c, 0x9b, 0x39, 0x71, 0x70, 0xf5, 0xd3, 0xd5, 0xa1, 0xc1, 0x3c, 0xd2, 0x20,
	0x7f, 0x10, 0x23, 0xe0, 0x34, 0x3d, 0x96, 0xe7, 0x20, 0x8e, 0x74, 0x7f, 0x81, 0x13, 0x84, 0xa1,
	0x4b, 0x0c, 0x24, 0x42, 0xde, 0x70, 0x03, 0x49, 0xa8, 0x09, 0x97, 0x47, 0x38, 0xfc, 0x0c, 0x3d,
	0xa6, 0x1b, 0x48, 0x07, 0xb1, 0xc7, 0x74, 0x03, 0x74, 0x0a, 0x45, 0x9b, 0xd8, 0xcc, 0x5b, 0x4a,
	0xf9, 0xc8, 0x99, 0x58, 0x48, 0x82, 0x8f, 0x7c, 0xce, 0x3c, 0xdd, 0x24, 0x52, 0x21, 0x3a, 0x48,
	0x4d, 0xf9, 0x2f, 0x01, 0x4e, 0x37, 0xa1, 0x5a, 0x9c, 0x7b, 0x74, 0x12, 0x70, 0xe2, 0xa3, 0x9f,
	0xa0, 0x34, 0x25, 0x33, 0x3d, 0xb0, 0xb8, 0x1f, 0xa1, 0x96, 0x9b, 0xb5, 0xfa, 0xf6, 0x8c, 0xf5,
	0x6c, 0x93, 0x78, 0x95, 0x81, 0x9e, 0x43, 0xd1, 0xa2, 0x36, 0xe5, 0x7e, 0xd4, 0xdf, 0x7d, 0x72,
	0x93, 0x78, 0xf9, 0x1f, 0x01, 0x1e, 0x77, 0xac, 0xc0, 0xe7, 0xc4, 0xdb, 0xd1, 0xd5, 0x1b, 0x00,
	0x7d, 0x65, 0x49, 0x42, 0x2d, 0x7f, 0x59, 0x6e, 0xfe, 0x90, 0xad, 0xbd, 0x37, 0xbd, 0xbe, 0xfe,
	0x54, 0x1c, 0xee, 0x2d, 0xf1, 0x46, 0xb1, 0xea, 0xcf, 0x70, 0x9c, 0x39, 0x0e, 0x57, 0xbc, 0x20,
	0xcb, 0x74, 0xe9, 0x0b, 0xb2, 0x44, 0x8f, 0xe0, 0xf0, 0x4e, 0xb7, 0x02, 0x92, 0xac, 0x3d, 0x36,
	0x7e, 0x3c, 0x78, 0x2e, 0xc8, 0x75, 0x90, 0x94, 0xf4, 0x1e, 0x6f, 0x02, 0x12, 0x6c, 0x76, 0x8d,
	0xa0, 0xc0, 0x75, 0x33, 0xee, 0xf7, 0x08, 0x47, 0xdf, 0xf2, 0x53, 0x38, 0x59, 0xc5, 0x27, 0x0d,
	0xf7, 0xf5, 0x09, 0xb1, 0xd6, 0x10, 0xc2, 0x06, 0x84, 0xfc, 0xbf, 0x00, 0x95, 0x6b, 0x2b, 0x30,
	0xa9, 0x33, 0xb8, 0x23, 0x9e, 0x47, 0xa7, 0x04, 0x9d, 0xc3, 0x11, 0xd7, 0xfd, 0x85, 0xc6, 0x97,
	0x6e, 0x1a, 0x5c, 0x0a, 0x1d, 0xa3, 0xa5, 0x1b, 0x1d, 0xba, 0x51, 0xb8, 0x46, 0xa7, 0xd2, 0x41,
	0x84, 0x5b, 0x8a, 0x1d, 0xea, 0x14, 0x59, 0x70, 0x66, 0x53, 0xdf, 0xa7, 0x8e, 0xa9, 0x25, 0x41,
	0x13, 0x32, 0xd7, 0xef, 0x28, 0xf3, 0x22, 0x82, 0x54, 0x9a, 0xdf, 0x67, 0x57, 0xba, 0x0d, 0x5d,
	0xbf, 0x8a, 0xb3, 0x63, 0x6f, 0x3b, 0xc9, 0xc5, 0x27, 0xf6, 0x2e, 0xb7, 0xdc, 0x84, 0x93, 0x9d,
	0xf1, 0xa8, 0x04, 0x85, 0x6e, 0x4b, 0xed, 0x8b, 0x39, 0x74, 0x0c, 0xe5, 0xf1, 0x50, 0xd1, 0x5e,
	0x28, 0xdd, 0xd6, 0xb8, 0x3f, 0x12, 0x05, 0xf9, 0xbf, 0x02, 0xa0, 0xab, 0xf0, 0xad, 0x51, 0xc7,
	0xdc, 0x58, 0xe4, 0x04, 0xa4, 0x68, 0xe4, 0xf4, 0xf1, 0x69, 0x5b, 0x64, 0x08, 0x89, 0xf6, 0xe5,
	0x87, 0x88, 0xb6, 0xae, 0xd4, 0xcb, 0xe1, 0x53, 0xbe, 0x9b, 0xf8, 0x0b, 0x38, 0x37, 0xe2, 0xfb,
	0xd8, 0x09, 0x13, 0xf3, 0xf9, 0xeb, 0x7b, 0x73, 0xae, 0x97, 0xc3, 0x8f, 0x8d, 0xbd, 0x7c, 0x9e,
	0x43, 0x75, 0xf5, 0xfa, 0xb5, 0xdb, 0x90, 0x36, 0x9b, 0x58, 0xf9, 0x08, 0xeb, 0x32, 0x8b, 0xb5,
	0x8f, 0x67, 0xbd, 0x1c, 0x96, 0xc8, 0x3e, 0x0e, 0x6a, 0x70, 0xb6, 0x46, 0x4a, 0x07, 0xb4, 0x42,
	0xc6, 0x45, 0x77, 0x5e, 0x6e, 0x7e, 0xb1, 0x17, 0x66, 0x93, 0x9e, 0xbd, 0x1c, 0x3e, 0x21, 0x3b,
	0x79, 0x3b, 0x00, 0x74, 0x1b, 0xe8, 0x16, 0xe5, 0x4b, 0x8d, 0xcd, 0x34, 0x9f, 0x78, 0x77, 0xd4,
	0x20, 0xd2, 0x61, 0x54, 0xfb, 0x62, 0x5d, 0x3b, 0x54, 0xbc, 0xfa, 0x4d, 0x1c, 0x38, 0x98, 0x0d,
	0xe3, 0xb0, 0x5e, 0x0e, 0x8b, 0xb7, 0x19, 0x1f, 0x52, 0xe1, 0x38, 0x61, 0x27, 0x4b, 0x78, 0x27,
	0x15, 0xa3, 0x6a, 0x4f, 0x3e, 0xcc, 0xce, 0x5e, 0x0e, 0x57, 0xdc, 0x2d, 0x4f, 0xbb, 0x04, 0x45,
	0xae, 0x7b, 0x26, 0xe1, 0xf2, 0xbf, 0x02, 0xd4, 0xae, 0x52, 0x11, 0x5f, 0xaf, 0xa7, 0xc3, 0x9c,
	0x19, 0x35, 0x03, 0x4f, 0x0f, 0xc7, 0x42, 0xed, 0x8c, 0xca, 0x84, 0xa0, 0x72, 0x16, 0xf4, 0x7d,
	0x7a, 0x6e, 0xca, 0x49, 0x28, 0xc6, 0x53, 0x66, 0xeb, 0xd4, 0x49, 0xa4, 0x22, 0xb1, 0x42, 0x31,
	0x76, 0x3d, 0xf6, 0x3b, 0x31, 0x78, 0xa2, 0xd2, 0xa9, 0x89, 0xaa, 0x50, 0x7a, 0xcb, 0xbc, 0xc5,
	0xcc, 0x62, 0x6f, 0x13, 0x9d, 0x5e, 0xd9, 0xe8, 0x02, 0xca, 0x96, 0x1e, 0x38, 0xc6, 0x5c, 0x73,
	0x2d, 0xdd, 0x89, 0xb6, 0x7a, 0x84, 0x21, 0x76, 0x5d, 0x5b, 0xba, 0x23, 0xcf, 0xe1, 0x49, 0x9f,
	0xfa, 0x7c, 0xc7, 0x68, 0x98, 0xdc, 0x06, 0xc4, 0xe7, 0xa8, 0x0b, 0x0f, 0x56, 0x7c, 0x5e, 0x49,
	0x46, 0xa5, 0xf9, 0xd9, 0xce, 0xb9, 0xc2, 0x12, 0x29, 0x5d, 0xf1, 0xc7, 0x69, 0x5e, 0xa8, 0x2c,
	0xf2, 0x9f, 0x70, 0xb1, 0x17, 0xc9, 0x77, 0x99, 0xe3, 0x13, 0xf4, 0x1a, 0x2a, 0xc6, 0xe6, 0x42,
	0x53, 0xa5, 0xfe, 0x76, 0x2f, 0xd6, 0x9e, 0x9b, 0xc0, 0x99, 0x3a, 0xdf, 0xfc, 0x2d, 0xc0, 0xc3,
	0xf7, 0x1a, 0x44, 0x0f, 0xe1, 0xc1, 0xa8, 0x35, 0xfc, 0x55, 0xc3, 0xca, 0x70, 0x30, 0xc6, 0x1d,
	0x45, 0xcc, 0xa1, 0x47, 0x20, 0x76, 0xfa, 0xe3, 0xe1, 0x48, 0xc1, 0x6b, 0xaf, 0x80, 0x3e, 0x81,
	0x63, 0xe5, 0xb5, 0xd2, 0x19, 0x8f, 0xd4, 0xc1, 0x4b, 0xed, 0x66, 0xac, 0x8c, 0x15, 0xf1, 0x00,
	0x9d, 0xc3, 0xd9, 0xda, 0x99, 0x26, 0xf5, 0x5b, 0x6d, 0xa5, 0x2f, 0xe6, 0xd1, 0xe7, 0x50, 0xbb,
	0x19, 0xb7, 0xfa, 0xea, 0xe8, 0x8d, 0x36, 0xe8, 0x6a, 0x43, 0x05, 0xbf, 0x52, 0x3b, 0x8a, 0x36,
	0xbc, 0x56, 0x3a, 0x6a, 0x57, 0xed, 0xb4, 0xc2, 0x1c, 0xb1, 0x10, 0xd6, 0xbd, 0xee, 0x8f, 0x7f,
	0x51, 0x5f, 0x6a, 0x83, 0x57, 0x0a, 0xc6, 0xea, 0x0b, 0x45, 0x3c, 0x6c, 0x3f, 0xfb, 0xed, 0x3b,
	0x93, 0xf2, 0x79, 0x30, 0xa9, 0x1b, 0xcc, 0x6e, 0x58, 0xcb, 0x19, 0x6f, 0xac, 0x7e, 0xf7, 0x4d,
	0xe2, 0x34, 0xdc, 0xc9, 0x53, 0x93, 0x35, 0xb6, 0xff, 0x27, 0x4c, 0x8a, 0xd1, 0x5f, 0x80, 0x67,
	0xef, 0x02, 0x00, 0x00, 0xff, 0xff, 0x12, 0x81, 0xff, 0x49, 0x79, 0x08, 0x00, 0x00,
}
