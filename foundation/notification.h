#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Notification_Alloc();

void* C_NSNotification_Init(void* ptr);
void* C_NSNotification_InitWithCoder(void* ptr, void* coder);
void* C_NSNotification_NotificationWithName_Object(void* aName, void* anObject);
void* C_NSNotification_AllocNotification();
void* C_NSNotification_NewNotification();
void* C_NSNotification_Autorelease(void* ptr);
void* C_NSNotification_Retain(void* ptr);
void* C_NSNotification_Name(void* ptr);
void* C_NSNotification_Object(void* ptr);
