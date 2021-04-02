#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

NSRect Box_BorderRect(void* ptr);
unsigned long Box_BoxType(void* ptr);
void Box_SetBoxType(void* ptr, unsigned long boxType);
bool Box_IsTransparent(void* ptr);
void Box_SetTransparent(void* ptr, bool transparent);
const char* Box_Title(void* ptr);
void Box_SetTitle(void* ptr, const char* title);
void* Box_TitleFont(void* ptr);
void Box_SetTitleFont(void* ptr, void* titleFont);
unsigned long Box_TitlePosition(void* ptr);
void Box_SetTitlePosition(void* ptr, unsigned long titlePosition);
void* Box_TitleCell(void* ptr);
NSRect Box_TitleRect(void* ptr);
void* Box_BorderColor(void* ptr);
void Box_SetBorderColor(void* ptr, void* borderColor);
double Box_BorderWidth(void* ptr);
void Box_SetBorderWidth(void* ptr, double borderWidth);
double Box_CornerRadius(void* ptr);
void Box_SetCornerRadius(void* ptr, double cornerRadius);
void* Box_FillColor(void* ptr);
void Box_SetFillColor(void* ptr, void* fillColor);
void* Box_ContentView(void* ptr);
void Box_SetContentView(void* ptr, void* contentView);
NSSize Box_ContentViewMargins(void* ptr);
void Box_SetContentViewMargins(void* ptr, NSSize contentViewMargins);

void* Box_NewBox(NSRect frame);
void Box_SetFrameFromContentFrame(void* ptr, NSRect contentFrame);
void Box_SizeToFit(void* ptr);
