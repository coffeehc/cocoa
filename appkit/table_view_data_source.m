#import "table_view_data_source.h"
#import "_cgo_export.h"

@implementation NSTableViewDataSourceAdaptor


- (NSInteger)numberOfRowsInTableView:(NSTableView*)tableView {
    int result_ = tableViewDataSource_NumberOfRowsInTableView([self goID], tableView);
    return result_;
}

- (id)tableView:(NSTableView*)tableView objectValueForTableColumn:(NSTableColumn*)tableColumn row:(NSInteger)row {
    void* result_ = tableViewDataSource_TableView_ObjectValueForTableColumn_Row([self goID], tableView, tableColumn, row);
    return (id)result_;
}

- (void)tableView:(NSTableView*)tableView setObjectValue:(id)object forTableColumn:(NSTableColumn*)tableColumn row:(NSInteger)row {
    tableViewDataSource_TableView_SetObjectValue_ForTableColumn_Row([self goID], tableView, object, tableColumn, row);
}

- (id)tableView:(NSTableView*)tableView pasteboardWriterForRow:(NSInteger)row {
    void* result_ = tableViewDataSource_TableView_PasteboardWriterForRow([self goID], tableView, row);
    return (id)result_;
}

- (BOOL)tableView:(NSTableView*)tableView acceptDrop:(id)info row:(NSInteger)row dropOperation:(NSTableViewDropOperation)dropOperation {
    bool result_ = tableViewDataSource_TableView_AcceptDrop_Row_DropOperation([self goID], tableView, info, row, dropOperation);
    return result_;
}

- (NSDragOperation)tableView:(NSTableView*)tableView validateDrop:(id)info proposedRow:(NSInteger)row proposedDropOperation:(NSTableViewDropOperation)dropOperation {
    unsigned int result_ = tableViewDataSource_TableView_ValidateDrop_ProposedRow_ProposedDropOperation([self goID], tableView, info, row, dropOperation);
    return result_;
}

- (void)tableView:(NSTableView*)tableView draggingSession:(NSDraggingSession*)session willBeginAtPoint:(NSPoint)screenPoint forRowIndexes:(NSIndexSet*)rowIndexes {
    tableViewDataSource_TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes([self goID], tableView, session, screenPoint, rowIndexes);
}

- (void)tableView:(NSTableView*)tableView updateDraggingItemsForDrag:(id)draggingInfo {
    tableViewDataSource_TableView_UpdateDraggingItemsForDrag([self goID], tableView, draggingInfo);
}

- (void)tableView:(NSTableView*)tableView draggingSession:(NSDraggingSession*)session endedAtPoint:(NSPoint)screenPoint operation:(NSDragOperation)operation {
    tableViewDataSource_TableView_DraggingSession_EndedAtPoint_Operation([self goID], tableView, session, screenPoint, operation);
}

- (void)tableView:(NSTableView*)tableView sortDescriptorsDidChange:(NSArray*)oldDescriptors {
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
    tableViewDataSource_TableView_SortDescriptorsDidChange([self goID], tableView, oldDescriptorsArray);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return TableViewDataSource_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapTableViewDataSource(uintptr_t goID) {
    NSTableViewDataSourceAdaptor* adaptor = [[NSTableViewDataSourceAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
