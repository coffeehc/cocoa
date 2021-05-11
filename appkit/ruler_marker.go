package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "ruler_marker.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type RulerMarker interface {
	objc.Object
	DrawRect(rect foundation.Rect)
	TrackMouse_Adding(mouseDownEvent Event, isAdding bool) bool
	Ruler() RulerView
	Image() Image
	SetImage(value Image)
	ImageOrigin() foundation.Point
	SetImageOrigin(value foundation.Point)
	ImageRectInRuler() foundation.Rect
	ThicknessRequiredInRuler() coregraphics.Float
	IsMovable() bool
	SetMovable(value bool)
	IsRemovable() bool
	SetRemovable(value bool)
	MarkerLocation() coregraphics.Float
	SetMarkerLocation(value coregraphics.Float)
	IsDragging() bool
}

type NSRulerMarker struct {
	objc.NSObject
}

func MakeRulerMarker(ptr unsafe.Pointer) *NSRulerMarker {
	if ptr == nil {
		return nil
	}
	return &NSRulerMarker{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocRulerMarker() *NSRulerMarker {
	return MakeRulerMarker(C.C_RulerMarker_Alloc())
}

func (n *NSRulerMarker) InitWithRulerView_MarkerLocation_Image_ImageOrigin(ruler RulerView, location coregraphics.Float, image Image, imageOrigin foundation.Point) RulerMarker {
	result := C.C_NSRulerMarker_InitWithRulerView_MarkerLocation_Image_ImageOrigin(n.Ptr(), objc.ExtractPtr(ruler), C.double(float64(location)), objc.ExtractPtr(image), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(imageOrigin))))
	return MakeRulerMarker(result)
}

func (n *NSRulerMarker) InitWithCoder(coder foundation.Coder) RulerMarker {
	result := C.C_NSRulerMarker_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeRulerMarker(result)
}

func (n *NSRulerMarker) DrawRect(rect foundation.Rect) {
	C.C_NSRulerMarker_DrawRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSRulerMarker) TrackMouse_Adding(mouseDownEvent Event, isAdding bool) bool {
	result := C.C_NSRulerMarker_TrackMouse_Adding(n.Ptr(), objc.ExtractPtr(mouseDownEvent), C.bool(isAdding))
	return bool(result)
}

func (n *NSRulerMarker) Ruler() RulerView {
	result := C.C_NSRulerMarker_Ruler(n.Ptr())
	return MakeRulerView(result)
}

func (n *NSRulerMarker) Image() Image {
	result := C.C_NSRulerMarker_Image(n.Ptr())
	return MakeImage(result)
}

func (n *NSRulerMarker) SetImage(value Image) {
	C.C_NSRulerMarker_SetImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSRulerMarker) ImageOrigin() foundation.Point {
	result := C.C_NSRulerMarker_ImageOrigin(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSRulerMarker) SetImageOrigin(value foundation.Point) {
	C.C_NSRulerMarker_SetImageOrigin(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(value))))
}

func (n *NSRulerMarker) ImageRectInRuler() foundation.Rect {
	result := C.C_NSRulerMarker_ImageRectInRuler(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSRulerMarker) ThicknessRequiredInRuler() coregraphics.Float {
	result := C.C_NSRulerMarker_ThicknessRequiredInRuler(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSRulerMarker) IsMovable() bool {
	result := C.C_NSRulerMarker_IsMovable(n.Ptr())
	return bool(result)
}

func (n *NSRulerMarker) SetMovable(value bool) {
	C.C_NSRulerMarker_SetMovable(n.Ptr(), C.bool(value))
}

func (n *NSRulerMarker) IsRemovable() bool {
	result := C.C_NSRulerMarker_IsRemovable(n.Ptr())
	return bool(result)
}

func (n *NSRulerMarker) SetRemovable(value bool) {
	C.C_NSRulerMarker_SetRemovable(n.Ptr(), C.bool(value))
}

func (n *NSRulerMarker) MarkerLocation() coregraphics.Float {
	result := C.C_NSRulerMarker_MarkerLocation(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSRulerMarker) SetMarkerLocation(value coregraphics.Float) {
	C.C_NSRulerMarker_SetMarkerLocation(n.Ptr(), C.double(float64(value)))
}

func (n *NSRulerMarker) IsDragging() bool {
	result := C.C_NSRulerMarker_IsDragging(n.Ptr())
	return bool(result)
}