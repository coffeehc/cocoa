#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "secure_text_field.h"
#import "secure_text_field_delegate.h"

@implementation MyNSSecureTextFieldDelegate

@end
void SecureTextField_SetDelegate(void *ptr, long goID) {
	NSSecureTextField* secureTextField = (NSSecureTextField*)ptr;
	MyNSSecureTextFieldDelegate* delegate = [[MyNSSecureTextFieldDelegate alloc] init];
	[delegate setGoID:goID];
	[secureTextField setDelegate:delegate];
}

void* SecureTextField_NewSecureTextField(NSRect frame) {
	NSSecureTextField* secureTextField = [NSSecureTextField alloc];
	return [[secureTextField initWithFrame:frame] autorelease];
}
