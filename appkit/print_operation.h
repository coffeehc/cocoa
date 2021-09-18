#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_PrintOperation_Alloc();

void* C_NSPrintOperation_AllocPrintOperation();
void* C_NSPrintOperation_Init(void* ptr);
void* C_NSPrintOperation_NewPrintOperation();
void* C_NSPrintOperation_Autorelease(void* ptr);
void* C_NSPrintOperation_Retain(void* ptr);
void* C_NSPrintOperation_PrintOperation_EPSOperationWithView_InsideRect_ToData(void* view, CGRect rect, void* data);
void* C_NSPrintOperation_PrintOperation_EPSOperationWithView_InsideRect_ToData_PrintInfo(void* view, CGRect rect, void* data, void* printInfo);
void* C_NSPrintOperation_PrintOperation_EPSOperationWithView_InsideRect_ToPath_PrintInfo(void* view, CGRect rect, void* path, void* printInfo);
void* C_NSPrintOperation_PrintOperation_PDFOperationWithView_InsideRect_ToData(void* view, CGRect rect, void* data);
void* C_NSPrintOperation_PrintOperation_PDFOperationWithView_InsideRect_ToData_PrintInfo(void* view, CGRect rect, void* data, void* printInfo);
void* C_NSPrintOperation_PrintOperation_PDFOperationWithView_InsideRect_ToPath_PrintInfo(void* view, CGRect rect, void* path, void* printInfo);
void* C_NSPrintOperation_PrintOperationWithView(void* view);
void* C_NSPrintOperation_PrintOperationWithView_PrintInfo(void* view, void* printInfo);
bool C_NSPrintOperation_RunOperation(void* ptr);
void C_NSPrintOperation_CleanUpOperation(void* ptr);
bool C_NSPrintOperation_DeliverResult(void* ptr);
void* C_NSPrintOperation_CreateContext(void* ptr);
void C_NSPrintOperation_DestroyContext(void* ptr);
void* C_NSPrintOperation_PrintOperation_CurrentOperation();
void C_NSPrintOperation_PrintOperation_SetCurrentOperation(void* value);
bool C_NSPrintOperation_IsCopyingOperation(void* ptr);
void* C_NSPrintOperation_PrintInfo(void* ptr);
void C_NSPrintOperation_SetPrintInfo(void* ptr, void* value);
void* C_NSPrintOperation_View(void* ptr);
int C_NSPrintOperation_PreferredRenderingQuality(void* ptr);
bool C_NSPrintOperation_ShowsPrintPanel(void* ptr);
void C_NSPrintOperation_SetShowsPrintPanel(void* ptr, bool value);
bool C_NSPrintOperation_ShowsProgressPanel(void* ptr);
void C_NSPrintOperation_SetShowsProgressPanel(void* ptr, bool value);
void* C_NSPrintOperation_JobTitle(void* ptr);
void C_NSPrintOperation_SetJobTitle(void* ptr, void* value);
void* C_NSPrintOperation_PrintPanel(void* ptr);
void C_NSPrintOperation_SetPrintPanel(void* ptr, void* value);
void* C_NSPrintOperation_PDFPanel(void* ptr);
void C_NSPrintOperation_SetPDFPanel(void* ptr, void* value);
void* C_NSPrintOperation_Context(void* ptr);
int C_NSPrintOperation_CurrentPage(void* ptr);
NSRange C_NSPrintOperation_PageRange(void* ptr);
int C_NSPrintOperation_PageOrder(void* ptr);
void C_NSPrintOperation_SetPageOrder(void* ptr, int value);
bool C_NSPrintOperation_CanSpawnSeparateThread(void* ptr);
void C_NSPrintOperation_SetCanSpawnSeparateThread(void* ptr, bool value);
