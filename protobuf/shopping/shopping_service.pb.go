// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.1
// source: playground/protobuf/shopping/shopping_service.proto

package shopping

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

type GetCatalogItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCatalogItemsRequest) Reset() {
	*x = GetCatalogItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogItemsRequest) ProtoMessage() {}

func (x *GetCatalogItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogItemsRequest.ProtoReflect.Descriptor instead.
func (*GetCatalogItemsRequest) Descriptor() ([]byte, []int) {
	return file_playground_protobuf_shopping_shopping_service_proto_rawDescGZIP(), []int{0}
}

type GetCatalogItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCatalogItemsResponse) Reset() {
	*x = GetCatalogItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogItemsResponse) ProtoMessage() {}

func (x *GetCatalogItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogItemsResponse.ProtoReflect.Descriptor instead.
func (*GetCatalogItemsResponse) Descriptor() ([]byte, []int) {
	return file_playground_protobuf_shopping_shopping_service_proto_rawDescGZIP(), []int{1}
}

type AddCatalogItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	PictureUrl  string `protobuf:"bytes,4,opt,name=picture_url,json=pictureUrl,proto3" json:"picture_url,omitempty"`
}

func (x *AddCatalogItemRequest) Reset() {
	*x = AddCatalogItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCatalogItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCatalogItemRequest) ProtoMessage() {}

func (x *AddCatalogItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCatalogItemRequest.ProtoReflect.Descriptor instead.
func (*AddCatalogItemRequest) Descriptor() ([]byte, []int) {
	return file_playground_protobuf_shopping_shopping_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddCatalogItemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddCatalogItemRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddCatalogItemRequest) GetPictureUrl() string {
	if x != nil {
		return x.PictureUrl
	}
	return ""
}

type AddCatalogItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddCatalogItemResponse) Reset() {
	*x = AddCatalogItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCatalogItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCatalogItemResponse) ProtoMessage() {}

func (x *AddCatalogItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCatalogItemResponse.ProtoReflect.Descriptor instead.
func (*AddCatalogItemResponse) Descriptor() ([]byte, []int) {
	return file_playground_protobuf_shopping_shopping_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddCatalogItemResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateCatalogItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	PictureUrl  string `protobuf:"bytes,4,opt,name=picture_url,json=pictureUrl,proto3" json:"picture_url,omitempty"`
}

func (x *UpdateCatalogItemRequest) Reset() {
	*x = UpdateCatalogItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCatalogItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCatalogItemRequest) ProtoMessage() {}

func (x *UpdateCatalogItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCatalogItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateCatalogItemRequest) Descriptor() ([]byte, []int) {
	return file_playground_protobuf_shopping_shopping_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCatalogItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCatalogItemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCatalogItemRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateCatalogItemRequest) GetPictureUrl() string {
	if x != nil {
		return x.PictureUrl
	}
	return ""
}

type UpdateCatalogItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCatalogItemResponse) Reset() {
	*x = UpdateCatalogItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCatalogItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCatalogItemResponse) ProtoMessage() {}

func (x *UpdateCatalogItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCatalogItemResponse.ProtoReflect.Descriptor instead.
func (*UpdateCatalogItemResponse) Descriptor() ([]byte, []int) {
	return file_playground_protobuf_shopping_shopping_service_proto_rawDescGZIP(), []int{5}
}

type DeleteCatalogItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteCatalogItemsRequest) Reset() {
	*x = DeleteCatalogItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCatalogItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCatalogItemsRequest) ProtoMessage() {}

func (x *DeleteCatalogItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCatalogItemsRequest.ProtoReflect.Descriptor instead.
func (*DeleteCatalogItemsRequest) Descriptor() ([]byte, []int) {
	return file_playground_protobuf_shopping_shopping_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCatalogItemsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteCatalogItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCatalogItemsResponse) Reset() {
	*x = DeleteCatalogItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCatalogItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCatalogItemsResponse) ProtoMessage() {}

func (x *DeleteCatalogItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCatalogItemsResponse.ProtoReflect.Descriptor instead.
func (*DeleteCatalogItemsResponse) Descriptor() ([]byte, []int) {
	return file_playground_protobuf_shopping_shopping_service_proto_rawDescGZIP(), []int{7}
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_playground_protobuf_shopping_shopping_service_proto_rawDescGZIP(), []int{8}
}

func (x *HelloRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseMessage string `protobuf:"bytes,1,opt,name=response_message,json=responseMessage,proto3" json:"response_message,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playground_protobuf_shopping_shopping_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_playground_protobuf_shopping_shopping_service_proto_rawDescGZIP(), []int{9}
}

func (x *HelloResponse) GetResponseMessage() string {
	if x != nil {
		return x.ResponseMessage
	}
	return ""
}

var File_playground_protobuf_shopping_shopping_service_proto protoreflect.FileDescriptor

var file_playground_protobuf_shopping_shopping_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x19, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6e, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x28, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x1b, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2d, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x28, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x0a, 0x0d, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x81, 0x05, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x05, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x2a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7e, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x34, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7b, 0x0a, 0x0e,
	0x41, 0x64, 0x64, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x33,
	0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x64,
	0x64, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x36, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x87, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x37, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x64, 0x0a, 0x1c, 0x70, 0x6c,
	0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x70, 0x6c, 0x61,
	0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_playground_protobuf_shopping_shopping_service_proto_rawDescOnce sync.Once
	file_playground_protobuf_shopping_shopping_service_proto_rawDescData = file_playground_protobuf_shopping_shopping_service_proto_rawDesc
)

func file_playground_protobuf_shopping_shopping_service_proto_rawDescGZIP() []byte {
	file_playground_protobuf_shopping_shopping_service_proto_rawDescOnce.Do(func() {
		file_playground_protobuf_shopping_shopping_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_playground_protobuf_shopping_shopping_service_proto_rawDescData)
	})
	return file_playground_protobuf_shopping_shopping_service_proto_rawDescData
}

var file_playground_protobuf_shopping_shopping_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_playground_protobuf_shopping_shopping_service_proto_goTypes = []interface{}{
	(*GetCatalogItemsRequest)(nil),     // 0: playground.protobuf.shopping.GetCatalogItemsRequest
	(*GetCatalogItemsResponse)(nil),    // 1: playground.protobuf.shopping.GetCatalogItemsResponse
	(*AddCatalogItemRequest)(nil),      // 2: playground.protobuf.shopping.AddCatalogItemRequest
	(*AddCatalogItemResponse)(nil),     // 3: playground.protobuf.shopping.AddCatalogItemResponse
	(*UpdateCatalogItemRequest)(nil),   // 4: playground.protobuf.shopping.UpdateCatalogItemRequest
	(*UpdateCatalogItemResponse)(nil),  // 5: playground.protobuf.shopping.UpdateCatalogItemResponse
	(*DeleteCatalogItemsRequest)(nil),  // 6: playground.protobuf.shopping.DeleteCatalogItemsRequest
	(*DeleteCatalogItemsResponse)(nil), // 7: playground.protobuf.shopping.DeleteCatalogItemsResponse
	(*HelloRequest)(nil),               // 8: playground.protobuf.shopping.HelloRequest
	(*HelloResponse)(nil),              // 9: playground.protobuf.shopping.HelloResponse
}
var file_playground_protobuf_shopping_shopping_service_proto_depIdxs = []int32{
	8, // 0: playground.protobuf.shopping.ShoppingService.Hello:input_type -> playground.protobuf.shopping.HelloRequest
	0, // 1: playground.protobuf.shopping.ShoppingService.GetCatalogItems:input_type -> playground.protobuf.shopping.GetCatalogItemsRequest
	2, // 2: playground.protobuf.shopping.ShoppingService.AddCatalogItem:input_type -> playground.protobuf.shopping.AddCatalogItemRequest
	4, // 3: playground.protobuf.shopping.ShoppingService.UpdateCatalogItem:input_type -> playground.protobuf.shopping.UpdateCatalogItemRequest
	6, // 4: playground.protobuf.shopping.ShoppingService.DeleteCatalogItems:input_type -> playground.protobuf.shopping.DeleteCatalogItemsRequest
	9, // 5: playground.protobuf.shopping.ShoppingService.Hello:output_type -> playground.protobuf.shopping.HelloResponse
	1, // 6: playground.protobuf.shopping.ShoppingService.GetCatalogItems:output_type -> playground.protobuf.shopping.GetCatalogItemsResponse
	3, // 7: playground.protobuf.shopping.ShoppingService.AddCatalogItem:output_type -> playground.protobuf.shopping.AddCatalogItemResponse
	5, // 8: playground.protobuf.shopping.ShoppingService.UpdateCatalogItem:output_type -> playground.protobuf.shopping.UpdateCatalogItemResponse
	7, // 9: playground.protobuf.shopping.ShoppingService.DeleteCatalogItems:output_type -> playground.protobuf.shopping.DeleteCatalogItemsResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_playground_protobuf_shopping_shopping_service_proto_init() }
func file_playground_protobuf_shopping_shopping_service_proto_init() {
	if File_playground_protobuf_shopping_shopping_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_playground_protobuf_shopping_shopping_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogItemsRequest); i {
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
		file_playground_protobuf_shopping_shopping_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogItemsResponse); i {
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
		file_playground_protobuf_shopping_shopping_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCatalogItemRequest); i {
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
		file_playground_protobuf_shopping_shopping_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCatalogItemResponse); i {
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
		file_playground_protobuf_shopping_shopping_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCatalogItemRequest); i {
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
		file_playground_protobuf_shopping_shopping_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCatalogItemResponse); i {
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
		file_playground_protobuf_shopping_shopping_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCatalogItemsRequest); i {
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
		file_playground_protobuf_shopping_shopping_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCatalogItemsResponse); i {
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
		file_playground_protobuf_shopping_shopping_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_playground_protobuf_shopping_shopping_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
			RawDescriptor: file_playground_protobuf_shopping_shopping_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_playground_protobuf_shopping_shopping_service_proto_goTypes,
		DependencyIndexes: file_playground_protobuf_shopping_shopping_service_proto_depIdxs,
		MessageInfos:      file_playground_protobuf_shopping_shopping_service_proto_msgTypes,
	}.Build()
	File_playground_protobuf_shopping_shopping_service_proto = out.File
	file_playground_protobuf_shopping_shopping_service_proto_rawDesc = nil
	file_playground_protobuf_shopping_shopping_service_proto_goTypes = nil
	file_playground_protobuf_shopping_shopping_service_proto_depIdxs = nil
}
