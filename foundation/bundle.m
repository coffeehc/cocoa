#import "bundle.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSBundle.h>
#import <Foundation/NSDictionary.h>

void* C_Bundle_Alloc() {
    return [NSBundle alloc];
}

void* C_NSBundle_BundleWithURL(void* url) {
    NSBundle* result_ = [NSBundle bundleWithURL:(NSURL*)url];
    return result_;
}

void* C_NSBundle_BundleWithPath(void* path) {
    NSBundle* result_ = [NSBundle bundleWithPath:(NSString*)path];
    return result_;
}

void* C_NSBundle_InitWithURL(void* ptr, void* url) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSBundle* result_ = [nSBundle initWithURL:(NSURL*)url];
    return result_;
}

void* C_NSBundle_InitWithPath(void* ptr, void* path) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSBundle* result_ = [nSBundle initWithPath:(NSString*)path];
    return result_;
}

void* C_NSBundle_AllocBundle() {
    NSBundle* result_ = [NSBundle alloc];
    return result_;
}

void* C_NSBundle_Init(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSBundle* result_ = [nSBundle init];
    return result_;
}

void* C_NSBundle_NewBundle() {
    NSBundle* result_ = [NSBundle new];
    return result_;
}

void* C_NSBundle_Autorelease(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSBundle* result_ = [nSBundle autorelease];
    return result_;
}

void* C_NSBundle_Retain(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSBundle* result_ = [nSBundle retain];
    return result_;
}

void* C_NSBundle_BundleWithIdentifier(void* identifier) {
    NSBundle* result_ = [NSBundle bundleWithIdentifier:(NSString*)identifier];
    return result_;
}

void* C_NSBundle_URLForResource_WithExtension_Subdirectory(void* ptr, void* name, void* ext, void* subpath) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSURL* result_ = [nSBundle URLForResource:(NSString*)name withExtension:(NSString*)ext subdirectory:(NSString*)subpath];
    return result_;
}

void* C_NSBundle_URLForResource_WithExtension(void* ptr, void* name, void* ext) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSURL* result_ = [nSBundle URLForResource:(NSString*)name withExtension:(NSString*)ext];
    return result_;
}

Array C_NSBundle_URLsForResourcesWithExtension_Subdirectory(void* ptr, void* ext, void* subpath) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSArray* result_ = [nSBundle URLsForResourcesWithExtension:(NSString*)ext subdirectory:(NSString*)subpath];
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

void* C_NSBundle_URLForResource_WithExtension_Subdirectory_Localization(void* ptr, void* name, void* ext, void* subpath, void* localizationName) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSURL* result_ = [nSBundle URLForResource:(NSString*)name withExtension:(NSString*)ext subdirectory:(NSString*)subpath localization:(NSString*)localizationName];
    return result_;
}

Array C_NSBundle_URLsForResourcesWithExtension_Subdirectory_Localization(void* ptr, void* ext, void* subpath, void* localizationName) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSArray* result_ = [nSBundle URLsForResourcesWithExtension:(NSString*)ext subdirectory:(NSString*)subpath localization:(NSString*)localizationName];
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

void* C_NSBundle_Bundle_URLForResource_WithExtension_Subdirectory_InBundleWithURL(void* name, void* ext, void* subpath, void* bundleURL) {
    NSURL* result_ = [NSBundle URLForResource:(NSString*)name withExtension:(NSString*)ext subdirectory:(NSString*)subpath inBundleWithURL:(NSURL*)bundleURL];
    return result_;
}

Array C_NSBundle_Bundle_URLsForResourcesWithExtension_Subdirectory_InBundleWithURL(void* ext, void* subpath, void* bundleURL) {
    NSArray* result_ = [NSBundle URLsForResourcesWithExtension:(NSString*)ext subdirectory:(NSString*)subpath inBundleWithURL:(NSURL*)bundleURL];
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

void* C_NSBundle_PathForResource_OfType(void* ptr, void* name, void* ext) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle pathForResource:(NSString*)name ofType:(NSString*)ext];
    return result_;
}

void* C_NSBundle_PathForResource_OfType_InDirectory(void* ptr, void* name, void* ext, void* subpath) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle pathForResource:(NSString*)name ofType:(NSString*)ext inDirectory:(NSString*)subpath];
    return result_;
}

void* C_NSBundle_PathForResource_OfType_InDirectory_ForLocalization(void* ptr, void* name, void* ext, void* subpath, void* localizationName) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle pathForResource:(NSString*)name ofType:(NSString*)ext inDirectory:(NSString*)subpath forLocalization:(NSString*)localizationName];
    return result_;
}

Array C_NSBundle_PathsForResourcesOfType_InDirectory(void* ptr, void* ext, void* subpath) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSArray* result_ = [nSBundle pathsForResourcesOfType:(NSString*)ext inDirectory:(NSString*)subpath];
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

Array C_NSBundle_PathsForResourcesOfType_InDirectory_ForLocalization(void* ptr, void* ext, void* subpath, void* localizationName) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSArray* result_ = [nSBundle pathsForResourcesOfType:(NSString*)ext inDirectory:(NSString*)subpath forLocalization:(NSString*)localizationName];
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

void* C_NSBundle_LocalizedStringForKey_Value_Table(void* ptr, void* key, void* value, void* tableName) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle localizedStringForKey:(NSString*)key value:(NSString*)value table:(NSString*)tableName];
    return result_;
}

void* C_NSBundle_URLForAuxiliaryExecutable(void* ptr, void* executableName) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSURL* result_ = [nSBundle URLForAuxiliaryExecutable:(NSString*)executableName];
    return result_;
}

void* C_NSBundle_PathForAuxiliaryExecutable(void* ptr, void* executableName) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle pathForAuxiliaryExecutable:(NSString*)executableName];
    return result_;
}

void* C_NSBundle_ObjectForInfoDictionaryKey(void* ptr, void* key) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    id result_ = [nSBundle objectForInfoDictionaryKey:(NSString*)key];
    return result_;
}

Array C_NSBundle_Bundle_PreferredLocalizationsFromArray(Array localizationsArray) {
    NSMutableArray* objcLocalizationsArray = [NSMutableArray arrayWithCapacity:localizationsArray.len];
    if (localizationsArray.len > 0) {
    	void** localizationsArrayData = (void**)localizationsArray.data;
    	for (int i = 0; i < localizationsArray.len; i++) {
    		void* p = localizationsArrayData[i];
    		[objcLocalizationsArray addObject:(NSString*)p];
    	}
    }
    NSArray* result_ = [NSBundle preferredLocalizationsFromArray:objcLocalizationsArray];
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

Array C_NSBundle_Bundle_PreferredLocalizationsFromArray_ForPreferences(Array localizationsArray, Array preferencesArray) {
    NSMutableArray* objcLocalizationsArray = [NSMutableArray arrayWithCapacity:localizationsArray.len];
    if (localizationsArray.len > 0) {
    	void** localizationsArrayData = (void**)localizationsArray.data;
    	for (int i = 0; i < localizationsArray.len; i++) {
    		void* p = localizationsArrayData[i];
    		[objcLocalizationsArray addObject:(NSString*)p];
    	}
    }
    NSMutableArray* objcPreferencesArray = [NSMutableArray arrayWithCapacity:preferencesArray.len];
    if (preferencesArray.len > 0) {
    	void** preferencesArrayData = (void**)preferencesArray.data;
    	for (int i = 0; i < preferencesArray.len; i++) {
    		void* p = preferencesArrayData[i];
    		[objcPreferencesArray addObject:(NSString*)p];
    	}
    }
    NSArray* result_ = [NSBundle preferredLocalizationsFromArray:objcLocalizationsArray forPreferences:objcPreferencesArray];
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

bool C_NSBundle_Load(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    BOOL result_ = [nSBundle load];
    return result_;
}

bool C_NSBundle_Unload(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    BOOL result_ = [nSBundle unload];
    return result_;
}

void* C_NSBundle_MainBundle() {
    NSBundle* result_ = [NSBundle mainBundle];
    return result_;
}

Array C_NSBundle_Bundle_AllFrameworks() {
    NSArray* result_ = [NSBundle allFrameworks];
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

Array C_NSBundle_AllBundles() {
    NSArray* result_ = [NSBundle allBundles];
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

void* C_NSBundle_ResourceURL(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSURL* result_ = [nSBundle resourceURL];
    return result_;
}

void* C_NSBundle_ExecutableURL(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSURL* result_ = [nSBundle executableURL];
    return result_;
}

void* C_NSBundle_PrivateFrameworksURL(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSURL* result_ = [nSBundle privateFrameworksURL];
    return result_;
}

void* C_NSBundle_SharedFrameworksURL(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSURL* result_ = [nSBundle sharedFrameworksURL];
    return result_;
}

void* C_NSBundle_BuiltInPlugInsURL(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSURL* result_ = [nSBundle builtInPlugInsURL];
    return result_;
}

void* C_NSBundle_SharedSupportURL(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSURL* result_ = [nSBundle sharedSupportURL];
    return result_;
}

void* C_NSBundle_AppStoreReceiptURL(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSURL* result_ = [nSBundle appStoreReceiptURL];
    return result_;
}

void* C_NSBundle_ResourcePath(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle resourcePath];
    return result_;
}

void* C_NSBundle_ExecutablePath(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle executablePath];
    return result_;
}

void* C_NSBundle_PrivateFrameworksPath(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle privateFrameworksPath];
    return result_;
}

void* C_NSBundle_SharedFrameworksPath(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle sharedFrameworksPath];
    return result_;
}

void* C_NSBundle_BuiltInPlugInsPath(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle builtInPlugInsPath];
    return result_;
}

void* C_NSBundle_SharedSupportPath(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle sharedSupportPath];
    return result_;
}

void* C_NSBundle_BundleURL(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSURL* result_ = [nSBundle bundleURL];
    return result_;
}

void* C_NSBundle_BundlePath(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle bundlePath];
    return result_;
}

void* C_NSBundle_BundleIdentifier(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle bundleIdentifier];
    return result_;
}

Dictionary C_NSBundle_InfoDictionary(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSDictionary* result_ = [nSBundle infoDictionary];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
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

Array C_NSBundle_Localizations(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSArray* result_ = [nSBundle localizations];
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

Array C_NSBundle_PreferredLocalizations(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSArray* result_ = [nSBundle preferredLocalizations];
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

void* C_NSBundle_DevelopmentLocalization(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSString* result_ = [nSBundle developmentLocalization];
    return result_;
}

Dictionary C_NSBundle_LocalizedInfoDictionary(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSDictionary* result_ = [nSBundle localizedInfoDictionary];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
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

Array C_NSBundle_ExecutableArchitectures(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    NSArray* result_ = [nSBundle executableArchitectures];
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

bool C_NSBundle_IsLoaded(void* ptr) {
    NSBundle* nSBundle = (NSBundle*)ptr;
    BOOL result_ = [nSBundle isLoaded];
    return result_;
}
