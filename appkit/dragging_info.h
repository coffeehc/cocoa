#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSDragging.h>

@interface NSDraggingInfoAdaptor : NSObject <NSDraggingInfo>
@property (assign) uintptr_t goID;
@end

void* WrapDraggingInfo(uintptr_t goID);
