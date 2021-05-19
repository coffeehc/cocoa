package webkit

// #include "user_content_controller.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type UserContentController interface {
	objc.Object
	AddUserScript(userScript UserScript)
	RemoveAllUserScripts()
	AddScriptMessageHandler_Name(scriptMessageHandler objc.Object, name string)
	AddScriptMessageHandler_ContentWorld_Name(scriptMessageHandler objc.Object, world ContentWorld, name string)
	AddScriptMessageHandlerWithReply_ContentWorld_Name(scriptMessageHandlerWithReply objc.Object, contentWorld ContentWorld, name string)
	RemoveScriptMessageHandlerForName(name string)
	RemoveScriptMessageHandlerForName_ContentWorld(name string, contentWorld ContentWorld)
	RemoveAllScriptMessageHandlersFromContentWorld(contentWorld ContentWorld)
	RemoveAllScriptMessageHandlers()
	AddContentRuleList(contentRuleList ContentRuleList)
	RemoveContentRuleList(contentRuleList ContentRuleList)
	RemoveAllContentRuleLists()
	UserScripts() []UserScript
}

type WKUserContentController struct {
	objc.NSObject
}

func MakeUserContentController(ptr unsafe.Pointer) *WKUserContentController {
	if ptr == nil {
		return nil
	}
	return &WKUserContentController{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocUserContentController() *WKUserContentController {
	return MakeUserContentController(C.C_UserContentController_Alloc())
}

func (w *WKUserContentController) Init() UserContentController {
	result_ := C.C_WKUserContentController_Init(w.Ptr())
	return MakeUserContentController(result_)
}

func (w *WKUserContentController) AddUserScript(userScript UserScript) {
	C.C_WKUserContentController_AddUserScript(w.Ptr(), objc.ExtractPtr(userScript))
}

func (w *WKUserContentController) RemoveAllUserScripts() {
	C.C_WKUserContentController_RemoveAllUserScripts(w.Ptr())
}

func (w *WKUserContentController) AddScriptMessageHandler_Name(scriptMessageHandler objc.Object, name string) {
	C.C_WKUserContentController_AddScriptMessageHandler_Name(w.Ptr(), objc.ExtractPtr(scriptMessageHandler), foundation.NewString(name).Ptr())
}

func (w *WKUserContentController) AddScriptMessageHandler_ContentWorld_Name(scriptMessageHandler objc.Object, world ContentWorld, name string) {
	C.C_WKUserContentController_AddScriptMessageHandler_ContentWorld_Name(w.Ptr(), objc.ExtractPtr(scriptMessageHandler), objc.ExtractPtr(world), foundation.NewString(name).Ptr())
}

func (w *WKUserContentController) AddScriptMessageHandlerWithReply_ContentWorld_Name(scriptMessageHandlerWithReply objc.Object, contentWorld ContentWorld, name string) {
	C.C_WKUserContentController_AddScriptMessageHandlerWithReply_ContentWorld_Name(w.Ptr(), objc.ExtractPtr(scriptMessageHandlerWithReply), objc.ExtractPtr(contentWorld), foundation.NewString(name).Ptr())
}

func (w *WKUserContentController) RemoveScriptMessageHandlerForName(name string) {
	C.C_WKUserContentController_RemoveScriptMessageHandlerForName(w.Ptr(), foundation.NewString(name).Ptr())
}

func (w *WKUserContentController) RemoveScriptMessageHandlerForName_ContentWorld(name string, contentWorld ContentWorld) {
	C.C_WKUserContentController_RemoveScriptMessageHandlerForName_ContentWorld(w.Ptr(), foundation.NewString(name).Ptr(), objc.ExtractPtr(contentWorld))
}

func (w *WKUserContentController) RemoveAllScriptMessageHandlersFromContentWorld(contentWorld ContentWorld) {
	C.C_WKUserContentController_RemoveAllScriptMessageHandlersFromContentWorld(w.Ptr(), objc.ExtractPtr(contentWorld))
}

func (w *WKUserContentController) RemoveAllScriptMessageHandlers() {
	C.C_WKUserContentController_RemoveAllScriptMessageHandlers(w.Ptr())
}

func (w *WKUserContentController) AddContentRuleList(contentRuleList ContentRuleList) {
	C.C_WKUserContentController_AddContentRuleList(w.Ptr(), objc.ExtractPtr(contentRuleList))
}

func (w *WKUserContentController) RemoveContentRuleList(contentRuleList ContentRuleList) {
	C.C_WKUserContentController_RemoveContentRuleList(w.Ptr(), objc.ExtractPtr(contentRuleList))
}

func (w *WKUserContentController) RemoveAllContentRuleLists() {
	C.C_WKUserContentController_RemoveAllContentRuleLists(w.Ptr())
}

func (w *WKUserContentController) UserScripts() []UserScript {
	result_ := C.C_WKUserContentController_UserScripts(w.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]UserScript, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeUserScript(r)
	}
	return goResult_
}
