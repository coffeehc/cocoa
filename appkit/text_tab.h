#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TextTab_Alloc();

void* C_NSTextTab_InitWithTextAlignment_Location_Options(void* ptr, int alignment, double loc, Dictionary options);
void* C_NSTextTab_AllocTextTab();
void* C_NSTextTab_Init(void* ptr);
void* C_NSTextTab_NewTextTab();
void* C_NSTextTab_Autorelease(void* ptr);
void* C_NSTextTab_Retain(void* ptr);
void* C_NSTextTab_TextTab_ColumnTerminatorsForLocale(void* aLocale);
double C_NSTextTab_Location(void* ptr);
int C_NSTextTab_Alignment(void* ptr);
Dictionary C_NSTextTab_Options(void* ptr);
