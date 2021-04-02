package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "layout_anchor.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type LayoutAnchor interface {
	objc.Object
	ConstraintsAffectingLayout() []LayoutConstraint
	HasAmbiguousLayout() bool
	Name() string
	Item() objc.Object
	ConstraintEqualToAnchor(anchor LayoutAnchor) LayoutConstraint
	ConstraintEqualToAnchor2(anchor LayoutAnchor, constant float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor2(anchor LayoutAnchor, constant float64) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor2(anchor LayoutAnchor, constant float64) LayoutConstraint
}

var _ LayoutAnchor = (*NSLayoutAnchor)(nil)

type NSLayoutAnchor struct {
	objc.NSObject
}

func MakeLayoutAnchor(ptr unsafe.Pointer) *NSLayoutAnchor {
	if ptr == nil {
		return nil
	}
	return &NSLayoutAnchor{
		NSObject: *objc.MakeObject(ptr),
	}
}

func (l *NSLayoutAnchor) ConstraintsAffectingLayout() []LayoutConstraint {
	var cArray C.Array = C.LayoutAnchor_ConstraintsAffectingLayout(l.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]LayoutConstraint, len(result))
	for idx, r := range result {
		goArray[idx] = MakeLayoutConstraint(r)
	}
	return goArray
}

func (l *NSLayoutAnchor) HasAmbiguousLayout() bool {
	return bool(C.LayoutAnchor_HasAmbiguousLayout(l.Ptr()))
}

func (l *NSLayoutAnchor) Name() string {
	return C.GoString(C.LayoutAnchor_Name(l.Ptr()))
}

func (l *NSLayoutAnchor) Item() objc.Object {
	return objc.MakeObject(C.LayoutAnchor_Item(l.Ptr()))
}

func (l *NSLayoutAnchor) ConstraintEqualToAnchor(anchor LayoutAnchor) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintEqualToAnchor(l.Ptr(), toPointer(anchor)))
}

func (l *NSLayoutAnchor) ConstraintEqualToAnchor2(anchor LayoutAnchor, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintEqualToAnchor2(l.Ptr(), toPointer(anchor), C.double(constant)))
}

func (l *NSLayoutAnchor) ConstraintGreaterThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintGreaterThanOrEqualToAnchor(l.Ptr(), toPointer(anchor)))
}

func (l *NSLayoutAnchor) ConstraintGreaterThanOrEqualToAnchor2(anchor LayoutAnchor, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintGreaterThanOrEqualToAnchor2(l.Ptr(), toPointer(anchor), C.double(constant)))
}

func (l *NSLayoutAnchor) ConstraintLessThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintLessThanOrEqualToAnchor(l.Ptr(), toPointer(anchor)))
}

func (l *NSLayoutAnchor) ConstraintLessThanOrEqualToAnchor2(anchor LayoutAnchor, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintLessThanOrEqualToAnchor2(l.Ptr(), toPointer(anchor), C.double(constant)))
}
