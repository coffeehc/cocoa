#import "touch_bar_item.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSTouchBarItem.h>

void* C_TouchBarItem_Alloc() {
    return [NSTouchBarItem alloc];
}

void* C_NSTouchBarItem_InitWithIdentifier(void* ptr, void* identifier) {
    NSTouchBarItem* nSTouchBarItem = (NSTouchBarItem*)ptr;
    NSTouchBarItem* result_ = [nSTouchBarItem initWithIdentifier:(NSString*)identifier];
    return result_;
}

void* C_NSTouchBarItem_InitWithCoder(void* ptr, void* coder) {
    NSTouchBarItem* nSTouchBarItem = (NSTouchBarItem*)ptr;
    NSTouchBarItem* result_ = [nSTouchBarItem initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTouchBarItem_AllocTouchBarItem() {
    NSTouchBarItem* result_ = [NSTouchBarItem alloc];
    return result_;
}

void* C_NSTouchBarItem_Autorelease(void* ptr) {
    NSTouchBarItem* nSTouchBarItem = (NSTouchBarItem*)ptr;
    NSTouchBarItem* result_ = [nSTouchBarItem autorelease];
    return result_;
}

void* C_NSTouchBarItem_Retain(void* ptr) {
    NSTouchBarItem* nSTouchBarItem = (NSTouchBarItem*)ptr;
    NSTouchBarItem* result_ = [nSTouchBarItem retain];
    return result_;
}

void* C_NSTouchBarItem_Identifier(void* ptr) {
    NSTouchBarItem* nSTouchBarItem = (NSTouchBarItem*)ptr;
    NSTouchBarItemIdentifier result_ = [nSTouchBarItem identifier];
    return result_;
}

float C_NSTouchBarItem_VisibilityPriority(void* ptr) {
    NSTouchBarItem* nSTouchBarItem = (NSTouchBarItem*)ptr;
    NSTouchBarItemPriority result_ = [nSTouchBarItem visibilityPriority];
    return result_;
}

void C_NSTouchBarItem_SetVisibilityPriority(void* ptr, float value) {
    NSTouchBarItem* nSTouchBarItem = (NSTouchBarItem*)ptr;
    [nSTouchBarItem setVisibilityPriority:value];
}

bool C_NSTouchBarItem_IsVisible(void* ptr) {
    NSTouchBarItem* nSTouchBarItem = (NSTouchBarItem*)ptr;
    BOOL result_ = [nSTouchBarItem isVisible];
    return result_;
}

void* C_NSTouchBarItem_CustomizationLabel(void* ptr) {
    NSTouchBarItem* nSTouchBarItem = (NSTouchBarItem*)ptr;
    NSString* result_ = [nSTouchBarItem customizationLabel];
    return result_;
}

void* C_NSTouchBarItem_ViewController(void* ptr) {
    NSTouchBarItem* nSTouchBarItem = (NSTouchBarItem*)ptr;
    NSViewController* result_ = [nSTouchBarItem viewController];
    return result_;
}

void* C_NSTouchBarItem_View(void* ptr) {
    NSTouchBarItem* nSTouchBarItem = (NSTouchBarItem*)ptr;
    NSView* result_ = [nSTouchBarItem view];
    return result_;
}
