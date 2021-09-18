#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSComboBox.h>

@interface NSComboBoxDelegateAdaptor : NSObject <NSComboBoxDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapComboBoxDelegate(uintptr_t goID);
