#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TableColumn_Alloc();

void* C_NSTableColumn_InitWithIdentifier(void* ptr, void* identifier);
void* C_NSTableColumn_InitWithCoder(void* ptr, void* coder);
void* C_NSTableColumn_AllocTableColumn();
void* C_NSTableColumn_Init(void* ptr);
void* C_NSTableColumn_NewTableColumn();
void* C_NSTableColumn_Autorelease(void* ptr);
void* C_NSTableColumn_Retain(void* ptr);
void C_NSTableColumn_SizeToFit(void* ptr);
void* C_NSTableColumn_TableView(void* ptr);
void C_NSTableColumn_SetTableView(void* ptr, void* value);
double C_NSTableColumn_Width(void* ptr);
void C_NSTableColumn_SetWidth(void* ptr, double value);
double C_NSTableColumn_MinWidth(void* ptr);
void C_NSTableColumn_SetMinWidth(void* ptr, double value);
double C_NSTableColumn_MaxWidth(void* ptr);
void C_NSTableColumn_SetMaxWidth(void* ptr, double value);
unsigned int C_NSTableColumn_ResizingMask(void* ptr);
void C_NSTableColumn_SetResizingMask(void* ptr, unsigned int value);
void* C_NSTableColumn_Title(void* ptr);
void C_NSTableColumn_SetTitle(void* ptr, void* value);
void* C_NSTableColumn_HeaderCell(void* ptr);
void C_NSTableColumn_SetHeaderCell(void* ptr, void* value);
void* C_NSTableColumn_Identifier(void* ptr);
void C_NSTableColumn_SetIdentifier(void* ptr, void* value);
bool C_NSTableColumn_IsEditable(void* ptr);
void C_NSTableColumn_SetEditable(void* ptr, bool value);
void* C_NSTableColumn_SortDescriptorPrototype(void* ptr);
void C_NSTableColumn_SetSortDescriptorPrototype(void* ptr, void* value);
bool C_NSTableColumn_IsHidden(void* ptr);
void C_NSTableColumn_SetHidden(void* ptr, bool value);
void* C_NSTableColumn_HeaderToolTip(void* ptr);
void C_NSTableColumn_SetHeaderToolTip(void* ptr, void* value);
