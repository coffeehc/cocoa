#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSTouchBarDelegateAdaptor : NSObject <NSTouchBarDelegate>
@property (assign) long goID;
@end

void* WrapTouchBarDelegate(long goID);
