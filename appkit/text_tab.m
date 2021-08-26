#import <Appkit/Appkit.h>
#import "text_tab.h"

void* C_TextTab_Alloc() {
    return [NSTextTab alloc];
}

void* C_NSTextTab_InitWithTextAlignment_Location_Options(void* ptr, int alignment, double loc, Dictionary options) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    NSMutableDictionary* objcOptions = [[NSMutableDictionary alloc] initWithCapacity: options.len];
    if (options.len > 0) {
    	void** optionsKeyData = (void**)options.key_data;
    	void** optionsValueData = (void**)options.value_data;
    	for (int i = 0; i < options.len; i++) {
    		void* kp = optionsKeyData[i];
    		void* vp = optionsValueData[i];
    		[objcOptions setObject:(NSTextTabOptionKey)(NSString*)kp forKey:(id)(NSString*)vp];
    	}
    }
    NSTextTab* result_ = [nSTextTab initWithTextAlignment:alignment location:loc options:objcOptions];
    return result_;
}

void* C_NSTextTab_Init(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    NSTextTab* result_ = [nSTextTab init];
    return result_;
}

void* C_NSTextTab_TextTab_ColumnTerminatorsForLocale(void* aLocale) {
    NSCharacterSet* result_ = [NSTextTab columnTerminatorsForLocale:(NSLocale*)aLocale];
    return result_;
}

double C_NSTextTab_Location(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    CGFloat result_ = [nSTextTab location];
    return result_;
}

int C_NSTextTab_Alignment(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    NSTextAlignment result_ = [nSTextTab alignment];
    return result_;
}

Dictionary C_NSTextTab_Options(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    NSDictionary* result_ = [nSTextTab options];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSTextTabOptionKey kp = [result_Keys objectAtIndex:i];
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
