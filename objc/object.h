#import <stdlib.h>
#import <stdint.h>

void* Object_Retain(void* ptr);
void Object_Release(void* ptr);
void* Object_Autorelease(void* ptr);
void Object_Dealloc(void* ptr);
void* Object_PerformSelector(void* ptr, void* sel_p);
void* Object_PerformSelector_WithObject(void* ptr, void* sel_p, void* param);
void* Object_PerformSelector_WithObject_WithObject(void* ptr, void* sel_p, void* param1, void* param2);
void Dealloc_AddHook(void* ptr, uintptr_t hookPtr);