#import <Appkit/Appkit.h>
#import "tracking_area.h"

void* C_TrackingArea_Alloc() {
    return [NSTrackingArea alloc];
}

void* C_NSTrackingArea_InitWithRect_Options_Owner_UserInfo(void* ptr, CGRect rect, unsigned int options, void* owner, Dictionary userInfo) {
    NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
    NSMutableDictionary* objcUserInfo = [NSMutableDictionary dictionaryWithCapacity:userInfo.len];
    if (userInfo.len > 0) {
    	void** userInfoKeyData = (void**)userInfo.key_data;
    	void** userInfoValueData = (void**)userInfo.value_data;
    	for (int i = 0; i < userInfo.len; i++) {
    		void* kp = userInfoKeyData[i];
    		void* vp = userInfoValueData[i];
    		[objcUserInfo setObject:(id)kp forKey:(id)vp];
    	}
    }
    NSTrackingArea* result_ = [nSTrackingArea initWithRect:rect options:options owner:(id)owner userInfo:objcUserInfo];
    return result_;
}

void* C_NSTrackingArea_AllocTrackingArea() {
    NSTrackingArea* result_ = [NSTrackingArea alloc];
    return result_;
}

void* C_NSTrackingArea_Init(void* ptr) {
    NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
    NSTrackingArea* result_ = [nSTrackingArea init];
    return result_;
}

void* C_NSTrackingArea_NewTrackingArea() {
    NSTrackingArea* result_ = [NSTrackingArea new];
    return result_;
}

void* C_NSTrackingArea_Autorelease(void* ptr) {
    NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
    NSTrackingArea* result_ = [nSTrackingArea autorelease];
    return result_;
}

void* C_NSTrackingArea_Retain(void* ptr) {
    NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
    NSTrackingArea* result_ = [nSTrackingArea retain];
    return result_;
}

unsigned int C_NSTrackingArea_Options(void* ptr) {
    NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
    NSTrackingAreaOptions result_ = [nSTrackingArea options];
    return result_;
}

void* C_NSTrackingArea_Owner(void* ptr) {
    NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
    id result_ = [nSTrackingArea owner];
    return result_;
}

CGRect C_NSTrackingArea_Rect(void* ptr) {
    NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
    NSRect result_ = [nSTrackingArea rect];
    return result_;
}

Dictionary C_NSTrackingArea_UserInfo(void* ptr) {
    NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
    NSDictionary* result_ = [nSTrackingArea userInfo];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		id kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}
