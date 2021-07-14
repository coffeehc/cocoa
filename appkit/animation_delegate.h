#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSAnimationDelegateAdaptor : NSObject <NSAnimationDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapAnimationDelegate(uintptr_t goID);
