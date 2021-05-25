#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

@interface NSViewControllerPresentationAnimatorAdaptor : NSObject <NSViewControllerPresentationAnimator>
@property (assign) long goID;
@end

void* WrapViewControllerPresentationAnimator(long goID);
