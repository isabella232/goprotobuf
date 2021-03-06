// Code generated by protoc-gen-go.
// source: test_objects.proto
// DO NOT EDIT!

package jsonpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf2 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf4 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Widget_Color int32

const (
	Widget_RED   Widget_Color = 0
	Widget_GREEN Widget_Color = 1
	Widget_BLUE  Widget_Color = 2
)

var Widget_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}
var Widget_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Widget_Color) Enum() *Widget_Color {
	p := new(Widget_Color)
	*p = x
	return p
}
func (x Widget_Color) String() string {
	return proto.EnumName(Widget_Color_name, int32(x))
}
func (x *Widget_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Widget_Color_value, data, "Widget_Color")
	if err != nil {
		return err
	}
	*x = Widget_Color(value)
	return nil
}
func (Widget_Color) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

// Test message for holding primitive types.
type Simple struct {
	OBool            *bool    `protobuf:"varint,1,opt,name=o_bool,json=oBool" json:"o_bool,omitempty"`
	OInt32           *int32   `protobuf:"varint,2,opt,name=o_int32,json=oInt32" json:"o_int32,omitempty"`
	OInt64           *int64   `protobuf:"varint,3,opt,name=o_int64,json=oInt64" json:"o_int64,omitempty"`
	OUint32          *uint32  `protobuf:"varint,4,opt,name=o_uint32,json=oUint32" json:"o_uint32,omitempty"`
	OUint64          *uint64  `protobuf:"varint,5,opt,name=o_uint64,json=oUint64" json:"o_uint64,omitempty"`
	OSint32          *int32   `protobuf:"zigzag32,6,opt,name=o_sint32,json=oSint32" json:"o_sint32,omitempty"`
	OSint64          *int64   `protobuf:"zigzag64,7,opt,name=o_sint64,json=oSint64" json:"o_sint64,omitempty"`
	OFloat           *float32 `protobuf:"fixed32,8,opt,name=o_float,json=oFloat" json:"o_float,omitempty"`
	ODouble          *float64 `protobuf:"fixed64,9,opt,name=o_double,json=oDouble" json:"o_double,omitempty"`
	OString          *string  `protobuf:"bytes,10,opt,name=o_string,json=oString" json:"o_string,omitempty"`
	OBytes           []byte   `protobuf:"bytes,11,opt,name=o_bytes,json=oBytes" json:"o_bytes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Simple) Reset()                    { *m = Simple{} }
func (m *Simple) String() string            { return proto.CompactTextString(m) }
func (*Simple) ProtoMessage()               {}
func (*Simple) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Simple) GetOBool() bool {
	if m != nil && m.OBool != nil {
		return *m.OBool
	}
	return false
}

func (m *Simple) GetOInt32() int32 {
	if m != nil && m.OInt32 != nil {
		return *m.OInt32
	}
	return 0
}

func (m *Simple) GetOInt64() int64 {
	if m != nil && m.OInt64 != nil {
		return *m.OInt64
	}
	return 0
}

func (m *Simple) GetOUint32() uint32 {
	if m != nil && m.OUint32 != nil {
		return *m.OUint32
	}
	return 0
}

func (m *Simple) GetOUint64() uint64 {
	if m != nil && m.OUint64 != nil {
		return *m.OUint64
	}
	return 0
}

func (m *Simple) GetOSint32() int32 {
	if m != nil && m.OSint32 != nil {
		return *m.OSint32
	}
	return 0
}

func (m *Simple) GetOSint64() int64 {
	if m != nil && m.OSint64 != nil {
		return *m.OSint64
	}
	return 0
}

func (m *Simple) GetOFloat() float32 {
	if m != nil && m.OFloat != nil {
		return *m.OFloat
	}
	return 0
}

func (m *Simple) GetODouble() float64 {
	if m != nil && m.ODouble != nil {
		return *m.ODouble
	}
	return 0
}

func (m *Simple) GetOString() string {
	if m != nil && m.OString != nil {
		return *m.OString
	}
	return ""
}

func (m *Simple) GetOBytes() []byte {
	if m != nil {
		return m.OBytes
	}
	return nil
}

// Test message for holding repeated primitives.
type Repeats struct {
	RBool            []bool    `protobuf:"varint,1,rep,name=r_bool,json=rBool" json:"r_bool,omitempty"`
	RInt32           []int32   `protobuf:"varint,2,rep,name=r_int32,json=rInt32" json:"r_int32,omitempty"`
	RInt64           []int64   `protobuf:"varint,3,rep,name=r_int64,json=rInt64" json:"r_int64,omitempty"`
	RUint32          []uint32  `protobuf:"varint,4,rep,name=r_uint32,json=rUint32" json:"r_uint32,omitempty"`
	RUint64          []uint64  `protobuf:"varint,5,rep,name=r_uint64,json=rUint64" json:"r_uint64,omitempty"`
	RSint32          []int32   `protobuf:"zigzag32,6,rep,name=r_sint32,json=rSint32" json:"r_sint32,omitempty"`
	RSint64          []int64   `protobuf:"zigzag64,7,rep,name=r_sint64,json=rSint64" json:"r_sint64,omitempty"`
	RFloat           []float32 `protobuf:"fixed32,8,rep,name=r_float,json=rFloat" json:"r_float,omitempty"`
	RDouble          []float64 `protobuf:"fixed64,9,rep,name=r_double,json=rDouble" json:"r_double,omitempty"`
	RString          []string  `protobuf:"bytes,10,rep,name=r_string,json=rString" json:"r_string,omitempty"`
	RBytes           [][]byte  `protobuf:"bytes,11,rep,name=r_bytes,json=rBytes" json:"r_bytes,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Repeats) Reset()                    { *m = Repeats{} }
func (m *Repeats) String() string            { return proto.CompactTextString(m) }
func (*Repeats) ProtoMessage()               {}
func (*Repeats) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Repeats) GetRBool() []bool {
	if m != nil {
		return m.RBool
	}
	return nil
}

func (m *Repeats) GetRInt32() []int32 {
	if m != nil {
		return m.RInt32
	}
	return nil
}

func (m *Repeats) GetRInt64() []int64 {
	if m != nil {
		return m.RInt64
	}
	return nil
}

func (m *Repeats) GetRUint32() []uint32 {
	if m != nil {
		return m.RUint32
	}
	return nil
}

func (m *Repeats) GetRUint64() []uint64 {
	if m != nil {
		return m.RUint64
	}
	return nil
}

func (m *Repeats) GetRSint32() []int32 {
	if m != nil {
		return m.RSint32
	}
	return nil
}

func (m *Repeats) GetRSint64() []int64 {
	if m != nil {
		return m.RSint64
	}
	return nil
}

func (m *Repeats) GetRFloat() []float32 {
	if m != nil {
		return m.RFloat
	}
	return nil
}

func (m *Repeats) GetRDouble() []float64 {
	if m != nil {
		return m.RDouble
	}
	return nil
}

func (m *Repeats) GetRString() []string {
	if m != nil {
		return m.RString
	}
	return nil
}

func (m *Repeats) GetRBytes() [][]byte {
	if m != nil {
		return m.RBytes
	}
	return nil
}

// Test message for holding enums and nested messages.
type Widget struct {
	Color            *Widget_Color  `protobuf:"varint,1,opt,name=color,enum=jsonpb.Widget_Color" json:"color,omitempty"`
	RColor           []Widget_Color `protobuf:"varint,2,rep,name=r_color,json=rColor,enum=jsonpb.Widget_Color" json:"r_color,omitempty"`
	Simple           *Simple        `protobuf:"bytes,10,opt,name=simple" json:"simple,omitempty"`
	RSimple          []*Simple      `protobuf:"bytes,11,rep,name=r_simple,json=rSimple" json:"r_simple,omitempty"`
	Repeats          *Repeats       `protobuf:"bytes,20,opt,name=repeats" json:"repeats,omitempty"`
	RRepeats         []*Repeats     `protobuf:"bytes,21,rep,name=r_repeats,json=rRepeats" json:"r_repeats,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Widget) Reset()                    { *m = Widget{} }
func (m *Widget) String() string            { return proto.CompactTextString(m) }
func (*Widget) ProtoMessage()               {}
func (*Widget) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Widget) GetColor() Widget_Color {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return Widget_RED
}

func (m *Widget) GetRColor() []Widget_Color {
	if m != nil {
		return m.RColor
	}
	return nil
}

func (m *Widget) GetSimple() *Simple {
	if m != nil {
		return m.Simple
	}
	return nil
}

func (m *Widget) GetRSimple() []*Simple {
	if m != nil {
		return m.RSimple
	}
	return nil
}

func (m *Widget) GetRepeats() *Repeats {
	if m != nil {
		return m.Repeats
	}
	return nil
}

func (m *Widget) GetRRepeats() []*Repeats {
	if m != nil {
		return m.RRepeats
	}
	return nil
}

type Maps struct {
	MInt64Str        map[int64]string `protobuf:"bytes,1,rep,name=m_int64_str,json=mInt64Str" json:"m_int64_str,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MBoolSimple      map[bool]*Simple `protobuf:"bytes,2,rep,name=m_bool_simple,json=mBoolSimple" json:"m_bool_simple,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Maps) Reset()                    { *m = Maps{} }
func (m *Maps) String() string            { return proto.CompactTextString(m) }
func (*Maps) ProtoMessage()               {}
func (*Maps) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Maps) GetMInt64Str() map[int64]string {
	if m != nil {
		return m.MInt64Str
	}
	return nil
}

func (m *Maps) GetMBoolSimple() map[bool]*Simple {
	if m != nil {
		return m.MBoolSimple
	}
	return nil
}

type MsgWithOneof struct {
	// Types that are valid to be assigned to Union:
	//	*MsgWithOneof_Title
	//	*MsgWithOneof_Salary
	//	*MsgWithOneof_Country
	Union            isMsgWithOneof_Union `protobuf_oneof:"union"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *MsgWithOneof) Reset()                    { *m = MsgWithOneof{} }
func (m *MsgWithOneof) String() string            { return proto.CompactTextString(m) }
func (*MsgWithOneof) ProtoMessage()               {}
func (*MsgWithOneof) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type isMsgWithOneof_Union interface {
	isMsgWithOneof_Union()
}

type MsgWithOneof_Title struct {
	Title string `protobuf:"bytes,1,opt,name=title,oneof"`
}
type MsgWithOneof_Salary struct {
	Salary int64 `protobuf:"varint,2,opt,name=salary,oneof"`
}
type MsgWithOneof_Country struct {
	Country string `protobuf:"bytes,3,opt,name=Country,json=country,oneof"`
}

func (*MsgWithOneof_Title) isMsgWithOneof_Union()   {}
func (*MsgWithOneof_Salary) isMsgWithOneof_Union()  {}
func (*MsgWithOneof_Country) isMsgWithOneof_Union() {}

func (m *MsgWithOneof) GetUnion() isMsgWithOneof_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *MsgWithOneof) GetTitle() string {
	if x, ok := m.GetUnion().(*MsgWithOneof_Title); ok {
		return x.Title
	}
	return ""
}

func (m *MsgWithOneof) GetSalary() int64 {
	if x, ok := m.GetUnion().(*MsgWithOneof_Salary); ok {
		return x.Salary
	}
	return 0
}

func (m *MsgWithOneof) GetCountry() string {
	if x, ok := m.GetUnion().(*MsgWithOneof_Country); ok {
		return x.Country
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MsgWithOneof) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MsgWithOneof_OneofMarshaler, _MsgWithOneof_OneofUnmarshaler, _MsgWithOneof_OneofSizer, []interface{}{
		(*MsgWithOneof_Title)(nil),
		(*MsgWithOneof_Salary)(nil),
		(*MsgWithOneof_Country)(nil),
	}
}

func _MsgWithOneof_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MsgWithOneof)
	// union
	switch x := m.Union.(type) {
	case *MsgWithOneof_Title:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Title)
	case *MsgWithOneof_Salary:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Salary))
	case *MsgWithOneof_Country:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Country)
	case nil:
	default:
		return fmt.Errorf("MsgWithOneof.Union has unexpected type %T", x)
	}
	return nil
}

func _MsgWithOneof_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MsgWithOneof)
	switch tag {
	case 1: // union.title
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Union = &MsgWithOneof_Title{x}
		return true, err
	case 2: // union.salary
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &MsgWithOneof_Salary{int64(x)}
		return true, err
	case 3: // union.Country
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Union = &MsgWithOneof_Country{x}
		return true, err
	default:
		return false, nil
	}
}

func _MsgWithOneof_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MsgWithOneof)
	// union
	switch x := m.Union.(type) {
	case *MsgWithOneof_Title:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Title)))
		n += len(x.Title)
	case *MsgWithOneof_Salary:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Salary))
	case *MsgWithOneof_Country:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Country)))
		n += len(x.Country)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Real struct {
	Value                        *float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *Real) Reset()                    { *m = Real{} }
func (m *Real) String() string            { return proto.CompactTextString(m) }
func (*Real) ProtoMessage()               {}
func (*Real) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

var extRange_Real = []proto.ExtensionRange{
	{100, 536870911},
}

func (*Real) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Real
}

func (m *Real) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Complex struct {
	Imaginary                    *float64 `protobuf:"fixed64,1,opt,name=imaginary" json:"imaginary,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *Complex) Reset()                    { *m = Complex{} }
func (m *Complex) String() string            { return proto.CompactTextString(m) }
func (*Complex) ProtoMessage()               {}
func (*Complex) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

var extRange_Complex = []proto.ExtensionRange{
	{100, 536870911},
}

func (*Complex) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Complex
}

func (m *Complex) GetImaginary() float64 {
	if m != nil && m.Imaginary != nil {
		return *m.Imaginary
	}
	return 0
}

var E_Complex_RealExtension = &proto.ExtensionDesc{
	ExtendedType:  (*Real)(nil),
	ExtensionType: (*Complex)(nil),
	Field:         123,
	Name:          "jsonpb.Complex.real_extension",
	Tag:           "bytes,123,opt,name=real_extension,json=realExtension",
}

type MsgWithListVal struct {
	Mapped           map[string]*google_protobuf2.ListValue `protobuf:"bytes,1,rep,name=mapped" json:"mapped,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Repeated         []*google_protobuf2.ListValue          `protobuf:"bytes,2,rep,name=repeated" json:"repeated,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *MsgWithListVal) Reset()                    { *m = MsgWithListVal{} }
func (m *MsgWithListVal) String() string            { return proto.CompactTextString(m) }
func (*MsgWithListVal) ProtoMessage()               {}
func (*MsgWithListVal) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *MsgWithListVal) GetMapped() map[string]*google_protobuf2.ListValue {
	if m != nil {
		return m.Mapped
	}
	return nil
}

func (m *MsgWithListVal) GetRepeated() []*google_protobuf2.ListValue {
	if m != nil {
		return m.Repeated
	}
	return nil
}

type KnownTypes struct {
	An               *google_protobuf.Any          `protobuf:"bytes,14,opt,name=an" json:"an,omitempty"`
	Dur              *google_protobuf1.Duration    `protobuf:"bytes,1,opt,name=dur" json:"dur,omitempty"`
	St               *google_protobuf2.Struct      `protobuf:"bytes,12,opt,name=st" json:"st,omitempty"`
	Ts               *google_protobuf3.Timestamp   `protobuf:"bytes,2,opt,name=ts" json:"ts,omitempty"`
	Dbl              *google_protobuf4.DoubleValue `protobuf:"bytes,3,opt,name=dbl" json:"dbl,omitempty"`
	Flt              *google_protobuf4.FloatValue  `protobuf:"bytes,4,opt,name=flt" json:"flt,omitempty"`
	I64              *google_protobuf4.Int64Value  `protobuf:"bytes,5,opt,name=i64" json:"i64,omitempty"`
	U64              *google_protobuf4.UInt64Value `protobuf:"bytes,6,opt,name=u64" json:"u64,omitempty"`
	I32              *google_protobuf4.Int32Value  `protobuf:"bytes,7,opt,name=i32" json:"i32,omitempty"`
	U32              *google_protobuf4.UInt32Value `protobuf:"bytes,8,opt,name=u32" json:"u32,omitempty"`
	Bool             *google_protobuf4.BoolValue   `protobuf:"bytes,9,opt,name=bool" json:"bool,omitempty"`
	Str              *google_protobuf4.StringValue `protobuf:"bytes,10,opt,name=str" json:"str,omitempty"`
	Bytes            *google_protobuf4.BytesValue  `protobuf:"bytes,11,opt,name=bytes" json:"bytes,omitempty"`
	Listval          *google_protobuf2.ListValue   `protobuf:"bytes,15,opt,name=listval" json:"listval,omitempty"`
	Val              *google_protobuf2.Value       `protobuf:"bytes,16,opt,name=val" json:"val,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *KnownTypes) Reset()                    { *m = KnownTypes{} }
func (m *KnownTypes) String() string            { return proto.CompactTextString(m) }
func (*KnownTypes) ProtoMessage()               {}
func (*KnownTypes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *KnownTypes) GetAn() *google_protobuf.Any {
	if m != nil {
		return m.An
	}
	return nil
}

func (m *KnownTypes) GetDur() *google_protobuf1.Duration {
	if m != nil {
		return m.Dur
	}
	return nil
}

func (m *KnownTypes) GetSt() *google_protobuf2.Struct {
	if m != nil {
		return m.St
	}
	return nil
}

func (m *KnownTypes) GetTs() *google_protobuf3.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *KnownTypes) GetDbl() *google_protobuf4.DoubleValue {
	if m != nil {
		return m.Dbl
	}
	return nil
}

func (m *KnownTypes) GetFlt() *google_protobuf4.FloatValue {
	if m != nil {
		return m.Flt
	}
	return nil
}

func (m *KnownTypes) GetI64() *google_protobuf4.Int64Value {
	if m != nil {
		return m.I64
	}
	return nil
}

func (m *KnownTypes) GetU64() *google_protobuf4.UInt64Value {
	if m != nil {
		return m.U64
	}
	return nil
}

func (m *KnownTypes) GetI32() *google_protobuf4.Int32Value {
	if m != nil {
		return m.I32
	}
	return nil
}

func (m *KnownTypes) GetU32() *google_protobuf4.UInt32Value {
	if m != nil {
		return m.U32
	}
	return nil
}

func (m *KnownTypes) GetBool() *google_protobuf4.BoolValue {
	if m != nil {
		return m.Bool
	}
	return nil
}

func (m *KnownTypes) GetStr() *google_protobuf4.StringValue {
	if m != nil {
		return m.Str
	}
	return nil
}

func (m *KnownTypes) GetBytes() *google_protobuf4.BytesValue {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *KnownTypes) GetListval() *google_protobuf2.ListValue {
	if m != nil {
		return m.Listval
	}
	return nil
}

func (m *KnownTypes) GetVal() *google_protobuf2.Value {
	if m != nil {
		return m.Val
	}
	return nil
}

var E_Name = &proto.ExtensionDesc{
	ExtendedType:  (*Real)(nil),
	ExtensionType: (*string)(nil),
	Field:         124,
	Name:          "jsonpb.name",
	Tag:           "bytes,124,opt,name=name",
}

func init() {
	proto.RegisterType((*Simple)(nil), "jsonpb.Simple")
	proto.RegisterType((*Repeats)(nil), "jsonpb.Repeats")
	proto.RegisterType((*Widget)(nil), "jsonpb.Widget")
	proto.RegisterType((*Maps)(nil), "jsonpb.Maps")
	proto.RegisterType((*MsgWithOneof)(nil), "jsonpb.MsgWithOneof")
	proto.RegisterType((*Real)(nil), "jsonpb.Real")
	proto.RegisterType((*Complex)(nil), "jsonpb.Complex")
	proto.RegisterType((*MsgWithListVal)(nil), "jsonpb.MsgWithListVal")
	proto.RegisterType((*KnownTypes)(nil), "jsonpb.KnownTypes")
	proto.RegisterEnum("jsonpb.Widget_Color", Widget_Color_name, Widget_Color_value)
	proto.RegisterExtension(E_Complex_RealExtension)
	proto.RegisterExtension(E_Name)
}

func init() { proto.RegisterFile("test_objects.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1106 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xae, 0x25, 0xcb, 0xb2, 0xd7, 0x49, 0x6a, 0x76, 0xd2, 0x56, 0x31, 0x01, 0x3a, 0x1e, 0x28,
	0xa5, 0x80, 0x0b, 0x6e, 0x26, 0xc3, 0x04, 0x6e, 0x9a, 0xc6, 0xfc, 0x0c, 0x4d, 0x99, 0xd9, 0x34,
	0xf4, 0xd2, 0x23, 0x27, 0x8a, 0x51, 0x91, 0xb5, 0x9e, 0xd5, 0x8a, 0xd4, 0x03, 0x17, 0x3c, 0x04,
	0xaf, 0x00, 0x8f, 0xc0, 0x13, 0xf0, 0x16, 0xbc, 0x0c, 0xe7, 0x9c, 0x5d, 0x49, 0xfe, 0xa9, 0x4b,
	0x6e, 0xe2, 0xa3, 0xef, 0x3b, 0x9f, 0x8e, 0xbe, 0x73, 0x76, 0x0f, 0xe3, 0x3a, 0xca, 0xf4, 0x48,
	0x8e, 0x5f, 0x46, 0x17, 0x3a, 0xeb, 0xcf, 0x94, 0xd4, 0x92, 0x37, 0x5e, 0x66, 0x32, 0x9d, 0x8d,
	0xbb, 0x7b, 0x13, 0x29, 0x27, 0x49, 0xf4, 0x90, 0x9e, 0x8e, 0xf3, 0xab, 0x87, 0x61, 0x3a, 0x37,
	0x94, 0xee, 0xbb, 0xab, 0xd0, 0x65, 0xae, 0x42, 0x1d, 0xcb, 0xd4, 0xe2, 0xfb, 0xab, 0x78, 0xa6,
	0x55, 0x7e, 0xa1, 0x2d, 0xfa, 0xde, 0x2a, 0xaa, 0xe3, 0x29, 0x94, 0x11, 0x4e, 0x67, 0x9b, 0xe4,
	0xaf, 0x55, 0x38, 0x9b, 0x45, 0xca, 0x56, 0xd8, 0xfb, 0xd3, 0x61, 0x8d, 0xb3, 0x78, 0x3a, 0x4b,
	0x22, 0x7e, 0x8b, 0x35, 0xe4, 0x68, 0x2c, 0x65, 0x12, 0xd4, 0xee, 0xd6, 0xee, 0x37, 0x85, 0x27,
	0x8f, 0x21, 0xe0, 0x77, 0x98, 0x2f, 0x47, 0x71, 0xaa, 0x1f, 0x0d, 0x02, 0x07, 0x9e, 0x7b, 0xa2,
	0x21, 0xbf, 0xc3, 0xa8, 0x04, 0x0e, 0x0f, 0x02, 0x17, 0x00, 0xd7, 0x00, 0x87, 0x07, 0x7c, 0x8f,
	0x35, 0xe5, 0x28, 0x37, 0x29, 0x75, 0x40, 0xb6, 0x85, 0x2f, 0xcf, 0x29, 0xac, 0x20, 0x48, 0xf2,
	0x00, 0xaa, 0x5b, 0xa8, 0xc8, 0xca, 0x4c, 0x56, 0x03, 0xa0, 0xb7, 0x00, 0x3a, 0x5b, 0xc8, 0xca,
	0x4c, 0x96, 0x0f, 0x10, 0xb7, 0x10, 0x64, 0x51, 0x11, 0x57, 0x89, 0x0c, 0x75, 0xd0, 0x04, 0xc4,
	0x81, 0x22, 0xbe, 0xc6, 0xc8, 0xe4, 0x5c, 0xca, 0x7c, 0x9c, 0x44, 0x41, 0x0b, 0x90, 0x1a, 0xe4,
	0x9c, 0x50, 0x68, 0xe5, 0xb4, 0x8a, 0xd3, 0x49, 0xc0, 0x00, 0x6a, 0xa1, 0x1c, 0x85, 0x46, 0x6e,
	0x3c, 0x87, 0x56, 0x06, 0x6d, 0x40, 0xb6, 0x40, 0xee, 0x18, 0xa3, 0xde, 0x5f, 0x0e, 0xf3, 0x45,
	0x34, 0x8b, 0x42, 0x9d, 0xa1, 0x51, 0xaa, 0x30, 0xca, 0x45, 0xa3, 0x54, 0x61, 0x94, 0x2a, 0x8d,
	0x72, 0xd1, 0x28, 0x55, 0x1a, 0xa5, 0x4a, 0xa3, 0x5c, 0x34, 0x4a, 0x95, 0x46, 0xa9, 0xca, 0x28,
	0x17, 0x8d, 0x52, 0x95, 0x51, 0xaa, 0x32, 0xca, 0x45, 0xa3, 0x54, 0x65, 0x94, 0xaa, 0x8c, 0x72,
	0xd1, 0x28, 0x75, 0xb6, 0x90, 0x55, 0x1a, 0xe5, 0xa2, 0x51, 0xaa, 0x32, 0x4a, 0x95, 0x46, 0xb9,
	0x68, 0x94, 0x2a, 0x8d, 0x52, 0x95, 0x51, 0x2e, 0x1a, 0xa5, 0x2a, 0xa3, 0x54, 0x65, 0x94, 0x8b,
	0x46, 0xa9, 0xca, 0x28, 0x55, 0x1a, 0xe5, 0xa2, 0x51, 0xca, 0x18, 0xf5, 0x37, 0x0c, 0xd4, 0x8b,
	0xf8, 0x72, 0x12, 0x69, 0xfe, 0x80, 0x79, 0x17, 0x32, 0x91, 0x8a, 0xe6, 0x69, 0x67, 0xb0, 0xdb,
	0x37, 0xa7, 0xa1, 0x6f, 0xe0, 0xfe, 0x13, 0xc4, 0x84, 0xa1, 0xf0, 0x4f, 0x51, 0xcf, 0xb0, 0xd1,
	0xbc, 0x4d, 0xec, 0x86, 0xa2, 0xff, 0xfc, 0x1e, 0x6b, 0x64, 0x34, 0xb5, 0xd4, 0xc0, 0xf6, 0x60,
	0xa7, 0x60, 0x9b, 0x59, 0x16, 0x16, 0xe5, 0x1f, 0x19, 0x43, 0x88, 0x89, 0x75, 0xae, 0x33, 0xd1,
	0x20, 0x4b, 0xf5, 0x95, 0x69, 0x70, 0xb0, 0x4b, 0x9a, 0x37, 0x0b, 0xa6, 0xed, 0xbb, 0x28, 0x70,
	0xfe, 0x09, 0x6b, 0xa9, 0x51, 0x41, 0xbe, 0x45, 0xb2, 0x6b, 0xe4, 0xa6, 0xb2, 0xbf, 0x7a, 0x1f,
	0x30, 0xcf, 0x14, 0xed, 0x33, 0x57, 0x0c, 0x4f, 0x3a, 0x37, 0x78, 0x8b, 0x79, 0xdf, 0x88, 0xe1,
	0xf0, 0x59, 0xa7, 0xc6, 0x9b, 0xac, 0x7e, 0xfc, 0xf4, 0x7c, 0xd8, 0x71, 0x7a, 0x7f, 0x38, 0xac,
	0x7e, 0x1a, 0xce, 0x32, 0xfe, 0x25, 0x6b, 0x4f, 0xcd, 0xb8, 0xa0, 0xf7, 0x34, 0x63, 0xed, 0xc1,
	0xdb, 0x85, 0x3e, 0x52, 0xfa, 0xa7, 0x34, 0x3f, 0xd0, 0x8a, 0x61, 0xaa, 0xd5, 0x5c, 0xb4, 0xa6,
	0x45, 0xcc, 0x1f, 0xb3, 0xed, 0x29, 0xcd, 0x66, 0xf1, 0xd5, 0x0e, 0xa5, 0xbf, 0xb3, 0x9c, 0x8e,
	0xf3, 0x6a, 0x3e, 0xdb, 0x08, 0xb4, 0xa7, 0xd5, 0x93, 0xee, 0x57, 0x6c, 0x67, 0x59, 0x9f, 0x77,
	0x98, 0xfb, 0x73, 0x34, 0xa7, 0x36, 0xba, 0x02, 0x7f, 0xf2, 0x5d, 0xe6, 0xfd, 0x12, 0x26, 0x79,
	0x44, 0x57, 0x42, 0x4b, 0x98, 0xe0, 0xc8, 0xf9, 0xa2, 0xd6, 0x7d, 0xc6, 0x3a, 0xab, 0xf2, 0x8b,
	0xf9, 0x4d, 0x93, 0xff, 0xfe, 0x62, 0xfe, 0x7a, 0x53, 0x2a, 0xbd, 0x5e, 0xc4, 0xb6, 0x4e, 0xb3,
	0xc9, 0x8b, 0x58, 0xff, 0xf4, 0x43, 0x1a, 0xc9, 0x2b, 0x7e, 0x9b, 0x79, 0x3a, 0xd6, 0xf0, 0x61,
	0xa8, 0xd6, 0xfa, 0xf6, 0x86, 0x30, 0x21, 0x0f, 0x60, 0x22, 0xc2, 0x24, 0x54, 0x73, 0x92, 0x74,
	0x01, 0xb0, 0x31, 0xef, 0x32, 0xff, 0x89, 0xcc, 0xb1, 0x10, 0xba, 0xa7, 0x30, 0xc7, 0xbf, 0x30,
	0x0f, 0x8e, 0x7d, 0xe6, 0xe5, 0x29, 0x5c, 0xb6, 0xbd, 0x7b, 0xac, 0x2e, 0xa2, 0x30, 0xa9, 0x3e,
	0xac, 0x46, 0x77, 0x86, 0x09, 0x1e, 0x34, 0x9b, 0x97, 0x9d, 0xdf, 0xe1, 0xcf, 0xe9, 0x5d, 0xa3,
	0x18, 0xd6, 0xf8, 0x8a, 0xef, 0xb3, 0x56, 0x3c, 0x0d, 0x27, 0x71, 0x8a, 0x2f, 0x35, 0xf4, 0xea,
	0x41, 0x95, 0x32, 0x38, 0x61, 0x3b, 0x0a, 0xa4, 0x47, 0xd1, 0x2b, 0x1d, 0xa5, 0x19, 0xbc, 0x8c,
	0x6f, 0x55, 0xc3, 0x12, 0x26, 0xc1, 0xaf, 0xcb, 0xd3, 0x66, 0xe5, 0xc5, 0x36, 0x26, 0x0d, 0x8b,
	0x9c, 0xde, 0xbf, 0x35, 0x68, 0x8b, 0x31, 0xe2, 0x69, 0x9c, 0xe9, 0x1f, 0xa1, 0xd6, 0x23, 0xd6,
	0x98, 0xe2, 0x65, 0x7e, 0x69, 0x67, 0xa4, 0x57, 0x36, 0x79, 0x89, 0x87, 0x3d, 0x07, 0x92, 0xe9,
	0xb4, 0xcd, 0xe0, 0x87, 0x70, 0x30, 0x68, 0x3e, 0x21, 0xdb, 0x8c, 0x48, 0xb7, 0x6f, 0x56, 0x45,
	0xbf, 0x58, 0x15, 0x7d, 0x9b, 0x9f, 0x47, 0xa2, 0xe4, 0x76, 0xcf, 0x59, 0x7b, 0x41, 0x6e, 0xb1,
	0xb3, 0x2d, 0xd3, 0xd9, 0xcf, 0x96, 0x3b, 0xfb, 0x26, 0xd5, 0x85, 0x2e, 0xff, 0xe3, 0x31, 0xf6,
	0x7d, 0x2a, 0xaf, 0xd3, 0xe7, 0xf3, 0x59, 0x94, 0xc1, 0x78, 0x38, 0x61, 0x1a, 0xec, 0x90, 0xc2,
	0xee, 0x9a, 0xc2, 0xe3, 0x74, 0x2e, 0x00, 0xe7, 0x1f, 0x33, 0x17, 0x96, 0x25, 0xbd, 0xbc, 0x3d,
	0xd8, 0x5b, 0xa3, 0x9d, 0xd8, 0x45, 0x2a, 0x90, 0xc5, 0x3f, 0x64, 0x4e, 0xa6, 0x83, 0x2d, 0xe2,
	0xde, 0x59, 0xe3, 0x9e, 0xd1, 0x52, 0x15, 0x40, 0x81, 0x5b, 0xcb, 0x81, 0x53, 0xbd, 0xa9, 0xfa,
	0xe7, 0xc5, 0x7e, 0x15, 0xc0, 0xe2, 0x7d, 0xa8, 0x60, 0x9c, 0xd0, 0x58, 0xb5, 0x07, 0xfb, 0xeb,
	0x15, 0xd0, 0x35, 0x6a, 0x3e, 0x16, 0x89, 0x70, 0xcb, 0xb9, 0x57, 0x89, 0xa6, 0xa5, 0x88, 0x47,
	0x7a, 0x95, 0x4f, 0x17, 0xb2, 0xa5, 0x03, 0x0f, 0xe9, 0xb1, 0x5d, 0x94, 0xaf, 0xa3, 0xd3, 0x21,
	0xb5, 0x74, 0xe0, 0x61, 0x35, 0x39, 0xd0, 0x1b, 0x1b, 0xaa, 0x39, 0x5f, 0xe4, 0x03, 0x91, 0xe4,
	0x61, 0x87, 0xf8, 0x9b, 0xe5, 0x1f, 0x0d, 0x0a, 0x79, 0x58, 0x2e, 0x28, 0x0f, 0xf4, 0xe6, 0x1b,
	0xe4, 0x4b, 0x7e, 0x4e, 0xfc, 0x3a, 0x2d, 0xc9, 0xd6, 0x06, 0x2b, 0xf1, 0x96, 0x30, 0x74, 0xe2,
	0xa1, 0x3e, 0xde, 0x77, 0x6c, 0x83, 0xbe, 0x59, 0x3c, 0x56, 0x1f, 0x88, 0xfc, 0x73, 0xe6, 0x55,
	0x9b, 0xfa, 0x75, 0x1f, 0x40, 0x0b, 0xc9, 0x8e, 0x1a, 0x31, 0xf9, 0x01, 0xf3, 0x13, 0x18, 0x3f,
	0x98, 0xbb, 0xe0, 0xe6, 0xff, 0x8e, 0x67, 0x41, 0xe5, 0xf7, 0x99, 0x8b, 0x19, 0x1d, 0xca, 0xb8,
	0xbd, 0x96, 0x61, 0x4b, 0x02, 0xca, 0xd1, 0x5d, 0x56, 0x4f, 0xc3, 0x69, 0xb4, 0x72, 0xc0, 0x7f,
	0xa3, 0xd3, 0x41, 0xc8, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xc5, 0x21, 0x70, 0x26, 0x0a,
	0x00, 0x00,
}
