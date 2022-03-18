// Copyright 2020 The PipeCD Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pkg/model/insight.proto

package model

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type InsightResultType int32

const (
	InsightResultType_MATRIX InsightResultType = 0
	InsightResultType_VECTOR InsightResultType = 1
)

// Enum value maps for InsightResultType.
var (
	InsightResultType_name = map[int32]string{
		0: "MATRIX",
		1: "VECTOR",
	}
	InsightResultType_value = map[string]int32{
		"MATRIX": 0,
		"VECTOR": 1,
	}
)

func (x InsightResultType) Enum() *InsightResultType {
	p := new(InsightResultType)
	*p = x
	return p
}

func (x InsightResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsightResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_model_insight_proto_enumTypes[0].Descriptor()
}

func (InsightResultType) Type() protoreflect.EnumType {
	return &file_pkg_model_insight_proto_enumTypes[0]
}

func (x InsightResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsightResultType.Descriptor instead.
func (InsightResultType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_model_insight_proto_rawDescGZIP(), []int{0}
}

type InsightMetricsKind int32

const (
	InsightMetricsKind_DEPLOYMENT_FREQUENCY InsightMetricsKind = 0
	InsightMetricsKind_CHANGE_FAILURE_RATE  InsightMetricsKind = 1
	InsightMetricsKind_MTTR                 InsightMetricsKind = 2
	InsightMetricsKind_LEAD_TIME            InsightMetricsKind = 3
	InsightMetricsKind_APPLICATIONS_COUNT   InsightMetricsKind = 4
)

// Enum value maps for InsightMetricsKind.
var (
	InsightMetricsKind_name = map[int32]string{
		0: "DEPLOYMENT_FREQUENCY",
		1: "CHANGE_FAILURE_RATE",
		2: "MTTR",
		3: "LEAD_TIME",
		4: "APPLICATIONS_COUNT",
	}
	InsightMetricsKind_value = map[string]int32{
		"DEPLOYMENT_FREQUENCY": 0,
		"CHANGE_FAILURE_RATE":  1,
		"MTTR":                 2,
		"LEAD_TIME":            3,
		"APPLICATIONS_COUNT":   4,
	}
)

func (x InsightMetricsKind) Enum() *InsightMetricsKind {
	p := new(InsightMetricsKind)
	*p = x
	return p
}

func (x InsightMetricsKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsightMetricsKind) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_model_insight_proto_enumTypes[1].Descriptor()
}

func (InsightMetricsKind) Type() protoreflect.EnumType {
	return &file_pkg_model_insight_proto_enumTypes[1]
}

func (x InsightMetricsKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsightMetricsKind.Descriptor instead.
func (InsightMetricsKind) EnumDescriptor() ([]byte, []int) {
	return file_pkg_model_insight_proto_rawDescGZIP(), []int{1}
}

type InsightApplicationCountLabelKey int32

const (
	InsightApplicationCountLabelKey_KIND          InsightApplicationCountLabelKey = 0
	InsightApplicationCountLabelKey_ACTIVE_STATUS InsightApplicationCountLabelKey = 1
)

// Enum value maps for InsightApplicationCountLabelKey.
var (
	InsightApplicationCountLabelKey_name = map[int32]string{
		0: "KIND",
		1: "ACTIVE_STATUS",
	}
	InsightApplicationCountLabelKey_value = map[string]int32{
		"KIND":          0,
		"ACTIVE_STATUS": 1,
	}
)

func (x InsightApplicationCountLabelKey) Enum() *InsightApplicationCountLabelKey {
	p := new(InsightApplicationCountLabelKey)
	*p = x
	return p
}

func (x InsightApplicationCountLabelKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsightApplicationCountLabelKey) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_model_insight_proto_enumTypes[2].Descriptor()
}

func (InsightApplicationCountLabelKey) Type() protoreflect.EnumType {
	return &file_pkg_model_insight_proto_enumTypes[2]
}

func (x InsightApplicationCountLabelKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsightApplicationCountLabelKey.Descriptor instead.
func (InsightApplicationCountLabelKey) EnumDescriptor() ([]byte, []int) {
	return file_pkg_model_insight_proto_rawDescGZIP(), []int{2}
}

type InsightSample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels    map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DataPoint *InsightDataPoint `protobuf:"bytes,2,opt,name=data_point,json=dataPoint,proto3" json:"data_point,omitempty"`
}

func (x *InsightSample) Reset() {
	*x = InsightSample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_insight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsightSample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsightSample) ProtoMessage() {}

func (x *InsightSample) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_insight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsightSample.ProtoReflect.Descriptor instead.
func (*InsightSample) Descriptor() ([]byte, []int) {
	return file_pkg_model_insight_proto_rawDescGZIP(), []int{0}
}

func (x *InsightSample) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *InsightSample) GetDataPoint() *InsightDataPoint {
	if x != nil {
		return x.DataPoint
	}
	return nil
}

type InsightSampleStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels     map[string]string   `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DataPoints []*InsightDataPoint `protobuf:"bytes,2,rep,name=data_points,json=dataPoints,proto3" json:"data_points,omitempty"`
}

func (x *InsightSampleStream) Reset() {
	*x = InsightSampleStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_insight_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsightSampleStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsightSampleStream) ProtoMessage() {}

func (x *InsightSampleStream) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_insight_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsightSampleStream.ProtoReflect.Descriptor instead.
func (*InsightSampleStream) Descriptor() ([]byte, []int) {
	return file_pkg_model_insight_proto_rawDescGZIP(), []int{1}
}

func (x *InsightSampleStream) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *InsightSampleStream) GetDataPoints() []*InsightDataPoint {
	if x != nil {
		return x.DataPoints
	}
	return nil
}

type InsightDataPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value     float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *InsightDataPoint) Reset() {
	*x = InsightDataPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_insight_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsightDataPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsightDataPoint) ProtoMessage() {}

func (x *InsightDataPoint) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_insight_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsightDataPoint.ProtoReflect.Descriptor instead.
func (*InsightDataPoint) Descriptor() ([]byte, []int) {
	return file_pkg_model_insight_proto_rawDescGZIP(), []int{2}
}

func (x *InsightDataPoint) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *InsightDataPoint) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type InsightApplicationCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Count  int32             `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *InsightApplicationCount) Reset() {
	*x = InsightApplicationCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_insight_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsightApplicationCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsightApplicationCount) ProtoMessage() {}

func (x *InsightApplicationCount) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_insight_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsightApplicationCount.ProtoReflect.Descriptor instead.
func (*InsightApplicationCount) Descriptor() ([]byte, []int) {
	return file_pkg_model_insight_proto_rawDescGZIP(), []int{3}
}

func (x *InsightApplicationCount) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *InsightApplicationCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DeploymentSubset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt int64 `protobuf:"varint,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64 `protobuf:"varint,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *DeploymentSubset) Reset() {
	*x = DeploymentSubset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_insight_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentSubset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentSubset) ProtoMessage() {}

func (x *DeploymentSubset) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_insight_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentSubset.ProtoReflect.Descriptor instead.
func (*DeploymentSubset) Descriptor() ([]byte, []int) {
	return file_pkg_model_insight_proto_rawDescGZIP(), []int{4}
}

func (x *DeploymentSubset) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *DeploymentSubset) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type DailyDeployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date      int64 `protobuf:"varint,1,opt,name=date,proto3" json:"date,omitempty"`
	CreatedAt int64 `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64 `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Daily DeploymentSubset sorted by created_at
	DailyDeployments []*DeploymentSubset `protobuf:"bytes,4,rep,name=daily_deployments,json=dailyDeployments,proto3" json:"daily_deployments,omitempty"`
}

func (x *DailyDeployment) Reset() {
	*x = DailyDeployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_insight_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyDeployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyDeployment) ProtoMessage() {}

func (x *DailyDeployment) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_insight_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyDeployment.ProtoReflect.Descriptor instead.
func (*DailyDeployment) Descriptor() ([]byte, []int) {
	return file_pkg_model_insight_proto_rawDescGZIP(), []int{5}
}

func (x *DailyDeployment) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *DailyDeployment) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *DailyDeployment) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *DailyDeployment) GetDailyDeployments() []*DeploymentSubset {
	if x != nil {
		return x.DailyDeployments
	}
	return nil
}

// DeploymentChunk represents a chunk that is used by insight
type DeploymentChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateRange *ChunkDateRange `protobuf:"bytes,1,opt,name=date_range,json=dateRange,proto3" json:"date_range,omitempty"`
	// Daily-deployments sorted by date
	Deployments []*DailyDeployment `protobuf:"bytes,2,rep,name=deployments,proto3" json:"deployments,omitempty"`
}

func (x *DeploymentChunk) Reset() {
	*x = DeploymentChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_insight_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentChunk) ProtoMessage() {}

func (x *DeploymentChunk) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_insight_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentChunk.ProtoReflect.Descriptor instead.
func (*DeploymentChunk) Descriptor() ([]byte, []int) {
	return file_pkg_model_insight_proto_rawDescGZIP(), []int{6}
}

func (x *DeploymentChunk) GetDateRange() *ChunkDateRange {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *DeploymentChunk) GetDeployments() []*DailyDeployment {
	if x != nil {
		return x.Deployments
	}
	return nil
}

type DeploymentChunkMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sorted by Date. date_range does not overlap
	Data []*DeploymentChunkMetaData_ChunkData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *DeploymentChunkMetaData) Reset() {
	*x = DeploymentChunkMetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_insight_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentChunkMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentChunkMetaData) ProtoMessage() {}

func (x *DeploymentChunkMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_insight_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentChunkMetaData.ProtoReflect.Descriptor instead.
func (*DeploymentChunkMetaData) Descriptor() ([]byte, []int) {
	return file_pkg_model_insight_proto_rawDescGZIP(), []int{7}
}

func (x *DeploymentChunkMetaData) GetData() []*DeploymentChunkMetaData_ChunkData {
	if x != nil {
		return x.Data
	}
	return nil
}

type ChunkDateRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From int64 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To   int64 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *ChunkDateRange) Reset() {
	*x = ChunkDateRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_insight_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkDateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkDateRange) ProtoMessage() {}

func (x *ChunkDateRange) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_insight_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkDateRange.ProtoReflect.Descriptor instead.
func (*ChunkDateRange) Descriptor() ([]byte, []int) {
	return file_pkg_model_insight_proto_rawDescGZIP(), []int{8}
}

func (x *ChunkDateRange) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *ChunkDateRange) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

type DeploymentChunkMetaData_ChunkData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateRange *ChunkDateRange `protobuf:"bytes,1,opt,name=date_range,json=dateRange,proto3" json:"date_range,omitempty"`
	// Unique chunk identifier
	ChunkKey string `protobuf:"bytes,2,opt,name=chunk_key,json=chunkKey,proto3" json:"chunk_key,omitempty"`
	// chunk size (byte)
	ChunkSize int64 `protobuf:"varint,3,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
}

func (x *DeploymentChunkMetaData_ChunkData) Reset() {
	*x = DeploymentChunkMetaData_ChunkData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_insight_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentChunkMetaData_ChunkData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentChunkMetaData_ChunkData) ProtoMessage() {}

func (x *DeploymentChunkMetaData_ChunkData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_insight_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentChunkMetaData_ChunkData.ProtoReflect.Descriptor instead.
func (*DeploymentChunkMetaData_ChunkData) Descriptor() ([]byte, []int) {
	return file_pkg_model_insight_proto_rawDescGZIP(), []int{7, 0}
}

func (x *DeploymentChunkMetaData_ChunkData) GetDateRange() *ChunkDateRange {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *DeploymentChunkMetaData_ChunkData) GetChunkKey() string {
	if x != nil {
		return x.ChunkKey
	}
	return ""
}

func (x *DeploymentChunkMetaData_ChunkData) GetChunkSize() int64 {
	if x != nil {
		return x.ChunkSize
	}
	return 0
}

var File_pkg_model_insight_proto protoreflect.FileDescriptor

var file_pkg_model_insight_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x69,
	0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x0d, 0x49, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x39, 0x0a,
	0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xca, 0x01, 0x0a, 0x13, 0x49, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x3e, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x38, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a,
	0x64, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5b, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x0a, 0x05, 0x25, 0x00, 0x00, 0x00, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x17, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x42,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x62, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x28, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x26, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x00, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc4, 0x01, 0x0a, 0x0f, 0x44, 0x61, 0x69, 0x6c,
	0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02,
	0x28, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x22, 0x02, 0x28, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x26, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x00, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x11, 0x64, 0x61, 0x69, 0x6c,
	0x79, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x52, 0x10, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x81,
	0x01, 0x0a, 0x0f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x34, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x17, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3c,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x7d, 0x0a, 0x09,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x46, 0x0a, 0x0e, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x28, 0x00, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x00, 0x52,
	0x02, 0x74, 0x6f, 0x2a, 0x2b, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x54, 0x52,
	0x49, 0x58, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x10, 0x01,
	0x2a, 0x78, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x52, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x00,
	0x12, 0x17, 0x0a, 0x13, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55,
	0x52, 0x45, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x54, 0x54,
	0x52, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x45, 0x41, 0x44, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x53, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x04, 0x2a, 0x3e, 0x0a, 0x1f, 0x49, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x08, 0x0a,
	0x04, 0x4b, 0x49, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x2d, 0x63, 0x64,
	0x2f, 0x70, 0x69, 0x70, 0x65, 0x63, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_model_insight_proto_rawDescOnce sync.Once
	file_pkg_model_insight_proto_rawDescData = file_pkg_model_insight_proto_rawDesc
)

func file_pkg_model_insight_proto_rawDescGZIP() []byte {
	file_pkg_model_insight_proto_rawDescOnce.Do(func() {
		file_pkg_model_insight_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_model_insight_proto_rawDescData)
	})
	return file_pkg_model_insight_proto_rawDescData
}

var file_pkg_model_insight_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_pkg_model_insight_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pkg_model_insight_proto_goTypes = []interface{}{
	(InsightResultType)(0),                    // 0: model.InsightResultType
	(InsightMetricsKind)(0),                   // 1: model.InsightMetricsKind
	(InsightApplicationCountLabelKey)(0),      // 2: model.InsightApplicationCountLabelKey
	(*InsightSample)(nil),                     // 3: model.InsightSample
	(*InsightSampleStream)(nil),               // 4: model.InsightSampleStream
	(*InsightDataPoint)(nil),                  // 5: model.InsightDataPoint
	(*InsightApplicationCount)(nil),           // 6: model.InsightApplicationCount
	(*DeploymentSubset)(nil),                  // 7: model.DeploymentSubset
	(*DailyDeployment)(nil),                   // 8: model.DailyDeployment
	(*DeploymentChunk)(nil),                   // 9: model.DeploymentChunk
	(*DeploymentChunkMetaData)(nil),           // 10: model.DeploymentChunkMetaData
	(*ChunkDateRange)(nil),                    // 11: model.ChunkDateRange
	nil,                                       // 12: model.InsightSample.LabelsEntry
	nil,                                       // 13: model.InsightSampleStream.LabelsEntry
	nil,                                       // 14: model.InsightApplicationCount.LabelsEntry
	(*DeploymentChunkMetaData_ChunkData)(nil), // 15: model.DeploymentChunkMetaData.ChunkData
}
var file_pkg_model_insight_proto_depIdxs = []int32{
	12, // 0: model.InsightSample.labels:type_name -> model.InsightSample.LabelsEntry
	5,  // 1: model.InsightSample.data_point:type_name -> model.InsightDataPoint
	13, // 2: model.InsightSampleStream.labels:type_name -> model.InsightSampleStream.LabelsEntry
	5,  // 3: model.InsightSampleStream.data_points:type_name -> model.InsightDataPoint
	14, // 4: model.InsightApplicationCount.labels:type_name -> model.InsightApplicationCount.LabelsEntry
	7,  // 5: model.DailyDeployment.daily_deployments:type_name -> model.DeploymentSubset
	11, // 6: model.DeploymentChunk.date_range:type_name -> model.ChunkDateRange
	8,  // 7: model.DeploymentChunk.deployments:type_name -> model.DailyDeployment
	15, // 8: model.DeploymentChunkMetaData.data:type_name -> model.DeploymentChunkMetaData.ChunkData
	11, // 9: model.DeploymentChunkMetaData.ChunkData.date_range:type_name -> model.ChunkDateRange
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pkg_model_insight_proto_init() }
func file_pkg_model_insight_proto_init() {
	if File_pkg_model_insight_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_model_insight_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsightSample); i {
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
		file_pkg_model_insight_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsightSampleStream); i {
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
		file_pkg_model_insight_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsightDataPoint); i {
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
		file_pkg_model_insight_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsightApplicationCount); i {
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
		file_pkg_model_insight_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentSubset); i {
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
		file_pkg_model_insight_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyDeployment); i {
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
		file_pkg_model_insight_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentChunk); i {
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
		file_pkg_model_insight_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentChunkMetaData); i {
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
		file_pkg_model_insight_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkDateRange); i {
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
		file_pkg_model_insight_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentChunkMetaData_ChunkData); i {
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
			RawDescriptor: file_pkg_model_insight_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_model_insight_proto_goTypes,
		DependencyIndexes: file_pkg_model_insight_proto_depIdxs,
		EnumInfos:         file_pkg_model_insight_proto_enumTypes,
		MessageInfos:      file_pkg_model_insight_proto_msgTypes,
	}.Build()
	File_pkg_model_insight_proto = out.File
	file_pkg_model_insight_proto_rawDesc = nil
	file_pkg_model_insight_proto_goTypes = nil
	file_pkg_model_insight_proto_depIdxs = nil
}
