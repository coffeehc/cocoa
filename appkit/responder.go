package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "responder.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Responder interface {
	objc.Object
	BecomeFirstResponder() bool
	ResignFirstResponder() bool
	ValidateProposedFirstResponder_ForEvent(responder Responder, event Event) bool
	MouseDown(event Event)
	MouseDragged(event Event)
	MouseUp(event Event)
	MouseMoved(event Event)
	MouseEntered(event Event)
	MouseExited(event Event)
	RightMouseDown(event Event)
	RightMouseDragged(event Event)
	RightMouseUp(event Event)
	OtherMouseDown(event Event)
	OtherMouseDragged(event Event)
	OtherMouseUp(event Event)
	KeyDown(event Event)
	KeyUp(event Event)
	InterpretKeyEvents(eventArray []Event)
	PerformKeyEquivalent(event Event) bool
	FlushBufferedKeyEvents()
	PressureChangeWithEvent(event Event)
	CursorUpdate(event Event)
	FlagsChanged(event Event)
	TabletPoint(event Event)
	TabletProximity(event Event)
	HelpRequested(eventPtr Event)
	ScrollWheel(event Event)
	QuickLookWithEvent(event Event)
	ChangeModeWithEvent(event Event)
	SupplementalTargetForAction_Sender(action *objc.Selector, sender objc.Object) objc.Object
	EncodeRestorableStateWithCoder(coder foundation.Coder)
	RestoreStateWithCoder(coder foundation.Coder)
	InvalidateRestorableState()
	UpdateUserActivityState(userActivity foundation.UserActivity)
	PresentError(error foundation.Error) bool
	WillPresentError(error foundation.Error) foundation.Error
	TryToPerform_With(action *objc.Selector, object objc.Object) bool
	ValidRequestorForSendType_ReturnType(sendType PasteboardType, returnType PasteboardType) objc.Object
	ShouldBeTreatedAsInkEvent(event Event) bool
	NoResponderFor(eventSelector *objc.Selector)
	BeginGestureWithEvent(event Event)
	EndGestureWithEvent(event Event)
	MagnifyWithEvent(event Event)
	RotateWithEvent(event Event)
	SwipeWithEvent(event Event)
	TouchesBeganWithEvent(event Event)
	TouchesMovedWithEvent(event Event)
	TouchesCancelledWithEvent(event Event)
	TouchesEndedWithEvent(event Event)
	WantsForwardedScrollEventsForAxis(axis EventGestureAxis) bool
	SmartMagnifyWithEvent(event Event)
	WantsScrollEventsForSwipeTrackingOnAxis(axis EventGestureAxis) bool
	PerformTextFinderAction(sender objc.Object)
	EncodeRestorableStateWithCoder_BackgroundQueue(coder foundation.Coder, queue foundation.OperationQueue)
	MakeTouchBar() TouchBar
	NewWindowForTab(sender objc.Object)
	ShowContextHelp(sender objc.Object)
	AcceptsFirstResponder() bool
	NextResponder() Responder
	SetNextResponder(value Responder)
	UserActivity() foundation.UserActivity
	SetUserActivity(value foundation.UserActivity)
	Menu() Menu
	SetMenu(value Menu)
	UndoManager() foundation.UndoManager
	TouchBar() TouchBar
	SetTouchBar(value TouchBar)
}

type NSResponder struct {
	objc.NSObject
}

func MakeResponder(ptr unsafe.Pointer) *NSResponder {
	if ptr == nil {
		return nil
	}
	return &NSResponder{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocResponder() *NSResponder {
	return MakeResponder(C.C_Responder_Alloc())
}

func (n *NSResponder) Init() Responder {
	result_ := C.C_NSResponder_Init(n.Ptr())
	return MakeResponder(result_)
}

func (n *NSResponder) InitWithCoder(coder foundation.Coder) Responder {
	result_ := C.C_NSResponder_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeResponder(result_)
}

func (n *NSResponder) BecomeFirstResponder() bool {
	result_ := C.C_NSResponder_BecomeFirstResponder(n.Ptr())
	return bool(result_)
}

func (n *NSResponder) ResignFirstResponder() bool {
	result_ := C.C_NSResponder_ResignFirstResponder(n.Ptr())
	return bool(result_)
}

func (n *NSResponder) ValidateProposedFirstResponder_ForEvent(responder Responder, event Event) bool {
	result_ := C.C_NSResponder_ValidateProposedFirstResponder_ForEvent(n.Ptr(), objc.ExtractPtr(responder), objc.ExtractPtr(event))
	return bool(result_)
}

func (n *NSResponder) MouseDown(event Event) {
	C.C_NSResponder_MouseDown(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) MouseDragged(event Event) {
	C.C_NSResponder_MouseDragged(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) MouseUp(event Event) {
	C.C_NSResponder_MouseUp(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) MouseMoved(event Event) {
	C.C_NSResponder_MouseMoved(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) MouseEntered(event Event) {
	C.C_NSResponder_MouseEntered(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) MouseExited(event Event) {
	C.C_NSResponder_MouseExited(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) RightMouseDown(event Event) {
	C.C_NSResponder_RightMouseDown(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) RightMouseDragged(event Event) {
	C.C_NSResponder_RightMouseDragged(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) RightMouseUp(event Event) {
	C.C_NSResponder_RightMouseUp(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) OtherMouseDown(event Event) {
	C.C_NSResponder_OtherMouseDown(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) OtherMouseDragged(event Event) {
	C.C_NSResponder_OtherMouseDragged(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) OtherMouseUp(event Event) {
	C.C_NSResponder_OtherMouseUp(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) KeyDown(event Event) {
	C.C_NSResponder_KeyDown(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) KeyUp(event Event) {
	C.C_NSResponder_KeyUp(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) InterpretKeyEvents(eventArray []Event) {
	cEventArrayData := make([]unsafe.Pointer, len(eventArray))
	for idx, v := range eventArray {
		cEventArrayData[idx] = objc.ExtractPtr(v)
	}
	cEventArray := C.Array{data: unsafe.Pointer(&cEventArrayData[0]), len: C.int(len(eventArray))}
	C.C_NSResponder_InterpretKeyEvents(n.Ptr(), cEventArray)
}

func (n *NSResponder) PerformKeyEquivalent(event Event) bool {
	result_ := C.C_NSResponder_PerformKeyEquivalent(n.Ptr(), objc.ExtractPtr(event))
	return bool(result_)
}

func (n *NSResponder) FlushBufferedKeyEvents() {
	C.C_NSResponder_FlushBufferedKeyEvents(n.Ptr())
}

func (n *NSResponder) PressureChangeWithEvent(event Event) {
	C.C_NSResponder_PressureChangeWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) CursorUpdate(event Event) {
	C.C_NSResponder_CursorUpdate(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) FlagsChanged(event Event) {
	C.C_NSResponder_FlagsChanged(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) TabletPoint(event Event) {
	C.C_NSResponder_TabletPoint(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) TabletProximity(event Event) {
	C.C_NSResponder_TabletProximity(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) HelpRequested(eventPtr Event) {
	C.C_NSResponder_HelpRequested(n.Ptr(), objc.ExtractPtr(eventPtr))
}

func (n *NSResponder) ScrollWheel(event Event) {
	C.C_NSResponder_ScrollWheel(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) QuickLookWithEvent(event Event) {
	C.C_NSResponder_QuickLookWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) ChangeModeWithEvent(event Event) {
	C.C_NSResponder_ChangeModeWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) SupplementalTargetForAction_Sender(action *objc.Selector, sender objc.Object) objc.Object {
	result_ := C.C_NSResponder_SupplementalTargetForAction_Sender(n.Ptr(), objc.ExtractPtr(action), objc.ExtractPtr(sender))
	return objc.MakeObject(result_)
}

func (n *NSResponder) EncodeRestorableStateWithCoder(coder foundation.Coder) {
	C.C_NSResponder_EncodeRestorableStateWithCoder(n.Ptr(), objc.ExtractPtr(coder))
}

func (n *NSResponder) RestoreStateWithCoder(coder foundation.Coder) {
	C.C_NSResponder_RestoreStateWithCoder(n.Ptr(), objc.ExtractPtr(coder))
}

func (n *NSResponder) InvalidateRestorableState() {
	C.C_NSResponder_InvalidateRestorableState(n.Ptr())
}

func (n *NSResponder) UpdateUserActivityState(userActivity foundation.UserActivity) {
	C.C_NSResponder_UpdateUserActivityState(n.Ptr(), objc.ExtractPtr(userActivity))
}

func (n *NSResponder) PresentError(error foundation.Error) bool {
	result_ := C.C_NSResponder_PresentError(n.Ptr(), objc.ExtractPtr(error))
	return bool(result_)
}

func (n *NSResponder) WillPresentError(error foundation.Error) foundation.Error {
	result_ := C.C_NSResponder_WillPresentError(n.Ptr(), objc.ExtractPtr(error))
	return foundation.MakeError(result_)
}

func (n *NSResponder) TryToPerform_With(action *objc.Selector, object objc.Object) bool {
	result_ := C.C_NSResponder_TryToPerform_With(n.Ptr(), objc.ExtractPtr(action), objc.ExtractPtr(object))
	return bool(result_)
}

func (n *NSResponder) ValidRequestorForSendType_ReturnType(sendType PasteboardType, returnType PasteboardType) objc.Object {
	result_ := C.C_NSResponder_ValidRequestorForSendType_ReturnType(n.Ptr(), foundation.NewString(string(sendType)).Ptr(), foundation.NewString(string(returnType)).Ptr())
	return objc.MakeObject(result_)
}

func (n *NSResponder) ShouldBeTreatedAsInkEvent(event Event) bool {
	result_ := C.C_NSResponder_ShouldBeTreatedAsInkEvent(n.Ptr(), objc.ExtractPtr(event))
	return bool(result_)
}

func (n *NSResponder) NoResponderFor(eventSelector *objc.Selector) {
	C.C_NSResponder_NoResponderFor(n.Ptr(), objc.ExtractPtr(eventSelector))
}

func (n *NSResponder) BeginGestureWithEvent(event Event) {
	C.C_NSResponder_BeginGestureWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) EndGestureWithEvent(event Event) {
	C.C_NSResponder_EndGestureWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) MagnifyWithEvent(event Event) {
	C.C_NSResponder_MagnifyWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) RotateWithEvent(event Event) {
	C.C_NSResponder_RotateWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) SwipeWithEvent(event Event) {
	C.C_NSResponder_SwipeWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) TouchesBeganWithEvent(event Event) {
	C.C_NSResponder_TouchesBeganWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) TouchesMovedWithEvent(event Event) {
	C.C_NSResponder_TouchesMovedWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) TouchesCancelledWithEvent(event Event) {
	C.C_NSResponder_TouchesCancelledWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) TouchesEndedWithEvent(event Event) {
	C.C_NSResponder_TouchesEndedWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) WantsForwardedScrollEventsForAxis(axis EventGestureAxis) bool {
	result_ := C.C_NSResponder_WantsForwardedScrollEventsForAxis(n.Ptr(), C.int(int(axis)))
	return bool(result_)
}

func (n *NSResponder) SmartMagnifyWithEvent(event Event) {
	C.C_NSResponder_SmartMagnifyWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSResponder) WantsScrollEventsForSwipeTrackingOnAxis(axis EventGestureAxis) bool {
	result_ := C.C_NSResponder_WantsScrollEventsForSwipeTrackingOnAxis(n.Ptr(), C.int(int(axis)))
	return bool(result_)
}

func (n *NSResponder) PerformTextFinderAction(sender objc.Object) {
	C.C_NSResponder_PerformTextFinderAction(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSResponder) EncodeRestorableStateWithCoder_BackgroundQueue(coder foundation.Coder, queue foundation.OperationQueue) {
	C.C_NSResponder_EncodeRestorableStateWithCoder_BackgroundQueue(n.Ptr(), objc.ExtractPtr(coder), objc.ExtractPtr(queue))
}

func (n *NSResponder) MakeTouchBar() TouchBar {
	result_ := C.C_NSResponder_MakeTouchBar(n.Ptr())
	return MakeTouchBar(result_)
}

func (n *NSResponder) NewWindowForTab(sender objc.Object) {
	C.C_NSResponder_NewWindowForTab(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSResponder) ShowContextHelp(sender objc.Object) {
	C.C_NSResponder_ShowContextHelp(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSResponder) AcceptsFirstResponder() bool {
	result_ := C.C_NSResponder_AcceptsFirstResponder(n.Ptr())
	return bool(result_)
}

func (n *NSResponder) NextResponder() Responder {
	result_ := C.C_NSResponder_NextResponder(n.Ptr())
	return MakeResponder(result_)
}

func (n *NSResponder) SetNextResponder(value Responder) {
	C.C_NSResponder_SetNextResponder(n.Ptr(), objc.ExtractPtr(value))
}

func Responder_RestorableStateKeyPaths() []string {
	result_ := C.C_NSResponder_Responder_RestorableStateKeyPaths()
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n *NSResponder) UserActivity() foundation.UserActivity {
	result_ := C.C_NSResponder_UserActivity(n.Ptr())
	return foundation.MakeUserActivity(result_)
}

func (n *NSResponder) SetUserActivity(value foundation.UserActivity) {
	C.C_NSResponder_SetUserActivity(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSResponder) Menu() Menu {
	result_ := C.C_NSResponder_Menu(n.Ptr())
	return MakeMenu(result_)
}

func (n *NSResponder) SetMenu(value Menu) {
	C.C_NSResponder_SetMenu(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSResponder) UndoManager() foundation.UndoManager {
	result_ := C.C_NSResponder_UndoManager(n.Ptr())
	return foundation.MakeUndoManager(result_)
}

func (n *NSResponder) TouchBar() TouchBar {
	result_ := C.C_NSResponder_TouchBar(n.Ptr())
	return MakeTouchBar(result_)
}

func (n *NSResponder) SetTouchBar(value TouchBar) {
	C.C_NSResponder_SetTouchBar(n.Ptr(), objc.ExtractPtr(value))
}
