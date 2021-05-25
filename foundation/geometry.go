package foundation

// #include "geometry.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"unsafe"
)

type Point = coregraphics.Point

func ToNSPointPointer(p Point) unsafe.Pointer {
	return coregraphics.ToCGPointPointer(p)
}

func FromNSPointerPointer(p unsafe.Pointer) Point {
	return coregraphics.FromCGPointPointer(p)
}

type Size = coregraphics.Size

func ToNSSizePointer(s Size) unsafe.Pointer {
	return coregraphics.ToCGSizePointer(s)
}

func FromNSSizePointer(p unsafe.Pointer) Size {
	return coregraphics.FromCGSizePointer(p)
}

type Rect = coregraphics.Rect

func ToNSRectPointer(r Rect) unsafe.Pointer {
	return coregraphics.ToCGRectPointer(r)
}

func FromNSRectPointer(p unsafe.Pointer) Rect {
	return coregraphics.FromCGRectPointer(p)
}

// MakeRect create a Rect struct
func MakeRect(x, y, width, height float64) Rect {
	return Rect{
		Origin: coregraphics.Point{X: x, Y: y},
		Size:   coregraphics.Size{Width: width, Height: height},
	}
}

// EdgeInsets is a description of the distance between the edges of two rectangles.
type EdgeInsets struct {
	Top    float64
	Left   float64
	Bottom float64
	Right  float64
}

func ToNSEdgeInsetsPointer(e EdgeInsets) unsafe.Pointer {
	return unsafe.Pointer(&C.NSEdgeInsets{
		C.double(e.Top),
		C.double(e.Left),
		C.double(e.Bottom),
		C.double(e.Right),
	})
}

func FromNSEdgeInsetsPointer(p unsafe.Pointer) EdgeInsets {
	ns := *(*C.NSEdgeInsets)(p)
	return EdgeInsets{
		Top:    float64(ns.top),
		Left:   float64(ns.left),
		Bottom: float64(ns.bottom),
		Right:  float64(ns.right),
	}
}

type RectEdge int

const (
	RectEdgeMinX RectEdge = 0
	RectEdgeMinY RectEdge = 1
	RectEdgeMaxX RectEdge = 2
	RectEdgeMaxY RectEdge = 3

	MinXEdge RectEdge = RectEdgeMinX
	MinYEdge RectEdge = RectEdgeMinY
	MaxXEdge RectEdge = RectEdgeMaxX
	MaxYEdge RectEdge = RectEdgeMaxY
)
