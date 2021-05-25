#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSDraggingSourceAdaptor : NSObject <NSDraggingSource>
@property (assign) long goID;
@end

void* WrapDraggingSource(long goID);
