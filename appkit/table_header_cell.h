#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_TableHeaderCell_Alloc();

void* C_NSTableHeaderCell_InitTextCell(void* ptr, void* _string);
void* C_NSTableHeaderCell_InitWithCoder(void* ptr, void* coder);
void* C_NSTableHeaderCell_Init(void* ptr);
void* C_NSTableHeaderCell_AllocTableHeaderCell();
void* C_NSTableHeaderCell_NewTableHeaderCell();
void* C_NSTableHeaderCell_Autorelease(void* ptr);
void* C_NSTableHeaderCell_Retain(void* ptr);
void C_NSTableHeaderCell_DrawSortIndicatorWithFrame_InView_Ascending_Priority(void* ptr, CGRect cellFrame, void* controlView, bool ascending, int priority);
CGRect C_NSTableHeaderCell_SortIndicatorRectForBounds(void* ptr, CGRect rect);
