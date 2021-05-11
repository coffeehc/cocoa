#import <Foundation/Foundation.h>
#import "notification.h"

void* C_Notification_Alloc() {
	return [NSNotification alloc];
}

void* C_NSNotification_Init(void* ptr) {
	NSNotification* nSNotification = (NSNotification*)ptr;
	NSNotification* result = [nSNotification init];
	return result;
}

void* C_NSNotification_InitWithCoder(void* ptr, void* coder) {
	NSNotification* nSNotification = (NSNotification*)ptr;
	NSNotification* result = [nSNotification initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSNotification_NotificationWithName_Object(void* aName, void* anObject) {
	NSNotification* result = [NSNotification notificationWithName:(NSString*)aName object:(id)anObject];
	return result;
}

void* C_NSNotification_Name(void* ptr) {
	NSNotification* nSNotification = (NSNotification*)ptr;
	NSString* result = [nSNotification name];
	return result;
}

void* C_NSNotification_Object(void* ptr) {
	NSNotification* nSNotification = (NSNotification*)ptr;
	id result = [nSNotification object];
	return result;
}
