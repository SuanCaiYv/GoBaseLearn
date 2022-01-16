// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: msg_template.proto

package rpc

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

type Msg1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Length    int64            `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	Content   []byte           `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Map1      map[string]int32 `protobuf:"bytes,5,rep,name=map1,proto3" json:"map1,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Kinds     []string         `protobuf:"bytes,6,rep,name=kinds,proto3" json:"kinds,omitempty"`
	Available bool             `protobuf:"varint,7,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *Msg1) Reset() {
	*x = Msg1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg1) ProtoMessage() {}

func (x *Msg1) ProtoReflect() protoreflect.Message {
	mi := &file_msg_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg1.ProtoReflect.Descriptor instead.
func (*Msg1) Descriptor() ([]byte, []int) {
	return file_msg_template_proto_rawDescGZIP(), []int{0}
}

func (x *Msg1) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Msg1) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Msg1) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Msg1) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Msg1) GetMap1() map[string]int32 {
	if x != nil {
		return x.Map1
	}
	return nil
}

func (x *Msg1) GetKinds() []string {
	if x != nil {
		return x.Kinds
	}
	return nil
}

func (x *Msg1) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

type Msg2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg1 *Msg1 `protobuf:"bytes,1,opt,name=msg1,proto3" json:"msg1,omitempty"`
	// Types that are assignable to Type:
	//	*Msg2_A
	//	*Msg2_B
	Type isMsg2_Type `protobuf_oneof:"type"`
}

func (x *Msg2) Reset() {
	*x = Msg2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg2) ProtoMessage() {}

func (x *Msg2) ProtoReflect() protoreflect.Message {
	mi := &file_msg_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg2.ProtoReflect.Descriptor instead.
func (*Msg2) Descriptor() ([]byte, []int) {
	return file_msg_template_proto_rawDescGZIP(), []int{1}
}

func (x *Msg2) GetMsg1() *Msg1 {
	if x != nil {
		return x.Msg1
	}
	return nil
}

func (m *Msg2) GetType() isMsg2_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Msg2) GetA() string {
	if x, ok := x.GetType().(*Msg2_A); ok {
		return x.A
	}
	return ""
}

func (x *Msg2) GetB() string {
	if x, ok := x.GetType().(*Msg2_B); ok {
		return x.B
	}
	return ""
}

type isMsg2_Type interface {
	isMsg2_Type()
}

type Msg2_A struct {
	A string `protobuf:"bytes,2,opt,name=a,proto3,oneof"`
}

type Msg2_B struct {
	B string `protobuf:"bytes,3,opt,name=b,proto3,oneof"`
}

func (*Msg2_A) isMsg2_Type() {}

func (*Msg2_B) isMsg2_Type() {}

var File_msg_template_proto protoreflect.FileDescriptor

var file_msg_template_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x04, 0x4d, 0x73, 0x67, 0x31, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x6d, 0x61, 0x70, 0x31, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x73, 0x67, 0x31, 0x2e, 0x4d, 0x61, 0x70, 0x31, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x6d, 0x61, 0x70, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x69, 0x6e, 0x64,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x69, 0x6e, 0x64, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x37, 0x0a, 0x09,
	0x4d, 0x61, 0x70, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x49, 0x0a, 0x04, 0x4d, 0x73, 0x67, 0x32, 0x12, 0x19, 0x0a,
	0x04, 0x6d, 0x73, 0x67, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4d, 0x73,
	0x67, 0x31, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x31, 0x12, 0x0e, 0x0a, 0x01, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x01, 0x61, 0x12, 0x0e, 0x0a, 0x01, 0x62, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x01, 0x62, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_template_proto_rawDescOnce sync.Once
	file_msg_template_proto_rawDescData = file_msg_template_proto_rawDesc
)

func file_msg_template_proto_rawDescGZIP() []byte {
	file_msg_template_proto_rawDescOnce.Do(func() {
		file_msg_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_template_proto_rawDescData)
	})
	return file_msg_template_proto_rawDescData
}

var file_msg_template_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_msg_template_proto_goTypes = []interface{}{
	(*Msg1)(nil), // 0: Msg1
	(*Msg2)(nil), // 1: Msg2
	nil,          // 2: Msg1.Map1Entry
}
var file_msg_template_proto_depIdxs = []int32{
	2, // 0: Msg1.map1:type_name -> Msg1.Map1Entry
	0, // 1: Msg2.msg1:type_name -> Msg1
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_msg_template_proto_init() }
func file_msg_template_proto_init() {
	if File_msg_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg1); i {
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
		file_msg_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg2); i {
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
	file_msg_template_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Msg2_A)(nil),
		(*Msg2_B)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_template_proto_goTypes,
		DependencyIndexes: file_msg_template_proto_depIdxs,
		MessageInfos:      file_msg_template_proto_msgTypes,
	}.Build()
	File_msg_template_proto = out.File
	file_msg_template_proto_rawDesc = nil
	file_msg_template_proto_goTypes = nil
	file_msg_template_proto_depIdxs = nil
}
