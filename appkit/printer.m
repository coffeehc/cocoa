#import <Appkit/Appkit.h>
#import "printer.h"

void* C_Printer_Alloc() {
    return [NSPrinter alloc];
}

void* C_NSPrinter_Init(void* ptr) {
    NSPrinter* nSPrinter = (NSPrinter*)ptr;
    NSPrinter* result_ = [nSPrinter init];
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
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

Array C_NSPrinter_PrinterTypes() {
    NSArray* result_ = [NSPrinter printerTypes];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
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
