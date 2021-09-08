#import <Appkit/Appkit.h>
#import "scroller.h"

void* C_Scroller_Alloc() {
    return [NSScroller alloc];
}

void* C_NSScroller_InitWithFrame(void* ptr, CGRect frameRect) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScroller* result_ = [nSScroller initWithFrame:frameRect];
    return result_;
}

void* C_NSScroller_InitWithCoder(void* ptr, void* coder) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScroller* result_ = [nSScroller initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSScroller_Init(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScroller* result_ = [nSScroller init];
    return result_;
}

void* C_NSScroller_AllocScroller() {
    NSScroller* result_ = [NSScroller alloc];
    return result_;
}

void* C_NSScroller_NewScroller() {
    NSScroller* result_ = [NSScroller new];
    return result_;
}

void* C_NSScroller_Autorelease(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScroller* result_ = [nSScroller autorelease];
    return result_;
}

void* C_NSScroller_Retain(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScroller* result_ = [nSScroller retain];
    return result_;
}

double C_NSScroller_ScrollerWidthForControlSize_ScrollerStyle(unsigned int controlSize, int scrollerStyle) {
    CGFloat result_ = [NSScroller scrollerWidthForControlSize:controlSize scrollerStyle:scrollerStyle];
    return result_;
}

CGRect C_NSScroller_RectForPart(void* ptr, unsigned int partCode) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSRect result_ = [nSScroller rectForPart:partCode];
    return result_;
}

unsigned int C_NSScroller_TestPart(void* ptr, CGPoint point) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScrollerPart result_ = [nSScroller testPart:point];
    return result_;
}

void C_NSScroller_CheckSpaceForParts(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    [nSScroller checkSpaceForParts];
}

void C_NSScroller_DrawKnobSlotInRect_Highlight(void* ptr, CGRect slotRect, bool flag) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    [nSScroller drawKnobSlotInRect:slotRect highlight:flag];
}

void C_NSScroller_DrawKnob(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    [nSScroller drawKnob];
}

void C_NSScroller_TrackKnob(void* ptr, void* event) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    [nSScroller trackKnob:(NSEvent*)event];
}

unsigned int C_NSScroller_UsableParts(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSUsableScrollerParts result_ = [nSScroller usableParts];
    return result_;
}

unsigned int C_NSScroller_HitPart(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScrollerPart result_ = [nSScroller hitPart];
    return result_;
}

int C_NSScroller_PreferredScrollerStyle() {
    NSScrollerStyle result_ = [NSScroller preferredScrollerStyle];
    return result_;
}

int C_NSScroller_ScrollerStyle(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScrollerStyle result_ = [nSScroller scrollerStyle];
    return result_;
}

void C_NSScroller_SetScrollerStyle(void* ptr, int value) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    [nSScroller setScrollerStyle:value];
}

int C_NSScroller_KnobStyle(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScrollerKnobStyle result_ = [nSScroller knobStyle];
    return result_;
}

void C_NSScroller_SetKnobStyle(void* ptr, int value) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    [nSScroller setKnobStyle:value];
}

double C_NSScroller_KnobProportion(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    CGFloat result_ = [nSScroller knobProportion];
    return result_;
}

bool C_NSScroller_CompatibleWithOverlayScrollers() {
    BOOL result_ = [NSScroller compatibleWithOverlayScrollers];
    return result_;
}
