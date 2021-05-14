package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
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
	URLForResource_WithExtension_Subdirectory_Localization(name string, ext string, subpath string, localizationName string) URL
	PathForResource_OfType(name string, ext string) string
	PathForResource_OfType_InDirectory(name string, ext string, subpath string) string
	PathForResource_OfType_InDirectory_ForLocalization(name string, ext string, subpath string, localizationName string) string
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
	DevelopmentLocalization() string
	IsLoaded() bool
}

type NSBundle struct {
	objc.NSObject
}

func MakeBundle(ptr unsafe.Pointer) *NSBundle {
	if ptr == nil {
		return nil
	}
	return &NSBundle{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocBundle() *NSBundle {
	return MakeBundle(C.C_Bundle_Alloc())
}

func (n *NSBundle) InitWithURL(url URL) Bundle {
	result := C.C_NSBundle_InitWithURL(n.Ptr(), objc.ExtractPtr(url))
	return MakeBundle(result)
}

func (n *NSBundle) InitWithPath(path string) Bundle {
	result := C.C_NSBundle_InitWithPath(n.Ptr(), NewString(path).Ptr())
	return MakeBundle(result)
}

func (n *NSBundle) Init() Bundle {
	result := C.C_NSBundle_Init(n.Ptr())
	return MakeBundle(result)
}

func BundleWithURL(url URL) Bundle {
	result := C.C_NSBundle_BundleWithURL(objc.ExtractPtr(url))
	return MakeBundle(result)
}

func BundleWithPath(path string) Bundle {
	result := C.C_NSBundle_BundleWithPath(NewString(path).Ptr())
	return MakeBundle(result)
}

func BundleWithIdentifier(identifier string) Bundle {
	result := C.C_NSBundle_BundleWithIdentifier(NewString(identifier).Ptr())
	return MakeBundle(result)
}

func (n *NSBundle) URLForResource_WithExtension_Subdirectory(name string, ext string, subpath string) URL {
	result := C.C_NSBundle_URLForResource_WithExtension_Subdirectory(n.Ptr(), NewString(name).Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr())
	return MakeURL(result)
}

func (n *NSBundle) URLForResource_WithExtension(name string, ext string) URL {
	result := C.C_NSBundle_URLForResource_WithExtension(n.Ptr(), NewString(name).Ptr(), NewString(ext).Ptr())
	return MakeURL(result)
}

func (n *NSBundle) URLForResource_WithExtension_Subdirectory_Localization(name string, ext string, subpath string, localizationName string) URL {
	result := C.C_NSBundle_URLForResource_WithExtension_Subdirectory_Localization(n.Ptr(), NewString(name).Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr(), NewString(localizationName).Ptr())
	return MakeURL(result)
}

func BundleURLForResource_WithExtension_Subdirectory_InBundleWithURL(name string, ext string, subpath string, bundleURL URL) URL {
	result := C.C_NSBundle_BundleURLForResource_WithExtension_Subdirectory_InBundleWithURL(NewString(name).Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr(), objc.ExtractPtr(bundleURL))
	return MakeURL(result)
}

func (n *NSBundle) PathForResource_OfType(name string, ext string) string {
	result := C.C_NSBundle_PathForResource_OfType(n.Ptr(), NewString(name).Ptr(), NewString(ext).Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) PathForResource_OfType_InDirectory(name string, ext string, subpath string) string {
	result := C.C_NSBundle_PathForResource_OfType_InDirectory(n.Ptr(), NewString(name).Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) PathForResource_OfType_InDirectory_ForLocalization(name string, ext string, subpath string, localizationName string) string {
	result := C.C_NSBundle_PathForResource_OfType_InDirectory_ForLocalization(n.Ptr(), NewString(name).Ptr(), NewString(ext).Ptr(), NewString(subpath).Ptr(), NewString(localizationName).Ptr())
	return MakeString(result).String()
}

func BundlePathForResource_OfType_InDirectory(name string, ext string, bundlePath string) string {
	result := C.C_NSBundle_BundlePathForResource_OfType_InDirectory(NewString(name).Ptr(), NewString(ext).Ptr(), NewString(bundlePath).Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) LocalizedStringForKey_Value_Table(key string, value string, tableName string) string {
	result := C.C_NSBundle_LocalizedStringForKey_Value_Table(n.Ptr(), NewString(key).Ptr(), NewString(value).Ptr(), NewString(tableName).Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) URLForAuxiliaryExecutable(executableName string) URL {
	result := C.C_NSBundle_URLForAuxiliaryExecutable(n.Ptr(), NewString(executableName).Ptr())
	return MakeURL(result)
}

func (n *NSBundle) PathForAuxiliaryExecutable(executableName string) string {
	result := C.C_NSBundle_PathForAuxiliaryExecutable(n.Ptr(), NewString(executableName).Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) ObjectForInfoDictionaryKey(key string) objc.Object {
	result := C.C_NSBundle_ObjectForInfoDictionaryKey(n.Ptr(), NewString(key).Ptr())
	return objc.MakeObject(result)
}

func (n *NSBundle) Load() bool {
	result := C.C_NSBundle_Load(n.Ptr())
	return bool(result)
}

func (n *NSBundle) Unload() bool {
	result := C.C_NSBundle_Unload(n.Ptr())
	return bool(result)
}

func (n *NSBundle) ResourceURL() URL {
	result := C.C_NSBundle_ResourceURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSBundle) ExecutableURL() URL {
	result := C.C_NSBundle_ExecutableURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSBundle) PrivateFrameworksURL() URL {
	result := C.C_NSBundle_PrivateFrameworksURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSBundle) SharedFrameworksURL() URL {
	result := C.C_NSBundle_SharedFrameworksURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSBundle) BuiltInPlugInsURL() URL {
	result := C.C_NSBundle_BuiltInPlugInsURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSBundle) SharedSupportURL() URL {
	result := C.C_NSBundle_SharedSupportURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSBundle) AppStoreReceiptURL() URL {
	result := C.C_NSBundle_AppStoreReceiptURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSBundle) ResourcePath() string {
	result := C.C_NSBundle_ResourcePath(n.Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) ExecutablePath() string {
	result := C.C_NSBundle_ExecutablePath(n.Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) PrivateFrameworksPath() string {
	result := C.C_NSBundle_PrivateFrameworksPath(n.Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) SharedFrameworksPath() string {
	result := C.C_NSBundle_SharedFrameworksPath(n.Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) BuiltInPlugInsPath() string {
	result := C.C_NSBundle_BuiltInPlugInsPath(n.Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) SharedSupportPath() string {
	result := C.C_NSBundle_SharedSupportPath(n.Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) BundleURL() URL {
	result := C.C_NSBundle_BundleURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSBundle) BundlePath() string {
	result := C.C_NSBundle_BundlePath(n.Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) BundleIdentifier() string {
	result := C.C_NSBundle_BundleIdentifier(n.Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) DevelopmentLocalization() string {
	result := C.C_NSBundle_DevelopmentLocalization(n.Ptr())
	return MakeString(result).String()
}

func (n *NSBundle) IsLoaded() bool {
	result := C.C_NSBundle_IsLoaded(n.Ptr())
	return bool(result)
}
