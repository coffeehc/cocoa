package appkit

// #include "layout_constraint.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type LayoutConstraint interface {
	objc.Object
	IsActive() bool
	SetActive(value bool)
	FirstItem() objc.Object
	FirstAttribute() LayoutAttribute
	Relation() LayoutRelation
	SecondItem() objc.Object
	SecondAttribute() LayoutAttribute
	Multiplier() coregraphics.Float
	Constant() coregraphics.Float
	SetConstant(value coregraphics.Float)
	FirstAnchor() LayoutAnchor
	SecondAnchor() LayoutAnchor
	Priority() LayoutPriority
	SetPriority(value LayoutPriority)
	Identifier() string
	SetIdentifier(value string)
	ShouldBeArchived() bool
	SetShouldBeArchived(value bool)
}

type NSLayoutConstraint struct {
	objc.NSObject
}

func MakeLayoutConstraint(ptr unsafe.Pointer) NSLayoutConstraint {
	return NSLayoutConstraint{
		NSObject: objc.MakeObject(ptr),
	}
}

func LayoutConstraint_ConstraintWithItem_Attribute_RelatedBy_ToItem_Attribute_Multiplier_Constant(view1 objc.Object, attr1 LayoutAttribute, relation LayoutRelation, view2 objc.Object, attr2 LayoutAttribute, multiplier coregraphics.Float, c coregraphics.Float) NSLayoutConstraint {
	result_ := C.C_NSLayoutConstraint_LayoutConstraint_ConstraintWithItem_Attribute_RelatedBy_ToItem_Attribute_Multiplier_Constant(objc.ExtractPtr(view1), C.int(int(attr1)), C.int(int(relation)), objc.ExtractPtr(view2), C.int(int(attr2)), C.double(float64(multiplier)), C.double(float64(c)))
	return MakeLayoutConstraint(result_)
}

func AllocLayoutConstraint() NSLayoutConstraint {
	result_ := C.C_NSLayoutConstraint_AllocLayoutConstraint()
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutConstraint) Init() NSLayoutConstraint {
	result_ := C.C_NSLayoutConstraint_Init(n.Ptr())
	return MakeLayoutConstraint(result_)
}

func NewLayoutConstraint() NSLayoutConstraint {
	result_ := C.C_NSLayoutConstraint_NewLayoutConstraint()
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutConstraint) Autorelease() NSLayoutConstraint {
	result_ := C.C_NSLayoutConstraint_Autorelease(n.Ptr())
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutConstraint) Retain() NSLayoutConstraint {
	result_ := C.C_NSLayoutConstraint_Retain(n.Ptr())
	return MakeLayoutConstraint(result_)
}

func LayoutConstraint_ConstraintsWithVisualFormat_Options_Metrics_Views(format string, opts LayoutFormatOptions, metrics map[string]objc.Object, views map[string]objc.Object) []LayoutConstraint {
	var cMetrics C.Dictionary
	if len(metrics) > 0 {
		cMetricsKeyData := make([]unsafe.Pointer, len(metrics))
		cMetricsValueData := make([]unsafe.Pointer, len(metrics))
		var idx = 0
		for k, v := range metrics {
			cMetricsKeyData[idx] = foundation.NewString(k).Ptr()
			cMetricsValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cMetrics.key_data = unsafe.Pointer(&cMetricsKeyData[0])
		cMetrics.value_data = unsafe.Pointer(&cMetricsValueData[0])
		cMetrics.len = C.int(len(metrics))
	}
	var cViews C.Dictionary
	if len(views) > 0 {
		cViewsKeyData := make([]unsafe.Pointer, len(views))
		cViewsValueData := make([]unsafe.Pointer, len(views))
		var idx = 0
		for k, v := range views {
			cViewsKeyData[idx] = foundation.NewString(k).Ptr()
			cViewsValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cViews.key_data = unsafe.Pointer(&cViewsKeyData[0])
		cViews.value_data = unsafe.Pointer(&cViewsValueData[0])
		cViews.len = C.int(len(views))
	}
	result_ := C.C_NSLayoutConstraint_LayoutConstraint_ConstraintsWithVisualFormat_Options_Metrics_Views(foundation.NewString(format).Ptr(), C.uint(uint(opts)), cMetrics, cViews)
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]LayoutConstraint, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeLayoutConstraint(r)
	}
	return goResult_
}

func LayoutConstraint_ActivateConstraints(constraints []LayoutConstraint) {
	var cConstraints C.Array
	if len(constraints) > 0 {
		cConstraintsData := make([]unsafe.Pointer, len(constraints))
		for idx, v := range constraints {
			cConstraintsData[idx] = objc.ExtractPtr(v)
		}
		cConstraints.data = unsafe.Pointer(&cConstraintsData[0])
		cConstraints.len = C.int(len(constraints))
	}
	C.C_NSLayoutConstraint_LayoutConstraint_ActivateConstraints(cConstraints)
}

func LayoutConstraint_DeactivateConstraints(constraints []LayoutConstraint) {
	var cConstraints C.Array
	if len(constraints) > 0 {
		cConstraintsData := make([]unsafe.Pointer, len(constraints))
		for idx, v := range constraints {
			cConstraintsData[idx] = objc.ExtractPtr(v)
		}
		cConstraints.data = unsafe.Pointer(&cConstraintsData[0])
		cConstraints.len = C.int(len(constraints))
	}
	C.C_NSLayoutConstraint_LayoutConstraint_DeactivateConstraints(cConstraints)
}

func (n NSLayoutConstraint) IsActive() bool {
	result_ := C.C_NSLayoutConstraint_IsActive(n.Ptr())
	return bool(result_)
}

func (n NSLayoutConstraint) SetActive(value bool) {
	C.C_NSLayoutConstraint_SetActive(n.Ptr(), C.bool(value))
}

func (n NSLayoutConstraint) FirstItem() objc.Object {
	result_ := C.C_NSLayoutConstraint_FirstItem(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSLayoutConstraint) FirstAttribute() LayoutAttribute {
	result_ := C.C_NSLayoutConstraint_FirstAttribute(n.Ptr())
	return LayoutAttribute(int(result_))
}

func (n NSLayoutConstraint) Relation() LayoutRelation {
	result_ := C.C_NSLayoutConstraint_Relation(n.Ptr())
	return LayoutRelation(int(result_))
}

func (n NSLayoutConstraint) SecondItem() objc.Object {
	result_ := C.C_NSLayoutConstraint_SecondItem(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSLayoutConstraint) SecondAttribute() LayoutAttribute {
	result_ := C.C_NSLayoutConstraint_SecondAttribute(n.Ptr())
	return LayoutAttribute(int(result_))
}

func (n NSLayoutConstraint) Multiplier() coregraphics.Float {
	result_ := C.C_NSLayoutConstraint_Multiplier(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSLayoutConstraint) Constant() coregraphics.Float {
	result_ := C.C_NSLayoutConstraint_Constant(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSLayoutConstraint) SetConstant(value coregraphics.Float) {
	C.C_NSLayoutConstraint_SetConstant(n.Ptr(), C.double(float64(value)))
}

func (n NSLayoutConstraint) FirstAnchor() LayoutAnchor {
	result_ := C.C_NSLayoutConstraint_FirstAnchor(n.Ptr())
	return MakeLayoutAnchor(result_)
}

func (n NSLayoutConstraint) SecondAnchor() LayoutAnchor {
	result_ := C.C_NSLayoutConstraint_SecondAnchor(n.Ptr())
	return MakeLayoutAnchor(result_)
}

func (n NSLayoutConstraint) Priority() LayoutPriority {
	result_ := C.C_NSLayoutConstraint_Priority(n.Ptr())
	return LayoutPriority(float32(result_))
}

func (n NSLayoutConstraint) SetPriority(value LayoutPriority) {
	C.C_NSLayoutConstraint_SetPriority(n.Ptr(), C.float(float32(value)))
}

func (n NSLayoutConstraint) Identifier() string {
	result_ := C.C_NSLayoutConstraint_Identifier(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSLayoutConstraint) SetIdentifier(value string) {
	C.C_NSLayoutConstraint_SetIdentifier(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSLayoutConstraint) ShouldBeArchived() bool {
	result_ := C.C_NSLayoutConstraint_ShouldBeArchived(n.Ptr())
	return bool(result_)
}

func (n NSLayoutConstraint) SetShouldBeArchived(value bool) {
	C.C_NSLayoutConstraint_SetShouldBeArchived(n.Ptr(), C.bool(value))
}
