#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSCollectionViewDataSourceAdaptor : NSObject <NSCollectionViewDataSource>
@property (assign) long goID;
@end

void* WrapCollectionViewDataSource(long goID);
