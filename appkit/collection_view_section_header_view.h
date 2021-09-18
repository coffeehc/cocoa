#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSCollectionView.h>

@interface NSCollectionViewSectionHeaderViewAdaptor : NSObject <NSCollectionViewSectionHeaderView>
@property (assign) uintptr_t goID;
@end

void* WrapCollectionViewSectionHeaderView(uintptr_t goID);
