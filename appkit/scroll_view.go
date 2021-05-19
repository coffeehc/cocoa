package appkit

// #include "scroll_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ScrollView interface {
	View
	AddFloatingSubview_ForAxis(view View, axis EventGestureAxis)
	Tile()
	FlashScrollers()
	MagnifyToFitRect(rect foundation.Rect)
	SetMagnification_CenteredAtPoint(magnification coregraphics.Float, point foundation.Point)
	ContentSize() foundation.Size
	DocumentVisibleRect() foundation.Rect
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	BorderType() BorderType
	SetBorderType(value BorderType)
	DocumentCursor() Cursor
	SetDocumentCursor(value Cursor)
	ContentView() ClipView
	SetContentView(value ClipView)
	DocumentView() View
	SetDocumentView(value View)
	HorizontalScroller() Scroller
	SetHorizontalScroller(value Scroller)
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
	VerticalScroller() Scroller
	SetVerticalScroller(value Scroller)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	AutohidesScrollers() bool
	SetAutohidesScrollers(value bool)
	HasHorizontalRuler() bool
	SetHasHorizontalRuler(value bool)
	HorizontalRulerView() RulerView
	SetHorizontalRulerView(value RulerView)
	HasVerticalRuler() bool
	SetHasVerticalRuler(value bool)
	VerticalRulerView() RulerView
	SetVerticalRulerView(value RulerView)
	RulersVisible() bool
	SetRulersVisible(value bool)
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(value bool)
	ContentInsets() foundation.EdgeInsets
	SetContentInsets(value foundation.EdgeInsets)
	ScrollerInsets() foundation.EdgeInsets
	SetScrollerInsets(value foundation.EdgeInsets)
	ScrollerKnobStyle() ScrollerKnobStyle
	SetScrollerKnobStyle(value ScrollerKnobStyle)
	ScrollerStyle() ScrollerStyle
	SetScrollerStyle(value ScrollerStyle)
	LineScroll() coregraphics.Float
	SetLineScroll(value coregraphics.Float)
	HorizontalLineScroll() coregraphics.Float
	SetHorizontalLineScroll(value coregraphics.Float)
	VerticalLineScroll() coregraphics.Float
	SetVerticalLineScroll(value coregraphics.Float)
	PageScroll() coregraphics.Float
	SetPageScroll(value coregraphics.Float)
	HorizontalPageScroll() coregraphics.Float
	SetHorizontalPageScroll(value coregraphics.Float)
	VerticalPageScroll() coregraphics.Float
	SetVerticalPageScroll(value coregraphics.Float)
	ScrollsDynamically() bool
	SetScrollsDynamically(value bool)
	FindBarPosition() ScrollViewFindBarPosition
	SetFindBarPosition(value ScrollViewFindBarPosition)
	UsesPredominantAxisScrolling() bool
	SetUsesPredominantAxisScrolling(value bool)
	HorizontalScrollElasticity() ScrollElasticity
	SetHorizontalScrollElasticity(value ScrollElasticity)
	VerticalScrollElasticity() ScrollElasticity
	SetVerticalScrollElasticity(value ScrollElasticity)
	AllowsMagnification() bool
	SetAllowsMagnification(value bool)
	Magnification() coregraphics.Float
	SetMagnification(value coregraphics.Float)
	MaxMagnification() coregraphics.Float
	SetMaxMagnification(value coregraphics.Float)
	MinMagnification() coregraphics.Float
	SetMinMagnification(value coregraphics.Float)
}

type NSScrollView struct {
	NSView
}

func MakeScrollView(ptr unsafe.Pointer) *NSScrollView {
	if ptr == nil {
		return nil
	}
	return &NSScrollView{
		NSView: *MakeView(ptr),
	}
}

func AllocScrollView() *NSScrollView {
	return MakeScrollView(C.C_ScrollView_Alloc())
}

func (n *NSScrollView) InitWithCoder(coder foundation.Coder) ScrollView {
	result_ := C.C_NSScrollView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeScrollView(result_)
}

func (n *NSScrollView) InitWithFrame(frameRect foundation.Rect) ScrollView {
	result_ := C.C_NSScrollView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeScrollView(result_)
}

func (n *NSScrollView) Init() ScrollView {
	result_ := C.C_NSScrollView_Init(n.Ptr())
	return MakeScrollView(result_)
}

func (n *NSScrollView) AddFloatingSubview_ForAxis(view View, axis EventGestureAxis) {
	C.C_NSScrollView_AddFloatingSubview_ForAxis(n.Ptr(), objc.ExtractPtr(view), C.int(int(axis)))
}

func (n *NSScrollView) Tile() {
	C.C_NSScrollView_Tile(n.Ptr())
}

func (n *NSScrollView) FlashScrollers() {
	C.C_NSScrollView_FlashScrollers(n.Ptr())
}

func (n *NSScrollView) MagnifyToFitRect(rect foundation.Rect) {
	C.C_NSScrollView_MagnifyToFitRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSScrollView) SetMagnification_CenteredAtPoint(magnification coregraphics.Float, point foundation.Point) {
	C.C_NSScrollView_SetMagnification_CenteredAtPoint(n.Ptr(), C.double(float64(magnification)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
}

func (n *NSScrollView) ContentSize() foundation.Size {
	result_ := C.C_NSScrollView_ContentSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSScrollView) DocumentVisibleRect() foundation.Rect {
	result_ := C.C_NSScrollView_DocumentVisibleRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSScrollView) BackgroundColor() Color {
	result_ := C.C_NSScrollView_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n *NSScrollView) SetBackgroundColor(value Color) {
	C.C_NSScrollView_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSScrollView) DrawsBackground() bool {
	result_ := C.C_NSScrollView_DrawsBackground(n.Ptr())
	return bool(result_)
}

func (n *NSScrollView) SetDrawsBackground(value bool) {
	C.C_NSScrollView_SetDrawsBackground(n.Ptr(), C.bool(value))
}

func (n *NSScrollView) BorderType() BorderType {
	result_ := C.C_NSScrollView_BorderType(n.Ptr())
	return BorderType(uint(result_))
}

func (n *NSScrollView) SetBorderType(value BorderType) {
	C.C_NSScrollView_SetBorderType(n.Ptr(), C.uint(uint(value)))
}

func (n *NSScrollView) DocumentCursor() Cursor {
	result_ := C.C_NSScrollView_DocumentCursor(n.Ptr())
	return MakeCursor(result_)
}

func (n *NSScrollView) SetDocumentCursor(value Cursor) {
	C.C_NSScrollView_SetDocumentCursor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSScrollView) ContentView() ClipView {
	result_ := C.C_NSScrollView_ContentView(n.Ptr())
	return MakeClipView(result_)
}

func (n *NSScrollView) SetContentView(value ClipView) {
	C.C_NSScrollView_SetContentView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSScrollView) DocumentView() View {
	result_ := C.C_NSScrollView_DocumentView(n.Ptr())
	return MakeView(result_)
}

func (n *NSScrollView) SetDocumentView(value View) {
	C.C_NSScrollView_SetDocumentView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSScrollView) HorizontalScroller() Scroller {
	result_ := C.C_NSScrollView_HorizontalScroller(n.Ptr())
	return MakeScroller(result_)
}

func (n *NSScrollView) SetHorizontalScroller(value Scroller) {
	C.C_NSScrollView_SetHorizontalScroller(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSScrollView) HasHorizontalScroller() bool {
	result_ := C.C_NSScrollView_HasHorizontalScroller(n.Ptr())
	return bool(result_)
}

func (n *NSScrollView) SetHasHorizontalScroller(value bool) {
	C.C_NSScrollView_SetHasHorizontalScroller(n.Ptr(), C.bool(value))
}

func (n *NSScrollView) VerticalScroller() Scroller {
	result_ := C.C_NSScrollView_VerticalScroller(n.Ptr())
	return MakeScroller(result_)
}

func (n *NSScrollView) SetVerticalScroller(value Scroller) {
	C.C_NSScrollView_SetVerticalScroller(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSScrollView) HasVerticalScroller() bool {
	result_ := C.C_NSScrollView_HasVerticalScroller(n.Ptr())
	return bool(result_)
}

func (n *NSScrollView) SetHasVerticalScroller(value bool) {
	C.C_NSScrollView_SetHasVerticalScroller(n.Ptr(), C.bool(value))
}

func (n *NSScrollView) AutohidesScrollers() bool {
	result_ := C.C_NSScrollView_AutohidesScrollers(n.Ptr())
	return bool(result_)
}

func (n *NSScrollView) SetAutohidesScrollers(value bool) {
	C.C_NSScrollView_SetAutohidesScrollers(n.Ptr(), C.bool(value))
}

func (n *NSScrollView) HasHorizontalRuler() bool {
	result_ := C.C_NSScrollView_HasHorizontalRuler(n.Ptr())
	return bool(result_)
}

func (n *NSScrollView) SetHasHorizontalRuler(value bool) {
	C.C_NSScrollView_SetHasHorizontalRuler(n.Ptr(), C.bool(value))
}

func (n *NSScrollView) HorizontalRulerView() RulerView {
	result_ := C.C_NSScrollView_HorizontalRulerView(n.Ptr())
	return MakeRulerView(result_)
}

func (n *NSScrollView) SetHorizontalRulerView(value RulerView) {
	C.C_NSScrollView_SetHorizontalRulerView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSScrollView) HasVerticalRuler() bool {
	result_ := C.C_NSScrollView_HasVerticalRuler(n.Ptr())
	return bool(result_)
}

func (n *NSScrollView) SetHasVerticalRuler(value bool) {
	C.C_NSScrollView_SetHasVerticalRuler(n.Ptr(), C.bool(value))
}

func (n *NSScrollView) VerticalRulerView() RulerView {
	result_ := C.C_NSScrollView_VerticalRulerView(n.Ptr())
	return MakeRulerView(result_)
}

func (n *NSScrollView) SetVerticalRulerView(value RulerView) {
	C.C_NSScrollView_SetVerticalRulerView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSScrollView) RulersVisible() bool {
	result_ := C.C_NSScrollView_RulersVisible(n.Ptr())
	return bool(result_)
}

func (n *NSScrollView) SetRulersVisible(value bool) {
	C.C_NSScrollView_SetRulersVisible(n.Ptr(), C.bool(value))
}

func (n *NSScrollView) AutomaticallyAdjustsContentInsets() bool {
	result_ := C.C_NSScrollView_AutomaticallyAdjustsContentInsets(n.Ptr())
	return bool(result_)
}

func (n *NSScrollView) SetAutomaticallyAdjustsContentInsets(value bool) {
	C.C_NSScrollView_SetAutomaticallyAdjustsContentInsets(n.Ptr(), C.bool(value))
}

func (n *NSScrollView) ContentInsets() foundation.EdgeInsets {
	result_ := C.C_NSScrollView_ContentInsets(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n *NSScrollView) SetContentInsets(value foundation.EdgeInsets) {
	C.C_NSScrollView_SetContentInsets(n.Ptr(), *(*C.NSEdgeInsets)(foundation.ToNSEdgeInsetsPointer(value)))
}

func (n *NSScrollView) ScrollerInsets() foundation.EdgeInsets {
	result_ := C.C_NSScrollView_ScrollerInsets(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n *NSScrollView) SetScrollerInsets(value foundation.EdgeInsets) {
	C.C_NSScrollView_SetScrollerInsets(n.Ptr(), *(*C.NSEdgeInsets)(foundation.ToNSEdgeInsetsPointer(value)))
}

func (n *NSScrollView) ScrollerKnobStyle() ScrollerKnobStyle {
	result_ := C.C_NSScrollView_ScrollerKnobStyle(n.Ptr())
	return ScrollerKnobStyle(int(result_))
}

func (n *NSScrollView) SetScrollerKnobStyle(value ScrollerKnobStyle) {
	C.C_NSScrollView_SetScrollerKnobStyle(n.Ptr(), C.int(int(value)))
}

func (n *NSScrollView) ScrollerStyle() ScrollerStyle {
	result_ := C.C_NSScrollView_ScrollerStyle(n.Ptr())
	return ScrollerStyle(int(result_))
}

func (n *NSScrollView) SetScrollerStyle(value ScrollerStyle) {
	C.C_NSScrollView_SetScrollerStyle(n.Ptr(), C.int(int(value)))
}

func (n *NSScrollView) LineScroll() coregraphics.Float {
	result_ := C.C_NSScrollView_LineScroll(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSScrollView) SetLineScroll(value coregraphics.Float) {
	C.C_NSScrollView_SetLineScroll(n.Ptr(), C.double(float64(value)))
}

func (n *NSScrollView) HorizontalLineScroll() coregraphics.Float {
	result_ := C.C_NSScrollView_HorizontalLineScroll(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSScrollView) SetHorizontalLineScroll(value coregraphics.Float) {
	C.C_NSScrollView_SetHorizontalLineScroll(n.Ptr(), C.double(float64(value)))
}

func (n *NSScrollView) VerticalLineScroll() coregraphics.Float {
	result_ := C.C_NSScrollView_VerticalLineScroll(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSScrollView) SetVerticalLineScroll(value coregraphics.Float) {
	C.C_NSScrollView_SetVerticalLineScroll(n.Ptr(), C.double(float64(value)))
}

func (n *NSScrollView) PageScroll() coregraphics.Float {
	result_ := C.C_NSScrollView_PageScroll(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSScrollView) SetPageScroll(value coregraphics.Float) {
	C.C_NSScrollView_SetPageScroll(n.Ptr(), C.double(float64(value)))
}

func (n *NSScrollView) HorizontalPageScroll() coregraphics.Float {
	result_ := C.C_NSScrollView_HorizontalPageScroll(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSScrollView) SetHorizontalPageScroll(value coregraphics.Float) {
	C.C_NSScrollView_SetHorizontalPageScroll(n.Ptr(), C.double(float64(value)))
}

func (n *NSScrollView) VerticalPageScroll() coregraphics.Float {
	result_ := C.C_NSScrollView_VerticalPageScroll(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSScrollView) SetVerticalPageScroll(value coregraphics.Float) {
	C.C_NSScrollView_SetVerticalPageScroll(n.Ptr(), C.double(float64(value)))
}

func (n *NSScrollView) ScrollsDynamically() bool {
	result_ := C.C_NSScrollView_ScrollsDynamically(n.Ptr())
	return bool(result_)
}

func (n *NSScrollView) SetScrollsDynamically(value bool) {
	C.C_NSScrollView_SetScrollsDynamically(n.Ptr(), C.bool(value))
}

func (n *NSScrollView) FindBarPosition() ScrollViewFindBarPosition {
	result_ := C.C_NSScrollView_FindBarPosition(n.Ptr())
	return ScrollViewFindBarPosition(int(result_))
}

func (n *NSScrollView) SetFindBarPosition(value ScrollViewFindBarPosition) {
	C.C_NSScrollView_SetFindBarPosition(n.Ptr(), C.int(int(value)))
}

func (n *NSScrollView) UsesPredominantAxisScrolling() bool {
	result_ := C.C_NSScrollView_UsesPredominantAxisScrolling(n.Ptr())
	return bool(result_)
}

func (n *NSScrollView) SetUsesPredominantAxisScrolling(value bool) {
	C.C_NSScrollView_SetUsesPredominantAxisScrolling(n.Ptr(), C.bool(value))
}

func (n *NSScrollView) HorizontalScrollElasticity() ScrollElasticity {
	result_ := C.C_NSScrollView_HorizontalScrollElasticity(n.Ptr())
	return ScrollElasticity(int(result_))
}

func (n *NSScrollView) SetHorizontalScrollElasticity(value ScrollElasticity) {
	C.C_NSScrollView_SetHorizontalScrollElasticity(n.Ptr(), C.int(int(value)))
}

func (n *NSScrollView) VerticalScrollElasticity() ScrollElasticity {
	result_ := C.C_NSScrollView_VerticalScrollElasticity(n.Ptr())
	return ScrollElasticity(int(result_))
}

func (n *NSScrollView) SetVerticalScrollElasticity(value ScrollElasticity) {
	C.C_NSScrollView_SetVerticalScrollElasticity(n.Ptr(), C.int(int(value)))
}

func (n *NSScrollView) AllowsMagnification() bool {
	result_ := C.C_NSScrollView_AllowsMagnification(n.Ptr())
	return bool(result_)
}

func (n *NSScrollView) SetAllowsMagnification(value bool) {
	C.C_NSScrollView_SetAllowsMagnification(n.Ptr(), C.bool(value))
}

func (n *NSScrollView) Magnification() coregraphics.Float {
	result_ := C.C_NSScrollView_Magnification(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSScrollView) SetMagnification(value coregraphics.Float) {
	C.C_NSScrollView_SetMagnification(n.Ptr(), C.double(float64(value)))
}

func (n *NSScrollView) MaxMagnification() coregraphics.Float {
	result_ := C.C_NSScrollView_MaxMagnification(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSScrollView) SetMaxMagnification(value coregraphics.Float) {
	C.C_NSScrollView_SetMaxMagnification(n.Ptr(), C.double(float64(value)))
}

func (n *NSScrollView) MinMagnification() coregraphics.Float {
	result_ := C.C_NSScrollView_MinMagnification(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSScrollView) SetMinMagnification(value coregraphics.Float) {
	C.C_NSScrollView_SetMinMagnification(n.Ptr(), C.double(float64(value)))
}
