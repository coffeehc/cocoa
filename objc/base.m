#import <dispatch/dispatch.h>
#import "base.h"
#include "_cgo_export.h"
#import <objc/runtime.h>


void Dispatch_MainQueueAsync(uintptr_t ptr) {
    dispatch_async(dispatch_get_main_queue(), ^{
            runTask(ptr);
    });
}

void Run_WithAutoreleasePool(uintptr_t ptr) {
    @autoreleasepool {
        runTask(ptr);
    }
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
