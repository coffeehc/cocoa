#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSCollectionViewDelegateFlowLayoutAdaptor : NSObject <NSCollectionViewDelegateFlowLayout>
@property (assign) uintptr_t goID;
@end

void* WrapCollectionViewDelegateFlowLayout(uintptr_t goID);
