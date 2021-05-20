#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TableCellView_Alloc();

void* C_NSTableCellView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSTableCellView_InitWithCoder(void* ptr, void* coder);
void* C_NSTableCellView_Init(void* ptr);
void* C_NSTableCellView_ObjectValue(void* ptr);
void C_NSTableCellView_SetObjectValue(void* ptr, void* value);
int C_NSTableCellView_BackgroundStyle(void* ptr);
void C_NSTableCellView_SetBackgroundStyle(void* ptr, int value);
int C_NSTableCellView_RowSizeStyle(void* ptr);
void C_NSTableCellView_SetRowSizeStyle(void* ptr, int value);
Array C_NSTableCellView_DraggingImageComponents(void* ptr);
