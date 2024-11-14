// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: report/report.proto

package report

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetSummaryReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Period string `protobuf:"bytes,2,opt,name=period,proto3" json:"period,omitempty"`
	Start  string `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End    string `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *GetSummaryReportRequest) Reset() {
	*x = GetSummaryReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSummaryReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSummaryReportRequest) ProtoMessage() {}

func (x *GetSummaryReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSummaryReportRequest.ProtoReflect.Descriptor instead.
func (*GetSummaryReportRequest) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{0}
}

func (x *GetSummaryReportRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetSummaryReportRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *GetSummaryReportRequest) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *GetSummaryReportRequest) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type GetSummaryReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report *ReportResponse `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *GetSummaryReportResponse) Reset() {
	*x = GetSummaryReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSummaryReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSummaryReportResponse) ProtoMessage() {}

func (x *GetSummaryReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSummaryReportResponse.ProtoReflect.Descriptor instead.
func (*GetSummaryReportResponse) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{1}
}

func (x *GetSummaryReportResponse) GetReport() *ReportResponse {
	if x != nil {
		return x.Report
	}
	return nil
}

type GetBudgetReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	BudgetId string `protobuf:"bytes,2,opt,name=budgetId,proto3" json:"budgetId,omitempty"`
}

func (x *GetBudgetReportRequest) Reset() {
	*x = GetBudgetReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBudgetReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBudgetReportRequest) ProtoMessage() {}

func (x *GetBudgetReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBudgetReportRequest.ProtoReflect.Descriptor instead.
func (*GetBudgetReportRequest) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{2}
}

func (x *GetBudgetReportRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetBudgetReportRequest) GetBudgetId() string {
	if x != nil {
		return x.BudgetId
	}
	return ""
}

type GetBudgetReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report *BudgetReport `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *GetBudgetReportResponse) Reset() {
	*x = GetBudgetReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBudgetReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBudgetReportResponse) ProtoMessage() {}

func (x *GetBudgetReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBudgetReportResponse.ProtoReflect.Descriptor instead.
func (*GetBudgetReportResponse) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{3}
}

func (x *GetBudgetReportResponse) GetReport() *BudgetReport {
	if x != nil {
		return x.Report
	}
	return nil
}

type ReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Period           string            `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
	TotalSpent       float32           `protobuf:"fixed32,2,opt,name=totalSpent,proto3" json:"totalSpent,omitempty"`
	TransactionCount int32             `protobuf:"varint,3,opt,name=transactionCount,proto3" json:"transactionCount,omitempty"`
	Categories       []*CategoryReport `protobuf:"bytes,4,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *ReportResponse) Reset() {
	*x = ReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResponse) ProtoMessage() {}

func (x *ReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResponse.ProtoReflect.Descriptor instead.
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{4}
}

func (x *ReportResponse) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *ReportResponse) GetTotalSpent() float32 {
	if x != nil {
		return x.TotalSpent
	}
	return 0
}

func (x *ReportResponse) GetTransactionCount() int32 {
	if x != nil {
		return x.TransactionCount
	}
	return 0
}

func (x *ReportResponse) GetCategories() []*CategoryReport {
	if x != nil {
		return x.Categories
	}
	return nil
}

type CategoryReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Total float32 `protobuf:"fixed32,2,opt,name=total,proto3" json:"total,omitempty"`
	Count int32   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CategoryReport) Reset() {
	*x = CategoryReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryReport) ProtoMessage() {}

func (x *CategoryReport) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryReport.ProtoReflect.Descriptor instead.
func (*CategoryReport) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{5}
}

func (x *CategoryReport) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryReport) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CategoryReport) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type RequiredCategoryReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Total float32 `protobuf:"fixed32,2,opt,name=total,proto3" json:"total,omitempty"`
	Limit float32 `protobuf:"fixed32,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Count int32   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *RequiredCategoryReport) Reset() {
	*x = RequiredCategoryReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequiredCategoryReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredCategoryReport) ProtoMessage() {}

func (x *RequiredCategoryReport) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequiredCategoryReport.ProtoReflect.Descriptor instead.
func (*RequiredCategoryReport) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{6}
}

func (x *RequiredCategoryReport) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequiredCategoryReport) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RequiredCategoryReport) GetLimit() float32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *RequiredCategoryReport) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type BudgetReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BudgetName       string                    `protobuf:"bytes,1,opt,name=budgetName,proto3" json:"budgetName,omitempty"`
	Period           string                    `protobuf:"bytes,2,opt,name=period,proto3" json:"period,omitempty"`
	LeftDays         float32                   `protobuf:"fixed32,3,opt,name=leftDays,proto3" json:"leftDays,omitempty"`
	Limit            float32                   `protobuf:"fixed32,4,opt,name=limit,proto3" json:"limit,omitempty"`
	TotalSpent       float32                   `protobuf:"fixed32,5,opt,name=totalSpent,proto3" json:"totalSpent,omitempty"`
	TransactionCount int32                     `protobuf:"varint,6,opt,name=transactionCount,proto3" json:"transactionCount,omitempty"`
	ReqCategories    []*RequiredCategoryReport `protobuf:"bytes,7,rep,name=reqCategories,proto3" json:"reqCategories,omitempty"`
	Categories       []*CategoryReport         `protobuf:"bytes,8,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *BudgetReport) Reset() {
	*x = BudgetReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetReport) ProtoMessage() {}

func (x *BudgetReport) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetReport.ProtoReflect.Descriptor instead.
func (*BudgetReport) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{7}
}

func (x *BudgetReport) GetBudgetName() string {
	if x != nil {
		return x.BudgetName
	}
	return ""
}

func (x *BudgetReport) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *BudgetReport) GetLeftDays() float32 {
	if x != nil {
		return x.LeftDays
	}
	return 0
}

func (x *BudgetReport) GetLimit() float32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *BudgetReport) GetTotalSpent() float32 {
	if x != nil {
		return x.TotalSpent
	}
	return 0
}

func (x *BudgetReport) GetTransactionCount() int32 {
	if x != nil {
		return x.TransactionCount
	}
	return 0
}

func (x *BudgetReport) GetReqCategories() []*RequiredCategoryReport {
	if x != nil {
		return x.ReqCategories
	}
	return nil
}

func (x *BudgetReport) GetCategories() []*CategoryReport {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_report_report_proto protoreflect.FileDescriptor

var file_report_report_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x22, 0x4a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x4c, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x22, 0xac, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x50, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x6e, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xc2, 0x02, 0x0a, 0x0c, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x65, 0x66, 0x74, 0x44, 0x61, 0x79, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x6c, 0x65, 0x66, 0x74, 0x44, 0x61, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x12,
	0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x0d, 0x72,
	0x65, 0x71, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x32, 0xba, 0x01, 0x0a, 0x0d, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x1f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x52, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x8d, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x75, 0x73, 0x74, 0x49, 0x47, 0x72, 0x65, 0x4b, 0x2f, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0xa2, 0x02, 0x03, 0x52, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0xca,
	0x02, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0xe2, 0x02, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_report_report_proto_rawDescOnce sync.Once
	file_report_report_proto_rawDescData = file_report_report_proto_rawDesc
)

func file_report_report_proto_rawDescGZIP() []byte {
	file_report_report_proto_rawDescOnce.Do(func() {
		file_report_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_report_report_proto_rawDescData)
	})
	return file_report_report_proto_rawDescData
}

var file_report_report_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_report_report_proto_goTypes = []interface{}{
	(*GetSummaryReportRequest)(nil),  // 0: report.GetSummaryReportRequest
	(*GetSummaryReportResponse)(nil), // 1: report.GetSummaryReportResponse
	(*GetBudgetReportRequest)(nil),   // 2: report.GetBudgetReportRequest
	(*GetBudgetReportResponse)(nil),  // 3: report.GetBudgetReportResponse
	(*ReportResponse)(nil),           // 4: report.ReportResponse
	(*CategoryReport)(nil),           // 5: report.CategoryReport
	(*RequiredCategoryReport)(nil),   // 6: report.RequiredCategoryReport
	(*BudgetReport)(nil),             // 7: report.BudgetReport
}
var file_report_report_proto_depIdxs = []int32{
	4, // 0: report.GetSummaryReportResponse.report:type_name -> report.ReportResponse
	7, // 1: report.GetBudgetReportResponse.report:type_name -> report.BudgetReport
	5, // 2: report.ReportResponse.categories:type_name -> report.CategoryReport
	6, // 3: report.BudgetReport.reqCategories:type_name -> report.RequiredCategoryReport
	5, // 4: report.BudgetReport.categories:type_name -> report.CategoryReport
	0, // 5: report.ReportService.GetSummaryReport:input_type -> report.GetSummaryReportRequest
	2, // 6: report.ReportService.GetBudgetReport:input_type -> report.GetBudgetReportRequest
	1, // 7: report.ReportService.GetSummaryReport:output_type -> report.GetSummaryReportResponse
	3, // 8: report.ReportService.GetBudgetReport:output_type -> report.GetBudgetReportResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_report_report_proto_init() }
func file_report_report_proto_init() {
	if File_report_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_report_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSummaryReportRequest); i {
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
		file_report_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSummaryReportResponse); i {
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
		file_report_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBudgetReportRequest); i {
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
		file_report_report_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBudgetReportResponse); i {
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
		file_report_report_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportResponse); i {
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
		file_report_report_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryReport); i {
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
		file_report_report_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequiredCategoryReport); i {
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
		file_report_report_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BudgetReport); i {
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
			RawDescriptor: file_report_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_report_report_proto_goTypes,
		DependencyIndexes: file_report_report_proto_depIdxs,
		MessageInfos:      file_report_report_proto_msgTypes,
	}.Build()
	File_report_report_proto = out.File
	file_report_report_proto_rawDesc = nil
	file_report_report_proto_goTypes = nil
	file_report_report_proto_depIdxs = nil
}
