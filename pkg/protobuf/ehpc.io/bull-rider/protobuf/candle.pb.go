// Code generated by protoc-gen-go. DO NOT EDIT.
// source: candle.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Candle is an exhange candle
type Candle struct {
	Exhange              int32    `protobuf:"varint,1,opt,name=exhange,proto3" json:"exhange,omitempty"`
	Pair                 Pair     `protobuf:"varint,2,opt,name=pair,proto3,enum=protobuf.Pair" json:"pair,omitempty"`
	Interval             int64    `protobuf:"varint,3,opt,name=interval,proto3" json:"interval,omitempty"`
	OpenTime             int64    `protobuf:"varint,4,opt,name=open_time,json=openTime,proto3" json:"open_time,omitempty"`
	CloseTime            int64    `protobuf:"varint,5,opt,name=close_time,json=closeTime,proto3" json:"close_time,omitempty"`
	Open                 float64  `protobuf:"fixed64,6,opt,name=open,proto3" json:"open,omitempty"`
	Close                float64  `protobuf:"fixed64,7,opt,name=close,proto3" json:"close,omitempty"`
	High                 float64  `protobuf:"fixed64,8,opt,name=high,proto3" json:"high,omitempty"`
	Low                  float64  `protobuf:"fixed64,9,opt,name=low,proto3" json:"low,omitempty"`
	Volume               float64  `protobuf:"fixed64,10,opt,name=volume,proto3" json:"volume,omitempty"`
	QuoteVolume          float64  `protobuf:"fixed64,11,opt,name=quote_volume,json=quoteVolume,proto3" json:"quote_volume,omitempty"`
	TrandesNumber        int64    `protobuf:"varint,12,opt,name=trandes_number,json=trandesNumber,proto3" json:"trandes_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Candle) Reset()         { *m = Candle{} }
func (m *Candle) String() string { return proto.CompactTextString(m) }
func (*Candle) ProtoMessage()    {}
func (*Candle) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd2c92c952efae58, []int{0}
}

func (m *Candle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candle.Unmarshal(m, b)
}
func (m *Candle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candle.Marshal(b, m, deterministic)
}
func (m *Candle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candle.Merge(m, src)
}
func (m *Candle) XXX_Size() int {
	return xxx_messageInfo_Candle.Size(m)
}
func (m *Candle) XXX_DiscardUnknown() {
	xxx_messageInfo_Candle.DiscardUnknown(m)
}

var xxx_messageInfo_Candle proto.InternalMessageInfo

func (m *Candle) GetExhange() int32 {
	if m != nil {
		return m.Exhange
	}
	return 0
}

func (m *Candle) GetPair() Pair {
	if m != nil {
		return m.Pair
	}
	return Pair_BTCUSDT
}

func (m *Candle) GetInterval() int64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *Candle) GetOpenTime() int64 {
	if m != nil {
		return m.OpenTime
	}
	return 0
}

func (m *Candle) GetCloseTime() int64 {
	if m != nil {
		return m.CloseTime
	}
	return 0
}

func (m *Candle) GetOpen() float64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *Candle) GetClose() float64 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *Candle) GetHigh() float64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *Candle) GetLow() float64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *Candle) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Candle) GetQuoteVolume() float64 {
	if m != nil {
		return m.QuoteVolume
	}
	return 0
}

func (m *Candle) GetTrandesNumber() int64 {
	if m != nil {
		return m.TrandesNumber
	}
	return 0
}

// Candles is an array of candles
type Candles struct {
	Candles              []*Candles `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Candles) Reset()         { *m = Candles{} }
func (m *Candles) String() string { return proto.CompactTextString(m) }
func (*Candles) ProtoMessage()    {}
func (*Candles) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd2c92c952efae58, []int{1}
}

func (m *Candles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candles.Unmarshal(m, b)
}
func (m *Candles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candles.Marshal(b, m, deterministic)
}
func (m *Candles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candles.Merge(m, src)
}
func (m *Candles) XXX_Size() int {
	return xxx_messageInfo_Candles.Size(m)
}
func (m *Candles) XXX_DiscardUnknown() {
	xxx_messageInfo_Candles.DiscardUnknown(m)
}

var xxx_messageInfo_Candles proto.InternalMessageInfo

func (m *Candles) GetCandles() []*Candles {
	if m != nil {
		return m.Candles
	}
	return nil
}

func init() {
	proto.RegisterType((*Candle)(nil), "protobuf.Candle")
	proto.RegisterType((*Candles)(nil), "protobuf.Candles")
}

func init() { proto.RegisterFile("candle.proto", fileDescriptor_cd2c92c952efae58) }

var fileDescriptor_cd2c92c952efae58 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0xc9, 0xba, 0xb5, 0xdb, 0xd9, 0x1c, 0x1a, 0x44, 0xc2, 0xc6, 0xa0, 0x0e, 0x84, 0x82,
	0xd8, 0xc1, 0x04, 0x1f, 0x40, 0xef, 0x45, 0x8a, 0x78, 0xe1, 0xcd, 0x68, 0xb7, 0xe3, 0x1a, 0x48,
	0x9b, 0x9a, 0xb6, 0xd3, 0xc7, 0xf4, 0x91, 0x24, 0x27, 0xad, 0x5e, 0xf5, 0xff, 0xbf, 0xff, 0x6b,
	0x29, 0x09, 0xcc, 0xf6, 0x69, 0x79, 0x50, 0x18, 0x57, 0x46, 0x37, 0x9a, 0x8f, 0xe9, 0x91, 0xb5,
	0x1f, 0x0b, 0xa8, 0x52, 0x69, 0x1c, 0x5d, 0xff, 0x0c, 0xc0, 0x7f, 0x22, 0x8d, 0x0b, 0x08, 0xf0,
	0x3b, 0x4f, 0xcb, 0x23, 0x0a, 0x16, 0xb2, 0x68, 0x94, 0xf4, 0x95, 0xaf, 0x61, 0x68, 0x5f, 0x11,
	0x83, 0x90, 0x45, 0xf3, 0xed, 0x3c, 0xee, 0xbf, 0x14, 0xbf, 0xa4, 0xd2, 0x24, 0xb4, 0xf1, 0x05,
	0x8c, 0x65, 0xd9, 0xa0, 0x39, 0xa5, 0x4a, 0x78, 0x21, 0x8b, 0xbc, 0xe4, 0xaf, 0xf3, 0x25, 0x4c,
	0x74, 0x85, 0xe5, 0xae, 0x91, 0x05, 0x8a, 0xa1, 0x1b, 0x2d, 0x78, 0x95, 0x05, 0xf2, 0x15, 0xc0,
	0x5e, 0xe9, 0x1a, 0xdd, 0x3a, 0xa2, 0x75, 0x42, 0x84, 0x66, 0x0e, 0x43, 0xab, 0x0a, 0x3f, 0x64,
	0x11, 0x4b, 0x28, 0xf3, 0x4b, 0x18, 0x91, 0x20, 0x02, 0x82, 0xae, 0x58, 0x33, 0x97, 0xc7, 0x5c,
	0x8c, 0x9d, 0x69, 0x33, 0x3f, 0x07, 0x4f, 0xe9, 0x2f, 0x31, 0x21, 0x64, 0x23, 0xbf, 0x02, 0xff,
	0xa4, 0x55, 0x5b, 0xa0, 0x00, 0x82, 0x5d, 0xe3, 0xd7, 0x30, 0xfb, 0x6c, 0x75, 0x83, 0xbb, 0x6e,
	0x9d, 0xd2, 0x3a, 0x25, 0xf6, 0xe6, 0x94, 0x1b, 0x98, 0x37, 0x26, 0x2d, 0x0f, 0x58, 0xef, 0xca,
	0xb6, 0xc8, 0xd0, 0x88, 0x19, 0xfd, 0xed, 0x59, 0x47, 0x9f, 0x09, 0xae, 0x1f, 0x20, 0x70, 0x27,
	0x5a, 0xf3, 0x5b, 0x08, 0xdc, 0x1d, 0xd4, 0x82, 0x85, 0x5e, 0x34, 0xdd, 0x5e, 0xfc, 0x9f, 0x5d,
	0xe7, 0x24, 0xbd, 0xf1, 0xb8, 0x7a, 0x5f, 0x62, 0x5e, 0xed, 0x63, 0xa9, 0x37, 0x59, 0xab, 0xd4,
	0x9d, 0x91, 0x07, 0x34, 0x9b, 0xde, 0xcf, 0x7c, 0x4a, 0xf7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xb4, 0x27, 0x11, 0xea, 0xd6, 0x01, 0x00, 0x00,
}