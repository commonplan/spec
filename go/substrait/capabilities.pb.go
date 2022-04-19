// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: substrait/capabilities.proto

package substrait

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

// Defines a set of Capabilities that a system (producer or consumer) supports.
type Capabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of Substrait versions this system supports
	SubstraitVersions []string `protobuf:"bytes,1,rep,name=substrait_versions,json=substraitVersions,proto3" json:"substrait_versions,omitempty"`
	// list of com.google.Any message types this system supports for advanced
	// extensions.
	AdvancedExtensionTypeUrls []string `protobuf:"bytes,2,rep,name=advanced_extension_type_urls,json=advancedExtensionTypeUrls,proto3" json:"advanced_extension_type_urls,omitempty"`
	// list of simple extensions this system supports.
	SimpleExtensions []*Capabilities_SimpleExtension `protobuf:"bytes,3,rep,name=simple_extensions,json=simpleExtensions,proto3" json:"simple_extensions,omitempty"`
}

func (x *Capabilities) Reset() {
	*x = Capabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_capabilities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capabilities) ProtoMessage() {}

func (x *Capabilities) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_capabilities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capabilities.ProtoReflect.Descriptor instead.
func (*Capabilities) Descriptor() ([]byte, []int) {
	return file_substrait_capabilities_proto_rawDescGZIP(), []int{0}
}

func (x *Capabilities) GetSubstraitVersions() []string {
	if x != nil {
		return x.SubstraitVersions
	}
	return nil
}

func (x *Capabilities) GetAdvancedExtensionTypeUrls() []string {
	if x != nil {
		return x.AdvancedExtensionTypeUrls
	}
	return nil
}

func (x *Capabilities) GetSimpleExtensions() []*Capabilities_SimpleExtension {
	if x != nil {
		return x.SimpleExtensions
	}
	return nil
}

type Capabilities_SimpleExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri               string   `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	FunctionKeys      []string `protobuf:"bytes,2,rep,name=function_keys,json=functionKeys,proto3" json:"function_keys,omitempty"`
	TypeKeys          []string `protobuf:"bytes,3,rep,name=type_keys,json=typeKeys,proto3" json:"type_keys,omitempty"`
	TypeVariationKeys []string `protobuf:"bytes,4,rep,name=type_variation_keys,json=typeVariationKeys,proto3" json:"type_variation_keys,omitempty"`
}

func (x *Capabilities_SimpleExtension) Reset() {
	*x = Capabilities_SimpleExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_capabilities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capabilities_SimpleExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capabilities_SimpleExtension) ProtoMessage() {}

func (x *Capabilities_SimpleExtension) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_capabilities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capabilities_SimpleExtension.ProtoReflect.Descriptor instead.
func (*Capabilities_SimpleExtension) Descriptor() ([]byte, []int) {
	return file_substrait_capabilities_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Capabilities_SimpleExtension) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Capabilities_SimpleExtension) GetFunctionKeys() []string {
	if x != nil {
		return x.FunctionKeys
	}
	return nil
}

func (x *Capabilities_SimpleExtension) GetTypeKeys() []string {
	if x != nil {
		return x.TypeKeys
	}
	return nil
}

func (x *Capabilities_SimpleExtension) GetTypeVariationKeys() []string {
	if x != nil {
		return x.TypeVariationKeys
	}
	return nil
}

var File_substrait_capabilities_proto protoreflect.FileDescriptor

var file_substrait_capabilities_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2f, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x22, 0xec, 0x02, 0x0a, 0x0c, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x75,
	0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x1c, 0x61, 0x64, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x19, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x54, 0x0a, 0x11, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x10,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x95, 0x01, 0x0a, 0x0f, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x79, 0x70, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x73, 0x42, 0x5b, 0x0a, 0x12, 0x69, 0x6f, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x62,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74,
	0xaa, 0x02, 0x12, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_substrait_capabilities_proto_rawDescOnce sync.Once
	file_substrait_capabilities_proto_rawDescData = file_substrait_capabilities_proto_rawDesc
)

func file_substrait_capabilities_proto_rawDescGZIP() []byte {
	file_substrait_capabilities_proto_rawDescOnce.Do(func() {
		file_substrait_capabilities_proto_rawDescData = protoimpl.X.CompressGZIP(file_substrait_capabilities_proto_rawDescData)
	})
	return file_substrait_capabilities_proto_rawDescData
}

var file_substrait_capabilities_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_substrait_capabilities_proto_goTypes = []interface{}{
	(*Capabilities)(nil),                 // 0: substrait.Capabilities
	(*Capabilities_SimpleExtension)(nil), // 1: substrait.Capabilities.SimpleExtension
}
var file_substrait_capabilities_proto_depIdxs = []int32{
	1, // 0: substrait.Capabilities.simple_extensions:type_name -> substrait.Capabilities.SimpleExtension
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_substrait_capabilities_proto_init() }
func file_substrait_capabilities_proto_init() {
	if File_substrait_capabilities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_substrait_capabilities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capabilities); i {
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
		file_substrait_capabilities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capabilities_SimpleExtension); i {
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
			RawDescriptor: file_substrait_capabilities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_substrait_capabilities_proto_goTypes,
		DependencyIndexes: file_substrait_capabilities_proto_depIdxs,
		MessageInfos:      file_substrait_capabilities_proto_msgTypes,
	}.Build()
	File_substrait_capabilities_proto = out.File
	file_substrait_capabilities_proto_rawDesc = nil
	file_substrait_capabilities_proto_goTypes = nil
	file_substrait_capabilities_proto_depIdxs = nil
}
