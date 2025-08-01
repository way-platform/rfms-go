// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: wayplatform/rfms/v5/charging_state.proto

package rfmsv5

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Charging state.
type ChargingState int32

const (
	// Default value. This value is not used.
	ChargingState_CHARGING_STATE_UNSPECIFIED ChargingState = 0
	// The charging state is unknown.
	ChargingState_CHARGING_STATE_UNKNOWN ChargingState = 1
	// The charging state is not available.
	ChargingState_CHARGING_STATE_NOT_AVAILABLE ChargingState = 2
	// An error occurred when charging.
	ChargingState_CHARGING_STATE_ERROR ChargingState = 3
	// The charging state is not charging.
	ChargingState_NOT_CHARGING ChargingState = 4
	// The charging state is charging. AC or DC is unknown.
	ChargingState_CHARGING ChargingState = 5
	// Charging is ongoing using AC.
	ChargingState_CHARGING_AC ChargingState = 6
	// Charging is ongoing using DC.
	ChargingState_CHARGING_DC ChargingState = 7
)

// Enum value maps for ChargingState.
var (
	ChargingState_name = map[int32]string{
		0: "CHARGING_STATE_UNSPECIFIED",
		1: "CHARGING_STATE_UNKNOWN",
		2: "CHARGING_STATE_NOT_AVAILABLE",
		3: "CHARGING_STATE_ERROR",
		4: "NOT_CHARGING",
		5: "CHARGING",
		6: "CHARGING_AC",
		7: "CHARGING_DC",
	}
	ChargingState_value = map[string]int32{
		"CHARGING_STATE_UNSPECIFIED":   0,
		"CHARGING_STATE_UNKNOWN":       1,
		"CHARGING_STATE_NOT_AVAILABLE": 2,
		"CHARGING_STATE_ERROR":         3,
		"NOT_CHARGING":                 4,
		"CHARGING":                     5,
		"CHARGING_AC":                  6,
		"CHARGING_DC":                  7,
	}
)

func (x ChargingState) Enum() *ChargingState {
	p := new(ChargingState)
	*p = x
	return p
}

func (x ChargingState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChargingState) Descriptor() protoreflect.EnumDescriptor {
	return file_wayplatform_rfms_v5_charging_state_proto_enumTypes[0].Descriptor()
}

func (ChargingState) Type() protoreflect.EnumType {
	return &file_wayplatform_rfms_v5_charging_state_proto_enumTypes[0]
}

func (x ChargingState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

var File_wayplatform_rfms_v5_charging_state_proto protoreflect.FileDescriptor

const file_wayplatform_rfms_v5_charging_state_proto_rawDesc = "" +
	"\n" +
	"(wayplatform/rfms/v5/charging_state.proto\x12\x13wayplatform.rfms.v5*\xc9\x01\n" +
	"\rChargingState\x12\x1e\n" +
	"\x1aCHARGING_STATE_UNSPECIFIED\x10\x00\x12\x1a\n" +
	"\x16CHARGING_STATE_UNKNOWN\x10\x01\x12 \n" +
	"\x1cCHARGING_STATE_NOT_AVAILABLE\x10\x02\x12\x18\n" +
	"\x14CHARGING_STATE_ERROR\x10\x03\x12\x10\n" +
	"\fNOT_CHARGING\x10\x04\x12\f\n" +
	"\bCHARGING\x10\x05\x12\x0f\n" +
	"\vCHARGING_AC\x10\x06\x12\x0f\n" +
	"\vCHARGING_DC\x10\aB\xe4\x01\n" +
	"\x17com.wayplatform.rfms.v5B\x12ChargingStateProtoP\x01ZGgithub.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5;rfmsv5\xa2\x02\x03WRX\xaa\x02\x13Wayplatform.Rfms.V5\xca\x02\x13Wayplatform\\Rfms\\V5\xe2\x02\x1fWayplatform\\Rfms\\V5\\GPBMetadata\xea\x02\x15Wayplatform::Rfms::V5b\beditionsp\xe8\a"

var file_wayplatform_rfms_v5_charging_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wayplatform_rfms_v5_charging_state_proto_goTypes = []any{
	(ChargingState)(0), // 0: wayplatform.rfms.v5.ChargingState
}
var file_wayplatform_rfms_v5_charging_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wayplatform_rfms_v5_charging_state_proto_init() }
func file_wayplatform_rfms_v5_charging_state_proto_init() {
	if File_wayplatform_rfms_v5_charging_state_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_wayplatform_rfms_v5_charging_state_proto_rawDesc), len(file_wayplatform_rfms_v5_charging_state_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wayplatform_rfms_v5_charging_state_proto_goTypes,
		DependencyIndexes: file_wayplatform_rfms_v5_charging_state_proto_depIdxs,
		EnumInfos:         file_wayplatform_rfms_v5_charging_state_proto_enumTypes,
	}.Build()
	File_wayplatform_rfms_v5_charging_state_proto = out.File
	file_wayplatform_rfms_v5_charging_state_proto_goTypes = nil
	file_wayplatform_rfms_v5_charging_state_proto_depIdxs = nil
}
