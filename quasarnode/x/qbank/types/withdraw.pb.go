// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qbank/withdraw.proto

package types

import (
	fmt "fmt"
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

type Withdraw struct {
	Id                  uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RiskProfile         string `protobuf:"bytes,2,opt,name=riskProfile,proto3" json:"riskProfile,omitempty"`
	VaultID             string `protobuf:"bytes,3,opt,name=vaultID,proto3" json:"vaultID,omitempty"`
	DepositorAccAddress string `protobuf:"bytes,4,opt,name=depositorAccAddress,proto3" json:"depositorAccAddress,omitempty"`
	Amount              string `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Denom               string `protobuf:"bytes,6,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *Withdraw) Reset()         { *m = Withdraw{} }
func (m *Withdraw) String() string { return proto.CompactTextString(m) }
func (*Withdraw) ProtoMessage()    {}
func (*Withdraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_77bcdb9426b09e8b, []int{0}
}
func (m *Withdraw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Withdraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Withdraw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Withdraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Withdraw.Merge(m, src)
}
func (m *Withdraw) XXX_Size() int {
	return m.Size()
}
func (m *Withdraw) XXX_DiscardUnknown() {
	xxx_messageInfo_Withdraw.DiscardUnknown(m)
}

var xxx_messageInfo_Withdraw proto.InternalMessageInfo

func (m *Withdraw) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Withdraw) GetRiskProfile() string {
	if m != nil {
		return m.RiskProfile
	}
	return ""
}

func (m *Withdraw) GetVaultID() string {
	if m != nil {
		return m.VaultID
	}
	return ""
}

func (m *Withdraw) GetDepositorAccAddress() string {
	if m != nil {
		return m.DepositorAccAddress
	}
	return ""
}

func (m *Withdraw) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *Withdraw) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*Withdraw)(nil), "abag.quasarnode.qbank.Withdraw")
}

func init() { proto.RegisterFile("qbank/withdraw.proto", fileDescriptor_77bcdb9426b09e8b) }

var fileDescriptor_77bcdb9426b09e8b = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0xe3, 0xd0, 0x06, 0x30, 0x12, 0x83, 0x29, 0xc8, 0x93, 0x15, 0x31, 0x65, 0x8a, 0x91,
	0x38, 0x41, 0x2b, 0x16, 0x36, 0xd4, 0x05, 0x89, 0xcd, 0xc9, 0x6f, 0x5a, 0xab, 0x4d, 0xfe, 0xd4,
	0x76, 0x28, 0xdc, 0x82, 0xd3, 0x70, 0x06, 0xc6, 0x8e, 0x8c, 0x28, 0xb9, 0x08, 0xc2, 0x0d, 0x82,
	0x81, 0xf1, 0x7b, 0xef, 0x5b, 0xde, 0xa3, 0x93, 0x4d, 0xa1, 0xea, 0x95, 0xdc, 0x1a, 0xbf, 0x04,
	0xab, 0xb6, 0x79, 0x63, 0xd1, 0x23, 0x3b, 0x57, 0x85, 0x5a, 0xe4, 0x9b, 0x56, 0x39, 0x65, 0x6b,
	0x04, 0x9d, 0x07, 0xeb, 0xf2, 0x8d, 0xd0, 0xa3, 0xfb, 0xc1, 0x64, 0xa7, 0x34, 0x36, 0xc0, 0x49,
	0x4a, 0xb2, 0xd1, 0x3c, 0x36, 0xc0, 0x52, 0x7a, 0x62, 0x8d, 0x5b, 0xdd, 0x59, 0x7c, 0x34, 0x6b,
	0xcd, 0xe3, 0x94, 0x64, 0xc7, 0xf3, 0xbf, 0x11, 0xe3, 0xf4, 0xf0, 0x49, 0xb5, 0x6b, 0x7f, 0x7b,
	0xc3, 0x0f, 0x42, 0xfb, 0x83, 0xec, 0x8a, 0x9e, 0x81, 0x6e, 0xd0, 0x19, 0x8f, 0x76, 0x5a, 0x96,
	0x53, 0x00, 0xab, 0x9d, 0xe3, 0xa3, 0x60, 0xfd, 0x57, 0xb1, 0x0b, 0x9a, 0xa8, 0x0a, 0xdb, 0xda,
	0xf3, 0x71, 0x90, 0x06, 0x62, 0x13, 0x3a, 0x06, 0x5d, 0x63, 0xc5, 0x93, 0x10, 0xef, 0x61, 0x36,
	0x7b, 0xef, 0x04, 0xd9, 0x75, 0x82, 0x7c, 0x76, 0x82, 0xbc, 0xf6, 0x22, 0xda, 0xf5, 0x22, 0xfa,
	0xe8, 0x45, 0xf4, 0x90, 0x2d, 0x8c, 0x5f, 0xb6, 0x45, 0x5e, 0x62, 0x25, 0xbf, 0x47, 0xcb, 0xdf,
	0xd1, 0xf2, 0x59, 0xee, 0xcf, 0xf1, 0x2f, 0x8d, 0x76, 0x45, 0x12, 0xae, 0xb9, 0xfe, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x8c, 0x91, 0x88, 0x5c, 0x32, 0x01, 0x00, 0x00,
}

func (m *Withdraw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Withdraw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Withdraw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DepositorAccAddress) > 0 {
		i -= len(m.DepositorAccAddress)
		copy(dAtA[i:], m.DepositorAccAddress)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.DepositorAccAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.VaultID) > 0 {
		i -= len(m.VaultID)
		copy(dAtA[i:], m.VaultID)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.VaultID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RiskProfile) > 0 {
		i -= len(m.RiskProfile)
		copy(dAtA[i:], m.RiskProfile)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.RiskProfile)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintWithdraw(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintWithdraw(dAtA []byte, offset int, v uint64) int {
	offset -= sovWithdraw(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Withdraw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovWithdraw(uint64(m.Id))
	}
	l = len(m.RiskProfile)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	l = len(m.VaultID)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	l = len(m.DepositorAccAddress)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	return n
}

func sovWithdraw(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWithdraw(x uint64) (n int) {
	return sovWithdraw(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Withdraw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWithdraw
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
			return fmt.Errorf("proto: Withdraw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Withdraw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RiskProfile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RiskProfile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VaultID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositorAccAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositorAccAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWithdraw(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWithdraw
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
func skipWithdraw(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWithdraw
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
					return 0, ErrIntOverflowWithdraw
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
					return 0, ErrIntOverflowWithdraw
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
				return 0, ErrInvalidLengthWithdraw
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWithdraw
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWithdraw
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWithdraw        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWithdraw          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWithdraw = fmt.Errorf("proto: unexpected end of group")
)