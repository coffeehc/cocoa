#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSCollectionViewDelegateAdaptor : NSObject <NSCollectionViewDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapCollectionViewDelegate(uintptr_t goID);
