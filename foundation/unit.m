#import <Foundation/Foundation.h>
#import "unit.h"

void* C_Unit_Alloc() {
    return [NSUnit alloc];
}

void* C_NSUnit_InitWithSymbol(void* ptr, void* symbol) {
    NSUnit* nSUnit = (NSUnit*)ptr;
    NSUnit* result_ = [nSUnit initWithSymbol:(NSString*)symbol];
    return result_;
}

void* C_NSUnit_Symbol(void* ptr) {
    NSUnit* nSUnit = (NSUnit*)ptr;
    NSString* result_ = [nSUnit symbol];
    return result_;
}
