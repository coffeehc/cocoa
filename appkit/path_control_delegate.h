#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSPathControlDelegateAdaptor : NSObject <NSPathControlDelegate>
@property (assign) long goID;
@end

void* WrapPathControlDelegate(long goID);
