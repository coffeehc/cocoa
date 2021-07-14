#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSTableViewDelegateAdaptor : NSObject <NSTableViewDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapTableViewDelegate(uintptr_t goID);
