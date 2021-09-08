#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Box_Alloc();

void* C_NSBox_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSBox_InitWithCoder(void* ptr, void* coder);
void* C_NSBox_Init(void* ptr);
void* C_NSBox_AllocBox();
void* C_NSBox_NewBox();
void* C_NSBox_Autorelease(void* ptr);
void* C_NSBox_Retain(void* ptr);
void C_NSBox_SetFrameFromContentFrame(void* ptr, CGRect contentFrame);
void C_NSBox_SizeToFit(void* ptr);
CGRect C_NSBox_BorderRect(void* ptr);
unsigned int C_NSBox_BoxType(void* ptr);
void C_NSBox_SetBoxType(void* ptr, unsigned int value);
bool C_NSBox_IsTransparent(void* ptr);
void C_NSBox_SetTransparent(void* ptr, bool value);
void* C_NSBox_Title(void* ptr);
void C_NSBox_SetTitle(void* ptr, void* value);
void* C_NSBox_TitleFont(void* ptr);
void C_NSBox_SetTitleFont(void* ptr, void* value);
unsigned int C_NSBox_TitlePosition(void* ptr);
void C_NSBox_SetTitlePosition(void* ptr, unsigned int value);
void* C_NSBox_TitleCell(void* ptr);
CGRect C_NSBox_TitleRect(void* ptr);
void* C_NSBox_BorderColor(void* ptr);
void C_NSBox_SetBorderColor(void* ptr, void* value);
double C_NSBox_BorderWidth(void* ptr);
void C_NSBox_SetBorderWidth(void* ptr, double value);
double C_NSBox_CornerRadius(void* ptr);
void C_NSBox_SetCornerRadius(void* ptr, double value);
void* C_NSBox_FillColor(void* ptr);
void C_NSBox_SetFillColor(void* ptr, void* value);
void* C_NSBox_ContentView(void* ptr);
void C_NSBox_SetContentView(void* ptr, void* value);
CGSize C_NSBox_ContentViewMargins(void* ptr);
void C_NSBox_SetContentViewMargins(void* ptr, CGSize value);
