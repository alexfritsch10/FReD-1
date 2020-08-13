// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: storage.proto
//lint:file-ignore SA1019 Ignore old grpc implementation, this is a generated file

package storageconnection

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keygroup string `protobuf:"bytes,1,opt,name=keygroup,proto3" json:"keygroup,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Val      string `protobuf:"bytes,3,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetKeygroup() string {
	if x != nil {
		return x.Keygroup
	}
	return ""
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

// A Key uniquely identifies data. In our case it contains a keygroup and the id
type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keygroup string `protobuf:"bytes,1,opt,name=keygroup,proto3" json:"keygroup,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{1}
}

func (x *Key) GetKeygroup() string {
	if x != nil {
		return x.Keygroup
	}
	return ""
}

func (x *Key) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Val struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val string `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Val) Reset() {
	*x = Val{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Val) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Val) ProtoMessage() {}

func (x *Val) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Val.ProtoReflect.Descriptor instead.
func (*Val) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{2}
}

func (x *Val) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type Keygroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keygroup string `protobuf:"bytes,1,opt,name=keygroup,proto3" json:"keygroup,omitempty"`
}

func (x *Keygroup) Reset() {
	*x = Keygroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keygroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keygroup) ProtoMessage() {}

func (x *Keygroup) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keygroup.ProtoReflect.Descriptor instead.
func (*Keygroup) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{3}
}

func (x *Keygroup) GetKeygroup() string {
	if x != nil {
		return x.Keygroup
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_storage_proto protoreflect.FileDescriptor

var file_storage_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x44, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x31, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x6b, 0x65, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6b, 0x65, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x03, 0x56, 0x61, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x22, 0x26, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3e, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x89, 0x02, 0x0a, 0x08, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x1b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x04,
	0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x14, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x04, 0x2e, 0x4b, 0x65, 0x79, 0x1a,
	0x04, 0x2e, 0x56, 0x61, 0x6c, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x41,
	0x6c, 0x6c, 0x12, 0x09, 0x2e, 0x4b, 0x65, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x05, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x30, 0x01, 0x12, 0x1a, 0x0a, 0x03, 0x49, 0x44, 0x73, 0x12,
	0x09, 0x2e, 0x4b, 0x65, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x04, 0x2e, 0x4b, 0x65, 0x79,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x04,
	0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x28, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x09, 0x2e, 0x4b, 0x65, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x09,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x09, 0x2e,
	0x4b, 0x65, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_proto_rawDescOnce sync.Once
	file_storage_proto_rawDescData = file_storage_proto_rawDesc
)

func file_storage_proto_rawDescGZIP() []byte {
	file_storage_proto_rawDescOnce.Do(func() {
		file_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_proto_rawDescData)
	})
	return file_storage_proto_rawDescData
}

var file_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_storage_proto_goTypes = []interface{}{
	(*Item)(nil),     // 0: Item
	(*Key)(nil),      // 1: Key
	(*Val)(nil),      // 2: Val
	(*Keygroup)(nil), // 3: Keygroup
	(*Response)(nil), // 4: Response
}
var file_storage_proto_depIdxs = []int32{
	0, // 0: Database.Update:input_type -> Item
	1, // 1: Database.Delete:input_type -> Key
	1, // 2: Database.Read:input_type -> Key
	3, // 3: Database.ReadAll:input_type -> Keygroup
	3, // 4: Database.IDs:input_type -> Keygroup
	1, // 5: Database.Exists:input_type -> Key
	3, // 6: Database.CreateKeygroup:input_type -> Keygroup
	3, // 7: Database.DeleteKeygroup:input_type -> Keygroup
	4, // 8: Database.Update:output_type -> Response
	4, // 9: Database.Delete:output_type -> Response
	2, // 10: Database.Read:output_type -> Val
	0, // 11: Database.ReadAll:output_type -> Item
	1, // 12: Database.IDs:output_type -> Key
	4, // 13: Database.Exists:output_type -> Response
	4, // 14: Database.CreateKeygroup:output_type -> Response
	4, // 15: Database.DeleteKeygroup:output_type -> Response
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_storage_proto_init() }
func file_storage_proto_init() {
	if File_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Val); i {
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
		file_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keygroup); i {
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
		file_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storage_proto_goTypes,
		DependencyIndexes: file_storage_proto_depIdxs,
		MessageInfos:      file_storage_proto_msgTypes,
	}.Build()
	File_storage_proto = out.File
	file_storage_proto_rawDesc = nil
	file_storage_proto_goTypes = nil
	file_storage_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseClient interface {
	Update(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Response, error)
	Read(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Val, error)
	ReadAll(ctx context.Context, in *Keygroup, opts ...grpc.CallOption) (Database_ReadAllClient, error)
	IDs(ctx context.Context, in *Keygroup, opts ...grpc.CallOption) (Database_IDsClient, error)
	Exists(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Response, error)
	CreateKeygroup(ctx context.Context, in *Keygroup, opts ...grpc.CallOption) (*Response, error)
	DeleteKeygroup(ctx context.Context, in *Keygroup, opts ...grpc.CallOption) (*Response, error)
}

type databaseClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseClient(cc grpc.ClientConnInterface) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) Update(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Database/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Database/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Read(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Val, error) {
	out := new(Val)
	err := c.cc.Invoke(ctx, "/Database/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) ReadAll(ctx context.Context, in *Keygroup, opts ...grpc.CallOption) (Database_ReadAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Database_serviceDesc.Streams[0], "/Database/ReadAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &databaseReadAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Database_ReadAllClient interface {
	Recv() (*Item, error)
	grpc.ClientStream
}

type databaseReadAllClient struct {
	grpc.ClientStream
}

func (x *databaseReadAllClient) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *databaseClient) IDs(ctx context.Context, in *Keygroup, opts ...grpc.CallOption) (Database_IDsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Database_serviceDesc.Streams[1], "/Database/IDs", opts...)
	if err != nil {
		return nil, err
	}
	x := &databaseIDsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Database_IDsClient interface {
	Recv() (*Key, error)
	grpc.ClientStream
}

type databaseIDsClient struct {
	grpc.ClientStream
}

func (x *databaseIDsClient) Recv() (*Key, error) {
	m := new(Key)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *databaseClient) Exists(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Database/Exists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) CreateKeygroup(ctx context.Context, in *Keygroup, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Database/CreateKeygroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) DeleteKeygroup(ctx context.Context, in *Keygroup, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Database/DeleteKeygroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServer is the server API for Database service.
type DatabaseServer interface {
	Update(context.Context, *Item) (*Response, error)
	Delete(context.Context, *Key) (*Response, error)
	Read(context.Context, *Key) (*Val, error)
	ReadAll(*Keygroup, Database_ReadAllServer) error
	IDs(*Keygroup, Database_IDsServer) error
	Exists(context.Context, *Key) (*Response, error)
	CreateKeygroup(context.Context, *Keygroup) (*Response, error)
	DeleteKeygroup(context.Context, *Keygroup) (*Response, error)
}

// UnimplementedDatabaseServer can be embedded to have forward compatible implementations.
type UnimplementedDatabaseServer struct {
}

func (*UnimplementedDatabaseServer) Update(context.Context, *Item) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedDatabaseServer) Delete(context.Context, *Key) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedDatabaseServer) Read(context.Context, *Key) (*Val, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedDatabaseServer) ReadAll(*Keygroup, Database_ReadAllServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}
func (*UnimplementedDatabaseServer) IDs(*Keygroup, Database_IDsServer) error {
	return status.Errorf(codes.Unimplemented, "method IDs not implemented")
}
func (*UnimplementedDatabaseServer) Exists(context.Context, *Key) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exists not implemented")
}
func (*UnimplementedDatabaseServer) CreateKeygroup(context.Context, *Keygroup) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKeygroup not implemented")
}
func (*UnimplementedDatabaseServer) DeleteKeygroup(context.Context, *Keygroup) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKeygroup not implemented")
}

func RegisterDatabaseServer(s *grpc.Server, srv DatabaseServer) {
	s.RegisterService(&_Database_serviceDesc, srv)
}

func _Database_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Database/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Update(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Database/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Delete(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Database/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Read(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_ReadAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Keygroup)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DatabaseServer).ReadAll(m, &databaseReadAllServer{stream})
}

type Database_ReadAllServer interface {
	Send(*Item) error
	grpc.ServerStream
}

type databaseReadAllServer struct {
	grpc.ServerStream
}

func (x *databaseReadAllServer) Send(m *Item) error {
	return x.ServerStream.SendMsg(m)
}

func _Database_IDs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Keygroup)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DatabaseServer).IDs(m, &databaseIDsServer{stream})
}

type Database_IDsServer interface {
	Send(*Key) error
	grpc.ServerStream
}

type databaseIDsServer struct {
	grpc.ServerStream
}

func (x *databaseIDsServer) Send(m *Key) error {
	return x.ServerStream.SendMsg(m)
}

func _Database_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Exists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Database/Exists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Exists(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_CreateKeygroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Keygroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).CreateKeygroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Database/CreateKeygroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).CreateKeygroup(ctx, req.(*Keygroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_DeleteKeygroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Keygroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).DeleteKeygroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Database/DeleteKeygroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).DeleteKeygroup(ctx, req.(*Keygroup))
	}
	return interceptor(ctx, in, info, handler)
}

var _Database_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _Database_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Database_Delete_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Database_Read_Handler,
		},
		{
			MethodName: "Exists",
			Handler:    _Database_Exists_Handler,
		},
		{
			MethodName: "CreateKeygroup",
			Handler:    _Database_CreateKeygroup_Handler,
		},
		{
			MethodName: "DeleteKeygroup",
			Handler:    _Database_DeleteKeygroup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadAll",
			Handler:       _Database_ReadAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "IDs",
			Handler:       _Database_IDs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "storage.proto",
}
