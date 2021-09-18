#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_GridCell_Alloc();

void* C_NSGridCell_AllocGridCell();
void* C_NSGridCell_Init(void* ptr);
void* C_NSGridCell_NewGridCell();
void* C_NSGridCell_Autorelease(void* ptr);
void* C_NSGridCell_Retain(void* ptr);
void* C_NSGridCell_Column(void* ptr);
void* C_NSGridCell_Row(void* ptr);
void* C_NSGridCell_ContentView(void* ptr);
void C_NSGridCell_SetContentView(void* ptr, void* value);
void* C_NSGridCell_GridCell_EmptyContentView();
Array C_NSGridCell_CustomPlacementConstraints(void* ptr);
void C_NSGridCell_SetCustomPlacementConstraints(void* ptr, Array value);
int C_NSGridCell_RowAlignment(void* ptr);
void C_NSGridCell_SetRowAlignment(void* ptr, int value);
int C_NSGridCell_XPlacement(void* ptr);
void C_NSGridCell_SetXPlacement(void* ptr, int value);
int C_NSGridCell_YPlacement(void* ptr);
void C_NSGridCell_SetYPlacement(void* ptr, int value);
