#import <Foundation/Foundation.h>
#import "extension_context.h"

void* C_ExtensionContext_Alloc() {
	return [NSExtensionContext alloc];
}

void* C_NSExtensionContext_Init(void* ptr) {
	NSExtensionContext* nSExtensionContext = (NSExtensionContext*)ptr;
	NSExtensionContext* result = [nSExtensionContext init];
	return result;
}

void C_NSExtensionContext_CancelRequestWithError(void* ptr, void* error) {
	NSExtensionContext* nSExtensionContext = (NSExtensionContext*)ptr;
	[nSExtensionContext cancelRequestWithError:(NSError*)error];
}

void C_NSExtensionContext_MediaPlayingStarted(void* ptr) {
	NSExtensionContext* nSExtensionContext = (NSExtensionContext*)ptr;
	[nSExtensionContext mediaPlayingStarted];
}

void C_NSExtensionContext_MediaPlayingPaused(void* ptr) {
	NSExtensionContext* nSExtensionContext = (NSExtensionContext*)ptr;
	[nSExtensionContext mediaPlayingPaused];
}

void C_NSExtensionContext_DismissNotificationContentExtension(void* ptr) {
	NSExtensionContext* nSExtensionContext = (NSExtensionContext*)ptr;
	[nSExtensionContext dismissNotificationContentExtension];
}

void C_NSExtensionContext_PerformNotificationDefaultAction(void* ptr) {
	NSExtensionContext* nSExtensionContext = (NSExtensionContext*)ptr;
	[nSExtensionContext performNotificationDefaultAction];
}
