#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_PDFInfo_Alloc();

void* C_NSPDFInfo_Init(void* ptr);
void* C_NSPDFInfo_URL(void* ptr);
void C_NSPDFInfo_SetURL(void* ptr, void* value);
bool C_NSPDFInfo_IsFileExtensionHidden(void* ptr);
void C_NSPDFInfo_SetFileExtensionHidden(void* ptr, bool value);
Array C_NSPDFInfo_TagNames(void* ptr);
void C_NSPDFInfo_SetTagNames(void* ptr, Array value);
int C_NSPDFInfo_Orientation(void* ptr);
void C_NSPDFInfo_SetOrientation(void* ptr, int value);
CGSize C_NSPDFInfo_PaperSize(void* ptr);
void C_NSPDFInfo_SetPaperSize(void* ptr, CGSize value);
