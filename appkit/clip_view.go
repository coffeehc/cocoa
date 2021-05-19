package appkit

// #include "clip_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ClipView interface {
	View
	ScrollToPoint(newOrigin foundation.Point)
	ConstrainBoundsRect(proposedBounds foundation.Rect) foundation.Rect
	ViewBoundsChanged(notification foundation.Notification)
	ViewFrameChanged(notification foundation.Notification)
	DocumentView() View
	SetDocumentView(value View)
	ContentInsets() foundation.EdgeInsets
	SetContentInsets(value foundation.EdgeInsets)
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(value bool)
	DocumentRect() foundation.Rect
	DocumentVisibleRect() foundation.Rect
	DocumentCursor() Cursor
	SetDocumentCursor(value Cursor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
}

type NSClipView struct {
	NSView
}

func MakeClipView(ptr unsafe.Pointer) *NSClipView {
	if ptr == nil {
		return nil
	}
	return &NSClipView{
		NSView: *MakeView(ptr),
	}
}

func AllocClipView() *NSClipView {
	return MakeClipView(C.C_ClipView_Alloc())
}

func (n *NSClipView) InitWithFrame(frameRect foundation.Rect) ClipView {
	result_ := C.C_NSClipView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeClipView(result_)
}

func (n *NSClipView) InitWithCoder(coder foundation.Coder) ClipView {
	result_ := C.C_NSClipView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeClipView(result_)
}

func (n *NSClipView) Init() ClipView {
	result_ := C.C_NSClipView_Init(n.Ptr())
	return MakeClipView(result_)
}

func (n *NSClipView) ScrollToPoint(newOrigin foundation.Point) {
	C.C_NSClipView_ScrollToPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(newOrigin))))
}

func (n *NSClipView) ConstrainBoundsRect(proposedBounds foundation.Rect) foundation.Rect {
	result_ := C.C_NSClipView_ConstrainBoundsRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(proposedBounds))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSClipView) ViewBoundsChanged(notification foundation.Notification) {
	C.C_NSClipView_ViewBoundsChanged(n.Ptr(), objc.ExtractPtr(notification))
}

func (n *NSClipView) ViewFrameChanged(notification foundation.Notification) {
	C.C_NSClipView_ViewFrameChanged(n.Ptr(), objc.ExtractPtr(notification))
}

func (n *NSClipView) DocumentView() View {
	result_ := C.C_NSClipView_DocumentView(n.Ptr())
	return MakeView(result_)
}

func (n *NSClipView) SetDocumentView(value View) {
	C.C_NSClipView_SetDocumentView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSClipView) ContentInsets() foundation.EdgeInsets {
	result_ := C.C_NSClipView_ContentInsets(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n *NSClipView) SetContentInsets(value foundation.EdgeInsets) {
	C.C_NSClipView_SetContentInsets(n.Ptr(), *(*C.NSEdgeInsets)(foundation.ToNSEdgeInsetsPointer(value)))
}

func (n *NSClipView) AutomaticallyAdjustsContentInsets() bool {
	result_ := C.C_NSClipView_AutomaticallyAdjustsContentInsets(n.Ptr())
	return bool(result_)
}

func (n *NSClipView) SetAutomaticallyAdjustsContentInsets(value bool) {
	C.C_NSClipView_SetAutomaticallyAdjustsContentInsets(n.Ptr(), C.bool(value))
}

func (n *NSClipView) DocumentRect() foundation.Rect {
	result_ := C.C_NSClipView_DocumentRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSClipView) DocumentVisibleRect() foundation.Rect {
	result_ := C.C_NSClipView_DocumentVisibleRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSClipView) DocumentCursor() Cursor {
	result_ := C.C_NSClipView_DocumentCursor(n.Ptr())
	return MakeCursor(result_)
}

func (n *NSClipView) SetDocumentCursor(value Cursor) {
	C.C_NSClipView_SetDocumentCursor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSClipView) DrawsBackground() bool {
	result_ := C.C_NSClipView_DrawsBackground(n.Ptr())
	return bool(result_)
}

func (n *NSClipView) SetDrawsBackground(value bool) {
	C.C_NSClipView_SetDrawsBackground(n.Ptr(), C.bool(value))
}

func (n *NSClipView) BackgroundColor() Color {
	result_ := C.C_NSClipView_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n *NSClipView) SetBackgroundColor(value Color) {
	C.C_NSClipView_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}
