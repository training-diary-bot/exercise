// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: exercise/v1/get_default_exercises.proto

package exercisev1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/training-diary-bot/apis/genproto/go/common/v1"
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

type GetExercisesByMuscleGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MuscleGroupId string `protobuf:"bytes,2,opt,name=muscle_group_id,json=muscleGroupId,proto3" json:"muscle_group_id,omitempty"`
}

func (x *GetExercisesByMuscleGroupRequest) Reset() {
	*x = GetExercisesByMuscleGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercise_v1_get_default_exercises_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExercisesByMuscleGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExercisesByMuscleGroupRequest) ProtoMessage() {}

func (x *GetExercisesByMuscleGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercise_v1_get_default_exercises_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExercisesByMuscleGroupRequest.ProtoReflect.Descriptor instead.
func (*GetExercisesByMuscleGroupRequest) Descriptor() ([]byte, []int) {
	return file_exercise_v1_get_default_exercises_proto_rawDescGZIP(), []int{0}
}

func (x *GetExercisesByMuscleGroupRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetExercisesByMuscleGroupRequest) GetMuscleGroupId() string {
	if x != nil {
		return x.MuscleGroupId
	}
	return ""
}

type GetExercisesByMuscleGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercises []*v1.Exercise `protobuf:"bytes,1,rep,name=exercises,proto3" json:"exercises,omitempty"`
}

func (x *GetExercisesByMuscleGroupResponse) Reset() {
	*x = GetExercisesByMuscleGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercise_v1_get_default_exercises_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExercisesByMuscleGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExercisesByMuscleGroupResponse) ProtoMessage() {}

func (x *GetExercisesByMuscleGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercise_v1_get_default_exercises_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExercisesByMuscleGroupResponse.ProtoReflect.Descriptor instead.
func (*GetExercisesByMuscleGroupResponse) Descriptor() ([]byte, []int) {
	return file_exercise_v1_get_default_exercises_proto_rawDescGZIP(), []int{1}
}

func (x *GetExercisesByMuscleGroupResponse) GetExercises() []*v1.Exercise {
	if x != nil {
		return x.Exercises
	}
	return nil
}

var File_exercise_v1_get_default_exercises_proto protoreflect.FileDescriptor

var file_exercise_v1_get_default_exercises_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65,
	0x74, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x20, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x75, 0x73, 0x63, 0x6c,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x30, 0x0a, 0x0f, 0x6d, 0x75, 0x73, 0x63, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03,
	0xb0, 0x01, 0x01, 0x52, 0x0d, 0x6d, 0x75, 0x73, 0x63, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x22, 0x56, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x73, 0x42, 0x79, 0x4d, 0x75, 0x73, 0x63, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52,
	0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x2d, 0x64, 0x69, 0x61, 0x72, 0x79, 0x2d, 0x62, 0x6f, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exercise_v1_get_default_exercises_proto_rawDescOnce sync.Once
	file_exercise_v1_get_default_exercises_proto_rawDescData = file_exercise_v1_get_default_exercises_proto_rawDesc
)

func file_exercise_v1_get_default_exercises_proto_rawDescGZIP() []byte {
	file_exercise_v1_get_default_exercises_proto_rawDescOnce.Do(func() {
		file_exercise_v1_get_default_exercises_proto_rawDescData = protoimpl.X.CompressGZIP(file_exercise_v1_get_default_exercises_proto_rawDescData)
	})
	return file_exercise_v1_get_default_exercises_proto_rawDescData
}

var file_exercise_v1_get_default_exercises_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_exercise_v1_get_default_exercises_proto_goTypes = []interface{}{
	(*GetExercisesByMuscleGroupRequest)(nil),  // 0: exercise.v1.GetExercisesByMuscleGroupRequest
	(*GetExercisesByMuscleGroupResponse)(nil), // 1: exercise.v1.GetExercisesByMuscleGroupResponse
	(*v1.Exercise)(nil),                       // 2: common.v1.Exercise
}
var file_exercise_v1_get_default_exercises_proto_depIdxs = []int32{
	2, // 0: exercise.v1.GetExercisesByMuscleGroupResponse.exercises:type_name -> common.v1.Exercise
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_exercise_v1_get_default_exercises_proto_init() }
func file_exercise_v1_get_default_exercises_proto_init() {
	if File_exercise_v1_get_default_exercises_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exercise_v1_get_default_exercises_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExercisesByMuscleGroupRequest); i {
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
		file_exercise_v1_get_default_exercises_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExercisesByMuscleGroupResponse); i {
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
			RawDescriptor: file_exercise_v1_get_default_exercises_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_exercise_v1_get_default_exercises_proto_goTypes,
		DependencyIndexes: file_exercise_v1_get_default_exercises_proto_depIdxs,
		MessageInfos:      file_exercise_v1_get_default_exercises_proto_msgTypes,
	}.Build()
	File_exercise_v1_get_default_exercises_proto = out.File
	file_exercise_v1_get_default_exercises_proto_rawDesc = nil
	file_exercise_v1_get_default_exercises_proto_goTypes = nil
	file_exercise_v1_get_default_exercises_proto_depIdxs = nil
}