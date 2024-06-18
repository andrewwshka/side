// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: side/btcbridge/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// AssetType defines the type of asset
type AssetType int32

const (
	// Unspecified asset type
	AssetType_ASSET_TYPE_UNSPECIFIED AssetType = 0
	// BTC
	AssetType_ASSET_TYPE_BTC AssetType = 1
	// BRC20: ordi, sats
	AssetType_ASSET_TYPE_BRC20 AssetType = 2
	// RUNE, dog*go*to*the*moon
	AssetType_ASSET_TYPE_RUNE AssetType = 3
)

var AssetType_name = map[int32]string{
	0: "ASSET_TYPE_UNSPECIFIED",
	1: "ASSET_TYPE_BTC",
	2: "ASSET_TYPE_BRC20",
	3: "ASSET_TYPE_RUNE",
}

var AssetType_value = map[string]int32{
	"ASSET_TYPE_UNSPECIFIED": 0,
	"ASSET_TYPE_BTC":         1,
	"ASSET_TYPE_BRC20":       2,
	"ASSET_TYPE_RUNE":        3,
}

func (x AssetType) String() string {
	return proto.EnumName(AssetType_name, int32(x))
}

func (AssetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1d33573cda8a6d2, []int{0}
}

// Params defines the parameters for the module.
type Params struct {
	// Only accept blocks sending from these addresses
	AuthorizedRelayers []string `protobuf:"bytes,1,rep,name=authorized_relayers,json=authorizedRelayers,proto3" json:"authorized_relayers,omitempty"`
	// The minimum number of confirmations required for a block to be accepted
	Confirmations int32 `protobuf:"varint,2,opt,name=confirmations,proto3" json:"confirmations,omitempty"`
	// Indicates the maximum depth or distance from the latest block up to which transactions are considered for acceptance.
	MaxAcceptableBlockDepth uint64 `protobuf:"varint,3,opt,name=max_acceptable_block_depth,json=maxAcceptableBlockDepth,proto3" json:"max_acceptable_block_depth,omitempty"`
	// the denomanation of the voucher
	BtcVoucherDenom string   `protobuf:"bytes,4,opt,name=btc_voucher_denom,json=btcVoucherDenom,proto3" json:"btc_voucher_denom,omitempty"`
	Vaults          []*Vault `protobuf:"bytes,5,rep,name=vaults,proto3" json:"vaults,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1d33573cda8a6d2, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAuthorizedRelayers() []string {
	if m != nil {
		return m.AuthorizedRelayers
	}
	return nil
}

func (m *Params) GetConfirmations() int32 {
	if m != nil {
		return m.Confirmations
	}
	return 0
}

func (m *Params) GetMaxAcceptableBlockDepth() uint64 {
	if m != nil {
		return m.MaxAcceptableBlockDepth
	}
	return 0
}

func (m *Params) GetBtcVoucherDenom() string {
	if m != nil {
		return m.BtcVoucherDenom
	}
	return ""
}

func (m *Params) GetVaults() []*Vault {
	if m != nil {
		return m.Vaults
	}
	return nil
}

// Vault defines the parameters for the module.
type Vault struct {
	// the depositor should send their btc to this address
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// the pub key to which the voucher is sent
	PubKey string `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	// the address to which the voucher is sent
	AssetType AssetType `protobuf:"varint,4,opt,name=asset_type,json=assetType,proto3,enum=side.btcbridge.AssetType" json:"asset_type,omitempty"`
}

func (m *Vault) Reset()         { *m = Vault{} }
func (m *Vault) String() string { return proto.CompactTextString(m) }
func (*Vault) ProtoMessage()    {}
func (*Vault) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1d33573cda8a6d2, []int{1}
}
func (m *Vault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vault.Merge(m, src)
}
func (m *Vault) XXX_Size() int {
	return m.Size()
}
func (m *Vault) XXX_DiscardUnknown() {
	xxx_messageInfo_Vault.DiscardUnknown(m)
}

var xxx_messageInfo_Vault proto.InternalMessageInfo

func (m *Vault) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Vault) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *Vault) GetAssetType() AssetType {
	if m != nil {
		return m.AssetType
	}
	return AssetType_ASSET_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("side.btcbridge.AssetType", AssetType_name, AssetType_value)
	proto.RegisterType((*Params)(nil), "side.btcbridge.Params")
	proto.RegisterType((*Vault)(nil), "side.btcbridge.Vault")
}

func init() { proto.RegisterFile("side/btcbridge/params.proto", fileDescriptor_f1d33573cda8a6d2) }

var fileDescriptor_f1d33573cda8a6d2 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcf, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xe3, 0xa6, 0x49, 0xb1, 0xc6, 0x52, 0x4f, 0xed, 0x56, 0x2f, 0x03, 0x63, 0xca, 0x0e,
	0xa6, 0x30, 0x7b, 0x64, 0x97, 0xc1, 0x4e, 0xf9, 0xe1, 0xb1, 0x32, 0x28, 0x41, 0x49, 0x0b, 0xdb,
	0x45, 0x48, 0xb2, 0x9a, 0x98, 0xda, 0x91, 0x91, 0xe4, 0x12, 0xef, 0xaf, 0xd8, 0x9f, 0xb5, 0x63,
	0x8f, 0x3b, 0x8e, 0xe4, 0x8f, 0xd8, 0x75, 0xc8, 0x6d, 0xb3, 0xa4, 0xb7, 0xf7, 0x3e, 0x9f, 0xaf,
	0xed, 0xf7, 0xcc, 0x03, 0x6f, 0x54, 0x9a, 0xf0, 0x88, 0x6a, 0x46, 0x65, 0x9a, 0xcc, 0x78, 0x54,
	0x10, 0x49, 0x72, 0x15, 0x16, 0x52, 0x68, 0x01, 0x3b, 0x46, 0x86, 0x1b, 0xd9, 0x3d, 0x9e, 0x89,
	0x99, 0xa8, 0x55, 0x64, 0xaa, 0xfb, 0xd4, 0xe9, 0x5f, 0x0b, 0xb4, 0xc7, 0xf5, 0x63, 0x30, 0x02,
	0x47, 0xa4, 0xd4, 0x73, 0x21, 0xd3, 0x1f, 0x3c, 0xc1, 0x92, 0x67, 0xa4, 0xe2, 0x52, 0xb9, 0x96,
	0xdf, 0x0c, 0x6c, 0x04, 0xff, 0x2b, 0xf4, 0x60, 0xe0, 0x5b, 0xf0, 0x9c, 0x89, 0xc5, 0x75, 0x2a,
	0x73, 0xa2, 0x53, 0xb1, 0x50, 0xee, 0x9e, 0x6f, 0x05, 0x2d, 0xb4, 0x0b, 0xe1, 0x27, 0xd0, 0xcd,
	0xc9, 0x12, 0x13, 0xc6, 0x78, 0xa1, 0x09, 0xcd, 0x38, 0xa6, 0x99, 0x60, 0x37, 0x38, 0xe1, 0x85,
	0x9e, 0xbb, 0x4d, 0xdf, 0x0a, 0xf6, 0xd1, 0x49, 0x4e, 0x96, 0xfd, 0x4d, 0x60, 0x60, 0xfc, 0xc8,
	0x68, 0x78, 0x06, 0x5e, 0x50, 0xcd, 0xf0, 0xad, 0x28, 0xd9, 0x9c, 0x4b, 0x9c, 0xf0, 0x85, 0xc8,
	0xdd, 0x7d, 0xdf, 0x0a, 0x6c, 0x74, 0x48, 0x35, 0xbb, 0xba, 0xe7, 0x23, 0x83, 0xe1, 0x3b, 0xd0,
	0xbe, 0x25, 0x65, 0xa6, 0x95, 0xdb, 0xf2, 0x9b, 0xc1, 0xb3, 0xde, 0xcb, 0x70, 0xf7, 0x0f, 0x84,
	0x57, 0xc6, 0xa2, 0x87, 0xd0, 0xa9, 0x06, 0xad, 0x1a, 0x40, 0x17, 0x1c, 0x90, 0x24, 0x91, 0x5c,
	0x99, 0x5d, 0xcd, 0x9b, 0x1f, 0x5b, 0x78, 0x02, 0x0e, 0x8a, 0x92, 0xe2, 0x1b, 0x5e, 0xd5, 0xab,
	0xd9, 0xa8, 0x5d, 0x94, 0xf4, 0x2b, 0xaf, 0xe0, 0x47, 0x00, 0x88, 0x52, 0x5c, 0x63, 0x5d, 0x15,
	0xbc, 0x9e, 0xa7, 0xd3, 0x7b, 0xfd, 0xf4, 0x73, 0x7d, 0x93, 0x98, 0x56, 0x05, 0x47, 0x36, 0x79,
	0x2c, 0xcf, 0xae, 0x81, 0xbd, 0xe1, 0xb0, 0x0b, 0x5e, 0xf5, 0x27, 0x93, 0x78, 0x8a, 0xa7, 0xdf,
	0xc6, 0x31, 0xbe, 0xbc, 0x98, 0x8c, 0xe3, 0xe1, 0xf9, 0xe7, 0xf3, 0x78, 0xe4, 0x34, 0x20, 0x04,
	0x9d, 0x2d, 0x37, 0x98, 0x0e, 0x1d, 0x0b, 0x1e, 0x03, 0x67, 0x9b, 0xa1, 0x61, 0xef, 0xbd, 0xb3,
	0x07, 0x8f, 0xc0, 0xe1, 0x16, 0x45, 0x97, 0x17, 0xb1, 0xd3, 0x1c, 0x7c, 0xf9, 0xb5, 0xf2, 0xac,
	0xbb, 0x95, 0x67, 0xfd, 0x59, 0x79, 0xd6, 0xcf, 0xb5, 0xd7, 0xb8, 0x5b, 0x7b, 0x8d, 0xdf, 0x6b,
	0xaf, 0xf1, 0x3d, 0x9c, 0xa5, 0x7a, 0x5e, 0xd2, 0x90, 0x89, 0x3c, 0x32, 0x13, 0xd7, 0x77, 0xc0,
	0x44, 0x56, 0x37, 0xd1, 0x72, 0xeb, 0x9c, 0xcc, 0x72, 0x8a, 0xb6, 0xeb, 0xc0, 0x87, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x87, 0xf7, 0x68, 0xef, 0x6d, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Vaults) > 0 {
		for iNdEx := len(m.Vaults) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Vaults[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.BtcVoucherDenom) > 0 {
		i -= len(m.BtcVoucherDenom)
		copy(dAtA[i:], m.BtcVoucherDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BtcVoucherDenom)))
		i--
		dAtA[i] = 0x22
	}
	if m.MaxAcceptableBlockDepth != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxAcceptableBlockDepth))
		i--
		dAtA[i] = 0x18
	}
	if m.Confirmations != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Confirmations))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AuthorizedRelayers) > 0 {
		for iNdEx := len(m.AuthorizedRelayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AuthorizedRelayers[iNdEx])
			copy(dAtA[i:], m.AuthorizedRelayers[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.AuthorizedRelayers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Vault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AssetType != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AssetType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AuthorizedRelayers) > 0 {
		for _, s := range m.AuthorizedRelayers {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.Confirmations != 0 {
		n += 1 + sovParams(uint64(m.Confirmations))
	}
	if m.MaxAcceptableBlockDepth != 0 {
		n += 1 + sovParams(uint64(m.MaxAcceptableBlockDepth))
	}
	l = len(m.BtcVoucherDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.Vaults) > 0 {
		for _, e := range m.Vaults {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *Vault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.AssetType != 0 {
		n += 1 + sovParams(uint64(m.AssetType))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedRelayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizedRelayers = append(m.AuthorizedRelayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Confirmations", wireType)
			}
			m.Confirmations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Confirmations |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAcceptableBlockDepth", wireType)
			}
			m.MaxAcceptableBlockDepth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAcceptableBlockDepth |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcVoucherDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BtcVoucherDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vaults", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vaults = append(m.Vaults, &Vault{})
			if err := m.Vaults[len(m.Vaults)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Vault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Vault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetType", wireType)
			}
			m.AssetType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetType |= AssetType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)