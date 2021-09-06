package foundation

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

func MakeTimeZone(ptr unsafe.Pointer) NSTimeZone {
	return NSTimeZone{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocTimeZone() NSTimeZone {
	return MakeTimeZone(C.C_TimeZone_Alloc())
}

func (n NSTimeZone) InitWithName(tzName string) TimeZone {
	result_ := C.C_NSTimeZone_InitWithName(n.Ptr(), NewString(tzName).Ptr())
	return MakeTimeZone(result_)
}

func (n NSTimeZone) InitWithName_Data(tzName string, aData []byte) TimeZone {
	result_ := C.C_NSTimeZone_InitWithName_Data(n.Ptr(), NewString(tzName).Ptr(), NewData(aData).Ptr())
	return MakeTimeZone(result_)
}

func ResetSystemTimeZone() {
	C.C_NSTimeZone_ResetSystemTimeZone()
}

func TimeZoneWithName(tzName string) TimeZone {
	result_ := C.C_NSTimeZone_TimeZoneWithName(NewString(tzName).Ptr())
	return MakeTimeZone(result_)
}

func TimeZoneWithName_Data(tzName string, aData []byte) TimeZone {
	result_ := C.C_NSTimeZone_TimeZoneWithName_Data(NewString(tzName).Ptr(), NewData(aData).Ptr())
	return MakeTimeZone(result_)
}

func TimeZoneWithAbbreviation(abbreviation string) TimeZone {
	result_ := C.C_NSTimeZone_TimeZoneWithAbbreviation(NewString(abbreviation).Ptr())
	return MakeTimeZone(result_)
}

func TimeZoneForSecondsFromGMT(seconds int) TimeZone {
	result_ := C.C_NSTimeZone_TimeZoneForSecondsFromGMT(C.int(seconds))
	return MakeTimeZone(result_)
}

func (n NSTimeZone) AbbreviationForDate(aDate Date) string {
	result_ := C.C_NSTimeZone_AbbreviationForDate(n.Ptr(), objc.ExtractPtr(aDate))
	return MakeString(result_).String()
}

func (n NSTimeZone) SecondsFromGMTForDate(aDate Date) int {
	result_ := C.C_NSTimeZone_SecondsFromGMTForDate(n.Ptr(), objc.ExtractPtr(aDate))
	return int(result_)
}

func (n NSTimeZone) IsDaylightSavingTimeForDate(aDate Date) bool {
	result_ := C.C_NSTimeZone_IsDaylightSavingTimeForDate(n.Ptr(), objc.ExtractPtr(aDate))
	return bool(result_)
}

func (n NSTimeZone) DaylightSavingTimeOffsetForDate(aDate Date) TimeInterval {
	result_ := C.C_NSTimeZone_DaylightSavingTimeOffsetForDate(n.Ptr(), objc.ExtractPtr(aDate))
	return TimeInterval(float64(result_))
}

func (n NSTimeZone) NextDaylightSavingTimeTransitionAfterDate(aDate Date) Date {
	result_ := C.C_NSTimeZone_NextDaylightSavingTimeTransitionAfterDate(n.Ptr(), objc.ExtractPtr(aDate))
	return MakeDate(result_)
}

func (n NSTimeZone) IsEqualToTimeZone(aTimeZone TimeZone) bool {
	result_ := C.C_NSTimeZone_IsEqualToTimeZone(n.Ptr(), objc.ExtractPtr(aTimeZone))
	return bool(result_)
}

func (n NSTimeZone) LocalizedName_Locale(style TimeZoneNameStyle, locale Locale) string {
	result_ := C.C_NSTimeZone_LocalizedName_Locale(n.Ptr(), C.int(int(style)), objc.ExtractPtr(locale))
	return MakeString(result_).String()
}

func LocalTimeZone() TimeZone {
	result_ := C.C_NSTimeZone_LocalTimeZone()
	return MakeTimeZone(result_)
}

func SystemTimeZone() TimeZone {
	result_ := C.C_NSTimeZone_SystemTimeZone()
	return MakeTimeZone(result_)
}

func DefaultTimeZone() TimeZone {
	result_ := C.C_NSTimeZone_DefaultTimeZone()
	return MakeTimeZone(result_)
}

func SetDefaultTimeZone(value TimeZone) {
	C.C_NSTimeZone_SetDefaultTimeZone(objc.ExtractPtr(value))
}

func KnownTimeZoneNames() []string {
	result_ := C.C_NSTimeZone_KnownTimeZoneNames()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func TimeZone_AbbreviationDictionary() map[string]string {
	result_ := C.C_NSTimeZone_TimeZone_AbbreviationDictionary()
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[string]string)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[MakeString(k).String()] = MakeString(v).String()
	}
	return goResult_
}

func TimeZone_SetAbbreviationDictionary(value map[string]string) {
	var cValue C.Dictionary
	if len(value) > 0 {
		cValueKeyData := make([]unsafe.Pointer, len(value))
		cValueValueData := make([]unsafe.Pointer, len(value))
		var idx = 0
		for k, v := range value {
			cValueKeyData[idx] = NewString(k).Ptr()
			cValueValueData[idx] = NewString(v).Ptr()
			idx++
		}
		cValue.key_data = unsafe.Pointer(&cValueKeyData[0])
		cValue.value_data = unsafe.Pointer(&cValueValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSTimeZone_TimeZone_SetAbbreviationDictionary(cValue)
}

func (n NSTimeZone) Name() string {
	result_ := C.C_NSTimeZone_Name(n.Ptr())
	return MakeString(result_).String()
}

func (n NSTimeZone) Abbreviation() string {
	result_ := C.C_NSTimeZone_Abbreviation(n.Ptr())
	return MakeString(result_).String()
}

func (n NSTimeZone) SecondsFromGMT() int {
	result_ := C.C_NSTimeZone_SecondsFromGMT(n.Ptr())
	return int(result_)
}

func (n NSTimeZone) Data() []byte {
	result_ := C.C_NSTimeZone_Data(n.Ptr())
	return MakeData(result_).ToBytes()
}

func TimeZoneDataVersion() string {
	result_ := C.C_NSTimeZone_TimeZoneDataVersion()
	return MakeString(result_).String()
}

func (n NSTimeZone) IsDaylightSavingTime() bool {
	result_ := C.C_NSTimeZone_IsDaylightSavingTime(n.Ptr())
	return bool(result_)
}

func (n NSTimeZone) DaylightSavingTimeOffset() TimeInterval {
	result_ := C.C_NSTimeZone_DaylightSavingTimeOffset(n.Ptr())
	return TimeInterval(float64(result_))
}

func (n NSTimeZone) NextDaylightSavingTimeTransition() Date {
	result_ := C.C_NSTimeZone_NextDaylightSavingTimeTransition(n.Ptr())
	return MakeDate(result_)
}

func (n NSTimeZone) Description() string {
	result_ := C.C_NSTimeZone_Description(n.Ptr())
	return MakeString(result_).String()
}
