#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TableViewRowAction_Alloc();

void* C_NSTableViewRowAction_Init(void* ptr);
int C_NSTableViewRowAction_Style(void* ptr);
void* C_NSTableViewRowAction_Title(void* ptr);
void C_NSTableViewRowAction_SetTitle(void* ptr, void* value);
void* C_NSTableViewRowAction_BackgroundColor(void* ptr);
void C_NSTableViewRowAction_SetBackgroundColor(void* ptr, void* value);
void* C_NSTableViewRowAction_Image(void* ptr);
void C_NSTableViewRowAction_SetImage(void* ptr, void* value);
