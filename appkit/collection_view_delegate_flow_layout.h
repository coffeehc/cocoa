#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSCollectionViewDelegateFlowLayoutAdaptor : NSObject <NSCollectionViewDelegateFlowLayout>
@property (assign) long goID;
@end

void* WrapCollectionViewDelegateFlowLayout(long goID);
