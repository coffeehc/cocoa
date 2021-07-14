#import <Foundation/NSObject.h>
#import "object.h"
#include "_cgo_export.h"
#import <objc/runtime.h>


void Object_Dealloc(void* ptr) {
    NSObject* obj = (NSObject*)ptr;
    [obj dealloc];
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
