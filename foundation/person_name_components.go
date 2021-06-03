package foundation

// #include "person_name_components.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PersonNameComponents interface {
	objc.Object
	NamePrefix() string
	SetNamePrefix(value string)
	GivenName() string
	SetGivenName(value string)
	MiddleName() string
	SetMiddleName(value string)
	FamilyName() string
	SetFamilyName(value string)
	NameSuffix() string
	SetNameSuffix(value string)
	Nickname() string
	SetNickname(value string)
	PhoneticRepresentation() PersonNameComponents
	SetPhoneticRepresentation(value PersonNameComponents)
}

type NSPersonNameComponents struct {
	objc.NSObject
}

func MakePersonNameComponents(ptr unsafe.Pointer) NSPersonNameComponents {
	return NSPersonNameComponents{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocPersonNameComponents() NSPersonNameComponents {
	return MakePersonNameComponents(C.C_PersonNameComponents_Alloc())
}

func (n NSPersonNameComponents) Init() PersonNameComponents {
	result_ := C.C_NSPersonNameComponents_Init(n.Ptr())
	return MakePersonNameComponents(result_)
}

func (n NSPersonNameComponents) NamePrefix() string {
	result_ := C.C_NSPersonNameComponents_NamePrefix(n.Ptr())
	return MakeString(result_).String()
}

func (n NSPersonNameComponents) SetNamePrefix(value string) {
	C.C_NSPersonNameComponents_SetNamePrefix(n.Ptr(), NewString(value).Ptr())
}

func (n NSPersonNameComponents) GivenName() string {
	result_ := C.C_NSPersonNameComponents_GivenName(n.Ptr())
	return MakeString(result_).String()
}

func (n NSPersonNameComponents) SetGivenName(value string) {
	C.C_NSPersonNameComponents_SetGivenName(n.Ptr(), NewString(value).Ptr())
}

func (n NSPersonNameComponents) MiddleName() string {
	result_ := C.C_NSPersonNameComponents_MiddleName(n.Ptr())
	return MakeString(result_).String()
}

func (n NSPersonNameComponents) SetMiddleName(value string) {
	C.C_NSPersonNameComponents_SetMiddleName(n.Ptr(), NewString(value).Ptr())
}

func (n NSPersonNameComponents) FamilyName() string {
	result_ := C.C_NSPersonNameComponents_FamilyName(n.Ptr())
	return MakeString(result_).String()
}

func (n NSPersonNameComponents) SetFamilyName(value string) {
	C.C_NSPersonNameComponents_SetFamilyName(n.Ptr(), NewString(value).Ptr())
}

func (n NSPersonNameComponents) NameSuffix() string {
	result_ := C.C_NSPersonNameComponents_NameSuffix(n.Ptr())
	return MakeString(result_).String()
}

func (n NSPersonNameComponents) SetNameSuffix(value string) {
	C.C_NSPersonNameComponents_SetNameSuffix(n.Ptr(), NewString(value).Ptr())
}

func (n NSPersonNameComponents) Nickname() string {
	result_ := C.C_NSPersonNameComponents_Nickname(n.Ptr())
	return MakeString(result_).String()
}

func (n NSPersonNameComponents) SetNickname(value string) {
	C.C_NSPersonNameComponents_SetNickname(n.Ptr(), NewString(value).Ptr())
}

func (n NSPersonNameComponents) PhoneticRepresentation() PersonNameComponents {
	result_ := C.C_NSPersonNameComponents_PhoneticRepresentation(n.Ptr())
	return MakePersonNameComponents(result_)
}

func (n NSPersonNameComponents) SetPhoneticRepresentation(value PersonNameComponents) {
	C.C_NSPersonNameComponents_SetPhoneticRepresentation(n.Ptr(), objc.ExtractPtr(value))
}
