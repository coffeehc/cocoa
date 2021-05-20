#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSOutlineViewDataSourceAdaptor : NSObject <NSOutlineViewDataSource>
@property (assign) long goID;
@end

void* WrapOutlineViewDataSource(long goID);
