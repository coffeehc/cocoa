package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "stack_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type StackView interface {
	View
	Views() []View
	DetachedViews() []View
	Orientation() UserInterfaceLayoutOrientation
	SetOrientation(orientation UserInterfaceLayoutOrientation)
	Alignment() LayoutAttribute
	SetAlignment(alignment LayoutAttribute)
	Spacing() float64
	SetSpacing(spacing float64)
	EdgeInsets() foundation.EdgeInsets
	SetEdgeInsets(edgeInsets foundation.EdgeInsets)
	Distribution() StackViewDistribution
	SetDistribution(distribution StackViewDistribution)
	ArrangedSubviews() []View
	DetachesHiddenViews() bool
	SetDetachesHiddenViews(detachesHiddenViews bool)
	ViewsInGravity(gravity StackViewGravity) []View
	AddView(view View, gravity StackViewGravity)
	InsertView(view View, index uint, gravity StackViewGravity)
	SetViews(views []View, gravity StackViewGravity)
	RemoveView(view View)
	AddArrangedSubview(view View)
	InsertArrangedSubview(view View, index uint)
	RemoveArrangedSubview(view View)
	ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	SetClippingResistancePriority(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation)
	HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	SetHuggingPriority(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation)
	CustomSpacingAfterView(view View) float64
	SetCustomSpacing(spacing float64, view View)
	VisibilityPriorityForView(view View) StackViewVisibilityPriority
	SetVisibilityPriority(priority StackViewVisibilityPriority, view View)
}

var _ StackView = (*NSStackView)(nil)

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

func (s *NSStackView) Views() []View {
	var cArray C.Array = C.StackView_Views(s.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]View, len(result))
	for idx, r := range result {
		goArray[idx] = MakeView(r)
	}
	return goArray
}

func (s *NSStackView) DetachedViews() []View {
	var cArray C.Array = C.StackView_DetachedViews(s.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]View, len(result))
	for idx, r := range result {
		goArray[idx] = MakeView(r)
	}
	return goArray
}

func (s *NSStackView) Orientation() UserInterfaceLayoutOrientation {
	return UserInterfaceLayoutOrientation(C.StackView_Orientation(s.Ptr()))
}

func (s *NSStackView) SetOrientation(orientation UserInterfaceLayoutOrientation) {
	C.StackView_SetOrientation(s.Ptr(), C.long(orientation))
}

func (s *NSStackView) Alignment() LayoutAttribute {
	return LayoutAttribute(C.StackView_Alignment(s.Ptr()))
}

func (s *NSStackView) SetAlignment(alignment LayoutAttribute) {
	C.StackView_SetAlignment(s.Ptr(), C.long(alignment))
}

func (s *NSStackView) Spacing() float64 {
	return float64(C.StackView_Spacing(s.Ptr()))
}

func (s *NSStackView) SetSpacing(spacing float64) {
	C.StackView_SetSpacing(s.Ptr(), C.double(spacing))
}

func (s *NSStackView) EdgeInsets() foundation.EdgeInsets {
	return toEdgeInsets(C.StackView_EdgeInsets(s.Ptr()))
}

func (s *NSStackView) SetEdgeInsets(edgeInsets foundation.EdgeInsets) {
	C.StackView_SetEdgeInsets(s.Ptr(), toNSEdgeInsets(edgeInsets))
}

func (s *NSStackView) Distribution() StackViewDistribution {
	return StackViewDistribution(C.StackView_Distribution(s.Ptr()))
}

func (s *NSStackView) SetDistribution(distribution StackViewDistribution) {
	C.StackView_SetDistribution(s.Ptr(), C.long(distribution))
}

func (s *NSStackView) ArrangedSubviews() []View {
	var cArray C.Array = C.StackView_ArrangedSubviews(s.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]View, len(result))
	for idx, r := range result {
		goArray[idx] = MakeView(r)
	}
	return goArray
}

func (s *NSStackView) DetachesHiddenViews() bool {
	return bool(C.StackView_DetachesHiddenViews(s.Ptr()))
}

func (s *NSStackView) SetDetachesHiddenViews(detachesHiddenViews bool) {
	C.StackView_SetDetachesHiddenViews(s.Ptr(), C.bool(detachesHiddenViews))
}

func StackViewWithViews(views []View) StackView {
	c_viewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		c_viewsData[idx] = v.Ptr()
	}
	c_views := C.Array{data: unsafe.Pointer(&c_viewsData[0]), len: C.int(len(views))}
	return MakeStackView(C.StackView_StackViewWithViews(c_views))
}

func (s *NSStackView) ViewsInGravity(gravity StackViewGravity) []View {
	var cArray C.Array = C.StackView_ViewsInGravity(s.Ptr(), C.long(gravity))
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]View, len(result))
	for idx, r := range result {
		goArray[idx] = MakeView(r)
	}
	return goArray
}

func (s *NSStackView) AddView(view View, gravity StackViewGravity) {
	C.StackView_AddView(s.Ptr(), toPointer(view), C.long(gravity))
}

func (s *NSStackView) InsertView(view View, index uint, gravity StackViewGravity) {
	C.StackView_InsertView(s.Ptr(), toPointer(view), C.ulong(index), C.long(gravity))
}

func (s *NSStackView) SetViews(views []View, gravity StackViewGravity) {
	c_viewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		c_viewsData[idx] = v.Ptr()
	}
	c_views := C.Array{data: unsafe.Pointer(&c_viewsData[0]), len: C.int(len(views))}
	C.StackView_SetViews(s.Ptr(), c_views, C.long(gravity))
}

func (s *NSStackView) RemoveView(view View) {
	C.StackView_RemoveView(s.Ptr(), toPointer(view))
}

func (s *NSStackView) AddArrangedSubview(view View) {
	C.StackView_AddArrangedSubview(s.Ptr(), toPointer(view))
}

func (s *NSStackView) InsertArrangedSubview(view View, index uint) {
	C.StackView_InsertArrangedSubview(s.Ptr(), toPointer(view), C.ulong(index))
}

func (s *NSStackView) RemoveArrangedSubview(view View) {
	C.StackView_RemoveArrangedSubview(s.Ptr(), toPointer(view))
}

func (s *NSStackView) ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	return LayoutPriority(C.StackView_ClippingResistancePriorityForOrientation(s.Ptr(), C.long(orientation)))
}

func (s *NSStackView) SetClippingResistancePriority(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation) {
	C.StackView_SetClippingResistancePriority(s.Ptr(), C.float(clippingResistancePriority), C.long(orientation))
}

func (s *NSStackView) HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	return LayoutPriority(C.StackView_HuggingPriorityForOrientation(s.Ptr(), C.long(orientation)))
}

func (s *NSStackView) SetHuggingPriority(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation) {
	C.StackView_SetHuggingPriority(s.Ptr(), C.float(huggingPriority), C.long(orientation))
}

func (s *NSStackView) CustomSpacingAfterView(view View) float64 {
	return float64(C.StackView_CustomSpacingAfterView(s.Ptr(), toPointer(view)))
}

func (s *NSStackView) SetCustomSpacing(spacing float64, view View) {
	C.StackView_SetCustomSpacing(s.Ptr(), C.double(spacing), toPointer(view))
}

func (s *NSStackView) VisibilityPriorityForView(view View) StackViewVisibilityPriority {
	return StackViewVisibilityPriority(C.StackView_VisibilityPriorityForView(s.Ptr(), toPointer(view)))
}

func (s *NSStackView) SetVisibilityPriority(priority StackViewVisibilityPriority, view View) {
	C.StackView_SetVisibilityPriority(s.Ptr(), C.float(priority), toPointer(view))
}
