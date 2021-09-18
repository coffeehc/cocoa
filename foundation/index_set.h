#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_IndexSet_Alloc();

void* C_NSIndexSet_IndexSet_();
void* C_NSIndexSet_IndexSetWithIndex(unsigned int value);
void* C_NSIndexSet_IndexSetWithIndexesInRange(NSRange _range);
void* C_NSIndexSet_InitWithIndex(void* ptr, unsigned int value);
void* C_NSIndexSet_InitWithIndexesInRange(void* ptr, NSRange _range);
void* C_NSIndexSet_InitWithIndexSet(void* ptr, void* indexSet);
void* C_NSIndexSet_AllocIndexSet();
void* C_NSIndexSet_Init(void* ptr);
void* C_NSIndexSet_NewIndexSet();
void* C_NSIndexSet_Autorelease(void* ptr);
void* C_NSIndexSet_Retain(void* ptr);
bool C_NSIndexSet_ContainsIndex(void* ptr, unsigned int value);
bool C_NSIndexSet_ContainsIndexes(void* ptr, void* indexSet);
bool C_NSIndexSet_ContainsIndexesInRange(void* ptr, NSRange _range);
bool C_NSIndexSet_IntersectsIndexesInRange(void* ptr, NSRange _range);
unsigned int C_NSIndexSet_CountOfIndexesInRange(void* ptr, NSRange _range);
bool C_NSIndexSet_IsEqualToIndexSet(void* ptr, void* indexSet);
unsigned int C_NSIndexSet_IndexLessThanIndex(void* ptr, unsigned int value);
unsigned int C_NSIndexSet_IndexLessThanOrEqualToIndex(void* ptr, unsigned int value);
unsigned int C_NSIndexSet_IndexGreaterThanOrEqualToIndex(void* ptr, unsigned int value);
unsigned int C_NSIndexSet_IndexGreaterThanIndex(void* ptr, unsigned int value);
unsigned int C_NSIndexSet_Count(void* ptr);
unsigned int C_NSIndexSet_FirstIndex(void* ptr);
unsigned int C_NSIndexSet_LastIndex(void* ptr);
