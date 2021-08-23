#import <Foundation/NSObject.h>
#import "object.h"
#include "_cgo_export.h"
#import <objc/runtime.h>


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

@interface Parasite : NSObject
@property(nonatomic, assign) uintptr_t hookPtr;
@end

@implementation Parasite
- (void)dealloc {
    runDeallocTask(self.hookPtr);
    [super dealloc];
}
@end

static void *kDeallocHookAssociation = &kDeallocHookAssociation;

void Dealloc_AddHook(void* ptr, uintptr_t hookPtr) {
    Parasite *parasite = [Parasite alloc];
    parasite.hookPtr = hookPtr;
    objc_setAssociatedObject((NSObject*)ptr, &kDeallocHookAssociation, parasite, OBJC_ASSOCIATION_RETAIN_NONATOMIC);
}
