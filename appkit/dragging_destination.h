#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSDraggingDestinationAdaptor : NSObject <NSDraggingDestination>
@property (assign) long goID;
@end

void* WrapDraggingDestination(long goID);
