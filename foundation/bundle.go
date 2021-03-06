package foundation

// #include "bundle.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Bundle interface {
	objc.Object
	URLForResource_WithExtension_Subdirectory(name string, ext string, subpath string) URL
	URLForResource_WithExtension(name string, ext string) URL
	URLsForResourcesWithExtension_Subdirectory(ext string, subpath string) []URL
	URLForResource_WithExtension_Subdirectory_Localization(name string, ext string, subpath string, localizationName string) URL
	URLsForResourcesWithExtension_Subdirectory_Localization(ext string, subpath string, localizationName string) []URL
	PathForResource_OfType(name string, ext string) string
	PathForResource_OfType_InDirectory(name string, ext string, subpath string) string
	PathForResource_OfType_InDirectory_ForLocalization(name string, ext string, subpath string, localizationName string) string
	PathsForResourcesOfType_InDirectory(ext string, subpath string) []string
	PathsForResourcesOfType_InDirectory_ForLocalization(ext string, subpath string, localizationName string) []string
	LocalizedStringForKey_Value_Table(key string, value string, tableName string) string
	URLForAuxiliaryExecutable(executableName string) URL
	PathForAuxiliaryExecutable(executableName string) string
	ObjectForInfoDictionaryKey(key string) objc.Object
	Load() bool
	Unload() bool
	ResourceURL() URL
	ExecutableURL() URL
	PrivateFrameworksURL() URL
	SharedFrameworksURL() URL
	BuiltInPlugInsURL() URL
	SharedSupportURL() URL
	AppStoreReceiptURL() URL
	ResourcePath() string
	ExecutablePath() string
	PrivateFrameworksPath() string
	SharedFrameworksPath() string
	BuiltInPlugInsPath() string
	SharedSupportPath() string
	BundleURL() URL
	BundlePath() string
	BundleIdentifier() string
	InfoDictionary() map[string]objc.Object
	Localizations() []string
	PreferredLocalizations() []string
	DevelopmentLocalization() string
	LocalizedInfoDictionary() map[string]objc.Object
	ExecutableArchitectures() []Number
	IsLoaded() bool
}

type NSBundle struct {
	objc.NSObject
}

func MakeBundle(ptr unsafe.Pointer) NSBundle {
	return NSBundle{
		NSObject: objc.MakeObject(ptr),
	}
}

func BundleWithURL(url URL) NSBundle {
	result_ := C.C_NSBundle_BundleWithURL(objc.ExtractPtr(url))
	return MakeBundle(result_)
}

func BundleWithPath(path string) NSBundle {
	result_ := C.C_NSBundle_BundleWithPath(NewString(path).Ptr())
	return MakeBundle(result_)
}

func (n NSBundle) InitWithURL(url URL) NSBundle {
	result_ := C.C_NSBundle_InitWithURL(n.Ptr(), objc.ExtractPtr(url))
	return MakeBundle(result_)
}

func (n NSBundle) InitWithPath(path string) NSBundle {
	result_ := C.C_NSBundle_InitWithPath(n.Ptr(), NewString(path).Ptr())
	return MakeBundle(result_)
}

func AllocBundle() NSBundle {
	result_ := C.C_NSBundle_AllocBundle()
	return MakeBundle(result_)
}

func (n NSBundle) Init() NSBundle {
	result_ := C.C_NSBundle_Init(n.Ptr())
	return MakeBundle(result_)
}

func NewBundle() NSBundle {
	result_ := C.C_NSBundle_NewBundle()
	return MakeBundle(result_)
}

func (n NSBundle) Autorelease() NSBundle {
	result_ := C.C_NSBundle_Autorelease(n.Ptr())
	return MakeBundle(result_)
}

func (n NSBundle) Retain() NSBundle {
	result_ := C.C_NSBundle_Retain(n.Ptr())
	return MakeBundle(result_)
}

func BundleWithIdentifier(identifier string) Bundle {
	result_ := C.C_NSBundle_BundleWithIdentifier(NewString(identifier).Ptr())
	return MakeBundle(result_)
}

func (n NSBundle) URLForResource_WithExtension_Subdirectory(name string, ext string, subpath string) URL {
	result_ := C.C_NSBundle_URLForResource_WithExtension_Subdirectory(n.Ptr(), NewString(name).Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr())
	return MakeURL(result_)
}

func (n NSBundle) URLForResource_WithExtension(name string, ext string) URL {
	result_ := C.C_NSBundle_URLForResource_WithExtension(n.Ptr(), NewString(name).Ptr(), NewString(ext).Ptr())
	return MakeURL(result_)
}

func (n NSBundle) URLsForResourcesWithExtension_Subdirectory(ext string, subpath string) []URL {
	result_ := C.C_NSBundle_URLsForResourcesWithExtension_Subdirectory(n.Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]URL, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeURL(r)
	}
	return goResult_
}

func (n NSBundle) URLForResource_WithExtension_Subdirectory_Localization(name string, ext string, subpath string, localizationName string) URL {
	result_ := C.C_NSBundle_URLForResource_WithExtension_Subdirectory_Localization(n.Ptr(), NewString(name).Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr(), NewString(localizationName).Ptr())
	return MakeURL(result_)
}

func (n NSBundle) URLsForResourcesWithExtension_Subdirectory_Localization(ext string, subpath string, localizationName string) []URL {
	result_ := C.C_NSBundle_URLsForResourcesWithExtension_Subdirectory_Localization(n.Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr(), NewString(localizationName).Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]URL, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeURL(r)
	}
	return goResult_
}

func Bundle_URLForResource_WithExtension_Subdirectory_InBundleWithURL(name string, ext string, subpath string, bundleURL URL) URL {
	result_ := C.C_NSBundle_Bundle_URLForResource_WithExtension_Subdirectory_InBundleWithURL(NewString(name).Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr(), objc.ExtractPtr(bundleURL))
	return MakeURL(result_)
}

func Bundle_URLsForResourcesWithExtension_Subdirectory_InBundleWithURL(ext string, subpath string, bundleURL URL) []URL {
	result_ := C.C_NSBundle_Bundle_URLsForResourcesWithExtension_Subdirectory_InBundleWithURL(NewString(ext).Ptr(), NewString(subpath).Ptr(), objc.ExtractPtr(bundleURL))
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]URL, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeURL(r)
	}
	return goResult_
}

func (n NSBundle) PathForResource_OfType(name string, ext string) string {
	result_ := C.C_NSBundle_PathForResource_OfType(n.Ptr(), NewString(name).Ptr(), NewString(ext).Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) PathForResource_OfType_InDirectory(name string, ext string, subpath string) string {
	result_ := C.C_NSBundle_PathForResource_OfType_InDirectory(n.Ptr(), NewString(name).Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) PathForResource_OfType_InDirectory_ForLocalization(name string, ext string, subpath string, localizationName string) string {
	result_ := C.C_NSBundle_PathForResource_OfType_InDirectory_ForLocalization(n.Ptr(), NewString(name).Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr(), NewString(localizationName).Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) PathsForResourcesOfType_InDirectory(ext string, subpath string) []string {
	result_ := C.C_NSBundle_PathsForResourcesOfType_InDirectory(n.Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSBundle) PathsForResourcesOfType_InDirectory_ForLocalization(ext string, subpath string, localizationName string) []string {
	result_ := C.C_NSBundle_PathsForResourcesOfType_InDirectory_ForLocalization(n.Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr(), NewString(localizationName).Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSBundle) LocalizedStringForKey_Value_Table(key string, value string, tableName string) string {
	result_ := C.C_NSBundle_LocalizedStringForKey_Value_Table(n.Ptr(), NewString(key).Ptr(), NewString(value).Ptr(), NewString(tableName).Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) URLForAuxiliaryExecutable(executableName string) URL {
	result_ := C.C_NSBundle_URLForAuxiliaryExecutable(n.Ptr(), NewString(executableName).Ptr())
	return MakeURL(result_)
}

func (n NSBundle) PathForAuxiliaryExecutable(executableName string) string {
	result_ := C.C_NSBundle_PathForAuxiliaryExecutable(n.Ptr(), NewString(executableName).Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) ObjectForInfoDictionaryKey(key string) objc.Object {
	result_ := C.C_NSBundle_ObjectForInfoDictionaryKey(n.Ptr(), NewString(key).Ptr())
	return objc.MakeObject(result_)
}

func Bundle_PreferredLocalizationsFromArray(localizationsArray []string) []string {
	var cLocalizationsArray C.Array
	if len(localizationsArray) > 0 {
		cLocalizationsArrayData := make([]unsafe.Pointer, len(localizationsArray))
		for idx, v := range localizationsArray {
			cLocalizationsArrayData[idx] = NewString(v).Ptr()
		}
		cLocalizationsArray.data = unsafe.Pointer(&cLocalizationsArrayData[0])
		cLocalizationsArray.len = C.int(len(localizationsArray))
	}
	result_ := C.C_NSBundle_Bundle_PreferredLocalizationsFromArray(cLocalizationsArray)
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func Bundle_PreferredLocalizationsFromArray_ForPreferences(localizationsArray []string, preferencesArray []string) []string {
	var cLocalizationsArray C.Array
	if len(localizationsArray) > 0 {
		cLocalizationsArrayData := make([]unsafe.Pointer, len(localizationsArray))
		for idx, v := range localizationsArray {
			cLocalizationsArrayData[idx] = NewString(v).Ptr()
		}
		cLocalizationsArray.data = unsafe.Pointer(&cLocalizationsArrayData[0])
		cLocalizationsArray.len = C.int(len(localizationsArray))
	}
	var cPreferencesArray C.Array
	if len(preferencesArray) > 0 {
		cPreferencesArrayData := make([]unsafe.Pointer, len(preferencesArray))
		for idx, v := range preferencesArray {
			cPreferencesArrayData[idx] = NewString(v).Ptr()
		}
		cPreferencesArray.data = unsafe.Pointer(&cPreferencesArrayData[0])
		cPreferencesArray.len = C.int(len(preferencesArray))
	}
	result_ := C.C_NSBundle_Bundle_PreferredLocalizationsFromArray_ForPreferences(cLocalizationsArray, cPreferencesArray)
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSBundle) Load() bool {
	result_ := C.C_NSBundle_Load(n.Ptr())
	return bool(result_)
}

func (n NSBundle) Unload() bool {
	result_ := C.C_NSBundle_Unload(n.Ptr())
	return bool(result_)
}

func MainBundle() Bundle {
	result_ := C.C_NSBundle_MainBundle()
	return MakeBundle(result_)
}

func Bundle_AllFrameworks() []Bundle {
	result_ := C.C_NSBundle_Bundle_AllFrameworks()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Bundle, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeBundle(r)
	}
	return goResult_
}

func AllBundles() []Bundle {
	result_ := C.C_NSBundle_AllBundles()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Bundle, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeBundle(r)
	}
	return goResult_
}

func (n NSBundle) ResourceURL() URL {
	result_ := C.C_NSBundle_ResourceURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSBundle) ExecutableURL() URL {
	result_ := C.C_NSBundle_ExecutableURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSBundle) PrivateFrameworksURL() URL {
	result_ := C.C_NSBundle_PrivateFrameworksURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSBundle) SharedFrameworksURL() URL {
	result_ := C.C_NSBundle_SharedFrameworksURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSBundle) BuiltInPlugInsURL() URL {
	result_ := C.C_NSBundle_BuiltInPlugInsURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSBundle) SharedSupportURL() URL {
	result_ := C.C_NSBundle_SharedSupportURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSBundle) AppStoreReceiptURL() URL {
	result_ := C.C_NSBundle_AppStoreReceiptURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSBundle) ResourcePath() string {
	result_ := C.C_NSBundle_ResourcePath(n.Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) ExecutablePath() string {
	result_ := C.C_NSBundle_ExecutablePath(n.Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) PrivateFrameworksPath() string {
	result_ := C.C_NSBundle_PrivateFrameworksPath(n.Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) SharedFrameworksPath() string {
	result_ := C.C_NSBundle_SharedFrameworksPath(n.Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) BuiltInPlugInsPath() string {
	result_ := C.C_NSBundle_BuiltInPlugInsPath(n.Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) SharedSupportPath() string {
	result_ := C.C_NSBundle_SharedSupportPath(n.Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) BundleURL() URL {
	result_ := C.C_NSBundle_BundleURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSBundle) BundlePath() string {
	result_ := C.C_NSBundle_BundlePath(n.Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) BundleIdentifier() string {
	result_ := C.C_NSBundle_BundleIdentifier(n.Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) InfoDictionary() map[string]objc.Object {
	result_ := C.C_NSBundle_InfoDictionary(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[string]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[MakeString(k).String()] = objc.MakeObject(v)
	}
	return goResult_
}

func (n NSBundle) Localizations() []string {
	result_ := C.C_NSBundle_Localizations(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSBundle) PreferredLocalizations() []string {
	result_ := C.C_NSBundle_PreferredLocalizations(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSBundle) DevelopmentLocalization() string {
	result_ := C.C_NSBundle_DevelopmentLocalization(n.Ptr())
	return MakeString(result_).String()
}

func (n NSBundle) LocalizedInfoDictionary() map[string]objc.Object {
	result_ := C.C_NSBundle_LocalizedInfoDictionary(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[string]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[MakeString(k).String()] = objc.MakeObject(v)
	}
	return goResult_
}

func (n NSBundle) ExecutableArchitectures() []Number {
	result_ := C.C_NSBundle_ExecutableArchitectures(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Number, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeNumber(r)
	}
	return goResult_
}

func (n NSBundle) IsLoaded() bool {
	result_ := C.C_NSBundle_IsLoaded(n.Ptr())
	return bool(result_)
}
