package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "tab_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TabView interface {
	View
	AddTabViewItem(tabViewItem TabViewItem)
	InsertTabViewItem_AtIndex(tabViewItem TabViewItem, index int)
	RemoveTabViewItem(tabViewItem TabViewItem)
	IndexOfTabViewItem(tabViewItem TabViewItem) int
	IndexOfTabViewItemWithIdentifier(identifier objc.Object) int
	TabViewItemAtIndex(index int) TabViewItem
	SelectFirstTabViewItem(sender objc.Object)
	SelectLastTabViewItem(sender objc.Object)
	SelectNextTabViewItem(sender objc.Object)
	SelectPreviousTabViewItem(sender objc.Object)
	SelectTabViewItem(tabViewItem TabViewItem)
	SelectTabViewItemAtIndex(index int)
	SelectTabViewItemWithIdentifier(identifier objc.Object)
	TakeSelectedTabViewItemFromSender(sender objc.Object)
	TabViewItemAtPoint(point foundation.Point) TabViewItem
	NumberOfTabViewItems() int
	TabViewItems() []TabViewItem
	SetTabViewItems(value []TabViewItem)
	TabViewType() TabViewType
	SetTabViewType(value TabViewType)
	TabPosition() TabPosition
	SetTabPosition(value TabPosition)
	TabViewBorderType() TabViewBorderType
	SetTabViewBorderType(value TabViewBorderType)
	SelectedTabViewItem() TabViewItem
	Font() Font
	SetFont(value Font)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	MinimumSize() foundation.Size
	ContentRect() foundation.Rect
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	AllowsTruncatedLabels() bool
	SetAllowsTruncatedLabels(value bool)
}

type NSTabView struct {
	NSView
}

func MakeTabView(ptr unsafe.Pointer) *NSTabView {
	if ptr == nil {
		return nil
	}
	return &NSTabView{
		NSView: *MakeView(ptr),
	}
}

func AllocTabView() *NSTabView {
	return MakeTabView(C.C_TabView_Alloc())
}

func (n *NSTabView) InitWithFrame(frameRect foundation.Rect) TabView {
	result := C.C_NSTabView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeTabView(result)
}

func (n *NSTabView) InitWithCoder(coder foundation.Coder) TabView {
	result := C.C_NSTabView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTabView(result)
}

func (n *NSTabView) Init() TabView {
	result := C.C_NSTabView_Init(n.Ptr())
	return MakeTabView(result)
}

func (n *NSTabView) AddTabViewItem(tabViewItem TabViewItem) {
	C.C_NSTabView_AddTabViewItem(n.Ptr(), objc.ExtractPtr(tabViewItem))
}

func (n *NSTabView) InsertTabViewItem_AtIndex(tabViewItem TabViewItem, index int) {
	C.C_NSTabView_InsertTabViewItem_AtIndex(n.Ptr(), objc.ExtractPtr(tabViewItem), C.int(index))
}

func (n *NSTabView) RemoveTabViewItem(tabViewItem TabViewItem) {
	C.C_NSTabView_RemoveTabViewItem(n.Ptr(), objc.ExtractPtr(tabViewItem))
}

func (n *NSTabView) IndexOfTabViewItem(tabViewItem TabViewItem) int {
	result := C.C_NSTabView_IndexOfTabViewItem(n.Ptr(), objc.ExtractPtr(tabViewItem))
	return int(result)
}

func (n *NSTabView) IndexOfTabViewItemWithIdentifier(identifier objc.Object) int {
	result := C.C_NSTabView_IndexOfTabViewItemWithIdentifier(n.Ptr(), objc.ExtractPtr(identifier))
	return int(result)
}

func (n *NSTabView) TabViewItemAtIndex(index int) TabViewItem {
	result := C.C_NSTabView_TabViewItemAtIndex(n.Ptr(), C.int(index))
	return MakeTabViewItem(result)
}

func (n *NSTabView) SelectFirstTabViewItem(sender objc.Object) {
	C.C_NSTabView_SelectFirstTabViewItem(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSTabView) SelectLastTabViewItem(sender objc.Object) {
	C.C_NSTabView_SelectLastTabViewItem(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSTabView) SelectNextTabViewItem(sender objc.Object) {
	C.C_NSTabView_SelectNextTabViewItem(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSTabView) SelectPreviousTabViewItem(sender objc.Object) {
	C.C_NSTabView_SelectPreviousTabViewItem(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSTabView) SelectTabViewItem(tabViewItem TabViewItem) {
	C.C_NSTabView_SelectTabViewItem(n.Ptr(), objc.ExtractPtr(tabViewItem))
}

func (n *NSTabView) SelectTabViewItemAtIndex(index int) {
	C.C_NSTabView_SelectTabViewItemAtIndex(n.Ptr(), C.int(index))
}

func (n *NSTabView) SelectTabViewItemWithIdentifier(identifier objc.Object) {
	C.C_NSTabView_SelectTabViewItemWithIdentifier(n.Ptr(), objc.ExtractPtr(identifier))
}

func (n *NSTabView) TakeSelectedTabViewItemFromSender(sender objc.Object) {
	C.C_NSTabView_TakeSelectedTabViewItemFromSender(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSTabView) TabViewItemAtPoint(point foundation.Point) TabViewItem {
	result := C.C_NSTabView_TabViewItemAtPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return MakeTabViewItem(result)
}

func (n *NSTabView) NumberOfTabViewItems() int {
	result := C.C_NSTabView_NumberOfTabViewItems(n.Ptr())
	return int(result)
}

func (n *NSTabView) TabViewItems() []TabViewItem {
	result := C.C_NSTabView_TabViewItems(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]TabViewItem, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeTabViewItem(r)
	}
	return goResult
}

func (n *NSTabView) SetTabViewItems(value []TabViewItem) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSTabView_SetTabViewItems(n.Ptr(), cValue)
}

func (n *NSTabView) TabViewType() TabViewType {
	result := C.C_NSTabView_TabViewType(n.Ptr())
	return TabViewType(uint(result))
}

func (n *NSTabView) SetTabViewType(value TabViewType) {
	C.C_NSTabView_SetTabViewType(n.Ptr(), C.uint(uint(value)))
}

func (n *NSTabView) TabPosition() TabPosition {
	result := C.C_NSTabView_TabPosition(n.Ptr())
	return TabPosition(uint(result))
}

func (n *NSTabView) SetTabPosition(value TabPosition) {
	C.C_NSTabView_SetTabPosition(n.Ptr(), C.uint(uint(value)))
}

func (n *NSTabView) TabViewBorderType() TabViewBorderType {
	result := C.C_NSTabView_TabViewBorderType(n.Ptr())
	return TabViewBorderType(uint(result))
}

func (n *NSTabView) SetTabViewBorderType(value TabViewBorderType) {
	C.C_NSTabView_SetTabViewBorderType(n.Ptr(), C.uint(uint(value)))
}

func (n *NSTabView) SelectedTabViewItem() TabViewItem {
	result := C.C_NSTabView_SelectedTabViewItem(n.Ptr())
	return MakeTabViewItem(result)
}

func (n *NSTabView) Font() Font {
	result := C.C_NSTabView_Font(n.Ptr())
	return MakeFont(result)
}

func (n *NSTabView) SetFont(value Font) {
	C.C_NSTabView_SetFont(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSTabView) DrawsBackground() bool {
	result := C.C_NSTabView_DrawsBackground(n.Ptr())
	return bool(result)
}

func (n *NSTabView) SetDrawsBackground(value bool) {
	C.C_NSTabView_SetDrawsBackground(n.Ptr(), C.bool(value))
}

func (n *NSTabView) MinimumSize() foundation.Size {
	result := C.C_NSTabView_MinimumSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSTabView) ContentRect() foundation.Rect {
	result := C.C_NSTabView_ContentRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSTabView) ControlSize() ControlSize {
	result := C.C_NSTabView_ControlSize(n.Ptr())
	return ControlSize(uint(result))
}

func (n *NSTabView) SetControlSize(value ControlSize) {
	C.C_NSTabView_SetControlSize(n.Ptr(), C.uint(uint(value)))
}

func (n *NSTabView) AllowsTruncatedLabels() bool {
	result := C.C_NSTabView_AllowsTruncatedLabels(n.Ptr())
	return bool(result)
}

func (n *NSTabView) SetAllowsTruncatedLabels(value bool) {
	C.C_NSTabView_SetAllowsTruncatedLabels(n.Ptr(), C.bool(value))
}
