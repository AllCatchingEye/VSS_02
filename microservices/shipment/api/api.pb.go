// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: microservices/shipment/api/api.proto

package shipmentApi

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

type DELIVERY_STATUS int32

const (
	DELIVERY_STATUS_NOT_SENT  DELIVERY_STATUS = 0
	DELIVERY_STATUS_UNDER_WAY DELIVERY_STATUS = 1
	DELIVERY_STATUS_DELIVERED DELIVERY_STATUS = 2
)

// Enum value maps for DELIVERY_STATUS.
var (
	DELIVERY_STATUS_name = map[int32]string{
		0: "NOT_SENT",
		1: "UNDER_WAY",
		2: "DELIVERED",
	}
	DELIVERY_STATUS_value = map[string]int32{
		"NOT_SENT":  0,
		"UNDER_WAY": 1,
		"DELIVERED": 2,
	}
)

func (x DELIVERY_STATUS) Enum() *DELIVERY_STATUS {
	p := new(DELIVERY_STATUS)
	*p = x
	return p
}

func (x DELIVERY_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DELIVERY_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_microservices_shipment_api_api_proto_enumTypes[0].Descriptor()
}

func (DELIVERY_STATUS) Type() protoreflect.EnumType {
	return &file_microservices_shipment_api_api_proto_enumTypes[0]
}

func (x DELIVERY_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DELIVERY_STATUS.Descriptor instead.
func (DELIVERY_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_microservices_shipment_api_api_proto_rawDescGZIP(), []int{0}
}

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address *Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_shipment_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_shipment_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_microservices_shipment_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street  string `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	Zip     string `protobuf:"bytes,3,opt,name=zip,proto3" json:"zip,omitempty"`
	City    string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Country string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_shipment_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_shipment_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_microservices_shipment_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer       *Customer        `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	Products       map[uint32]int32 `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	OrderStatus    bool             `protobuf:"varint,3,opt,name=orderStatus,proto3" json:"orderStatus,omitempty"`
	PaymentStatus  bool             `protobuf:"varint,4,opt,name=paymentStatus,proto3" json:"paymentStatus,omitempty"`
	DeliveryStatus DELIVERY_STATUS  `protobuf:"varint,5,opt,name=deliveryStatus,proto3,enum=shipmentApi.DELIVERY_STATUS" json:"deliveryStatus,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_shipment_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_shipment_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_microservices_shipment_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *Order) GetProducts() map[uint32]int32 {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *Order) GetOrderStatus() bool {
	if x != nil {
		return x.OrderStatus
	}
	return false
}

func (x *Order) GetPaymentStatus() bool {
	if x != nil {
		return x.PaymentStatus
	}
	return false
}

func (x *Order) GetDeliveryStatus() DELIVERY_STATUS {
	if x != nil {
		return x.DeliveryStatus
	}
	return DELIVERY_STATUS_NOT_SENT
}

type Deliverer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelivererId uint32   `protobuf:"varint,1,opt,name=delivererId,proto3" json:"delivererId,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address     *Address `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Deliverer) Reset() {
	*x = Deliverer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_shipment_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deliverer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deliverer) ProtoMessage() {}

func (x *Deliverer) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_shipment_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deliverer.ProtoReflect.Descriptor instead.
func (*Deliverer) Descriptor() ([]byte, []int) {
	return file_microservices_shipment_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *Deliverer) GetDelivererId() uint32 {
	if x != nil {
		return x.DelivererId
	}
	return 0
}

func (x *Deliverer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Deliverer) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

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
		mi := &file_microservices_shipment_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipMyOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipMyOrderRequest) ProtoMessage() {}

func (x *ShipMyOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_shipment_api_api_proto_msgTypes[4]
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
	return file_microservices_shipment_api_api_proto_rawDescGZIP(), []int{4}
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

	OrderId uint32   `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Address *Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ShipMyOrderReply) Reset() {
	*x = ShipMyOrderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_shipment_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipMyOrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipMyOrderReply) ProtoMessage() {}

func (x *ShipMyOrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_shipment_api_api_proto_msgTypes[5]
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
	return file_microservices_shipment_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *ShipMyOrderReply) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *ShipMyOrderReply) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

var File_microservices_shipment_api_api_proto protoreflect.FileDescriptor

var file_microservices_shipment_api_api_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x70, 0x69, 0x22, 0x4e, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x61, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x7a, 0x69, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x7a, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0xc3, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x31, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x44, 0x0a, 0x0e, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x2e,
	0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52,
	0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a,
	0x3b, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x71, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2e, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x4e, 0x0a, 0x12, 0x53, 0x68, 0x69, 0x70, 0x4d, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x5c, 0x0a, 0x10, 0x53, 0x68, 0x69, 0x70, 0x4d, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2a, 0x3d, 0x0a,
	0x0f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x57, 0x41, 0x59, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x02, 0x32, 0x64, 0x0a, 0x0f,
	0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x51, 0x0a, 0x0d, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x1f, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x2e, 0x53,
	0x68, 0x69, 0x70, 0x4d, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x2e,
	0x53, 0x68, 0x69, 0x70, 0x4d, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x3b, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_microservices_shipment_api_api_proto_rawDescOnce sync.Once
	file_microservices_shipment_api_api_proto_rawDescData = file_microservices_shipment_api_api_proto_rawDesc
)

func file_microservices_shipment_api_api_proto_rawDescGZIP() []byte {
	file_microservices_shipment_api_api_proto_rawDescOnce.Do(func() {
		file_microservices_shipment_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_microservices_shipment_api_api_proto_rawDescData)
	})
	return file_microservices_shipment_api_api_proto_rawDescData
}

var file_microservices_shipment_api_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_microservices_shipment_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_microservices_shipment_api_api_proto_goTypes = []interface{}{
	(DELIVERY_STATUS)(0),       // 0: shipmentApi.DELIVERY_STATUS
	(*Customer)(nil),           // 1: shipmentApi.Customer
	(*Address)(nil),            // 2: shipmentApi.Address
	(*Order)(nil),              // 3: shipmentApi.Order
	(*Deliverer)(nil),          // 4: shipmentApi.Deliverer
	(*ShipMyOrderRequest)(nil), // 5: shipmentApi.ShipMyOrderRequest
	(*ShipMyOrderReply)(nil),   // 6: shipmentApi.ShipMyOrderReply
	nil,                        // 7: shipmentApi.Order.ProductsEntry
}
var file_microservices_shipment_api_api_proto_depIdxs = []int32{
	2, // 0: shipmentApi.Customer.address:type_name -> shipmentApi.Address
	1, // 1: shipmentApi.Order.customer:type_name -> shipmentApi.Customer
	7, // 2: shipmentApi.Order.products:type_name -> shipmentApi.Order.ProductsEntry
	0, // 3: shipmentApi.Order.deliveryStatus:type_name -> shipmentApi.DELIVERY_STATUS
	2, // 4: shipmentApi.Deliverer.address:type_name -> shipmentApi.Address
	2, // 5: shipmentApi.ShipMyOrderReply.address:type_name -> shipmentApi.Address
	5, // 6: shipmentApi.ShipmentService.ShipmentOrder:input_type -> shipmentApi.ShipMyOrderRequest
	6, // 7: shipmentApi.ShipmentService.ShipmentOrder:output_type -> shipmentApi.ShipMyOrderReply
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_microservices_shipment_api_api_proto_init() }
func file_microservices_shipment_api_api_proto_init() {
	if File_microservices_shipment_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_microservices_shipment_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_microservices_shipment_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_microservices_shipment_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_microservices_shipment_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deliverer); i {
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
		file_microservices_shipment_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_microservices_shipment_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_microservices_shipment_api_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_microservices_shipment_api_api_proto_goTypes,
		DependencyIndexes: file_microservices_shipment_api_api_proto_depIdxs,
		EnumInfos:         file_microservices_shipment_api_api_proto_enumTypes,
		MessageInfos:      file_microservices_shipment_api_api_proto_msgTypes,
	}.Build()
	File_microservices_shipment_api_api_proto = out.File
	file_microservices_shipment_api_api_proto_rawDesc = nil
	file_microservices_shipment_api_api_proto_goTypes = nil
	file_microservices_shipment_api_api_proto_depIdxs = nil
}
