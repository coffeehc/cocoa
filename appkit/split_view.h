#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_SplitView_Alloc();

void* C_NSSplitView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSSplitView_InitWithCoder(void* ptr, void* coder);
void* C_NSSplitView_Init(void* ptr);
void* C_NSSplitView_AllocSplitView();
void* C_NSSplitView_NewSplitView();
void* C_NSSplitView_Autorelease(void* ptr);
void* C_NSSplitView_Retain(void* ptr);
void C_NSSplitView_AddArrangedSubview(void* ptr, void* view);
void C_NSSplitView_InsertArrangedSubview_AtIndex(void* ptr, void* view, int index);
void C_NSSplitView_RemoveArrangedSubview(void* ptr, void* view);
void C_NSSplitView_AdjustSubviews(void* ptr);
bool C_NSSplitView_IsSubviewCollapsed(void* ptr, void* subview);
float C_NSSplitView_HoldingPriorityForSubviewAtIndex(void* ptr, int subviewIndex);
void C_NSSplitView_SetHoldingPriority_ForSubviewAtIndex(void* ptr, float priority, int subviewIndex);
void C_NSSplitView_DrawDividerInRect(void* ptr, CGRect rect);
double C_NSSplitView_MinPossiblePositionOfDividerAtIndex(void* ptr, int dividerIndex);
double C_NSSplitView_MaxPossiblePositionOfDividerAtIndex(void* ptr, int dividerIndex);
void C_NSSplitView_SetPosition_OfDividerAtIndex(void* ptr, double position, int dividerIndex);
void* C_NSSplitView_Delegate(void* ptr);
void C_NSSplitView_SetDelegate(void* ptr, void* value);
bool C_NSSplitView_ArrangesAllSubviews(void* ptr);
void C_NSSplitView_SetArrangesAllSubviews(void* ptr, bool value);
Array C_NSSplitView_ArrangedSubviews(void* ptr);
bool C_NSSplitView_IsVertical(void* ptr);
void C_NSSplitView_SetVertical(void* ptr, bool value);
int C_NSSplitView_DividerStyle(void* ptr);
void C_NSSplitView_SetDividerStyle(void* ptr, int value);
void* C_NSSplitView_DividerColor(void* ptr);
double C_NSSplitView_DividerThickness(void* ptr);
void* C_NSSplitView_AutosaveName(void* ptr);
void C_NSSplitView_SetAutosaveName(void* ptr, void* value);
