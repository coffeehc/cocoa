#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSCollectionViewDelegateAdaptor : NSObject <NSCollectionViewDelegate>
@property (assign) long goID;
@end

void* WrapCollectionViewDelegate(long goID);
