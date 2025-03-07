// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: user/private/biz/cache.proto

package biz

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

type ForgetPasswordTwoStepTokenPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ForgetPasswordTwoStepTokenPayload) Reset() {
	*x = ForgetPasswordTwoStepTokenPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_private_biz_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgetPasswordTwoStepTokenPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgetPasswordTwoStepTokenPayload) ProtoMessage() {}

func (x *ForgetPasswordTwoStepTokenPayload) ProtoReflect() protoreflect.Message {
	mi := &file_user_private_biz_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgetPasswordTwoStepTokenPayload.ProtoReflect.Descriptor instead.
func (*ForgetPasswordTwoStepTokenPayload) Descriptor() ([]byte, []int) {
	return file_user_private_biz_cache_proto_rawDescGZIP(), []int{0}
}

func (x *ForgetPasswordTwoStepTokenPayload) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserRoleCacheItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role []*UserRoleCacheItem_UserRole `protobuf:"bytes,1,rep,name=role,proto3" json:"role,omitempty"`
}

func (x *UserRoleCacheItem) Reset() {
	*x = UserRoleCacheItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_private_biz_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRoleCacheItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRoleCacheItem) ProtoMessage() {}

func (x *UserRoleCacheItem) ProtoReflect() protoreflect.Message {
	mi := &file_user_private_biz_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRoleCacheItem.ProtoReflect.Descriptor instead.
func (*UserRoleCacheItem) Descriptor() ([]byte, []int) {
	return file_user_private_biz_cache_proto_rawDescGZIP(), []int{1}
}

func (x *UserRoleCacheItem) GetRole() []*UserRoleCacheItem_UserRole {
	if x != nil {
		return x.Role
	}
	return nil
}

type UserRoleCacheItem_UserRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId   string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	TenantId string `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *UserRoleCacheItem_UserRole) Reset() {
	*x = UserRoleCacheItem_UserRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_private_biz_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRoleCacheItem_UserRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRoleCacheItem_UserRole) ProtoMessage() {}

func (x *UserRoleCacheItem_UserRole) ProtoReflect() protoreflect.Message {
	mi := &file_user_private_biz_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRoleCacheItem_UserRole.ProtoReflect.Descriptor instead.
func (*UserRoleCacheItem_UserRole) Descriptor() ([]byte, []int) {
	return file_user_private_biz_cache_proto_rawDescGZIP(), []int{1, 0}
}

func (x *UserRoleCacheItem_UserRole) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *UserRoleCacheItem_UserRole) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

var File_user_private_biz_cache_proto protoreflect.FileDescriptor

var file_user_private_biz_cache_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x62,
	0x69, 0x7a, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x22, 0x3c, 0x0a, 0x21, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x77, 0x6f, 0x53, 0x74, 0x65, 0x70, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x46, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x1a, 0x40, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x62, 0x69, 0x7a, 0x3b,
	0x62, 0x69, 0x7a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_private_biz_cache_proto_rawDescOnce sync.Once
	file_user_private_biz_cache_proto_rawDescData = file_user_private_biz_cache_proto_rawDesc
)

func file_user_private_biz_cache_proto_rawDescGZIP() []byte {
	file_user_private_biz_cache_proto_rawDescOnce.Do(func() {
		file_user_private_biz_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_private_biz_cache_proto_rawDescData)
	})
	return file_user_private_biz_cache_proto_rawDescData
}

var file_user_private_biz_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_user_private_biz_cache_proto_goTypes = []interface{}{
	(*ForgetPasswordTwoStepTokenPayload)(nil), // 0: user.private.biz.cache.ForgetPasswordTwoStepTokenPayload
	(*UserRoleCacheItem)(nil),                 // 1: user.private.biz.cache.UserRoleCacheItem
	(*UserRoleCacheItem_UserRole)(nil),        // 2: user.private.biz.cache.UserRoleCacheItem.UserRole
}
var file_user_private_biz_cache_proto_depIdxs = []int32{
	2, // 0: user.private.biz.cache.UserRoleCacheItem.role:type_name -> user.private.biz.cache.UserRoleCacheItem.UserRole
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_private_biz_cache_proto_init() }
func file_user_private_biz_cache_proto_init() {
	if File_user_private_biz_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_private_biz_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgetPasswordTwoStepTokenPayload); i {
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
		file_user_private_biz_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRoleCacheItem); i {
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
		file_user_private_biz_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRoleCacheItem_UserRole); i {
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
			RawDescriptor: file_user_private_biz_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_private_biz_cache_proto_goTypes,
		DependencyIndexes: file_user_private_biz_cache_proto_depIdxs,
		MessageInfos:      file_user_private_biz_cache_proto_msgTypes,
	}.Build()
	File_user_private_biz_cache_proto = out.File
	file_user_private_biz_cache_proto_rawDesc = nil
	file_user_private_biz_cache_proto_goTypes = nil
	file_user_private_biz_cache_proto_depIdxs = nil
}
