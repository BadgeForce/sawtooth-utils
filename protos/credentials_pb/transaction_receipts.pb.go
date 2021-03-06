// Code generated by protoc-gen-go. DO NOT EDIT.
// source: credentials/transaction_receipts.proto

package credentials_pb // import "github.com/BadgeForce/sawtooth-utils/protos/credentials_pb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IssuingEngineReceipt struct {
	Date                 int64       `protobuf:"varint,1,opt,name=date,proto3" json:"date,omitempty"`
	StateAddress         string      `protobuf:"bytes,2,opt,name=state_address,json=stateAddress,proto3" json:"state_address,omitempty"`
	RpcMethod            Method      `protobuf:"varint,3,opt,name=rpc_method,json=rpcMethod,proto3,enum=credential_pb.Method" json:"rpc_method,omitempty"`
	Credential           *Credential `protobuf:"bytes,4,opt,name=credential,proto3" json:"credential,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *IssuingEngineReceipt) Reset()         { *m = IssuingEngineReceipt{} }
func (m *IssuingEngineReceipt) String() string { return proto.CompactTextString(m) }
func (*IssuingEngineReceipt) ProtoMessage()    {}
func (*IssuingEngineReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_receipts_1768d7e1008a8037, []int{0}
}
func (m *IssuingEngineReceipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssuingEngineReceipt.Unmarshal(m, b)
}
func (m *IssuingEngineReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssuingEngineReceipt.Marshal(b, m, deterministic)
}
func (dst *IssuingEngineReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssuingEngineReceipt.Merge(dst, src)
}
func (m *IssuingEngineReceipt) XXX_Size() int {
	return xxx_messageInfo_IssuingEngineReceipt.Size(m)
}
func (m *IssuingEngineReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_IssuingEngineReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_IssuingEngineReceipt proto.InternalMessageInfo

func (m *IssuingEngineReceipt) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *IssuingEngineReceipt) GetStateAddress() string {
	if m != nil {
		return m.StateAddress
	}
	return ""
}

func (m *IssuingEngineReceipt) GetRpcMethod() Method {
	if m != nil {
		return m.RpcMethod
	}
	return Method_ISSUE
}

func (m *IssuingEngineReceipt) GetCredential() *Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

func init() {
	proto.RegisterType((*IssuingEngineReceipt)(nil), "credential_pb.IssuingEngineReceipt")
}

func init() {
	proto.RegisterFile("credentials/transaction_receipts.proto", fileDescriptor_transaction_receipts_1768d7e1008a8037)
}

var fileDescriptor_transaction_receipts_1768d7e1008a8037 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xfc, 0x30,
	0x14, 0xc4, 0xc9, 0x7f, 0x97, 0x3f, 0x6c, 0x74, 0x3d, 0x04, 0x85, 0xee, 0xe2, 0xa1, 0x28, 0x48,
	0x2f, 0xb6, 0xb0, 0x7a, 0x51, 0xbc, 0xb8, 0xa2, 0xe0, 0xc1, 0x4b, 0x8f, 0x5e, 0xca, 0x6b, 0xf2,
	0x68, 0x03, 0xbb, 0x49, 0xc8, 0x7b, 0x45, 0xfc, 0x76, 0x7e, 0x34, 0xd9, 0x16, 0x6c, 0xf5, 0x36,
	0x99, 0xf9, 0x0d, 0xc9, 0x44, 0x5e, 0xe9, 0x88, 0x06, 0x1d, 0x5b, 0xd8, 0x51, 0xc1, 0x11, 0x1c,
	0x81, 0x66, 0xeb, 0x5d, 0x15, 0x51, 0xa3, 0x0d, 0x4c, 0x79, 0x88, 0x9e, 0xbd, 0x5a, 0x8e, 0x5c,
	0x15, 0xea, 0xf5, 0xf9, 0xb4, 0x36, 0xea, 0x01, 0x5e, 0xaf, 0xa6, 0x69, 0x80, 0xcf, 0x9d, 0x07,
	0x33, 0x44, 0x17, 0x5f, 0x42, 0x9e, 0xbe, 0x12, 0x75, 0xd6, 0x35, 0xcf, 0xae, 0xb1, 0x0e, 0xcb,
	0xe1, 0x1e, 0xa5, 0xe4, 0xdc, 0x00, 0x63, 0x22, 0x52, 0x91, 0xcd, 0xca, 0x5e, 0xab, 0x4b, 0xb9,
	0x24, 0x06, 0xc6, 0x0a, 0x8c, 0x89, 0x48, 0x94, 0xfc, 0x4b, 0x45, 0xb6, 0x28, 0x8f, 0x7b, 0xf3,
	0x71, 0xf0, 0xd4, 0xad, 0x94, 0x31, 0xe8, 0x6a, 0x8f, 0xdc, 0x7a, 0x93, 0xcc, 0x52, 0x91, 0x9d,
	0x6c, 0xce, 0xf2, 0x5f, 0xcf, 0xcd, 0xdf, 0xfa, 0xb0, 0x5c, 0xc4, 0xa0, 0x07, 0xa9, 0xee, 0xa4,
	0x1c, 0x91, 0x64, 0x9e, 0x8a, 0xec, 0x68, 0xb3, 0xfa, 0xd3, 0x7a, 0xfa, 0x39, 0x95, 0x13, 0x78,
	0xfb, 0xf0, 0x7e, 0xdf, 0x58, 0x6e, 0xbb, 0x3a, 0xd7, 0x7e, 0x5f, 0x6c, 0xc1, 0x34, 0xf8, 0xe2,
	0xa3, 0xc6, 0x82, 0xe0, 0x83, 0xbd, 0xe7, 0xf6, 0xba, 0x63, 0x7b, 0x18, 0x7e, 0x18, 0x3c, 0xfd,
	0x1d, 0xaa, 0x42, 0x5d, 0xff, 0xef, 0xed, 0x9b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xc1,
	0x55, 0xae, 0x79, 0x01, 0x00, 0x00,
}
