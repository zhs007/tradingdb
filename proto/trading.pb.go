// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trading.proto

package tradingdb

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

type Candle struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Open                 int64    `protobuf:"varint,3,opt,name=open,proto3" json:"open,omitempty"`
	Close                int64    `protobuf:"varint,4,opt,name=close,proto3" json:"close,omitempty"`
	High                 int64    `protobuf:"varint,5,opt,name=high,proto3" json:"high,omitempty"`
	Low                  int64    `protobuf:"varint,6,opt,name=low,proto3" json:"low,omitempty"`
	CurTime              int64    `protobuf:"varint,7,opt,name=curTime,proto3" json:"curTime,omitempty"`
	Volume               int64    `protobuf:"varint,8,opt,name=volume,proto3" json:"volume,omitempty"`
	OpenInterest         int64    `protobuf:"varint,9,opt,name=openInterest,proto3" json:"openInterest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Candle) Reset()         { *m = Candle{} }
func (m *Candle) String() string { return proto.CompactTextString(m) }
func (*Candle) ProtoMessage()    {}
func (*Candle) Descriptor() ([]byte, []int) {
	return fileDescriptor_trading_5ff53483390b0ad2, []int{0}
}
func (m *Candle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candle.Unmarshal(m, b)
}
func (m *Candle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candle.Marshal(b, m, deterministic)
}
func (dst *Candle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candle.Merge(dst, src)
}
func (m *Candle) XXX_Size() int {
	return xxx_messageInfo_Candle.Size(m)
}
func (m *Candle) XXX_DiscardUnknown() {
	xxx_messageInfo_Candle.DiscardUnknown(m)
}

var xxx_messageInfo_Candle proto.InternalMessageInfo

func (m *Candle) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Candle) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Candle) GetOpen() int64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *Candle) GetClose() int64 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *Candle) GetHigh() int64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *Candle) GetLow() int64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *Candle) GetCurTime() int64 {
	if m != nil {
		return m.CurTime
	}
	return 0
}

func (m *Candle) GetVolume() int64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Candle) GetOpenInterest() int64 {
	if m != nil {
		return m.OpenInterest
	}
	return 0
}

type CandleChunk struct {
	Code                 string    `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StartTime            int64     `protobuf:"varint,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              int64     `protobuf:"varint,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
	KeyID                string    `protobuf:"bytes,5,opt,name=keyID,proto3" json:"keyID,omitempty"`
	Candles              []*Candle `protobuf:"bytes,6,rep,name=candles,proto3" json:"candles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CandleChunk) Reset()         { *m = CandleChunk{} }
func (m *CandleChunk) String() string { return proto.CompactTextString(m) }
func (*CandleChunk) ProtoMessage()    {}
func (*CandleChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_trading_5ff53483390b0ad2, []int{1}
}
func (m *CandleChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CandleChunk.Unmarshal(m, b)
}
func (m *CandleChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CandleChunk.Marshal(b, m, deterministic)
}
func (dst *CandleChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CandleChunk.Merge(dst, src)
}
func (m *CandleChunk) XXX_Size() int {
	return xxx_messageInfo_CandleChunk.Size(m)
}
func (m *CandleChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_CandleChunk.DiscardUnknown(m)
}

var xxx_messageInfo_CandleChunk proto.InternalMessageInfo

func (m *CandleChunk) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CandleChunk) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CandleChunk) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *CandleChunk) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *CandleChunk) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *CandleChunk) GetCandles() []*Candle {
	if m != nil {
		return m.Candles
	}
	return nil
}

func init() {
	proto.RegisterType((*Candle)(nil), "tradingdb.Candle")
	proto.RegisterType((*CandleChunk)(nil), "tradingdb.CandleChunk")
}

func init() { proto.RegisterFile("trading.proto", fileDescriptor_trading_5ff53483390b0ad2) }

var fileDescriptor_trading_5ff53483390b0ad2 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0xdd, 0x4d, 0xcd, 0xac, 0x82, 0x06, 0x91, 0x39, 0x78, 0x28, 0x3d, 0x15, 0x84,
	0x1e, 0xf4, 0x11, 0xd6, 0xcb, 0x5e, 0x8b, 0x2f, 0xd0, 0x6d, 0x86, 0x6d, 0xd9, 0x36, 0x59, 0xd2,
	0x54, 0xf1, 0xad, 0x7c, 0x1a, 0x9f, 0x47, 0x32, 0xd9, 0x2a, 0xde, 0xbc, 0xfd, 0xff, 0x37, 0x03,
	0xf9, 0xff, 0x09, 0x5c, 0x7b, 0xd7, 0xe8, 0xde, 0x1c, 0xaa, 0x93, 0xb3, 0xde, 0x2a, 0x79, 0xb6,
	0x7a, 0x5f, 0x7c, 0x25, 0x20, 0xb6, 0x8d, 0xd1, 0x03, 0x29, 0x05, 0xab, 0xd6, 0x6a, 0xc2, 0x24,
	0x4f, 0x4a, 0x59, 0xb3, 0x0e, 0xcc, 0x34, 0x23, 0xe1, 0x45, 0x64, 0x41, 0x07, 0x66, 0x4f, 0x64,
	0x30, 0xcd, 0x93, 0x32, 0xad, 0x59, 0xab, 0x3b, 0x58, 0xb7, 0x83, 0x9d, 0x08, 0x57, 0x0c, 0xa3,
	0x09, 0x9b, 0x5d, 0x7f, 0xe8, 0x70, 0x1d, 0x37, 0x83, 0x56, 0x37, 0x90, 0x0e, 0xf6, 0x1d, 0x05,
	0xa3, 0x20, 0x15, 0x42, 0xd6, 0xce, 0xee, 0xb5, 0x1f, 0x09, 0x33, 0xa6, 0x8b, 0x55, 0xf7, 0x20,
	0xde, 0xec, 0x30, 0x8f, 0x84, 0x97, 0x3c, 0x38, 0x3b, 0x55, 0xc0, 0x55, 0x78, 0x75, 0x67, 0x3c,
	0x39, 0x9a, 0x3c, 0x4a, 0x9e, 0xfe, 0x61, 0xc5, 0x67, 0x02, 0x9b, 0x58, 0x6c, 0xdb, 0xcd, 0xe6,
	0xf8, 0xef, 0x76, 0x0f, 0x20, 0x27, 0xdf, 0x38, 0xcf, 0x79, 0x62, 0xc5, 0x5f, 0x10, 0xb2, 0x92,
	0xd1, 0x3c, 0x8b, 0x4d, 0x17, 0x1b, 0x2e, 0x70, 0xa4, 0x8f, 0xdd, 0x0b, 0x97, 0x95, 0x75, 0x34,
	0xea, 0x11, 0xb2, 0x96, 0x43, 0x4c, 0x28, 0xf2, 0xb4, 0xdc, 0x3c, 0xdd, 0x56, 0x3f, 0xb7, 0xaf,
	0x62, 0xbc, 0x7a, 0xd9, 0xd8, 0x0b, 0xfe, 0x9d, 0xe7, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4f,
	0xfe, 0x20, 0x01, 0xae, 0x01, 0x00, 0x00,
}
