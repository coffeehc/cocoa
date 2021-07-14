#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSDraggingDestinationAdaptor : NSObject <NSDraggingDestination>
@property (assign) uintptr_t goID;
@end

void* WrapDraggingDestination(uintptr_t goID);
