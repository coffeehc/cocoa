package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "number.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Number interface {
	Value
	InitWithBool(value bool) Number
	InitWithDouble(value float64) Number
	InitWithFloat(value float32) Number
	InitWithInteger(value int) Number
	InitWithUnsignedInteger(value uint) Number
	DescriptionWithLocale(locale objc.Object) string
	IsEqualToNumber(number Number) bool
	BoolValue() bool
	DecimalValue() Decimal
	DoubleValue() float64
	FloatValue() float32
	IntegerValue() int
	UnsignedIntegerValue() uint
	StringValue() string
}

type NSNumber struct {
	NSValue
}

func MakeNumber(ptr unsafe.Pointer) *NSNumber {
	if ptr == nil {
		return nil
	}
	return &NSNumber{
		NSValue: *MakeValue(ptr),
	}
}

func AllocNumber() *NSNumber {
	return MakeNumber(C.C_Number_Alloc())
}

func (n *NSNumber) InitWithCoder(coder Coder) Number {
	result := C.C_NSNumber_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeNumber(result)
}

func (n *NSNumber) Init() Number {
	result := C.C_NSNumber_Init(n.Ptr())
	return MakeNumber(result)
}

func NumberWithBool(value bool) Number {
	result := C.C_NSNumber_NumberWithBool(C.bool(value))
	return MakeNumber(result)
}

func NumberWithDouble(value float64) Number {
	result := C.C_NSNumber_NumberWithDouble(C.double(value))
	return MakeNumber(result)
}

func NumberWithFloat(value float32) Number {
	result := C.C_NSNumber_NumberWithFloat(C.float(value))
	return MakeNumber(result)
}

func NumberWithInteger(value int) Number {
	result := C.C_NSNumber_NumberWithInteger(C.int(value))
	return MakeNumber(result)
}

func NumberWithUnsignedInteger(value uint) Number {
	result := C.C_NSNumber_NumberWithUnsignedInteger(C.uint(value))
	return MakeNumber(result)
}

func (n *NSNumber) InitWithBool(value bool) Number {
	result := C.C_NSNumber_InitWithBool(n.Ptr(), C.bool(value))
	return MakeNumber(result)
}

func (n *NSNumber) InitWithDouble(value float64) Number {
	result := C.C_NSNumber_InitWithDouble(n.Ptr(), C.double(value))
	return MakeNumber(result)
}

func (n *NSNumber) InitWithFloat(value float32) Number {
	result := C.C_NSNumber_InitWithFloat(n.Ptr(), C.float(value))
	return MakeNumber(result)
}

func (n *NSNumber) InitWithInteger(value int) Number {
	result := C.C_NSNumber_InitWithInteger(n.Ptr(), C.int(value))
	return MakeNumber(result)
}

func (n *NSNumber) InitWithUnsignedInteger(value uint) Number {
	result := C.C_NSNumber_InitWithUnsignedInteger(n.Ptr(), C.uint(value))
	return MakeNumber(result)
}

func (n *NSNumber) DescriptionWithLocale(locale objc.Object) string {
	result := C.C_NSNumber_DescriptionWithLocale(n.Ptr(), objc.ExtractPtr(locale))
	return MakeString(result).String()
}

func (n *NSNumber) IsEqualToNumber(number Number) bool {
	result := C.C_NSNumber_IsEqualToNumber(n.Ptr(), objc.ExtractPtr(number))
	return bool(result)
}

func (n *NSNumber) BoolValue() bool {
	result := C.C_NSNumber_BoolValue(n.Ptr())
	return bool(result)
}

func (n *NSNumber) DecimalValue() Decimal {
	result := C.C_NSNumber_DecimalValue(n.Ptr())
	return FromNSDecimalPointer(unsafe.Pointer(&result))
}

func (n *NSNumber) DoubleValue() float64 {
	result := C.C_NSNumber_DoubleValue(n.Ptr())
	return float64(result)
}

func (n *NSNumber) FloatValue() float32 {
	result := C.C_NSNumber_FloatValue(n.Ptr())
	return float32(result)
}

func (n *NSNumber) IntegerValue() int {
	result := C.C_NSNumber_IntegerValue(n.Ptr())
	return int(result)
}

func (n *NSNumber) UnsignedIntegerValue() uint {
	result := C.C_NSNumber_UnsignedIntegerValue(n.Ptr())
	return uint(result)
}

func (n *NSNumber) StringValue() string {
	result := C.C_NSNumber_StringValue(n.Ptr())
	return MakeString(result).String()
}
