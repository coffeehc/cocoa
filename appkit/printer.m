#import "printer.h"
#import <AppKit/NSPrinter.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_Printer_Alloc() {
    return [NSPrinter alloc];
}

void* C_NSPrinter_AllocPrinter() {
    NSPrinter* result_ = [NSPrinter alloc];
    return result_;
}

void* C_NSPrinter_Init(void* ptr) {
    NSPrinter* nSPrinter = (NSPrinter*)ptr;
    NSPrinter* result_ = [nSPrinter init];
    return result_;
}

void* C_NSPrinter_NewPrinter() {
    NSPrinter* result_ = [NSPrinter new];
    return result_;
}

void* C_NSPrinter_Autorelease(void* ptr) {
    NSPrinter* nSPrinter = (NSPrinter*)ptr;
    NSPrinter* result_ = [nSPrinter autorelease];
    return result_;
}

void* C_NSPrinter_Retain(void* ptr) {
    NSPrinter* nSPrinter = (NSPrinter*)ptr;
    NSPrinter* result_ = [nSPrinter retain];
    return result_;
}

void* C_NSPrinter_PrinterWithName(void* name) {
    NSPrinter* result_ = [NSPrinter printerWithName:(NSString*)name];
    return result_;
}

void* C_NSPrinter_PrinterWithType(void* _type) {
    NSPrinter* result_ = [NSPrinter printerWithType:(NSString*)_type];
    return result_;
}

CGSize C_NSPrinter_PageSizeForPaper(void* ptr, void* paperName) {
    NSPrinter* nSPrinter = (NSPrinter*)ptr;
    NSSize result_ = [nSPrinter pageSizeForPaper:(NSString*)paperName];
    return result_;
}

Array C_NSPrinter_PrinterNames() {
    NSArray* result_ = [NSPrinter printerNames];
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

Array C_NSPrinter_PrinterTypes() {
    NSArray* result_ = [NSPrinter printerTypes];
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

void* C_NSPrinter_Name(void* ptr) {
    NSPrinter* nSPrinter = (NSPrinter*)ptr;
    NSString* result_ = [nSPrinter name];
    return result_;
}

void* C_NSPrinter_Type(void* ptr) {
    NSPrinter* nSPrinter = (NSPrinter*)ptr;
    NSPrinterTypeName result_ = [nSPrinter type];
    return result_;
}

int C_NSPrinter_LanguageLevel(void* ptr) {
    NSPrinter* nSPrinter = (NSPrinter*)ptr;
    NSInteger result_ = [nSPrinter languageLevel];
    return result_;
}

Dictionary C_NSPrinter_DeviceDescription(void* ptr) {
    NSPrinter* nSPrinter = (NSPrinter*)ptr;
    NSDictionary* result_ = [nSPrinter deviceDescription];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSDeviceDescriptionKey kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}
