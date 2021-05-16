package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "stack_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type StackView interface {
	View
	AddView_InGravity(view View, gravity StackViewGravity)
	InsertView_AtIndex_InGravity(view View, index uint, gravity StackViewGravity)
	SetViews_InGravity(views []View, gravity StackViewGravity)
	RemoveView(view View)
	AddArrangedSubview(view View)
	InsertArrangedSubview_AtIndex(view View, index int)
	RemoveArrangedSubview(view View)
	ViewsInGravity(gravity StackViewGravity) []View
	ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	CustomSpacingAfterView(view View) coregraphics.Float
	SetCustomSpacing_AfterView(spacing coregraphics.Float, view View)
	VisibilityPriorityForView(view View) StackViewVisibilityPriority
	SetVisibilityPriority_ForView(priority StackViewVisibilityPriority, view View)
	SetClippingResistancePriority_ForOrientation(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation)
	SetHuggingPriority_ForOrientation(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation)
	ArrangedSubviews() []View
	Views() []View
	DetachedViews() []View
	Orientation() UserInterfaceLayoutOrientation
	SetOrientation(value UserInterfaceLayoutOrientation)
	Alignment() LayoutAttribute
	SetAlignment(value LayoutAttribute)
	Spacing() coregraphics.Float
	SetSpacing(value coregraphics.Float)
	EdgeInsets() foundation.EdgeInsets
	SetEdgeInsets(value foundation.EdgeInsets)
	Distribution() StackViewDistribution
	SetDistribution(value StackViewDistribution)
	DetachesHiddenViews() bool
	SetDetachesHiddenViews(value bool)
}

type NSStackView struct {
	NSView
}

func MakeStackView(ptr unsafe.Pointer) *NSStackView {
	if ptr == nil {
		return nil
	}
	return &NSStackView{
		NSView: *MakeView(ptr),
	}
}

func AllocStackView() *NSStackView {
	return MakeStackView(C.C_StackView_Alloc())
}

func (n *NSStackView) InitWithFrame(frameRect foundation.Rect) StackView {
	result := C.C_NSStackView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeStackView(result)
}

func (n *NSStackView) InitWithCoder(coder foundation.Coder) StackView {
	result := C.C_NSStackView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeStackView(result)
}

func (n *NSStackView) Init() StackView {
	result := C.C_NSStackView_Init(n.Ptr())
	return MakeStackView(result)
}

func StackViewWithViews(views []View) StackView {
	cViewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		cViewsData[idx] = objc.ExtractPtr(v)
	}
	cViews := C.Array{data: unsafe.Pointer(&cViewsData[0]), len: C.int(len(views))}
	result := C.C_NSStackView_StackViewWithViews(cViews)
	return MakeStackView(result)
}

func (n *NSStackView) AddView_InGravity(view View, gravity StackViewGravity) {
	C.C_NSStackView_AddView_InGravity(n.Ptr(), objc.ExtractPtr(view), C.int(int(gravity)))
}

func (n *NSStackView) InsertView_AtIndex_InGravity(view View, index uint, gravity StackViewGravity) {
	C.C_NSStackView_InsertView_AtIndex_InGravity(n.Ptr(), objc.ExtractPtr(view), C.uint(index), C.int(int(gravity)))
}

func (n *NSStackView) SetViews_InGravity(views []View, gravity StackViewGravity) {
	cViewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		cViewsData[idx] = objc.ExtractPtr(v)
	}
	cViews := C.Array{data: unsafe.Pointer(&cViewsData[0]), len: C.int(len(views))}
	C.C_NSStackView_SetViews_InGravity(n.Ptr(), cViews, C.int(int(gravity)))
}

func (n *NSStackView) RemoveView(view View) {
	C.C_NSStackView_RemoveView(n.Ptr(), objc.ExtractPtr(view))
}

func (n *NSStackView) AddArrangedSubview(view View) {
	C.C_NSStackView_AddArrangedSubview(n.Ptr(), objc.ExtractPtr(view))
}

func (n *NSStackView) InsertArrangedSubview_AtIndex(view View, index int) {
	C.C_NSStackView_InsertArrangedSubview_AtIndex(n.Ptr(), objc.ExtractPtr(view), C.int(index))
}

func (n *NSStackView) RemoveArrangedSubview(view View) {
	C.C_NSStackView_RemoveArrangedSubview(n.Ptr(), objc.ExtractPtr(view))
}

func (n *NSStackView) ViewsInGravity(gravity StackViewGravity) []View {
	result := C.C_NSStackView_ViewsInGravity(n.Ptr(), C.int(int(gravity)))
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]View, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeView(r)
	}
	return goResult
}

func (n *NSStackView) ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	result := C.C_NSStackView_ClippingResistancePriorityForOrientation(n.Ptr(), C.int(int(orientation)))
	return LayoutPriority(float32(result))
}

func (n *NSStackView) HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	result := C.C_NSStackView_HuggingPriorityForOrientation(n.Ptr(), C.int(int(orientation)))
	return LayoutPriority(float32(result))
}

func (n *NSStackView) CustomSpacingAfterView(view View) coregraphics.Float {
	result := C.C_NSStackView_CustomSpacingAfterView(n.Ptr(), objc.ExtractPtr(view))
	return coregraphics.Float(float64(result))
}

func (n *NSStackView) SetCustomSpacing_AfterView(spacing coregraphics.Float, view View) {
	C.C_NSStackView_SetCustomSpacing_AfterView(n.Ptr(), C.double(float64(spacing)), objc.ExtractPtr(view))
}

func (n *NSStackView) VisibilityPriorityForView(view View) StackViewVisibilityPriority {
	result := C.C_NSStackView_VisibilityPriorityForView(n.Ptr(), objc.ExtractPtr(view))
	return StackViewVisibilityPriority(float32(result))
}

func (n *NSStackView) SetVisibilityPriority_ForView(priority StackViewVisibilityPriority, view View) {
	C.C_NSStackView_SetVisibilityPriority_ForView(n.Ptr(), C.float(float32(priority)), objc.ExtractPtr(view))
}

func (n *NSStackView) SetClippingResistancePriority_ForOrientation(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation) {
	C.C_NSStackView_SetClippingResistancePriority_ForOrientation(n.Ptr(), C.float(float32(clippingResistancePriority)), C.int(int(orientation)))
}

func (n *NSStackView) SetHuggingPriority_ForOrientation(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation) {
	C.C_NSStackView_SetHuggingPriority_ForOrientation(n.Ptr(), C.float(float32(huggingPriority)), C.int(int(orientation)))
}

func (n *NSStackView) ArrangedSubviews() []View {
	result := C.C_NSStackView_ArrangedSubviews(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]View, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeView(r)
	}
	return goResult
}

func (n *NSStackView) Views() []View {
	result := C.C_NSStackView_Views(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]View, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeView(r)
	}
	return goResult
}

func (n *NSStackView) DetachedViews() []View {
	result := C.C_NSStackView_DetachedViews(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]View, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeView(r)
	}
	return goResult
}

func (n *NSStackView) Orientation() UserInterfaceLayoutOrientation {
	result := C.C_NSStackView_Orientation(n.Ptr())
	return UserInterfaceLayoutOrientation(int(result))
}

func (n *NSStackView) SetOrientation(value UserInterfaceLayoutOrientation) {
	C.C_NSStackView_SetOrientation(n.Ptr(), C.int(int(value)))
}

func (n *NSStackView) Alignment() LayoutAttribute {
	result := C.C_NSStackView_Alignment(n.Ptr())
	return LayoutAttribute(int(result))
}

func (n *NSStackView) SetAlignment(value LayoutAttribute) {
	C.C_NSStackView_SetAlignment(n.Ptr(), C.int(int(value)))
}

func (n *NSStackView) Spacing() coregraphics.Float {
	result := C.C_NSStackView_Spacing(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSStackView) SetSpacing(value coregraphics.Float) {
	C.C_NSStackView_SetSpacing(n.Ptr(), C.double(float64(value)))
}

func (n *NSStackView) EdgeInsets() foundation.EdgeInsets {
	result := C.C_NSStackView_EdgeInsets(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result))
}

func (n *NSStackView) SetEdgeInsets(value foundation.EdgeInsets) {
	C.C_NSStackView_SetEdgeInsets(n.Ptr(), *(*C.NSEdgeInsets)(foundation.ToNSEdgeInsetsPointer(value)))
}

func (n *NSStackView) Distribution() StackViewDistribution {
	result := C.C_NSStackView_Distribution(n.Ptr())
	return StackViewDistribution(int(result))
}

func (n *NSStackView) SetDistribution(value StackViewDistribution) {
	C.C_NSStackView_SetDistribution(n.Ptr(), C.int(int(value)))
}

func (n *NSStackView) DetachesHiddenViews() bool {
	result := C.C_NSStackView_DetachesHiddenViews(n.Ptr())
	return bool(result)
}

func (n *NSStackView) SetDetachesHiddenViews(value bool) {
	C.C_NSStackView_SetDetachesHiddenViews(n.Ptr(), C.bool(value))
}
