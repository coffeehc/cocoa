#import <stdlib.h>
#import <stdint.h>

void Object_Dealloc(void* ptr);
void Dealloc_AddHook(void* ptr, uintptr_t hookPtr);