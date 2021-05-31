#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_ByteCountFormatter_Alloc();

void* C_NSByteCountFormatter_Init(void* ptr);
void* C_NSByteCountFormatter_StringFromMeasurement(void* ptr, void* measurement);
void* C_NSByteCountFormatter_ByteCountFormatter_StringFromMeasurement_CountStyle(void* measurement, int countStyle);
int C_NSByteCountFormatter_FormattingContext(void* ptr);
void C_NSByteCountFormatter_SetFormattingContext(void* ptr, int value);
int C_NSByteCountFormatter_CountStyle(void* ptr);
void C_NSByteCountFormatter_SetCountStyle(void* ptr, int value);
bool C_NSByteCountFormatter_AllowsNonnumericFormatting(void* ptr);
void C_NSByteCountFormatter_SetAllowsNonnumericFormatting(void* ptr, bool value);
bool C_NSByteCountFormatter_IncludesActualByteCount(void* ptr);
void C_NSByteCountFormatter_SetIncludesActualByteCount(void* ptr, bool value);
bool C_NSByteCountFormatter_IsAdaptive(void* ptr);
void C_NSByteCountFormatter_SetAdaptive(void* ptr, bool value);
unsigned int C_NSByteCountFormatter_AllowedUnits(void* ptr);
void C_NSByteCountFormatter_SetAllowedUnits(void* ptr, unsigned int value);
bool C_NSByteCountFormatter_IncludesCount(void* ptr);
void C_NSByteCountFormatter_SetIncludesCount(void* ptr, bool value);
bool C_NSByteCountFormatter_IncludesUnit(void* ptr);
void C_NSByteCountFormatter_SetIncludesUnit(void* ptr, bool value);
bool C_NSByteCountFormatter_ZeroPadsFractionDigits(void* ptr);
void C_NSByteCountFormatter_SetZeroPadsFractionDigits(void* ptr, bool value);
