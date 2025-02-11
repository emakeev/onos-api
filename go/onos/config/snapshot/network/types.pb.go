// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onos/config/snapshot/network/types.proto

package network

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	snapshot "github.com/onosproject/onos-api/go/onos/config/snapshot"
	github_com_onosproject_onos_api_go_onos_config_snapshot_device "github.com/onosproject/onos-api/go/onos/config/snapshot/device"
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

// NetworkSnapshot is a snapshot of all network changes
type NetworkSnapshot struct {
	// 'id' is the unique snapshot identifier
	ID ID `protobuf:"bytes,1,opt,name=id,proto3,casttype=ID" json:"id,omitempty"`
	// 'index' is a monotonically increasing, globally unique snapshot index
	Index Index `protobuf:"varint,2,opt,name=index,proto3,casttype=Index" json:"index,omitempty"`
	// 'revision' is the request revision number
	Revision Revision `protobuf:"varint,3,opt,name=revision,proto3,casttype=Revision" json:"revision,omitempty"`
	// 'status' is the snapshot status
	Status snapshot.Status `protobuf:"bytes,4,opt,name=status,proto3" json:"status"`
	// 'retention' specifies the duration for which to retain changes
	Retention snapshot.RetentionOptions `protobuf:"bytes,6,opt,name=retention,proto3" json:"retention"`
	// 'created' is the time at which the configuration was created
	Created time.Time `protobuf:"bytes,8,opt,name=created,proto3,stdtime" json:"created"`
	// 'updated' is the time at which the configuration was last updated
	Updated time.Time `protobuf:"bytes,9,opt,name=updated,proto3,stdtime" json:"updated"`
	// 'refs' is a set of references to stored device snapshots
	Refs []*DeviceSnapshotRef `protobuf:"bytes,10,rep,name=refs,proto3" json:"refs,omitempty"`
}

func (m *NetworkSnapshot) Reset()         { *m = NetworkSnapshot{} }
func (m *NetworkSnapshot) String() string { return proto.CompactTextString(m) }
func (*NetworkSnapshot) ProtoMessage()    {}
func (*NetworkSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ecd5f2907ec8de9, []int{0}
}
func (m *NetworkSnapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkSnapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkSnapshot.Merge(m, src)
}
func (m *NetworkSnapshot) XXX_Size() int {
	return m.Size()
}
func (m *NetworkSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkSnapshot proto.InternalMessageInfo

func (m *NetworkSnapshot) GetID() ID {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *NetworkSnapshot) GetIndex() Index {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *NetworkSnapshot) GetRevision() Revision {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *NetworkSnapshot) GetStatus() snapshot.Status {
	if m != nil {
		return m.Status
	}
	return snapshot.Status{}
}

func (m *NetworkSnapshot) GetRetention() snapshot.RetentionOptions {
	if m != nil {
		return m.Retention
	}
	return snapshot.RetentionOptions{}
}

func (m *NetworkSnapshot) GetCreated() time.Time {
	if m != nil {
		return m.Created
	}
	return time.Time{}
}

func (m *NetworkSnapshot) GetUpdated() time.Time {
	if m != nil {
		return m.Updated
	}
	return time.Time{}
}

func (m *NetworkSnapshot) GetRefs() []*DeviceSnapshotRef {
	if m != nil {
		return m.Refs
	}
	return nil
}

// DeviceSnapshotRef is a reference to a device snapshot
type DeviceSnapshotRef struct {
	// 'device_snapshot_id' is the unique identifier of the device snapshot
	DeviceSnapshotID github_com_onosproject_onos_api_go_onos_config_snapshot_device.ID `protobuf:"bytes,1,opt,name=device_snapshot_id,json=deviceSnapshotId,proto3,casttype=github.com/onosproject/onos-api/go/onos/config/snapshot/device.ID" json:"device_snapshot_id,omitempty"`
}

func (m *DeviceSnapshotRef) Reset()         { *m = DeviceSnapshotRef{} }
func (m *DeviceSnapshotRef) String() string { return proto.CompactTextString(m) }
func (*DeviceSnapshotRef) ProtoMessage()    {}
func (*DeviceSnapshotRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ecd5f2907ec8de9, []int{1}
}
func (m *DeviceSnapshotRef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceSnapshotRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceSnapshotRef.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceSnapshotRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceSnapshotRef.Merge(m, src)
}
func (m *DeviceSnapshotRef) XXX_Size() int {
	return m.Size()
}
func (m *DeviceSnapshotRef) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceSnapshotRef.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceSnapshotRef proto.InternalMessageInfo

func (m *DeviceSnapshotRef) GetDeviceSnapshotID() github_com_onosproject_onos_api_go_onos_config_snapshot_device.ID {
	if m != nil {
		return m.DeviceSnapshotID
	}
	return ""
}

func init() {
	proto.RegisterType((*NetworkSnapshot)(nil), "onos.config.snapshot.network.NetworkSnapshot")
	proto.RegisterType((*DeviceSnapshotRef)(nil), "onos.config.snapshot.network.DeviceSnapshotRef")
}

func init() {
	proto.RegisterFile("onos/config/snapshot/network/types.proto", fileDescriptor_3ecd5f2907ec8de9)
}

var fileDescriptor_3ecd5f2907ec8de9 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0xb6, 0x57, 0x5a, 0xdf, 0x49, 0x1c, 0x16, 0x43, 0x54, 0x55, 0x49, 0x74, 0x03,
	0xca, 0x82, 0x2d, 0x1d, 0x1b, 0x03, 0x12, 0xa1, 0x4b, 0x19, 0x40, 0xf2, 0xb1, 0x30, 0x9d, 0xd2,
	0xd8, 0xcd, 0x19, 0x68, 0x6c, 0xc5, 0xee, 0x01, 0x9f, 0x80, 0xf5, 0xbe, 0x00, 0x9f, 0x87, 0x1b,
	0x6f, 0x64, 0x0a, 0x28, 0xfd, 0x16, 0x9d, 0x90, 0x9d, 0xf8, 0xe0, 0x44, 0x85, 0xc4, 0x12, 0x3d,
	0xbd, 0xf7, 0xff, 0xbd, 0x7f, 0xf4, 0x7f, 0x86, 0x89, 0x2c, 0xa5, 0x26, 0xb9, 0x2c, 0x57, 0xa2,
	0x20, 0xba, 0xcc, 0x94, 0xbe, 0x90, 0x86, 0x94, 0xdc, 0x7c, 0x94, 0xd5, 0x7b, 0x62, 0x3e, 0x2b,
	0xae, 0xb1, 0xaa, 0xa4, 0x91, 0x68, 0x66, 0x95, 0xb8, 0x55, 0x62, 0xaf, 0xc4, 0x9d, 0x72, 0x1a,
	0x15, 0x52, 0x16, 0x1f, 0x38, 0x71, 0xda, 0xe5, 0x66, 0x45, 0x8c, 0x58, 0x73, 0x6d, 0xb2, 0xb5,
	0x6a, 0xf1, 0xe9, 0xc3, 0x42, 0x16, 0xd2, 0x95, 0xc4, 0x56, 0x5d, 0x37, 0xde, 0x6b, 0xff, 0x87,
	0xed, 0xc9, 0xb7, 0x01, 0xbc, 0xff, 0xaa, 0x35, 0x39, 0xeb, 0xe6, 0x68, 0x06, 0xfb, 0x82, 0x05,
	0x20, 0x06, 0xc9, 0x24, 0x3d, 0x6a, 0xea, 0xa8, 0xbf, 0x98, 0xef, 0xdc, 0x97, 0xf6, 0x05, 0x43,
	0x11, 0x3c, 0x10, 0x25, 0xe3, 0x9f, 0x82, 0x7e, 0x0c, 0x92, 0x61, 0x3a, 0xd9, 0xd5, 0xd1, 0xc1,
	0xc2, 0x36, 0x68, 0xdb, 0x47, 0x09, 0x1c, 0x57, 0xfc, 0x52, 0x68, 0x21, 0xcb, 0x60, 0xe0, 0x34,
	0x47, 0xbb, 0x3a, 0x1a, 0xd3, 0xae, 0x47, 0x6f, 0xa7, 0xe8, 0x29, 0x1c, 0x69, 0x93, 0x99, 0x8d,
	0x0e, 0x86, 0x31, 0x48, 0x0e, 0x4f, 0x67, 0x78, 0x6f, 0x08, 0x67, 0x4e, 0x93, 0x0e, 0xaf, 0xeb,
	0xa8, 0x47, 0x3b, 0x02, 0xbd, 0x84, 0x93, 0x8a, 0x1b, 0x5e, 0x1a, 0x6b, 0x33, 0x72, 0xf8, 0xa3,
	0xfd, 0x38, 0xf5, 0xb2, 0xd7, 0xca, 0x7e, 0xfd, 0xa2, 0xdf, 0x38, 0x7a, 0x06, 0xef, 0xe5, 0x15,
	0xcf, 0x0c, 0x67, 0xc1, 0xd8, 0x6d, 0x9a, 0xe2, 0x36, 0x6f, 0xec, 0xf3, 0xc6, 0x6f, 0x7c, 0xde,
	0xe9, 0xd8, 0xd2, 0x57, 0x3f, 0x22, 0x40, 0x3d, 0x64, 0xf9, 0x8d, 0x62, 0x8e, 0x9f, 0xfc, 0x0f,
	0xdf, 0x41, 0xe8, 0x05, 0x1c, 0x56, 0x7c, 0xa5, 0x03, 0x18, 0x0f, 0x92, 0xc3, 0x53, 0x82, 0xff,
	0xf5, 0x14, 0xf0, 0x9c, 0x5f, 0x8a, 0x9c, 0xfb, 0x63, 0x51, 0xbe, 0xa2, 0x0e, 0x3e, 0xf9, 0x0a,
	0xe0, 0x83, 0xbf, 0x66, 0xe8, 0x0b, 0x80, 0x88, 0xb9, 0xee, 0xb9, 0xdf, 0x74, 0x7e, 0x7b, 0xdc,
	0xb7, 0x4d, 0x1d, 0x1d, 0xdf, 0x65, 0xdc, 0xa9, 0x9f, 0x17, 0xc2, 0x5c, 0x6c, 0x96, 0x38, 0x97,
	0x6b, 0x62, 0xff, 0x45, 0x55, 0xf2, 0x1d, 0xcf, 0x8d, 0xab, 0x1f, 0x67, 0x4a, 0x90, 0x42, 0x92,
	0xbd, 0x2f, 0xab, 0x75, 0xc2, 0x8b, 0x39, 0x3d, 0x66, 0x77, 0xd7, 0xb2, 0x34, 0xb8, 0x6e, 0x42,
	0x70, 0xd3, 0x84, 0xe0, 0x67, 0x13, 0x82, 0xab, 0x6d, 0xd8, 0xbb, 0xd9, 0x86, 0xbd, 0xef, 0xdb,
	0xb0, 0xb7, 0x1c, 0xb9, 0x94, 0x9e, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x48, 0x78, 0x26,
	0x2d, 0x03, 0x00, 0x00,
}

func (m *NetworkSnapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkSnapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkSnapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Refs) > 0 {
		for iNdEx := len(m.Refs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Refs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Updated, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Updated):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTypes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x4a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Created, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Created):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x42
	{
		size, err := m.Retention.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Revision != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Revision))
		i--
		dAtA[i] = 0x18
	}
	if m.Index != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeviceSnapshotRef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceSnapshotRef) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceSnapshotRef) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DeviceSnapshotID) > 0 {
		i -= len(m.DeviceSnapshotID)
		copy(dAtA[i:], m.DeviceSnapshotID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DeviceSnapshotID)))
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
func (m *NetworkSnapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovTypes(uint64(m.Index))
	}
	if m.Revision != 0 {
		n += 1 + sovTypes(uint64(m.Revision))
	}
	l = m.Status.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Retention.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Created)
	n += 1 + l + sovTypes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Updated)
	n += 1 + l + sovTypes(uint64(l))
	if len(m.Refs) > 0 {
		for _, e := range m.Refs {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *DeviceSnapshotRef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DeviceSnapshotID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkSnapshot) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NetworkSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
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
			m.ID = ID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= Index(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= Revision(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Retention", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Retention.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Created, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Updated, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Refs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Refs = append(m.Refs, &DeviceSnapshotRef{})
			if err := m.Refs[len(m.Refs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *DeviceSnapshotRef) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DeviceSnapshotRef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceSnapshotRef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceSnapshotID", wireType)
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
			m.DeviceSnapshotID = github_com_onosproject_onos_api_go_onos_config_snapshot_device.ID(dAtA[iNdEx:postIndex])
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
