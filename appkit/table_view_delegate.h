#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSTableViewDelegateAdaptor : NSObject <NSTableViewDelegate>
@property (assign) long goID;
@end

void* WrapTableViewDelegate(long goID);
