#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_Window_Alloc();

void* C_NSWindow_WindowWithContentViewController(void* contentViewController);
void* C_NSWindow_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag);
void* C_NSWindow_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen);
void* C_NSWindow_Init(void* ptr);
void* C_NSWindow_AllocWindow();
void* C_NSWindow_NewWindow();
void* C_NSWindow_Autorelease(void* ptr);
void* C_NSWindow_Retain(void* ptr);
void C_NSWindow_ToggleFullScreen(void* ptr, void* sender);
void C_NSWindow_InvalidateShadow(void* ptr);
bool C_NSWindow_AutorecalculatesContentBorderThicknessForEdge(void* ptr, unsigned int edge);
void C_NSWindow_SetAutorecalculatesContentBorderThickness_ForEdge(void* ptr, bool flag, unsigned int edge);
double C_NSWindow_ContentBorderThicknessForEdge(void* ptr, unsigned int edge);
void C_NSWindow_SetContentBorderThickness_ForEdge(void* ptr, double thickness, unsigned int edge);
Array C_NSWindow_WindowNumbersWithOptions(unsigned int options);
CGRect C_NSWindow_Window_ContentRectForFrameRect_StyleMask(CGRect fRect, unsigned int style);
CGRect C_NSWindow_Window_FrameRectForContentRect_StyleMask(CGRect cRect, unsigned int style);
double C_NSWindow_Window_MinFrameWidthWithTitle_StyleMask(void* title, unsigned int style);
CGRect C_NSWindow_ContentRectForFrameRect(void* ptr, CGRect frameRect);
CGRect C_NSWindow_FrameRectForContentRect(void* ptr, CGRect contentRect);
void C_NSWindow_EndSheet(void* ptr, void* sheetWindow);
void C_NSWindow_EndSheet_ReturnCode(void* ptr, void* sheetWindow, int returnCode);
void C_NSWindow_SetFrameOrigin(void* ptr, CGPoint point);
void C_NSWindow_SetFrameTopLeftPoint(void* ptr, CGPoint point);
CGRect C_NSWindow_ConstrainFrameRect_ToScreen(void* ptr, CGRect frameRect, void* screen);
CGPoint C_NSWindow_CascadeTopLeftFromPoint(void* ptr, CGPoint topLeftPoint);
void C_NSWindow_SetFrame_Display(void* ptr, CGRect frameRect, bool flag);
void C_NSWindow_SetFrame_Display_Animate(void* ptr, CGRect frameRect, bool displayFlag, bool animateFlag);
double C_NSWindow_AnimationResizeTime(void* ptr, CGRect newFrame);
void C_NSWindow_PerformZoom(void* ptr, void* sender);
void C_NSWindow_Zoom(void* ptr, void* sender);
void C_NSWindow_SetContentSize(void* ptr, CGSize size);
void C_NSWindow_OrderOut(void* ptr, void* sender);
void C_NSWindow_OrderBack(void* ptr, void* sender);
void C_NSWindow_OrderFront(void* ptr, void* sender);
void C_NSWindow_OrderFrontRegardless(void* ptr);
void C_NSWindow_OrderWindow_RelativeTo(void* ptr, int place, int otherWin);
void C_NSWindow_Window_RemoveFrameUsingName(void* name);
bool C_NSWindow_SetFrameUsingName(void* ptr, void* name);
bool C_NSWindow_SetFrameUsingName_Force(void* ptr, void* name, bool force);
void C_NSWindow_SaveFrameUsingName(void* ptr, void* name);
void C_NSWindow_SetFrameFromString(void* ptr, void* _string);
void C_NSWindow_MakeKeyWindow(void* ptr);
void C_NSWindow_MakeKeyAndOrderFront(void* ptr, void* sender);
void C_NSWindow_BecomeKeyWindow(void* ptr);
void C_NSWindow_ResignKeyWindow(void* ptr);
void C_NSWindow_MakeMainWindow(void* ptr);
void C_NSWindow_BecomeMainWindow(void* ptr);
void C_NSWindow_ResignMainWindow(void* ptr);
void C_NSWindow_ToggleToolbarShown(void* ptr, void* sender);
void C_NSWindow_RunToolbarCustomizationPalette(void* ptr, void* sender);
void C_NSWindow_AddChildWindow_Ordered(void* ptr, void* childWin, int place);
void C_NSWindow_RemoveChildWindow(void* ptr, void* childWin);
void C_NSWindow_EnableKeyEquivalentForDefaultButtonCell(void* ptr);
void C_NSWindow_DisableKeyEquivalentForDefaultButtonCell(void* ptr);
void* C_NSWindow_FieldEditor_ForObject(void* ptr, bool createFlag, void* object);
void C_NSWindow_EndEditingFor(void* ptr, void* object);
void C_NSWindow_EnableCursorRects(void* ptr);
void C_NSWindow_DisableCursorRects(void* ptr);
void C_NSWindow_DiscardCursorRects(void* ptr);
void C_NSWindow_InvalidateCursorRectsForView(void* ptr, void* view);
void C_NSWindow_ResetCursorRects(void* ptr);
void* C_NSWindow_StandardWindowButton_ForStyleMask(unsigned int b, unsigned int styleMask);
void* C_NSWindow_StandardWindowButton(void* ptr, unsigned int b);
void C_NSWindow_AddTitlebarAccessoryViewController(void* ptr, void* childViewController);
void C_NSWindow_InsertTitlebarAccessoryViewController_AtIndex(void* ptr, void* childViewController, int index);
void C_NSWindow_RemoveTitlebarAccessoryViewControllerAtIndex(void* ptr, int index);
void C_NSWindow_AddTabbedWindow_Ordered(void* ptr, void* window, int ordered);
void C_NSWindow_SelectNextTab(void* ptr, void* sender);
void C_NSWindow_SelectPreviousTab(void* ptr, void* sender);
void C_NSWindow_MoveTabToNewWindow(void* ptr, void* sender);
void C_NSWindow_ToggleTabBar(void* ptr, void* sender);
void C_NSWindow_ToggleTabOverview(void* ptr, void* sender);
void C_NSWindow_PostEvent_AtStart(void* ptr, void* event, bool flag);
void C_NSWindow_SendEvent(void* ptr, void* event);
bool C_NSWindow_MakeFirstResponder(void* ptr, void* responder);
void C_NSWindow_SelectKeyViewPrecedingView(void* ptr, void* view);
void C_NSWindow_SelectKeyViewFollowingView(void* ptr, void* view);
void C_NSWindow_SelectPreviousKeyView(void* ptr, void* sender);
void C_NSWindow_SelectNextKeyView(void* ptr, void* sender);
void C_NSWindow_RecalculateKeyViewLoop(void* ptr);
int C_NSWindow_WindowNumberAtPoint_BelowWindowWithWindowNumber(CGPoint point, int windowNumber);
void C_NSWindow_PerformWindowDragWithEvent(void* ptr, void* event);
void C_NSWindow_DisableSnapshotRestoration(void* ptr);
void C_NSWindow_EnableSnapshotRestoration(void* ptr);
void C_NSWindow_Display(void* ptr);
void C_NSWindow_DisplayIfNeeded(void* ptr);
void C_NSWindow_DisableScreenUpdatesUntilFlush(void* ptr);
void C_NSWindow_Update(void* ptr);
void C_NSWindow_DragImage_At_Offset_Event_Pasteboard_Source_SlideBack(void* ptr, void* image, CGPoint baseLocation, CGSize initialOffset, void* event, void* pboard, void* sourceObj, bool slideFlag);
void C_NSWindow_RegisterForDraggedTypes(void* ptr, Array newTypes);
void C_NSWindow_UnregisterDraggedTypes(void* ptr);
CGRect C_NSWindow_ConvertRectFromBacking(void* ptr, CGRect rect);
CGRect C_NSWindow_ConvertRectToBacking(void* ptr, CGRect rect);
CGRect C_NSWindow_ConvertRectToScreen(void* ptr, CGRect rect);
CGRect C_NSWindow_ConvertRectFromScreen(void* ptr, CGRect rect);
void C_NSWindow_SetTitleWithRepresentedFilename(void* ptr, void* filename);
void C_NSWindow_Center(void* ptr);
void C_NSWindow_PerformClose(void* ptr, void* sender);
void C_NSWindow_Close(void* ptr);
void C_NSWindow_PerformMiniaturize(void* ptr, void* sender);
void C_NSWindow_Miniaturize(void* ptr, void* sender);
void C_NSWindow_Deminiaturize(void* ptr, void* sender);
void C_NSWindow_Print(void* ptr, void* sender);
void* C_NSWindow_DataWithEPSInsideRect(void* ptr, CGRect rect);
void* C_NSWindow_DataWithPDFInsideRect(void* ptr, CGRect rect);
void C_NSWindow_UpdateConstraintsIfNeeded(void* ptr);
void C_NSWindow_LayoutIfNeeded(void* ptr);
void C_NSWindow_VisualizeConstraints(void* ptr, Array constraints);
int C_NSWindow_AnchorAttributeForOrientation(void* ptr, int orientation);
void C_NSWindow_SetAnchorAttribute_ForOrientation(void* ptr, int attr, int orientation);
void C_NSWindow_SetIsMiniaturized(void* ptr, bool flag);
void C_NSWindow_SetIsVisible(void* ptr, bool flag);
void C_NSWindow_SetIsZoomed(void* ptr, bool flag);
void* C_NSWindow_HandleCloseScriptCommand(void* ptr, void* command);
void* C_NSWindow_HandlePrintScriptCommand(void* ptr, void* command);
void* C_NSWindow_HandleSaveScriptCommand(void* ptr, void* command);
bool C_NSWindow_CanRepresentDisplayGamut(void* ptr, int displayGamut);
CGPoint C_NSWindow_ConvertPointFromScreen(void* ptr, CGPoint point);
CGPoint C_NSWindow_ConvertPointToScreen(void* ptr, CGPoint point);
CGPoint C_NSWindow_ConvertPointFromBacking(void* ptr, CGPoint point);
CGPoint C_NSWindow_ConvertPointToBacking(void* ptr, CGPoint point);
void C_NSWindow_MergeAllWindows(void* ptr, void* sender);
void C_NSWindow_SetDynamicDepthLimit(void* ptr, bool flag);
bool C_NSWindow_SetFrameAutosaveName(void* ptr, void* name);
void* C_NSWindow_Delegate(void* ptr);
void C_NSWindow_SetDelegate(void* ptr, void* value);
void* C_NSWindow_ContentViewController(void* ptr);
void C_NSWindow_SetContentViewController(void* ptr, void* value);
void* C_NSWindow_ContentView(void* ptr);
void C_NSWindow_SetContentView(void* ptr, void* value);
bool C_NSWindow_WorksWhenModal(void* ptr);
double C_NSWindow_AlphaValue(void* ptr);
void C_NSWindow_SetAlphaValue(void* ptr, double value);
void* C_NSWindow_BackgroundColor(void* ptr);
void C_NSWindow_SetBackgroundColor(void* ptr, void* value);
void* C_NSWindow_ColorSpace(void* ptr);
void C_NSWindow_SetColorSpace(void* ptr, void* value);
bool C_NSWindow_CanHide(void* ptr);
void C_NSWindow_SetCanHide(void* ptr, bool value);
bool C_NSWindow_IsOnActiveSpace(void* ptr);
bool C_NSWindow_HidesOnDeactivate(void* ptr);
void C_NSWindow_SetHidesOnDeactivate(void* ptr, bool value);
unsigned int C_NSWindow_CollectionBehavior(void* ptr);
void C_NSWindow_SetCollectionBehavior(void* ptr, unsigned int value);
bool C_NSWindow_IsOpaque(void* ptr);
void C_NSWindow_SetOpaque(void* ptr, bool value);
bool C_NSWindow_HasShadow(void* ptr);
void C_NSWindow_SetHasShadow(void* ptr, bool value);
bool C_NSWindow_PreventsApplicationTerminationWhenModal(void* ptr);
void C_NSWindow_SetPreventsApplicationTerminationWhenModal(void* ptr, bool value);
int32_t C_NSWindow_Window_DefaultDepthLimit();
int C_NSWindow_WindowNumber(void* ptr);
Dictionary C_NSWindow_DeviceDescription(void* ptr);
bool C_NSWindow_CanBecomeVisibleWithoutLogin(void* ptr);
void C_NSWindow_SetCanBecomeVisibleWithoutLogin(void* ptr, bool value);
unsigned int C_NSWindow_SharingType(void* ptr);
void C_NSWindow_SetSharingType(void* ptr, unsigned int value);
unsigned int C_NSWindow_BackingType(void* ptr);
void C_NSWindow_SetBackingType(void* ptr, unsigned int value);
int32_t C_NSWindow_DepthLimit(void* ptr);
void C_NSWindow_SetDepthLimit(void* ptr, int32_t value);
bool C_NSWindow_HasDynamicDepthLimit(void* ptr);
void* C_NSWindow_WindowController(void* ptr);
void C_NSWindow_SetWindowController(void* ptr, void* value);
void* C_NSWindow_AttachedSheet(void* ptr);
bool C_NSWindow_IsSheet(void* ptr);
void* C_NSWindow_SheetParent(void* ptr);
Array C_NSWindow_Sheets(void* ptr);
CGRect C_NSWindow_Frame(void* ptr);
CGSize C_NSWindow_AspectRatio(void* ptr);
void C_NSWindow_SetAspectRatio(void* ptr, CGSize value);
CGSize C_NSWindow_MinSize(void* ptr);
void C_NSWindow_SetMinSize(void* ptr, CGSize value);
CGSize C_NSWindow_MaxSize(void* ptr);
void C_NSWindow_SetMaxSize(void* ptr, CGSize value);
bool C_NSWindow_IsZoomed(void* ptr);
unsigned int C_NSWindow_ResizeFlags(void* ptr);
CGSize C_NSWindow_ResizeIncrements(void* ptr);
void C_NSWindow_SetResizeIncrements(void* ptr, CGSize value);
bool C_NSWindow_PreservesContentDuringLiveResize(void* ptr);
void C_NSWindow_SetPreservesContentDuringLiveResize(void* ptr, bool value);
bool C_NSWindow_InLiveResize(void* ptr);
CGSize C_NSWindow_ContentAspectRatio(void* ptr);
void C_NSWindow_SetContentAspectRatio(void* ptr, CGSize value);
CGSize C_NSWindow_ContentMinSize(void* ptr);
void C_NSWindow_SetContentMinSize(void* ptr, CGSize value);
CGSize C_NSWindow_ContentMaxSize(void* ptr);
void C_NSWindow_SetContentMaxSize(void* ptr, CGSize value);
CGSize C_NSWindow_ContentResizeIncrements(void* ptr);
void C_NSWindow_SetContentResizeIncrements(void* ptr, CGSize value);
void* C_NSWindow_ContentLayoutGuide(void* ptr);
CGRect C_NSWindow_ContentLayoutRect(void* ptr);
CGSize C_NSWindow_MaxFullScreenContentSize(void* ptr);
void C_NSWindow_SetMaxFullScreenContentSize(void* ptr, CGSize value);
CGSize C_NSWindow_MinFullScreenContentSize(void* ptr);
void C_NSWindow_SetMinFullScreenContentSize(void* ptr, CGSize value);
int C_NSWindow_Level(void* ptr);
void C_NSWindow_SetLevel(void* ptr, int value);
bool C_NSWindow_IsVisible(void* ptr);
unsigned int C_NSWindow_OcclusionState(void* ptr);
void* C_NSWindow_FrameAutosaveName(void* ptr);
void* C_NSWindow_StringWithSavedFrame(void* ptr);
bool C_NSWindow_IsKeyWindow(void* ptr);
bool C_NSWindow_CanBecomeKeyWindow(void* ptr);
bool C_NSWindow_IsMainWindow(void* ptr);
bool C_NSWindow_CanBecomeMainWindow(void* ptr);
void* C_NSWindow_Toolbar(void* ptr);
void C_NSWindow_SetToolbar(void* ptr, void* value);
Array C_NSWindow_ChildWindows(void* ptr);
void* C_NSWindow_ParentWindow(void* ptr);
void C_NSWindow_SetParentWindow(void* ptr, void* value);
void* C_NSWindow_DefaultButtonCell(void* ptr);
void C_NSWindow_SetDefaultButtonCell(void* ptr, void* value);
bool C_NSWindow_IsExcludedFromWindowsMenu(void* ptr);
void C_NSWindow_SetExcludedFromWindowsMenu(void* ptr, bool value);
bool C_NSWindow_AreCursorRectsEnabled(void* ptr);
bool C_NSWindow_ShowsToolbarButton(void* ptr);
void C_NSWindow_SetShowsToolbarButton(void* ptr, bool value);
bool C_NSWindow_TitlebarAppearsTransparent(void* ptr);
void C_NSWindow_SetTitlebarAppearsTransparent(void* ptr, bool value);
int C_NSWindow_ToolbarStyle(void* ptr);
void C_NSWindow_SetToolbarStyle(void* ptr, int value);
int C_NSWindow_TitlebarSeparatorStyle(void* ptr);
void C_NSWindow_SetTitlebarSeparatorStyle(void* ptr, int value);
Array C_NSWindow_TitlebarAccessoryViewControllers(void* ptr);
void C_NSWindow_SetTitlebarAccessoryViewControllers(void* ptr, Array value);
bool C_NSWindow_AllowsAutomaticWindowTabbing();
void C_NSWindow_SetAllowsAutomaticWindowTabbing(bool value);
int C_NSWindow_Window_UserTabbingPreference();
void* C_NSWindow_Tab(void* ptr);
void* C_NSWindow_TabbingIdentifier(void* ptr);
void C_NSWindow_SetTabbingIdentifier(void* ptr, void* value);
int C_NSWindow_TabbingMode(void* ptr);
void C_NSWindow_SetTabbingMode(void* ptr, int value);
Array C_NSWindow_TabbedWindows(void* ptr);
void* C_NSWindow_TabGroup(void* ptr);
bool C_NSWindow_AllowsToolTipsWhenApplicationIsInactive(void* ptr);
void C_NSWindow_SetAllowsToolTipsWhenApplicationIsInactive(void* ptr, bool value);
void* C_NSWindow_CurrentEvent(void* ptr);
void* C_NSWindow_InitialFirstResponder(void* ptr);
void C_NSWindow_SetInitialFirstResponder(void* ptr, void* value);
void* C_NSWindow_FirstResponder(void* ptr);
unsigned int C_NSWindow_KeyViewSelectionDirection(void* ptr);
bool C_NSWindow_AutorecalculatesKeyViewLoop(void* ptr);
void C_NSWindow_SetAutorecalculatesKeyViewLoop(void* ptr, bool value);
bool C_NSWindow_AcceptsMouseMovedEvents(void* ptr);
void C_NSWindow_SetAcceptsMouseMovedEvents(void* ptr, bool value);
bool C_NSWindow_IgnoresMouseEvents(void* ptr);
void C_NSWindow_SetIgnoresMouseEvents(void* ptr, bool value);
CGPoint C_NSWindow_MouseLocationOutsideOfEventStream(void* ptr);
bool C_NSWindow_IsRestorable(void* ptr);
void C_NSWindow_SetRestorable(void* ptr, bool value);
bool C_NSWindow_ViewsNeedDisplay(void* ptr);
void C_NSWindow_SetViewsNeedDisplay(void* ptr, bool value);
bool C_NSWindow_AllowsConcurrentViewDrawing(void* ptr);
void C_NSWindow_SetAllowsConcurrentViewDrawing(void* ptr, bool value);
int C_NSWindow_AnimationBehavior(void* ptr);
void C_NSWindow_SetAnimationBehavior(void* ptr, int value);
bool C_NSWindow_IsDocumentEdited(void* ptr);
void C_NSWindow_SetDocumentEdited(void* ptr, bool value);
double C_NSWindow_BackingScaleFactor(void* ptr);
void* C_NSWindow_Title(void* ptr);
void C_NSWindow_SetTitle(void* ptr, void* value);
void* C_NSWindow_Subtitle(void* ptr);
void C_NSWindow_SetSubtitle(void* ptr, void* value);
int C_NSWindow_TitleVisibility(void* ptr);
void C_NSWindow_SetTitleVisibility(void* ptr, int value);
void* C_NSWindow_RepresentedFilename(void* ptr);
void C_NSWindow_SetRepresentedFilename(void* ptr, void* value);
void* C_NSWindow_RepresentedURL(void* ptr);
void C_NSWindow_SetRepresentedURL(void* ptr, void* value);
void* C_NSWindow_Screen(void* ptr);
void* C_NSWindow_DeepestScreen(void* ptr);
bool C_NSWindow_DisplaysWhenScreenProfileChanges(void* ptr);
void C_NSWindow_SetDisplaysWhenScreenProfileChanges(void* ptr, bool value);
bool C_NSWindow_IsMovableByWindowBackground(void* ptr);
void C_NSWindow_SetMovableByWindowBackground(void* ptr, bool value);
bool C_NSWindow_IsMovable(void* ptr);
void C_NSWindow_SetMovable(void* ptr, bool value);
bool C_NSWindow_IsReleasedWhenClosed(void* ptr);
void C_NSWindow_SetReleasedWhenClosed(void* ptr, bool value);
bool C_NSWindow_IsMiniaturized(void* ptr);
void* C_NSWindow_MiniwindowImage(void* ptr);
void C_NSWindow_SetMiniwindowImage(void* ptr, void* value);
void* C_NSWindow_MiniwindowTitle(void* ptr);
void C_NSWindow_SetMiniwindowTitle(void* ptr, void* value);
void* C_NSWindow_DockTile(void* ptr);
bool C_NSWindow_HasCloseBox(void* ptr);
bool C_NSWindow_HasTitleBar(void* ptr);
int C_NSWindow_OrderedIndex(void* ptr);
void C_NSWindow_SetOrderedIndex(void* ptr, int value);
void* C_NSWindow_AppearanceSource(void* ptr);
void C_NSWindow_SetAppearanceSource(void* ptr, void* value);
bool C_NSWindow_IsFloatingPanel(void* ptr);
bool C_NSWindow_IsMiniaturizable(void* ptr);
bool C_NSWindow_IsModalPanel(void* ptr);
bool C_NSWindow_IsResizable(void* ptr);
unsigned int C_NSWindow_StyleMask(void* ptr);
void C_NSWindow_SetStyleMask(void* ptr, unsigned int value);
int C_NSWindow_WindowTitlebarLayoutDirection(void* ptr);
bool C_NSWindow_IsZoomable(void* ptr);
