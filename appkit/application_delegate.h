#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSApplicationDelegateAdaptor : NSObject <NSApplicationDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapApplicationDelegate(uintptr_t goID);
