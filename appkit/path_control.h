#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_PathControl_Alloc();

void* C_NSPathControl_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSPathControl_InitWithCoder(void* ptr, void* coder);
void* C_NSPathControl_Init(void* ptr);
void* C_NSPathControl_AllocPathControl();
void* C_NSPathControl_NewPathControl();
void* C_NSPathControl_Autorelease(void* ptr);
void* C_NSPathControl_Retain(void* ptr);
void C_NSPathControl_SetDraggingSourceOperationMask_ForLocal(void* ptr, unsigned int mask, bool isLocal);
int C_NSPathControl_PathStyle(void* ptr);
void C_NSPathControl_SetPathStyle(void* ptr, int value);
void* C_NSPathControl_BackgroundColor(void* ptr);
void C_NSPathControl_SetBackgroundColor(void* ptr, void* value);
void* C_NSPathControl_DoubleAction(void* ptr);
void C_NSPathControl_SetDoubleAction(void* ptr, void* value);
void* C_NSPathControl_URL(void* ptr);
void C_NSPathControl_SetURL(void* ptr, void* value);
void* C_NSPathControl_Delegate(void* ptr);
void C_NSPathControl_SetDelegate(void* ptr, void* value);
Array C_NSPathControl_AllowedTypes(void* ptr);
void C_NSPathControl_SetAllowedTypes(void* ptr, Array value);
void* C_NSPathControl_ClickedPathItem(void* ptr);
bool C_NSPathControl_IsEditable(void* ptr);
void C_NSPathControl_SetEditable(void* ptr, bool value);
Array C_NSPathControl_PathItems(void* ptr);
void C_NSPathControl_SetPathItems(void* ptr, Array value);
void* C_NSPathControl_PlaceholderAttributedString(void* ptr);
void C_NSPathControl_SetPlaceholderAttributedString(void* ptr, void* value);
void* C_NSPathControl_PlaceholderString(void* ptr);
void C_NSPathControl_SetPlaceholderString(void* ptr, void* value);
