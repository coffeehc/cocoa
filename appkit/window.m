#import <Appkit/Appkit.h>
#import "window.h"

void* C_Window_Alloc() {
	return [NSWindow alloc];
}

void* C_NSWindow_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSWindow* result = [nSWindow initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag];
	return result;
}

void* C_NSWindow_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSWindow* result = [nSWindow initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag screen:(NSScreen*)screen];
	return result;
}

void* C_NSWindow_Init(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSWindow* result = [nSWindow init];
	return result;
}

void* C_NSWindow_WindowWithContentViewController(void* contentViewController) {
	NSWindow* result = [NSWindow windowWithContentViewController:(NSViewController*)contentViewController];
	return result;
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
	bool result = [nSWindow autorecalculatesContentBorderThicknessForEdge:edge];
	return result;
}

void C_NSWindow_SetAutorecalculatesContentBorderThickness_ForEdge(void* ptr, bool flag, unsigned int edge) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setAutorecalculatesContentBorderThickness:flag forEdge:edge];
}

double C_NSWindow_ContentBorderThicknessForEdge(void* ptr, unsigned int edge) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	double result = [nSWindow contentBorderThicknessForEdge:edge];
	return result;
}

void C_NSWindow_SetContentBorderThickness_ForEdge(void* ptr, double thickness, unsigned int edge) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setContentBorderThickness:thickness forEdge:edge];
}

CGRect C_NSWindow_WindowContentRectForFrameRect_StyleMask(CGRect fRect, unsigned int style) {
	CGRect result = [NSWindow contentRectForFrameRect:fRect styleMask:style];
	return result;
}

CGRect C_NSWindow_WindowFrameRectForContentRect_StyleMask(CGRect cRect, unsigned int style) {
	CGRect result = [NSWindow frameRectForContentRect:cRect styleMask:style];
	return result;
}

double C_NSWindow_WindowMinFrameWidthWithTitle_StyleMask(void* title, unsigned int style) {
	double result = [NSWindow minFrameWidthWithTitle:(NSString*)title styleMask:style];
	return result;
}

CGRect C_NSWindow_ContentRectForFrameRect(void* ptr, CGRect frameRect) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGRect result = [nSWindow contentRectForFrameRect:frameRect];
	return result;
}

CGRect C_NSWindow_FrameRectForContentRect(void* ptr, CGRect contentRect) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGRect result = [nSWindow frameRectForContentRect:contentRect];
	return result;
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
	CGRect result = [nSWindow constrainFrameRect:frameRect toScreen:(NSScreen*)screen];
	return result;
}

CGPoint C_NSWindow_CascadeTopLeftFromPoint(void* ptr, CGPoint topLeftPoint) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGPoint result = [nSWindow cascadeTopLeftFromPoint:topLeftPoint];
	return result;
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
	double result = [nSWindow animationResizeTime:newFrame];
	return result;
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

void C_NSWindow_WindowRemoveFrameUsingName(void* name) {
	[NSWindow removeFrameUsingName:(NSString*)name];
}

bool C_NSWindow_SetFrameUsingName(void* ptr, void* name) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow setFrameUsingName:(NSString*)name];
	return result;
}

bool C_NSWindow_SetFrameUsingName_Force(void* ptr, void* name, bool force) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow setFrameUsingName:(NSString*)name force:force];
	return result;
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
	NSText* result = [nSWindow fieldEditor:createFlag forObject:(id)object];
	return result;
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

void* C_NSWindow_WindowStandardWindowButton_ForStyleMask(unsigned int b, unsigned int styleMask) {
	NSButton* result = [NSWindow standardWindowButton:b forStyleMask:styleMask];
	return result;
}

void* C_NSWindow_StandardWindowButton(void* ptr, unsigned int b) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSButton* result = [nSWindow standardWindowButton:b];
	return result;
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
	bool result = [nSWindow makeFirstResponder:(NSResponder*)responder];
	return result;
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
	int result = [NSWindow windowNumberAtPoint:point belowWindowWithWindowNumber:windowNumber];
	return result;
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

void C_NSWindow_UnregisterDraggedTypes(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow unregisterDraggedTypes];
}

CGRect C_NSWindow_ConvertRectFromBacking(void* ptr, CGRect rect) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGRect result = [nSWindow convertRectFromBacking:rect];
	return result;
}

CGRect C_NSWindow_ConvertRectToBacking(void* ptr, CGRect rect) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGRect result = [nSWindow convertRectToBacking:rect];
	return result;
}

CGRect C_NSWindow_ConvertRectToScreen(void* ptr, CGRect rect) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGRect result = [nSWindow convertRectToScreen:rect];
	return result;
}

CGRect C_NSWindow_ConvertRectFromScreen(void* ptr, CGRect rect) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGRect result = [nSWindow convertRectFromScreen:rect];
	return result;
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
	NSData* result = [nSWindow dataWithEPSInsideRect:rect];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}

Array C_NSWindow_DataWithPDFInsideRect(void* ptr, CGRect rect) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSData* result = [nSWindow dataWithPDFInsideRect:rect];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}

void C_NSWindow_UpdateConstraintsIfNeeded(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow updateConstraintsIfNeeded];
}

void C_NSWindow_LayoutIfNeeded(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow layoutIfNeeded];
}

int C_NSWindow_AnchorAttributeForOrientation(void* ptr, int orientation) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	int result = [nSWindow anchorAttributeForOrientation:orientation];
	return result;
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
	id result = [nSWindow handleCloseScriptCommand:(NSCloseCommand*)command];
	return result;
}

void* C_NSWindow_HandlePrintScriptCommand(void* ptr, void* command) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	id result = [nSWindow handlePrintScriptCommand:(NSScriptCommand*)command];
	return result;
}

void* C_NSWindow_HandleSaveScriptCommand(void* ptr, void* command) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	id result = [nSWindow handleSaveScriptCommand:(NSScriptCommand*)command];
	return result;
}

bool C_NSWindow_CanRepresentDisplayGamut(void* ptr, int displayGamut) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow canRepresentDisplayGamut:displayGamut];
	return result;
}

CGPoint C_NSWindow_ConvertPointFromScreen(void* ptr, CGPoint point) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGPoint result = [nSWindow convertPointFromScreen:point];
	return result;
}

CGPoint C_NSWindow_ConvertPointToScreen(void* ptr, CGPoint point) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGPoint result = [nSWindow convertPointToScreen:point];
	return result;
}

CGPoint C_NSWindow_ConvertPointFromBacking(void* ptr, CGPoint point) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGPoint result = [nSWindow convertPointFromBacking:point];
	return result;
}

CGPoint C_NSWindow_ConvertPointToBacking(void* ptr, CGPoint point) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGPoint result = [nSWindow convertPointToBacking:point];
	return result;
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
	bool result = [nSWindow setFrameAutosaveName:(NSString*)name];
	return result;
}

void* C_NSWindow_Delegate(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	id result = [nSWindow delegate];
	return result;
}

void C_NSWindow_SetDelegate(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setDelegate:(id)value];
}

void* C_NSWindow_ContentViewController(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSViewController* result = [nSWindow contentViewController];
	return result;
}

void C_NSWindow_SetContentViewController(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setContentViewController:(NSViewController*)value];
}

void* C_NSWindow_ContentView(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSView* result = [nSWindow contentView];
	return result;
}

void C_NSWindow_SetContentView(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setContentView:(NSView*)value];
}

bool C_NSWindow_WorksWhenModal(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow worksWhenModal];
	return result;
}

double C_NSWindow_AlphaValue(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	double result = [nSWindow alphaValue];
	return result;
}

void C_NSWindow_SetAlphaValue(void* ptr, double value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setAlphaValue:value];
}

void* C_NSWindow_BackgroundColor(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSColor* result = [nSWindow backgroundColor];
	return result;
}

void C_NSWindow_SetBackgroundColor(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setBackgroundColor:(NSColor*)value];
}

void* C_NSWindow_ColorSpace(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSColorSpace* result = [nSWindow colorSpace];
	return result;
}

void C_NSWindow_SetColorSpace(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setColorSpace:(NSColorSpace*)value];
}

bool C_NSWindow_CanHide(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow canHide];
	return result;
}

void C_NSWindow_SetCanHide(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setCanHide:value];
}

bool C_NSWindow_IsOnActiveSpace(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isOnActiveSpace];
	return result;
}

bool C_NSWindow_HidesOnDeactivate(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow hidesOnDeactivate];
	return result;
}

void C_NSWindow_SetHidesOnDeactivate(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setHidesOnDeactivate:value];
}

unsigned int C_NSWindow_CollectionBehavior(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	unsigned int result = [nSWindow collectionBehavior];
	return result;
}

void C_NSWindow_SetCollectionBehavior(void* ptr, unsigned int value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setCollectionBehavior:value];
}

bool C_NSWindow_IsOpaque(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isOpaque];
	return result;
}

void C_NSWindow_SetOpaque(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setOpaque:value];
}

bool C_NSWindow_HasShadow(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow hasShadow];
	return result;
}

void C_NSWindow_SetHasShadow(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setHasShadow:value];
}

bool C_NSWindow_PreventsApplicationTerminationWhenModal(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow preventsApplicationTerminationWhenModal];
	return result;
}

void C_NSWindow_SetPreventsApplicationTerminationWhenModal(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setPreventsApplicationTerminationWhenModal:value];
}

int C_NSWindow_WindowNumber(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	int result = [nSWindow windowNumber];
	return result;
}

bool C_NSWindow_CanBecomeVisibleWithoutLogin(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow canBecomeVisibleWithoutLogin];
	return result;
}

void C_NSWindow_SetCanBecomeVisibleWithoutLogin(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setCanBecomeVisibleWithoutLogin:value];
}

unsigned int C_NSWindow_SharingType(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	unsigned int result = [nSWindow sharingType];
	return result;
}

void C_NSWindow_SetSharingType(void* ptr, unsigned int value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setSharingType:value];
}

unsigned int C_NSWindow_BackingType(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	unsigned int result = [nSWindow backingType];
	return result;
}

void C_NSWindow_SetBackingType(void* ptr, unsigned int value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setBackingType:value];
}

int32_t C_NSWindow_DepthLimit(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	int32_t result = [nSWindow depthLimit];
	return result;
}

void C_NSWindow_SetDepthLimit(void* ptr, int32_t value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setDepthLimit:value];
}

bool C_NSWindow_HasDynamicDepthLimit(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow hasDynamicDepthLimit];
	return result;
}

void* C_NSWindow_WindowController(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSWindowController* result = [nSWindow windowController];
	return result;
}

void C_NSWindow_SetWindowController(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setWindowController:(NSWindowController*)value];
}

void* C_NSWindow_AttachedSheet(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSWindow* result = [nSWindow attachedSheet];
	return result;
}

bool C_NSWindow_IsSheet(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isSheet];
	return result;
}

void* C_NSWindow_SheetParent(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSWindow* result = [nSWindow sheetParent];
	return result;
}

CGRect C_NSWindow_Frame(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGRect result = [nSWindow frame];
	return result;
}

CGSize C_NSWindow_AspectRatio(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGSize result = [nSWindow aspectRatio];
	return result;
}

void C_NSWindow_SetAspectRatio(void* ptr, CGSize value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setAspectRatio:value];
}

CGSize C_NSWindow_MinSize(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGSize result = [nSWindow minSize];
	return result;
}

void C_NSWindow_SetMinSize(void* ptr, CGSize value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setMinSize:value];
}

CGSize C_NSWindow_MaxSize(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGSize result = [nSWindow maxSize];
	return result;
}

void C_NSWindow_SetMaxSize(void* ptr, CGSize value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setMaxSize:value];
}

bool C_NSWindow_IsZoomed(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isZoomed];
	return result;
}

unsigned int C_NSWindow_ResizeFlags(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	unsigned int result = [nSWindow resizeFlags];
	return result;
}

CGSize C_NSWindow_ResizeIncrements(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGSize result = [nSWindow resizeIncrements];
	return result;
}

void C_NSWindow_SetResizeIncrements(void* ptr, CGSize value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setResizeIncrements:value];
}

bool C_NSWindow_PreservesContentDuringLiveResize(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow preservesContentDuringLiveResize];
	return result;
}

void C_NSWindow_SetPreservesContentDuringLiveResize(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setPreservesContentDuringLiveResize:value];
}

bool C_NSWindow_InLiveResize(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow inLiveResize];
	return result;
}

CGSize C_NSWindow_ContentAspectRatio(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGSize result = [nSWindow contentAspectRatio];
	return result;
}

void C_NSWindow_SetContentAspectRatio(void* ptr, CGSize value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setContentAspectRatio:value];
}

CGSize C_NSWindow_ContentMinSize(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGSize result = [nSWindow contentMinSize];
	return result;
}

void C_NSWindow_SetContentMinSize(void* ptr, CGSize value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setContentMinSize:value];
}

CGSize C_NSWindow_ContentMaxSize(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGSize result = [nSWindow contentMaxSize];
	return result;
}

void C_NSWindow_SetContentMaxSize(void* ptr, CGSize value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setContentMaxSize:value];
}

CGSize C_NSWindow_ContentResizeIncrements(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGSize result = [nSWindow contentResizeIncrements];
	return result;
}

void C_NSWindow_SetContentResizeIncrements(void* ptr, CGSize value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setContentResizeIncrements:value];
}

void* C_NSWindow_ContentLayoutGuide(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	id result = [nSWindow contentLayoutGuide];
	return result;
}

CGRect C_NSWindow_ContentLayoutRect(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGRect result = [nSWindow contentLayoutRect];
	return result;
}

CGSize C_NSWindow_MaxFullScreenContentSize(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGSize result = [nSWindow maxFullScreenContentSize];
	return result;
}

void C_NSWindow_SetMaxFullScreenContentSize(void* ptr, CGSize value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setMaxFullScreenContentSize:value];
}

CGSize C_NSWindow_MinFullScreenContentSize(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGSize result = [nSWindow minFullScreenContentSize];
	return result;
}

void C_NSWindow_SetMinFullScreenContentSize(void* ptr, CGSize value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setMinFullScreenContentSize:value];
}

int C_NSWindow_Level(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	int result = [nSWindow level];
	return result;
}

void C_NSWindow_SetLevel(void* ptr, int value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setLevel:value];
}

bool C_NSWindow_IsVisible(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isVisible];
	return result;
}

unsigned int C_NSWindow_OcclusionState(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	unsigned int result = [nSWindow occlusionState];
	return result;
}

void* C_NSWindow_FrameAutosaveName(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSString* result = [nSWindow frameAutosaveName];
	return result;
}

void* C_NSWindow_StringWithSavedFrame(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSString* result = [nSWindow stringWithSavedFrame];
	return result;
}

bool C_NSWindow_IsKeyWindow(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isKeyWindow];
	return result;
}

bool C_NSWindow_CanBecomeKeyWindow(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow canBecomeKeyWindow];
	return result;
}

bool C_NSWindow_IsMainWindow(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isMainWindow];
	return result;
}

bool C_NSWindow_CanBecomeMainWindow(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow canBecomeMainWindow];
	return result;
}

void* C_NSWindow_Toolbar(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSToolbar* result = [nSWindow toolbar];
	return result;
}

void C_NSWindow_SetToolbar(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setToolbar:(NSToolbar*)value];
}

void* C_NSWindow_ParentWindow(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSWindow* result = [nSWindow parentWindow];
	return result;
}

void C_NSWindow_SetParentWindow(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setParentWindow:(NSWindow*)value];
}

void* C_NSWindow_DefaultButtonCell(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSButtonCell* result = [nSWindow defaultButtonCell];
	return result;
}

void C_NSWindow_SetDefaultButtonCell(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setDefaultButtonCell:(NSButtonCell*)value];
}

bool C_NSWindow_IsExcludedFromWindowsMenu(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isExcludedFromWindowsMenu];
	return result;
}

void C_NSWindow_SetExcludedFromWindowsMenu(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setExcludedFromWindowsMenu:value];
}

bool C_NSWindow_AreCursorRectsEnabled(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow areCursorRectsEnabled];
	return result;
}

bool C_NSWindow_ShowsToolbarButton(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow showsToolbarButton];
	return result;
}

void C_NSWindow_SetShowsToolbarButton(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setShowsToolbarButton:value];
}

bool C_NSWindow_TitlebarAppearsTransparent(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow titlebarAppearsTransparent];
	return result;
}

void C_NSWindow_SetTitlebarAppearsTransparent(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setTitlebarAppearsTransparent:value];
}

int C_NSWindow_ToolbarStyle(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	int result = [nSWindow toolbarStyle];
	return result;
}

void C_NSWindow_SetToolbarStyle(void* ptr, int value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setToolbarStyle:value];
}

int C_NSWindow_TitlebarSeparatorStyle(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	int result = [nSWindow titlebarSeparatorStyle];
	return result;
}

void C_NSWindow_SetTitlebarSeparatorStyle(void* ptr, int value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setTitlebarSeparatorStyle:value];
}

void* C_NSWindow_Tab(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSWindowTab* result = [nSWindow tab];
	return result;
}

void* C_NSWindow_TabbingIdentifier(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSString* result = [nSWindow tabbingIdentifier];
	return result;
}

void C_NSWindow_SetTabbingIdentifier(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setTabbingIdentifier:(NSString*)value];
}

int C_NSWindow_TabbingMode(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	int result = [nSWindow tabbingMode];
	return result;
}

void C_NSWindow_SetTabbingMode(void* ptr, int value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setTabbingMode:value];
}

void* C_NSWindow_TabGroup(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSWindowTabGroup* result = [nSWindow tabGroup];
	return result;
}

bool C_NSWindow_AllowsToolTipsWhenApplicationIsInactive(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow allowsToolTipsWhenApplicationIsInactive];
	return result;
}

void C_NSWindow_SetAllowsToolTipsWhenApplicationIsInactive(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setAllowsToolTipsWhenApplicationIsInactive:value];
}

void* C_NSWindow_CurrentEvent(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSEvent* result = [nSWindow currentEvent];
	return result;
}

void* C_NSWindow_InitialFirstResponder(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSView* result = [nSWindow initialFirstResponder];
	return result;
}

void C_NSWindow_SetInitialFirstResponder(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setInitialFirstResponder:(NSView*)value];
}

void* C_NSWindow_FirstResponder(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSResponder* result = [nSWindow firstResponder];
	return result;
}

unsigned int C_NSWindow_KeyViewSelectionDirection(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	unsigned int result = [nSWindow keyViewSelectionDirection];
	return result;
}

bool C_NSWindow_AutorecalculatesKeyViewLoop(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow autorecalculatesKeyViewLoop];
	return result;
}

void C_NSWindow_SetAutorecalculatesKeyViewLoop(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setAutorecalculatesKeyViewLoop:value];
}

bool C_NSWindow_AcceptsMouseMovedEvents(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow acceptsMouseMovedEvents];
	return result;
}

void C_NSWindow_SetAcceptsMouseMovedEvents(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setAcceptsMouseMovedEvents:value];
}

bool C_NSWindow_IgnoresMouseEvents(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow ignoresMouseEvents];
	return result;
}

void C_NSWindow_SetIgnoresMouseEvents(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setIgnoresMouseEvents:value];
}

CGPoint C_NSWindow_MouseLocationOutsideOfEventStream(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	CGPoint result = [nSWindow mouseLocationOutsideOfEventStream];
	return result;
}

bool C_NSWindow_IsRestorable(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isRestorable];
	return result;
}

void C_NSWindow_SetRestorable(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setRestorable:value];
}

bool C_NSWindow_ViewsNeedDisplay(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow viewsNeedDisplay];
	return result;
}

void C_NSWindow_SetViewsNeedDisplay(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setViewsNeedDisplay:value];
}

bool C_NSWindow_AllowsConcurrentViewDrawing(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow allowsConcurrentViewDrawing];
	return result;
}

void C_NSWindow_SetAllowsConcurrentViewDrawing(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setAllowsConcurrentViewDrawing:value];
}

int C_NSWindow_AnimationBehavior(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	int result = [nSWindow animationBehavior];
	return result;
}

void C_NSWindow_SetAnimationBehavior(void* ptr, int value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setAnimationBehavior:value];
}

bool C_NSWindow_IsDocumentEdited(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isDocumentEdited];
	return result;
}

void C_NSWindow_SetDocumentEdited(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setDocumentEdited:value];
}

double C_NSWindow_BackingScaleFactor(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	double result = [nSWindow backingScaleFactor];
	return result;
}

void* C_NSWindow_Title(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSString* result = [nSWindow title];
	return result;
}

void C_NSWindow_SetTitle(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setTitle:(NSString*)value];
}

void* C_NSWindow_Subtitle(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSString* result = [nSWindow subtitle];
	return result;
}

void C_NSWindow_SetSubtitle(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setSubtitle:(NSString*)value];
}

int C_NSWindow_TitleVisibility(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	int result = [nSWindow titleVisibility];
	return result;
}

void C_NSWindow_SetTitleVisibility(void* ptr, int value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setTitleVisibility:value];
}

void* C_NSWindow_RepresentedFilename(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSString* result = [nSWindow representedFilename];
	return result;
}

void C_NSWindow_SetRepresentedFilename(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setRepresentedFilename:(NSString*)value];
}

void* C_NSWindow_RepresentedURL(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSURL* result = [nSWindow representedURL];
	return result;
}

void C_NSWindow_SetRepresentedURL(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setRepresentedURL:(NSURL*)value];
}

void* C_NSWindow_Screen(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSScreen* result = [nSWindow screen];
	return result;
}

void* C_NSWindow_DeepestScreen(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSScreen* result = [nSWindow deepestScreen];
	return result;
}

bool C_NSWindow_DisplaysWhenScreenProfileChanges(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow displaysWhenScreenProfileChanges];
	return result;
}

void C_NSWindow_SetDisplaysWhenScreenProfileChanges(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setDisplaysWhenScreenProfileChanges:value];
}

bool C_NSWindow_IsMovableByWindowBackground(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isMovableByWindowBackground];
	return result;
}

void C_NSWindow_SetMovableByWindowBackground(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setMovableByWindowBackground:value];
}

bool C_NSWindow_IsMovable(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isMovable];
	return result;
}

void C_NSWindow_SetMovable(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setMovable:value];
}

bool C_NSWindow_IsReleasedWhenClosed(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isReleasedWhenClosed];
	return result;
}

void C_NSWindow_SetReleasedWhenClosed(void* ptr, bool value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setReleasedWhenClosed:value];
}

bool C_NSWindow_IsMiniaturized(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isMiniaturized];
	return result;
}

void* C_NSWindow_MiniwindowImage(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSImage* result = [nSWindow miniwindowImage];
	return result;
}

void C_NSWindow_SetMiniwindowImage(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setMiniwindowImage:(NSImage*)value];
}

void* C_NSWindow_MiniwindowTitle(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSString* result = [nSWindow miniwindowTitle];
	return result;
}

void C_NSWindow_SetMiniwindowTitle(void* ptr, void* value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setMiniwindowTitle:(NSString*)value];
}

void* C_NSWindow_DockTile(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	NSDockTile* result = [nSWindow dockTile];
	return result;
}

bool C_NSWindow_HasCloseBox(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow hasCloseBox];
	return result;
}

bool C_NSWindow_HasTitleBar(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow hasTitleBar];
	return result;
}

int C_NSWindow_OrderedIndex(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	int result = [nSWindow orderedIndex];
	return result;
}

void C_NSWindow_SetOrderedIndex(void* ptr, int value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setOrderedIndex:value];
}

bool C_NSWindow_IsFloatingPanel(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isFloatingPanel];
	return result;
}

bool C_NSWindow_IsMiniaturizable(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isMiniaturizable];
	return result;
}

bool C_NSWindow_IsModalPanel(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isModalPanel];
	return result;
}

bool C_NSWindow_IsResizable(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isResizable];
	return result;
}

unsigned int C_NSWindow_StyleMask(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	unsigned int result = [nSWindow styleMask];
	return result;
}

void C_NSWindow_SetStyleMask(void* ptr, unsigned int value) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	[nSWindow setStyleMask:value];
}

int C_NSWindow_WindowTitlebarLayoutDirection(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	int result = [nSWindow windowTitlebarLayoutDirection];
	return result;
}

bool C_NSWindow_IsZoomable(void* ptr) {
	NSWindow* nSWindow = (NSWindow*)ptr;
	bool result = [nSWindow isZoomable];
	return result;
}
