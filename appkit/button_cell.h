#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_ButtonCell_Alloc();

void* C_NSButtonCell_InitWithCoder(void* ptr, void* coder);
void* C_NSButtonCell_InitImageCell(void* ptr, void* image);
void* C_NSButtonCell_InitTextCell(void* ptr, void* _string);
void* C_NSButtonCell_Init(void* ptr);
void* C_NSButtonCell_AllocButtonCell();
void* C_NSButtonCell_NewButtonCell();
void* C_NSButtonCell_Autorelease(void* ptr);
void* C_NSButtonCell_Retain(void* ptr);
void C_NSButtonCell_SetPeriodicDelay_Interval(void* ptr, float delay, float interval);
void C_NSButtonCell_SetButtonType(void* ptr, unsigned int _type);
void C_NSButtonCell_MouseEntered(void* ptr, void* event);
void C_NSButtonCell_MouseExited(void* ptr, void* event);
void C_NSButtonCell_DrawBezelWithFrame_InView(void* ptr, CGRect frame, void* controlView);
void C_NSButtonCell_DrawImage_WithFrame_InView(void* ptr, void* image, CGRect frame, void* controlView);
CGRect C_NSButtonCell_DrawTitle_WithFrame_InView(void* ptr, void* title, CGRect frame, void* controlView);
void* C_NSButtonCell_AlternateTitle(void* ptr);
void C_NSButtonCell_SetAlternateTitle(void* ptr, void* value);
void* C_NSButtonCell_AttributedAlternateTitle(void* ptr);
void C_NSButtonCell_SetAttributedAlternateTitle(void* ptr, void* value);
void* C_NSButtonCell_AttributedTitle(void* ptr);
void C_NSButtonCell_SetAttributedTitle(void* ptr, void* value);
void* C_NSButtonCell_AlternateImage(void* ptr);
void C_NSButtonCell_SetAlternateImage(void* ptr, void* value);
unsigned int C_NSButtonCell_ImagePosition(void* ptr);
void C_NSButtonCell_SetImagePosition(void* ptr, unsigned int value);
unsigned int C_NSButtonCell_ImageScaling(void* ptr);
void C_NSButtonCell_SetImageScaling(void* ptr, unsigned int value);
void C_NSButtonCell_SetKeyEquivalent(void* ptr, void* value);
unsigned int C_NSButtonCell_KeyEquivalentModifierMask(void* ptr);
void C_NSButtonCell_SetKeyEquivalentModifierMask(void* ptr, unsigned int value);
void* C_NSButtonCell_BackgroundColor(void* ptr);
void C_NSButtonCell_SetBackgroundColor(void* ptr, void* value);
unsigned int C_NSButtonCell_BezelStyle(void* ptr);
void C_NSButtonCell_SetBezelStyle(void* ptr, unsigned int value);
bool C_NSButtonCell_ImageDimsWhenDisabled(void* ptr);
void C_NSButtonCell_SetImageDimsWhenDisabled(void* ptr, bool value);
bool C_NSButtonCell_IsTransparent(void* ptr);
void C_NSButtonCell_SetTransparent(void* ptr, bool value);
bool C_NSButtonCell_ShowsBorderOnlyWhileMouseInside(void* ptr);
void C_NSButtonCell_SetShowsBorderOnlyWhileMouseInside(void* ptr, bool value);
unsigned int C_NSButtonCell_HighlightsBy(void* ptr);
void C_NSButtonCell_SetHighlightsBy(void* ptr, unsigned int value);
unsigned int C_NSButtonCell_ShowsStateBy(void* ptr);
void C_NSButtonCell_SetShowsStateBy(void* ptr, unsigned int value);
void* C_NSButtonCell_Sound(void* ptr);
void C_NSButtonCell_SetSound(void* ptr, void* value);
