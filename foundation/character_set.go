package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
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

func MakeCharacterSet(ptr unsafe.Pointer) *NSCharacterSet {
	if ptr == nil {
		return nil
	}
	return &NSCharacterSet{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocCharacterSet() *NSCharacterSet {
	return MakeCharacterSet(C.C_CharacterSet_Alloc())
}

func (n *NSCharacterSet) InitWithCoder(coder Coder) CharacterSet {
	result := C.C_NSCharacterSet_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeCharacterSet(result)
}

func CharacterSetWithCharactersInString(aString string) CharacterSet {
	result := C.C_NSCharacterSet_CharacterSetWithCharactersInString(NewString(aString).Ptr())
	return MakeCharacterSet(result)
}

func CharacterSetWithRange(aRange Range) CharacterSet {
	result := C.C_NSCharacterSet_CharacterSetWithRange(*(*C.NSRange)(ToNSRangePointer(aRange)))
	return MakeCharacterSet(result)
}

func CharacterSetWithBitmapRepresentation(data []byte) CharacterSet {
	result := C.C_NSCharacterSet_CharacterSetWithBitmapRepresentation(C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
	return MakeCharacterSet(result)
}

func CharacterSetWithContentsOfFile(fName string) CharacterSet {
	result := C.C_NSCharacterSet_CharacterSetWithContentsOfFile(NewString(fName).Ptr())
	return MakeCharacterSet(result)
}

func (n *NSCharacterSet) IsSupersetOfSet(theOtherSet CharacterSet) bool {
	result := C.C_NSCharacterSet_IsSupersetOfSet(n.Ptr(), objc.ExtractPtr(theOtherSet))
	return bool(result)
}

func AlphanumericCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_AlphanumericCharacterSet()
	return MakeCharacterSet(result)
}

func CapitalizedLetterCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_CapitalizedLetterCharacterSet()
	return MakeCharacterSet(result)
}

func ControlCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_ControlCharacterSet()
	return MakeCharacterSet(result)
}

func DecimalDigitCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_DecimalDigitCharacterSet()
	return MakeCharacterSet(result)
}

func DecomposableCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_DecomposableCharacterSet()
	return MakeCharacterSet(result)
}

func IllegalCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_IllegalCharacterSet()
	return MakeCharacterSet(result)
}

func LetterCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_LetterCharacterSet()
	return MakeCharacterSet(result)
}

func LowercaseLetterCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_LowercaseLetterCharacterSet()
	return MakeCharacterSet(result)
}

func NewlineCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_NewlineCharacterSet()
	return MakeCharacterSet(result)
}

func NonBaseCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_NonBaseCharacterSet()
	return MakeCharacterSet(result)
}

func PunctuationCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_PunctuationCharacterSet()
	return MakeCharacterSet(result)
}

func SymbolCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_SymbolCharacterSet()
	return MakeCharacterSet(result)
}

func UppercaseLetterCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_UppercaseLetterCharacterSet()
	return MakeCharacterSet(result)
}

func WhitespaceAndNewlineCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_WhitespaceAndNewlineCharacterSet()
	return MakeCharacterSet(result)
}

func WhitespaceCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_WhitespaceCharacterSet()
	return MakeCharacterSet(result)
}

func URLFragmentAllowedCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_URLFragmentAllowedCharacterSet()
	return MakeCharacterSet(result)
}

func URLHostAllowedCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_URLHostAllowedCharacterSet()
	return MakeCharacterSet(result)
}

func URLPasswordAllowedCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_URLPasswordAllowedCharacterSet()
	return MakeCharacterSet(result)
}

func URLPathAllowedCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_URLPathAllowedCharacterSet()
	return MakeCharacterSet(result)
}

func URLQueryAllowedCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_URLQueryAllowedCharacterSet()
	return MakeCharacterSet(result)
}

func URLUserAllowedCharacterSet() CharacterSet {
	result := C.C_NSCharacterSet_URLUserAllowedCharacterSet()
	return MakeCharacterSet(result)
}

func (n *NSCharacterSet) BitmapRepresentation() []byte {
	result := C.C_NSCharacterSet_BitmapRepresentation(n.Ptr())
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func (n *NSCharacterSet) InvertedSet() CharacterSet {
	result := C.C_NSCharacterSet_InvertedSet(n.Ptr())
	return MakeCharacterSet(result)
}
