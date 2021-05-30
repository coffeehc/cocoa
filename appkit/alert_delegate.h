#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSAlertDelegateAdaptor : NSObject <NSAlertDelegate>
@property (assign) long goID;
@end

void* WrapAlertDelegate(long goID);
