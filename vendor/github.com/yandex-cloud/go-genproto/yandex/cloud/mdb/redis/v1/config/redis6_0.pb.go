// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: yandex/cloud/mdb/redis/v1/config/redis6_0.proto

package redis

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type RedisConfig6_0_MaxmemoryPolicy int32

const (
	RedisConfig6_0_MAXMEMORY_POLICY_UNSPECIFIED RedisConfig6_0_MaxmemoryPolicy = 0
	// Try to remove less recently used (LRU) keys with `expire set`.
	RedisConfig6_0_VOLATILE_LRU RedisConfig6_0_MaxmemoryPolicy = 1
	// Remove less recently used (LRU) keys.
	RedisConfig6_0_ALLKEYS_LRU RedisConfig6_0_MaxmemoryPolicy = 2
	// Try to remove least frequently used (LFU) keys with `expire set`.
	RedisConfig6_0_VOLATILE_LFU RedisConfig6_0_MaxmemoryPolicy = 3
	// Remove least frequently used (LFU) keys.
	RedisConfig6_0_ALLKEYS_LFU RedisConfig6_0_MaxmemoryPolicy = 4
	// Try to remove keys with `expire set` randomly.
	RedisConfig6_0_VOLATILE_RANDOM RedisConfig6_0_MaxmemoryPolicy = 5
	// Remove keys randomly.
	RedisConfig6_0_ALLKEYS_RANDOM RedisConfig6_0_MaxmemoryPolicy = 6
	// Try to remove less recently used (LRU) keys with `expire set`
	// and shorter TTL first.
	RedisConfig6_0_VOLATILE_TTL RedisConfig6_0_MaxmemoryPolicy = 7
	// Return errors when memory limit was reached and commands could require
	// more memory to be used.
	RedisConfig6_0_NOEVICTION RedisConfig6_0_MaxmemoryPolicy = 8
)

// Enum value maps for RedisConfig6_0_MaxmemoryPolicy.
var (
	RedisConfig6_0_MaxmemoryPolicy_name = map[int32]string{
		0: "MAXMEMORY_POLICY_UNSPECIFIED",
		1: "VOLATILE_LRU",
		2: "ALLKEYS_LRU",
		3: "VOLATILE_LFU",
		4: "ALLKEYS_LFU",
		5: "VOLATILE_RANDOM",
		6: "ALLKEYS_RANDOM",
		7: "VOLATILE_TTL",
		8: "NOEVICTION",
	}
	RedisConfig6_0_MaxmemoryPolicy_value = map[string]int32{
		"MAXMEMORY_POLICY_UNSPECIFIED": 0,
		"VOLATILE_LRU":                 1,
		"ALLKEYS_LRU":                  2,
		"VOLATILE_LFU":                 3,
		"ALLKEYS_LFU":                  4,
		"VOLATILE_RANDOM":              5,
		"ALLKEYS_RANDOM":               6,
		"VOLATILE_TTL":                 7,
		"NOEVICTION":                   8,
	}
)

func (x RedisConfig6_0_MaxmemoryPolicy) Enum() *RedisConfig6_0_MaxmemoryPolicy {
	p := new(RedisConfig6_0_MaxmemoryPolicy)
	*p = x
	return p
}

func (x RedisConfig6_0_MaxmemoryPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RedisConfig6_0_MaxmemoryPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_enumTypes[0].Descriptor()
}

func (RedisConfig6_0_MaxmemoryPolicy) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_enumTypes[0]
}

func (x RedisConfig6_0_MaxmemoryPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RedisConfig6_0_MaxmemoryPolicy.Descriptor instead.
func (RedisConfig6_0_MaxmemoryPolicy) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescGZIP(), []int{0, 0}
}

// Fields and structure of `RedisConfig` reflects Redis configuration file
// parameters.
type RedisConfig6_0 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Redis key eviction policy for a dataset that reaches maximum memory,
	// available to the host. Redis maxmemory setting depends on Managed
	// Service for Redis [host class](/docs/managed-redis/concepts/instance-types).
	//
	// All policies are described in detail in [Redis documentation](https://redis.io/topics/lru-cache).
	MaxmemoryPolicy RedisConfig6_0_MaxmemoryPolicy `protobuf:"varint,1,opt,name=maxmemory_policy,json=maxmemoryPolicy,proto3,enum=yandex.cloud.mdb.redis.v1.config.RedisConfig6_0_MaxmemoryPolicy" json:"maxmemory_policy,omitempty"`
	// Time that Redis keeps the connection open while the client is idle.
	// If no new command is sent during that time, the connection is closed.
	Timeout *wrappers.Int64Value `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Authentication password.
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RedisConfig6_0) Reset() {
	*x = RedisConfig6_0{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisConfig6_0) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisConfig6_0) ProtoMessage() {}

func (x *RedisConfig6_0) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisConfig6_0.ProtoReflect.Descriptor instead.
func (*RedisConfig6_0) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescGZIP(), []int{0}
}

func (x *RedisConfig6_0) GetMaxmemoryPolicy() RedisConfig6_0_MaxmemoryPolicy {
	if x != nil {
		return x.MaxmemoryPolicy
	}
	return RedisConfig6_0_MAXMEMORY_POLICY_UNSPECIFIED
}

func (x *RedisConfig6_0) GetTimeout() *wrappers.Int64Value {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *RedisConfig6_0) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RedisConfigSet6_0 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Effective settings for a Redis 6.0 cluster (a combination of settings
	// defined in [user_config] and [default_config]).
	EffectiveConfig *RedisConfig6_0 `protobuf:"bytes,1,opt,name=effective_config,json=effectiveConfig,proto3" json:"effective_config,omitempty"`
	// User-defined settings for a Redis 6.0 cluster.
	UserConfig *RedisConfig6_0 `protobuf:"bytes,2,opt,name=user_config,json=userConfig,proto3" json:"user_config,omitempty"`
	// Default configuration for a Redis 6.0 cluster.
	DefaultConfig *RedisConfig6_0 `protobuf:"bytes,3,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
}

func (x *RedisConfigSet6_0) Reset() {
	*x = RedisConfigSet6_0{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisConfigSet6_0) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisConfigSet6_0) ProtoMessage() {}

func (x *RedisConfigSet6_0) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisConfigSet6_0.ProtoReflect.Descriptor instead.
func (*RedisConfigSet6_0) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescGZIP(), []int{1}
}

func (x *RedisConfigSet6_0) GetEffectiveConfig() *RedisConfig6_0 {
	if x != nil {
		return x.EffectiveConfig
	}
	return nil
}

func (x *RedisConfigSet6_0) GetUserConfig() *RedisConfig6_0 {
	if x != nil {
		return x.UserConfig
	}
	return nil
}

func (x *RedisConfigSet6_0) GetDefaultConfig() *RedisConfig6_0 {
	if x != nil {
		return x.DefaultConfig
	}
	return nil
}

var File_yandex_cloud_mdb_redis_v1_config_redis6_0_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x36, 0x5f, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbf, 0x03, 0x0a, 0x0e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x36, 0x5f, 0x30, 0x12, 0x6b, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x40, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d,
	0x64, 0x62, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x36, 0x5f,
	0x30, 0x2e, 0x4d, 0x61, 0x78, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x35, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x42, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xf2, 0xc7, 0x31,
	0x22, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x40, 0x3d, 0x2b, 0x3f, 0x2a,
	0x2e, 0x2c, 0x21, 0x26, 0x23, 0x24, 0x5e, 0x3c, 0x3e, 0x5f, 0x2d, 0x5d, 0x7b, 0x38, 0x2c, 0x31,
	0x32, 0x38, 0x7d, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xc4, 0x01,
	0x0a, 0x0f, 0x4d, 0x61, 0x78, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x41, 0x58, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x5f, 0x50,
	0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x56, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4c, 0x45, 0x5f,
	0x4c, 0x52, 0x55, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x4c, 0x4c, 0x4b, 0x45, 0x59, 0x53,
	0x5f, 0x4c, 0x52, 0x55, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x56, 0x4f, 0x4c, 0x41, 0x54, 0x49,
	0x4c, 0x45, 0x5f, 0x4c, 0x46, 0x55, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x4c, 0x4c, 0x4b,
	0x45, 0x59, 0x53, 0x5f, 0x4c, 0x46, 0x55, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x56, 0x4f, 0x4c,
	0x41, 0x54, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x10, 0x05, 0x12, 0x12,
	0x0a, 0x0e, 0x41, 0x4c, 0x4c, 0x4b, 0x45, 0x59, 0x53, 0x5f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d,
	0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x56, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4c, 0x45, 0x5f, 0x54,
	0x54, 0x4c, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x45, 0x56, 0x49, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x08, 0x22, 0x9c, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x74, 0x36, 0x5f, 0x30, 0x12, 0x5b, 0x0a, 0x10, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x36, 0x5f, 0x30, 0x52, 0x0f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x51, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x36, 0x5f, 0x30, 0x52, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x57, 0x0a, 0x0e, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x36, 0x5f, 0x30, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x72, 0x0a, 0x24, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5a, 0x4a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64,
	0x62, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x3b, 0x72, 0x65, 0x64, 0x69, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescData = file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDesc
)

func file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescData
}

var file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_goTypes = []interface{}{
	(RedisConfig6_0_MaxmemoryPolicy)(0), // 0: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.MaxmemoryPolicy
	(*RedisConfig6_0)(nil),              // 1: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0
	(*RedisConfigSet6_0)(nil),           // 2: yandex.cloud.mdb.redis.v1.config.RedisConfigSet6_0
	(*wrappers.Int64Value)(nil),         // 3: google.protobuf.Int64Value
}
var file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.maxmemory_policy:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.MaxmemoryPolicy
	3, // 1: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.timeout:type_name -> google.protobuf.Int64Value
	1, // 2: yandex.cloud.mdb.redis.v1.config.RedisConfigSet6_0.effective_config:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig6_0
	1, // 3: yandex.cloud.mdb.redis.v1.config.RedisConfigSet6_0.user_config:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig6_0
	1, // 4: yandex.cloud.mdb.redis.v1.config.RedisConfigSet6_0.default_config:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig6_0
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_init() }
func file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_init() {
	if File_yandex_cloud_mdb_redis_v1_config_redis6_0_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisConfig6_0); i {
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
		file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisConfigSet6_0); i {
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
			RawDescriptor: file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_redis_v1_config_redis6_0_proto = out.File
	file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDesc = nil
	file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_goTypes = nil
	file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_depIdxs = nil
}