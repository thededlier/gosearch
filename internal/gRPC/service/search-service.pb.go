// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        (unknown)
// source: gosearch/internal/proto/service/search-service.proto

package service

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	domain "gosearch/internal/gRPC/domain"
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

type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info  *domain.ElasticsearchInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Error *Error                    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosearch_internal_proto_service_search_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gosearch_internal_proto_service_search_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_gosearch_internal_proto_service_search_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetInfoResponse) GetInfo() *domain.ElasticsearchInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *GetInfoResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type IndexDocumentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentsList []*domain.Document `protobuf:"bytes,1,rep,name=documentsList,proto3" json:"documentsList,omitempty"`
}

func (x *IndexDocumentsRequest) Reset() {
	*x = IndexDocumentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosearch_internal_proto_service_search_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexDocumentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexDocumentsRequest) ProtoMessage() {}

func (x *IndexDocumentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gosearch_internal_proto_service_search_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexDocumentsRequest.ProtoReflect.Descriptor instead.
func (*IndexDocumentsRequest) Descriptor() ([]byte, []int) {
	return file_gosearch_internal_proto_service_search_service_proto_rawDescGZIP(), []int{1}
}

func (x *IndexDocumentsRequest) GetDocumentsList() []*domain.Document {
	if x != nil {
		return x.DocumentsList
	}
	return nil
}

type IndexDocumentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentsList []*domain.Document `protobuf:"bytes,1,rep,name=documentsList,proto3" json:"documentsList,omitempty"`
	Failed        []string           `protobuf:"bytes,2,rep,name=failed,proto3" json:"failed,omitempty"`
	Error         *Error             `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *IndexDocumentsResponse) Reset() {
	*x = IndexDocumentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosearch_internal_proto_service_search_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexDocumentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexDocumentsResponse) ProtoMessage() {}

func (x *IndexDocumentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gosearch_internal_proto_service_search_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexDocumentsResponse.ProtoReflect.Descriptor instead.
func (*IndexDocumentsResponse) Descriptor() ([]byte, []int) {
	return file_gosearch_internal_proto_service_search_service_proto_rawDescGZIP(), []int{2}
}

func (x *IndexDocumentsResponse) GetDocumentsList() []*domain.Document {
	if x != nil {
		return x.DocumentsList
	}
	return nil
}

func (x *IndexDocumentsResponse) GetFailed() []string {
	if x != nil {
		return x.Failed
	}
	return nil
}

func (x *IndexDocumentsResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosearch_internal_proto_service_search_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_gosearch_internal_proto_service_search_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_gosearch_internal_proto_service_search_service_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalHits int32              `protobuf:"varint,1,opt,name=totalHits,proto3" json:"totalHits,omitempty"`
	Hits      []*domain.Document `protobuf:"bytes,2,rep,name=hits,proto3" json:"hits,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosearch_internal_proto_service_search_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gosearch_internal_proto_service_search_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_gosearch_internal_proto_service_search_service_proto_rawDescGZIP(), []int{4}
}

func (x *SearchResponse) GetTotalHits() int32 {
	if x != nil {
		return x.TotalHits
	}
	return 0
}

func (x *SearchResponse) GetHits() []*domain.Document {
	if x != nil {
		return x.Hits
	}
	return nil
}

type CustomSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indexes string `protobuf:"bytes,1,opt,name=indexes,proto3" json:"indexes,omitempty"`
	Query   []byte `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *CustomSearchRequest) Reset() {
	*x = CustomSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosearch_internal_proto_service_search_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomSearchRequest) ProtoMessage() {}

func (x *CustomSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gosearch_internal_proto_service_search_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomSearchRequest.ProtoReflect.Descriptor instead.
func (*CustomSearchRequest) Descriptor() ([]byte, []int) {
	return file_gosearch_internal_proto_service_search_service_proto_rawDescGZIP(), []int{5}
}

func (x *CustomSearchRequest) GetIndexes() string {
	if x != nil {
		return x.Indexes
	}
	return ""
}

func (x *CustomSearchRequest) GetQuery() []byte {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_gosearch_internal_proto_service_search_service_proto protoreflect.FileDescriptor

var file_gosearch_internal_proto_service_search_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x2b, 0x67, 0x6f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x24,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x4f, 0x0a, 0x15, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x12, 0x24, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x54, 0x0a,
	0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x48, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x48, 0x69, 0x74, 0x73, 0x12, 0x24, 0x0a,
	0x04, 0x68, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x68,
	0x69, 0x74, 0x73, 0x22, 0x45, 0x0a, 0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x32, 0xe4, 0x01, 0x0a, 0x0d, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x07,
	0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x6f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gosearch_internal_proto_service_search_service_proto_rawDescOnce sync.Once
	file_gosearch_internal_proto_service_search_service_proto_rawDescData = file_gosearch_internal_proto_service_search_service_proto_rawDesc
)

func file_gosearch_internal_proto_service_search_service_proto_rawDescGZIP() []byte {
	file_gosearch_internal_proto_service_search_service_proto_rawDescOnce.Do(func() {
		file_gosearch_internal_proto_service_search_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_gosearch_internal_proto_service_search_service_proto_rawDescData)
	})
	return file_gosearch_internal_proto_service_search_service_proto_rawDescData
}

var file_gosearch_internal_proto_service_search_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gosearch_internal_proto_service_search_service_proto_goTypes = []interface{}{
	(*GetInfoResponse)(nil),          // 0: service.GetInfoResponse
	(*IndexDocumentsRequest)(nil),    // 1: service.IndexDocumentsRequest
	(*IndexDocumentsResponse)(nil),   // 2: service.IndexDocumentsResponse
	(*Error)(nil),                    // 3: service.Error
	(*SearchResponse)(nil),           // 4: service.SearchResponse
	(*CustomSearchRequest)(nil),      // 5: service.CustomSearchRequest
	(*domain.ElasticsearchInfo)(nil), // 6: domain.ElasticsearchInfo
	(*domain.Document)(nil),          // 7: domain.Document
	(*domain.EmptyRequest)(nil),      // 8: domain.EmptyRequest
}
var file_gosearch_internal_proto_service_search_service_proto_depIdxs = []int32{
	6, // 0: service.GetInfoResponse.info:type_name -> domain.ElasticsearchInfo
	3, // 1: service.GetInfoResponse.error:type_name -> service.Error
	7, // 2: service.IndexDocumentsRequest.documentsList:type_name -> domain.Document
	7, // 3: service.IndexDocumentsResponse.documentsList:type_name -> domain.Document
	3, // 4: service.IndexDocumentsResponse.error:type_name -> service.Error
	7, // 5: service.SearchResponse.hits:type_name -> domain.Document
	8, // 6: service.SearchService.getInfo:input_type -> domain.EmptyRequest
	1, // 7: service.SearchService.indexDocuments:input_type -> service.IndexDocumentsRequest
	5, // 8: service.SearchService.customSearch:input_type -> service.CustomSearchRequest
	0, // 9: service.SearchService.getInfo:output_type -> service.GetInfoResponse
	2, // 10: service.SearchService.indexDocuments:output_type -> service.IndexDocumentsResponse
	4, // 11: service.SearchService.customSearch:output_type -> service.SearchResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_gosearch_internal_proto_service_search_service_proto_init() }
func file_gosearch_internal_proto_service_search_service_proto_init() {
	if File_gosearch_internal_proto_service_search_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gosearch_internal_proto_service_search_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResponse); i {
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
		file_gosearch_internal_proto_service_search_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexDocumentsRequest); i {
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
		file_gosearch_internal_proto_service_search_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexDocumentsResponse); i {
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
		file_gosearch_internal_proto_service_search_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_gosearch_internal_proto_service_search_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_gosearch_internal_proto_service_search_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomSearchRequest); i {
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
			RawDescriptor: file_gosearch_internal_proto_service_search_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gosearch_internal_proto_service_search_service_proto_goTypes,
		DependencyIndexes: file_gosearch_internal_proto_service_search_service_proto_depIdxs,
		MessageInfos:      file_gosearch_internal_proto_service_search_service_proto_msgTypes,
	}.Build()
	File_gosearch_internal_proto_service_search_service_proto = out.File
	file_gosearch_internal_proto_service_search_service_proto_rawDesc = nil
	file_gosearch_internal_proto_service_search_service_proto_goTypes = nil
	file_gosearch_internal_proto_service_search_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchServiceClient interface {
	GetInfo(ctx context.Context, in *domain.EmptyRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
	IndexDocuments(ctx context.Context, in *IndexDocumentsRequest, opts ...grpc.CallOption) (*IndexDocumentsResponse, error)
	CustomSearch(ctx context.Context, in *CustomSearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) GetInfo(ctx context.Context, in *domain.EmptyRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, "/service.SearchService/getInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) IndexDocuments(ctx context.Context, in *IndexDocumentsRequest, opts ...grpc.CallOption) (*IndexDocumentsResponse, error) {
	out := new(IndexDocumentsResponse)
	err := c.cc.Invoke(ctx, "/service.SearchService/indexDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) CustomSearch(ctx context.Context, in *CustomSearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/service.SearchService/customSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
type SearchServiceServer interface {
	GetInfo(context.Context, *domain.EmptyRequest) (*GetInfoResponse, error)
	IndexDocuments(context.Context, *IndexDocumentsRequest) (*IndexDocumentsResponse, error)
	CustomSearch(context.Context, *CustomSearchRequest) (*SearchResponse, error)
}

// UnimplementedSearchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (*UnimplementedSearchServiceServer) GetInfo(context.Context, *domain.EmptyRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (*UnimplementedSearchServiceServer) IndexDocuments(context.Context, *IndexDocumentsRequest) (*IndexDocumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IndexDocuments not implemented")
}
func (*UnimplementedSearchServiceServer) CustomSearch(context.Context, *CustomSearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomSearch not implemented")
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SearchService/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetInfo(ctx, req.(*domain.EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_IndexDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).IndexDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SearchService/IndexDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).IndexDocuments(ctx, req.(*IndexDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_CustomSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).CustomSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SearchService/CustomSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).CustomSearch(ctx, req.(*CustomSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getInfo",
			Handler:    _SearchService_GetInfo_Handler,
		},
		{
			MethodName: "indexDocuments",
			Handler:    _SearchService_IndexDocuments_Handler,
		},
		{
			MethodName: "customSearch",
			Handler:    _SearchService_CustomSearch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gosearch/internal/proto/service/search-service.proto",
}
