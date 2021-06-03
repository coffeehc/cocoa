package appkit

// #include "event.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Event interface {
	objc.Object
	TouchesMatchingPhase_InView(phase TouchPhase, view View) foundation.Set
	AllTouches() foundation.Set
	TouchesForView(view View) foundation.Set
	CoalescedTouchesForTouch(touch Touch) []Touch
	CharactersByApplyingModifiers(modifiers EventModifierFlags) string
	Type() EventType
	ModifierFlags() EventModifierFlags
	LocationInWindow() foundation.Point
	Timestamp() foundation.TimeInterval
	Window() Window
	WindowNumber() int
	Characters() string
	CharactersIgnoringModifiers() string
	IsARepeat() bool
	ButtonNumber() int
	ClickCount() int
	EventNumber() int
	TrackingNumber() int
	TrackingArea() TrackingArea
	Data1() int
	Data2() int
	DeltaX() coregraphics.Float
	DeltaY() coregraphics.Float
	DeltaZ() coregraphics.Float
	Pressure() float32
	Stage() int
	StageTransition() coregraphics.Float
	PressureBehavior() PressureBehavior
	CapabilityMask() uint
	DeviceID() uint
	IsEnteringProximity() bool
	PointingDeviceID() uint
	PointingDeviceSerialNumber() uint
	PointingDeviceType() PointingDeviceType
	SystemTabletID() uint
	TabletID() uint
	VendorID() uint
	VendorPointingDeviceType() uint
	AbsoluteX() int
	AbsoluteY() int
	AbsoluteZ() int
	ButtonMask() EventButtonMask
	Rotation() float32
	TangentialPressure() float32
	Tilt() foundation.Point
	VendorDefined() objc.Object
	Magnification() coregraphics.Float
	HasPreciseScrollingDeltas() bool
	ScrollingDeltaX() coregraphics.Float
	ScrollingDeltaY() coregraphics.Float
	MomentumPhase() EventPhase
	Phase() EventPhase
	IsDirectionInvertedFromDevice() bool
}

type NSEvent struct {
	objc.NSObject
}

func MakeEvent(ptr unsafe.Pointer) NSEvent {
	return NSEvent{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocEvent() NSEvent {
	return MakeEvent(C.C_Event_Alloc())
}

func (n NSEvent) Init() Event {
	result_ := C.C_NSEvent_Init(n.Ptr())
	return MakeEvent(result_)
}

func Event_MouseEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_EventNumber_ClickCount_Pressure(_type EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil GraphicsContext, eNum int, cNum int, pressure float32) Event {
	result_ := C.C_NSEvent_Event_MouseEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_EventNumber_ClickCount_Pressure(C.uint(uint(_type)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(location))), C.uint(uint(flags)), C.double(float64(time)), C.int(wNum), objc.ExtractPtr(unusedPassNil), C.int(eNum), C.int(cNum), C.float(pressure))
	return MakeEvent(result_)
}

func Event_StartPeriodicEventsAfterDelay_WithPeriod(delay foundation.TimeInterval, period foundation.TimeInterval) {
	C.C_NSEvent_Event_StartPeriodicEventsAfterDelay_WithPeriod(C.double(float64(delay)), C.double(float64(period)))
}

func Event_StopPeriodicEvents() {
	C.C_NSEvent_Event_StopPeriodicEvents()
}

func (n NSEvent) TouchesMatchingPhase_InView(phase TouchPhase, view View) foundation.Set {
	result_ := C.C_NSEvent_TouchesMatchingPhase_InView(n.Ptr(), C.uint(uint(phase)), objc.ExtractPtr(view))
	return foundation.MakeSet(result_)
}

func (n NSEvent) AllTouches() foundation.Set {
	result_ := C.C_NSEvent_AllTouches(n.Ptr())
	return foundation.MakeSet(result_)
}

func (n NSEvent) TouchesForView(view View) foundation.Set {
	result_ := C.C_NSEvent_TouchesForView(n.Ptr(), objc.ExtractPtr(view))
	return foundation.MakeSet(result_)
}

func (n NSEvent) CoalescedTouchesForTouch(touch Touch) []Touch {
	result_ := C.C_NSEvent_CoalescedTouchesForTouch(n.Ptr(), objc.ExtractPtr(touch))
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]Touch, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTouch(r)
	}
	return goResult_
}

func Event_RemoveMonitor(eventMonitor objc.Object) {
	C.C_NSEvent_Event_RemoveMonitor(objc.ExtractPtr(eventMonitor))
}

func (n NSEvent) CharactersByApplyingModifiers(modifiers EventModifierFlags) string {
	result_ := C.C_NSEvent_CharactersByApplyingModifiers(n.Ptr(), C.uint(uint(modifiers)))
	return foundation.MakeString(result_).String()
}

func (n NSEvent) Type() EventType {
	result_ := C.C_NSEvent_Type(n.Ptr())
	return EventType(uint(result_))
}

func (n NSEvent) ModifierFlags() EventModifierFlags {
	result_ := C.C_NSEvent_ModifierFlags(n.Ptr())
	return EventModifierFlags(uint(result_))
}

func (n NSEvent) LocationInWindow() foundation.Point {
	result_ := C.C_NSEvent_LocationInWindow(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSEvent) Timestamp() foundation.TimeInterval {
	result_ := C.C_NSEvent_Timestamp(n.Ptr())
	return foundation.TimeInterval(float64(result_))
}

func (n NSEvent) Window() Window {
	result_ := C.C_NSEvent_Window(n.Ptr())
	return MakeWindow(result_)
}

func (n NSEvent) WindowNumber() int {
	result_ := C.C_NSEvent_WindowNumber(n.Ptr())
	return int(result_)
}

func Event_KeyRepeatDelay() foundation.TimeInterval {
	result_ := C.C_NSEvent_Event_KeyRepeatDelay()
	return foundation.TimeInterval(float64(result_))
}

func Event_KeyRepeatInterval() foundation.TimeInterval {
	result_ := C.C_NSEvent_Event_KeyRepeatInterval()
	return foundation.TimeInterval(float64(result_))
}

func (n NSEvent) Characters() string {
	result_ := C.C_NSEvent_Characters(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSEvent) CharactersIgnoringModifiers() string {
	result_ := C.C_NSEvent_CharactersIgnoringModifiers(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSEvent) IsARepeat() bool {
	result_ := C.C_NSEvent_IsARepeat(n.Ptr())
	return bool(result_)
}

func Event_PressedMouseButtons() uint {
	result_ := C.C_NSEvent_Event_PressedMouseButtons()
	return uint(result_)
}

func Event_DoubleClickInterval() foundation.TimeInterval {
	result_ := C.C_NSEvent_Event_DoubleClickInterval()
	return foundation.TimeInterval(float64(result_))
}

func Event_MouseLocation() foundation.Point {
	result_ := C.C_NSEvent_Event_MouseLocation()
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSEvent) ButtonNumber() int {
	result_ := C.C_NSEvent_ButtonNumber(n.Ptr())
	return int(result_)
}

func (n NSEvent) ClickCount() int {
	result_ := C.C_NSEvent_ClickCount(n.Ptr())
	return int(result_)
}

func (n NSEvent) EventNumber() int {
	result_ := C.C_NSEvent_EventNumber(n.Ptr())
	return int(result_)
}

func (n NSEvent) TrackingNumber() int {
	result_ := C.C_NSEvent_TrackingNumber(n.Ptr())
	return int(result_)
}

func (n NSEvent) TrackingArea() TrackingArea {
	result_ := C.C_NSEvent_TrackingArea(n.Ptr())
	return MakeTrackingArea(result_)
}

func (n NSEvent) Data1() int {
	result_ := C.C_NSEvent_Data1(n.Ptr())
	return int(result_)
}

func (n NSEvent) Data2() int {
	result_ := C.C_NSEvent_Data2(n.Ptr())
	return int(result_)
}

func (n NSEvent) DeltaX() coregraphics.Float {
	result_ := C.C_NSEvent_DeltaX(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSEvent) DeltaY() coregraphics.Float {
	result_ := C.C_NSEvent_DeltaY(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSEvent) DeltaZ() coregraphics.Float {
	result_ := C.C_NSEvent_DeltaZ(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSEvent) Pressure() float32 {
	result_ := C.C_NSEvent_Pressure(n.Ptr())
	return float32(result_)
}

func (n NSEvent) Stage() int {
	result_ := C.C_NSEvent_Stage(n.Ptr())
	return int(result_)
}

func (n NSEvent) StageTransition() coregraphics.Float {
	result_ := C.C_NSEvent_StageTransition(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSEvent) PressureBehavior() PressureBehavior {
	result_ := C.C_NSEvent_PressureBehavior(n.Ptr())
	return PressureBehavior(int(result_))
}

func (n NSEvent) CapabilityMask() uint {
	result_ := C.C_NSEvent_CapabilityMask(n.Ptr())
	return uint(result_)
}

func (n NSEvent) DeviceID() uint {
	result_ := C.C_NSEvent_DeviceID(n.Ptr())
	return uint(result_)
}

func (n NSEvent) IsEnteringProximity() bool {
	result_ := C.C_NSEvent_IsEnteringProximity(n.Ptr())
	return bool(result_)
}

func (n NSEvent) PointingDeviceID() uint {
	result_ := C.C_NSEvent_PointingDeviceID(n.Ptr())
	return uint(result_)
}

func (n NSEvent) PointingDeviceSerialNumber() uint {
	result_ := C.C_NSEvent_PointingDeviceSerialNumber(n.Ptr())
	return uint(result_)
}

func (n NSEvent) PointingDeviceType() PointingDeviceType {
	result_ := C.C_NSEvent_PointingDeviceType(n.Ptr())
	return PointingDeviceType(uint(result_))
}

func (n NSEvent) SystemTabletID() uint {
	result_ := C.C_NSEvent_SystemTabletID(n.Ptr())
	return uint(result_)
}

func (n NSEvent) TabletID() uint {
	result_ := C.C_NSEvent_TabletID(n.Ptr())
	return uint(result_)
}

func (n NSEvent) VendorID() uint {
	result_ := C.C_NSEvent_VendorID(n.Ptr())
	return uint(result_)
}

func (n NSEvent) VendorPointingDeviceType() uint {
	result_ := C.C_NSEvent_VendorPointingDeviceType(n.Ptr())
	return uint(result_)
}

func (n NSEvent) AbsoluteX() int {
	result_ := C.C_NSEvent_AbsoluteX(n.Ptr())
	return int(result_)
}

func (n NSEvent) AbsoluteY() int {
	result_ := C.C_NSEvent_AbsoluteY(n.Ptr())
	return int(result_)
}

func (n NSEvent) AbsoluteZ() int {
	result_ := C.C_NSEvent_AbsoluteZ(n.Ptr())
	return int(result_)
}

func (n NSEvent) ButtonMask() EventButtonMask {
	result_ := C.C_NSEvent_ButtonMask(n.Ptr())
	return EventButtonMask(uint(result_))
}

func (n NSEvent) Rotation() float32 {
	result_ := C.C_NSEvent_Rotation(n.Ptr())
	return float32(result_)
}

func (n NSEvent) TangentialPressure() float32 {
	result_ := C.C_NSEvent_TangentialPressure(n.Ptr())
	return float32(result_)
}

func (n NSEvent) Tilt() foundation.Point {
	result_ := C.C_NSEvent_Tilt(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSEvent) VendorDefined() objc.Object {
	result_ := C.C_NSEvent_VendorDefined(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSEvent) Magnification() coregraphics.Float {
	result_ := C.C_NSEvent_Magnification(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func Event_MouseCoalescingEnabled() bool {
	result_ := C.C_NSEvent_Event_MouseCoalescingEnabled()
	return bool(result_)
}

func Event_SetMouseCoalescingEnabled(value bool) {
	C.C_NSEvent_Event_SetMouseCoalescingEnabled(C.bool(value))
}

func (n NSEvent) HasPreciseScrollingDeltas() bool {
	result_ := C.C_NSEvent_HasPreciseScrollingDeltas(n.Ptr())
	return bool(result_)
}

func (n NSEvent) ScrollingDeltaX() coregraphics.Float {
	result_ := C.C_NSEvent_ScrollingDeltaX(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSEvent) ScrollingDeltaY() coregraphics.Float {
	result_ := C.C_NSEvent_ScrollingDeltaY(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSEvent) MomentumPhase() EventPhase {
	result_ := C.C_NSEvent_MomentumPhase(n.Ptr())
	return EventPhase(uint(result_))
}

func (n NSEvent) Phase() EventPhase {
	result_ := C.C_NSEvent_Phase(n.Ptr())
	return EventPhase(uint(result_))
}

func Event_SwipeTrackingFromScrollEventsEnabled() bool {
	result_ := C.C_NSEvent_Event_SwipeTrackingFromScrollEventsEnabled()
	return bool(result_)
}

func (n NSEvent) IsDirectionInvertedFromDevice() bool {
	result_ := C.C_NSEvent_IsDirectionInvertedFromDevice(n.Ptr())
	return bool(result_)
}
