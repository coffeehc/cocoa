package appkit

// #include "layout_anchor.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type LayoutAnchor interface {
	objc.Object
	ConstraintEqualToAnchor(anchor LayoutAnchor) LayoutConstraint
	ConstraintEqualToAnchor_Constant(anchor LayoutAnchor, c coregraphics.Float) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor_Constant(anchor LayoutAnchor, c coregraphics.Float) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor_Constant(anchor LayoutAnchor, c coregraphics.Float) LayoutConstraint
	ConstraintsAffectingLayout() []LayoutConstraint
	HasAmbiguousLayout() bool
	Name() string
	Item() objc.Object
}

type NSLayoutAnchor struct {
	objc.NSObject
}

func MakeLayoutAnchor(ptr unsafe.Pointer) NSLayoutAnchor {
	return NSLayoutAnchor{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocLayoutAnchor() NSLayoutAnchor {
	result_ := C.C_NSLayoutAnchor_AllocLayoutAnchor()
	return MakeLayoutAnchor(result_)
}

func (n NSLayoutAnchor) Init() NSLayoutAnchor {
	result_ := C.C_NSLayoutAnchor_Init(n.Ptr())
	return MakeLayoutAnchor(result_)
}

func NewLayoutAnchor() NSLayoutAnchor {
	result_ := C.C_NSLayoutAnchor_NewLayoutAnchor()
	return MakeLayoutAnchor(result_)
}

func (n NSLayoutAnchor) Autorelease() NSLayoutAnchor {
	result_ := C.C_NSLayoutAnchor_Autorelease(n.Ptr())
	return MakeLayoutAnchor(result_)
}

func (n NSLayoutAnchor) Retain() NSLayoutAnchor {
	result_ := C.C_NSLayoutAnchor_Retain(n.Ptr())
	return MakeLayoutAnchor(result_)
}

func (n NSLayoutAnchor) ConstraintEqualToAnchor(anchor LayoutAnchor) LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintEqualToAnchor(n.Ptr(), objc.ExtractPtr(anchor))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutAnchor) ConstraintEqualToAnchor_Constant(anchor LayoutAnchor, c coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintEqualToAnchor_Constant(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(c)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutAnchor) ConstraintGreaterThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintGreaterThanOrEqualToAnchor(n.Ptr(), objc.ExtractPtr(anchor))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutAnchor) ConstraintGreaterThanOrEqualToAnchor_Constant(anchor LayoutAnchor, c coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintGreaterThanOrEqualToAnchor_Constant(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(c)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutAnchor) ConstraintLessThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintLessThanOrEqualToAnchor(n.Ptr(), objc.ExtractPtr(anchor))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutAnchor) ConstraintLessThanOrEqualToAnchor_Constant(anchor LayoutAnchor, c coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintLessThanOrEqualToAnchor_Constant(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(c)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutAnchor) ConstraintsAffectingLayout() []LayoutConstraint {
	result_ := C.C_NSLayoutAnchor_ConstraintsAffectingLayout(n.Ptr())
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

func (n NSLayoutAnchor) HasAmbiguousLayout() bool {
	result_ := C.C_NSLayoutAnchor_HasAmbiguousLayout(n.Ptr())
	return bool(result_)
}

func (n NSLayoutAnchor) Name() string {
	result_ := C.C_NSLayoutAnchor_Name(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSLayoutAnchor) Item() objc.Object {
	result_ := C.C_NSLayoutAnchor_Item(n.Ptr())
	return objc.MakeObject(result_)
}

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

func MakeLayoutDimension(ptr unsafe.Pointer) NSLayoutDimension {
	return NSLayoutDimension{
		NSLayoutAnchor: MakeLayoutAnchor(ptr),
	}
}

func AllocLayoutDimension() NSLayoutDimension {
	result_ := C.C_NSLayoutDimension_AllocLayoutDimension()
	return MakeLayoutDimension(result_)
}

func (n NSLayoutDimension) Init() NSLayoutDimension {
	result_ := C.C_NSLayoutDimension_Init(n.Ptr())
	return MakeLayoutDimension(result_)
}

func NewLayoutDimension() NSLayoutDimension {
	result_ := C.C_NSLayoutDimension_NewLayoutDimension()
	return MakeLayoutDimension(result_)
}

func (n NSLayoutDimension) Autorelease() NSLayoutDimension {
	result_ := C.C_NSLayoutDimension_Autorelease(n.Ptr())
	return MakeLayoutDimension(result_)
}

func (n NSLayoutDimension) Retain() NSLayoutDimension {
	result_ := C.C_NSLayoutDimension_Retain(n.Ptr())
	return MakeLayoutDimension(result_)
}

func (n NSLayoutDimension) ConstraintEqualToAnchor_Multiplier(anchor LayoutDimension, m coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutDimension_ConstraintEqualToAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(m)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutDimension) ConstraintEqualToAnchor_Multiplier_Constant(anchor LayoutDimension, m coregraphics.Float, c coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutDimension_ConstraintEqualToAnchor_Multiplier_Constant(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(m)), C.double(float64(c)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutDimension) ConstraintEqualToConstant(c coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutDimension_ConstraintEqualToConstant(n.Ptr(), C.double(float64(c)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutDimension) ConstraintGreaterThanOrEqualToAnchor_Multiplier(anchor LayoutDimension, m coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutDimension_ConstraintGreaterThanOrEqualToAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(m)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutDimension) ConstraintGreaterThanOrEqualToAnchor_Multiplier_Constant(anchor LayoutDimension, m coregraphics.Float, c coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutDimension_ConstraintGreaterThanOrEqualToAnchor_Multiplier_Constant(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(m)), C.double(float64(c)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutDimension) ConstraintGreaterThanOrEqualToConstant(c coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutDimension_ConstraintGreaterThanOrEqualToConstant(n.Ptr(), C.double(float64(c)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutDimension) ConstraintLessThanOrEqualToAnchor_Multiplier(anchor LayoutDimension, m coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutDimension_ConstraintLessThanOrEqualToAnchor_Multiplier(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(m)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutDimension) ConstraintLessThanOrEqualToAnchor_Multiplier_Constant(anchor LayoutDimension, m coregraphics.Float, c coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutDimension_ConstraintLessThanOrEqualToAnchor_Multiplier_Constant(n.Ptr(), objc.ExtractPtr(anchor), C.double(float64(m)), C.double(float64(c)))
	return MakeLayoutConstraint(result_)
}

func (n NSLayoutDimension) ConstraintLessThanOrEqualToConstant(c coregraphics.Float) LayoutConstraint {
	result_ := C.C_NSLayoutDimension_ConstraintLessThanOrEqualToConstant(n.Ptr(), C.double(float64(c)))
	return MakeLayoutConstraint(result_)
}

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
	result_ := C.C_NSLayoutXAxisAnchor_AllocLayoutXAxisAnchor()
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSLayoutXAxisAnchor) Init() NSLayoutXAxisAnchor {
	result_ := C.C_NSLayoutXAxisAnchor_Init(n.Ptr())
	return MakeLayoutXAxisAnchor(result_)
}

func NewLayoutXAxisAnchor() NSLayoutXAxisAnchor {
	result_ := C.C_NSLayoutXAxisAnchor_NewLayoutXAxisAnchor()
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSLayoutXAxisAnchor) Autorelease() NSLayoutXAxisAnchor {
	result_ := C.C_NSLayoutXAxisAnchor_Autorelease(n.Ptr())
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSLayoutXAxisAnchor) Retain() NSLayoutXAxisAnchor {
	result_ := C.C_NSLayoutXAxisAnchor_Retain(n.Ptr())
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
