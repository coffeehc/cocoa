#import <Foundation/Foundation.h>
#import "undo_manager.h"

void* C_UndoManager_Alloc() {
	return [NSUndoManager alloc];
}

void* C_NSUndoManager_Init(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	NSUndoManager* result = [nSUndoManager init];
	return result;
}

void C_NSUndoManager_RegisterUndoWithTarget_Selector_Object(void* ptr, void* target, void* selector, void* anObject) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager registerUndoWithTarget:(id)target selector:(SEL)selector object:(id)anObject];
}

void* C_NSUndoManager_PrepareWithInvocationTarget(void* ptr, void* target) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	id result = [nSUndoManager prepareWithInvocationTarget:(id)target];
	return result;
}

void C_NSUndoManager_Undo(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager undo];
}

void C_NSUndoManager_UndoNestedGroup(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager undoNestedGroup];
}

void C_NSUndoManager_Redo(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager redo];
}

void C_NSUndoManager_BeginUndoGrouping(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager beginUndoGrouping];
}

void C_NSUndoManager_EndUndoGrouping(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager endUndoGrouping];
}

void C_NSUndoManager_DisableUndoRegistration(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager disableUndoRegistration];
}

void C_NSUndoManager_EnableUndoRegistration(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager enableUndoRegistration];
}

void C_NSUndoManager_RemoveAllActions(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager removeAllActions];
}

void C_NSUndoManager_RemoveAllActionsWithTarget(void* ptr, void* target) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager removeAllActionsWithTarget:(id)target];
}

void C_NSUndoManager_SetActionName(void* ptr, void* actionName) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager setActionName:(NSString*)actionName];
}

void* C_NSUndoManager_UndoMenuTitleForUndoActionName(void* ptr, void* actionName) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	NSString* result = [nSUndoManager undoMenuTitleForUndoActionName:(NSString*)actionName];
	return result;
}

void* C_NSUndoManager_RedoMenuTitleForUndoActionName(void* ptr, void* actionName) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	NSString* result = [nSUndoManager redoMenuTitleForUndoActionName:(NSString*)actionName];
	return result;
}

void C_NSUndoManager_SetActionIsDiscardable(void* ptr, bool discardable) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager setActionIsDiscardable:discardable];
}

bool C_NSUndoManager_CanUndo(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	bool result = [nSUndoManager canUndo];
	return result;
}

bool C_NSUndoManager_CanRedo(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	bool result = [nSUndoManager canRedo];
	return result;
}

unsigned int C_NSUndoManager_LevelsOfUndo(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	unsigned int result = [nSUndoManager levelsOfUndo];
	return result;
}

void C_NSUndoManager_SetLevelsOfUndo(void* ptr, unsigned int value) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager setLevelsOfUndo:value];
}

bool C_NSUndoManager_GroupsByEvent(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	bool result = [nSUndoManager groupsByEvent];
	return result;
}

void C_NSUndoManager_SetGroupsByEvent(void* ptr, bool value) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	[nSUndoManager setGroupsByEvent:value];
}

int C_NSUndoManager_GroupingLevel(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	int result = [nSUndoManager groupingLevel];
	return result;
}

bool C_NSUndoManager_IsUndoRegistrationEnabled(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	bool result = [nSUndoManager isUndoRegistrationEnabled];
	return result;
}

bool C_NSUndoManager_IsUndoing(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	bool result = [nSUndoManager isUndoing];
	return result;
}

bool C_NSUndoManager_IsRedoing(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	bool result = [nSUndoManager isRedoing];
	return result;
}

void* C_NSUndoManager_UndoActionName(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	NSString* result = [nSUndoManager undoActionName];
	return result;
}

void* C_NSUndoManager_RedoActionName(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	NSString* result = [nSUndoManager redoActionName];
	return result;
}

void* C_NSUndoManager_UndoMenuItemTitle(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	NSString* result = [nSUndoManager undoMenuItemTitle];
	return result;
}

void* C_NSUndoManager_RedoMenuItemTitle(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	NSString* result = [nSUndoManager redoMenuItemTitle];
	return result;
}

bool C_NSUndoManager_UndoActionIsDiscardable(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	bool result = [nSUndoManager undoActionIsDiscardable];
	return result;
}

bool C_NSUndoManager_RedoActionIsDiscardable(void* ptr) {
	NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
	bool result = [nSUndoManager redoActionIsDiscardable];
	return result;
}
