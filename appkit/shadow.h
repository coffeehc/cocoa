#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Shadow_Alloc();

void* C_NSShadow_Init(void* ptr);
void C_NSShadow_Set(void* ptr);
CGSize C_NSShadow_ShadowOffset(void* ptr);
void C_NSShadow_SetShadowOffset(void* ptr, CGSize value);
double C_NSShadow_ShadowBlurRadius(void* ptr);
void C_NSShadow_SetShadowBlurRadius(void* ptr, double value);
void* C_NSShadow_ShadowColor(void* ptr);
void C_NSShadow_SetShadowColor(void* ptr, void* value);
