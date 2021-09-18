#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSCollectionView.h>

@interface NSCollectionViewElementAdaptor : NSObject <NSCollectionViewElement>
@property (assign) uintptr_t goID;
@end

void* WrapCollectionViewElement(uintptr_t goID);
