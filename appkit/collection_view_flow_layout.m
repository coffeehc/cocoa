#import "collection_view_flow_layout.h"
#import <AppKit/NSCollectionViewFlowLayout.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

void* C_CollectionViewFlowLayout_Alloc() {
    return [NSCollectionViewFlowLayout alloc];
}

void* C_NSCollectionViewFlowLayout_AllocCollectionViewFlowLayout() {
    NSCollectionViewFlowLayout* result_ = [NSCollectionViewFlowLayout alloc];
    return result_;
}

void* C_NSCollectionViewFlowLayout_Init(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSCollectionViewFlowLayout* result_ = [nSCollectionViewFlowLayout init];
    return result_;
}

void* C_NSCollectionViewFlowLayout_NewCollectionViewFlowLayout() {
    NSCollectionViewFlowLayout* result_ = [NSCollectionViewFlowLayout new];
    return result_;
}

void* C_NSCollectionViewFlowLayout_Autorelease(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSCollectionViewFlowLayout* result_ = [nSCollectionViewFlowLayout autorelease];
    return result_;
}

void* C_NSCollectionViewFlowLayout_Retain(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSCollectionViewFlowLayout* result_ = [nSCollectionViewFlowLayout retain];
    return result_;
}

void C_NSCollectionViewFlowLayout_CollapseSectionAtIndex(void* ptr, unsigned int sectionIndex) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout collapseSectionAtIndex:sectionIndex];
}

void C_NSCollectionViewFlowLayout_ExpandSectionAtIndex(void* ptr, unsigned int sectionIndex) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout expandSectionAtIndex:sectionIndex];
}

bool C_NSCollectionViewFlowLayout_SectionAtIndexIsCollapsed(void* ptr, unsigned int sectionIndex) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    BOOL result_ = [nSCollectionViewFlowLayout sectionAtIndexIsCollapsed:sectionIndex];
    return result_;
}

int C_NSCollectionViewFlowLayout_ScrollDirection(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSCollectionViewScrollDirection result_ = [nSCollectionViewFlowLayout scrollDirection];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetScrollDirection(void* ptr, int value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setScrollDirection:value];
}

double C_NSCollectionViewFlowLayout_MinimumLineSpacing(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    CGFloat result_ = [nSCollectionViewFlowLayout minimumLineSpacing];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetMinimumLineSpacing(void* ptr, double value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setMinimumLineSpacing:value];
}

double C_NSCollectionViewFlowLayout_MinimumInteritemSpacing(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    CGFloat result_ = [nSCollectionViewFlowLayout minimumInteritemSpacing];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetMinimumInteritemSpacing(void* ptr, double value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setMinimumInteritemSpacing:value];
}

CGSize C_NSCollectionViewFlowLayout_EstimatedItemSize(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSSize result_ = [nSCollectionViewFlowLayout estimatedItemSize];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetEstimatedItemSize(void* ptr, CGSize value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setEstimatedItemSize:value];
}

CGSize C_NSCollectionViewFlowLayout_ItemSize(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSSize result_ = [nSCollectionViewFlowLayout itemSize];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetItemSize(void* ptr, CGSize value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setItemSize:value];
}

NSEdgeInsets C_NSCollectionViewFlowLayout_SectionInset(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSEdgeInsets result_ = [nSCollectionViewFlowLayout sectionInset];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetSectionInset(void* ptr, NSEdgeInsets value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setSectionInset:value];
}

CGSize C_NSCollectionViewFlowLayout_HeaderReferenceSize(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSSize result_ = [nSCollectionViewFlowLayout headerReferenceSize];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetHeaderReferenceSize(void* ptr, CGSize value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setHeaderReferenceSize:value];
}

CGSize C_NSCollectionViewFlowLayout_FooterReferenceSize(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSSize result_ = [nSCollectionViewFlowLayout footerReferenceSize];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetFooterReferenceSize(void* ptr, CGSize value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setFooterReferenceSize:value];
}

bool C_NSCollectionViewFlowLayout_SectionFootersPinToVisibleBounds(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    BOOL result_ = [nSCollectionViewFlowLayout sectionFootersPinToVisibleBounds];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetSectionFootersPinToVisibleBounds(void* ptr, bool value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setSectionFootersPinToVisibleBounds:value];
}

bool C_NSCollectionViewFlowLayout_SectionHeadersPinToVisibleBounds(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    BOOL result_ = [nSCollectionViewFlowLayout sectionHeadersPinToVisibleBounds];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetSectionHeadersPinToVisibleBounds(void* ptr, bool value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setSectionHeadersPinToVisibleBounds:value];
}

void* C_CollectionViewFlowLayoutInvalidationContext_Alloc() {
    return [NSCollectionViewFlowLayoutInvalidationContext alloc];
}

void* C_NSCollectionViewFlowLayoutInvalidationContext_AllocCollectionViewFlowLayoutInvalidationContext() {
    NSCollectionViewFlowLayoutInvalidationContext* result_ = [NSCollectionViewFlowLayoutInvalidationContext alloc];
    return result_;
}

void* C_NSCollectionViewFlowLayoutInvalidationContext_Init(void* ptr) {
    NSCollectionViewFlowLayoutInvalidationContext* nSCollectionViewFlowLayoutInvalidationContext = (NSCollectionViewFlowLayoutInvalidationContext*)ptr;
    NSCollectionViewFlowLayoutInvalidationContext* result_ = [nSCollectionViewFlowLayoutInvalidationContext init];
    return result_;
}

void* C_NSCollectionViewFlowLayoutInvalidationContext_NewCollectionViewFlowLayoutInvalidationContext() {
    NSCollectionViewFlowLayoutInvalidationContext* result_ = [NSCollectionViewFlowLayoutInvalidationContext new];
    return result_;
}

void* C_NSCollectionViewFlowLayoutInvalidationContext_Autorelease(void* ptr) {
    NSCollectionViewFlowLayoutInvalidationContext* nSCollectionViewFlowLayoutInvalidationContext = (NSCollectionViewFlowLayoutInvalidationContext*)ptr;
    NSCollectionViewFlowLayoutInvalidationContext* result_ = [nSCollectionViewFlowLayoutInvalidationContext autorelease];
    return result_;
}

void* C_NSCollectionViewFlowLayoutInvalidationContext_Retain(void* ptr) {
    NSCollectionViewFlowLayoutInvalidationContext* nSCollectionViewFlowLayoutInvalidationContext = (NSCollectionViewFlowLayoutInvalidationContext*)ptr;
    NSCollectionViewFlowLayoutInvalidationContext* result_ = [nSCollectionViewFlowLayoutInvalidationContext retain];
    return result_;
}

bool C_NSCollectionViewFlowLayoutInvalidationContext_InvalidateFlowLayoutAttributes(void* ptr) {
    NSCollectionViewFlowLayoutInvalidationContext* nSCollectionViewFlowLayoutInvalidationContext = (NSCollectionViewFlowLayoutInvalidationContext*)ptr;
    BOOL result_ = [nSCollectionViewFlowLayoutInvalidationContext invalidateFlowLayoutAttributes];
    return result_;
}

void C_NSCollectionViewFlowLayoutInvalidationContext_SetInvalidateFlowLayoutAttributes(void* ptr, bool value) {
    NSCollectionViewFlowLayoutInvalidationContext* nSCollectionViewFlowLayoutInvalidationContext = (NSCollectionViewFlowLayoutInvalidationContext*)ptr;
    [nSCollectionViewFlowLayoutInvalidationContext setInvalidateFlowLayoutAttributes:value];
}

bool C_NSCollectionViewFlowLayoutInvalidationContext_InvalidateFlowLayoutDelegateMetrics(void* ptr) {
    NSCollectionViewFlowLayoutInvalidationContext* nSCollectionViewFlowLayoutInvalidationContext = (NSCollectionViewFlowLayoutInvalidationContext*)ptr;
    BOOL result_ = [nSCollectionViewFlowLayoutInvalidationContext invalidateFlowLayoutDelegateMetrics];
    return result_;
}

void C_NSCollectionViewFlowLayoutInvalidationContext_SetInvalidateFlowLayoutDelegateMetrics(void* ptr, bool value) {
    NSCollectionViewFlowLayoutInvalidationContext* nSCollectionViewFlowLayoutInvalidationContext = (NSCollectionViewFlowLayoutInvalidationContext*)ptr;
    [nSCollectionViewFlowLayoutInvalidationContext setInvalidateFlowLayoutDelegateMetrics:value];
}

@interface NSCollectionViewDelegateFlowLayoutAdaptor : NSObject <NSCollectionViewDelegateFlowLayout>
@property (assign) uintptr_t goID;
@end

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
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapCollectionViewDelegateFlowLayout(uintptr_t goID) {
    NSCollectionViewDelegateFlowLayoutAdaptor* adaptor = [[NSCollectionViewDelegateFlowLayoutAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
