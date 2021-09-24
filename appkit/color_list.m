#import "color_list.h"
#import <AppKit/NSColorList.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_ColorList_Alloc() {
    return [NSColorList alloc];
}

void* C_NSColorList_InitWithName(void* ptr, void* name) {
    NSColorList* nSColorList = (NSColorList*)ptr;
    NSColorList* result_ = [nSColorList initWithName:(NSString*)name];
    return result_;
}

void* C_NSColorList_InitWithName_FromFile(void* ptr, void* name, void* path) {
    NSColorList* nSColorList = (NSColorList*)ptr;
    NSColorList* result_ = [nSColorList initWithName:(NSString*)name fromFile:(NSString*)path];
    return result_;
}

void* C_NSColorList_AllocColorList() {
    NSColorList* result_ = [NSColorList alloc];
    return result_;
}

void* C_NSColorList_Init(void* ptr) {
    NSColorList* nSColorList = (NSColorList*)ptr;
    NSColorList* result_ = [nSColorList init];
    return result_;
}

void* C_NSColorList_NewColorList() {
    NSColorList* result_ = [NSColorList new];
    return result_;
}

void* C_NSColorList_Autorelease(void* ptr) {
    NSColorList* nSColorList = (NSColorList*)ptr;
    NSColorList* result_ = [nSColorList autorelease];
    return result_;
}

void* C_NSColorList_Retain(void* ptr) {
    NSColorList* nSColorList = (NSColorList*)ptr;
    NSColorList* result_ = [nSColorList retain];
    return result_;
}

void* C_NSColorList_ColorListNamed(void* name) {
    NSColorList* result_ = [NSColorList colorListNamed:(NSString*)name];
    return result_;
}

void* C_NSColorList_ColorWithKey(void* ptr, void* key) {
    NSColorList* nSColorList = (NSColorList*)ptr;
    NSColor* result_ = [nSColorList colorWithKey:(NSString*)key];
    return result_;
}

void C_NSColorList_InsertColor_Key_AtIndex(void* ptr, void* color, void* key, unsigned int loc) {
    NSColorList* nSColorList = (NSColorList*)ptr;
    [nSColorList insertColor:(NSColor*)color key:(NSString*)key atIndex:loc];
}

void C_NSColorList_RemoveColorWithKey(void* ptr, void* key) {
    NSColorList* nSColorList = (NSColorList*)ptr;
    [nSColorList removeColorWithKey:(NSString*)key];
}

void C_NSColorList_SetColor_ForKey(void* ptr, void* color, void* key) {
    NSColorList* nSColorList = (NSColorList*)ptr;
    [nSColorList setColor:(NSColor*)color forKey:(NSString*)key];
}

void C_NSColorList_RemoveFile(void* ptr) {
    NSColorList* nSColorList = (NSColorList*)ptr;
    [nSColorList removeFile];
}

Array C_NSColorList_AvailableColorLists() {
    NSArray* result_ = [NSColorList availableColorLists];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void* C_NSColorList_Name(void* ptr) {
    NSColorList* nSColorList = (NSColorList*)ptr;
    NSColorListName result_ = [nSColorList name];
    return result_;
}

bool C_NSColorList_IsEditable(void* ptr) {
    NSColorList* nSColorList = (NSColorList*)ptr;
    BOOL result_ = [nSColorList isEditable];
    return result_;
}

Array C_NSColorList_AllKeys(void* ptr) {
    NSColorList* nSColorList = (NSColorList*)ptr;
    NSArray* result_ = [nSColorList allKeys];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}
