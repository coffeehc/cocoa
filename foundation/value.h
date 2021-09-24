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

void* C_Number_Alloc();

void* C_NSNumber_InitWithCoder(void* ptr, void* coder);
void* C_NSNumber_AllocNumber();
void* C_NSNumber_Init(void* ptr);
void* C_NSNumber_NewNumber();
void* C_NSNumber_Autorelease(void* ptr);
void* C_NSNumber_Retain(void* ptr);
void* C_NSNumber_NumberWithBool(bool value);
void* C_NSNumber_NumberWithDouble(double value);
void* C_NSNumber_NumberWithFloat(float value);
void* C_NSNumber_NumberWithInteger(int value);
void* C_NSNumber_NumberWithUnsignedInteger(unsigned int value);
void* C_NSNumber_InitWithBool(void* ptr, bool value);
void* C_NSNumber_InitWithDouble(void* ptr, double value);
void* C_NSNumber_InitWithFloat(void* ptr, float value);
void* C_NSNumber_InitWithInteger(void* ptr, int value);
void* C_NSNumber_InitWithUnsignedInteger(void* ptr, unsigned int value);
void* C_NSNumber_DescriptionWithLocale(void* ptr, void* locale);
int C_NSNumber_Compare(void* ptr, void* otherNumber);
bool C_NSNumber_IsEqualToNumber(void* ptr, void* number);
bool C_NSNumber_BoolValue(void* ptr);
double C_NSNumber_DoubleValue(void* ptr);
float C_NSNumber_FloatValue(void* ptr);
int C_NSNumber_IntegerValue(void* ptr);
unsigned int C_NSNumber_UnsignedIntegerValue(void* ptr);
void* C_NSNumber_StringValue(void* ptr);
