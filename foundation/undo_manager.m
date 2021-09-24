#import "undo_manager.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <Foundation/NSUndoManager.h>

void* C_UndoManager_Alloc() {
    return [NSUndoManager alloc];
}

void* C_NSUndoManager_AllocUndoManager() {
    NSUndoManager* result_ = [NSUndoManager alloc];
    return result_;
}

void* C_NSUndoManager_Init(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    NSUndoManager* result_ = [nSUndoManager init];
    return result_;
}

void* C_NSUndoManager_NewUndoManager() {
    NSUndoManager* result_ = [NSUndoManager new];
    return result_;
}

void* C_NSUndoManager_Autorelease(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    NSUndoManager* result_ = [nSUndoManager autorelease];
    return result_;
}

void* C_NSUndoManager_Retain(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    NSUndoManager* result_ = [nSUndoManager retain];
    return result_;
}

void C_NSUndoManager_RegisterUndoWithTarget_Selector_Object(void* ptr, void* target, void* selector, void* anObject) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    [nSUndoManager registerUndoWithTarget:(id)target selector:(SEL)selector object:(id)anObject];
}

void* C_NSUndoManager_PrepareWithInvocationTarget(void* ptr, void* target) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    id result_ = [nSUndoManager prepareWithInvocationTarget:(id)target];
    return result_;
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
    NSString* result_ = [nSUndoManager undoMenuTitleForUndoActionName:(NSString*)actionName];
    return result_;
}

void* C_NSUndoManager_RedoMenuTitleForUndoActionName(void* ptr, void* actionName) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    NSString* result_ = [nSUndoManager redoMenuTitleForUndoActionName:(NSString*)actionName];
    return result_;
}

void C_NSUndoManager_SetActionIsDiscardable(void* ptr, bool discardable) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    [nSUndoManager setActionIsDiscardable:discardable];
}

bool C_NSUndoManager_CanUndo(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    BOOL result_ = [nSUndoManager canUndo];
    return result_;
}

bool C_NSUndoManager_CanRedo(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    BOOL result_ = [nSUndoManager canRedo];
    return result_;
}

unsigned int C_NSUndoManager_LevelsOfUndo(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    NSUInteger result_ = [nSUndoManager levelsOfUndo];
    return result_;
}

void C_NSUndoManager_SetLevelsOfUndo(void* ptr, unsigned int value) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    [nSUndoManager setLevelsOfUndo:value];
}

bool C_NSUndoManager_GroupsByEvent(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    BOOL result_ = [nSUndoManager groupsByEvent];
    return result_;
}

void C_NSUndoManager_SetGroupsByEvent(void* ptr, bool value) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    [nSUndoManager setGroupsByEvent:value];
}

int C_NSUndoManager_GroupingLevel(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    NSInteger result_ = [nSUndoManager groupingLevel];
    return result_;
}

bool C_NSUndoManager_IsUndoRegistrationEnabled(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    BOOL result_ = [nSUndoManager isUndoRegistrationEnabled];
    return result_;
}

bool C_NSUndoManager_IsUndoing(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    BOOL result_ = [nSUndoManager isUndoing];
    return result_;
}

bool C_NSUndoManager_IsRedoing(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    BOOL result_ = [nSUndoManager isRedoing];
    return result_;
}

void* C_NSUndoManager_UndoActionName(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    NSString* result_ = [nSUndoManager undoActionName];
    return result_;
}

void* C_NSUndoManager_RedoActionName(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    NSString* result_ = [nSUndoManager redoActionName];
    return result_;
}

void* C_NSUndoManager_UndoMenuItemTitle(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    NSString* result_ = [nSUndoManager undoMenuItemTitle];
    return result_;
}

void* C_NSUndoManager_RedoMenuItemTitle(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    NSString* result_ = [nSUndoManager redoMenuItemTitle];
    return result_;
}

Array C_NSUndoManager_RunLoopModes(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    NSArray* result_ = [nSUndoManager runLoopModes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void C_NSUndoManager_SetRunLoopModes(void* ptr, Array value) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSString*)p];
    	}
    }
    [nSUndoManager setRunLoopModes:objcValue];
}

bool C_NSUndoManager_UndoActionIsDiscardable(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    BOOL result_ = [nSUndoManager undoActionIsDiscardable];
    return result_;
}

bool C_NSUndoManager_RedoActionIsDiscardable(void* ptr) {
    NSUndoManager* nSUndoManager = (NSUndoManager*)ptr;
    BOOL result_ = [nSUndoManager redoActionIsDiscardable];
    return result_;
}
