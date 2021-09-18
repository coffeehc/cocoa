#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <AppKit/NSAnimation.h>

@interface NSAnimatablePropertyContainerAdaptor : NSObject <NSAnimatablePropertyContainer>
@property (assign) uintptr_t goID;
@end

void* WrapAnimatablePropertyContainer(uintptr_t goID);
