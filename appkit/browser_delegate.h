#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSBrowserDelegateAdaptor : NSObject <NSBrowserDelegate>
@property (assign) long goID;
@end

void* WrapBrowserDelegate(long goID);
