#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSApplication.h>

@interface NSApplicationDelegateAdaptor : NSObject <NSApplicationDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapApplicationDelegate(uintptr_t goID);
