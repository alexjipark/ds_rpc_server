// Code generated by protoc-gen-go.
// source: ProtoTest.proto
// DO NOT EDIT!

/*
Package ProtobufTest is a generated protocol buffer package.

It is generated from these files:
	ProtoTest.proto

It has these top-level messages:
	DataSTResponse
	DataSTRequest
	TestMessage
*/
package ProtobufTest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DataSTRequest_RequestType int32

const (
	DataSTRequest_ReqCheckBalance           DataSTRequest_RequestType = 0
	DataSTRequest_ReqTransferCurrency       DataSTRequest_RequestType = 1
	DataSTRequest_ReqTransferDataOwnerhship DataSTRequest_RequestType = 2
	DataSTRequest_ReqGetDataList            DataSTRequest_RequestType = 3
)

var DataSTRequest_RequestType_name = map[int32]string{
	0: "ReqCheckBalance",
	1: "ReqTransferCurrency",
	2: "ReqTransferDataOwnerhship",
	3: "ReqGetDataList",
}
var DataSTRequest_RequestType_value = map[string]int32{
	"ReqCheckBalance":           0,
	"ReqTransferCurrency":       1,
	"ReqTransferDataOwnerhship": 2,
	"ReqGetDataList":            3,
}

func (x DataSTRequest_RequestType) Enum() *DataSTRequest_RequestType {
	p := new(DataSTRequest_RequestType)
	*p = x
	return p
}
func (x DataSTRequest_RequestType) String() string {
	return proto.EnumName(DataSTRequest_RequestType_name, int32(x))
}
func (x *DataSTRequest_RequestType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DataSTRequest_RequestType_value, data, "DataSTRequest_RequestType")
	if err != nil {
		return err
	}
	*x = DataSTRequest_RequestType(value)
	return nil
}
func (DataSTRequest_RequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type TestMessage_ItemType int32

const (
	TestMessage_TypeX TestMessage_ItemType = 0
	TestMessage_TypeY TestMessage_ItemType = 1
	TestMessage_TypeZ TestMessage_ItemType = 2
)

var TestMessage_ItemType_name = map[int32]string{
	0: "TypeX",
	1: "TypeY",
	2: "TypeZ",
}
var TestMessage_ItemType_value = map[string]int32{
	"TypeX": 0,
	"TypeY": 1,
	"TypeZ": 2,
}

func (x TestMessage_ItemType) Enum() *TestMessage_ItemType {
	p := new(TestMessage_ItemType)
	*p = x
	return p
}
func (x TestMessage_ItemType) String() string {
	return proto.EnumName(TestMessage_ItemType_name, int32(x))
}
func (x *TestMessage_ItemType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TestMessage_ItemType_value, data, "TestMessage_ItemType")
	if err != nil {
		return err
	}
	*x = TestMessage_ItemType(value)
	return nil
}
func (TestMessage_ItemType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type DataSTResponse struct {
	Address          *string                         `protobuf:"bytes,1,req,name=address" json:"address,omitempty"`
	Balance          *int32                          `protobuf:"varint,2,req,name=balance" json:"balance,omitempty"`
	OwnedDataList    []*DataSTResponse_DataOwnership `protobuf:"bytes,3,rep,name=ownedDataList" json:"ownedDataList,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *DataSTResponse) Reset()                    { *m = DataSTResponse{} }
func (m *DataSTResponse) String() string            { return proto.CompactTextString(m) }
func (*DataSTResponse) ProtoMessage()               {}
func (*DataSTResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DataSTResponse) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *DataSTResponse) GetBalance() int32 {
	if m != nil && m.Balance != nil {
		return *m.Balance
	}
	return 0
}

func (m *DataSTResponse) GetOwnedDataList() []*DataSTResponse_DataOwnership {
	if m != nil {
		return m.OwnedDataList
	}
	return nil
}

type DataSTResponse_DataOwnership struct {
	DataHash         *string `protobuf:"bytes,1,req,name=dataHash" json:"dataHash,omitempty"`
	OriginAddr       *string `protobuf:"bytes,2,req,name=originAddr" json:"originAddr,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DataSTResponse_DataOwnership) Reset()                    { *m = DataSTResponse_DataOwnership{} }
func (m *DataSTResponse_DataOwnership) String() string            { return proto.CompactTextString(m) }
func (*DataSTResponse_DataOwnership) ProtoMessage()               {}
func (*DataSTResponse_DataOwnership) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *DataSTResponse_DataOwnership) GetDataHash() string {
	if m != nil && m.DataHash != nil {
		return *m.DataHash
	}
	return ""
}

func (m *DataSTResponse_DataOwnership) GetOriginAddr() string {
	if m != nil && m.OriginAddr != nil {
		return *m.OriginAddr
	}
	return ""
}

type DataSTRequest struct {
	ReqType          *DataSTRequest_RequestType `protobuf:"varint,1,req,name=reqType,enum=ProtobufTest.DataSTRequest_RequestType" json:"reqType,omitempty"`
	Address          *string                    `protobuf:"bytes,2,req,name=address" json:"address,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *DataSTRequest) Reset()                    { *m = DataSTRequest{} }
func (m *DataSTRequest) String() string            { return proto.CompactTextString(m) }
func (*DataSTRequest) ProtoMessage()               {}
func (*DataSTRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DataSTRequest) GetReqType() DataSTRequest_RequestType {
	if m != nil && m.ReqType != nil {
		return *m.ReqType
	}
	return DataSTRequest_ReqCheckBalance
}

func (m *DataSTRequest) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

type TestMessage struct {
	ClientName       *string                `protobuf:"bytes,1,req,name=clientName" json:"clientName,omitempty"`
	ClientId         *int32                 `protobuf:"varint,2,req,name=clientId" json:"clientId,omitempty"`
	Description      *string                `protobuf:"bytes,3,opt,name=description,def=NONE" json:"description,omitempty"`
	Messageitems     []*TestMessage_MsgItem `protobuf:"bytes,4,rep,name=messageitems" json:"messageitems,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *TestMessage) Reset()                    { *m = TestMessage{} }
func (m *TestMessage) String() string            { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()               {}
func (*TestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

const Default_TestMessage_Description string = "NONE"

func (m *TestMessage) GetClientName() string {
	if m != nil && m.ClientName != nil {
		return *m.ClientName
	}
	return ""
}

func (m *TestMessage) GetClientId() int32 {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return 0
}

func (m *TestMessage) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return Default_TestMessage_Description
}

func (m *TestMessage) GetMessageitems() []*TestMessage_MsgItem {
	if m != nil {
		return m.Messageitems
	}
	return nil
}

type TestMessage_MsgItem struct {
	Id               *int32                `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	ItemName         *string               `protobuf:"bytes,2,opt,name=itemName" json:"itemName,omitempty"`
	ItemValue        *int32                `protobuf:"varint,3,opt,name=itemValue" json:"itemValue,omitempty"`
	ItemType         *TestMessage_ItemType `protobuf:"varint,4,opt,name=itemType,enum=ProtobufTest.TestMessage_ItemType" json:"itemType,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *TestMessage_MsgItem) Reset()                    { *m = TestMessage_MsgItem{} }
func (m *TestMessage_MsgItem) String() string            { return proto.CompactTextString(m) }
func (*TestMessage_MsgItem) ProtoMessage()               {}
func (*TestMessage_MsgItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *TestMessage_MsgItem) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TestMessage_MsgItem) GetItemName() string {
	if m != nil && m.ItemName != nil {
		return *m.ItemName
	}
	return ""
}

func (m *TestMessage_MsgItem) GetItemValue() int32 {
	if m != nil && m.ItemValue != nil {
		return *m.ItemValue
	}
	return 0
}

func (m *TestMessage_MsgItem) GetItemType() TestMessage_ItemType {
	if m != nil && m.ItemType != nil {
		return *m.ItemType
	}
	return TestMessage_TypeX
}

func init() {
	proto.RegisterType((*DataSTResponse)(nil), "ProtobufTest.DataSTResponse")
	proto.RegisterType((*DataSTResponse_DataOwnership)(nil), "ProtobufTest.DataSTResponse.DataOwnership")
	proto.RegisterType((*DataSTRequest)(nil), "ProtobufTest.DataSTRequest")
	proto.RegisterType((*TestMessage)(nil), "ProtobufTest.TestMessage")
	proto.RegisterType((*TestMessage_MsgItem)(nil), "ProtobufTest.TestMessage.MsgItem")
	proto.RegisterEnum("ProtobufTest.DataSTRequest_RequestType", DataSTRequest_RequestType_name, DataSTRequest_RequestType_value)
	proto.RegisterEnum("ProtobufTest.TestMessage_ItemType", TestMessage_ItemType_name, TestMessage_ItemType_value)
}

func init() { proto.RegisterFile("ProtoTest.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x51, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xed, 0x3a, 0x89, 0xd2, 0x4c, 0xda, 0x74, 0xe5, 0x1e, 0x58, 0x22, 0x40, 0x61, 0x0f, 0x10,
	0x81, 0x94, 0x43, 0x8e, 0x1c, 0x90, 0x4a, 0xa9, 0x20, 0x82, 0xa6, 0x95, 0x89, 0x10, 0x70, 0x73,
	0xd7, 0xd3, 0xc4, 0x22, 0xf1, 0xee, 0xda, 0x8e, 0xaa, 0xfe, 0x0e, 0xfe, 0x1e, 0x37, 0x7e, 0x06,
	0x17, 0x64, 0xef, 0x07, 0xbb, 0x12, 0x9c, 0x3c, 0xef, 0xcd, 0xd7, 0xf3, 0x1b, 0x38, 0xb9, 0xd6,
	0xa9, 0x4d, 0x57, 0x68, 0xec, 0x2c, 0x73, 0x11, 0x3d, 0xf2, 0xc4, 0xcd, 0xfe, 0xd6, 0x71, 0xf1,
	0xaf, 0x00, 0x46, 0x6f, 0xb9, 0xe5, 0x9f, 0x56, 0x0c, 0x4d, 0x96, 0x2a, 0x83, 0x34, 0x82, 0x3e,
	0x17, 0x42, 0xa3, 0x31, 0x51, 0x30, 0x21, 0xd3, 0x01, 0xab, 0xa0, 0xcb, 0xdc, 0xf0, 0x2d, 0x57,
	0x09, 0x46, 0x64, 0x42, 0xa6, 0x3d, 0x56, 0x41, 0x7a, 0x0d, 0xc7, 0xe9, 0x9d, 0x42, 0xe1, 0x46,
	0x7d, 0x94, 0xc6, 0x46, 0x9d, 0x49, 0x67, 0x3a, 0x9c, 0xbf, 0x98, 0x35, 0x97, 0xcd, 0xda, 0x8b,
	0x3c, 0xbc, 0xba, 0x53, 0xa8, 0xcd, 0x46, 0x66, 0xac, 0x3d, 0x60, 0xfc, 0x01, 0x8e, 0x5b, 0x79,
	0x3a, 0x86, 0x43, 0xc1, 0x2d, 0x7f, 0xcf, 0xcd, 0xa6, 0xd4, 0x55, 0x63, 0xfa, 0x04, 0x20, 0xd5,
	0x72, 0x2d, 0xd5, 0x99, 0x10, 0xda, 0x6b, 0x1b, 0xb0, 0x06, 0x13, 0xff, 0x0c, 0x8a, 0x69, 0x6e,
	0x79, 0xbe, 0x47, 0x63, 0xe9, 0x19, 0xf4, 0x35, 0xe6, 0xab, 0xfb, 0x0c, 0xfd, 0xb0, 0xd1, 0xfc,
	0xf9, 0xbf, 0xa5, 0xfa, 0xea, 0x59, 0xf9, 0xba, 0x72, 0x56, 0xf5, 0x35, 0x7d, 0x22, 0x2d, 0x9f,
	0x62, 0x05, 0xc3, 0x46, 0x07, 0x3d, 0x85, 0x13, 0x86, 0xf9, 0xf9, 0x06, 0x93, 0xef, 0x6f, 0x0a,
	0xbf, 0xc2, 0x03, 0xfa, 0x00, 0x4e, 0x19, 0xe6, 0x2b, 0xcd, 0x95, 0xb9, 0x45, 0x7d, 0xbe, 0xd7,
	0x1a, 0x55, 0x72, 0x1f, 0x06, 0xf4, 0x31, 0x3c, 0x6c, 0x24, 0x6a, 0x0f, 0x36, 0xce, 0x84, 0x90,
	0x50, 0x0a, 0x23, 0x86, 0xf9, 0x3b, 0xb4, 0x95, 0x53, 0x61, 0x27, 0xfe, 0x4d, 0x60, 0xe8, 0x54,
	0x5f, 0xa2, 0x31, 0x7c, 0x8d, 0xce, 0x8e, 0x64, 0x2b, 0x51, 0xd9, 0x25, 0xdf, 0x61, 0x69, 0x56,
	0x83, 0x71, 0x56, 0x16, 0x68, 0x21, 0xca, 0x43, 0xd6, 0x98, 0x3e, 0x83, 0xa1, 0x40, 0x93, 0x68,
	0x99, 0x59, 0x99, 0xaa, 0xa8, 0x33, 0x09, 0xa6, 0x83, 0x57, 0xdd, 0xe5, 0xd5, 0xf2, 0x82, 0x35,
	0x13, 0xf4, 0x02, 0x8e, 0x76, 0xc5, 0x3a, 0x69, 0x71, 0x67, 0xa2, 0xae, 0x3f, 0xf8, 0xd3, 0xb6,
	0x8b, 0x0d, 0x51, 0xb3, 0x4b, 0xb3, 0x5e, 0x58, 0xdc, 0xb1, 0x56, 0xdb, 0xf8, 0x47, 0x00, 0xfd,
	0x32, 0x43, 0x47, 0x40, 0xa4, 0xf0, 0x72, 0x7b, 0x8c, 0x48, 0xe1, 0x64, 0xba, 0x22, 0xff, 0x09,
	0xe2, 0x74, 0xb0, 0x1a, 0xd3, 0x47, 0x30, 0x70, 0xf1, 0x67, 0xbe, 0xdd, 0xa3, 0x17, 0xd9, 0x63,
	0x7f, 0x09, 0xfa, 0xba, 0xe8, 0xf4, 0xe7, 0xed, 0x4e, 0x82, 0xe9, 0x68, 0x1e, 0xff, 0x5f, 0xd8,
	0xa2, 0xac, 0x64, 0x75, 0x4f, 0xfc, 0x12, 0x0e, 0x2b, 0x96, 0x0e, 0xa0, 0xe7, 0xde, 0x2f, 0xe1,
	0x41, 0x15, 0x7e, 0x0d, 0x83, 0x2a, 0xfc, 0x16, 0x92, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0xb8, 0x5d, 0x8b, 0x62, 0x03, 0x00, 0x00,
}
