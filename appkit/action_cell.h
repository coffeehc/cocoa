#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_ActionCell_Alloc();

void* C_NSActionCell_InitImageCell(void* ptr, void* image);
void* C_NSActionCell_InitTextCell(void* ptr, void* _string);
void* C_NSActionCell_Init(void* ptr);
void* C_NSActionCell_InitWithCoder(void* ptr, void* coder);
