#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Number_Alloc();

void* C_NSNumber_InitWithCoder(void* ptr, void* coder);
void* C_NSNumber_Init(void* ptr);
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
bool C_NSNumber_IsEqualToNumber(void* ptr, void* number);
bool C_NSNumber_BoolValue(void* ptr);
NSDecimal C_NSNumber_DecimalValue(void* ptr);
double C_NSNumber_DoubleValue(void* ptr);
float C_NSNumber_FloatValue(void* ptr);
int C_NSNumber_IntegerValue(void* ptr);
unsigned int C_NSNumber_UnsignedIntegerValue(void* ptr);
void* C_NSNumber_StringValue(void* ptr);
