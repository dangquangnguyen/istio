// Code generated by protoc-gen-gogo.
// source: bazel-out/local-fastbuild/genfiles/mixer/template/sample/report/go_default_library_tmpl.proto
// DO NOT EDIT!

/*
	Package istio_mixer_adapter_sample_report is a generated protocol buffer package.

	It is generated from these files:
		bazel-out/local-fastbuild/genfiles/mixer/template/sample/report/go_default_library_tmpl.proto

	It has these top-level messages:
		Type
		InstanceParam
*/
package istio_mixer_adapter_sample_report

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "istio.io/api/mixer/v1/template"
import istio_mixer_v1_config_descriptor "istio.io/api/mixer/v1/config/descriptor"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Type struct {
	Value      istio_mixer_v1_config_descriptor.ValueType            `protobuf:"varint,1,opt,name=value,proto3,enum=istio.mixer.v1.config.descriptor.ValueType" json:"value,omitempty"`
	Dimensions map[string]istio_mixer_v1_config_descriptor.ValueType `protobuf:"bytes,2,rep,name=dimensions" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=istio.mixer.v1.config.descriptor.ValueType"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptorGoDefaultLibraryTmpl, []int{0} }

func (m *Type) GetValue() istio_mixer_v1_config_descriptor.ValueType {
	if m != nil {
		return m.Value
	}
	return istio_mixer_v1_config_descriptor.VALUE_TYPE_UNSPECIFIED
}

func (m *Type) GetDimensions() map[string]istio_mixer_v1_config_descriptor.ValueType {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

type InstanceParam struct {
	Value           string            `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Dimensions      map[string]string `protobuf:"bytes,2,rep,name=dimensions" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Int64Primitive  string            `protobuf:"bytes,3,opt,name=int64Primitive,proto3" json:"int64Primitive,omitempty"`
	BoolPrimitive   string            `protobuf:"bytes,4,opt,name=boolPrimitive,proto3" json:"boolPrimitive,omitempty"`
	DoublePrimitive string            `protobuf:"bytes,5,opt,name=doublePrimitive,proto3" json:"doublePrimitive,omitempty"`
	StringPrimitive string            `protobuf:"bytes,6,opt,name=stringPrimitive,proto3" json:"stringPrimitive,omitempty"`
	Int64Map        map[string]string `protobuf:"bytes,7,rep,name=int64Map" json:"int64Map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TimeStamp       string            `protobuf:"bytes,9,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	Duration        string            `protobuf:"bytes,10,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (m *InstanceParam) Reset()      { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage() {}
func (*InstanceParam) Descriptor() ([]byte, []int) {
	return fileDescriptorGoDefaultLibraryTmpl, []int{1}
}

func (m *InstanceParam) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *InstanceParam) GetDimensions() map[string]string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *InstanceParam) GetInt64Primitive() string {
	if m != nil {
		return m.Int64Primitive
	}
	return ""
}

func (m *InstanceParam) GetBoolPrimitive() string {
	if m != nil {
		return m.BoolPrimitive
	}
	return ""
}

func (m *InstanceParam) GetDoublePrimitive() string {
	if m != nil {
		return m.DoublePrimitive
	}
	return ""
}

func (m *InstanceParam) GetStringPrimitive() string {
	if m != nil {
		return m.StringPrimitive
	}
	return ""
}

func (m *InstanceParam) GetInt64Map() map[string]string {
	if m != nil {
		return m.Int64Map
	}
	return nil
}

func (m *InstanceParam) GetTimeStamp() string {
	if m != nil {
		return m.TimeStamp
	}
	return ""
}

func (m *InstanceParam) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func init() {
	proto.RegisterType((*Type)(nil), "istio.mixer.adapter.sample.report.Type")
	proto.RegisterType((*InstanceParam)(nil), "istio.mixer.adapter.sample.report.InstanceParam")
}
func (this *Type) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Type)
	if !ok {
		that2, ok := that.(Type)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if len(this.Dimensions) != len(that1.Dimensions) {
		return false
	}
	for i := range this.Dimensions {
		if this.Dimensions[i] != that1.Dimensions[i] {
			return false
		}
	}
	return true
}
func (this *InstanceParam) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*InstanceParam)
	if !ok {
		that2, ok := that.(InstanceParam)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if len(this.Dimensions) != len(that1.Dimensions) {
		return false
	}
	for i := range this.Dimensions {
		if this.Dimensions[i] != that1.Dimensions[i] {
			return false
		}
	}
	if this.Int64Primitive != that1.Int64Primitive {
		return false
	}
	if this.BoolPrimitive != that1.BoolPrimitive {
		return false
	}
	if this.DoublePrimitive != that1.DoublePrimitive {
		return false
	}
	if this.StringPrimitive != that1.StringPrimitive {
		return false
	}
	if len(this.Int64Map) != len(that1.Int64Map) {
		return false
	}
	for i := range this.Int64Map {
		if this.Int64Map[i] != that1.Int64Map[i] {
			return false
		}
	}
	if this.TimeStamp != that1.TimeStamp {
		return false
	}
	if this.Duration != that1.Duration {
		return false
	}
	return true
}
func (this *Type) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&istio_mixer_adapter_sample_report.Type{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]istio_mixer_v1_config_descriptor.ValueType{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%#v: %#v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	if this.Dimensions != nil {
		s = append(s, "Dimensions: "+mapStringForDimensions+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InstanceParam) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 13)
	s = append(s, "&istio_mixer_adapter_sample_report.InstanceParam{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]string{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%#v: %#v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	if this.Dimensions != nil {
		s = append(s, "Dimensions: "+mapStringForDimensions+",\n")
	}
	s = append(s, "Int64Primitive: "+fmt.Sprintf("%#v", this.Int64Primitive)+",\n")
	s = append(s, "BoolPrimitive: "+fmt.Sprintf("%#v", this.BoolPrimitive)+",\n")
	s = append(s, "DoublePrimitive: "+fmt.Sprintf("%#v", this.DoublePrimitive)+",\n")
	s = append(s, "StringPrimitive: "+fmt.Sprintf("%#v", this.StringPrimitive)+",\n")
	keysForInt64Map := make([]string, 0, len(this.Int64Map))
	for k, _ := range this.Int64Map {
		keysForInt64Map = append(keysForInt64Map, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForInt64Map)
	mapStringForInt64Map := "map[string]string{"
	for _, k := range keysForInt64Map {
		mapStringForInt64Map += fmt.Sprintf("%#v: %#v,", k, this.Int64Map[k])
	}
	mapStringForInt64Map += "}"
	if this.Int64Map != nil {
		s = append(s, "Int64Map: "+mapStringForInt64Map+",\n")
	}
	s = append(s, "TimeStamp: "+fmt.Sprintf("%#v", this.TimeStamp)+",\n")
	s = append(s, "Duration: "+fmt.Sprintf("%#v", this.Duration)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGoDefaultLibraryTmpl(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Type) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Type) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(m.Value))
	}
	if len(m.Dimensions) > 0 {
		for k, _ := range m.Dimensions {
			dAtA[i] = 0x12
			i++
			v := m.Dimensions[k]
			mapSize := 1 + len(k) + sovGoDefaultLibraryTmpl(uint64(len(k))) + 1 + sovGoDefaultLibraryTmpl(uint64(v))
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(v))
		}
	}
	return i, nil
}

func (m *InstanceParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceParam) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if len(m.Dimensions) > 0 {
		for k, _ := range m.Dimensions {
			dAtA[i] = 0x12
			i++
			v := m.Dimensions[k]
			mapSize := 1 + len(k) + sovGoDefaultLibraryTmpl(uint64(len(k))) + 1 + len(v) + sovGoDefaultLibraryTmpl(uint64(len(v)))
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Int64Primitive) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.Int64Primitive)))
		i += copy(dAtA[i:], m.Int64Primitive)
	}
	if len(m.BoolPrimitive) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.BoolPrimitive)))
		i += copy(dAtA[i:], m.BoolPrimitive)
	}
	if len(m.DoublePrimitive) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.DoublePrimitive)))
		i += copy(dAtA[i:], m.DoublePrimitive)
	}
	if len(m.StringPrimitive) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.StringPrimitive)))
		i += copy(dAtA[i:], m.StringPrimitive)
	}
	if len(m.Int64Map) > 0 {
		for k, _ := range m.Int64Map {
			dAtA[i] = 0x3a
			i++
			v := m.Int64Map[k]
			mapSize := 1 + len(k) + sovGoDefaultLibraryTmpl(uint64(len(k))) + 1 + len(v) + sovGoDefaultLibraryTmpl(uint64(len(v)))
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.TimeStamp) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.TimeStamp)))
		i += copy(dAtA[i:], m.TimeStamp)
	}
	if len(m.Duration) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.Duration)))
		i += copy(dAtA[i:], m.Duration)
	}
	return i, nil
}

func encodeFixed64GoDefaultLibraryTmpl(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32GoDefaultLibraryTmpl(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGoDefaultLibraryTmpl(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Type) Size() (n int) {
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovGoDefaultLibraryTmpl(uint64(m.Value))
	}
	if len(m.Dimensions) > 0 {
		for k, v := range m.Dimensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGoDefaultLibraryTmpl(uint64(len(k))) + 1 + sovGoDefaultLibraryTmpl(uint64(v))
			n += mapEntrySize + 1 + sovGoDefaultLibraryTmpl(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *InstanceParam) Size() (n int) {
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	if len(m.Dimensions) > 0 {
		for k, v := range m.Dimensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGoDefaultLibraryTmpl(uint64(len(k))) + 1 + len(v) + sovGoDefaultLibraryTmpl(uint64(len(v)))
			n += mapEntrySize + 1 + sovGoDefaultLibraryTmpl(uint64(mapEntrySize))
		}
	}
	l = len(m.Int64Primitive)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.BoolPrimitive)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.DoublePrimitive)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.StringPrimitive)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	if len(m.Int64Map) > 0 {
		for k, v := range m.Int64Map {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGoDefaultLibraryTmpl(uint64(len(k))) + 1 + len(v) + sovGoDefaultLibraryTmpl(uint64(len(v)))
			n += mapEntrySize + 1 + sovGoDefaultLibraryTmpl(uint64(mapEntrySize))
		}
	}
	l = len(m.TimeStamp)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.Duration)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	return n
}

func sovGoDefaultLibraryTmpl(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGoDefaultLibraryTmpl(x uint64) (n int) {
	return sovGoDefaultLibraryTmpl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Type) String() string {
	if this == nil {
		return "nil"
	}
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]istio_mixer_v1_config_descriptor.ValueType{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%v: %v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	s := strings.Join([]string{`&Type{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`Dimensions:` + mapStringForDimensions + `,`,
		`}`,
	}, "")
	return s
}
func (this *InstanceParam) String() string {
	if this == nil {
		return "nil"
	}
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]string{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%v: %v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	keysForInt64Map := make([]string, 0, len(this.Int64Map))
	for k, _ := range this.Int64Map {
		keysForInt64Map = append(keysForInt64Map, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForInt64Map)
	mapStringForInt64Map := "map[string]string{"
	for _, k := range keysForInt64Map {
		mapStringForInt64Map += fmt.Sprintf("%v: %v,", k, this.Int64Map[k])
	}
	mapStringForInt64Map += "}"
	s := strings.Join([]string{`&InstanceParam{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`Dimensions:` + mapStringForDimensions + `,`,
		`Int64Primitive:` + fmt.Sprintf("%v", this.Int64Primitive) + `,`,
		`BoolPrimitive:` + fmt.Sprintf("%v", this.BoolPrimitive) + `,`,
		`DoublePrimitive:` + fmt.Sprintf("%v", this.DoublePrimitive) + `,`,
		`StringPrimitive:` + fmt.Sprintf("%v", this.StringPrimitive) + `,`,
		`Int64Map:` + mapStringForInt64Map + `,`,
		`TimeStamp:` + fmt.Sprintf("%v", this.TimeStamp) + `,`,
		`Duration:` + fmt.Sprintf("%v", this.Duration) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGoDefaultLibraryTmpl(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Type) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGoDefaultLibraryTmpl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Type: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Type: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= (istio_mixer_v1_config_descriptor.ValueType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dimensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Dimensions == nil {
				m.Dimensions = make(map[string]istio_mixer_v1_config_descriptor.ValueType)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGoDefaultLibraryTmpl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapvalue istio_mixer_v1_config_descriptor.ValueType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGoDefaultLibraryTmpl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapvalue |= (istio_mixer_v1_config_descriptor.ValueType(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Dimensions[mapkey] = mapvalue
			} else {
				var mapvalue istio_mixer_v1_config_descriptor.ValueType
				m.Dimensions[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGoDefaultLibraryTmpl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InstanceParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGoDefaultLibraryTmpl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstanceParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dimensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Dimensions == nil {
				m.Dimensions = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGoDefaultLibraryTmpl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGoDefaultLibraryTmpl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthGoDefaultLibraryTmpl
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Dimensions[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Dimensions[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int64Primitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Int64Primitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoolPrimitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BoolPrimitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DoublePrimitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DoublePrimitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringPrimitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StringPrimitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int64Map", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Int64Map == nil {
				m.Int64Map = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGoDefaultLibraryTmpl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGoDefaultLibraryTmpl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthGoDefaultLibraryTmpl
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Int64Map[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Int64Map[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeStamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Duration = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGoDefaultLibraryTmpl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGoDefaultLibraryTmpl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGoDefaultLibraryTmpl
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGoDefaultLibraryTmpl
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGoDefaultLibraryTmpl
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGoDefaultLibraryTmpl(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGoDefaultLibraryTmpl = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGoDefaultLibraryTmpl   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("bazel-out/local-fastbuild/genfiles/mixer/template/sample/report/go_default_library_tmpl.proto", fileDescriptorGoDefaultLibraryTmpl)
}

var fileDescriptorGoDefaultLibraryTmpl = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x77, 0xb6, 0x7f, 0xec, 0x8e, 0xb4, 0x95, 0xa1, 0x87, 0x10, 0x64, 0xa8, 0x8b, 0xc8,
	0x42, 0xe9, 0x84, 0xae, 0xa2, 0xa2, 0x28, 0x2a, 0x7a, 0xe8, 0x41, 0x28, 0xab, 0x28, 0x08, 0xb2,
	0x4e, 0x36, 0xb3, 0xcb, 0xe8, 0x24, 0x33, 0x4c, 0xde, 0x2c, 0x5d, 0x4f, 0x7e, 0x02, 0x11, 0xfc,
	0x12, 0x7e, 0x07, 0xbf, 0x80, 0xc7, 0xe2, 0xc9, 0xa3, 0x1b, 0x3d, 0x78, 0xec, 0xd1, 0xa3, 0x64,
	0x52, 0x37, 0xe9, 0x22, 0xda, 0x7a, 0x4b, 0xde, 0x79, 0xde, 0xdf, 0x3c, 0xcf, 0x9b, 0x37, 0xf8,
	0x79, 0xc8, 0x5f, 0x0b, 0xb5, 0xad, 0x33, 0x08, 0x94, 0x1e, 0x70, 0xb5, 0x3d, 0xe4, 0x29, 0x84,
	0x99, 0x54, 0x51, 0x30, 0x12, 0xc9, 0x50, 0x2a, 0x91, 0x06, 0xb1, 0xdc, 0x17, 0x36, 0x00, 0x11,
	0x1b, 0xc5, 0x41, 0x04, 0x29, 0x8f, 0x8d, 0x12, 0x81, 0x15, 0x46, 0x5b, 0x08, 0x46, 0xba, 0x1f,
	0x89, 0x21, 0xcf, 0x14, 0xf4, 0x95, 0x0c, 0x2d, 0xb7, 0x93, 0x3e, 0xc4, 0x46, 0x31, 0x63, 0x35,
	0x68, 0x72, 0x41, 0xa6, 0x20, 0x35, 0x73, 0x04, 0xc6, 0x23, 0x6e, 0x40, 0x58, 0x56, 0x02, 0x58,
	0x09, 0xf0, 0xdb, 0x25, 0x7e, 0xbc, 0x53, 0xdd, 0x20, 0xf6, 0x41, 0x24, 0xa9, 0xd4, 0x49, 0x5a,
	0x62, 0xfc, 0xad, 0x99, 0x66, 0xa0, 0x93, 0xa1, 0x1c, 0x05, 0x91, 0x48, 0x07, 0x56, 0x1a, 0xd0,
	0x36, 0x18, 0x73, 0x95, 0x89, 0x3e, 0x4c, 0x8c, 0x28, 0xc5, 0xed, 0xb7, 0x4d, 0xbc, 0xf8, 0x78,
	0x62, 0x04, 0xb9, 0x8b, 0x97, 0xdc, 0xa1, 0x87, 0x36, 0x51, 0x67, 0xad, 0xbb, 0xc5, 0xea, 0x66,
	0xc6, 0x3b, 0xac, 0x64, 0xb1, 0x8a, 0xc5, 0x9e, 0x14, 0xf2, 0xa2, 0xb7, 0x57, 0x76, 0x92, 0xa7,
	0x18, 0x47, 0x32, 0x3e, 0x32, 0xe3, 0x35, 0x37, 0x17, 0x3a, 0x67, 0xbb, 0xd7, 0xd8, 0x3f, 0x43,
	0xb1, 0x82, 0xc1, 0xee, 0xcf, 0x3a, 0x1f, 0x24, 0x60, 0x27, 0xbd, 0x1a, 0xca, 0x7f, 0x89, 0xd7,
	0xe7, 0x8e, 0xc9, 0x39, 0xbc, 0xf0, 0x4a, 0x4c, 0x9c, 0xd9, 0x56, 0xaf, 0x78, 0xac, 0x02, 0x34,
	0xff, 0x37, 0xc0, 0x8d, 0xe6, 0x75, 0xd4, 0xfe, 0xb8, 0x88, 0x57, 0x77, 0x93, 0x14, 0x78, 0x32,
	0x10, 0x7b, 0xdc, 0xf2, 0x98, 0x6c, 0xd4, 0x27, 0xd3, 0xfa, 0x1d, 0xf6, 0xc5, 0x1f, 0xc2, 0xde,
	0x39, 0x41, 0xd8, 0x63, 0xec, 0xbf, 0xa5, 0x26, 0x97, 0xf0, 0x9a, 0x4c, 0xe0, 0xea, 0x95, 0x3d,
	0x2b, 0x63, 0x09, 0x72, 0x2c, 0xbc, 0x05, 0x67, 0x60, 0xae, 0x4a, 0x2e, 0xe2, 0xd5, 0x50, 0x6b,
	0x55, 0xc9, 0x16, 0x9d, 0xec, 0x78, 0x91, 0x74, 0xf0, 0x7a, 0xa4, 0xb3, 0x50, 0x89, 0x4a, 0xb7,
	0xe4, 0x74, 0xf3, 0xe5, 0x42, 0x99, 0x82, 0x95, 0xc9, 0xa8, 0x52, 0x2e, 0x97, 0xca, 0xb9, 0x32,
	0x79, 0x86, 0x57, 0x9c, 0x97, 0x87, 0xdc, 0x78, 0x67, 0xdc, 0x04, 0x6e, 0x9f, 0x7a, 0x02, 0xbb,
	0x47, 0x80, 0x32, 0xff, 0x8c, 0x47, 0xce, 0xe3, 0x16, 0xc8, 0x58, 0x3c, 0x02, 0x1e, 0x1b, 0xaf,
	0xe5, 0xee, 0xaf, 0x0a, 0xc4, 0xc7, 0x2b, 0x51, 0x66, 0x39, 0x48, 0x9d, 0x78, 0xd8, 0x1d, 0xce,
	0xde, 0xfd, 0x5b, 0x27, 0xd9, 0x96, 0x8d, 0xfa, 0xb6, 0xb4, 0x6a, 0x0b, 0xe0, 0xdf, 0x2c, 0xbe,
	0x7f, 0xcd, 0xd3, 0x69, 0x9a, 0xef, 0x75, 0x0f, 0xa6, 0xb4, 0xf1, 0x65, 0x4a, 0x1b, 0x87, 0x53,
	0x8a, 0xde, 0xe4, 0x14, 0x7d, 0xc8, 0x29, 0xfa, 0x94, 0x53, 0x74, 0x90, 0x53, 0xf4, 0x35, 0xa7,
	0xe8, 0x47, 0x4e, 0x1b, 0x87, 0x39, 0x45, 0xef, 0xbe, 0xd1, 0xc6, 0xcf, 0xcf, 0xdf, 0xdf, 0x37,
	0x51, 0xb8, 0xec, 0xfe, 0xc4, 0xcb, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x28, 0x29, 0x81,
	0x5e, 0x04, 0x00, 0x00,
}
