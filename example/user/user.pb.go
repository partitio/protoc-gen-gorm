// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/user/user.proto

package user // import "github.com/infobloxopen/protoc-gen-gorm/example/user"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import resource "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
import _ "github.com/infobloxopen/protoc-gen-gorm/options"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                   *resource.Identifier `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Birthday             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=birthday" json:"birthday,omitempty"`
	Age                  uint32               `protobuf:"varint,5,opt,name=age" json:"age,omitempty"`
	Num                  uint32               `protobuf:"varint,6,opt,name=num" json:"num,omitempty"`
	CreditCard           *CreditCard          `protobuf:"bytes,7,opt,name=credit_card,json=creditCard" json:"credit_card,omitempty"`
	Emails               []*Email             `protobuf:"bytes,8,rep,name=emails" json:"emails,omitempty"`
	Tasks                []*Task              `protobuf:"bytes,9,rep,name=tasks" json:"tasks,omitempty"`
	BillingAddress       *Address             `protobuf:"bytes,10,opt,name=billing_address,json=billingAddress" json:"billing_address,omitempty"`
	ShippingAddress      *Address             `protobuf:"bytes,11,opt,name=shipping_address,json=shippingAddress" json:"shipping_address,omitempty"`
	Languages            []*Language          `protobuf:"bytes,12,rep,name=languages" json:"languages,omitempty"`
	Friends              []*User              `protobuf:"bytes,13,rep,name=friends" json:"friends,omitempty"`
	ShippingAddressId    *resource.Identifier `protobuf:"bytes,14,opt,name=shipping_address_id,json=shippingAddressId" json:"shipping_address_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_71ee7f1cdb44a642, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
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

type Email struct {
	Id                   *resource.Identifier `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Email                string               `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Subscribed           bool                 `protobuf:"varint,3,opt,name=subscribed" json:"subscribed,omitempty"`
	UserId               *resource.Identifier `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Email) Reset()         { *m = Email{} }
func (m *Email) String() string { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()    {}
func (*Email) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_71ee7f1cdb44a642, []int{1}
}
func (m *Email) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Email.Unmarshal(m, b)
}
func (m *Email) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Email.Marshal(b, m, deterministic)
}
func (dst *Email) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Email.Merge(dst, src)
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

type Address struct {
	Id                   *resource.Identifier `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Address_1            string               `protobuf:"bytes,2,opt,name=address_1,json=address1" json:"address_1,omitempty"`
	Address_2            string               `protobuf:"bytes,3,opt,name=address_2,json=address2" json:"address_2,omitempty"`
	Post                 string               `protobuf:"bytes,4,opt,name=post" json:"post,omitempty"`
	External             *resource.Identifier `protobuf:"bytes,5,opt,name=external" json:"external,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_71ee7f1cdb44a642, []int{2}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (dst *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(dst, src)
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

type Language struct {
	Id                   *resource.Identifier `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Code                 string               `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Language) Reset()         { *m = Language{} }
func (m *Language) String() string { return proto.CompactTextString(m) }
func (*Language) ProtoMessage()    {}
func (*Language) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_71ee7f1cdb44a642, []int{3}
}
func (m *Language) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Language.Unmarshal(m, b)
}
func (m *Language) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Language.Marshal(b, m, deterministic)
}
func (dst *Language) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Language.Merge(dst, src)
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

type CreditCard struct {
	Id                   *resource.Identifier `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Number               string               `protobuf:"bytes,4,opt,name=number" json:"number,omitempty"`
	UserId               *resource.Identifier `protobuf:"bytes,5,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreditCard) Reset()         { *m = CreditCard{} }
func (m *CreditCard) String() string { return proto.CompactTextString(m) }
func (*CreditCard) ProtoMessage()    {}
func (*CreditCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_71ee7f1cdb44a642, []int{4}
}
func (m *CreditCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreditCard.Unmarshal(m, b)
}
func (m *CreditCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreditCard.Marshal(b, m, deterministic)
}
func (dst *CreditCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreditCard.Merge(dst, src)
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
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Priority             int64    `protobuf:"varint,3,opt,name=priority" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_71ee7f1cdb44a642, []int{5}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
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

func init() { proto.RegisterFile("example/user/user.proto", fileDescriptor_user_71ee7f1cdb44a642) }

var fileDescriptor_user_71ee7f1cdb44a642 = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0x4d, 0x6e, 0xeb, 0x36,
	0x10, 0x7e, 0xb2, 0x65, 0x5b, 0x1a, 0x3f, 0x3b, 0x2e, 0xfb, 0xda, 0xca, 0x2e, 0xd0, 0x1a, 0xee,
	0x26, 0x28, 0x6a, 0x09, 0xf1, 0x2b, 0x02, 0xbc, 0x04, 0x28, 0x12, 0x07, 0x59, 0x04, 0xe8, 0x4a,
	0x48, 0x37, 0xdd, 0x18, 0x94, 0x48, 0x2b, 0x6c, 0x24, 0x51, 0x20, 0x29, 0x20, 0xb9, 0x43, 0x8f,
	0xd2, 0x0b, 0xd8, 0x57, 0x69, 0x0f, 0x53, 0x88, 0xfa, 0xb1, 0x1a, 0x20, 0x49, 0x9b, 0x4d, 0x37,
	0xc2, 0x70, 0xe6, 0x9b, 0xe1, 0x37, 0xc3, 0x4f, 0x24, 0x7c, 0x45, 0x1f, 0x70, 0x92, 0xc5, 0xd4,
	0xcb, 0x25, 0x15, 0xfa, 0xe3, 0x66, 0x82, 0x2b, 0x8e, 0xcc, 0xc2, 0x9e, 0x9d, 0x45, 0x4c, 0xdd,
	0xe5, 0x81, 0x1b, 0xf2, 0xc4, 0x63, 0xe9, 0x96, 0x07, 0x31, 0x7f, 0xe0, 0x19, 0x4d, 0x3d, 0x0d,
	0x0a, 0x97, 0x11, 0x4d, 0x97, 0x11, 0x17, 0x89, 0xc7, 0x33, 0xc5, 0x78, 0x2a, 0xbd, 0x62, 0x51,
	0x56, 0x98, 0x7d, 0x1b, 0x71, 0x1e, 0xc5, 0xb4, 0x84, 0x06, 0xf9, 0xd6, 0x53, 0x2c, 0xa1, 0x52,
	0xe1, 0x24, 0xab, 0x00, 0xd7, 0xcf, 0x15, 0xc7, 0x2a, 0xc6, 0x72, 0x89, 0xb3, 0x6c, 0xa9, 0x38,
	0x8f, 0xef, 0x99, 0xf2, 0x44, 0x16, 0x7a, 0x82, 0x4a, 0x9e, 0x8b, 0x90, 0x36, 0x46, 0x59, 0x66,
	0xf1, 0x57, 0x0f, 0xcc, 0x5f, 0x24, 0x15, 0xe8, 0x23, 0x74, 0x18, 0x71, 0x8c, 0xb9, 0x71, 0x3c,
	0x5c, 0x7d, 0xe1, 0xea, 0x22, 0xae, 0xc8, 0x42, 0xf7, 0x86, 0xd0, 0x54, 0xb1, 0x2d, 0xa3, 0x62,
	0xfd, 0x7e, 0xbf, 0x9b, 0x5a, 0xd0, 0x47, 0x66, 0x9e, 0x33, 0xe2, 0x77, 0x18, 0x41, 0x9f, 0x00,
	0x42, 0x41, 0xb1, 0xa2, 0x64, 0x83, 0x95, 0xd3, 0xd1, 0xc9, 0x33, 0xb7, 0xa4, 0xee, 0xd6, 0xd4,
	0xdd, 0xdb, 0x9a, 0xba, 0x6f, 0x57, 0xe8, 0x4b, 0x55, 0xa4, 0xe6, 0x19, 0xa9, 0x53, 0xbb, 0xaf,
	0xa7, 0x56, 0xe8, 0x4b, 0x85, 0x4e, 0xc1, 0x0a, 0x98, 0x50, 0x77, 0x04, 0x3f, 0x3a, 0xe6, 0xab,
	0x89, 0x0d, 0x16, 0x39, 0xd0, 0xc5, 0x11, 0x75, 0x7a, 0x73, 0xe3, 0x78, 0xb4, 0xee, 0xef, 0x77,
	0xd3, 0xce, 0xc4, 0xf0, 0x0b, 0x17, 0x9a, 0x40, 0x37, 0xcd, 0x13, 0xa7, 0x5f, 0x44, 0xfc, 0xc2,
	0x44, 0x27, 0x30, 0x0c, 0x05, 0x25, 0x4c, 0x6d, 0x42, 0x2c, 0x88, 0x33, 0xd0, 0xdb, 0x4c, 0x5c,
	0x7d, 0xc6, 0x57, 0x3a, 0x70, 0x85, 0x05, 0xf1, 0x21, 0x6c, 0x6c, 0xf4, 0x1d, 0xf4, 0x69, 0x82,
	0x59, 0x2c, 0x1d, 0x6b, 0xde, 0x3d, 0x1e, 0xae, 0x86, 0x25, 0xfa, 0xba, 0xf0, 0xf9, 0x55, 0x08,
	0x9d, 0x42, 0x4f, 0x61, 0x79, 0x2f, 0x1d, 0x5b, 0x63, 0xa0, 0xc4, 0xdc, 0x62, 0x79, 0xbf, 0xfe,
	0xb0, 0xdf, 0x4d, 0x27, 0xdf, 0x8f, 0x51, 0xe7, 0xc2, 0x58, 0x58, 0x99, 0x60, 0x5c, 0x30, 0xf5,
	0xe8, 0x97, 0x70, 0xf4, 0x13, 0x1c, 0x05, 0x2c, 0x8e, 0x59, 0x1a, 0x6d, 0x30, 0x21, 0x82, 0x4a,
	0xe9, 0x80, 0xe6, 0x34, 0x2a, 0x2b, 0x5c, 0x96, 0xce, 0xb2, 0xad, 0xc5, 0x3b, 0x7f, 0x5c, 0xa1,
	0x2b, 0x3f, 0xba, 0x80, 0x89, 0xbc, 0x63, 0x59, 0xd6, 0x2e, 0x30, 0x7c, 0xa9, 0xc0, 0x51, 0x0d,
	0xaf, 0x2b, 0xfc, 0x08, 0x76, 0x8c, 0xd3, 0x28, 0xc7, 0x11, 0x95, 0xce, 0x7b, 0xcd, 0x7e, 0x5c,
	0xa6, 0xfe, 0x5c, 0xb9, 0xcb, 0xdc, 0xd5, 0x3b, 0xff, 0x00, 0x44, 0x3f, 0xc0, 0x60, 0x2b, 0x18,
	0x4d, 0x89, 0x74, 0x46, 0xed, 0x8e, 0x0b, 0xcd, 0x35, 0xf8, 0x1a, 0x82, 0xae, 0xe1, 0xf3, 0xa7,
	0x2c, 0x37, 0x8c, 0x38, 0xe3, 0x17, 0x54, 0xe9, 0x7f, 0xf6, 0x84, 0xe8, 0x0d, 0x39, 0xb3, 0xf6,
	0xbb, 0xa9, 0x69, 0x19, 0x73, 0x63, 0xf1, 0x87, 0x01, 0x3d, 0x7d, 0x00, 0x6f, 0xd3, 0xf7, 0x07,
	0xe8, 0xe9, 0x73, 0xd3, 0xd2, 0xb6, 0xfd, 0x72, 0x81, 0xbe, 0x01, 0x90, 0x79, 0x20, 0x43, 0xc1,
	0x02, 0x4a, 0xb4, 0x74, 0x2d, 0xbf, 0xe5, 0x41, 0x2e, 0x0c, 0x8a, 0x1e, 0x0b, 0xe6, 0xe6, 0x4b,
	0xcc, 0xfb, 0x05, 0xea, 0x1f, 0x74, 0xff, 0x34, 0x60, 0x50, 0xcf, 0xfb, 0xf4, 0x75, 0xc2, 0x47,
	0xfb, 0xdd, 0x74, 0x08, 0x36, 0x1a, 0xb0, 0x54, 0xd1, 0x88, 0x0a, 0xcd, 0xf9, 0x6b, 0xb0, 0xeb,
	0xd1, 0x9d, 0x54, 0xbc, 0xad, 0xca, 0x71, 0xd2, 0x0e, 0xae, 0x34, 0xf3, 0x43, 0x70, 0x85, 0x10,
	0x98, 0x19, 0x97, 0x4a, 0x93, 0xb6, 0x7d, 0x6d, 0xa3, 0x0b, 0xb0, 0xe8, 0x83, 0xa2, 0x22, 0xc5,
	0xb1, 0xfe, 0x71, 0x9e, 0xe5, 0x32, 0xda, 0xef, 0xa6, 0x36, 0x0c, 0x50, 0xef, 0x37, 0xc9, 0xd3,
	0xc0, 0x6f, 0xb2, 0x5a, 0xdd, 0x29, 0xb0, 0x6a, 0xa9, 0xbc, 0xb9, 0x3b, 0x04, 0x66, 0x8a, 0x13,
	0x5a, 0x35, 0xa6, 0xed, 0xc2, 0x17, 0x72, 0x42, 0xab, 0x7e, 0xb4, 0xdd, 0xda, 0xf5, 0xf7, 0x0e,
	0xc0, 0xe1, 0x8f, 0x7d, 0xf3, 0xc6, 0xff, 0xcf, 0x55, 0xf7, 0x25, 0xf4, 0xd3, 0x3c, 0x09, 0xa8,
	0xa8, 0x0e, 0xa5, 0x5a, 0xb5, 0x25, 0xd6, 0xfb, 0x6f, 0x12, 0x0b, 0xc0, 0x2c, 0x6e, 0x9b, 0x66,
	0x90, 0x46, 0x6b, 0x90, 0x73, 0x18, 0x12, 0x5a, 0xa8, 0x58, 0xbf, 0x47, 0xd5, 0x8c, 0xdb, 0x2e,
	0x34, 0x83, 0xe6, 0x66, 0xd2, 0x8d, 0x74, 0xfd, 0x66, 0x7d, 0xd8, 0x63, 0x7d, 0xfe, 0xeb, 0xa7,
	0x7f, 0xfb, 0xf4, 0xb5, 0x5f, 0xd0, 0xf3, 0xe2, 0x13, 0xf4, 0x35, 0xe4, 0xe3, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x8c, 0xac, 0x67, 0x7d, 0x5d, 0x07, 0x00, 0x00,
}
