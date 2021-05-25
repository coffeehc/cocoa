#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSDraggingInfoAdaptor : NSObject <NSDraggingInfo>
@property (assign) long goID;
@end

void* WrapDraggingInfo(long goID);
