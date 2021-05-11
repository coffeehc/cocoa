package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "candidate_list_touch_bar_item.h"
import "C"
import (
	"unsafe"
)

type CandidateListTouchBarItem interface {
}

type NSCandidateListTouchBarItem struct {
}

func MakeCandidateListTouchBarItem(ptr unsafe.Pointer) *NSCandidateListTouchBarItem {
	if ptr == nil {
		return nil
	}
	return &NSCandidateListTouchBarItem{}
}

func AllocCandidateListTouchBarItem() *NSCandidateListTouchBarItem {
	return MakeCandidateListTouchBarItem(C.C_CandidateListTouchBarItem_Alloc())
}
