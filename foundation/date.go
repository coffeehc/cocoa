package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "date.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Date interface {
	objc.Object
	IsEqualToDate(otherDate Date) bool
	EarlierDate(anotherDate Date) Date
	LaterDate(anotherDate Date) Date
	TimeIntervalSinceDate(anotherDate Date) TimeInterval
	DescriptionWithLocale(locale objc.Object) string
	TimeIntervalSinceNow() TimeInterval
	TimeIntervalSinceReferenceDate() TimeInterval
	TimeIntervalSince1970() TimeInterval
	Description() string
}

type NSDate struct {
	objc.NSObject
}

func MakeDate(ptr unsafe.Pointer) *NSDate {
	if ptr == nil {
		return nil
	}
	return &NSDate{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocDate() *NSDate {
	return MakeDate(C.C_Date_Alloc())
}

func (n *NSDate) Init() Date {
	result := C.C_NSDate_Init(n.Ptr())
	return MakeDate(result)
}

func (n *NSDate) InitWithTimeIntervalSinceNow(secs TimeInterval) Date {
	result := C.C_NSDate_InitWithTimeIntervalSinceNow(n.Ptr(), C.double(float64(secs)))
	return MakeDate(result)
}

func (n *NSDate) InitWithTimeInterval_SinceDate(secsToBeAdded TimeInterval, date Date) Date {
	result := C.C_NSDate_InitWithTimeInterval_SinceDate(n.Ptr(), C.double(float64(secsToBeAdded)), objc.ExtractPtr(date))
	return MakeDate(result)
}

func (n *NSDate) InitWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	result := C.C_NSDate_InitWithTimeIntervalSinceReferenceDate(n.Ptr(), C.double(float64(ti)))
	return MakeDate(result)
}

func (n *NSDate) InitWithTimeIntervalSince1970(secs TimeInterval) Date {
	result := C.C_NSDate_InitWithTimeIntervalSince1970(n.Ptr(), C.double(float64(secs)))
	return MakeDate(result)
}

func (n *NSDate) InitWithCoder(coder Coder) Date {
	result := C.C_NSDate_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeDate(result)
}

func (n *NSDate) DateByAddingTimeInterval(ti TimeInterval) Date {
	result := C.C_NSDate_DateByAddingTimeInterval(n.Ptr(), C.double(float64(ti)))
	return MakeDate(result)
}

func CurrentDate() Date {
	result := C.C_NSDate_CurrentDate()
	return MakeDate(result)
}

func DateWithTimeIntervalSinceNow(secs TimeInterval) Date {
	result := C.C_NSDate_DateWithTimeIntervalSinceNow(C.double(float64(secs)))
	return MakeDate(result)
}

func DateWithTimeInterval_SinceDate(secsToBeAdded TimeInterval, date Date) Date {
	result := C.C_NSDate_DateWithTimeInterval_SinceDate(C.double(float64(secsToBeAdded)), objc.ExtractPtr(date))
	return MakeDate(result)
}

func DateWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	result := C.C_NSDate_DateWithTimeIntervalSinceReferenceDate(C.double(float64(ti)))
	return MakeDate(result)
}

func DateWithTimeIntervalSince1970(secs TimeInterval) Date {
	result := C.C_NSDate_DateWithTimeIntervalSince1970(C.double(float64(secs)))
	return MakeDate(result)
}

func (n *NSDate) IsEqualToDate(otherDate Date) bool {
	result := C.C_NSDate_IsEqualToDate(n.Ptr(), objc.ExtractPtr(otherDate))
	return bool(result)
}

func (n *NSDate) EarlierDate(anotherDate Date) Date {
	result := C.C_NSDate_EarlierDate(n.Ptr(), objc.ExtractPtr(anotherDate))
	return MakeDate(result)
}

func (n *NSDate) LaterDate(anotherDate Date) Date {
	result := C.C_NSDate_LaterDate(n.Ptr(), objc.ExtractPtr(anotherDate))
	return MakeDate(result)
}

func (n *NSDate) TimeIntervalSinceDate(anotherDate Date) TimeInterval {
	result := C.C_NSDate_TimeIntervalSinceDate(n.Ptr(), objc.ExtractPtr(anotherDate))
	return TimeInterval(float64(result))
}

func (n *NSDate) DescriptionWithLocale(locale objc.Object) string {
	result := C.C_NSDate_DescriptionWithLocale(n.Ptr(), objc.ExtractPtr(locale))
	return MakeString(result).String()
}

func (n *NSDate) TimeIntervalSinceNow() TimeInterval {
	result := C.C_NSDate_TimeIntervalSinceNow(n.Ptr())
	return TimeInterval(float64(result))
}

func (n *NSDate) TimeIntervalSinceReferenceDate() TimeInterval {
	result := C.C_NSDate_TimeIntervalSinceReferenceDate(n.Ptr())
	return TimeInterval(float64(result))
}

func (n *NSDate) TimeIntervalSince1970() TimeInterval {
	result := C.C_NSDate_TimeIntervalSince1970(n.Ptr())
	return TimeInterval(float64(result))
}

func (n *NSDate) Description() string {
	result := C.C_NSDate_Description(n.Ptr())
	return MakeString(result).String()
}
