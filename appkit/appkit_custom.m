#import "appkit_custom.h"
#include "_cgo_export.h"
#import <Appkit/NSImage.h>
#import <Appkit/NSSavePanel.h>

void* C_NSImage_CGImageForProposedRect_Context_Hints(void* ptr) {
    NSImage* image = (NSImage*)ptr;
    return [image CGImageForProposedRect:NULL context:nil hints:nil];
}

void SavePanel_BeginWithCompletionHandler(void* ptr, uintptr_t handler) {
    NSSavePanel* panel = (NSSavePanel*)ptr;
    [panel beginWithCompletionHandler:^(NSModalResponse result) {
        callModalResponseHandler(handler, result);
    }];
}
void SavePanel_BeginSheetModalForWindow(void* ptr, void* winPtr, uintptr_t handler) {
    NSSavePanel* panel = (NSSavePanel*)ptr;
    [panel beginSheetModalForWindow: (NSWindow*)winPtr completionHandler:^(NSModalResponse result) {
        callModalResponseHandler(handler, result);
    }];
}