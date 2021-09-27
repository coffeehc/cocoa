#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* WrapDraggingSource(uintptr_t goID);

void* WrapDraggingDestination(uintptr_t goID);

void* WrapSpringLoadingDestination(uintptr_t goID);

void* WrapDraggingInfo(uintptr_t goID);
