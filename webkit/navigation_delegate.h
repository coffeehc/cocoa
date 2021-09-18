#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WKNavigationDelegate.h>

@interface WKNavigationDelegateAdaptor : NSObject <WKNavigationDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapNavigationDelegate(uintptr_t goID);
