#import <Foundation/Foundation.h>
#import "notification.h"

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
