// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proofs/proto/proof.proto

package proofspb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MerkleHash struct {
	Left                 []byte   `protobuf:"bytes,1,opt,name=left,proto3" json:"left,omitempty"`
	Right                []byte   `protobuf:"bytes,2,opt,name=right,proto3" json:"right,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerkleHash) Reset()         { *m = MerkleHash{} }
func (m *MerkleHash) String() string { return proto.CompactTextString(m) }
func (*MerkleHash) ProtoMessage()    {}
func (*MerkleHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_proof_627568b69a19d335, []int{0}
}
func (m *MerkleHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerkleHash.Unmarshal(m, b)
}
func (m *MerkleHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerkleHash.Marshal(b, m, deterministic)
}
func (dst *MerkleHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerkleHash.Merge(dst, src)
}
func (m *MerkleHash) XXX_Size() int {
	return xxx_messageInfo_MerkleHash.Size(m)
}
func (m *MerkleHash) XXX_DiscardUnknown() {
	xxx_messageInfo_MerkleHash.DiscardUnknown(m)
}

var xxx_messageInfo_MerkleHash proto.InternalMessageInfo

func (m *MerkleHash) GetLeft() []byte {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *MerkleHash) GetRight() []byte {
	if m != nil {
		return m.Right
	}
	return nil
}

type Proof struct {
	Property string `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Salt     []byte `protobuf:"bytes,3,opt,name=salt,proto3" json:"salt,omitempty"`
	// hash is filled if value & salt are not available
	Hash []byte `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	// Fills either 'hashes' for standard Merkle trees or 'sortedHashes' for a lexicograhical ordered of a node hash
	// not both
	Hashes               []*MerkleHash `protobuf:"bytes,4,rep,name=hashes" json:"hashes,omitempty"`
	SortedHashes         [][]byte      `protobuf:"bytes,5,rep,name=sortedHashes,proto3" json:"sortedHashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_proof_627568b69a19d335, []int{1}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proof.Unmarshal(m, b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
}
func (dst *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(dst, src)
}
func (m *Proof) XXX_Size() int {
	return xxx_messageInfo_Proof.Size(m)
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func (m *Proof) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

func (m *Proof) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Proof) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *Proof) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Proof) GetHashes() []*MerkleHash {
	if m != nil {
		return m.Hashes
	}
	return nil
}

func (m *Proof) GetSortedHashes() [][]byte {
	if m != nil {
		return m.SortedHashes
	}
	return nil
}

var E_ExcludeFromTree = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         2862100,
	Name:          "proofs.exclude_from_tree",
	Tag:           "varint,2862100,opt,name=exclude_from_tree,json=excludeFromTree",
	Filename:      "proofs/proto/proof.proto",
}

var E_HashedField = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         2862101,
	Name:          "proofs.hashed_field",
	Tag:           "varint,2862101,opt,name=hashed_field,json=hashedField",
	Filename:      "proofs/proto/proof.proto",
}

func init() {
	proto.RegisterType((*MerkleHash)(nil), "proofs.MerkleHash")
	proto.RegisterType((*Proof)(nil), "proofs.Proof")
	proto.RegisterExtension(E_ExcludeFromTree)
	proto.RegisterExtension(E_HashedField)
}

func init() { proto.RegisterFile("proofs/proto/proof.proto", fileDescriptor_proof_627568b69a19d335) }

var fileDescriptor_proof_627568b69a19d335 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xc9, 0xbf, 0x6d, 0x68, 0xa7, 0x81, 0x3f, 0x2e, 0x1e, 0x96, 0x82, 0x10, 0x8a, 0x87,
	0xe2, 0x21, 0x05, 0x05, 0x0f, 0x3d, 0x56, 0x28, 0x3d, 0x28, 0x96, 0xe0, 0xc9, 0x4b, 0x49, 0x9b,
	0x49, 0x13, 0xdc, 0x3a, 0xcb, 0xec, 0x56, 0xf4, 0x3d, 0xf4, 0x35, 0x04, 0xdf, 0xc6, 0xc7, 0x91,
	0xdd, 0x6d, 0x15, 0x4f, 0x9e, 0x32, 0xdf, 0xb7, 0xf3, 0x9b, 0x2f, 0x33, 0x20, 0x35, 0x13, 0x55,
	0x66, 0xac, 0x99, 0x2c, 0x8d, 0xbd, 0xc8, 0x7c, 0x2d, 0xe2, 0xf0, 0x32, 0x48, 0x37, 0x44, 0x1b,
	0x85, 0xa1, 0x63, 0xb5, 0xab, 0xc6, 0x25, 0x9a, 0x35, 0x37, 0xda, 0x12, 0x87, 0xce, 0xe1, 0x25,
	0xc0, 0x0d, 0xf2, 0x83, 0xc2, 0x79, 0x61, 0x6a, 0x21, 0xa0, 0xad, 0xb0, 0xb2, 0x32, 0x4a, 0xa3,
	0x51, 0x92, 0xfb, 0x5a, 0x1c, 0x43, 0x87, 0x9b, 0x4d, 0x6d, 0xe5, 0x3f, 0x6f, 0x06, 0x31, 0xfc,
	0x88, 0xa0, 0xb3, 0x70, 0x21, 0x62, 0x00, 0x5d, 0xcd, 0xa4, 0x91, 0xed, 0x8b, 0xe7, 0x7a, 0xf9,
	0xb7, 0x76, 0xec, 0x53, 0xa1, 0x76, 0xe8, 0xd9, 0x5e, 0x1e, 0x84, 0x4b, 0x31, 0x85, 0xb2, 0xb2,
	0x15, 0x52, 0x5c, 0xed, 0xbc, 0xba, 0x30, 0xb5, 0x8c, 0x83, 0xe7, 0x6a, 0x71, 0x06, 0xb1, 0xfb,
	0xa2, 0x91, 0xed, 0xb4, 0x35, 0xea, 0x9f, 0x8b, 0x2c, 0xac, 0x95, 0xfd, 0xfc, 0x71, 0xbe, 0xef,
	0x10, 0x43, 0x48, 0x0c, 0xb1, 0xc5, 0x72, 0x1e, 0x88, 0x4e, 0xda, 0x1a, 0x25, 0xf9, 0x2f, 0x6f,
	0x72, 0x0d, 0x47, 0xf8, 0xbc, 0x56, 0xbb, 0x12, 0x97, 0x15, 0xd3, 0x76, 0x69, 0x19, 0x51, 0x9c,
	0x64, 0xe1, 0x46, 0xd9, 0xe1, 0x46, 0xd9, 0xac, 0x41, 0x55, 0xde, 0x6a, 0xdb, 0xd0, 0xa3, 0x91,
	0xaf, 0x9f, 0xef, 0x6e, 0xa9, 0x6e, 0xfe, 0x7f, 0x8f, 0xce, 0x98, 0xb6, 0x77, 0x8c, 0x38, 0xb9,
	0x82, 0xc4, 0x67, 0x97, 0xcb, 0xca, 0x01, 0x7f, 0x0d, 0x7a, 0x3b, 0x0c, 0xea, 0x07, 0xca, 0x3f,
	0x4e, 0x4f, 0x01, 0xd6, 0xb4, 0xdd, 0xef, 0x35, 0x05, 0x7f, 0xd1, 0x85, 0xe3, 0x17, 0xd1, 0x7d,
	0x37, 0xb8, 0x7a, 0xb5, 0x8a, 0xfd, 0xc8, 0x8b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0x82,
	0xa0, 0xd6, 0xf1, 0x01, 0x00, 0x00,
}