#import <Foundation/Foundation.h>
#import "decimal_number_handler.h"

void* C_DecimalNumberHandler_Alloc() {
    return [NSDecimalNumberHandler alloc];
}

void* C_NSDecimalNumberHandler_Init(void* ptr) {
    NSDecimalNumberHandler* nSDecimalNumberHandler = (NSDecimalNumberHandler*)ptr;
    NSDecimalNumberHandler* result_ = [nSDecimalNumberHandler init];
    return result_;
}

void* C_NSDecimalNumberHandler_DefaultDecimalNumberHandler() {
    NSDecimalNumberHandler* result_ = [NSDecimalNumberHandler defaultDecimalNumberHandler];
    return result_;
}
