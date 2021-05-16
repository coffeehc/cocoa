package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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

func MakeLayoutGuide(ptr unsafe.Pointer) *NSLayoutGuide {
	if ptr == nil {
		return nil
	}
	return &NSLayoutGuide{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocLayoutGuide() *NSLayoutGuide {
	return MakeLayoutGuide(C.C_LayoutGuide_Alloc())
}

func (n *NSLayoutGuide) Init() LayoutGuide {
	result := C.C_NSLayoutGuide_Init(n.Ptr())
	return MakeLayoutGuide(result)
}

func (n *NSLayoutGuide) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	result := C.C_NSLayoutGuide_ConstraintsAffectingLayoutForOrientation(n.Ptr(), C.int(int(orientation)))
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]LayoutConstraint, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeLayoutConstraint(r)
	}
	return goResult
}

func (n *NSLayoutGuide) Identifier() UserInterfaceItemIdentifier {
	result := C.C_NSLayoutGuide_Identifier(n.Ptr())
	return UserInterfaceItemIdentifier(foundation.MakeString(result).String())
}

func (n *NSLayoutGuide) SetIdentifier(value UserInterfaceItemIdentifier) {
	C.C_NSLayoutGuide_SetIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n *NSLayoutGuide) Frame() foundation.Rect {
	result := C.C_NSLayoutGuide_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSLayoutGuide) OwningView() View {
	result := C.C_NSLayoutGuide_OwningView(n.Ptr())
	return MakeView(result)
}

func (n *NSLayoutGuide) SetOwningView(value View) {
	C.C_NSLayoutGuide_SetOwningView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSLayoutGuide) BottomAnchor() LayoutYAxisAnchor {
	result := C.C_NSLayoutGuide_BottomAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result)
}

func (n *NSLayoutGuide) CenterXAnchor() LayoutXAxisAnchor {
	result := C.C_NSLayoutGuide_CenterXAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result)
}

func (n *NSLayoutGuide) CenterYAnchor() LayoutYAxisAnchor {
	result := C.C_NSLayoutGuide_CenterYAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result)
}

func (n *NSLayoutGuide) HeightAnchor() LayoutDimension {
	result := C.C_NSLayoutGuide_HeightAnchor(n.Ptr())
	return MakeLayoutDimension(result)
}

func (n *NSLayoutGuide) LeadingAnchor() LayoutXAxisAnchor {
	result := C.C_NSLayoutGuide_LeadingAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result)
}

func (n *NSLayoutGuide) LeftAnchor() LayoutXAxisAnchor {
	result := C.C_NSLayoutGuide_LeftAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result)
}

func (n *NSLayoutGuide) RightAnchor() LayoutXAxisAnchor {
	result := C.C_NSLayoutGuide_RightAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result)
}

func (n *NSLayoutGuide) TopAnchor() LayoutYAxisAnchor {
	result := C.C_NSLayoutGuide_TopAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result)
}

func (n *NSLayoutGuide) TrailingAnchor() LayoutXAxisAnchor {
	result := C.C_NSLayoutGuide_TrailingAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result)
}

func (n *NSLayoutGuide) WidthAnchor() LayoutDimension {
	result := C.C_NSLayoutGuide_WidthAnchor(n.Ptr())
	return MakeLayoutDimension(result)
}

func (n *NSLayoutGuide) HasAmbiguousLayout() bool {
	result := C.C_NSLayoutGuide_HasAmbiguousLayout(n.Ptr())
	return bool(result)
}
