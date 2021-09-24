#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Printer_Alloc();

void* C_NSPrinter_AllocPrinter();
void* C_NSPrinter_Init(void* ptr);
void* C_NSPrinter_NewPrinter();
void* C_NSPrinter_Autorelease(void* ptr);
void* C_NSPrinter_Retain(void* ptr);
void* C_NSPrinter_PrinterWithName(void* name);
void* C_NSPrinter_PrinterWithType(void* _type);
CGSize C_NSPrinter_PageSizeForPaper(void* ptr, void* paperName);
Array C_NSPrinter_PrinterNames();
Array C_NSPrinter_PrinterTypes();
void* C_NSPrinter_Name(void* ptr);
void* C_NSPrinter_Type(void* ptr);
int C_NSPrinter_LanguageLevel(void* ptr);
Dictionary C_NSPrinter_DeviceDescription(void* ptr);
