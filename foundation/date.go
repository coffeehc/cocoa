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
	TimeIntervalSinceNow() float64
	TimeIntervalSinceReferenceDate() float64
	TimeIntervalSince1970() float64
	Description() string
	DescriptionWithLocale(locale objc.Object) string
	IsEqualToDate(otherDate Date) bool
	EarlierDate(otherDate Date) Date
	LaterDate(otherDate Date) Date
	Compare(other Date) ComparisonResult
	TimeIntervalSinceDate(other Date) float64
}

var _ Date = (*NSDate)(nil)

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

func DistantFuture() Date {
	return MakeDate(C.Date_DistantFuture())
}

func DistantPast() Date {
	return MakeDate(C.Date_DistantPast())
}

func (d *NSDate) TimeIntervalSinceNow() float64 {
	return float64(C.Date_TimeIntervalSinceNow(d.Ptr()))
}

func (d *NSDate) TimeIntervalSinceReferenceDate() float64 {
	return float64(C.Date_TimeIntervalSinceReferenceDate(d.Ptr()))
}

func (d *NSDate) TimeIntervalSince1970() float64 {
	return float64(C.Date_TimeIntervalSince1970(d.Ptr()))
}

func (d *NSDate) Description() string {
	return C.GoString(C.Date_Description(d.Ptr()))
}

func NewDate() Date {
	return MakeDate(C.Date_NewDate())
}

func CurrentDate() Date {
	return MakeDate(C.Date_CurrentDate())
}

func DateWithTimeIntervalSinceNow(secs float64) Date {
	return MakeDate(C.Date_DateWithTimeIntervalSinceNow(C.double(secs)))
}

func dateWithTimeIntervalSinceDate(secsToBeAdded float64, date Date) Date {
	return MakeDate(C.Date_dateWithTimeIntervalSinceDate(C.double(secsToBeAdded), toPointer(date)))
}

func DateWithTimeIntervalSinceReferenceDate(secs float64) Date {
	return MakeDate(C.Date_DateWithTimeIntervalSinceReferenceDate(C.double(secs)))
}

func DateWithTimeIntervalSince1970(secs float64) Date {
	return MakeDate(C.Date_DateWithTimeIntervalSince1970(C.double(secs)))
}

func (d *NSDate) DescriptionWithLocale(locale objc.Object) string {
	return C.GoString(C.Date_DescriptionWithLocale(d.Ptr(), toPointer(locale)))
}

func (d *NSDate) IsEqualToDate(otherDate Date) bool {
	return bool(C.Date_IsEqualToDate(d.Ptr(), toPointer(otherDate)))
}

func (d *NSDate) EarlierDate(otherDate Date) Date {
	return MakeDate(C.Date_EarlierDate(d.Ptr(), toPointer(otherDate)))
}

func (d *NSDate) LaterDate(otherDate Date) Date {
	return MakeDate(C.Date_LaterDate(d.Ptr(), toPointer(otherDate)))
}

func (d *NSDate) Compare(other Date) ComparisonResult {
	return ComparisonResult(C.Date_Compare(d.Ptr(), toPointer(other)))
}

func (d *NSDate) TimeIntervalSinceDate(other Date) float64 {
	return float64(C.Date_TimeIntervalSinceDate(d.Ptr(), toPointer(other)))
}
