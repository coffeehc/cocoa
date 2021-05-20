#import <Appkit/Appkit.h>
#import "table_view_row_action.h"

void* C_TableViewRowAction_Alloc() {
    return [NSTableViewRowAction alloc];
}

void* C_NSTableViewRowAction_Init(void* ptr) {
    NSTableViewRowAction* nSTableViewRowAction = (NSTableViewRowAction*)ptr;
    NSTableViewRowAction* result_ = [nSTableViewRowAction init];
    return result_;
}

int C_NSTableViewRowAction_Style(void* ptr) {
    NSTableViewRowAction* nSTableViewRowAction = (NSTableViewRowAction*)ptr;
    NSTableViewRowActionStyle result_ = [nSTableViewRowAction style];
    return result_;
}

void* C_NSTableViewRowAction_Title(void* ptr) {
    NSTableViewRowAction* nSTableViewRowAction = (NSTableViewRowAction*)ptr;
    NSString* result_ = [nSTableViewRowAction title];
    return result_;
}

void C_NSTableViewRowAction_SetTitle(void* ptr, void* value) {
    NSTableViewRowAction* nSTableViewRowAction = (NSTableViewRowAction*)ptr;
    [nSTableViewRowAction setTitle:(NSString*)value];
}

void* C_NSTableViewRowAction_BackgroundColor(void* ptr) {
    NSTableViewRowAction* nSTableViewRowAction = (NSTableViewRowAction*)ptr;
    NSColor* result_ = [nSTableViewRowAction backgroundColor];
    return result_;
}

void C_NSTableViewRowAction_SetBackgroundColor(void* ptr, void* value) {
    NSTableViewRowAction* nSTableViewRowAction = (NSTableViewRowAction*)ptr;
    [nSTableViewRowAction setBackgroundColor:(NSColor*)value];
}

void* C_NSTableViewRowAction_Image(void* ptr) {
    NSTableViewRowAction* nSTableViewRowAction = (NSTableViewRowAction*)ptr;
    NSImage* result_ = [nSTableViewRowAction image];
    return result_;
}

void C_NSTableViewRowAction_SetImage(void* ptr, void* value) {
    NSTableViewRowAction* nSTableViewRowAction = (NSTableViewRowAction*)ptr;
    [nSTableViewRowAction setImage:(NSImage*)value];
}
