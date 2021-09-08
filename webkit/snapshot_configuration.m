#import <WebKit/WebKit.h>
#import "snapshot_configuration.h"

void* C_SnapshotConfiguration_Alloc() {
    return [WKSnapshotConfiguration alloc];
}

void* C_WKSnapshotConfiguration_AllocSnapshotConfiguration() {
    WKSnapshotConfiguration* result_ = [WKSnapshotConfiguration alloc];
    return result_;
}

void* C_WKSnapshotConfiguration_Init(void* ptr) {
    WKSnapshotConfiguration* wKSnapshotConfiguration = (WKSnapshotConfiguration*)ptr;
    WKSnapshotConfiguration* result_ = [wKSnapshotConfiguration init];
    return result_;
}

void* C_WKSnapshotConfiguration_NewSnapshotConfiguration() {
    WKSnapshotConfiguration* result_ = [WKSnapshotConfiguration new];
    return result_;
}

void* C_WKSnapshotConfiguration_Autorelease(void* ptr) {
    WKSnapshotConfiguration* wKSnapshotConfiguration = (WKSnapshotConfiguration*)ptr;
    WKSnapshotConfiguration* result_ = [wKSnapshotConfiguration autorelease];
    return result_;
}

void* C_WKSnapshotConfiguration_Retain(void* ptr) {
    WKSnapshotConfiguration* wKSnapshotConfiguration = (WKSnapshotConfiguration*)ptr;
    WKSnapshotConfiguration* result_ = [wKSnapshotConfiguration retain];
    return result_;
}

CGRect C_WKSnapshotConfiguration_Rect(void* ptr) {
    WKSnapshotConfiguration* wKSnapshotConfiguration = (WKSnapshotConfiguration*)ptr;
    CGRect result_ = [wKSnapshotConfiguration rect];
    return result_;
}

void C_WKSnapshotConfiguration_SetRect(void* ptr, CGRect value) {
    WKSnapshotConfiguration* wKSnapshotConfiguration = (WKSnapshotConfiguration*)ptr;
    [wKSnapshotConfiguration setRect:value];
}

void* C_WKSnapshotConfiguration_SnapshotWidth(void* ptr) {
    WKSnapshotConfiguration* wKSnapshotConfiguration = (WKSnapshotConfiguration*)ptr;
    NSNumber* result_ = [wKSnapshotConfiguration snapshotWidth];
    return result_;
}

void C_WKSnapshotConfiguration_SetSnapshotWidth(void* ptr, void* value) {
    WKSnapshotConfiguration* wKSnapshotConfiguration = (WKSnapshotConfiguration*)ptr;
    [wKSnapshotConfiguration setSnapshotWidth:(NSNumber*)value];
}

bool C_WKSnapshotConfiguration_AfterScreenUpdates(void* ptr) {
    WKSnapshotConfiguration* wKSnapshotConfiguration = (WKSnapshotConfiguration*)ptr;
    BOOL result_ = [wKSnapshotConfiguration afterScreenUpdates];
    return result_;
}

void C_WKSnapshotConfiguration_SetAfterScreenUpdates(void* ptr, bool value) {
    WKSnapshotConfiguration* wKSnapshotConfiguration = (WKSnapshotConfiguration*)ptr;
    [wKSnapshotConfiguration setAfterScreenUpdates:value];
}
