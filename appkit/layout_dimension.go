package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "layout_dimension.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type LayoutDimension interface {
	LayoutAnchor
	ConstraintEqualToAnchor_Multiplier(anchor LayoutDimension, m coregraphics.Float) LayoutConstraint
	ConstraintEqualToAnchor_Multiplier_Constant(anchor LayoutDimension, m coregraphics.Float, c coregraphics.Float) LayoutConstraint
	ConstraintEqualToConstant(c coregraphics.Float) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor_Multiplier(anchor LayoutDimension, m coregraphics.Float) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor_Multiplier_Constant(anchor LayoutDimension, m coregraphics.Float, c coregraphics.Float) LayoutConstraint
	ConstraintGreaterThanOrEqualToConstant(c coregraphics.Float) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor_Multiplier(anchor LayoutDimension, m coregraphics.Float) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor_Multiplier_Constant(anchor LayoutDimension, m coregraphics.Float, c coregraphics.Float) LayoutConstraint
	ConstraintLessThanOrEqualToConstant(c coregraphics.Float) LayoutConstraint
}

type NSLayoutDimension struct {
	NSLayoutAnchor
}

func MakeLayoutDimension(ptr unsafe.Pointer) *NSLayoutDimension {
	if ptr == nil {
		return nil
	}
	return &NSLayoutDimension{
		NSLayoutAnchor: *MakeLayoutAnchor(ptr),
	}
}

func AllocLayoutDimension() *NSLayoutDimension {
	return MakeLayoutDimension(C.C_LayoutDimension_Alloc())
}

func (n *NSLayoutDimension) Init() LayoutDimension {
	result := C.C_NSLayoutDimension_Init(n.Ptr())
	return MakeLayoutDimension(result)
}

func (n *NSLayoutDimension) ConstraintEqualToAnchor_Multiplier(anchor LayoutDimension, m coregraphics.Float) LayoutConstraint {
	result := C.C_NSLayoutDimension_ConstraintEqualToAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(m)))
	return MakeLayoutConstraint(result)
}

func (n *NSLayoutDimension) ConstraintEqualToAnchor_Multiplier_Constant(anchor LayoutDimension, m coregraphics.Float, c coregraphics.Float) LayoutConstraint {
	result := C.C_NSLayoutDimension_ConstraintEqualToAnchor_Multiplier_Constant(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(m)), C.double(float64(c)))
	return MakeLayoutConstraint(result)
}

func (n *NSLayoutDimension) ConstraintEqualToConstant(c coregraphics.Float) LayoutConstraint {
	result := C.C_NSLayoutDimension_ConstraintEqualToConstant(n.Ptr(), C.double(float64(c)))
	return MakeLayoutConstraint(result)
}

func (n *NSLayoutDimension) ConstraintGreaterThanOrEqualToAnchor_Multiplier(anchor LayoutDimension, m coregraphics.Float) LayoutConstraint {
	result := C.C_NSLayoutDimension_ConstraintGreaterThanOrEqualToAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(m)))
	return MakeLayoutConstraint(result)
}

func (n *NSLayoutDimension) ConstraintGreaterThanOrEqualToAnchor_Multiplier_Constant(anchor LayoutDimension, m coregraphics.Float, c coregraphics.Float) LayoutConstraint {
	result := C.C_NSLayoutDimension_ConstraintGreaterThanOrEqualToAnchor_Multiplier_Constant(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(m)), C.double(float64(c)))
	return MakeLayoutConstraint(result)
}

func (n *NSLayoutDimension) ConstraintGreaterThanOrEqualToConstant(c coregraphics.Float) LayoutConstraint {
	result := C.C_NSLayoutDimension_ConstraintGreaterThanOrEqualToConstant(n.Ptr(), C.double(float64(c)))
	return MakeLayoutConstraint(result)
}

func (n *NSLayoutDimension) ConstraintLessThanOrEqualToAnchor_Multiplier(anchor LayoutDimension, m coregraphics.Float) LayoutConstraint {
	result := C.C_NSLayoutDimension_ConstraintLessThanOrEqualToAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(m)))
	return MakeLayoutConstraint(result)
}

func (n *NSLayoutDimension) ConstraintLessThanOrEqualToAnchor_Multiplier_Constant(anchor LayoutDimension, m coregraphics.Float, c coregraphics.Float) LayoutConstraint {
	result := C.C_NSLayoutDimension_ConstraintLessThanOrEqualToAnchor_Multiplier_Constant(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(m)), C.double(float64(c)))
	return MakeLayoutConstraint(result)
}

func (n *NSLayoutDimension) ConstraintLessThanOrEqualToConstant(c coregraphics.Float) LayoutConstraint {
	result := C.C_NSLayoutDimension_ConstraintLessThanOrEqualToConstant(n.Ptr(), C.double(float64(c)))
	return MakeLayoutConstraint(result)
}
