// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: space/employees.proto

package spacev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Employee struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	Id        int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	// optional bytes image_uuid = 6;
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	BlockedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=blocked_at,json=blockedAt,proto3,oneof" json:"blocked_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Employee) Reset() {
	*x = Employee{}
	mi := &file_space_employees_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_space_employees_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_space_employees_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Employee) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Employee) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Employee) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Employee) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Employee) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Employee) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Employee) GetBlockedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.BlockedAt
	}
	return nil
}

type GetEmployeesParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEmployeesParams) Reset() {
	*x = GetEmployeesParams{}
	mi := &file_space_employees_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEmployeesParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmployeesParams) ProtoMessage() {}

func (x *GetEmployeesParams) ProtoReflect() protoreflect.Message {
	mi := &file_space_employees_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmployeesParams.ProtoReflect.Descriptor instead.
func (*GetEmployeesParams) Descriptor() ([]byte, []int) {
	return file_space_employees_proto_rawDescGZIP(), []int{1}
}

type GetEmployeesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*Employee            `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEmployeesResponse) Reset() {
	*x = GetEmployeesResponse{}
	mi := &file_space_employees_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEmployeesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmployeesResponse) ProtoMessage() {}

func (x *GetEmployeesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_space_employees_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmployeesResponse.ProtoReflect.Descriptor instead.
func (*GetEmployeesResponse) Descriptor() ([]byte, []int) {
	return file_space_employees_proto_rawDescGZIP(), []int{2}
}

func (x *GetEmployeesResponse) GetList() []*Employee {
	if x != nil {
		return x.List
	}
	return nil
}

type CreateEmployeeParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName     string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email         string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateEmployeeParams) Reset() {
	*x = CreateEmployeeParams{}
	mi := &file_space_employees_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateEmployeeParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmployeeParams) ProtoMessage() {}

func (x *CreateEmployeeParams) ProtoReflect() protoreflect.Message {
	mi := &file_space_employees_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmployeeParams.ProtoReflect.Descriptor instead.
func (*CreateEmployeeParams) Descriptor() ([]byte, []int) {
	return file_space_employees_proto_rawDescGZIP(), []int{3}
}

func (x *CreateEmployeeParams) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateEmployeeParams) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateEmployeeParams) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateEmployeeParams) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UpdateEmployeeParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName     string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"` //  optional bytes image_uuid = 5;
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateEmployeeParams) Reset() {
	*x = UpdateEmployeeParams{}
	mi := &file_space_employees_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEmployeeParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmployeeParams) ProtoMessage() {}

func (x *UpdateEmployeeParams) ProtoReflect() protoreflect.Message {
	mi := &file_space_employees_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmployeeParams.ProtoReflect.Descriptor instead.
func (*UpdateEmployeeParams) Descriptor() ([]byte, []int) {
	return file_space_employees_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEmployeeParams) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateEmployeeParams) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateEmployeeParams) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateEmployeeParams) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CheckEmployeeAvailableUsernameParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckEmployeeAvailableUsernameParams) Reset() {
	*x = CheckEmployeeAvailableUsernameParams{}
	mi := &file_space_employees_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckEmployeeAvailableUsernameParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEmployeeAvailableUsernameParams) ProtoMessage() {}

func (x *CheckEmployeeAvailableUsernameParams) ProtoReflect() protoreflect.Message {
	mi := &file_space_employees_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEmployeeAvailableUsernameParams.ProtoReflect.Descriptor instead.
func (*CheckEmployeeAvailableUsernameParams) Descriptor() ([]byte, []int) {
	return file_space_employees_proto_rawDescGZIP(), []int{5}
}

func (x *CheckEmployeeAvailableUsernameParams) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type BlockEmployeeParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BlockEmployeeParams) Reset() {
	*x = BlockEmployeeParams{}
	mi := &file_space_employees_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlockEmployeeParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockEmployeeParams) ProtoMessage() {}

func (x *BlockEmployeeParams) ProtoReflect() protoreflect.Message {
	mi := &file_space_employees_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockEmployeeParams.ProtoReflect.Descriptor instead.
func (*BlockEmployeeParams) Descriptor() ([]byte, []int) {
	return file_space_employees_proto_rawDescGZIP(), []int{6}
}

func (x *BlockEmployeeParams) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_space_employees_proto protoreflect.FileDescriptor

const file_space_employees_proto_rawDesc = "" +
	"\n" +
	"\x15space/employees.proto\x12\x05space\x1a\x1fgoogle/protobuf/timestamp.proto\"\xcd\x02\n" +
	"\bEmployee\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x1d\n" +
	"\n" +
	"first_name\x18\x03 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x04 \x01(\tR\blastName\x12\x14\n" +
	"\x05email\x18\x05 \x01(\tR\x05email\x129\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12>\n" +
	"\n" +
	"blocked_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampH\x00R\tblockedAt\x88\x01\x01B\r\n" +
	"\v_blocked_at\"\x14\n" +
	"\x12GetEmployeesParams\";\n" +
	"\x14GetEmployeesResponse\x12#\n" +
	"\x04list\x18\x01 \x03(\v2\x0f.space.EmployeeR\x04list\"\x84\x01\n" +
	"\x14CreateEmployeeParams\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x1d\n" +
	"\n" +
	"first_name\x18\x03 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x04 \x01(\tR\blastName\x12\x14\n" +
	"\x05email\x18\x05 \x01(\tR\x05email\"x\n" +
	"\x14UpdateEmployeeParams\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x1d\n" +
	"\n" +
	"first_name\x18\x02 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x03 \x01(\tR\blastName\x12\x14\n" +
	"\x05email\x18\x04 \x01(\tR\x05email\"B\n" +
	"$CheckEmployeeAvailableUsernameParams\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\"%\n" +
	"\x13BlockEmployeeParams\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02idB4Z2github.com/market-4-test/contract-go/space;spacev1b\x06proto3"

var (
	file_space_employees_proto_rawDescOnce sync.Once
	file_space_employees_proto_rawDescData []byte
)

func file_space_employees_proto_rawDescGZIP() []byte {
	file_space_employees_proto_rawDescOnce.Do(func() {
		file_space_employees_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_space_employees_proto_rawDesc), len(file_space_employees_proto_rawDesc)))
	})
	return file_space_employees_proto_rawDescData
}

var file_space_employees_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_space_employees_proto_goTypes = []any{
	(*Employee)(nil),                             // 0: space.Employee
	(*GetEmployeesParams)(nil),                   // 1: space.GetEmployeesParams
	(*GetEmployeesResponse)(nil),                 // 2: space.GetEmployeesResponse
	(*CreateEmployeeParams)(nil),                 // 3: space.CreateEmployeeParams
	(*UpdateEmployeeParams)(nil),                 // 4: space.UpdateEmployeeParams
	(*CheckEmployeeAvailableUsernameParams)(nil), // 5: space.CheckEmployeeAvailableUsernameParams
	(*BlockEmployeeParams)(nil),                  // 6: space.BlockEmployeeParams
	(*timestamppb.Timestamp)(nil),                // 7: google.protobuf.Timestamp
}
var file_space_employees_proto_depIdxs = []int32{
	7, // 0: space.Employee.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: space.Employee.updated_at:type_name -> google.protobuf.Timestamp
	7, // 2: space.Employee.blocked_at:type_name -> google.protobuf.Timestamp
	0, // 3: space.GetEmployeesResponse.list:type_name -> space.Employee
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_space_employees_proto_init() }
func file_space_employees_proto_init() {
	if File_space_employees_proto != nil {
		return
	}
	file_space_employees_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_space_employees_proto_rawDesc), len(file_space_employees_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_space_employees_proto_goTypes,
		DependencyIndexes: file_space_employees_proto_depIdxs,
		MessageInfos:      file_space_employees_proto_msgTypes,
	}.Build()
	File_space_employees_proto = out.File
	file_space_employees_proto_goTypes = nil
	file_space_employees_proto_depIdxs = nil
}
