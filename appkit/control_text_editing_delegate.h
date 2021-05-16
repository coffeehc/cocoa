#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSControlTextEditingDelegateAdaptor : NSObject <NSControlTextEditingDelegate>
@property (assign) long goID;
@end

void* WrapControlTextEditingDelegate(long goID);
