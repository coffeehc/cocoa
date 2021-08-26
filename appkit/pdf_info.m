#import <Appkit/Appkit.h>
#import "pdf_info.h"

void* C_PDFInfo_Alloc() {
    return [NSPDFInfo alloc];
}

void* C_NSPDFInfo_Init(void* ptr) {
    NSPDFInfo* nSPDFInfo = (NSPDFInfo*)ptr;
    NSPDFInfo* result_ = [nSPDFInfo init];
    return result_;
}

void* C_NSPDFInfo_URL(void* ptr) {
    NSPDFInfo* nSPDFInfo = (NSPDFInfo*)ptr;
    NSURL* result_ = [nSPDFInfo URL];
    return result_;
}

void C_NSPDFInfo_SetURL(void* ptr, void* value) {
    NSPDFInfo* nSPDFInfo = (NSPDFInfo*)ptr;
    [nSPDFInfo setURL:(NSURL*)value];
}

bool C_NSPDFInfo_IsFileExtensionHidden(void* ptr) {
    NSPDFInfo* nSPDFInfo = (NSPDFInfo*)ptr;
    BOOL result_ = [nSPDFInfo isFileExtensionHidden];
    return result_;
}

void C_NSPDFInfo_SetFileExtensionHidden(void* ptr, bool value) {
    NSPDFInfo* nSPDFInfo = (NSPDFInfo*)ptr;
    [nSPDFInfo setFileExtensionHidden:value];
}

Array C_NSPDFInfo_TagNames(void* ptr) {
    NSPDFInfo* nSPDFInfo = (NSPDFInfo*)ptr;
    NSArray* result_ = [nSPDFInfo tagNames];
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

void C_NSPDFInfo_SetTagNames(void* ptr, Array value) {
    NSPDFInfo* nSPDFInfo = (NSPDFInfo*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSString*)(NSString*)p];
    	}
    }
    [nSPDFInfo setTagNames:objcValue];
}

int C_NSPDFInfo_Orientation(void* ptr) {
    NSPDFInfo* nSPDFInfo = (NSPDFInfo*)ptr;
    NSPaperOrientation result_ = [nSPDFInfo orientation];
    return result_;
}

void C_NSPDFInfo_SetOrientation(void* ptr, int value) {
    NSPDFInfo* nSPDFInfo = (NSPDFInfo*)ptr;
    [nSPDFInfo setOrientation:value];
}

CGSize C_NSPDFInfo_PaperSize(void* ptr) {
    NSPDFInfo* nSPDFInfo = (NSPDFInfo*)ptr;
    NSSize result_ = [nSPDFInfo paperSize];
    return result_;
}

void C_NSPDFInfo_SetPaperSize(void* ptr, CGSize value) {
    NSPDFInfo* nSPDFInfo = (NSPDFInfo*)ptr;
    [nSPDFInfo setPaperSize:value];
}

Dictionary C_NSPDFInfo_Attributes(void* ptr) {
    NSPDFInfo* nSPDFInfo = (NSPDFInfo*)ptr;
    NSDictionary* result_ = [nSPDFInfo attributes];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSPrintInfoAttributeKey kp = [result_Keys objectAtIndex:i];
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
