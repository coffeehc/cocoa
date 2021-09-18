#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSControl.h>

@interface NSControlTextEditingDelegateAdaptor : NSObject <NSControlTextEditingDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapControlTextEditingDelegate(uintptr_t goID);
