#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSDraggingSourceAdaptor : NSObject <NSDraggingSource>
@property (assign) uintptr_t goID;
@end

void* WrapDraggingSource(uintptr_t goID);
