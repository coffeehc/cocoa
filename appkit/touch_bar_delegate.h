#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSTouchBar.h>

@interface NSTouchBarDelegateAdaptor : NSObject <NSTouchBarDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapTouchBarDelegate(uintptr_t goID);
