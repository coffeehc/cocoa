#import <Appkit/Appkit.h>
#import "collection_view_layout_invalidation_context.h"

void* C_CollectionViewLayoutInvalidationContext_Alloc() {
    return [NSCollectionViewLayoutInvalidationContext alloc];
}

void* C_NSCollectionViewLayoutInvalidationContext_Init(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    NSCollectionViewLayoutInvalidationContext* result_ = [nSCollectionViewLayoutInvalidationContext init];
    return result_;
}

void C_NSCollectionViewLayoutInvalidationContext_InvalidateItemsAtIndexPaths(void* ptr, void* indexPaths) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    [nSCollectionViewLayoutInvalidationContext invalidateItemsAtIndexPaths:(NSSet*)indexPaths];
}

void C_NSCollectionViewLayoutInvalidationContext_InvalidateSupplementaryElementsOfKind_AtIndexPaths(void* ptr, void* elementKind, void* indexPaths) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    [nSCollectionViewLayoutInvalidationContext invalidateSupplementaryElementsOfKind:(NSString*)elementKind atIndexPaths:(NSSet*)indexPaths];
}

void C_NSCollectionViewLayoutInvalidationContext_InvalidateDecorationElementsOfKind_AtIndexPaths(void* ptr, void* elementKind, void* indexPaths) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    [nSCollectionViewLayoutInvalidationContext invalidateDecorationElementsOfKind:(NSString*)elementKind atIndexPaths:(NSSet*)indexPaths];
}

bool C_NSCollectionViewLayoutInvalidationContext_InvalidateEverything(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    BOOL result_ = [nSCollectionViewLayoutInvalidationContext invalidateEverything];
    return result_;
}

bool C_NSCollectionViewLayoutInvalidationContext_InvalidateDataSourceCounts(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    BOOL result_ = [nSCollectionViewLayoutInvalidationContext invalidateDataSourceCounts];
    return result_;
}

CGPoint C_NSCollectionViewLayoutInvalidationContext_ContentOffsetAdjustment(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    NSPoint result_ = [nSCollectionViewLayoutInvalidationContext contentOffsetAdjustment];
    return result_;
}

void C_NSCollectionViewLayoutInvalidationContext_SetContentOffsetAdjustment(void* ptr, CGPoint value) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    [nSCollectionViewLayoutInvalidationContext setContentOffsetAdjustment:value];
}

CGSize C_NSCollectionViewLayoutInvalidationContext_ContentSizeAdjustment(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    NSSize result_ = [nSCollectionViewLayoutInvalidationContext contentSizeAdjustment];
    return result_;
}

void C_NSCollectionViewLayoutInvalidationContext_SetContentSizeAdjustment(void* ptr, CGSize value) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    [nSCollectionViewLayoutInvalidationContext setContentSizeAdjustment:value];
}

void* C_NSCollectionViewLayoutInvalidationContext_InvalidatedItemIndexPaths(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    NSSet* result_ = [nSCollectionViewLayoutInvalidationContext invalidatedItemIndexPaths];
    return result_;
}
