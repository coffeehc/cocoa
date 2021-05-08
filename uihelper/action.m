#import <Foundation/NSObject.h>
#import "object.h"
#include "_cgo_export.h"
#import <objc/runtime.h>

@interface ActionHandler : NSObject
@property (assign) long goID;
@end

@implementation ActionHandler

- (IBAction)onAction:(id)sender {
	return callAction([self goID], sender);
}
- (void)dealloc {
    deleteAction([self goID]);
    [super dealloc];
}
@end

void* C_NewAction(long id) {
    ActionHandler* handler = [[ActionHandler alloc] init];
	[handler setGoID:id];
	return handler;
}