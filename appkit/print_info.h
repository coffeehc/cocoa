#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_PrintInfo_Alloc();

void* C_NSPrintInfo_Init(void* ptr);
void* C_NSPrintInfo_InitWithCoder(void* ptr, void* coder);
void C_NSPrintInfo_SetUpPrintOperationDefaultValues(void* ptr);
void C_NSPrintInfo_UpdateFromPMPageFormat(void* ptr);
void C_NSPrintInfo_UpdateFromPMPrintSettings(void* ptr);
void C_NSPrintInfo_TakeSettingsFromPDFInfo(void* ptr, void* inPDFInfo);
void* C_NSPrintInfo_SharedPrintInfo();
void C_NSPrintInfo_SetSharedPrintInfo(void* value);
CGSize C_NSPrintInfo_PaperSize(void* ptr);
void C_NSPrintInfo_SetPaperSize(void* ptr, CGSize value);
double C_NSPrintInfo_TopMargin(void* ptr);
void C_NSPrintInfo_SetTopMargin(void* ptr, double value);
double C_NSPrintInfo_BottomMargin(void* ptr);
void C_NSPrintInfo_SetBottomMargin(void* ptr, double value);
double C_NSPrintInfo_LeftMargin(void* ptr);
void C_NSPrintInfo_SetLeftMargin(void* ptr, double value);
double C_NSPrintInfo_RightMargin(void* ptr);
void C_NSPrintInfo_SetRightMargin(void* ptr, double value);
CGRect C_NSPrintInfo_ImageablePageBounds(void* ptr);
int C_NSPrintInfo_Orientation(void* ptr);
void C_NSPrintInfo_SetOrientation(void* ptr, int value);
void* C_NSPrintInfo_PaperName(void* ptr);
void C_NSPrintInfo_SetPaperName(void* ptr, void* value);
void* C_NSPrintInfo_LocalizedPaperName(void* ptr);
unsigned int C_NSPrintInfo_HorizontalPagination(void* ptr);
void C_NSPrintInfo_SetHorizontalPagination(void* ptr, unsigned int value);
unsigned int C_NSPrintInfo_VerticalPagination(void* ptr);
void C_NSPrintInfo_SetVerticalPagination(void* ptr, unsigned int value);
bool C_NSPrintInfo_IsHorizontallyCentered(void* ptr);
void C_NSPrintInfo_SetHorizontallyCentered(void* ptr, bool value);
bool C_NSPrintInfo_IsVerticallyCentered(void* ptr);
void C_NSPrintInfo_SetVerticallyCentered(void* ptr, bool value);
void* C_NSPrintInfo_Printer(void* ptr);
void C_NSPrintInfo_SetPrinter(void* ptr, void* value);
void* C_NSPrintInfo_JobDisposition(void* ptr);
void C_NSPrintInfo_SetJobDisposition(void* ptr, void* value);
bool C_NSPrintInfo_IsSelectionOnly(void* ptr);
void C_NSPrintInfo_SetSelectionOnly(void* ptr, bool value);
double C_NSPrintInfo_ScalingFactor(void* ptr);
void C_NSPrintInfo_SetScalingFactor(void* ptr, double value);
void* C_NSPrintInfo_PrintInfo_DefaultPrinter();
