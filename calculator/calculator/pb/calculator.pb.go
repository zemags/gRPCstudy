// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: api/calculator/calculator.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstValue  int32 `protobuf:"varint,1,opt,name=first_value,json=firstValue,proto3" json:"first_value,omitempty"`
	SecondValue int32 `protobuf:"varint,2,opt,name=second_value,json=secondValue,proto3" json:"second_value,omitempty"`
}

func (x *Values) Reset() {
	*x = Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calculator_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Values) ProtoMessage() {}

func (x *Values) ProtoReflect() protoreflect.Message {
	mi := &file_api_calculator_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Values.ProtoReflect.Descriptor instead.
func (*Values) Descriptor() ([]byte, []int) {
	return file_api_calculator_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *Values) GetFirstValue() int32 {
	if x != nil {
		return x.FirstValue
	}
	return 0
}

func (x *Values) GetSecondValue() int32 {
	if x != nil {
		return x.SecondValue
	}
	return 0
}

type CalculatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values *Values `protobuf:"bytes,1,opt,name=values,proto3" json:"values,omitempty"`
}

func (x *CalculatorRequest) Reset() {
	*x = CalculatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calculator_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest) ProtoMessage() {}

func (x *CalculatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_calculator_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest.ProtoReflect.Descriptor instead.
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return file_api_calculator_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculatorRequest) GetValues() *Values {
	if x != nil {
		return x.Values
	}
	return nil
}

type CalculatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *CalculatorResponse) Reset() {
	*x = CalculatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calculator_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorResponse) ProtoMessage() {}

func (x *CalculatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_calculator_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorResponse.ProtoReflect.Descriptor instead.
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return file_api_calculator_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *CalculatorResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositiveInteger int32 `protobuf:"varint,1,opt,name=positive_integer,json=positiveInteger,proto3" json:"positive_integer,omitempty"`
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calculator_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_calculator_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_api_calculator_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *StreamRequest) GetPositiveInteger() int32 {
	if x != nil {
		return x.PositiveInteger
	}
	return 0
}

type StreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Factor int32 `protobuf:"varint,1,opt,name=factor,proto3" json:"factor,omitempty"`
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calculator_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_calculator_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_api_calculator_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *StreamResponse) GetFactor() int32 {
	if x != nil {
		return x.Factor
	}
	return 0
}

type ResponseDouble struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average float64 `protobuf:"fixed64,1,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *ResponseDouble) Reset() {
	*x = ResponseDouble{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calculator_calculator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDouble) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDouble) ProtoMessage() {}

func (x *ResponseDouble) ProtoReflect() protoreflect.Message {
	mi := &file_api_calculator_calculator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseDouble.ProtoReflect.Descriptor instead.
func (*ResponseDouble) Descriptor() ([]byte, []int) {
	return file_api_calculator_calculator_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseDouble) GetAverage() float64 {
	if x != nil {
		return x.Average
	}
	return 0
}

type ResponseMaximum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maximum int32 `protobuf:"varint,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
}

func (x *ResponseMaximum) Reset() {
	*x = ResponseMaximum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calculator_calculator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMaximum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMaximum) ProtoMessage() {}

func (x *ResponseMaximum) ProtoReflect() protoreflect.Message {
	mi := &file_api_calculator_calculator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMaximum.ProtoReflect.Descriptor instead.
func (*ResponseMaximum) Descriptor() ([]byte, []int) {
	return file_api_calculator_calculator_proto_rawDescGZIP(), []int{6}
}

func (x *ResponseMaximum) GetMaximum() int32 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

var File_api_calculator_calculator_proto protoreflect.FileDescriptor

var file_api_calculator_calculator_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x4c, 0x0a,
	0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3f, 0x0a, 0x11, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x12,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x73, 0x75, 0x6d, 0x22, 0x3a, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x22, 0x28, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x2a, 0x0a, 0x0e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78,
	0x69, 0x6d, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x32, 0xb0, 0x02, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x47, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x63, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0d, 0x44,
	0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x47, 0x0a,
	0x07, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_calculator_calculator_proto_rawDescOnce sync.Once
	file_api_calculator_calculator_proto_rawDescData = file_api_calculator_calculator_proto_rawDesc
)

func file_api_calculator_calculator_proto_rawDescGZIP() []byte {
	file_api_calculator_calculator_proto_rawDescOnce.Do(func() {
		file_api_calculator_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_calculator_calculator_proto_rawDescData)
	})
	return file_api_calculator_calculator_proto_rawDescData
}

var file_api_calculator_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_calculator_calculator_proto_goTypes = []interface{}{
	(*Values)(nil),             // 0: calculator.Values
	(*CalculatorRequest)(nil),  // 1: calculator.CalculatorRequest
	(*CalculatorResponse)(nil), // 2: calculator.CalculatorResponse
	(*StreamRequest)(nil),      // 3: calculator.StreamRequest
	(*StreamResponse)(nil),     // 4: calculator.StreamResponse
	(*ResponseDouble)(nil),     // 5: calculator.ResponseDouble
	(*ResponseMaximum)(nil),    // 6: calculator.ResponseMaximum
}
var file_api_calculator_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.CalculatorRequest.values:type_name -> calculator.Values
	1, // 1: calculator.Calculator.Calc:input_type -> calculator.CalculatorRequest
	3, // 2: calculator.Calculator.Decomposition:input_type -> calculator.StreamRequest
	3, // 3: calculator.Calculator.Average:input_type -> calculator.StreamRequest
	3, // 4: calculator.Calculator.Maximum:input_type -> calculator.StreamRequest
	2, // 5: calculator.Calculator.Calc:output_type -> calculator.CalculatorResponse
	4, // 6: calculator.Calculator.Decomposition:output_type -> calculator.StreamResponse
	5, // 7: calculator.Calculator.Average:output_type -> calculator.ResponseDouble
	6, // 8: calculator.Calculator.Maximum:output_type -> calculator.ResponseMaximum
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_calculator_calculator_proto_init() }
func file_api_calculator_calculator_proto_init() {
	if File_api_calculator_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_calculator_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Values); i {
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
		file_api_calculator_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorRequest); i {
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
		file_api_calculator_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorResponse); i {
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
		file_api_calculator_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
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
		file_api_calculator_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponse); i {
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
		file_api_calculator_calculator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseDouble); i {
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
		file_api_calculator_calculator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMaximum); i {
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
			RawDescriptor: file_api_calculator_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_calculator_calculator_proto_goTypes,
		DependencyIndexes: file_api_calculator_calculator_proto_depIdxs,
		MessageInfos:      file_api_calculator_calculator_proto_msgTypes,
	}.Build()
	File_api_calculator_calculator_proto = out.File
	file_api_calculator_calculator_proto_rawDesc = nil
	file_api_calculator_calculator_proto_goTypes = nil
	file_api_calculator_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	// Implement calculator unary
	Calc(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	// Calculator server streaming
	Decomposition(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Calculator_DecompositionClient, error)
	// Calculator client streaming
	Average(ctx context.Context, opts ...grpc.CallOption) (Calculator_AverageClient, error)
	// Calculator bi-directional streaming
	Maximum(ctx context.Context, opts ...grpc.CallOption) (Calculator_MaximumClient, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Calc(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Calc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Decomposition(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Calculator_DecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Calculator_serviceDesc.Streams[0], "/calculator.Calculator/Decomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Calculator_DecompositionClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type calculatorDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorDecompositionClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) Average(ctx context.Context, opts ...grpc.CallOption) (Calculator_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Calculator_serviceDesc.Streams[1], "/calculator.Calculator/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorAverageClient{stream}
	return x, nil
}

type Calculator_AverageClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*ResponseDouble, error)
	grpc.ClientStream
}

type calculatorAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorAverageClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorAverageClient) CloseAndRecv() (*ResponseDouble, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ResponseDouble)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) Maximum(ctx context.Context, opts ...grpc.CallOption) (Calculator_MaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Calculator_serviceDesc.Streams[2], "/calculator.Calculator/Maximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorMaximumClient{stream}
	return x, nil
}

type Calculator_MaximumClient interface {
	Send(*StreamRequest) error
	Recv() (*ResponseMaximum, error)
	grpc.ClientStream
}

type calculatorMaximumClient struct {
	grpc.ClientStream
}

func (x *calculatorMaximumClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorMaximumClient) Recv() (*ResponseMaximum, error) {
	m := new(ResponseMaximum)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	// Implement calculator unary
	Calc(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	// Calculator server streaming
	Decomposition(*StreamRequest, Calculator_DecompositionServer) error
	// Calculator client streaming
	Average(Calculator_AverageServer) error
	// Calculator bi-directional streaming
	Maximum(Calculator_MaximumServer) error
}

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (*UnimplementedCalculatorServer) Calc(context.Context, *CalculatorRequest) (*CalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calc not implemented")
}
func (*UnimplementedCalculatorServer) Decomposition(*StreamRequest, Calculator_DecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method Decomposition not implemented")
}
func (*UnimplementedCalculatorServer) Average(Calculator_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (*UnimplementedCalculatorServer) Maximum(Calculator_MaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method Maximum not implemented")
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Calc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Calc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Calc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Calc(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Decomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServer).Decomposition(m, &calculatorDecompositionServer{stream})
}

type Calculator_DecompositionServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type calculatorDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorDecompositionServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Calculator_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).Average(&calculatorAverageServer{stream})
}

type Calculator_AverageServer interface {
	SendAndClose(*ResponseDouble) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type calculatorAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorAverageServer) SendAndClose(m *ResponseDouble) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorAverageServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Calculator_Maximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).Maximum(&calculatorMaximumServer{stream})
}

type Calculator_MaximumServer interface {
	Send(*ResponseMaximum) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type calculatorMaximumServer struct {
	grpc.ServerStream
}

func (x *calculatorMaximumServer) Send(m *ResponseMaximum) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorMaximumServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calc",
			Handler:    _Calculator_Calc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Decomposition",
			Handler:       _Calculator_Decomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _Calculator_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Maximum",
			Handler:       _Calculator_Maximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/calculator/calculator.proto",
}
