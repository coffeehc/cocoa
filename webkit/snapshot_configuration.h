#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_SnapshotConfiguration_Alloc();

void* C_WKSnapshotConfiguration_Init(void* ptr);
CGRect C_WKSnapshotConfiguration_Rect(void* ptr);
void C_WKSnapshotConfiguration_SetRect(void* ptr, CGRect value);
void* C_WKSnapshotConfiguration_SnapshotWidth(void* ptr);
void C_WKSnapshotConfiguration_SetSnapshotWidth(void* ptr, void* value);
bool C_WKSnapshotConfiguration_AfterScreenUpdates(void* ptr);
void C_WKSnapshotConfiguration_SetAfterScreenUpdates(void* ptr, bool value);
