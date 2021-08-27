package appkit

// #include "layout_guide.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type LayoutGuide interface {
	objc.Object
	ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint
	Identifier() UserInterfaceItemIdentifier
	SetIdentifier(value UserInterfaceItemIdentifier)
	Frame() foundation.Rect
	OwningView() View
	SetOwningView(value View)
	BottomAnchor() LayoutYAxisAnchor
	CenterXAnchor() LayoutXAxisAnchor
	CenterYAnchor() LayoutYAxisAnchor
	HeightAnchor() LayoutDimension
	LeadingAnchor() LayoutXAxisAnchor
	LeftAnchor() LayoutXAxisAnchor
	RightAnchor() LayoutXAxisAnchor
	TopAnchor() LayoutYAxisAnchor
	TrailingAnchor() LayoutXAxisAnchor
	WidthAnchor() LayoutDimension
	HasAmbiguousLayout() bool
}

type NSLayoutGuide struct {
	objc.NSObject
}

func MakeLayoutGuide(ptr unsafe.Pointer) NSLayoutGuide {
	return NSLayoutGuide{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocLayoutGuide() NSLayoutGuide {
	return MakeLayoutGuide(C.C_LayoutGuide_Alloc())
}

func (n NSLayoutGuide) Init() LayoutGuide {
	result_ := C.C_NSLayoutGuide_Init(n.Ptr())
	return MakeLayoutGuide(result_)
}

func (n NSLayoutGuide) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	result_ := C.C_NSLayoutGuide_ConstraintsAffectingLayoutForOrientation(n.Ptr(), C.int(int(orientation)))
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

func (n NSLayoutGuide) Identifier() UserInterfaceItemIdentifier {
	result_ := C.C_NSLayoutGuide_Identifier(n.Ptr())
	return UserInterfaceItemIdentifier(foundation.MakeString(result_).String())
}

func (n NSLayoutGuide) SetIdentifier(value UserInterfaceItemIdentifier) {
	C.C_NSLayoutGuide_SetIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSLayoutGuide) Frame() foundation.Rect {
	result_ := C.C_NSLayoutGuide_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSLayoutGuide) OwningView() View {
	result_ := C.C_NSLayoutGuide_OwningView(n.Ptr())
	return MakeView(result_)
}

func (n NSLayoutGuide) SetOwningView(value View) {
	C.C_NSLayoutGuide_SetOwningView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSLayoutGuide) BottomAnchor() LayoutYAxisAnchor {
	result_ := C.C_NSLayoutGuide_BottomAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result_)
}

func (n NSLayoutGuide) CenterXAnchor() LayoutXAxisAnchor {
	result_ := C.C_NSLayoutGuide_CenterXAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSLayoutGuide) CenterYAnchor() LayoutYAxisAnchor {
	result_ := C.C_NSLayoutGuide_CenterYAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result_)
}

func (n NSLayoutGuide) HeightAnchor() LayoutDimension {
	result_ := C.C_NSLayoutGuide_HeightAnchor(n.Ptr())
	return MakeLayoutDimension(result_)
}

func (n NSLayoutGuide) LeadingAnchor() LayoutXAxisAnchor {
	result_ := C.C_NSLayoutGuide_LeadingAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSLayoutGuide) LeftAnchor() LayoutXAxisAnchor {
	result_ := C.C_NSLayoutGuide_LeftAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSLayoutGuide) RightAnchor() LayoutXAxisAnchor {
	result_ := C.C_NSLayoutGuide_RightAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSLayoutGuide) TopAnchor() LayoutYAxisAnchor {
	result_ := C.C_NSLayoutGuide_TopAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result_)
}

func (n NSLayoutGuide) TrailingAnchor() LayoutXAxisAnchor {
	result_ := C.C_NSLayoutGuide_TrailingAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSLayoutGuide) WidthAnchor() LayoutDimension {
	result_ := C.C_NSLayoutGuide_WidthAnchor(n.Ptr())
	return MakeLayoutDimension(result_)
}

func (n NSLayoutGuide) HasAmbiguousLayout() bool {
	result_ := C.C_NSLayoutGuide_HasAmbiguousLayout(n.Ptr())
	return bool(result_)
}
