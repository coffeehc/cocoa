#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSWindowDelegateAdaptor : NSObject <NSWindowDelegate>
@property (assign) long goID;
@end

void* WrapWindowDelegate(long goID);
