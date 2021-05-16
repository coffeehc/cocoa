package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "scroller.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Scroller interface {
	Control
	RectForPart(partCode ScrollerPart) foundation.Rect
	TestPart(point foundation.Point) ScrollerPart
	CheckSpaceForParts()
	DrawKnobSlotInRect_Highlight(slotRect foundation.Rect, flag bool)
	DrawKnob()
	TrackKnob(event Event)
	UsableParts() UsableScrollerParts
	HitPart() ScrollerPart
	ScrollerStyle() ScrollerStyle
	SetScrollerStyle(value ScrollerStyle)
	KnobStyle() ScrollerKnobStyle
	SetKnobStyle(value ScrollerKnobStyle)
	KnobProportion() coregraphics.Float
	SetKnobProportion(value coregraphics.Float)
}

type NSScroller struct {
	NSControl
}

func MakeScroller(ptr unsafe.Pointer) *NSScroller {
	if ptr == nil {
		return nil
	}
	return &NSScroller{
		NSControl: *MakeControl(ptr),
	}
}

func AllocScroller() *NSScroller {
	return MakeScroller(C.C_Scroller_Alloc())
}

func (n *NSScroller) InitWithFrame(frameRect foundation.Rect) Scroller {
	result := C.C_NSScroller_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeScroller(result)
}

func (n *NSScroller) InitWithCoder(coder foundation.Coder) Scroller {
	result := C.C_NSScroller_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeScroller(result)
}

func (n *NSScroller) Init() Scroller {
	result := C.C_NSScroller_Init(n.Ptr())
	return MakeScroller(result)
}

func ScrollerWidthForControlSize_ScrollerStyle(controlSize ControlSize, scrollerStyle ScrollerStyle) coregraphics.Float {
	result := C.C_NSScroller_ScrollerWidthForControlSize_ScrollerStyle(C.uint(uint(controlSize)), C.int(int(scrollerStyle)))
	return coregraphics.Float(float64(result))
}

func (n *NSScroller) RectForPart(partCode ScrollerPart) foundation.Rect {
	result := C.C_NSScroller_RectForPart(n.Ptr(), C.uint(uint(partCode)))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSScroller) TestPart(point foundation.Point) ScrollerPart {
	result := C.C_NSScroller_TestPart(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return ScrollerPart(uint(result))
}

func (n *NSScroller) CheckSpaceForParts() {
	C.C_NSScroller_CheckSpaceForParts(n.Ptr())
}

func (n *NSScroller) DrawKnobSlotInRect_Highlight(slotRect foundation.Rect, flag bool) {
	C.C_NSScroller_DrawKnobSlotInRect_Highlight(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(slotRect))), C.bool(flag))
}

func (n *NSScroller) DrawKnob() {
	C.C_NSScroller_DrawKnob(n.Ptr())
}

func (n *NSScroller) TrackKnob(event Event) {
	C.C_NSScroller_TrackKnob(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSScroller) UsableParts() UsableScrollerParts {
	result := C.C_NSScroller_UsableParts(n.Ptr())
	return UsableScrollerParts(uint(result))
}

func (n *NSScroller) HitPart() ScrollerPart {
	result := C.C_NSScroller_HitPart(n.Ptr())
	return ScrollerPart(uint(result))
}

func Scroller_PreferredScrollerStyle() ScrollerStyle {
	result := C.C_NSScroller_Scroller_PreferredScrollerStyle()
	return ScrollerStyle(int(result))
}

func (n *NSScroller) ScrollerStyle() ScrollerStyle {
	result := C.C_NSScroller_ScrollerStyle(n.Ptr())
	return ScrollerStyle(int(result))
}

func (n *NSScroller) SetScrollerStyle(value ScrollerStyle) {
	C.C_NSScroller_SetScrollerStyle(n.Ptr(), C.int(int(value)))
}

func (n *NSScroller) KnobStyle() ScrollerKnobStyle {
	result := C.C_NSScroller_KnobStyle(n.Ptr())
	return ScrollerKnobStyle(int(result))
}

func (n *NSScroller) SetKnobStyle(value ScrollerKnobStyle) {
	C.C_NSScroller_SetKnobStyle(n.Ptr(), C.int(int(value)))
}

func (n *NSScroller) KnobProportion() coregraphics.Float {
	result := C.C_NSScroller_KnobProportion(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSScroller) SetKnobProportion(value coregraphics.Float) {
	C.C_NSScroller_SetKnobProportion(n.Ptr(), C.double(float64(value)))
}

func Scroller_CompatibleWithOverlayScrollers() bool {
	result := C.C_NSScroller_Scroller_CompatibleWithOverlayScrollers()
	return bool(result)
}
