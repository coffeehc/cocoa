package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include <Foundation/NSGeometry.h>
// #include <Foundation/NSRange.h>
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"github.com/hsiafan/glow/unsafex"
	"unsafe"
)

func toNSRect(rect Rect) C.NSRect {
	return C.NSRect{
		C.NSPoint{C.double(rect.X), C.double(rect.Y)},
		C.NSSize{C.double(rect.Width), C.double(rect.Height)},
	}
}
func toRect(nsRect C.NSRect) Rect {
	return MakeRect(float64(nsRect.origin.x), float64(nsRect.origin.y),
		float64(nsRect.size.width), float64(nsRect.size.height))
}

func toNSPoint(point Point) C.NSPoint {
	return C.NSPoint{C.double(point.X), C.double(point.Y)}
}
func toPoint(point C.NSPoint) Point {
	return Point{X: float64(point.x), Y: float64(point.y)}
}

func toNSSize(size Size) C.NSSize {
	return C.NSSize{C.double(size.Width), C.double(size.Height)}
}
func toSize(size C.NSSize) Size {
	return Size{Width: float64(size.width), Height: float64(size.height)}
}

func toNSEdgeInsets(insets EdgeInsets) C.NSEdgeInsets {
	return C.NSEdgeInsets{C.double(insets.Top), C.double(insets.Left), C.double(insets.Bottom), C.double(insets.Right)}
}
func toEdgeInsets(insets C.NSEdgeInsets) EdgeInsets {
	return EdgeInsets{
		Top:    float64(insets.top),
		Left:   float64(insets.left),
		Bottom: float64(insets.bottom),
		Right:  float64(insets.right),
	}
}

func toNSRange(r Range) C.NSRange {
	return C.NSRange{
		location: C.ulong(r.Location),
		length:   C.ulong(r.Length),
	}
}

func toRange(r C.NSRange) Range {
	return Range{
		Location: uint(r.location),
		Length:   uint(r.length),
	}
}

func toPointer(o objc.Object) unsafe.Pointer {
	if unsafex.InterfaceIsNil(o) {
		return nil
	}
	return o.Ptr()
}
