package foundation

// #include "text_checking_result.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TextCheckingResult interface {
	objc.Object
	RangeAtIndex(idx uint) Range
	ResultByAdjustingRangesWithOffset(offset int) TextCheckingResult
	RangeWithName(name string) Range
	Range() Range
	NumberOfRanges() uint
	ReplacementString() string
	RegularExpression() RegularExpression
	Components() map[TextCheckingKey]string
	URL() URL
	AddressComponents() map[TextCheckingKey]string
	PhoneNumber() string
	Date() Date
	Duration() TimeInterval
	TimeZone() TimeZone
	Orthography() Orthography
	AlternativeStrings() []string
}

type NSTextCheckingResult struct {
	objc.NSObject
}

func MakeTextCheckingResult(ptr unsafe.Pointer) NSTextCheckingResult {
	return NSTextCheckingResult{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocTextCheckingResult() NSTextCheckingResult {
	return MakeTextCheckingResult(C.C_TextCheckingResult_Alloc())
}

func (n NSTextCheckingResult) Init() TextCheckingResult {
	result_ := C.C_NSTextCheckingResult_Init(n.Ptr())
	return MakeTextCheckingResult(result_)
}

func (n NSTextCheckingResult) RangeAtIndex(idx uint) Range {
	result_ := C.C_NSTextCheckingResult_RangeAtIndex(n.Ptr(), C.uint(idx))
	return FromNSRangePointer(unsafe.Pointer(&result_))
}

func TextCheckingResult_ReplacementCheckingResultWithRange_ReplacementString(_range Range, replacementString string) TextCheckingResult {
	result_ := C.C_NSTextCheckingResult_TextCheckingResult_ReplacementCheckingResultWithRange_ReplacementString(*(*C.NSRange)(ToNSRangePointer(_range)), NewString(replacementString).Ptr())
	return MakeTextCheckingResult(result_)
}

func TextCheckingResult_LinkCheckingResultWithRange_URL(_range Range, url URL) TextCheckingResult {
	result_ := C.C_NSTextCheckingResult_TextCheckingResult_LinkCheckingResultWithRange_URL(*(*C.NSRange)(ToNSRangePointer(_range)), objc.ExtractPtr(url))
	return MakeTextCheckingResult(result_)
}

func TextCheckingResult_AddressCheckingResultWithRange_Components(_range Range, components map[TextCheckingKey]string) TextCheckingResult {
	var cComponents C.Dictionary
	if len(components) > 0 {
		cComponentsKeyData := make([]unsafe.Pointer, len(components))
		cComponentsValueData := make([]unsafe.Pointer, len(components))
		var idx = 0
		for k, v := range components {
			cComponentsKeyData[idx] = NewString(string(k)).Ptr()
			cComponentsValueData[idx] = NewString(v).Ptr()
			idx++
		}
		cComponents.key_data = unsafe.Pointer(&cComponentsKeyData[0])
		cComponents.value_data = unsafe.Pointer(&cComponentsValueData[0])
		cComponents.len = C.int(len(components))
	}
	result_ := C.C_NSTextCheckingResult_TextCheckingResult_AddressCheckingResultWithRange_Components(*(*C.NSRange)(ToNSRangePointer(_range)), cComponents)
	return MakeTextCheckingResult(result_)
}

func TextCheckingResult_TransitInformationCheckingResultWithRange_Components(_range Range, components map[TextCheckingKey]string) TextCheckingResult {
	var cComponents C.Dictionary
	if len(components) > 0 {
		cComponentsKeyData := make([]unsafe.Pointer, len(components))
		cComponentsValueData := make([]unsafe.Pointer, len(components))
		var idx = 0
		for k, v := range components {
			cComponentsKeyData[idx] = NewString(string(k)).Ptr()
			cComponentsValueData[idx] = NewString(v).Ptr()
			idx++
		}
		cComponents.key_data = unsafe.Pointer(&cComponentsKeyData[0])
		cComponents.value_data = unsafe.Pointer(&cComponentsValueData[0])
		cComponents.len = C.int(len(components))
	}
	result_ := C.C_NSTextCheckingResult_TextCheckingResult_TransitInformationCheckingResultWithRange_Components(*(*C.NSRange)(ToNSRangePointer(_range)), cComponents)
	return MakeTextCheckingResult(result_)
}

func TextCheckingResult_PhoneNumberCheckingResultWithRange_PhoneNumber(_range Range, phoneNumber string) TextCheckingResult {
	result_ := C.C_NSTextCheckingResult_TextCheckingResult_PhoneNumberCheckingResultWithRange_PhoneNumber(*(*C.NSRange)(ToNSRangePointer(_range)), NewString(phoneNumber).Ptr())
	return MakeTextCheckingResult(result_)
}

func TextCheckingResult_DateCheckingResultWithRange_Date(_range Range, date Date) TextCheckingResult {
	result_ := C.C_NSTextCheckingResult_TextCheckingResult_DateCheckingResultWithRange_Date(*(*C.NSRange)(ToNSRangePointer(_range)), objc.ExtractPtr(date))
	return MakeTextCheckingResult(result_)
}

func TextCheckingResult_DateCheckingResultWithRange_Date_TimeZone_Duration(_range Range, date Date, timeZone TimeZone, duration TimeInterval) TextCheckingResult {
	result_ := C.C_NSTextCheckingResult_TextCheckingResult_DateCheckingResultWithRange_Date_TimeZone_Duration(*(*C.NSRange)(ToNSRangePointer(_range)), objc.ExtractPtr(date), objc.ExtractPtr(timeZone), C.double(float64(duration)))
	return MakeTextCheckingResult(result_)
}

func TextCheckingResult_DashCheckingResultWithRange_ReplacementString(_range Range, replacementString string) TextCheckingResult {
	result_ := C.C_NSTextCheckingResult_TextCheckingResult_DashCheckingResultWithRange_ReplacementString(*(*C.NSRange)(ToNSRangePointer(_range)), NewString(replacementString).Ptr())
	return MakeTextCheckingResult(result_)
}

func TextCheckingResult_QuoteCheckingResultWithRange_ReplacementString(_range Range, replacementString string) TextCheckingResult {
	result_ := C.C_NSTextCheckingResult_TextCheckingResult_QuoteCheckingResultWithRange_ReplacementString(*(*C.NSRange)(ToNSRangePointer(_range)), NewString(replacementString).Ptr())
	return MakeTextCheckingResult(result_)
}

func TextCheckingResult_SpellCheckingResultWithRange(_range Range) TextCheckingResult {
	result_ := C.C_NSTextCheckingResult_TextCheckingResult_SpellCheckingResultWithRange(*(*C.NSRange)(ToNSRangePointer(_range)))
	return MakeTextCheckingResult(result_)
}

func TextCheckingResult_CorrectionCheckingResultWithRange_ReplacementString(_range Range, replacementString string) TextCheckingResult {
	result_ := C.C_NSTextCheckingResult_TextCheckingResult_CorrectionCheckingResultWithRange_ReplacementString(*(*C.NSRange)(ToNSRangePointer(_range)), NewString(replacementString).Ptr())
	return MakeTextCheckingResult(result_)
}

func TextCheckingResult_OrthographyCheckingResultWithRange_Orthography(_range Range, orthography Orthography) TextCheckingResult {
	result_ := C.C_NSTextCheckingResult_TextCheckingResult_OrthographyCheckingResultWithRange_Orthography(*(*C.NSRange)(ToNSRangePointer(_range)), objc.ExtractPtr(orthography))
	return MakeTextCheckingResult(result_)
}

func (n NSTextCheckingResult) ResultByAdjustingRangesWithOffset(offset int) TextCheckingResult {
	result_ := C.C_NSTextCheckingResult_ResultByAdjustingRangesWithOffset(n.Ptr(), C.int(offset))
	return MakeTextCheckingResult(result_)
}

func (n NSTextCheckingResult) RangeWithName(name string) Range {
	result_ := C.C_NSTextCheckingResult_RangeWithName(n.Ptr(), NewString(name).Ptr())
	return FromNSRangePointer(unsafe.Pointer(&result_))
}

func TextCheckingResult_CorrectionCheckingResultWithRange_ReplacementString_AlternativeStrings(_range Range, replacementString string, alternativeStrings []string) TextCheckingResult {
	var cAlternativeStrings C.Array
	if len(alternativeStrings) > 0 {
		cAlternativeStringsData := make([]unsafe.Pointer, len(alternativeStrings))
		for idx, v := range alternativeStrings {
			cAlternativeStringsData[idx] = NewString(v).Ptr()
		}
		cAlternativeStrings.data = unsafe.Pointer(&cAlternativeStringsData[0])
		cAlternativeStrings.len = C.int(len(alternativeStrings))
	}
	result_ := C.C_NSTextCheckingResult_TextCheckingResult_CorrectionCheckingResultWithRange_ReplacementString_AlternativeStrings(*(*C.NSRange)(ToNSRangePointer(_range)), NewString(replacementString).Ptr(), cAlternativeStrings)
	return MakeTextCheckingResult(result_)
}

func (n NSTextCheckingResult) Range() Range {
	result_ := C.C_NSTextCheckingResult_Range(n.Ptr())
	return FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSTextCheckingResult) NumberOfRanges() uint {
	result_ := C.C_NSTextCheckingResult_NumberOfRanges(n.Ptr())
	return uint(result_)
}

func (n NSTextCheckingResult) ReplacementString() string {
	result_ := C.C_NSTextCheckingResult_ReplacementString(n.Ptr())
	return MakeString(result_).String()
}

func (n NSTextCheckingResult) RegularExpression() RegularExpression {
	result_ := C.C_NSTextCheckingResult_RegularExpression(n.Ptr())
	return MakeRegularExpression(result_)
}

func (n NSTextCheckingResult) Components() map[TextCheckingKey]string {
	result_ := C.C_NSTextCheckingResult_Components(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[TextCheckingKey]string)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[TextCheckingKey(MakeString(k).String())] = MakeString(v).String()
	}
	return goResult_
}

func (n NSTextCheckingResult) URL() URL {
	result_ := C.C_NSTextCheckingResult_URL(n.Ptr())
	return MakeURL(result_)
}

func (n NSTextCheckingResult) AddressComponents() map[TextCheckingKey]string {
	result_ := C.C_NSTextCheckingResult_AddressComponents(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[TextCheckingKey]string)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[TextCheckingKey(MakeString(k).String())] = MakeString(v).String()
	}
	return goResult_
}

func (n NSTextCheckingResult) PhoneNumber() string {
	result_ := C.C_NSTextCheckingResult_PhoneNumber(n.Ptr())
	return MakeString(result_).String()
}

func (n NSTextCheckingResult) Date() Date {
	result_ := C.C_NSTextCheckingResult_Date(n.Ptr())
	return MakeDate(result_)
}

func (n NSTextCheckingResult) Duration() TimeInterval {
	result_ := C.C_NSTextCheckingResult_Duration(n.Ptr())
	return TimeInterval(float64(result_))
}

func (n NSTextCheckingResult) TimeZone() TimeZone {
	result_ := C.C_NSTextCheckingResult_TimeZone(n.Ptr())
	return MakeTimeZone(result_)
}

func (n NSTextCheckingResult) Orthography() Orthography {
	result_ := C.C_NSTextCheckingResult_Orthography(n.Ptr())
	return MakeOrthography(result_)
}

func (n NSTextCheckingResult) AlternativeStrings() []string {
	result_ := C.C_NSTextCheckingResult_AlternativeStrings(n.Ptr())
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
