// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sync.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import pb "github.com/BOXFoundation/boxd/core/pb"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type LocateHeaders struct {
	// locate hashes, formally it is as follows:
	// n, n-1, ... n-k, n-k-2, n-k-5, n-k-10, ... n-k-(2^m+m-1), ... genesis
	// n is tail height, k is sequence part length, m is distance factor
	// n-k-(2^m+m-1) is the (k+m)th element
	// to ensure the closer blocks get to genesis, the sparser the locator becomes
	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes" json:"hashes,omitempty"`
}

func (m *LocateHeaders) Reset()         { *m = LocateHeaders{} }
func (m *LocateHeaders) String() string { return proto.CompactTextString(m) }
func (*LocateHeaders) ProtoMessage()    {}
func (*LocateHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_5244f3e500ccc9a7, []int{0}
}
func (m *LocateHeaders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocateHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocateHeaders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *LocateHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocateHeaders.Merge(dst, src)
}
func (m *LocateHeaders) XXX_Size() int {
	return m.Size()
}
func (m *LocateHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_LocateHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_LocateHeaders proto.InternalMessageInfo

func (m *LocateHeaders) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type SyncHeaders struct {
	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes" json:"hashes,omitempty"`
}

func (m *SyncHeaders) Reset()         { *m = SyncHeaders{} }
func (m *SyncHeaders) String() string { return proto.CompactTextString(m) }
func (*SyncHeaders) ProtoMessage()    {}
func (*SyncHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_5244f3e500ccc9a7, []int{1}
}
func (m *SyncHeaders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyncHeaders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SyncHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncHeaders.Merge(dst, src)
}
func (m *SyncHeaders) XXX_Size() int {
	return m.Size()
}
func (m *SyncHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_SyncHeaders proto.InternalMessageInfo

func (m *SyncHeaders) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type CheckHash struct {
	BeginHash []byte `protobuf:"bytes,1,opt,name=begin_hash,json=beginHash,proto3" json:"begin_hash,omitempty"`
	Length    uint32 `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (m *CheckHash) Reset()         { *m = CheckHash{} }
func (m *CheckHash) String() string { return proto.CompactTextString(m) }
func (*CheckHash) ProtoMessage()    {}
func (*CheckHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_5244f3e500ccc9a7, []int{2}
}
func (m *CheckHash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CheckHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CheckHash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CheckHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckHash.Merge(dst, src)
}
func (m *CheckHash) XXX_Size() int {
	return m.Size()
}
func (m *CheckHash) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckHash.DiscardUnknown(m)
}

var xxx_messageInfo_CheckHash proto.InternalMessageInfo

func (m *CheckHash) GetBeginHash() []byte {
	if m != nil {
		return m.BeginHash
	}
	return nil
}

func (m *CheckHash) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type SyncCheckHash struct {
	// it is a root hash for headers between start header and end header
	RootHash []byte `protobuf:"bytes,1,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
}

func (m *SyncCheckHash) Reset()         { *m = SyncCheckHash{} }
func (m *SyncCheckHash) String() string { return proto.CompactTextString(m) }
func (*SyncCheckHash) ProtoMessage()    {}
func (*SyncCheckHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_5244f3e500ccc9a7, []int{3}
}
func (m *SyncCheckHash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncCheckHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyncCheckHash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SyncCheckHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncCheckHash.Merge(dst, src)
}
func (m *SyncCheckHash) XXX_Size() int {
	return m.Size()
}
func (m *SyncCheckHash) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncCheckHash.DiscardUnknown(m)
}

var xxx_messageInfo_SyncCheckHash proto.InternalMessageInfo

func (m *SyncCheckHash) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

type FetchBlockHeaders struct {
	Idx       uint32 `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
	BeginHash []byte `protobuf:"bytes,2,opt,name=begin_hash,json=beginHash,proto3" json:"begin_hash,omitempty"`
	Length    uint32 `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
}

func (m *FetchBlockHeaders) Reset()         { *m = FetchBlockHeaders{} }
func (m *FetchBlockHeaders) String() string { return proto.CompactTextString(m) }
func (*FetchBlockHeaders) ProtoMessage()    {}
func (*FetchBlockHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_5244f3e500ccc9a7, []int{4}
}
func (m *FetchBlockHeaders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchBlockHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchBlockHeaders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *FetchBlockHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchBlockHeaders.Merge(dst, src)
}
func (m *FetchBlockHeaders) XXX_Size() int {
	return m.Size()
}
func (m *FetchBlockHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchBlockHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_FetchBlockHeaders proto.InternalMessageInfo

func (m *FetchBlockHeaders) GetIdx() uint32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

func (m *FetchBlockHeaders) GetBeginHash() []byte {
	if m != nil {
		return m.BeginHash
	}
	return nil
}

func (m *FetchBlockHeaders) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type SyncBlocks struct {
	Idx    uint32      `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
	Blocks []*pb.Block `protobuf:"bytes,2,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *SyncBlocks) Reset()         { *m = SyncBlocks{} }
func (m *SyncBlocks) String() string { return proto.CompactTextString(m) }
func (*SyncBlocks) ProtoMessage()    {}
func (*SyncBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_5244f3e500ccc9a7, []int{5}
}
func (m *SyncBlocks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyncBlocks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SyncBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncBlocks.Merge(dst, src)
}
func (m *SyncBlocks) XXX_Size() int {
	return m.Size()
}
func (m *SyncBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_SyncBlocks proto.InternalMessageInfo

func (m *SyncBlocks) GetIdx() uint32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

func (m *SyncBlocks) GetBlocks() []*pb.Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func init() {
	proto.RegisterType((*LocateHeaders)(nil), "pb.LocateHeaders")
	proto.RegisterType((*SyncHeaders)(nil), "pb.SyncHeaders")
	proto.RegisterType((*CheckHash)(nil), "pb.CheckHash")
	proto.RegisterType((*SyncCheckHash)(nil), "pb.SyncCheckHash")
	proto.RegisterType((*FetchBlockHeaders)(nil), "pb.FetchBlockHeaders")
	proto.RegisterType((*SyncBlocks)(nil), "pb.SyncBlocks")
}
func (m *LocateHeaders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocateHeaders) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for _, b := range m.Hashes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSync(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	return i, nil
}

func (m *SyncHeaders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncHeaders) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for _, b := range m.Hashes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSync(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	return i, nil
}

func (m *CheckHash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckHash) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.BeginHash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSync(dAtA, i, uint64(len(m.BeginHash)))
		i += copy(dAtA[i:], m.BeginHash)
	}
	if m.Length != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSync(dAtA, i, uint64(m.Length))
	}
	return i, nil
}

func (m *SyncCheckHash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncCheckHash) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RootHash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSync(dAtA, i, uint64(len(m.RootHash)))
		i += copy(dAtA[i:], m.RootHash)
	}
	return i, nil
}

func (m *FetchBlockHeaders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchBlockHeaders) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Idx != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSync(dAtA, i, uint64(m.Idx))
	}
	if len(m.BeginHash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSync(dAtA, i, uint64(len(m.BeginHash)))
		i += copy(dAtA[i:], m.BeginHash)
	}
	if m.Length != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSync(dAtA, i, uint64(m.Length))
	}
	return i, nil
}

func (m *SyncBlocks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncBlocks) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Idx != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSync(dAtA, i, uint64(m.Idx))
	}
	if len(m.Blocks) > 0 {
		for _, msg := range m.Blocks {
			dAtA[i] = 0x12
			i++
			i = encodeVarintSync(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintSync(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LocateHeaders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for _, b := range m.Hashes {
			l = len(b)
			n += 1 + l + sovSync(uint64(l))
		}
	}
	return n
}

func (m *SyncHeaders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for _, b := range m.Hashes {
			l = len(b)
			n += 1 + l + sovSync(uint64(l))
		}
	}
	return n
}

func (m *CheckHash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BeginHash)
	if l > 0 {
		n += 1 + l + sovSync(uint64(l))
	}
	if m.Length != 0 {
		n += 1 + sovSync(uint64(m.Length))
	}
	return n
}

func (m *SyncCheckHash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RootHash)
	if l > 0 {
		n += 1 + l + sovSync(uint64(l))
	}
	return n
}

func (m *FetchBlockHeaders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Idx != 0 {
		n += 1 + sovSync(uint64(m.Idx))
	}
	l = len(m.BeginHash)
	if l > 0 {
		n += 1 + l + sovSync(uint64(l))
	}
	if m.Length != 0 {
		n += 1 + sovSync(uint64(m.Length))
	}
	return n
}

func (m *SyncBlocks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Idx != 0 {
		n += 1 + sovSync(uint64(m.Idx))
	}
	if len(m.Blocks) > 0 {
		for _, e := range m.Blocks {
			l = e.Size()
			n += 1 + l + sovSync(uint64(l))
		}
	}
	return n
}

func sovSync(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSync(x uint64) (n int) {
	return sovSync(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LocateHeaders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSync
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LocateHeaders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocateHeaders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSync
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashes = append(m.Hashes, make([]byte, postIndex-iNdEx))
			copy(m.Hashes[len(m.Hashes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSync
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
func (m *SyncHeaders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSync
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SyncHeaders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncHeaders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSync
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashes = append(m.Hashes, make([]byte, postIndex-iNdEx))
			copy(m.Hashes[len(m.Hashes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSync
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
func (m *CheckHash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSync
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CheckHash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckHash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSync
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeginHash = append(m.BeginHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BeginHash == nil {
				m.BeginHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSync
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
func (m *SyncCheckHash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSync
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SyncCheckHash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncCheckHash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSync
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootHash = append(m.RootHash[:0], dAtA[iNdEx:postIndex]...)
			if m.RootHash == nil {
				m.RootHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSync
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
func (m *FetchBlockHeaders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSync
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FetchBlockHeaders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchBlockHeaders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Idx", wireType)
			}
			m.Idx = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Idx |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSync
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeginHash = append(m.BeginHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BeginHash == nil {
				m.BeginHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSync
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
func (m *SyncBlocks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSync
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SyncBlocks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncBlocks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Idx", wireType)
			}
			m.Idx = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Idx |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSync
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocks = append(m.Blocks, &pb.Block{})
			if err := m.Blocks[len(m.Blocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSync
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
func skipSync(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSync
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
					return 0, ErrIntOverflowSync
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
					return 0, ErrIntOverflowSync
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSync
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSync
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
				next, err := skipSync(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthSync = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSync   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sync.proto", fileDescriptor_sync_5244f3e500ccc9a7) }

var fileDescriptor_sync_5244f3e500ccc9a7 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x97, 0x16, 0x8a, 0x7b, 0xb7, 0x82, 0xee, 0x30, 0x8a, 0x62, 0x28, 0x85, 0x61, 0x0f,
	0xd2, 0xa2, 0x7e, 0x83, 0x8a, 0x63, 0x07, 0x41, 0xa8, 0x17, 0x0f, 0x82, 0x34, 0x69, 0x58, 0xca,
	0x66, 0x52, 0x9a, 0x0c, 0xb6, 0x6f, 0xe1, 0xc7, 0xf2, 0xb8, 0xa3, 0x47, 0x59, 0xbf, 0x88, 0x24,
	0x56, 0xfc, 0x83, 0xe2, 0x2d, 0xef, 0xfb, 0xbc, 0xef, 0x93, 0x5f, 0x9e, 0x00, 0xa8, 0x8d, 0xa0,
	0x49, 0xdd, 0x48, 0x2d, 0x47, 0x4e, 0x4d, 0x0e, 0xcf, 0xe6, 0x95, 0xe6, 0x2b, 0x92, 0x50, 0xf9,
	0x98, 0x66, 0x37, 0x77, 0x53, 0xb9, 0x12, 0x65, 0xa1, 0x2b, 0x29, 0x52, 0x22, 0xd7, 0x65, 0x4a,
	0x65, 0xc3, 0xd2, 0x9a, 0xa4, 0x64, 0x29, 0xe9, 0xe2, 0x7d, 0x2d, 0x3a, 0x01, 0xff, 0x5a, 0xd2,
	0x42, 0xb3, 0x19, 0x2b, 0x4a, 0xd6, 0xa8, 0xd1, 0x18, 0x3c, 0x5e, 0x28, 0xce, 0x54, 0x80, 0x42,
	0x37, 0x1e, 0xe6, 0x5d, 0x15, 0x4d, 0x60, 0x70, 0xbb, 0x11, 0xf4, 0xbf, 0xb1, 0x0c, 0xfa, 0x97,
	0x9c, 0xd1, 0xc5, 0xac, 0x50, 0x7c, 0x74, 0x0c, 0x40, 0xd8, 0xbc, 0x12, 0x0f, 0x46, 0x0c, 0x50,
	0x88, 0xe2, 0x61, 0xde, 0xb7, 0x1d, 0x2b, 0x8f, 0xc1, 0x5b, 0x32, 0x31, 0xd7, 0x3c, 0x70, 0x42,
	0x14, 0xfb, 0x79, 0x57, 0x45, 0xa7, 0xe0, 0x9b, 0xab, 0x3e, 0x7d, 0x8e, 0xa0, 0xdf, 0x48, 0xa9,
	0xbf, 0xda, 0xec, 0x99, 0x86, 0x11, 0xa3, 0x7b, 0x38, 0x98, 0x32, 0x4d, 0x79, 0x66, 0x5e, 0xf5,
	0x81, 0xb7, 0x0f, 0x6e, 0x55, 0xae, 0xed, 0xac, 0x9f, 0x9b, 0xe3, 0x0f, 0x16, 0xe7, 0x6f, 0x16,
	0xf7, 0x1b, 0xcb, 0x15, 0x80, 0x61, 0xb1, 0xe6, 0xbf, 0xd9, 0x4e, 0xc0, 0xb3, 0x71, 0xaa, 0xc0,
	0x09, 0xdd, 0x78, 0x70, 0xee, 0x27, 0x26, 0xe5, 0x9a, 0x24, 0x76, 0x23, 0xef, 0xc4, 0x2c, 0x78,
	0xde, 0x61, 0xb4, 0xdd, 0x61, 0xf4, 0xba, 0xc3, 0xe8, 0xa9, 0xc5, 0xbd, 0x6d, 0x8b, 0x7b, 0x2f,
	0x2d, 0xee, 0x11, 0xcf, 0xfe, 0xc3, 0xc5, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0x26, 0x2b,
	0xb0, 0xcc, 0x01, 0x00, 0x00,
}