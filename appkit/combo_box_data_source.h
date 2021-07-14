#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSComboBoxDataSourceAdaptor : NSObject <NSComboBoxDataSource>
@property (assign) uintptr_t goID;
@end

void* WrapComboBoxDataSource(uintptr_t goID);
