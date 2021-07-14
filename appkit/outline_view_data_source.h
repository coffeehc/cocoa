#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSOutlineViewDataSourceAdaptor : NSObject <NSOutlineViewDataSource>
@property (assign) uintptr_t goID;
@end

void* WrapOutlineViewDataSource(uintptr_t goID);
