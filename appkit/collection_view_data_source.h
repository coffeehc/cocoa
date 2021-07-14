#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSCollectionViewDataSourceAdaptor : NSObject <NSCollectionViewDataSource>
@property (assign) uintptr_t goID;
@end

void* WrapCollectionViewDataSource(uintptr_t goID);
