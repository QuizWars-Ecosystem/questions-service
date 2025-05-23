// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: external/questions/v1/questions.proto

package questionsv1

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

// *
// Represents a light filter for getting desired questions
type GetQuestionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Difficulty    Difficulty             `protobuf:"varint,1,opt,name=difficulty,proto3,enum=questionsservice.v1.Difficulty" json:"difficulty,omitempty"` // Desired difficulty
	Language      string                 `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`                                          // Desired question language
	CategoryId    int32                  `protobuf:"varint,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`                   // Desired question category id
	Amount        int32                  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`                                             // Amount of desired questions
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetQuestionsRequest) Reset() {
	*x = GetQuestionsRequest{}
	mi := &file_external_questions_v1_questions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuestionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionsRequest) ProtoMessage() {}

func (x *GetQuestionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_questions_v1_questions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionsRequest.ProtoReflect.Descriptor instead.
func (*GetQuestionsRequest) Descriptor() ([]byte, []int) {
	return file_external_questions_v1_questions_proto_rawDescGZIP(), []int{0}
}

func (x *GetQuestionsRequest) GetDifficulty() Difficulty {
	if x != nil {
		return x.Difficulty
	}
	return Difficulty_DIFFICULTY_UNSPECIFIED
}

func (x *GetQuestionsRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *GetQuestionsRequest) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *GetQuestionsRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// *
// Represents a heavy filtered request for getting questions
type GetQuestionBatchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Types         []Type                 `protobuf:"varint,1,rep,packed,name=types,proto3,enum=questionsservice.v1.Type" json:"types,omitempty"`                     // List of desired question's types
	Sources       []Source               `protobuf:"varint,2,rep,packed,name=sources,proto3,enum=questionsservice.v1.Source" json:"sources,omitempty"`               // List of desired questions' sources
	Difficulties  []Difficulty           `protobuf:"varint,3,rep,packed,name=difficulties,proto3,enum=questionsservice.v1.Difficulty" json:"difficulties,omitempty"` // List of desired question's difficulties
	CategoriesIds []int32                `protobuf:"varint,4,rep,packed,name=categories_ids,json=categoriesIds,proto3" json:"categories_ids,omitempty"`              // List of desired question's categories ids
	Language      string                 `protobuf:"bytes,5,opt,name=language,proto3" json:"language,omitempty"`                                                     // Language of requesting question
	Amount        int32                  `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`                                                        // Amount of desired questions
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetQuestionBatchRequest) Reset() {
	*x = GetQuestionBatchRequest{}
	mi := &file_external_questions_v1_questions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuestionBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionBatchRequest) ProtoMessage() {}

func (x *GetQuestionBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_questions_v1_questions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionBatchRequest.ProtoReflect.Descriptor instead.
func (*GetQuestionBatchRequest) Descriptor() ([]byte, []int) {
	return file_external_questions_v1_questions_proto_rawDescGZIP(), []int{1}
}

func (x *GetQuestionBatchRequest) GetTypes() []Type {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *GetQuestionBatchRequest) GetSources() []Source {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *GetQuestionBatchRequest) GetDifficulties() []Difficulty {
	if x != nil {
		return x.Difficulties
	}
	return nil
}

func (x *GetQuestionBatchRequest) GetCategoriesIds() []int32 {
	if x != nil {
		return x.CategoriesIds
	}
	return nil
}

func (x *GetQuestionBatchRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *GetQuestionBatchRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// *
// Represents of requested, filtered questions list
type QuestionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Questions     []*Question            `protobuf:"bytes,1,rep,name=questions,proto3" json:"questions,omitempty"` // List of resulted questions
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QuestionsResponse) Reset() {
	*x = QuestionsResponse{}
	mi := &file_external_questions_v1_questions_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QuestionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionsResponse) ProtoMessage() {}

func (x *QuestionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_questions_v1_questions_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionsResponse.ProtoReflect.Descriptor instead.
func (*QuestionsResponse) Descriptor() ([]byte, []int) {
	return file_external_questions_v1_questions_proto_rawDescGZIP(), []int{2}
}

func (x *QuestionsResponse) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

var File_external_questions_v1_questions_proto protoreflect.FileDescriptor

var file_external_questions_v1_questions_proto_rawDesc = string([]byte{
	0x0a, 0x25, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xab, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x52, 0x0a, 0x64,
	0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa1,
	0x02, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x07, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x12, 0x43, 0x0a, 0x0c, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x52, 0x0c, 0x64, 0x69, 0x66, 0x66, 0x69,
	0x63, 0x75, 0x6c, 0x74, 0x69, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x50, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x32, 0xde, 0x01, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x2e, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x2c, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_external_questions_v1_questions_proto_rawDescOnce sync.Once
	file_external_questions_v1_questions_proto_rawDescData []byte
)

func file_external_questions_v1_questions_proto_rawDescGZIP() []byte {
	file_external_questions_v1_questions_proto_rawDescOnce.Do(func() {
		file_external_questions_v1_questions_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_external_questions_v1_questions_proto_rawDesc), len(file_external_questions_v1_questions_proto_rawDesc)))
	})
	return file_external_questions_v1_questions_proto_rawDescData
}

var file_external_questions_v1_questions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_external_questions_v1_questions_proto_goTypes = []any{
	(*GetQuestionsRequest)(nil),     // 0: questionsservice.v1.GetQuestionsRequest
	(*GetQuestionBatchRequest)(nil), // 1: questionsservice.v1.GetQuestionBatchRequest
	(*QuestionsResponse)(nil),       // 2: questionsservice.v1.QuestionsResponse
	(Difficulty)(0),                 // 3: questionsservice.v1.Difficulty
	(Type)(0),                       // 4: questionsservice.v1.Type
	(Source)(0),                     // 5: questionsservice.v1.Source
	(*Question)(nil),                // 6: questionsservice.v1.Question
}
var file_external_questions_v1_questions_proto_depIdxs = []int32{
	3, // 0: questionsservice.v1.GetQuestionsRequest.difficulty:type_name -> questionsservice.v1.Difficulty
	4, // 1: questionsservice.v1.GetQuestionBatchRequest.types:type_name -> questionsservice.v1.Type
	5, // 2: questionsservice.v1.GetQuestionBatchRequest.sources:type_name -> questionsservice.v1.Source
	3, // 3: questionsservice.v1.GetQuestionBatchRequest.difficulties:type_name -> questionsservice.v1.Difficulty
	6, // 4: questionsservice.v1.QuestionsResponse.questions:type_name -> questionsservice.v1.Question
	0, // 5: questionsservice.v1.QuestionsService.GetQuestions:input_type -> questionsservice.v1.GetQuestionsRequest
	1, // 6: questionsservice.v1.QuestionsService.GetQuestionBatch:input_type -> questionsservice.v1.GetQuestionBatchRequest
	2, // 7: questionsservice.v1.QuestionsService.GetQuestions:output_type -> questionsservice.v1.QuestionsResponse
	2, // 8: questionsservice.v1.QuestionsService.GetQuestionBatch:output_type -> questionsservice.v1.QuestionsResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_external_questions_v1_questions_proto_init() }
func file_external_questions_v1_questions_proto_init() {
	if File_external_questions_v1_questions_proto != nil {
		return
	}
	file_external_questions_v1_shared_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_external_questions_v1_questions_proto_rawDesc), len(file_external_questions_v1_questions_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_questions_v1_questions_proto_goTypes,
		DependencyIndexes: file_external_questions_v1_questions_proto_depIdxs,
		MessageInfos:      file_external_questions_v1_questions_proto_msgTypes,
	}.Build()
	File_external_questions_v1_questions_proto = out.File
	file_external_questions_v1_questions_proto_goTypes = nil
	file_external_questions_v1_questions_proto_depIdxs = nil
}
