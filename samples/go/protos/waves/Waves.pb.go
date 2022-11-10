// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: Waves.proto

package waves

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

// Transfer transaction
type TransferMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// amount
	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// asset ID
	Asset string `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	// minimum 0.001 Waves (100000 Wavelets) for now
	Fee int64 `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
	// asset of the fee
	FeeAsset string `protobuf:"bytes,4,opt,name=fee_asset,json=feeAsset,proto3" json:"fee_asset,omitempty"`
	// destination address
	To string `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	// any 140 bytes payload, will be displayed to the client as utf-8 string
	Attachment []byte `protobuf:"bytes,6,opt,name=attachment,proto3" json:"attachment,omitempty"`
}

func (x *TransferMessage) Reset() {
	*x = TransferMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Waves_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferMessage) ProtoMessage() {}

func (x *TransferMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Waves_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferMessage.ProtoReflect.Descriptor instead.
func (*TransferMessage) Descriptor() ([]byte, []int) {
	return file_Waves_proto_rawDescGZIP(), []int{0}
}

func (x *TransferMessage) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransferMessage) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *TransferMessage) GetFee() int64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *TransferMessage) GetFeeAsset() string {
	if x != nil {
		return x.FeeAsset
	}
	return ""
}

func (x *TransferMessage) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TransferMessage) GetAttachment() []byte {
	if x != nil {
		return x.Attachment
	}
	return nil
}

// Lease transaction
type LeaseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// amount
	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// destination
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// minimum 0.001 Waves (100000 Wavelets) for now
	Fee int64 `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (x *LeaseMessage) Reset() {
	*x = LeaseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Waves_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaseMessage) ProtoMessage() {}

func (x *LeaseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Waves_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaseMessage.ProtoReflect.Descriptor instead.
func (*LeaseMessage) Descriptor() ([]byte, []int) {
	return file_Waves_proto_rawDescGZIP(), []int{1}
}

func (x *LeaseMessage) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *LeaseMessage) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *LeaseMessage) GetFee() int64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

// Cancel Lease transaction
type CancelLeaseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Lease ID to cancel
	LeaseId string `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
	// Fee used
	Fee int64 `protobuf:"varint,2,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (x *CancelLeaseMessage) Reset() {
	*x = CancelLeaseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Waves_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelLeaseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelLeaseMessage) ProtoMessage() {}

func (x *CancelLeaseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Waves_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelLeaseMessage.ProtoReflect.Descriptor instead.
func (*CancelLeaseMessage) Descriptor() ([]byte, []int) {
	return file_Waves_proto_rawDescGZIP(), []int{2}
}

func (x *CancelLeaseMessage) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

func (x *CancelLeaseMessage) GetFee() int64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

// Input data necessary to create a signed transaction.
type SigningInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// in millis
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The secret private key used for signing (32 bytes).
	PrivateKey []byte `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Payload message
	//
	// Types that are assignable to MessageOneof:
	//	*SigningInput_TransferMessage
	//	*SigningInput_LeaseMessage
	//	*SigningInput_CancelLeaseMessage
	MessageOneof isSigningInput_MessageOneof `protobuf_oneof:"message_oneof"`
}

func (x *SigningInput) Reset() {
	*x = SigningInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Waves_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningInput) ProtoMessage() {}

func (x *SigningInput) ProtoReflect() protoreflect.Message {
	mi := &file_Waves_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigningInput.ProtoReflect.Descriptor instead.
func (*SigningInput) Descriptor() ([]byte, []int) {
	return file_Waves_proto_rawDescGZIP(), []int{3}
}

func (x *SigningInput) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *SigningInput) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (m *SigningInput) GetMessageOneof() isSigningInput_MessageOneof {
	if m != nil {
		return m.MessageOneof
	}
	return nil
}

func (x *SigningInput) GetTransferMessage() *TransferMessage {
	if x, ok := x.GetMessageOneof().(*SigningInput_TransferMessage); ok {
		return x.TransferMessage
	}
	return nil
}

func (x *SigningInput) GetLeaseMessage() *LeaseMessage {
	if x, ok := x.GetMessageOneof().(*SigningInput_LeaseMessage); ok {
		return x.LeaseMessage
	}
	return nil
}

func (x *SigningInput) GetCancelLeaseMessage() *CancelLeaseMessage {
	if x, ok := x.GetMessageOneof().(*SigningInput_CancelLeaseMessage); ok {
		return x.CancelLeaseMessage
	}
	return nil
}

type isSigningInput_MessageOneof interface {
	isSigningInput_MessageOneof()
}

type SigningInput_TransferMessage struct {
	TransferMessage *TransferMessage `protobuf:"bytes,3,opt,name=transfer_message,json=transferMessage,proto3,oneof"`
}

type SigningInput_LeaseMessage struct {
	LeaseMessage *LeaseMessage `protobuf:"bytes,4,opt,name=lease_message,json=leaseMessage,proto3,oneof"`
}

type SigningInput_CancelLeaseMessage struct {
	CancelLeaseMessage *CancelLeaseMessage `protobuf:"bytes,5,opt,name=cancel_lease_message,json=cancelLeaseMessage,proto3,oneof"`
}

func (*SigningInput_TransferMessage) isSigningInput_MessageOneof() {}

func (*SigningInput_LeaseMessage) isSigningInput_MessageOneof() {}

func (*SigningInput_CancelLeaseMessage) isSigningInput_MessageOneof() {}

// Result containing the signed and encoded transaction.
type SigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signature data
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// transaction in JSON format
	Json string `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *SigningOutput) Reset() {
	*x = SigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Waves_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningOutput) ProtoMessage() {}

func (x *SigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_Waves_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigningOutput.ProtoReflect.Descriptor instead.
func (*SigningOutput) Descriptor() ([]byte, []int) {
	return file_Waves_proto_rawDescGZIP(), []int{4}
}

func (x *SigningOutput) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SigningOutput) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

var File_Waves_proto protoreflect.FileDescriptor

var file_Waves_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x57, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x54,
	0x57, 0x2e, 0x57, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01,
	0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x66, 0x65,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x65, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x48,
	0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x66, 0x65, 0x65, 0x22, 0x41, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x66, 0x65, 0x65, 0x22, 0xc9, 0x02, 0x0a, 0x0c,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x4c, 0x0a, 0x10, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x54, 0x57, 0x2e, 0x57, 0x61, 0x76, 0x65, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x54, 0x57, 0x2e, 0x57, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00,
	0x52, 0x0c, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x56,
	0x0a, 0x14, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x54,
	0x57, 0x2e, 0x57, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x12, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x69,
	0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x42, 0x17, 0x0a, 0x15, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6a, 0x6e, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Waves_proto_rawDescOnce sync.Once
	file_Waves_proto_rawDescData = file_Waves_proto_rawDesc
)

func file_Waves_proto_rawDescGZIP() []byte {
	file_Waves_proto_rawDescOnce.Do(func() {
		file_Waves_proto_rawDescData = protoimpl.X.CompressGZIP(file_Waves_proto_rawDescData)
	})
	return file_Waves_proto_rawDescData
}

var file_Waves_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Waves_proto_goTypes = []interface{}{
	(*TransferMessage)(nil),    // 0: TW.Waves.Proto.TransferMessage
	(*LeaseMessage)(nil),       // 1: TW.Waves.Proto.LeaseMessage
	(*CancelLeaseMessage)(nil), // 2: TW.Waves.Proto.CancelLeaseMessage
	(*SigningInput)(nil),       // 3: TW.Waves.Proto.SigningInput
	(*SigningOutput)(nil),      // 4: TW.Waves.Proto.SigningOutput
}
var file_Waves_proto_depIdxs = []int32{
	0, // 0: TW.Waves.Proto.SigningInput.transfer_message:type_name -> TW.Waves.Proto.TransferMessage
	1, // 1: TW.Waves.Proto.SigningInput.lease_message:type_name -> TW.Waves.Proto.LeaseMessage
	2, // 2: TW.Waves.Proto.SigningInput.cancel_lease_message:type_name -> TW.Waves.Proto.CancelLeaseMessage
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Waves_proto_init() }
func file_Waves_proto_init() {
	if File_Waves_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Waves_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferMessage); i {
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
		file_Waves_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaseMessage); i {
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
		file_Waves_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelLeaseMessage); i {
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
		file_Waves_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigningInput); i {
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
		file_Waves_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigningOutput); i {
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
	file_Waves_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SigningInput_TransferMessage)(nil),
		(*SigningInput_LeaseMessage)(nil),
		(*SigningInput_CancelLeaseMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Waves_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Waves_proto_goTypes,
		DependencyIndexes: file_Waves_proto_depIdxs,
		MessageInfos:      file_Waves_proto_msgTypes,
	}.Build()
	File_Waves_proto = out.File
	file_Waves_proto_rawDesc = nil
	file_Waves_proto_goTypes = nil
	file_Waves_proto_depIdxs = nil
}
