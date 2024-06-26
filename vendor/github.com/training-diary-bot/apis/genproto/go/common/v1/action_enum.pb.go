// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: common/v1/action_enum.proto

package commonv1

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

type ActionID int32

const (
	ActionID_ACTION_ID_UNSPECIFIED                       ActionID = 0
	ActionID_ACTION_ID_START                             ActionID = 1
	ActionID_ACTION_ID_MENU                              ActionID = 2
	ActionID_ACTION_ID_CREATING_TRAINING_DATE            ActionID = 3
	ActionID_ACTION_ID_CREATING_TRAINING_MUSCLE_GROUP    ActionID = 4
	ActionID_ACTION_ID_CREATING_TRAINING_EXERCISE        ActionID = 5
	ActionID_ACTION_ID_CREATING_TRAINING_EXERCISE_REPEAT ActionID = 6
	ActionID_ACTION_ID_CREATING_TRAINING_EXERCISE_WEIGHT ActionID = 7
	ActionID_ACTION_ID_CREATING_TRAINING_SUCCESS         ActionID = 8
	ActionID_ACTION_ID_VIEWING_TRAININGS_DATE            ActionID = 9
	ActionID_ACTION_ID_VIEWING_TRAINING                  ActionID = 10
	ActionID_ACTION_ID_DELETING_TRAINING_CONFIRM         ActionID = 11
	ActionID_ACTION_ID_DELETING_TRAINING_SUCCESS         ActionID = 12
	ActionID_ACTION_ID_DELETING_TRAINING_NOT_SUCCESS     ActionID = 13
	ActionID_ACTION_ID_STARTING_TRAINING                 ActionID = 14
	ActionID_ACTION_ID_COMPLETING_EXERCISE               ActionID = 15
	ActionID_ACTION_ID_FINISHING_TRAINING                ActionID = 16
)

// Enum value maps for ActionID.
var (
	ActionID_name = map[int32]string{
		0:  "ACTION_ID_UNSPECIFIED",
		1:  "ACTION_ID_START",
		2:  "ACTION_ID_MENU",
		3:  "ACTION_ID_CREATING_TRAINING_DATE",
		4:  "ACTION_ID_CREATING_TRAINING_MUSCLE_GROUP",
		5:  "ACTION_ID_CREATING_TRAINING_EXERCISE",
		6:  "ACTION_ID_CREATING_TRAINING_EXERCISE_REPEAT",
		7:  "ACTION_ID_CREATING_TRAINING_EXERCISE_WEIGHT",
		8:  "ACTION_ID_CREATING_TRAINING_SUCCESS",
		9:  "ACTION_ID_VIEWING_TRAININGS_DATE",
		10: "ACTION_ID_VIEWING_TRAINING",
		11: "ACTION_ID_DELETING_TRAINING_CONFIRM",
		12: "ACTION_ID_DELETING_TRAINING_SUCCESS",
		13: "ACTION_ID_DELETING_TRAINING_NOT_SUCCESS",
		14: "ACTION_ID_STARTING_TRAINING",
		15: "ACTION_ID_COMPLETING_EXERCISE",
		16: "ACTION_ID_FINISHING_TRAINING",
	}
	ActionID_value = map[string]int32{
		"ACTION_ID_UNSPECIFIED":                       0,
		"ACTION_ID_START":                             1,
		"ACTION_ID_MENU":                              2,
		"ACTION_ID_CREATING_TRAINING_DATE":            3,
		"ACTION_ID_CREATING_TRAINING_MUSCLE_GROUP":    4,
		"ACTION_ID_CREATING_TRAINING_EXERCISE":        5,
		"ACTION_ID_CREATING_TRAINING_EXERCISE_REPEAT": 6,
		"ACTION_ID_CREATING_TRAINING_EXERCISE_WEIGHT": 7,
		"ACTION_ID_CREATING_TRAINING_SUCCESS":         8,
		"ACTION_ID_VIEWING_TRAININGS_DATE":            9,
		"ACTION_ID_VIEWING_TRAINING":                  10,
		"ACTION_ID_DELETING_TRAINING_CONFIRM":         11,
		"ACTION_ID_DELETING_TRAINING_SUCCESS":         12,
		"ACTION_ID_DELETING_TRAINING_NOT_SUCCESS":     13,
		"ACTION_ID_STARTING_TRAINING":                 14,
		"ACTION_ID_COMPLETING_EXERCISE":               15,
		"ACTION_ID_FINISHING_TRAINING":                16,
	}
)

func (x ActionID) Enum() *ActionID {
	p := new(ActionID)
	*p = x
	return p
}

func (x ActionID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionID) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_action_enum_proto_enumTypes[0].Descriptor()
}

func (ActionID) Type() protoreflect.EnumType {
	return &file_common_v1_action_enum_proto_enumTypes[0]
}

func (x ActionID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionID.Descriptor instead.
func (ActionID) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_action_enum_proto_rawDescGZIP(), []int{0}
}

var File_common_v1_action_enum_proto protoreflect.FileDescriptor

var file_common_v1_action_enum_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2a, 0x82, 0x05, 0x0a, 0x08, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x49, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x49, 0x44, 0x5f, 0x4d, 0x45, 0x4e, 0x55, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12,
	0x2c, 0x0a, 0x28, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4d,
	0x55, 0x53, 0x43, 0x4c, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x04, 0x12, 0x28, 0x0a,
	0x24, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x49, 0x4e, 0x47, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x58, 0x45,
	0x52, 0x43, 0x49, 0x53, 0x45, 0x10, 0x05, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x52,
	0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x58, 0x45, 0x52, 0x43, 0x49, 0x53, 0x45, 0x5f,
	0x52, 0x45, 0x50, 0x45, 0x41, 0x54, 0x10, 0x06, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x54,
	0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x58, 0x45, 0x52, 0x43, 0x49, 0x53, 0x45,
	0x5f, 0x57, 0x45, 0x49, 0x47, 0x48, 0x54, 0x10, 0x07, 0x12, 0x27, 0x0a, 0x23, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x08, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f,
	0x56, 0x49, 0x45, 0x57, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47,
	0x53, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x52,
	0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x27, 0x0a, 0x23, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x54,
	0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x10,
	0x0b, 0x12, 0x27, 0x0a, 0x23, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x0c, 0x12, 0x2b, 0x0a, 0x27, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47,
	0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x0d, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x52,
	0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x0e, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47,
	0x5f, 0x45, 0x58, 0x45, 0x52, 0x43, 0x49, 0x53, 0x45, 0x10, 0x0f, 0x12, 0x20, 0x0a, 0x1c, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x49,
	0x4e, 0x47, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x10, 0x42, 0x43, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x64, 0x69, 0x61, 0x72, 0x79, 0x2d, 0x62, 0x6f, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_v1_action_enum_proto_rawDescOnce sync.Once
	file_common_v1_action_enum_proto_rawDescData = file_common_v1_action_enum_proto_rawDesc
)

func file_common_v1_action_enum_proto_rawDescGZIP() []byte {
	file_common_v1_action_enum_proto_rawDescOnce.Do(func() {
		file_common_v1_action_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v1_action_enum_proto_rawDescData)
	})
	return file_common_v1_action_enum_proto_rawDescData
}

var file_common_v1_action_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_v1_action_enum_proto_goTypes = []interface{}{
	(ActionID)(0), // 0: common.v1.ActionID
}
var file_common_v1_action_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_v1_action_enum_proto_init() }
func file_common_v1_action_enum_proto_init() {
	if File_common_v1_action_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_v1_action_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_action_enum_proto_goTypes,
		DependencyIndexes: file_common_v1_action_enum_proto_depIdxs,
		EnumInfos:         file_common_v1_action_enum_proto_enumTypes,
	}.Build()
	File_common_v1_action_enum_proto = out.File
	file_common_v1_action_enum_proto_rawDesc = nil
	file_common_v1_action_enum_proto_goTypes = nil
	file_common_v1_action_enum_proto_depIdxs = nil
}
