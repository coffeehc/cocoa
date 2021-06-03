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

func AllocCandidateListTouchBarItem() NSCandidateListTouchBarItem {
	return MakeCandidateListTouchBarItem(C.C_CandidateListTouchBarItem_Alloc())
}

func (n NSCandidateListTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) CandidateListTouchBarItem {
	result_ := C.C_NSCandidateListTouchBarItem_InitWithIdentifier(n.Ptr(), foundation.NewString(string(identifier)).Ptr())
	return MakeCandidateListTouchBarItem(result_)
}

func (n NSCandidateListTouchBarItem) InitWithCoder(coder foundation.Coder) CandidateListTouchBarItem {
	result_ := C.C_NSCandidateListTouchBarItem_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeCandidateListTouchBarItem(result_)
}
