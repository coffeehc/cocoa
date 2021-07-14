#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSPathControlDelegateAdaptor : NSObject <NSPathControlDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapPathControlDelegate(uintptr_t goID);
