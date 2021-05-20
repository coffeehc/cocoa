#import <Appkit/Appkit.h>
#import "outline_view_data_source.h"
#import "_cgo_export.h"

@implementation NSOutlineViewDataSourceAdaptor


- (BOOL)outlineView:(NSOutlineView*)outlineView acceptDrop:(id)info item:(id)item childIndex:(NSInteger)index {
    bool result_ = OutlineViewDataSource_OutlineView_AcceptDrop_Item_ChildIndex([self goID], outlineView, info, item, index);
    return result_;
}

- (id)outlineView:(NSOutlineView*)outlineView child:(NSInteger)index ofItem:(id)item {
    void* result_ = OutlineViewDataSource_OutlineView_Child_OfItem([self goID], outlineView, index, item);
    return (id)result_;
}

- (void)outlineView:(NSOutlineView*)outlineView draggingSession:(NSDraggingSession*)session endedAtPoint:(NSPoint)screenPoint operation:(NSDragOperation)operation {
    OutlineViewDataSource_OutlineView_DraggingSession_EndedAtPoint_Operation([self goID], outlineView, session, screenPoint, operation);
}

- (void)outlineView:(NSOutlineView*)outlineView draggingSession:(NSDraggingSession*)session willBeginAtPoint:(NSPoint)screenPoint forItems:(NSArray*)draggedItems {
    int draggedItemscount = [draggedItems count];
    void** draggedItemsData = malloc(draggedItemscount * sizeof(void*));
    for (int i = 0; i < draggedItemscount; i++) {
    	 void* p = [draggedItems objectAtIndex:i];
    	 draggedItemsData[i] = p;
    }
    Array draggedItemsArray;
    draggedItemsArray.data = draggedItemsData;
    draggedItemsArray.len = draggedItemscount;
    OutlineViewDataSource_OutlineView_DraggingSession_WillBeginAtPoint_ForItems([self goID], outlineView, session, screenPoint, draggedItemsArray);
}

- (BOOL)outlineView:(NSOutlineView*)outlineView isItemExpandable:(id)item {
    bool result_ = OutlineViewDataSource_OutlineView_IsItemExpandable([self goID], outlineView, item);
    return result_;
}

- (id)outlineView:(NSOutlineView*)outlineView itemForPersistentObject:(id)object {
    void* result_ = OutlineViewDataSource_OutlineView_ItemForPersistentObject([self goID], outlineView, object);
    return (id)result_;
}

- (NSInteger)outlineView:(NSOutlineView*)outlineView numberOfChildrenOfItem:(id)item {
    int result_ = OutlineViewDataSource_OutlineView_NumberOfChildrenOfItem([self goID], outlineView, item);
    return result_;
}

- (id)outlineView:(NSOutlineView*)outlineView objectValueForTableColumn:(NSTableColumn*)tableColumn byItem:(id)item {
    void* result_ = OutlineViewDataSource_OutlineView_ObjectValueForTableColumn_ByItem([self goID], outlineView, tableColumn, item);
    return (id)result_;
}

- (id)outlineView:(NSOutlineView*)outlineView pasteboardWriterForItem:(id)item {
    void* result_ = OutlineViewDataSource_OutlineView_PasteboardWriterForItem([self goID], outlineView, item);
    return (id)result_;
}

- (id)outlineView:(NSOutlineView*)outlineView persistentObjectForItem:(id)item {
    void* result_ = OutlineViewDataSource_OutlineView_PersistentObjectForItem([self goID], outlineView, item);
    return (id)result_;
}

- (void)outlineView:(NSOutlineView*)outlineView setObjectValue:(id)object forTableColumn:(NSTableColumn*)tableColumn byItem:(id)item {
    OutlineViewDataSource_OutlineView_SetObjectValue_ForTableColumn_ByItem([self goID], outlineView, object, tableColumn, item);
}

- (void)outlineView:(NSOutlineView*)outlineView sortDescriptorsDidChange:(NSArray*)oldDescriptors {
    int oldDescriptorscount = [oldDescriptors count];
    void** oldDescriptorsData = malloc(oldDescriptorscount * sizeof(void*));
    for (int i = 0; i < oldDescriptorscount; i++) {
    	 void* p = [oldDescriptors objectAtIndex:i];
    	 oldDescriptorsData[i] = p;
    }
    Array oldDescriptorsArray;
    oldDescriptorsArray.data = oldDescriptorsData;
    oldDescriptorsArray.len = oldDescriptorscount;
    OutlineViewDataSource_OutlineView_SortDescriptorsDidChange([self goID], outlineView, oldDescriptorsArray);
}

- (void)outlineView:(NSOutlineView*)outlineView updateDraggingItemsForDrag:(id)draggingInfo {
    OutlineViewDataSource_OutlineView_UpdateDraggingItemsForDrag([self goID], outlineView, draggingInfo);
}

- (NSDragOperation)outlineView:(NSOutlineView*)outlineView validateDrop:(id)info proposedItem:(id)item proposedChildIndex:(NSInteger)index {
    unsigned int result_ = OutlineViewDataSource_OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex([self goID], outlineView, info, item, index);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return OutlineViewDataSource_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteOutlineViewDataSource([self goID]);
	[super dealloc];
}
@end

void* WrapOutlineViewDataSource(long goID) {
    NSOutlineViewDataSourceAdaptor* adaptor = [[NSOutlineViewDataSourceAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
