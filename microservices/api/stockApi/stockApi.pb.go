// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: microservices/api/stockApi/stockApi.proto

package stockApi

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

type AddProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*types.Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *AddProductsRequest) Reset() {
	*x = AddProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductsRequest) ProtoMessage() {}

func (x *AddProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductsRequest.ProtoReflect.Descriptor instead.
func (*AddProductsRequest) Descriptor() ([]byte, []int) {
	return file_microservices_api_stockApi_stockApi_proto_rawDescGZIP(), []int{0}
}

func (x *AddProductsRequest) GetProducts() []*types.Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type AddProductsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductIds []uint32 `protobuf:"varint,1,rep,packed,name=productIds,proto3" json:"productIds,omitempty"`
}

func (x *AddProductsReply) Reset() {
	*x = AddProductsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductsReply) ProtoMessage() {}

func (x *AddProductsReply) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductsReply.ProtoReflect.Descriptor instead.
func (*AddProductsReply) Descriptor() ([]byte, []int) {
	return file_microservices_api_stockApi_stockApi_proto_rawDescGZIP(), []int{1}
}

func (x *AddProductsReply) GetProductIds() []uint32 {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

type GetProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductIds []uint32 `protobuf:"varint,1,rep,packed,name=productIds,proto3" json:"productIds,omitempty"`
}

func (x *GetProductsRequest) Reset() {
	*x = GetProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsRequest) ProtoMessage() {}

func (x *GetProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsRequest.ProtoReflect.Descriptor instead.
func (*GetProductsRequest) Descriptor() ([]byte, []int) {
	return file_microservices_api_stockApi_stockApi_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductsRequest) GetProductIds() []uint32 {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

type GetProductsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*types.Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *GetProductsReply) Reset() {
	*x = GetProductsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsReply) ProtoMessage() {}

func (x *GetProductsReply) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsReply.ProtoReflect.Descriptor instead.
func (*GetProductsReply) Descriptor() ([]byte, []int) {
	return file_microservices_api_stockApi_stockApi_proto_rawDescGZIP(), []int{3}
}

func (x *GetProductsReply) GetProducts() []*types.Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type RemoveProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint32 `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
}

func (x *RemoveProductRequest) Reset() {
	*x = RemoveProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProductRequest) ProtoMessage() {}

func (x *RemoveProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProductRequest.ProtoReflect.Descriptor instead.
func (*RemoveProductRequest) Descriptor() ([]byte, []int) {
	return file_microservices_api_stockApi_stockApi_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveProductRequest) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type RemoveProductReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *types.Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *RemoveProductReply) Reset() {
	*x = RemoveProductReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProductReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProductReply) ProtoMessage() {}

func (x *RemoveProductReply) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProductReply.ProtoReflect.Descriptor instead.
func (*RemoveProductReply) Descriptor() ([]byte, []int) {
	return file_microservices_api_stockApi_stockApi_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveProductReply) GetProduct() *types.Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type OrderProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId  uint32            `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Products map[uint32]uint32 `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *OrderProductsRequest) Reset() {
	*x = OrderProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProductsRequest) ProtoMessage() {}

func (x *OrderProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProductsRequest.ProtoReflect.Descriptor instead.
func (*OrderProductsRequest) Descriptor() ([]byte, []int) {
	return file_microservices_api_stockApi_stockApi_proto_rawDescGZIP(), []int{6}
}

func (x *OrderProductsRequest) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderProductsRequest) GetProducts() map[uint32]uint32 {
	if x != nil {
		return x.Products
	}
	return nil
}

type OrderProductsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Received bool `protobuf:"varint,1,opt,name=received,proto3" json:"received,omitempty"`
}

func (x *OrderProductsReply) Reset() {
	*x = OrderProductsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProductsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProductsReply) ProtoMessage() {}

func (x *OrderProductsReply) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProductsReply.ProtoReflect.Descriptor instead.
func (*OrderProductsReply) Descriptor() ([]byte, []int) {
	return file_microservices_api_stockApi_stockApi_proto_rawDescGZIP(), []int{7}
}

func (x *OrderProductsReply) GetReceived() bool {
	if x != nil {
		return x.Received
	}
	return false
}

type DecreaseProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint32 `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Amount    uint32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DecreaseProductRequest) Reset() {
	*x = DecreaseProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecreaseProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseProductRequest) ProtoMessage() {}

func (x *DecreaseProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseProductRequest.ProtoReflect.Descriptor instead.
func (*DecreaseProductRequest) Descriptor() ([]byte, []int) {
	return file_microservices_api_stockApi_stockApi_proto_rawDescGZIP(), []int{8}
}

func (x *DecreaseProductRequest) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *DecreaseProductRequest) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type DecreaseProductReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *types.Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *DecreaseProductReply) Reset() {
	*x = DecreaseProductReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecreaseProductReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseProductReply) ProtoMessage() {}

func (x *DecreaseProductReply) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_api_stockApi_stockApi_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseProductReply.ProtoReflect.Descriptor instead.
func (*DecreaseProductReply) Descriptor() ([]byte, []int) {
	return file_microservices_api_stockApi_stockApi_proto_rawDescGZIP(), []int{9}
}

func (x *DecreaseProductReply) GetProduct() *types.Product {
	if x != nil {
		return x.Product
	}
	return nil
}

var File_microservices_api_stockApi_stockApi_proto protoreflect.FileDescriptor

var file_microservices_api_stockApi_stockApi_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x41, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x41, 0x70, 0x69, 0x1a, 0x23, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x12, 0x41, 0x64,
	0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x32, 0x0a, 0x10,
	0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73,
	0x22, 0x34, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x22, 0x3e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x34, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x12,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0xb7, 0x01, 0x0a,
	0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x48, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x30, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x22, 0x4e, 0x0a, 0x16, 0x44, 0x65, 0x63, 0x72,
	0x65, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x40, 0x0a, 0x14, 0x44, 0x65, 0x63, 0x72,
	0x65, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x28, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6c, 0x72, 0x7a, 0x2e, 0x64, 0x65, 0x2f, 0x76, 0x73, 0x73, 0x2f,
	0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x6f, 0x62, 0x2d, 0x32, 0x33, 0x73, 0x73,
	0x2f, 0x62, 0x6c, 0x61, 0x74, 0x74, 0x2d, 0x32, 0x2f, 0x62, 0x6c, 0x61, 0x74, 0x74, 0x32, 0x2d,
	0x67, 0x72, 0x70, 0x30, 0x36, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x41, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_microservices_api_stockApi_stockApi_proto_rawDescOnce sync.Once
	file_microservices_api_stockApi_stockApi_proto_rawDescData = file_microservices_api_stockApi_stockApi_proto_rawDesc
)

func file_microservices_api_stockApi_stockApi_proto_rawDescGZIP() []byte {
	file_microservices_api_stockApi_stockApi_proto_rawDescOnce.Do(func() {
		file_microservices_api_stockApi_stockApi_proto_rawDescData = protoimpl.X.CompressGZIP(file_microservices_api_stockApi_stockApi_proto_rawDescData)
	})
	return file_microservices_api_stockApi_stockApi_proto_rawDescData
}

var file_microservices_api_stockApi_stockApi_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_microservices_api_stockApi_stockApi_proto_goTypes = []interface{}{
	(*AddProductsRequest)(nil),     // 0: stockApi.AddProductsRequest
	(*AddProductsReply)(nil),       // 1: stockApi.AddProductsReply
	(*GetProductsRequest)(nil),     // 2: stockApi.GetProductsRequest
	(*GetProductsReply)(nil),       // 3: stockApi.GetProductsReply
	(*RemoveProductRequest)(nil),   // 4: stockApi.RemoveProductRequest
	(*RemoveProductReply)(nil),     // 5: stockApi.RemoveProductReply
	(*OrderProductsRequest)(nil),   // 6: stockApi.OrderProductsRequest
	(*OrderProductsReply)(nil),     // 7: stockApi.OrderProductsReply
	(*DecreaseProductRequest)(nil), // 8: stockApi.DecreaseProductRequest
	(*DecreaseProductReply)(nil),   // 9: stockApi.DecreaseProductReply
	nil,                            // 10: stockApi.OrderProductsRequest.ProductsEntry
	(*types.Product)(nil),          // 11: types.Product
}
var file_microservices_api_stockApi_stockApi_proto_depIdxs = []int32{
	11, // 0: stockApi.AddProductsRequest.products:type_name -> types.Product
	11, // 1: stockApi.GetProductsReply.products:type_name -> types.Product
	11, // 2: stockApi.RemoveProductReply.product:type_name -> types.Product
	10, // 3: stockApi.OrderProductsRequest.products:type_name -> stockApi.OrderProductsRequest.ProductsEntry
	11, // 4: stockApi.DecreaseProductReply.product:type_name -> types.Product
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_microservices_api_stockApi_stockApi_proto_init() }
func file_microservices_api_stockApi_stockApi_proto_init() {
	if File_microservices_api_stockApi_stockApi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_microservices_api_stockApi_stockApi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProductsRequest); i {
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
		file_microservices_api_stockApi_stockApi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProductsReply); i {
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
		file_microservices_api_stockApi_stockApi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductsRequest); i {
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
		file_microservices_api_stockApi_stockApi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductsReply); i {
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
		file_microservices_api_stockApi_stockApi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProductRequest); i {
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
		file_microservices_api_stockApi_stockApi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProductReply); i {
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
		file_microservices_api_stockApi_stockApi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderProductsRequest); i {
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
		file_microservices_api_stockApi_stockApi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderProductsReply); i {
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
		file_microservices_api_stockApi_stockApi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecreaseProductRequest); i {
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
		file_microservices_api_stockApi_stockApi_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecreaseProductReply); i {
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
			RawDescriptor: file_microservices_api_stockApi_stockApi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_microservices_api_stockApi_stockApi_proto_goTypes,
		DependencyIndexes: file_microservices_api_stockApi_stockApi_proto_depIdxs,
		MessageInfos:      file_microservices_api_stockApi_stockApi_proto_msgTypes,
	}.Build()
	File_microservices_api_stockApi_stockApi_proto = out.File
	file_microservices_api_stockApi_stockApi_proto_rawDesc = nil
	file_microservices_api_stockApi_stockApi_proto_goTypes = nil
	file_microservices_api_stockApi_stockApi_proto_depIdxs = nil
}
