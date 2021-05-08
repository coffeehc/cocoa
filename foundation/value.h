#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* C_Value_Alloc();

void* C_NSValue_InitWithCoder(void* ptr, void* coder);
void* C_NSValue_Init(void* ptr);
void* C_NSValue_ValueWithNonretainedObject(void* anObject);
bool C_NSValue_IsEqualToValue(void* ptr, void* value);
void* C_NSValue_NonretainedObjectValue(void* ptr);
