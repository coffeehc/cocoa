#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSFontPanel.h>

@interface NSFontChangingAdaptor : NSObject <NSFontChanging>
@property (assign) uintptr_t goID;
@end

void* WrapFontChanging(uintptr_t goID);
