#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSColorPickingDefaultAdaptor : NSObject <NSColorPickingDefault>
@property (assign) long goID;
@end

void* WrapColorPickingDefault(long goID);
