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
	objc.Object
	BoolValue() bool
	FloatValue() float32
	IntValue() int
	LongLongValue() int64
	UnsignedIntValue() uint
	UnsignedLongLongValue() uint64
}

var _ Number = (*NSNumber)(nil)

type NSNumber struct {
	objc.NSObject
}

func MakeNumber(ptr unsafe.Pointer) *NSNumber {
	if ptr == nil {
		return nil
	}
	return &NSNumber{
		NSObject: *objc.MakeObject(ptr),
	}
}

func (n *NSNumber) BoolValue() bool {
	return bool(C.Number_BoolValue(n.Ptr()))
}

func (n *NSNumber) FloatValue() float32 {
	return float32(C.Number_FloatValue(n.Ptr()))
}

func (n *NSNumber) IntValue() int {
	return int(C.Number_IntValue(n.Ptr()))
}

func (n *NSNumber) LongLongValue() int64 {
	return int64(C.Number_LongLongValue(n.Ptr()))
}

func (n *NSNumber) UnsignedIntValue() uint {
	return uint(C.Number_UnsignedIntValue(n.Ptr()))
}

func (n *NSNumber) UnsignedLongLongValue() uint64 {
	return uint64(C.Number_UnsignedLongLongValue(n.Ptr()))
}

func NumberWithBool(value bool) Number {
	return MakeNumber(C.Number_NumberWithBool(C.bool(value)))
}

func NumberWithDouble(value float64) Number {
	return MakeNumber(C.Number_NumberWithDouble(C.double(value)))
}

func NumberWithFloat(value float32) Number {
	return MakeNumber(C.Number_NumberWithFloat(C.float(value)))
}

func NumberWithInt(value int) Number {
	return MakeNumber(C.Number_NumberWithInt(C.long(value)))
}

func NumberWithLongLong(value int64) Number {
	return MakeNumber(C.Number_NumberWithLongLong(C.long(value)))
}

func NumberWithUnsignedInt(value uint) Number {
	return MakeNumber(C.Number_NumberWithUnsignedInt(C.ulong(value)))
}

func NumberWithUnsignedLongLong(value uint64) Number {
	return MakeNumber(C.Number_NumberWithUnsignedLongLong(C.ulong(value)))
}
