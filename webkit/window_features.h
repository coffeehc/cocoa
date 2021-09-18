#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_WindowFeatures_Alloc();

void* C_WKWindowFeatures_AllocWindowFeatures();
void* C_WKWindowFeatures_Init(void* ptr);
void* C_WKWindowFeatures_NewWindowFeatures();
void* C_WKWindowFeatures_Autorelease(void* ptr);
void* C_WKWindowFeatures_Retain(void* ptr);
void* C_WKWindowFeatures_AllowsResizing(void* ptr);
void* C_WKWindowFeatures_Height(void* ptr);
void* C_WKWindowFeatures_Width(void* ptr);
void* C_WKWindowFeatures_X(void* ptr);
void* C_WKWindowFeatures_Y(void* ptr);
void* C_WKWindowFeatures_MenuBarVisibility(void* ptr);
void* C_WKWindowFeatures_StatusBarVisibility(void* ptr);
void* C_WKWindowFeatures_ToolbarsVisibility(void* ptr);
