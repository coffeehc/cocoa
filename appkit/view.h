#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_View_Alloc();

void* C_NSView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSView_InitWithCoder(void* ptr, void* coder);
void* C_NSView_Init(void* ptr);
void C_NSView_PrepareForReuse(void* ptr);
void C_NSView_AddSubview(void* ptr, void* view);
void C_NSView_AddSubview_Positioned_RelativeTo(void* ptr, void* view, int place, void* otherView);
void C_NSView_DidAddSubview(void* ptr, void* subview);
void C_NSView_RemoveFromSuperview(void* ptr);
void C_NSView_RemoveFromSuperviewWithoutNeedingDisplay(void* ptr);
void C_NSView_ReplaceSubview_With(void* ptr, void* oldView, void* newView);
bool C_NSView_IsDescendantOf(void* ptr, void* view);
void* C_NSView_AncestorSharedWithView(void* ptr, void* view);
void C_NSView_ViewDidMoveToSuperview(void* ptr);
void C_NSView_ViewDidMoveToWindow(void* ptr);
void C_NSView_ViewWillMoveToSuperview(void* ptr, void* newSuperview);
void C_NSView_ViewWillMoveToWindow(void* ptr, void* newWindow);
void C_NSView_WillRemoveSubview(void* ptr, void* subview);
void C_NSView_SetFrameOrigin(void* ptr, CGPoint newOrigin);
void C_NSView_SetFrameSize(void* ptr, CGSize newSize);
void C_NSView_SetBoundsOrigin(void* ptr, CGPoint newOrigin);
void C_NSView_SetBoundsSize(void* ptr, CGSize newSize);
void C_NSView_UpdateLayer(void* ptr);
void C_NSView_DrawRect(void* ptr, CGRect dirtyRect);
bool C_NSView_NeedsToDrawRect(void* ptr, CGRect rect);
void* C_NSView_BitmapImageRepForCachingDisplayInRect(void* ptr, CGRect rect);
void C_NSView_CacheDisplayInRect_ToBitmapImageRep(void* ptr, CGRect rect, void* bitmapImageRep);
void C_NSView_Print(void* ptr, void* sender);
void C_NSView_BeginPageInRect_AtPlacement(void* ptr, CGRect rect, CGPoint location);
Array C_NSView_DataWithEPSInsideRect(void* ptr, CGRect rect);
Array C_NSView_DataWithPDFInsideRect(void* ptr, CGRect rect);
void C_NSView_WriteEPSInsideRect_ToPasteboard(void* ptr, CGRect rect, void* pasteboard);
void C_NSView_WritePDFInsideRect_ToPasteboard(void* ptr, CGRect rect, void* pasteboard);
void C_NSView_DrawPageBorderWithSize(void* ptr, CGSize borderSize);
CGRect C_NSView_RectForPage(void* ptr, int page);
CGPoint C_NSView_LocationOfPrintRect(void* ptr, CGRect rect);
void C_NSView_SetNeedsDisplayInRect(void* ptr, CGRect invalidRect);
void C_NSView_Display(void* ptr);
void C_NSView_DisplayRect(void* ptr, CGRect rect);
void C_NSView_DisplayRectIgnoringOpacity(void* ptr, CGRect rect);
void C_NSView_DisplayRectIgnoringOpacity_InContext(void* ptr, CGRect rect, void* context);
void C_NSView_DisplayIfNeeded(void* ptr);
void C_NSView_DisplayIfNeededInRect(void* ptr, CGRect rect);
void C_NSView_DisplayIfNeededIgnoringOpacity(void* ptr);
void C_NSView_DisplayIfNeededInRectIgnoringOpacity(void* ptr, CGRect rect);
void C_NSView_TranslateRectsNeedingDisplayInRect_By(void* ptr, CGRect clipRect, CGSize delta);
void C_NSView_ViewWillDraw(void* ptr);
CGPoint C_NSView_ConvertPointFromBacking(void* ptr, CGPoint point);
CGPoint C_NSView_ConvertPointToBacking(void* ptr, CGPoint point);
CGPoint C_NSView_ConvertPointFromLayer(void* ptr, CGPoint point);
CGPoint C_NSView_ConvertPointToLayer(void* ptr, CGPoint point);
CGRect C_NSView_ConvertRectFromBacking(void* ptr, CGRect rect);
CGRect C_NSView_ConvertRectToBacking(void* ptr, CGRect rect);
CGRect C_NSView_ConvertRectFromLayer(void* ptr, CGRect rect);
CGRect C_NSView_ConvertRectToLayer(void* ptr, CGRect rect);
CGSize C_NSView_ConvertSizeFromBacking(void* ptr, CGSize size);
CGSize C_NSView_ConvertSizeToBacking(void* ptr, CGSize size);
CGSize C_NSView_ConvertSizeFromLayer(void* ptr, CGSize size);
CGSize C_NSView_ConvertSizeToLayer(void* ptr, CGSize size);
CGPoint C_NSView_ConvertPoint_FromView(void* ptr, CGPoint point, void* view);
CGPoint C_NSView_ConvertPoint_ToView(void* ptr, CGPoint point, void* view);
CGSize C_NSView_ConvertSize_FromView(void* ptr, CGSize size, void* view);
CGSize C_NSView_ConvertSize_ToView(void* ptr, CGSize size, void* view);
CGRect C_NSView_ConvertRect_FromView(void* ptr, CGRect rect, void* view);
CGRect C_NSView_ConvertRect_ToView(void* ptr, CGRect rect, void* view);
CGRect C_NSView_CenterScanRect(void* ptr, CGRect rect);
void C_NSView_TranslateOriginToPoint(void* ptr, CGPoint translation);
void C_NSView_ScaleUnitSquareToSize(void* ptr, CGSize newUnitSize);
void C_NSView_RotateByAngle(void* ptr, double angle);
void C_NSView_ResizeSubviewsWithOldSize(void* ptr, CGSize oldSize);
void C_NSView_ResizeWithOldSuperviewSize(void* ptr, CGSize oldSize);
void C_NSView_AddConstraint(void* ptr, void* constraint);
void C_NSView_AddConstraints(void* ptr, Array constraints);
void C_NSView_RemoveConstraint(void* ptr, void* constraint);
void C_NSView_RemoveConstraints(void* ptr, Array constraints);
void C_NSView_AddLayoutGuide(void* ptr, void* guide);
void C_NSView_RemoveLayoutGuide(void* ptr, void* guide);
void C_NSView_InvalidateIntrinsicContentSize(void* ptr);
float C_NSView_ContentCompressionResistancePriorityForOrientation(void* ptr, int orientation);
void C_NSView_SetContentCompressionResistancePriority_ForOrientation(void* ptr, float priority, int orientation);
float C_NSView_ContentHuggingPriorityForOrientation(void* ptr, int orientation);
void C_NSView_SetContentHuggingPriority_ForOrientation(void* ptr, float priority, int orientation);
CGRect C_NSView_AlignmentRectForFrame(void* ptr, CGRect frame);
CGRect C_NSView_FrameForAlignmentRect(void* ptr, CGRect alignmentRect);
void C_NSView_Layout(void* ptr);
void C_NSView_LayoutSubtreeIfNeeded(void* ptr);
void C_NSView_UpdateConstraints(void* ptr);
void C_NSView_UpdateConstraintsForSubtreeIfNeeded(void* ptr);
Array C_NSView_ConstraintsAffectingLayoutForOrientation(void* ptr, int orientation);
void C_NSView_ExerciseAmbiguityInLayout(void* ptr);
void C_NSView_DrawFocusRingMask(void* ptr);
void C_NSView_NoteFocusRingMaskChanged(void* ptr);
void C_NSView_SetKeyboardFocusRingNeedsDisplayInRect(void* ptr, CGRect rect);
void C_NSView_ViewDidHide(void* ptr);
void C_NSView_ViewDidUnhide(void* ptr);
void C_NSView_ViewWillStartLiveResize(void* ptr);
void C_NSView_ViewDidEndLiveResize(void* ptr);
void C_NSView_AddGestureRecognizer(void* ptr, void* gestureRecognizer);
void C_NSView_RemoveGestureRecognizer(void* ptr, void* gestureRecognizer);
bool C_NSView_AcceptsFirstMouse(void* ptr, void* event);
void* C_NSView_HitTest(void* ptr, CGPoint point);
bool C_NSView_Mouse_InRect(void* ptr, CGPoint point, CGRect rect);
void C_NSView_PrepareContentInRect(void* ptr, CGRect rect);
void C_NSView_ScrollPoint(void* ptr, CGPoint point);
bool C_NSView_ScrollRectToVisible(void* ptr, CGRect rect);
bool C_NSView_Autoscroll(void* ptr, void* event);
CGRect C_NSView_AdjustScroll(void* ptr, CGRect newVisible);
void C_NSView_ScrollClipView_ToPoint(void* ptr, void* clipView, CGPoint point);
void C_NSView_ReflectScrolledClipView(void* ptr, void* clipView);
void C_NSView_RegisterForDraggedTypes(void* ptr, Array newTypes);
void C_NSView_UnregisterDraggedTypes(void* ptr);
void* C_NSView_BeginDraggingSessionWithItems_Event_Source(void* ptr, Array items, void* event, void* source);
bool C_NSView_ShouldDelayWindowOrderingForEvent(void* ptr, void* event);
CGRect C_NSView_RectForSmartMagnificationAtPoint_InRect(void* ptr, CGPoint location, CGRect visibleRect);
void C_NSView_ViewDidChangeBackingProperties(void* ptr);
void* C_NSView_ViewWithTag(void* ptr, int tag);
void C_NSView_RemoveAllToolTips(void* ptr);
void C_NSView_RemoveToolTip(void* ptr, int tag);
void C_NSView_RemoveTrackingRect(void* ptr, int tag);
void C_NSView_AddTrackingArea(void* ptr, void* trackingArea);
void C_NSView_RemoveTrackingArea(void* ptr, void* trackingArea);
void C_NSView_UpdateTrackingAreas(void* ptr);
void C_NSView_AddCursorRect_Cursor(void* ptr, CGRect rect, void* object);
void C_NSView_RemoveCursorRect_Cursor(void* ptr, CGRect rect, void* object);
void C_NSView_DiscardCursorRects(void* ptr);
void C_NSView_ResetCursorRects(void* ptr);
void* C_NSView_MenuForEvent(void* ptr, void* event);
void C_NSView_WillOpenMenu_WithEvent(void* ptr, void* menu, void* event);
void C_NSView_DidCloseMenu_WithEvent(void* ptr, void* menu, void* event);
void C_NSView_BeginDocument(void* ptr);
void C_NSView_EndDocument(void* ptr);
void C_NSView_EndPage(void* ptr);
void C_NSView_ShowDefinitionForAttributedString_AtPoint(void* ptr, void* attrString, CGPoint textBaselineOrigin);
void C_NSView_RulerView_DidAddMarker(void* ptr, void* ruler, void* marker);
void C_NSView_RulerView_DidMoveMarker(void* ptr, void* ruler, void* marker);
void C_NSView_RulerView_DidRemoveMarker(void* ptr, void* ruler, void* marker);
void C_NSView_RulerView_HandleMouseDown(void* ptr, void* ruler, void* event);
double C_NSView_RulerView_LocationForPoint(void* ptr, void* ruler, CGPoint point);
CGPoint C_NSView_RulerView_PointForLocation(void* ptr, void* ruler, double point);
bool C_NSView_RulerView_ShouldAddMarker(void* ptr, void* ruler, void* marker);
bool C_NSView_RulerView_ShouldMoveMarker(void* ptr, void* ruler, void* marker);
bool C_NSView_RulerView_ShouldRemoveMarker(void* ptr, void* ruler, void* marker);
double C_NSView_RulerView_WillAddMarker_AtLocation(void* ptr, void* ruler, void* marker, double location);
double C_NSView_RulerView_WillMoveMarker_ToLocation(void* ptr, void* ruler, void* marker, double location);
void C_NSView_RulerView_WillSetClientView(void* ptr, void* ruler, void* newClient);
void C_NSView_ViewDidChangeEffectiveAppearance(void* ptr);
void* C_NSView_Superview(void* ptr);
Array C_NSView_Subviews(void* ptr);
void C_NSView_SetSubviews(void* ptr, Array value);
void* C_NSView_Window(void* ptr);
void* C_NSView_OpaqueAncestor(void* ptr);
void* C_NSView_EnclosingMenuItem(void* ptr);
CGRect C_NSView_Frame(void* ptr);
void C_NSView_SetFrame(void* ptr, CGRect value);
double C_NSView_FrameRotation(void* ptr);
void C_NSView_SetFrameRotation(void* ptr, double value);
CGRect C_NSView_Bounds(void* ptr);
void C_NSView_SetBounds(void* ptr, CGRect value);
double C_NSView_BoundsRotation(void* ptr);
void C_NSView_SetBoundsRotation(void* ptr, double value);
bool C_NSView_WantsLayer(void* ptr);
void C_NSView_SetWantsLayer(void* ptr, bool value);
bool C_NSView_WantsUpdateLayer(void* ptr);
int C_NSView_LayerContentsPlacement(void* ptr);
void C_NSView_SetLayerContentsPlacement(void* ptr, int value);
int C_NSView_LayerContentsRedrawPolicy(void* ptr);
void C_NSView_SetLayerContentsRedrawPolicy(void* ptr, int value);
bool C_NSView_CanDrawSubviewsIntoLayer(void* ptr);
void C_NSView_SetCanDrawSubviewsIntoLayer(void* ptr, bool value);
bool C_NSView_LayerUsesCoreImageFilters(void* ptr);
void C_NSView_SetLayerUsesCoreImageFilters(void* ptr, bool value);
double C_NSView_AlphaValue(void* ptr);
void C_NSView_SetAlphaValue(void* ptr, double value);
double C_NSView_FrameCenterRotation(void* ptr);
void C_NSView_SetFrameCenterRotation(void* ptr, double value);
void* C_NSView_Shadow(void* ptr);
void C_NSView_SetShadow(void* ptr, void* value);
bool C_NSView_CanDrawConcurrently(void* ptr);
void C_NSView_SetCanDrawConcurrently(void* ptr, bool value);
CGRect C_NSView_VisibleRect(void* ptr);
bool C_NSView_WantsDefaultClipping(void* ptr);
void* C_NSView_PrintJobTitle(void* ptr);
void* C_NSView_PageHeader(void* ptr);
void* C_NSView_PageFooter(void* ptr);
double C_NSView_HeightAdjustLimit(void* ptr);
double C_NSView_WidthAdjustLimit(void* ptr);
bool C_NSView_NeedsDisplay(void* ptr);
void C_NSView_SetNeedsDisplay(void* ptr, bool value);
bool C_NSView_IsOpaque(void* ptr);
bool C_NSView_IsFlipped(void* ptr);
bool C_NSView_IsRotatedFromBase(void* ptr);
bool C_NSView_IsRotatedOrScaledFromBase(void* ptr);
bool C_NSView_AutoresizesSubviews(void* ptr);
void C_NSView_SetAutoresizesSubviews(void* ptr, bool value);
unsigned int C_NSView_AutoresizingMask(void* ptr);
void C_NSView_SetAutoresizingMask(void* ptr, unsigned int value);
void* C_NSView_BottomAnchor(void* ptr);
void* C_NSView_CenterXAnchor(void* ptr);
void* C_NSView_CenterYAnchor(void* ptr);
void* C_NSView_FirstBaselineAnchor(void* ptr);
void* C_NSView_HeightAnchor(void* ptr);
void* C_NSView_LastBaselineAnchor(void* ptr);
void* C_NSView_LeadingAnchor(void* ptr);
void* C_NSView_LeftAnchor(void* ptr);
void* C_NSView_RightAnchor(void* ptr);
void* C_NSView_TopAnchor(void* ptr);
void* C_NSView_TrailingAnchor(void* ptr);
void* C_NSView_WidthAnchor(void* ptr);
Array C_NSView_Constraints(void* ptr);
Array C_NSView_LayoutGuides(void* ptr);
void* C_NSView_LayoutMarginsGuide(void* ptr);
void* C_NSView_SafeAreaLayoutGuide(void* ptr);
CGSize C_NSView_FittingSize(void* ptr);
CGSize C_NSView_IntrinsicContentSize(void* ptr);
NSEdgeInsets C_NSView_AlignmentRectInsets(void* ptr);
double C_NSView_BaselineOffsetFromBottom(void* ptr);
double C_NSView_FirstBaselineOffsetFromTop(void* ptr);
double C_NSView_LastBaselineOffsetFromBottom(void* ptr);
bool C_NSView_NeedsLayout(void* ptr);
void C_NSView_SetNeedsLayout(void* ptr, bool value);
bool C_NSView_NeedsUpdateConstraints(void* ptr);
void C_NSView_SetNeedsUpdateConstraints(void* ptr, bool value);
bool C_NSView_View_RequiresConstraintBasedLayout();
bool C_NSView_TranslatesAutoresizingMaskIntoConstraints(void* ptr);
void C_NSView_SetTranslatesAutoresizingMaskIntoConstraints(void* ptr, bool value);
bool C_NSView_HasAmbiguousLayout(void* ptr);
void* C_NSView_FocusView();
unsigned int C_NSView_FocusRingType(void* ptr);
void C_NSView_SetFocusRingType(void* ptr, unsigned int value);
CGRect C_NSView_FocusRingMaskBounds(void* ptr);
unsigned int C_NSView_View_DefaultFocusRingType();
bool C_NSView_AllowsVibrancy(void* ptr);
bool C_NSView_IsInFullScreenMode(void* ptr);
bool C_NSView_IsHidden(void* ptr);
void C_NSView_SetHidden(void* ptr, bool value);
bool C_NSView_IsHiddenOrHasHiddenAncestor(void* ptr);
bool C_NSView_InLiveResize(void* ptr);
bool C_NSView_PreservesContentDuringLiveResize(void* ptr);
CGRect C_NSView_RectPreservedDuringLiveResize(void* ptr);
Array C_NSView_GestureRecognizers(void* ptr);
void C_NSView_SetGestureRecognizers(void* ptr, Array value);
bool C_NSView_MouseDownCanMoveWindow(void* ptr);
void* C_NSView_InputContext(void* ptr);
bool C_NSView_WantsRestingTouches(void* ptr);
void C_NSView_SetWantsRestingTouches(void* ptr, bool value);
bool C_NSView_CanBecomeKeyView(void* ptr);
bool C_NSView_NeedsPanelToBecomeKey(void* ptr);
void* C_NSView_NextKeyView(void* ptr);
void C_NSView_SetNextKeyView(void* ptr, void* value);
void* C_NSView_NextValidKeyView(void* ptr);
void* C_NSView_PreviousKeyView(void* ptr);
void* C_NSView_PreviousValidKeyView(void* ptr);
CGRect C_NSView_PreparedContentRect(void* ptr);
void C_NSView_SetPreparedContentRect(void* ptr, CGRect value);
void* C_NSView_EnclosingScrollView(void* ptr);
Array C_NSView_RegisteredDraggedTypes(void* ptr);
bool C_NSView_PostsFrameChangedNotifications(void* ptr);
void C_NSView_SetPostsFrameChangedNotifications(void* ptr, bool value);
bool C_NSView_PostsBoundsChangedNotifications(void* ptr);
void C_NSView_SetPostsBoundsChangedNotifications(void* ptr, bool value);
int C_NSView_Tag(void* ptr);
void* C_NSView_ToolTip(void* ptr);
void C_NSView_SetToolTip(void* ptr, void* value);
Array C_NSView_TrackingAreas(void* ptr);
void* C_NSView_View_DefaultMenu();
bool C_NSView_IsDrawingFindIndicator(void* ptr);
int C_NSView_UserInterfaceLayoutDirection(void* ptr);
void C_NSView_SetUserInterfaceLayoutDirection(void* ptr, int value);
void* C_NSView_PressureConfiguration(void* ptr);
void C_NSView_SetPressureConfiguration(void* ptr, void* value);
NSEdgeInsets C_NSView_AdditionalSafeAreaInsets(void* ptr);
void C_NSView_SetAdditionalSafeAreaInsets(void* ptr, NSEdgeInsets value);
unsigned int C_NSView_AllowedTouchTypes(void* ptr);
void C_NSView_SetAllowedTouchTypes(void* ptr, unsigned int value);
void* C_NSView_CandidateListTouchBarItem(void* ptr);
bool C_NSView_IsHorizontalContentSizeConstraintActive(void* ptr);
void C_NSView_SetHorizontalContentSizeConstraintActive(void* ptr, bool value);
bool C_NSView_IsVerticalContentSizeConstraintActive(void* ptr);
void C_NSView_SetVerticalContentSizeConstraintActive(void* ptr, bool value);
NSEdgeInsets C_NSView_SafeAreaInsets(void* ptr);
CGRect C_NSView_SafeAreaRect(void* ptr);
bool C_NSView_View_CompatibleWithResponsiveScrolling();
