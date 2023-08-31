// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/cloud/dialogflow/cx/v3/safety_settings.proto

package cxpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Settings for Generative Safety.
type SafetySettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Banned phrases for generated text.
	BannedPhrases []*SafetySettings_Phrase `protobuf:"bytes,1,rep,name=banned_phrases,json=bannedPhrases,proto3" json:"banned_phrases,omitempty"`
}

func (x *SafetySettings) Reset() {
	*x = SafetySettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3_safety_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SafetySettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SafetySettings) ProtoMessage() {}

func (x *SafetySettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3_safety_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SafetySettings.ProtoReflect.Descriptor instead.
func (*SafetySettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3_safety_settings_proto_rawDescGZIP(), []int{0}
}

func (x *SafetySettings) GetBannedPhrases() []*SafetySettings_Phrase {
	if x != nil {
		return x.BannedPhrases
	}
	return nil
}

// Text input which can be used for prompt or banned phrases.
type SafetySettings_Phrase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Text input which can be used for prompt or banned phrases.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Required. Language code of the phrase.
	LanguageCode string `protobuf:"bytes,2,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
}

func (x *SafetySettings_Phrase) Reset() {
	*x = SafetySettings_Phrase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3_safety_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SafetySettings_Phrase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SafetySettings_Phrase) ProtoMessage() {}

func (x *SafetySettings_Phrase) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3_safety_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SafetySettings_Phrase.ProtoReflect.Descriptor instead.
func (*SafetySettings_Phrase) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3_safety_settings_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SafetySettings_Phrase) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SafetySettings_Phrase) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

var File_google_cloud_dialogflow_cx_v3_safety_settings_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_cx_v3_safety_settings_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33, 0x2f,
	0x73, 0x61, 0x66, 0x65, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63,
	0x78, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x0e, 0x53, 0x61, 0x66, 0x65, 0x74, 0x79,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x5b, 0x0a, 0x0e, 0x62, 0x61, 0x6e, 0x6e,
	0x65, 0x64, 0x5f, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33,
	0x2e, 0x53, 0x61, 0x66, 0x65, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x52, 0x0d, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x50, 0x68,
	0x72, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x4b, 0x0a, 0x06, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x12,
	0x17, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x28, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x42, 0x95, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x42, 0x13, 0x53, 0x61, 0x66, 0x65, 0x74, 0x79,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x63, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2f, 0x63, 0x78, 0x70, 0x62, 0x3b, 0x63, 0x78,
	0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x02, 0x44, 0x46, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x78, 0x2e, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_dialogflow_cx_v3_safety_settings_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_cx_v3_safety_settings_proto_rawDescData = file_google_cloud_dialogflow_cx_v3_safety_settings_proto_rawDesc
)

func file_google_cloud_dialogflow_cx_v3_safety_settings_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_cx_v3_safety_settings_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_cx_v3_safety_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_cx_v3_safety_settings_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_cx_v3_safety_settings_proto_rawDescData
}

var file_google_cloud_dialogflow_cx_v3_safety_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_dialogflow_cx_v3_safety_settings_proto_goTypes = []interface{}{
	(*SafetySettings)(nil),        // 0: google.cloud.dialogflow.cx.v3.SafetySettings
	(*SafetySettings_Phrase)(nil), // 1: google.cloud.dialogflow.cx.v3.SafetySettings.Phrase
}
var file_google_cloud_dialogflow_cx_v3_safety_settings_proto_depIdxs = []int32{
	1, // 0: google.cloud.dialogflow.cx.v3.SafetySettings.banned_phrases:type_name -> google.cloud.dialogflow.cx.v3.SafetySettings.Phrase
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_cx_v3_safety_settings_proto_init() }
func file_google_cloud_dialogflow_cx_v3_safety_settings_proto_init() {
	if File_google_cloud_dialogflow_cx_v3_safety_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dialogflow_cx_v3_safety_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SafetySettings); i {
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
		file_google_cloud_dialogflow_cx_v3_safety_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SafetySettings_Phrase); i {
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
			RawDescriptor: file_google_cloud_dialogflow_cx_v3_safety_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dialogflow_cx_v3_safety_settings_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_cx_v3_safety_settings_proto_depIdxs,
		MessageInfos:      file_google_cloud_dialogflow_cx_v3_safety_settings_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_cx_v3_safety_settings_proto = out.File
	file_google_cloud_dialogflow_cx_v3_safety_settings_proto_rawDesc = nil
	file_google_cloud_dialogflow_cx_v3_safety_settings_proto_goTypes = nil
	file_google_cloud_dialogflow_cx_v3_safety_settings_proto_depIdxs = nil
}
