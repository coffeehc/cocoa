#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TabViewItem_Alloc();

void* C_NSTabViewItem_InitWithIdentifier(void* ptr, void* identifier);
void* C_NSTabViewItem_Init(void* ptr);
void C_NSTabViewItem_DrawLabel_InRect(void* ptr, bool shouldTruncateLabel, CGRect labelRect);
CGSize C_NSTabViewItem_SizeOfLabel(void* ptr, bool computeMin);
void* C_NSTabViewItem_TabViewItemWithViewController(void* viewController);
void* C_NSTabViewItem_Label(void* ptr);
void C_NSTabViewItem_SetLabel(void* ptr, void* value);
unsigned int C_NSTabViewItem_TabState(void* ptr);
void* C_NSTabViewItem_Identifier(void* ptr);
void C_NSTabViewItem_SetIdentifier(void* ptr, void* value);
void* C_NSTabViewItem_Color(void* ptr);
void C_NSTabViewItem_SetColor(void* ptr, void* value);
void* C_NSTabViewItem_View(void* ptr);
void C_NSTabViewItem_SetView(void* ptr, void* value);
void* C_NSTabViewItem_InitialFirstResponder(void* ptr);
void C_NSTabViewItem_SetInitialFirstResponder(void* ptr, void* value);
void* C_NSTabViewItem_TabView(void* ptr);
void* C_NSTabViewItem_ToolTip(void* ptr);
void C_NSTabViewItem_SetToolTip(void* ptr, void* value);
void* C_NSTabViewItem_Image(void* ptr);
void C_NSTabViewItem_SetImage(void* ptr, void* value);
void* C_NSTabViewItem_ViewController(void* ptr);
void C_NSTabViewItem_SetViewController(void* ptr, void* value);
