package coregraphics

// #include "geometry.h"
import "C"
import "unsafe"

type Float = float64

type Point struct {
	X float64
	Y float64
}

func ToCGPointPointer(p Point) unsafe.Pointer {
	return unsafe.Pointer(&C.CGPoint{C.double(p.X), C.double(p.Y)})
}

func FromCGPointPointer(p unsafe.Pointer) Point {
	nsPoint := *(*C.CGPoint)(p)
	return Point{
		X: float64(nsPoint.x),
		Y: float64(nsPoint.y),
	}
}

type Size struct {
	Width  float64
	Height float64
}

func ToCGSizePointer(s Size) unsafe.Pointer {
	return unsafe.Pointer(&C.CGSize{C.double(s.Width), C.double(s.Height)})
}

func FromCGSizePointer(p unsafe.Pointer) Size {
	ns := *(*C.CGSize)(p)
	return Size{
		Width:  float64(ns.width),
		Height: float64(ns.height),
	}
}

type Rect struct {
	Origin Point
	Size   Size
}

func ToCGRectPointer(r Rect) unsafe.Pointer {
	return unsafe.Pointer(&C.CGRect{
		C.CGPoint{C.double(r.Origin.X), C.double(r.Origin.Y)},
		C.CGSize{C.double(r.Size.Width), C.double(r.Size.Height)},
	})
}

func FromCGRectPointer(p unsafe.Pointer) Rect {
	ns := *(*C.CGRect)(p)
	return Rect{
		Origin: Point{
			X: float64(ns.origin.x),
			Y: float64(ns.origin.y),
		},
		Size: Size{
			Width:  float64(ns.size.width),
			Height: float64(ns.size.height),
		},
	}
}
