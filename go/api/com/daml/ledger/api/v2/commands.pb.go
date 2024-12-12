// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: com/daml/ledger/api/v2/commands.proto

package v2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Command:
	//
	//	*Command_Create
	//	*Command_Exercise
	//	*Command_ExerciseByKey
	//	*Command_CreateAndExercise
	Command isCommand_Command `protobuf_oneof:"command"`
}

func (x *Command) Reset() {
	*x = Command{}
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_commands_proto_rawDescGZIP(), []int{0}
}

func (m *Command) GetCommand() isCommand_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (x *Command) GetCreate() *CreateCommand {
	if x, ok := x.GetCommand().(*Command_Create); ok {
		return x.Create
	}
	return nil
}

func (x *Command) GetExercise() *ExerciseCommand {
	if x, ok := x.GetCommand().(*Command_Exercise); ok {
		return x.Exercise
	}
	return nil
}

func (x *Command) GetExerciseByKey() *ExerciseByKeyCommand {
	if x, ok := x.GetCommand().(*Command_ExerciseByKey); ok {
		return x.ExerciseByKey
	}
	return nil
}

func (x *Command) GetCreateAndExercise() *CreateAndExerciseCommand {
	if x, ok := x.GetCommand().(*Command_CreateAndExercise); ok {
		return x.CreateAndExercise
	}
	return nil
}

type isCommand_Command interface {
	isCommand_Command()
}

type Command_Create struct {
	Create *CreateCommand `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type Command_Exercise struct {
	Exercise *ExerciseCommand `protobuf:"bytes,2,opt,name=exercise,proto3,oneof"`
}

type Command_ExerciseByKey struct {
	ExerciseByKey *ExerciseByKeyCommand `protobuf:"bytes,4,opt,name=exercise_by_key,json=exerciseByKey,proto3,oneof"`
}

type Command_CreateAndExercise struct {
	CreateAndExercise *CreateAndExerciseCommand `protobuf:"bytes,3,opt,name=create_and_exercise,json=createAndExercise,proto3,oneof"`
}

func (*Command_Create) isCommand_Command() {}

func (*Command_Exercise) isCommand_Command() {}

func (*Command_ExerciseByKey) isCommand_Command() {}

func (*Command_CreateAndExercise) isCommand_Command() {}

type CreateCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId      *Identifier `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	CreateArguments *Record     `protobuf:"bytes,2,opt,name=create_arguments,json=createArguments,proto3" json:"create_arguments,omitempty"`
}

func (x *CreateCommand) Reset() {
	*x = CreateCommand{}
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommand) ProtoMessage() {}

func (x *CreateCommand) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommand.ProtoReflect.Descriptor instead.
func (*CreateCommand) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_commands_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCommand) GetTemplateId() *Identifier {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *CreateCommand) GetCreateArguments() *Record {
	if x != nil {
		return x.CreateArguments
	}
	return nil
}

type ExerciseCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId     *Identifier `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	ContractId     string      `protobuf:"bytes,2,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	Choice         string      `protobuf:"bytes,3,opt,name=choice,proto3" json:"choice,omitempty"`
	ChoiceArgument *Value      `protobuf:"bytes,4,opt,name=choice_argument,json=choiceArgument,proto3" json:"choice_argument,omitempty"`
}

func (x *ExerciseCommand) Reset() {
	*x = ExerciseCommand{}
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExerciseCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExerciseCommand) ProtoMessage() {}

func (x *ExerciseCommand) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExerciseCommand.ProtoReflect.Descriptor instead.
func (*ExerciseCommand) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_commands_proto_rawDescGZIP(), []int{2}
}

func (x *ExerciseCommand) GetTemplateId() *Identifier {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *ExerciseCommand) GetContractId() string {
	if x != nil {
		return x.ContractId
	}
	return ""
}

func (x *ExerciseCommand) GetChoice() string {
	if x != nil {
		return x.Choice
	}
	return ""
}

func (x *ExerciseCommand) GetChoiceArgument() *Value {
	if x != nil {
		return x.ChoiceArgument
	}
	return nil
}

type ExerciseByKeyCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId     *Identifier `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	ContractKey    *Value      `protobuf:"bytes,2,opt,name=contract_key,json=contractKey,proto3" json:"contract_key,omitempty"`
	Choice         string      `protobuf:"bytes,3,opt,name=choice,proto3" json:"choice,omitempty"`
	ChoiceArgument *Value      `protobuf:"bytes,4,opt,name=choice_argument,json=choiceArgument,proto3" json:"choice_argument,omitempty"`
}

func (x *ExerciseByKeyCommand) Reset() {
	*x = ExerciseByKeyCommand{}
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExerciseByKeyCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExerciseByKeyCommand) ProtoMessage() {}

func (x *ExerciseByKeyCommand) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExerciseByKeyCommand.ProtoReflect.Descriptor instead.
func (*ExerciseByKeyCommand) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_commands_proto_rawDescGZIP(), []int{3}
}

func (x *ExerciseByKeyCommand) GetTemplateId() *Identifier {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *ExerciseByKeyCommand) GetContractKey() *Value {
	if x != nil {
		return x.ContractKey
	}
	return nil
}

func (x *ExerciseByKeyCommand) GetChoice() string {
	if x != nil {
		return x.Choice
	}
	return ""
}

func (x *ExerciseByKeyCommand) GetChoiceArgument() *Value {
	if x != nil {
		return x.ChoiceArgument
	}
	return nil
}

type CreateAndExerciseCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId      *Identifier `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	CreateArguments *Record     `protobuf:"bytes,2,opt,name=create_arguments,json=createArguments,proto3" json:"create_arguments,omitempty"`
	Choice          string      `protobuf:"bytes,3,opt,name=choice,proto3" json:"choice,omitempty"`
	ChoiceArgument  *Value      `protobuf:"bytes,4,opt,name=choice_argument,json=choiceArgument,proto3" json:"choice_argument,omitempty"`
}

func (x *CreateAndExerciseCommand) Reset() {
	*x = CreateAndExerciseCommand{}
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAndExerciseCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAndExerciseCommand) ProtoMessage() {}

func (x *CreateAndExerciseCommand) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAndExerciseCommand.ProtoReflect.Descriptor instead.
func (*CreateAndExerciseCommand) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_commands_proto_rawDescGZIP(), []int{4}
}

func (x *CreateAndExerciseCommand) GetTemplateId() *Identifier {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *CreateAndExerciseCommand) GetCreateArguments() *Record {
	if x != nil {
		return x.CreateArguments
	}
	return nil
}

func (x *CreateAndExerciseCommand) GetChoice() string {
	if x != nil {
		return x.Choice
	}
	return ""
}

func (x *CreateAndExerciseCommand) GetChoiceArgument() *Value {
	if x != nil {
		return x.ChoiceArgument
	}
	return nil
}

type DisclosedContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId       *Identifier `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	ContractId       string      `protobuf:"bytes,2,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	CreatedEventBlob []byte      `protobuf:"bytes,3,opt,name=created_event_blob,json=createdEventBlob,proto3" json:"created_event_blob,omitempty"`
	DomainId         string      `protobuf:"bytes,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *DisclosedContract) Reset() {
	*x = DisclosedContract{}
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DisclosedContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisclosedContract) ProtoMessage() {}

func (x *DisclosedContract) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisclosedContract.ProtoReflect.Descriptor instead.
func (*DisclosedContract) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_commands_proto_rawDescGZIP(), []int{5}
}

func (x *DisclosedContract) GetTemplateId() *Identifier {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *DisclosedContract) GetContractId() string {
	if x != nil {
		return x.ContractId
	}
	return ""
}

func (x *DisclosedContract) GetCreatedEventBlob() []byte {
	if x != nil {
		return x.CreatedEventBlob
	}
	return nil
}

func (x *DisclosedContract) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type Commands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId    string     `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	ApplicationId string     `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	CommandId     string     `protobuf:"bytes,3,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	Commands      []*Command `protobuf:"bytes,4,rep,name=commands,proto3" json:"commands,omitempty"`
	// Types that are assignable to DeduplicationPeriod:
	//
	//	*Commands_DeduplicationDuration
	//	*Commands_DeduplicationOffset
	DeduplicationPeriod          isCommands_DeduplicationPeriod `protobuf_oneof:"deduplication_period"`
	MinLedgerTimeAbs             *timestamppb.Timestamp         `protobuf:"bytes,7,opt,name=min_ledger_time_abs,json=minLedgerTimeAbs,proto3" json:"min_ledger_time_abs,omitempty"`
	MinLedgerTimeRel             *durationpb.Duration           `protobuf:"bytes,8,opt,name=min_ledger_time_rel,json=minLedgerTimeRel,proto3" json:"min_ledger_time_rel,omitempty"`
	ActAs                        []string                       `protobuf:"bytes,9,rep,name=act_as,json=actAs,proto3" json:"act_as,omitempty"`
	ReadAs                       []string                       `protobuf:"bytes,10,rep,name=read_as,json=readAs,proto3" json:"read_as,omitempty"`
	SubmissionId                 string                         `protobuf:"bytes,11,opt,name=submission_id,json=submissionId,proto3" json:"submission_id,omitempty"`
	DisclosedContracts           []*DisclosedContract           `protobuf:"bytes,12,rep,name=disclosed_contracts,json=disclosedContracts,proto3" json:"disclosed_contracts,omitempty"`
	DomainId                     string                         `protobuf:"bytes,13,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	PackageIdSelectionPreference []string                       `protobuf:"bytes,14,rep,name=package_id_selection_preference,json=packageIdSelectionPreference,proto3" json:"package_id_selection_preference,omitempty"`
}

func (x *Commands) Reset() {
	*x = Commands{}
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Commands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commands) ProtoMessage() {}

func (x *Commands) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_commands_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commands.ProtoReflect.Descriptor instead.
func (*Commands) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_commands_proto_rawDescGZIP(), []int{6}
}

func (x *Commands) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *Commands) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *Commands) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

func (x *Commands) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

func (m *Commands) GetDeduplicationPeriod() isCommands_DeduplicationPeriod {
	if m != nil {
		return m.DeduplicationPeriod
	}
	return nil
}

func (x *Commands) GetDeduplicationDuration() *durationpb.Duration {
	if x, ok := x.GetDeduplicationPeriod().(*Commands_DeduplicationDuration); ok {
		return x.DeduplicationDuration
	}
	return nil
}

func (x *Commands) GetDeduplicationOffset() int64 {
	if x, ok := x.GetDeduplicationPeriod().(*Commands_DeduplicationOffset); ok {
		return x.DeduplicationOffset
	}
	return 0
}

func (x *Commands) GetMinLedgerTimeAbs() *timestamppb.Timestamp {
	if x != nil {
		return x.MinLedgerTimeAbs
	}
	return nil
}

func (x *Commands) GetMinLedgerTimeRel() *durationpb.Duration {
	if x != nil {
		return x.MinLedgerTimeRel
	}
	return nil
}

func (x *Commands) GetActAs() []string {
	if x != nil {
		return x.ActAs
	}
	return nil
}

func (x *Commands) GetReadAs() []string {
	if x != nil {
		return x.ReadAs
	}
	return nil
}

func (x *Commands) GetSubmissionId() string {
	if x != nil {
		return x.SubmissionId
	}
	return ""
}

func (x *Commands) GetDisclosedContracts() []*DisclosedContract {
	if x != nil {
		return x.DisclosedContracts
	}
	return nil
}

func (x *Commands) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *Commands) GetPackageIdSelectionPreference() []string {
	if x != nil {
		return x.PackageIdSelectionPreference
	}
	return nil
}

type isCommands_DeduplicationPeriod interface {
	isCommands_DeduplicationPeriod()
}

type Commands_DeduplicationDuration struct {
	DeduplicationDuration *durationpb.Duration `protobuf:"bytes,5,opt,name=deduplication_duration,json=deduplicationDuration,proto3,oneof"`
}

type Commands_DeduplicationOffset struct {
	DeduplicationOffset int64 `protobuf:"varint,6,opt,name=deduplication_offset,json=deduplicationOffset,proto3,oneof"`
}

func (*Commands_DeduplicationDuration) isCommands_DeduplicationPeriod() {}

func (*Commands_DeduplicationOffset) isCommands_DeduplicationPeriod() {}

var File_com_daml_ledger_api_v2_commands_proto protoreflect.FileDescriptor

var file_com_daml_ledger_api_v2_commands_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d,
	0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a,
	0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x3f, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x45, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x08,
	0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x5f, 0x62, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48,
	0x00, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79,
	0x12, 0x62, 0x0a, 0x13, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x65,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64,
	0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48,
	0x00, 0x52, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22,
	0x9f, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x43, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d,
	0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0xd7, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x43, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x63, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xfd, 0x01, 0x0a, 0x14,
	0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x43, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x63, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x8a, 0x02, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x43, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x49, 0x0a,
	0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61,
	0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x12, 0x46, 0x0a, 0x0f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xc4, 0x01, 0x0a, 0x11, 0x44, 0x69, 0x73,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x43,
	0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x6c,
	0x6f, 0x62, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22,
	0xf9, 0x05, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x12, 0x52, 0x0a, 0x16, 0x64, 0x65, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x15, 0x64,
	0x65, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x14, 0x64, 0x65, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x13, 0x64, 0x65, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x49, 0x0a, 0x13, 0x6d, 0x69, 0x6e,
	0x5f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x61, 0x62, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x41, 0x62, 0x73, 0x12, 0x48, 0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6d, 0x69,
	0x6e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6c, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x63, 0x74, 0x41, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x61, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x64, 0x41, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x5a, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x12, 0x64, 0x69, 0x73,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x1f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x42, 0x16, 0x0a, 0x14, 0x64, 0x65, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x42, 0x8c, 0x01, 0x0a, 0x16,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x42, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x4f, 0x75, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x38, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x32, 0xaa, 0x02, 0x16, 0x43, 0x6f, 0x6d, 0x2e, 0x44, 0x61, 0x6d, 0x6c, 0x2e, 0x4c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_daml_ledger_api_v2_commands_proto_rawDescOnce sync.Once
	file_com_daml_ledger_api_v2_commands_proto_rawDescData = file_com_daml_ledger_api_v2_commands_proto_rawDesc
)

func file_com_daml_ledger_api_v2_commands_proto_rawDescGZIP() []byte {
	file_com_daml_ledger_api_v2_commands_proto_rawDescOnce.Do(func() {
		file_com_daml_ledger_api_v2_commands_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_daml_ledger_api_v2_commands_proto_rawDescData)
	})
	return file_com_daml_ledger_api_v2_commands_proto_rawDescData
}

var file_com_daml_ledger_api_v2_commands_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_com_daml_ledger_api_v2_commands_proto_goTypes = []any{
	(*Command)(nil),                  // 0: com.daml.ledger.api.v2.Command
	(*CreateCommand)(nil),            // 1: com.daml.ledger.api.v2.CreateCommand
	(*ExerciseCommand)(nil),          // 2: com.daml.ledger.api.v2.ExerciseCommand
	(*ExerciseByKeyCommand)(nil),     // 3: com.daml.ledger.api.v2.ExerciseByKeyCommand
	(*CreateAndExerciseCommand)(nil), // 4: com.daml.ledger.api.v2.CreateAndExerciseCommand
	(*DisclosedContract)(nil),        // 5: com.daml.ledger.api.v2.DisclosedContract
	(*Commands)(nil),                 // 6: com.daml.ledger.api.v2.Commands
	(*Identifier)(nil),               // 7: com.daml.ledger.api.v2.Identifier
	(*Record)(nil),                   // 8: com.daml.ledger.api.v2.Record
	(*Value)(nil),                    // 9: com.daml.ledger.api.v2.Value
	(*durationpb.Duration)(nil),      // 10: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil),    // 11: google.protobuf.Timestamp
}
var file_com_daml_ledger_api_v2_commands_proto_depIdxs = []int32{
	1,  // 0: com.daml.ledger.api.v2.Command.create:type_name -> com.daml.ledger.api.v2.CreateCommand
	2,  // 1: com.daml.ledger.api.v2.Command.exercise:type_name -> com.daml.ledger.api.v2.ExerciseCommand
	3,  // 2: com.daml.ledger.api.v2.Command.exercise_by_key:type_name -> com.daml.ledger.api.v2.ExerciseByKeyCommand
	4,  // 3: com.daml.ledger.api.v2.Command.create_and_exercise:type_name -> com.daml.ledger.api.v2.CreateAndExerciseCommand
	7,  // 4: com.daml.ledger.api.v2.CreateCommand.template_id:type_name -> com.daml.ledger.api.v2.Identifier
	8,  // 5: com.daml.ledger.api.v2.CreateCommand.create_arguments:type_name -> com.daml.ledger.api.v2.Record
	7,  // 6: com.daml.ledger.api.v2.ExerciseCommand.template_id:type_name -> com.daml.ledger.api.v2.Identifier
	9,  // 7: com.daml.ledger.api.v2.ExerciseCommand.choice_argument:type_name -> com.daml.ledger.api.v2.Value
	7,  // 8: com.daml.ledger.api.v2.ExerciseByKeyCommand.template_id:type_name -> com.daml.ledger.api.v2.Identifier
	9,  // 9: com.daml.ledger.api.v2.ExerciseByKeyCommand.contract_key:type_name -> com.daml.ledger.api.v2.Value
	9,  // 10: com.daml.ledger.api.v2.ExerciseByKeyCommand.choice_argument:type_name -> com.daml.ledger.api.v2.Value
	7,  // 11: com.daml.ledger.api.v2.CreateAndExerciseCommand.template_id:type_name -> com.daml.ledger.api.v2.Identifier
	8,  // 12: com.daml.ledger.api.v2.CreateAndExerciseCommand.create_arguments:type_name -> com.daml.ledger.api.v2.Record
	9,  // 13: com.daml.ledger.api.v2.CreateAndExerciseCommand.choice_argument:type_name -> com.daml.ledger.api.v2.Value
	7,  // 14: com.daml.ledger.api.v2.DisclosedContract.template_id:type_name -> com.daml.ledger.api.v2.Identifier
	0,  // 15: com.daml.ledger.api.v2.Commands.commands:type_name -> com.daml.ledger.api.v2.Command
	10, // 16: com.daml.ledger.api.v2.Commands.deduplication_duration:type_name -> google.protobuf.Duration
	11, // 17: com.daml.ledger.api.v2.Commands.min_ledger_time_abs:type_name -> google.protobuf.Timestamp
	10, // 18: com.daml.ledger.api.v2.Commands.min_ledger_time_rel:type_name -> google.protobuf.Duration
	5,  // 19: com.daml.ledger.api.v2.Commands.disclosed_contracts:type_name -> com.daml.ledger.api.v2.DisclosedContract
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_com_daml_ledger_api_v2_commands_proto_init() }
func file_com_daml_ledger_api_v2_commands_proto_init() {
	if File_com_daml_ledger_api_v2_commands_proto != nil {
		return
	}
	file_com_daml_ledger_api_v2_value_proto_init()
	file_com_daml_ledger_api_v2_commands_proto_msgTypes[0].OneofWrappers = []any{
		(*Command_Create)(nil),
		(*Command_Exercise)(nil),
		(*Command_ExerciseByKey)(nil),
		(*Command_CreateAndExercise)(nil),
	}
	file_com_daml_ledger_api_v2_commands_proto_msgTypes[6].OneofWrappers = []any{
		(*Commands_DeduplicationDuration)(nil),
		(*Commands_DeduplicationOffset)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_daml_ledger_api_v2_commands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_daml_ledger_api_v2_commands_proto_goTypes,
		DependencyIndexes: file_com_daml_ledger_api_v2_commands_proto_depIdxs,
		MessageInfos:      file_com_daml_ledger_api_v2_commands_proto_msgTypes,
	}.Build()
	File_com_daml_ledger_api_v2_commands_proto = out.File
	file_com_daml_ledger_api_v2_commands_proto_rawDesc = nil
	file_com_daml_ledger_api_v2_commands_proto_goTypes = nil
	file_com_daml_ledger_api_v2_commands_proto_depIdxs = nil
}
