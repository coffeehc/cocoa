#import <Appkit/Appkit.h>
#import "scroll_view.h"

void* C_ScrollView_Alloc() {
    return [NSScrollView alloc];
}

void* C_NSScrollView_InitWithCoder(void* ptr, void* coder) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSScrollView* result = [nSScrollView initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSScrollView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSScrollView* result = [nSScrollView initWithFrame:frameRect];
    return result;
}

void* C_NSScrollView_Init(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSScrollView* result = [nSScrollView init];
    return result;
}

void C_NSScrollView_AddFloatingSubview_ForAxis(void* ptr, void* view, int axis) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView addFloatingSubview:(NSView*)view forAxis:axis];
}

void C_NSScrollView_Tile(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView tile];
}

void C_NSScrollView_FlashScrollers(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView flashScrollers];
}

void C_NSScrollView_MagnifyToFitRect(void* ptr, CGRect rect) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView magnifyToFitRect:rect];
}

void C_NSScrollView_SetMagnification_CenteredAtPoint(void* ptr, double magnification, CGPoint point) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setMagnification:magnification centeredAtPoint:point];
}

CGSize C_NSScrollView_ContentSize(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSSize result = [nSScrollView contentSize];
    return result;
}

CGRect C_NSScrollView_DocumentVisibleRect(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSRect result = [nSScrollView documentVisibleRect];
    return result;
}

void* C_NSScrollView_BackgroundColor(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSColor* result = [nSScrollView backgroundColor];
    return result;
}

void C_NSScrollView_SetBackgroundColor(void* ptr, void* value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setBackgroundColor:(NSColor*)value];
}

bool C_NSScrollView_DrawsBackground(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    BOOL result = [nSScrollView drawsBackground];
    return result;
}

void C_NSScrollView_SetDrawsBackground(void* ptr, bool value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setDrawsBackground:value];
}

unsigned int C_NSScrollView_BorderType(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSBorderType result = [nSScrollView borderType];
    return result;
}

void C_NSScrollView_SetBorderType(void* ptr, unsigned int value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setBorderType:value];
}

void* C_NSScrollView_DocumentCursor(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSCursor* result = [nSScrollView documentCursor];
    return result;
}

void C_NSScrollView_SetDocumentCursor(void* ptr, void* value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setDocumentCursor:(NSCursor*)value];
}

void* C_NSScrollView_ContentView(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSClipView* result = [nSScrollView contentView];
    return result;
}

void C_NSScrollView_SetContentView(void* ptr, void* value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setContentView:(NSClipView*)value];
}

void* C_NSScrollView_DocumentView(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSView* result = [nSScrollView documentView];
    return result;
}

void C_NSScrollView_SetDocumentView(void* ptr, void* value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setDocumentView:(NSView*)value];
}

void* C_NSScrollView_HorizontalScroller(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSScroller* result = [nSScrollView horizontalScroller];
    return result;
}

void C_NSScrollView_SetHorizontalScroller(void* ptr, void* value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setHorizontalScroller:(NSScroller*)value];
}

bool C_NSScrollView_HasHorizontalScroller(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    BOOL result = [nSScrollView hasHorizontalScroller];
    return result;
}

void C_NSScrollView_SetHasHorizontalScroller(void* ptr, bool value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setHasHorizontalScroller:value];
}

void* C_NSScrollView_VerticalScroller(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSScroller* result = [nSScrollView verticalScroller];
    return result;
}

void C_NSScrollView_SetVerticalScroller(void* ptr, void* value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setVerticalScroller:(NSScroller*)value];
}

bool C_NSScrollView_HasVerticalScroller(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    BOOL result = [nSScrollView hasVerticalScroller];
    return result;
}

void C_NSScrollView_SetHasVerticalScroller(void* ptr, bool value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setHasVerticalScroller:value];
}

bool C_NSScrollView_AutohidesScrollers(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    BOOL result = [nSScrollView autohidesScrollers];
    return result;
}

void C_NSScrollView_SetAutohidesScrollers(void* ptr, bool value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setAutohidesScrollers:value];
}

bool C_NSScrollView_HasHorizontalRuler(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    BOOL result = [nSScrollView hasHorizontalRuler];
    return result;
}

void C_NSScrollView_SetHasHorizontalRuler(void* ptr, bool value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setHasHorizontalRuler:value];
}

void* C_NSScrollView_HorizontalRulerView(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSRulerView* result = [nSScrollView horizontalRulerView];
    return result;
}

void C_NSScrollView_SetHorizontalRulerView(void* ptr, void* value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setHorizontalRulerView:(NSRulerView*)value];
}

bool C_NSScrollView_HasVerticalRuler(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    BOOL result = [nSScrollView hasVerticalRuler];
    return result;
}

void C_NSScrollView_SetHasVerticalRuler(void* ptr, bool value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setHasVerticalRuler:value];
}

void* C_NSScrollView_VerticalRulerView(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSRulerView* result = [nSScrollView verticalRulerView];
    return result;
}

void C_NSScrollView_SetVerticalRulerView(void* ptr, void* value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setVerticalRulerView:(NSRulerView*)value];
}

bool C_NSScrollView_RulersVisible(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    BOOL result = [nSScrollView rulersVisible];
    return result;
}

void C_NSScrollView_SetRulersVisible(void* ptr, bool value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setRulersVisible:value];
}

bool C_NSScrollView_AutomaticallyAdjustsContentInsets(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    BOOL result = [nSScrollView automaticallyAdjustsContentInsets];
    return result;
}

void C_NSScrollView_SetAutomaticallyAdjustsContentInsets(void* ptr, bool value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setAutomaticallyAdjustsContentInsets:value];
}

NSEdgeInsets C_NSScrollView_ContentInsets(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSEdgeInsets result = [nSScrollView contentInsets];
    return result;
}

void C_NSScrollView_SetContentInsets(void* ptr, NSEdgeInsets value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setContentInsets:value];
}

NSEdgeInsets C_NSScrollView_ScrollerInsets(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSEdgeInsets result = [nSScrollView scrollerInsets];
    return result;
}

void C_NSScrollView_SetScrollerInsets(void* ptr, NSEdgeInsets value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setScrollerInsets:value];
}

int C_NSScrollView_ScrollerKnobStyle(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSScrollerKnobStyle result = [nSScrollView scrollerKnobStyle];
    return result;
}

void C_NSScrollView_SetScrollerKnobStyle(void* ptr, int value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setScrollerKnobStyle:value];
}

int C_NSScrollView_ScrollerStyle(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSScrollerStyle result = [nSScrollView scrollerStyle];
    return result;
}

void C_NSScrollView_SetScrollerStyle(void* ptr, int value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setScrollerStyle:value];
}

double C_NSScrollView_LineScroll(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    CGFloat result = [nSScrollView lineScroll];
    return result;
}

void C_NSScrollView_SetLineScroll(void* ptr, double value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setLineScroll:value];
}

double C_NSScrollView_HorizontalLineScroll(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    CGFloat result = [nSScrollView horizontalLineScroll];
    return result;
}

void C_NSScrollView_SetHorizontalLineScroll(void* ptr, double value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setHorizontalLineScroll:value];
}

double C_NSScrollView_VerticalLineScroll(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    CGFloat result = [nSScrollView verticalLineScroll];
    return result;
}

void C_NSScrollView_SetVerticalLineScroll(void* ptr, double value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setVerticalLineScroll:value];
}

double C_NSScrollView_PageScroll(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    CGFloat result = [nSScrollView pageScroll];
    return result;
}

void C_NSScrollView_SetPageScroll(void* ptr, double value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setPageScroll:value];
}

double C_NSScrollView_HorizontalPageScroll(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    CGFloat result = [nSScrollView horizontalPageScroll];
    return result;
}

void C_NSScrollView_SetHorizontalPageScroll(void* ptr, double value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setHorizontalPageScroll:value];
}

double C_NSScrollView_VerticalPageScroll(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    CGFloat result = [nSScrollView verticalPageScroll];
    return result;
}

void C_NSScrollView_SetVerticalPageScroll(void* ptr, double value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setVerticalPageScroll:value];
}

bool C_NSScrollView_ScrollsDynamically(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    BOOL result = [nSScrollView scrollsDynamically];
    return result;
}

void C_NSScrollView_SetScrollsDynamically(void* ptr, bool value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setScrollsDynamically:value];
}

int C_NSScrollView_FindBarPosition(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSScrollViewFindBarPosition result = [nSScrollView findBarPosition];
    return result;
}

void C_NSScrollView_SetFindBarPosition(void* ptr, int value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setFindBarPosition:value];
}

bool C_NSScrollView_UsesPredominantAxisScrolling(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    BOOL result = [nSScrollView usesPredominantAxisScrolling];
    return result;
}

void C_NSScrollView_SetUsesPredominantAxisScrolling(void* ptr, bool value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setUsesPredominantAxisScrolling:value];
}

int C_NSScrollView_HorizontalScrollElasticity(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSScrollElasticity result = [nSScrollView horizontalScrollElasticity];
    return result;
}

void C_NSScrollView_SetHorizontalScrollElasticity(void* ptr, int value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setHorizontalScrollElasticity:value];
}

int C_NSScrollView_VerticalScrollElasticity(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    NSScrollElasticity result = [nSScrollView verticalScrollElasticity];
    return result;
}

void C_NSScrollView_SetVerticalScrollElasticity(void* ptr, int value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setVerticalScrollElasticity:value];
}

bool C_NSScrollView_AllowsMagnification(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    BOOL result = [nSScrollView allowsMagnification];
    return result;
}

void C_NSScrollView_SetAllowsMagnification(void* ptr, bool value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setAllowsMagnification:value];
}

double C_NSScrollView_Magnification(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    CGFloat result = [nSScrollView magnification];
    return result;
}

void C_NSScrollView_SetMagnification(void* ptr, double value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setMagnification:value];
}

double C_NSScrollView_MaxMagnification(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    CGFloat result = [nSScrollView maxMagnification];
    return result;
}

void C_NSScrollView_SetMaxMagnification(void* ptr, double value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setMaxMagnification:value];
}

double C_NSScrollView_MinMagnification(void* ptr) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    CGFloat result = [nSScrollView minMagnification];
    return result;
}

void C_NSScrollView_SetMinMagnification(void* ptr, double value) {
    NSScrollView* nSScrollView = (NSScrollView*)ptr;
    [nSScrollView setMinMagnification:value];
}
