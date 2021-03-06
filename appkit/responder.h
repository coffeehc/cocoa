#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Responder_Alloc();

void* C_NSResponder_Init(void* ptr);
void* C_NSResponder_InitWithCoder(void* ptr, void* coder);
void* C_NSResponder_AllocResponder();
void* C_NSResponder_NewResponder();
void* C_NSResponder_Autorelease(void* ptr);
void* C_NSResponder_Retain(void* ptr);
bool C_NSResponder_BecomeFirstResponder(void* ptr);
bool C_NSResponder_ResignFirstResponder(void* ptr);
bool C_NSResponder_ValidateProposedFirstResponder_ForEvent(void* ptr, void* responder, void* event);
void C_NSResponder_MouseDown(void* ptr, void* event);
void C_NSResponder_MouseDragged(void* ptr, void* event);
void C_NSResponder_MouseUp(void* ptr, void* event);
void C_NSResponder_MouseMoved(void* ptr, void* event);
void C_NSResponder_MouseEntered(void* ptr, void* event);
void C_NSResponder_MouseExited(void* ptr, void* event);
void C_NSResponder_RightMouseDown(void* ptr, void* event);
void C_NSResponder_RightMouseDragged(void* ptr, void* event);
void C_NSResponder_RightMouseUp(void* ptr, void* event);
void C_NSResponder_OtherMouseDown(void* ptr, void* event);
void C_NSResponder_OtherMouseDragged(void* ptr, void* event);
void C_NSResponder_OtherMouseUp(void* ptr, void* event);
void C_NSResponder_KeyDown(void* ptr, void* event);
void C_NSResponder_KeyUp(void* ptr, void* event);
void C_NSResponder_InterpretKeyEvents(void* ptr, Array eventArray);
bool C_NSResponder_PerformKeyEquivalent(void* ptr, void* event);
void C_NSResponder_FlushBufferedKeyEvents(void* ptr);
void C_NSResponder_PressureChangeWithEvent(void* ptr, void* event);
void C_NSResponder_CursorUpdate(void* ptr, void* event);
void C_NSResponder_FlagsChanged(void* ptr, void* event);
void C_NSResponder_TabletPoint(void* ptr, void* event);
void C_NSResponder_TabletProximity(void* ptr, void* event);
void C_NSResponder_HelpRequested(void* ptr, void* eventPtr);
void C_NSResponder_ScrollWheel(void* ptr, void* event);
void C_NSResponder_QuickLookWithEvent(void* ptr, void* event);
void C_NSResponder_ChangeModeWithEvent(void* ptr, void* event);
void* C_NSResponder_SupplementalTargetForAction_Sender(void* ptr, void* action, void* sender);
void C_NSResponder_EncodeRestorableStateWithCoder(void* ptr, void* coder);
void C_NSResponder_RestoreStateWithCoder(void* ptr, void* coder);
void C_NSResponder_InvalidateRestorableState(void* ptr);
void C_NSResponder_UpdateUserActivityState(void* ptr, void* userActivity);
bool C_NSResponder_PresentError(void* ptr, void* error);
void* C_NSResponder_WillPresentError(void* ptr, void* error);
bool C_NSResponder_TryToPerform_With(void* ptr, void* action, void* object);
void* C_NSResponder_ValidRequestorForSendType_ReturnType(void* ptr, void* sendType, void* returnType);
bool C_NSResponder_ShouldBeTreatedAsInkEvent(void* ptr, void* event);
void C_NSResponder_NoResponderFor(void* ptr, void* eventSelector);
void C_NSResponder_BeginGestureWithEvent(void* ptr, void* event);
void C_NSResponder_EndGestureWithEvent(void* ptr, void* event);
void C_NSResponder_MagnifyWithEvent(void* ptr, void* event);
void C_NSResponder_RotateWithEvent(void* ptr, void* event);
void C_NSResponder_SwipeWithEvent(void* ptr, void* event);
void C_NSResponder_TouchesBeganWithEvent(void* ptr, void* event);
void C_NSResponder_TouchesMovedWithEvent(void* ptr, void* event);
void C_NSResponder_TouchesCancelledWithEvent(void* ptr, void* event);
void C_NSResponder_TouchesEndedWithEvent(void* ptr, void* event);
bool C_NSResponder_WantsForwardedScrollEventsForAxis(void* ptr, int axis);
void C_NSResponder_SmartMagnifyWithEvent(void* ptr, void* event);
bool C_NSResponder_WantsScrollEventsForSwipeTrackingOnAxis(void* ptr, int axis);
void C_NSResponder_PerformTextFinderAction(void* ptr, void* sender);
void C_NSResponder_EncodeRestorableStateWithCoder_BackgroundQueue(void* ptr, void* coder, void* queue);
void* C_NSResponder_MakeTouchBar(void* ptr);
void C_NSResponder_NewWindowForTab(void* ptr, void* sender);
void C_NSResponder_ShowContextHelp(void* ptr, void* sender);
bool C_NSResponder_AcceptsFirstResponder(void* ptr);
void* C_NSResponder_NextResponder(void* ptr);
void C_NSResponder_SetNextResponder(void* ptr, void* value);
Array C_NSResponder_Responder_RestorableStateKeyPaths();
void* C_NSResponder_UserActivity(void* ptr);
void C_NSResponder_SetUserActivity(void* ptr, void* value);
void* C_NSResponder_Menu(void* ptr);
void C_NSResponder_SetMenu(void* ptr, void* value);
void* C_NSResponder_UndoManager(void* ptr);
void* C_NSResponder_TouchBar(void* ptr);
void C_NSResponder_SetTouchBar(void* ptr, void* value);
