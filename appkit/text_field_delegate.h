#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSTextFieldDelegateAdaptor : NSObject <NSTextFieldDelegate>
@property (assign) long goID;
@end

void* WrapTextFieldDelegate(long goID);
