package appkit

// #include "appearance.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Appearance interface {
	objc.Object
	BestMatchFromAppearancesWithNames(appearances []AppearanceName) AppearanceName
	Name() AppearanceName
	AllowsVibrancy() bool
}

type NSAppearance struct {
	objc.NSObject
}

func MakeAppearance(ptr unsafe.Pointer) NSAppearance {
	return NSAppearance{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocAppearance() NSAppearance {
	return MakeAppearance(C.C_Appearance_Alloc())
}

func (n NSAppearance) InitWithAppearanceNamed_Bundle(name AppearanceName, bundle foundation.Bundle) Appearance {
	result_ := C.C_NSAppearance_InitWithAppearanceNamed_Bundle(n.Ptr(), foundation.NewString(string(name)).Ptr(), objc.ExtractPtr(bundle))
	return MakeAppearance(result_)
}

func (n NSAppearance) InitWithCoder(coder foundation.Coder) Appearance {
	result_ := C.C_NSAppearance_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeAppearance(result_)
}

func (n NSAppearance) Init() Appearance {
	result_ := C.C_NSAppearance_Init(n.Ptr())
	return MakeAppearance(result_)
}

func AppearanceNamed(name AppearanceName) Appearance {
	result_ := C.C_NSAppearance_AppearanceNamed(foundation.NewString(string(name)).Ptr())
	return MakeAppearance(result_)
}

func (n NSAppearance) BestMatchFromAppearancesWithNames(appearances []AppearanceName) AppearanceName {
	cAppearancesData := make([]unsafe.Pointer, len(appearances))
	for idx, v := range appearances {
		cAppearancesData[idx] = foundation.NewString(string(v)).Ptr()
	}
	cAppearances := C.Array{data: unsafe.Pointer(&cAppearancesData[0]), len: C.int(len(appearances))}
	result_ := C.C_NSAppearance_BestMatchFromAppearancesWithNames(n.Ptr(), cAppearances)
	return AppearanceName(foundation.MakeString(result_).String())
}

func (n NSAppearance) Name() AppearanceName {
	result_ := C.C_NSAppearance_Name(n.Ptr())
	return AppearanceName(foundation.MakeString(result_).String())
}

func CurrentDrawingAppearance() Appearance {
	result_ := C.C_NSAppearance_CurrentDrawingAppearance()
	return MakeAppearance(result_)
}

func (n NSAppearance) AllowsVibrancy() bool {
	result_ := C.C_NSAppearance_AllowsVibrancy(n.Ptr())
	return bool(result_)
}
