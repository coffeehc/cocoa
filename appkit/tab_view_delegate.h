#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSTabView.h>

@interface NSTabViewDelegateAdaptor : NSObject <NSTabViewDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapTabViewDelegate(uintptr_t goID);
