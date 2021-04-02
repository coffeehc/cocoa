package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "layout_guide.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type LayoutGuide interface {
	objc.Object
	Frame() foundation.Rect
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
	OwningView() View
}

var _ LayoutGuide = (*NSLayoutGuide)(nil)

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

func (l *NSLayoutGuide) Frame() foundation.Rect {
	return toRect(C.LayoutGuide_Frame(l.Ptr()))
}

func (l *NSLayoutGuide) BottomAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.LayoutGuide_BottomAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) CenterXAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.LayoutGuide_CenterXAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) CenterYAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.LayoutGuide_CenterYAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) HeightAnchor() LayoutDimension {
	return MakeLayoutDimension(C.LayoutGuide_HeightAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) LeadingAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.LayoutGuide_LeadingAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) LeftAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.LayoutGuide_LeftAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) RightAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.LayoutGuide_RightAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) TopAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.LayoutGuide_TopAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) TrailingAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.LayoutGuide_TrailingAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) WidthAnchor() LayoutDimension {
	return MakeLayoutDimension(C.LayoutGuide_WidthAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) OwningView() View {
	return MakeView(C.LayoutGuide_OwningView(l.Ptr()))
}

func NewLayoutGuide() LayoutGuide {
	return MakeLayoutGuide(C.LayoutGuide_NewLayoutGuide())
}
