// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: error_view.proto

package errors

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

// ErrorViewType represents how the error has to be presented to the user
type ErrorViewType int32

const (
	ErrorViewType_ERROR_TYPE_UNSPECIFIED ErrorViewType = 0
	ErrorViewType_FULL_SCREEN            ErrorViewType = 1
	ErrorViewType_BOTTOM_SHEET           ErrorViewType = 2
)

// Enum value maps for ErrorViewType.
var (
	ErrorViewType_name = map[int32]string{
		0: "ERROR_TYPE_UNSPECIFIED",
		1: "FULL_SCREEN",
		2: "BOTTOM_SHEET",
	}
	ErrorViewType_value = map[string]int32{
		"ERROR_TYPE_UNSPECIFIED": 0,
		"FULL_SCREEN":            1,
		"BOTTOM_SHEET":           2,
	}
)

func (x ErrorViewType) Enum() *ErrorViewType {
	p := new(ErrorViewType)
	*p = x
	return p
}

func (x ErrorViewType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorViewType) Descriptor() protoreflect.EnumDescriptor {
	return file_error_view_proto_enumTypes[0].Descriptor()
}

func (ErrorViewType) Type() protoreflect.EnumType {
	return &file_error_view_proto_enumTypes[0]
}

func (x ErrorViewType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorViewType.Descriptor instead.
func (ErrorViewType) EnumDescriptor() ([]byte, []int) {
	return file_error_view_proto_rawDescGZIP(), []int{0}
}

// ErrorView represents the error UI for the client. It stores the
// error view type and the corresponding options for the ErrorViewType
// It contains the information that is needs to be shown on the
// client for graceful handling of failures/errors.
type ErrorView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ErrorViewType `protobuf:"varint,1,opt,name=type,proto3,enum=frontend.errors.ErrorViewType" json:"type,omitempty"`
	// screen_options is the payload that is used by the error view type
	// options map one-to-one with ErrorViewType
	//
	// Types that are assignable to Options:
	//
	//	*ErrorView_FullScreenErrorView
	//	*ErrorView_BottomSheetErrorView
	Options isErrorView_Options `protobuf_oneof:"options"`
}

func (x *ErrorView) Reset() {
	*x = ErrorView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorView) ProtoMessage() {}

func (x *ErrorView) ProtoReflect() protoreflect.Message {
	mi := &file_error_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorView.ProtoReflect.Descriptor instead.
func (*ErrorView) Descriptor() ([]byte, []int) {
	return file_error_view_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorView) GetType() ErrorViewType {
	if x != nil {
		return x.Type
	}
	return ErrorViewType_ERROR_TYPE_UNSPECIFIED
}

func (m *ErrorView) GetOptions() isErrorView_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (x *ErrorView) GetFullScreenErrorView() *FullScreenErrorView {
	if x, ok := x.GetOptions().(*ErrorView_FullScreenErrorView); ok {
		return x.FullScreenErrorView
	}
	return nil
}

func (x *ErrorView) GetBottomSheetErrorView() *BottomSheetErrorView {
	if x, ok := x.GetOptions().(*ErrorView_BottomSheetErrorView); ok {
		return x.BottomSheetErrorView
	}
	return nil
}

type isErrorView_Options interface {
	isErrorView_Options()
}

type ErrorView_FullScreenErrorView struct {
	FullScreenErrorView *FullScreenErrorView `protobuf:"bytes,2,opt,name=full_screen_error_view,json=fullScreenErrorView,proto3,oneof"`
}

type ErrorView_BottomSheetErrorView struct {
	BottomSheetErrorView *BottomSheetErrorView `protobuf:"bytes,3,opt,name=bottom_sheet_error_view,json=bottomSheetErrorView,proto3,oneof"`
}

func (*ErrorView_FullScreenErrorView) isErrorView_Options() {}

func (*ErrorView_BottomSheetErrorView) isErrorView_Options() {}

type FullScreenErrorView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	SubTitle string `protobuf:"bytes,2,opt,name=sub_title,json=subTitle,proto3" json:"sub_title,omitempty"`
}

func (x *FullScreenErrorView) Reset() {
	*x = FullScreenErrorView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_view_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullScreenErrorView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullScreenErrorView) ProtoMessage() {}

func (x *FullScreenErrorView) ProtoReflect() protoreflect.Message {
	mi := &file_error_view_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullScreenErrorView.ProtoReflect.Descriptor instead.
func (*FullScreenErrorView) Descriptor() ([]byte, []int) {
	return file_error_view_proto_rawDescGZIP(), []int{1}
}

func (x *FullScreenErrorView) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FullScreenErrorView) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

type BottomSheetErrorView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	SubTitle string `protobuf:"bytes,2,opt,name=sub_title,json=subTitle,proto3" json:"sub_title,omitempty"`
}

func (x *BottomSheetErrorView) Reset() {
	*x = BottomSheetErrorView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_view_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BottomSheetErrorView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BottomSheetErrorView) ProtoMessage() {}

func (x *BottomSheetErrorView) ProtoReflect() protoreflect.Message {
	mi := &file_error_view_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BottomSheetErrorView.ProtoReflect.Descriptor instead.
func (*BottomSheetErrorView) Descriptor() ([]byte, []int) {
	return file_error_view_proto_rawDescGZIP(), []int{2}
}

func (x *BottomSheetErrorView) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BottomSheetErrorView) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

var File_error_view_proto protoreflect.FileDescriptor

var file_error_view_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x22, 0x87, 0x02, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x56, 0x69, 0x65,
	0x77, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x5b, 0x0a, 0x16, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x53, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x48, 0x00, 0x52, 0x13, 0x66,
	0x75, 0x6c, 0x6c, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x5e, 0x0a, 0x17, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x53, 0x68, 0x65, 0x65,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x48, 0x00, 0x52, 0x14, 0x62, 0x6f,
	0x74, 0x74, 0x6f, 0x6d, 0x53, 0x68, 0x65, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x56, 0x69,
	0x65, 0x77, 0x42, 0x09, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x48, 0x0a,
	0x13, 0x46, 0x75, 0x6c, 0x6c, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x56, 0x69, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75,
	0x62, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x49, 0x0a, 0x14, 0x42, 0x6f, 0x74, 0x74, 0x6f,
	0x6d, 0x53, 0x68, 0x65, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x2a, 0x4e, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x53, 0x43, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x42, 0x4f, 0x54, 0x54, 0x4f, 0x4d, 0x5f, 0x53, 0x48, 0x45, 0x45, 0x54,
	0x10, 0x02, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x6a, 0x61, 0x7a, 0x62, 0x61, 0x61,
	0x74, 0x69, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_view_proto_rawDescOnce sync.Once
	file_error_view_proto_rawDescData = file_error_view_proto_rawDesc
)

func file_error_view_proto_rawDescGZIP() []byte {
	file_error_view_proto_rawDescOnce.Do(func() {
		file_error_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_view_proto_rawDescData)
	})
	return file_error_view_proto_rawDescData
}

var file_error_view_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_view_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_error_view_proto_goTypes = []interface{}{
	(ErrorViewType)(0),           // 0: frontend.errors.ErrorViewType
	(*ErrorView)(nil),            // 1: frontend.errors.ErrorView
	(*FullScreenErrorView)(nil),  // 2: frontend.errors.FullScreenErrorView
	(*BottomSheetErrorView)(nil), // 3: frontend.errors.BottomSheetErrorView
}
var file_error_view_proto_depIdxs = []int32{
	0, // 0: frontend.errors.ErrorView.type:type_name -> frontend.errors.ErrorViewType
	2, // 1: frontend.errors.ErrorView.full_screen_error_view:type_name -> frontend.errors.FullScreenErrorView
	3, // 2: frontend.errors.ErrorView.bottom_sheet_error_view:type_name -> frontend.errors.BottomSheetErrorView
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_error_view_proto_init() }
func file_error_view_proto_init() {
	if File_error_view_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_error_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorView); i {
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
		file_error_view_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullScreenErrorView); i {
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
		file_error_view_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BottomSheetErrorView); i {
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
	file_error_view_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ErrorView_FullScreenErrorView)(nil),
		(*ErrorView_BottomSheetErrorView)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_view_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_view_proto_goTypes,
		DependencyIndexes: file_error_view_proto_depIdxs,
		EnumInfos:         file_error_view_proto_enumTypes,
		MessageInfos:      file_error_view_proto_msgTypes,
	}.Build()
	File_error_view_proto = out.File
	file_error_view_proto_rawDesc = nil
	file_error_view_proto_goTypes = nil
	file_error_view_proto_depIdxs = nil
}
