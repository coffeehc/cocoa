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

func MakeContentRuleList(ptr unsafe.Pointer) WKContentRuleList {
	return WKContentRuleList{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocContentRuleList() WKContentRuleList {
	result_ := C.C_WKContentRuleList_AllocContentRuleList()
	return MakeContentRuleList(result_)
}

func (w WKContentRuleList) Init() WKContentRuleList {
	result_ := C.C_WKContentRuleList_Init(w.Ptr())
	return MakeContentRuleList(result_)
}

func NewContentRuleList() WKContentRuleList {
	result_ := C.C_WKContentRuleList_NewContentRuleList()
	return MakeContentRuleList(result_)
}

func (w WKContentRuleList) Autorelease() WKContentRuleList {
	result_ := C.C_WKContentRuleList_Autorelease(w.Ptr())
	return MakeContentRuleList(result_)
}

func (w WKContentRuleList) Retain() WKContentRuleList {
	result_ := C.C_WKContentRuleList_Retain(w.Ptr())
	return MakeContentRuleList(result_)
}

func (w WKContentRuleList) Identifier() string {
	result_ := C.C_WKContentRuleList_Identifier(w.Ptr())
	return foundation.MakeString(result_).String()
}
