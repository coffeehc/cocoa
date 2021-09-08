package foundation

// #include "regular_expression.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type RegularExpression interface {
	objc.Object
	NumberOfMatchesInString_Options_Range(_string string, options MatchingOptions, _range Range) uint
	MatchesInString_Options_Range(_string string, options MatchingOptions, _range Range) []TextCheckingResult
	FirstMatchInString_Options_Range(_string string, options MatchingOptions, _range Range) TextCheckingResult
	RangeOfFirstMatchInString_Options_Range(_string string, options MatchingOptions, _range Range) Range
	ReplaceMatchesInString_Options_Range_WithTemplate(_string MutableString, options MatchingOptions, _range Range, templ string) uint
	StringByReplacingMatchesInString_Options_Range_WithTemplate(_string string, options MatchingOptions, _range Range, templ string) string
	ReplacementStringForResult_InString_Offset_Template(result TextCheckingResult, _string string, offset int, templ string) string
	Pattern() string
	Options() RegularExpressionOptions
	NumberOfCaptureGroups() uint
}

type NSRegularExpression struct {
	objc.NSObject
}

func MakeRegularExpression(ptr unsafe.Pointer) NSRegularExpression {
	return NSRegularExpression{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocRegularExpression() NSRegularExpression {
	result_ := C.C_NSRegularExpression_AllocRegularExpression()
	return MakeRegularExpression(result_)
}

func (n NSRegularExpression) Init() NSRegularExpression {
	result_ := C.C_NSRegularExpression_Init(n.Ptr())
	return MakeRegularExpression(result_)
}

func NewRegularExpression() NSRegularExpression {
	result_ := C.C_NSRegularExpression_NewRegularExpression()
	return MakeRegularExpression(result_)
}

func (n NSRegularExpression) Autorelease() NSRegularExpression {
	result_ := C.C_NSRegularExpression_Autorelease(n.Ptr())
	return MakeRegularExpression(result_)
}

func (n NSRegularExpression) Retain() NSRegularExpression {
	result_ := C.C_NSRegularExpression_Retain(n.Ptr())
	return MakeRegularExpression(result_)
}

func (n NSRegularExpression) NumberOfMatchesInString_Options_Range(_string string, options MatchingOptions, _range Range) uint {
	result_ := C.C_NSRegularExpression_NumberOfMatchesInString_Options_Range(n.Ptr(), NewString(_string).Ptr(), C.uint(uint(options)), *(*C.NSRange)(ToNSRangePointer(_range)))
	return uint(result_)
}

func (n NSRegularExpression) MatchesInString_Options_Range(_string string, options MatchingOptions, _range Range) []TextCheckingResult {
	result_ := C.C_NSRegularExpression_MatchesInString_Options_Range(n.Ptr(), NewString(_string).Ptr(), C.uint(uint(options)), *(*C.NSRange)(ToNSRangePointer(_range)))
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]TextCheckingResult, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTextCheckingResult(r)
	}
	return goResult_
}

func (n NSRegularExpression) FirstMatchInString_Options_Range(_string string, options MatchingOptions, _range Range) TextCheckingResult {
	result_ := C.C_NSRegularExpression_FirstMatchInString_Options_Range(n.Ptr(), NewString(_string).Ptr(), C.uint(uint(options)), *(*C.NSRange)(ToNSRangePointer(_range)))
	return MakeTextCheckingResult(result_)
}

func (n NSRegularExpression) RangeOfFirstMatchInString_Options_Range(_string string, options MatchingOptions, _range Range) Range {
	result_ := C.C_NSRegularExpression_RangeOfFirstMatchInString_Options_Range(n.Ptr(), NewString(_string).Ptr(), C.uint(uint(options)), *(*C.NSRange)(ToNSRangePointer(_range)))
	return FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSRegularExpression) ReplaceMatchesInString_Options_Range_WithTemplate(_string MutableString, options MatchingOptions, _range Range, templ string) uint {
	result_ := C.C_NSRegularExpression_ReplaceMatchesInString_Options_Range_WithTemplate(n.Ptr(), objc.ExtractPtr(_string), C.uint(uint(options)), *(*C.NSRange)(ToNSRangePointer(_range)), NewString(templ).Ptr())
	return uint(result_)
}

func (n NSRegularExpression) StringByReplacingMatchesInString_Options_Range_WithTemplate(_string string, options MatchingOptions, _range Range, templ string) string {
	result_ := C.C_NSRegularExpression_StringByReplacingMatchesInString_Options_Range_WithTemplate(n.Ptr(), NewString(_string).Ptr(), C.uint(uint(options)), *(*C.NSRange)(ToNSRangePointer(_range)), NewString(templ).Ptr())
	return MakeString(result_).String()
}

func RegularExpression_EscapedTemplateForString(_string string) string {
	result_ := C.C_NSRegularExpression_RegularExpression_EscapedTemplateForString(NewString(_string).Ptr())
	return MakeString(result_).String()
}

func RegularExpression_EscapedPatternForString(_string string) string {
	result_ := C.C_NSRegularExpression_RegularExpression_EscapedPatternForString(NewString(_string).Ptr())
	return MakeString(result_).String()
}

func (n NSRegularExpression) ReplacementStringForResult_InString_Offset_Template(result TextCheckingResult, _string string, offset int, templ string) string {
	result_ := C.C_NSRegularExpression_ReplacementStringForResult_InString_Offset_Template(n.Ptr(), objc.ExtractPtr(result), NewString(_string).Ptr(), C.int(offset), NewString(templ).Ptr())
	return MakeString(result_).String()
}

func (n NSRegularExpression) Pattern() string {
	result_ := C.C_NSRegularExpression_Pattern(n.Ptr())
	return MakeString(result_).String()
}

func (n NSRegularExpression) Options() RegularExpressionOptions {
	result_ := C.C_NSRegularExpression_Options(n.Ptr())
	return RegularExpressionOptions(uint(result_))
}

func (n NSRegularExpression) NumberOfCaptureGroups() uint {
	result_ := C.C_NSRegularExpression_NumberOfCaptureGroups(n.Ptr())
	return uint(result_)
}
