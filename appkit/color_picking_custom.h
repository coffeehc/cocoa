#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSColorPickingCustomAdaptor : NSObject <NSColorPickingCustom>
@property (assign) uintptr_t goID;
@end

void* WrapColorPickingCustom(uintptr_t goID);
