#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSColorPickingCustomAdaptor : NSObject <NSColorPickingCustom>
@property (assign) long goID;
@end

void* WrapColorPickingCustom(long goID);
