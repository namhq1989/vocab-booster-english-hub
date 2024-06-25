// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: exercisepb/hub.proto

package exercisepb

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

type NewExerciseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VocabularyExampleId string               `protobuf:"bytes,1,opt,name=vocabularyExampleId,proto3" json:"vocabularyExampleId,omitempty"`
	Level               string               `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	Content             string               `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Translated          *TranslatedLanguages `protobuf:"bytes,4,opt,name=translated,proto3" json:"translated,omitempty"`
	Vocabulary          string               `protobuf:"bytes,5,opt,name=vocabulary,proto3" json:"vocabulary,omitempty"`
	CorrectAnswer       string               `protobuf:"bytes,6,opt,name=correctAnswer,proto3" json:"correctAnswer,omitempty"`
	Options             []string             `protobuf:"bytes,7,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *NewExerciseRequest) Reset() {
	*x = NewExerciseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_hub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewExerciseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewExerciseRequest) ProtoMessage() {}

func (x *NewExerciseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_hub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewExerciseRequest.ProtoReflect.Descriptor instead.
func (*NewExerciseRequest) Descriptor() ([]byte, []int) {
	return file_exercisepb_hub_proto_rawDescGZIP(), []int{0}
}

func (x *NewExerciseRequest) GetVocabularyExampleId() string {
	if x != nil {
		return x.VocabularyExampleId
	}
	return ""
}

func (x *NewExerciseRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *NewExerciseRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NewExerciseRequest) GetTranslated() *TranslatedLanguages {
	if x != nil {
		return x.Translated
	}
	return nil
}

func (x *NewExerciseRequest) GetVocabulary() string {
	if x != nil {
		return x.Vocabulary
	}
	return ""
}

func (x *NewExerciseRequest) GetCorrectAnswer() string {
	if x != nil {
		return x.CorrectAnswer
	}
	return ""
}

func (x *NewExerciseRequest) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

type NewExerciseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewExerciseResponse) Reset() {
	*x = NewExerciseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_hub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewExerciseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewExerciseResponse) ProtoMessage() {}

func (x *NewExerciseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_hub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewExerciseResponse.ProtoReflect.Descriptor instead.
func (*NewExerciseResponse) Descriptor() ([]byte, []int) {
	return file_exercisepb_hub_proto_rawDescGZIP(), []int{1}
}

func (x *NewExerciseResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateExerciseAudioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VocabularyExampleId string `protobuf:"bytes,1,opt,name=vocabularyExampleId,proto3" json:"vocabularyExampleId,omitempty"`
	Audio               string `protobuf:"bytes,2,opt,name=audio,proto3" json:"audio,omitempty"`
}

func (x *UpdateExerciseAudioRequest) Reset() {
	*x = UpdateExerciseAudioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_hub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExerciseAudioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExerciseAudioRequest) ProtoMessage() {}

func (x *UpdateExerciseAudioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_hub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExerciseAudioRequest.ProtoReflect.Descriptor instead.
func (*UpdateExerciseAudioRequest) Descriptor() ([]byte, []int) {
	return file_exercisepb_hub_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateExerciseAudioRequest) GetVocabularyExampleId() string {
	if x != nil {
		return x.VocabularyExampleId
	}
	return ""
}

func (x *UpdateExerciseAudioRequest) GetAudio() string {
	if x != nil {
		return x.Audio
	}
	return ""
}

type UpdateExerciseAudioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateExerciseAudioResponse) Reset() {
	*x = UpdateExerciseAudioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_hub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExerciseAudioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExerciseAudioResponse) ProtoMessage() {}

func (x *UpdateExerciseAudioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_hub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExerciseAudioResponse.ProtoReflect.Descriptor instead.
func (*UpdateExerciseAudioResponse) Descriptor() ([]byte, []int) {
	return file_exercisepb_hub_proto_rawDescGZIP(), []int{3}
}

type AnswerExerciseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExerciseId string `protobuf:"bytes,1,opt,name=exerciseId,proto3" json:"exerciseId,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	IsCorrect  bool   `protobuf:"varint,3,opt,name=isCorrect,proto3" json:"isCorrect,omitempty"`
}

func (x *AnswerExerciseRequest) Reset() {
	*x = AnswerExerciseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_hub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswerExerciseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswerExerciseRequest) ProtoMessage() {}

func (x *AnswerExerciseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_hub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswerExerciseRequest.ProtoReflect.Descriptor instead.
func (*AnswerExerciseRequest) Descriptor() ([]byte, []int) {
	return file_exercisepb_hub_proto_rawDescGZIP(), []int{4}
}

func (x *AnswerExerciseRequest) GetExerciseId() string {
	if x != nil {
		return x.ExerciseId
	}
	return ""
}

func (x *AnswerExerciseRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AnswerExerciseRequest) GetIsCorrect() bool {
	if x != nil {
		return x.IsCorrect
	}
	return false
}

type AnswerExerciseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnswerExerciseResponse) Reset() {
	*x = AnswerExerciseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_hub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswerExerciseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswerExerciseResponse) ProtoMessage() {}

func (x *AnswerExerciseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_hub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswerExerciseResponse.ProtoReflect.Descriptor instead.
func (*AnswerExerciseResponse) Descriptor() ([]byte, []int) {
	return file_exercisepb_hub_proto_rawDescGZIP(), []int{5}
}

type GetUserExercisesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Level  string `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	Lang   string `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *GetUserExercisesRequest) Reset() {
	*x = GetUserExercisesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_hub_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserExercisesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserExercisesRequest) ProtoMessage() {}

func (x *GetUserExercisesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_hub_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserExercisesRequest.ProtoReflect.Descriptor instead.
func (*GetUserExercisesRequest) Descriptor() ([]byte, []int) {
	return file_exercisepb_hub_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserExercisesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserExercisesRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *GetUserExercisesRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

type GetUserExercisesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercises []*UserExercise `protobuf:"bytes,1,rep,name=exercises,proto3" json:"exercises,omitempty"`
}

func (x *GetUserExercisesResponse) Reset() {
	*x = GetUserExercisesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_hub_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserExercisesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserExercisesResponse) ProtoMessage() {}

func (x *GetUserExercisesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_hub_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserExercisesResponse.ProtoReflect.Descriptor instead.
func (*GetUserExercisesResponse) Descriptor() ([]byte, []int) {
	return file_exercisepb_hub_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserExercisesResponse) GetExercises() []*UserExercise {
	if x != nil {
		return x.Exercises
	}
	return nil
}

type CountUserReadyToReviewExercisesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CountUserReadyToReviewExercisesRequest) Reset() {
	*x = CountUserReadyToReviewExercisesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_hub_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUserReadyToReviewExercisesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUserReadyToReviewExercisesRequest) ProtoMessage() {}

func (x *CountUserReadyToReviewExercisesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_hub_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUserReadyToReviewExercisesRequest.ProtoReflect.Descriptor instead.
func (*CountUserReadyToReviewExercisesRequest) Descriptor() ([]byte, []int) {
	return file_exercisepb_hub_proto_rawDescGZIP(), []int{8}
}

func (x *CountUserReadyToReviewExercisesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CountUserReadyToReviewExercisesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *CountUserReadyToReviewExercisesResponse) Reset() {
	*x = CountUserReadyToReviewExercisesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_hub_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUserReadyToReviewExercisesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUserReadyToReviewExercisesResponse) ProtoMessage() {}

func (x *CountUserReadyToReviewExercisesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_hub_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUserReadyToReviewExercisesResponse.ProtoReflect.Descriptor instead.
func (*CountUserReadyToReviewExercisesResponse) Descriptor() ([]byte, []int) {
	return file_exercisepb_hub_proto_rawDescGZIP(), []int{9}
}

func (x *CountUserReadyToReviewExercisesResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetUserReadyToReviewExercisesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Level  string `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	Lang   string `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *GetUserReadyToReviewExercisesRequest) Reset() {
	*x = GetUserReadyToReviewExercisesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_hub_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReadyToReviewExercisesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReadyToReviewExercisesRequest) ProtoMessage() {}

func (x *GetUserReadyToReviewExercisesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_hub_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReadyToReviewExercisesRequest.ProtoReflect.Descriptor instead.
func (*GetUserReadyToReviewExercisesRequest) Descriptor() ([]byte, []int) {
	return file_exercisepb_hub_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserReadyToReviewExercisesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserReadyToReviewExercisesRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *GetUserReadyToReviewExercisesRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

type GetUserReadyToReviewExercisesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercises []*UserExercise `protobuf:"bytes,1,rep,name=exercises,proto3" json:"exercises,omitempty"`
}

func (x *GetUserReadyToReviewExercisesResponse) Reset() {
	*x = GetUserReadyToReviewExercisesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_hub_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReadyToReviewExercisesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReadyToReviewExercisesResponse) ProtoMessage() {}

func (x *GetUserReadyToReviewExercisesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_hub_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReadyToReviewExercisesResponse.ProtoReflect.Descriptor instead.
func (*GetUserReadyToReviewExercisesResponse) Descriptor() ([]byte, []int) {
	return file_exercisepb_hub_proto_rawDescGZIP(), []int{11}
}

func (x *GetUserReadyToReviewExercisesResponse) GetExercises() []*UserExercise {
	if x != nil {
		return x.Exercises
	}
	return nil
}

var File_exercisepb_hub_proto protoreflect.FileDescriptor

var file_exercisepb_hub_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2f, 0x68, 0x75, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x70, 0x62, 0x1a, 0x18, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a,
	0x12, 0x4e, 0x65, 0x77, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72,
	0x79, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75,
	0x6c, 0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x6f, 0x63, 0x61,
	0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x45, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x64, 0x0a,
	0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x76,
	0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75,
	0x6c, 0x61, 0x72, 0x79, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x75,
	0x64, 0x69, 0x6f, 0x22, 0x1d, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x6d, 0x0a, 0x15, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x45, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x22, 0x18, 0x0a, 0x16, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x45, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x52, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x52, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x22, 0x40, 0x0a, 0x26,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3f,
	0x0a, 0x27, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79,
	0x54, 0x6f, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x68, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54,
	0x6f, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x5f, 0x0a, 0x25, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52,
	0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x32, 0xa1, 0x05, 0x0a, 0x0f, 0x45,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50,
	0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x1e, 0x2e,
	0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x59, 0x0a, 0x0e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x12, 0x21, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2e,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x70, 0x62, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x41, 0x75, 0x64,
	0x69, 0x6f, 0x12, 0x26, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x41, 0x75,
	0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x65, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8c, 0x01, 0x0a, 0x1f, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x12, 0x32, 0x2e, 0x65, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33,
	0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x86, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x12, 0x30, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x54, 0x6f, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x65, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x61, 0x64, 0x79, 0x54, 0x6f, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x45, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xa0,
	0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70,
	0x62, 0x42, 0x08, 0x48, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x76,
	0x6f, 0x63, 0x61, 0x62, 0x2d, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x65, 0x6e, 0x67,
	0x6c, 0x69, 0x73, 0x68, 0x2d, 0x68, 0x75, 0x62, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x2f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62,
	0x2f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x45, 0x58,
	0x58, 0xaa, 0x02, 0x0a, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0xca, 0x02,
	0x0a, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0xe2, 0x02, 0x16, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exercisepb_hub_proto_rawDescOnce sync.Once
	file_exercisepb_hub_proto_rawDescData = file_exercisepb_hub_proto_rawDesc
)

func file_exercisepb_hub_proto_rawDescGZIP() []byte {
	file_exercisepb_hub_proto_rawDescOnce.Do(func() {
		file_exercisepb_hub_proto_rawDescData = protoimpl.X.CompressGZIP(file_exercisepb_hub_proto_rawDescData)
	})
	return file_exercisepb_hub_proto_rawDescData
}

var file_exercisepb_hub_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_exercisepb_hub_proto_goTypes = []interface{}{
	(*NewExerciseRequest)(nil),                      // 0: exercisepb.NewExerciseRequest
	(*NewExerciseResponse)(nil),                     // 1: exercisepb.NewExerciseResponse
	(*UpdateExerciseAudioRequest)(nil),              // 2: exercisepb.UpdateExerciseAudioRequest
	(*UpdateExerciseAudioResponse)(nil),             // 3: exercisepb.UpdateExerciseAudioResponse
	(*AnswerExerciseRequest)(nil),                   // 4: exercisepb.AnswerExerciseRequest
	(*AnswerExerciseResponse)(nil),                  // 5: exercisepb.AnswerExerciseResponse
	(*GetUserExercisesRequest)(nil),                 // 6: exercisepb.GetUserExercisesRequest
	(*GetUserExercisesResponse)(nil),                // 7: exercisepb.GetUserExercisesResponse
	(*CountUserReadyToReviewExercisesRequest)(nil),  // 8: exercisepb.CountUserReadyToReviewExercisesRequest
	(*CountUserReadyToReviewExercisesResponse)(nil), // 9: exercisepb.CountUserReadyToReviewExercisesResponse
	(*GetUserReadyToReviewExercisesRequest)(nil),    // 10: exercisepb.GetUserReadyToReviewExercisesRequest
	(*GetUserReadyToReviewExercisesResponse)(nil),   // 11: exercisepb.GetUserReadyToReviewExercisesResponse
	(*TranslatedLanguages)(nil),                     // 12: exercisepb.TranslatedLanguages
	(*UserExercise)(nil),                            // 13: exercisepb.UserExercise
}
var file_exercisepb_hub_proto_depIdxs = []int32{
	12, // 0: exercisepb.NewExerciseRequest.translated:type_name -> exercisepb.TranslatedLanguages
	13, // 1: exercisepb.GetUserExercisesResponse.exercises:type_name -> exercisepb.UserExercise
	13, // 2: exercisepb.GetUserReadyToReviewExercisesResponse.exercises:type_name -> exercisepb.UserExercise
	0,  // 3: exercisepb.ExerciseService.NewExercise:input_type -> exercisepb.NewExerciseRequest
	4,  // 4: exercisepb.ExerciseService.AnswerExercise:input_type -> exercisepb.AnswerExerciseRequest
	2,  // 5: exercisepb.ExerciseService.UpdateExerciseAudio:input_type -> exercisepb.UpdateExerciseAudioRequest
	6,  // 6: exercisepb.ExerciseService.GetUserExercises:input_type -> exercisepb.GetUserExercisesRequest
	8,  // 7: exercisepb.ExerciseService.CountUserReadyToReviewExercises:input_type -> exercisepb.CountUserReadyToReviewExercisesRequest
	10, // 8: exercisepb.ExerciseService.GetUserReadyToReviewExercises:input_type -> exercisepb.GetUserReadyToReviewExercisesRequest
	1,  // 9: exercisepb.ExerciseService.NewExercise:output_type -> exercisepb.NewExerciseResponse
	5,  // 10: exercisepb.ExerciseService.AnswerExercise:output_type -> exercisepb.AnswerExerciseResponse
	3,  // 11: exercisepb.ExerciseService.UpdateExerciseAudio:output_type -> exercisepb.UpdateExerciseAudioResponse
	7,  // 12: exercisepb.ExerciseService.GetUserExercises:output_type -> exercisepb.GetUserExercisesResponse
	9,  // 13: exercisepb.ExerciseService.CountUserReadyToReviewExercises:output_type -> exercisepb.CountUserReadyToReviewExercisesResponse
	11, // 14: exercisepb.ExerciseService.GetUserReadyToReviewExercises:output_type -> exercisepb.GetUserReadyToReviewExercisesResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_exercisepb_hub_proto_init() }
func file_exercisepb_hub_proto_init() {
	if File_exercisepb_hub_proto != nil {
		return
	}
	file_exercisepb_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_exercisepb_hub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewExerciseRequest); i {
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
		file_exercisepb_hub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewExerciseResponse); i {
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
		file_exercisepb_hub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateExerciseAudioRequest); i {
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
		file_exercisepb_hub_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateExerciseAudioResponse); i {
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
		file_exercisepb_hub_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswerExerciseRequest); i {
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
		file_exercisepb_hub_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswerExerciseResponse); i {
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
		file_exercisepb_hub_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserExercisesRequest); i {
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
		file_exercisepb_hub_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserExercisesResponse); i {
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
		file_exercisepb_hub_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUserReadyToReviewExercisesRequest); i {
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
		file_exercisepb_hub_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUserReadyToReviewExercisesResponse); i {
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
		file_exercisepb_hub_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReadyToReviewExercisesRequest); i {
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
		file_exercisepb_hub_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReadyToReviewExercisesResponse); i {
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
			RawDescriptor: file_exercisepb_hub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exercisepb_hub_proto_goTypes,
		DependencyIndexes: file_exercisepb_hub_proto_depIdxs,
		MessageInfos:      file_exercisepb_hub_proto_msgTypes,
	}.Build()
	File_exercisepb_hub_proto = out.File
	file_exercisepb_hub_proto_rawDesc = nil
	file_exercisepb_hub_proto_goTypes = nil
	file_exercisepb_hub_proto_depIdxs = nil
}
