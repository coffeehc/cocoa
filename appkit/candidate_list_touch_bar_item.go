package appkit

// #include "candidate_list_touch_bar_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CandidateListTouchBarItem interface {
	TouchBarItem
}

type NSCandidateListTouchBarItem struct {
	NSTouchBarItem
}

func MakeCandidateListTouchBarItem(ptr unsafe.Pointer) NSCandidateListTouchBarItem {
	return NSCandidateListTouchBarItem{
		NSTouchBarItem: MakeTouchBarItem(ptr),
	}
}

func (n NSCandidateListTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) NSCandidateListTouchBarItem {
	result_ := C.C_NSCandidateListTouchBarItem_InitWithIdentifier(n.Ptr(), foundation.NewString(string(identifier)).Ptr())
	return MakeCandidateListTouchBarItem(result_)
}

func (n NSCandidateListTouchBarItem) InitWithCoder(coder foundation.Coder) NSCandidateListTouchBarItem {
	result_ := C.C_NSCandidateListTouchBarItem_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeCandidateListTouchBarItem(result_)
}

func AllocCandidateListTouchBarItem() NSCandidateListTouchBarItem {
	result_ := C.C_NSCandidateListTouchBarItem_AllocCandidateListTouchBarItem()
	return MakeCandidateListTouchBarItem(result_)
}

func (n NSCandidateListTouchBarItem) Autorelease() NSCandidateListTouchBarItem {
	result_ := C.C_NSCandidateListTouchBarItem_Autorelease(n.Ptr())
	return MakeCandidateListTouchBarItem(result_)
}

func (n NSCandidateListTouchBarItem) Retain() NSCandidateListTouchBarItem {
	result_ := C.C_NSCandidateListTouchBarItem_Retain(n.Ptr())
	return MakeCandidateListTouchBarItem(result_)
}
