// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: cdr.proto

package cdr

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Uses numeric values from this table
// https://github.com/multiformats/multicodec/blob/master/table.csv
type HashAlgo int32

const (
	HashAlgo_UNKNOWN_HASH HashAlgo = 0
	HashAlgo_BLAKE2b_256  HashAlgo = 45600
)

// Enum value maps for HashAlgo.
var (
	HashAlgo_name = map[int32]string{
		0:     "UNKNOWN_HASH",
		45600: "BLAKE2b_256",
	}
	HashAlgo_value = map[string]int32{
		"UNKNOWN_HASH": 0,
		"BLAKE2b_256":  45600,
	}
)

func (x HashAlgo) Enum() *HashAlgo {
	p := new(HashAlgo)
	*p = x
	return p
}

func (x HashAlgo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HashAlgo) Descriptor() protoreflect.EnumDescriptor {
	return file_cdr_proto_enumTypes[0].Descriptor()
}

func (HashAlgo) Type() protoreflect.EnumType {
	return &file_cdr_proto_enumTypes[0]
}

func (x HashAlgo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HashAlgo.Descriptor instead.
func (HashAlgo) EnumDescriptor() ([]byte, []int) {
	return file_cdr_proto_rawDescGZIP(), []int{0}
}

type CipherAlgo int32

const (
	CipherAlgo_UNKNOWN_CIPHER CipherAlgo = 0
	CipherAlgo_CHACHA20       CipherAlgo = 1
)

// Enum value maps for CipherAlgo.
var (
	CipherAlgo_name = map[int32]string{
		0: "UNKNOWN_CIPHER",
		1: "CHACHA20",
	}
	CipherAlgo_value = map[string]int32{
		"UNKNOWN_CIPHER": 0,
		"CHACHA20":       1,
	}
)

func (x CipherAlgo) Enum() *CipherAlgo {
	p := new(CipherAlgo)
	*p = x
	return p
}

func (x CipherAlgo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CipherAlgo) Descriptor() protoreflect.EnumDescriptor {
	return file_cdr_proto_enumTypes[1].Descriptor()
}

func (CipherAlgo) Type() protoreflect.EnumType {
	return &file_cdr_proto_enumTypes[1]
}

func (x CipherAlgo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CipherAlgo.Descriptor instead.
func (CipherAlgo) EnumDescriptor() ([]byte, []int) {
	return file_cdr_proto_rawDescGZIP(), []int{1}
}

type CompressAlgo int32

const (
	CompressAlgo_UNKNOWN_COMPRESS CompressAlgo = 0
	CompressAlgo_GZIP             CompressAlgo = 1
)

// Enum value maps for CompressAlgo.
var (
	CompressAlgo_name = map[int32]string{
		0: "UNKNOWN_COMPRESS",
		1: "GZIP",
	}
	CompressAlgo_value = map[string]int32{
		"UNKNOWN_COMPRESS": 0,
		"GZIP":             1,
	}
)

func (x CompressAlgo) Enum() *CompressAlgo {
	p := new(CompressAlgo)
	*p = x
	return p
}

func (x CompressAlgo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompressAlgo) Descriptor() protoreflect.EnumDescriptor {
	return file_cdr_proto_enumTypes[2].Descriptor()
}

func (CompressAlgo) Type() protoreflect.EnumType {
	return &file_cdr_proto_enumTypes[2]
}

func (x CompressAlgo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompressAlgo.Descriptor instead.
func (CompressAlgo) EnumDescriptor() ([]byte, []int) {
	return file_cdr_proto_rawDescGZIP(), []int{2}
}

type Ref struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//	*Ref_Http
	//	*Ref_ContentHash
	//	*Ref_SizeLimits
	//	*Ref_Cipher
	//	*Ref_Compress
	//	*Ref_Slice
	//	*Ref_Concat
	Body isRef_Body `protobuf_oneof:"body"`
}

func (x *Ref) Reset() {
	*x = Ref{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cdr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ref) ProtoMessage() {}

func (x *Ref) ProtoReflect() protoreflect.Message {
	mi := &file_cdr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ref.ProtoReflect.Descriptor instead.
func (*Ref) Descriptor() ([]byte, []int) {
	return file_cdr_proto_rawDescGZIP(), []int{0}
}

func (m *Ref) GetBody() isRef_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Ref) GetHttp() *HTTP {
	if x, ok := x.GetBody().(*Ref_Http); ok {
		return x.Http
	}
	return nil
}

func (x *Ref) GetContentHash() *ContentHash {
	if x, ok := x.GetBody().(*Ref_ContentHash); ok {
		return x.ContentHash
	}
	return nil
}

func (x *Ref) GetSizeLimits() *SizeLimits {
	if x, ok := x.GetBody().(*Ref_SizeLimits); ok {
		return x.SizeLimits
	}
	return nil
}

func (x *Ref) GetCipher() *Cipher {
	if x, ok := x.GetBody().(*Ref_Cipher); ok {
		return x.Cipher
	}
	return nil
}

func (x *Ref) GetCompress() *Compress {
	if x, ok := x.GetBody().(*Ref_Compress); ok {
		return x.Compress
	}
	return nil
}

func (x *Ref) GetSlice() *Slice {
	if x, ok := x.GetBody().(*Ref_Slice); ok {
		return x.Slice
	}
	return nil
}

func (x *Ref) GetConcat() *Concat {
	if x, ok := x.GetBody().(*Ref_Concat); ok {
		return x.Concat
	}
	return nil
}

type isRef_Body interface {
	isRef_Body()
}

type Ref_Http struct {
	// Sources
	Http *HTTP `protobuf:"bytes,1,opt,name=http,proto3,oneof"`
}

type Ref_ContentHash struct {
	// Constraints
	ContentHash *ContentHash `protobuf:"bytes,32,opt,name=content_hash,json=contentHash,proto3,oneof"`
}

type Ref_SizeLimits struct {
	SizeLimits *SizeLimits `protobuf:"bytes,33,opt,name=size_limits,json=sizeLimits,proto3,oneof"`
}

type Ref_Cipher struct {
	// 1:1 Transforms
	Cipher *Cipher `protobuf:"bytes,64,opt,name=cipher,proto3,oneof"`
}

type Ref_Compress struct {
	Compress *Compress `protobuf:"bytes,65,opt,name=compress,proto3,oneof"`
}

type Ref_Slice struct {
	Slice *Slice `protobuf:"bytes,66,opt,name=slice,proto3,oneof"`
}

type Ref_Concat struct {
	// Many:1 Transforms
	Concat *Concat `protobuf:"bytes,96,opt,name=concat,proto3,oneof"`
}

func (*Ref_Http) isRef_Body() {}

func (*Ref_ContentHash) isRef_Body() {}

func (*Ref_SizeLimits) isRef_Body() {}

func (*Ref_Cipher) isRef_Body() {}

func (*Ref_Compress) isRef_Body() {}

func (*Ref_Slice) isRef_Body() {}

func (*Ref_Concat) isRef_Body() {}

type HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string            `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HTTP) Reset() {
	*x = HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cdr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTP) ProtoMessage() {}

func (x *HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_cdr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTP.ProtoReflect.Descriptor instead.
func (*HTTP) Descriptor() ([]byte, []int) {
	return file_cdr_proto_rawDescGZIP(), []int{1}
}

func (x *HTTP) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HTTP) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

// Contraints
type ContentHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner *Ref     `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
	Algo  HashAlgo `protobuf:"varint,2,opt,name=algo,proto3,enum=pachyderm.cdr.HashAlgo" json:"algo,omitempty"`
	Hash  []byte   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *ContentHash) Reset() {
	*x = ContentHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cdr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentHash) ProtoMessage() {}

func (x *ContentHash) ProtoReflect() protoreflect.Message {
	mi := &file_cdr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentHash.ProtoReflect.Descriptor instead.
func (*ContentHash) Descriptor() ([]byte, []int) {
	return file_cdr_proto_rawDescGZIP(), []int{2}
}

func (x *ContentHash) GetInner() *Ref {
	if x != nil {
		return x.Inner
	}
	return nil
}

func (x *ContentHash) GetAlgo() HashAlgo {
	if x != nil {
		return x.Algo
	}
	return HashAlgo_UNKNOWN_HASH
}

func (x *ContentHash) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type SizeLimits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner *Ref  `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
	Min   int64 `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
	Max   int64 `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *SizeLimits) Reset() {
	*x = SizeLimits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cdr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SizeLimits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SizeLimits) ProtoMessage() {}

func (x *SizeLimits) ProtoReflect() protoreflect.Message {
	mi := &file_cdr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SizeLimits.ProtoReflect.Descriptor instead.
func (*SizeLimits) Descriptor() ([]byte, []int) {
	return file_cdr_proto_rawDescGZIP(), []int{3}
}

func (x *SizeLimits) GetInner() *Ref {
	if x != nil {
		return x.Inner
	}
	return nil
}

func (x *SizeLimits) GetMin() int64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *SizeLimits) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

// 1:1 Transforms
type Cipher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner *Ref       `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
	Algo  CipherAlgo `protobuf:"varint,2,opt,name=algo,proto3,enum=pachyderm.cdr.CipherAlgo" json:"algo,omitempty"`
	Key   []byte     `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Nonce []byte     `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *Cipher) Reset() {
	*x = Cipher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cdr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cipher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cipher) ProtoMessage() {}

func (x *Cipher) ProtoReflect() protoreflect.Message {
	mi := &file_cdr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cipher.ProtoReflect.Descriptor instead.
func (*Cipher) Descriptor() ([]byte, []int) {
	return file_cdr_proto_rawDescGZIP(), []int{4}
}

func (x *Cipher) GetInner() *Ref {
	if x != nil {
		return x.Inner
	}
	return nil
}

func (x *Cipher) GetAlgo() CipherAlgo {
	if x != nil {
		return x.Algo
	}
	return CipherAlgo_UNKNOWN_CIPHER
}

func (x *Cipher) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Cipher) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

type Compress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner *Ref         `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
	Algo  CompressAlgo `protobuf:"varint,2,opt,name=algo,proto3,enum=pachyderm.cdr.CompressAlgo" json:"algo,omitempty"`
}

func (x *Compress) Reset() {
	*x = Compress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cdr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Compress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Compress) ProtoMessage() {}

func (x *Compress) ProtoReflect() protoreflect.Message {
	mi := &file_cdr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Compress.ProtoReflect.Descriptor instead.
func (*Compress) Descriptor() ([]byte, []int) {
	return file_cdr_proto_rawDescGZIP(), []int{5}
}

func (x *Compress) GetInner() *Ref {
	if x != nil {
		return x.Inner
	}
	return nil
}

func (x *Compress) GetAlgo() CompressAlgo {
	if x != nil {
		return x.Algo
	}
	return CompressAlgo_UNKNOWN_COMPRESS
}

// 1:1 Transforms
type Slice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner *Ref   `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
	Start uint64 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End   uint64 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Slice) Reset() {
	*x = Slice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cdr_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slice) ProtoMessage() {}

func (x *Slice) ProtoReflect() protoreflect.Message {
	mi := &file_cdr_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slice.ProtoReflect.Descriptor instead.
func (*Slice) Descriptor() ([]byte, []int) {
	return file_cdr_proto_rawDescGZIP(), []int{6}
}

func (x *Slice) GetInner() *Ref {
	if x != nil {
		return x.Inner
	}
	return nil
}

func (x *Slice) GetStart() uint64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Slice) GetEnd() uint64 {
	if x != nil {
		return x.End
	}
	return 0
}

// Many:1 Transforms
type Concat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Refs []*Ref `protobuf:"bytes,1,rep,name=refs,proto3" json:"refs,omitempty"`
}

func (x *Concat) Reset() {
	*x = Concat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cdr_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Concat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Concat) ProtoMessage() {}

func (x *Concat) ProtoReflect() protoreflect.Message {
	mi := &file_cdr_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Concat.ProtoReflect.Descriptor instead.
func (*Concat) Descriptor() ([]byte, []int) {
	return file_cdr_proto_rawDescGZIP(), []int{7}
}

func (x *Concat) GetRefs() []*Ref {
	if x != nil {
		return x.Refs
	}
	return nil
}

var File_cdr_proto protoreflect.FileDescriptor

var file_cdr_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x64, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x61, 0x63,
	0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e, 0x63, 0x64, 0x72, 0x22, 0xfe, 0x02, 0x0a, 0x03, 0x52,
	0x65, 0x66, 0x12, 0x29, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e, 0x63, 0x64, 0x72,
	0x2e, 0x48, 0x54, 0x54, 0x50, 0x48, 0x00, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x3f, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e,
	0x63, 0x64, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x48,
	0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x3c,
	0x0a, 0x0b, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x21, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e,
	0x63, 0x64, 0x72, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x48, 0x00,
	0x52, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x06,
	0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x18, 0x40, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e, 0x63, 0x64, 0x72, 0x2e, 0x43, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x12, 0x35, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x18, 0x41, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e, 0x63, 0x64, 0x72, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x42, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e,
	0x63, 0x64, 0x72, 0x2e, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x48, 0x00, 0x52, 0x05, 0x73, 0x6c, 0x69,
	0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x18, 0x60, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e, 0x63,
	0x64, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x63, 0x61, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x90, 0x01, 0x0a, 0x04,
	0x48, 0x54, 0x54, 0x50, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x3a, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64,
	0x65, 0x72, 0x6d, 0x2e, 0x63, 0x64, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x78,
	0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x28, 0x0a,
	0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e, 0x63, 0x64, 0x72, 0x2e, 0x52, 0x65, 0x66,
	0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72,
	0x6d, 0x2e, 0x63, 0x64, 0x72, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x52, 0x04,
	0x61, 0x6c, 0x67, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x5a, 0x0a, 0x0a, 0x53, 0x69, 0x7a, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72,
	0x6d, 0x2e, 0x63, 0x64, 0x72, 0x2e, 0x52, 0x65, 0x66, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d,
	0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6d, 0x61, 0x78, 0x22, 0x89, 0x01, 0x0a, 0x06, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x12,
	0x28, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e, 0x63, 0x64, 0x72, 0x2e, 0x52,
	0x65, 0x66, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x61, 0x6c, 0x67,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64,
	0x65, 0x72, 0x6d, 0x2e, 0x63, 0x64, 0x72, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x41, 0x6c,
	0x67, 0x6f, 0x52, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x22, 0x65, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x05,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61,
	0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e, 0x63, 0x64, 0x72, 0x2e, 0x52, 0x65, 0x66, 0x52,
	0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d,
	0x2e, 0x63, 0x64, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x41, 0x6c, 0x67,
	0x6f, 0x52, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x22, 0x59, 0x0a, 0x05, 0x53, 0x6c, 0x69, 0x63, 0x65,
	0x12, 0x28, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e, 0x63, 0x64, 0x72, 0x2e,
	0x52, 0x65, 0x66, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x22, 0x30, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x12, 0x26, 0x0a, 0x04,
	0x72, 0x65, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x63,
	0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e, 0x63, 0x64, 0x72, 0x2e, 0x52, 0x65, 0x66, 0x52, 0x04,
	0x72, 0x65, 0x66, 0x73, 0x2a, 0x2f, 0x0a, 0x08, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f,
	0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x48, 0x41, 0x53, 0x48,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0b, 0x42, 0x4c, 0x41, 0x4b, 0x45, 0x32, 0x62, 0x5f, 0x32, 0x35,
	0x36, 0x10, 0xa0, 0xe4, 0x02, 0x2a, 0x2e, 0x0a, 0x0a, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x41,
	0x6c, 0x67, 0x6f, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x43,
	0x49, 0x50, 0x48, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x48, 0x41, 0x43, 0x48,
	0x41, 0x32, 0x30, 0x10, 0x01, 0x2a, 0x2e, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x47,
	0x5a, 0x49, 0x50, 0x10, 0x01, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2f, 0x63, 0x64,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cdr_proto_rawDescOnce sync.Once
	file_cdr_proto_rawDescData = file_cdr_proto_rawDesc
)

func file_cdr_proto_rawDescGZIP() []byte {
	file_cdr_proto_rawDescOnce.Do(func() {
		file_cdr_proto_rawDescData = protoimpl.X.CompressGZIP(file_cdr_proto_rawDescData)
	})
	return file_cdr_proto_rawDescData
}

var file_cdr_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_cdr_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cdr_proto_goTypes = []interface{}{
	(HashAlgo)(0),       // 0: pachyderm.cdr.HashAlgo
	(CipherAlgo)(0),     // 1: pachyderm.cdr.CipherAlgo
	(CompressAlgo)(0),   // 2: pachyderm.cdr.CompressAlgo
	(*Ref)(nil),         // 3: pachyderm.cdr.Ref
	(*HTTP)(nil),        // 4: pachyderm.cdr.HTTP
	(*ContentHash)(nil), // 5: pachyderm.cdr.ContentHash
	(*SizeLimits)(nil),  // 6: pachyderm.cdr.SizeLimits
	(*Cipher)(nil),      // 7: pachyderm.cdr.Cipher
	(*Compress)(nil),    // 8: pachyderm.cdr.Compress
	(*Slice)(nil),       // 9: pachyderm.cdr.Slice
	(*Concat)(nil),      // 10: pachyderm.cdr.Concat
	nil,                 // 11: pachyderm.cdr.HTTP.HeadersEntry
}
var file_cdr_proto_depIdxs = []int32{
	4,  // 0: pachyderm.cdr.Ref.http:type_name -> pachyderm.cdr.HTTP
	5,  // 1: pachyderm.cdr.Ref.content_hash:type_name -> pachyderm.cdr.ContentHash
	6,  // 2: pachyderm.cdr.Ref.size_limits:type_name -> pachyderm.cdr.SizeLimits
	7,  // 3: pachyderm.cdr.Ref.cipher:type_name -> pachyderm.cdr.Cipher
	8,  // 4: pachyderm.cdr.Ref.compress:type_name -> pachyderm.cdr.Compress
	9,  // 5: pachyderm.cdr.Ref.slice:type_name -> pachyderm.cdr.Slice
	10, // 6: pachyderm.cdr.Ref.concat:type_name -> pachyderm.cdr.Concat
	11, // 7: pachyderm.cdr.HTTP.headers:type_name -> pachyderm.cdr.HTTP.HeadersEntry
	3,  // 8: pachyderm.cdr.ContentHash.inner:type_name -> pachyderm.cdr.Ref
	0,  // 9: pachyderm.cdr.ContentHash.algo:type_name -> pachyderm.cdr.HashAlgo
	3,  // 10: pachyderm.cdr.SizeLimits.inner:type_name -> pachyderm.cdr.Ref
	3,  // 11: pachyderm.cdr.Cipher.inner:type_name -> pachyderm.cdr.Ref
	1,  // 12: pachyderm.cdr.Cipher.algo:type_name -> pachyderm.cdr.CipherAlgo
	3,  // 13: pachyderm.cdr.Compress.inner:type_name -> pachyderm.cdr.Ref
	2,  // 14: pachyderm.cdr.Compress.algo:type_name -> pachyderm.cdr.CompressAlgo
	3,  // 15: pachyderm.cdr.Slice.inner:type_name -> pachyderm.cdr.Ref
	3,  // 16: pachyderm.cdr.Concat.refs:type_name -> pachyderm.cdr.Ref
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_cdr_proto_init() }
func file_cdr_proto_init() {
	if File_cdr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cdr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ref); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cdr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTP); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cdr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentHash); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cdr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SizeLimits); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cdr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cipher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cdr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Compress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cdr_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cdr_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Concat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_cdr_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Ref_Http)(nil),
		(*Ref_ContentHash)(nil),
		(*Ref_SizeLimits)(nil),
		(*Ref_Cipher)(nil),
		(*Ref_Compress)(nil),
		(*Ref_Slice)(nil),
		(*Ref_Concat)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cdr_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cdr_proto_goTypes,
		DependencyIndexes: file_cdr_proto_depIdxs,
		EnumInfos:         file_cdr_proto_enumTypes,
		MessageInfos:      file_cdr_proto_msgTypes,
	}.Build()
	File_cdr_proto = out.File
	file_cdr_proto_rawDesc = nil
	file_cdr_proto_goTypes = nil
	file_cdr_proto_depIdxs = nil
}
