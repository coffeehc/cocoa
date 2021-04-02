#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool Number_BoolValue(void* ptr);
float Number_FloatValue(void* ptr);
long Number_IntValue(void* ptr);
long Number_LongLongValue(void* ptr);
unsigned long Number_UnsignedIntValue(void* ptr);
unsigned long Number_UnsignedLongLongValue(void* ptr);

void* Number_NumberWithBool(bool value);
void* Number_NumberWithDouble(double value);
void* Number_NumberWithFloat(float value);
void* Number_NumberWithInt(long value);
void* Number_NumberWithLongLong(long value);
void* Number_NumberWithUnsignedInt(unsigned long value);
void* Number_NumberWithUnsignedLongLong(unsigned long value);
