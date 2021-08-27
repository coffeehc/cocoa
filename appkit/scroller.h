#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Scroller_Alloc();

void* C_NSScroller_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSScroller_InitWithCoder(void* ptr, void* coder);
void* C_NSScroller_Init(void* ptr);
double C_NSScroller_ScrollerWidthForControlSize_ScrollerStyle(unsigned int controlSize, int scrollerStyle);
CGRect C_NSScroller_RectForPart(void* ptr, unsigned int partCode);
unsigned int C_NSScroller_TestPart(void* ptr, CGPoint point);
void C_NSScroller_CheckSpaceForParts(void* ptr);
void C_NSScroller_DrawKnobSlotInRect_Highlight(void* ptr, CGRect slotRect, bool flag);
void C_NSScroller_DrawKnob(void* ptr);
void C_NSScroller_TrackKnob(void* ptr, void* event);
unsigned int C_NSScroller_UsableParts(void* ptr);
unsigned int C_NSScroller_HitPart(void* ptr);
int C_NSScroller_PreferredScrollerStyle();
int C_NSScroller_ScrollerStyle(void* ptr);
void C_NSScroller_SetScrollerStyle(void* ptr, int value);
int C_NSScroller_KnobStyle(void* ptr);
void C_NSScroller_SetKnobStyle(void* ptr, int value);
double C_NSScroller_KnobProportion(void* ptr);
bool C_NSScroller_CompatibleWithOverlayScrollers();
