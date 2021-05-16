package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "orthography.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Orthography interface {
	objc.Object
	DominantLanguageForScript(script string) string
	LanguagesForScript(script string) []string
	DominantLanguage() string
	DominantScript() string
	AllScripts() []string
	AllLanguages() []string
}

type NSOrthography struct {
	objc.NSObject
}

func MakeOrthography(ptr unsafe.Pointer) *NSOrthography {
	if ptr == nil {
		return nil
	}
	return &NSOrthography{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocOrthography() *NSOrthography {
	return MakeOrthography(C.C_Orthography_Alloc())
}

func (n *NSOrthography) InitWithCoder(coder Coder) Orthography {
	result_ := C.C_NSOrthography_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeOrthography(result_)
}

func (n *NSOrthography) Init() Orthography {
	result_ := C.C_NSOrthography_Init(n.Ptr())
	return MakeOrthography(result_)
}

func Orthography_DefaultOrthographyForLanguage(language string) Orthography {
	result_ := C.C_NSOrthography_Orthography_DefaultOrthographyForLanguage(NewString(language).Ptr())
	return MakeOrthography(result_)
}

func (n *NSOrthography) DominantLanguageForScript(script string) string {
	result_ := C.C_NSOrthography_DominantLanguageForScript(n.Ptr(), NewString(script).Ptr())
	return MakeString(result_).String()
}

func (n *NSOrthography) LanguagesForScript(script string) []string {
	result_ := C.C_NSOrthography_LanguagesForScript(n.Ptr(), NewString(script).Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSOrthography) DominantLanguage() string {
	result_ := C.C_NSOrthography_DominantLanguage(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSOrthography) DominantScript() string {
	result_ := C.C_NSOrthography_DominantScript(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSOrthography) AllScripts() []string {
	result_ := C.C_NSOrthography_AllScripts(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSOrthography) AllLanguages() []string {
	result_ := C.C_NSOrthography_AllLanguages(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}
