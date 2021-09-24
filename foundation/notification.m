#import "notification.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <Foundation/NSNotification.h>

void* C_Notification_Alloc() {
    return [NSNotification alloc];
}

void* C_NSNotification_Init(void* ptr) {
    NSNotification* nSNotification = (NSNotification*)ptr;
    NSNotification* result_ = [nSNotification init];
    return result_;
}

void* C_NSNotification_InitWithCoder(void* ptr, void* coder) {
    NSNotification* nSNotification = (NSNotification*)ptr;
    NSNotification* result_ = [nSNotification initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSNotification_NotificationWithName_Object(void* aName, void* anObject) {
    NSNotification* result_ = [NSNotification notificationWithName:(NSString*)aName object:(id)anObject];
    return result_;
}

void* C_NSNotification_AllocNotification() {
    NSNotification* result_ = [NSNotification alloc];
    return result_;
}

void* C_NSNotification_NewNotification() {
    NSNotification* result_ = [NSNotification new];
    return result_;
}

void* C_NSNotification_Autorelease(void* ptr) {
    NSNotification* nSNotification = (NSNotification*)ptr;
    NSNotification* result_ = [nSNotification autorelease];
    return result_;
}

void* C_NSNotification_Retain(void* ptr) {
    NSNotification* nSNotification = (NSNotification*)ptr;
    NSNotification* result_ = [nSNotification retain];
    return result_;
}

void* C_NSNotification_Name(void* ptr) {
    NSNotification* nSNotification = (NSNotification*)ptr;
    NSNotificationName result_ = [nSNotification name];
    return result_;
}

void* C_NSNotification_Object(void* ptr) {
    NSNotification* nSNotification = (NSNotification*)ptr;
    id result_ = [nSNotification object];
    return result_;
}
