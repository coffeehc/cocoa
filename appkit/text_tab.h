#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TextTab_Alloc();

void* C_NSTextTab_Init(void* ptr);
void* C_NSTextTab_TextTab_ColumnTerminatorsForLocale(void* aLocale);
double C_NSTextTab_Location(void* ptr);
int C_NSTextTab_Alignment(void* ptr);
