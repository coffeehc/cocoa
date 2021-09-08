#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_PathComponentCell_Alloc();

void* C_NSPathComponentCell_InitTextCell(void* ptr, void* _string);
void* C_NSPathComponentCell_InitWithCoder(void* ptr, void* coder);
void* C_NSPathComponentCell_Init(void* ptr);
void* C_NSPathComponentCell_AllocPathComponentCell();
void* C_NSPathComponentCell_NewPathComponentCell();
void* C_NSPathComponentCell_Autorelease(void* ptr);
void* C_NSPathComponentCell_Retain(void* ptr);
void* C_NSPathComponentCell_URL(void* ptr);
void C_NSPathComponentCell_SetURL(void* ptr, void* value);
