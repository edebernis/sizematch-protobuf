// Code generated by protoc-gen-go. DO NOT EDIT.
// source: items/items.proto

package items

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

type Lang int32

const (
	Lang_EN Lang = 0
	Lang_FR Lang = 1
)

var Lang_name = map[int32]string{
	0: "EN",
	1: "FR",
}

var Lang_value = map[string]int32{
	"EN": 0,
	"FR": 1,
}

func (x Lang) String() string {
	return proto.EnumName(Lang_name, int32(x))
}

func (Lang) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2532e29ffc1dd680, []int{0}
}

type Dimension_Name int32

const (
	Dimension_HEIGHT Dimension_Name = 0
	Dimension_WIDTH  Dimension_Name = 1
	Dimension_DEPTH  Dimension_Name = 2
	Dimension_WEIGHT Dimension_Name = 3
)

var Dimension_Name_name = map[int32]string{
	0: "HEIGHT",
	1: "WIDTH",
	2: "DEPTH",
	3: "WEIGHT",
}

var Dimension_Name_value = map[string]int32{
	"HEIGHT": 0,
	"WIDTH":  1,
	"DEPTH":  2,
	"WEIGHT": 3,
}

func (x Dimension_Name) String() string {
	return proto.EnumName(Dimension_Name_name, int32(x))
}

func (Dimension_Name) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2532e29ffc1dd680, []int{2, 0}
}

type Dimension_Unit int32

const (
	Dimension_MM Dimension_Unit = 0
	Dimension_CM Dimension_Unit = 1
	Dimension_M  Dimension_Unit = 2
	Dimension_G  Dimension_Unit = 3
	Dimension_KG Dimension_Unit = 4
)

var Dimension_Unit_name = map[int32]string{
	0: "MM",
	1: "CM",
	2: "M",
	3: "G",
	4: "KG",
}

var Dimension_Unit_value = map[string]int32{
	"MM": 0,
	"CM": 1,
	"M":  2,
	"G":  3,
	"KG": 4,
}

func (x Dimension_Unit) String() string {
	return proto.EnumName(Dimension_Unit_name, int32(x))
}

func (Dimension_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2532e29ffc1dd680, []int{2, 1}
}

type Price_Currency int32

const (
	Price_EUR Price_Currency = 0
	Price_GBP Price_Currency = 1
)

var Price_Currency_name = map[int32]string{
	0: "EUR",
	1: "GBP",
}

var Price_Currency_value = map[string]int32{
	"EUR": 0,
	"GBP": 1,
}

func (x Price_Currency) String() string {
	return proto.EnumName(Price_Currency_name, int32(x))
}

func (Price_Currency) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2532e29ffc1dd680, []int{3, 0}
}

type Item struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Source               string            `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Lang                 string            `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
	Urls                 []string          `protobuf:"bytes,4,rep,name=urls,proto3" json:"urls,omitempty"`
	Name                 string            `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description          string            `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Categories           []string          `protobuf:"bytes,7,rep,name=categories,proto3" json:"categories,omitempty"`
	ImageUrls            []string          `protobuf:"bytes,8,rep,name=image_urls,json=imageUrls,proto3" json:"image_urls,omitempty"`
	Dimensions           map[string]string `protobuf:"bytes,9,rep,name=dimensions,proto3" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Price                uint64            `protobuf:"varint,10,opt,name=price,proto3" json:"price,omitempty"`
	PriceCurrency        string            `protobuf:"bytes,11,opt,name=price_currency,json=priceCurrency,proto3" json:"price_currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_2532e29ffc1dd680, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Item) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *Item) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Item) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Item) GetImageUrls() []string {
	if m != nil {
		return m.ImageUrls
	}
	return nil
}

func (m *Item) GetDimensions() map[string]string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Item) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Item) GetPriceCurrency() string {
	if m != nil {
		return m.PriceCurrency
	}
	return ""
}

type NormalizedItem struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Source               string       `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Lang                 Lang         `protobuf:"varint,3,opt,name=lang,proto3,enum=sizematch.protobuf.items.Lang" json:"lang,omitempty"`
	Urls                 []string     `protobuf:"bytes,4,rep,name=urls,proto3" json:"urls,omitempty"`
	Name                 string       `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description          string       `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Categories           []string     `protobuf:"bytes,7,rep,name=categories,proto3" json:"categories,omitempty"`
	ImageUrls            []string     `protobuf:"bytes,8,rep,name=image_urls,json=imageUrls,proto3" json:"image_urls,omitempty"`
	Dimensions           []*Dimension `protobuf:"bytes,9,rep,name=dimensions,proto3" json:"dimensions,omitempty"`
	Price                *Price       `protobuf:"bytes,10,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *NormalizedItem) Reset()         { *m = NormalizedItem{} }
func (m *NormalizedItem) String() string { return proto.CompactTextString(m) }
func (*NormalizedItem) ProtoMessage()    {}
func (*NormalizedItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_2532e29ffc1dd680, []int{1}
}

func (m *NormalizedItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NormalizedItem.Unmarshal(m, b)
}
func (m *NormalizedItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NormalizedItem.Marshal(b, m, deterministic)
}
func (m *NormalizedItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NormalizedItem.Merge(m, src)
}
func (m *NormalizedItem) XXX_Size() int {
	return xxx_messageInfo_NormalizedItem.Size(m)
}
func (m *NormalizedItem) XXX_DiscardUnknown() {
	xxx_messageInfo_NormalizedItem.DiscardUnknown(m)
}

var xxx_messageInfo_NormalizedItem proto.InternalMessageInfo

func (m *NormalizedItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NormalizedItem) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *NormalizedItem) GetLang() Lang {
	if m != nil {
		return m.Lang
	}
	return Lang_EN
}

func (m *NormalizedItem) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

func (m *NormalizedItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NormalizedItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NormalizedItem) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *NormalizedItem) GetImageUrls() []string {
	if m != nil {
		return m.ImageUrls
	}
	return nil
}

func (m *NormalizedItem) GetDimensions() []*Dimension {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *NormalizedItem) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

type Dimension struct {
	Name                 Dimension_Name `protobuf:"varint,1,opt,name=name,proto3,enum=sizematch.protobuf.items.Dimension_Name" json:"name,omitempty"`
	Value                float64        `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	Unit                 Dimension_Unit `protobuf:"varint,3,opt,name=unit,proto3,enum=sizematch.protobuf.items.Dimension_Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Dimension) Reset()         { *m = Dimension{} }
func (m *Dimension) String() string { return proto.CompactTextString(m) }
func (*Dimension) ProtoMessage()    {}
func (*Dimension) Descriptor() ([]byte, []int) {
	return fileDescriptor_2532e29ffc1dd680, []int{2}
}

func (m *Dimension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dimension.Unmarshal(m, b)
}
func (m *Dimension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dimension.Marshal(b, m, deterministic)
}
func (m *Dimension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dimension.Merge(m, src)
}
func (m *Dimension) XXX_Size() int {
	return xxx_messageInfo_Dimension.Size(m)
}
func (m *Dimension) XXX_DiscardUnknown() {
	xxx_messageInfo_Dimension.DiscardUnknown(m)
}

var xxx_messageInfo_Dimension proto.InternalMessageInfo

func (m *Dimension) GetName() Dimension_Name {
	if m != nil {
		return m.Name
	}
	return Dimension_HEIGHT
}

func (m *Dimension) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Dimension) GetUnit() Dimension_Unit {
	if m != nil {
		return m.Unit
	}
	return Dimension_MM
}

type Price struct {
	Price                float64        `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty"`
	Currency             Price_Currency `protobuf:"varint,2,opt,name=currency,proto3,enum=sizematch.protobuf.items.Price_Currency" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_2532e29ffc1dd680, []int{3}
}

func (m *Price) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Price.Unmarshal(m, b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Price.Marshal(b, m, deterministic)
}
func (m *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(m, src)
}
func (m *Price) XXX_Size() int {
	return xxx_messageInfo_Price.Size(m)
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func (m *Price) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Price) GetCurrency() Price_Currency {
	if m != nil {
		return m.Currency
	}
	return Price_EUR
}

func init() {
	proto.RegisterEnum("sizematch.protobuf.items.Lang", Lang_name, Lang_value)
	proto.RegisterEnum("sizematch.protobuf.items.Dimension_Name", Dimension_Name_name, Dimension_Name_value)
	proto.RegisterEnum("sizematch.protobuf.items.Dimension_Unit", Dimension_Unit_name, Dimension_Unit_value)
	proto.RegisterEnum("sizematch.protobuf.items.Price_Currency", Price_Currency_name, Price_Currency_value)
	proto.RegisterType((*Item)(nil), "sizematch.protobuf.items.Item")
	proto.RegisterMapType((map[string]string)(nil), "sizematch.protobuf.items.Item.DimensionsEntry")
	proto.RegisterType((*NormalizedItem)(nil), "sizematch.protobuf.items.NormalizedItem")
	proto.RegisterType((*Dimension)(nil), "sizematch.protobuf.items.Dimension")
	proto.RegisterType((*Price)(nil), "sizematch.protobuf.items.Price")
}

func init() {
	proto.RegisterFile("items/items.proto", fileDescriptor_2532e29ffc1dd680)
}

var fileDescriptor_2532e29ffc1dd680 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xdd, 0x8a, 0x13, 0x4d,
	0x10, 0x4d, 0xcf, 0x4f, 0x7e, 0x2a, 0x7c, 0xf9, 0xda, 0x46, 0x96, 0x46, 0x74, 0x0d, 0x23, 0x42,
	0x10, 0x19, 0x21, 0x2a, 0x88, 0xe8, 0xcd, 0x26, 0x31, 0x09, 0x9a, 0x10, 0x86, 0x84, 0x05, 0x6f,
	0x96, 0xd9, 0x49, 0x1b, 0x1b, 0x33, 0x3d, 0xa1, 0x67, 0x46, 0xc8, 0xde, 0xed, 0x53, 0xf8, 0x9c,
	0xbe, 0x81, 0x54, 0xcf, 0x24, 0x66, 0xc5, 0xe0, 0x7a, 0xe7, 0x4d, 0xa6, 0xfa, 0x74, 0xd5, 0xa9,
	0xae, 0x73, 0x2a, 0x70, 0x47, 0x66, 0x22, 0x4e, 0x9f, 0x99, 0x5f, 0x7f, 0xa3, 0x93, 0x2c, 0x61,
	0x3c, 0x95, 0x57, 0x22, 0x0e, 0xb3, 0xe8, 0x73, 0x01, 0x5c, 0xe6, 0x9f, 0x7c, 0x73, 0xef, 0x7d,
	0xb3, 0xc1, 0x19, 0x67, 0x22, 0x66, 0x2d, 0xb0, 0xe4, 0x92, 0x93, 0x36, 0xe9, 0x34, 0x02, 0x4b,
	0x2e, 0xd9, 0x09, 0x54, 0xd3, 0x24, 0xd7, 0x91, 0xe0, 0x96, 0xc1, 0xca, 0x13, 0x63, 0xe0, 0xac,
	0x43, 0xb5, 0xe2, 0xb6, 0x41, 0x4d, 0x8c, 0x58, 0xae, 0xd7, 0x29, 0x77, 0xda, 0x36, 0x62, 0x18,
	0x23, 0xa6, 0xc2, 0x58, 0x70, 0xb7, 0xc8, 0xc3, 0x98, 0xb5, 0xa1, 0xb9, 0x14, 0x69, 0xa4, 0xe5,
	0x26, 0x93, 0x89, 0xe2, 0x55, 0x73, 0x75, 0x08, 0xb1, 0x53, 0x80, 0x28, 0xcc, 0xc4, 0x2a, 0xd1,
	0x52, 0xa4, 0xbc, 0x66, 0xf8, 0x0e, 0x10, 0xf6, 0x00, 0x40, 0xc6, 0xe1, 0x4a, 0x5c, 0x98, 0x7e,
	0x75, 0x73, 0xdf, 0x30, 0xc8, 0x02, 0x9b, 0x4e, 0x01, 0x96, 0x32, 0x16, 0x2a, 0x95, 0x89, 0x4a,
	0x79, 0xa3, 0x6d, 0x77, 0x9a, 0x5d, 0xdf, 0x3f, 0x36, 0xbc, 0x8f, 0x83, 0xfb, 0xfd, 0x7d, 0xc1,
	0x40, 0x65, 0x7a, 0x1b, 0x1c, 0x30, 0xb0, 0xbb, 0xe0, 0x6e, 0xb4, 0x8c, 0x04, 0x87, 0x36, 0xe9,
	0x38, 0x41, 0x71, 0x60, 0x8f, 0xa1, 0x65, 0x82, 0x8b, 0x28, 0xd7, 0x5a, 0xa8, 0x68, 0xcb, 0x9b,
	0x66, 0x92, 0xff, 0x0c, 0xda, 0x2b, 0xc1, 0x7b, 0x6f, 0xe1, 0xff, 0x5f, 0xb8, 0x19, 0x05, 0xfb,
	0x8b, 0xd8, 0x96, 0x2a, 0x63, 0x88, 0x1d, 0xbe, 0x86, 0xeb, 0x7c, 0xa7, 0x72, 0x71, 0x78, 0x6d,
	0xbd, 0x22, 0xde, 0x77, 0x0b, 0x5a, 0xd3, 0x44, 0xc7, 0xe1, 0x5a, 0x5e, 0x89, 0xe5, 0x5f, 0x79,
	0xd4, 0x3d, 0xf0, 0xa8, 0xd5, 0x3d, 0x3d, 0x2e, 0xc0, 0x87, 0x50, 0xad, 0xfe, 0x35, 0x0f, 0x7b,
	0xbf, 0xf1, 0xf0, 0xd1, 0xf1, 0x11, 0xf6, 0x12, 0xdf, 0x30, 0xee, 0xe5, 0xa1, 0x71, 0xcd, 0xee,
	0xc3, 0xe3, 0xf5, 0x33, 0x4c, 0x2b, 0x9d, 0xf5, 0xae, 0x2d, 0x68, 0xec, 0x09, 0xd9, 0x9b, 0x72,
	0x7c, 0x62, 0x64, 0xec, 0xdc, 0xe2, 0x0d, 0xfe, 0x34, 0x8c, 0x45, 0x29, 0xd4, 0x0d, 0x67, 0x49,
	0xe9, 0x2c, 0x72, 0xe6, 0x4a, 0x66, 0xa5, 0x35, 0xb7, 0xe2, 0x5c, 0x28, 0x99, 0x05, 0xa6, 0xca,
	0x7b, 0x01, 0x0e, 0x76, 0x60, 0x00, 0xd5, 0xd1, 0x60, 0x3c, 0x1c, 0xcd, 0x69, 0x85, 0x35, 0xc0,
	0x3d, 0x1f, 0xf7, 0xe7, 0x23, 0x4a, 0x30, 0xec, 0x0f, 0x66, 0xf3, 0x11, 0xb5, 0x30, 0xe3, 0xbc,
	0xc8, 0xb0, 0xbd, 0xa7, 0xe0, 0x20, 0x07, 0xab, 0x82, 0x35, 0x99, 0xd0, 0x0a, 0x7e, 0x7b, 0x13,
	0x4a, 0x98, 0x0b, 0x64, 0x42, 0x2d, 0xfc, 0x0c, 0xa9, 0x8d, 0xe8, 0xfb, 0x21, 0x75, 0xbc, 0x6b,
	0x02, 0xae, 0x11, 0xe5, 0xe7, 0xf6, 0x93, 0x62, 0x82, 0x62, 0xfb, 0xfb, 0x50, 0xdf, 0xef, 0xbd,
	0xf5, 0xa7, 0x29, 0x0c, 0x91, 0xbf, 0xfb, 0x4b, 0x04, 0xfb, 0x4a, 0xef, 0x3e, 0xd4, 0x77, 0x28,
	0xab, 0x81, 0x3d, 0x58, 0x04, 0xb4, 0x82, 0xc1, 0xf0, 0x6c, 0x46, 0xc9, 0x93, 0x13, 0x70, 0x70,
	0x35, 0xf1, 0x4d, 0x83, 0x69, 0xf1, 0xe2, 0x77, 0x01, 0x25, 0x67, 0xb5, 0x8f, 0xae, 0xe1, 0xbd,
	0xac, 0x9a, 0x3e, 0xcf, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x62, 0x6a, 0xed, 0xa0, 0xec, 0x04,
	0x00, 0x00,
}
