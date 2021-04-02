package uti

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework UniformTypeIdentifiers
// #include "ut_type.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type UTType interface {
	objc.Object
	Identifier() string
	PreferredFilenameExtension() string
	PreferredMIMEType() string
	IsDeclared() bool
	IsDynamic() bool
	IsPublicType() bool
	ReferenceURL() foundation.URL
	ConformsToType(type_ UTType) bool
	IsSubtypeOfType(type_ UTType) bool
	IsSupertypeOfType(type_ UTType) bool
}

var _ UTType = (*NSUTType)(nil)

type NSUTType struct {
	objc.NSObject
}

func MakeUTType(ptr unsafe.Pointer) *NSUTType {
	if ptr == nil {
		return nil
	}
	return &NSUTType{
		NSObject: *objc.MakeObject(ptr),
	}
}

func (u *NSUTType) Identifier() string {
	return C.GoString(C.UTType_Identifier(u.Ptr()))
}

func (u *NSUTType) PreferredFilenameExtension() string {
	return C.GoString(C.UTType_PreferredFilenameExtension(u.Ptr()))
}

func (u *NSUTType) PreferredMIMEType() string {
	return C.GoString(C.UTType_PreferredMIMEType(u.Ptr()))
}

func (u *NSUTType) IsDeclared() bool {
	return bool(C.UTType_IsDeclared(u.Ptr()))
}

func (u *NSUTType) IsDynamic() bool {
	return bool(C.UTType_IsDynamic(u.Ptr()))
}

func (u *NSUTType) IsPublicType() bool {
	return bool(C.UTType_IsPublicType(u.Ptr()))
}

func (u *NSUTType) ReferenceURL() foundation.URL {
	return foundation.MakeURL(C.UTType_ReferenceURL(u.Ptr()))
}

func typesWithTagConformingToType(tag string, tagClass string, supertype UTType) []UTType {
	cTag := C.CString(tag)
	defer C.free(unsafe.Pointer(cTag))
	cTagClass := C.CString(tagClass)
	defer C.free(unsafe.Pointer(cTagClass))
	var cArray C.Array = C.UTType_typesWithTagConformingToType(cTag, cTagClass, toPointer(supertype))
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]UTType, len(result))
	for idx, r := range result {
		goArray[idx] = MakeUTType(r)
	}
	return goArray
}

func ExportedTypeWithIdentifier(identifier string) UTType {
	cIdentifier := C.CString(identifier)
	defer C.free(unsafe.Pointer(cIdentifier))
	return MakeUTType(C.UTType_ExportedTypeWithIdentifier(cIdentifier))
}

func exportedTypeWithIdentifierConformingToType(identifier string, parentType UTType) UTType {
	cIdentifier := C.CString(identifier)
	defer C.free(unsafe.Pointer(cIdentifier))
	return MakeUTType(C.UTType_exportedTypeWithIdentifierConformingToType(cIdentifier, toPointer(parentType)))
}

func ImportedTypeWithIdentifier(identifier string) UTType {
	cIdentifier := C.CString(identifier)
	defer C.free(unsafe.Pointer(cIdentifier))
	return MakeUTType(C.UTType_ImportedTypeWithIdentifier(cIdentifier))
}

func importedTypeWithIdentifierConformingToType(identifier string, parentType UTType) UTType {
	cIdentifier := C.CString(identifier)
	defer C.free(unsafe.Pointer(cIdentifier))
	return MakeUTType(C.UTType_importedTypeWithIdentifierConformingToType(cIdentifier, toPointer(parentType)))
}

func TypeWithIdentifier(identifier string) UTType {
	cIdentifier := C.CString(identifier)
	defer C.free(unsafe.Pointer(cIdentifier))
	return MakeUTType(C.UTType_TypeWithIdentifier(cIdentifier))
}

func TypeWithFilenameExtension(filenameExtension string) UTType {
	cFilenameExtension := C.CString(filenameExtension)
	defer C.free(unsafe.Pointer(cFilenameExtension))
	return MakeUTType(C.UTType_TypeWithFilenameExtension(cFilenameExtension))
}

func typeWithFilenameExtensionConformingToType(filenameExtension string, supertype UTType) UTType {
	cFilenameExtension := C.CString(filenameExtension)
	defer C.free(unsafe.Pointer(cFilenameExtension))
	return MakeUTType(C.UTType_typeWithFilenameExtensionConformingToType(cFilenameExtension, toPointer(supertype)))
}

func TypeWithMIMEType(mimeType string) UTType {
	cMimeType := C.CString(mimeType)
	defer C.free(unsafe.Pointer(cMimeType))
	return MakeUTType(C.UTType_TypeWithMIMEType(cMimeType))
}

func typeWithMIMETypeConformingToType(mimeType string, supertype UTType) UTType {
	cMimeType := C.CString(mimeType)
	defer C.free(unsafe.Pointer(cMimeType))
	return MakeUTType(C.UTType_typeWithMIMETypeConformingToType(cMimeType, toPointer(supertype)))
}

func typeWithTagConformingToType(tag string, tagClass string, supertype UTType) UTType {
	cTag := C.CString(tag)
	defer C.free(unsafe.Pointer(cTag))
	cTagClass := C.CString(tagClass)
	defer C.free(unsafe.Pointer(cTagClass))
	return MakeUTType(C.UTType_typeWithTagConformingToType(cTag, cTagClass, toPointer(supertype)))
}

func (u *NSUTType) ConformsToType(type_ UTType) bool {
	return bool(C.UTType_ConformsToType(u.Ptr(), toPointer(type_)))
}

func (u *NSUTType) IsSubtypeOfType(type_ UTType) bool {
	return bool(C.UTType_IsSubtypeOfType(u.Ptr(), toPointer(type_)))
}

func (u *NSUTType) IsSupertypeOfType(type_ UTType) bool {
	return bool(C.UTType_IsSupertypeOfType(u.Ptr(), toPointer(type_)))
}
