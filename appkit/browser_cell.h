#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_BrowserCell_Alloc();

void* C_NSBrowserCell_InitWithCoder(void* ptr, void* coder);
void* C_NSBrowserCell_InitImageCell(void* ptr, void* image);
void* C_NSBrowserCell_InitTextCell(void* ptr, void* _string);
void* C_NSBrowserCell_Init(void* ptr);
void* C_NSBrowserCell_AllocBrowserCell();
void* C_NSBrowserCell_NewBrowserCell();
void* C_NSBrowserCell_Autorelease(void* ptr);
void* C_NSBrowserCell_Retain(void* ptr);
void C_NSBrowserCell_Reset(void* ptr);
void C_NSBrowserCell_Set(void* ptr);
void* C_NSBrowserCell_HighlightColorInView(void* ptr, void* controlView);
void* C_NSBrowserCell_BrowserCell_BranchImage();
void* C_NSBrowserCell_BrowserCell_HighlightedBranchImage();
void* C_NSBrowserCell_AlternateImage(void* ptr);
void C_NSBrowserCell_SetAlternateImage(void* ptr, void* value);
bool C_NSBrowserCell_IsLeaf(void* ptr);
void C_NSBrowserCell_SetLeaf(void* ptr, bool value);
bool C_NSBrowserCell_IsLoaded(void* ptr);
void C_NSBrowserCell_SetLoaded(void* ptr, bool value);
