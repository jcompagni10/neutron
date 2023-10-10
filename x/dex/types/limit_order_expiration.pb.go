// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/dex/limit_order_expiration.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LimitOrderExpiration struct {
	// see limitOrderTranche.proto for details on goodTilDate
	ExpirationTime time.Time `protobuf:"bytes,1,opt,name=expirationTime,proto3,stdtime" json:"expirationTime"`
	TrancheRef     []byte    `protobuf:"bytes,2,opt,name=trancheRef,proto3" json:"trancheRef,omitempty"`
}

func (m *LimitOrderExpiration) Reset()         { *m = LimitOrderExpiration{} }
func (m *LimitOrderExpiration) String() string { return proto.CompactTextString(m) }
func (*LimitOrderExpiration) ProtoMessage()    {}
func (*LimitOrderExpiration) Descriptor() ([]byte, []int) {
	return fileDescriptor_61264397cad6ae82, []int{0}
}
func (m *LimitOrderExpiration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitOrderExpiration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitOrderExpiration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitOrderExpiration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderExpiration.Merge(m, src)
}
func (m *LimitOrderExpiration) XXX_Size() int {
	return m.Size()
}
func (m *LimitOrderExpiration) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderExpiration.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderExpiration proto.InternalMessageInfo

func (m *LimitOrderExpiration) GetExpirationTime() time.Time {
	if m != nil {
		return m.ExpirationTime
	}
	return time.Time{}
}

func (m *LimitOrderExpiration) GetTrancheRef() []byte {
	if m != nil {
		return m.TrancheRef
	}
	return nil
}

func init() {
	proto.RegisterType((*LimitOrderExpiration)(nil), "neutron.dex.LimitOrderExpiration")
}

func init() {
	proto.RegisterFile("neutron/dex/limit_order_expiration.proto", fileDescriptor_61264397cad6ae82)
}

var fileDescriptor_61264397cad6ae82 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc8, 0x4b, 0x2d, 0x2d,
	0x29, 0xca, 0xcf, 0xd3, 0x4f, 0x49, 0xad, 0xd0, 0xcf, 0xc9, 0xcc, 0xcd, 0x2c, 0x89, 0xcf, 0x2f,
	0x4a, 0x49, 0x2d, 0x8a, 0x4f, 0xad, 0x28, 0xc8, 0x2c, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0xaa, 0xd4, 0x4b, 0x49, 0xad, 0x90, 0x92, 0x4f, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x4b, 0x25, 0x95, 0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6, 0x16,
	0x97, 0x24, 0xe6, 0x16, 0x40, 0x54, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x99, 0xfa, 0x20,
	0x16, 0x44, 0x54, 0xa9, 0x85, 0x91, 0x4b, 0xc4, 0x07, 0x64, 0x89, 0x3f, 0xc8, 0x0e, 0x57, 0xb8,
	0x15, 0x42, 0x3e, 0x5c, 0x7c, 0x08, 0x0b, 0x43, 0x32, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0xb8, 0x8d, 0xa4, 0xf4, 0x20, 0x16, 0xe9, 0xc1, 0x2c, 0xd2, 0x0b, 0x81, 0x59, 0xe4, 0xc4, 0x71,
	0xe2, 0x9e, 0x3c, 0xc3, 0x84, 0xfb, 0xf2, 0x8c, 0x41, 0x68, 0x7a, 0x85, 0xe4, 0xb8, 0xb8, 0x4a,
	0x8a, 0x12, 0xf3, 0x92, 0x33, 0x52, 0x83, 0x52, 0xd3, 0x24, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82,
	0x90, 0x44, 0x9c, 0x5c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39,
	0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x2b,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xea, 0x5f, 0xdd, 0xfc, 0xa2,
	0x74, 0x18, 0x5b, 0xbf, 0x02, 0x1c, 0x4e, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x37,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xd9, 0xbe, 0xa6, 0x43, 0x01, 0x00, 0x00,
}

func (m *LimitOrderExpiration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitOrderExpiration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitOrderExpiration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TrancheRef) > 0 {
		i -= len(m.TrancheRef)
		copy(dAtA[i:], m.TrancheRef)
		i = encodeVarintLimitOrderExpiration(dAtA, i, uint64(len(m.TrancheRef)))
		i--
		dAtA[i] = 0x12
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.ExpirationTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.ExpirationTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLimitOrderExpiration(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintLimitOrderExpiration(dAtA []byte, offset int, v uint64) int {
	offset -= sovLimitOrderExpiration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LimitOrderExpiration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.ExpirationTime)
	n += 1 + l + sovLimitOrderExpiration(uint64(l))
	l = len(m.TrancheRef)
	if l > 0 {
		n += 1 + l + sovLimitOrderExpiration(uint64(l))
	}
	return n
}

func sovLimitOrderExpiration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLimitOrderExpiration(x uint64) (n int) {
	return sovLimitOrderExpiration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LimitOrderExpiration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitOrderExpiration
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
			return fmt.Errorf("proto: LimitOrderExpiration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitOrderExpiration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderExpiration
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
				return ErrInvalidLengthLimitOrderExpiration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderExpiration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.ExpirationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrancheRef", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderExpiration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLimitOrderExpiration
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderExpiration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrancheRef = append(m.TrancheRef[:0], dAtA[iNdEx:postIndex]...)
			if m.TrancheRef == nil {
				m.TrancheRef = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLimitOrderExpiration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitOrderExpiration
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
func skipLimitOrderExpiration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLimitOrderExpiration
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
					return 0, ErrIntOverflowLimitOrderExpiration
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
					return 0, ErrIntOverflowLimitOrderExpiration
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
				return 0, ErrInvalidLengthLimitOrderExpiration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLimitOrderExpiration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLimitOrderExpiration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLimitOrderExpiration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLimitOrderExpiration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLimitOrderExpiration = fmt.Errorf("proto: unexpected end of group")
)
