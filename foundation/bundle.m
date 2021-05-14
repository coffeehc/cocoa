#import <Foundation/Foundation.h>
#import "bundle.h"

void* C_Bundle_Alloc() {
	return [NSBundle alloc];
}

void* C_NSBundle_InitWithURL(void* ptr, void* url) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSBundle* result = [nSBundle initWithURL:(NSURL*)url];
	return result;
}

void* C_NSBundle_InitWithPath(void* ptr, void* path) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSBundle* result = [nSBundle initWithPath:(NSString*)path];
	return result;
}

void* C_NSBundle_Init(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSBundle* result = [nSBundle init];
	return result;
}

void* C_NSBundle_BundleWithURL(void* url) {
	NSBundle* result = [NSBundle bundleWithURL:(NSURL*)url];
	return result;
}

void* C_NSBundle_BundleWithPath(void* path) {
	NSBundle* result = [NSBundle bundleWithPath:(NSString*)path];
	return result;
}

void* C_NSBundle_BundleWithIdentifier(void* identifier) {
	NSBundle* result = [NSBundle bundleWithIdentifier:(NSString*)identifier];
	return result;
}

void* C_NSBundle_URLForResource_WithExtension_Subdirectory(void* ptr, void* name, void* ext, void* subpath) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSURL* result = [nSBundle URLForResource:(NSString*)name withExtension:(NSString*)ext subdirectory:(NSString*)subpath];
	return result;
}

void* C_NSBundle_URLForResource_WithExtension(void* ptr, void* name, void* ext) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSURL* result = [nSBundle URLForResource:(NSString*)name withExtension:(NSString*)ext];
	return result;
}

void* C_NSBundle_URLForResource_WithExtension_Subdirectory_Localization(void* ptr, void* name, void* ext, void* subpath, void* localizationName) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSURL* result = [nSBundle URLForResource:(NSString*)name withExtension:(NSString*)ext subdirectory:(NSString*)subpath localization:(NSString*)localizationName];
	return result;
}

void* C_NSBundle_BundleURLForResource_WithExtension_Subdirectory_InBundleWithURL(void* name, void* ext, void* subpath, void* bundleURL) {
	NSURL* result = [NSBundle URLForResource:(NSString*)name withExtension:(NSString*)ext subdirectory:(NSString*)subpath inBundleWithURL:(NSURL*)bundleURL];
	return result;
}

void* C_NSBundle_PathForResource_OfType(void* ptr, void* name, void* ext) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle pathForResource:(NSString*)name ofType:(NSString*)ext];
	return result;
}

void* C_NSBundle_PathForResource_OfType_InDirectory(void* ptr, void* name, void* ext, void* subpath) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle pathForResource:(NSString*)name ofType:(NSString*)ext inDirectory:(NSString*)subpath];
	return result;
}

void* C_NSBundle_PathForResource_OfType_InDirectory_ForLocalization(void* ptr, void* name, void* ext, void* subpath, void* localizationName) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle pathForResource:(NSString*)name ofType:(NSString*)ext inDirectory:(NSString*)subpath forLocalization:(NSString*)localizationName];
	return result;
}

void* C_NSBundle_BundlePathForResource_OfType_InDirectory(void* name, void* ext, void* bundlePath) {
	NSString* result = [NSBundle pathForResource:(NSString*)name ofType:(NSString*)ext inDirectory:(NSString*)bundlePath];
	return result;
}

void* C_NSBundle_LocalizedStringForKey_Value_Table(void* ptr, void* key, void* value, void* tableName) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle localizedStringForKey:(NSString*)key value:(NSString*)value table:(NSString*)tableName];
	return result;
}

void* C_NSBundle_URLForAuxiliaryExecutable(void* ptr, void* executableName) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSURL* result = [nSBundle URLForAuxiliaryExecutable:(NSString*)executableName];
	return result;
}

void* C_NSBundle_PathForAuxiliaryExecutable(void* ptr, void* executableName) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle pathForAuxiliaryExecutable:(NSString*)executableName];
	return result;
}

void* C_NSBundle_ObjectForInfoDictionaryKey(void* ptr, void* key) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	id result = [nSBundle objectForInfoDictionaryKey:(NSString*)key];
	return result;
}

bool C_NSBundle_Load(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	bool result = [nSBundle load];
	return result;
}

bool C_NSBundle_Unload(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	bool result = [nSBundle unload];
	return result;
}

void* C_NSBundle_ResourceURL(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSURL* result = [nSBundle resourceURL];
	return result;
}

void* C_NSBundle_ExecutableURL(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSURL* result = [nSBundle executableURL];
	return result;
}

void* C_NSBundle_PrivateFrameworksURL(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSURL* result = [nSBundle privateFrameworksURL];
	return result;
}

void* C_NSBundle_SharedFrameworksURL(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSURL* result = [nSBundle sharedFrameworksURL];
	return result;
}

void* C_NSBundle_BuiltInPlugInsURL(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSURL* result = [nSBundle builtInPlugInsURL];
	return result;
}

void* C_NSBundle_SharedSupportURL(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSURL* result = [nSBundle sharedSupportURL];
	return result;
}

void* C_NSBundle_AppStoreReceiptURL(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSURL* result = [nSBundle appStoreReceiptURL];
	return result;
}

void* C_NSBundle_ResourcePath(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle resourcePath];
	return result;
}

void* C_NSBundle_ExecutablePath(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle executablePath];
	return result;
}

void* C_NSBundle_PrivateFrameworksPath(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle privateFrameworksPath];
	return result;
}

void* C_NSBundle_SharedFrameworksPath(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle sharedFrameworksPath];
	return result;
}

void* C_NSBundle_BuiltInPlugInsPath(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle builtInPlugInsPath];
	return result;
}

void* C_NSBundle_SharedSupportPath(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle sharedSupportPath];
	return result;
}

void* C_NSBundle_BundleURL(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSURL* result = [nSBundle bundleURL];
	return result;
}

void* C_NSBundle_BundlePath(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle bundlePath];
	return result;
}

void* C_NSBundle_BundleIdentifier(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle bundleIdentifier];
	return result;
}

void* C_NSBundle_DevelopmentLocalization(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	NSString* result = [nSBundle developmentLocalization];
	return result;
}

bool C_NSBundle_IsLoaded(void* ptr) {
	NSBundle* nSBundle = (NSBundle*)ptr;
	bool result = [nSBundle isLoaded];
	return result;
}
