#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_TextList_Alloc();

void* C_NSTextList_InitWithMarkerFormat_Options(void* ptr, void* format, unsigned int mask);
void* C_NSTextList_AllocTextList();
void* C_NSTextList_Init(void* ptr);
void* C_NSTextList_NewTextList();
void* C_NSTextList_Autorelease(void* ptr);
void* C_NSTextList_Retain(void* ptr);
void* C_NSTextList_MarkerForItemNumber(void* ptr, int itemNum);
void* C_NSTextList_MarkerFormat(void* ptr);
unsigned int C_NSTextList_ListOptions(void* ptr);
int C_NSTextList_StartingItemNumber(void* ptr);
void C_NSTextList_SetStartingItemNumber(void* ptr, int value);
