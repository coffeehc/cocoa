package appkit

// #include "gesture_recognizer.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type GestureRecognizer interface {
	objc.Object
	LocationInView(view View) foundation.Point
	Reset()
	MouseDown(event Event)
	MouseDragged(event Event)
	MouseUp(event Event)
	OtherMouseDown(event Event)
	OtherMouseDragged(event Event)
	OtherMouseUp(event Event)
	RightMouseDown(event Event)
	RightMouseDragged(event Event)
	RightMouseUp(event Event)
	MagnifyWithEvent(event Event)
	RotateWithEvent(event Event)
	CanBePreventedByGestureRecognizer(preventingGestureRecognizer GestureRecognizer) bool
	CanPreventGestureRecognizer(preventedGestureRecognizer GestureRecognizer) bool
	ShouldBeRequiredToFailByGestureRecognizer(otherGestureRecognizer GestureRecognizer) bool
	ShouldRequireFailureOfGestureRecognizer(otherGestureRecognizer GestureRecognizer) bool
	KeyDown(event Event)
	KeyUp(event Event)
	TabletPoint(event Event)
	FlagsChanged(event Event)
	PressureChangeWithEvent(event Event)
	TouchesBeganWithEvent(event Event)
	TouchesCancelledWithEvent(event Event)
	TouchesEndedWithEvent(event Event)
	TouchesMovedWithEvent(event Event)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.Object)
	State() GestureRecognizerState
	View() View
	IsEnabled() bool
	SetEnabled(value bool)
	DelaysPrimaryMouseButtonEvents() bool
	SetDelaysPrimaryMouseButtonEvents(value bool)
	DelaysSecondaryMouseButtonEvents() bool
	SetDelaysSecondaryMouseButtonEvents(value bool)
	DelaysOtherMouseButtonEvents() bool
	SetDelaysOtherMouseButtonEvents(value bool)
	DelaysKeyEvents() bool
	SetDelaysKeyEvents(value bool)
	DelaysMagnificationEvents() bool
	SetDelaysMagnificationEvents(value bool)
	DelaysRotationEvents() bool
	SetDelaysRotationEvents(value bool)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	PressureConfiguration() PressureConfiguration
	SetPressureConfiguration(value PressureConfiguration)
	AllowedTouchTypes() TouchTypeMask
	SetAllowedTouchTypes(value TouchTypeMask)
}

type NSGestureRecognizer struct {
	objc.NSObject
}

func MakeGestureRecognizer(ptr unsafe.Pointer) NSGestureRecognizer {
	return NSGestureRecognizer{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSGestureRecognizer) InitWithTarget_Action(target objc.Object, action objc.Selector) NSGestureRecognizer {
	result_ := C.C_NSGestureRecognizer_InitWithTarget_Action(n.Ptr(), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakeGestureRecognizer(result_)
}

func (n NSGestureRecognizer) InitWithCoder(coder foundation.Coder) NSGestureRecognizer {
	result_ := C.C_NSGestureRecognizer_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeGestureRecognizer(result_)
}

func AllocGestureRecognizer() NSGestureRecognizer {
	result_ := C.C_NSGestureRecognizer_AllocGestureRecognizer()
	return MakeGestureRecognizer(result_)
}

func (n NSGestureRecognizer) Init() NSGestureRecognizer {
	result_ := C.C_NSGestureRecognizer_Init(n.Ptr())
	return MakeGestureRecognizer(result_)
}

func NewGestureRecognizer() NSGestureRecognizer {
	result_ := C.C_NSGestureRecognizer_NewGestureRecognizer()
	return MakeGestureRecognizer(result_)
}

func (n NSGestureRecognizer) Autorelease() NSGestureRecognizer {
	result_ := C.C_NSGestureRecognizer_Autorelease(n.Ptr())
	return MakeGestureRecognizer(result_)
}

func (n NSGestureRecognizer) Retain() NSGestureRecognizer {
	result_ := C.C_NSGestureRecognizer_Retain(n.Ptr())
	return MakeGestureRecognizer(result_)
}

func (n NSGestureRecognizer) LocationInView(view View) foundation.Point {
	result_ := C.C_NSGestureRecognizer_LocationInView(n.Ptr(), objc.ExtractPtr(view))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSGestureRecognizer) Reset() {
	C.C_NSGestureRecognizer_Reset(n.Ptr())
}

func (n NSGestureRecognizer) MouseDown(event Event) {
	C.C_NSGestureRecognizer_MouseDown(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) MouseDragged(event Event) {
	C.C_NSGestureRecognizer_MouseDragged(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) MouseUp(event Event) {
	C.C_NSGestureRecognizer_MouseUp(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) OtherMouseDown(event Event) {
	C.C_NSGestureRecognizer_OtherMouseDown(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) OtherMouseDragged(event Event) {
	C.C_NSGestureRecognizer_OtherMouseDragged(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) OtherMouseUp(event Event) {
	C.C_NSGestureRecognizer_OtherMouseUp(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) RightMouseDown(event Event) {
	C.C_NSGestureRecognizer_RightMouseDown(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) RightMouseDragged(event Event) {
	C.C_NSGestureRecognizer_RightMouseDragged(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) RightMouseUp(event Event) {
	C.C_NSGestureRecognizer_RightMouseUp(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) MagnifyWithEvent(event Event) {
	C.C_NSGestureRecognizer_MagnifyWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) RotateWithEvent(event Event) {
	C.C_NSGestureRecognizer_RotateWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) CanBePreventedByGestureRecognizer(preventingGestureRecognizer GestureRecognizer) bool {
	result_ := C.C_NSGestureRecognizer_CanBePreventedByGestureRecognizer(n.Ptr(), objc.ExtractPtr(preventingGestureRecognizer))
	return bool(result_)
}

func (n NSGestureRecognizer) CanPreventGestureRecognizer(preventedGestureRecognizer GestureRecognizer) bool {
	result_ := C.C_NSGestureRecognizer_CanPreventGestureRecognizer(n.Ptr(), objc.ExtractPtr(preventedGestureRecognizer))
	return bool(result_)
}

func (n NSGestureRecognizer) ShouldBeRequiredToFailByGestureRecognizer(otherGestureRecognizer GestureRecognizer) bool {
	result_ := C.C_NSGestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(n.Ptr(), objc.ExtractPtr(otherGestureRecognizer))
	return bool(result_)
}

func (n NSGestureRecognizer) ShouldRequireFailureOfGestureRecognizer(otherGestureRecognizer GestureRecognizer) bool {
	result_ := C.C_NSGestureRecognizer_ShouldRequireFailureOfGestureRecognizer(n.Ptr(), objc.ExtractPtr(otherGestureRecognizer))
	return bool(result_)
}

func (n NSGestureRecognizer) KeyDown(event Event) {
	C.C_NSGestureRecognizer_KeyDown(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) KeyUp(event Event) {
	C.C_NSGestureRecognizer_KeyUp(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) TabletPoint(event Event) {
	C.C_NSGestureRecognizer_TabletPoint(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) FlagsChanged(event Event) {
	C.C_NSGestureRecognizer_FlagsChanged(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) PressureChangeWithEvent(event Event) {
	C.C_NSGestureRecognizer_PressureChangeWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) TouchesBeganWithEvent(event Event) {
	C.C_NSGestureRecognizer_TouchesBeganWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) TouchesCancelledWithEvent(event Event) {
	C.C_NSGestureRecognizer_TouchesCancelledWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) TouchesEndedWithEvent(event Event) {
	C.C_NSGestureRecognizer_TouchesEndedWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) TouchesMovedWithEvent(event Event) {
	C.C_NSGestureRecognizer_TouchesMovedWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSGestureRecognizer) Action() objc.Selector {
	result_ := C.C_NSGestureRecognizer_Action(n.Ptr())
	return objc.Selector(result_)
}

func (n NSGestureRecognizer) SetAction(value objc.Selector) {
	C.C_NSGestureRecognizer_SetAction(n.Ptr(), unsafe.Pointer(value))
}

func (n NSGestureRecognizer) Target() objc.Object {
	result_ := C.C_NSGestureRecognizer_Target(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSGestureRecognizer) SetTarget(value objc.Object) {
	C.C_NSGestureRecognizer_SetTarget(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSGestureRecognizer) State() GestureRecognizerState {
	result_ := C.C_NSGestureRecognizer_State(n.Ptr())
	return GestureRecognizerState(int(result_))
}

func (n NSGestureRecognizer) View() View {
	result_ := C.C_NSGestureRecognizer_View(n.Ptr())
	return MakeView(result_)
}

func (n NSGestureRecognizer) IsEnabled() bool {
	result_ := C.C_NSGestureRecognizer_IsEnabled(n.Ptr())
	return bool(result_)
}

func (n NSGestureRecognizer) SetEnabled(value bool) {
	C.C_NSGestureRecognizer_SetEnabled(n.Ptr(), C.bool(value))
}

func (n NSGestureRecognizer) DelaysPrimaryMouseButtonEvents() bool {
	result_ := C.C_NSGestureRecognizer_DelaysPrimaryMouseButtonEvents(n.Ptr())
	return bool(result_)
}

func (n NSGestureRecognizer) SetDelaysPrimaryMouseButtonEvents(value bool) {
	C.C_NSGestureRecognizer_SetDelaysPrimaryMouseButtonEvents(n.Ptr(), C.bool(value))
}

func (n NSGestureRecognizer) DelaysSecondaryMouseButtonEvents() bool {
	result_ := C.C_NSGestureRecognizer_DelaysSecondaryMouseButtonEvents(n.Ptr())
	return bool(result_)
}

func (n NSGestureRecognizer) SetDelaysSecondaryMouseButtonEvents(value bool) {
	C.C_NSGestureRecognizer_SetDelaysSecondaryMouseButtonEvents(n.Ptr(), C.bool(value))
}

func (n NSGestureRecognizer) DelaysOtherMouseButtonEvents() bool {
	result_ := C.C_NSGestureRecognizer_DelaysOtherMouseButtonEvents(n.Ptr())
	return bool(result_)
}

func (n NSGestureRecognizer) SetDelaysOtherMouseButtonEvents(value bool) {
	C.C_NSGestureRecognizer_SetDelaysOtherMouseButtonEvents(n.Ptr(), C.bool(value))
}

func (n NSGestureRecognizer) DelaysKeyEvents() bool {
	result_ := C.C_NSGestureRecognizer_DelaysKeyEvents(n.Ptr())
	return bool(result_)
}

func (n NSGestureRecognizer) SetDelaysKeyEvents(value bool) {
	C.C_NSGestureRecognizer_SetDelaysKeyEvents(n.Ptr(), C.bool(value))
}

func (n NSGestureRecognizer) DelaysMagnificationEvents() bool {
	result_ := C.C_NSGestureRecognizer_DelaysMagnificationEvents(n.Ptr())
	return bool(result_)
}

func (n NSGestureRecognizer) SetDelaysMagnificationEvents(value bool) {
	C.C_NSGestureRecognizer_SetDelaysMagnificationEvents(n.Ptr(), C.bool(value))
}

func (n NSGestureRecognizer) DelaysRotationEvents() bool {
	result_ := C.C_NSGestureRecognizer_DelaysRotationEvents(n.Ptr())
	return bool(result_)
}

func (n NSGestureRecognizer) SetDelaysRotationEvents(value bool) {
	C.C_NSGestureRecognizer_SetDelaysRotationEvents(n.Ptr(), C.bool(value))
}

func (n NSGestureRecognizer) Delegate() objc.Object {
	result_ := C.C_NSGestureRecognizer_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSGestureRecognizer) SetDelegate(value objc.Object) {
	C.C_NSGestureRecognizer_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSGestureRecognizer) PressureConfiguration() PressureConfiguration {
	result_ := C.C_NSGestureRecognizer_PressureConfiguration(n.Ptr())
	return MakePressureConfiguration(result_)
}

func (n NSGestureRecognizer) SetPressureConfiguration(value PressureConfiguration) {
	C.C_NSGestureRecognizer_SetPressureConfiguration(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSGestureRecognizer) AllowedTouchTypes() TouchTypeMask {
	result_ := C.C_NSGestureRecognizer_AllowedTouchTypes(n.Ptr())
	return TouchTypeMask(uint(result_))
}

func (n NSGestureRecognizer) SetAllowedTouchTypes(value TouchTypeMask) {
	C.C_NSGestureRecognizer_SetAllowedTouchTypes(n.Ptr(), C.uint(uint(value)))
}
