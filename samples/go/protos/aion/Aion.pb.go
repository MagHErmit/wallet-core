// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: Aion.proto

package aion

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

// Input data necessary to create a signed transaction.
type SigningInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Nonce (uint256, serialized little endian)
	Nonce []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Gas price (uint256, serialized little endian)
	GasPrice []byte `protobuf:"bytes,2,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	// Gas limit (uint256, serialized little endian)
	GasLimit []byte `protobuf:"bytes,3,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	// Recipient's address.
	ToAddress string `protobuf:"bytes,4,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	// Amount to send in wei (uint256, serialized little endian)
	Amount []byte `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
	// Optional payload
	Payload []byte `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	// The secret private key used for signing (32 bytes).
	PrivateKey []byte `protobuf:"bytes,7,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Timestamp
	Timestamp uint64 `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SigningInput) Reset() {
	*x = SigningInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Aion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningInput) ProtoMessage() {}

func (x *SigningInput) ProtoReflect() protoreflect.Message {
	mi := &file_Aion_proto_msgTypes[0]
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
	return file_Aion_proto_rawDescGZIP(), []int{0}
}

func (x *SigningInput) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *SigningInput) GetGasPrice() []byte {
	if x != nil {
		return x.GasPrice
	}
	return nil
}

func (x *SigningInput) GetGasLimit() []byte {
	if x != nil {
		return x.GasLimit
	}
	return nil
}

func (x *SigningInput) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *SigningInput) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *SigningInput) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SigningInput) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *SigningInput) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Result containing the signed and encoded transaction.
type SigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Signed and encoded transaction bytes.
	Encoded []byte `protobuf:"bytes,1,opt,name=encoded,proto3" json:"encoded,omitempty"`
	// Signature.
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SigningOutput) Reset() {
	*x = SigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Aion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningOutput) ProtoMessage() {}

func (x *SigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_Aion_proto_msgTypes[1]
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
	return file_Aion_proto_rawDescGZIP(), []int{1}
}

func (x *SigningOutput) GetEncoded() []byte {
	if x != nil {
		return x.Encoded
	}
	return nil
}

func (x *SigningOutput) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_Aion_proto protoreflect.FileDescriptor

var file_Aion_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x41, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x54, 0x57,
	0x2e, 0x41, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x0c,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x47, 0x0a, 0x0d,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x17, 0x0a, 0x15, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6a, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Aion_proto_rawDescOnce sync.Once
	file_Aion_proto_rawDescData = file_Aion_proto_rawDesc
)

func file_Aion_proto_rawDescGZIP() []byte {
	file_Aion_proto_rawDescOnce.Do(func() {
		file_Aion_proto_rawDescData = protoimpl.X.CompressGZIP(file_Aion_proto_rawDescData)
	})
	return file_Aion_proto_rawDescData
}

var file_Aion_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Aion_proto_goTypes = []interface{}{
	(*SigningInput)(nil),  // 0: TW.Aion.Proto.SigningInput
	(*SigningOutput)(nil), // 1: TW.Aion.Proto.SigningOutput
}
var file_Aion_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Aion_proto_init() }
func file_Aion_proto_init() {
	if File_Aion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Aion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Aion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Aion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Aion_proto_goTypes,
		DependencyIndexes: file_Aion_proto_depIdxs,
		MessageInfos:      file_Aion_proto_msgTypes,
	}.Build()
	File_Aion_proto = out.File
	file_Aion_proto_rawDesc = nil
	file_Aion_proto_goTypes = nil
	file_Aion_proto_depIdxs = nil
}
