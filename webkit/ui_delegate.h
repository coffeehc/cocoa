#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WKUIDelegate.h>

@interface WKUIDelegateAdaptor : NSObject <WKUIDelegate>
@property (assign) uintptr_t goID;
@end

void* WrapUIDelegate(uintptr_t goID);
