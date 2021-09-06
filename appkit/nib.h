#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Nib_Alloc();

void* C_NSNib_InitWithNibNamed_Bundle(void* ptr, void* nibName, void* bundle);
void* C_NSNib_InitWithNibData_Bundle(void* ptr, void* nibData, void* bundle);
void* C_NSNib_Init(void* ptr);
