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

type GetQuestionBatchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Categories    []string               `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	Difficulty    Difficulty             `protobuf:"varint,2,opt,name=difficulty,proto3,enum=questions.v1.Difficulty" json:"difficulty,omitempty"`
	Language      string                 `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
	Count         int32                  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetQuestionBatchRequest) Reset() {
	*x = GetQuestionBatchRequest{}
	mi := &file_external_questions_v1_questions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuestionBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionBatchRequest) ProtoMessage() {}

func (x *GetQuestionBatchRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetQuestionBatchRequest.ProtoReflect.Descriptor instead.
func (*GetQuestionBatchRequest) Descriptor() ([]byte, []int) {
	return file_external_questions_v1_questions_proto_rawDescGZIP(), []int{0}
}

func (x *GetQuestionBatchRequest) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *GetQuestionBatchRequest) GetDifficulty() Difficulty {
	if x != nil {
		return x.Difficulty
	}
	return Difficulty_DIFFICULTY_UNSPECIFIED
}

func (x *GetQuestionBatchRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *GetQuestionBatchRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetQuestionsBatchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Questions     []*Question            `protobuf:"bytes,1,rep,name=questions,proto3" json:"questions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetQuestionsBatchResponse) Reset() {
	*x = GetQuestionsBatchResponse{}
	mi := &file_external_questions_v1_questions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuestionsBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionsBatchResponse) ProtoMessage() {}

func (x *GetQuestionsBatchResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetQuestionsBatchResponse.ProtoReflect.Descriptor instead.
func (*GetQuestionsBatchResponse) Descriptor() ([]byte, []int) {
	return file_external_questions_v1_questions_proto_rawDescGZIP(), []int{1}
}

func (x *GetQuestionsBatchResponse) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

var File_external_questions_v1_questions_proto protoreflect.FileDescriptor

var file_external_questions_v1_questions_proto_rawDesc = string([]byte{
	0x0a, 0x25, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75,
	0x6c, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75,
	0x6c, 0x74, 0x79, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x51, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x32, 0x76, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x25, 0x2e, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_external_questions_v1_questions_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_external_questions_v1_questions_proto_goTypes = []any{
	(*GetQuestionBatchRequest)(nil),   // 0: questions.v1.GetQuestionBatchRequest
	(*GetQuestionsBatchResponse)(nil), // 1: questions.v1.GetQuestionsBatchResponse
	(Difficulty)(0),                   // 2: questions.v1.Difficulty
	(*Question)(nil),                  // 3: questions.v1.Question
}
var file_external_questions_v1_questions_proto_depIdxs = []int32{
	2, // 0: questions.v1.GetQuestionBatchRequest.difficulty:type_name -> questions.v1.Difficulty
	3, // 1: questions.v1.GetQuestionsBatchResponse.questions:type_name -> questions.v1.Question
	0, // 2: questions.v1.QuestionsService.GetQuestionBatch:input_type -> questions.v1.GetQuestionBatchRequest
	1, // 3: questions.v1.QuestionsService.GetQuestionBatch:output_type -> questions.v1.GetQuestionsBatchResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
			NumMessages:   2,
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
