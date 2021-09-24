#import "pdf_panel.h"
#import <AppKit/NSPDFPanel.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_PDFPanel_Alloc() {
    return [NSPDFPanel alloc];
}

void* C_NSPDFPanel_AllocPDFPanel() {
    NSPDFPanel* result_ = [NSPDFPanel alloc];
    return result_;
}

void* C_NSPDFPanel_Init(void* ptr) {
    NSPDFPanel* nSPDFPanel = (NSPDFPanel*)ptr;
    NSPDFPanel* result_ = [nSPDFPanel init];
    return result_;
}

void* C_NSPDFPanel_NewPDFPanel() {
    NSPDFPanel* result_ = [NSPDFPanel new];
    return result_;
}

void* C_NSPDFPanel_Autorelease(void* ptr) {
    NSPDFPanel* nSPDFPanel = (NSPDFPanel*)ptr;
    NSPDFPanel* result_ = [nSPDFPanel autorelease];
    return result_;
}

void* C_NSPDFPanel_Retain(void* ptr) {
    NSPDFPanel* nSPDFPanel = (NSPDFPanel*)ptr;
    NSPDFPanel* result_ = [nSPDFPanel retain];
    return result_;
}

void* C_NSPDFPanel_PDFPanel_Panel() {
    NSPDFPanel* result_ = [NSPDFPanel panel];
    return result_;
}

void* C_NSPDFPanel_AccessoryController(void* ptr) {
    NSPDFPanel* nSPDFPanel = (NSPDFPanel*)ptr;
    NSViewController* result_ = [nSPDFPanel accessoryController];
    return result_;
}

void C_NSPDFPanel_SetAccessoryController(void* ptr, void* value) {
    NSPDFPanel* nSPDFPanel = (NSPDFPanel*)ptr;
    [nSPDFPanel setAccessoryController:(NSViewController*)value];
}

int C_NSPDFPanel_Options(void* ptr) {
    NSPDFPanel* nSPDFPanel = (NSPDFPanel*)ptr;
    NSPDFPanelOptions result_ = [nSPDFPanel options];
    return result_;
}

void C_NSPDFPanel_SetOptions(void* ptr, int value) {
    NSPDFPanel* nSPDFPanel = (NSPDFPanel*)ptr;
    [nSPDFPanel setOptions:value];
}

void* C_NSPDFPanel_DefaultFileName(void* ptr) {
    NSPDFPanel* nSPDFPanel = (NSPDFPanel*)ptr;
    NSString* result_ = [nSPDFPanel defaultFileName];
    return result_;
}

void C_NSPDFPanel_SetDefaultFileName(void* ptr, void* value) {
    NSPDFPanel* nSPDFPanel = (NSPDFPanel*)ptr;
    [nSPDFPanel setDefaultFileName:(NSString*)value];
}
