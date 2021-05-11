package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "user_activity.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type UserActivity interface {
	objc.Object
	BecomeCurrent()
	ResignCurrent()
	Invalidate()
	ActivityType() string
	Title() string
	SetTitle(value string)
	TargetContentIdentifier() string
	SetTargetContentIdentifier(value string)
	NeedsSave() bool
	SetNeedsSave(value bool)
	PersistentIdentifier() UserActivityPersistentIdentifier
	SetPersistentIdentifier(value UserActivityPersistentIdentifier)
	IsEligibleForHandoff() bool
	SetEligibleForHandoff(value bool)
	IsEligibleForSearch() bool
	SetEligibleForSearch(value bool)
	IsEligibleForPublicIndexing() bool
	SetEligibleForPublicIndexing(value bool)
	ExpirationDate() Date
	SetExpirationDate(value Date)
	SupportsContinuationStreams() bool
	SetSupportsContinuationStreams(value bool)
	WebpageURL() URL
	SetWebpageURL(value URL)
	ReferrerURL() URL
	SetReferrerURL(value URL)
}

type NSUserActivity struct {
	objc.NSObject
}

func MakeUserActivity(ptr unsafe.Pointer) *NSUserActivity {
	if ptr == nil {
		return nil
	}
	return &NSUserActivity{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocUserActivity() *NSUserActivity {
	return MakeUserActivity(C.C_UserActivity_Alloc())
}

func (n *NSUserActivity) InitWithActivityType(activityType string) UserActivity {
	result := C.C_NSUserActivity_InitWithActivityType(n.Ptr(), NewString(activityType).Ptr())
	return MakeUserActivity(result)
}

func (n *NSUserActivity) BecomeCurrent() {
	C.C_NSUserActivity_BecomeCurrent(n.Ptr())
}

func (n *NSUserActivity) ResignCurrent() {
	C.C_NSUserActivity_ResignCurrent(n.Ptr())
}

func (n *NSUserActivity) Invalidate() {
	C.C_NSUserActivity_Invalidate(n.Ptr())
}

func (n *NSUserActivity) ActivityType() string {
	result := C.C_NSUserActivity_ActivityType(n.Ptr())
	return MakeString(result).String()
}

func (n *NSUserActivity) Title() string {
	result := C.C_NSUserActivity_Title(n.Ptr())
	return MakeString(result).String()
}

func (n *NSUserActivity) SetTitle(value string) {
	C.C_NSUserActivity_SetTitle(n.Ptr(), NewString(value).Ptr())
}

func (n *NSUserActivity) TargetContentIdentifier() string {
	result := C.C_NSUserActivity_TargetContentIdentifier(n.Ptr())
	return MakeString(result).String()
}

func (n *NSUserActivity) SetTargetContentIdentifier(value string) {
	C.C_NSUserActivity_SetTargetContentIdentifier(n.Ptr(), NewString(value).Ptr())
}

func (n *NSUserActivity) NeedsSave() bool {
	result := C.C_NSUserActivity_NeedsSave(n.Ptr())
	return bool(result)
}

func (n *NSUserActivity) SetNeedsSave(value bool) {
	C.C_NSUserActivity_SetNeedsSave(n.Ptr(), C.bool(value))
}

func (n *NSUserActivity) PersistentIdentifier() UserActivityPersistentIdentifier {
	result := C.C_NSUserActivity_PersistentIdentifier(n.Ptr())
	return UserActivityPersistentIdentifier(MakeString(result).String())
}

func (n *NSUserActivity) SetPersistentIdentifier(value UserActivityPersistentIdentifier) {
	C.C_NSUserActivity_SetPersistentIdentifier(n.Ptr(), NewString(string(value)).Ptr())
}

func (n *NSUserActivity) IsEligibleForHandoff() bool {
	result := C.C_NSUserActivity_IsEligibleForHandoff(n.Ptr())
	return bool(result)
}

func (n *NSUserActivity) SetEligibleForHandoff(value bool) {
	C.C_NSUserActivity_SetEligibleForHandoff(n.Ptr(), C.bool(value))
}

func (n *NSUserActivity) IsEligibleForSearch() bool {
	result := C.C_NSUserActivity_IsEligibleForSearch(n.Ptr())
	return bool(result)
}

func (n *NSUserActivity) SetEligibleForSearch(value bool) {
	C.C_NSUserActivity_SetEligibleForSearch(n.Ptr(), C.bool(value))
}

func (n *NSUserActivity) IsEligibleForPublicIndexing() bool {
	result := C.C_NSUserActivity_IsEligibleForPublicIndexing(n.Ptr())
	return bool(result)
}

func (n *NSUserActivity) SetEligibleForPublicIndexing(value bool) {
	C.C_NSUserActivity_SetEligibleForPublicIndexing(n.Ptr(), C.bool(value))
}

func (n *NSUserActivity) ExpirationDate() Date {
	result := C.C_NSUserActivity_ExpirationDate(n.Ptr())
	return MakeDate(result)
}

func (n *NSUserActivity) SetExpirationDate(value Date) {
	C.C_NSUserActivity_SetExpirationDate(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSUserActivity) SupportsContinuationStreams() bool {
	result := C.C_NSUserActivity_SupportsContinuationStreams(n.Ptr())
	return bool(result)
}

func (n *NSUserActivity) SetSupportsContinuationStreams(value bool) {
	C.C_NSUserActivity_SetSupportsContinuationStreams(n.Ptr(), C.bool(value))
}

func (n *NSUserActivity) WebpageURL() URL {
	result := C.C_NSUserActivity_WebpageURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSUserActivity) SetWebpageURL(value URL) {
	C.C_NSUserActivity_SetWebpageURL(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSUserActivity) ReferrerURL() URL {
	result := C.C_NSUserActivity_ReferrerURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSUserActivity) SetReferrerURL(value URL) {
	C.C_NSUserActivity_SetReferrerURL(n.Ptr(), objc.ExtractPtr(value))
}
