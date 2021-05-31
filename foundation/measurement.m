#import <Foundation/Foundation.h>
#import "measurement.h"

void* C_Measurement_Alloc() {
    return [NSMeasurement alloc];
}

bool C_NSMeasurement_CanBeConvertedToUnit(void* ptr, void* unit) {
    NSMeasurement* nSMeasurement = (NSMeasurement*)ptr;
    BOOL result_ = [nSMeasurement canBeConvertedToUnit:(NSUnit*)unit];
    return result_;
}

void* C_NSMeasurement_MeasurementByConvertingToUnit(void* ptr, void* unit) {
    NSMeasurement* nSMeasurement = (NSMeasurement*)ptr;
    NSMeasurement* result_ = [nSMeasurement measurementByConvertingToUnit:(NSUnit*)unit];
    return result_;
}

void* C_NSMeasurement_MeasurementByAddingMeasurement(void* ptr, void* measurement) {
    NSMeasurement* nSMeasurement = (NSMeasurement*)ptr;
    NSMeasurement* result_ = [nSMeasurement measurementByAddingMeasurement:(NSMeasurement*)measurement];
    return result_;
}

void* C_NSMeasurement_MeasurementBySubtractingMeasurement(void* ptr, void* measurement) {
    NSMeasurement* nSMeasurement = (NSMeasurement*)ptr;
    NSMeasurement* result_ = [nSMeasurement measurementBySubtractingMeasurement:(NSMeasurement*)measurement];
    return result_;
}

double C_NSMeasurement_DoubleValue(void* ptr) {
    NSMeasurement* nSMeasurement = (NSMeasurement*)ptr;
    double result_ = [nSMeasurement doubleValue];
    return result_;
}
