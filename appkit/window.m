#import <Appkit/Appkit.h>
#import "window.h"

void* C_Window_Alloc() {
    return [NSWindow alloc];
}

void* C_NSWindow_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindow* result_ = [nSWindow initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag];
    return result_;
}

void* C_NSWindow_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindow* result_ = [nSWindow initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag screen:(NSScreen*)screen];
    return result_;
}

void* C_NSWindow_Init(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindow* result_ = [nSWindow init];
    return result_;
}

void* C_NSWindow_WindowWithContentViewController(void* contentViewController) {
    NSWindow* result_ = [NSWindow windowWithContentViewController:(NSViewController*)contentViewController];
    return result_;
}

void C_NSWindow_ToggleFullScreen(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow toggleFullScreen:(id)sender];
}

void C_NSWindow_InvalidateShadow(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow invalidateShadow];
}

bool C_NSWindow_AutorecalculatesContentBorderThicknessForEdge(void* ptr, unsigned int edge) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow autorecalculatesContentBorderThicknessForEdge:edge];
    return result_;
}

void C_NSWindow_SetAutorecalculatesContentBorderThickness_ForEdge(void* ptr, bool flag, unsigned int edge) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setAutorecalculatesContentBorderThickness:flag forEdge:edge];
}

double C_NSWindow_ContentBorderThicknessForEdge(void* ptr, unsigned int edge) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    CGFloat result_ = [nSWindow contentBorderThicknessForEdge:edge];
    return result_;
}

void C_NSWindow_SetContentBorderThickness_ForEdge(void* ptr, double thickness, unsigned int edge) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setContentBorderThickness:thickness forEdge:edge];
}

Array C_NSWindow_WindowNumbersWithOptions(unsigned int options) {
    NSArray* result_ = [NSWindow windowNumbersWithOptions:options];
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

CGRect C_NSWindow_Window_ContentRectForFrameRect_StyleMask(CGRect fRect, unsigned int style) {
    NSRect result_ = [NSWindow contentRectForFrameRect:fRect styleMask:style];
    return result_;
}

CGRect C_NSWindow_Window_FrameRectForContentRect_StyleMask(CGRect cRect, unsigned int style) {
    NSRect result_ = [NSWindow frameRectForContentRect:cRect styleMask:style];
    return result_;
}

double C_NSWindow_Window_MinFrameWidthWithTitle_StyleMask(void* title, unsigned int style) {
    CGFloat result_ = [NSWindow minFrameWidthWithTitle:(NSString*)title styleMask:style];
    return result_;
}

CGRect C_NSWindow_ContentRectForFrameRect(void* ptr, CGRect frameRect) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSRect result_ = [nSWindow contentRectForFrameRect:frameRect];
    return result_;
}

CGRect C_NSWindow_FrameRectForContentRect(void* ptr, CGRect contentRect) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSRect result_ = [nSWindow frameRectForContentRect:contentRect];
    return result_;
}

void C_NSWindow_EndSheet(void* ptr, void* sheetWindow) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow endSheet:(NSWindow*)sheetWindow];
}

void C_NSWindow_EndSheet_ReturnCode(void* ptr, void* sheetWindow, int returnCode) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow endSheet:(NSWindow*)sheetWindow returnCode:returnCode];
}

void C_NSWindow_SetFrameOrigin(void* ptr, CGPoint point) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setFrameOrigin:point];
}

void C_NSWindow_SetFrameTopLeftPoint(void* ptr, CGPoint point) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setFrameTopLeftPoint:point];
}

CGRect C_NSWindow_ConstrainFrameRect_ToScreen(void* ptr, CGRect frameRect, void* screen) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSRect result_ = [nSWindow constrainFrameRect:frameRect toScreen:(NSScreen*)screen];
    return result_;
}

CGPoint C_NSWindow_CascadeTopLeftFromPoint(void* ptr, CGPoint topLeftPoint) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSPoint result_ = [nSWindow cascadeTopLeftFromPoint:topLeftPoint];
    return result_;
}

void C_NSWindow_SetFrame_Display(void* ptr, CGRect frameRect, bool flag) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setFrame:frameRect display:flag];
}

void C_NSWindow_SetFrame_Display_Animate(void* ptr, CGRect frameRect, bool displayFlag, bool animateFlag) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setFrame:frameRect display:displayFlag animate:animateFlag];
}

double C_NSWindow_AnimationResizeTime(void* ptr, CGRect newFrame) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSTimeInterval result_ = [nSWindow animationResizeTime:newFrame];
    return result_;
}

void C_NSWindow_PerformZoom(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow performZoom:(id)sender];
}

void C_NSWindow_Zoom(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow zoom:(id)sender];
}

void C_NSWindow_SetContentSize(void* ptr, CGSize size) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setContentSize:size];
}

void C_NSWindow_OrderOut(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow orderOut:(id)sender];
}

void C_NSWindow_OrderBack(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow orderBack:(id)sender];
}

void C_NSWindow_OrderFront(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow orderFront:(id)sender];
}

void C_NSWindow_OrderFrontRegardless(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow orderFrontRegardless];
}

void C_NSWindow_OrderWindow_RelativeTo(void* ptr, int place, int otherWin) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow orderWindow:place relativeTo:otherWin];
}

void C_NSWindow_Window_RemoveFrameUsingName(void* name) {
    [NSWindow removeFrameUsingName:(NSString*)name];
}

bool C_NSWindow_SetFrameUsingName(void* ptr, void* name) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow setFrameUsingName:(NSString*)name];
    return result_;
}

bool C_NSWindow_SetFrameUsingName_Force(void* ptr, void* name, bool force) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow setFrameUsingName:(NSString*)name force:force];
    return result_;
}

void C_NSWindow_SaveFrameUsingName(void* ptr, void* name) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow saveFrameUsingName:(NSString*)name];
}

void C_NSWindow_SetFrameFromString(void* ptr, void* _string) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setFrameFromString:(NSString*)_string];
}

void C_NSWindow_MakeKeyWindow(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow makeKeyWindow];
}

void C_NSWindow_MakeKeyAndOrderFront(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow makeKeyAndOrderFront:(id)sender];
}

void C_NSWindow_BecomeKeyWindow(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow becomeKeyWindow];
}

void C_NSWindow_ResignKeyWindow(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow resignKeyWindow];
}

void C_NSWindow_MakeMainWindow(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow makeMainWindow];
}

void C_NSWindow_BecomeMainWindow(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow becomeMainWindow];
}

void C_NSWindow_ResignMainWindow(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow resignMainWindow];
}

void C_NSWindow_ToggleToolbarShown(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow toggleToolbarShown:(id)sender];
}

void C_NSWindow_RunToolbarCustomizationPalette(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow runToolbarCustomizationPalette:(id)sender];
}

void C_NSWindow_AddChildWindow_Ordered(void* ptr, void* childWin, int place) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow addChildWindow:(NSWindow*)childWin ordered:place];
}

void C_NSWindow_RemoveChildWindow(void* ptr, void* childWin) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow removeChildWindow:(NSWindow*)childWin];
}

void C_NSWindow_EnableKeyEquivalentForDefaultButtonCell(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow enableKeyEquivalentForDefaultButtonCell];
}

void C_NSWindow_DisableKeyEquivalentForDefaultButtonCell(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow disableKeyEquivalentForDefaultButtonCell];
}

void* C_NSWindow_FieldEditor_ForObject(void* ptr, bool createFlag, void* object) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSText* result_ = [nSWindow fieldEditor:createFlag forObject:(id)object];
    return result_;
}

void C_NSWindow_EndEditingFor(void* ptr, void* object) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow endEditingFor:(id)object];
}

void C_NSWindow_EnableCursorRects(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow enableCursorRects];
}

void C_NSWindow_DisableCursorRects(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow disableCursorRects];
}

void C_NSWindow_DiscardCursorRects(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow discardCursorRects];
}

void C_NSWindow_InvalidateCursorRectsForView(void* ptr, void* view) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow invalidateCursorRectsForView:(NSView*)view];
}

void C_NSWindow_ResetCursorRects(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow resetCursorRects];
}

void* C_NSWindow_Window_StandardWindowButton_ForStyleMask(unsigned int b, unsigned int styleMask) {
    NSButton* result_ = [NSWindow standardWindowButton:b forStyleMask:styleMask];
    return result_;
}

void* C_NSWindow_StandardWindowButton(void* ptr, unsigned int b) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSButton* result_ = [nSWindow standardWindowButton:b];
    return result_;
}

void C_NSWindow_AddTitlebarAccessoryViewController(void* ptr, void* childViewController) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow addTitlebarAccessoryViewController:(NSTitlebarAccessoryViewController*)childViewController];
}

void C_NSWindow_InsertTitlebarAccessoryViewController_AtIndex(void* ptr, void* childViewController, int index) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow insertTitlebarAccessoryViewController:(NSTitlebarAccessoryViewController*)childViewController atIndex:index];
}

void C_NSWindow_RemoveTitlebarAccessoryViewControllerAtIndex(void* ptr, int index) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow removeTitlebarAccessoryViewControllerAtIndex:index];
}

void C_NSWindow_AddTabbedWindow_Ordered(void* ptr, void* window, int ordered) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow addTabbedWindow:(NSWindow*)window ordered:ordered];
}

void C_NSWindow_SelectNextTab(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow selectNextTab:(id)sender];
}

void C_NSWindow_SelectPreviousTab(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow selectPreviousTab:(id)sender];
}

void C_NSWindow_MoveTabToNewWindow(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow moveTabToNewWindow:(id)sender];
}

void C_NSWindow_ToggleTabBar(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow toggleTabBar:(id)sender];
}

void C_NSWindow_ToggleTabOverview(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow toggleTabOverview:(id)sender];
}

void C_NSWindow_PostEvent_AtStart(void* ptr, void* event, bool flag) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow postEvent:(NSEvent*)event atStart:flag];
}

void C_NSWindow_SendEvent(void* ptr, void* event) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow sendEvent:(NSEvent*)event];
}

bool C_NSWindow_MakeFirstResponder(void* ptr, void* responder) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow makeFirstResponder:(NSResponder*)responder];
    return result_;
}

void C_NSWindow_SelectKeyViewPrecedingView(void* ptr, void* view) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow selectKeyViewPrecedingView:(NSView*)view];
}

void C_NSWindow_SelectKeyViewFollowingView(void* ptr, void* view) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow selectKeyViewFollowingView:(NSView*)view];
}

void C_NSWindow_SelectPreviousKeyView(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow selectPreviousKeyView:(id)sender];
}

void C_NSWindow_SelectNextKeyView(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow selectNextKeyView:(id)sender];
}

void C_NSWindow_RecalculateKeyViewLoop(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow recalculateKeyViewLoop];
}

int C_NSWindow_WindowNumberAtPoint_BelowWindowWithWindowNumber(CGPoint point, int windowNumber) {
    NSInteger result_ = [NSWindow windowNumberAtPoint:point belowWindowWithWindowNumber:windowNumber];
    return result_;
}

void C_NSWindow_PerformWindowDragWithEvent(void* ptr, void* event) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow performWindowDragWithEvent:(NSEvent*)event];
}

void C_NSWindow_DisableSnapshotRestoration(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow disableSnapshotRestoration];
}

void C_NSWindow_EnableSnapshotRestoration(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow enableSnapshotRestoration];
}

void C_NSWindow_Display(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow display];
}

void C_NSWindow_DisplayIfNeeded(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow displayIfNeeded];
}

void C_NSWindow_DisableScreenUpdatesUntilFlush(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow disableScreenUpdatesUntilFlush];
}

void C_NSWindow_Update(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow update];
}

void C_NSWindow_DragImage_At_Offset_Event_Pasteboard_Source_SlideBack(void* ptr, void* image, CGPoint baseLocation, CGSize initialOffset, void* event, void* pboard, void* sourceObj, bool slideFlag) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow dragImage:(NSImage*)image at:baseLocation offset:initialOffset event:(NSEvent*)event pasteboard:(NSPasteboard*)pboard source:(id)sourceObj slideBack:slideFlag];
}

void C_NSWindow_RegisterForDraggedTypes(void* ptr, Array newTypes) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSMutableArray* objcNewTypes = [[NSMutableArray alloc] init];
    if (newTypes.len > 0) {
    	void** newTypesData = (void**)newTypes.data;
    	for (int i = 0; i < newTypes.len; i++) {
    		void* p = newTypesData[i];
    		[objcNewTypes addObject:(NSPasteboardType)(NSString*)p];
    	}
    }
    [nSWindow registerForDraggedTypes:objcNewTypes];
}

void C_NSWindow_UnregisterDraggedTypes(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow unregisterDraggedTypes];
}

CGRect C_NSWindow_ConvertRectFromBacking(void* ptr, CGRect rect) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSRect result_ = [nSWindow convertRectFromBacking:rect];
    return result_;
}

CGRect C_NSWindow_ConvertRectToBacking(void* ptr, CGRect rect) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSRect result_ = [nSWindow convertRectToBacking:rect];
    return result_;
}

CGRect C_NSWindow_ConvertRectToScreen(void* ptr, CGRect rect) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSRect result_ = [nSWindow convertRectToScreen:rect];
    return result_;
}

CGRect C_NSWindow_ConvertRectFromScreen(void* ptr, CGRect rect) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSRect result_ = [nSWindow convertRectFromScreen:rect];
    return result_;
}

void C_NSWindow_SetTitleWithRepresentedFilename(void* ptr, void* filename) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setTitleWithRepresentedFilename:(NSString*)filename];
}

void C_NSWindow_Center(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow center];
}

void C_NSWindow_PerformClose(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow performClose:(id)sender];
}

void C_NSWindow_Close(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow close];
}

void C_NSWindow_PerformMiniaturize(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow performMiniaturize:(id)sender];
}

void C_NSWindow_Miniaturize(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow miniaturize:(id)sender];
}

void C_NSWindow_Deminiaturize(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow deminiaturize:(id)sender];
}

void C_NSWindow_Print(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow print:(id)sender];
}

Array C_NSWindow_DataWithEPSInsideRect(void* ptr, CGRect rect) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSData* result_ = [nSWindow dataWithEPSInsideRect:rect];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

Array C_NSWindow_DataWithPDFInsideRect(void* ptr, CGRect rect) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSData* result_ = [nSWindow dataWithPDFInsideRect:rect];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

void C_NSWindow_UpdateConstraintsIfNeeded(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow updateConstraintsIfNeeded];
}

void C_NSWindow_LayoutIfNeeded(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow layoutIfNeeded];
}

void C_NSWindow_VisualizeConstraints(void* ptr, Array constraints) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSMutableArray* objcConstraints = [[NSMutableArray alloc] init];
    if (constraints.len > 0) {
    	void** constraintsData = (void**)constraints.data;
    	for (int i = 0; i < constraints.len; i++) {
    		void* p = constraintsData[i];
    		[objcConstraints addObject:(NSLayoutConstraint*)(NSLayoutConstraint*)p];
    	}
    }
    [nSWindow visualizeConstraints:objcConstraints];
}

int C_NSWindow_AnchorAttributeForOrientation(void* ptr, int orientation) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSLayoutAttribute result_ = [nSWindow anchorAttributeForOrientation:orientation];
    return result_;
}

void C_NSWindow_SetAnchorAttribute_ForOrientation(void* ptr, int attr, int orientation) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setAnchorAttribute:attr forOrientation:orientation];
}

void C_NSWindow_SetIsMiniaturized(void* ptr, bool flag) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setIsMiniaturized:flag];
}

void C_NSWindow_SetIsVisible(void* ptr, bool flag) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setIsVisible:flag];
}

void C_NSWindow_SetIsZoomed(void* ptr, bool flag) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setIsZoomed:flag];
}

void* C_NSWindow_HandleCloseScriptCommand(void* ptr, void* command) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    id result_ = [nSWindow handleCloseScriptCommand:(NSCloseCommand*)command];
    return result_;
}

void* C_NSWindow_HandlePrintScriptCommand(void* ptr, void* command) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    id result_ = [nSWindow handlePrintScriptCommand:(NSScriptCommand*)command];
    return result_;
}

void* C_NSWindow_HandleSaveScriptCommand(void* ptr, void* command) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    id result_ = [nSWindow handleSaveScriptCommand:(NSScriptCommand*)command];
    return result_;
}

bool C_NSWindow_CanRepresentDisplayGamut(void* ptr, int displayGamut) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow canRepresentDisplayGamut:displayGamut];
    return result_;
}

CGPoint C_NSWindow_ConvertPointFromScreen(void* ptr, CGPoint point) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSPoint result_ = [nSWindow convertPointFromScreen:point];
    return result_;
}

CGPoint C_NSWindow_ConvertPointToScreen(void* ptr, CGPoint point) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSPoint result_ = [nSWindow convertPointToScreen:point];
    return result_;
}

CGPoint C_NSWindow_ConvertPointFromBacking(void* ptr, CGPoint point) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSPoint result_ = [nSWindow convertPointFromBacking:point];
    return result_;
}

CGPoint C_NSWindow_ConvertPointToBacking(void* ptr, CGPoint point) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSPoint result_ = [nSWindow convertPointToBacking:point];
    return result_;
}

void C_NSWindow_MergeAllWindows(void* ptr, void* sender) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow mergeAllWindows:(id)sender];
}

void C_NSWindow_SetDynamicDepthLimit(void* ptr, bool flag) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setDynamicDepthLimit:flag];
}

bool C_NSWindow_SetFrameAutosaveName(void* ptr, void* name) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow setFrameAutosaveName:(NSString*)name];
    return result_;
}

void* C_NSWindow_Delegate(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    id result_ = [nSWindow delegate];
    return result_;
}

void C_NSWindow_SetDelegate(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setDelegate:(id)value];
}

void* C_NSWindow_ContentViewController(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSViewController* result_ = [nSWindow contentViewController];
    return result_;
}

void C_NSWindow_SetContentViewController(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setContentViewController:(NSViewController*)value];
}

void* C_NSWindow_ContentView(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSView* result_ = [nSWindow contentView];
    return result_;
}

void C_NSWindow_SetContentView(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setContentView:(NSView*)value];
}

bool C_NSWindow_WorksWhenModal(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow worksWhenModal];
    return result_;
}

double C_NSWindow_AlphaValue(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    CGFloat result_ = [nSWindow alphaValue];
    return result_;
}

void C_NSWindow_SetAlphaValue(void* ptr, double value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setAlphaValue:value];
}

void* C_NSWindow_BackgroundColor(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSColor* result_ = [nSWindow backgroundColor];
    return result_;
}

void C_NSWindow_SetBackgroundColor(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setBackgroundColor:(NSColor*)value];
}

void* C_NSWindow_ColorSpace(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSColorSpace* result_ = [nSWindow colorSpace];
    return result_;
}

void C_NSWindow_SetColorSpace(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setColorSpace:(NSColorSpace*)value];
}

bool C_NSWindow_CanHide(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow canHide];
    return result_;
}

void C_NSWindow_SetCanHide(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setCanHide:value];
}

bool C_NSWindow_IsOnActiveSpace(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isOnActiveSpace];
    return result_;
}

bool C_NSWindow_HidesOnDeactivate(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow hidesOnDeactivate];
    return result_;
}

void C_NSWindow_SetHidesOnDeactivate(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setHidesOnDeactivate:value];
}

unsigned int C_NSWindow_CollectionBehavior(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowCollectionBehavior result_ = [nSWindow collectionBehavior];
    return result_;
}

void C_NSWindow_SetCollectionBehavior(void* ptr, unsigned int value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setCollectionBehavior:value];
}

bool C_NSWindow_IsOpaque(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isOpaque];
    return result_;
}

void C_NSWindow_SetOpaque(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setOpaque:value];
}

bool C_NSWindow_HasShadow(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow hasShadow];
    return result_;
}

void C_NSWindow_SetHasShadow(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setHasShadow:value];
}

bool C_NSWindow_PreventsApplicationTerminationWhenModal(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow preventsApplicationTerminationWhenModal];
    return result_;
}

void C_NSWindow_SetPreventsApplicationTerminationWhenModal(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setPreventsApplicationTerminationWhenModal:value];
}

int32_t C_NSWindow_Window_DefaultDepthLimit() {
    NSWindowDepth result_ = [NSWindow defaultDepthLimit];
    return result_;
}

int C_NSWindow_WindowNumber(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSInteger result_ = [nSWindow windowNumber];
    return result_;
}

Dictionary C_NSWindow_DeviceDescription(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSDictionary* result_ = [nSWindow deviceDescription];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSDeviceDescriptionKey kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

bool C_NSWindow_CanBecomeVisibleWithoutLogin(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow canBecomeVisibleWithoutLogin];
    return result_;
}

void C_NSWindow_SetCanBecomeVisibleWithoutLogin(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setCanBecomeVisibleWithoutLogin:value];
}

unsigned int C_NSWindow_SharingType(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowSharingType result_ = [nSWindow sharingType];
    return result_;
}

void C_NSWindow_SetSharingType(void* ptr, unsigned int value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setSharingType:value];
}

unsigned int C_NSWindow_BackingType(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSBackingStoreType result_ = [nSWindow backingType];
    return result_;
}

void C_NSWindow_SetBackingType(void* ptr, unsigned int value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setBackingType:value];
}

int32_t C_NSWindow_DepthLimit(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowDepth result_ = [nSWindow depthLimit];
    return result_;
}

void C_NSWindow_SetDepthLimit(void* ptr, int32_t value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setDepthLimit:value];
}

bool C_NSWindow_HasDynamicDepthLimit(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow hasDynamicDepthLimit];
    return result_;
}

void* C_NSWindow_WindowController(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowController* result_ = [nSWindow windowController];
    return result_;
}

void C_NSWindow_SetWindowController(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setWindowController:(NSWindowController*)value];
}

void* C_NSWindow_AttachedSheet(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindow* result_ = [nSWindow attachedSheet];
    return result_;
}

bool C_NSWindow_IsSheet(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isSheet];
    return result_;
}

void* C_NSWindow_SheetParent(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindow* result_ = [nSWindow sheetParent];
    return result_;
}

Array C_NSWindow_Sheets(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSArray* result_ = [nSWindow sheets];
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

CGRect C_NSWindow_Frame(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSRect result_ = [nSWindow frame];
    return result_;
}

CGSize C_NSWindow_AspectRatio(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSSize result_ = [nSWindow aspectRatio];
    return result_;
}

void C_NSWindow_SetAspectRatio(void* ptr, CGSize value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setAspectRatio:value];
}

CGSize C_NSWindow_MinSize(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSSize result_ = [nSWindow minSize];
    return result_;
}

void C_NSWindow_SetMinSize(void* ptr, CGSize value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setMinSize:value];
}

CGSize C_NSWindow_MaxSize(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSSize result_ = [nSWindow maxSize];
    return result_;
}

void C_NSWindow_SetMaxSize(void* ptr, CGSize value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setMaxSize:value];
}

bool C_NSWindow_IsZoomed(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isZoomed];
    return result_;
}

unsigned int C_NSWindow_ResizeFlags(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSEventModifierFlags result_ = [nSWindow resizeFlags];
    return result_;
}

CGSize C_NSWindow_ResizeIncrements(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSSize result_ = [nSWindow resizeIncrements];
    return result_;
}

void C_NSWindow_SetResizeIncrements(void* ptr, CGSize value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setResizeIncrements:value];
}

bool C_NSWindow_PreservesContentDuringLiveResize(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow preservesContentDuringLiveResize];
    return result_;
}

void C_NSWindow_SetPreservesContentDuringLiveResize(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setPreservesContentDuringLiveResize:value];
}

bool C_NSWindow_InLiveResize(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow inLiveResize];
    return result_;
}

CGSize C_NSWindow_ContentAspectRatio(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSSize result_ = [nSWindow contentAspectRatio];
    return result_;
}

void C_NSWindow_SetContentAspectRatio(void* ptr, CGSize value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setContentAspectRatio:value];
}

CGSize C_NSWindow_ContentMinSize(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSSize result_ = [nSWindow contentMinSize];
    return result_;
}

void C_NSWindow_SetContentMinSize(void* ptr, CGSize value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setContentMinSize:value];
}

CGSize C_NSWindow_ContentMaxSize(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSSize result_ = [nSWindow contentMaxSize];
    return result_;
}

void C_NSWindow_SetContentMaxSize(void* ptr, CGSize value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setContentMaxSize:value];
}

CGSize C_NSWindow_ContentResizeIncrements(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSSize result_ = [nSWindow contentResizeIncrements];
    return result_;
}

void C_NSWindow_SetContentResizeIncrements(void* ptr, CGSize value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setContentResizeIncrements:value];
}

void* C_NSWindow_ContentLayoutGuide(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    id result_ = [nSWindow contentLayoutGuide];
    return result_;
}

CGRect C_NSWindow_ContentLayoutRect(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSRect result_ = [nSWindow contentLayoutRect];
    return result_;
}

CGSize C_NSWindow_MaxFullScreenContentSize(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSSize result_ = [nSWindow maxFullScreenContentSize];
    return result_;
}

void C_NSWindow_SetMaxFullScreenContentSize(void* ptr, CGSize value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setMaxFullScreenContentSize:value];
}

CGSize C_NSWindow_MinFullScreenContentSize(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSSize result_ = [nSWindow minFullScreenContentSize];
    return result_;
}

void C_NSWindow_SetMinFullScreenContentSize(void* ptr, CGSize value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setMinFullScreenContentSize:value];
}

int C_NSWindow_Level(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowLevel result_ = [nSWindow level];
    return result_;
}

void C_NSWindow_SetLevel(void* ptr, int value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setLevel:value];
}

bool C_NSWindow_IsVisible(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isVisible];
    return result_;
}

unsigned int C_NSWindow_OcclusionState(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowOcclusionState result_ = [nSWindow occlusionState];
    return result_;
}

void* C_NSWindow_FrameAutosaveName(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowFrameAutosaveName result_ = [nSWindow frameAutosaveName];
    return result_;
}

void* C_NSWindow_StringWithSavedFrame(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowPersistableFrameDescriptor result_ = [nSWindow stringWithSavedFrame];
    return result_;
}

bool C_NSWindow_IsKeyWindow(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isKeyWindow];
    return result_;
}

bool C_NSWindow_CanBecomeKeyWindow(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow canBecomeKeyWindow];
    return result_;
}

bool C_NSWindow_IsMainWindow(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isMainWindow];
    return result_;
}

bool C_NSWindow_CanBecomeMainWindow(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow canBecomeMainWindow];
    return result_;
}

void* C_NSWindow_Toolbar(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSToolbar* result_ = [nSWindow toolbar];
    return result_;
}

void C_NSWindow_SetToolbar(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setToolbar:(NSToolbar*)value];
}

Array C_NSWindow_ChildWindows(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSArray* result_ = [nSWindow childWindows];
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

void* C_NSWindow_ParentWindow(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindow* result_ = [nSWindow parentWindow];
    return result_;
}

void C_NSWindow_SetParentWindow(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setParentWindow:(NSWindow*)value];
}

void* C_NSWindow_DefaultButtonCell(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSButtonCell* result_ = [nSWindow defaultButtonCell];
    return result_;
}

void C_NSWindow_SetDefaultButtonCell(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setDefaultButtonCell:(NSButtonCell*)value];
}

bool C_NSWindow_IsExcludedFromWindowsMenu(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isExcludedFromWindowsMenu];
    return result_;
}

void C_NSWindow_SetExcludedFromWindowsMenu(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setExcludedFromWindowsMenu:value];
}

bool C_NSWindow_AreCursorRectsEnabled(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow areCursorRectsEnabled];
    return result_;
}

bool C_NSWindow_ShowsToolbarButton(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow showsToolbarButton];
    return result_;
}

void C_NSWindow_SetShowsToolbarButton(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setShowsToolbarButton:value];
}

bool C_NSWindow_TitlebarAppearsTransparent(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow titlebarAppearsTransparent];
    return result_;
}

void C_NSWindow_SetTitlebarAppearsTransparent(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setTitlebarAppearsTransparent:value];
}

int C_NSWindow_ToolbarStyle(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowToolbarStyle result_ = [nSWindow toolbarStyle];
    return result_;
}

void C_NSWindow_SetToolbarStyle(void* ptr, int value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setToolbarStyle:value];
}

int C_NSWindow_TitlebarSeparatorStyle(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSTitlebarSeparatorStyle result_ = [nSWindow titlebarSeparatorStyle];
    return result_;
}

void C_NSWindow_SetTitlebarSeparatorStyle(void* ptr, int value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setTitlebarSeparatorStyle:value];
}

Array C_NSWindow_TitlebarAccessoryViewControllers(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSArray* result_ = [nSWindow titlebarAccessoryViewControllers];
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

void C_NSWindow_SetTitlebarAccessoryViewControllers(void* ptr, Array value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSTitlebarAccessoryViewController*)(NSTitlebarAccessoryViewController*)p];
    	}
    }
    [nSWindow setTitlebarAccessoryViewControllers:objcValue];
}

bool C_NSWindow_Window_AllowsAutomaticWindowTabbing() {
    BOOL result_ = [NSWindow allowsAutomaticWindowTabbing];
    return result_;
}

void C_NSWindow_Window_SetAllowsAutomaticWindowTabbing(bool value) {
    [NSWindow setAllowsAutomaticWindowTabbing:value];
}

int C_NSWindow_Window_UserTabbingPreference() {
    NSWindowUserTabbingPreference result_ = [NSWindow userTabbingPreference];
    return result_;
}

void* C_NSWindow_Tab(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowTab* result_ = [nSWindow tab];
    return result_;
}

void* C_NSWindow_TabbingIdentifier(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowTabbingIdentifier result_ = [nSWindow tabbingIdentifier];
    return result_;
}

void C_NSWindow_SetTabbingIdentifier(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setTabbingIdentifier:(NSString*)value];
}

int C_NSWindow_TabbingMode(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowTabbingMode result_ = [nSWindow tabbingMode];
    return result_;
}

void C_NSWindow_SetTabbingMode(void* ptr, int value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setTabbingMode:value];
}

Array C_NSWindow_TabbedWindows(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSArray* result_ = [nSWindow tabbedWindows];
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

void* C_NSWindow_TabGroup(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowTabGroup* result_ = [nSWindow tabGroup];
    return result_;
}

bool C_NSWindow_AllowsToolTipsWhenApplicationIsInactive(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow allowsToolTipsWhenApplicationIsInactive];
    return result_;
}

void C_NSWindow_SetAllowsToolTipsWhenApplicationIsInactive(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setAllowsToolTipsWhenApplicationIsInactive:value];
}

void* C_NSWindow_CurrentEvent(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSEvent* result_ = [nSWindow currentEvent];
    return result_;
}

void* C_NSWindow_InitialFirstResponder(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSView* result_ = [nSWindow initialFirstResponder];
    return result_;
}

void C_NSWindow_SetInitialFirstResponder(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setInitialFirstResponder:(NSView*)value];
}

void* C_NSWindow_FirstResponder(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSResponder* result_ = [nSWindow firstResponder];
    return result_;
}

unsigned int C_NSWindow_KeyViewSelectionDirection(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSSelectionDirection result_ = [nSWindow keyViewSelectionDirection];
    return result_;
}

bool C_NSWindow_AutorecalculatesKeyViewLoop(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow autorecalculatesKeyViewLoop];
    return result_;
}

void C_NSWindow_SetAutorecalculatesKeyViewLoop(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setAutorecalculatesKeyViewLoop:value];
}

bool C_NSWindow_AcceptsMouseMovedEvents(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow acceptsMouseMovedEvents];
    return result_;
}

void C_NSWindow_SetAcceptsMouseMovedEvents(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setAcceptsMouseMovedEvents:value];
}

bool C_NSWindow_IgnoresMouseEvents(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow ignoresMouseEvents];
    return result_;
}

void C_NSWindow_SetIgnoresMouseEvents(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setIgnoresMouseEvents:value];
}

CGPoint C_NSWindow_MouseLocationOutsideOfEventStream(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSPoint result_ = [nSWindow mouseLocationOutsideOfEventStream];
    return result_;
}

bool C_NSWindow_IsRestorable(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isRestorable];
    return result_;
}

void C_NSWindow_SetRestorable(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setRestorable:value];
}

bool C_NSWindow_ViewsNeedDisplay(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow viewsNeedDisplay];
    return result_;
}

void C_NSWindow_SetViewsNeedDisplay(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setViewsNeedDisplay:value];
}

bool C_NSWindow_AllowsConcurrentViewDrawing(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow allowsConcurrentViewDrawing];
    return result_;
}

void C_NSWindow_SetAllowsConcurrentViewDrawing(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setAllowsConcurrentViewDrawing:value];
}

int C_NSWindow_AnimationBehavior(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowAnimationBehavior result_ = [nSWindow animationBehavior];
    return result_;
}

void C_NSWindow_SetAnimationBehavior(void* ptr, int value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setAnimationBehavior:value];
}

bool C_NSWindow_IsDocumentEdited(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isDocumentEdited];
    return result_;
}

void C_NSWindow_SetDocumentEdited(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setDocumentEdited:value];
}

double C_NSWindow_BackingScaleFactor(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    CGFloat result_ = [nSWindow backingScaleFactor];
    return result_;
}

void* C_NSWindow_Title(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSString* result_ = [nSWindow title];
    return result_;
}

void C_NSWindow_SetTitle(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setTitle:(NSString*)value];
}

void* C_NSWindow_Subtitle(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSString* result_ = [nSWindow subtitle];
    return result_;
}

void C_NSWindow_SetSubtitle(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setSubtitle:(NSString*)value];
}

int C_NSWindow_TitleVisibility(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowTitleVisibility result_ = [nSWindow titleVisibility];
    return result_;
}

void C_NSWindow_SetTitleVisibility(void* ptr, int value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setTitleVisibility:value];
}

void* C_NSWindow_RepresentedFilename(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSString* result_ = [nSWindow representedFilename];
    return result_;
}

void C_NSWindow_SetRepresentedFilename(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setRepresentedFilename:(NSString*)value];
}

void* C_NSWindow_RepresentedURL(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSURL* result_ = [nSWindow representedURL];
    return result_;
}

void C_NSWindow_SetRepresentedURL(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setRepresentedURL:(NSURL*)value];
}

void* C_NSWindow_Screen(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSScreen* result_ = [nSWindow screen];
    return result_;
}

void* C_NSWindow_DeepestScreen(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSScreen* result_ = [nSWindow deepestScreen];
    return result_;
}

bool C_NSWindow_DisplaysWhenScreenProfileChanges(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow displaysWhenScreenProfileChanges];
    return result_;
}

void C_NSWindow_SetDisplaysWhenScreenProfileChanges(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setDisplaysWhenScreenProfileChanges:value];
}

bool C_NSWindow_IsMovableByWindowBackground(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isMovableByWindowBackground];
    return result_;
}

void C_NSWindow_SetMovableByWindowBackground(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setMovableByWindowBackground:value];
}

bool C_NSWindow_IsMovable(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isMovable];
    return result_;
}

void C_NSWindow_SetMovable(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setMovable:value];
}

bool C_NSWindow_IsReleasedWhenClosed(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isReleasedWhenClosed];
    return result_;
}

void C_NSWindow_SetReleasedWhenClosed(void* ptr, bool value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setReleasedWhenClosed:value];
}

bool C_NSWindow_IsMiniaturized(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isMiniaturized];
    return result_;
}

void* C_NSWindow_MiniwindowImage(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSImage* result_ = [nSWindow miniwindowImage];
    return result_;
}

void C_NSWindow_SetMiniwindowImage(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setMiniwindowImage:(NSImage*)value];
}

void* C_NSWindow_MiniwindowTitle(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSString* result_ = [nSWindow miniwindowTitle];
    return result_;
}

void C_NSWindow_SetMiniwindowTitle(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setMiniwindowTitle:(NSString*)value];
}

void* C_NSWindow_DockTile(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSDockTile* result_ = [nSWindow dockTile];
    return result_;
}

bool C_NSWindow_HasCloseBox(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow hasCloseBox];
    return result_;
}

bool C_NSWindow_HasTitleBar(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow hasTitleBar];
    return result_;
}

int C_NSWindow_OrderedIndex(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSInteger result_ = [nSWindow orderedIndex];
    return result_;
}

void C_NSWindow_SetOrderedIndex(void* ptr, int value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setOrderedIndex:value];
}

void* C_NSWindow_AppearanceSource(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSObject* result_ = [nSWindow appearanceSource];
    return result_;
}

void C_NSWindow_SetAppearanceSource(void* ptr, void* value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setAppearanceSource:(NSObject*)value];
}

bool C_NSWindow_IsFloatingPanel(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isFloatingPanel];
    return result_;
}

bool C_NSWindow_IsMiniaturizable(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isMiniaturizable];
    return result_;
}

bool C_NSWindow_IsModalPanel(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isModalPanel];
    return result_;
}

bool C_NSWindow_IsResizable(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isResizable];
    return result_;
}

unsigned int C_NSWindow_StyleMask(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSWindowStyleMask result_ = [nSWindow styleMask];
    return result_;
}

void C_NSWindow_SetStyleMask(void* ptr, unsigned int value) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    [nSWindow setStyleMask:value];
}

int C_NSWindow_WindowTitlebarLayoutDirection(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    NSUserInterfaceLayoutDirection result_ = [nSWindow windowTitlebarLayoutDirection];
    return result_;
}

bool C_NSWindow_IsZoomable(void* ptr) {
    NSWindow* nSWindow = (NSWindow*)ptr;
    BOOL result_ = [nSWindow isZoomable];
    return result_;
}
