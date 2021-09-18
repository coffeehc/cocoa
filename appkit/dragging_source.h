#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSDragging.h>

@interface NSDraggingSourceAdaptor : NSObject <NSDraggingSource>
@property (assign) uintptr_t goID;
@end

void* WrapDraggingSource(uintptr_t goID);
