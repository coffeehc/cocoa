#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSAnimation.h>

@interface NSAnimationDelegateAdaptor : NSObject <NSAnimationDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapAnimationDelegate(uintptr_t goID);
