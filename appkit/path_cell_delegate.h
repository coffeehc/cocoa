#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSPathCellDelegateAdaptor : NSObject <NSPathCellDelegate>
@property (assign) long goID;
@end

void* WrapPathCellDelegate(long goID);
