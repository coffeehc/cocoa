#import <Appkit/Appkit.h>
#import "view.h"

void* C_View_Alloc() {
	return [NSView alloc];
}

void* C_NSView_InitWithFrame(void* ptr, CGRect frameRect) {
	NSView* nSView = (NSView*)ptr;
	NSView* result = [nSView initWithFrame:frameRect];
	return result;
}

void* C_NSView_InitWithCoder(void* ptr, void* coder) {
	NSView* nSView = (NSView*)ptr;
	NSView* result = [nSView initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSView_Init(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSView* result = [nSView init];
	return result;
}

void C_NSView_PrepareForReuse(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView prepareForReuse];
}

void C_NSView_AddSubview(void* ptr, void* view) {
	NSView* nSView = (NSView*)ptr;
	[nSView addSubview:(NSView*)view];
}

void C_NSView_DidAddSubview(void* ptr, void* subview) {
	NSView* nSView = (NSView*)ptr;
	[nSView didAddSubview:(NSView*)subview];
}

void C_NSView_RemoveFromSuperview(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView removeFromSuperview];
}

void C_NSView_RemoveFromSuperviewWithoutNeedingDisplay(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView removeFromSuperviewWithoutNeedingDisplay];
}

void C_NSView_ReplaceSubview_With(void* ptr, void* oldView, void* newView) {
	NSView* nSView = (NSView*)ptr;
	[nSView replaceSubview:(NSView*)oldView with:(NSView*)newView];
}

bool C_NSView_IsDescendantOf(void* ptr, void* view) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView isDescendantOf:(NSView*)view];
	return result;
}

void* C_NSView_AncestorSharedWithView(void* ptr, void* view) {
	NSView* nSView = (NSView*)ptr;
	NSView* result = [nSView ancestorSharedWithView:(NSView*)view];
	return result;
}

void C_NSView_ViewDidMoveToSuperview(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView viewDidMoveToSuperview];
}

void C_NSView_ViewDidMoveToWindow(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView viewDidMoveToWindow];
}

void C_NSView_ViewWillMoveToSuperview(void* ptr, void* newSuperview) {
	NSView* nSView = (NSView*)ptr;
	[nSView viewWillMoveToSuperview:(NSView*)newSuperview];
}

void C_NSView_ViewWillMoveToWindow(void* ptr, void* newWindow) {
	NSView* nSView = (NSView*)ptr;
	[nSView viewWillMoveToWindow:(NSWindow*)newWindow];
}

void C_NSView_WillRemoveSubview(void* ptr, void* subview) {
	NSView* nSView = (NSView*)ptr;
	[nSView willRemoveSubview:(NSView*)subview];
}

void C_NSView_SetFrameOrigin(void* ptr, CGPoint newOrigin) {
	NSView* nSView = (NSView*)ptr;
	[nSView setFrameOrigin:newOrigin];
}

void C_NSView_SetFrameSize(void* ptr, CGSize newSize) {
	NSView* nSView = (NSView*)ptr;
	[nSView setFrameSize:newSize];
}

void C_NSView_SetBoundsOrigin(void* ptr, CGPoint newOrigin) {
	NSView* nSView = (NSView*)ptr;
	[nSView setBoundsOrigin:newOrigin];
}

void C_NSView_SetBoundsSize(void* ptr, CGSize newSize) {
	NSView* nSView = (NSView*)ptr;
	[nSView setBoundsSize:newSize];
}

void C_NSView_UpdateLayer(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView updateLayer];
}

void C_NSView_DrawRect(void* ptr, CGRect dirtyRect) {
	NSView* nSView = (NSView*)ptr;
	[nSView drawRect:dirtyRect];
}

bool C_NSView_NeedsToDrawRect(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView needsToDrawRect:rect];
	return result;
}

void* C_NSView_BitmapImageRepForCachingDisplayInRect(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	NSBitmapImageRep* result = [nSView bitmapImageRepForCachingDisplayInRect:rect];
	return result;
}

void C_NSView_CacheDisplayInRect_ToBitmapImageRep(void* ptr, CGRect rect, void* bitmapImageRep) {
	NSView* nSView = (NSView*)ptr;
	[nSView cacheDisplayInRect:rect toBitmapImageRep:(NSBitmapImageRep*)bitmapImageRep];
}

void C_NSView_Print(void* ptr, void* sender) {
	NSView* nSView = (NSView*)ptr;
	[nSView print:(id)sender];
}

void C_NSView_BeginPageInRect_AtPlacement(void* ptr, CGRect rect, CGPoint location) {
	NSView* nSView = (NSView*)ptr;
	[nSView beginPageInRect:rect atPlacement:location];
}

Array C_NSView_DataWithEPSInsideRect(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	NSData* result = [nSView dataWithEPSInsideRect:rect];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}

Array C_NSView_DataWithPDFInsideRect(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	NSData* result = [nSView dataWithPDFInsideRect:rect];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}

void C_NSView_WriteEPSInsideRect_ToPasteboard(void* ptr, CGRect rect, void* pasteboard) {
	NSView* nSView = (NSView*)ptr;
	[nSView writeEPSInsideRect:rect toPasteboard:(NSPasteboard*)pasteboard];
}

void C_NSView_WritePDFInsideRect_ToPasteboard(void* ptr, CGRect rect, void* pasteboard) {
	NSView* nSView = (NSView*)ptr;
	[nSView writePDFInsideRect:rect toPasteboard:(NSPasteboard*)pasteboard];
}

void C_NSView_DrawPageBorderWithSize(void* ptr, CGSize borderSize) {
	NSView* nSView = (NSView*)ptr;
	[nSView drawPageBorderWithSize:borderSize];
}

CGRect C_NSView_RectForPage(void* ptr, int page) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView rectForPage:page];
	return result;
}

CGPoint C_NSView_LocationOfPrintRect(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	CGPoint result = [nSView locationOfPrintRect:rect];
	return result;
}

void C_NSView_SetNeedsDisplayInRect(void* ptr, CGRect invalidRect) {
	NSView* nSView = (NSView*)ptr;
	[nSView setNeedsDisplayInRect:invalidRect];
}

void C_NSView_Display(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView display];
}

void C_NSView_DisplayRect(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	[nSView displayRect:rect];
}

void C_NSView_DisplayRectIgnoringOpacity(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	[nSView displayRectIgnoringOpacity:rect];
}

void C_NSView_DisplayRectIgnoringOpacity_InContext(void* ptr, CGRect rect, void* context) {
	NSView* nSView = (NSView*)ptr;
	[nSView displayRectIgnoringOpacity:rect inContext:(NSGraphicsContext*)context];
}

void C_NSView_DisplayIfNeeded(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView displayIfNeeded];
}

void C_NSView_DisplayIfNeededInRect(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	[nSView displayIfNeededInRect:rect];
}

void C_NSView_DisplayIfNeededIgnoringOpacity(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView displayIfNeededIgnoringOpacity];
}

void C_NSView_DisplayIfNeededInRectIgnoringOpacity(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	[nSView displayIfNeededInRectIgnoringOpacity:rect];
}

void C_NSView_TranslateRectsNeedingDisplayInRect_By(void* ptr, CGRect clipRect, CGSize delta) {
	NSView* nSView = (NSView*)ptr;
	[nSView translateRectsNeedingDisplayInRect:clipRect by:delta];
}

void C_NSView_ViewWillDraw(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView viewWillDraw];
}

CGPoint C_NSView_ConvertPointFromBacking(void* ptr, CGPoint point) {
	NSView* nSView = (NSView*)ptr;
	CGPoint result = [nSView convertPointFromBacking:point];
	return result;
}

CGPoint C_NSView_ConvertPointToBacking(void* ptr, CGPoint point) {
	NSView* nSView = (NSView*)ptr;
	CGPoint result = [nSView convertPointToBacking:point];
	return result;
}

CGPoint C_NSView_ConvertPointFromLayer(void* ptr, CGPoint point) {
	NSView* nSView = (NSView*)ptr;
	CGPoint result = [nSView convertPointFromLayer:point];
	return result;
}

CGPoint C_NSView_ConvertPointToLayer(void* ptr, CGPoint point) {
	NSView* nSView = (NSView*)ptr;
	CGPoint result = [nSView convertPointToLayer:point];
	return result;
}

CGRect C_NSView_ConvertRectFromBacking(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView convertRectFromBacking:rect];
	return result;
}

CGRect C_NSView_ConvertRectToBacking(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView convertRectToBacking:rect];
	return result;
}

CGRect C_NSView_ConvertRectFromLayer(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView convertRectFromLayer:rect];
	return result;
}

CGRect C_NSView_ConvertRectToLayer(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView convertRectToLayer:rect];
	return result;
}

CGSize C_NSView_ConvertSizeFromBacking(void* ptr, CGSize size) {
	NSView* nSView = (NSView*)ptr;
	CGSize result = [nSView convertSizeFromBacking:size];
	return result;
}

CGSize C_NSView_ConvertSizeToBacking(void* ptr, CGSize size) {
	NSView* nSView = (NSView*)ptr;
	CGSize result = [nSView convertSizeToBacking:size];
	return result;
}

CGSize C_NSView_ConvertSizeFromLayer(void* ptr, CGSize size) {
	NSView* nSView = (NSView*)ptr;
	CGSize result = [nSView convertSizeFromLayer:size];
	return result;
}

CGSize C_NSView_ConvertSizeToLayer(void* ptr, CGSize size) {
	NSView* nSView = (NSView*)ptr;
	CGSize result = [nSView convertSizeToLayer:size];
	return result;
}

CGPoint C_NSView_ConvertPoint_FromView(void* ptr, CGPoint point, void* view) {
	NSView* nSView = (NSView*)ptr;
	CGPoint result = [nSView convertPoint:point fromView:(NSView*)view];
	return result;
}

CGPoint C_NSView_ConvertPoint_ToView(void* ptr, CGPoint point, void* view) {
	NSView* nSView = (NSView*)ptr;
	CGPoint result = [nSView convertPoint:point toView:(NSView*)view];
	return result;
}

CGSize C_NSView_ConvertSize_FromView(void* ptr, CGSize size, void* view) {
	NSView* nSView = (NSView*)ptr;
	CGSize result = [nSView convertSize:size fromView:(NSView*)view];
	return result;
}

CGSize C_NSView_ConvertSize_ToView(void* ptr, CGSize size, void* view) {
	NSView* nSView = (NSView*)ptr;
	CGSize result = [nSView convertSize:size toView:(NSView*)view];
	return result;
}

CGRect C_NSView_ConvertRect_FromView(void* ptr, CGRect rect, void* view) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView convertRect:rect fromView:(NSView*)view];
	return result;
}

CGRect C_NSView_ConvertRect_ToView(void* ptr, CGRect rect, void* view) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView convertRect:rect toView:(NSView*)view];
	return result;
}

CGRect C_NSView_CenterScanRect(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView centerScanRect:rect];
	return result;
}

void C_NSView_TranslateOriginToPoint(void* ptr, CGPoint translation) {
	NSView* nSView = (NSView*)ptr;
	[nSView translateOriginToPoint:translation];
}

void C_NSView_ScaleUnitSquareToSize(void* ptr, CGSize newUnitSize) {
	NSView* nSView = (NSView*)ptr;
	[nSView scaleUnitSquareToSize:newUnitSize];
}

void C_NSView_RotateByAngle(void* ptr, double angle) {
	NSView* nSView = (NSView*)ptr;
	[nSView rotateByAngle:angle];
}

void C_NSView_ResizeSubviewsWithOldSize(void* ptr, CGSize oldSize) {
	NSView* nSView = (NSView*)ptr;
	[nSView resizeSubviewsWithOldSize:oldSize];
}

void C_NSView_ResizeWithOldSuperviewSize(void* ptr, CGSize oldSize) {
	NSView* nSView = (NSView*)ptr;
	[nSView resizeWithOldSuperviewSize:oldSize];
}

void C_NSView_AddConstraint(void* ptr, void* constraint) {
	NSView* nSView = (NSView*)ptr;
	[nSView addConstraint:(NSLayoutConstraint*)constraint];
}

void C_NSView_RemoveConstraint(void* ptr, void* constraint) {
	NSView* nSView = (NSView*)ptr;
	[nSView removeConstraint:(NSLayoutConstraint*)constraint];
}

void C_NSView_AddLayoutGuide(void* ptr, void* guide) {
	NSView* nSView = (NSView*)ptr;
	[nSView addLayoutGuide:(NSLayoutGuide*)guide];
}

void C_NSView_RemoveLayoutGuide(void* ptr, void* guide) {
	NSView* nSView = (NSView*)ptr;
	[nSView removeLayoutGuide:(NSLayoutGuide*)guide];
}

void C_NSView_InvalidateIntrinsicContentSize(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView invalidateIntrinsicContentSize];
}

CGRect C_NSView_AlignmentRectForFrame(void* ptr, CGRect frame) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView alignmentRectForFrame:frame];
	return result;
}

CGRect C_NSView_FrameForAlignmentRect(void* ptr, CGRect alignmentRect) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView frameForAlignmentRect:alignmentRect];
	return result;
}

void C_NSView_Layout(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView layout];
}

void C_NSView_LayoutSubtreeIfNeeded(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView layoutSubtreeIfNeeded];
}

void C_NSView_UpdateConstraints(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView updateConstraints];
}

void C_NSView_UpdateConstraintsForSubtreeIfNeeded(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView updateConstraintsForSubtreeIfNeeded];
}

void C_NSView_ExerciseAmbiguityInLayout(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView exerciseAmbiguityInLayout];
}

void C_NSView_DrawFocusRingMask(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView drawFocusRingMask];
}

void C_NSView_NoteFocusRingMaskChanged(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView noteFocusRingMaskChanged];
}

void C_NSView_SetKeyboardFocusRingNeedsDisplayInRect(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	[nSView setKeyboardFocusRingNeedsDisplayInRect:rect];
}

void C_NSView_ViewDidHide(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView viewDidHide];
}

void C_NSView_ViewDidUnhide(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView viewDidUnhide];
}

void C_NSView_ViewWillStartLiveResize(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView viewWillStartLiveResize];
}

void C_NSView_ViewDidEndLiveResize(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView viewDidEndLiveResize];
}

void C_NSView_AddGestureRecognizer(void* ptr, void* gestureRecognizer) {
	NSView* nSView = (NSView*)ptr;
	[nSView addGestureRecognizer:(NSGestureRecognizer*)gestureRecognizer];
}

void C_NSView_RemoveGestureRecognizer(void* ptr, void* gestureRecognizer) {
	NSView* nSView = (NSView*)ptr;
	[nSView removeGestureRecognizer:(NSGestureRecognizer*)gestureRecognizer];
}

bool C_NSView_AcceptsFirstMouse(void* ptr, void* event) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView acceptsFirstMouse:(NSEvent*)event];
	return result;
}

void* C_NSView_HitTest(void* ptr, CGPoint point) {
	NSView* nSView = (NSView*)ptr;
	NSView* result = [nSView hitTest:point];
	return result;
}

bool C_NSView_Mouse_InRect(void* ptr, CGPoint point, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView mouse:point inRect:rect];
	return result;
}

void C_NSView_PrepareContentInRect(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	[nSView prepareContentInRect:rect];
}

void C_NSView_ScrollPoint(void* ptr, CGPoint point) {
	NSView* nSView = (NSView*)ptr;
	[nSView scrollPoint:point];
}

bool C_NSView_ScrollRectToVisible(void* ptr, CGRect rect) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView scrollRectToVisible:rect];
	return result;
}

bool C_NSView_Autoscroll(void* ptr, void* event) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView autoscroll:(NSEvent*)event];
	return result;
}

CGRect C_NSView_AdjustScroll(void* ptr, CGRect newVisible) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView adjustScroll:newVisible];
	return result;
}

void C_NSView_ScrollClipView_ToPoint(void* ptr, void* clipView, CGPoint point) {
	NSView* nSView = (NSView*)ptr;
	[nSView scrollClipView:(NSClipView*)clipView toPoint:point];
}

void C_NSView_ReflectScrolledClipView(void* ptr, void* clipView) {
	NSView* nSView = (NSView*)ptr;
	[nSView reflectScrolledClipView:(NSClipView*)clipView];
}

void C_NSView_UnregisterDraggedTypes(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView unregisterDraggedTypes];
}

bool C_NSView_ShouldDelayWindowOrderingForEvent(void* ptr, void* event) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView shouldDelayWindowOrderingForEvent:(NSEvent*)event];
	return result;
}

CGRect C_NSView_RectForSmartMagnificationAtPoint_InRect(void* ptr, CGPoint location, CGRect visibleRect) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView rectForSmartMagnificationAtPoint:location inRect:visibleRect];
	return result;
}

void C_NSView_ViewDidChangeBackingProperties(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView viewDidChangeBackingProperties];
}

void C_NSView_RemoveAllToolTips(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView removeAllToolTips];
}

void C_NSView_RemoveToolTip(void* ptr, int tag) {
	NSView* nSView = (NSView*)ptr;
	[nSView removeToolTip:tag];
}

void C_NSView_RemoveTrackingRect(void* ptr, int tag) {
	NSView* nSView = (NSView*)ptr;
	[nSView removeTrackingRect:tag];
}

void C_NSView_AddTrackingArea(void* ptr, void* trackingArea) {
	NSView* nSView = (NSView*)ptr;
	[nSView addTrackingArea:(NSTrackingArea*)trackingArea];
}

void C_NSView_RemoveTrackingArea(void* ptr, void* trackingArea) {
	NSView* nSView = (NSView*)ptr;
	[nSView removeTrackingArea:(NSTrackingArea*)trackingArea];
}

void C_NSView_UpdateTrackingAreas(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView updateTrackingAreas];
}

void C_NSView_AddCursorRect_Cursor(void* ptr, CGRect rect, void* object) {
	NSView* nSView = (NSView*)ptr;
	[nSView addCursorRect:rect cursor:(NSCursor*)object];
}

void C_NSView_RemoveCursorRect_Cursor(void* ptr, CGRect rect, void* object) {
	NSView* nSView = (NSView*)ptr;
	[nSView removeCursorRect:rect cursor:(NSCursor*)object];
}

void C_NSView_DiscardCursorRects(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView discardCursorRects];
}

void C_NSView_ResetCursorRects(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView resetCursorRects];
}

void* C_NSView_MenuForEvent(void* ptr, void* event) {
	NSView* nSView = (NSView*)ptr;
	NSMenu* result = [nSView menuForEvent:(NSEvent*)event];
	return result;
}

void C_NSView_WillOpenMenu_WithEvent(void* ptr, void* menu, void* event) {
	NSView* nSView = (NSView*)ptr;
	[nSView willOpenMenu:(NSMenu*)menu withEvent:(NSEvent*)event];
}

void C_NSView_DidCloseMenu_WithEvent(void* ptr, void* menu, void* event) {
	NSView* nSView = (NSView*)ptr;
	[nSView didCloseMenu:(NSMenu*)menu withEvent:(NSEvent*)event];
}

void C_NSView_BeginDocument(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView beginDocument];
}

void C_NSView_EndDocument(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView endDocument];
}

void C_NSView_EndPage(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView endPage];
}

void C_NSView_ShowDefinitionForAttributedString_AtPoint(void* ptr, void* attrString, CGPoint textBaselineOrigin) {
	NSView* nSView = (NSView*)ptr;
	[nSView showDefinitionForAttributedString:(NSAttributedString*)attrString atPoint:textBaselineOrigin];
}

void C_NSView_RulerView_DidAddMarker(void* ptr, void* ruler, void* marker) {
	NSView* nSView = (NSView*)ptr;
	[nSView rulerView:(NSRulerView*)ruler didAddMarker:(NSRulerMarker*)marker];
}

void C_NSView_RulerView_DidMoveMarker(void* ptr, void* ruler, void* marker) {
	NSView* nSView = (NSView*)ptr;
	[nSView rulerView:(NSRulerView*)ruler didMoveMarker:(NSRulerMarker*)marker];
}

void C_NSView_RulerView_DidRemoveMarker(void* ptr, void* ruler, void* marker) {
	NSView* nSView = (NSView*)ptr;
	[nSView rulerView:(NSRulerView*)ruler didRemoveMarker:(NSRulerMarker*)marker];
}

void C_NSView_RulerView_HandleMouseDown(void* ptr, void* ruler, void* event) {
	NSView* nSView = (NSView*)ptr;
	[nSView rulerView:(NSRulerView*)ruler handleMouseDown:(NSEvent*)event];
}

double C_NSView_RulerView_LocationForPoint(void* ptr, void* ruler, CGPoint point) {
	NSView* nSView = (NSView*)ptr;
	double result = [nSView rulerView:(NSRulerView*)ruler locationForPoint:point];
	return result;
}

CGPoint C_NSView_RulerView_PointForLocation(void* ptr, void* ruler, double point) {
	NSView* nSView = (NSView*)ptr;
	CGPoint result = [nSView rulerView:(NSRulerView*)ruler pointForLocation:point];
	return result;
}

bool C_NSView_RulerView_ShouldAddMarker(void* ptr, void* ruler, void* marker) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView rulerView:(NSRulerView*)ruler shouldAddMarker:(NSRulerMarker*)marker];
	return result;
}

bool C_NSView_RulerView_ShouldMoveMarker(void* ptr, void* ruler, void* marker) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView rulerView:(NSRulerView*)ruler shouldMoveMarker:(NSRulerMarker*)marker];
	return result;
}

bool C_NSView_RulerView_ShouldRemoveMarker(void* ptr, void* ruler, void* marker) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView rulerView:(NSRulerView*)ruler shouldRemoveMarker:(NSRulerMarker*)marker];
	return result;
}

double C_NSView_RulerView_WillAddMarker_AtLocation(void* ptr, void* ruler, void* marker, double location) {
	NSView* nSView = (NSView*)ptr;
	double result = [nSView rulerView:(NSRulerView*)ruler willAddMarker:(NSRulerMarker*)marker atLocation:location];
	return result;
}

double C_NSView_RulerView_WillMoveMarker_ToLocation(void* ptr, void* ruler, void* marker, double location) {
	NSView* nSView = (NSView*)ptr;
	double result = [nSView rulerView:(NSRulerView*)ruler willMoveMarker:(NSRulerMarker*)marker toLocation:location];
	return result;
}

void C_NSView_RulerView_WillSetClientView(void* ptr, void* ruler, void* newClient) {
	NSView* nSView = (NSView*)ptr;
	[nSView rulerView:(NSRulerView*)ruler willSetClientView:(NSView*)newClient];
}

void C_NSView_ViewDidChangeEffectiveAppearance(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	[nSView viewDidChangeEffectiveAppearance];
}

void* C_NSView_Superview(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSView* result = [nSView superview];
	return result;
}

void* C_NSView_Window(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSWindow* result = [nSView window];
	return result;
}

void* C_NSView_OpaqueAncestor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSView* result = [nSView opaqueAncestor];
	return result;
}

void* C_NSView_EnclosingMenuItem(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSMenuItem* result = [nSView enclosingMenuItem];
	return result;
}

CGRect C_NSView_Frame(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView frame];
	return result;
}

void C_NSView_SetFrame(void* ptr, CGRect value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setFrame:value];
}

double C_NSView_FrameRotation(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	double result = [nSView frameRotation];
	return result;
}

void C_NSView_SetFrameRotation(void* ptr, double value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setFrameRotation:value];
}

CGRect C_NSView_Bounds(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView bounds];
	return result;
}

void C_NSView_SetBounds(void* ptr, CGRect value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setBounds:value];
}

double C_NSView_BoundsRotation(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	double result = [nSView boundsRotation];
	return result;
}

void C_NSView_SetBoundsRotation(void* ptr, double value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setBoundsRotation:value];
}

bool C_NSView_WantsLayer(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView wantsLayer];
	return result;
}

void C_NSView_SetWantsLayer(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setWantsLayer:value];
}

bool C_NSView_WantsUpdateLayer(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView wantsUpdateLayer];
	return result;
}

bool C_NSView_CanDrawSubviewsIntoLayer(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView canDrawSubviewsIntoLayer];
	return result;
}

void C_NSView_SetCanDrawSubviewsIntoLayer(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setCanDrawSubviewsIntoLayer:value];
}

bool C_NSView_LayerUsesCoreImageFilters(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView layerUsesCoreImageFilters];
	return result;
}

void C_NSView_SetLayerUsesCoreImageFilters(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setLayerUsesCoreImageFilters:value];
}

double C_NSView_AlphaValue(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	double result = [nSView alphaValue];
	return result;
}

void C_NSView_SetAlphaValue(void* ptr, double value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setAlphaValue:value];
}

double C_NSView_FrameCenterRotation(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	double result = [nSView frameCenterRotation];
	return result;
}

void C_NSView_SetFrameCenterRotation(void* ptr, double value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setFrameCenterRotation:value];
}

void* C_NSView_Shadow(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSShadow* result = [nSView shadow];
	return result;
}

void C_NSView_SetShadow(void* ptr, void* value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setShadow:(NSShadow*)value];
}

bool C_NSView_CanDrawConcurrently(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView canDrawConcurrently];
	return result;
}

void C_NSView_SetCanDrawConcurrently(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setCanDrawConcurrently:value];
}

CGRect C_NSView_VisibleRect(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView visibleRect];
	return result;
}

bool C_NSView_WantsDefaultClipping(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView wantsDefaultClipping];
	return result;
}

void* C_NSView_PrintJobTitle(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSString* result = [nSView printJobTitle];
	return result;
}

void* C_NSView_PageHeader(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSAttributedString* result = [nSView pageHeader];
	return result;
}

void* C_NSView_PageFooter(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSAttributedString* result = [nSView pageFooter];
	return result;
}

double C_NSView_HeightAdjustLimit(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	double result = [nSView heightAdjustLimit];
	return result;
}

double C_NSView_WidthAdjustLimit(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	double result = [nSView widthAdjustLimit];
	return result;
}

bool C_NSView_NeedsDisplay(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView needsDisplay];
	return result;
}

void C_NSView_SetNeedsDisplay(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setNeedsDisplay:value];
}

bool C_NSView_IsOpaque(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView isOpaque];
	return result;
}

bool C_NSView_IsFlipped(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView isFlipped];
	return result;
}

bool C_NSView_IsRotatedFromBase(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView isRotatedFromBase];
	return result;
}

bool C_NSView_IsRotatedOrScaledFromBase(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView isRotatedOrScaledFromBase];
	return result;
}

bool C_NSView_AutoresizesSubviews(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView autoresizesSubviews];
	return result;
}

void C_NSView_SetAutoresizesSubviews(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setAutoresizesSubviews:value];
}

void* C_NSView_BottomAnchor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutYAxisAnchor* result = [nSView bottomAnchor];
	return result;
}

void* C_NSView_CenterXAnchor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutXAxisAnchor* result = [nSView centerXAnchor];
	return result;
}

void* C_NSView_CenterYAnchor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutYAxisAnchor* result = [nSView centerYAnchor];
	return result;
}

void* C_NSView_FirstBaselineAnchor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutYAxisAnchor* result = [nSView firstBaselineAnchor];
	return result;
}

void* C_NSView_HeightAnchor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutDimension* result = [nSView heightAnchor];
	return result;
}

void* C_NSView_LastBaselineAnchor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutYAxisAnchor* result = [nSView lastBaselineAnchor];
	return result;
}

void* C_NSView_LeadingAnchor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutXAxisAnchor* result = [nSView leadingAnchor];
	return result;
}

void* C_NSView_LeftAnchor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutXAxisAnchor* result = [nSView leftAnchor];
	return result;
}

void* C_NSView_RightAnchor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutXAxisAnchor* result = [nSView rightAnchor];
	return result;
}

void* C_NSView_TopAnchor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutYAxisAnchor* result = [nSView topAnchor];
	return result;
}

void* C_NSView_TrailingAnchor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutXAxisAnchor* result = [nSView trailingAnchor];
	return result;
}

void* C_NSView_WidthAnchor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutDimension* result = [nSView widthAnchor];
	return result;
}

void* C_NSView_LayoutMarginsGuide(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutGuide* result = [nSView layoutMarginsGuide];
	return result;
}

void* C_NSView_SafeAreaLayoutGuide(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSLayoutGuide* result = [nSView safeAreaLayoutGuide];
	return result;
}

CGSize C_NSView_FittingSize(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	CGSize result = [nSView fittingSize];
	return result;
}

CGSize C_NSView_IntrinsicContentSize(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	CGSize result = [nSView intrinsicContentSize];
	return result;
}

NSEdgeInsets C_NSView_AlignmentRectInsets(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSEdgeInsets result = [nSView alignmentRectInsets];
	return result;
}

double C_NSView_BaselineOffsetFromBottom(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	double result = [nSView baselineOffsetFromBottom];
	return result;
}

double C_NSView_FirstBaselineOffsetFromTop(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	double result = [nSView firstBaselineOffsetFromTop];
	return result;
}

double C_NSView_LastBaselineOffsetFromBottom(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	double result = [nSView lastBaselineOffsetFromBottom];
	return result;
}

bool C_NSView_NeedsLayout(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView needsLayout];
	return result;
}

void C_NSView_SetNeedsLayout(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setNeedsLayout:value];
}

bool C_NSView_NeedsUpdateConstraints(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView needsUpdateConstraints];
	return result;
}

void C_NSView_SetNeedsUpdateConstraints(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setNeedsUpdateConstraints:value];
}

bool C_NSView_TranslatesAutoresizingMaskIntoConstraints(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView translatesAutoresizingMaskIntoConstraints];
	return result;
}

void C_NSView_SetTranslatesAutoresizingMaskIntoConstraints(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setTranslatesAutoresizingMaskIntoConstraints:value];
}

bool C_NSView_HasAmbiguousLayout(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView hasAmbiguousLayout];
	return result;
}

CGRect C_NSView_FocusRingMaskBounds(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView focusRingMaskBounds];
	return result;
}

bool C_NSView_AllowsVibrancy(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView allowsVibrancy];
	return result;
}

bool C_NSView_IsInFullScreenMode(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView isInFullScreenMode];
	return result;
}

bool C_NSView_IsHidden(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView isHidden];
	return result;
}

void C_NSView_SetHidden(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setHidden:value];
}

bool C_NSView_IsHiddenOrHasHiddenAncestor(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView isHiddenOrHasHiddenAncestor];
	return result;
}

bool C_NSView_InLiveResize(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView inLiveResize];
	return result;
}

bool C_NSView_PreservesContentDuringLiveResize(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView preservesContentDuringLiveResize];
	return result;
}

CGRect C_NSView_RectPreservedDuringLiveResize(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView rectPreservedDuringLiveResize];
	return result;
}

bool C_NSView_MouseDownCanMoveWindow(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView mouseDownCanMoveWindow];
	return result;
}

void* C_NSView_InputContext(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSTextInputContext* result = [nSView inputContext];
	return result;
}

bool C_NSView_WantsRestingTouches(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView wantsRestingTouches];
	return result;
}

void C_NSView_SetWantsRestingTouches(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setWantsRestingTouches:value];
}

bool C_NSView_CanBecomeKeyView(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView canBecomeKeyView];
	return result;
}

bool C_NSView_NeedsPanelToBecomeKey(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView needsPanelToBecomeKey];
	return result;
}

void* C_NSView_NextKeyView(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSView* result = [nSView nextKeyView];
	return result;
}

void C_NSView_SetNextKeyView(void* ptr, void* value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setNextKeyView:(NSView*)value];
}

void* C_NSView_NextValidKeyView(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSView* result = [nSView nextValidKeyView];
	return result;
}

void* C_NSView_PreviousKeyView(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSView* result = [nSView previousKeyView];
	return result;
}

void* C_NSView_PreviousValidKeyView(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSView* result = [nSView previousValidKeyView];
	return result;
}

CGRect C_NSView_PreparedContentRect(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView preparedContentRect];
	return result;
}

void C_NSView_SetPreparedContentRect(void* ptr, CGRect value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setPreparedContentRect:value];
}

void* C_NSView_EnclosingScrollView(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSScrollView* result = [nSView enclosingScrollView];
	return result;
}

bool C_NSView_PostsFrameChangedNotifications(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView postsFrameChangedNotifications];
	return result;
}

void C_NSView_SetPostsFrameChangedNotifications(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setPostsFrameChangedNotifications:value];
}

bool C_NSView_PostsBoundsChangedNotifications(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView postsBoundsChangedNotifications];
	return result;
}

void C_NSView_SetPostsBoundsChangedNotifications(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setPostsBoundsChangedNotifications:value];
}

int C_NSView_Tag(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	int result = [nSView tag];
	return result;
}

void* C_NSView_ToolTip(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSString* result = [nSView toolTip];
	return result;
}

void C_NSView_SetToolTip(void* ptr, void* value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setToolTip:(NSString*)value];
}

bool C_NSView_IsDrawingFindIndicator(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView isDrawingFindIndicator];
	return result;
}

void* C_NSView_PressureConfiguration(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSPressureConfiguration* result = [nSView pressureConfiguration];
	return result;
}

void C_NSView_SetPressureConfiguration(void* ptr, void* value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setPressureConfiguration:(NSPressureConfiguration*)value];
}

NSEdgeInsets C_NSView_AdditionalSafeAreaInsets(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSEdgeInsets result = [nSView additionalSafeAreaInsets];
	return result;
}

void C_NSView_SetAdditionalSafeAreaInsets(void* ptr, NSEdgeInsets value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setAdditionalSafeAreaInsets:value];
}

void* C_NSView_CandidateListTouchBarItem(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSCandidateListTouchBarItem* result = [nSView candidateListTouchBarItem];
	return result;
}

bool C_NSView_IsHorizontalContentSizeConstraintActive(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView isHorizontalContentSizeConstraintActive];
	return result;
}

void C_NSView_SetHorizontalContentSizeConstraintActive(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setHorizontalContentSizeConstraintActive:value];
}

bool C_NSView_IsVerticalContentSizeConstraintActive(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	bool result = [nSView isVerticalContentSizeConstraintActive];
	return result;
}

void C_NSView_SetVerticalContentSizeConstraintActive(void* ptr, bool value) {
	NSView* nSView = (NSView*)ptr;
	[nSView setVerticalContentSizeConstraintActive:value];
}

NSEdgeInsets C_NSView_SafeAreaInsets(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	NSEdgeInsets result = [nSView safeAreaInsets];
	return result;
}

CGRect C_NSView_SafeAreaRect(void* ptr) {
	NSView* nSView = (NSView*)ptr;
	CGRect result = [nSView safeAreaRect];
	return result;
}
