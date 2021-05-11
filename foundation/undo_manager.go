package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "undo_manager.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type UndoManager interface {
	objc.Object
	RegisterUndoWithTarget_Selector_Object(target objc.Object, selector *objc.Selector, anObject objc.Object)
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
	UndoActionIsDiscardable() bool
	RedoActionIsDiscardable() bool
}

type NSUndoManager struct {
	objc.NSObject
}

func MakeUndoManager(ptr unsafe.Pointer) *NSUndoManager {
	if ptr == nil {
		return nil
	}
	return &NSUndoManager{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocUndoManager() *NSUndoManager {
	return MakeUndoManager(C.C_UndoManager_Alloc())
}

func (n *NSUndoManager) Init() UndoManager {
	result := C.C_NSUndoManager_Init(n.Ptr())
	return MakeUndoManager(result)
}

func (n *NSUndoManager) RegisterUndoWithTarget_Selector_Object(target objc.Object, selector *objc.Selector, anObject objc.Object) {
	C.C_NSUndoManager_RegisterUndoWithTarget_Selector_Object(n.Ptr(), objc.ExtractPtr(target), objc.ExtractPtr(selector), objc.ExtractPtr(anObject))
}

func (n *NSUndoManager) PrepareWithInvocationTarget(target objc.Object) objc.Object {
	result := C.C_NSUndoManager_PrepareWithInvocationTarget(n.Ptr(), objc.ExtractPtr(target))
	return objc.MakeObject(result)
}

func (n *NSUndoManager) Undo() {
	C.C_NSUndoManager_Undo(n.Ptr())
}

func (n *NSUndoManager) UndoNestedGroup() {
	C.C_NSUndoManager_UndoNestedGroup(n.Ptr())
}

func (n *NSUndoManager) Redo() {
	C.C_NSUndoManager_Redo(n.Ptr())
}

func (n *NSUndoManager) BeginUndoGrouping() {
	C.C_NSUndoManager_BeginUndoGrouping(n.Ptr())
}

func (n *NSUndoManager) EndUndoGrouping() {
	C.C_NSUndoManager_EndUndoGrouping(n.Ptr())
}

func (n *NSUndoManager) DisableUndoRegistration() {
	C.C_NSUndoManager_DisableUndoRegistration(n.Ptr())
}

func (n *NSUndoManager) EnableUndoRegistration() {
	C.C_NSUndoManager_EnableUndoRegistration(n.Ptr())
}

func (n *NSUndoManager) RemoveAllActions() {
	C.C_NSUndoManager_RemoveAllActions(n.Ptr())
}

func (n *NSUndoManager) RemoveAllActionsWithTarget(target objc.Object) {
	C.C_NSUndoManager_RemoveAllActionsWithTarget(n.Ptr(), objc.ExtractPtr(target))
}

func (n *NSUndoManager) SetActionName(actionName string) {
	C.C_NSUndoManager_SetActionName(n.Ptr(), NewString(actionName).Ptr())
}

func (n *NSUndoManager) UndoMenuTitleForUndoActionName(actionName string) string {
	result := C.C_NSUndoManager_UndoMenuTitleForUndoActionName(n.Ptr(), NewString(actionName).Ptr())
	return MakeString(result).String()
}

func (n *NSUndoManager) RedoMenuTitleForUndoActionName(actionName string) string {
	result := C.C_NSUndoManager_RedoMenuTitleForUndoActionName(n.Ptr(), NewString(actionName).Ptr())
	return MakeString(result).String()
}

func (n *NSUndoManager) SetActionIsDiscardable(discardable bool) {
	C.C_NSUndoManager_SetActionIsDiscardable(n.Ptr(), C.bool(discardable))
}

func (n *NSUndoManager) CanUndo() bool {
	result := C.C_NSUndoManager_CanUndo(n.Ptr())
	return bool(result)
}

func (n *NSUndoManager) CanRedo() bool {
	result := C.C_NSUndoManager_CanRedo(n.Ptr())
	return bool(result)
}

func (n *NSUndoManager) LevelsOfUndo() uint {
	result := C.C_NSUndoManager_LevelsOfUndo(n.Ptr())
	return uint(result)
}

func (n *NSUndoManager) SetLevelsOfUndo(value uint) {
	C.C_NSUndoManager_SetLevelsOfUndo(n.Ptr(), C.uint(value))
}

func (n *NSUndoManager) GroupsByEvent() bool {
	result := C.C_NSUndoManager_GroupsByEvent(n.Ptr())
	return bool(result)
}

func (n *NSUndoManager) SetGroupsByEvent(value bool) {
	C.C_NSUndoManager_SetGroupsByEvent(n.Ptr(), C.bool(value))
}

func (n *NSUndoManager) GroupingLevel() int {
	result := C.C_NSUndoManager_GroupingLevel(n.Ptr())
	return int(result)
}

func (n *NSUndoManager) IsUndoRegistrationEnabled() bool {
	result := C.C_NSUndoManager_IsUndoRegistrationEnabled(n.Ptr())
	return bool(result)
}

func (n *NSUndoManager) IsUndoing() bool {
	result := C.C_NSUndoManager_IsUndoing(n.Ptr())
	return bool(result)
}

func (n *NSUndoManager) IsRedoing() bool {
	result := C.C_NSUndoManager_IsRedoing(n.Ptr())
	return bool(result)
}

func (n *NSUndoManager) UndoActionName() string {
	result := C.C_NSUndoManager_UndoActionName(n.Ptr())
	return MakeString(result).String()
}

func (n *NSUndoManager) RedoActionName() string {
	result := C.C_NSUndoManager_RedoActionName(n.Ptr())
	return MakeString(result).String()
}

func (n *NSUndoManager) UndoMenuItemTitle() string {
	result := C.C_NSUndoManager_UndoMenuItemTitle(n.Ptr())
	return MakeString(result).String()
}

func (n *NSUndoManager) RedoMenuItemTitle() string {
	result := C.C_NSUndoManager_RedoMenuItemTitle(n.Ptr())
	return MakeString(result).String()
}

func (n *NSUndoManager) UndoActionIsDiscardable() bool {
	result := C.C_NSUndoManager_UndoActionIsDiscardable(n.Ptr())
	return bool(result)
}

func (n *NSUndoManager) RedoActionIsDiscardable() bool {
	result := C.C_NSUndoManager_RedoActionIsDiscardable(n.Ptr())
	return bool(result)
}
