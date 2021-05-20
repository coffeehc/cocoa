#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSTableViewDataSourceAdaptor : NSObject <NSTableViewDataSource>
@property (assign) long goID;
@end

void* WrapTableViewDataSource(long goID);
