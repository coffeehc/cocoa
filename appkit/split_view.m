#import <Appkit/Appkit.h>
#import "split_view.h"

void* C_SplitView_Alloc() {
    return [NSSplitView alloc];
}

void* C_NSSplitView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSSplitView* result = [nSSplitView initWithFrame:frameRect];
    return result;
}

void* C_NSSplitView_InitWithCoder(void* ptr, void* coder) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSSplitView* result = [nSSplitView initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSSplitView_Init(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSSplitView* result = [nSSplitView init];
    return result;
}

void C_NSSplitView_AddArrangedSubview(void* ptr, void* view) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView addArrangedSubview:(NSView*)view];
}

void C_NSSplitView_InsertArrangedSubview_AtIndex(void* ptr, void* view, int index) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView insertArrangedSubview:(NSView*)view atIndex:index];
}

void C_NSSplitView_RemoveArrangedSubview(void* ptr, void* view) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView removeArrangedSubview:(NSView*)view];
}

void C_NSSplitView_AdjustSubviews(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView adjustSubviews];
}

bool C_NSSplitView_IsSubviewCollapsed(void* ptr, void* subview) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    BOOL result = [nSSplitView isSubviewCollapsed:(NSView*)subview];
    return result;
}

float C_NSSplitView_HoldingPriorityForSubviewAtIndex(void* ptr, int subviewIndex) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSLayoutPriority result = [nSSplitView holdingPriorityForSubviewAtIndex:subviewIndex];
    return result;
}

void C_NSSplitView_SetHoldingPriority_ForSubviewAtIndex(void* ptr, float priority, int subviewIndex) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView setHoldingPriority:priority forSubviewAtIndex:subviewIndex];
}

void C_NSSplitView_DrawDividerInRect(void* ptr, CGRect rect) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView drawDividerInRect:rect];
}

double C_NSSplitView_MinPossiblePositionOfDividerAtIndex(void* ptr, int dividerIndex) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    CGFloat result = [nSSplitView minPossiblePositionOfDividerAtIndex:dividerIndex];
    return result;
}

double C_NSSplitView_MaxPossiblePositionOfDividerAtIndex(void* ptr, int dividerIndex) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    CGFloat result = [nSSplitView maxPossiblePositionOfDividerAtIndex:dividerIndex];
    return result;
}

void C_NSSplitView_SetPosition_OfDividerAtIndex(void* ptr, double position, int dividerIndex) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView setPosition:position ofDividerAtIndex:dividerIndex];
}

bool C_NSSplitView_ArrangesAllSubviews(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    BOOL result = [nSSplitView arrangesAllSubviews];
    return result;
}

void C_NSSplitView_SetArrangesAllSubviews(void* ptr, bool value) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView setArrangesAllSubviews:value];
}

Array C_NSSplitView_ArrangedSubviews(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSArray* result = [nSSplitView arrangedSubviews];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

bool C_NSSplitView_IsVertical(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    BOOL result = [nSSplitView isVertical];
    return result;
}

void C_NSSplitView_SetVertical(void* ptr, bool value) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView setVertical:value];
}

int C_NSSplitView_DividerStyle(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSSplitViewDividerStyle result = [nSSplitView dividerStyle];
    return result;
}

void C_NSSplitView_SetDividerStyle(void* ptr, int value) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView setDividerStyle:value];
}

void* C_NSSplitView_DividerColor(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSColor* result = [nSSplitView dividerColor];
    return result;
}

double C_NSSplitView_DividerThickness(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    CGFloat result = [nSSplitView dividerThickness];
    return result;
}

void* C_NSSplitView_AutosaveName(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSSplitViewAutosaveName result = [nSSplitView autosaveName];
    return result;
}

void C_NSSplitView_SetAutosaveName(void* ptr, void* value) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView setAutosaveName:(NSString*)value];
}
