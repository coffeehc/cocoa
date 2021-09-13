#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_IndexPath_Alloc();

void* C_NSIndexPath_IndexPathWithIndex(unsigned int index);
void* C_NSIndexPath_InitWithIndex(void* ptr, unsigned int index);
void* C_NSIndexPath_AllocIndexPath();
void* C_NSIndexPath_Init(void* ptr);
void* C_NSIndexPath_NewIndexPath();
void* C_NSIndexPath_Autorelease(void* ptr);
void* C_NSIndexPath_Retain(void* ptr);
void* C_NSIndexPath_IndexPathByAddingIndex(void* ptr, unsigned int index);
void* C_NSIndexPath_IndexPathByRemovingLastIndex(void* ptr);
int C_NSIndexPath_Compare(void* ptr, void* otherObject);
unsigned int C_NSIndexPath_IndexAtPosition(void* ptr, unsigned int position);
unsigned int C_NSIndexPath_Length(void* ptr);
