#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSTextDelegateAdaptor : NSObject <NSTextDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapTextDelegate(uintptr_t goID);
