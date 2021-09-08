package appkit

// #include "layout_y_axis_anchor.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type LayoutYAxisAnchor interface {
	LayoutAnchor
	ConstraintEqualToSystemSpacingBelowAnchor_Multiplier(anchor LayoutYAxisAnchor, multiplier coregraphics.Float) LayoutConstraint
	ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor_Multiplier(anchor LayoutYAxisAnchor, multiplier coregraphics.Float) LayoutConstraint
	ConstraintLessThanOrEqualToSystemSpacingBelowAnchor_Multiplier(anchor LayoutYAxisAnchor, multiplier coregraphics.Float) LayoutConstraint
	AnchorWithOffsetToAnchor(otherAnchor LayoutYAxisAnchor) LayoutDimension
}

type NSLayoutYAxisAnchor struct {
	NSLayoutAnchor
}

func MakeLayoutYAxisAnchor(ptr unsafe.Pointer) NSLayoutYAxisAnchor {
	return NSLayoutYAxisAnchor{
		NSLayoutAnchor: MakeLayoutAnchor(ptr),
	}
}

func AllocLayoutYAxisAnchor() NSLayoutYAxisAnchor {
	result_ := C.C_NSLayoutYAxisAnchor_AllocLayoutYAxisAnchor()
	return MakeLayoutYAxisAnchor(result_)
}

func (n NSLayoutYAxisAnchor) Init() NSLayoutYAxisAnchor {
	result_ := C.C_NSLayoutYAxisAnchor_Init(n.Ptr())
	return MakeLayoutYAxisAnchor(result_)
}

func NewLayoutYAxisAnchor() NSLayoutYAxisAnchor {
	result_ := C.C_NSLayoutYAxisAnchor_NewLayoutYAxisAnchor()
	return MakeLayoutYAxisAnchor(result_)
}

func (n NSLayoutYAxisAnchor) Autorelease() NSLayoutYAxisAnchor {
	result_ := C.C_NSLayoutYAxisAnchor_Autorelease(n.Ptr())
	return MakeLayoutYAxisAnchor(result_)
}

func (n NSLayoutYAxisAnchor) Retain() NSLayoutYAxisAnchor {
	result_ := C.C_NSLayoutYAxisAnchor_Retain(n.Ptr())
	return MakeLayoutYAxisAnchor(result_)
}

func (n NSLayoutYAxisAnchor) ConstraintEqualToSystemSpacingBelowAnchor_Multiplier(anchor LayoutYAxisAnchor, multiplier coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutYAxisAnchor_ConstraintEqualToSystemSpacingBelowAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(multiplier)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutYAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor_Multiplier(anchor LayoutYAxisAnchor, multiplier coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutYAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(multiplier)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutYAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingBelowAnchor_Multiplier(anchor LayoutYAxisAnchor, multiplier coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutYAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingBelowAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(multiplier)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutYAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor LayoutYAxisAnchor) LayoutDimension {
	result_ := C.C_NSLayoutYAxisAnchor_AnchorWithOffsetToAnchor(n.Ptr(), objc.ExtractPtr(otherAnchor))
	return MakeLayoutDimension(result_)
}
