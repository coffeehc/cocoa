#import <stdlib.h>
#import <stdint.h>
#import <stdbool.h>

void* C_NSObject_AllocObject();
void* C_NSObject_Init(void* ptr);
void* C_NSObject_NewObject();
void* Object_Retain(void* ptr);
void Object_Release(void* ptr);
void* Object_Autorelease(void* ptr);
void* C_NSObject_Copy(void* ptr);
void* C_NSObject_MutableCopy(void* ptr);
void Object_Dealloc(void* ptr);
void* Object_PerformSelector(void* ptr, void* sel_p);
void* Object_PerformSelector_WithObject(void* ptr, void* sel_p, void* param);
void* Object_PerformSelector_WithObject_WithObject(void* ptr, void* sel_p, void* param1, void* param2);