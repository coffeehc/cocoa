#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSTextDelegateAdaptor : NSObject <NSTextDelegate>
@property (assign) long goID;
@end

void* WrapTextDelegate(long goID);
