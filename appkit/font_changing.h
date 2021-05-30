#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSFontChangingAdaptor : NSObject <NSFontChanging>
@property (assign) long goID;
@end

void* WrapFontChanging(long goID);
