package foundation

// #include "measurement.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Measurement interface {
	objc.Object
	CanBeConvertedToUnit(unit Unit) bool
	MeasurementByConvertingToUnit(unit Unit) Measurement
	MeasurementByAddingMeasurement(measurement Measurement) Measurement
	MeasurementBySubtractingMeasurement(measurement Measurement) Measurement
	DoubleValue() float64
}

type NSMeasurement struct {
	objc.NSObject
}

func MakeMeasurement(ptr unsafe.Pointer) NSMeasurement {
	return NSMeasurement{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocMeasurement() NSMeasurement {
	return MakeMeasurement(C.C_Measurement_Alloc())
}

func (n NSMeasurement) CanBeConvertedToUnit(unit Unit) bool {
	result_ := C.C_NSMeasurement_CanBeConvertedToUnit(n.Ptr(), objc.ExtractPtr(unit))
	return bool(result_)
}

func (n NSMeasurement) MeasurementByConvertingToUnit(unit Unit) Measurement {
	result_ := C.C_NSMeasurement_MeasurementByConvertingToUnit(n.Ptr(), objc.ExtractPtr(unit))
	return MakeMeasurement(result_)
}

func (n NSMeasurement) MeasurementByAddingMeasurement(measurement Measurement) Measurement {
	result_ := C.C_NSMeasurement_MeasurementByAddingMeasurement(n.Ptr(), objc.ExtractPtr(measurement))
	return MakeMeasurement(result_)
}

func (n NSMeasurement) MeasurementBySubtractingMeasurement(measurement Measurement) Measurement {
	result_ := C.C_NSMeasurement_MeasurementBySubtractingMeasurement(n.Ptr(), objc.ExtractPtr(measurement))
	return MakeMeasurement(result_)
}

func (n NSMeasurement) DoubleValue() float64 {
	result_ := C.C_NSMeasurement_DoubleValue(n.Ptr())
	return float64(result_)
}
