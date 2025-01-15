// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: payment/rpc/payment.proto

package payment

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

type CreditCardInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreditCardNumber          string `protobuf:"bytes,1,opt,name=credit_card_number,json=creditCardNumber,proto3" json:"credit_card_number,omitempty"`
	CreditCardCvv             int32  `protobuf:"varint,2,opt,name=credit_card_cvv,json=creditCardCvv,proto3" json:"credit_card_cvv,omitempty"`
	CreditCardExpirationYear  int32  `protobuf:"varint,3,opt,name=credit_card_expiration_year,json=creditCardExpirationYear,proto3" json:"credit_card_expiration_year,omitempty"`
	CreditCardExpirationMonth int32  `protobuf:"varint,4,opt,name=credit_card_expiration_month,json=creditCardExpirationMonth,proto3" json:"credit_card_expiration_month,omitempty"`
}

func (x *CreditCardInfo) Reset() {
	*x = CreditCardInfo{}
	mi := &file_payment_rpc_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreditCardInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditCardInfo) ProtoMessage() {}

func (x *CreditCardInfo) ProtoReflect() protoreflect.Message {
	mi := &file_payment_rpc_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditCardInfo.ProtoReflect.Descriptor instead.
func (*CreditCardInfo) Descriptor() ([]byte, []int) {
	return file_payment_rpc_payment_proto_rawDescGZIP(), []int{0}
}

func (x *CreditCardInfo) GetCreditCardNumber() string {
	if x != nil {
		return x.CreditCardNumber
	}
	return ""
}

func (x *CreditCardInfo) GetCreditCardCvv() int32 {
	if x != nil {
		return x.CreditCardCvv
	}
	return 0
}

func (x *CreditCardInfo) GetCreditCardExpirationYear() int32 {
	if x != nil {
		return x.CreditCardExpirationYear
	}
	return 0
}

func (x *CreditCardInfo) GetCreditCardExpirationMonth() int32 {
	if x != nil {
		return x.CreditCardExpirationMonth
	}
	return 0
}

type ChargeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount     float32         `protobuf:"fixed32,1,opt,name=amount,proto3" json:"amount,omitempty"`
	CreditCard *CreditCardInfo `protobuf:"bytes,2,opt,name=credit_card,json=creditCard,proto3" json:"credit_card,omitempty"`
	OrderId    uint32          `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId     uint32          `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ChargeReq) Reset() {
	*x = ChargeReq{}
	mi := &file_payment_rpc_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChargeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeReq) ProtoMessage() {}

func (x *ChargeReq) ProtoReflect() protoreflect.Message {
	mi := &file_payment_rpc_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargeReq.ProtoReflect.Descriptor instead.
func (*ChargeReq) Descriptor() ([]byte, []int) {
	return file_payment_rpc_payment_proto_rawDescGZIP(), []int{1}
}

func (x *ChargeReq) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ChargeReq) GetCreditCard() *CreditCardInfo {
	if x != nil {
		return x.CreditCard
	}
	return nil
}

func (x *ChargeReq) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *ChargeReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ChargeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId uint64 `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Success       string `protobuf:"bytes,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ChargeResp) Reset() {
	*x = ChargeResp{}
	mi := &file_payment_rpc_payment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChargeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeResp) ProtoMessage() {}

func (x *ChargeResp) ProtoReflect() protoreflect.Message {
	mi := &file_payment_rpc_payment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargeResp.ProtoReflect.Descriptor instead.
func (*ChargeResp) Descriptor() ([]byte, []int) {
	return file_payment_rpc_payment_proto_rawDescGZIP(), []int{2}
}

func (x *ChargeResp) GetTransactionId() uint64 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

func (x *ChargeResp) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

type CancelPaymentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId uint64 `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"` // 支付交易ID
	OrderId       uint32 `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`                   // 可选，关联订单ID
}

func (x *CancelPaymentReq) Reset() {
	*x = CancelPaymentReq{}
	mi := &file_payment_rpc_payment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelPaymentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelPaymentReq) ProtoMessage() {}

func (x *CancelPaymentReq) ProtoReflect() protoreflect.Message {
	mi := &file_payment_rpc_payment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelPaymentReq.ProtoReflect.Descriptor instead.
func (*CancelPaymentReq) Descriptor() ([]byte, []int) {
	return file_payment_rpc_payment_proto_rawDescGZIP(), []int{3}
}

func (x *CancelPaymentReq) GetTransactionId() uint64 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

func (x *CancelPaymentReq) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

// 取消支付请求和响应： 用于取消特定交易。
type CancelPaymentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 是否取消成功
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // 错误或状态信息
}

func (x *CancelPaymentResp) Reset() {
	*x = CancelPaymentResp{}
	mi := &file_payment_rpc_payment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelPaymentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelPaymentResp) ProtoMessage() {}

func (x *CancelPaymentResp) ProtoReflect() protoreflect.Message {
	mi := &file_payment_rpc_payment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelPaymentResp.ProtoReflect.Descriptor instead.
func (*CancelPaymentResp) Descriptor() ([]byte, []int) {
	return file_payment_rpc_payment_proto_rawDescGZIP(), []int{4}
}

func (x *CancelPaymentResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CancelPaymentResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type QueryPaymentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId uint64 `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"` // 支付交易ID
}

func (x *QueryPaymentReq) Reset() {
	*x = QueryPaymentReq{}
	mi := &file_payment_rpc_payment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryPaymentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPaymentReq) ProtoMessage() {}

func (x *QueryPaymentReq) ProtoReflect() protoreflect.Message {
	mi := &file_payment_rpc_payment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPaymentReq.ProtoReflect.Descriptor instead.
func (*QueryPaymentReq) Descriptor() ([]byte, []int) {
	return file_payment_rpc_payment_proto_rawDescGZIP(), []int{5}
}

func (x *QueryPaymentReq) GetTransactionId() uint64 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

// 查询支付状态： 用于检查支付交易的当前状态（成功、失败、处理中）。
type QueryPaymentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`                   // 支付状态，例如 "success", "failed", "pending"
	Amount  float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`                 // 支付金额
	OrderId uint32  `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"` // 关联订单ID
}

func (x *QueryPaymentResp) Reset() {
	*x = QueryPaymentResp{}
	mi := &file_payment_rpc_payment_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryPaymentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPaymentResp) ProtoMessage() {}

func (x *QueryPaymentResp) ProtoReflect() protoreflect.Message {
	mi := &file_payment_rpc_payment_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPaymentResp.ProtoReflect.Descriptor instead.
func (*QueryPaymentResp) Descriptor() ([]byte, []int) {
	return file_payment_rpc_payment_proto_rawDescGZIP(), []int{6}
}

func (x *QueryPaymentResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *QueryPaymentResp) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *QueryPaymentResp) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

var File_payment_rpc_payment_proto protoreflect.FileDescriptor

var file_payment_rpc_payment_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0xe6, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43,
	0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x76, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x43, 0x76, 0x76, 0x12, 0x3d, 0x0a,
	0x1b, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x18, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x12, 0x3f, 0x0a, 0x1c,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x19, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x22, 0x91, 0x01,
	0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x4d, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x54, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x38, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x10, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x32, 0xd6, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x43,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x48, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_rpc_payment_proto_rawDescOnce sync.Once
	file_payment_rpc_payment_proto_rawDescData = file_payment_rpc_payment_proto_rawDesc
)

func file_payment_rpc_payment_proto_rawDescGZIP() []byte {
	file_payment_rpc_payment_proto_rawDescOnce.Do(func() {
		file_payment_rpc_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_rpc_payment_proto_rawDescData)
	})
	return file_payment_rpc_payment_proto_rawDescData
}

var file_payment_rpc_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_payment_rpc_payment_proto_goTypes = []any{
	(*CreditCardInfo)(nil),    // 0: payment.CreditCardInfo
	(*ChargeReq)(nil),         // 1: payment.ChargeReq
	(*ChargeResp)(nil),        // 2: payment.ChargeResp
	(*CancelPaymentReq)(nil),  // 3: payment.CancelPaymentReq
	(*CancelPaymentResp)(nil), // 4: payment.CancelPaymentResp
	(*QueryPaymentReq)(nil),   // 5: payment.QueryPaymentReq
	(*QueryPaymentResp)(nil),  // 6: payment.QueryPaymentResp
}
var file_payment_rpc_payment_proto_depIdxs = []int32{
	0, // 0: payment.ChargeReq.credit_card:type_name -> payment.CreditCardInfo
	1, // 1: payment.PaymentService.Charge:input_type -> payment.ChargeReq
	3, // 2: payment.PaymentService.CancelPayment:input_type -> payment.CancelPaymentReq
	5, // 3: payment.PaymentService.QueryPayment:input_type -> payment.QueryPaymentReq
	2, // 4: payment.PaymentService.Charge:output_type -> payment.ChargeResp
	4, // 5: payment.PaymentService.CancelPayment:output_type -> payment.CancelPaymentResp
	6, // 6: payment.PaymentService.QueryPayment:output_type -> payment.QueryPaymentResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payment_rpc_payment_proto_init() }
func file_payment_rpc_payment_proto_init() {
	if File_payment_rpc_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payment_rpc_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_rpc_payment_proto_goTypes,
		DependencyIndexes: file_payment_rpc_payment_proto_depIdxs,
		MessageInfos:      file_payment_rpc_payment_proto_msgTypes,
	}.Build()
	File_payment_rpc_payment_proto = out.File
	file_payment_rpc_payment_proto_rawDesc = nil
	file_payment_rpc_payment_proto_goTypes = nil
	file_payment_rpc_payment_proto_depIdxs = nil
}
