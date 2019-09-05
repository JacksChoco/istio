// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/opa/config/config.proto

// The `opa` adapter exposes an [Open Policy Agent](http://www.openpolicyagent.org) engine
// that provides sophisticated access control mechanisms.
//
// This adapter supports the [authorization template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/authorization/).

package config

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Configuration format for the `opa` adapter.
//
// Example configuration:
// ```yaml
// policy:
//   - |+
//     package mixerauthz
//     policy = [
//       {
//         "rule": {
//           "verbs": [
//             "storage.buckets.get"
//           ],
//           "users": [
//             "bucket-admins"
//           ]
//         }
//       }
//     ]
//
//     default allow = false
//
//     allow = true {
//       rule = policy[_].rule
//       input.subject.user = rule.users[_]
//       input.action.method = rule.verbs[_]
//     }
// checkMethod: "data.mixerauthz.allow"
// failClose: true
// ```
type Params struct {
	// List of OPA policies
	Policy []string `protobuf:"bytes,1,rep,name=policy,proto3" json:"policy,omitempty"`
	// Query method to check.
	// Format: data.<package name>.<method name>
	CheckMethod string `protobuf:"bytes,2,opt,name=check_method,json=checkMethod,proto3" json:"check_method,omitempty"`
	// Close the client request when adapter has a issue.
	// If failClose is set to true and there is a runtime error,
	// instead of disabling the adapter, close the client request
	FailClose bool `protobuf:"varint,3,opt,name=fail_close,json=failClose,proto3" json:"fail_close,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_05827bfc1c8a686c, []int{0}
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

func init() {
	proto.RegisterType((*Params)(nil), "adapter.opa.config.Params")
}

func init() {
	proto.RegisterFile("mixer/adapter/opa/config/config.proto", fileDescriptor_05827bfc1c8a686c)
}

var fileDescriptor_05827bfc1c8a686c = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0xcd, 0xac, 0x48,
	0x2d, 0xd2, 0x4f, 0x4c, 0x49, 0x2c, 0x28, 0x49, 0x2d, 0xd2, 0xcf, 0x2f, 0x48, 0xd4, 0x4f, 0xce,
	0xcf, 0x4b, 0xcb, 0x4c, 0x87, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x42, 0x50, 0x05,
	0x7a, 0xf9, 0x05, 0x89, 0x7a, 0x10, 0x19, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0xb4, 0x3e,
	0x88, 0x05, 0x51, 0xa9, 0x94, 0xc4, 0xc5, 0x16, 0x90, 0x58, 0x94, 0x98, 0x5b, 0x2c, 0x24, 0xc6,
	0xc5, 0x56, 0x90, 0x9f, 0x93, 0x99, 0x5c, 0x29, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x19, 0x04, 0xe5,
	0x09, 0x29, 0x72, 0xf1, 0x24, 0x67, 0xa4, 0x26, 0x67, 0xc7, 0xe7, 0xa6, 0x96, 0x64, 0xe4, 0xa7,
	0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x71, 0x83, 0xc5, 0x7c, 0xc1, 0x42, 0x42, 0xb2, 0x5c,
	0x5c, 0x69, 0x89, 0x99, 0x39, 0xf1, 0xc9, 0x39, 0xf9, 0xc5, 0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a,
	0x1c, 0x41, 0x9c, 0x20, 0x11, 0x67, 0x90, 0x80, 0x93, 0xc3, 0x89, 0x87, 0x72, 0x0c, 0x17, 0x1e,
	0xca, 0x31, 0xdc, 0x78, 0x28, 0xc7, 0xf0, 0xe1, 0xa1, 0x1c, 0x43, 0xc3, 0x23, 0x39, 0xc6, 0x15,
	0x8f, 0xe4, 0x18, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x17, 0x8f, 0xe4, 0x18, 0x3e, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39,
	0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xd8, 0x20, 0x6e, 0x4f, 0x62, 0x03, 0x3b, 0xd6, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x8e, 0xb7, 0xf6, 0xf2, 0xff, 0x00, 0x00, 0x00,
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
	if m.FailClose {
		i--
		if m.FailClose {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.CheckMethod) > 0 {
		i -= len(m.CheckMethod)
		copy(dAtA[i:], m.CheckMethod)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.CheckMethod)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Policy) > 0 {
		for iNdEx := len(m.Policy) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Policy[iNdEx])
			copy(dAtA[i:], m.Policy[iNdEx])
			i = encodeVarintConfig(dAtA, i, uint64(len(m.Policy[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
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
	if len(m.Policy) > 0 {
		for _, s := range m.Policy {
			l = len(s)
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	l = len(m.CheckMethod)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.FailClose {
		n += 2
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`Policy:` + fmt.Sprintf("%v", this.Policy) + `,`,
		`CheckMethod:` + fmt.Sprintf("%v", this.CheckMethod) + `,`,
		`FailClose:` + fmt.Sprintf("%v", this.FailClose) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Policy = append(m.Policy, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CheckMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailClose", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FailClose = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthConfig
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)
