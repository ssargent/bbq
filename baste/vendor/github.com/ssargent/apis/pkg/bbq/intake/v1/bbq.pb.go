// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: bbq/intake/v1/bbq.proto

package intakev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sensor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Manufacturer    string `protobuf:"bytes,2,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	SensorCount     int32  `protobuf:"varint,3,opt,name=sensor_count,json=sensorCount,proto3" json:"sensor_count,omitempty"`
	TemperatureUnit string `protobuf:"bytes,4,opt,name=temperature_unit,json=temperatureUnit,proto3" json:"temperature_unit,omitempty"`
	SensorId        string `protobuf:"bytes,5,opt,name=sensor_id,json=sensorId,proto3" json:"sensor_id,omitempty"`
}

func (x *Sensor) Reset() {
	*x = Sensor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbq_intake_v1_bbq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sensor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sensor) ProtoMessage() {}

func (x *Sensor) ProtoReflect() protoreflect.Message {
	mi := &file_bbq_intake_v1_bbq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sensor.ProtoReflect.Descriptor instead.
func (*Sensor) Descriptor() ([]byte, []int) {
	return file_bbq_intake_v1_bbq_proto_rawDescGZIP(), []int{0}
}

func (x *Sensor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sensor) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *Sensor) GetSensorCount() int32 {
	if x != nil {
		return x.SensorCount
	}
	return 0
}

func (x *Sensor) GetTemperatureUnit() string {
	if x != nil {
		return x.TemperatureUnit
	}
	return ""
}

func (x *Sensor) GetSensorId() string {
	if x != nil {
		return x.SensorId
	}
	return ""
}

type SensorReading struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensorNumber int32   `protobuf:"varint,1,opt,name=sensor_number,json=sensorNumber,proto3" json:"sensor_number,omitempty"`
	Temperature  float32 `protobuf:"fixed32,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *SensorReading) Reset() {
	*x = SensorReading{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbq_intake_v1_bbq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensorReading) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorReading) ProtoMessage() {}

func (x *SensorReading) ProtoReflect() protoreflect.Message {
	mi := &file_bbq_intake_v1_bbq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorReading.ProtoReflect.Descriptor instead.
func (*SensorReading) Descriptor() ([]byte, []int) {
	return file_bbq_intake_v1_bbq_proto_rawDescGZIP(), []int{1}
}

func (x *SensorReading) GetSensorNumber() int32 {
	if x != nil {
		return x.SensorNumber
	}
	return 0
}

func (x *SensorReading) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

type Reading struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sensor     *Sensor                `protobuf:"bytes,1,opt,name=sensor,proto3" json:"sensor,omitempty"`
	Readings   []*SensorReading       `protobuf:"bytes,2,rep,name=readings,proto3" json:"readings,omitempty"`
	RecordedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=recorded_at,json=recordedAt,proto3" json:"recorded_at,omitempty"`
}

func (x *Reading) Reset() {
	*x = Reading{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbq_intake_v1_bbq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reading) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reading) ProtoMessage() {}

func (x *Reading) ProtoReflect() protoreflect.Message {
	mi := &file_bbq_intake_v1_bbq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reading.ProtoReflect.Descriptor instead.
func (*Reading) Descriptor() ([]byte, []int) {
	return file_bbq_intake_v1_bbq_proto_rawDescGZIP(), []int{2}
}

func (x *Reading) GetSensor() *Sensor {
	if x != nil {
		return x.Sensor
	}
	return nil
}

func (x *Reading) GetReadings() []*SensorReading {
	if x != nil {
		return x.Readings
	}
	return nil
}

func (x *Reading) GetRecordedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RecordedAt
	}
	return nil
}

type SessionDataRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sensors           int32 `protobuf:"varint,1,opt,name=sensors,proto3" json:"sensors,omitempty"`
	MaxReadingsMinute int32 `protobuf:"varint,2,opt,name=max_readings_minute,json=maxReadingsMinute,proto3" json:"max_readings_minute,omitempty"`
}

func (x *SessionDataRate) Reset() {
	*x = SessionDataRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbq_intake_v1_bbq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionDataRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionDataRate) ProtoMessage() {}

func (x *SessionDataRate) ProtoReflect() protoreflect.Message {
	mi := &file_bbq_intake_v1_bbq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionDataRate.ProtoReflect.Descriptor instead.
func (*SessionDataRate) Descriptor() ([]byte, []int) {
	return file_bbq_intake_v1_bbq_proto_rawDescGZIP(), []int{3}
}

func (x *SessionDataRate) GetSensors() int32 {
	if x != nil {
		return x.Sensors
	}
	return 0
}

func (x *SessionDataRate) GetMaxReadingsMinute() int32 {
	if x != nil {
		return x.MaxReadingsMinute
	}
	return 0
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DataRate *SessionDataRate `protobuf:"bytes,3,opt,name=data_rate,json=dataRate,proto3" json:"data_rate,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbq_intake_v1_bbq_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_bbq_intake_v1_bbq_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_bbq_intake_v1_bbq_proto_rawDescGZIP(), []int{4}
}

func (x *Session) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Session) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Session) GetDataRate() *SessionDataRate {
	if x != nil {
		return x.DataRate
	}
	return nil
}

var File_bbq_intake_v1_bbq_proto protoreflect.FileDescriptor

var file_bbq_intake_v1_bbq_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x62, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x62, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x62, 0x62, 0x71, 0x2e, 0x69,
	0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x06, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75,
	0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x29, 0x0a, 0x10, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22,
	0xaf, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2d, 0x0a, 0x06, 0x73,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x62,
	0x71, 0x2e, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62,
	0x62, 0x71, 0x2e, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x5b, 0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x12, 0x2e,
	0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x61, 0x78,
	0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x22, 0x6a,
	0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x62, 0x62, 0x71, 0x2e, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x61, 0x74, 0x65,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x52, 0x61, 0x74, 0x65, 0x42, 0xa8, 0x01, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x2e, 0x62, 0x62, 0x71, 0x2e, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x08, 0x42, 0x62, 0x71, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x73, 0x61, 0x72, 0x67, 0x65, 0x6e,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x62, 0x71, 0x2f, 0x69,
	0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x42, 0x49, 0x58, 0xaa, 0x02, 0x0d, 0x42, 0x62, 0x71, 0x2e, 0x49, 0x6e,
	0x74, 0x61, 0x6b, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x42, 0x62, 0x71, 0x5c, 0x49, 0x6e,
	0x74, 0x61, 0x6b, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x42, 0x62, 0x71, 0x5c, 0x49, 0x6e,
	0x74, 0x61, 0x6b, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x42, 0x62, 0x71, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x61, 0x6b,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bbq_intake_v1_bbq_proto_rawDescOnce sync.Once
	file_bbq_intake_v1_bbq_proto_rawDescData = file_bbq_intake_v1_bbq_proto_rawDesc
)

func file_bbq_intake_v1_bbq_proto_rawDescGZIP() []byte {
	file_bbq_intake_v1_bbq_proto_rawDescOnce.Do(func() {
		file_bbq_intake_v1_bbq_proto_rawDescData = protoimpl.X.CompressGZIP(file_bbq_intake_v1_bbq_proto_rawDescData)
	})
	return file_bbq_intake_v1_bbq_proto_rawDescData
}

var file_bbq_intake_v1_bbq_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bbq_intake_v1_bbq_proto_goTypes = []interface{}{
	(*Sensor)(nil),                // 0: bbq.intake.v1.Sensor
	(*SensorReading)(nil),         // 1: bbq.intake.v1.SensorReading
	(*Reading)(nil),               // 2: bbq.intake.v1.Reading
	(*SessionDataRate)(nil),       // 3: bbq.intake.v1.SessionDataRate
	(*Session)(nil),               // 4: bbq.intake.v1.Session
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_bbq_intake_v1_bbq_proto_depIdxs = []int32{
	0, // 0: bbq.intake.v1.Reading.sensor:type_name -> bbq.intake.v1.Sensor
	1, // 1: bbq.intake.v1.Reading.readings:type_name -> bbq.intake.v1.SensorReading
	5, // 2: bbq.intake.v1.Reading.recorded_at:type_name -> google.protobuf.Timestamp
	3, // 3: bbq.intake.v1.Session.data_rate:type_name -> bbq.intake.v1.SessionDataRate
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bbq_intake_v1_bbq_proto_init() }
func file_bbq_intake_v1_bbq_proto_init() {
	if File_bbq_intake_v1_bbq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bbq_intake_v1_bbq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sensor); i {
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
		file_bbq_intake_v1_bbq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensorReading); i {
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
		file_bbq_intake_v1_bbq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reading); i {
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
		file_bbq_intake_v1_bbq_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionDataRate); i {
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
		file_bbq_intake_v1_bbq_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
			RawDescriptor: file_bbq_intake_v1_bbq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bbq_intake_v1_bbq_proto_goTypes,
		DependencyIndexes: file_bbq_intake_v1_bbq_proto_depIdxs,
		MessageInfos:      file_bbq_intake_v1_bbq_proto_msgTypes,
	}.Build()
	File_bbq_intake_v1_bbq_proto = out.File
	file_bbq_intake_v1_bbq_proto_rawDesc = nil
	file_bbq_intake_v1_bbq_proto_goTypes = nil
	file_bbq_intake_v1_bbq_proto_depIdxs = nil
}
