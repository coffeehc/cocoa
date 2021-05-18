package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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

func MakeLayoutYAxisAnchor(ptr unsafe.Pointer) *NSLayoutYAxisAnchor {
	if ptr == nil {
		return nil
	}
	return &NSLayoutYAxisAnchor{
		NSLayoutAnchor: *MakeLayoutAnchor(ptr),
	}
}

func AllocLayoutYAxisAnchor() *NSLayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.C_LayoutYAxisAnchor_Alloc())
}

func (n *NSLayoutYAxisAnchor) Init() LayoutYAxisAnchor {
	result_ := C.C_NSLayoutYAxisAnchor_Init(n.Ptr())
	return MakeLayoutYAxisAnchor(result_)
}

func (n *NSLayoutYAxisAnchor) ConstraintEqualToSystemSpacingBelowAnchor_Multiplier(anchor LayoutYAxisAnchor, multiplier coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutYAxisAnchor_ConstraintEqualToSystemSpacingBelowAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(multiplier)))
	return MakeLayoutConstraint(result_)
}

func (n *NSLayoutYAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor_Multiplier(anchor LayoutYAxisAnchor, multiplier coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutYAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(multiplier)))
	return MakeLayoutConstraint(result_)
}

func (n *NSLayoutYAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingBelowAnchor_Multiplier(anchor LayoutYAxisAnchor, multiplier coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutYAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingBelowAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(multiplier)))
	return MakeLayoutConstraint(result_)
}

func (n *NSLayoutYAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor LayoutYAxisAnchor) LayoutDimension {
	result_ := C.C_NSLayoutYAxisAnchor_AnchorWithOffsetToAnchor(n.Ptr(), objc.ExtractPtr(otherAnchor))
	return MakeLayoutDimension(result_)
}
