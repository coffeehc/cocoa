#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSTableView.h>

@interface NSTableViewDataSourceAdaptor : NSObject <NSTableViewDataSource>
@property (assign) uintptr_t goID;
@end

void* WrapTableViewDataSource(uintptr_t goID);
