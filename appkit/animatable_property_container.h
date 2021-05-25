#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSAnimatablePropertyContainerAdaptor : NSObject <NSAnimatablePropertyContainer>
@property (assign) long goID;
@end

void* WrapAnimatablePropertyContainer(long goID);
