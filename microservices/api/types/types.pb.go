// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: microservices/api/types/types.proto

package types

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
	return file_microservices_api_types_types_proto_enumTypes[0].Descriptor()
}

func (DELIVERY_STATUS) Type() protoreflect.EnumType {
	return &file_microservices_api_types_types_proto_enumTypes[0]
}

func (x DELIVERY_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DELIVERY_STATUS.Descriptor instead.
func (DELIVERY_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_microservices_api_types_types_proto_rawDescGZIP(), []int{0}
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
		mi := &file_microservices_api_types_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_types_types_proto_msgTypes[0]
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
	return file_microservices_api_types_types_proto_rawDescGZIP(), []int{0}
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
		mi := &file_microservices_api_types_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_types_types_proto_msgTypes[1]
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
	return file_microservices_api_types_types_proto_rawDescGZIP(), []int{1}
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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer       uint32            `protobuf:"varint,1,opt,name=customer,proto3" json:"customer,omitempty"`
	Products       map[uint32]uint32 `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ProductsStatus map[uint32]bool   `protobuf:"bytes,3,rep,name=productsStatus,proto3" json:"productsStatus,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	OrderStatus    bool              `protobuf:"varint,4,opt,name=orderStatus,proto3" json:"orderStatus,omitempty"`
	PaymentStatus  bool              `protobuf:"varint,5,opt,name=paymentStatus,proto3" json:"paymentStatus,omitempty"`
	DeliveryStatus DELIVERY_STATUS   `protobuf:"varint,6,opt,name=deliveryStatus,proto3,enum=types.DELIVERY_STATUS" json:"deliveryStatus,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_types_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_types_types_proto_msgTypes[2]
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
	return file_microservices_api_types_types_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetCustomer() uint32 {
	if x != nil {
		return x.Customer
	}
	return 0
}

func (x *Order) GetProducts() map[uint32]uint32 {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *Order) GetProductsStatus() map[uint32]bool {
	if x != nil {
		return x.ProductsStatus
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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   uint32  `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Amount      uint32  `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Supplier    uint32  `protobuf:"varint,6,opt,name=supplier,proto3" json:"supplier,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_types_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_types_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_microservices_api_types_types_proto_rawDescGZIP(), []int{3}
}

func (x *Product) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Product) GetSupplier() uint32 {
	if x != nil {
		return x.Supplier
	}
	return 0
}

type Supplier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplierId uint32     `protobuf:"varint,1,opt,name=supplierId,proto3" json:"supplierId,omitempty"`
	Name       string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address    *Address   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Products   []*Product `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *Supplier) Reset() {
	*x = Supplier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_types_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Supplier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Supplier) ProtoMessage() {}

func (x *Supplier) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_types_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Supplier.ProtoReflect.Descriptor instead.
func (*Supplier) Descriptor() ([]byte, []int) {
	return file_microservices_api_types_types_proto_rawDescGZIP(), []int{4}
}

func (x *Supplier) GetSupplierId() uint32 {
	if x != nil {
		return x.SupplierId
	}
	return 0
}

func (x *Supplier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Supplier) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Supplier) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
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
		mi := &file_microservices_api_types_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deliverer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deliverer) ProtoMessage() {}

func (x *Deliverer) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_types_types_proto_msgTypes[5]
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
	return file_microservices_api_types_types_proto_rawDescGZIP(), []int{5}
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

var File_microservices_api_types_types_proto protoreflect.FileDescriptor

var file_microservices_api_types_types_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x61, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x7a, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x7a, 0x69,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22,
	0x48, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x28, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xad, 0x03, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x36, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x48, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x0e, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45,
	0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x41, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa7, 0x01, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x22, 0x94, 0x01, 0x0a, 0x08, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x6b, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2a, 0x3d, 0x0a, 0x0f, 0x44, 0x45, 0x4c, 0x49, 0x56,
	0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f,
	0x54, 0x5f, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45,
	0x52, 0x5f, 0x57, 0x41, 0x59, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49, 0x56,
	0x45, 0x52, 0x45, 0x44, 0x10, 0x02, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x6c, 0x72, 0x7a, 0x2e, 0x64, 0x65, 0x2f, 0x76, 0x73, 0x73, 0x2f, 0x73, 0x65, 0x6d, 0x65,
	0x73, 0x74, 0x65, 0x72, 0x2f, 0x6f, 0x62, 0x2d, 0x32, 0x33, 0x73, 0x73, 0x2f, 0x62, 0x6c, 0x61,
	0x74, 0x74, 0x2d, 0x32, 0x2f, 0x62, 0x6c, 0x61, 0x74, 0x74, 0x32, 0x2d, 0x67, 0x72, 0x70, 0x30,
	0x36, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_microservices_api_types_types_proto_rawDescOnce sync.Once
	file_microservices_api_types_types_proto_rawDescData = file_microservices_api_types_types_proto_rawDesc
)

func file_microservices_api_types_types_proto_rawDescGZIP() []byte {
	file_microservices_api_types_types_proto_rawDescOnce.Do(func() {
		file_microservices_api_types_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_microservices_api_types_types_proto_rawDescData)
	})
	return file_microservices_api_types_types_proto_rawDescData
}

var file_microservices_api_types_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_microservices_api_types_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_microservices_api_types_types_proto_goTypes = []interface{}{
	(DELIVERY_STATUS)(0), // 0: types.DELIVERY_STATUS
	(*Address)(nil),      // 1: types.Address
	(*Customer)(nil),     // 2: types.Customer
	(*Order)(nil),        // 3: types.Order
	(*Product)(nil),      // 4: types.Product
	(*Supplier)(nil),     // 5: types.Supplier
	(*Deliverer)(nil),    // 6: types.Deliverer
	nil,                  // 7: types.Order.ProductsEntry
	nil,                  // 8: types.Order.ProductsStatusEntry
}
var file_microservices_api_types_types_proto_depIdxs = []int32{
	1, // 0: types.Customer.address:type_name -> types.Address
	7, // 1: types.Order.products:type_name -> types.Order.ProductsEntry
	8, // 2: types.Order.productsStatus:type_name -> types.Order.ProductsStatusEntry
	0, // 3: types.Order.deliveryStatus:type_name -> types.DELIVERY_STATUS
	1, // 4: types.Supplier.address:type_name -> types.Address
	4, // 5: types.Supplier.products:type_name -> types.Product
	1, // 6: types.Deliverer.address:type_name -> types.Address
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_microservices_api_types_types_proto_init() }
func file_microservices_api_types_types_proto_init() {
	if File_microservices_api_types_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_microservices_api_types_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_microservices_api_types_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_microservices_api_types_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_microservices_api_types_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_microservices_api_types_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Supplier); i {
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
		file_microservices_api_types_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_microservices_api_types_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_microservices_api_types_types_proto_goTypes,
		DependencyIndexes: file_microservices_api_types_types_proto_depIdxs,
		EnumInfos:         file_microservices_api_types_types_proto_enumTypes,
		MessageInfos:      file_microservices_api_types_types_proto_msgTypes,
	}.Build()
	File_microservices_api_types_types_proto = out.File
	file_microservices_api_types_types_proto_rawDesc = nil
	file_microservices_api_types_types_proto_goTypes = nil
	file_microservices_api_types_types_proto_depIdxs = nil
}
