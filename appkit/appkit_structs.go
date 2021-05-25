package appkit

// #include "appkit_structs.h"
import "C"
import "unsafe"

type DirectionalEdgeInsets struct {
	Top      float64
	Leading  float64
	Bottom   float64
	Trailing float64
}

func MakeDirectionalEdgeInsets(top, leading, bottom, trailing float64) DirectionalEdgeInsets {
	return DirectionalEdgeInsets{
		Top:      top,
		Leading:  leading,
		Bottom:   bottom,
		Trailing: trailing,
	}
}

func ToNSDirectionalEdgeInsetsPointer(e DirectionalEdgeInsets) unsafe.Pointer {
	return unsafe.Pointer(&C.NSDirectionalEdgeInsets{
		C.double(e.Top),
		C.double(e.Leading),
		C.double(e.Bottom),
		C.double(e.Trailing),
	})
}

func FromNSDirectionalEdgeInsetsPointer(p unsafe.Pointer) DirectionalEdgeInsets {
	ns := *(*C.NSDirectionalEdgeInsets)(p)
	return DirectionalEdgeInsets{
		Top:      float64(ns.top),
		Leading:  float64(ns.leading),
		Bottom:   float64(ns.bottom),
		Trailing: float64(ns.trailing),
	}
}
