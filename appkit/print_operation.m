#import <Appkit/Appkit.h>
#import "print_operation.h"

void* C_PrintOperation_Alloc() {
    return [NSPrintOperation alloc];
}

void* C_NSPrintOperation_Init(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    NSPrintOperation* result_ = [nSPrintOperation init];
    return result_;
}

void* C_NSPrintOperation_PrintOperation_EPSOperationWithView_InsideRect_ToData(void* view, CGRect rect, void* data) {
    NSPrintOperation* result_ = [NSPrintOperation EPSOperationWithView:(NSView*)view insideRect:rect toData:(NSData*)data];
    return result_;
}

void* C_NSPrintOperation_PrintOperation_EPSOperationWithView_InsideRect_ToData_PrintInfo(void* view, CGRect rect, void* data, void* printInfo) {
    NSPrintOperation* result_ = [NSPrintOperation EPSOperationWithView:(NSView*)view insideRect:rect toData:(NSData*)data printInfo:(NSPrintInfo*)printInfo];
    return result_;
}

void* C_NSPrintOperation_PrintOperation_EPSOperationWithView_InsideRect_ToPath_PrintInfo(void* view, CGRect rect, void* path, void* printInfo) {
    NSPrintOperation* result_ = [NSPrintOperation EPSOperationWithView:(NSView*)view insideRect:rect toPath:(NSString*)path printInfo:(NSPrintInfo*)printInfo];
    return result_;
}

void* C_NSPrintOperation_PrintOperation_PDFOperationWithView_InsideRect_ToData(void* view, CGRect rect, void* data) {
    NSPrintOperation* result_ = [NSPrintOperation PDFOperationWithView:(NSView*)view insideRect:rect toData:(NSData*)data];
    return result_;
}

void* C_NSPrintOperation_PrintOperation_PDFOperationWithView_InsideRect_ToData_PrintInfo(void* view, CGRect rect, void* data, void* printInfo) {
    NSPrintOperation* result_ = [NSPrintOperation PDFOperationWithView:(NSView*)view insideRect:rect toData:(NSData*)data printInfo:(NSPrintInfo*)printInfo];
    return result_;
}

void* C_NSPrintOperation_PrintOperation_PDFOperationWithView_InsideRect_ToPath_PrintInfo(void* view, CGRect rect, void* path, void* printInfo) {
    NSPrintOperation* result_ = [NSPrintOperation PDFOperationWithView:(NSView*)view insideRect:rect toPath:(NSString*)path printInfo:(NSPrintInfo*)printInfo];
    return result_;
}

void* C_NSPrintOperation_PrintOperationWithView(void* view) {
    NSPrintOperation* result_ = [NSPrintOperation printOperationWithView:(NSView*)view];
    return result_;
}

void* C_NSPrintOperation_PrintOperationWithView_PrintInfo(void* view, void* printInfo) {
    NSPrintOperation* result_ = [NSPrintOperation printOperationWithView:(NSView*)view printInfo:(NSPrintInfo*)printInfo];
    return result_;
}

bool C_NSPrintOperation_RunOperation(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    BOOL result_ = [nSPrintOperation runOperation];
    return result_;
}

void C_NSPrintOperation_CleanUpOperation(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    [nSPrintOperation cleanUpOperation];
}

bool C_NSPrintOperation_DeliverResult(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    BOOL result_ = [nSPrintOperation deliverResult];
    return result_;
}

void* C_NSPrintOperation_CreateContext(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    NSGraphicsContext* result_ = [nSPrintOperation createContext];
    return result_;
}

void C_NSPrintOperation_DestroyContext(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    [nSPrintOperation destroyContext];
}

void* C_NSPrintOperation_PrintOperation_CurrentOperation() {
    NSPrintOperation* result_ = [NSPrintOperation currentOperation];
    return result_;
}

void C_NSPrintOperation_PrintOperation_SetCurrentOperation(void* value) {
    [NSPrintOperation setCurrentOperation:(NSPrintOperation*)value];
}

bool C_NSPrintOperation_IsCopyingOperation(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    BOOL result_ = [nSPrintOperation isCopyingOperation];
    return result_;
}

void* C_NSPrintOperation_PrintInfo(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    NSPrintInfo* result_ = [nSPrintOperation printInfo];
    return result_;
}

void C_NSPrintOperation_SetPrintInfo(void* ptr, void* value) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    [nSPrintOperation setPrintInfo:(NSPrintInfo*)value];
}

void* C_NSPrintOperation_View(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    NSView* result_ = [nSPrintOperation view];
    return result_;
}

int C_NSPrintOperation_PreferredRenderingQuality(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    NSPrintRenderingQuality result_ = [nSPrintOperation preferredRenderingQuality];
    return result_;
}

bool C_NSPrintOperation_ShowsPrintPanel(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    BOOL result_ = [nSPrintOperation showsPrintPanel];
    return result_;
}

void C_NSPrintOperation_SetShowsPrintPanel(void* ptr, bool value) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    [nSPrintOperation setShowsPrintPanel:value];
}

bool C_NSPrintOperation_ShowsProgressPanel(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    BOOL result_ = [nSPrintOperation showsProgressPanel];
    return result_;
}

void C_NSPrintOperation_SetShowsProgressPanel(void* ptr, bool value) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    [nSPrintOperation setShowsProgressPanel:value];
}

void* C_NSPrintOperation_JobTitle(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    NSString* result_ = [nSPrintOperation jobTitle];
    return result_;
}

void C_NSPrintOperation_SetJobTitle(void* ptr, void* value) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    [nSPrintOperation setJobTitle:(NSString*)value];
}

void* C_NSPrintOperation_PrintPanel(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    NSPrintPanel* result_ = [nSPrintOperation printPanel];
    return result_;
}

void C_NSPrintOperation_SetPrintPanel(void* ptr, void* value) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    [nSPrintOperation setPrintPanel:(NSPrintPanel*)value];
}

void* C_NSPrintOperation_PDFPanel(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    NSPDFPanel* result_ = [nSPrintOperation PDFPanel];
    return result_;
}

void C_NSPrintOperation_SetPDFPanel(void* ptr, void* value) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    [nSPrintOperation setPDFPanel:(NSPDFPanel*)value];
}

void* C_NSPrintOperation_Context(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    NSGraphicsContext* result_ = [nSPrintOperation context];
    return result_;
}

int C_NSPrintOperation_CurrentPage(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    NSInteger result_ = [nSPrintOperation currentPage];
    return result_;
}

NSRange C_NSPrintOperation_PageRange(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    NSRange result_ = [nSPrintOperation pageRange];
    return result_;
}

int C_NSPrintOperation_PageOrder(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    NSPrintingPageOrder result_ = [nSPrintOperation pageOrder];
    return result_;
}

void C_NSPrintOperation_SetPageOrder(void* ptr, int value) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    [nSPrintOperation setPageOrder:value];
}

bool C_NSPrintOperation_CanSpawnSeparateThread(void* ptr) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    BOOL result_ = [nSPrintOperation canSpawnSeparateThread];
    return result_;
}

void C_NSPrintOperation_SetCanSpawnSeparateThread(void* ptr, bool value) {
    NSPrintOperation* nSPrintOperation = (NSPrintOperation*)ptr;
    [nSPrintOperation setCanSpawnSeparateThread:value];
}
