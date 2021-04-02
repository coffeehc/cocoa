package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type View interface {
	Responder
	Frame() foundation.Rect
	SetFrame(frame foundation.Rect)
	AutoresizingMask() AutoresizingMaskOptions
	SetAutoresizingMask(autoresizingMask AutoresizingMaskOptions)
	NeedsDisplay() bool
	SetNeedsDisplay(needsDisplay bool)
	TranslatesAutoresizingMaskIntoConstraints() bool
	SetTranslatesAutoresizingMaskIntoConstraints(translatesAutoresizingMaskIntoConstraints bool)
	BottomAnchor() LayoutYAxisAnchor
	CenterXAnchor() LayoutXAxisAnchor
	CenterYAnchor() LayoutYAxisAnchor
	FirstBaselineAnchor() LayoutYAxisAnchor
	HeightAnchor() LayoutDimension
	LastBaselineAnchor() LayoutYAxisAnchor
	LeadingAnchor() LayoutXAxisAnchor
	LeftAnchor() LayoutXAxisAnchor
	RightAnchor() LayoutXAxisAnchor
	TopAnchor() LayoutYAxisAnchor
	TrailingAnchor() LayoutXAxisAnchor
	WidthAnchor() LayoutDimension
	AddSubview(subView View)
}

var _ View = (*NSView)(nil)

type NSView struct {
	NSResponder
}

func MakeView(ptr unsafe.Pointer) *NSView {
	if ptr == nil {
		return nil
	}
	return &NSView{
		NSResponder: *MakeResponder(ptr),
	}
}

func (v *NSView) Frame() foundation.Rect {
	return toRect(C.View_Frame(v.Ptr()))
}

func (v *NSView) SetFrame(frame foundation.Rect) {
	C.View_SetFrame(v.Ptr(), toNSRect(frame))
}

func (v *NSView) AutoresizingMask() AutoresizingMaskOptions {
	return AutoresizingMaskOptions(C.View_AutoresizingMask(v.Ptr()))
}

func (v *NSView) SetAutoresizingMask(autoresizingMask AutoresizingMaskOptions) {
	C.View_SetAutoresizingMask(v.Ptr(), C.ulong(autoresizingMask))
}

func (v *NSView) NeedsDisplay() bool {
	return bool(C.View_NeedsDisplay(v.Ptr()))
}

func (v *NSView) SetNeedsDisplay(needsDisplay bool) {
	C.View_SetNeedsDisplay(v.Ptr(), C.bool(needsDisplay))
}

func (v *NSView) TranslatesAutoresizingMaskIntoConstraints() bool {
	return bool(C.View_TranslatesAutoresizingMaskIntoConstraints(v.Ptr()))
}

func (v *NSView) SetTranslatesAutoresizingMaskIntoConstraints(translatesAutoresizingMaskIntoConstraints bool) {
	C.View_SetTranslatesAutoresizingMaskIntoConstraints(v.Ptr(), C.bool(translatesAutoresizingMaskIntoConstraints))
}

func (v *NSView) BottomAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.View_BottomAnchor(v.Ptr()))
}

func (v *NSView) CenterXAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.View_CenterXAnchor(v.Ptr()))
}

func (v *NSView) CenterYAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.View_CenterYAnchor(v.Ptr()))
}

func (v *NSView) FirstBaselineAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.View_FirstBaselineAnchor(v.Ptr()))
}

func (v *NSView) HeightAnchor() LayoutDimension {
	return MakeLayoutDimension(C.View_HeightAnchor(v.Ptr()))
}

func (v *NSView) LastBaselineAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.View_LastBaselineAnchor(v.Ptr()))
}

func (v *NSView) LeadingAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.View_LeadingAnchor(v.Ptr()))
}

func (v *NSView) LeftAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.View_LeftAnchor(v.Ptr()))
}

func (v *NSView) RightAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.View_RightAnchor(v.Ptr()))
}

func (v *NSView) TopAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.View_TopAnchor(v.Ptr()))
}

func (v *NSView) TrailingAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.View_TrailingAnchor(v.Ptr()))
}

func (v *NSView) WidthAnchor() LayoutDimension {
	return MakeLayoutDimension(C.View_WidthAnchor(v.Ptr()))
}

func (v *NSView) AddSubview(subView View) {
	C.View_AddSubview(v.Ptr(), toPointer(subView))
}
