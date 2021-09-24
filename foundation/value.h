#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Value_Alloc();

void* C_NSValue_InitWithCoder(void* ptr, void* coder);
void* C_NSValue_AllocValue();
void* C_NSValue_Init(void* ptr);
void* C_NSValue_NewValue();
void* C_NSValue_Autorelease(void* ptr);
void* C_NSValue_Retain(void* ptr);
void* C_NSValue_ValueWithNonretainedObject(void* anObject);
void* C_NSValue_ValueWithRange(NSRange _range);
void* C_NSValue_ValueWithPoint(CGPoint point);
void* C_NSValue_ValueWithSize(CGSize size);
void* C_NSValue_ValueWithRect(CGRect rect);
bool C_NSValue_IsEqualToValue(void* ptr, void* value);
void* C_NSValue_ValueWithEdgeInsets(NSEdgeInsets insets);
void* C_NSValue_NonretainedObjectValue(void* ptr);
NSRange C_NSValue_RangeValue(void* ptr);
CGPoint C_NSValue_PointValue(void* ptr);
CGSize C_NSValue_SizeValue(void* ptr);
CGRect C_NSValue_RectValue(void* ptr);
NSEdgeInsets C_NSValue_EdgeInsetsValue(void* ptr);
