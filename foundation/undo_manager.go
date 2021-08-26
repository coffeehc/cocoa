package foundation

// #include "undo_manager.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type UndoManager interface {
	objc.Object
	RegisterUndoWithTarget_Selector_Object(target objc.Object, selector objc.Selector, anObject objc.Object)
	PrepareWithInvocationTarget(target objc.Object) objc.Object
	Undo()
	UndoNestedGroup()
	Redo()
	BeginUndoGrouping()
	EndUndoGrouping()
	DisableUndoRegistration()
	EnableUndoRegistration()
	RemoveAllActions()
	RemoveAllActionsWithTarget(target objc.Object)
	SetActionName(actionName string)
	UndoMenuTitleForUndoActionName(actionName string) string
	RedoMenuTitleForUndoActionName(actionName string) string
	SetActionIsDiscardable(discardable bool)
	CanUndo() bool
	CanRedo() bool
	LevelsOfUndo() uint
	SetLevelsOfUndo(value uint)
	GroupsByEvent() bool
	SetGroupsByEvent(value bool)
	GroupingLevel() int
	IsUndoRegistrationEnabled() bool
	IsUndoing() bool
	IsRedoing() bool
	UndoActionName() string
	RedoActionName() string
	UndoMenuItemTitle() string
	RedoMenuItemTitle() string
	RunLoopModes() []RunLoopMode
	SetRunLoopModes(value []RunLoopMode)
	UndoActionIsDiscardable() bool
	RedoActionIsDiscardable() bool
}

type NSUndoManager struct {
	objc.NSObject
}

func MakeUndoManager(ptr unsafe.Pointer) NSUndoManager {
	return NSUndoManager{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocUndoManager() NSUndoManager {
	return MakeUndoManager(C.C_UndoManager_Alloc())
}

func (n NSUndoManager) Init() UndoManager {
	result_ := C.C_NSUndoManager_Init(n.Ptr())
	return MakeUndoManager(result_)
}

func (n NSUndoManager) RegisterUndoWithTarget_Selector_Object(target objc.Object, selector objc.Selector, anObject objc.Object) {
	C.C_NSUndoManager_RegisterUndoWithTarget_Selector_Object(n.Ptr(), objc.ExtractPtr(target), unsafe.Pointer(selector), objc.ExtractPtr(anObject))
}

func (n NSUndoManager) PrepareWithInvocationTarget(target objc.Object) objc.Object {
	result_ := C.C_NSUndoManager_PrepareWithInvocationTarget(n.Ptr(), objc.ExtractPtr(target))
	return objc.MakeObject(result_)
}

func (n NSUndoManager) Undo() {
	C.C_NSUndoManager_Undo(n.Ptr())
}

func (n NSUndoManager) UndoNestedGroup() {
	C.C_NSUndoManager_UndoNestedGroup(n.Ptr())
}

func (n NSUndoManager) Redo() {
	C.C_NSUndoManager_Redo(n.Ptr())
}

func (n NSUndoManager) BeginUndoGrouping() {
	C.C_NSUndoManager_BeginUndoGrouping(n.Ptr())
}

func (n NSUndoManager) EndUndoGrouping() {
	C.C_NSUndoManager_EndUndoGrouping(n.Ptr())
}

func (n NSUndoManager) DisableUndoRegistration() {
	C.C_NSUndoManager_DisableUndoRegistration(n.Ptr())
}

func (n NSUndoManager) EnableUndoRegistration() {
	C.C_NSUndoManager_EnableUndoRegistration(n.Ptr())
}

func (n NSUndoManager) RemoveAllActions() {
	C.C_NSUndoManager_RemoveAllActions(n.Ptr())
}

func (n NSUndoManager) RemoveAllActionsWithTarget(target objc.Object) {
	C.C_NSUndoManager_RemoveAllActionsWithTarget(n.Ptr(), objc.ExtractPtr(target))
}

func (n NSUndoManager) SetActionName(actionName string) {
	C.C_NSUndoManager_SetActionName(n.Ptr(), NewString(actionName).Ptr())
}

func (n NSUndoManager) UndoMenuTitleForUndoActionName(actionName string) string {
	result_ := C.C_NSUndoManager_UndoMenuTitleForUndoActionName(n.Ptr(), NewString(actionName).Ptr())
	return MakeString(result_).String()
}

func (n NSUndoManager) RedoMenuTitleForUndoActionName(actionName string) string {
	result_ := C.C_NSUndoManager_RedoMenuTitleForUndoActionName(n.Ptr(), NewString(actionName).Ptr())
	return MakeString(result_).String()
}

func (n NSUndoManager) SetActionIsDiscardable(discardable bool) {
	C.C_NSUndoManager_SetActionIsDiscardable(n.Ptr(), C.bool(discardable))
}

func (n NSUndoManager) CanUndo() bool {
	result_ := C.C_NSUndoManager_CanUndo(n.Ptr())
	return bool(result_)
}

func (n NSUndoManager) CanRedo() bool {
	result_ := C.C_NSUndoManager_CanRedo(n.Ptr())
	return bool(result_)
}

func (n NSUndoManager) LevelsOfUndo() uint {
	result_ := C.C_NSUndoManager_LevelsOfUndo(n.Ptr())
	return uint(result_)
}

func (n NSUndoManager) SetLevelsOfUndo(value uint) {
	C.C_NSUndoManager_SetLevelsOfUndo(n.Ptr(), C.uint(value))
}

func (n NSUndoManager) GroupsByEvent() bool {
	result_ := C.C_NSUndoManager_GroupsByEvent(n.Ptr())
	return bool(result_)
}

func (n NSUndoManager) SetGroupsByEvent(value bool) {
	C.C_NSUndoManager_SetGroupsByEvent(n.Ptr(), C.bool(value))
}

func (n NSUndoManager) GroupingLevel() int {
	result_ := C.C_NSUndoManager_GroupingLevel(n.Ptr())
	return int(result_)
}

func (n NSUndoManager) IsUndoRegistrationEnabled() bool {
	result_ := C.C_NSUndoManager_IsUndoRegistrationEnabled(n.Ptr())
	return bool(result_)
}

func (n NSUndoManager) IsUndoing() bool {
	result_ := C.C_NSUndoManager_IsUndoing(n.Ptr())
	return bool(result_)
}

func (n NSUndoManager) IsRedoing() bool {
	result_ := C.C_NSUndoManager_IsRedoing(n.Ptr())
	return bool(result_)
}

func (n NSUndoManager) UndoActionName() string {
	result_ := C.C_NSUndoManager_UndoActionName(n.Ptr())
	return MakeString(result_).String()
}

func (n NSUndoManager) RedoActionName() string {
	result_ := C.C_NSUndoManager_RedoActionName(n.Ptr())
	return MakeString(result_).String()
}

func (n NSUndoManager) UndoMenuItemTitle() string {
	result_ := C.C_NSUndoManager_UndoMenuItemTitle(n.Ptr())
	return MakeString(result_).String()
}

func (n NSUndoManager) RedoMenuItemTitle() string {
	result_ := C.C_NSUndoManager_RedoMenuItemTitle(n.Ptr())
	return MakeString(result_).String()
}

func (n NSUndoManager) RunLoopModes() []RunLoopMode {
	result_ := C.C_NSUndoManager_RunLoopModes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]RunLoopMode, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = RunLoopMode(MakeString(r).String())
	}
	return goResult_
}

func (n NSUndoManager) SetRunLoopModes(value []RunLoopMode) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(string(v)).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSUndoManager_SetRunLoopModes(n.Ptr(), cValue)
}

func (n NSUndoManager) UndoActionIsDiscardable() bool {
	result_ := C.C_NSUndoManager_UndoActionIsDiscardable(n.Ptr())
	return bool(result_)
}

func (n NSUndoManager) RedoActionIsDiscardable() bool {
	result_ := C.C_NSUndoManager_RedoActionIsDiscardable(n.Ptr())
	return bool(result_)
}
