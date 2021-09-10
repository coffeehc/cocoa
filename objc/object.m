#import <objc/NSObject.h>
#import "object.h"
#include "_cgo_export.h"
#import <objc/runtime.h>

void* C_NSObject_AllocObject() {
    NSObject* result_ = [NSObject alloc];
    return result_;
}

void* C_NSObject_Init(void* ptr) {
    NSObject* nSObject = (NSObject*)ptr;
    NSObject* result_ = [nSObject init];
    return result_;
}

void* C_NSObject_NewObject() {
    NSObject* result_ = [NSObject new];
    return result_;
}

void* Object_Retain(void* ptr) {
    NSObject* obj = (NSObject*)ptr;
    return [obj retain];
}

void Object_Release(void* ptr) {
    NSObject* obj = (NSObject*)ptr;
    [obj release];
}

void* Object_Autorelease(void* ptr) {
    NSObject* obj = (NSObject*)ptr;
    return [obj retain];
}

void* C_NSObject_Copy(void* ptr) {
    NSObject* nSObject = (NSObject*)ptr;
    id result_ = [nSObject copy];
    return result_;
}

void* C_NSObject_MutableCopy(void* ptr) {
    NSObject* nSObject = (NSObject*)ptr;
    id result_ = [nSObject mutableCopy];
    return result_;
}

void Object_Dealloc(void* ptr) {
    NSObject* obj = (NSObject*)ptr;
    [obj dealloc];
}

void* Object_PerformSelector(void* ptr, void* sel_p) {
    NSObject* obj = (NSObject*)ptr;
    return [obj performSelector:(SEL)sel_p];
}

void* Object_PerformSelector_WithObject(void* ptr, void* sel_p, void* param) {
    NSObject* obj = (NSObject*)ptr;
    return [obj performSelector:(SEL)sel_p withObject:(NSObject*)param];
}

void* Object_PerformSelector_WithObject_WithObject(void* ptr, void* sel_p, void* param1, void* param2) {
    NSObject* obj = (NSObject*)ptr;
    return [obj performSelector:(SEL)sel_p withObject:(NSObject*)param1  withObject:(NSObject*)param2];
}
