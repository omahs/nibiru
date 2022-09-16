// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: perp/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/NibiruChain/nibiru/x/common"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// GenesisState defines the perp module's genesis state.
type GenesisState struct {
	Params               Params            `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PairMetadata         []*PairMetadata   `protobuf:"bytes,2,rep,name=pair_metadata,json=pairMetadata,proto3" json:"pair_metadata,omitempty"`
	Positions            []*Position       `protobuf:"bytes,3,rep,name=positions,proto3" json:"positions,omitempty"`
	PrepaidBadDebts      []*PrepaidBadDebt `protobuf:"bytes,4,rep,name=prepaid_bad_debts,json=prepaidBadDebts,proto3" json:"prepaid_bad_debts,omitempty"`
	WhitelistedAddresses []string          `protobuf:"bytes,5,rep,name=whitelisted_addresses,json=whitelistedAddresses,proto3" json:"whitelisted_addresses,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_24e163498ed621a8, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPairMetadata() []*PairMetadata {
	if m != nil {
		return m.PairMetadata
	}
	return nil
}

func (m *GenesisState) GetPositions() []*Position {
	if m != nil {
		return m.Positions
	}
	return nil
}

func (m *GenesisState) GetPrepaidBadDebts() []*PrepaidBadDebt {
	if m != nil {
		return m.PrepaidBadDebts
	}
	return nil
}

func (m *GenesisState) GetWhitelistedAddresses() []string {
	if m != nil {
		return m.WhitelistedAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nibiru.perp.v1.GenesisState")
}

func init() { proto.RegisterFile("perp/v1/genesis.proto", fileDescriptor_24e163498ed621a8) }

var fileDescriptor_24e163498ed621a8 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0xab, 0x13, 0x31,
	0x14, 0x85, 0x67, 0x5e, 0x9f, 0x0f, 0xde, 0xbc, 0xaa, 0x38, 0xb6, 0x32, 0x94, 0x12, 0x8b, 0xab,
	0xe2, 0x62, 0xc2, 0xb4, 0xe2, 0xbe, 0xb5, 0x20, 0x08, 0x8a, 0x8c, 0x3b, 0x37, 0xc3, 0xcd, 0xe4,
	0x32, 0x0d, 0x74, 0x92, 0x90, 0xa4, 0x55, 0xff, 0x85, 0x3f, 0xc9, 0x65, 0x97, 0x5d, 0xba, 0x12,
	0x69, 0xff, 0x88, 0x34, 0x33, 0xc5, 0x96, 0xae, 0x92, 0x7b, 0xce, 0x77, 0xc2, 0x09, 0x37, 0xea,
	0x6b, 0x34, 0x9a, 0x6e, 0x32, 0x5a, 0xa1, 0x44, 0x2b, 0x6c, 0xaa, 0x8d, 0x72, 0x2a, 0x7e, 0x22,
	0x05, 0x13, 0x66, 0x9d, 0x1e, 0xdd, 0x74, 0x93, 0x0d, 0x7a, 0x95, 0xaa, 0x94, 0xb7, 0xe8, 0xf1,
	0xd6, 0x50, 0x83, 0x61, 0xa5, 0x54, 0xb5, 0x42, 0x0a, 0x5a, 0x50, 0x90, 0x52, 0x39, 0x70, 0x42,
	0xc9, 0xf6, 0x8d, 0x01, 0x29, 0x95, 0xad, 0x95, 0xa5, 0x0c, 0x2c, 0xd2, 0x4d, 0xc6, 0xd0, 0x41,
	0x46, 0x4b, 0x25, 0x64, 0xeb, 0x3f, 0x2f, 0x55, 0x5d, 0x2b, 0x49, 0x9b, 0xe3, 0x24, 0x9e, 0xfa,
	0x58, 0x07, 0x0e, 0x1b, 0xf1, 0xd5, 0xaf, 0x9b, 0xa8, 0xfb, 0xbe, 0xe9, 0xf7, 0xe5, 0x28, 0xc7,
	0x6f, 0xa2, 0x3b, 0x0d, 0x06, 0x6a, 0x9b, 0x84, 0xa3, 0x70, 0xfc, 0x30, 0x79, 0x91, 0x5e, 0xf6,
	0x4d, 0x3f, 0x7b, 0x77, 0x7e, 0xbb, 0xfd, 0xf3, 0x32, 0xc8, 0x5b, 0x36, 0x9e, 0x45, 0x8f, 0x35,
	0x08, 0x53, 0xd4, 0xe8, 0x80, 0x83, 0x83, 0xe4, 0x66, 0xd4, 0x19, 0x3f, 0x4c, 0x86, 0xd7, 0x61,
	0x61, 0x3e, 0xb6, 0x4c, 0xde, 0xd5, 0x67, 0x53, 0xfc, 0x36, 0xba, 0xd7, 0xca, 0x0a, 0xff, 0xcd,
	0xa4, 0xe3, 0xe3, 0xc9, 0x55, 0xbc, 0x05, 0xf2, 0xff, 0x68, 0xfc, 0x21, 0x7a, 0xa6, 0x0d, 0x6a,
	0x10, 0xbc, 0x60, 0xc0, 0x0b, 0x8e, 0xcc, 0xd9, 0xe4, 0xd6, 0xe7, 0xc9, 0x55, 0xbe, 0x01, 0xe7,
	0xc0, 0x17, 0xc8, 0x5c, 0xfe, 0x54, 0x5f, 0xcc, 0x36, 0x9e, 0x46, 0xfd, 0x6f, 0x4b, 0xe1, 0x70,
	0x25, 0xac, 0x43, 0x5e, 0x00, 0xe7, 0x06, 0xad, 0x45, 0x9b, 0x3c, 0x1a, 0x75, 0xc6, 0xf7, 0x79,
	0xef, 0xcc, 0x9c, 0x9d, 0xbc, 0xf9, 0x62, 0xbb, 0x27, 0xe1, 0x6e, 0x4f, 0xc2, 0xbf, 0x7b, 0x12,
	0xfe, 0x3c, 0x90, 0x60, 0x77, 0x20, 0xc1, 0xef, 0x03, 0x09, 0xbe, 0xbe, 0xae, 0x84, 0x5b, 0xae,
	0x59, 0x5a, 0xaa, 0x9a, 0x7e, 0xf2, 0x4d, 0xde, 0x2d, 0x41, 0x48, 0xda, 0xb4, 0xa2, 0xdf, 0xa9,
	0xdf, 0x88, 0xfb, 0xa1, 0xd1, 0xb2, 0x3b, 0xbf, 0x8f, 0xe9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x2d, 0x98, 0x28, 0x86, 0x36, 0x02, 0x00, 0x00,
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
	if len(m.WhitelistedAddresses) > 0 {
		for iNdEx := len(m.WhitelistedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WhitelistedAddresses[iNdEx])
			copy(dAtA[i:], m.WhitelistedAddresses[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.WhitelistedAddresses[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PrepaidBadDebts) > 0 {
		for iNdEx := len(m.PrepaidBadDebts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PrepaidBadDebts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Positions) > 0 {
		for iNdEx := len(m.Positions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Positions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PairMetadata) > 0 {
		for iNdEx := len(m.PairMetadata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PairMetadata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.PairMetadata) > 0 {
		for _, e := range m.PairMetadata {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Positions) > 0 {
		for _, e := range m.Positions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PrepaidBadDebts) > 0 {
		for _, e := range m.PrepaidBadDebts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.WhitelistedAddresses) > 0 {
		for _, s := range m.WhitelistedAddresses {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairMetadata = append(m.PairMetadata, &PairMetadata{})
			if err := m.PairMetadata[len(m.PairMetadata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Positions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Positions = append(m.Positions, &Position{})
			if err := m.Positions[len(m.Positions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrepaidBadDebts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrepaidBadDebts = append(m.PrepaidBadDebts, &PrepaidBadDebt{})
			if err := m.PrepaidBadDebts[len(m.PrepaidBadDebts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhitelistedAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhitelistedAddresses = append(m.WhitelistedAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)