#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_ExtensionContext_Alloc();

void* C_NSExtensionContext_Init(void* ptr);
void C_NSExtensionContext_CancelRequestWithError(void* ptr, void* error);
void C_NSExtensionContext_MediaPlayingStarted(void* ptr);
void C_NSExtensionContext_MediaPlayingPaused(void* ptr);
void C_NSExtensionContext_DismissNotificationContentExtension(void* ptr);
void C_NSExtensionContext_PerformNotificationDefaultAction(void* ptr);
