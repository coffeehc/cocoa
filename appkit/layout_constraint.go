package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "layout_constraint.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type LayoutConstraint interface {
	objc.Object
	IsActive() bool
	SetActive(active bool)
	FirstItem() objc.Object
	FirstAttribute() LayoutAttribute
	Relation() LayoutRelation
	SecondItem() objc.Object
	SecondAttribute() LayoutAttribute
	Multiplier() float64
	Constant() float64
	FirstAnchor() LayoutAnchor
	SecondAnchor() LayoutAnchor
	Priority() LayoutPriority
	SetPriority(priority LayoutPriority)
	Identifier() string
	SetIdentifier(identifier string)
	ShouldBeArchived() bool
	SetShouldBeArchived(shouldBeArchived bool)
}

var _ LayoutConstraint = (*NSLayoutConstraint)(nil)

type NSLayoutConstraint struct {
	objc.NSObject
}

func MakeLayoutConstraint(ptr unsafe.Pointer) *NSLayoutConstraint {
	if ptr == nil {
		return nil
	}
	return &NSLayoutConstraint{
		NSObject: *objc.MakeObject(ptr),
	}
}

func (l *NSLayoutConstraint) IsActive() bool {
	return bool(C.LayoutConstraint_IsActive(l.Ptr()))
}

func (l *NSLayoutConstraint) SetActive(active bool) {
	C.LayoutConstraint_SetActive(l.Ptr(), C.bool(active))
}

func (l *NSLayoutConstraint) FirstItem() objc.Object {
	return objc.MakeObject(C.LayoutConstraint_FirstItem(l.Ptr()))
}

func (l *NSLayoutConstraint) FirstAttribute() LayoutAttribute {
	return LayoutAttribute(C.LayoutConstraint_FirstAttribute(l.Ptr()))
}

func (l *NSLayoutConstraint) Relation() LayoutRelation {
	return LayoutRelation(C.LayoutConstraint_Relation(l.Ptr()))
}

func (l *NSLayoutConstraint) SecondItem() objc.Object {
	return objc.MakeObject(C.LayoutConstraint_SecondItem(l.Ptr()))
}

func (l *NSLayoutConstraint) SecondAttribute() LayoutAttribute {
	return LayoutAttribute(C.LayoutConstraint_SecondAttribute(l.Ptr()))
}

func (l *NSLayoutConstraint) Multiplier() float64 {
	return float64(C.LayoutConstraint_Multiplier(l.Ptr()))
}

func (l *NSLayoutConstraint) Constant() float64 {
	return float64(C.LayoutConstraint_Constant(l.Ptr()))
}

func (l *NSLayoutConstraint) FirstAnchor() LayoutAnchor {
	return MakeLayoutAnchor(C.LayoutConstraint_FirstAnchor(l.Ptr()))
}

func (l *NSLayoutConstraint) SecondAnchor() LayoutAnchor {
	return MakeLayoutAnchor(C.LayoutConstraint_SecondAnchor(l.Ptr()))
}

func (l *NSLayoutConstraint) Priority() LayoutPriority {
	return LayoutPriority(C.LayoutConstraint_Priority(l.Ptr()))
}

func (l *NSLayoutConstraint) SetPriority(priority LayoutPriority) {
	C.LayoutConstraint_SetPriority(l.Ptr(), C.float(priority))
}

func (l *NSLayoutConstraint) Identifier() string {
	return C.GoString(C.LayoutConstraint_Identifier(l.Ptr()))
}

func (l *NSLayoutConstraint) SetIdentifier(identifier string) {
	cIdentifier := C.CString(identifier)
	defer C.free(unsafe.Pointer(cIdentifier))
	C.LayoutConstraint_SetIdentifier(l.Ptr(), cIdentifier)
}

func (l *NSLayoutConstraint) ShouldBeArchived() bool {
	return bool(C.LayoutConstraint_ShouldBeArchived(l.Ptr()))
}

func (l *NSLayoutConstraint) SetShouldBeArchived(shouldBeArchived bool) {
	C.LayoutConstraint_SetShouldBeArchived(l.Ptr(), C.bool(shouldBeArchived))
}

func ConstraintWithItem(view1 objc.Object, attr1 LayoutAttribute, relation LayoutRelation, view2 objc.Object, attr2 LayoutAttribute, multiplier float64, c float64) {
	C.LayoutConstraint_ConstraintWithItem(toPointer(view1), C.long(attr1), C.long(relation), toPointer(view2), C.long(attr2), C.double(multiplier), C.double(c))
}

func ActivateConstraints(constraints []LayoutConstraint) {
	c_constraintsData := make([]unsafe.Pointer, len(constraints))
	for idx, v := range constraints {
		c_constraintsData[idx] = v.Ptr()
	}
	c_constraints := C.Array{data: unsafe.Pointer(&c_constraintsData[0]), len: C.int(len(constraints))}
	C.LayoutConstraint_ActivateConstraints(c_constraints)
}

func DeactivateConstraints(constraints []LayoutConstraint) {
	c_constraintsData := make([]unsafe.Pointer, len(constraints))
	for idx, v := range constraints {
		c_constraintsData[idx] = v.Ptr()
	}
	c_constraints := C.Array{data: unsafe.Pointer(&c_constraintsData[0]), len: C.int(len(constraints))}
	C.LayoutConstraint_DeactivateConstraints(c_constraints)
}
