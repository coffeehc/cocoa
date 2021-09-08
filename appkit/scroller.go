package appkit

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
}

type NSScroller struct {
	NSControl
}

func MakeScroller(ptr unsafe.Pointer) NSScroller {
	return NSScroller{
		NSControl: MakeControl(ptr),
	}
}

func (n NSScroller) InitWithFrame(frameRect foundation.Rect) NSScroller {
	result_ := C.C_NSScroller_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeScroller(result_)
}

func (n NSScroller) InitWithCoder(coder foundation.Coder) NSScroller {
	result_ := C.C_NSScroller_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeScroller(result_)
}

func (n NSScroller) Init() NSScroller {
	result_ := C.C_NSScroller_Init(n.Ptr())
	return MakeScroller(result_)
}

func AllocScroller() NSScroller {
	result_ := C.C_NSScroller_AllocScroller()
	return MakeScroller(result_)
}

func NewScroller() NSScroller {
	result_ := C.C_NSScroller_NewScroller()
	return MakeScroller(result_)
}

func (n NSScroller) Autorelease() NSScroller {
	result_ := C.C_NSScroller_Autorelease(n.Ptr())
	return MakeScroller(result_)
}

func (n NSScroller) Retain() NSScroller {
	result_ := C.C_NSScroller_Retain(n.Ptr())
	return MakeScroller(result_)
}

func ScrollerWidthForControlSize_ScrollerStyle(controlSize ControlSize, scrollerStyle ScrollerStyle) coregraphics.Float {
	result_ := C.C_NSScroller_ScrollerWidthForControlSize_ScrollerStyle(C.uint(uint(controlSize)), C.int(int(scrollerStyle)))
	return coregraphics.Float(float64(result_))
}

func (n NSScroller) RectForPart(partCode ScrollerPart) foundation.Rect {
	result_ := C.C_NSScroller_RectForPart(n.Ptr(), C.uint(uint(partCode)))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSScroller) TestPart(point foundation.Point) ScrollerPart {
	result_ := C.C_NSScroller_TestPart(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return ScrollerPart(uint(result_))
}

func (n NSScroller) CheckSpaceForParts() {
	C.C_NSScroller_CheckSpaceForParts(n.Ptr())
}

func (n NSScroller) DrawKnobSlotInRect_Highlight(slotRect foundation.Rect, flag bool) {
	C.C_NSScroller_DrawKnobSlotInRect_Highlight(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(slotRect))), C.bool(flag))
}

func (n NSScroller) DrawKnob() {
	C.C_NSScroller_DrawKnob(n.Ptr())
}

func (n NSScroller) TrackKnob(event Event) {
	C.C_NSScroller_TrackKnob(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSScroller) UsableParts() UsableScrollerParts {
	result_ := C.C_NSScroller_UsableParts(n.Ptr())
	return UsableScrollerParts(uint(result_))
}

func (n NSScroller) HitPart() ScrollerPart {
	result_ := C.C_NSScroller_HitPart(n.Ptr())
	return ScrollerPart(uint(result_))
}

func PreferredScrollerStyle() ScrollerStyle {
	result_ := C.C_NSScroller_PreferredScrollerStyle()
	return ScrollerStyle(int(result_))
}

func (n NSScroller) ScrollerStyle() ScrollerStyle {
	result_ := C.C_NSScroller_ScrollerStyle(n.Ptr())
	return ScrollerStyle(int(result_))
}

func (n NSScroller) SetScrollerStyle(value ScrollerStyle) {
	C.C_NSScroller_SetScrollerStyle(n.Ptr(), C.int(int(value)))
}

func (n NSScroller) KnobStyle() ScrollerKnobStyle {
	result_ := C.C_NSScroller_KnobStyle(n.Ptr())
	return ScrollerKnobStyle(int(result_))
}

func (n NSScroller) SetKnobStyle(value ScrollerKnobStyle) {
	C.C_NSScroller_SetKnobStyle(n.Ptr(), C.int(int(value)))
}

func (n NSScroller) KnobProportion() coregraphics.Float {
	result_ := C.C_NSScroller_KnobProportion(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func CompatibleWithOverlayScrollers() bool {
	result_ := C.C_NSScroller_CompatibleWithOverlayScrollers()
	return bool(result_)
}
