#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_UndoManager_Alloc();

void* C_NSUndoManager_AllocUndoManager();
void* C_NSUndoManager_Init(void* ptr);
void* C_NSUndoManager_NewUndoManager();
void* C_NSUndoManager_Autorelease(void* ptr);
void* C_NSUndoManager_Retain(void* ptr);
void C_NSUndoManager_RegisterUndoWithTarget_Selector_Object(void* ptr, void* target, void* selector, void* anObject);
void* C_NSUndoManager_PrepareWithInvocationTarget(void* ptr, void* target);
void C_NSUndoManager_Undo(void* ptr);
void C_NSUndoManager_UndoNestedGroup(void* ptr);
void C_NSUndoManager_Redo(void* ptr);
void C_NSUndoManager_BeginUndoGrouping(void* ptr);
void C_NSUndoManager_EndUndoGrouping(void* ptr);
void C_NSUndoManager_DisableUndoRegistration(void* ptr);
void C_NSUndoManager_EnableUndoRegistration(void* ptr);
void C_NSUndoManager_RemoveAllActions(void* ptr);
void C_NSUndoManager_RemoveAllActionsWithTarget(void* ptr, void* target);
void C_NSUndoManager_SetActionName(void* ptr, void* actionName);
void* C_NSUndoManager_UndoMenuTitleForUndoActionName(void* ptr, void* actionName);
void* C_NSUndoManager_RedoMenuTitleForUndoActionName(void* ptr, void* actionName);
void C_NSUndoManager_SetActionIsDiscardable(void* ptr, bool discardable);
bool C_NSUndoManager_CanUndo(void* ptr);
bool C_NSUndoManager_CanRedo(void* ptr);
unsigned int C_NSUndoManager_LevelsOfUndo(void* ptr);
void C_NSUndoManager_SetLevelsOfUndo(void* ptr, unsigned int value);
bool C_NSUndoManager_GroupsByEvent(void* ptr);
void C_NSUndoManager_SetGroupsByEvent(void* ptr, bool value);
int C_NSUndoManager_GroupingLevel(void* ptr);
bool C_NSUndoManager_IsUndoRegistrationEnabled(void* ptr);
bool C_NSUndoManager_IsUndoing(void* ptr);
bool C_NSUndoManager_IsRedoing(void* ptr);
void* C_NSUndoManager_UndoActionName(void* ptr);
void* C_NSUndoManager_RedoActionName(void* ptr);
void* C_NSUndoManager_UndoMenuItemTitle(void* ptr);
void* C_NSUndoManager_RedoMenuItemTitle(void* ptr);
Array C_NSUndoManager_RunLoopModes(void* ptr);
void C_NSUndoManager_SetRunLoopModes(void* ptr, Array value);
bool C_NSUndoManager_UndoActionIsDiscardable(void* ptr);
bool C_NSUndoManager_RedoActionIsDiscardable(void* ptr);
