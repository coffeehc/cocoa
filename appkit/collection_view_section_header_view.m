#import "collection_view_section_header_view.h"
#import "_cgo_export.h"
#import <AppKit/NSCollectionView.h>

@interface NSCollectionViewSectionHeaderViewAdaptor : NSObject <NSCollectionViewSectionHeaderView>
@property (assign) uintptr_t goID;
@end

@implementation NSCollectionViewSectionHeaderViewAdaptor


- (void)prepareForReuse {
    collectionViewSectionHeaderView_PrepareForReuse([self goID]);
}

- (NSCollectionViewLayoutAttributes*)preferredLayoutAttributesFittingAttributes:(NSCollectionViewLayoutAttributes*)layoutAttributes {
    void* result_ = collectionViewSectionHeaderView_PreferredLayoutAttributesFittingAttributes([self goID], layoutAttributes);
    return (NSCollectionViewLayoutAttributes*)result_;
}

- (void)applyLayoutAttributes:(NSCollectionViewLayoutAttributes*)layoutAttributes {
    collectionViewSectionHeaderView_ApplyLayoutAttributes([self goID], layoutAttributes);
}

- (void)willTransitionFromLayout:(NSCollectionViewLayout*)oldLayout toLayout:(NSCollectionViewLayout*)newLayout {
    collectionViewSectionHeaderView_WillTransitionFromLayout_ToLayout([self goID], oldLayout, newLayout);
}

- (void)didTransitionFromLayout:(NSCollectionViewLayout*)oldLayout toLayout:(NSCollectionViewLayout*)newLayout {
    collectionViewSectionHeaderView_DidTransitionFromLayout_ToLayout([self goID], oldLayout, newLayout);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return CollectionViewSectionHeaderView_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapCollectionViewSectionHeaderView(uintptr_t goID) {
    NSCollectionViewSectionHeaderViewAdaptor* adaptor = [[NSCollectionViewSectionHeaderViewAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
