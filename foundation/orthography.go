package foundation

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

func MakeOrthography(ptr unsafe.Pointer) NSOrthography {
	return NSOrthography{
		NSObject: objc.MakeObject(ptr),
	}
}

func DefaultOrthographyForLanguage(language string) NSOrthography {
	result_ := C.C_NSOrthography_DefaultOrthographyForLanguage(NewString(language).Ptr())
	return MakeOrthography(result_)
}

func (n NSOrthography) InitWithCoder(coder Coder) NSOrthography {
	result_ := C.C_NSOrthography_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeOrthography(result_)
}

func AllocOrthography() NSOrthography {
	result_ := C.C_NSOrthography_AllocOrthography()
	return MakeOrthography(result_)
}

func (n NSOrthography) Init() NSOrthography {
	result_ := C.C_NSOrthography_Init(n.Ptr())
	return MakeOrthography(result_)
}

func NewOrthography() NSOrthography {
	result_ := C.C_NSOrthography_NewOrthography()
	return MakeOrthography(result_)
}

func (n NSOrthography) Autorelease() NSOrthography {
	result_ := C.C_NSOrthography_Autorelease(n.Ptr())
	return MakeOrthography(result_)
}

func (n NSOrthography) Retain() NSOrthography {
	result_ := C.C_NSOrthography_Retain(n.Ptr())
	return MakeOrthography(result_)
}

func (n NSOrthography) DominantLanguageForScript(script string) string {
	result_ := C.C_NSOrthography_DominantLanguageForScript(n.Ptr(), NewString(script).Ptr())
	return MakeString(result_).String()
}

func (n NSOrthography) LanguagesForScript(script string) []string {
	result_ := C.C_NSOrthography_LanguagesForScript(n.Ptr(), NewString(script).Ptr())
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

func (n NSOrthography) DominantLanguage() string {
	result_ := C.C_NSOrthography_DominantLanguage(n.Ptr())
	return MakeString(result_).String()
}

func (n NSOrthography) DominantScript() string {
	result_ := C.C_NSOrthography_DominantScript(n.Ptr())
	return MakeString(result_).String()
}

func (n NSOrthography) AllScripts() []string {
	result_ := C.C_NSOrthography_AllScripts(n.Ptr())
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

func (n NSOrthography) AllLanguages() []string {
	result_ := C.C_NSOrthography_AllLanguages(n.Ptr())
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
