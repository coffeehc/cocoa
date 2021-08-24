#import <Appkit/Appkit.h>
#import "appkit_custom.h"

void* C_NSImage_CGImageForProposedRect_Context_Hints(void* ptr) {
    NSImage* image = (NSImage*)ptr;
    return [image CGImageForProposedRect:NULL context:nil hints:nil];
}