#import "split_view.h"
#import <AppKit/NSSplitView.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_SplitView_Alloc() {
    return [NSSplitView alloc];
}

void* C_NSSplitView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSSplitView* result_ = [nSSplitView initWithFrame:frameRect];
    return result_;
}

void* C_NSSplitView_InitWithCoder(void* ptr, void* coder) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSSplitView* result_ = [nSSplitView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSSplitView_Init(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSSplitView* result_ = [nSSplitView init];
    return result_;
}

void* C_NSSplitView_AllocSplitView() {
    NSSplitView* result_ = [NSSplitView alloc];
    return result_;
}

void* C_NSSplitView_NewSplitView() {
    NSSplitView* result_ = [NSSplitView new];
    return result_;
}

void* C_NSSplitView_Autorelease(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSSplitView* result_ = [nSSplitView autorelease];
    return result_;
}

void* C_NSSplitView_Retain(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSSplitView* result_ = [nSSplitView retain];
    return result_;
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
    BOOL result_ = [nSSplitView isSubviewCollapsed:(NSView*)subview];
    return result_;
}

float C_NSSplitView_HoldingPriorityForSubviewAtIndex(void* ptr, int subviewIndex) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSLayoutPriority result_ = [nSSplitView holdingPriorityForSubviewAtIndex:subviewIndex];
    return result_;
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
    CGFloat result_ = [nSSplitView minPossiblePositionOfDividerAtIndex:dividerIndex];
    return result_;
}

double C_NSSplitView_MaxPossiblePositionOfDividerAtIndex(void* ptr, int dividerIndex) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    CGFloat result_ = [nSSplitView maxPossiblePositionOfDividerAtIndex:dividerIndex];
    return result_;
}

void C_NSSplitView_SetPosition_OfDividerAtIndex(void* ptr, double position, int dividerIndex) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView setPosition:position ofDividerAtIndex:dividerIndex];
}

void* C_NSSplitView_Delegate(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    id result_ = [nSSplitView delegate];
    return result_;
}

void C_NSSplitView_SetDelegate(void* ptr, void* value) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView setDelegate:(id)value];
}

bool C_NSSplitView_ArrangesAllSubviews(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    BOOL result_ = [nSSplitView arrangesAllSubviews];
    return result_;
}

void C_NSSplitView_SetArrangesAllSubviews(void* ptr, bool value) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView setArrangesAllSubviews:value];
}

Array C_NSSplitView_ArrangedSubviews(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSArray* result_ = [nSSplitView arrangedSubviews];
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

bool C_NSSplitView_IsVertical(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    BOOL result_ = [nSSplitView isVertical];
    return result_;
}

void C_NSSplitView_SetVertical(void* ptr, bool value) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView setVertical:value];
}

int C_NSSplitView_DividerStyle(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSSplitViewDividerStyle result_ = [nSSplitView dividerStyle];
    return result_;
}

void C_NSSplitView_SetDividerStyle(void* ptr, int value) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView setDividerStyle:value];
}

void* C_NSSplitView_DividerColor(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSColor* result_ = [nSSplitView dividerColor];
    return result_;
}

double C_NSSplitView_DividerThickness(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    CGFloat result_ = [nSSplitView dividerThickness];
    return result_;
}

void* C_NSSplitView_AutosaveName(void* ptr) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    NSSplitViewAutosaveName result_ = [nSSplitView autosaveName];
    return result_;
}

void C_NSSplitView_SetAutosaveName(void* ptr, void* value) {
    NSSplitView* nSSplitView = (NSSplitView*)ptr;
    [nSSplitView setAutosaveName:(NSString*)value];
}
