#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_GridColumn_Alloc();

void* C_NSGridColumn_AllocGridColumn();
void* C_NSGridColumn_Init(void* ptr);
void* C_NSGridColumn_NewGridColumn();
void* C_NSGridColumn_Autorelease(void* ptr);
void* C_NSGridColumn_Retain(void* ptr);
void* C_NSGridColumn_CellAtIndex(void* ptr, int index);
void C_NSGridColumn_MergeCellsInRange(void* ptr, NSRange _range);
void* C_NSGridColumn_GridView(void* ptr);
bool C_NSGridColumn_IsHidden(void* ptr);
void C_NSGridColumn_SetHidden(void* ptr, bool value);
double C_NSGridColumn_LeadingPadding(void* ptr);
void C_NSGridColumn_SetLeadingPadding(void* ptr, double value);
int C_NSGridColumn_NumberOfCells(void* ptr);
double C_NSGridColumn_TrailingPadding(void* ptr);
void C_NSGridColumn_SetTrailingPadding(void* ptr, double value);
double C_NSGridColumn_Width(void* ptr);
void C_NSGridColumn_SetWidth(void* ptr, double value);
int C_NSGridColumn_XPlacement(void* ptr);
void C_NSGridColumn_SetXPlacement(void* ptr, int value);
