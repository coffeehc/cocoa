#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSAnimationDelegateAdaptor : NSObject <NSAnimationDelegate>
@property (assign) long goID;
@end

void* WrapAnimationDelegate(long goID);
