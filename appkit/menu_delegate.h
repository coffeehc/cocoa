#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSMenuDelegateAdaptor : NSObject <NSMenuDelegate>
@property (assign) long goID;
@end

void* WrapMenuDelegate(long goID);
