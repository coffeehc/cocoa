#import <Foundation/Foundation.h>
#import "exception.h"

void* C_Exception_Alloc() {
    return [NSException alloc];
}

void* C_NSException_Init(void* ptr) {
    NSException* nSException = (NSException*)ptr;
    NSException* result = [nSException init];
    return result;
}

void C_NSException_Raise(void* ptr) {
    NSException* nSException = (NSException*)ptr;
    [nSException raise];
}

void* C_NSException_Name(void* ptr) {
    NSException* nSException = (NSException*)ptr;
    NSExceptionName result = [nSException name];
    return result;
}

void* C_NSException_Reason(void* ptr) {
    NSException* nSException = (NSException*)ptr;
    NSString* result = [nSException reason];
    return result;
}
