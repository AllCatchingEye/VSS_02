// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: microservices/api/shipmentApi/shipmentApi.proto

package shipmentApi

import (
	types "gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/types"
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

type ShipMyOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint32 `protobuf:"varint,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	OrderId    uint32 `protobuf:"varint,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *ShipMyOrderRequest) Reset() {
	*x = ShipMyOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipMyOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipMyOrderRequest) ProtoMessage() {}

func (x *ShipMyOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipMyOrderRequest.ProtoReflect.Descriptor instead.
func (*ShipMyOrderRequest) Descriptor() ([]byte, []int) {
	return file_microservices_api_shipmentApi_shipmentApi_proto_rawDescGZIP(), []int{0}
}

func (x *ShipMyOrderRequest) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *ShipMyOrderRequest) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type ShipMyOrderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint32         `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Address *types.Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ShipMyOrderReply) Reset() {
	*x = ShipMyOrderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipMyOrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipMyOrderReply) ProtoMessage() {}

func (x *ShipMyOrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipMyOrderReply.ProtoReflect.Descriptor instead.
func (*ShipMyOrderReply) Descriptor() ([]byte, []int) {
	return file_microservices_api_shipmentApi_shipmentApi_proto_rawDescGZIP(), []int{1}
}

func (x *ShipMyOrderReply) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *ShipMyOrderReply) GetAddress() *types.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type IsOrderShippedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint32 `protobuf:"varint,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	OrderId    uint32 `protobuf:"varint,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *IsOrderShippedRequest) Reset() {
	*x = IsOrderShippedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsOrderShippedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsOrderShippedRequest) ProtoMessage() {}

func (x *IsOrderShippedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsOrderShippedRequest.ProtoReflect.Descriptor instead.
func (*IsOrderShippedRequest) Descriptor() ([]byte, []int) {
	return file_microservices_api_shipmentApi_shipmentApi_proto_rawDescGZIP(), []int{2}
}

func (x *IsOrderShippedRequest) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *IsOrderShippedRequest) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type IsOrderShippedReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsShipped bool `protobuf:"varint,1,opt,name=isShipped,proto3" json:"isShipped,omitempty"`
}

func (x *IsOrderShippedReply) Reset() {
	*x = IsOrderShippedReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsOrderShippedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsOrderShippedReply) ProtoMessage() {}

func (x *IsOrderShippedReply) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsOrderShippedReply.ProtoReflect.Descriptor instead.
func (*IsOrderShippedReply) Descriptor() ([]byte, []int) {
	return file_microservices_api_shipmentApi_shipmentApi_proto_rawDescGZIP(), []int{3}
}

func (x *IsOrderShippedReply) GetIsShipped() bool {
	if x != nil {
		return x.IsShipped
	}
	return false
}

type RetoureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint32 `protobuf:"varint,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	OrderId    uint32 `protobuf:"varint,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	WantRefund bool   `protobuf:"varint,3,opt,name=wantRefund,proto3" json:"wantRefund,omitempty"`
	Product    uint32 `protobuf:"varint,4,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *RetoureRequest) Reset() {
	*x = RetoureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetoureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetoureRequest) ProtoMessage() {}

func (x *RetoureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetoureRequest.ProtoReflect.Descriptor instead.
func (*RetoureRequest) Descriptor() ([]byte, []int) {
	return file_microservices_api_shipmentApi_shipmentApi_proto_rawDescGZIP(), []int{4}
}

func (x *RetoureRequest) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *RetoureRequest) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *RetoureRequest) GetWantRefund() bool {
	if x != nil {
		return x.WantRefund
	}
	return false
}

func (x *RetoureRequest) GetProduct() uint32 {
	if x != nil {
		return x.Product
	}
	return 0
}

type RetoureReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RetoureReply) Reset() {
	*x = RetoureReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetoureReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetoureReply) ProtoMessage() {}

func (x *RetoureReply) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetoureReply.ProtoReflect.Descriptor instead.
func (*RetoureReply) Descriptor() ([]byte, []int) {
	return file_microservices_api_shipmentApi_shipmentApi_proto_rawDescGZIP(), []int{5}
}

func (x *RetoureReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_microservices_api_shipmentApi_shipmentApi_proto protoreflect.FileDescriptor

var file_microservices_api_shipmentApi_shipmentApi_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x2f,
	0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x1a, 0x23,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x12, 0x53, 0x68, 0x69, 0x70, 0x4d, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x10, 0x53, 0x68, 0x69, 0x70, 0x4d, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x51, 0x0a, 0x15, 0x49,
	0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x33,
	0x0a, 0x13, 0x49, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x68, 0x69, 0x70, 0x70,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x68, 0x69, 0x70,
	0x70, 0x65, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x74, 0x6f, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x77, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x28, 0x0a, 0x0c, 0x52, 0x65,
	0x74, 0x6f, 0x75, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6c,
	0x72, 0x7a, 0x2e, 0x64, 0x65, 0x2f, 0x76, 0x73, 0x73, 0x2f, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x2f, 0x6f, 0x62, 0x2d, 0x32, 0x33, 0x73, 0x73, 0x2f, 0x62, 0x6c, 0x61, 0x74, 0x74,
	0x2d, 0x32, 0x2f, 0x62, 0x6c, 0x61, 0x74, 0x74, 0x32, 0x2d, 0x67, 0x72, 0x70, 0x30, 0x36, 0x2f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_microservices_api_shipmentApi_shipmentApi_proto_rawDescOnce sync.Once
	file_microservices_api_shipmentApi_shipmentApi_proto_rawDescData = file_microservices_api_shipmentApi_shipmentApi_proto_rawDesc
)

func file_microservices_api_shipmentApi_shipmentApi_proto_rawDescGZIP() []byte {
	file_microservices_api_shipmentApi_shipmentApi_proto_rawDescOnce.Do(func() {
		file_microservices_api_shipmentApi_shipmentApi_proto_rawDescData = protoimpl.X.CompressGZIP(file_microservices_api_shipmentApi_shipmentApi_proto_rawDescData)
	})
	return file_microservices_api_shipmentApi_shipmentApi_proto_rawDescData
}

var file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_microservices_api_shipmentApi_shipmentApi_proto_goTypes = []interface{}{
	(*ShipMyOrderRequest)(nil),    // 0: shipmentApi.ShipMyOrderRequest
	(*ShipMyOrderReply)(nil),      // 1: shipmentApi.ShipMyOrderReply
	(*IsOrderShippedRequest)(nil), // 2: shipmentApi.IsOrderShippedRequest
	(*IsOrderShippedReply)(nil),   // 3: shipmentApi.IsOrderShippedReply
	(*RetoureRequest)(nil),        // 4: shipmentApi.RetoureRequest
	(*RetoureReply)(nil),          // 5: shipmentApi.RetoureReply
	(*types.Address)(nil),         // 6: types.Address
}
var file_microservices_api_shipmentApi_shipmentApi_proto_depIdxs = []int32{
	6, // 0: shipmentApi.ShipMyOrderReply.address:type_name -> types.Address
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_microservices_api_shipmentApi_shipmentApi_proto_init() }
func file_microservices_api_shipmentApi_shipmentApi_proto_init() {
	if File_microservices_api_shipmentApi_shipmentApi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipMyOrderRequest); i {
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
		file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipMyOrderReply); i {
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
		file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsOrderShippedRequest); i {
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
		file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsOrderShippedReply); i {
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
		file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetoureRequest); i {
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
		file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetoureReply); i {
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
			RawDescriptor: file_microservices_api_shipmentApi_shipmentApi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_microservices_api_shipmentApi_shipmentApi_proto_goTypes,
		DependencyIndexes: file_microservices_api_shipmentApi_shipmentApi_proto_depIdxs,
		MessageInfos:      file_microservices_api_shipmentApi_shipmentApi_proto_msgTypes,
	}.Build()
	File_microservices_api_shipmentApi_shipmentApi_proto = out.File
	file_microservices_api_shipmentApi_shipmentApi_proto_rawDesc = nil
	file_microservices_api_shipmentApi_shipmentApi_proto_goTypes = nil
	file_microservices_api_shipmentApi_shipmentApi_proto_depIdxs = nil
}
