package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "layout_dimension.h"
import "C"
import (
	"unsafe"
)

type LayoutDimension interface {
	LayoutAnchor
	ConstraintEqualToConstant(constant float64) LayoutConstraint
	ConstraintEqualToAnchor3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToConstant(constant float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint
	ConstraintLessThanOrEqualToConstant(constant float64) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint
}

var _ LayoutDimension = (*NSLayoutDimension)(nil)

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

func (l *NSLayoutDimension) ConstraintEqualToConstant(constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintEqualToConstant(l.Ptr(), C.double(constant)))
}

func (l *NSLayoutDimension) ConstraintEqualToAnchor3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintEqualToAnchor3(l.Ptr(), toPointer(anchor), C.double(multiplier), C.double(constant)))
}

func (l *NSLayoutDimension) ConstraintGreaterThanOrEqualToConstant(constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintGreaterThanOrEqualToConstant(l.Ptr(), C.double(constant)))
}

func (l *NSLayoutDimension) ConstraintGreaterThanOrEqualToAnchor3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintGreaterThanOrEqualToAnchor3(l.Ptr(), toPointer(anchor), C.double(multiplier), C.double(constant)))
}

func (l *NSLayoutDimension) ConstraintLessThanOrEqualToConstant(constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintLessThanOrEqualToConstant(l.Ptr(), C.double(constant)))
}

func (l *NSLayoutDimension) ConstraintLessThanOrEqualToAnchor3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintLessThanOrEqualToAnchor3(l.Ptr(), toPointer(anchor), C.double(multiplier), C.double(constant)))
}
