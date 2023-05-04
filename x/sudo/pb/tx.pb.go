// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sudo/v1/tx.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgEditSudoers: Msg to update the "Sudoers" state.
type MsgEditSudoers struct {
	// Action: identifier for the type of edit that will take place. Using this
	//   action field prevents us from needing to create several similar message
	//   types.
	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	// Contracts: An input payload.
	Contracts []string `protobuf:"bytes,2,rep,name=contracts,proto3" json:"contracts,omitempty"`
	// Sender: Address for the signer of the transaction.
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgEditSudoers) Reset()         { *m = MsgEditSudoers{} }
func (m *MsgEditSudoers) String() string { return proto.CompactTextString(m) }
func (*MsgEditSudoers) ProtoMessage()    {}
func (*MsgEditSudoers) Descriptor() ([]byte, []int) {
	return fileDescriptor_14aeb0702eb5e8b0, []int{0}
}
func (m *MsgEditSudoers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEditSudoers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEditSudoers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEditSudoers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEditSudoers.Merge(m, src)
}
func (m *MsgEditSudoers) XXX_Size() int {
	return m.Size()
}
func (m *MsgEditSudoers) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEditSudoers.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEditSudoers proto.InternalMessageInfo

func (m *MsgEditSudoers) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *MsgEditSudoers) GetContracts() []string {
	if m != nil {
		return m.Contracts
	}
	return nil
}

func (m *MsgEditSudoers) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// MsgEditSudoersResponse indicates the successful execution of MsgEditSudeors.
type MsgEditSudoersResponse struct {
}

func (m *MsgEditSudoersResponse) Reset()         { *m = MsgEditSudoersResponse{} }
func (m *MsgEditSudoersResponse) String() string { return proto.CompactTextString(m) }
func (*MsgEditSudoersResponse) ProtoMessage()    {}
func (*MsgEditSudoersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_14aeb0702eb5e8b0, []int{1}
}
func (m *MsgEditSudoersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEditSudoersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEditSudoersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEditSudoersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEditSudoersResponse.Merge(m, src)
}
func (m *MsgEditSudoersResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgEditSudoersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEditSudoersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEditSudoersResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgEditSudoers)(nil), "nibiru.sudo.v1.MsgEditSudoers")
	proto.RegisterType((*MsgEditSudoersResponse)(nil), "nibiru.sudo.v1.MsgEditSudoersResponse")
}

func init() { proto.RegisterFile("sudo/v1/tx.proto", fileDescriptor_14aeb0702eb5e8b0) }

var fileDescriptor_14aeb0702eb5e8b0 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xf3, 0x30,
	0x18, 0xc7, 0xdb, 0x15, 0x06, 0xcd, 0x0b, 0xe3, 0xa5, 0xc8, 0xa8, 0x75, 0x84, 0xd9, 0x83, 0xee,
	0x94, 0x30, 0xfd, 0x04, 0x2a, 0x1e, 0xe7, 0xa1, 0xde, 0x3c, 0x28, 0x69, 0x1b, 0xb2, 0x80, 0xe6,
	0x29, 0x4d, 0x3a, 0x7a, 0x13, 0xfc, 0x04, 0x82, 0x5f, 0xca, 0xe3, 0xc0, 0x8b, 0x47, 0x69, 0xfd,
	0x20, 0xd2, 0xc6, 0xa1, 0xbd, 0x78, 0xcb, 0x93, 0xff, 0xff, 0xf7, 0x0b, 0x79, 0xd0, 0x7f, 0x5d,
	0xe5, 0x40, 0x37, 0x4b, 0x6a, 0x6a, 0x52, 0x94, 0x60, 0x20, 0x98, 0x28, 0x99, 0xca, 0xb2, 0x22,
	0x5d, 0x40, 0x36, 0xcb, 0x68, 0x4f, 0x80, 0x80, 0x3e, 0xa2, 0xdd, 0xc9, 0xb6, 0xa2, 0x99, 0x00,
	0x10, 0xf7, 0x9c, 0xb2, 0x42, 0x52, 0xa6, 0x14, 0x18, 0x66, 0x24, 0x28, 0x6d, 0xd3, 0xf8, 0x16,
	0x4d, 0x56, 0x5a, 0x5c, 0xe6, 0xd2, 0x5c, 0x57, 0x39, 0xf0, 0x52, 0x07, 0x53, 0x34, 0x66, 0x59,
	0x57, 0x09, 0xdd, 0xb9, 0xbb, 0xf0, 0x93, 0xef, 0x29, 0x98, 0x21, 0x3f, 0x03, 0x65, 0x4a, 0x96,
	0x19, 0x1d, 0x8e, 0xe6, 0xde, 0xc2, 0x4f, 0x7e, 0x2e, 0x3a, 0x4a, 0x73, 0x95, 0xf3, 0x32, 0xf4,
	0x2c, 0x65, 0xa7, 0x38, 0x44, 0xd3, 0xa1, 0x3f, 0xe1, 0xba, 0x00, 0xa5, 0xf9, 0xc9, 0x23, 0xf2,
	0x56, 0x5a, 0x04, 0x35, 0xfa, 0xf7, 0xfb, 0x75, 0x4c, 0x86, 0x9f, 0x22, 0x43, 0x3a, 0x3a, 0xfa,
	0x3b, 0xdf, 0xd9, 0xe3, 0xc3, 0xa7, 0xb7, 0xcf, 0x97, 0xd1, 0x41, 0xbc, 0x4f, 0x6d, 0x9f, 0xf6,
	0xdb, 0xe3, 0xb9, 0x34, 0x77, 0xda, 0x56, 0xcf, 0xcf, 0x5e, 0x1b, 0xec, 0x6e, 0x1b, 0xec, 0x7e,
	0x34, 0xd8, 0x7d, 0x6e, 0xb1, 0xb3, 0x6d, 0xb1, 0xf3, 0xde, 0x62, 0xe7, 0xe6, 0x58, 0x48, 0xb3,
	0xae, 0x52, 0x92, 0xc1, 0x03, 0xbd, 0xea, 0xf1, 0x8b, 0x35, 0x93, 0x6a, 0xa7, 0xaa, 0xad, 0xac,
	0x48, 0xd3, 0x71, 0xbf, 0xc4, 0xd3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0xce, 0x06, 0x0b,
	0x9c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// EditSudoers updates the "Sudoers" state
	EditSudoers(ctx context.Context, in *MsgEditSudoers, opts ...grpc.CallOption) (*MsgEditSudoersResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) EditSudoers(ctx context.Context, in *MsgEditSudoers, opts ...grpc.CallOption) (*MsgEditSudoersResponse, error) {
	out := new(MsgEditSudoersResponse)
	err := c.cc.Invoke(ctx, "/nibiru.sudo.v1.Msg/EditSudoers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// EditSudoers updates the "Sudoers" state
	EditSudoers(context.Context, *MsgEditSudoers) (*MsgEditSudoersResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) EditSudoers(ctx context.Context, req *MsgEditSudoers) (*MsgEditSudoersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditSudoers not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_EditSudoers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditSudoers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditSudoers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nibiru.sudo.v1.Msg/EditSudoers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditSudoers(ctx, req.(*MsgEditSudoers))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nibiru.sudo.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EditSudoers",
			Handler:    _Msg_EditSudoers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sudo/v1/tx.proto",
}

func (m *MsgEditSudoers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEditSudoers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEditSudoers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contracts) > 0 {
		for iNdEx := len(m.Contracts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contracts[iNdEx])
			copy(dAtA[i:], m.Contracts[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Contracts[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Action) > 0 {
		i -= len(m.Action)
		copy(dAtA[i:], m.Action)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Action)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgEditSudoersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEditSudoersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEditSudoersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgEditSudoers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Contracts) > 0 {
		for _, s := range m.Contracts {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgEditSudoersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgEditSudoers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgEditSudoers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEditSudoers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contracts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contracts = append(m.Contracts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgEditSudoersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgEditSudoersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEditSudoersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
