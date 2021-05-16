package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "split_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type SplitView interface {
	View
	AddArrangedSubview(view View)
	InsertArrangedSubview_AtIndex(view View, index int)
	RemoveArrangedSubview(view View)
	AdjustSubviews()
	IsSubviewCollapsed(subview View) bool
	HoldingPriorityForSubviewAtIndex(subviewIndex int) LayoutPriority
	SetHoldingPriority_ForSubviewAtIndex(priority LayoutPriority, subviewIndex int)
	DrawDividerInRect(rect foundation.Rect)
	MinPossiblePositionOfDividerAtIndex(dividerIndex int) coregraphics.Float
	MaxPossiblePositionOfDividerAtIndex(dividerIndex int) coregraphics.Float
	SetPosition_OfDividerAtIndex(position coregraphics.Float, dividerIndex int)
	ArrangesAllSubviews() bool
	SetArrangesAllSubviews(value bool)
	ArrangedSubviews() []View
	IsVertical() bool
	SetVertical(value bool)
	DividerStyle() SplitViewDividerStyle
	SetDividerStyle(value SplitViewDividerStyle)
	DividerColor() Color
	DividerThickness() coregraphics.Float
	AutosaveName() SplitViewAutosaveName
	SetAutosaveName(value SplitViewAutosaveName)
}

type NSSplitView struct {
	NSView
}

func MakeSplitView(ptr unsafe.Pointer) *NSSplitView {
	if ptr == nil {
		return nil
	}
	return &NSSplitView{
		NSView: *MakeView(ptr),
	}
}

func AllocSplitView() *NSSplitView {
	return MakeSplitView(C.C_SplitView_Alloc())
}

func (n *NSSplitView) InitWithFrame(frameRect foundation.Rect) SplitView {
	result := C.C_NSSplitView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeSplitView(result)
}

func (n *NSSplitView) InitWithCoder(coder foundation.Coder) SplitView {
	result := C.C_NSSplitView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeSplitView(result)
}

func (n *NSSplitView) Init() SplitView {
	result := C.C_NSSplitView_Init(n.Ptr())
	return MakeSplitView(result)
}

func (n *NSSplitView) AddArrangedSubview(view View) {
	C.C_NSSplitView_AddArrangedSubview(n.Ptr(), objc.ExtractPtr(view))
}

func (n *NSSplitView) InsertArrangedSubview_AtIndex(view View, index int) {
	C.C_NSSplitView_InsertArrangedSubview_AtIndex(n.Ptr(), objc.ExtractPtr(view), C.int(index))
}

func (n *NSSplitView) RemoveArrangedSubview(view View) {
	C.C_NSSplitView_RemoveArrangedSubview(n.Ptr(), objc.ExtractPtr(view))
}

func (n *NSSplitView) AdjustSubviews() {
	C.C_NSSplitView_AdjustSubviews(n.Ptr())
}

func (n *NSSplitView) IsSubviewCollapsed(subview View) bool {
	result := C.C_NSSplitView_IsSubviewCollapsed(n.Ptr(), objc.ExtractPtr(subview))
	return bool(result)
}

func (n *NSSplitView) HoldingPriorityForSubviewAtIndex(subviewIndex int) LayoutPriority {
	result := C.C_NSSplitView_HoldingPriorityForSubviewAtIndex(n.Ptr(), C.int(subviewIndex))
	return LayoutPriority(float32(result))
}

func (n *NSSplitView) SetHoldingPriority_ForSubviewAtIndex(priority LayoutPriority, subviewIndex int) {
	C.C_NSSplitView_SetHoldingPriority_ForSubviewAtIndex(n.Ptr(), C.float(float32(priority)), C.int(subviewIndex))
}

func (n *NSSplitView) DrawDividerInRect(rect foundation.Rect) {
	C.C_NSSplitView_DrawDividerInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSSplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) coregraphics.Float {
	result := C.C_NSSplitView_MinPossiblePositionOfDividerAtIndex(n.Ptr(), C.int(dividerIndex))
	return coregraphics.Float(float64(result))
}

func (n *NSSplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) coregraphics.Float {
	result := C.C_NSSplitView_MaxPossiblePositionOfDividerAtIndex(n.Ptr(), C.int(dividerIndex))
	return coregraphics.Float(float64(result))
}

func (n *NSSplitView) SetPosition_OfDividerAtIndex(position coregraphics.Float, dividerIndex int) {
	C.C_NSSplitView_SetPosition_OfDividerAtIndex(n.Ptr(), C.double(float64(position)), C.int(dividerIndex))
}

func (n *NSSplitView) ArrangesAllSubviews() bool {
	result := C.C_NSSplitView_ArrangesAllSubviews(n.Ptr())
	return bool(result)
}

func (n *NSSplitView) SetArrangesAllSubviews(value bool) {
	C.C_NSSplitView_SetArrangesAllSubviews(n.Ptr(), C.bool(value))
}

func (n *NSSplitView) ArrangedSubviews() []View {
	result := C.C_NSSplitView_ArrangedSubviews(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]View, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeView(r)
	}
	return goResult
}

func (n *NSSplitView) IsVertical() bool {
	result := C.C_NSSplitView_IsVertical(n.Ptr())
	return bool(result)
}

func (n *NSSplitView) SetVertical(value bool) {
	C.C_NSSplitView_SetVertical(n.Ptr(), C.bool(value))
}

func (n *NSSplitView) DividerStyle() SplitViewDividerStyle {
	result := C.C_NSSplitView_DividerStyle(n.Ptr())
	return SplitViewDividerStyle(int(result))
}

func (n *NSSplitView) SetDividerStyle(value SplitViewDividerStyle) {
	C.C_NSSplitView_SetDividerStyle(n.Ptr(), C.int(int(value)))
}

func (n *NSSplitView) DividerColor() Color {
	result := C.C_NSSplitView_DividerColor(n.Ptr())
	return MakeColor(result)
}

func (n *NSSplitView) DividerThickness() coregraphics.Float {
	result := C.C_NSSplitView_DividerThickness(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSSplitView) AutosaveName() SplitViewAutosaveName {
	result := C.C_NSSplitView_AutosaveName(n.Ptr())
	return SplitViewAutosaveName(foundation.MakeString(result).String())
}

func (n *NSSplitView) SetAutosaveName(value SplitViewAutosaveName) {
	C.C_NSSplitView_SetAutosaveName(n.Ptr(), foundation.NewString(string(value)).Ptr())
}
