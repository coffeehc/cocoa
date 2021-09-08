package foundation

// #include "character_set.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CharacterSet interface {
	objc.Object
	IsSupersetOfSet(theOtherSet CharacterSet) bool
	BitmapRepresentation() []byte
	InvertedSet() CharacterSet
}

type NSCharacterSet struct {
	objc.NSObject
}

func MakeCharacterSet(ptr unsafe.Pointer) NSCharacterSet {
	return NSCharacterSet{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSCharacterSet) InitWithCoder(coder Coder) NSCharacterSet {
	result_ := C.C_NSCharacterSet_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeCharacterSet(result_)
}

func AllocCharacterSet() NSCharacterSet {
	result_ := C.C_NSCharacterSet_AllocCharacterSet()
	return MakeCharacterSet(result_)
}

func (n NSCharacterSet) Autorelease() NSCharacterSet {
	result_ := C.C_NSCharacterSet_Autorelease(n.Ptr())
	return MakeCharacterSet(result_)
}

func (n NSCharacterSet) Retain() NSCharacterSet {
	result_ := C.C_NSCharacterSet_Retain(n.Ptr())
	return MakeCharacterSet(result_)
}

func CharacterSetWithCharactersInString(aString string) CharacterSet {
	result_ := C.C_NSCharacterSet_CharacterSetWithCharactersInString(NewString(aString).Ptr())
	return MakeCharacterSet(result_)
}

func CharacterSetWithRange(aRange Range) CharacterSet {
	result_ := C.C_NSCharacterSet_CharacterSetWithRange(*(*C.NSRange)(ToNSRangePointer(aRange)))
	return MakeCharacterSet(result_)
}

func CharacterSetWithBitmapRepresentation(data []byte) CharacterSet {
	result_ := C.C_NSCharacterSet_CharacterSetWithBitmapRepresentation(NewData(data).Ptr())
	return MakeCharacterSet(result_)
}

func CharacterSetWithContentsOfFile(fName string) CharacterSet {
	result_ := C.C_NSCharacterSet_CharacterSetWithContentsOfFile(NewString(fName).Ptr())
	return MakeCharacterSet(result_)
}

func (n NSCharacterSet) IsSupersetOfSet(theOtherSet CharacterSet) bool {
	result_ := C.C_NSCharacterSet_IsSupersetOfSet(n.Ptr(), objc.ExtractPtr(theOtherSet))
	return bool(result_)
}

func AlphanumericCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_AlphanumericCharacterSet()
	return MakeCharacterSet(result_)
}

func CapitalizedLetterCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_CapitalizedLetterCharacterSet()
	return MakeCharacterSet(result_)
}

func ControlCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_ControlCharacterSet()
	return MakeCharacterSet(result_)
}

func DecimalDigitCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_DecimalDigitCharacterSet()
	return MakeCharacterSet(result_)
}

func DecomposableCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_DecomposableCharacterSet()
	return MakeCharacterSet(result_)
}

func IllegalCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_IllegalCharacterSet()
	return MakeCharacterSet(result_)
}

func LetterCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_LetterCharacterSet()
	return MakeCharacterSet(result_)
}

func LowercaseLetterCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_LowercaseLetterCharacterSet()
	return MakeCharacterSet(result_)
}

func NewlineCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_NewlineCharacterSet()
	return MakeCharacterSet(result_)
}

func NonBaseCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_NonBaseCharacterSet()
	return MakeCharacterSet(result_)
}

func PunctuationCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_PunctuationCharacterSet()
	return MakeCharacterSet(result_)
}

func SymbolCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_SymbolCharacterSet()
	return MakeCharacterSet(result_)
}

func UppercaseLetterCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_UppercaseLetterCharacterSet()
	return MakeCharacterSet(result_)
}

func WhitespaceAndNewlineCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_WhitespaceAndNewlineCharacterSet()
	return MakeCharacterSet(result_)
}

func WhitespaceCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_WhitespaceCharacterSet()
	return MakeCharacterSet(result_)
}

func URLFragmentAllowedCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_URLFragmentAllowedCharacterSet()
	return MakeCharacterSet(result_)
}

func URLHostAllowedCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_URLHostAllowedCharacterSet()
	return MakeCharacterSet(result_)
}

func URLPasswordAllowedCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_URLPasswordAllowedCharacterSet()
	return MakeCharacterSet(result_)
}

func URLPathAllowedCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_URLPathAllowedCharacterSet()
	return MakeCharacterSet(result_)
}

func URLQueryAllowedCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_URLQueryAllowedCharacterSet()
	return MakeCharacterSet(result_)
}

func URLUserAllowedCharacterSet() CharacterSet {
	result_ := C.C_NSCharacterSet_URLUserAllowedCharacterSet()
	return MakeCharacterSet(result_)
}

func (n NSCharacterSet) BitmapRepresentation() []byte {
	result_ := C.C_NSCharacterSet_BitmapRepresentation(n.Ptr())
	return MakeData(result_).ToBytes()
}

func (n NSCharacterSet) InvertedSet() CharacterSet {
	result_ := C.C_NSCharacterSet_InvertedSet(n.Ptr())
	return MakeCharacterSet(result_)
}
