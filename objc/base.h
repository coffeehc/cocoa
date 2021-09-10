#import <stdlib.h>
#import <stdint.h>

void Dispatch_MainQueueAsync(uintptr_t ptr);

void Run_WithAutoreleasePool(uintptr_t ptr);

void Dealloc_AddHook(void* ptr, uintptr_t hookPtr);