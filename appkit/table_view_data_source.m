#import <Appkit/Appkit.h>
#import "table_view_data_source.h"
#import "_cgo_export.h"

@implementation NSTableViewDataSourceAdaptor


- (NSInteger)numberOfRowsInTableView:(NSTableView*)tableView {
    int result_ = TableViewDataSource_NumberOfRowsInTableView([self goID], tableView);
    return result_;
}

- (id)tableView:(NSTableView*)tableView objectValueForTableColumn:(NSTableColumn*)tableColumn row:(NSInteger)row {
    void* result_ = TableViewDataSource_TableView_ObjectValueForTableColumn_Row([self goID], tableView, tableColumn, row);
    return (id)result_;
}

- (void)tableView:(NSTableView*)tableView setObjectValue:(id)object forTableColumn:(NSTableColumn*)tableColumn row:(NSInteger)row {
    TableViewDataSource_TableView_SetObjectValue_ForTableColumn_Row([self goID], tableView, object, tableColumn, row);
}

- (id)tableView:(NSTableView*)tableView pasteboardWriterForRow:(NSInteger)row {
    void* result_ = TableViewDataSource_TableView_PasteboardWriterForRow([self goID], tableView, row);
    return (id)result_;
}

- (BOOL)tableView:(NSTableView*)tableView acceptDrop:(id)info row:(NSInteger)row dropOperation:(NSTableViewDropOperation)dropOperation {
    bool result_ = TableViewDataSource_TableView_AcceptDrop_Row_DropOperation([self goID], tableView, info, row, dropOperation);
    return result_;
}

- (NSDragOperation)tableView:(NSTableView*)tableView validateDrop:(id)info proposedRow:(NSInteger)row proposedDropOperation:(NSTableViewDropOperation)dropOperation {
    unsigned int result_ = TableViewDataSource_TableView_ValidateDrop_ProposedRow_ProposedDropOperation([self goID], tableView, info, row, dropOperation);
    return result_;
}

- (void)tableView:(NSTableView*)tableView draggingSession:(NSDraggingSession*)session willBeginAtPoint:(NSPoint)screenPoint forRowIndexes:(NSIndexSet*)rowIndexes {
    TableViewDataSource_TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes([self goID], tableView, session, screenPoint, rowIndexes);
}

- (void)tableView:(NSTableView*)tableView updateDraggingItemsForDrag:(id)draggingInfo {
    TableViewDataSource_TableView_UpdateDraggingItemsForDrag([self goID], tableView, draggingInfo);
}

- (void)tableView:(NSTableView*)tableView draggingSession:(NSDraggingSession*)session endedAtPoint:(NSPoint)screenPoint operation:(NSDragOperation)operation {
    TableViewDataSource_TableView_DraggingSession_EndedAtPoint_Operation([self goID], tableView, session, screenPoint, operation);
}

- (void)tableView:(NSTableView*)tableView sortDescriptorsDidChange:(NSArray*)oldDescriptors {
    int oldDescriptorscount = [oldDescriptors count];
    void** oldDescriptorsData = malloc(oldDescriptorscount * sizeof(void*));
    for (int i = 0; i < oldDescriptorscount; i++) {
    	 void* p = [oldDescriptors objectAtIndex:i];
    	 oldDescriptorsData[i] = p;
    }
    Array oldDescriptorsArray;
    oldDescriptorsArray.data = oldDescriptorsData;
    oldDescriptorsArray.len = oldDescriptorscount;
    TableViewDataSource_TableView_SortDescriptorsDidChange([self goID], tableView, oldDescriptorsArray);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return TableViewDataSource_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteTableViewDataSource([self goID]);
	[super dealloc];
}
@end

void* WrapTableViewDataSource(long goID) {
    NSTableViewDataSourceAdaptor* adaptor = [[NSTableViewDataSourceAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
