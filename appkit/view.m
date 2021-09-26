#import "view.h"
#import <AppKit/NSCandidateListTouchBarItem.h>
#import <AppKit/NSClipView.h>
#import <AppKit/NSLayoutConstraint.h>
#import <AppKit/NSLayoutGuide.h>
#import <AppKit/NSPressureConfiguration.h>
#import <AppKit/NSRulerView.h>
#import <AppKit/NSView.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_View_Alloc() {
    return [NSView alloc];
}

void* C_NSView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSView* nSView = (NSView*)ptr;
    NSView* result_ = [nSView initWithFrame:frameRect];
    return result_;
}

void* C_NSView_InitWithCoder(void* ptr, void* coder) {
    NSView* nSView = (NSView*)ptr;
    NSView* result_ = [nSView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSView_Init(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSView* result_ = [nSView init];
    return result_;
}

void* C_NSView_AllocView() {
    NSView* result_ = [NSView alloc];
    return result_;
}

void* C_NSView_NewView() {
    NSView* result_ = [NSView new];
    return result_;
}

void* C_NSView_Autorelease(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSView* result_ = [nSView autorelease];
    return result_;
}

void* C_NSView_Retain(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSView* result_ = [nSView retain];
    return result_;
}

void C_NSView_PrepareForReuse(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    [nSView prepareForReuse];
}

void C_NSView_AddSubview(void* ptr, void* view) {
    NSView* nSView = (NSView*)ptr;
    [nSView addSubview:(NSView*)view];
}

void C_NSView_AddSubview_Positioned_RelativeTo(void* ptr, void* view, int place, void* otherView) {
    NSView* nSView = (NSView*)ptr;
    [nSView addSubview:(NSView*)view positioned:place relativeTo:(NSView*)otherView];
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
    BOOL result_ = [nSView isDescendantOf:(NSView*)view];
    return result_;
}

void* C_NSView_AncestorSharedWithView(void* ptr, void* view) {
    NSView* nSView = (NSView*)ptr;
    NSView* result_ = [nSView ancestorSharedWithView:(NSView*)view];
    return result_;
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
    BOOL result_ = [nSView needsToDrawRect:rect];
    return result_;
}

void* C_NSView_BitmapImageRepForCachingDisplayInRect(void* ptr, CGRect rect) {
    NSView* nSView = (NSView*)ptr;
    NSBitmapImageRep* result_ = [nSView bitmapImageRepForCachingDisplayInRect:rect];
    return result_;
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

void* C_NSView_DataWithEPSInsideRect(void* ptr, CGRect rect) {
    NSView* nSView = (NSView*)ptr;
    NSData* result_ = [nSView dataWithEPSInsideRect:rect];
    return result_;
}

void* C_NSView_DataWithPDFInsideRect(void* ptr, CGRect rect) {
    NSView* nSView = (NSView*)ptr;
    NSData* result_ = [nSView dataWithPDFInsideRect:rect];
    return result_;
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
    NSRect result_ = [nSView rectForPage:page];
    return result_;
}

CGPoint C_NSView_LocationOfPrintRect(void* ptr, CGRect rect) {
    NSView* nSView = (NSView*)ptr;
    NSPoint result_ = [nSView locationOfPrintRect:rect];
    return result_;
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
    NSPoint result_ = [nSView convertPointFromBacking:point];
    return result_;
}

CGPoint C_NSView_ConvertPointToBacking(void* ptr, CGPoint point) {
    NSView* nSView = (NSView*)ptr;
    NSPoint result_ = [nSView convertPointToBacking:point];
    return result_;
}

CGPoint C_NSView_ConvertPointFromLayer(void* ptr, CGPoint point) {
    NSView* nSView = (NSView*)ptr;
    NSPoint result_ = [nSView convertPointFromLayer:point];
    return result_;
}

CGPoint C_NSView_ConvertPointToLayer(void* ptr, CGPoint point) {
    NSView* nSView = (NSView*)ptr;
    NSPoint result_ = [nSView convertPointToLayer:point];
    return result_;
}

CGRect C_NSView_ConvertRectFromBacking(void* ptr, CGRect rect) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView convertRectFromBacking:rect];
    return result_;
}

CGRect C_NSView_ConvertRectToBacking(void* ptr, CGRect rect) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView convertRectToBacking:rect];
    return result_;
}

CGRect C_NSView_ConvertRectFromLayer(void* ptr, CGRect rect) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView convertRectFromLayer:rect];
    return result_;
}

CGRect C_NSView_ConvertRectToLayer(void* ptr, CGRect rect) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView convertRectToLayer:rect];
    return result_;
}

CGSize C_NSView_ConvertSizeFromBacking(void* ptr, CGSize size) {
    NSView* nSView = (NSView*)ptr;
    NSSize result_ = [nSView convertSizeFromBacking:size];
    return result_;
}

CGSize C_NSView_ConvertSizeToBacking(void* ptr, CGSize size) {
    NSView* nSView = (NSView*)ptr;
    NSSize result_ = [nSView convertSizeToBacking:size];
    return result_;
}

CGSize C_NSView_ConvertSizeFromLayer(void* ptr, CGSize size) {
    NSView* nSView = (NSView*)ptr;
    NSSize result_ = [nSView convertSizeFromLayer:size];
    return result_;
}

CGSize C_NSView_ConvertSizeToLayer(void* ptr, CGSize size) {
    NSView* nSView = (NSView*)ptr;
    NSSize result_ = [nSView convertSizeToLayer:size];
    return result_;
}

CGPoint C_NSView_ConvertPoint_FromView(void* ptr, CGPoint point, void* view) {
    NSView* nSView = (NSView*)ptr;
    NSPoint result_ = [nSView convertPoint:point fromView:(NSView*)view];
    return result_;
}

CGPoint C_NSView_ConvertPoint_ToView(void* ptr, CGPoint point, void* view) {
    NSView* nSView = (NSView*)ptr;
    NSPoint result_ = [nSView convertPoint:point toView:(NSView*)view];
    return result_;
}

CGSize C_NSView_ConvertSize_FromView(void* ptr, CGSize size, void* view) {
    NSView* nSView = (NSView*)ptr;
    NSSize result_ = [nSView convertSize:size fromView:(NSView*)view];
    return result_;
}

CGSize C_NSView_ConvertSize_ToView(void* ptr, CGSize size, void* view) {
    NSView* nSView = (NSView*)ptr;
    NSSize result_ = [nSView convertSize:size toView:(NSView*)view];
    return result_;
}

CGRect C_NSView_ConvertRect_FromView(void* ptr, CGRect rect, void* view) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView convertRect:rect fromView:(NSView*)view];
    return result_;
}

CGRect C_NSView_ConvertRect_ToView(void* ptr, CGRect rect, void* view) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView convertRect:rect toView:(NSView*)view];
    return result_;
}

CGRect C_NSView_CenterScanRect(void* ptr, CGRect rect) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView centerScanRect:rect];
    return result_;
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

void C_NSView_AddConstraints(void* ptr, Array constraints) {
    NSView* nSView = (NSView*)ptr;
    NSMutableArray* objcConstraints = [NSMutableArray arrayWithCapacity:constraints.len];
    if (constraints.len > 0) {
    	void** constraintsData = (void**)constraints.data;
    	for (int i = 0; i < constraints.len; i++) {
    		void* p = constraintsData[i];
    		[objcConstraints addObject:(NSLayoutConstraint*)p];
    	}
    }
    [nSView addConstraints:objcConstraints];
}

void C_NSView_RemoveConstraint(void* ptr, void* constraint) {
    NSView* nSView = (NSView*)ptr;
    [nSView removeConstraint:(NSLayoutConstraint*)constraint];
}

void C_NSView_RemoveConstraints(void* ptr, Array constraints) {
    NSView* nSView = (NSView*)ptr;
    NSMutableArray* objcConstraints = [NSMutableArray arrayWithCapacity:constraints.len];
    if (constraints.len > 0) {
    	void** constraintsData = (void**)constraints.data;
    	for (int i = 0; i < constraints.len; i++) {
    		void* p = constraintsData[i];
    		[objcConstraints addObject:(NSLayoutConstraint*)p];
    	}
    }
    [nSView removeConstraints:objcConstraints];
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

float C_NSView_ContentCompressionResistancePriorityForOrientation(void* ptr, int orientation) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutPriority result_ = [nSView contentCompressionResistancePriorityForOrientation:orientation];
    return result_;
}

void C_NSView_SetContentCompressionResistancePriority_ForOrientation(void* ptr, float priority, int orientation) {
    NSView* nSView = (NSView*)ptr;
    [nSView setContentCompressionResistancePriority:priority forOrientation:orientation];
}

float C_NSView_ContentHuggingPriorityForOrientation(void* ptr, int orientation) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutPriority result_ = [nSView contentHuggingPriorityForOrientation:orientation];
    return result_;
}

void C_NSView_SetContentHuggingPriority_ForOrientation(void* ptr, float priority, int orientation) {
    NSView* nSView = (NSView*)ptr;
    [nSView setContentHuggingPriority:priority forOrientation:orientation];
}

CGRect C_NSView_AlignmentRectForFrame(void* ptr, CGRect frame) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView alignmentRectForFrame:frame];
    return result_;
}

CGRect C_NSView_FrameForAlignmentRect(void* ptr, CGRect alignmentRect) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView frameForAlignmentRect:alignmentRect];
    return result_;
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

Array C_NSView_ConstraintsAffectingLayoutForOrientation(void* ptr, int orientation) {
    NSView* nSView = (NSView*)ptr;
    NSArray* result_ = [nSView constraintsAffectingLayoutForOrientation:orientation];
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

bool C_NSView_EnterFullScreenMode_WithOptions(void* ptr, void* screen, Dictionary options) {
    NSView* nSView = (NSView*)ptr;
    NSMutableDictionary* objcOptions = [NSMutableDictionary dictionaryWithCapacity:options.len];
    if (options.len > 0) {
    	void** optionsKeyData = (void**)options.key_data;
    	void** optionsValueData = (void**)options.value_data;
    	for (int i = 0; i < options.len; i++) {
    		void* kp = optionsKeyData[i];
    		void* vp = optionsValueData[i];
    		[objcOptions setObject:(id)vp forKey:(id<NSCopying>)(NSString*)kp];
    	}
    }
    BOOL result_ = [nSView enterFullScreenMode:(NSScreen*)screen withOptions:objcOptions];
    return result_;
}

void C_NSView_ExitFullScreenModeWithOptions(void* ptr, Dictionary options) {
    NSView* nSView = (NSView*)ptr;
    NSMutableDictionary* objcOptions = [NSMutableDictionary dictionaryWithCapacity:options.len];
    if (options.len > 0) {
    	void** optionsKeyData = (void**)options.key_data;
    	void** optionsValueData = (void**)options.value_data;
    	for (int i = 0; i < options.len; i++) {
    		void* kp = optionsKeyData[i];
    		void* vp = optionsValueData[i];
    		[objcOptions setObject:(id)vp forKey:(id<NSCopying>)(NSString*)kp];
    	}
    }
    [nSView exitFullScreenModeWithOptions:objcOptions];
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
    BOOL result_ = [nSView acceptsFirstMouse:(NSEvent*)event];
    return result_;
}

void* C_NSView_HitTest(void* ptr, CGPoint point) {
    NSView* nSView = (NSView*)ptr;
    NSView* result_ = [nSView hitTest:point];
    return result_;
}

bool C_NSView_Mouse_InRect(void* ptr, CGPoint point, CGRect rect) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView mouse:point inRect:rect];
    return result_;
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
    BOOL result_ = [nSView scrollRectToVisible:rect];
    return result_;
}

bool C_NSView_Autoscroll(void* ptr, void* event) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView autoscroll:(NSEvent*)event];
    return result_;
}

CGRect C_NSView_AdjustScroll(void* ptr, CGRect newVisible) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView adjustScroll:newVisible];
    return result_;
}

void C_NSView_ScrollClipView_ToPoint(void* ptr, void* clipView, CGPoint point) {
    NSView* nSView = (NSView*)ptr;
    [nSView scrollClipView:(NSClipView*)clipView toPoint:point];
}

void C_NSView_ReflectScrolledClipView(void* ptr, void* clipView) {
    NSView* nSView = (NSView*)ptr;
    [nSView reflectScrolledClipView:(NSClipView*)clipView];
}

void C_NSView_RegisterForDraggedTypes(void* ptr, Array newTypes) {
    NSView* nSView = (NSView*)ptr;
    NSMutableArray* objcNewTypes = [NSMutableArray arrayWithCapacity:newTypes.len];
    if (newTypes.len > 0) {
    	void** newTypesData = (void**)newTypes.data;
    	for (int i = 0; i < newTypes.len; i++) {
    		void* p = newTypesData[i];
    		[objcNewTypes addObject:(NSString*)p];
    	}
    }
    [nSView registerForDraggedTypes:objcNewTypes];
}

void C_NSView_UnregisterDraggedTypes(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    [nSView unregisterDraggedTypes];
}

void* C_NSView_BeginDraggingSessionWithItems_Event_Source(void* ptr, Array items, void* event, void* source) {
    NSView* nSView = (NSView*)ptr;
    NSMutableArray* objcItems = [NSMutableArray arrayWithCapacity:items.len];
    if (items.len > 0) {
    	void** itemsData = (void**)items.data;
    	for (int i = 0; i < items.len; i++) {
    		void* p = itemsData[i];
    		[objcItems addObject:(NSDraggingItem*)p];
    	}
    }
    NSDraggingSession* result_ = [nSView beginDraggingSessionWithItems:objcItems event:(NSEvent*)event source:(id)source];
    return result_;
}

bool C_NSView_ShouldDelayWindowOrderingForEvent(void* ptr, void* event) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView shouldDelayWindowOrderingForEvent:(NSEvent*)event];
    return result_;
}

CGRect C_NSView_RectForSmartMagnificationAtPoint_InRect(void* ptr, CGPoint location, CGRect visibleRect) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView rectForSmartMagnificationAtPoint:location inRect:visibleRect];
    return result_;
}

void C_NSView_ViewDidChangeBackingProperties(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    [nSView viewDidChangeBackingProperties];
}

void* C_NSView_ViewWithTag(void* ptr, int tag) {
    NSView* nSView = (NSView*)ptr;
    NSView* result_ = [nSView viewWithTag:tag];
    return result_;
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
    NSMenu* result_ = [nSView menuForEvent:(NSEvent*)event];
    return result_;
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
    CGFloat result_ = [nSView rulerView:(NSRulerView*)ruler locationForPoint:point];
    return result_;
}

CGPoint C_NSView_RulerView_PointForLocation(void* ptr, void* ruler, double point) {
    NSView* nSView = (NSView*)ptr;
    NSPoint result_ = [nSView rulerView:(NSRulerView*)ruler pointForLocation:point];
    return result_;
}

bool C_NSView_RulerView_ShouldAddMarker(void* ptr, void* ruler, void* marker) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView rulerView:(NSRulerView*)ruler shouldAddMarker:(NSRulerMarker*)marker];
    return result_;
}

bool C_NSView_RulerView_ShouldMoveMarker(void* ptr, void* ruler, void* marker) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView rulerView:(NSRulerView*)ruler shouldMoveMarker:(NSRulerMarker*)marker];
    return result_;
}

bool C_NSView_RulerView_ShouldRemoveMarker(void* ptr, void* ruler, void* marker) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView rulerView:(NSRulerView*)ruler shouldRemoveMarker:(NSRulerMarker*)marker];
    return result_;
}

double C_NSView_RulerView_WillAddMarker_AtLocation(void* ptr, void* ruler, void* marker, double location) {
    NSView* nSView = (NSView*)ptr;
    CGFloat result_ = [nSView rulerView:(NSRulerView*)ruler willAddMarker:(NSRulerMarker*)marker atLocation:location];
    return result_;
}

double C_NSView_RulerView_WillMoveMarker_ToLocation(void* ptr, void* ruler, void* marker, double location) {
    NSView* nSView = (NSView*)ptr;
    CGFloat result_ = [nSView rulerView:(NSRulerView*)ruler willMoveMarker:(NSRulerMarker*)marker toLocation:location];
    return result_;
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
    NSView* result_ = [nSView superview];
    return result_;
}

Array C_NSView_Subviews(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSArray* result_ = [nSView subviews];
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

void C_NSView_SetSubviews(void* ptr, Array value) {
    NSView* nSView = (NSView*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSView*)p];
    	}
    }
    [nSView setSubviews:objcValue];
}

void* C_NSView_Window(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSWindow* result_ = [nSView window];
    return result_;
}

void* C_NSView_OpaqueAncestor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSView* result_ = [nSView opaqueAncestor];
    return result_;
}

void* C_NSView_EnclosingMenuItem(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSMenuItem* result_ = [nSView enclosingMenuItem];
    return result_;
}

CGRect C_NSView_Frame(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView frame];
    return result_;
}

void C_NSView_SetFrame(void* ptr, CGRect value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setFrame:value];
}

double C_NSView_FrameRotation(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    CGFloat result_ = [nSView frameRotation];
    return result_;
}

void C_NSView_SetFrameRotation(void* ptr, double value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setFrameRotation:value];
}

CGRect C_NSView_Bounds(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView bounds];
    return result_;
}

void C_NSView_SetBounds(void* ptr, CGRect value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setBounds:value];
}

double C_NSView_BoundsRotation(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    CGFloat result_ = [nSView boundsRotation];
    return result_;
}

void C_NSView_SetBoundsRotation(void* ptr, double value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setBoundsRotation:value];
}

bool C_NSView_WantsLayer(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView wantsLayer];
    return result_;
}

void C_NSView_SetWantsLayer(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setWantsLayer:value];
}

bool C_NSView_WantsUpdateLayer(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView wantsUpdateLayer];
    return result_;
}

int C_NSView_LayerContentsPlacement(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSViewLayerContentsPlacement result_ = [nSView layerContentsPlacement];
    return result_;
}

void C_NSView_SetLayerContentsPlacement(void* ptr, int value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setLayerContentsPlacement:value];
}

int C_NSView_LayerContentsRedrawPolicy(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSViewLayerContentsRedrawPolicy result_ = [nSView layerContentsRedrawPolicy];
    return result_;
}

void C_NSView_SetLayerContentsRedrawPolicy(void* ptr, int value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setLayerContentsRedrawPolicy:value];
}

bool C_NSView_CanDrawSubviewsIntoLayer(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView canDrawSubviewsIntoLayer];
    return result_;
}

void C_NSView_SetCanDrawSubviewsIntoLayer(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setCanDrawSubviewsIntoLayer:value];
}

bool C_NSView_LayerUsesCoreImageFilters(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView layerUsesCoreImageFilters];
    return result_;
}

void C_NSView_SetLayerUsesCoreImageFilters(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setLayerUsesCoreImageFilters:value];
}

double C_NSView_AlphaValue(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    CGFloat result_ = [nSView alphaValue];
    return result_;
}

void C_NSView_SetAlphaValue(void* ptr, double value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setAlphaValue:value];
}

double C_NSView_FrameCenterRotation(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    CGFloat result_ = [nSView frameCenterRotation];
    return result_;
}

void C_NSView_SetFrameCenterRotation(void* ptr, double value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setFrameCenterRotation:value];
}

void* C_NSView_Shadow(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSShadow* result_ = [nSView shadow];
    return result_;
}

void C_NSView_SetShadow(void* ptr, void* value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setShadow:(NSShadow*)value];
}

bool C_NSView_CanDrawConcurrently(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView canDrawConcurrently];
    return result_;
}

void C_NSView_SetCanDrawConcurrently(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setCanDrawConcurrently:value];
}

CGRect C_NSView_VisibleRect(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView visibleRect];
    return result_;
}

bool C_NSView_WantsDefaultClipping(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView wantsDefaultClipping];
    return result_;
}

void* C_NSView_PrintJobTitle(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSString* result_ = [nSView printJobTitle];
    return result_;
}

void* C_NSView_PageHeader(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSAttributedString* result_ = [nSView pageHeader];
    return result_;
}

void* C_NSView_PageFooter(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSAttributedString* result_ = [nSView pageFooter];
    return result_;
}

double C_NSView_HeightAdjustLimit(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    CGFloat result_ = [nSView heightAdjustLimit];
    return result_;
}

double C_NSView_WidthAdjustLimit(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    CGFloat result_ = [nSView widthAdjustLimit];
    return result_;
}

bool C_NSView_NeedsDisplay(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView needsDisplay];
    return result_;
}

void C_NSView_SetNeedsDisplay(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setNeedsDisplay:value];
}

bool C_NSView_IsOpaque(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView isOpaque];
    return result_;
}

bool C_NSView_IsFlipped(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView isFlipped];
    return result_;
}

bool C_NSView_IsRotatedFromBase(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView isRotatedFromBase];
    return result_;
}

bool C_NSView_IsRotatedOrScaledFromBase(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView isRotatedOrScaledFromBase];
    return result_;
}

bool C_NSView_AutoresizesSubviews(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView autoresizesSubviews];
    return result_;
}

void C_NSView_SetAutoresizesSubviews(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setAutoresizesSubviews:value];
}

unsigned int C_NSView_AutoresizingMask(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSAutoresizingMaskOptions result_ = [nSView autoresizingMask];
    return result_;
}

void C_NSView_SetAutoresizingMask(void* ptr, unsigned int value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setAutoresizingMask:value];
}

void* C_NSView_BottomAnchor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutYAxisAnchor* result_ = [nSView bottomAnchor];
    return result_;
}

void* C_NSView_CenterXAnchor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutXAxisAnchor* result_ = [nSView centerXAnchor];
    return result_;
}

void* C_NSView_CenterYAnchor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutYAxisAnchor* result_ = [nSView centerYAnchor];
    return result_;
}

void* C_NSView_FirstBaselineAnchor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutYAxisAnchor* result_ = [nSView firstBaselineAnchor];
    return result_;
}

void* C_NSView_HeightAnchor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutDimension* result_ = [nSView heightAnchor];
    return result_;
}

void* C_NSView_LastBaselineAnchor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutYAxisAnchor* result_ = [nSView lastBaselineAnchor];
    return result_;
}

void* C_NSView_LeadingAnchor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutXAxisAnchor* result_ = [nSView leadingAnchor];
    return result_;
}

void* C_NSView_LeftAnchor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutXAxisAnchor* result_ = [nSView leftAnchor];
    return result_;
}

void* C_NSView_RightAnchor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutXAxisAnchor* result_ = [nSView rightAnchor];
    return result_;
}

void* C_NSView_TopAnchor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutYAxisAnchor* result_ = [nSView topAnchor];
    return result_;
}

void* C_NSView_TrailingAnchor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutXAxisAnchor* result_ = [nSView trailingAnchor];
    return result_;
}

void* C_NSView_WidthAnchor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutDimension* result_ = [nSView widthAnchor];
    return result_;
}

Array C_NSView_Constraints(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSArray* result_ = [nSView constraints];
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

Array C_NSView_LayoutGuides(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSArray* result_ = [nSView layoutGuides];
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

void* C_NSView_LayoutMarginsGuide(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutGuide* result_ = [nSView layoutMarginsGuide];
    return result_;
}

void* C_NSView_SafeAreaLayoutGuide(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSLayoutGuide* result_ = [nSView safeAreaLayoutGuide];
    return result_;
}

CGSize C_NSView_FittingSize(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSSize result_ = [nSView fittingSize];
    return result_;
}

CGSize C_NSView_IntrinsicContentSize(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSSize result_ = [nSView intrinsicContentSize];
    return result_;
}

NSEdgeInsets C_NSView_AlignmentRectInsets(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSEdgeInsets result_ = [nSView alignmentRectInsets];
    return result_;
}

double C_NSView_BaselineOffsetFromBottom(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    CGFloat result_ = [nSView baselineOffsetFromBottom];
    return result_;
}

double C_NSView_FirstBaselineOffsetFromTop(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    CGFloat result_ = [nSView firstBaselineOffsetFromTop];
    return result_;
}

double C_NSView_LastBaselineOffsetFromBottom(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    CGFloat result_ = [nSView lastBaselineOffsetFromBottom];
    return result_;
}

bool C_NSView_NeedsLayout(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView needsLayout];
    return result_;
}

void C_NSView_SetNeedsLayout(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setNeedsLayout:value];
}

bool C_NSView_NeedsUpdateConstraints(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView needsUpdateConstraints];
    return result_;
}

void C_NSView_SetNeedsUpdateConstraints(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setNeedsUpdateConstraints:value];
}

bool C_NSView_View_RequiresConstraintBasedLayout() {
    BOOL result_ = [NSView requiresConstraintBasedLayout];
    return result_;
}

bool C_NSView_TranslatesAutoresizingMaskIntoConstraints(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView translatesAutoresizingMaskIntoConstraints];
    return result_;
}

void C_NSView_SetTranslatesAutoresizingMaskIntoConstraints(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setTranslatesAutoresizingMaskIntoConstraints:value];
}

bool C_NSView_HasAmbiguousLayout(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView hasAmbiguousLayout];
    return result_;
}

void* C_NSView_FocusView() {
    NSView* result_ = [NSView focusView];
    return result_;
}

unsigned int C_NSView_FocusRingType(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSFocusRingType result_ = [nSView focusRingType];
    return result_;
}

void C_NSView_SetFocusRingType(void* ptr, unsigned int value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setFocusRingType:value];
}

CGRect C_NSView_FocusRingMaskBounds(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView focusRingMaskBounds];
    return result_;
}

unsigned int C_NSView_View_DefaultFocusRingType() {
    NSFocusRingType result_ = [NSView defaultFocusRingType];
    return result_;
}

bool C_NSView_AllowsVibrancy(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView allowsVibrancy];
    return result_;
}

bool C_NSView_IsInFullScreenMode(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView isInFullScreenMode];
    return result_;
}

bool C_NSView_IsHidden(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView isHidden];
    return result_;
}

void C_NSView_SetHidden(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setHidden:value];
}

bool C_NSView_IsHiddenOrHasHiddenAncestor(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView isHiddenOrHasHiddenAncestor];
    return result_;
}

bool C_NSView_InLiveResize(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView inLiveResize];
    return result_;
}

bool C_NSView_PreservesContentDuringLiveResize(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView preservesContentDuringLiveResize];
    return result_;
}

CGRect C_NSView_RectPreservedDuringLiveResize(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView rectPreservedDuringLiveResize];
    return result_;
}

Array C_NSView_GestureRecognizers(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSArray* result_ = [nSView gestureRecognizers];
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

void C_NSView_SetGestureRecognizers(void* ptr, Array value) {
    NSView* nSView = (NSView*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSGestureRecognizer*)p];
    	}
    }
    [nSView setGestureRecognizers:objcValue];
}

bool C_NSView_MouseDownCanMoveWindow(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView mouseDownCanMoveWindow];
    return result_;
}

void* C_NSView_InputContext(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSTextInputContext* result_ = [nSView inputContext];
    return result_;
}

bool C_NSView_WantsRestingTouches(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView wantsRestingTouches];
    return result_;
}

void C_NSView_SetWantsRestingTouches(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setWantsRestingTouches:value];
}

bool C_NSView_CanBecomeKeyView(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView canBecomeKeyView];
    return result_;
}

bool C_NSView_NeedsPanelToBecomeKey(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView needsPanelToBecomeKey];
    return result_;
}

void* C_NSView_NextKeyView(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSView* result_ = [nSView nextKeyView];
    return result_;
}

void C_NSView_SetNextKeyView(void* ptr, void* value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setNextKeyView:(NSView*)value];
}

void* C_NSView_NextValidKeyView(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSView* result_ = [nSView nextValidKeyView];
    return result_;
}

void* C_NSView_PreviousKeyView(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSView* result_ = [nSView previousKeyView];
    return result_;
}

void* C_NSView_PreviousValidKeyView(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSView* result_ = [nSView previousValidKeyView];
    return result_;
}

CGRect C_NSView_PreparedContentRect(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView preparedContentRect];
    return result_;
}

void C_NSView_SetPreparedContentRect(void* ptr, CGRect value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setPreparedContentRect:value];
}

void* C_NSView_EnclosingScrollView(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSScrollView* result_ = [nSView enclosingScrollView];
    return result_;
}

Array C_NSView_RegisteredDraggedTypes(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSArray* result_ = [nSView registeredDraggedTypes];
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

bool C_NSView_PostsFrameChangedNotifications(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView postsFrameChangedNotifications];
    return result_;
}

void C_NSView_SetPostsFrameChangedNotifications(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setPostsFrameChangedNotifications:value];
}

bool C_NSView_PostsBoundsChangedNotifications(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView postsBoundsChangedNotifications];
    return result_;
}

void C_NSView_SetPostsBoundsChangedNotifications(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setPostsBoundsChangedNotifications:value];
}

int C_NSView_Tag(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSInteger result_ = [nSView tag];
    return result_;
}

void* C_NSView_ToolTip(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSString* result_ = [nSView toolTip];
    return result_;
}

void C_NSView_SetToolTip(void* ptr, void* value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setToolTip:(NSString*)value];
}

Array C_NSView_TrackingAreas(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSArray* result_ = [nSView trackingAreas];
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

void* C_NSView_View_DefaultMenu() {
    NSMenu* result_ = [NSView defaultMenu];
    return result_;
}

bool C_NSView_IsDrawingFindIndicator(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView isDrawingFindIndicator];
    return result_;
}

int C_NSView_UserInterfaceLayoutDirection(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSUserInterfaceLayoutDirection result_ = [nSView userInterfaceLayoutDirection];
    return result_;
}

void C_NSView_SetUserInterfaceLayoutDirection(void* ptr, int value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setUserInterfaceLayoutDirection:value];
}

void* C_NSView_PressureConfiguration(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSPressureConfiguration* result_ = [nSView pressureConfiguration];
    return result_;
}

void C_NSView_SetPressureConfiguration(void* ptr, void* value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setPressureConfiguration:(NSPressureConfiguration*)value];
}

NSEdgeInsets C_NSView_AdditionalSafeAreaInsets(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSEdgeInsets result_ = [nSView additionalSafeAreaInsets];
    return result_;
}

void C_NSView_SetAdditionalSafeAreaInsets(void* ptr, NSEdgeInsets value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setAdditionalSafeAreaInsets:value];
}

unsigned int C_NSView_AllowedTouchTypes(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSTouchTypeMask result_ = [nSView allowedTouchTypes];
    return result_;
}

void C_NSView_SetAllowedTouchTypes(void* ptr, unsigned int value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setAllowedTouchTypes:value];
}

void* C_NSView_CandidateListTouchBarItem(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSCandidateListTouchBarItem* result_ = [nSView candidateListTouchBarItem];
    return result_;
}

bool C_NSView_IsHorizontalContentSizeConstraintActive(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView isHorizontalContentSizeConstraintActive];
    return result_;
}

void C_NSView_SetHorizontalContentSizeConstraintActive(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setHorizontalContentSizeConstraintActive:value];
}

bool C_NSView_IsVerticalContentSizeConstraintActive(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    BOOL result_ = [nSView isVerticalContentSizeConstraintActive];
    return result_;
}

void C_NSView_SetVerticalContentSizeConstraintActive(void* ptr, bool value) {
    NSView* nSView = (NSView*)ptr;
    [nSView setVerticalContentSizeConstraintActive:value];
}

NSEdgeInsets C_NSView_SafeAreaInsets(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSEdgeInsets result_ = [nSView safeAreaInsets];
    return result_;
}

CGRect C_NSView_SafeAreaRect(void* ptr) {
    NSView* nSView = (NSView*)ptr;
    NSRect result_ = [nSView safeAreaRect];
    return result_;
}

bool C_NSView_View_IsCompatibleWithResponsiveScrolling() {
    BOOL result_ = [NSView isCompatibleWithResponsiveScrolling];
    return result_;
}
