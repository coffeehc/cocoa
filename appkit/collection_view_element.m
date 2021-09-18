#import "collection_view_element.h"
#import "_cgo_export.h"
#import <AppKit/NSCollectionView.h>

@interface NSCollectionViewElementAdaptor : NSObject <NSCollectionViewElement>
@property (assign) uintptr_t goID;
@end

@implementation NSCollectionViewElementAdaptor


- (void)prepareForReuse {
    collectionViewElement_PrepareForReuse([self goID]);
}

- (NSCollectionViewLayoutAttributes*)preferredLayoutAttributesFittingAttributes:(NSCollectionViewLayoutAttributes*)layoutAttributes {
    void* result_ = collectionViewElement_PreferredLayoutAttributesFittingAttributes([self goID], layoutAttributes);
    return (NSCollectionViewLayoutAttributes*)result_;
}

- (void)applyLayoutAttributes:(NSCollectionViewLayoutAttributes*)layoutAttributes {
    collectionViewElement_ApplyLayoutAttributes([self goID], layoutAttributes);
}

- (void)willTransitionFromLayout:(NSCollectionViewLayout*)oldLayout toLayout:(NSCollectionViewLayout*)newLayout {
    collectionViewElement_WillTransitionFromLayout_ToLayout([self goID], oldLayout, newLayout);
}

- (void)didTransitionFromLayout:(NSCollectionViewLayout*)oldLayout toLayout:(NSCollectionViewLayout*)newLayout {
    collectionViewElement_DidTransitionFromLayout_ToLayout([self goID], oldLayout, newLayout);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return CollectionViewElement_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapCollectionViewElement(uintptr_t goID) {
    NSCollectionViewElementAdaptor* adaptor = [[NSCollectionViewElementAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
