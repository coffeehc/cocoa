package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "decimal.h"
import "C"
import "unsafe"

const DecimalMaxSize = 8

type Decimal struct {
	meta     int32
	mantissa [DecimalMaxSize]uint16
}

func (d *Decimal) ToNSDecimalPointer() unsafe.Pointer {
	// TODO: to be implemented
	panic("to be implemented")
}

func FromNSDecimalPointer(p unsafe.Pointer) Decimal {
	_ = *(*C.NSDecimal)(p)
	// TODO: to be implemented
	panic("to be implemented")
}
