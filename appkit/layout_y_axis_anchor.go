package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "layout_y_axis_anchor.h"
import "C"
import (
	"unsafe"
)

type LayoutYAxisAnchor interface {
	LayoutAnchor
	ConstraintEqualToSystemSpacingBelowAnchor(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintLessThanOrEqualToSystemSpacingBelowAnchor(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint
	AnchorWithOffsetToAnchor(otherAnchor LayoutYAxisAnchor) LayoutYAxisAnchor
}

var _ LayoutYAxisAnchor = (*NSLayoutYAxisAnchor)(nil)

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

func (l *NSLayoutYAxisAnchor) ConstraintEqualToSystemSpacingBelowAnchor(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutYAxisAnchor_ConstraintEqualToSystemSpacingBelowAnchor(l.Ptr(), toPointer(anchor), C.double(multiplier)))
}

func (l *NSLayoutYAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutYAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor(l.Ptr(), toPointer(anchor), C.double(multiplier)))
}

func (l *NSLayoutYAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingBelowAnchor(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutYAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingBelowAnchor(l.Ptr(), toPointer(anchor), C.double(multiplier)))
}

func (l *NSLayoutYAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor LayoutYAxisAnchor) LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.LayoutYAxisAnchor_AnchorWithOffsetToAnchor(l.Ptr(), toPointer(otherAnchor)))
}
