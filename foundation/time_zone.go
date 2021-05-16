package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "time_zone.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TimeZone interface {
	objc.Object
	AbbreviationForDate(aDate Date) string
	SecondsFromGMTForDate(aDate Date) int
	IsDaylightSavingTimeForDate(aDate Date) bool
	DaylightSavingTimeOffsetForDate(aDate Date) TimeInterval
	NextDaylightSavingTimeTransitionAfterDate(aDate Date) Date
	IsEqualToTimeZone(aTimeZone TimeZone) bool
	LocalizedName_Locale(style TimeZoneNameStyle, locale Locale) string
	Name() string
	Abbreviation() string
	SecondsFromGMT() int
	Data() []byte
	IsDaylightSavingTime() bool
	DaylightSavingTimeOffset() TimeInterval
	NextDaylightSavingTimeTransition() Date
	Description() string
}

type NSTimeZone struct {
	objc.NSObject
}

func MakeTimeZone(ptr unsafe.Pointer) *NSTimeZone {
	if ptr == nil {
		return nil
	}
	return &NSTimeZone{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocTimeZone() *NSTimeZone {
	return MakeTimeZone(C.C_TimeZone_Alloc())
}

func (n *NSTimeZone) InitWithName(tzName string) TimeZone {
	result := C.C_NSTimeZone_InitWithName(n.Ptr(), NewString(tzName).Ptr())
	return MakeTimeZone(result)
}

func (n *NSTimeZone) InitWithName_Data(tzName string, aData []byte) TimeZone {
	result := C.C_NSTimeZone_InitWithName_Data(n.Ptr(), NewString(tzName).Ptr(), C.Array{data: unsafe.Pointer(&aData[0]), len: C.int(len(aData))})
	return MakeTimeZone(result)
}

func ResetSystemTimeZone() {
	C.C_NSTimeZone_ResetSystemTimeZone()
}

func TimeZoneWithName(tzName string) TimeZone {
	result := C.C_NSTimeZone_TimeZoneWithName(NewString(tzName).Ptr())
	return MakeTimeZone(result)
}

func TimeZoneWithName_Data(tzName string, aData []byte) TimeZone {
	result := C.C_NSTimeZone_TimeZoneWithName_Data(NewString(tzName).Ptr(), C.Array{data: unsafe.Pointer(&aData[0]), len: C.int(len(aData))})
	return MakeTimeZone(result)
}

func TimeZoneWithAbbreviation(abbreviation string) TimeZone {
	result := C.C_NSTimeZone_TimeZoneWithAbbreviation(NewString(abbreviation).Ptr())
	return MakeTimeZone(result)
}

func TimeZoneForSecondsFromGMT(seconds int) TimeZone {
	result := C.C_NSTimeZone_TimeZoneForSecondsFromGMT(C.int(seconds))
	return MakeTimeZone(result)
}

func (n *NSTimeZone) AbbreviationForDate(aDate Date) string {
	result := C.C_NSTimeZone_AbbreviationForDate(n.Ptr(), objc.ExtractPtr(aDate))
	return MakeString(result).String()
}

func (n *NSTimeZone) SecondsFromGMTForDate(aDate Date) int {
	result := C.C_NSTimeZone_SecondsFromGMTForDate(n.Ptr(), objc.ExtractPtr(aDate))
	return int(result)
}

func (n *NSTimeZone) IsDaylightSavingTimeForDate(aDate Date) bool {
	result := C.C_NSTimeZone_IsDaylightSavingTimeForDate(n.Ptr(), objc.ExtractPtr(aDate))
	return bool(result)
}

func (n *NSTimeZone) DaylightSavingTimeOffsetForDate(aDate Date) TimeInterval {
	result := C.C_NSTimeZone_DaylightSavingTimeOffsetForDate(n.Ptr(), objc.ExtractPtr(aDate))
	return TimeInterval(float64(result))
}

func (n *NSTimeZone) NextDaylightSavingTimeTransitionAfterDate(aDate Date) Date {
	result := C.C_NSTimeZone_NextDaylightSavingTimeTransitionAfterDate(n.Ptr(), objc.ExtractPtr(aDate))
	return MakeDate(result)
}

func (n *NSTimeZone) IsEqualToTimeZone(aTimeZone TimeZone) bool {
	result := C.C_NSTimeZone_IsEqualToTimeZone(n.Ptr(), objc.ExtractPtr(aTimeZone))
	return bool(result)
}

func (n *NSTimeZone) LocalizedName_Locale(style TimeZoneNameStyle, locale Locale) string {
	result := C.C_NSTimeZone_LocalizedName_Locale(n.Ptr(), C.int(int(style)), objc.ExtractPtr(locale))
	return MakeString(result).String()
}

func LocalTimeZone() TimeZone {
	result := C.C_NSTimeZone_LocalTimeZone()
	return MakeTimeZone(result)
}

func SystemTimeZone() TimeZone {
	result := C.C_NSTimeZone_SystemTimeZone()
	return MakeTimeZone(result)
}

func DefaultTimeZone() TimeZone {
	result := C.C_NSTimeZone_DefaultTimeZone()
	return MakeTimeZone(result)
}

func SetDefaultTimeZone(value TimeZone) {
	C.C_NSTimeZone_SetDefaultTimeZone(objc.ExtractPtr(value))
}

func TimeZone_KnownTimeZoneNames() []string {
	result := C.C_NSTimeZone_TimeZone_KnownTimeZoneNames()
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]string, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeString(r).String()
	}
	return goResult
}

func (n *NSTimeZone) Name() string {
	result := C.C_NSTimeZone_Name(n.Ptr())
	return MakeString(result).String()
}

func (n *NSTimeZone) Abbreviation() string {
	result := C.C_NSTimeZone_Abbreviation(n.Ptr())
	return MakeString(result).String()
}

func (n *NSTimeZone) SecondsFromGMT() int {
	result := C.C_NSTimeZone_SecondsFromGMT(n.Ptr())
	return int(result)
}

func (n *NSTimeZone) Data() []byte {
	result := C.C_NSTimeZone_Data(n.Ptr())
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func TimeZoneDataVersion() string {
	result := C.C_NSTimeZone_TimeZoneDataVersion()
	return MakeString(result).String()
}

func (n *NSTimeZone) IsDaylightSavingTime() bool {
	result := C.C_NSTimeZone_IsDaylightSavingTime(n.Ptr())
	return bool(result)
}

func (n *NSTimeZone) DaylightSavingTimeOffset() TimeInterval {
	result := C.C_NSTimeZone_DaylightSavingTimeOffset(n.Ptr())
	return TimeInterval(float64(result))
}

func (n *NSTimeZone) NextDaylightSavingTimeTransition() Date {
	result := C.C_NSTimeZone_NextDaylightSavingTimeTransition(n.Ptr())
	return MakeDate(result)
}

func (n *NSTimeZone) Description() string {
	result := C.C_NSTimeZone_Description(n.Ptr())
	return MakeString(result).String()
}
