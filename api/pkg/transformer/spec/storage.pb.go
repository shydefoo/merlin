// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: transformer/spec/storage.proto

package spec

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// ServingType indicate type of interaction to Feast
type ServingType int32

const (
	ServingType_FEAST_GRPC     ServingType = 0 // Features are retrieved through GRPC web service
	ServingType_DIRECT_STORAGE ServingType = 1 // Features are retrieved by fetching feast storage directly
)

// Enum value maps for ServingType.
var (
	ServingType_name = map[int32]string{
		0: "FEAST_GRPC",
		1: "DIRECT_STORAGE",
	}
	ServingType_value = map[string]int32{
		"FEAST_GRPC":     0,
		"DIRECT_STORAGE": 1,
	}
)

func (x ServingType) Enum() *ServingType {
	p := new(ServingType)
	*p = x
	return p
}

func (x ServingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServingType) Descriptor() protoreflect.EnumDescriptor {
	return file_transformer_spec_storage_proto_enumTypes[0].Descriptor()
}

func (ServingType) Type() protoreflect.EnumType {
	return &file_transformer_spec_storage_proto_enumTypes[0]
}

func (x ServingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServingType.Descriptor instead.
func (ServingType) EnumDescriptor() ([]byte, []int) {
	return file_transformer_spec_storage_proto_rawDescGZIP(), []int{0}
}

type OnlineStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServingType ServingType `protobuf:"varint,1,opt,name=servingType,proto3,enum=merlin.transformer.ServingType" json:"servingType,omitempty"` // Type of interaction to feast
	// Types that are assignable to Storage:
	//	*OnlineStorage_Redis
	//	*OnlineStorage_RedisCluster
	//	*OnlineStorage_Bigtable
	Storage isOnlineStorage_Storage `protobuf_oneof:"storage"`
}

func (x *OnlineStorage) Reset() {
	*x = OnlineStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineStorage) ProtoMessage() {}

func (x *OnlineStorage) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineStorage.ProtoReflect.Descriptor instead.
func (*OnlineStorage) Descriptor() ([]byte, []int) {
	return file_transformer_spec_storage_proto_rawDescGZIP(), []int{0}
}

func (x *OnlineStorage) GetServingType() ServingType {
	if x != nil {
		return x.ServingType
	}
	return ServingType_FEAST_GRPC
}

func (m *OnlineStorage) GetStorage() isOnlineStorage_Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (x *OnlineStorage) GetRedis() *RedisStorage {
	if x, ok := x.GetStorage().(*OnlineStorage_Redis); ok {
		return x.Redis
	}
	return nil
}

func (x *OnlineStorage) GetRedisCluster() *RedisClusterStorage {
	if x, ok := x.GetStorage().(*OnlineStorage_RedisCluster); ok {
		return x.RedisCluster
	}
	return nil
}

func (x *OnlineStorage) GetBigtable() *BigTableStorage {
	if x, ok := x.GetStorage().(*OnlineStorage_Bigtable); ok {
		return x.Bigtable
	}
	return nil
}

type isOnlineStorage_Storage interface {
	isOnlineStorage_Storage()
}

type OnlineStorage_Redis struct {
	Redis *RedisStorage `protobuf:"bytes,2,opt,name=redis,proto3,oneof"`
}

type OnlineStorage_RedisCluster struct {
	RedisCluster *RedisClusterStorage `protobuf:"bytes,3,opt,name=redisCluster,proto3,oneof"`
}

type OnlineStorage_Bigtable struct {
	Bigtable *BigTableStorage `protobuf:"bytes,4,opt,name=bigtable,proto3,oneof"`
}

func (*OnlineStorage_Redis) isOnlineStorage_Storage() {}

func (*OnlineStorage_RedisCluster) isOnlineStorage_Storage() {}

func (*OnlineStorage_Bigtable) isOnlineStorage_Storage() {}

type BigTableStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeastServingUrl string `protobuf:"bytes,1,opt,name=feastServingUrl,proto3" json:"feastServingUrl,omitempty"` // Feast GRPC URL
}

func (x *BigTableStorage) Reset() {
	*x = BigTableStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigTableStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigTableStorage) ProtoMessage() {}

func (x *BigTableStorage) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigTableStorage.ProtoReflect.Descriptor instead.
func (*BigTableStorage) Descriptor() ([]byte, []int) {
	return file_transformer_spec_storage_proto_rawDescGZIP(), []int{1}
}

func (x *BigTableStorage) GetFeastServingUrl() string {
	if x != nil {
		return x.FeastServingUrl
	}
	return ""
}

type RedisStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeastServingUrl string       `protobuf:"bytes,1,opt,name=feastServingUrl,proto3" json:"feastServingUrl,omitempty"` // Feast GRPC URL
	RedisAddress    string       `protobuf:"bytes,2,opt,name=redisAddress,proto3" json:"redisAddress,omitempty"`       // Address of single redis
	Option          *RedisOption `protobuf:"bytes,3,opt,name=option,proto3" json:"option,omitempty"`                   // Redis option
}

func (x *RedisStorage) Reset() {
	*x = RedisStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisStorage) ProtoMessage() {}

func (x *RedisStorage) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisStorage.ProtoReflect.Descriptor instead.
func (*RedisStorage) Descriptor() ([]byte, []int) {
	return file_transformer_spec_storage_proto_rawDescGZIP(), []int{2}
}

func (x *RedisStorage) GetFeastServingUrl() string {
	if x != nil {
		return x.FeastServingUrl
	}
	return ""
}

func (x *RedisStorage) GetRedisAddress() string {
	if x != nil {
		return x.RedisAddress
	}
	return ""
}

func (x *RedisStorage) GetOption() *RedisOption {
	if x != nil {
		return x.Option
	}
	return nil
}

type RedisClusterStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeastServingUrl string       `protobuf:"bytes,1,opt,name=feastServingUrl,proto3" json:"feastServingUrl,omitempty"` // Feast GRPC URL
	RedisAddress    []string     `protobuf:"bytes,2,rep,name=redisAddress,proto3" json:"redisAddress,omitempty"`       // List of addresses of redis cluster
	Option          *RedisOption `protobuf:"bytes,3,opt,name=option,proto3" json:"option,omitempty"`                   // Redis option
}

func (x *RedisClusterStorage) Reset() {
	*x = RedisClusterStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisClusterStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisClusterStorage) ProtoMessage() {}

func (x *RedisClusterStorage) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisClusterStorage.ProtoReflect.Descriptor instead.
func (*RedisClusterStorage) Descriptor() ([]byte, []int) {
	return file_transformer_spec_storage_proto_rawDescGZIP(), []int{3}
}

func (x *RedisClusterStorage) GetFeastServingUrl() string {
	if x != nil {
		return x.FeastServingUrl
	}
	return ""
}

func (x *RedisClusterStorage) GetRedisAddress() []string {
	if x != nil {
		return x.RedisAddress
	}
	return nil
}

func (x *RedisClusterStorage) GetOption() *RedisOption {
	if x != nil {
		return x.Option
	}
	return nil
}

// options to be passed to the goredis client
type RedisOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxRetries         int32                `protobuf:"varint,1,opt,name=maxRetries,proto3" json:"maxRetries,omitempty"`                 // Number of maximum retry
	MinRetryBackoff    *durationpb.Duration `protobuf:"bytes,2,opt,name=minRetryBackoff,proto3" json:"minRetryBackoff,omitempty"`        // minimum duration of backoff between retry
	DialTimeout        *durationpb.Duration `protobuf:"bytes,3,opt,name=dialTimeout,proto3" json:"dialTimeout,omitempty"`                // dial timeout for establishing new connection
	ReadTimeout        *durationpb.Duration `protobuf:"bytes,4,opt,name=readTimeout,proto3" json:"readTimeout,omitempty"`                // read timeout duration for redis command
	WriteTimeout       *durationpb.Duration `protobuf:"bytes,5,opt,name=writeTimeout,proto3" json:"writeTimeout,omitempty"`              // write timeout duration for redis command
	PoolSize           int32                `protobuf:"varint,6,opt,name=poolSize,proto3" json:"poolSize,omitempty"`                     // maximum number of connection created that later on will be reused
	MaxConnAge         *durationpb.Duration `protobuf:"bytes,7,opt,name=maxConnAge,proto3" json:"maxConnAge,omitempty"`                  // maximum age of redis connection
	PoolTimeout        *durationpb.Duration `protobuf:"bytes,8,opt,name=poolTimeout,proto3" json:"poolTimeout,omitempty"`                // amount of time clients wait for connection if all connection are busy before returning error
	IdleTimeout        *durationpb.Duration `protobuf:"bytes,9,opt,name=idleTimeout,proto3" json:"idleTimeout,omitempty"`                // amount of time after which client closes idle connections
	IdleCheckFrequency *durationpb.Duration `protobuf:"bytes,10,opt,name=idleCheckFrequency,proto3" json:"idleCheckFrequency,omitempty"` // frequency of idle check
}

func (x *RedisOption) Reset() {
	*x = RedisOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisOption) ProtoMessage() {}

func (x *RedisOption) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisOption.ProtoReflect.Descriptor instead.
func (*RedisOption) Descriptor() ([]byte, []int) {
	return file_transformer_spec_storage_proto_rawDescGZIP(), []int{4}
}

func (x *RedisOption) GetMaxRetries() int32 {
	if x != nil {
		return x.MaxRetries
	}
	return 0
}

func (x *RedisOption) GetMinRetryBackoff() *durationpb.Duration {
	if x != nil {
		return x.MinRetryBackoff
	}
	return nil
}

func (x *RedisOption) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

func (x *RedisOption) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *RedisOption) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *RedisOption) GetPoolSize() int32 {
	if x != nil {
		return x.PoolSize
	}
	return 0
}

func (x *RedisOption) GetMaxConnAge() *durationpb.Duration {
	if x != nil {
		return x.MaxConnAge
	}
	return nil
}

func (x *RedisOption) GetPoolTimeout() *durationpb.Duration {
	if x != nil {
		return x.PoolTimeout
	}
	return nil
}

func (x *RedisOption) GetIdleTimeout() *durationpb.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *RedisOption) GetIdleCheckFrequency() *durationpb.Duration {
	if x != nil {
		return x.IdleCheckFrequency
	}
	return nil
}

var File_transformer_spec_storage_proto protoreflect.FileDescriptor

var file_transformer_spec_storage_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2f, 0x73, 0x70,
	0x65, 0x63, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x65, 0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x02, 0x0a, 0x0d, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x65,
	0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69,
	0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65,
	0x64, 0x69, 0x73, 0x12, 0x4d, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x65, 0x72, 0x6c,
	0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x41, 0x0a, 0x08, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x67, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x08, 0x62, 0x69, 0x67,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x22, 0x3b, 0x0a, 0x0f, 0x42, 0x69, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x65,
	0x61, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x22, 0x95, 0x01,
	0x0a, 0x0c, 0x52, 0x65, 0x64, 0x69, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x55, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x37, 0x0a, 0x06,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d,
	0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9c, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x73,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x64, 0x69, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x37, 0x0a, 0x06, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x65,
	0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc7, 0x04, 0x0a, 0x0b, 0x52, 0x65, 0x64, 0x69, 0x73, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x0f, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x74, 0x72, 0x79,
	0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x74,
	0x72, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x69, 0x61,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x41, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6d,
	0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x41, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x70, 0x6f, 0x6f,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x6f, 0x6f, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x49, 0x0a, 0x12, 0x69, 0x64, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x69, 0x64, 0x6c, 0x65,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x2a, 0x31,
	0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x0a, 0x46, 0x45, 0x41, 0x53, 0x54, 0x5f, 0x47, 0x52, 0x50, 0x43, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10,
	0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x6a, 0x65, 0x6b, 0x2f, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2f, 0x73, 0x70, 0x65,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transformer_spec_storage_proto_rawDescOnce sync.Once
	file_transformer_spec_storage_proto_rawDescData = file_transformer_spec_storage_proto_rawDesc
)

func file_transformer_spec_storage_proto_rawDescGZIP() []byte {
	file_transformer_spec_storage_proto_rawDescOnce.Do(func() {
		file_transformer_spec_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_transformer_spec_storage_proto_rawDescData)
	})
	return file_transformer_spec_storage_proto_rawDescData
}

var file_transformer_spec_storage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_transformer_spec_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_transformer_spec_storage_proto_goTypes = []interface{}{
	(ServingType)(0),            // 0: merlin.transformer.ServingType
	(*OnlineStorage)(nil),       // 1: merlin.transformer.OnlineStorage
	(*BigTableStorage)(nil),     // 2: merlin.transformer.BigTableStorage
	(*RedisStorage)(nil),        // 3: merlin.transformer.RedisStorage
	(*RedisClusterStorage)(nil), // 4: merlin.transformer.RedisClusterStorage
	(*RedisOption)(nil),         // 5: merlin.transformer.RedisOption
	(*durationpb.Duration)(nil), // 6: google.protobuf.Duration
}
var file_transformer_spec_storage_proto_depIdxs = []int32{
	0,  // 0: merlin.transformer.OnlineStorage.servingType:type_name -> merlin.transformer.ServingType
	3,  // 1: merlin.transformer.OnlineStorage.redis:type_name -> merlin.transformer.RedisStorage
	4,  // 2: merlin.transformer.OnlineStorage.redisCluster:type_name -> merlin.transformer.RedisClusterStorage
	2,  // 3: merlin.transformer.OnlineStorage.bigtable:type_name -> merlin.transformer.BigTableStorage
	5,  // 4: merlin.transformer.RedisStorage.option:type_name -> merlin.transformer.RedisOption
	5,  // 5: merlin.transformer.RedisClusterStorage.option:type_name -> merlin.transformer.RedisOption
	6,  // 6: merlin.transformer.RedisOption.minRetryBackoff:type_name -> google.protobuf.Duration
	6,  // 7: merlin.transformer.RedisOption.dialTimeout:type_name -> google.protobuf.Duration
	6,  // 8: merlin.transformer.RedisOption.readTimeout:type_name -> google.protobuf.Duration
	6,  // 9: merlin.transformer.RedisOption.writeTimeout:type_name -> google.protobuf.Duration
	6,  // 10: merlin.transformer.RedisOption.maxConnAge:type_name -> google.protobuf.Duration
	6,  // 11: merlin.transformer.RedisOption.poolTimeout:type_name -> google.protobuf.Duration
	6,  // 12: merlin.transformer.RedisOption.idleTimeout:type_name -> google.protobuf.Duration
	6,  // 13: merlin.transformer.RedisOption.idleCheckFrequency:type_name -> google.protobuf.Duration
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_transformer_spec_storage_proto_init() }
func file_transformer_spec_storage_proto_init() {
	if File_transformer_spec_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transformer_spec_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineStorage); i {
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
		file_transformer_spec_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigTableStorage); i {
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
		file_transformer_spec_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisStorage); i {
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
		file_transformer_spec_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisClusterStorage); i {
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
		file_transformer_spec_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisOption); i {
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
	file_transformer_spec_storage_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OnlineStorage_Redis)(nil),
		(*OnlineStorage_RedisCluster)(nil),
		(*OnlineStorage_Bigtable)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transformer_spec_storage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transformer_spec_storage_proto_goTypes,
		DependencyIndexes: file_transformer_spec_storage_proto_depIdxs,
		EnumInfos:         file_transformer_spec_storage_proto_enumTypes,
		MessageInfos:      file_transformer_spec_storage_proto_msgTypes,
	}.Build()
	File_transformer_spec_storage_proto = out.File
	file_transformer_spec_storage_proto_rawDesc = nil
	file_transformer_spec_storage_proto_goTypes = nil
	file_transformer_spec_storage_proto_depIdxs = nil
}
