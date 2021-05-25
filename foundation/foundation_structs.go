package foundation

// #include <Foundation/NSAffineTransform.h>
import "C"
import "unsafe"

type AffineTransformStruct struct {
	M11 float64
	M12 float64
	M21 float64
	M22 float64
	TX  float64
	TY  float64
}

func FromNSAffineTransformStructPointer(p unsafe.Pointer) AffineTransformStruct {
	ns := *(*C.NSAffineTransformStruct)(p)
	return AffineTransformStruct{
		M11: float64(ns.m11),
		M12: float64(ns.m12),
		M21: float64(ns.m21),
		M22: float64(ns.m22),
		TX:  float64(ns.tX),
		TY:  float64(ns.tY),
	}

}

func ToNSAffineTransformStructPointer(s AffineTransformStruct) unsafe.Pointer {
	return unsafe.Pointer(&C.NSAffineTransformStruct{
		m11: C.double(s.M11),
		m12: C.double(s.M12),
		m21: C.double(s.M21),
		m22: C.double(s.M22),
		tX:  C.double(s.TX),
		tY:  C.double(s.TY),
	})
}
