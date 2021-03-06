#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_TableHeaderView_Alloc();

void* C_NSTableHeaderView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSTableHeaderView_InitWithCoder(void* ptr, void* coder);
void* C_NSTableHeaderView_Init(void* ptr);
void* C_NSTableHeaderView_AllocTableHeaderView();
void* C_NSTableHeaderView_NewTableHeaderView();
void* C_NSTableHeaderView_Autorelease(void* ptr);
void* C_NSTableHeaderView_Retain(void* ptr);
int C_NSTableHeaderView_ColumnAtPoint(void* ptr, CGPoint point);
CGRect C_NSTableHeaderView_HeaderRectOfColumn(void* ptr, int column);
void* C_NSTableHeaderView_TableView(void* ptr);
void C_NSTableHeaderView_SetTableView(void* ptr, void* value);
int C_NSTableHeaderView_DraggedColumn(void* ptr);
double C_NSTableHeaderView_DraggedDistance(void* ptr);
int C_NSTableHeaderView_ResizedColumn(void* ptr);
