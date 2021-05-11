package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "geometry.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"unsafe"
)

// ZeroRect for convenient
var ZeroRect Rect

type Point coregraphics.Point

func ToNSPointPointer(p Point) unsafe.Pointer {
	return unsafe.Pointer(&C.NSPoint{C.double(p.X), C.double(p.Y)})
}

func FromNSPointerPointer(p unsafe.Pointer) Point {
	nsPoint := *(*C.NSPoint)(p)
	return Point{
		X: float64(nsPoint.x),
		Y: float64(nsPoint.y),
	}
}

type Size coregraphics.Size

func ToNSSizePointer(s Size) unsafe.Pointer {
	return unsafe.Pointer(&C.NSSize{C.double(s.Width), C.double(s.Height)})
}

func FromNSSizePointer(p unsafe.Pointer) Size {
	ns := *(*C.NSSize)(p)
	return Size{
		Width:  float64(ns.width),
		Height: float64(ns.height),
	}
}

type Rect coregraphics.Rect

func ToNSRectPointer(r Rect) unsafe.Pointer {
	return unsafe.Pointer(&C.NSRect{
		C.NSPoint{C.double(r.Origin.X), C.double(r.Origin.Y)},
		C.NSSize{C.double(r.Size.Width), C.double(r.Size.Height)},
	})
}

func FromNSRectPointer(p unsafe.Pointer) Rect {
	ns := *(*C.NSRect)(p)
	return Rect{
		Origin: coregraphics.Point{
			X: float64(ns.origin.x),
			Y: float64(ns.origin.y),
		},
		Size: coregraphics.Size{
			Width:  float64(ns.size.width),
			Height: float64(ns.size.height),
		},
	}
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
