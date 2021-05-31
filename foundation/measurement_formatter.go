package foundation

// #include "measurement_formatter.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type MeasurementFormatter interface {
	Formatter
	StringFromMeasurement(measurement Measurement) string
	StringFromUnit(unit Unit) string
	UnitOptions() MeasurementFormatterUnitOptions
	SetUnitOptions(value MeasurementFormatterUnitOptions)
	UnitStyle() FormattingUnitStyle
	SetUnitStyle(value FormattingUnitStyle)
	Locale() Locale
	SetLocale(value Locale)
	NumberFormatter() NumberFormatter
	SetNumberFormatter(value NumberFormatter)
}

type NSMeasurementFormatter struct {
	NSFormatter
}

func MakeMeasurementFormatter(ptr unsafe.Pointer) *NSMeasurementFormatter {
	if ptr == nil {
		return nil
	}
	return &NSMeasurementFormatter{
		NSFormatter: *MakeFormatter(ptr),
	}
}

func AllocMeasurementFormatter() *NSMeasurementFormatter {
	return MakeMeasurementFormatter(C.C_MeasurementFormatter_Alloc())
}

func (n *NSMeasurementFormatter) Init() MeasurementFormatter {
	result_ := C.C_NSMeasurementFormatter_Init(n.Ptr())
	return MakeMeasurementFormatter(result_)
}

func (n *NSMeasurementFormatter) StringFromMeasurement(measurement Measurement) string {
	result_ := C.C_NSMeasurementFormatter_StringFromMeasurement(n.Ptr(), objc.ExtractPtr(measurement))
	return MakeString(result_).String()
}

func (n *NSMeasurementFormatter) StringFromUnit(unit Unit) string {
	result_ := C.C_NSMeasurementFormatter_StringFromUnit(n.Ptr(), objc.ExtractPtr(unit))
	return MakeString(result_).String()
}

func (n *NSMeasurementFormatter) UnitOptions() MeasurementFormatterUnitOptions {
	result_ := C.C_NSMeasurementFormatter_UnitOptions(n.Ptr())
	return MeasurementFormatterUnitOptions(uint(result_))
}

func (n *NSMeasurementFormatter) SetUnitOptions(value MeasurementFormatterUnitOptions) {
	C.C_NSMeasurementFormatter_SetUnitOptions(n.Ptr(), C.uint(uint(value)))
}

func (n *NSMeasurementFormatter) UnitStyle() FormattingUnitStyle {
	result_ := C.C_NSMeasurementFormatter_UnitStyle(n.Ptr())
	return FormattingUnitStyle(int(result_))
}

func (n *NSMeasurementFormatter) SetUnitStyle(value FormattingUnitStyle) {
	C.C_NSMeasurementFormatter_SetUnitStyle(n.Ptr(), C.int(int(value)))
}

func (n *NSMeasurementFormatter) Locale() Locale {
	result_ := C.C_NSMeasurementFormatter_Locale(n.Ptr())
	return MakeLocale(result_)
}

func (n *NSMeasurementFormatter) SetLocale(value Locale) {
	C.C_NSMeasurementFormatter_SetLocale(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMeasurementFormatter) NumberFormatter() NumberFormatter {
	result_ := C.C_NSMeasurementFormatter_NumberFormatter(n.Ptr())
	return MakeNumberFormatter(result_)
}

func (n *NSMeasurementFormatter) SetNumberFormatter(value NumberFormatter) {
	C.C_NSMeasurementFormatter_SetNumberFormatter(n.Ptr(), objc.ExtractPtr(value))
}
