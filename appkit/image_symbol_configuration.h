#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_ImageSymbolConfiguration_Alloc();

void* C_NSImageSymbolConfiguration_AllocImageSymbolConfiguration();
void* C_NSImageSymbolConfiguration_Init(void* ptr);
void* C_NSImageSymbolConfiguration_NewImageSymbolConfiguration();
void* C_NSImageSymbolConfiguration_Autorelease(void* ptr);
void* C_NSImageSymbolConfiguration_Retain(void* ptr);
void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithPointSize_Weight(double pointSize, double weight);
void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithPointSize_Weight_Scale(double pointSize, double weight, int scale);
void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithTextStyle(void* style);
void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithTextStyle_Scale(void* style, int scale);
void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithScale(int scale);
