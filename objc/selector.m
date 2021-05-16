#import "selector.h"
#import <objc/runtime.h>


void* Selector_SEL_RegisterName(const char* name) {
    return (void*)sel_registerName(name);
}

const char* Selector_SEL_GetName(void* ptr) {
    return sel_getName((SEL)ptr);
}