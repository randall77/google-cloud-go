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
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: google/cloud/dialogflow/v2beta1/conversation_event.proto

package dialogflowpb

import (
	reflect "reflect"
	sync "sync"

	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enumeration of the types of events available.
type ConversationEvent_Type int32

const (
	// Type not set.
	ConversationEvent_TYPE_UNSPECIFIED ConversationEvent_Type = 0
	// A new conversation has been opened. This is fired when a telephone call
	// is answered, or a conversation is created via the API.
	ConversationEvent_CONVERSATION_STARTED ConversationEvent_Type = 1
	// An existing conversation has closed. This is fired when a telephone call
	// is terminated, or a conversation is closed via the API.
	ConversationEvent_CONVERSATION_FINISHED ConversationEvent_Type = 2
	// An existing conversation has received notification from Dialogflow that
	// human intervention is required.
	ConversationEvent_HUMAN_INTERVENTION_NEEDED ConversationEvent_Type = 3
	// An existing conversation has received a new message, either from API or
	// telephony. It is configured in
	// [ConversationProfile.new_message_event_notification_config][google.cloud.dialogflow.v2beta1.ConversationProfile.new_message_event_notification_config]
	ConversationEvent_NEW_MESSAGE ConversationEvent_Type = 5
	// Unrecoverable error during a telephone call.
	//
	// In general non-recoverable errors only occur if something was
	// misconfigured in the ConversationProfile corresponding to the call. After
	// a non-recoverable error, Dialogflow may stop responding.
	//
	// We don't fire this event:
	//
	// * in an API call because we can directly return the error, or,
	// * when we can recover from an error.
	ConversationEvent_UNRECOVERABLE_ERROR ConversationEvent_Type = 4
)

// Enum value maps for ConversationEvent_Type.
var (
	ConversationEvent_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "CONVERSATION_STARTED",
		2: "CONVERSATION_FINISHED",
		3: "HUMAN_INTERVENTION_NEEDED",
		5: "NEW_MESSAGE",
		4: "UNRECOVERABLE_ERROR",
	}
	ConversationEvent_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":          0,
		"CONVERSATION_STARTED":      1,
		"CONVERSATION_FINISHED":     2,
		"HUMAN_INTERVENTION_NEEDED": 3,
		"NEW_MESSAGE":               5,
		"UNRECOVERABLE_ERROR":       4,
	}
)

func (x ConversationEvent_Type) Enum() *ConversationEvent_Type {
	p := new(ConversationEvent_Type)
	*p = x
	return p
}

func (x ConversationEvent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversationEvent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dialogflow_v2beta1_conversation_event_proto_enumTypes[0].Descriptor()
}

func (ConversationEvent_Type) Type() protoreflect.EnumType {
	return &file_google_cloud_dialogflow_v2beta1_conversation_event_proto_enumTypes[0]
}

func (x ConversationEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversationEvent_Type.Descriptor instead.
func (ConversationEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2beta1_conversation_event_proto_rawDescGZIP(), []int{0, 0}
}

// Represents a notification sent to Pub/Sub subscribers for conversation
// lifecycle events.
type ConversationEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The unique identifier of the conversation this notification
	// refers to.
	// Format: `projects/<Project ID>/conversations/<Conversation ID>`.
	Conversation string `protobuf:"bytes,1,opt,name=conversation,proto3" json:"conversation,omitempty"`
	// Required. The type of the event that this notification refers to.
	Type ConversationEvent_Type `protobuf:"varint,2,opt,name=type,proto3,enum=google.cloud.dialogflow.v2beta1.ConversationEvent_Type" json:"type,omitempty"`
	// Optional. More detailed information about an error. Only set for type
	// UNRECOVERABLE_ERROR_IN_PHONE_CALL.
	ErrorStatus *status.Status `protobuf:"bytes,3,opt,name=error_status,json=errorStatus,proto3" json:"error_status,omitempty"`
	// Payload of conversation event.
	//
	// Types that are assignable to Payload:
	//	*ConversationEvent_NewMessagePayload
	Payload isConversationEvent_Payload `protobuf_oneof:"payload"`
}

func (x *ConversationEvent) Reset() {
	*x = ConversationEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_v2beta1_conversation_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationEvent) ProtoMessage() {}

func (x *ConversationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2beta1_conversation_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationEvent.ProtoReflect.Descriptor instead.
func (*ConversationEvent) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2beta1_conversation_event_proto_rawDescGZIP(), []int{0}
}

func (x *ConversationEvent) GetConversation() string {
	if x != nil {
		return x.Conversation
	}
	return ""
}

func (x *ConversationEvent) GetType() ConversationEvent_Type {
	if x != nil {
		return x.Type
	}
	return ConversationEvent_TYPE_UNSPECIFIED
}

func (x *ConversationEvent) GetErrorStatus() *status.Status {
	if x != nil {
		return x.ErrorStatus
	}
	return nil
}

func (m *ConversationEvent) GetPayload() isConversationEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *ConversationEvent) GetNewMessagePayload() *Message {
	if x, ok := x.GetPayload().(*ConversationEvent_NewMessagePayload); ok {
		return x.NewMessagePayload
	}
	return nil
}

type isConversationEvent_Payload interface {
	isConversationEvent_Payload()
}

type ConversationEvent_NewMessagePayload struct {
	// Payload of NEW_MESSAGE event.
	NewMessagePayload *Message `protobuf:"bytes,4,opt,name=new_message_payload,json=newMessagePayload,proto3,oneof"`
}

func (*ConversationEvent_NewMessagePayload) isConversationEvent_Payload() {}

var File_google_cloud_dialogflow_v2beta1_conversation_event_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_v2beta1_conversation_event_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x31, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x03, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x4b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x35,
	0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5a, 0x0a, 0x13, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x11,
	0x6e, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x9a, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f,
	0x4e, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53,
	0x48, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x48, 0x55, 0x4d, 0x41, 0x4e, 0x5f, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x56, 0x45, 0x4e, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x45, 0x45, 0x44,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x45, 0x57, 0x5f, 0x4d, 0x45, 0x53, 0x53,
	0x41, 0x47, 0x45, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x4e, 0x52, 0x45, 0x43, 0x4f, 0x56,
	0x45, 0x52, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x42, 0x09,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0xae, 0x01, 0x0a, 0x23, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x42, 0x16, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f,
	0x77, 0x70, 0x62, 0x3b, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x70, 0x62,
	0xf8, 0x01, 0x01, 0xa2, 0x02, 0x02, 0x44, 0x46, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x56, 0x32, 0x42, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_dialogflow_v2beta1_conversation_event_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_v2beta1_conversation_event_proto_rawDescData = file_google_cloud_dialogflow_v2beta1_conversation_event_proto_rawDesc
)

func file_google_cloud_dialogflow_v2beta1_conversation_event_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_v2beta1_conversation_event_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_v2beta1_conversation_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_v2beta1_conversation_event_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_v2beta1_conversation_event_proto_rawDescData
}

var file_google_cloud_dialogflow_v2beta1_conversation_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_dialogflow_v2beta1_conversation_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_dialogflow_v2beta1_conversation_event_proto_goTypes = []interface{}{
	(ConversationEvent_Type)(0), // 0: google.cloud.dialogflow.v2beta1.ConversationEvent.Type
	(*ConversationEvent)(nil),   // 1: google.cloud.dialogflow.v2beta1.ConversationEvent
	(*status.Status)(nil),       // 2: google.rpc.Status
	(*Message)(nil),             // 3: google.cloud.dialogflow.v2beta1.Message
}
var file_google_cloud_dialogflow_v2beta1_conversation_event_proto_depIdxs = []int32{
	0, // 0: google.cloud.dialogflow.v2beta1.ConversationEvent.type:type_name -> google.cloud.dialogflow.v2beta1.ConversationEvent.Type
	2, // 1: google.cloud.dialogflow.v2beta1.ConversationEvent.error_status:type_name -> google.rpc.Status
	3, // 2: google.cloud.dialogflow.v2beta1.ConversationEvent.new_message_payload:type_name -> google.cloud.dialogflow.v2beta1.Message
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_v2beta1_conversation_event_proto_init() }
func file_google_cloud_dialogflow_v2beta1_conversation_event_proto_init() {
	if File_google_cloud_dialogflow_v2beta1_conversation_event_proto != nil {
		return
	}
	file_google_cloud_dialogflow_v2beta1_participant_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dialogflow_v2beta1_conversation_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationEvent); i {
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
	file_google_cloud_dialogflow_v2beta1_conversation_event_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ConversationEvent_NewMessagePayload)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_dialogflow_v2beta1_conversation_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dialogflow_v2beta1_conversation_event_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_v2beta1_conversation_event_proto_depIdxs,
		EnumInfos:         file_google_cloud_dialogflow_v2beta1_conversation_event_proto_enumTypes,
		MessageInfos:      file_google_cloud_dialogflow_v2beta1_conversation_event_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_v2beta1_conversation_event_proto = out.File
	file_google_cloud_dialogflow_v2beta1_conversation_event_proto_rawDesc = nil
	file_google_cloud_dialogflow_v2beta1_conversation_event_proto_goTypes = nil
	file_google_cloud_dialogflow_v2beta1_conversation_event_proto_depIdxs = nil
}
