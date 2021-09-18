#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_MeasurementFormatter_Alloc();

void* C_NSMeasurementFormatter_AllocMeasurementFormatter();
void* C_NSMeasurementFormatter_Init(void* ptr);
void* C_NSMeasurementFormatter_NewMeasurementFormatter();
void* C_NSMeasurementFormatter_Autorelease(void* ptr);
void* C_NSMeasurementFormatter_Retain(void* ptr);
void* C_NSMeasurementFormatter_StringFromMeasurement(void* ptr, void* measurement);
void* C_NSMeasurementFormatter_StringFromUnit(void* ptr, void* unit);
unsigned int C_NSMeasurementFormatter_UnitOptions(void* ptr);
void C_NSMeasurementFormatter_SetUnitOptions(void* ptr, unsigned int value);
int C_NSMeasurementFormatter_UnitStyle(void* ptr);
void C_NSMeasurementFormatter_SetUnitStyle(void* ptr, int value);
void* C_NSMeasurementFormatter_Locale(void* ptr);
void C_NSMeasurementFormatter_SetLocale(void* ptr, void* value);
void* C_NSMeasurementFormatter_NumberFormatter(void* ptr);
void C_NSMeasurementFormatter_SetNumberFormatter(void* ptr, void* value);
