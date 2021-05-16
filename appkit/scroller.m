#import <Appkit/Appkit.h>
#import "scroller.h"

void* C_Scroller_Alloc() {
    return [NSScroller alloc];
}

void* C_NSScroller_InitWithFrame(void* ptr, CGRect frameRect) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScroller* result = [nSScroller initWithFrame:frameRect];
    return result;
}

void* C_NSScroller_InitWithCoder(void* ptr, void* coder) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScroller* result = [nSScroller initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSScroller_Init(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScroller* result = [nSScroller init];
    return result;
}

double C_NSScroller_ScrollerWidthForControlSize_ScrollerStyle(unsigned int controlSize, int scrollerStyle) {
    CGFloat result = [NSScroller scrollerWidthForControlSize:controlSize scrollerStyle:scrollerStyle];
    return result;
}

CGRect C_NSScroller_RectForPart(void* ptr, unsigned int partCode) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSRect result = [nSScroller rectForPart:partCode];
    return result;
}

unsigned int C_NSScroller_TestPart(void* ptr, CGPoint point) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScrollerPart result = [nSScroller testPart:point];
    return result;
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
    NSUsableScrollerParts result = [nSScroller usableParts];
    return result;
}

unsigned int C_NSScroller_HitPart(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScrollerPart result = [nSScroller hitPart];
    return result;
}

int C_NSScroller_Scroller_PreferredScrollerStyle() {
    NSScrollerStyle result = [NSScroller preferredScrollerStyle];
    return result;
}

int C_NSScroller_ScrollerStyle(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScrollerStyle result = [nSScroller scrollerStyle];
    return result;
}

void C_NSScroller_SetScrollerStyle(void* ptr, int value) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    [nSScroller setScrollerStyle:value];
}

int C_NSScroller_KnobStyle(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    NSScrollerKnobStyle result = [nSScroller knobStyle];
    return result;
}

void C_NSScroller_SetKnobStyle(void* ptr, int value) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    [nSScroller setKnobStyle:value];
}

double C_NSScroller_KnobProportion(void* ptr) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    CGFloat result = [nSScroller knobProportion];
    return result;
}

void C_NSScroller_SetKnobProportion(void* ptr, double value) {
    NSScroller* nSScroller = (NSScroller*)ptr;
    [nSScroller setKnobProportion:value];
}

bool C_NSScroller_Scroller_CompatibleWithOverlayScrollers() {
    BOOL result = [NSScroller compatibleWithOverlayScrollers];
    return result;
}
