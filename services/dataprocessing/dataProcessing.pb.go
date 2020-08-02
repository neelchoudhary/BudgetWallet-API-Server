// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/dataprocessing/dataProcessing.proto

package dataprocessing

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AccountDailySnapshot struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	StartDayBalance      float64  `protobuf:"fixed64,2,opt,name=start_day_balance,json=startDayBalance,proto3" json:"start_day_balance,omitempty"`
	EndDayBalance        float64  `protobuf:"fixed64,3,opt,name=end_day_balance,json=endDayBalance,proto3" json:"end_day_balance,omitempty"`
	CashOut              float64  `protobuf:"fixed64,4,opt,name=cash_out,json=cashOut,proto3" json:"cash_out,omitempty"`
	CashIn               float64  `protobuf:"fixed64,5,opt,name=cash_in,json=cashIn,proto3" json:"cash_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountDailySnapshot) Reset()         { *m = AccountDailySnapshot{} }
func (m *AccountDailySnapshot) String() string { return proto.CompactTextString(m) }
func (*AccountDailySnapshot) ProtoMessage()    {}
func (*AccountDailySnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d0ea3f8dbe2a8, []int{0}
}

func (m *AccountDailySnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountDailySnapshot.Unmarshal(m, b)
}
func (m *AccountDailySnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountDailySnapshot.Marshal(b, m, deterministic)
}
func (m *AccountDailySnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountDailySnapshot.Merge(m, src)
}
func (m *AccountDailySnapshot) XXX_Size() int {
	return xxx_messageInfo_AccountDailySnapshot.Size(m)
}
func (m *AccountDailySnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountDailySnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_AccountDailySnapshot proto.InternalMessageInfo

func (m *AccountDailySnapshot) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *AccountDailySnapshot) GetStartDayBalance() float64 {
	if m != nil {
		return m.StartDayBalance
	}
	return 0
}

func (m *AccountDailySnapshot) GetEndDayBalance() float64 {
	if m != nil {
		return m.EndDayBalance
	}
	return 0
}

func (m *AccountDailySnapshot) GetCashOut() float64 {
	if m != nil {
		return m.CashOut
	}
	return 0
}

func (m *AccountDailySnapshot) GetCashIn() float64 {
	if m != nil {
		return m.CashIn
	}
	return 0
}

type AccountMonthlySnapshot struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	StartMonthBalance    float64  `protobuf:"fixed64,2,opt,name=start_month_balance,json=startMonthBalance,proto3" json:"start_month_balance,omitempty"`
	EndMonthBalance      float64  `protobuf:"fixed64,3,opt,name=end_month_balance,json=endMonthBalance,proto3" json:"end_month_balance,omitempty"`
	CashOut              float64  `protobuf:"fixed64,4,opt,name=cash_out,json=cashOut,proto3" json:"cash_out,omitempty"`
	CashIn               float64  `protobuf:"fixed64,5,opt,name=cash_in,json=cashIn,proto3" json:"cash_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountMonthlySnapshot) Reset()         { *m = AccountMonthlySnapshot{} }
func (m *AccountMonthlySnapshot) String() string { return proto.CompactTextString(m) }
func (*AccountMonthlySnapshot) ProtoMessage()    {}
func (*AccountMonthlySnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d0ea3f8dbe2a8, []int{1}
}

func (m *AccountMonthlySnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountMonthlySnapshot.Unmarshal(m, b)
}
func (m *AccountMonthlySnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountMonthlySnapshot.Marshal(b, m, deterministic)
}
func (m *AccountMonthlySnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountMonthlySnapshot.Merge(m, src)
}
func (m *AccountMonthlySnapshot) XXX_Size() int {
	return xxx_messageInfo_AccountMonthlySnapshot.Size(m)
}
func (m *AccountMonthlySnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountMonthlySnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_AccountMonthlySnapshot proto.InternalMessageInfo

func (m *AccountMonthlySnapshot) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *AccountMonthlySnapshot) GetStartMonthBalance() float64 {
	if m != nil {
		return m.StartMonthBalance
	}
	return 0
}

func (m *AccountMonthlySnapshot) GetEndMonthBalance() float64 {
	if m != nil {
		return m.EndMonthBalance
	}
	return 0
}

func (m *AccountMonthlySnapshot) GetCashOut() float64 {
	if m != nil {
		return m.CashOut
	}
	return 0
}

func (m *AccountMonthlySnapshot) GetCashIn() float64 {
	if m != nil {
		return m.CashIn
	}
	return 0
}

type GetAccountDailySnapshotsRequest struct {
	AccountId            int64    `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountDailySnapshotsRequest) Reset()         { *m = GetAccountDailySnapshotsRequest{} }
func (m *GetAccountDailySnapshotsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountDailySnapshotsRequest) ProtoMessage()    {}
func (*GetAccountDailySnapshotsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d0ea3f8dbe2a8, []int{2}
}

func (m *GetAccountDailySnapshotsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountDailySnapshotsRequest.Unmarshal(m, b)
}
func (m *GetAccountDailySnapshotsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountDailySnapshotsRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountDailySnapshotsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountDailySnapshotsRequest.Merge(m, src)
}
func (m *GetAccountDailySnapshotsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountDailySnapshotsRequest.Size(m)
}
func (m *GetAccountDailySnapshotsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountDailySnapshotsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountDailySnapshotsRequest proto.InternalMessageInfo

func (m *GetAccountDailySnapshotsRequest) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

type GetAccountDailySnapshotsResponse struct {
	AccountDailySnapshots []*AccountDailySnapshot `protobuf:"bytes,1,rep,name=account_daily_snapshots,json=accountDailySnapshots,proto3" json:"account_daily_snapshots,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                `json:"-"`
	XXX_unrecognized      []byte                  `json:"-"`
	XXX_sizecache         int32                   `json:"-"`
}

func (m *GetAccountDailySnapshotsResponse) Reset()         { *m = GetAccountDailySnapshotsResponse{} }
func (m *GetAccountDailySnapshotsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccountDailySnapshotsResponse) ProtoMessage()    {}
func (*GetAccountDailySnapshotsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d0ea3f8dbe2a8, []int{3}
}

func (m *GetAccountDailySnapshotsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountDailySnapshotsResponse.Unmarshal(m, b)
}
func (m *GetAccountDailySnapshotsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountDailySnapshotsResponse.Marshal(b, m, deterministic)
}
func (m *GetAccountDailySnapshotsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountDailySnapshotsResponse.Merge(m, src)
}
func (m *GetAccountDailySnapshotsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAccountDailySnapshotsResponse.Size(m)
}
func (m *GetAccountDailySnapshotsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountDailySnapshotsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountDailySnapshotsResponse proto.InternalMessageInfo

func (m *GetAccountDailySnapshotsResponse) GetAccountDailySnapshots() []*AccountDailySnapshot {
	if m != nil {
		return m.AccountDailySnapshots
	}
	return nil
}

type GetAccountMonthlySnapshotsRequest struct {
	AccountId            int64    `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountMonthlySnapshotsRequest) Reset()         { *m = GetAccountMonthlySnapshotsRequest{} }
func (m *GetAccountMonthlySnapshotsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountMonthlySnapshotsRequest) ProtoMessage()    {}
func (*GetAccountMonthlySnapshotsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d0ea3f8dbe2a8, []int{4}
}

func (m *GetAccountMonthlySnapshotsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountMonthlySnapshotsRequest.Unmarshal(m, b)
}
func (m *GetAccountMonthlySnapshotsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountMonthlySnapshotsRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountMonthlySnapshotsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountMonthlySnapshotsRequest.Merge(m, src)
}
func (m *GetAccountMonthlySnapshotsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountMonthlySnapshotsRequest.Size(m)
}
func (m *GetAccountMonthlySnapshotsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountMonthlySnapshotsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountMonthlySnapshotsRequest proto.InternalMessageInfo

func (m *GetAccountMonthlySnapshotsRequest) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

type GetAccountMonthlySnapshotsResponse struct {
	AccountMonthlySnapshots []*AccountMonthlySnapshot `protobuf:"bytes,1,rep,name=account_monthly_snapshots,json=accountMonthlySnapshots,proto3" json:"account_monthly_snapshots,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                  `json:"-"`
	XXX_unrecognized        []byte                    `json:"-"`
	XXX_sizecache           int32                     `json:"-"`
}

func (m *GetAccountMonthlySnapshotsResponse) Reset()         { *m = GetAccountMonthlySnapshotsResponse{} }
func (m *GetAccountMonthlySnapshotsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccountMonthlySnapshotsResponse) ProtoMessage()    {}
func (*GetAccountMonthlySnapshotsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d0ea3f8dbe2a8, []int{5}
}

func (m *GetAccountMonthlySnapshotsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountMonthlySnapshotsResponse.Unmarshal(m, b)
}
func (m *GetAccountMonthlySnapshotsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountMonthlySnapshotsResponse.Marshal(b, m, deterministic)
}
func (m *GetAccountMonthlySnapshotsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountMonthlySnapshotsResponse.Merge(m, src)
}
func (m *GetAccountMonthlySnapshotsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAccountMonthlySnapshotsResponse.Size(m)
}
func (m *GetAccountMonthlySnapshotsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountMonthlySnapshotsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountMonthlySnapshotsResponse proto.InternalMessageInfo

func (m *GetAccountMonthlySnapshotsResponse) GetAccountMonthlySnapshots() []*AccountMonthlySnapshot {
	if m != nil {
		return m.AccountMonthlySnapshots
	}
	return nil
}

type GetCategoryMonthlySnapshotsRequest struct {
	CategoryId           int64    `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCategoryMonthlySnapshotsRequest) Reset()         { *m = GetCategoryMonthlySnapshotsRequest{} }
func (m *GetCategoryMonthlySnapshotsRequest) String() string { return proto.CompactTextString(m) }
func (*GetCategoryMonthlySnapshotsRequest) ProtoMessage()    {}
func (*GetCategoryMonthlySnapshotsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d0ea3f8dbe2a8, []int{6}
}

func (m *GetCategoryMonthlySnapshotsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCategoryMonthlySnapshotsRequest.Unmarshal(m, b)
}
func (m *GetCategoryMonthlySnapshotsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCategoryMonthlySnapshotsRequest.Marshal(b, m, deterministic)
}
func (m *GetCategoryMonthlySnapshotsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCategoryMonthlySnapshotsRequest.Merge(m, src)
}
func (m *GetCategoryMonthlySnapshotsRequest) XXX_Size() int {
	return xxx_messageInfo_GetCategoryMonthlySnapshotsRequest.Size(m)
}
func (m *GetCategoryMonthlySnapshotsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCategoryMonthlySnapshotsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCategoryMonthlySnapshotsRequest proto.InternalMessageInfo

func (m *GetCategoryMonthlySnapshotsRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type GetCategoryMonthlySnapshotsResponse struct {
	AccountMonthlySnapshots []*AccountMonthlySnapshot `protobuf:"bytes,1,rep,name=account_monthly_snapshots,json=accountMonthlySnapshots,proto3" json:"account_monthly_snapshots,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                  `json:"-"`
	XXX_unrecognized        []byte                    `json:"-"`
	XXX_sizecache           int32                     `json:"-"`
}

func (m *GetCategoryMonthlySnapshotsResponse) Reset()         { *m = GetCategoryMonthlySnapshotsResponse{} }
func (m *GetCategoryMonthlySnapshotsResponse) String() string { return proto.CompactTextString(m) }
func (*GetCategoryMonthlySnapshotsResponse) ProtoMessage()    {}
func (*GetCategoryMonthlySnapshotsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d0ea3f8dbe2a8, []int{7}
}

func (m *GetCategoryMonthlySnapshotsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCategoryMonthlySnapshotsResponse.Unmarshal(m, b)
}
func (m *GetCategoryMonthlySnapshotsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCategoryMonthlySnapshotsResponse.Marshal(b, m, deterministic)
}
func (m *GetCategoryMonthlySnapshotsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCategoryMonthlySnapshotsResponse.Merge(m, src)
}
func (m *GetCategoryMonthlySnapshotsResponse) XXX_Size() int {
	return xxx_messageInfo_GetCategoryMonthlySnapshotsResponse.Size(m)
}
func (m *GetCategoryMonthlySnapshotsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCategoryMonthlySnapshotsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCategoryMonthlySnapshotsResponse proto.InternalMessageInfo

func (m *GetCategoryMonthlySnapshotsResponse) GetAccountMonthlySnapshots() []*AccountMonthlySnapshot {
	if m != nil {
		return m.AccountMonthlySnapshots
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d0ea3f8dbe2a8, []int{8}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AccountDailySnapshot)(nil), "dataprocessing.AccountDailySnapshot")
	proto.RegisterType((*AccountMonthlySnapshot)(nil), "dataprocessing.AccountMonthlySnapshot")
	proto.RegisterType((*GetAccountDailySnapshotsRequest)(nil), "dataprocessing.GetAccountDailySnapshotsRequest")
	proto.RegisterType((*GetAccountDailySnapshotsResponse)(nil), "dataprocessing.GetAccountDailySnapshotsResponse")
	proto.RegisterType((*GetAccountMonthlySnapshotsRequest)(nil), "dataprocessing.GetAccountMonthlySnapshotsRequest")
	proto.RegisterType((*GetAccountMonthlySnapshotsResponse)(nil), "dataprocessing.GetAccountMonthlySnapshotsResponse")
	proto.RegisterType((*GetCategoryMonthlySnapshotsRequest)(nil), "dataprocessing.GetCategoryMonthlySnapshotsRequest")
	proto.RegisterType((*GetCategoryMonthlySnapshotsResponse)(nil), "dataprocessing.GetCategoryMonthlySnapshotsResponse")
	proto.RegisterType((*Empty)(nil), "dataprocessing.Empty")
}

func init() {
	proto.RegisterFile("services/dataprocessing/dataProcessing.proto", fileDescriptor_0a5d0ea3f8dbe2a8)
}

var fileDescriptor_0a5d0ea3f8dbe2a8 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0x74, 0x1f, 0xec, 0x4e, 0x30, 0xcd, 0x30, 0x96, 0x15, 0xa1, 0x95, 0x80, 0xa6, 0x6a,
	0x62, 0x09, 0x74, 0x7f, 0x00, 0x4a, 0x27, 0x28, 0x12, 0x02, 0x65, 0x6f, 0x08, 0x29, 0x72, 0x6c,
	0xab, 0x89, 0x94, 0xda, 0x21, 0x76, 0x40, 0x79, 0x82, 0x17, 0xa4, 0xf1, 0x5b, 0x78, 0xe3, 0x99,
	0x1f, 0x87, 0xe2, 0x34, 0x1b, 0x09, 0x4d, 0xbb, 0xf1, 0xc0, 0x5b, 0xae, 0x7d, 0x7c, 0xee, 0x39,
	0xf7, 0x5e, 0xdd, 0xc0, 0x63, 0xc5, 0xd3, 0x4f, 0x11, 0xe5, 0xca, 0x65, 0x44, 0x93, 0x24, 0x95,
	0x94, 0x2b, 0x15, 0x89, 0x89, 0x09, 0xdf, 0x9d, 0x87, 0x4e, 0x92, 0x4a, 0x2d, 0xf1, 0xad, 0x3a,
	0xc8, 0xfe, 0x89, 0xe0, 0xce, 0x73, 0x4a, 0x65, 0x26, 0xf4, 0x88, 0x44, 0x71, 0x7e, 0x2a, 0x48,
	0xa2, 0x42, 0xa9, 0x31, 0x86, 0x15, 0x46, 0x34, 0xb7, 0x50, 0x0f, 0xf5, 0x37, 0x3c, 0xf3, 0x8d,
	0x0f, 0x61, 0x5b, 0x69, 0x92, 0x6a, 0x9f, 0x91, 0xdc, 0x0f, 0x48, 0x4c, 0x04, 0xe5, 0xd6, 0xf5,
	0x1e, 0xea, 0x23, 0x6f, 0xcb, 0x5c, 0x8c, 0x48, 0x3e, 0x2c, 0x8f, 0xf1, 0x01, 0x6c, 0x71, 0xc1,
	0x6a, 0xc8, 0x8e, 0x41, 0xde, 0xe4, 0x82, 0xfd, 0x81, 0xdb, 0x83, 0x1b, 0x94, 0xa8, 0xd0, 0x97,
	0x99, 0xb6, 0x56, 0x0c, 0x60, 0xbd, 0x88, 0xdf, 0x66, 0x1a, 0xef, 0x82, 0xf9, 0xf4, 0x23, 0x61,
	0xad, 0x9a, 0x9b, 0xb5, 0x22, 0x1c, 0x0b, 0xfb, 0x17, 0x82, 0xbb, 0x33, 0xd1, 0x6f, 0xa4, 0xd0,
	0xe1, 0x12, 0xd9, 0x0e, 0xdc, 0x2e, 0x65, 0x4f, 0x0b, 0x70, 0x43, 0x78, 0xe9, 0xc8, 0xd0, 0x54,
	0x92, 0x0e, 0x61, 0xbb, 0x90, 0x5e, 0x47, 0x97, 0xe2, 0x0b, 0x4f, 0x35, 0xec, 0xbf, 0xc8, 0x7f,
	0x06, 0xfb, 0x2f, 0xb9, 0x9e, 0x57, 0x75, 0xe5, 0xf1, 0x8f, 0x19, 0x57, 0x1a, 0xdf, 0x07, 0x20,
	0xe5, 0xbd, 0x1f, 0x31, 0x63, 0xa6, 0xe3, 0x6d, 0xcc, 0x4e, 0xc6, 0xcc, 0xfe, 0x8a, 0xa0, 0xd7,
	0x4e, 0xa1, 0x12, 0x29, 0x14, 0xc7, 0x1f, 0x60, 0xb7, 0xe2, 0x60, 0x05, 0xc2, 0x57, 0x15, 0xc4,
	0x42, 0xbd, 0x4e, 0x7f, 0x73, 0xf0, 0xc8, 0xa9, 0x0f, 0x83, 0x33, 0x8f, 0xcf, 0xdb, 0x21, 0xf3,
	0xb2, 0xd8, 0x43, 0x78, 0x70, 0xa1, 0xa0, 0xd1, 0x85, 0xcb, 0xda, 0x38, 0x43, 0x60, 0x2f, 0x22,
	0x99, 0x19, 0x09, 0x60, 0xaf, 0x62, 0x99, 0x96, 0x98, 0xbf, 0xac, 0x1c, 0xb4, 0x58, 0x69, 0x70,
	0x7a, 0x55, 0x45, 0x9a, 0xb9, 0xec, 0x13, 0xa3, 0xe4, 0x05, 0xd1, 0x7c, 0x22, 0xd3, 0xbc, 0xcd,
	0xcf, 0x3e, 0x6c, 0xd2, 0x19, 0xe4, 0xc2, 0x10, 0x54, 0x47, 0x63, 0x66, 0x7f, 0x47, 0xf0, 0x70,
	0x21, 0xcf, 0x7f, 0xb4, 0xb4, 0x0e, 0xab, 0x27, 0xd3, 0x44, 0xe7, 0x83, 0x1f, 0x1d, 0xd8, 0x19,
	0xd5, 0x96, 0xc1, 0x69, 0xb9, 0x32, 0xf0, 0x17, 0xb0, 0xda, 0xc6, 0x08, 0xbb, 0xcd, 0xfc, 0x4b,
	0x66, 0xb6, 0xfb, 0xe4, 0xf2, 0x0f, 0xca, 0x2a, 0xd8, 0xd7, 0xf0, 0x37, 0x04, 0xdd, 0xf6, 0x09,
	0xc0, 0x4f, 0xdb, 0x29, 0x5b, 0x5a, 0xd4, 0x1d, 0x5c, 0xe5, 0xc9, 0xb9, 0x8e, 0x33, 0x04, 0xf7,
	0x16, 0xf4, 0x0d, 0xcf, 0x63, 0x5d, 0x32, 0x2c, 0xdd, 0xe3, 0x2b, 0xbd, 0xa9, 0xa4, 0x0c, 0x5f,
	0xbf, 0x7f, 0x35, 0x89, 0x74, 0x98, 0x05, 0x0e, 0x95, 0x53, 0x57, 0x70, 0x1e, 0xd3, 0x50, 0x66,
	0x2c, 0x24, 0x69, 0xee, 0x06, 0x19, 0x9b, 0x70, 0xfd, 0x99, 0xc4, 0x31, 0xd7, 0x47, 0x24, 0x89,
	0x8e, 0x8a, 0xdd, 0xcf, 0x53, 0xb7, 0xe5, 0x17, 0x10, 0xac, 0x99, 0xa5, 0x7f, 0xfc, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x2b, 0x6a, 0xaf, 0x62, 0x24, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataProcessingServiceClient is the client API for DataProcessingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataProcessingServiceClient interface {
	GetAccountDailySnapshots(ctx context.Context, in *GetAccountDailySnapshotsRequest, opts ...grpc.CallOption) (*GetAccountDailySnapshotsResponse, error)
	GetAccountMonthlySnapshots(ctx context.Context, in *GetAccountMonthlySnapshotsRequest, opts ...grpc.CallOption) (*GetAccountMonthlySnapshotsResponse, error)
	GetCategoryMonthlySnapshots(ctx context.Context, in *GetCategoryMonthlySnapshotsRequest, opts ...grpc.CallOption) (*GetCategoryMonthlySnapshotsResponse, error)
}

type dataProcessingServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataProcessingServiceClient(cc *grpc.ClientConn) DataProcessingServiceClient {
	return &dataProcessingServiceClient{cc}
}

func (c *dataProcessingServiceClient) GetAccountDailySnapshots(ctx context.Context, in *GetAccountDailySnapshotsRequest, opts ...grpc.CallOption) (*GetAccountDailySnapshotsResponse, error) {
	out := new(GetAccountDailySnapshotsResponse)
	err := c.cc.Invoke(ctx, "/dataprocessing.DataProcessingService/GetAccountDailySnapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataProcessingServiceClient) GetAccountMonthlySnapshots(ctx context.Context, in *GetAccountMonthlySnapshotsRequest, opts ...grpc.CallOption) (*GetAccountMonthlySnapshotsResponse, error) {
	out := new(GetAccountMonthlySnapshotsResponse)
	err := c.cc.Invoke(ctx, "/dataprocessing.DataProcessingService/GetAccountMonthlySnapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataProcessingServiceClient) GetCategoryMonthlySnapshots(ctx context.Context, in *GetCategoryMonthlySnapshotsRequest, opts ...grpc.CallOption) (*GetCategoryMonthlySnapshotsResponse, error) {
	out := new(GetCategoryMonthlySnapshotsResponse)
	err := c.cc.Invoke(ctx, "/dataprocessing.DataProcessingService/GetCategoryMonthlySnapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataProcessingServiceServer is the server API for DataProcessingService service.
type DataProcessingServiceServer interface {
	GetAccountDailySnapshots(context.Context, *GetAccountDailySnapshotsRequest) (*GetAccountDailySnapshotsResponse, error)
	GetAccountMonthlySnapshots(context.Context, *GetAccountMonthlySnapshotsRequest) (*GetAccountMonthlySnapshotsResponse, error)
	GetCategoryMonthlySnapshots(context.Context, *GetCategoryMonthlySnapshotsRequest) (*GetCategoryMonthlySnapshotsResponse, error)
}

// UnimplementedDataProcessingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDataProcessingServiceServer struct {
}

func (*UnimplementedDataProcessingServiceServer) GetAccountDailySnapshots(ctx context.Context, req *GetAccountDailySnapshotsRequest) (*GetAccountDailySnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountDailySnapshots not implemented")
}
func (*UnimplementedDataProcessingServiceServer) GetAccountMonthlySnapshots(ctx context.Context, req *GetAccountMonthlySnapshotsRequest) (*GetAccountMonthlySnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountMonthlySnapshots not implemented")
}
func (*UnimplementedDataProcessingServiceServer) GetCategoryMonthlySnapshots(ctx context.Context, req *GetCategoryMonthlySnapshotsRequest) (*GetCategoryMonthlySnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryMonthlySnapshots not implemented")
}

func RegisterDataProcessingServiceServer(s *grpc.Server, srv DataProcessingServiceServer) {
	s.RegisterService(&_DataProcessingService_serviceDesc, srv)
}

func _DataProcessingService_GetAccountDailySnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountDailySnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataProcessingServiceServer).GetAccountDailySnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataprocessing.DataProcessingService/GetAccountDailySnapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataProcessingServiceServer).GetAccountDailySnapshots(ctx, req.(*GetAccountDailySnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataProcessingService_GetAccountMonthlySnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountMonthlySnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataProcessingServiceServer).GetAccountMonthlySnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataprocessing.DataProcessingService/GetAccountMonthlySnapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataProcessingServiceServer).GetAccountMonthlySnapshots(ctx, req.(*GetAccountMonthlySnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataProcessingService_GetCategoryMonthlySnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryMonthlySnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataProcessingServiceServer).GetCategoryMonthlySnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataprocessing.DataProcessingService/GetCategoryMonthlySnapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataProcessingServiceServer).GetCategoryMonthlySnapshots(ctx, req.(*GetCategoryMonthlySnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataProcessingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataprocessing.DataProcessingService",
	HandlerType: (*DataProcessingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountDailySnapshots",
			Handler:    _DataProcessingService_GetAccountDailySnapshots_Handler,
		},
		{
			MethodName: "GetAccountMonthlySnapshots",
			Handler:    _DataProcessingService_GetAccountMonthlySnapshots_Handler,
		},
		{
			MethodName: "GetCategoryMonthlySnapshots",
			Handler:    _DataProcessingService_GetCategoryMonthlySnapshots_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/dataprocessing/dataProcessing.proto",
}
