package appkit

// #include "layout_x_axis_anchor.h"
import "C"
import (
	"unsafe"
)

type LayoutXAxisAnchor interface {
	LayoutAnchor
	ConstraintEqualToSystemSpacingAfterAnchor(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintLessThanOrEqualToSystemSpacingAfterAnchor(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint
	AnchorWithOffsetToAnchor(otherAnchor LayoutXAxisAnchor) LayoutDimension
}

var _ LayoutXAxisAnchor = (*NSLayoutXAxisAnchor)(nil)

type NSLayoutXAxisAnchor struct {
	NSLayoutAnchor
}

func MakeLayoutXAxisAnchor(ptr unsafe.Pointer) *NSLayoutXAxisAnchor {
	if ptr == nil {
		return nil
	}
	return &NSLayoutXAxisAnchor{
		NSLayoutAnchor: *MakeLayoutAnchor(ptr),
	}
}

func (l *NSLayoutXAxisAnchor) ConstraintEqualToSystemSpacingAfterAnchor(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutXAxisAnchor_ConstraintEqualToSystemSpacingAfterAnchor(l.Ptr(), toPointer(anchor), C.double(multiplier)))
}

func (l *NSLayoutXAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutXAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor(l.Ptr(), toPointer(anchor), C.double(multiplier)))
}

func (l *NSLayoutXAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingAfterAnchor(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutXAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingAfterAnchor(l.Ptr(), toPointer(anchor), C.double(multiplier)))
}

func (l *NSLayoutXAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor LayoutXAxisAnchor) LayoutDimension {
	return MakeLayoutDimension(C.LayoutXAxisAnchor_AnchorWithOffsetToAnchor(l.Ptr(), toPointer(otherAnchor)))
}
