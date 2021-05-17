#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSTabViewDelegateAdaptor : NSObject <NSTabViewDelegate>
@property (assign) long goID;
@end

void* WrapTabViewDelegate(long goID);
