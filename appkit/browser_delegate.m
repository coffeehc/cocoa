#import "browser_delegate.h"
#import <AppKit/NSBrowser.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

@interface NSBrowserDelegateAdaptor : NSObject <NSBrowserDelegate>
@property (assign) uintptr_t goID;
@end

@implementation NSBrowserDelegateAdaptor


- (BOOL)browser:(NSBrowser*)sender isColumnValid:(NSInteger)column {
    bool result_ = browserDelegate_Browser_IsColumnValid([self goID], sender, column);
    return result_;
}

- (NSInteger)browser:(NSBrowser*)sender numberOfRowsInColumn:(NSInteger)column {
    int result_ = browserDelegate_Browser_NumberOfRowsInColumn([self goID], sender, column);
    return result_;
}

- (NSInteger)browser:(NSBrowser*)browser numberOfChildrenOfItem:(id)item {
    int result_ = browserDelegate_Browser_NumberOfChildrenOfItem([self goID], browser, item);
    return result_;
}

- (NSString*)browser:(NSBrowser*)sender titleOfColumn:(NSInteger)column {
    void* result_ = browserDelegate_Browser_TitleOfColumn([self goID], sender, column);
    return (NSString*)result_;
}

- (BOOL)browser:(NSBrowser*)browser shouldTypeSelectForEvent:(NSEvent*)event withCurrentSearchString:(NSString*)searchString {
    bool result_ = browserDelegate_Browser_ShouldTypeSelectForEvent_WithCurrentSearchString([self goID], browser, event, searchString);
    return result_;
}

- (NSString*)browser:(NSBrowser*)browser typeSelectStringForRow:(NSInteger)row inColumn:(NSInteger)column {
    void* result_ = browserDelegate_Browser_TypeSelectStringForRow_InColumn([self goID], browser, row, column);
    return (NSString*)result_;
}

- (NSInteger)browser:(NSBrowser*)browser nextTypeSelectMatchFromRow:(NSInteger)startRow toRow:(NSInteger)endRow inColumn:(NSInteger)column forString:(NSString*)searchString {
    int result_ = browserDelegate_Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString([self goID], browser, startRow, endRow, column, searchString);
    return result_;
}

- (BOOL)browser:(NSBrowser*)sender selectCellWithString:(NSString*)title inColumn:(NSInteger)column {
    bool result_ = browserDelegate_Browser_SelectCellWithString_InColumn([self goID], sender, title, column);
    return result_;
}

- (BOOL)browser:(NSBrowser*)sender selectRow:(NSInteger)row inColumn:(NSInteger)column {
    bool result_ = browserDelegate_Browser_SelectRow_InColumn([self goID], sender, row, column);
    return result_;
}

- (NSIndexSet*)browser:(NSBrowser*)browser selectionIndexesForProposedSelection:(NSIndexSet*)proposedSelectionIndexes inColumn:(NSInteger)column {
    void* result_ = browserDelegate_Browser_SelectionIndexesForProposedSelection_InColumn([self goID], browser, proposedSelectionIndexes, column);
    return (NSIndexSet*)result_;
}

- (id)browser:(NSBrowser*)browser child:(NSInteger)index ofItem:(id)item {
    void* result_ = browserDelegate_Browser_Child_OfItem([self goID], browser, index, item);
    return (id)result_;
}

- (BOOL)browser:(NSBrowser*)browser isLeafItem:(id)item {
    bool result_ = browserDelegate_Browser_IsLeafItem([self goID], browser, item);
    return result_;
}

- (BOOL)browser:(NSBrowser*)browser shouldEditItem:(id)item {
    bool result_ = browserDelegate_Browser_ShouldEditItem([self goID], browser, item);
    return result_;
}

- (id)browser:(NSBrowser*)browser objectValueForItem:(id)item {
    void* result_ = browserDelegate_Browser_ObjectValueForItem([self goID], browser, item);
    return (id)result_;
}

- (void)browser:(NSBrowser*)browser setObjectValue:(id)object forItem:(id)item {
    browserDelegate_Browser_SetObjectValue_ForItem([self goID], browser, object, item);
}

- (id)rootItemForBrowser:(NSBrowser*)browser {
    void* result_ = browserDelegate_RootItemForBrowser([self goID], browser);
    return (id)result_;
}

- (NSViewController*)browser:(NSBrowser*)browser previewViewControllerForLeafItem:(id)item {
    void* result_ = browserDelegate_Browser_PreviewViewControllerForLeafItem([self goID], browser, item);
    return (NSViewController*)result_;
}

- (NSViewController*)browser:(NSBrowser*)browser headerViewControllerForItem:(id)item {
    void* result_ = browserDelegate_Browser_HeaderViewControllerForItem([self goID], browser, item);
    return (NSViewController*)result_;
}

- (void)browser:(NSBrowser*)sender createRowsForColumn:(NSInteger)column inMatrix:(NSMatrix*)matrix {
    browserDelegate_Browser_CreateRowsForColumn_InMatrix([self goID], sender, column, matrix);
}

- (void)browser:(NSBrowser*)sender willDisplayCell:(id)cell atRow:(NSInteger)row column:(NSInteger)column {
    browserDelegate_Browser_WillDisplayCell_AtRow_Column([self goID], sender, cell, row, column);
}

- (void)browser:(NSBrowser*)browser didChangeLastColumn:(NSInteger)oldLastColumn toColumn:(NSInteger)column {
    browserDelegate_Browser_DidChangeLastColumn_ToColumn([self goID], browser, oldLastColumn, column);
}

- (void)browserWillScroll:(NSBrowser*)sender {
    browserDelegate_BrowserWillScroll([self goID], sender);
}

- (void)browserDidScroll:(NSBrowser*)sender {
    browserDelegate_BrowserDidScroll([self goID], sender);
}

- (BOOL)browser:(NSBrowser*)browser canDragRowsWithIndexes:(NSIndexSet*)rowIndexes inColumn:(NSInteger)column withEvent:(NSEvent*)event {
    bool result_ = browserDelegate_Browser_CanDragRowsWithIndexes_InColumn_WithEvent([self goID], browser, rowIndexes, column, event);
    return result_;
}

- (BOOL)browser:(NSBrowser*)browser acceptDrop:(id)info atRow:(NSInteger)row column:(NSInteger)column dropOperation:(NSBrowserDropOperation)dropOperation {
    bool result_ = browserDelegate_Browser_AcceptDrop_AtRow_Column_DropOperation([self goID], browser, info, row, column, dropOperation);
    return result_;
}

- (BOOL)browser:(NSBrowser*)browser writeRowsWithIndexes:(NSIndexSet*)rowIndexes inColumn:(NSInteger)column toPasteboard:(NSPasteboard*)pasteboard {
    bool result_ = browserDelegate_Browser_WriteRowsWithIndexes_InColumn_ToPasteboard([self goID], browser, rowIndexes, column, pasteboard);
    return result_;
}

- (CGFloat)browser:(NSBrowser*)browser shouldSizeColumn:(NSInteger)columnIndex forUserResize:(BOOL)forUserResize toWidth:(CGFloat)suggestedWidth {
    double result_ = browserDelegate_Browser_ShouldSizeColumn_ForUserResize_ToWidth([self goID], browser, columnIndex, forUserResize, suggestedWidth);
    return result_;
}

- (CGFloat)browser:(NSBrowser*)browser sizeToFitWidthOfColumn:(NSInteger)columnIndex {
    double result_ = browserDelegate_Browser_SizeToFitWidthOfColumn([self goID], browser, columnIndex);
    return result_;
}

- (void)browserColumnConfigurationDidChange:(NSNotification*)notification {
    browserDelegate_BrowserColumnConfigurationDidChange([self goID], notification);
}

- (CGFloat)browser:(NSBrowser*)browser heightOfRow:(NSInteger)row inColumn:(NSInteger)columnIndex {
    double result_ = browserDelegate_Browser_HeightOfRow_InColumn([self goID], browser, row, columnIndex);
    return result_;
}

- (BOOL)browser:(NSBrowser*)browser shouldShowCellExpansionForRow:(NSInteger)row column:(NSInteger)column {
    bool result_ = browserDelegate_Browser_ShouldShowCellExpansionForRow_Column([self goID], browser, row, column);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return BrowserDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapBrowserDelegate(uintptr_t goID) {
    NSBrowserDelegateAdaptor* adaptor = [[NSBrowserDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
