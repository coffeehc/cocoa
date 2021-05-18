#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Bundle_Alloc();

void* C_NSBundle_InitWithURL(void* ptr, void* url);
void* C_NSBundle_InitWithPath(void* ptr, void* path);
void* C_NSBundle_Init(void* ptr);
void* C_NSBundle_BundleWithURL(void* url);
void* C_NSBundle_BundleWithPath(void* path);
void* C_NSBundle_BundleWithIdentifier(void* identifier);
void* C_NSBundle_URLForResource_WithExtension_Subdirectory(void* ptr, void* name, void* ext, void* subpath);
void* C_NSBundle_URLForResource_WithExtension(void* ptr, void* name, void* ext);
Array C_NSBundle_URLsForResourcesWithExtension_Subdirectory(void* ptr, void* ext, void* subpath);
void* C_NSBundle_URLForResource_WithExtension_Subdirectory_Localization(void* ptr, void* name, void* ext, void* subpath, void* localizationName);
Array C_NSBundle_URLsForResourcesWithExtension_Subdirectory_Localization(void* ptr, void* ext, void* subpath, void* localizationName);
void* C_NSBundle_Bundle_URLForResource_WithExtension_Subdirectory_InBundleWithURL(void* name, void* ext, void* subpath, void* bundleURL);
Array C_NSBundle_Bundle_URLsForResourcesWithExtension_Subdirectory_InBundleWithURL(void* ext, void* subpath, void* bundleURL);
void* C_NSBundle_PathForResource_OfType(void* ptr, void* name, void* ext);
void* C_NSBundle_PathForResource_OfType_InDirectory(void* ptr, void* name, void* ext, void* subpath);
void* C_NSBundle_PathForResource_OfType_InDirectory_ForLocalization(void* ptr, void* name, void* ext, void* subpath, void* localizationName);
Array C_NSBundle_PathsForResourcesOfType_InDirectory(void* ptr, void* ext, void* subpath);
Array C_NSBundle_PathsForResourcesOfType_InDirectory_ForLocalization(void* ptr, void* ext, void* subpath, void* localizationName);
void* C_NSBundle_LocalizedStringForKey_Value_Table(void* ptr, void* key, void* value, void* tableName);
void* C_NSBundle_URLForAuxiliaryExecutable(void* ptr, void* executableName);
void* C_NSBundle_PathForAuxiliaryExecutable(void* ptr, void* executableName);
void* C_NSBundle_ObjectForInfoDictionaryKey(void* ptr, void* key);
Array C_NSBundle_Bundle_PreferredLocalizationsFromArray(Array localizationsArray);
Array C_NSBundle_Bundle_PreferredLocalizationsFromArray_ForPreferences(Array localizationsArray, Array preferencesArray);
bool C_NSBundle_Load(void* ptr);
bool C_NSBundle_Unload(void* ptr);
void* C_NSBundle_MainBundle();
Array C_NSBundle_Bundle_AllFrameworks();
Array C_NSBundle_Bundle_AllBundles();
void* C_NSBundle_ResourceURL(void* ptr);
void* C_NSBundle_ExecutableURL(void* ptr);
void* C_NSBundle_PrivateFrameworksURL(void* ptr);
void* C_NSBundle_SharedFrameworksURL(void* ptr);
void* C_NSBundle_BuiltInPlugInsURL(void* ptr);
void* C_NSBundle_SharedSupportURL(void* ptr);
void* C_NSBundle_AppStoreReceiptURL(void* ptr);
void* C_NSBundle_ResourcePath(void* ptr);
void* C_NSBundle_ExecutablePath(void* ptr);
void* C_NSBundle_PrivateFrameworksPath(void* ptr);
void* C_NSBundle_SharedFrameworksPath(void* ptr);
void* C_NSBundle_BuiltInPlugInsPath(void* ptr);
void* C_NSBundle_SharedSupportPath(void* ptr);
void* C_NSBundle_BundleURL(void* ptr);
void* C_NSBundle_BundlePath(void* ptr);
void* C_NSBundle_BundleIdentifier(void* ptr);
Array C_NSBundle_Localizations(void* ptr);
Array C_NSBundle_PreferredLocalizations(void* ptr);
void* C_NSBundle_DevelopmentLocalization(void* ptr);
Array C_NSBundle_ExecutableArchitectures(void* ptr);
bool C_NSBundle_IsLoaded(void* ptr);
