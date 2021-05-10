#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Value_Alloc();

void* C_NSValue_InitWithCoder(void* ptr, void* coder);
void* C_NSValue_Init(void* ptr);
void* C_NSValue_ValueWithNonretainedObject(void* anObject);
void* C_NSValue_ValueWithRange(NSRange _range);
bool C_NSValue_IsEqualToValue(void* ptr, void* value);
void* C_NSValue_ValueWithEdgeInsets(NSEdgeInsets insets);
void* C_NSValue_NonretainedObjectValue(void* ptr);
NSRange C_NSValue_RangeValue(void* ptr);
NSEdgeInsets C_NSValue_EdgeInsetsValue(void* ptr);
