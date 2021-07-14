#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSTextViewDelegateAdaptor : NSObject <NSTextViewDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapTextViewDelegate(uintptr_t goID);
