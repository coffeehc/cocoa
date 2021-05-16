#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSTextViewDelegateAdaptor : NSObject <NSTextViewDelegate>
@property (assign) long goID;
@end

void* WrapTextViewDelegate(long goID);
