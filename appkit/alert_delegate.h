#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSAlert.h>

@interface NSAlertDelegateAdaptor : NSObject <NSAlertDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapAlertDelegate(uintptr_t goID);
