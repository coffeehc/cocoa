package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "layout_manager.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type LayoutManager interface {
	objc.Object
	AllowsNonContiguousLayout() bool
	SetAllowsNonContiguousLayout(allowsNonContiguousLayout bool)
	HasNonContiguousLayout() bool
	ShowsInvisibleCharacters() bool
	SetShowsInvisibleCharacters(showsInvisibleCharacters bool)
	ShowsControlCharacters() bool
	SetShowsControlCharacters(showsControlCharacters bool)
	UsesFontLeading() bool
	SetUsesFontLeading(usesFontLeading bool)
	BackgroundLayoutEnabled() bool
	SetBackgroundLayoutEnabled(backgroundLayoutEnabled bool)
	LimitsLayoutForSuspiciousContents() bool
	SetLimitsLayoutForSuspiciousContents(limitsLayoutForSuspiciousContents bool)
	UsesDefaultHyphenation() bool
	SetUsesDefaultHyphenation(usesDefaultHyphenation bool)
}

var _ LayoutManager = (*NSLayoutManager)(nil)

type NSLayoutManager struct {
	objc.NSObject
}

func MakeLayoutManager(ptr unsafe.Pointer) *NSLayoutManager {
	if ptr == nil {
		return nil
	}
	return &NSLayoutManager{
		NSObject: *objc.MakeObject(ptr),
	}
}

func (l *NSLayoutManager) AllowsNonContiguousLayout() bool {
	return bool(C.LayoutManager_AllowsNonContiguousLayout(l.Ptr()))
}

func (l *NSLayoutManager) SetAllowsNonContiguousLayout(allowsNonContiguousLayout bool) {
	C.LayoutManager_SetAllowsNonContiguousLayout(l.Ptr(), C.bool(allowsNonContiguousLayout))
}

func (l *NSLayoutManager) HasNonContiguousLayout() bool {
	return bool(C.LayoutManager_HasNonContiguousLayout(l.Ptr()))
}

func (l *NSLayoutManager) ShowsInvisibleCharacters() bool {
	return bool(C.LayoutManager_ShowsInvisibleCharacters(l.Ptr()))
}

func (l *NSLayoutManager) SetShowsInvisibleCharacters(showsInvisibleCharacters bool) {
	C.LayoutManager_SetShowsInvisibleCharacters(l.Ptr(), C.bool(showsInvisibleCharacters))
}

func (l *NSLayoutManager) ShowsControlCharacters() bool {
	return bool(C.LayoutManager_ShowsControlCharacters(l.Ptr()))
}

func (l *NSLayoutManager) SetShowsControlCharacters(showsControlCharacters bool) {
	C.LayoutManager_SetShowsControlCharacters(l.Ptr(), C.bool(showsControlCharacters))
}

func (l *NSLayoutManager) UsesFontLeading() bool {
	return bool(C.LayoutManager_UsesFontLeading(l.Ptr()))
}

func (l *NSLayoutManager) SetUsesFontLeading(usesFontLeading bool) {
	C.LayoutManager_SetUsesFontLeading(l.Ptr(), C.bool(usesFontLeading))
}

func (l *NSLayoutManager) BackgroundLayoutEnabled() bool {
	return bool(C.LayoutManager_BackgroundLayoutEnabled(l.Ptr()))
}

func (l *NSLayoutManager) SetBackgroundLayoutEnabled(backgroundLayoutEnabled bool) {
	C.LayoutManager_SetBackgroundLayoutEnabled(l.Ptr(), C.bool(backgroundLayoutEnabled))
}

func (l *NSLayoutManager) LimitsLayoutForSuspiciousContents() bool {
	return bool(C.LayoutManager_LimitsLayoutForSuspiciousContents(l.Ptr()))
}

func (l *NSLayoutManager) SetLimitsLayoutForSuspiciousContents(limitsLayoutForSuspiciousContents bool) {
	C.LayoutManager_SetLimitsLayoutForSuspiciousContents(l.Ptr(), C.bool(limitsLayoutForSuspiciousContents))
}

func (l *NSLayoutManager) UsesDefaultHyphenation() bool {
	return bool(C.LayoutManager_UsesDefaultHyphenation(l.Ptr()))
}

func (l *NSLayoutManager) SetUsesDefaultHyphenation(usesDefaultHyphenation bool) {
	C.LayoutManager_SetUsesDefaultHyphenation(l.Ptr(), C.bool(usesDefaultHyphenation))
}

func NewLayoutManager() LayoutManager {
	return MakeLayoutManager(C.LayoutManager_NewLayoutManager())
}
