#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* ColorWell_Color(void* ptr);
void ColorWell_SetColor(void* ptr, void* color);
bool ColorWell_IsActive(void* ptr);
bool ColorWell_IsBordered(void* ptr);
void ColorWell_SetBordered(void* ptr, bool bordered);

void* ColorWell_NewColorWell(NSRect frame);
void ColorWell_TakeColorFrom(void* ptr, void* sender);
void ColorWell_Activate(void* ptr, bool exclusive);
void ColorWell_Deactivate(void* ptr);
void ColorWell_DrawWellInside(void* ptr, NSRect insideRect);
