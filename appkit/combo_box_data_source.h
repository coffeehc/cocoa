#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSComboBoxDataSourceAdaptor : NSObject <NSComboBoxDataSource>
@property (assign) long goID;
@end

void* WrapComboBoxDataSource(long goID);
