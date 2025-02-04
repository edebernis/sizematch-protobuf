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
	Dimension_HEIGHT    Dimension_Name = 0
	Dimension_WIDTH     Dimension_Name = 1
	Dimension_DEPTH     Dimension_Name = 2
	Dimension_WEIGHT    Dimension_Name = 3
	Dimension_LENGTH    Dimension_Name = 4
	Dimension_DIAMETER  Dimension_Name = 5
	Dimension_VOLUME    Dimension_Name = 6
	Dimension_THICKNESS Dimension_Name = 7
)

var Dimension_Name_name = map[int32]string{
	0: "HEIGHT",
	1: "WIDTH",
	2: "DEPTH",
	3: "WEIGHT",
	4: "LENGTH",
	5: "DIAMETER",
	6: "VOLUME",
	7: "THICKNESS",
}

var Dimension_Name_value = map[string]int32{
	"HEIGHT":    0,
	"WIDTH":     1,
	"DEPTH":     2,
	"WEIGHT":    3,
	"LENGTH":    4,
	"DIAMETER":  5,
	"VOLUME":    6,
	"THICKNESS": 7,
}

func (x Dimension_Name) String() string {
	return proto.EnumName(Dimension_Name_name, int32(x))
}

func (Dimension_Name) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2532e29ffc1dd680, []int{2, 0}
}

type Dimension_Unit int32

const (
	Dimension_MM  Dimension_Unit = 0
	Dimension_CM  Dimension_Unit = 1
	Dimension_M   Dimension_Unit = 2
	Dimension_G   Dimension_Unit = 3
	Dimension_KG  Dimension_Unit = 4
	Dimension_MM2 Dimension_Unit = 5
	Dimension_CM2 Dimension_Unit = 6
	Dimension_M2  Dimension_Unit = 7
	Dimension_MM3 Dimension_Unit = 8
	Dimension_CM3 Dimension_Unit = 9
	Dimension_M3  Dimension_Unit = 10
	Dimension_L   Dimension_Unit = 11
	Dimension_CL  Dimension_Unit = 12
	Dimension_ML  Dimension_Unit = 13
)

var Dimension_Unit_name = map[int32]string{
	0:  "MM",
	1:  "CM",
	2:  "M",
	3:  "G",
	4:  "KG",
	5:  "MM2",
	6:  "CM2",
	7:  "M2",
	8:  "MM3",
	9:  "CM3",
	10: "M3",
	11: "L",
	12: "CL",
	13: "ML",
}

var Dimension_Unit_value = map[string]int32{
	"MM":  0,
	"CM":  1,
	"M":   2,
	"G":   3,
	"KG":  4,
	"MM2": 5,
	"CM2": 6,
	"M2":  7,
	"MM3": 8,
	"CM3": 9,
	"M3":  10,
	"L":   11,
	"CL":  12,
	"ML":  13,
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
	Lang                 Lang              `protobuf:"varint,3,opt,name=lang,proto3,enum=sizematch.protobuf.items.Lang" json:"lang,omitempty"`
	Urls                 []string          `protobuf:"bytes,4,rep,name=urls,proto3" json:"urls,omitempty"`
	Name                 string            `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description          string            `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Categories           []string          `protobuf:"bytes,7,rep,name=categories,proto3" json:"categories,omitempty"`
	ImageUrls            []string          `protobuf:"bytes,8,rep,name=image_urls,json=imageUrls,proto3" json:"image_urls,omitempty"`
	Dimensions           map[string]string `protobuf:"bytes,9,rep,name=dimensions,proto3" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Price                float64           `protobuf:"fixed64,10,opt,name=price,proto3" json:"price,omitempty"`
	PriceCurrency        string            `protobuf:"bytes,11,opt,name=price_currency,json=priceCurrency,proto3" json:"price_currency,omitempty"`
	Brand                string            `protobuf:"bytes,12,opt,name=brand,proto3" json:"brand,omitempty"`
	Colors               []string          `protobuf:"bytes,13,rep,name=colors,proto3" json:"colors,omitempty"`
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

func (m *Item) GetLang() Lang {
	if m != nil {
		return m.Lang
	}
	return Lang_EN
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

func (m *Item) GetPrice() float64 {
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

func (m *Item) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Item) GetColors() []string {
	if m != nil {
		return m.Colors
	}
	return nil
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
	Brand                string       `protobuf:"bytes,11,opt,name=brand,proto3" json:"brand,omitempty"`
	Colors               []string     `protobuf:"bytes,12,rep,name=colors,proto3" json:"colors,omitempty"`
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

func (m *NormalizedItem) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *NormalizedItem) GetColors() []string {
	if m != nil {
		return m.Colors
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
	Amount               float64        `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
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

func (m *Price) GetAmount() float64 {
	if m != nil {
		return m.Amount
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
	// 662 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xdd, 0xc9, 0x47, 0xdb, 0xdc, 0x7e, 0x38, 0x0e, 0xb2, 0x04, 0xd1, 0xb5, 0x54, 0x84, 0xe2,
	0x43, 0x84, 0x14, 0x41, 0x44, 0x1f, 0xdc, 0x36, 0xb6, 0x65, 0x9b, 0xba, 0x64, 0x5b, 0x17, 0x7c,
	0x59, 0xb2, 0xe9, 0x58, 0x07, 0x9b, 0x64, 0xc9, 0x87, 0xb2, 0xfb, 0x26, 0xfe, 0x10, 0x7f, 0x9d,
	0xff, 0x43, 0xe6, 0x26, 0x1b, 0xba, 0xb2, 0xc5, 0x7d, 0xf7, 0xa5, 0xb9, 0xf7, 0x9c, 0xb9, 0xe7,
	0xce, 0xdc, 0x33, 0x53, 0xb8, 0x2f, 0x32, 0x1e, 0xa6, 0x2f, 0xf0, 0xd7, 0xba, 0x48, 0xe2, 0x2c,
	0x66, 0x66, 0x2a, 0xae, 0x78, 0xe8, 0x67, 0xc1, 0x97, 0x02, 0x38, 0xcf, 0x3f, 0x5b, 0xc8, 0xf7,
	0x7e, 0xab, 0xa0, 0x4d, 0x33, 0x1e, 0xb2, 0x0e, 0x28, 0x62, 0x65, 0x92, 0x2e, 0xe9, 0x1b, 0x9e,
	0x22, 0x56, 0x6c, 0x1f, 0x6a, 0x69, 0x9c, 0x27, 0x01, 0x37, 0x15, 0xc4, 0xca, 0x8c, 0xd9, 0xa0,
	0x6d, 0xfc, 0x68, 0x6d, 0xaa, 0x5d, 0xd2, 0xef, 0xd8, 0x07, 0xd6, 0x2e, 0x65, 0x6b, 0xe6, 0x47,
	0x6b, 0x0f, 0xd7, 0x32, 0x06, 0x5a, 0x9e, 0x6c, 0x52, 0x53, 0xeb, 0xaa, 0x7d, 0xc3, 0xc3, 0x58,
	0x62, 0x91, 0x1f, 0x72, 0x53, 0x47, 0x75, 0x8c, 0x59, 0x17, 0x9a, 0x2b, 0x9e, 0x06, 0x89, 0xb8,
	0xc8, 0x44, 0x1c, 0x99, 0x35, 0xa4, 0xb6, 0x21, 0x76, 0x00, 0x10, 0xf8, 0x19, 0x5f, 0xc7, 0x89,
	0xe0, 0xa9, 0x59, 0x47, 0xbd, 0x2d, 0x84, 0x3d, 0x06, 0x10, 0xa1, 0xbf, 0xe6, 0x67, 0xd8, 0xaf,
	0x81, 0xbc, 0x81, 0xc8, 0x52, 0x36, 0x9d, 0x03, 0xac, 0x44, 0xc8, 0xa3, 0x54, 0xc4, 0x51, 0x6a,
	0x1a, 0x5d, 0xb5, 0xdf, 0xb4, 0xad, 0xdd, 0x47, 0x90, 0x83, 0xb1, 0x46, 0x55, 0x81, 0x13, 0x65,
	0xc9, 0xa5, 0xb7, 0xa5, 0xc0, 0x1e, 0x80, 0x7e, 0x91, 0x88, 0x80, 0x9b, 0xd0, 0x25, 0x7d, 0xe2,
	0x15, 0x09, 0x7b, 0x06, 0x1d, 0x0c, 0xce, 0x82, 0x3c, 0x49, 0x78, 0x14, 0x5c, 0x9a, 0x4d, 0x3c,
	0x49, 0x1b, 0xd1, 0x61, 0x09, 0xca, 0xe2, 0xf3, 0xc4, 0x8f, 0x56, 0x66, 0x0b, 0xd9, 0x22, 0x91,
	0x73, 0x0f, 0xe2, 0x4d, 0x9c, 0xa4, 0x66, 0x1b, 0x77, 0x5f, 0x66, 0x0f, 0xdf, 0xc2, 0xbd, 0xbf,
	0x76, 0xc2, 0x28, 0xa8, 0x5f, 0xf9, 0x65, 0xe9, 0x99, 0x0c, 0xa5, 0xe4, 0x37, 0x7f, 0x93, 0x5f,
	0x7b, 0x56, 0x24, 0xaf, 0x95, 0x57, 0xa4, 0xf7, 0x4b, 0x85, 0xce, 0x3c, 0x4e, 0x42, 0x7f, 0x23,
	0xae, 0xf8, 0xea, 0xff, 0x71, 0x7c, 0x78, 0x8b, 0xe3, 0x4f, 0x77, 0x1f, 0xa1, 0x1a, 0xf1, 0x0d,
	0x9b, 0x5f, 0x6e, 0xdb, 0xdc, 0xb4, 0x9f, 0xec, 0xae, 0x3f, 0x96, 0xcb, 0xae, 0xef, 0x41, 0x65,
	0x70, 0xf3, 0x76, 0x83, 0x5b, 0xdb, 0x06, 0xf7, 0x7e, 0xa8, 0x60, 0x54, 0xed, 0xd9, 0x9b, 0x72,
	0x58, 0x04, 0x87, 0xde, 0xbf, 0xc3, 0x8e, 0xad, 0xb9, 0x1f, 0xf2, 0x72, 0xac, 0x37, 0xee, 0x01,
	0x29, 0xef, 0x81, 0xd4, 0xcc, 0x23, 0x91, 0x95, 0x46, 0xde, 0x49, 0x73, 0x19, 0x89, 0xcc, 0xc3,
	0xaa, 0x9e, 0x00, 0x4d, 0x76, 0x60, 0x00, 0xb5, 0x89, 0x33, 0x1d, 0x4f, 0x16, 0x74, 0x8f, 0x19,
	0xa0, 0x9f, 0x4e, 0x47, 0x8b, 0x09, 0x25, 0x32, 0x1c, 0x39, 0xc7, 0x8b, 0x09, 0x55, 0xe4, 0x8a,
	0xd3, 0x62, 0x85, 0x2a, 0xe3, 0x99, 0x33, 0x1f, 0x2f, 0x26, 0x54, 0x63, 0x2d, 0x68, 0x8c, 0xa6,
	0xef, 0x5c, 0x67, 0xe1, 0x78, 0x54, 0x97, 0xcc, 0xc7, 0x0f, 0xb3, 0xa5, 0xeb, 0xd0, 0x1a, 0x6b,
	0x83, 0xb1, 0x98, 0x4c, 0x87, 0x47, 0x73, 0xe7, 0xe4, 0x84, 0xd6, 0x7b, 0xdf, 0x41, 0x93, 0x8d,
	0x59, 0x0d, 0x14, 0xd7, 0xa5, 0x7b, 0xf2, 0x3b, 0x74, 0x29, 0x61, 0x3a, 0x10, 0x97, 0x2a, 0xf2,
	0x33, 0xa6, 0xaa, 0x44, 0x8f, 0xc6, 0x54, 0x63, 0x75, 0x50, 0x5d, 0xd7, 0xa6, 0xba, 0x0c, 0x86,
	0xae, 0x4d, 0x6b, 0x58, 0x67, 0xd3, 0x7a, 0xc1, 0x0c, 0x68, 0xa3, 0x60, 0x06, 0xd4, 0x40, 0x66,
	0x40, 0x41, 0x4a, 0xcc, 0x68, 0x13, 0x85, 0x67, 0xb4, 0x85, 0xf0, 0x8c, 0xb6, 0x7b, 0x3f, 0x09,
	0xe8, 0x68, 0xa1, 0x74, 0xc9, 0x0f, 0xe3, 0x3c, 0xca, 0xd0, 0x01, 0xe2, 0x95, 0x19, 0x1b, 0x41,
	0xa3, 0x7a, 0xd5, 0xca, 0xbf, 0xe6, 0x88, 0x52, 0xd6, 0xf5, 0x83, 0xf7, 0xaa, 0xca, 0xde, 0x23,
	0x68, 0x54, 0x7f, 0x03, 0x75, 0x50, 0x9d, 0xa5, 0x47, 0xf7, 0x64, 0x30, 0x3e, 0x3c, 0xa6, 0xe4,
	0xf9, 0x3e, 0x68, 0xf2, 0x29, 0xc9, 0x5d, 0x39, 0xf3, 0xe2, 0xf8, 0xef, 0x3d, 0x4a, 0x0e, 0xeb,
	0x9f, 0x74, 0xd4, 0x3d, 0xaf, 0x61, 0x9f, 0xc1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x0e,
	0x51, 0xdd, 0xea, 0x05, 0x00, 0x00,
}
