// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/ethbridge/v1/types.proto

package types

import (
	fmt "fmt"
	types "github.com/Sifchain/sifnode/x/oracle/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Claim type enum
type ClaimType int32

const (
	// Unspecified claim type
	ClaimType_CLAIM_TYPE_UNSPECIFIED ClaimType = 0
	// Burn claim type
	ClaimType_CLAIM_TYPE_BURN ClaimType = 1
	// Lock claim type
	ClaimType_CLAIM_TYPE_LOCK ClaimType = 2
)

var ClaimType_name = map[int32]string{
	0: "CLAIM_TYPE_UNSPECIFIED",
	1: "CLAIM_TYPE_BURN",
	2: "CLAIM_TYPE_LOCK",
}

var ClaimType_value = map[string]int32{
	"CLAIM_TYPE_UNSPECIFIED": 0,
	"CLAIM_TYPE_BURN":        1,
	"CLAIM_TYPE_LOCK":        2,
}

func (x ClaimType) String() string {
	return proto.EnumName(ClaimType_name, int32(x))
}

func (ClaimType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4cb34f678c9ed59f, []int{0}
}

// EthBridgeClaim is a structure that contains all the data for a particular
// bridge claim
type EthBridgeClaim struct {
	NetworkDescriptor types.NetworkDescriptor `protobuf:"varint,1,opt,name=network_descriptor,json=networkDescriptor,proto3,enum=sifnode.oracle.v1.NetworkDescriptor" json:"network_descriptor,omitempty" yaml:"network_descriptor"`
	// bridge_contract_address is an EthereumAddress
	BridgeContractAddress string `protobuf:"bytes,2,opt,name=bridge_contract_address,json=bridgeContractAddress,proto3" json:"bridge_contract_address,omitempty" yaml:"bridge_contract_address"`
	Nonce                 int64  `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty" yaml:"nonce"`
	Symbol                string `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty" yaml:"symbol"`
	// token_contract_address is an EthereumAddress
	TokenContractAddress string `protobuf:"bytes,5,opt,name=token_contract_address,json=tokenContractAddress,proto3" json:"token_contract_address,omitempty" yaml:"token_contract_address"`
	// ethereum_sender is an EthereumAddress
	EthereumSender string `protobuf:"bytes,6,opt,name=ethereum_sender,json=ethereumSender,proto3" json:"ethereum_sender,omitempty" yaml:"ethereum_sender"`
	// cosmos_receiver is an sdk.AccAddress
	CosmosReceiver string `protobuf:"bytes,7,opt,name=cosmos_receiver,json=cosmosReceiver,proto3" json:"cosmos_receiver,omitempty" yaml:"cosmos_receiver"`
	// validator_address is an sdk.ValAddress
	ValidatorAddress string                                 `protobuf:"bytes,8,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`
	Amount           github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount" yaml:"amount"`
	ClaimType        ClaimType                              `protobuf:"varint,10,opt,name=claim_type,json=claimType,proto3,enum=sifnode.ethbridge.v1.ClaimType" json:"claim_type,omitempty"`
	// Types that are valid to be assigned to XTokenName:
	//	*EthBridgeClaim_TokenName
	XTokenName isEthBridgeClaim_XTokenName `protobuf_oneof:"_token_name"`
	// Types that are valid to be assigned to XDecimals:
	//	*EthBridgeClaim_Decimals
	XDecimals isEthBridgeClaim_XDecimals `protobuf_oneof:"_decimals"`
	// Types that are valid to be assigned to XDenomHash:
	//	*EthBridgeClaim_DenomHash
	XDenomHash isEthBridgeClaim_XDenomHash `protobuf_oneof:"_denom_hash"`
}

func (m *EthBridgeClaim) Reset()         { *m = EthBridgeClaim{} }
func (m *EthBridgeClaim) String() string { return proto.CompactTextString(m) }
func (*EthBridgeClaim) ProtoMessage()    {}
func (*EthBridgeClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cb34f678c9ed59f, []int{0}
}
func (m *EthBridgeClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthBridgeClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthBridgeClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthBridgeClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthBridgeClaim.Merge(m, src)
}
func (m *EthBridgeClaim) XXX_Size() int {
	return m.Size()
}
func (m *EthBridgeClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_EthBridgeClaim.DiscardUnknown(m)
}

var xxx_messageInfo_EthBridgeClaim proto.InternalMessageInfo

type isEthBridgeClaim_XTokenName interface {
	isEthBridgeClaim_XTokenName()
	MarshalTo([]byte) (int, error)
	Size() int
}
type isEthBridgeClaim_XDecimals interface {
	isEthBridgeClaim_XDecimals()
	MarshalTo([]byte) (int, error)
	Size() int
}
type isEthBridgeClaim_XDenomHash interface {
	isEthBridgeClaim_XDenomHash()
	MarshalTo([]byte) (int, error)
	Size() int
}

type EthBridgeClaim_TokenName struct {
	TokenName string `protobuf:"bytes,11,opt,name=token_name,json=tokenName,proto3,oneof" json:"token_name,omitempty" yaml:"token_name"`
}
type EthBridgeClaim_Decimals struct {
	Decimals int32 `protobuf:"varint,12,opt,name=decimals,proto3,oneof" json:"decimals,omitempty" yaml:"token_decimals"`
}
type EthBridgeClaim_DenomHash struct {
	DenomHash string `protobuf:"bytes,13,opt,name=denom_hash,json=denomHash,proto3,oneof" json:"denom_hash,omitempty" yaml:"denom_hash"`
}

func (*EthBridgeClaim_TokenName) isEthBridgeClaim_XTokenName() {}
func (*EthBridgeClaim_Decimals) isEthBridgeClaim_XDecimals()   {}
func (*EthBridgeClaim_DenomHash) isEthBridgeClaim_XDenomHash() {}

func (m *EthBridgeClaim) GetXTokenName() isEthBridgeClaim_XTokenName {
	if m != nil {
		return m.XTokenName
	}
	return nil
}
func (m *EthBridgeClaim) GetXDecimals() isEthBridgeClaim_XDecimals {
	if m != nil {
		return m.XDecimals
	}
	return nil
}
func (m *EthBridgeClaim) GetXDenomHash() isEthBridgeClaim_XDenomHash {
	if m != nil {
		return m.XDenomHash
	}
	return nil
}

func (m *EthBridgeClaim) GetNetworkDescriptor() types.NetworkDescriptor {
	if m != nil {
		return m.NetworkDescriptor
	}
	return types.NetworkDescriptor_NETWORK_DESCRIPTOR_UNSPECIFIED
}

func (m *EthBridgeClaim) GetBridgeContractAddress() string {
	if m != nil {
		return m.BridgeContractAddress
	}
	return ""
}

func (m *EthBridgeClaim) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *EthBridgeClaim) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *EthBridgeClaim) GetTokenContractAddress() string {
	if m != nil {
		return m.TokenContractAddress
	}
	return ""
}

func (m *EthBridgeClaim) GetEthereumSender() string {
	if m != nil {
		return m.EthereumSender
	}
	return ""
}

func (m *EthBridgeClaim) GetCosmosReceiver() string {
	if m != nil {
		return m.CosmosReceiver
	}
	return ""
}

func (m *EthBridgeClaim) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *EthBridgeClaim) GetClaimType() ClaimType {
	if m != nil {
		return m.ClaimType
	}
	return ClaimType_CLAIM_TYPE_UNSPECIFIED
}

func (m *EthBridgeClaim) GetTokenName() string {
	if x, ok := m.GetXTokenName().(*EthBridgeClaim_TokenName); ok {
		return x.TokenName
	}
	return ""
}

func (m *EthBridgeClaim) GetDecimals() int32 {
	if x, ok := m.GetXDecimals().(*EthBridgeClaim_Decimals); ok {
		return x.Decimals
	}
	return 0
}

func (m *EthBridgeClaim) GetDenomHash() string {
	if x, ok := m.GetXDenomHash().(*EthBridgeClaim_DenomHash); ok {
		return x.DenomHash
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EthBridgeClaim) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EthBridgeClaim_TokenName)(nil),
		(*EthBridgeClaim_Decimals)(nil),
		(*EthBridgeClaim_DenomHash)(nil),
	}
}

type PeggyTokens struct {
	Tokens []string `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (m *PeggyTokens) Reset()         { *m = PeggyTokens{} }
func (m *PeggyTokens) String() string { return proto.CompactTextString(m) }
func (*PeggyTokens) ProtoMessage()    {}
func (*PeggyTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cb34f678c9ed59f, []int{1}
}
func (m *PeggyTokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeggyTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeggyTokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeggyTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeggyTokens.Merge(m, src)
}
func (m *PeggyTokens) XXX_Size() int {
	return m.Size()
}
func (m *PeggyTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_PeggyTokens.DiscardUnknown(m)
}

var xxx_messageInfo_PeggyTokens proto.InternalMessageInfo

func (m *PeggyTokens) GetTokens() []string {
	if m != nil {
		return m.Tokens
	}
	return nil
}

// GenesisState for ethbridge
type GenesisState struct {
	CrosschainFeeReceiveAccount string   `protobuf:"bytes,1,opt,name=crosschain_fee_receive_account,json=crosschainFeeReceiveAccount,proto3" json:"crosschain_fee_receive_account,omitempty"`
	PeggyTokens                 []string `protobuf:"bytes,2,rep,name=peggy_tokens,json=peggyTokens,proto3" json:"peggy_tokens,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cb34f678c9ed59f, []int{2}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetCrosschainFeeReceiveAccount() string {
	if m != nil {
		return m.CrosschainFeeReceiveAccount
	}
	return ""
}

func (m *GenesisState) GetPeggyTokens() []string {
	if m != nil {
		return m.PeggyTokens
	}
	return nil
}

func init() {
	proto.RegisterEnum("sifnode.ethbridge.v1.ClaimType", ClaimType_name, ClaimType_value)
	proto.RegisterType((*EthBridgeClaim)(nil), "sifnode.ethbridge.v1.EthBridgeClaim")
	proto.RegisterType((*PeggyTokens)(nil), "sifnode.ethbridge.v1.PeggyTokens")
	proto.RegisterType((*GenesisState)(nil), "sifnode.ethbridge.v1.GenesisState")
}

func init() { proto.RegisterFile("sifnode/ethbridge/v1/types.proto", fileDescriptor_4cb34f678c9ed59f) }

var fileDescriptor_4cb34f678c9ed59f = []byte{
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x41, 0x6f, 0xdb, 0x36,
	0x14, 0xc7, 0xad, 0x64, 0x71, 0x2b, 0x3a, 0x49, 0x1d, 0x2e, 0xf1, 0x54, 0x6f, 0x95, 0x5c, 0x62,
	0x2b, 0xbc, 0x02, 0x95, 0x91, 0xed, 0xd4, 0x1e, 0x36, 0x44, 0xae, 0x5b, 0x1b, 0xeb, 0xbc, 0x8c,
	0x4e, 0x51, 0xac, 0x17, 0x81, 0x96, 0x18, 0x4b, 0x88, 0x25, 0x1a, 0x22, 0xe3, 0xcd, 0xdf, 0xa0,
	0xc7, 0x7d, 0xac, 0x1e, 0x7b, 0xdb, 0xb0, 0x83, 0x30, 0x24, 0xdf, 0x40, 0x9f, 0x60, 0x10, 0x29,
	0xd9, 0x86, 0x9d, 0x9d, 0x2c, 0xbe, 0xff, 0xef, 0xfd, 0xdf, 0x23, 0xde, 0x33, 0x41, 0x8b, 0x87,
	0x97, 0x31, 0xf3, 0x69, 0x87, 0x8a, 0x60, 0x9c, 0x84, 0xfe, 0x84, 0x76, 0xe6, 0xa7, 0x1d, 0xb1,
	0x98, 0x51, 0x6e, 0xcf, 0x12, 0x26, 0x18, 0x3c, 0x2e, 0x08, 0x7b, 0x49, 0xd8, 0xf3, 0xd3, 0xe6,
	0xf1, 0x84, 0x4d, 0x98, 0x04, 0x3a, 0xf9, 0x97, 0x62, 0x9b, 0x4f, 0x4b, 0x37, 0x96, 0x10, 0x6f,
	0x2a, 0xad, 0x62, 0x2a, 0x7e, 0x67, 0xc9, 0x95, 0xeb, 0x53, 0xee, 0x25, 0xe1, 0x4c, 0xb0, 0x44,
	0xb1, 0xe8, 0xaf, 0x7b, 0xe0, 0xb0, 0x27, 0x02, 0x47, 0x5a, 0x76, 0xa7, 0x24, 0x8c, 0x60, 0x02,
	0xe0, 0x36, 0x6e, 0x68, 0x2d, 0xad, 0x7d, 0xf8, 0xdd, 0xd7, 0x76, 0xd9, 0x87, 0xf2, 0xb6, 0xe7,
	0xa7, 0xf6, 0x50, 0xc1, 0x2f, 0x97, 0xac, 0xf3, 0x28, 0x4b, 0xad, 0x87, 0x0b, 0x12, 0x4d, 0x5f,
	0xa0, 0x6d, 0x27, 0x84, 0x8f, 0xe2, 0xcd, 0x0c, 0xf8, 0x1e, 0x7c, 0xa1, 0x6e, 0xe5, 0x7a, 0x2c,
	0x16, 0x09, 0xf1, 0x84, 0x4b, 0x7c, 0x3f, 0xa1, 0x9c, 0x1b, 0x3b, 0x2d, 0xad, 0xad, 0x3b, 0x28,
	0x4b, 0x2d, 0x53, 0x59, 0xfe, 0x0f, 0x88, 0xf0, 0x89, 0x52, 0xba, 0x85, 0x70, 0xa6, 0xe2, 0xf0,
	0x09, 0xd8, 0x8b, 0x59, 0xec, 0x51, 0x63, 0xb7, 0xa5, 0xb5, 0x77, 0x9d, 0x7a, 0x96, 0x5a, 0xfb,
	0x45, 0x73, 0x79, 0x18, 0x61, 0x25, 0xc3, 0x6f, 0x41, 0x95, 0x2f, 0xa2, 0x31, 0x9b, 0x1a, 0x9f,
	0xc9, 0x92, 0x47, 0x59, 0x6a, 0x1d, 0x28, 0x50, 0xc5, 0x11, 0x2e, 0x00, 0xf8, 0x0e, 0x34, 0x04,
	0xbb, 0xa2, 0xf1, 0x76, 0xb7, 0x7b, 0x32, 0xf5, 0x71, 0x96, 0x5a, 0x8f, 0x54, 0xea, 0xdd, 0x1c,
	0xc2, 0xc7, 0x52, 0xd8, 0xec, 0xb5, 0x0b, 0x1e, 0x50, 0x11, 0xd0, 0x84, 0x5e, 0x47, 0x2e, 0xa7,
	0xb1, 0x4f, 0x13, 0xa3, 0x2a, 0x1d, 0x9b, 0x59, 0x6a, 0x35, 0x94, 0xe3, 0x06, 0x80, 0xf0, 0x61,
	0x19, 0x19, 0xc9, 0x40, 0x6e, 0xe2, 0x31, 0x1e, 0x31, 0xee, 0x26, 0xd4, 0xa3, 0xe1, 0x9c, 0x26,
	0xc6, 0xbd, 0x4d, 0x93, 0x0d, 0x00, 0xe1, 0x43, 0x15, 0xc1, 0x45, 0x00, 0x0e, 0xc0, 0xd1, 0x9c,
	0x4c, 0x43, 0x9f, 0x08, 0x96, 0x2c, 0x6f, 0x77, 0x5f, 0xda, 0x7c, 0x95, 0xa5, 0x96, 0xa1, 0x6c,
	0xb6, 0x10, 0x84, 0xeb, 0xcb, 0x58, 0x79, 0xa9, 0x77, 0xa0, 0x4a, 0x22, 0x76, 0x1d, 0x0b, 0x43,
	0x97, 0xf9, 0x3f, 0x7e, 0x4c, 0xad, 0xca, 0x3f, 0xa9, 0xf5, 0x64, 0x12, 0x8a, 0xe0, 0x7a, 0x6c,
	0x7b, 0x2c, 0xea, 0xa8, 0xea, 0xc5, 0xcf, 0x33, 0xee, 0x5f, 0x15, 0xdb, 0x3f, 0x88, 0xc5, 0x6a,
	0x0c, 0xca, 0x05, 0xe1, 0xc2, 0x0e, 0xfe, 0x00, 0x80, 0x97, 0xaf, 0xac, 0x9b, 0xb3, 0x06, 0x90,
	0x1b, 0x6a, 0xd9, 0x77, 0xfd, 0x53, 0x6c, 0xb9, 0xda, 0x17, 0x8b, 0x19, 0xc5, 0xba, 0x57, 0x7e,
	0xc2, 0xe7, 0x00, 0xa8, 0xf1, 0xc4, 0x24, 0xa2, 0x46, 0x4d, 0x36, 0x77, 0x92, 0xa5, 0xd6, 0xd1,
	0xfa, 0xe8, 0x72, 0x0d, 0xf5, 0x2b, 0x58, 0x97, 0xc7, 0x21, 0x89, 0xe8, 0x07, 0x4d, 0x83, 0x2f,
	0xc0, 0x7d, 0x9f, 0x7a, 0x61, 0x44, 0xa6, 0xdc, 0xd8, 0x6f, 0x69, 0xed, 0x3d, 0xe7, 0x61, 0x96,
	0x5a, 0x27, 0xeb, 0x89, 0xa5, 0x8e, 0xfa, 0x1a, 0x5e, 0xc2, 0x79, 0xee, 0x73, 0x00, 0x7c, 0x1a,
	0xb3, 0xc8, 0x0d, 0x08, 0x0f, 0x8c, 0x83, 0xcd, 0xb2, 0x2b, 0x0d, 0xf5, 0x77, 0xb0, 0x2e, 0x8f,
	0x7d, 0xc2, 0x83, 0x0f, 0x9a, 0xe6, 0x1c, 0x80, 0x9a, 0xbb, 0x6a, 0xcb, 0xa9, 0x01, 0x7d, 0x59,
	0x46, 0x6a, 0x6b, 0xb9, 0xdf, 0x80, 0xda, 0x39, 0x9d, 0x4c, 0x16, 0x17, 0x39, 0xce, 0x61, 0x03,
	0x54, 0x65, 0x22, 0x37, 0xb4, 0xd6, 0x6e, 0x5b, 0xc7, 0xc5, 0x09, 0xcd, 0xc1, 0xfe, 0x6b, 0x1a,
	0x53, 0x1e, 0xf2, 0x91, 0x20, 0x82, 0xc2, 0x2e, 0x30, 0xbd, 0x84, 0x71, 0xee, 0x05, 0x24, 0x8c,
	0xdd, 0x4b, 0x4a, 0xcb, 0x1d, 0x71, 0x89, 0xe7, 0xc9, 0x21, 0xe6, 0x2f, 0x81, 0x8e, 0xbf, 0x5c,
	0x51, 0xaf, 0x28, 0x2d, 0xd6, 0xe6, 0x4c, 0x21, 0xf0, 0x31, 0xd8, 0x9f, 0xe5, 0xb5, 0xdd, 0xa2,
	0xe4, 0x8e, 0x2c, 0x59, 0x9b, 0xad, 0xfa, 0x79, 0xfa, 0x2b, 0xd0, 0x97, 0x33, 0x81, 0x4d, 0xd0,
	0xe8, 0xbe, 0x39, 0x1b, 0xfc, 0xec, 0x5e, 0xfc, 0x76, 0xde, 0x73, 0xdf, 0x0e, 0x47, 0xe7, 0xbd,
	0xee, 0xe0, 0xd5, 0xa0, 0xf7, 0xb2, 0x5e, 0x81, 0x9f, 0x83, 0x07, 0x6b, 0x9a, 0xf3, 0x16, 0x0f,
	0xeb, 0xda, 0x46, 0xf0, 0xcd, 0x2f, 0xdd, 0x9f, 0xea, 0x3b, 0xce, 0xeb, 0x8f, 0x37, 0xa6, 0xf6,
	0xe9, 0xc6, 0xd4, 0xfe, 0xbd, 0x31, 0xb5, 0x3f, 0x6f, 0xcd, 0xca, 0xa7, 0x5b, 0xb3, 0xf2, 0xf7,
	0xad, 0x59, 0x79, 0xff, 0x6c, 0x6d, 0xd3, 0x46, 0xe1, 0xa5, 0xec, 0xba, 0x53, 0xbe, 0x92, 0x7f,
	0xac, 0xbd, 0xba, 0x72, 0xe9, 0xc6, 0x55, 0xf9, 0x36, 0x7e, 0xff, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xea, 0xcd, 0x7c, 0x22, 0x97, 0x05, 0x00, 0x00,
}

func (m *EthBridgeClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthBridgeClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthBridgeClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XDenomHash != nil {
		{
			size := m.XDenomHash.Size()
			i -= size
			if _, err := m.XDenomHash.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.XDecimals != nil {
		{
			size := m.XDecimals.Size()
			i -= size
			if _, err := m.XDecimals.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.XTokenName != nil {
		{
			size := m.XTokenName.Size()
			i -= size
			if _, err := m.XTokenName.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.ClaimType != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ClaimType))
		i--
		dAtA[i] = 0x50
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.CosmosReceiver) > 0 {
		i -= len(m.CosmosReceiver)
		copy(dAtA[i:], m.CosmosReceiver)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CosmosReceiver)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.EthereumSender) > 0 {
		i -= len(m.EthereumSender)
		copy(dAtA[i:], m.EthereumSender)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.EthereumSender)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TokenContractAddress) > 0 {
		i -= len(m.TokenContractAddress)
		copy(dAtA[i:], m.TokenContractAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TokenContractAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x22
	}
	if m.Nonce != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x18
	}
	if len(m.BridgeContractAddress) > 0 {
		i -= len(m.BridgeContractAddress)
		copy(dAtA[i:], m.BridgeContractAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BridgeContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.NetworkDescriptor != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.NetworkDescriptor))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EthBridgeClaim_TokenName) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthBridgeClaim_TokenName) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.TokenName)
	copy(dAtA[i:], m.TokenName)
	i = encodeVarintTypes(dAtA, i, uint64(len(m.TokenName)))
	i--
	dAtA[i] = 0x5a
	return len(dAtA) - i, nil
}
func (m *EthBridgeClaim_Decimals) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthBridgeClaim_Decimals) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintTypes(dAtA, i, uint64(m.Decimals))
	i--
	dAtA[i] = 0x60
	return len(dAtA) - i, nil
}
func (m *EthBridgeClaim_DenomHash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthBridgeClaim_DenomHash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.DenomHash)
	copy(dAtA[i:], m.DenomHash)
	i = encodeVarintTypes(dAtA, i, uint64(len(m.DenomHash)))
	i--
	dAtA[i] = 0x6a
	return len(dAtA) - i, nil
}
func (m *PeggyTokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeggyTokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeggyTokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tokens[iNdEx])
			copy(dAtA[i:], m.Tokens[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Tokens[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PeggyTokens) > 0 {
		for iNdEx := len(m.PeggyTokens) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PeggyTokens[iNdEx])
			copy(dAtA[i:], m.PeggyTokens[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.PeggyTokens[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.CrosschainFeeReceiveAccount) > 0 {
		i -= len(m.CrosschainFeeReceiveAccount)
		copy(dAtA[i:], m.CrosschainFeeReceiveAccount)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CrosschainFeeReceiveAccount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthBridgeClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NetworkDescriptor != 0 {
		n += 1 + sovTypes(uint64(m.NetworkDescriptor))
	}
	l = len(m.BridgeContractAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovTypes(uint64(m.Nonce))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.TokenContractAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.EthereumSender)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.CosmosReceiver)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.ClaimType != 0 {
		n += 1 + sovTypes(uint64(m.ClaimType))
	}
	if m.XTokenName != nil {
		n += m.XTokenName.Size()
	}
	if m.XDecimals != nil {
		n += m.XDecimals.Size()
	}
	if m.XDenomHash != nil {
		n += m.XDenomHash.Size()
	}
	return n
}

func (m *EthBridgeClaim_TokenName) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenName)
	n += 1 + l + sovTypes(uint64(l))
	return n
}
func (m *EthBridgeClaim_Decimals) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovTypes(uint64(m.Decimals))
	return n
}
func (m *EthBridgeClaim_DenomHash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DenomHash)
	n += 1 + l + sovTypes(uint64(l))
	return n
}
func (m *PeggyTokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		for _, s := range m.Tokens {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CrosschainFeeReceiveAccount)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.PeggyTokens) > 0 {
		for _, s := range m.PeggyTokens {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthBridgeClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EthBridgeClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthBridgeClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkDescriptor", wireType)
			}
			m.NetworkDescriptor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NetworkDescriptor |= types.NetworkDescriptor(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumSender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthereumSender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosReceiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CosmosReceiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimType", wireType)
			}
			m.ClaimType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimType |= ClaimType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.XTokenName = &EthBridgeClaim_TokenName{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.XDecimals = &EthBridgeClaim_Decimals{v}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.XDenomHash = &EthBridgeClaim_DenomHash{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PeggyTokens) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PeggyTokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeggyTokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrosschainFeeReceiveAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CrosschainFeeReceiveAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeggyTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeggyTokens = append(m.PeggyTokens, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
