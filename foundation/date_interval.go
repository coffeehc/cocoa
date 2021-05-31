package foundation

// #include "date_interval.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type DateInterval interface {
	objc.Object
	Compare(dateInterval DateInterval) ComparisonResult
	IsEqualToDateInterval(dateInterval DateInterval) bool
	IntersectsDateInterval(dateInterval DateInterval) bool
	IntersectionWithDateInterval(dateInterval DateInterval) DateInterval
	ContainsDate(date Date) bool
	StartDate() Date
	EndDate() Date
	Duration() TimeInterval
}

type NSDateInterval struct {
	objc.NSObject
}

func MakeDateInterval(ptr unsafe.Pointer) *NSDateInterval {
	if ptr == nil {
		return nil
	}
	return &NSDateInterval{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocDateInterval() *NSDateInterval {
	return MakeDateInterval(C.C_DateInterval_Alloc())
}

func (n *NSDateInterval) Init() DateInterval {
	result_ := C.C_NSDateInterval_Init(n.Ptr())
	return MakeDateInterval(result_)
}

func (n *NSDateInterval) InitWithStartDate_Duration(startDate Date, duration TimeInterval) DateInterval {
	result_ := C.C_NSDateInterval_InitWithStartDate_Duration(n.Ptr(), objc.ExtractPtr(startDate), C.double(float64(duration)))
	return MakeDateInterval(result_)
}

func (n *NSDateInterval) InitWithStartDate_EndDate(startDate Date, endDate Date) DateInterval {
	result_ := C.C_NSDateInterval_InitWithStartDate_EndDate(n.Ptr(), objc.ExtractPtr(startDate), objc.ExtractPtr(endDate))
	return MakeDateInterval(result_)
}

func (n *NSDateInterval) InitWithCoder(coder Coder) DateInterval {
	result_ := C.C_NSDateInterval_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeDateInterval(result_)
}

func (n *NSDateInterval) Compare(dateInterval DateInterval) ComparisonResult {
	result_ := C.C_NSDateInterval_Compare(n.Ptr(), objc.ExtractPtr(dateInterval))
	return ComparisonResult(int(result_))
}

func (n *NSDateInterval) IsEqualToDateInterval(dateInterval DateInterval) bool {
	result_ := C.C_NSDateInterval_IsEqualToDateInterval(n.Ptr(), objc.ExtractPtr(dateInterval))
	return bool(result_)
}

func (n *NSDateInterval) IntersectsDateInterval(dateInterval DateInterval) bool {
	result_ := C.C_NSDateInterval_IntersectsDateInterval(n.Ptr(), objc.ExtractPtr(dateInterval))
	return bool(result_)
}

func (n *NSDateInterval) IntersectionWithDateInterval(dateInterval DateInterval) DateInterval {
	result_ := C.C_NSDateInterval_IntersectionWithDateInterval(n.Ptr(), objc.ExtractPtr(dateInterval))
	return MakeDateInterval(result_)
}

func (n *NSDateInterval) ContainsDate(date Date) bool {
	result_ := C.C_NSDateInterval_ContainsDate(n.Ptr(), objc.ExtractPtr(date))
	return bool(result_)
}

func (n *NSDateInterval) StartDate() Date {
	result_ := C.C_NSDateInterval_StartDate(n.Ptr())
	return MakeDate(result_)
}

func (n *NSDateInterval) EndDate() Date {
	result_ := C.C_NSDateInterval_EndDate(n.Ptr())
	return MakeDate(result_)
}

func (n *NSDateInterval) Duration() TimeInterval {
	result_ := C.C_NSDateInterval_Duration(n.Ptr())
	return TimeInterval(float64(result_))
}
