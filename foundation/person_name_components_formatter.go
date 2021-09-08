package foundation

// #include "person_name_components_formatter.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PersonNameComponentsFormatter interface {
	Formatter
	StringFromPersonNameComponents(components PersonNameComponents) string
	AnnotatedStringFromPersonNameComponents(components PersonNameComponents) AttributedString
	PersonNameComponentsFromString(_string string) PersonNameComponents
	Style() PersonNameComponentsFormatterStyle
	SetStyle(value PersonNameComponentsFormatterStyle)
	IsPhonetic() bool
	SetPhonetic(value bool)
}

type NSPersonNameComponentsFormatter struct {
	NSFormatter
}

func MakePersonNameComponentsFormatter(ptr unsafe.Pointer) NSPersonNameComponentsFormatter {
	return NSPersonNameComponentsFormatter{
		NSFormatter: MakeFormatter(ptr),
	}
}

func AllocPersonNameComponentsFormatter() NSPersonNameComponentsFormatter {
	result_ := C.C_NSPersonNameComponentsFormatter_AllocPersonNameComponentsFormatter()
	return MakePersonNameComponentsFormatter(result_)
}

func (n NSPersonNameComponentsFormatter) Init() NSPersonNameComponentsFormatter {
	result_ := C.C_NSPersonNameComponentsFormatter_Init(n.Ptr())
	return MakePersonNameComponentsFormatter(result_)
}

func NewPersonNameComponentsFormatter() NSPersonNameComponentsFormatter {
	result_ := C.C_NSPersonNameComponentsFormatter_NewPersonNameComponentsFormatter()
	return MakePersonNameComponentsFormatter(result_)
}

func (n NSPersonNameComponentsFormatter) Autorelease() NSPersonNameComponentsFormatter {
	result_ := C.C_NSPersonNameComponentsFormatter_Autorelease(n.Ptr())
	return MakePersonNameComponentsFormatter(result_)
}

func (n NSPersonNameComponentsFormatter) Retain() NSPersonNameComponentsFormatter {
	result_ := C.C_NSPersonNameComponentsFormatter_Retain(n.Ptr())
	return MakePersonNameComponentsFormatter(result_)
}

func PersonNameComponentsFormatter_LocalizedStringFromPersonNameComponents_Style_Options(components PersonNameComponents, nameFormatStyle PersonNameComponentsFormatterStyle, nameOptions PersonNameComponentsFormatterOptions) string {
	result_ := C.C_NSPersonNameComponentsFormatter_PersonNameComponentsFormatter_LocalizedStringFromPersonNameComponents_Style_Options(objc.ExtractPtr(components), C.int(int(nameFormatStyle)), C.uint(uint(nameOptions)))
	return MakeString(result_).String()
}

func (n NSPersonNameComponentsFormatter) StringFromPersonNameComponents(components PersonNameComponents) string {
	result_ := C.C_NSPersonNameComponentsFormatter_StringFromPersonNameComponents(n.Ptr(), objc.ExtractPtr(components))
	return MakeString(result_).String()
}

func (n NSPersonNameComponentsFormatter) AnnotatedStringFromPersonNameComponents(components PersonNameComponents) AttributedString {
	result_ := C.C_NSPersonNameComponentsFormatter_AnnotatedStringFromPersonNameComponents(n.Ptr(), objc.ExtractPtr(components))
	return MakeAttributedString(result_)
}

func (n NSPersonNameComponentsFormatter) PersonNameComponentsFromString(_string string) PersonNameComponents {
	result_ := C.C_NSPersonNameComponentsFormatter_PersonNameComponentsFromString(n.Ptr(), NewString(_string).Ptr())
	return MakePersonNameComponents(result_)
}

func (n NSPersonNameComponentsFormatter) Style() PersonNameComponentsFormatterStyle {
	result_ := C.C_NSPersonNameComponentsFormatter_Style(n.Ptr())
	return PersonNameComponentsFormatterStyle(int(result_))
}

func (n NSPersonNameComponentsFormatter) SetStyle(value PersonNameComponentsFormatterStyle) {
	C.C_NSPersonNameComponentsFormatter_SetStyle(n.Ptr(), C.int(int(value)))
}

func (n NSPersonNameComponentsFormatter) IsPhonetic() bool {
	result_ := C.C_NSPersonNameComponentsFormatter_IsPhonetic(n.Ptr())
	return bool(result_)
}

func (n NSPersonNameComponentsFormatter) SetPhonetic(value bool) {
	C.C_NSPersonNameComponentsFormatter_SetPhonetic(n.Ptr(), C.bool(value))
}
