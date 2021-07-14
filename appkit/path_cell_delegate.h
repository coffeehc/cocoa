#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSPathCellDelegateAdaptor : NSObject <NSPathCellDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapPathCellDelegate(uintptr_t goID);
