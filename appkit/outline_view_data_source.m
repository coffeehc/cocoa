#import "outline_view_data_source.h"
#import "_cgo_export.h"
#import <AppKit/NSOutlineView.h>

@interface NSOutlineViewDataSourceAdaptor : NSObject <NSOutlineViewDataSource>
@property (assign) uintptr_t goID;
@end

@implementation NSOutlineViewDataSourceAdaptor


- (BOOL)outlineView:(NSOutlineView*)outlineView acceptDrop:(id)info item:(id)item childIndex:(NSInteger)index {
    bool result_ = outlineViewDataSource_OutlineView_AcceptDrop_Item_ChildIndex([self goID], outlineView, info, item, index);
    return result_;
}

- (id)outlineView:(NSOutlineView*)outlineView child:(NSInteger)index ofItem:(id)item {
    void* result_ = outlineViewDataSource_OutlineView_Child_OfItem([self goID], outlineView, index, item);
    return (id)result_;
}

- (void)outlineView:(NSOutlineView*)outlineView draggingSession:(NSDraggingSession*)session endedAtPoint:(NSPoint)screenPoint operation:(NSDragOperation)operation {
    outlineViewDataSource_OutlineView_DraggingSession_EndedAtPoint_Operation([self goID], outlineView, session, screenPoint, operation);
}

- (void)outlineView:(NSOutlineView*)outlineView draggingSession:(NSDraggingSession*)session willBeginAtPoint:(NSPoint)screenPoint forItems:(NSArray*)draggedItems {
    Array draggedItemsArray;
    int draggedItemscount = [draggedItems count];
    if (draggedItemscount > 0) {
    	void** draggedItemsData = malloc(draggedItemscount * sizeof(void*));
    	for (int i = 0; i < draggedItemscount; i++) {
    		 void* p = [draggedItems objectAtIndex:i];
    		 draggedItemsData[i] = p;
    	}
    	draggedItemsArray.data = draggedItemsData;
    	draggedItemsArray.len = draggedItemscount;
    }
    outlineViewDataSource_OutlineView_DraggingSession_WillBeginAtPoint_ForItems([self goID], outlineView, session, screenPoint, draggedItemsArray);
}

- (BOOL)outlineView:(NSOutlineView*)outlineView isItemExpandable:(id)item {
    bool result_ = outlineViewDataSource_OutlineView_IsItemExpandable([self goID], outlineView, item);
    return result_;
}

- (id)outlineView:(NSOutlineView*)outlineView itemForPersistentObject:(id)object {
    void* result_ = outlineViewDataSource_OutlineView_ItemForPersistentObject([self goID], outlineView, object);
    return (id)result_;
}

- (NSInteger)outlineView:(NSOutlineView*)outlineView numberOfChildrenOfItem:(id)item {
    int result_ = outlineViewDataSource_OutlineView_NumberOfChildrenOfItem([self goID], outlineView, item);
    return result_;
}

- (id)outlineView:(NSOutlineView*)outlineView objectValueForTableColumn:(NSTableColumn*)tableColumn byItem:(id)item {
    void* result_ = outlineViewDataSource_OutlineView_ObjectValueForTableColumn_ByItem([self goID], outlineView, tableColumn, item);
    return (id)result_;
}

- (id)outlineView:(NSOutlineView*)outlineView pasteboardWriterForItem:(id)item {
    void* result_ = outlineViewDataSource_OutlineView_PasteboardWriterForItem([self goID], outlineView, item);
    return (id)result_;
}

- (id)outlineView:(NSOutlineView*)outlineView persistentObjectForItem:(id)item {
    void* result_ = outlineViewDataSource_OutlineView_PersistentObjectForItem([self goID], outlineView, item);
    return (id)result_;
}

- (void)outlineView:(NSOutlineView*)outlineView setObjectValue:(id)object forTableColumn:(NSTableColumn*)tableColumn byItem:(id)item {
    outlineViewDataSource_OutlineView_SetObjectValue_ForTableColumn_ByItem([self goID], outlineView, object, tableColumn, item);
}

- (void)outlineView:(NSOutlineView*)outlineView sortDescriptorsDidChange:(NSArray*)oldDescriptors {
    Array oldDescriptorsArray;
    int oldDescriptorscount = [oldDescriptors count];
    if (oldDescriptorscount > 0) {
    	void** oldDescriptorsData = malloc(oldDescriptorscount * sizeof(void*));
    	for (int i = 0; i < oldDescriptorscount; i++) {
    		 void* p = [oldDescriptors objectAtIndex:i];
    		 oldDescriptorsData[i] = p;
    	}
    	oldDescriptorsArray.data = oldDescriptorsData;
    	oldDescriptorsArray.len = oldDescriptorscount;
    }
    outlineViewDataSource_OutlineView_SortDescriptorsDidChange([self goID], outlineView, oldDescriptorsArray);
}

- (void)outlineView:(NSOutlineView*)outlineView updateDraggingItemsForDrag:(id)draggingInfo {
    outlineViewDataSource_OutlineView_UpdateDraggingItemsForDrag([self goID], outlineView, draggingInfo);
}

- (NSDragOperation)outlineView:(NSOutlineView*)outlineView validateDrop:(id)info proposedItem:(id)item proposedChildIndex:(NSInteger)index {
    unsigned int result_ = outlineViewDataSource_OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex([self goID], outlineView, info, item, index);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return OutlineViewDataSource_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapOutlineViewDataSource(uintptr_t goID) {
    NSOutlineViewDataSourceAdaptor* adaptor = [[NSOutlineViewDataSourceAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
