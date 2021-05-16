package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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
	CoalescedTouchesForTouch(touch Touch) []Touch
	CharactersByApplyingModifiers(modifiers EventModifierFlags) string
	Type() EventType
	ModifierFlags() EventModifierFlags
	LocationInWindow() foundation.Point
	Timestamp() foundation.TimeInterval
	Window() Window
	WindowNumber() int
	CGEvent() coregraphics.EventRef
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

func MakeEvent(ptr unsafe.Pointer) *NSEvent {
	if ptr == nil {
		return nil
	}
	return &NSEvent{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocEvent() *NSEvent {
	return MakeEvent(C.C_Event_Alloc())
}

func (n *NSEvent) Init() Event {
	result := C.C_NSEvent_Init(n.Ptr())
	return MakeEvent(result)
}

func Event_MouseEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_EventNumber_ClickCount_Pressure(_type EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil GraphicsContext, eNum int, cNum int, pressure float32) Event {
	result := C.C_NSEvent_Event_MouseEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_EventNumber_ClickCount_Pressure(C.uint(uint(_type)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(location))), C.uint(uint(flags)), C.double(float64(time)), C.int(wNum), objc.ExtractPtr(unusedPassNil), C.int(eNum), C.int(cNum), C.float(pressure))
	return MakeEvent(result)
}

func EventWithCGEvent(cgEvent coregraphics.EventRef) Event {
	result := C.C_NSEvent_EventWithCGEvent(*(*C.CGEventRef)(coregraphics.ToCGEventRefPointer(cgEvent)))
	return MakeEvent(result)
}

func Event_StartPeriodicEventsAfterDelay_WithPeriod(delay foundation.TimeInterval, period foundation.TimeInterval) {
	C.C_NSEvent_Event_StartPeriodicEventsAfterDelay_WithPeriod(C.double(float64(delay)), C.double(float64(period)))
}

func Event_StopPeriodicEvents() {
	C.C_NSEvent_Event_StopPeriodicEvents()
}

func (n *NSEvent) CoalescedTouchesForTouch(touch Touch) []Touch {
	result := C.C_NSEvent_CoalescedTouchesForTouch(n.Ptr(), objc.ExtractPtr(touch))
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]Touch, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeTouch(r)
	}
	return goResult
}

func Event_RemoveMonitor(eventMonitor objc.Object) {
	C.C_NSEvent_Event_RemoveMonitor(objc.ExtractPtr(eventMonitor))
}

func (n *NSEvent) CharactersByApplyingModifiers(modifiers EventModifierFlags) string {
	result := C.C_NSEvent_CharactersByApplyingModifiers(n.Ptr(), C.uint(uint(modifiers)))
	return foundation.MakeString(result).String()
}

func (n *NSEvent) Type() EventType {
	result := C.C_NSEvent_Type(n.Ptr())
	return EventType(uint(result))
}

func (n *NSEvent) ModifierFlags() EventModifierFlags {
	result := C.C_NSEvent_ModifierFlags(n.Ptr())
	return EventModifierFlags(uint(result))
}

func (n *NSEvent) LocationInWindow() foundation.Point {
	result := C.C_NSEvent_LocationInWindow(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSEvent) Timestamp() foundation.TimeInterval {
	result := C.C_NSEvent_Timestamp(n.Ptr())
	return foundation.TimeInterval(float64(result))
}

func (n *NSEvent) Window() Window {
	result := C.C_NSEvent_Window(n.Ptr())
	return MakeWindow(result)
}

func (n *NSEvent) WindowNumber() int {
	result := C.C_NSEvent_WindowNumber(n.Ptr())
	return int(result)
}

func (n *NSEvent) CGEvent() coregraphics.EventRef {
	result := C.C_NSEvent_CGEvent(n.Ptr())
	return coregraphics.FromCGEventRefPointer(unsafe.Pointer(&result))
}

func Event_KeyRepeatDelay() foundation.TimeInterval {
	result := C.C_NSEvent_Event_KeyRepeatDelay()
	return foundation.TimeInterval(float64(result))
}

func Event_KeyRepeatInterval() foundation.TimeInterval {
	result := C.C_NSEvent_Event_KeyRepeatInterval()
	return foundation.TimeInterval(float64(result))
}

func (n *NSEvent) Characters() string {
	result := C.C_NSEvent_Characters(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSEvent) CharactersIgnoringModifiers() string {
	result := C.C_NSEvent_CharactersIgnoringModifiers(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSEvent) IsARepeat() bool {
	result := C.C_NSEvent_IsARepeat(n.Ptr())
	return bool(result)
}

func Event_PressedMouseButtons() uint {
	result := C.C_NSEvent_Event_PressedMouseButtons()
	return uint(result)
}

func Event_DoubleClickInterval() foundation.TimeInterval {
	result := C.C_NSEvent_Event_DoubleClickInterval()
	return foundation.TimeInterval(float64(result))
}

func Event_MouseLocation() foundation.Point {
	result := C.C_NSEvent_Event_MouseLocation()
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSEvent) ButtonNumber() int {
	result := C.C_NSEvent_ButtonNumber(n.Ptr())
	return int(result)
}

func (n *NSEvent) ClickCount() int {
	result := C.C_NSEvent_ClickCount(n.Ptr())
	return int(result)
}

func (n *NSEvent) EventNumber() int {
	result := C.C_NSEvent_EventNumber(n.Ptr())
	return int(result)
}

func (n *NSEvent) TrackingNumber() int {
	result := C.C_NSEvent_TrackingNumber(n.Ptr())
	return int(result)
}

func (n *NSEvent) TrackingArea() TrackingArea {
	result := C.C_NSEvent_TrackingArea(n.Ptr())
	return MakeTrackingArea(result)
}

func (n *NSEvent) Data1() int {
	result := C.C_NSEvent_Data1(n.Ptr())
	return int(result)
}

func (n *NSEvent) Data2() int {
	result := C.C_NSEvent_Data2(n.Ptr())
	return int(result)
}

func (n *NSEvent) DeltaX() coregraphics.Float {
	result := C.C_NSEvent_DeltaX(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSEvent) DeltaY() coregraphics.Float {
	result := C.C_NSEvent_DeltaY(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSEvent) DeltaZ() coregraphics.Float {
	result := C.C_NSEvent_DeltaZ(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSEvent) Pressure() float32 {
	result := C.C_NSEvent_Pressure(n.Ptr())
	return float32(result)
}

func (n *NSEvent) Stage() int {
	result := C.C_NSEvent_Stage(n.Ptr())
	return int(result)
}

func (n *NSEvent) StageTransition() coregraphics.Float {
	result := C.C_NSEvent_StageTransition(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSEvent) PressureBehavior() PressureBehavior {
	result := C.C_NSEvent_PressureBehavior(n.Ptr())
	return PressureBehavior(int(result))
}

func (n *NSEvent) CapabilityMask() uint {
	result := C.C_NSEvent_CapabilityMask(n.Ptr())
	return uint(result)
}

func (n *NSEvent) DeviceID() uint {
	result := C.C_NSEvent_DeviceID(n.Ptr())
	return uint(result)
}

func (n *NSEvent) IsEnteringProximity() bool {
	result := C.C_NSEvent_IsEnteringProximity(n.Ptr())
	return bool(result)
}

func (n *NSEvent) PointingDeviceID() uint {
	result := C.C_NSEvent_PointingDeviceID(n.Ptr())
	return uint(result)
}

func (n *NSEvent) PointingDeviceSerialNumber() uint {
	result := C.C_NSEvent_PointingDeviceSerialNumber(n.Ptr())
	return uint(result)
}

func (n *NSEvent) PointingDeviceType() PointingDeviceType {
	result := C.C_NSEvent_PointingDeviceType(n.Ptr())
	return PointingDeviceType(uint(result))
}

func (n *NSEvent) SystemTabletID() uint {
	result := C.C_NSEvent_SystemTabletID(n.Ptr())
	return uint(result)
}

func (n *NSEvent) TabletID() uint {
	result := C.C_NSEvent_TabletID(n.Ptr())
	return uint(result)
}

func (n *NSEvent) VendorID() uint {
	result := C.C_NSEvent_VendorID(n.Ptr())
	return uint(result)
}

func (n *NSEvent) VendorPointingDeviceType() uint {
	result := C.C_NSEvent_VendorPointingDeviceType(n.Ptr())
	return uint(result)
}

func (n *NSEvent) AbsoluteX() int {
	result := C.C_NSEvent_AbsoluteX(n.Ptr())
	return int(result)
}

func (n *NSEvent) AbsoluteY() int {
	result := C.C_NSEvent_AbsoluteY(n.Ptr())
	return int(result)
}

func (n *NSEvent) AbsoluteZ() int {
	result := C.C_NSEvent_AbsoluteZ(n.Ptr())
	return int(result)
}

func (n *NSEvent) ButtonMask() EventButtonMask {
	result := C.C_NSEvent_ButtonMask(n.Ptr())
	return EventButtonMask(uint(result))
}

func (n *NSEvent) Rotation() float32 {
	result := C.C_NSEvent_Rotation(n.Ptr())
	return float32(result)
}

func (n *NSEvent) TangentialPressure() float32 {
	result := C.C_NSEvent_TangentialPressure(n.Ptr())
	return float32(result)
}

func (n *NSEvent) Tilt() foundation.Point {
	result := C.C_NSEvent_Tilt(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSEvent) VendorDefined() objc.Object {
	result := C.C_NSEvent_VendorDefined(n.Ptr())
	return objc.MakeObject(result)
}

func (n *NSEvent) Magnification() coregraphics.Float {
	result := C.C_NSEvent_Magnification(n.Ptr())
	return coregraphics.Float(float64(result))
}

func Event_MouseCoalescingEnabled() bool {
	result := C.C_NSEvent_Event_MouseCoalescingEnabled()
	return bool(result)
}

func Event_SetMouseCoalescingEnabled(value bool) {
	C.C_NSEvent_Event_SetMouseCoalescingEnabled(C.bool(value))
}

func (n *NSEvent) HasPreciseScrollingDeltas() bool {
	result := C.C_NSEvent_HasPreciseScrollingDeltas(n.Ptr())
	return bool(result)
}

func (n *NSEvent) ScrollingDeltaX() coregraphics.Float {
	result := C.C_NSEvent_ScrollingDeltaX(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSEvent) ScrollingDeltaY() coregraphics.Float {
	result := C.C_NSEvent_ScrollingDeltaY(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSEvent) MomentumPhase() EventPhase {
	result := C.C_NSEvent_MomentumPhase(n.Ptr())
	return EventPhase(uint(result))
}

func (n *NSEvent) Phase() EventPhase {
	result := C.C_NSEvent_Phase(n.Ptr())
	return EventPhase(uint(result))
}

func Event_SwipeTrackingFromScrollEventsEnabled() bool {
	result := C.C_NSEvent_Event_SwipeTrackingFromScrollEventsEnabled()
	return bool(result)
}

func (n *NSEvent) IsDirectionInvertedFromDevice() bool {
	result := C.C_NSEvent_IsDirectionInvertedFromDevice(n.Ptr())
	return bool(result)
}
