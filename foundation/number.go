package foundation

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
	Compare(otherNumber Number) ComparisonResult
	IsEqualToNumber(number Number) bool
	BoolValue() bool
	DoubleValue() float64
	FloatValue() float32
	IntegerValue() int
	UnsignedIntegerValue() uint
	StringValue() string
}

type NSNumber struct {
	NSValue
}

func MakeNumber(ptr unsafe.Pointer) NSNumber {
	return NSNumber{
		NSValue: MakeValue(ptr),
	}
}

func (n NSNumber) InitWithCoder(coder Coder) NSNumber {
	result_ := C.C_NSNumber_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeNumber(result_)
}

func AllocNumber() NSNumber {
	result_ := C.C_NSNumber_AllocNumber()
	return MakeNumber(result_)
}

func (n NSNumber) Init() NSNumber {
	result_ := C.C_NSNumber_Init(n.Ptr())
	return MakeNumber(result_)
}

func NewNumber() NSNumber {
	result_ := C.C_NSNumber_NewNumber()
	return MakeNumber(result_)
}

func (n NSNumber) Autorelease() NSNumber {
	result_ := C.C_NSNumber_Autorelease(n.Ptr())
	return MakeNumber(result_)
}

func (n NSNumber) Retain() NSNumber {
	result_ := C.C_NSNumber_Retain(n.Ptr())
	return MakeNumber(result_)
}

func NumberWithBool(value bool) Number {
	result_ := C.C_NSNumber_NumberWithBool(C.bool(value))
	return MakeNumber(result_)
}

func NumberWithDouble(value float64) Number {
	result_ := C.C_NSNumber_NumberWithDouble(C.double(value))
	return MakeNumber(result_)
}

func NumberWithFloat(value float32) Number {
	result_ := C.C_NSNumber_NumberWithFloat(C.float(value))
	return MakeNumber(result_)
}

func NumberWithInteger(value int) Number {
	result_ := C.C_NSNumber_NumberWithInteger(C.int(value))
	return MakeNumber(result_)
}

func NumberWithUnsignedInteger(value uint) Number {
	result_ := C.C_NSNumber_NumberWithUnsignedInteger(C.uint(value))
	return MakeNumber(result_)
}

func (n NSNumber) InitWithBool(value bool) Number {
	result_ := C.C_NSNumber_InitWithBool(n.Ptr(), C.bool(value))
	return MakeNumber(result_)
}

func (n NSNumber) InitWithDouble(value float64) Number {
	result_ := C.C_NSNumber_InitWithDouble(n.Ptr(), C.double(value))
	return MakeNumber(result_)
}

func (n NSNumber) InitWithFloat(value float32) Number {
	result_ := C.C_NSNumber_InitWithFloat(n.Ptr(), C.float(value))
	return MakeNumber(result_)
}

func (n NSNumber) InitWithInteger(value int) Number {
	result_ := C.C_NSNumber_InitWithInteger(n.Ptr(), C.int(value))
	return MakeNumber(result_)
}

func (n NSNumber) InitWithUnsignedInteger(value uint) Number {
	result_ := C.C_NSNumber_InitWithUnsignedInteger(n.Ptr(), C.uint(value))
	return MakeNumber(result_)
}

func (n NSNumber) DescriptionWithLocale(locale objc.Object) string {
	result_ := C.C_NSNumber_DescriptionWithLocale(n.Ptr(), objc.ExtractPtr(locale))
	return MakeString(result_).String()
}

func (n NSNumber) Compare(otherNumber Number) ComparisonResult {
	result_ := C.C_NSNumber_Compare(n.Ptr(), objc.ExtractPtr(otherNumber))
	return ComparisonResult(int(result_))
}

func (n NSNumber) IsEqualToNumber(number Number) bool {
	result_ := C.C_NSNumber_IsEqualToNumber(n.Ptr(), objc.ExtractPtr(number))
	return bool(result_)
}

func (n NSNumber) BoolValue() bool {
	result_ := C.C_NSNumber_BoolValue(n.Ptr())
	return bool(result_)
}

func (n NSNumber) DoubleValue() float64 {
	result_ := C.C_NSNumber_DoubleValue(n.Ptr())
	return float64(result_)
}

func (n NSNumber) FloatValue() float32 {
	result_ := C.C_NSNumber_FloatValue(n.Ptr())
	return float32(result_)
}

func (n NSNumber) IntegerValue() int {
	result_ := C.C_NSNumber_IntegerValue(n.Ptr())
	return int(result_)
}

func (n NSNumber) UnsignedIntegerValue() uint {
	result_ := C.C_NSNumber_UnsignedIntegerValue(n.Ptr())
	return uint(result_)
}

func (n NSNumber) StringValue() string {
	result_ := C.C_NSNumber_StringValue(n.Ptr())
	return MakeString(result_).String()
}
