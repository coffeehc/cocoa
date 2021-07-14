#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSMenuDelegateAdaptor : NSObject <NSMenuDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapMenuDelegate(uintptr_t goID);
