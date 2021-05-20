#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSOutlineViewDelegateAdaptor : NSObject <NSOutlineViewDelegate>
@property (assign) long goID;
@end

void* WrapOutlineViewDelegate(long goID);
