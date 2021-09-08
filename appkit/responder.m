#import <Appkit/Appkit.h>
#import "responder.h"

void* C_Responder_Alloc() {
    return [NSResponder alloc];
}

void* C_NSResponder_Init(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    NSResponder* result_ = [nSResponder init];
    return result_;
}

void* C_NSResponder_InitWithCoder(void* ptr, void* coder) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    NSResponder* result_ = [nSResponder initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSResponder_AllocResponder() {
    NSResponder* result_ = [NSResponder alloc];
    return result_;
}

void* C_NSResponder_NewResponder() {
    NSResponder* result_ = [NSResponder new];
    return result_;
}

void* C_NSResponder_Autorelease(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    NSResponder* result_ = [nSResponder autorelease];
    return result_;
}

void* C_NSResponder_Retain(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    NSResponder* result_ = [nSResponder retain];
    return result_;
}

bool C_NSResponder_BecomeFirstResponder(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    BOOL result_ = [nSResponder becomeFirstResponder];
    return result_;
}

bool C_NSResponder_ResignFirstResponder(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    BOOL result_ = [nSResponder resignFirstResponder];
    return result_;
}

bool C_NSResponder_ValidateProposedFirstResponder_ForEvent(void* ptr, void* responder, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    BOOL result_ = [nSResponder validateProposedFirstResponder:(NSResponder*)responder forEvent:(NSEvent*)event];
    return result_;
}

void C_NSResponder_MouseDown(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder mouseDown:(NSEvent*)event];
}

void C_NSResponder_MouseDragged(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder mouseDragged:(NSEvent*)event];
}

void C_NSResponder_MouseUp(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder mouseUp:(NSEvent*)event];
}

void C_NSResponder_MouseMoved(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder mouseMoved:(NSEvent*)event];
}

void C_NSResponder_MouseEntered(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder mouseEntered:(NSEvent*)event];
}

void C_NSResponder_MouseExited(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder mouseExited:(NSEvent*)event];
}

void C_NSResponder_RightMouseDown(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder rightMouseDown:(NSEvent*)event];
}

void C_NSResponder_RightMouseDragged(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder rightMouseDragged:(NSEvent*)event];
}

void C_NSResponder_RightMouseUp(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder rightMouseUp:(NSEvent*)event];
}

void C_NSResponder_OtherMouseDown(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder otherMouseDown:(NSEvent*)event];
}

void C_NSResponder_OtherMouseDragged(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder otherMouseDragged:(NSEvent*)event];
}

void C_NSResponder_OtherMouseUp(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder otherMouseUp:(NSEvent*)event];
}

void C_NSResponder_KeyDown(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder keyDown:(NSEvent*)event];
}

void C_NSResponder_KeyUp(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder keyUp:(NSEvent*)event];
}

void C_NSResponder_InterpretKeyEvents(void* ptr, Array eventArray) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    NSMutableArray* objcEventArray = [NSMutableArray arrayWithCapacity:eventArray.len];
    if (eventArray.len > 0) {
    	void** eventArrayData = (void**)eventArray.data;
    	for (int i = 0; i < eventArray.len; i++) {
    		void* p = eventArrayData[i];
    		[objcEventArray addObject:(NSEvent*)(NSEvent*)p];
    	}
    }
    [nSResponder interpretKeyEvents:objcEventArray];
}

bool C_NSResponder_PerformKeyEquivalent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    BOOL result_ = [nSResponder performKeyEquivalent:(NSEvent*)event];
    return result_;
}

void C_NSResponder_FlushBufferedKeyEvents(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder flushBufferedKeyEvents];
}

void C_NSResponder_PressureChangeWithEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder pressureChangeWithEvent:(NSEvent*)event];
}

void C_NSResponder_CursorUpdate(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder cursorUpdate:(NSEvent*)event];
}

void C_NSResponder_FlagsChanged(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder flagsChanged:(NSEvent*)event];
}

void C_NSResponder_TabletPoint(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder tabletPoint:(NSEvent*)event];
}

void C_NSResponder_TabletProximity(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder tabletProximity:(NSEvent*)event];
}

void C_NSResponder_HelpRequested(void* ptr, void* eventPtr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder helpRequested:(NSEvent*)eventPtr];
}

void C_NSResponder_ScrollWheel(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder scrollWheel:(NSEvent*)event];
}

void C_NSResponder_QuickLookWithEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder quickLookWithEvent:(NSEvent*)event];
}

void C_NSResponder_ChangeModeWithEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder changeModeWithEvent:(NSEvent*)event];
}

void* C_NSResponder_SupplementalTargetForAction_Sender(void* ptr, void* action, void* sender) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    id result_ = [nSResponder supplementalTargetForAction:(SEL)action sender:(id)sender];
    return result_;
}

void C_NSResponder_EncodeRestorableStateWithCoder(void* ptr, void* coder) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder encodeRestorableStateWithCoder:(NSCoder*)coder];
}

void C_NSResponder_RestoreStateWithCoder(void* ptr, void* coder) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder restoreStateWithCoder:(NSCoder*)coder];
}

void C_NSResponder_InvalidateRestorableState(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder invalidateRestorableState];
}

void C_NSResponder_UpdateUserActivityState(void* ptr, void* userActivity) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder updateUserActivityState:(NSUserActivity*)userActivity];
}

bool C_NSResponder_PresentError(void* ptr, void* error) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    BOOL result_ = [nSResponder presentError:(NSError*)error];
    return result_;
}

void* C_NSResponder_WillPresentError(void* ptr, void* error) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    NSError* result_ = [nSResponder willPresentError:(NSError*)error];
    return result_;
}

bool C_NSResponder_TryToPerform_With(void* ptr, void* action, void* object) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    BOOL result_ = [nSResponder tryToPerform:(SEL)action with:(id)object];
    return result_;
}

void* C_NSResponder_ValidRequestorForSendType_ReturnType(void* ptr, void* sendType, void* returnType) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    id result_ = [nSResponder validRequestorForSendType:(NSString*)sendType returnType:(NSString*)returnType];
    return result_;
}

bool C_NSResponder_ShouldBeTreatedAsInkEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    BOOL result_ = [nSResponder shouldBeTreatedAsInkEvent:(NSEvent*)event];
    return result_;
}

void C_NSResponder_NoResponderFor(void* ptr, void* eventSelector) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder noResponderFor:(SEL)eventSelector];
}

void C_NSResponder_BeginGestureWithEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder beginGestureWithEvent:(NSEvent*)event];
}

void C_NSResponder_EndGestureWithEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder endGestureWithEvent:(NSEvent*)event];
}

void C_NSResponder_MagnifyWithEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder magnifyWithEvent:(NSEvent*)event];
}

void C_NSResponder_RotateWithEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder rotateWithEvent:(NSEvent*)event];
}

void C_NSResponder_SwipeWithEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder swipeWithEvent:(NSEvent*)event];
}

void C_NSResponder_TouchesBeganWithEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder touchesBeganWithEvent:(NSEvent*)event];
}

void C_NSResponder_TouchesMovedWithEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder touchesMovedWithEvent:(NSEvent*)event];
}

void C_NSResponder_TouchesCancelledWithEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder touchesCancelledWithEvent:(NSEvent*)event];
}

void C_NSResponder_TouchesEndedWithEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder touchesEndedWithEvent:(NSEvent*)event];
}

bool C_NSResponder_WantsForwardedScrollEventsForAxis(void* ptr, int axis) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    BOOL result_ = [nSResponder wantsForwardedScrollEventsForAxis:axis];
    return result_;
}

void C_NSResponder_SmartMagnifyWithEvent(void* ptr, void* event) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder smartMagnifyWithEvent:(NSEvent*)event];
}

bool C_NSResponder_WantsScrollEventsForSwipeTrackingOnAxis(void* ptr, int axis) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    BOOL result_ = [nSResponder wantsScrollEventsForSwipeTrackingOnAxis:axis];
    return result_;
}

void C_NSResponder_PerformTextFinderAction(void* ptr, void* sender) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder performTextFinderAction:(id)sender];
}

void C_NSResponder_EncodeRestorableStateWithCoder_BackgroundQueue(void* ptr, void* coder, void* queue) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder encodeRestorableStateWithCoder:(NSCoder*)coder backgroundQueue:(NSOperationQueue*)queue];
}

void* C_NSResponder_MakeTouchBar(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    NSTouchBar* result_ = [nSResponder makeTouchBar];
    return result_;
}

void C_NSResponder_NewWindowForTab(void* ptr, void* sender) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder newWindowForTab:(id)sender];
}

void C_NSResponder_ShowContextHelp(void* ptr, void* sender) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder showContextHelp:(id)sender];
}

bool C_NSResponder_AcceptsFirstResponder(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    BOOL result_ = [nSResponder acceptsFirstResponder];
    return result_;
}

void* C_NSResponder_NextResponder(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    NSResponder* result_ = [nSResponder nextResponder];
    return result_;
}

void C_NSResponder_SetNextResponder(void* ptr, void* value) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder setNextResponder:(NSResponder*)value];
}

Array C_NSResponder_Responder_RestorableStateKeyPaths() {
    NSArray* result_ = [NSResponder restorableStateKeyPaths];
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

void* C_NSResponder_UserActivity(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    NSUserActivity* result_ = [nSResponder userActivity];
    return result_;
}

void C_NSResponder_SetUserActivity(void* ptr, void* value) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder setUserActivity:(NSUserActivity*)value];
}

void* C_NSResponder_Menu(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    NSMenu* result_ = [nSResponder menu];
    return result_;
}

void C_NSResponder_SetMenu(void* ptr, void* value) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder setMenu:(NSMenu*)value];
}

void* C_NSResponder_UndoManager(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    NSUndoManager* result_ = [nSResponder undoManager];
    return result_;
}

void* C_NSResponder_TouchBar(void* ptr) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    NSTouchBar* result_ = [nSResponder touchBar];
    return result_;
}

void C_NSResponder_SetTouchBar(void* ptr, void* value) {
    NSResponder* nSResponder = (NSResponder*)ptr;
    [nSResponder setTouchBar:(NSTouchBar*)value];
}
