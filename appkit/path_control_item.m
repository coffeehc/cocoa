#import "path_control_item.h"
#import <AppKit/NSPathControlItem.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_PathControlItem_Alloc() {
    return [NSPathControlItem alloc];
}

void* C_NSPathControlItem_AllocPathControlItem() {
    NSPathControlItem* result_ = [NSPathControlItem alloc];
    return result_;
}

void* C_NSPathControlItem_Init(void* ptr) {
    NSPathControlItem* nSPathControlItem = (NSPathControlItem*)ptr;
    NSPathControlItem* result_ = [nSPathControlItem init];
    return result_;
}

void* C_NSPathControlItem_NewPathControlItem() {
    NSPathControlItem* result_ = [NSPathControlItem new];
    return result_;
}

void* C_NSPathControlItem_Autorelease(void* ptr) {
    NSPathControlItem* nSPathControlItem = (NSPathControlItem*)ptr;
    NSPathControlItem* result_ = [nSPathControlItem autorelease];
    return result_;
}

void* C_NSPathControlItem_Retain(void* ptr) {
    NSPathControlItem* nSPathControlItem = (NSPathControlItem*)ptr;
    NSPathControlItem* result_ = [nSPathControlItem retain];
    return result_;
}

void* C_NSPathControlItem_URL(void* ptr) {
    NSPathControlItem* nSPathControlItem = (NSPathControlItem*)ptr;
    NSURL* result_ = [nSPathControlItem URL];
    return result_;
}

void* C_NSPathControlItem_AttributedTitle(void* ptr) {
    NSPathControlItem* nSPathControlItem = (NSPathControlItem*)ptr;
    NSAttributedString* result_ = [nSPathControlItem attributedTitle];
    return result_;
}

void C_NSPathControlItem_SetAttributedTitle(void* ptr, void* value) {
    NSPathControlItem* nSPathControlItem = (NSPathControlItem*)ptr;
    [nSPathControlItem setAttributedTitle:(NSAttributedString*)value];
}

void* C_NSPathControlItem_Image(void* ptr) {
    NSPathControlItem* nSPathControlItem = (NSPathControlItem*)ptr;
    NSImage* result_ = [nSPathControlItem image];
    return result_;
}

void C_NSPathControlItem_SetImage(void* ptr, void* value) {
    NSPathControlItem* nSPathControlItem = (NSPathControlItem*)ptr;
    [nSPathControlItem setImage:(NSImage*)value];
}

void* C_NSPathControlItem_Title(void* ptr) {
    NSPathControlItem* nSPathControlItem = (NSPathControlItem*)ptr;
    NSString* result_ = [nSPathControlItem title];
    return result_;
}

void C_NSPathControlItem_SetTitle(void* ptr, void* value) {
    NSPathControlItem* nSPathControlItem = (NSPathControlItem*)ptr;
    [nSPathControlItem setTitle:(NSString*)value];
}
