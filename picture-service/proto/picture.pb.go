// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/picture.proto

package picture

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenreData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenreData) Reset() {
	*x = GenreData{}
	mi := &file_proto_picture_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenreData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenreData) ProtoMessage() {}

func (x *GenreData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenreData.ProtoReflect.Descriptor instead.
func (*GenreData) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{0}
}

func (x *GenreData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GenreData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenreData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GenreData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type AuthorData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName      string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	DateOfBirth   string                 `protobuf:"bytes,3,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	DateOfDeath   string                 `protobuf:"bytes,4,opt,name=date_of_death,json=dateOfDeath,proto3" json:"date_of_death,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthorData) Reset() {
	*x = AuthorData{}
	mi := &file_proto_picture_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorData) ProtoMessage() {}

func (x *AuthorData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorData.ProtoReflect.Descriptor instead.
func (*AuthorData) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthorData) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *AuthorData) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *AuthorData) GetDateOfDeath() string {
	if x != nil {
		return x.DateOfDeath
	}
	return ""
}

func (x *AuthorData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type PictureData struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DateOfPainting string                 `protobuf:"bytes,3,opt,name=date_of_painting,json=dateOfPainting,proto3" json:"date_of_painting,omitempty"`
	GenresIds      []string               `protobuf:"bytes,4,rep,name=genres_ids,json=genresIds,proto3" json:"genres_ids,omitempty"`
	AuthorsIds     []string               `protobuf:"bytes,5,rep,name=authors_ids,json=authorsIds,proto3" json:"authors_ids,omitempty"`
	Authors        []*AuthorData          `protobuf:"bytes,6,rep,name=authors,proto3" json:"authors,omitempty"`
	Genres         []*GenreData           `protobuf:"bytes,7,rep,name=genres,proto3" json:"genres,omitempty"`
	ExhibitionId   string                 `protobuf:"bytes,8,opt,name=exhibition_id,json=exhibitionId,proto3" json:"exhibition_id,omitempty"`
	Cost           float64                `protobuf:"fixed64,9,opt,name=cost,proto3" json:"cost,omitempty"`
	Location       string                 `protobuf:"bytes,10,opt,name=location,proto3" json:"location,omitempty"`
	CreatedAt      string                 `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *PictureData) Reset() {
	*x = PictureData{}
	mi := &file_proto_picture_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PictureData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PictureData) ProtoMessage() {}

func (x *PictureData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PictureData.ProtoReflect.Descriptor instead.
func (*PictureData) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{2}
}

func (x *PictureData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PictureData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PictureData) GetDateOfPainting() string {
	if x != nil {
		return x.DateOfPainting
	}
	return ""
}

func (x *PictureData) GetGenresIds() []string {
	if x != nil {
		return x.GenresIds
	}
	return nil
}

func (x *PictureData) GetAuthorsIds() []string {
	if x != nil {
		return x.AuthorsIds
	}
	return nil
}

func (x *PictureData) GetAuthors() []*AuthorData {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *PictureData) GetGenres() []*GenreData {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *PictureData) GetExhibitionId() string {
	if x != nil {
		return x.ExhibitionId
	}
	return ""
}

func (x *PictureData) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *PictureData) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *PictureData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetAllRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PageNumber    int32                  `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"` // номер страницы, начиная с 1
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	mi := &file_proto_picture_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *GetAllRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetAllResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pictures      []*PictureData         `protobuf:"bytes,1,rep,name=pictures,proto3" json:"pictures,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	TotalPages    int32                  `protobuf:"varint,3,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	mi := &file_proto_picture_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllResponse) GetPictures() []*PictureData {
	if x != nil {
		return x.Pictures
	}
	return nil
}

func (x *GetAllResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetAllResponse) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

type GetPictureRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PictureId     string                 `protobuf:"bytes,1,opt,name=picture_id,json=pictureId,proto3" json:"picture_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPictureRequest) Reset() {
	*x = GetPictureRequest{}
	mi := &file_proto_picture_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPictureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPictureRequest) ProtoMessage() {}

func (x *GetPictureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPictureRequest.ProtoReflect.Descriptor instead.
func (*GetPictureRequest) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{5}
}

func (x *GetPictureRequest) GetPictureId() string {
	if x != nil {
		return x.PictureId
	}
	return ""
}

type GetPictureResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Picture       *PictureData           `protobuf:"bytes,1,opt,name=picture,proto3" json:"picture,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPictureResponse) Reset() {
	*x = GetPictureResponse{}
	mi := &file_proto_picture_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPictureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPictureResponse) ProtoMessage() {}

func (x *GetPictureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPictureResponse.ProtoReflect.Descriptor instead.
func (*GetPictureResponse) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{6}
}

func (x *GetPictureResponse) GetPicture() *PictureData {
	if x != nil {
		return x.Picture
	}
	return nil
}

type CreatePictureRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Name           string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DateOfPainting string                 `protobuf:"bytes,2,opt,name=date_of_painting,json=dateOfPainting,proto3" json:"date_of_painting,omitempty"`
	AuthorsIds     []string               `protobuf:"bytes,3,rep,name=authors_ids,json=authorsIds,proto3" json:"authors_ids,omitempty"`
	GenresIds      []string               `protobuf:"bytes,4,rep,name=genres_ids,json=genresIds,proto3" json:"genres_ids,omitempty"`
	ExhibitionId   string                 `protobuf:"bytes,5,opt,name=exhibition_id,json=exhibitionId,proto3" json:"exhibition_id,omitempty"`
	Cost           float64                `protobuf:"fixed64,6,opt,name=cost,proto3" json:"cost,omitempty"`
	Location       string                 `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CreatePictureRequest) Reset() {
	*x = CreatePictureRequest{}
	mi := &file_proto_picture_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePictureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePictureRequest) ProtoMessage() {}

func (x *CreatePictureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePictureRequest.ProtoReflect.Descriptor instead.
func (*CreatePictureRequest) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{7}
}

func (x *CreatePictureRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePictureRequest) GetDateOfPainting() string {
	if x != nil {
		return x.DateOfPainting
	}
	return ""
}

func (x *CreatePictureRequest) GetAuthorsIds() []string {
	if x != nil {
		return x.AuthorsIds
	}
	return nil
}

func (x *CreatePictureRequest) GetGenresIds() []string {
	if x != nil {
		return x.GenresIds
	}
	return nil
}

func (x *CreatePictureRequest) GetExhibitionId() string {
	if x != nil {
		return x.ExhibitionId
	}
	return ""
}

func (x *CreatePictureRequest) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *CreatePictureRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type CreatePictureResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Picture       *PictureData           `protobuf:"bytes,1,opt,name=picture,proto3" json:"picture,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePictureResponse) Reset() {
	*x = CreatePictureResponse{}
	mi := &file_proto_picture_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePictureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePictureResponse) ProtoMessage() {}

func (x *CreatePictureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePictureResponse.ProtoReflect.Descriptor instead.
func (*CreatePictureResponse) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{8}
}

func (x *CreatePictureResponse) GetPicture() *PictureData {
	if x != nil {
		return x.Picture
	}
	return nil
}

func (x *CreatePictureResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdatePictureRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	PictureId      string                 `protobuf:"bytes,1,opt,name=picture_id,json=pictureId,proto3" json:"picture_id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DateOfPainting string                 `protobuf:"bytes,3,opt,name=date_of_painting,json=dateOfPainting,proto3" json:"date_of_painting,omitempty"`
	AuthorsIds     []string               `protobuf:"bytes,4,rep,name=authors_ids,json=authorsIds,proto3" json:"authors_ids,omitempty"`
	GenresIds      []string               `protobuf:"bytes,5,rep,name=genres_ids,json=genresIds,proto3" json:"genres_ids,omitempty"`
	ExhibitionId   string                 `protobuf:"bytes,6,opt,name=exhibition_id,json=exhibitionId,proto3" json:"exhibition_id,omitempty"`
	Cost           float64                `protobuf:"fixed64,7,opt,name=cost,proto3" json:"cost,omitempty"`
	Location       string                 `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	UserId         string                 `protobuf:"bytes,9,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UpdatePictureRequest) Reset() {
	*x = UpdatePictureRequest{}
	mi := &file_proto_picture_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePictureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePictureRequest) ProtoMessage() {}

func (x *UpdatePictureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePictureRequest.ProtoReflect.Descriptor instead.
func (*UpdatePictureRequest) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{9}
}

func (x *UpdatePictureRequest) GetPictureId() string {
	if x != nil {
		return x.PictureId
	}
	return ""
}

func (x *UpdatePictureRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePictureRequest) GetDateOfPainting() string {
	if x != nil {
		return x.DateOfPainting
	}
	return ""
}

func (x *UpdatePictureRequest) GetAuthorsIds() []string {
	if x != nil {
		return x.AuthorsIds
	}
	return nil
}

func (x *UpdatePictureRequest) GetGenresIds() []string {
	if x != nil {
		return x.GenresIds
	}
	return nil
}

func (x *UpdatePictureRequest) GetExhibitionId() string {
	if x != nil {
		return x.ExhibitionId
	}
	return ""
}

func (x *UpdatePictureRequest) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *UpdatePictureRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UpdatePictureRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdatePictureResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Picture       *PictureData           `protobuf:"bytes,1,opt,name=picture,proto3" json:"picture,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePictureResponse) Reset() {
	*x = UpdatePictureResponse{}
	mi := &file_proto_picture_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePictureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePictureResponse) ProtoMessage() {}

func (x *UpdatePictureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePictureResponse.ProtoReflect.Descriptor instead.
func (*UpdatePictureResponse) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{10}
}

func (x *UpdatePictureResponse) GetPicture() *PictureData {
	if x != nil {
		return x.Picture
	}
	return nil
}

type DeletePictureRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PictureId     string                 `protobuf:"bytes,1,opt,name=picture_id,json=pictureId,proto3" json:"picture_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePictureRequest) Reset() {
	*x = DeletePictureRequest{}
	mi := &file_proto_picture_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePictureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePictureRequest) ProtoMessage() {}

func (x *DeletePictureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePictureRequest.ProtoReflect.Descriptor instead.
func (*DeletePictureRequest) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{11}
}

func (x *DeletePictureRequest) GetPictureId() string {
	if x != nil {
		return x.PictureId
	}
	return ""
}

type DeletePictureResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePictureResponse) Reset() {
	*x = DeletePictureResponse{}
	mi := &file_proto_picture_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePictureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePictureResponse) ProtoMessage() {}

func (x *DeletePictureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePictureResponse.ProtoReflect.Descriptor instead.
func (*DeletePictureResponse) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{12}
}

func (x *DeletePictureResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ErrorResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Errors        map[string]string      `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	mi := &file_proto_picture_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_picture_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_proto_picture_proto_rawDescGZIP(), []int{13}
}

func (x *ErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorResponse) GetErrors() map[string]string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_proto_picture_proto protoreflect.FileDescriptor

const file_proto_picture_proto_rawDesc = "" +
	"\n" +
	"\x13proto/picture.proto\x12\apicture\"p\n" +
	"\tGenreData\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1d\n" +
	"\n" +
	"created_at\x18\x04 \x01(\tR\tcreatedAt\"\xa0\x01\n" +
	"\n" +
	"AuthorData\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tfull_name\x18\x02 \x01(\tR\bfullName\x12\"\n" +
	"\rdate_of_birth\x18\x03 \x01(\tR\vdateOfBirth\x12\"\n" +
	"\rdate_of_death\x18\x04 \x01(\tR\vdateOfDeath\x12\x1d\n" +
	"\n" +
	"created_at\x18\x05 \x01(\tR\tcreatedAt\"\xea\x02\n" +
	"\vPictureData\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12(\n" +
	"\x10date_of_painting\x18\x03 \x01(\tR\x0edateOfPainting\x12\x1d\n" +
	"\n" +
	"genres_ids\x18\x04 \x03(\tR\tgenresIds\x12\x1f\n" +
	"\vauthors_ids\x18\x05 \x03(\tR\n" +
	"authorsIds\x12-\n" +
	"\aauthors\x18\x06 \x03(\v2\x13.picture.AuthorDataR\aauthors\x12*\n" +
	"\x06genres\x18\a \x03(\v2\x12.picture.GenreDataR\x06genres\x12#\n" +
	"\rexhibition_id\x18\b \x01(\tR\fexhibitionId\x12\x12\n" +
	"\x04cost\x18\t \x01(\x01R\x04cost\x12\x1a\n" +
	"\blocation\x18\n" +
	" \x01(\tR\blocation\x12\x1d\n" +
	"\n" +
	"created_at\x18\v \x01(\tR\tcreatedAt\"M\n" +
	"\rGetAllRequest\x12\x1f\n" +
	"\vpage_number\x18\x01 \x01(\x05R\n" +
	"pageNumber\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x05R\bpageSize\"y\n" +
	"\x0eGetAllResponse\x120\n" +
	"\bpictures\x18\x01 \x03(\v2\x14.picture.PictureDataR\bpictures\x12\x14\n" +
	"\x05total\x18\x02 \x01(\x05R\x05total\x12\x1f\n" +
	"\vtotal_pages\x18\x03 \x01(\x05R\n" +
	"totalPages\"2\n" +
	"\x11GetPictureRequest\x12\x1d\n" +
	"\n" +
	"picture_id\x18\x01 \x01(\tR\tpictureId\"D\n" +
	"\x12GetPictureResponse\x12.\n" +
	"\apicture\x18\x01 \x01(\v2\x14.picture.PictureDataR\apicture\"\xe9\x01\n" +
	"\x14CreatePictureRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12(\n" +
	"\x10date_of_painting\x18\x02 \x01(\tR\x0edateOfPainting\x12\x1f\n" +
	"\vauthors_ids\x18\x03 \x03(\tR\n" +
	"authorsIds\x12\x1d\n" +
	"\n" +
	"genres_ids\x18\x04 \x03(\tR\tgenresIds\x12#\n" +
	"\rexhibition_id\x18\x05 \x01(\tR\fexhibitionId\x12\x12\n" +
	"\x04cost\x18\x06 \x01(\x01R\x04cost\x12\x1a\n" +
	"\blocation\x18\a \x01(\tR\blocation\"a\n" +
	"\x15CreatePictureResponse\x12.\n" +
	"\apicture\x18\x01 \x01(\v2\x14.picture.PictureDataR\apicture\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"\xa1\x02\n" +
	"\x14UpdatePictureRequest\x12\x1d\n" +
	"\n" +
	"picture_id\x18\x01 \x01(\tR\tpictureId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12(\n" +
	"\x10date_of_painting\x18\x03 \x01(\tR\x0edateOfPainting\x12\x1f\n" +
	"\vauthors_ids\x18\x04 \x03(\tR\n" +
	"authorsIds\x12\x1d\n" +
	"\n" +
	"genres_ids\x18\x05 \x03(\tR\tgenresIds\x12#\n" +
	"\rexhibition_id\x18\x06 \x01(\tR\fexhibitionId\x12\x12\n" +
	"\x04cost\x18\a \x01(\x01R\x04cost\x12\x1a\n" +
	"\blocation\x18\b \x01(\tR\blocation\x12\x17\n" +
	"\auser_id\x18\t \x01(\tR\x06userId\"G\n" +
	"\x15UpdatePictureResponse\x12.\n" +
	"\apicture\x18\x01 \x01(\v2\x14.picture.PictureDataR\apicture\"5\n" +
	"\x14DeletePictureRequest\x12\x1d\n" +
	"\n" +
	"picture_id\x18\x01 \x01(\tR\tpictureId\"1\n" +
	"\x15DeletePictureResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"\xa0\x01\n" +
	"\rErrorResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12:\n" +
	"\x06errors\x18\x02 \x03(\v2\".picture.ErrorResponse.ErrorsEntryR\x06errors\x1a9\n" +
	"\vErrorsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x012\x82\x03\n" +
	"\x0ePictureService\x12N\n" +
	"\rCreatePicture\x12\x1d.picture.CreatePictureRequest\x1a\x1e.picture.CreatePictureResponse\x129\n" +
	"\x06GetAll\x12\x16.picture.GetAllRequest\x1a\x17.picture.GetAllResponse\x12E\n" +
	"\n" +
	"GetPicture\x12\x1a.picture.GetPictureRequest\x1a\x1b.picture.GetPictureResponse\x12N\n" +
	"\rUpdatePicture\x12\x1d.picture.UpdatePictureRequest\x1a\x1e.picture.UpdatePictureResponse\x12N\n" +
	"\rDeletePicture\x12\x1d.picture.DeletePictureRequest\x1a\x1e.picture.DeletePictureResponseBOZMgithub.com/Ostap00034/siproject-beercut-backend/picture-service/proto;pictureb\x06proto3"

var (
	file_proto_picture_proto_rawDescOnce sync.Once
	file_proto_picture_proto_rawDescData []byte
)

func file_proto_picture_proto_rawDescGZIP() []byte {
	file_proto_picture_proto_rawDescOnce.Do(func() {
		file_proto_picture_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_picture_proto_rawDesc), len(file_proto_picture_proto_rawDesc)))
	})
	return file_proto_picture_proto_rawDescData
}

var file_proto_picture_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_picture_proto_goTypes = []any{
	(*GenreData)(nil),             // 0: picture.GenreData
	(*AuthorData)(nil),            // 1: picture.AuthorData
	(*PictureData)(nil),           // 2: picture.PictureData
	(*GetAllRequest)(nil),         // 3: picture.GetAllRequest
	(*GetAllResponse)(nil),        // 4: picture.GetAllResponse
	(*GetPictureRequest)(nil),     // 5: picture.GetPictureRequest
	(*GetPictureResponse)(nil),    // 6: picture.GetPictureResponse
	(*CreatePictureRequest)(nil),  // 7: picture.CreatePictureRequest
	(*CreatePictureResponse)(nil), // 8: picture.CreatePictureResponse
	(*UpdatePictureRequest)(nil),  // 9: picture.UpdatePictureRequest
	(*UpdatePictureResponse)(nil), // 10: picture.UpdatePictureResponse
	(*DeletePictureRequest)(nil),  // 11: picture.DeletePictureRequest
	(*DeletePictureResponse)(nil), // 12: picture.DeletePictureResponse
	(*ErrorResponse)(nil),         // 13: picture.ErrorResponse
	nil,                           // 14: picture.ErrorResponse.ErrorsEntry
}
var file_proto_picture_proto_depIdxs = []int32{
	1,  // 0: picture.PictureData.authors:type_name -> picture.AuthorData
	0,  // 1: picture.PictureData.genres:type_name -> picture.GenreData
	2,  // 2: picture.GetAllResponse.pictures:type_name -> picture.PictureData
	2,  // 3: picture.GetPictureResponse.picture:type_name -> picture.PictureData
	2,  // 4: picture.CreatePictureResponse.picture:type_name -> picture.PictureData
	2,  // 5: picture.UpdatePictureResponse.picture:type_name -> picture.PictureData
	14, // 6: picture.ErrorResponse.errors:type_name -> picture.ErrorResponse.ErrorsEntry
	7,  // 7: picture.PictureService.CreatePicture:input_type -> picture.CreatePictureRequest
	3,  // 8: picture.PictureService.GetAll:input_type -> picture.GetAllRequest
	5,  // 9: picture.PictureService.GetPicture:input_type -> picture.GetPictureRequest
	9,  // 10: picture.PictureService.UpdatePicture:input_type -> picture.UpdatePictureRequest
	11, // 11: picture.PictureService.DeletePicture:input_type -> picture.DeletePictureRequest
	8,  // 12: picture.PictureService.CreatePicture:output_type -> picture.CreatePictureResponse
	4,  // 13: picture.PictureService.GetAll:output_type -> picture.GetAllResponse
	6,  // 14: picture.PictureService.GetPicture:output_type -> picture.GetPictureResponse
	10, // 15: picture.PictureService.UpdatePicture:output_type -> picture.UpdatePictureResponse
	12, // 16: picture.PictureService.DeletePicture:output_type -> picture.DeletePictureResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_picture_proto_init() }
func file_proto_picture_proto_init() {
	if File_proto_picture_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_picture_proto_rawDesc), len(file_proto_picture_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_picture_proto_goTypes,
		DependencyIndexes: file_proto_picture_proto_depIdxs,
		MessageInfos:      file_proto_picture_proto_msgTypes,
	}.Build()
	File_proto_picture_proto = out.File
	file_proto_picture_proto_goTypes = nil
	file_proto_picture_proto_depIdxs = nil
}
