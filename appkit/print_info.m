#import "print_info.h"
#import <AppKit/NSPrintInfo.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_PrintInfo_Alloc() {
    return [NSPrintInfo alloc];
}

void* C_NSPrintInfo_InitWithDictionary(void* ptr, Dictionary attributes) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSMutableDictionary* objcAttributes = [NSMutableDictionary dictionaryWithCapacity:attributes.len];
    if (attributes.len > 0) {
    	void** attributesKeyData = (void**)attributes.key_data;
    	void** attributesValueData = (void**)attributes.value_data;
    	for (int i = 0; i < attributes.len; i++) {
    		void* kp = attributesKeyData[i];
    		void* vp = attributesValueData[i];
    		[objcAttributes setObject:(id)vp forKey:(id<NSCopying>)(NSString*)kp];
    	}
    }
    NSPrintInfo* result_ = [nSPrintInfo initWithDictionary:objcAttributes];
    return result_;
}

void* C_NSPrintInfo_Init(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSPrintInfo* result_ = [nSPrintInfo init];
    return result_;
}

void* C_NSPrintInfo_InitWithCoder(void* ptr, void* coder) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSPrintInfo* result_ = [nSPrintInfo initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSPrintInfo_AllocPrintInfo() {
    NSPrintInfo* result_ = [NSPrintInfo alloc];
    return result_;
}

void* C_NSPrintInfo_NewPrintInfo() {
    NSPrintInfo* result_ = [NSPrintInfo new];
    return result_;
}

void* C_NSPrintInfo_Autorelease(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSPrintInfo* result_ = [nSPrintInfo autorelease];
    return result_;
}

void* C_NSPrintInfo_Retain(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSPrintInfo* result_ = [nSPrintInfo retain];
    return result_;
}

void C_NSPrintInfo_SetUpPrintOperationDefaultValues(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setUpPrintOperationDefaultValues];
}

Dictionary C_NSPrintInfo_Dictionary(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSDictionary* result_ = [nSPrintInfo dictionary];
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

void C_NSPrintInfo_UpdateFromPMPageFormat(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo updateFromPMPageFormat];
}

void C_NSPrintInfo_UpdateFromPMPrintSettings(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo updateFromPMPrintSettings];
}

void C_NSPrintInfo_TakeSettingsFromPDFInfo(void* ptr, void* inPDFInfo) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo takeSettingsFromPDFInfo:(NSPDFInfo*)inPDFInfo];
}

void* C_NSPrintInfo_SharedPrintInfo() {
    NSPrintInfo* result_ = [NSPrintInfo sharedPrintInfo];
    return result_;
}

void C_NSPrintInfo_SetSharedPrintInfo(void* value) {
    [NSPrintInfo setSharedPrintInfo:(NSPrintInfo*)value];
}

CGSize C_NSPrintInfo_PaperSize(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSSize result_ = [nSPrintInfo paperSize];
    return result_;
}

void C_NSPrintInfo_SetPaperSize(void* ptr, CGSize value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setPaperSize:value];
}

double C_NSPrintInfo_TopMargin(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    CGFloat result_ = [nSPrintInfo topMargin];
    return result_;
}

void C_NSPrintInfo_SetTopMargin(void* ptr, double value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setTopMargin:value];
}

double C_NSPrintInfo_BottomMargin(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    CGFloat result_ = [nSPrintInfo bottomMargin];
    return result_;
}

void C_NSPrintInfo_SetBottomMargin(void* ptr, double value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setBottomMargin:value];
}

double C_NSPrintInfo_LeftMargin(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    CGFloat result_ = [nSPrintInfo leftMargin];
    return result_;
}

void C_NSPrintInfo_SetLeftMargin(void* ptr, double value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setLeftMargin:value];
}

double C_NSPrintInfo_RightMargin(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    CGFloat result_ = [nSPrintInfo rightMargin];
    return result_;
}

void C_NSPrintInfo_SetRightMargin(void* ptr, double value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setRightMargin:value];
}

CGRect C_NSPrintInfo_ImageablePageBounds(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSRect result_ = [nSPrintInfo imageablePageBounds];
    return result_;
}

int C_NSPrintInfo_Orientation(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSPaperOrientation result_ = [nSPrintInfo orientation];
    return result_;
}

void C_NSPrintInfo_SetOrientation(void* ptr, int value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setOrientation:value];
}

void* C_NSPrintInfo_PaperName(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSPrinterPaperName result_ = [nSPrintInfo paperName];
    return result_;
}

void C_NSPrintInfo_SetPaperName(void* ptr, void* value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setPaperName:(NSString*)value];
}

void* C_NSPrintInfo_LocalizedPaperName(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSString* result_ = [nSPrintInfo localizedPaperName];
    return result_;
}

unsigned int C_NSPrintInfo_HorizontalPagination(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSPrintingPaginationMode result_ = [nSPrintInfo horizontalPagination];
    return result_;
}

void C_NSPrintInfo_SetHorizontalPagination(void* ptr, unsigned int value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setHorizontalPagination:value];
}

unsigned int C_NSPrintInfo_VerticalPagination(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSPrintingPaginationMode result_ = [nSPrintInfo verticalPagination];
    return result_;
}

void C_NSPrintInfo_SetVerticalPagination(void* ptr, unsigned int value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setVerticalPagination:value];
}

bool C_NSPrintInfo_IsHorizontallyCentered(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    BOOL result_ = [nSPrintInfo isHorizontallyCentered];
    return result_;
}

void C_NSPrintInfo_SetHorizontallyCentered(void* ptr, bool value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setHorizontallyCentered:value];
}

bool C_NSPrintInfo_IsVerticallyCentered(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    BOOL result_ = [nSPrintInfo isVerticallyCentered];
    return result_;
}

void C_NSPrintInfo_SetVerticallyCentered(void* ptr, bool value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setVerticallyCentered:value];
}

void* C_NSPrintInfo_Printer(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSPrinter* result_ = [nSPrintInfo printer];
    return result_;
}

void C_NSPrintInfo_SetPrinter(void* ptr, void* value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setPrinter:(NSPrinter*)value];
}

void* C_NSPrintInfo_JobDisposition(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSPrintJobDispositionValue result_ = [nSPrintInfo jobDisposition];
    return result_;
}

void C_NSPrintInfo_SetJobDisposition(void* ptr, void* value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setJobDisposition:(NSString*)value];
}

bool C_NSPrintInfo_IsSelectionOnly(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    BOOL result_ = [nSPrintInfo isSelectionOnly];
    return result_;
}

void C_NSPrintInfo_SetSelectionOnly(void* ptr, bool value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setSelectionOnly:value];
}

double C_NSPrintInfo_ScalingFactor(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    CGFloat result_ = [nSPrintInfo scalingFactor];
    return result_;
}

void C_NSPrintInfo_SetScalingFactor(void* ptr, double value) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    [nSPrintInfo setScalingFactor:value];
}

Dictionary C_NSPrintInfo_PrintSettings(void* ptr) {
    NSPrintInfo* nSPrintInfo = (NSPrintInfo*)ptr;
    NSDictionary* result_ = [nSPrintInfo printSettings];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSPrintInfoSettingKey kp = [result_Keys objectAtIndex:i];
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

void* C_NSPrintInfo_PrintInfo_DefaultPrinter() {
    NSPrinter* result_ = [NSPrintInfo defaultPrinter];
    return result_;
}
