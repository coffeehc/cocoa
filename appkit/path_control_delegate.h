#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSPathControl.h>

@interface NSPathControlDelegateAdaptor : NSObject <NSPathControlDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapPathControlDelegate(uintptr_t goID);
