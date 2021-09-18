#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSColorPicking.h>

@interface NSColorPickingDefaultAdaptor : NSObject <NSColorPickingDefault>
@property (assign) uintptr_t goID;
@end

void* WrapColorPickingDefault(uintptr_t goID);
