#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_Measurement_Alloc();

void* C_NSMeasurement_AllocMeasurement();
void* C_NSMeasurement_Autorelease(void* ptr);
void* C_NSMeasurement_Retain(void* ptr);
bool C_NSMeasurement_CanBeConvertedToUnit(void* ptr, void* unit);
void* C_NSMeasurement_MeasurementByConvertingToUnit(void* ptr, void* unit);
void* C_NSMeasurement_MeasurementByAddingMeasurement(void* ptr, void* measurement);
void* C_NSMeasurement_MeasurementBySubtractingMeasurement(void* ptr, void* measurement);
double C_NSMeasurement_DoubleValue(void* ptr);
