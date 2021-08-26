#import <Foundation/Foundation.h>
#import "extension_context.h"

void* C_ExtensionContext_Alloc() {
    return [NSExtensionContext alloc];
}

void* C_NSExtensionContext_Init(void* ptr) {
    NSExtensionContext* nSExtensionContext = (NSExtensionContext*)ptr;
    NSExtensionContext* result_ = [nSExtensionContext init];
    return result_;
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

Array C_NSExtensionContext_InputItems(void* ptr) {
    NSExtensionContext* nSExtensionContext = (NSExtensionContext*)ptr;
    NSArray* result_ = [nSExtensionContext inputItems];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}
