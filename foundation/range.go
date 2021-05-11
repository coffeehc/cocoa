package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "range.h"
import "C"
import "unsafe"

// Range is a structure used to describe a portion of a series, such as characters in a string or objects in an array
type Range struct {
	Location uint64
	Length   uint64
}

type RangePointer *Range

func ToNSRangePointer(r Range) unsafe.Pointer {
	return unsafe.Pointer(&C.NSRange{
		location: C.ulong(r.Location),
		length:   C.ulong(r.Length),
	})
}

func FromNSRangePointer(p unsafe.Pointer) Range {
	ns := *(*C.NSRange)(p)
	return Range{
		Location: uint64(ns.location),
		Length:   uint64(ns.length),
	}
}
