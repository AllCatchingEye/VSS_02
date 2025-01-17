// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: microservices/api/paymentApi/paymentApi.proto

package paymentApi

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

type PayMyOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint32 `protobuf:"varint,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	OrderId    uint32 `protobuf:"varint,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *PayMyOrderRequest) Reset() {
	*x = PayMyOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_paymentApi_paymentApi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayMyOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayMyOrderRequest) ProtoMessage() {}

func (x *PayMyOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_paymentApi_paymentApi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayMyOrderRequest.ProtoReflect.Descriptor instead.
func (*PayMyOrderRequest) Descriptor() ([]byte, []int) {
	return file_microservices_api_paymentApi_paymentApi_proto_rawDescGZIP(), []int{0}
}

func (x *PayMyOrderRequest) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *PayMyOrderRequest) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

// TODO: add more fields?
type PayMyOrderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint32 `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *PayMyOrderReply) Reset() {
	*x = PayMyOrderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_paymentApi_paymentApi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayMyOrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayMyOrderReply) ProtoMessage() {}

func (x *PayMyOrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_paymentApi_paymentApi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayMyOrderReply.ProtoReflect.Descriptor instead.
func (*PayMyOrderReply) Descriptor() ([]byte, []int) {
	return file_microservices_api_paymentApi_paymentApi_proto_rawDescGZIP(), []int{1}
}

func (x *PayMyOrderReply) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type IsOrderPayedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint32 `protobuf:"varint,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	OrderId    uint32 `protobuf:"varint,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *IsOrderPayedRequest) Reset() {
	*x = IsOrderPayedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_paymentApi_paymentApi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsOrderPayedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsOrderPayedRequest) ProtoMessage() {}

func (x *IsOrderPayedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_paymentApi_paymentApi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsOrderPayedRequest.ProtoReflect.Descriptor instead.
func (*IsOrderPayedRequest) Descriptor() ([]byte, []int) {
	return file_microservices_api_paymentApi_paymentApi_proto_rawDescGZIP(), []int{2}
}

func (x *IsOrderPayedRequest) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *IsOrderPayedRequest) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type IsOrderPayedReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsPayed bool `protobuf:"varint,1,opt,name=isPayed,proto3" json:"isPayed,omitempty"`
}

func (x *IsOrderPayedReply) Reset() {
	*x = IsOrderPayedReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_paymentApi_paymentApi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsOrderPayedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsOrderPayedReply) ProtoMessage() {}

func (x *IsOrderPayedReply) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_paymentApi_paymentApi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsOrderPayedReply.ProtoReflect.Descriptor instead.
func (*IsOrderPayedReply) Descriptor() ([]byte, []int) {
	return file_microservices_api_paymentApi_paymentApi_proto_rawDescGZIP(), []int{3}
}

func (x *IsOrderPayedReply) GetIsPayed() bool {
	if x != nil {
		return x.IsPayed
	}
	return false
}

type RefundMyOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint32 `protobuf:"varint,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	OrderId    uint32 `protobuf:"varint,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *RefundMyOrderRequest) Reset() {
	*x = RefundMyOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_paymentApi_paymentApi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundMyOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundMyOrderRequest) ProtoMessage() {}

func (x *RefundMyOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_paymentApi_paymentApi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundMyOrderRequest.ProtoReflect.Descriptor instead.
func (*RefundMyOrderRequest) Descriptor() ([]byte, []int) {
	return file_microservices_api_paymentApi_paymentApi_proto_rawDescGZIP(), []int{4}
}

func (x *RefundMyOrderRequest) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *RefundMyOrderRequest) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type RefundMyOrderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefundSuccess bool `protobuf:"varint,1,opt,name=refundSuccess,proto3" json:"refundSuccess,omitempty"`
}

func (x *RefundMyOrderReply) Reset() {
	*x = RefundMyOrderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_paymentApi_paymentApi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundMyOrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundMyOrderReply) ProtoMessage() {}

func (x *RefundMyOrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_paymentApi_paymentApi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundMyOrderReply.ProtoReflect.Descriptor instead.
func (*RefundMyOrderReply) Descriptor() ([]byte, []int) {
	return file_microservices_api_paymentApi_paymentApi_proto_rawDescGZIP(), []int{5}
}

func (x *RefundMyOrderReply) GetRefundSuccess() bool {
	if x != nil {
		return x.RefundSuccess
	}
	return false
}

var File_microservices_api_paymentApi_paymentApi_proto protoreflect.FileDescriptor

var file_microservices_api_paymentApi_paymentApi_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x2f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x22, 0x4d, 0x0a, 0x11, 0x50,
	0x61, 0x79, 0x4d, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x50, 0x61,
	0x79, 0x4d, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x13, 0x49, 0x73, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x50, 0x61, 0x79, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x49, 0x73, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x73, 0x50, 0x61, 0x79, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x50, 0x61, 0x79, 0x65, 0x64, 0x22, 0x50, 0x0a, 0x14, 0x52, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x4d, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x12, 0x52, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x4d, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x24, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x6c, 0x72, 0x7a, 0x2e, 0x64, 0x65, 0x2f, 0x76, 0x73, 0x73, 0x2f, 0x73, 0x65, 0x6d, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x2f, 0x6f, 0x62, 0x2d, 0x32, 0x33, 0x73, 0x73, 0x2f, 0x62, 0x6c, 0x61, 0x74,
	0x74, 0x2d, 0x32, 0x2f, 0x62, 0x6c, 0x61, 0x74, 0x74, 0x32, 0x2d, 0x67, 0x72, 0x70, 0x30, 0x36,
	0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_microservices_api_paymentApi_paymentApi_proto_rawDescOnce sync.Once
	file_microservices_api_paymentApi_paymentApi_proto_rawDescData = file_microservices_api_paymentApi_paymentApi_proto_rawDesc
)

func file_microservices_api_paymentApi_paymentApi_proto_rawDescGZIP() []byte {
	file_microservices_api_paymentApi_paymentApi_proto_rawDescOnce.Do(func() {
		file_microservices_api_paymentApi_paymentApi_proto_rawDescData = protoimpl.X.CompressGZIP(file_microservices_api_paymentApi_paymentApi_proto_rawDescData)
	})
	return file_microservices_api_paymentApi_paymentApi_proto_rawDescData
}

var file_microservices_api_paymentApi_paymentApi_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_microservices_api_paymentApi_paymentApi_proto_goTypes = []interface{}{
	(*PayMyOrderRequest)(nil),    // 0: paymentApi.PayMyOrderRequest
	(*PayMyOrderReply)(nil),      // 1: paymentApi.PayMyOrderReply
	(*IsOrderPayedRequest)(nil),  // 2: paymentApi.IsOrderPayedRequest
	(*IsOrderPayedReply)(nil),    // 3: paymentApi.IsOrderPayedReply
	(*RefundMyOrderRequest)(nil), // 4: paymentApi.RefundMyOrderRequest
	(*RefundMyOrderReply)(nil),   // 5: paymentApi.RefundMyOrderReply
}
var file_microservices_api_paymentApi_paymentApi_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_microservices_api_paymentApi_paymentApi_proto_init() }
func file_microservices_api_paymentApi_paymentApi_proto_init() {
	if File_microservices_api_paymentApi_paymentApi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_microservices_api_paymentApi_paymentApi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayMyOrderRequest); i {
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
		file_microservices_api_paymentApi_paymentApi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayMyOrderReply); i {
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
		file_microservices_api_paymentApi_paymentApi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsOrderPayedRequest); i {
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
		file_microservices_api_paymentApi_paymentApi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsOrderPayedReply); i {
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
		file_microservices_api_paymentApi_paymentApi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundMyOrderRequest); i {
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
		file_microservices_api_paymentApi_paymentApi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundMyOrderReply); i {
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
			RawDescriptor: file_microservices_api_paymentApi_paymentApi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_microservices_api_paymentApi_paymentApi_proto_goTypes,
		DependencyIndexes: file_microservices_api_paymentApi_paymentApi_proto_depIdxs,
		MessageInfos:      file_microservices_api_paymentApi_paymentApi_proto_msgTypes,
	}.Build()
	File_microservices_api_paymentApi_paymentApi_proto = out.File
	file_microservices_api_paymentApi_paymentApi_proto_rawDesc = nil
	file_microservices_api_paymentApi_paymentApi_proto_goTypes = nil
	file_microservices_api_paymentApi_paymentApi_proto_depIdxs = nil
}
