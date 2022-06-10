// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ride/stored_ride.proto

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

type StoredRide struct {
	Index       string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Driver      string `protobuf:"bytes,3,opt,name=driver,proto3" json:"driver,omitempty"`
	Passenger   string `protobuf:"bytes,4,opt,name=passenger,proto3" json:"passenger,omitempty"`
	MutualStake uint64 `protobuf:"varint,5,opt,name=mutualStake,proto3" json:"mutualStake,omitempty"`
	PayPerHour  uint64 `protobuf:"varint,6,opt,name=payPerHour,proto3" json:"payPerHour,omitempty"`
	DistanceTip uint64 `protobuf:"varint,7,opt,name=distanceTip,proto3" json:"distanceTip,omitempty"`
}

func (m *StoredRide) Reset()         { *m = StoredRide{} }
func (m *StoredRide) String() string { return proto.CompactTextString(m) }
func (*StoredRide) ProtoMessage()    {}
func (*StoredRide) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d539ef8589d7f08, []int{0}
}
func (m *StoredRide) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredRide) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredRide.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredRide) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredRide.Merge(m, src)
}
func (m *StoredRide) XXX_Size() int {
	return m.Size()
}
func (m *StoredRide) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredRide.DiscardUnknown(m)
}

var xxx_messageInfo_StoredRide proto.InternalMessageInfo

func (m *StoredRide) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StoredRide) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *StoredRide) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *StoredRide) GetPassenger() string {
	if m != nil {
		return m.Passenger
	}
	return ""
}

func (m *StoredRide) GetMutualStake() uint64 {
	if m != nil {
		return m.MutualStake
	}
	return 0
}

func (m *StoredRide) GetPayPerHour() uint64 {
	if m != nil {
		return m.PayPerHour
	}
	return 0
}

func (m *StoredRide) GetDistanceTip() uint64 {
	if m != nil {
		return m.DistanceTip
	}
	return 0
}

func init() {
	proto.RegisterType((*StoredRide)(nil), "smarshallspitzbart.ride.ride.StoredRide")
}

func init() { proto.RegisterFile("ride/stored_ride.proto", fileDescriptor_4d539ef8589d7f08) }

var fileDescriptor_4d539ef8589d7f08 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x63, 0x68, 0x83, 0x6a, 0x36, 0x0b, 0x55, 0x1e, 0x2a, 0x2b, 0x62, 0xea, 0x42, 0x82,
	0xc4, 0x0d, 0x98, 0x10, 0x13, 0x4a, 0x99, 0x58, 0x90, 0x53, 0x3f, 0xb5, 0x16, 0x49, 0x6c, 0xd9,
	0x2f, 0xa8, 0xe5, 0x14, 0x1c, 0x8b, 0xb1, 0x23, 0x13, 0x42, 0xc9, 0x45, 0x50, 0x1c, 0xd4, 0x66,
	0xb1, 0xfd, 0xbe, 0xdf, 0xfe, 0x65, 0x7d, 0x74, 0xee, 0xb4, 0x82, 0xcc, 0xa3, 0x71, 0xa0, 0x5e,
	0xfb, 0x73, 0x6a, 0x9d, 0x41, 0xc3, 0x16, 0xbe, 0x92, 0xce, 0x6f, 0x65, 0x59, 0x7a, 0xab, 0xf1,
	0xa3, 0x90, 0x0e, 0xd3, 0x10, 0xf7, 0xcb, 0xf5, 0x0f, 0xa1, 0x74, 0x15, 0xde, 0xe4, 0x5a, 0x01,
	0xbb, 0xa2, 0x53, 0x5d, 0x2b, 0xd8, 0x71, 0x92, 0x90, 0xe5, 0x2c, 0x1f, 0x06, 0x96, 0xd0, 0x4b,
	0x05, 0x1e, 0x75, 0x2d, 0x51, 0x9b, 0x9a, 0x9f, 0x85, 0x6c, 0x8c, 0xd8, 0x9c, 0xc6, 0xca, 0xe9,
	0x77, 0x70, 0xfc, 0x3c, 0x84, 0xff, 0x13, 0x5b, 0xd0, 0x99, 0x95, 0xde, 0x43, 0xbd, 0x01, 0xc7,
	0x27, 0x21, 0x3a, 0x81, 0xbe, 0xb7, 0x6a, 0xb0, 0x91, 0xe5, 0x0a, 0xe5, 0x1b, 0xf0, 0x69, 0x42,
	0x96, 0x93, 0x7c, 0x8c, 0x98, 0xa0, 0xd4, 0xca, 0xfd, 0x13, 0xb8, 0x07, 0xd3, 0x38, 0x1e, 0x87,
	0x0b, 0x23, 0x12, 0x7e, 0xa6, 0x3d, 0xca, 0x7a, 0x0d, 0xcf, 0xda, 0xf2, 0x8b, 0xa1, 0x61, 0x84,
	0xee, 0x1f, 0xbf, 0x5a, 0x41, 0x0e, 0xad, 0x20, 0xbf, 0xad, 0x20, 0x9f, 0x9d, 0x88, 0x0e, 0x9d,
	0x88, 0xbe, 0x3b, 0x11, 0xbd, 0xdc, 0x6e, 0x34, 0x6e, 0x9b, 0x22, 0x5d, 0x9b, 0x2a, 0x3b, 0x3a,
	0xba, 0x39, 0x4a, 0xca, 0x82, 0xcf, 0xdd, 0xb0, 0xe1, 0xde, 0x82, 0x2f, 0xe2, 0x60, 0xf4, 0xee,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0xde, 0x1c, 0x51, 0x23, 0x6b, 0x01, 0x00, 0x00,
}

func (m *StoredRide) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredRide) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredRide) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DistanceTip != 0 {
		i = encodeVarintStoredRide(dAtA, i, uint64(m.DistanceTip))
		i--
		dAtA[i] = 0x38
	}
	if m.PayPerHour != 0 {
		i = encodeVarintStoredRide(dAtA, i, uint64(m.PayPerHour))
		i--
		dAtA[i] = 0x30
	}
	if m.MutualStake != 0 {
		i = encodeVarintStoredRide(dAtA, i, uint64(m.MutualStake))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Passenger) > 0 {
		i -= len(m.Passenger)
		copy(dAtA[i:], m.Passenger)
		i = encodeVarintStoredRide(dAtA, i, uint64(len(m.Passenger)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Driver) > 0 {
		i -= len(m.Driver)
		copy(dAtA[i:], m.Driver)
		i = encodeVarintStoredRide(dAtA, i, uint64(len(m.Driver)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Destination) > 0 {
		i -= len(m.Destination)
		copy(dAtA[i:], m.Destination)
		i = encodeVarintStoredRide(dAtA, i, uint64(len(m.Destination)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStoredRide(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStoredRide(dAtA []byte, offset int, v uint64) int {
	offset -= sovStoredRide(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoredRide) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStoredRide(uint64(l))
	}
	l = len(m.Destination)
	if l > 0 {
		n += 1 + l + sovStoredRide(uint64(l))
	}
	l = len(m.Driver)
	if l > 0 {
		n += 1 + l + sovStoredRide(uint64(l))
	}
	l = len(m.Passenger)
	if l > 0 {
		n += 1 + l + sovStoredRide(uint64(l))
	}
	if m.MutualStake != 0 {
		n += 1 + sovStoredRide(uint64(m.MutualStake))
	}
	if m.PayPerHour != 0 {
		n += 1 + sovStoredRide(uint64(m.PayPerHour))
	}
	if m.DistanceTip != 0 {
		n += 1 + sovStoredRide(uint64(m.DistanceTip))
	}
	return n
}

func sovStoredRide(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStoredRide(x uint64) (n int) {
	return sovStoredRide(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoredRide) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStoredRide
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
			return fmt.Errorf("proto: StoredRide: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredRide: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredRide
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
				return ErrInvalidLengthStoredRide
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredRide
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destination", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredRide
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
				return ErrInvalidLengthStoredRide
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredRide
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destination = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Driver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredRide
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
				return ErrInvalidLengthStoredRide
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredRide
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Driver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Passenger", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredRide
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
				return ErrInvalidLengthStoredRide
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredRide
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Passenger = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutualStake", wireType)
			}
			m.MutualStake = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredRide
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MutualStake |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayPerHour", wireType)
			}
			m.PayPerHour = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredRide
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PayPerHour |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistanceTip", wireType)
			}
			m.DistanceTip = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredRide
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistanceTip |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStoredRide(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStoredRide
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
func skipStoredRide(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStoredRide
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
					return 0, ErrIntOverflowStoredRide
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
					return 0, ErrIntOverflowStoredRide
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
				return 0, ErrInvalidLengthStoredRide
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStoredRide
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStoredRide
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStoredRide        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStoredRide          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStoredRide = fmt.Errorf("proto: unexpected end of group")
)
