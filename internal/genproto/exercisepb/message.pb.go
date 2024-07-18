// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: exercisepb/message.proto

package exercisepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TranslatedLanguages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vietnamese string `protobuf:"bytes,1,opt,name=vietnamese,proto3" json:"vietnamese,omitempty"`
}

func (x *TranslatedLanguages) Reset() {
	*x = TranslatedLanguages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslatedLanguages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslatedLanguages) ProtoMessage() {}

func (x *TranslatedLanguages) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslatedLanguages.ProtoReflect.Descriptor instead.
func (*TranslatedLanguages) Descriptor() ([]byte, []int) {
	return file_exercisepb_message_proto_rawDescGZIP(), []int{0}
}

func (x *TranslatedLanguages) GetVietnamese() string {
	if x != nil {
		return x.Vietnamese
	}
	return ""
}

type UserExercise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Audio         string                 `protobuf:"bytes,2,opt,name=audio,proto3" json:"audio,omitempty"`
	Level         string                 `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Translated    string                 `protobuf:"bytes,5,opt,name=translated,proto3" json:"translated,omitempty"`
	Vocabulary    string                 `protobuf:"bytes,6,opt,name=vocabulary,proto3" json:"vocabulary,omitempty"`
	CorrectAnswer string                 `protobuf:"bytes,7,opt,name=correctAnswer,proto3" json:"correctAnswer,omitempty"`
	Options       []string               `protobuf:"bytes,8,rep,name=options,proto3" json:"options,omitempty"`
	CorrectStreak int32                  `protobuf:"varint,9,opt,name=correctStreak,proto3" json:"correctStreak,omitempty"`
	IsFavorite    bool                   `protobuf:"varint,10,opt,name=isFavorite,proto3" json:"isFavorite,omitempty"`
	IsMastered    bool                   `protobuf:"varint,11,opt,name=isMastered,proto3" json:"isMastered,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	NextReviewAt  *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=nextReviewAt,proto3" json:"nextReviewAt,omitempty"`
}

func (x *UserExercise) Reset() {
	*x = UserExercise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserExercise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserExercise) ProtoMessage() {}

func (x *UserExercise) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserExercise.ProtoReflect.Descriptor instead.
func (*UserExercise) Descriptor() ([]byte, []int) {
	return file_exercisepb_message_proto_rawDescGZIP(), []int{1}
}

func (x *UserExercise) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserExercise) GetAudio() string {
	if x != nil {
		return x.Audio
	}
	return ""
}

func (x *UserExercise) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *UserExercise) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UserExercise) GetTranslated() string {
	if x != nil {
		return x.Translated
	}
	return ""
}

func (x *UserExercise) GetVocabulary() string {
	if x != nil {
		return x.Vocabulary
	}
	return ""
}

func (x *UserExercise) GetCorrectAnswer() string {
	if x != nil {
		return x.CorrectAnswer
	}
	return ""
}

func (x *UserExercise) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *UserExercise) GetCorrectStreak() int32 {
	if x != nil {
		return x.CorrectStreak
	}
	return 0
}

func (x *UserExercise) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *UserExercise) GetIsMastered() bool {
	if x != nil {
		return x.IsMastered
	}
	return false
}

func (x *UserExercise) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UserExercise) GetNextReviewAt() *timestamppb.Timestamp {
	if x != nil {
		return x.NextReviewAt
	}
	return nil
}

type ExerciseCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slug            string `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Translated      string `protobuf:"bytes,4,opt,name=translated,proto3" json:"translated,omitempty"`
	Image           string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	StatsExercises  int32  `protobuf:"varint,7,opt,name=statsExercises,proto3" json:"statsExercises,omitempty"`
	StatsInteracted int32  `protobuf:"varint,8,opt,name=statsInteracted,proto3" json:"statsInteracted,omitempty"`
}

func (x *ExerciseCollection) Reset() {
	*x = ExerciseCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExerciseCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExerciseCollection) ProtoMessage() {}

func (x *ExerciseCollection) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExerciseCollection.ProtoReflect.Descriptor instead.
func (*ExerciseCollection) Descriptor() ([]byte, []int) {
	return file_exercisepb_message_proto_rawDescGZIP(), []int{2}
}

func (x *ExerciseCollection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExerciseCollection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExerciseCollection) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *ExerciseCollection) GetTranslated() string {
	if x != nil {
		return x.Translated
	}
	return ""
}

func (x *ExerciseCollection) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ExerciseCollection) GetStatsExercises() int32 {
	if x != nil {
		return x.StatsExercises
	}
	return 0
}

func (x *ExerciseCollection) GetStatsInteracted() int32 {
	if x != nil {
		return x.StatsInteracted
	}
	return 0
}

type UserAggregatedExercise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date     string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Exercise int64  `protobuf:"varint,2,opt,name=exercise,proto3" json:"exercise,omitempty"`
}

func (x *UserAggregatedExercise) Reset() {
	*x = UserAggregatedExercise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercisepb_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAggregatedExercise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAggregatedExercise) ProtoMessage() {}

func (x *UserAggregatedExercise) ProtoReflect() protoreflect.Message {
	mi := &file_exercisepb_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAggregatedExercise.ProtoReflect.Descriptor instead.
func (*UserAggregatedExercise) Descriptor() ([]byte, []int) {
	return file_exercisepb_message_proto_rawDescGZIP(), []int{3}
}

func (x *UserAggregatedExercise) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *UserAggregatedExercise) GetExercise() int64 {
	if x != nil {
		return x.Exercise
	}
	return 0
}

var File_exercisepb_message_proto protoreflect.FileDescriptor

var file_exercisepb_message_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x76, 0x69, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x76, 0x69, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x22, 0xc4,
	0x03, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c,
	0x61, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x6f, 0x63, 0x61, 0x62,
	0x75, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x63, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x73, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x73, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x41, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x41, 0x74, 0x22, 0xd4, 0x01, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x6c, 0x75, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x22, 0x48, 0x0a, 0x16,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x42, 0xa4, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x65,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x76, 0x6f, 0x63, 0x61, 0x62,
	0x2d, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68,
	0x2d, 0x68, 0x75, 0x62, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x2f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x0a,
	0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0xca, 0x02, 0x0a, 0x45, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0xe2, 0x02, 0x16, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0a, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exercisepb_message_proto_rawDescOnce sync.Once
	file_exercisepb_message_proto_rawDescData = file_exercisepb_message_proto_rawDesc
)

func file_exercisepb_message_proto_rawDescGZIP() []byte {
	file_exercisepb_message_proto_rawDescOnce.Do(func() {
		file_exercisepb_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_exercisepb_message_proto_rawDescData)
	})
	return file_exercisepb_message_proto_rawDescData
}

var file_exercisepb_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_exercisepb_message_proto_goTypes = []interface{}{
	(*TranslatedLanguages)(nil),    // 0: exercisepb.TranslatedLanguages
	(*UserExercise)(nil),           // 1: exercisepb.UserExercise
	(*ExerciseCollection)(nil),     // 2: exercisepb.ExerciseCollection
	(*UserAggregatedExercise)(nil), // 3: exercisepb.UserAggregatedExercise
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
}
var file_exercisepb_message_proto_depIdxs = []int32{
	4, // 0: exercisepb.UserExercise.updatedAt:type_name -> google.protobuf.Timestamp
	4, // 1: exercisepb.UserExercise.nextReviewAt:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_exercisepb_message_proto_init() }
func file_exercisepb_message_proto_init() {
	if File_exercisepb_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exercisepb_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslatedLanguages); i {
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
		file_exercisepb_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserExercise); i {
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
		file_exercisepb_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExerciseCollection); i {
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
		file_exercisepb_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAggregatedExercise); i {
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
			RawDescriptor: file_exercisepb_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_exercisepb_message_proto_goTypes,
		DependencyIndexes: file_exercisepb_message_proto_depIdxs,
		MessageInfos:      file_exercisepb_message_proto_msgTypes,
	}.Build()
	File_exercisepb_message_proto = out.File
	file_exercisepb_message_proto_rawDesc = nil
	file_exercisepb_message_proto_goTypes = nil
	file_exercisepb_message_proto_depIdxs = nil
}
