package foundation

// #include "byte_count_formatter.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ByteCountFormatter interface {
	Formatter
	StringFromMeasurement(measurement Measurement) string
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	CountStyle() ByteCountFormatterCountStyle
	SetCountStyle(value ByteCountFormatterCountStyle)
	AllowsNonnumericFormatting() bool
	SetAllowsNonnumericFormatting(value bool)
	IncludesActualByteCount() bool
	SetIncludesActualByteCount(value bool)
	IsAdaptive() bool
	SetAdaptive(value bool)
	AllowedUnits() ByteCountFormatterUnits
	SetAllowedUnits(value ByteCountFormatterUnits)
	IncludesCount() bool
	SetIncludesCount(value bool)
	IncludesUnit() bool
	SetIncludesUnit(value bool)
	ZeroPadsFractionDigits() bool
	SetZeroPadsFractionDigits(value bool)
}

type NSByteCountFormatter struct {
	NSFormatter
}

func MakeByteCountFormatter(ptr unsafe.Pointer) *NSByteCountFormatter {
	if ptr == nil {
		return nil
	}
	return &NSByteCountFormatter{
		NSFormatter: *MakeFormatter(ptr),
	}
}

func AllocByteCountFormatter() *NSByteCountFormatter {
	return MakeByteCountFormatter(C.C_ByteCountFormatter_Alloc())
}

func (n *NSByteCountFormatter) Init() ByteCountFormatter {
	result_ := C.C_NSByteCountFormatter_Init(n.Ptr())
	return MakeByteCountFormatter(result_)
}

func (n *NSByteCountFormatter) StringFromMeasurement(measurement Measurement) string {
	result_ := C.C_NSByteCountFormatter_StringFromMeasurement(n.Ptr(), objc.ExtractPtr(measurement))
	return MakeString(result_).String()
}

func ByteCountFormatter_StringFromMeasurement_CountStyle(measurement Measurement, countStyle ByteCountFormatterCountStyle) string {
	result_ := C.C_NSByteCountFormatter_ByteCountFormatter_StringFromMeasurement_CountStyle(objc.ExtractPtr(measurement), C.int(int(countStyle)))
	return MakeString(result_).String()
}

func (n *NSByteCountFormatter) FormattingContext() FormattingContext {
	result_ := C.C_NSByteCountFormatter_FormattingContext(n.Ptr())
	return FormattingContext(int(result_))
}

func (n *NSByteCountFormatter) SetFormattingContext(value FormattingContext) {
	C.C_NSByteCountFormatter_SetFormattingContext(n.Ptr(), C.int(int(value)))
}

func (n *NSByteCountFormatter) CountStyle() ByteCountFormatterCountStyle {
	result_ := C.C_NSByteCountFormatter_CountStyle(n.Ptr())
	return ByteCountFormatterCountStyle(int(result_))
}

func (n *NSByteCountFormatter) SetCountStyle(value ByteCountFormatterCountStyle) {
	C.C_NSByteCountFormatter_SetCountStyle(n.Ptr(), C.int(int(value)))
}

func (n *NSByteCountFormatter) AllowsNonnumericFormatting() bool {
	result_ := C.C_NSByteCountFormatter_AllowsNonnumericFormatting(n.Ptr())
	return bool(result_)
}

func (n *NSByteCountFormatter) SetAllowsNonnumericFormatting(value bool) {
	C.C_NSByteCountFormatter_SetAllowsNonnumericFormatting(n.Ptr(), C.bool(value))
}

func (n *NSByteCountFormatter) IncludesActualByteCount() bool {
	result_ := C.C_NSByteCountFormatter_IncludesActualByteCount(n.Ptr())
	return bool(result_)
}

func (n *NSByteCountFormatter) SetIncludesActualByteCount(value bool) {
	C.C_NSByteCountFormatter_SetIncludesActualByteCount(n.Ptr(), C.bool(value))
}

func (n *NSByteCountFormatter) IsAdaptive() bool {
	result_ := C.C_NSByteCountFormatter_IsAdaptive(n.Ptr())
	return bool(result_)
}

func (n *NSByteCountFormatter) SetAdaptive(value bool) {
	C.C_NSByteCountFormatter_SetAdaptive(n.Ptr(), C.bool(value))
}

func (n *NSByteCountFormatter) AllowedUnits() ByteCountFormatterUnits {
	result_ := C.C_NSByteCountFormatter_AllowedUnits(n.Ptr())
	return ByteCountFormatterUnits(uint(result_))
}

func (n *NSByteCountFormatter) SetAllowedUnits(value ByteCountFormatterUnits) {
	C.C_NSByteCountFormatter_SetAllowedUnits(n.Ptr(), C.uint(uint(value)))
}

func (n *NSByteCountFormatter) IncludesCount() bool {
	result_ := C.C_NSByteCountFormatter_IncludesCount(n.Ptr())
	return bool(result_)
}

func (n *NSByteCountFormatter) SetIncludesCount(value bool) {
	C.C_NSByteCountFormatter_SetIncludesCount(n.Ptr(), C.bool(value))
}

func (n *NSByteCountFormatter) IncludesUnit() bool {
	result_ := C.C_NSByteCountFormatter_IncludesUnit(n.Ptr())
	return bool(result_)
}

func (n *NSByteCountFormatter) SetIncludesUnit(value bool) {
	C.C_NSByteCountFormatter_SetIncludesUnit(n.Ptr(), C.bool(value))
}

func (n *NSByteCountFormatter) ZeroPadsFractionDigits() bool {
	result_ := C.C_NSByteCountFormatter_ZeroPadsFractionDigits(n.Ptr())
	return bool(result_)
}

func (n *NSByteCountFormatter) SetZeroPadsFractionDigits(value bool) {
	C.C_NSByteCountFormatter_SetZeroPadsFractionDigits(n.Ptr(), C.bool(value))
}
