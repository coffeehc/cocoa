package appkit

// #include "layout_x_axis_anchor.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type LayoutXAxisAnchor interface {
	LayoutAnchor
	ConstraintEqualToSystemSpacingAfterAnchor_Multiplier(anchor LayoutXAxisAnchor, multiplier coregraphics.Float) LayoutConstraint
	ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor_Multiplier(anchor LayoutXAxisAnchor, multiplier coregraphics.Float) LayoutConstraint
	ConstraintLessThanOrEqualToSystemSpacingAfterAnchor_Multiplier(anchor LayoutXAxisAnchor, multiplier coregraphics.Float) LayoutConstraint
	AnchorWithOffsetToAnchor(otherAnchor LayoutXAxisAnchor) LayoutDimension
}

type NSLayoutXAxisAnchor struct {
	NSLayoutAnchor
}

func MakeLayoutXAxisAnchor(ptr unsafe.Pointer) NSLayoutXAxisAnchor {
	return NSLayoutXAxisAnchor{
		NSLayoutAnchor: MakeLayoutAnchor(ptr),
	}
}

func AllocLayoutXAxisAnchor() NSLayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.C_LayoutXAxisAnchor_Alloc())
}

func (n NSLayoutXAxisAnchor) Init() LayoutXAxisAnchor {
	result_ := C.C_NSLayoutXAxisAnchor_Init(n.Ptr())
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSLayoutXAxisAnchor) ConstraintEqualToSystemSpacingAfterAnchor_Multiplier(anchor LayoutXAxisAnchor, multiplier coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutXAxisAnchor_ConstraintEqualToSystemSpacingAfterAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(multiplier)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutXAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor_Multiplier(anchor LayoutXAxisAnchor, multiplier coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutXAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(multiplier)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutXAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingAfterAnchor_Multiplier(anchor LayoutXAxisAnchor, multiplier coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutXAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingAfterAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(multiplier)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutXAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor LayoutXAxisAnchor) LayoutDimension {
	result_ := C.C_NSLayoutXAxisAnchor_AnchorWithOffsetToAnchor(n.Ptr(), objc.ExtractPtr(otherAnchor))
	return MakeLayoutDimension(result_)
}
