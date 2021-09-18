#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSDragging.h>

@interface NSSpringLoadingDestinationAdaptor : NSObject <NSSpringLoadingDestination>
@property (assign) uintptr_t goID;
@end

void* WrapSpringLoadingDestination(uintptr_t goID);
