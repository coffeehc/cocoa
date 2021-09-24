#import "tab_view_item.h"
#import <AppKit/NSTabViewItem.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_TabViewItem_Alloc() {
    return [NSTabViewItem alloc];
}

void* C_NSTabViewItem_InitWithIdentifier(void* ptr, void* identifier) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSTabViewItem* result_ = [nSTabViewItem initWithIdentifier:(id)identifier];
    return result_;
}

void* C_NSTabViewItem_TabViewItemWithViewController(void* viewController) {
    NSTabViewItem* result_ = [NSTabViewItem tabViewItemWithViewController:(NSViewController*)viewController];
    return result_;
}

void* C_NSTabViewItem_AllocTabViewItem() {
    NSTabViewItem* result_ = [NSTabViewItem alloc];
    return result_;
}

void* C_NSTabViewItem_Init(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSTabViewItem* result_ = [nSTabViewItem init];
    return result_;
}

void* C_NSTabViewItem_NewTabViewItem() {
    NSTabViewItem* result_ = [NSTabViewItem new];
    return result_;
}

void* C_NSTabViewItem_Autorelease(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSTabViewItem* result_ = [nSTabViewItem autorelease];
    return result_;
}

void* C_NSTabViewItem_Retain(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSTabViewItem* result_ = [nSTabViewItem retain];
    return result_;
}

void C_NSTabViewItem_DrawLabel_InRect(void* ptr, bool shouldTruncateLabel, CGRect labelRect) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem drawLabel:shouldTruncateLabel inRect:labelRect];
}

CGSize C_NSTabViewItem_SizeOfLabel(void* ptr, bool computeMin) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSSize result_ = [nSTabViewItem sizeOfLabel:computeMin];
    return result_;
}

void* C_NSTabViewItem_Label(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSString* result_ = [nSTabViewItem label];
    return result_;
}

void C_NSTabViewItem_SetLabel(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setLabel:(NSString*)value];
}

unsigned int C_NSTabViewItem_TabState(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSTabState result_ = [nSTabViewItem tabState];
    return result_;
}

void* C_NSTabViewItem_Identifier(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    id result_ = [nSTabViewItem identifier];
    return result_;
}

void C_NSTabViewItem_SetIdentifier(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setIdentifier:(id)value];
}

void* C_NSTabViewItem_Color(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSColor* result_ = [nSTabViewItem color];
    return result_;
}

void C_NSTabViewItem_SetColor(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setColor:(NSColor*)value];
}

void* C_NSTabViewItem_View(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSView* result_ = [nSTabViewItem view];
    return result_;
}

void C_NSTabViewItem_SetView(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setView:(NSView*)value];
}

void* C_NSTabViewItem_InitialFirstResponder(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSView* result_ = [nSTabViewItem initialFirstResponder];
    return result_;
}

void C_NSTabViewItem_SetInitialFirstResponder(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setInitialFirstResponder:(NSView*)value];
}

void* C_NSTabViewItem_TabView(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSTabView* result_ = [nSTabViewItem tabView];
    return result_;
}

void* C_NSTabViewItem_ToolTip(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSString* result_ = [nSTabViewItem toolTip];
    return result_;
}

void C_NSTabViewItem_SetToolTip(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setToolTip:(NSString*)value];
}

void* C_NSTabViewItem_Image(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSImage* result_ = [nSTabViewItem image];
    return result_;
}

void C_NSTabViewItem_SetImage(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setImage:(NSImage*)value];
}

void* C_NSTabViewItem_ViewController(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSViewController* result_ = [nSTabViewItem viewController];
    return result_;
}

void C_NSTabViewItem_SetViewController(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setViewController:(NSViewController*)value];
}
