package webkit

// #include "content_rule_list.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ContentRuleList interface {
	objc.Object
	Identifier() string
}

type WKContentRuleList struct {
	objc.NSObject
}

func MakeContentRuleList(ptr unsafe.Pointer) *WKContentRuleList {
	if ptr == nil {
		return nil
	}
	return &WKContentRuleList{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocContentRuleList() *WKContentRuleList {
	return MakeContentRuleList(C.C_ContentRuleList_Alloc())
}

func (w *WKContentRuleList) Init() ContentRuleList {
	result_ := C.C_WKContentRuleList_Init(w.Ptr())
	return MakeContentRuleList(result_)
}

func (w *WKContentRuleList) Identifier() string {
	result_ := C.C_WKContentRuleList_Identifier(w.Ptr())
	return foundation.MakeString(result_).String()
}
