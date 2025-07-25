// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.2
// source: balance/balance.proto

package balancev1

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

// Запрос баланса
type GetBalanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	mi := &file_balance_balance_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_balance_balance_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_balance_balance_proto_rawDescGZIP(), []int{0}
}

func (x *GetBalanceRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Ответ с балансом
type GetBalanceResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	CurrentBalance   float64                `protobuf:"fixed64,1,opt,name=current_balance,json=currentBalance,proto3" json:"current_balance,omitempty"`
	WithdrawnBalance float64                `protobuf:"fixed64,2,opt,name=withdrawn_balance,json=withdrawnBalance,proto3" json:"withdrawn_balance,omitempty"`
	TotalEarned      float64                `protobuf:"fixed64,3,opt,name=total_earned,json=totalEarned,proto3" json:"total_earned,omitempty"` // current + withdrawn
	Currency         string                 `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`                            // RUB, USD и т.д.
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	mi := &file_balance_balance_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_balance_balance_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_balance_balance_proto_rawDescGZIP(), []int{1}
}

func (x *GetBalanceResponse) GetCurrentBalance() float64 {
	if x != nil {
		return x.CurrentBalance
	}
	return 0
}

func (x *GetBalanceResponse) GetWithdrawnBalance() float64 {
	if x != nil {
		return x.WithdrawnBalance
	}
	return 0
}

func (x *GetBalanceResponse) GetTotalEarned() float64 {
	if x != nil {
		return x.TotalEarned
	}
	return 0
}

func (x *GetBalanceResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

// Запрос на пополнение
type AddFundsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount        float64                `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency      string                 `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	ReferenceId   string                 `protobuf:"bytes,4,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"` // ID транзакции/реферала для идемпотентности
	PaymentMethod string                 `protobuf:"bytes,5,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddFundsRequest) Reset() {
	*x = AddFundsRequest{}
	mi := &file_balance_balance_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddFundsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFundsRequest) ProtoMessage() {}

func (x *AddFundsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_balance_balance_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFundsRequest.ProtoReflect.Descriptor instead.
func (*AddFundsRequest) Descriptor() ([]byte, []int) {
	return file_balance_balance_proto_rawDescGZIP(), []int{2}
}

func (x *AddFundsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddFundsRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *AddFundsRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *AddFundsRequest) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *AddFundsRequest) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

// Ответ на пополнение
type AddFundsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	TransactionId string                 `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	NewBalance    float64                `protobuf:"fixed64,3,opt,name=new_balance,json=newBalance,proto3" json:"new_balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddFundsResponse) Reset() {
	*x = AddFundsResponse{}
	mi := &file_balance_balance_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddFundsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFundsResponse) ProtoMessage() {}

func (x *AddFundsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_balance_balance_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFundsResponse.ProtoReflect.Descriptor instead.
func (*AddFundsResponse) Descriptor() ([]byte, []int) {
	return file_balance_balance_proto_rawDescGZIP(), []int{3}
}

func (x *AddFundsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddFundsResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *AddFundsResponse) GetNewBalance() float64 {
	if x != nil {
		return x.NewBalance
	}
	return 0
}

// Запрос на вывод
type WithdrawRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount        float64                `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	PaymentMethod string                 `protobuf:"bytes,3,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"` // card, crypto, etc.
	Currency      string                 `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WithdrawRequest) Reset() {
	*x = WithdrawRequest{}
	mi := &file_balance_balance_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawRequest) ProtoMessage() {}

func (x *WithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_balance_balance_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawRequest.ProtoReflect.Descriptor instead.
func (*WithdrawRequest) Descriptor() ([]byte, []int) {
	return file_balance_balance_proto_rawDescGZIP(), []int{4}
}

func (x *WithdrawRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *WithdrawRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *WithdrawRequest) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *WithdrawRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

// Ответ на вывод
type WithdrawResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	TransactionId string                 `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	NewBalance    float64                `protobuf:"fixed64,3,opt,name=new_balance,json=newBalance,proto3" json:"new_balance,omitempty"`
	NewWithdrawn  float64                `protobuf:"fixed64,4,opt,name=new_withdrawn,json=newWithdrawn,proto3" json:"new_withdrawn,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WithdrawResponse) Reset() {
	*x = WithdrawResponse{}
	mi := &file_balance_balance_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WithdrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawResponse) ProtoMessage() {}

func (x *WithdrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_balance_balance_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawResponse.ProtoReflect.Descriptor instead.
func (*WithdrawResponse) Descriptor() ([]byte, []int) {
	return file_balance_balance_proto_rawDescGZIP(), []int{5}
}

func (x *WithdrawResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *WithdrawResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *WithdrawResponse) GetNewBalance() float64 {
	if x != nil {
		return x.NewBalance
	}
	return 0
}

func (x *WithdrawResponse) GetNewWithdrawn() float64 {
	if x != nil {
		return x.NewWithdrawn
	}
	return 0
}

var File_balance_balance_proto protoreflect.FileDescriptor

var file_balance_balance_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61,
	0x6c, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xa9, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x11, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x6e, 0x5f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x77, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0xa8, 0x01, 0x0a, 0x0f,
	0x41, 0x64, 0x64, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x74, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x46, 0x75, 0x6e,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6e,
	0x65, 0x77, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0a, 0x6e, 0x65, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x85, 0x01, 0x0a,
	0x0f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x22, 0x99, 0x01, 0x0a, 0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65,
	0x77, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x6e, 0x65, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6e,
	0x65, 0x77, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x6e,
	0x32, 0xdd, 0x01, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x72, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x46, 0x75, 0x6e, 0x64,
	0x73, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x2e, 0x41, 0x64, 0x64,
	0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x75, 0x6e, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x72, 0x61, 0x6c, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x2e,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x1d, 0x5a, 0x1b, 0x62, 0x79, 0x73, 0x74, 0x72, 0x6f, 0x76, 0x2e, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x31, 0x3b, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_balance_balance_proto_rawDescOnce sync.Once
	file_balance_balance_proto_rawDescData = file_balance_balance_proto_rawDesc
)

func file_balance_balance_proto_rawDescGZIP() []byte {
	file_balance_balance_proto_rawDescOnce.Do(func() {
		file_balance_balance_proto_rawDescData = protoimpl.X.CompressGZIP(file_balance_balance_proto_rawDescData)
	})
	return file_balance_balance_proto_rawDescData
}

var file_balance_balance_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_balance_balance_proto_goTypes = []any{
	(*GetBalanceRequest)(nil),  // 0: referral.GetBalanceRequest
	(*GetBalanceResponse)(nil), // 1: referral.GetBalanceResponse
	(*AddFundsRequest)(nil),    // 2: referral.AddFundsRequest
	(*AddFundsResponse)(nil),   // 3: referral.AddFundsResponse
	(*WithdrawRequest)(nil),    // 4: referral.WithdrawRequest
	(*WithdrawResponse)(nil),   // 5: referral.WithdrawResponse
}
var file_balance_balance_proto_depIdxs = []int32{
	0, // 0: referral.Balance.GetBalance:input_type -> referral.GetBalanceRequest
	2, // 1: referral.Balance.AddFunds:input_type -> referral.AddFundsRequest
	4, // 2: referral.Balance.WithdrawFunds:input_type -> referral.WithdrawRequest
	1, // 3: referral.Balance.GetBalance:output_type -> referral.GetBalanceResponse
	3, // 4: referral.Balance.AddFunds:output_type -> referral.AddFundsResponse
	5, // 5: referral.Balance.WithdrawFunds:output_type -> referral.WithdrawResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_balance_balance_proto_init() }
func file_balance_balance_proto_init() {
	if File_balance_balance_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_balance_balance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_balance_balance_proto_goTypes,
		DependencyIndexes: file_balance_balance_proto_depIdxs,
		MessageInfos:      file_balance_balance_proto_msgTypes,
	}.Build()
	File_balance_balance_proto = out.File
	file_balance_balance_proto_rawDesc = nil
	file_balance_balance_proto_goTypes = nil
	file_balance_balance_proto_depIdxs = nil
}
