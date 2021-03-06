// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	resource "github.com/partitio/atlas-app-toolkit/rpc/resource"
	_ "github.com/partitio/protoc-gen-gorm/options"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   *resource.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Birthday             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Age                  uint32               `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	Num                  uint32               `protobuf:"varint,6,opt,name=num,proto3" json:"num,omitempty"`
	CreditCard           *CreditCard          `protobuf:"bytes,7,opt,name=credit_card,json=creditCard,proto3" json:"credit_card,omitempty"`
	Emails               []*Email             `protobuf:"bytes,8,rep,name=emails,proto3" json:"emails,omitempty"`
	Tasks                []*Task              `protobuf:"bytes,9,rep,name=tasks,proto3" json:"tasks,omitempty"`
	BillingAddress       *Address             `protobuf:"bytes,10,opt,name=billing_address,json=billingAddress,proto3" json:"billing_address,omitempty"`
	ShippingAddress      *Address             `protobuf:"bytes,11,opt,name=shipping_address,json=shippingAddress,proto3" json:"shipping_address,omitempty"`
	Languages            []*Language          `protobuf:"bytes,12,rep,name=languages,proto3" json:"languages,omitempty"`
	Friends              []*User              `protobuf:"bytes,13,rep,name=friends,proto3" json:"friends,omitempty"`
	ShippingAddressId    *resource.Identifier `protobuf:"bytes,14,opt,name=shipping_address_id,json=shippingAddressId,proto3" json:"shipping_address_id,omitempty"`
	ExternalUuid         *resource.Identifier `protobuf:"bytes,15,opt,name=external_uuid,json=externalUuid,proto3" json:"external_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a81b87a58f0b229, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() *resource.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *User) GetBirthday() *timestamp.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *User) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetNum() uint32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *User) GetCreditCard() *CreditCard {
	if m != nil {
		return m.CreditCard
	}
	return nil
}

func (m *User) GetEmails() []*Email {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *User) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *User) GetBillingAddress() *Address {
	if m != nil {
		return m.BillingAddress
	}
	return nil
}

func (m *User) GetShippingAddress() *Address {
	if m != nil {
		return m.ShippingAddress
	}
	return nil
}

func (m *User) GetLanguages() []*Language {
	if m != nil {
		return m.Languages
	}
	return nil
}

func (m *User) GetFriends() []*User {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *User) GetShippingAddressId() *resource.Identifier {
	if m != nil {
		return m.ShippingAddressId
	}
	return nil
}

func (m *User) GetExternalUuid() *resource.Identifier {
	if m != nil {
		return m.ExternalUuid
	}
	return nil
}

type Email struct {
	Id                   *resource.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Subscribed           bool                 `protobuf:"varint,3,opt,name=subscribed,proto3" json:"subscribed,omitempty"`
	UserId               *resource.Identifier `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ExternalNotNull      *resource.Identifier `protobuf:"bytes,5,opt,name=external_not_null,json=externalNotNull,proto3" json:"external_not_null,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Email) Reset()         { *m = Email{} }
func (m *Email) String() string { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()    {}
func (*Email) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a81b87a58f0b229, []int{1}
}

func (m *Email) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Email.Unmarshal(m, b)
}
func (m *Email) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Email.Marshal(b, m, deterministic)
}
func (m *Email) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Email.Merge(m, src)
}
func (m *Email) XXX_Size() int {
	return xxx_messageInfo_Email.Size(m)
}
func (m *Email) XXX_DiscardUnknown() {
	xxx_messageInfo_Email.DiscardUnknown(m)
}

var xxx_messageInfo_Email proto.InternalMessageInfo

func (m *Email) GetId() *resource.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Email) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Email) GetSubscribed() bool {
	if m != nil {
		return m.Subscribed
	}
	return false
}

func (m *Email) GetUserId() *resource.Identifier {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Email) GetExternalNotNull() *resource.Identifier {
	if m != nil {
		return m.ExternalNotNull
	}
	return nil
}

type Address struct {
	Id                   *resource.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address_1            string               `protobuf:"bytes,2,opt,name=address_1,json=address1,proto3" json:"address_1,omitempty"`
	Address_2            string               `protobuf:"bytes,3,opt,name=address_2,json=address2,proto3" json:"address_2,omitempty"`
	Post                 string               `protobuf:"bytes,4,opt,name=post,proto3" json:"post,omitempty"`
	External             *resource.Identifier `protobuf:"bytes,5,opt,name=external,proto3" json:"external,omitempty"`
	ImplicitFk           *resource.Identifier `protobuf:"bytes,6,opt,name=implicit_fk,json=implicitFk,proto3" json:"implicit_fk,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a81b87a58f0b229, []int{2}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetId() *resource.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Address) GetAddress_1() string {
	if m != nil {
		return m.Address_1
	}
	return ""
}

func (m *Address) GetAddress_2() string {
	if m != nil {
		return m.Address_2
	}
	return ""
}

func (m *Address) GetPost() string {
	if m != nil {
		return m.Post
	}
	return ""
}

func (m *Address) GetExternal() *resource.Identifier {
	if m != nil {
		return m.External
	}
	return nil
}

func (m *Address) GetImplicitFk() *resource.Identifier {
	if m != nil {
		return m.ImplicitFk
	}
	return nil
}

type Language struct {
	Id                   *resource.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string               `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	ExternalInt          *resource.Identifier `protobuf:"bytes,4,opt,name=external_int,json=externalInt,proto3" json:"external_int,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Language) Reset()         { *m = Language{} }
func (m *Language) String() string { return proto.CompactTextString(m) }
func (*Language) ProtoMessage()    {}
func (*Language) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a81b87a58f0b229, []int{3}
}

func (m *Language) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Language.Unmarshal(m, b)
}
func (m *Language) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Language.Marshal(b, m, deterministic)
}
func (m *Language) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Language.Merge(m, src)
}
func (m *Language) XXX_Size() int {
	return xxx_messageInfo_Language.Size(m)
}
func (m *Language) XXX_DiscardUnknown() {
	xxx_messageInfo_Language.DiscardUnknown(m)
}

var xxx_messageInfo_Language proto.InternalMessageInfo

func (m *Language) GetId() *resource.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Language) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Language) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Language) GetExternalInt() *resource.Identifier {
	if m != nil {
		return m.ExternalInt
	}
	return nil
}

type CreditCard struct {
	Id                   *resource.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Number               string               `protobuf:"bytes,4,opt,name=number,proto3" json:"number,omitempty"`
	UserId               *resource.Identifier `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreditCard) Reset()         { *m = CreditCard{} }
func (m *CreditCard) String() string { return proto.CompactTextString(m) }
func (*CreditCard) ProtoMessage()    {}
func (*CreditCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a81b87a58f0b229, []int{4}
}

func (m *CreditCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreditCard.Unmarshal(m, b)
}
func (m *CreditCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreditCard.Marshal(b, m, deterministic)
}
func (m *CreditCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreditCard.Merge(m, src)
}
func (m *CreditCard) XXX_Size() int {
	return xxx_messageInfo_CreditCard.Size(m)
}
func (m *CreditCard) XXX_DiscardUnknown() {
	xxx_messageInfo_CreditCard.DiscardUnknown(m)
}

var xxx_messageInfo_CreditCard proto.InternalMessageInfo

func (m *CreditCard) GetId() *resource.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CreditCard) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *CreditCard) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *CreditCard) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *CreditCard) GetUserId() *resource.Identifier {
	if m != nil {
		return m.UserId
	}
	return nil
}

type Task struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Priority             int64    `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a81b87a58f0b229, []int{5}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Task) GetPriority() int64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Email)(nil), "user.Email")
	proto.RegisterType((*Address)(nil), "user.Address")
	proto.RegisterType((*Language)(nil), "user.Language")
	proto.RegisterType((*CreditCard)(nil), "user.CreditCard")
	proto.RegisterType((*Task)(nil), "user.Task")
}

func init() { proto.RegisterFile("example/user/user.proto", fileDescriptor_5a81b87a58f0b229) }

var fileDescriptor_5a81b87a58f0b229 = []byte{
	// 885 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xdd, 0x6e, 0xe4, 0x34,
	0x14, 0xde, 0x4c, 0xe7, 0x27, 0x73, 0xd2, 0xf9, 0xa9, 0x77, 0x81, 0x4c, 0x91, 0x60, 0x54, 0x6e,
	0x2a, 0x44, 0x33, 0xea, 0xc0, 0x16, 0x58, 0x24, 0xd4, 0xce, 0x6a, 0x11, 0x23, 0xd0, 0x4a, 0x58,
	0xbb, 0x37, 0xdc, 0x8c, 0x9c, 0xd8, 0x4d, 0xcd, 0x24, 0x71, 0x64, 0x3b, 0x52, 0xf7, 0x92, 0x17,
	0xe0, 0x39, 0x10, 0x6f, 0x30, 0xb9, 0xe3, 0xcd, 0x50, 0x9c, 0x9f, 0x46, 0x2b, 0x98, 0x56, 0xec,
	0xc5, 0xde, 0x44, 0xc7, 0xc7, 0xdf, 0x39, 0x39, 0xe7, 0xf3, 0x77, 0x9c, 0xc0, 0x47, 0xec, 0x96,
	0xc4, 0x69, 0xc4, 0x16, 0x99, 0x62, 0xd2, 0x3c, 0xbc, 0x54, 0x0a, 0x2d, 0x50, 0xb7, 0xb0, 0x8f,
	0x2f, 0x42, 0xae, 0x6f, 0x32, 0xdf, 0x0b, 0x44, 0xbc, 0x48, 0x89, 0xd4, 0x5c, 0x73, 0xb1, 0x30,
	0x80, 0xe0, 0x2c, 0x64, 0xc9, 0x59, 0x28, 0x64, 0xbc, 0x10, 0xa9, 0xe6, 0x22, 0x51, 0x8b, 0x62,
	0x51, 0x46, 0x1f, 0x7f, 0x1a, 0x0a, 0x11, 0x46, 0xac, 0x84, 0xfa, 0xd9, 0xf5, 0x42, 0xf3, 0x98,
	0x29, 0x4d, 0xe2, 0xb4, 0x02, 0x5c, 0xfd, 0x5b, 0x62, 0xa2, 0x23, 0xa2, 0xce, 0x48, 0x9a, 0x9e,
	0x69, 0x21, 0xa2, 0x2d, 0xd7, 0x0b, 0x99, 0x06, 0x0b, 0xc9, 0x94, 0xc8, 0x64, 0xc0, 0x1a, 0xa3,
	0x4c, 0x71, 0xf2, 0x67, 0x1f, 0xba, 0xaf, 0x15, 0x93, 0xe8, 0x29, 0x74, 0x38, 0x75, 0xad, 0xb9,
	0x75, 0xea, 0x2c, 0x3f, 0xf0, 0x4c, 0x12, 0x4f, 0xa6, 0x81, 0xb7, 0xa6, 0x2c, 0xd1, 0xfc, 0x9a,
	0x33, 0xb9, 0x1a, 0xe7, 0xbb, 0x19, 0x80, 0x8d, 0xba, 0x59, 0xc6, 0xe9, 0xa9, 0x85, 0x3b, 0x9c,
	0xa2, 0x6f, 0x01, 0x02, 0xc9, 0x88, 0x66, 0x74, 0x43, 0xb4, 0xdb, 0x31, 0xe1, 0xc7, 0x5e, 0x59,
	0xb8, 0x57, 0x17, 0xee, 0xbd, 0xaa, 0x0b, 0xc7, 0xc3, 0x0a, 0x7d, 0xa5, 0x8b, 0xd0, 0x2c, 0xa5,
	0x75, 0xe8, 0xc1, 0xfd, 0xa1, 0x15, 0xfa, 0x4a, 0xa3, 0x0b, 0xb0, 0x7d, 0x2e, 0xf5, 0x0d, 0x25,
	0x6f, 0xdc, 0xee, 0xbd, 0x81, 0x0d, 0x16, 0xb9, 0x70, 0x40, 0x42, 0xe6, 0xf6, 0xe6, 0xd6, 0xe9,
	0x68, 0xd5, 0xcf, 0x77, 0xb3, 0xce, 0xd4, 0xc2, 0x85, 0x0b, 0x4d, 0xe1, 0x20, 0xc9, 0x62, 0xb7,
	0x5f, 0xec, 0xe0, 0xc2, 0x44, 0xe7, 0xe0, 0x04, 0x92, 0x51, 0xae, 0x37, 0x01, 0x91, 0xd4, 0x1d,
	0x98, 0xd7, 0x4c, 0x3d, 0x73, 0xba, 0xcf, 0xcd, 0xc6, 0x73, 0x22, 0x29, 0x86, 0xa0, 0xb1, 0xd1,
	0x67, 0xd0, 0x67, 0x31, 0xe1, 0x91, 0x72, 0xed, 0xf9, 0xc1, 0xa9, 0xb3, 0x74, 0x4a, 0xf4, 0x8b,
	0xc2, 0x87, 0xab, 0x2d, 0x74, 0x01, 0x3d, 0x4d, 0xd4, 0x56, 0xb9, 0x43, 0x83, 0x81, 0x12, 0xf3,
	0x8a, 0xa8, 0xed, 0xea, 0x49, 0xbe, 0x9b, 0x4d, 0x3f, 0x1f, 0xa3, 0xce, 0xa5, 0x75, 0x62, 0xa7,
	0x92, 0x0b, 0xc9, 0xf5, 0x1b, 0x5c, 0xc2, 0xd1, 0xf7, 0x30, 0xf1, 0x79, 0x14, 0xf1, 0x24, 0xdc,
	0x10, 0x4a, 0x25, 0x53, 0xca, 0x05, 0x53, 0xd3, 0xa8, 0xcc, 0x70, 0x55, 0x3a, 0xcb, 0xb6, 0x4e,
	0x1e, 0xe1, 0x71, 0x85, 0xae, 0xfc, 0xe8, 0x12, 0xa6, 0xea, 0x86, 0xa7, 0x69, 0x3b, 0x81, 0xb3,
	0x2f, 0xc1, 0xa4, 0x86, 0xd7, 0x19, 0xbe, 0x82, 0x61, 0x44, 0x92, 0x30, 0x23, 0x21, 0x53, 0xee,
	0xa1, 0xa9, 0x7e, 0x5c, 0x86, 0xfe, 0x5c, 0xb9, 0xcb, 0xd8, 0xe5, 0x23, 0x7c, 0x07, 0x44, 0x5f,
	0xc0, 0xe0, 0x5a, 0x72, 0x96, 0x50, 0xe5, 0x8e, 0xda, 0x1d, 0x17, 0xaa, 0x6b, 0xf0, 0x35, 0x04,
	0xbd, 0x80, 0xc7, 0x6f, 0x57, 0xb9, 0xe1, 0xd4, 0x1d, 0xef, 0xd1, 0x25, 0x3e, 0x7a, 0xab, 0xd0,
	0x35, 0x45, 0x3f, 0xc2, 0x88, 0xdd, 0x6a, 0x26, 0x13, 0x12, 0x6d, 0x0a, 0xb5, 0xba, 0x93, 0x7d,
	0xc2, 0x3e, 0xcc, 0x77, 0x33, 0x1b, 0xfa, 0xa5, 0xb0, 0xf1, 0x61, 0x1d, 0xf9, 0x3a, 0xe3, 0xf4,
	0x99, 0x9d, 0xef, 0x66, 0x5d, 0xdb, 0x9a, 0x5b, 0x27, 0xbf, 0x77, 0xa0, 0x67, 0x8e, 0xf2, 0xff,
	0xce, 0xca, 0x13, 0xe8, 0x19, 0x0d, 0x98, 0x31, 0x19, 0xe2, 0x72, 0x81, 0x3e, 0x01, 0x50, 0x99,
	0xaf, 0x02, 0xc9, 0x7d, 0x46, 0xcd, 0x18, 0xd8, 0xb8, 0xe5, 0x41, 0x1e, 0x0c, 0x0a, 0xbe, 0x0a,
	0x16, 0xba, 0xfb, 0x58, 0xe8, 0x17, 0xa8, 0x35, 0x45, 0xbf, 0xc0, 0x51, 0xd3, 0x7a, 0x22, 0xf4,
	0x26, 0xc9, 0xa2, 0xc8, 0x28, 0xfe, 0x61, 0xb5, 0x5e, 0x5a, 0x78, 0x52, 0xc7, 0xbf, 0x14, 0xfa,
	0x65, 0x16, 0x45, 0x2d, 0x0e, 0xfe, 0xea, 0xc0, 0xa0, 0x96, 0xc3, 0x37, 0xf7, 0xb3, 0x70, 0x94,
	0xef, 0x66, 0x23, 0x70, 0xd0, 0x80, 0x27, 0x9a, 0x85, 0x4c, 0x56, 0x44, 0x7c, 0x0c, 0xc3, 0xfa,
	0x6c, 0xcf, 0x2b, 0x32, 0xec, 0xca, 0x71, 0xde, 0xde, 0x5c, 0x1a, 0x3a, 0xee, 0x36, 0x97, 0x08,
	0x41, 0x37, 0x15, 0x4a, 0x1b, 0x26, 0x86, 0xd8, 0xd8, 0xe8, 0x12, 0xec, 0xba, 0xe0, 0xfd, 0x7d,
	0x8e, 0xf2, 0xdd, 0x6c, 0x08, 0x03, 0xd4, 0xfb, 0x4d, 0x89, 0xc4, 0xc7, 0x4d, 0x14, 0xfa, 0x09,
	0x1c, 0x1e, 0xa7, 0x11, 0x0f, 0xb8, 0xde, 0x5c, 0x6f, 0xcd, 0x25, 0xf0, 0x9f, 0x49, 0x1e, 0xe7,
	0xbb, 0xd9, 0xa4, 0xd0, 0x8a, 0x66, 0xb7, 0xfa, 0x59, 0x29, 0x08, 0x0c, 0x75, 0xf8, 0x0f, 0xdb,
	0x16, 0x59, 0x7f, 0x5b, 0x60, 0xd7, 0x93, 0xf1, 0x0e, 0x6c, 0x21, 0xe8, 0x26, 0x24, 0x66, 0x15,
	0x51, 0xc6, 0x2e, 0x7c, 0x81, 0xa0, 0xac, 0xe2, 0xc7, 0xd8, 0x68, 0x0d, 0x8d, 0x72, 0x37, 0x3c,
	0xd1, 0x7b, 0xd5, 0xb2, 0x9a, 0xe4, 0xbb, 0x99, 0x03, 0xc3, 0xe6, 0x5d, 0xd8, 0xa9, 0x63, 0xd7,
	0x89, 0x6e, 0xf5, 0xf0, 0x47, 0x07, 0xe0, 0xee, 0xb6, 0x7b, 0x87, 0x2e, 0xde, 0xcf, 0x87, 0xe2,
	0x43, 0xe8, 0x27, 0x59, 0xec, 0x33, 0x59, 0x29, 0xa6, 0x5a, 0xb5, 0x87, 0xaa, 0xf7, 0x80, 0xa1,
	0x6a, 0x11, 0xe2, 0x43, 0xb7, 0xb8, 0xab, 0x9b, 0x53, 0xb1, 0x5a, 0xa7, 0x32, 0x07, 0x87, 0xb2,
	0x62, 0x6e, 0xcd, 0xb7, 0xbc, 0x3a, 0xb0, 0xb6, 0x0b, 0x1d, 0x43, 0x73, 0xaf, 0x9b, 0x46, 0x0e,
	0x70, 0xb3, 0xbe, 0x7b, 0xc7, 0xea, 0xeb, 0x5f, 0x9f, 0x3e, 0xe4, 0x97, 0xa1, 0xfd, 0xd7, 0xf1,
	0x5d, 0xf1, 0xf0, 0xfb, 0x06, 0xf2, 0xe5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x64, 0xf3,
	0xa3, 0x91, 0x08, 0x00, 0x00,
}
