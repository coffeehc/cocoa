#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSApplicationDelegateAdaptor : NSObject <NSApplicationDelegate>
@property (assign) long goID;
@end

void* WrapApplicationDelegate(long goID);
