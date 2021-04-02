package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "clip_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type ClipView interface {
	View
	ContentInsets() foundation.EdgeInsets
	SetContentInsets(contentInsets foundation.EdgeInsets)
	DocumentView() View
	SetDocumentView(documentView View)
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(automaticallyAdjustsContentInsets bool)
	DocumentRect() foundation.Rect
	DocumentVisibleRect() foundation.Rect
	DrawsBackground() bool
	SetDrawsBackground(drawsBackground bool)
	BackgroundColor() Color
	SetBackgroundColor(backgroundColor Color)
	ScrollToPoint(newOrigin foundation.Point)
	Autoscroll(event Event)
	ConstrainBoundsRect(proposedBounds foundation.Rect) foundation.Rect
}

var _ ClipView = (*NSClipView)(nil)

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

func (c *NSClipView) ContentInsets() foundation.EdgeInsets {
	return toEdgeInsets(C.ClipView_ContentInsets(c.Ptr()))
}

func (c *NSClipView) SetContentInsets(contentInsets foundation.EdgeInsets) {
	C.ClipView_SetContentInsets(c.Ptr(), toNSEdgeInsets(contentInsets))
}

func (c *NSClipView) DocumentView() View {
	return MakeView(C.ClipView_DocumentView(c.Ptr()))
}

func (c *NSClipView) SetDocumentView(documentView View) {
	C.ClipView_SetDocumentView(c.Ptr(), toPointer(documentView))
}

func (c *NSClipView) AutomaticallyAdjustsContentInsets() bool {
	return bool(C.ClipView_AutomaticallyAdjustsContentInsets(c.Ptr()))
}

func (c *NSClipView) SetAutomaticallyAdjustsContentInsets(automaticallyAdjustsContentInsets bool) {
	C.ClipView_SetAutomaticallyAdjustsContentInsets(c.Ptr(), C.bool(automaticallyAdjustsContentInsets))
}

func (c *NSClipView) DocumentRect() foundation.Rect {
	return toRect(C.ClipView_DocumentRect(c.Ptr()))
}

func (c *NSClipView) DocumentVisibleRect() foundation.Rect {
	return toRect(C.ClipView_DocumentVisibleRect(c.Ptr()))
}

func (c *NSClipView) DrawsBackground() bool {
	return bool(C.ClipView_DrawsBackground(c.Ptr()))
}

func (c *NSClipView) SetDrawsBackground(drawsBackground bool) {
	C.ClipView_SetDrawsBackground(c.Ptr(), C.bool(drawsBackground))
}

func (c *NSClipView) BackgroundColor() Color {
	return MakeColor(C.ClipView_BackgroundColor(c.Ptr()))
}

func (c *NSClipView) SetBackgroundColor(backgroundColor Color) {
	C.ClipView_SetBackgroundColor(c.Ptr(), toPointer(backgroundColor))
}

func (c *NSClipView) ScrollToPoint(newOrigin foundation.Point) {
	C.ClipView_ScrollToPoint(c.Ptr(), toNSPoint(newOrigin))
}

func (c *NSClipView) Autoscroll(event Event) {
	C.ClipView_Autoscroll(c.Ptr(), toPointer(event))
}

func (c *NSClipView) ConstrainBoundsRect(proposedBounds foundation.Rect) foundation.Rect {
	return toRect(C.ClipView_ConstrainBoundsRect(c.Ptr(), toNSRect(proposedBounds)))
}
