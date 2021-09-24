#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_PathCell_Alloc();

void* C_NSPathCell_InitImageCell(void* ptr, void* image);
void* C_NSPathCell_InitTextCell(void* ptr, void* _string);
void* C_NSPathCell_Init(void* ptr);
void* C_NSPathCell_InitWithCoder(void* ptr, void* coder);
void* C_NSPathCell_AllocPathCell();
void* C_NSPathCell_NewPathCell();
void* C_NSPathCell_Autorelease(void* ptr);
void* C_NSPathCell_Retain(void* ptr);
void C_NSPathCell_MouseEntered_WithFrame_InView(void* ptr, void* event, CGRect frame, void* view);
void C_NSPathCell_MouseExited_WithFrame_InView(void* ptr, void* event, CGRect frame, void* view);
CGRect C_NSPathCell_RectOfPathComponentCell_WithFrame_InView(void* ptr, void* cell, CGRect frame, void* view);
void* C_NSPathCell_PathComponentCellAtPoint_WithFrame_InView(void* ptr, CGPoint point, CGRect frame, void* view);
Array C_NSPathCell_AllowedTypes(void* ptr);
void C_NSPathCell_SetAllowedTypes(void* ptr, Array value);
int C_NSPathCell_PathStyle(void* ptr);
void C_NSPathCell_SetPathStyle(void* ptr, int value);
void* C_NSPathCell_PlaceholderAttributedString(void* ptr);
void C_NSPathCell_SetPlaceholderAttributedString(void* ptr, void* value);
void* C_NSPathCell_PlaceholderString(void* ptr);
void C_NSPathCell_SetPlaceholderString(void* ptr, void* value);
void* C_NSPathCell_BackgroundColor(void* ptr);
void C_NSPathCell_SetBackgroundColor(void* ptr, void* value);
void* C_NSPathCell_ClickedPathComponentCell(void* ptr);
Array C_NSPathCell_PathComponentCells(void* ptr);
void C_NSPathCell_SetPathComponentCells(void* ptr, Array value);
void* C_NSPathCell_DoubleAction(void* ptr);
void C_NSPathCell_SetDoubleAction(void* ptr, void* value);
void* C_NSPathCell_URL(void* ptr);
void C_NSPathCell_SetURL(void* ptr, void* value);
void* C_NSPathCell_Delegate(void* ptr);
void C_NSPathCell_SetDelegate(void* ptr, void* value);

void* WrapPathCellDelegate(uintptr_t goID);
