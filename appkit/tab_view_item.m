#import <Appkit/Appkit.h>
#import "tab_view_item.h"

void* C_TabViewItem_Alloc() {
    return [NSTabViewItem alloc];
}

void* C_NSTabViewItem_InitWithIdentifier(void* ptr, void* identifier) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSTabViewItem* result = [nSTabViewItem initWithIdentifier:(id)identifier];
    return result;
}

void* C_NSTabViewItem_Init(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSTabViewItem* result = [nSTabViewItem init];
    return result;
}

void C_NSTabViewItem_DrawLabel_InRect(void* ptr, bool shouldTruncateLabel, CGRect labelRect) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem drawLabel:shouldTruncateLabel inRect:labelRect];
}

CGSize C_NSTabViewItem_SizeOfLabel(void* ptr, bool computeMin) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSSize result = [nSTabViewItem sizeOfLabel:computeMin];
    return result;
}

void* C_NSTabViewItem_TabViewItemWithViewController(void* viewController) {
    NSTabViewItem* result = [NSTabViewItem tabViewItemWithViewController:(NSViewController*)viewController];
    return result;
}

void* C_NSTabViewItem_Label(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSString* result = [nSTabViewItem label];
    return result;
}

void C_NSTabViewItem_SetLabel(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setLabel:(NSString*)value];
}

unsigned int C_NSTabViewItem_TabState(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSTabState result = [nSTabViewItem tabState];
    return result;
}

void* C_NSTabViewItem_Identifier(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    id result = [nSTabViewItem identifier];
    return result;
}

void C_NSTabViewItem_SetIdentifier(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setIdentifier:(id)value];
}

void* C_NSTabViewItem_Color(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSColor* result = [nSTabViewItem color];
    return result;
}

void C_NSTabViewItem_SetColor(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setColor:(NSColor*)value];
}

void* C_NSTabViewItem_View(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSView* result = [nSTabViewItem view];
    return result;
}

void C_NSTabViewItem_SetView(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setView:(NSView*)value];
}

void* C_NSTabViewItem_InitialFirstResponder(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSView* result = [nSTabViewItem initialFirstResponder];
    return result;
}

void C_NSTabViewItem_SetInitialFirstResponder(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setInitialFirstResponder:(NSView*)value];
}

void* C_NSTabViewItem_TabView(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSTabView* result = [nSTabViewItem tabView];
    return result;
}

void* C_NSTabViewItem_ToolTip(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSString* result = [nSTabViewItem toolTip];
    return result;
}

void C_NSTabViewItem_SetToolTip(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setToolTip:(NSString*)value];
}

void* C_NSTabViewItem_Image(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSImage* result = [nSTabViewItem image];
    return result;
}

void C_NSTabViewItem_SetImage(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setImage:(NSImage*)value];
}

void* C_NSTabViewItem_ViewController(void* ptr) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    NSViewController* result = [nSTabViewItem viewController];
    return result;
}

void C_NSTabViewItem_SetViewController(void* ptr, void* value) {
    NSTabViewItem* nSTabViewItem = (NSTabViewItem*)ptr;
    [nSTabViewItem setViewController:(NSViewController*)value];
}
