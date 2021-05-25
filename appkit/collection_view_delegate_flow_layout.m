#import <Appkit/Appkit.h>
#import "collection_view_delegate_flow_layout.h"
#import "_cgo_export.h"

@implementation NSCollectionViewDelegateFlowLayoutAdaptor


- (NSSize)collectionView:(NSCollectionView*)collectionView layout:(NSCollectionViewLayout*)collectionViewLayout sizeForItemAtIndexPath:(NSIndexPath*)indexPath {
    CGSize result_ = collectionViewDelegateFlowLayout_CollectionView_Layout_SizeForItemAtIndexPath([self goID], collectionView, collectionViewLayout, indexPath);
    return result_;
}

- (NSEdgeInsets)collectionView:(NSCollectionView*)collectionView layout:(NSCollectionViewLayout*)collectionViewLayout insetForSectionAtIndex:(NSInteger)section {
    NSEdgeInsets result_ = collectionViewDelegateFlowLayout_CollectionView_Layout_InsetForSectionAtIndex([self goID], collectionView, collectionViewLayout, section);
    return result_;
}

- (CGFloat)collectionView:(NSCollectionView*)collectionView layout:(NSCollectionViewLayout*)collectionViewLayout minimumLineSpacingForSectionAtIndex:(NSInteger)section {
    double result_ = collectionViewDelegateFlowLayout_CollectionView_Layout_MinimumLineSpacingForSectionAtIndex([self goID], collectionView, collectionViewLayout, section);
    return result_;
}

- (CGFloat)collectionView:(NSCollectionView*)collectionView layout:(NSCollectionViewLayout*)collectionViewLayout minimumInteritemSpacingForSectionAtIndex:(NSInteger)section {
    double result_ = collectionViewDelegateFlowLayout_CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex([self goID], collectionView, collectionViewLayout, section);
    return result_;
}

- (NSSize)collectionView:(NSCollectionView*)collectionView layout:(NSCollectionViewLayout*)collectionViewLayout referenceSizeForHeaderInSection:(NSInteger)section {
    CGSize result_ = collectionViewDelegateFlowLayout_CollectionView_Layout_ReferenceSizeForHeaderInSection([self goID], collectionView, collectionViewLayout, section);
    return result_;
}

- (NSSize)collectionView:(NSCollectionView*)collectionView layout:(NSCollectionViewLayout*)collectionViewLayout referenceSizeForFooterInSection:(NSInteger)section {
    CGSize result_ = collectionViewDelegateFlowLayout_CollectionView_Layout_ReferenceSizeForFooterInSection([self goID], collectionView, collectionViewLayout, section);
    return result_;
}

- (NSSet*)collectionView:(NSCollectionView*)collectionView shouldSelectItemsAtIndexPaths:(NSSet*)indexPaths {
    void* result_ = collectionViewDelegateFlowLayout_CollectionView_ShouldSelectItemsAtIndexPaths([self goID], collectionView, indexPaths);
    return (NSSet*)result_;
}

- (void)collectionView:(NSCollectionView*)collectionView didSelectItemsAtIndexPaths:(NSSet*)indexPaths {
    collectionViewDelegateFlowLayout_CollectionView_DidSelectItemsAtIndexPaths([self goID], collectionView, indexPaths);
}

- (NSSet*)collectionView:(NSCollectionView*)collectionView shouldDeselectItemsAtIndexPaths:(NSSet*)indexPaths {
    void* result_ = collectionViewDelegateFlowLayout_CollectionView_ShouldDeselectItemsAtIndexPaths([self goID], collectionView, indexPaths);
    return (NSSet*)result_;
}

- (void)collectionView:(NSCollectionView*)collectionView didDeselectItemsAtIndexPaths:(NSSet*)indexPaths {
    collectionViewDelegateFlowLayout_CollectionView_DidDeselectItemsAtIndexPaths([self goID], collectionView, indexPaths);
}

- (NSSet*)collectionView:(NSCollectionView*)collectionView shouldChangeItemsAtIndexPaths:(NSSet*)indexPaths toHighlightState:(NSCollectionViewItemHighlightState)highlightState {
    void* result_ = collectionViewDelegateFlowLayout_CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState([self goID], collectionView, indexPaths, highlightState);
    return (NSSet*)result_;
}

- (void)collectionView:(NSCollectionView*)collectionView didChangeItemsAtIndexPaths:(NSSet*)indexPaths toHighlightState:(NSCollectionViewItemHighlightState)highlightState {
    collectionViewDelegateFlowLayout_CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState([self goID], collectionView, indexPaths, highlightState);
}

- (void)collectionView:(NSCollectionView*)collectionView willDisplayItem:(NSCollectionViewItem*)item forRepresentedObjectAtIndexPath:(NSIndexPath*)indexPath {
    collectionViewDelegateFlowLayout_CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath([self goID], collectionView, item, indexPath);
}

- (void)collectionView:(NSCollectionView*)collectionView didEndDisplayingItem:(NSCollectionViewItem*)item forRepresentedObjectAtIndexPath:(NSIndexPath*)indexPath {
    collectionViewDelegateFlowLayout_CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath([self goID], collectionView, item, indexPath);
}

- (void)collectionView:(NSCollectionView*)collectionView willDisplaySupplementaryView:(NSView*)view forElementKind:(NSCollectionViewSupplementaryElementKind)elementKind atIndexPath:(NSIndexPath*)indexPath {
    collectionViewDelegateFlowLayout_CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath([self goID], collectionView, view, elementKind, indexPath);
}

- (void)collectionView:(NSCollectionView*)collectionView didEndDisplayingSupplementaryView:(NSView*)view forElementOfKind:(NSCollectionViewSupplementaryElementKind)elementKind atIndexPath:(NSIndexPath*)indexPath {
    collectionViewDelegateFlowLayout_CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath([self goID], collectionView, view, elementKind, indexPath);
}

- (NSCollectionViewTransitionLayout*)collectionView:(NSCollectionView*)collectionView transitionLayoutForOldLayout:(NSCollectionViewLayout*)fromLayout newLayout:(NSCollectionViewLayout*)toLayout {
    void* result_ = collectionViewDelegateFlowLayout_CollectionView_TransitionLayoutForOldLayout_NewLayout([self goID], collectionView, fromLayout, toLayout);
    return (NSCollectionViewTransitionLayout*)result_;
}

- (BOOL)collectionView:(NSCollectionView*)collectionView canDragItemsAtIndexPaths:(NSSet*)indexPaths withEvent:(NSEvent*)event {
    bool result_ = collectionViewDelegateFlowLayout_CollectionView_CanDragItemsAtIndexPaths_WithEvent([self goID], collectionView, indexPaths, event);
    return result_;
}

- (id)collectionView:(NSCollectionView*)collectionView pasteboardWriterForItemAtIndexPath:(NSIndexPath*)indexPath {
    void* result_ = collectionViewDelegateFlowLayout_CollectionView_PasteboardWriterForItemAtIndexPath([self goID], collectionView, indexPath);
    return (id)result_;
}

- (void)collectionView:(NSCollectionView*)collectionView draggingSession:(NSDraggingSession*)session willBeginAtPoint:(NSPoint)screenPoint forItemsAtIndexPaths:(NSSet*)indexPaths {
    collectionViewDelegateFlowLayout_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths([self goID], collectionView, session, screenPoint, indexPaths);
}

- (void)collectionView:(NSCollectionView*)collectionView draggingSession:(NSDraggingSession*)session endedAtPoint:(NSPoint)screenPoint dragOperation:(NSDragOperation)operation {
    collectionViewDelegateFlowLayout_CollectionView_DraggingSession_EndedAtPoint_DragOperation([self goID], collectionView, session, screenPoint, operation);
}

- (void)collectionView:(NSCollectionView*)collectionView updateDraggingItemsForDrag:(id)draggingInfo {
    collectionViewDelegateFlowLayout_CollectionView_UpdateDraggingItemsForDrag([self goID], collectionView, draggingInfo);
}

- (BOOL)collectionView:(NSCollectionView*)collectionView acceptDrop:(id)draggingInfo indexPath:(NSIndexPath*)indexPath dropOperation:(NSCollectionViewDropOperation)dropOperation {
    bool result_ = collectionViewDelegateFlowLayout_CollectionView_AcceptDrop_IndexPath_DropOperation([self goID], collectionView, draggingInfo, indexPath, dropOperation);
    return result_;
}

- (BOOL)collectionView:(NSCollectionView*)collectionView canDragItemsAtIndexes:(NSIndexSet*)indexes withEvent:(NSEvent*)event {
    bool result_ = collectionViewDelegateFlowLayout_CollectionView_CanDragItemsAtIndexes_WithEvent([self goID], collectionView, indexes, event);
    return result_;
}

- (id)collectionView:(NSCollectionView*)collectionView pasteboardWriterForItemAtIndex:(NSUInteger)index {
    void* result_ = collectionViewDelegateFlowLayout_CollectionView_PasteboardWriterForItemAtIndex([self goID], collectionView, index);
    return (id)result_;
}

- (void)collectionView:(NSCollectionView*)collectionView draggingSession:(NSDraggingSession*)session willBeginAtPoint:(NSPoint)screenPoint forItemsAtIndexes:(NSIndexSet*)indexes {
    collectionViewDelegateFlowLayout_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes([self goID], collectionView, session, screenPoint, indexes);
}

- (BOOL)collectionView:(NSCollectionView*)collectionView acceptDrop:(id)draggingInfo index:(NSInteger)index dropOperation:(NSCollectionViewDropOperation)dropOperation {
    bool result_ = collectionViewDelegateFlowLayout_CollectionView_AcceptDrop_Index_DropOperation([self goID], collectionView, draggingInfo, index, dropOperation);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return CollectionViewDelegateFlowLayout_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteCollectionViewDelegateFlowLayout([self goID]);
	[super dealloc];
}
@end

void* WrapCollectionViewDelegateFlowLayout(long goID) {
    NSCollectionViewDelegateFlowLayoutAdaptor* adaptor = [[NSCollectionViewDelegateFlowLayoutAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
