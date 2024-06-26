// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: exercise/v1/get_exercises_by_ids.proto

package exercisev1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type GetExercisesByIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetExercisesByIDsRequest) Reset() {
	*x = GetExercisesByIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercise_v1_get_exercises_by_ids_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExercisesByIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExercisesByIDsRequest) ProtoMessage() {}

func (x *GetExercisesByIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercise_v1_get_exercises_by_ids_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExercisesByIDsRequest.ProtoReflect.Descriptor instead.
func (*GetExercisesByIDsRequest) Descriptor() ([]byte, []int) {
	return file_exercise_v1_get_exercises_by_ids_proto_rawDescGZIP(), []int{0}
}

func (x *GetExercisesByIDsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetExercisesByIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercises []*ExerciseEntity `protobuf:"bytes,1,rep,name=exercises,proto3" json:"exercises,omitempty"`
}

func (x *GetExercisesByIDsResponse) Reset() {
	*x = GetExercisesByIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercise_v1_get_exercises_by_ids_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExercisesByIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExercisesByIDsResponse) ProtoMessage() {}

func (x *GetExercisesByIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercise_v1_get_exercises_by_ids_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExercisesByIDsResponse.ProtoReflect.Descriptor instead.
func (*GetExercisesByIDsResponse) Descriptor() ([]byte, []int) {
	return file_exercise_v1_get_exercises_by_ids_proto_rawDescGZIP(), []int{1}
}

func (x *GetExercisesByIDsResponse) GetExercises() []*ExerciseEntity {
	if x != nil {
		return x.Exercises
	}
	return nil
}

var File_exercise_v1_get_exercises_by_ids_proto protoreflect.FileDescriptor

var file_exercise_v1_get_exercises_by_ids_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65,
	0x74, 0x5f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x69,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x72, 0x03, 0xb0,
	0x01, 0x01, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x56, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x42,
	0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x64, 0x69, 0x61, 0x72, 0x79, 0x2d, 0x62, 0x6f, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exercise_v1_get_exercises_by_ids_proto_rawDescOnce sync.Once
	file_exercise_v1_get_exercises_by_ids_proto_rawDescData = file_exercise_v1_get_exercises_by_ids_proto_rawDesc
)

func file_exercise_v1_get_exercises_by_ids_proto_rawDescGZIP() []byte {
	file_exercise_v1_get_exercises_by_ids_proto_rawDescOnce.Do(func() {
		file_exercise_v1_get_exercises_by_ids_proto_rawDescData = protoimpl.X.CompressGZIP(file_exercise_v1_get_exercises_by_ids_proto_rawDescData)
	})
	return file_exercise_v1_get_exercises_by_ids_proto_rawDescData
}

var file_exercise_v1_get_exercises_by_ids_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_exercise_v1_get_exercises_by_ids_proto_goTypes = []interface{}{
	(*GetExercisesByIDsRequest)(nil),  // 0: exercise.v1.GetExercisesByIDsRequest
	(*GetExercisesByIDsResponse)(nil), // 1: exercise.v1.GetExercisesByIDsResponse
	(*ExerciseEntity)(nil),            // 2: exercise.v1.ExerciseEntity
}
var file_exercise_v1_get_exercises_by_ids_proto_depIdxs = []int32{
	2, // 0: exercise.v1.GetExercisesByIDsResponse.exercises:type_name -> exercise.v1.ExerciseEntity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_exercise_v1_get_exercises_by_ids_proto_init() }
func file_exercise_v1_get_exercises_by_ids_proto_init() {
	if File_exercise_v1_get_exercises_by_ids_proto != nil {
		return
	}
	file_exercise_v1_entities_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_exercise_v1_get_exercises_by_ids_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExercisesByIDsRequest); i {
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
		file_exercise_v1_get_exercises_by_ids_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExercisesByIDsResponse); i {
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
			RawDescriptor: file_exercise_v1_get_exercises_by_ids_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_exercise_v1_get_exercises_by_ids_proto_goTypes,
		DependencyIndexes: file_exercise_v1_get_exercises_by_ids_proto_depIdxs,
		MessageInfos:      file_exercise_v1_get_exercises_by_ids_proto_msgTypes,
	}.Build()
	File_exercise_v1_get_exercises_by_ids_proto = out.File
	file_exercise_v1_get_exercises_by_ids_proto_rawDesc = nil
	file_exercise_v1_get_exercises_by_ids_proto_goTypes = nil
	file_exercise_v1_get_exercises_by_ids_proto_depIdxs = nil
}
