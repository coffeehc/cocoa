#import "table_column.h"
#import <AppKit/NSTableColumn.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_TableColumn_Alloc() {
    return [NSTableColumn alloc];
}

void* C_NSTableColumn_InitWithIdentifier(void* ptr, void* identifier) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    NSTableColumn* result_ = [nSTableColumn initWithIdentifier:(NSString*)identifier];
    return result_;
}

void* C_NSTableColumn_InitWithCoder(void* ptr, void* coder) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    NSTableColumn* result_ = [nSTableColumn initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTableColumn_AllocTableColumn() {
    NSTableColumn* result_ = [NSTableColumn alloc];
    return result_;
}

void* C_NSTableColumn_Init(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    NSTableColumn* result_ = [nSTableColumn init];
    return result_;
}

void* C_NSTableColumn_NewTableColumn() {
    NSTableColumn* result_ = [NSTableColumn new];
    return result_;
}

void* C_NSTableColumn_Autorelease(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    NSTableColumn* result_ = [nSTableColumn autorelease];
    return result_;
}

void* C_NSTableColumn_Retain(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    NSTableColumn* result_ = [nSTableColumn retain];
    return result_;
}

void C_NSTableColumn_SizeToFit(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    [nSTableColumn sizeToFit];
}

void* C_NSTableColumn_TableView(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    NSTableView* result_ = [nSTableColumn tableView];
    return result_;
}

void C_NSTableColumn_SetTableView(void* ptr, void* value) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    [nSTableColumn setTableView:(NSTableView*)value];
}

double C_NSTableColumn_Width(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    CGFloat result_ = [nSTableColumn width];
    return result_;
}

void C_NSTableColumn_SetWidth(void* ptr, double value) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    [nSTableColumn setWidth:value];
}

double C_NSTableColumn_MinWidth(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    CGFloat result_ = [nSTableColumn minWidth];
    return result_;
}

void C_NSTableColumn_SetMinWidth(void* ptr, double value) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    [nSTableColumn setMinWidth:value];
}

double C_NSTableColumn_MaxWidth(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    CGFloat result_ = [nSTableColumn maxWidth];
    return result_;
}

void C_NSTableColumn_SetMaxWidth(void* ptr, double value) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    [nSTableColumn setMaxWidth:value];
}

unsigned int C_NSTableColumn_ResizingMask(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    NSTableColumnResizingOptions result_ = [nSTableColumn resizingMask];
    return result_;
}

void C_NSTableColumn_SetResizingMask(void* ptr, unsigned int value) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    [nSTableColumn setResizingMask:value];
}

void* C_NSTableColumn_Title(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    NSString* result_ = [nSTableColumn title];
    return result_;
}

void C_NSTableColumn_SetTitle(void* ptr, void* value) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    [nSTableColumn setTitle:(NSString*)value];
}

void* C_NSTableColumn_HeaderCell(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    NSTableHeaderCell* result_ = [nSTableColumn headerCell];
    return result_;
}

void C_NSTableColumn_SetHeaderCell(void* ptr, void* value) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    [nSTableColumn setHeaderCell:(NSTableHeaderCell*)value];
}

void* C_NSTableColumn_Identifier(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    NSUserInterfaceItemIdentifier result_ = [nSTableColumn identifier];
    return result_;
}

void C_NSTableColumn_SetIdentifier(void* ptr, void* value) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    [nSTableColumn setIdentifier:(NSString*)value];
}

bool C_NSTableColumn_IsEditable(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    BOOL result_ = [nSTableColumn isEditable];
    return result_;
}

void C_NSTableColumn_SetEditable(void* ptr, bool value) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    [nSTableColumn setEditable:value];
}

void* C_NSTableColumn_SortDescriptorPrototype(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    NSSortDescriptor* result_ = [nSTableColumn sortDescriptorPrototype];
    return result_;
}

void C_NSTableColumn_SetSortDescriptorPrototype(void* ptr, void* value) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    [nSTableColumn setSortDescriptorPrototype:(NSSortDescriptor*)value];
}

bool C_NSTableColumn_IsHidden(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    BOOL result_ = [nSTableColumn isHidden];
    return result_;
}

void C_NSTableColumn_SetHidden(void* ptr, bool value) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    [nSTableColumn setHidden:value];
}

void* C_NSTableColumn_HeaderToolTip(void* ptr) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    NSString* result_ = [nSTableColumn headerToolTip];
    return result_;
}

void C_NSTableColumn_SetHeaderToolTip(void* ptr, void* value) {
    NSTableColumn* nSTableColumn = (NSTableColumn*)ptr;
    [nSTableColumn setHeaderToolTip:(NSString*)value];
}
