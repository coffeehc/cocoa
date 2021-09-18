#import "print_panel.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSPrintPanel.h>

void* C_PrintPanel_Alloc() {
    return [NSPrintPanel alloc];
}

void* C_NSPrintPanel_AllocPrintPanel() {
    NSPrintPanel* result_ = [NSPrintPanel alloc];
    return result_;
}

void* C_NSPrintPanel_Init(void* ptr) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    NSPrintPanel* result_ = [nSPrintPanel init];
    return result_;
}

void* C_NSPrintPanel_NewPrintPanel() {
    NSPrintPanel* result_ = [NSPrintPanel new];
    return result_;
}

void* C_NSPrintPanel_Autorelease(void* ptr) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    NSPrintPanel* result_ = [nSPrintPanel autorelease];
    return result_;
}

void* C_NSPrintPanel_Retain(void* ptr) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    NSPrintPanel* result_ = [nSPrintPanel retain];
    return result_;
}

void* C_NSPrintPanel_PrintPanel_() {
    NSPrintPanel* result_ = [NSPrintPanel printPanel];
    return result_;
}

void* C_NSPrintPanel_DefaultButtonTitle(void* ptr) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    NSString* result_ = [nSPrintPanel defaultButtonTitle];
    return result_;
}

void C_NSPrintPanel_SetDefaultButtonTitle(void* ptr, void* defaultButtonTitle) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    [nSPrintPanel setDefaultButtonTitle:(NSString*)defaultButtonTitle];
}

void C_NSPrintPanel_AddAccessoryController(void* ptr, void* accessoryController) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    [nSPrintPanel addAccessoryController:(NSViewController*)accessoryController];
}

void C_NSPrintPanel_RemoveAccessoryController(void* ptr, void* accessoryController) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    [nSPrintPanel removeAccessoryController:(NSViewController*)accessoryController];
}

int C_NSPrintPanel_RunModal(void* ptr) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    NSInteger result_ = [nSPrintPanel runModal];
    return result_;
}

int C_NSPrintPanel_RunModalWithPrintInfo(void* ptr, void* printInfo) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    NSInteger result_ = [nSPrintPanel runModalWithPrintInfo:(NSPrintInfo*)printInfo];
    return result_;
}

void* C_NSPrintPanel_JobStyleHint(void* ptr) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    NSPrintPanelJobStyleHint result_ = [nSPrintPanel jobStyleHint];
    return result_;
}

void C_NSPrintPanel_SetJobStyleHint(void* ptr, void* value) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    [nSPrintPanel setJobStyleHint:(NSString*)value];
}

unsigned int C_NSPrintPanel_Options(void* ptr) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    NSPrintPanelOptions result_ = [nSPrintPanel options];
    return result_;
}

void C_NSPrintPanel_SetOptions(void* ptr, unsigned int value) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    [nSPrintPanel setOptions:value];
}

void* C_NSPrintPanel_HelpAnchor(void* ptr) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    NSHelpAnchorName result_ = [nSPrintPanel helpAnchor];
    return result_;
}

void C_NSPrintPanel_SetHelpAnchor(void* ptr, void* value) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    [nSPrintPanel setHelpAnchor:(NSString*)value];
}

Array C_NSPrintPanel_AccessoryControllers(void* ptr) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    NSArray* result_ = [nSPrintPanel accessoryControllers];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void* C_NSPrintPanel_PrintInfo(void* ptr) {
    NSPrintPanel* nSPrintPanel = (NSPrintPanel*)ptr;
    NSPrintInfo* result_ = [nSPrintPanel printInfo];
    return result_;
}
