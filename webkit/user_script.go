package webkit

// #include "user_script.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type UserScript interface {
	objc.Object
	Source() string
	InjectionTime() UserScriptInjectionTime
	IsForMainFrameOnly() bool
}

type WKUserScript struct {
	objc.NSObject
}

func MakeUserScript(ptr unsafe.Pointer) *WKUserScript {
	if ptr == nil {
		return nil
	}
	return &WKUserScript{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocUserScript() *WKUserScript {
	return MakeUserScript(C.C_UserScript_Alloc())
}

func (w *WKUserScript) InitWithSource_InjectionTime_ForMainFrameOnly(source string, injectionTime UserScriptInjectionTime, forMainFrameOnly bool) UserScript {
	result_ := C.C_WKUserScript_InitWithSource_InjectionTime_ForMainFrameOnly(w.Ptr(), foundation.NewString(source).Ptr(), C.int(int(injectionTime)), C.bool(forMainFrameOnly))
	return MakeUserScript(result_)
}

func (w *WKUserScript) InitWithSource_InjectionTime_ForMainFrameOnly_InContentWorld(source string, injectionTime UserScriptInjectionTime, forMainFrameOnly bool, contentWorld ContentWorld) UserScript {
	result_ := C.C_WKUserScript_InitWithSource_InjectionTime_ForMainFrameOnly_InContentWorld(w.Ptr(), foundation.NewString(source).Ptr(), C.int(int(injectionTime)), C.bool(forMainFrameOnly), objc.ExtractPtr(contentWorld))
	return MakeUserScript(result_)
}

func (w *WKUserScript) Init() UserScript {
	result_ := C.C_WKUserScript_Init(w.Ptr())
	return MakeUserScript(result_)
}

func (w *WKUserScript) Source() string {
	result_ := C.C_WKUserScript_Source(w.Ptr())
	return foundation.MakeString(result_).String()
}

func (w *WKUserScript) InjectionTime() UserScriptInjectionTime {
	result_ := C.C_WKUserScript_InjectionTime(w.Ptr())
	return UserScriptInjectionTime(int(result_))
}

func (w *WKUserScript) IsForMainFrameOnly() bool {
	result_ := C.C_WKUserScript_IsForMainFrameOnly(w.Ptr())
	return bool(result_)
}
