#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_Navigation_Alloc();

void* C_WKNavigation_Init(void* ptr);
int C_WKNavigation_EffectiveContentMode(void* ptr);
