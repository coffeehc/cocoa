package appkit

// #include "layout_anchor.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type LayoutAnchor interface {
	objc.Object
	ConstraintEqualToAnchor(anchor LayoutAnchor) LayoutConstraint
	ConstraintEqualToAnchor_Constant(anchor LayoutAnchor, c coregraphics.Float) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor_Constant(anchor LayoutAnchor, c coregraphics.Float) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor_Constant(anchor LayoutAnchor, c coregraphics.Float) LayoutConstraint
	ConstraintsAffectingLayout() []LayoutConstraint
	HasAmbiguousLayout() bool
	Name() string
	Item() objc.Object
}

type NSLayoutAnchor struct {
	objc.NSObject
}

func MakeLayoutAnchor(ptr unsafe.Pointer) NSLayoutAnchor {
	return NSLayoutAnchor{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocLayoutAnchor() NSLayoutAnchor {
	return MakeLayoutAnchor(C.C_LayoutAnchor_Alloc())
}

func (n NSLayoutAnchor) Init() LayoutAnchor {
	result_ := C.C_NSLayoutAnchor_Init(n.Ptr())
	return MakeLayoutAnchor(result_)
}

func (n NSLayoutAnchor) ConstraintEqualToAnchor(anchor LayoutAnchor) LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintEqualToAnchor(n.Ptr(), objc.ExtractPtr(anchor))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutAnchor) ConstraintEqualToAnchor_Constant(anchor LayoutAnchor, c coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintEqualToAnchor_Constant(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(c)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutAnchor) ConstraintGreaterThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintGreaterThanOrEqualToAnchor(n.Ptr(), objc.ExtractPtr(anchor))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutAnchor) ConstraintGreaterThanOrEqualToAnchor_Constant(anchor LayoutAnchor, c coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintGreaterThanOrEqualToAnchor_Constant(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(c)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutAnchor) ConstraintLessThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintLessThanOrEqualToAnchor(n.Ptr(), objc.ExtractPtr(anchor))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutAnchor) ConstraintLessThanOrEqualToAnchor_Constant(anchor LayoutAnchor, c coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintLessThanOrEqualToAnchor_Constant(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(c)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutAnchor) ConstraintsAffectingLayout() []LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintsAffectingLayout(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]LayoutConstraint, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeLayoutConstraint(r)
	}
	return goResult_
}

func (n NSLayoutAnchor) HasAmbiguousLayout() bool {
	result_ := C.C_NSLayoutAnchor_HasAmbiguousLayout(n.Ptr())
	return bool(result_)
}

func (n NSLayoutAnchor) Name() string {
	result_ := C.C_NSLayoutAnchor_Name(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSLayoutAnchor) Item() objc.Object {
	result_ := C.C_NSLayoutAnchor_Item(n.Ptr())
	return objc.MakeObject(result_)
}
