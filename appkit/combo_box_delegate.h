#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSComboBoxDelegateAdaptor : NSObject <NSComboBoxDelegate>
@property (assign) long goID;
@end

void* WrapComboBoxDelegate(long goID);
