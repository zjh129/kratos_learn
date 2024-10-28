// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.4
// source: conf/conf_driver.proto

package conf

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

type ConfDriver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string      `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Apollo *ConfApollo `protobuf:"bytes,2,opt,name=apollo,proto3" json:"apollo,omitempty"`
	Consul *ConfConsul `protobuf:"bytes,3,opt,name=consul,proto3" json:"consul,omitempty"`
	Nacos  *ConfNacos  `protobuf:"bytes,4,opt,name=nacos,proto3" json:"nacos,omitempty"`
}

func (x *ConfDriver) Reset() {
	*x = ConfDriver{}
	mi := &file_conf_conf_driver_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfDriver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfDriver) ProtoMessage() {}

func (x *ConfDriver) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_driver_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfDriver.ProtoReflect.Descriptor instead.
func (*ConfDriver) Descriptor() ([]byte, []int) {
	return file_conf_conf_driver_proto_rawDescGZIP(), []int{0}
}

func (x *ConfDriver) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ConfDriver) GetApollo() *ConfApollo {
	if x != nil {
		return x.Apollo
	}
	return nil
}

func (x *ConfDriver) GetConsul() *ConfConsul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *ConfDriver) GetNacos() *ConfNacos {
	if x != nil {
		return x.Nacos
	}
	return nil
}

// apollo config
type ConfApollo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId         string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Cluster       string `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Endpoint      string `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	NamespaceName string `protobuf:"bytes,4,opt,name=namespace_name,json=namespaceName,proto3" json:"namespace_name,omitempty"`
	Secret        string `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *ConfApollo) Reset() {
	*x = ConfApollo{}
	mi := &file_conf_conf_driver_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfApollo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfApollo) ProtoMessage() {}

func (x *ConfApollo) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_driver_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfApollo.ProtoReflect.Descriptor instead.
func (*ConfApollo) Descriptor() ([]byte, []int) {
	return file_conf_conf_driver_proto_rawDescGZIP(), []int{1}
}

func (x *ConfApollo) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *ConfApollo) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *ConfApollo) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *ConfApollo) GetNamespaceName() string {
	if x != nil {
		return x.NamespaceName
	}
	return ""
}

func (x *ConfApollo) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

// consul config
type ConfConsul struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Scheme     string `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	PathPrefix string `protobuf:"bytes,3,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	Token      string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Path       string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ConfConsul) Reset() {
	*x = ConfConsul{}
	mi := &file_conf_conf_driver_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfConsul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfConsul) ProtoMessage() {}

func (x *ConfConsul) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_driver_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfConsul.ProtoReflect.Descriptor instead.
func (*ConfConsul) Descriptor() ([]byte, []int) {
	return file_conf_conf_driver_proto_rawDescGZIP(), []int{2}
}

func (x *ConfConsul) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ConfConsul) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *ConfConsul) GetPathPrefix() string {
	if x != nil {
		return x.PathPrefix
	}
	return ""
}

func (x *ConfConsul) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ConfConsul) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// nacos config
type ConfNacos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addrs     []string `protobuf:"bytes,1,rep,name=addrs,proto3" json:"addrs,omitempty"`
	Namespace string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Username  string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password  string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	DataId    string   `protobuf:"bytes,5,opt,name=data_id,json=dataId,proto3" json:"data_id,omitempty"`
	Group     string   `protobuf:"bytes,6,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *ConfNacos) Reset() {
	*x = ConfNacos{}
	mi := &file_conf_conf_driver_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfNacos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfNacos) ProtoMessage() {}

func (x *ConfNacos) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_driver_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfNacos.ProtoReflect.Descriptor instead.
func (*ConfNacos) Descriptor() ([]byte, []int) {
	return file_conf_conf_driver_proto_rawDescGZIP(), []int{3}
}

func (x *ConfNacos) GetAddrs() []string {
	if x != nil {
		return x.Addrs
	}
	return nil
}

func (x *ConfNacos) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ConfNacos) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ConfNacos) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ConfNacos) GetDataId() string {
	if x != nil {
		return x.DataId
	}
	return ""
}

func (x *ConfNacos) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

var File_conf_conf_driver_proto protoreflect.FileDescriptor

var file_conf_conf_driver_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x22, 0xad, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x61, 0x70, 0x6f, 0x6c, 0x6c,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x41, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x52,
	0x06, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x12, 0x2e, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x2b, 0x0a, 0x05, 0x6e, 0x61, 0x63, 0x6f, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x4e, 0x61, 0x63, 0x6f, 0x73, 0x52, 0x05, 0x6e,
	0x61, 0x63, 0x6f, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x41, 0x70, 0x6f,
	0x6c, 0x6c, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22,
	0x89, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xa6, 0x01, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x66, 0x4e, 0x61, 0x63, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x61, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x42, 0x21, 0x5a, 0x1f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x6c,
	0x65, 0x61, 0x72, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_conf_driver_proto_rawDescOnce sync.Once
	file_conf_conf_driver_proto_rawDescData = file_conf_conf_driver_proto_rawDesc
)

func file_conf_conf_driver_proto_rawDescGZIP() []byte {
	file_conf_conf_driver_proto_rawDescOnce.Do(func() {
		file_conf_conf_driver_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_conf_driver_proto_rawDescData)
	})
	return file_conf_conf_driver_proto_rawDescData
}

var file_conf_conf_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_conf_conf_driver_proto_goTypes = []any{
	(*ConfDriver)(nil), // 0: kratos.api.ConfDriver
	(*ConfApollo)(nil), // 1: kratos.api.ConfApollo
	(*ConfConsul)(nil), // 2: kratos.api.ConfConsul
	(*ConfNacos)(nil),  // 3: kratos.api.ConfNacos
}
var file_conf_conf_driver_proto_depIdxs = []int32{
	1, // 0: kratos.api.ConfDriver.apollo:type_name -> kratos.api.ConfApollo
	2, // 1: kratos.api.ConfDriver.consul:type_name -> kratos.api.ConfConsul
	3, // 2: kratos.api.ConfDriver.nacos:type_name -> kratos.api.ConfNacos
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_conf_conf_driver_proto_init() }
func file_conf_conf_driver_proto_init() {
	if File_conf_conf_driver_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conf_conf_driver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_conf_driver_proto_goTypes,
		DependencyIndexes: file_conf_conf_driver_proto_depIdxs,
		MessageInfos:      file_conf_conf_driver_proto_msgTypes,
	}.Build()
	File_conf_conf_driver_proto = out.File
	file_conf_conf_driver_proto_rawDesc = nil
	file_conf_conf_driver_proto_goTypes = nil
	file_conf_conf_driver_proto_depIdxs = nil
}
