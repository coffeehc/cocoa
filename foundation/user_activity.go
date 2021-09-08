package foundation

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
	RequiredUserInfoKeys() Set
	SetRequiredUserInfoKeys(value Set)
	TargetContentIdentifier() string
	SetTargetContentIdentifier(value string)
	NeedsSave() bool
	SetNeedsSave(value bool)
	Keywords() Set
	SetKeywords(value Set)
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
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	SupportsContinuationStreams() bool
	SetSupportsContinuationStreams(value bool)
	WebpageURL() URL
	SetWebpageURL(value URL)
	ReferrerURL() URL
	SetReferrerURL(value URL)
	ContextIdentifierPath() []string
}

type NSUserActivity struct {
	objc.NSObject
}

func MakeUserActivity(ptr unsafe.Pointer) NSUserActivity {
	return NSUserActivity{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSUserActivity) InitWithActivityType(activityType string) NSUserActivity {
	result_ := C.C_NSUserActivity_InitWithActivityType(n.Ptr(), NewString(activityType).Ptr())
	return MakeUserActivity(result_)
}

func AllocUserActivity() NSUserActivity {
	result_ := C.C_NSUserActivity_AllocUserActivity()
	return MakeUserActivity(result_)
}

func NewUserActivity() NSUserActivity {
	result_ := C.C_NSUserActivity_NewUserActivity()
	return MakeUserActivity(result_)
}

func (n NSUserActivity) Autorelease() NSUserActivity {
	result_ := C.C_NSUserActivity_Autorelease(n.Ptr())
	return MakeUserActivity(result_)
}

func (n NSUserActivity) Retain() NSUserActivity {
	result_ := C.C_NSUserActivity_Retain(n.Ptr())
	return MakeUserActivity(result_)
}

func (n NSUserActivity) BecomeCurrent() {
	C.C_NSUserActivity_BecomeCurrent(n.Ptr())
}

func (n NSUserActivity) ResignCurrent() {
	C.C_NSUserActivity_ResignCurrent(n.Ptr())
}

func (n NSUserActivity) Invalidate() {
	C.C_NSUserActivity_Invalidate(n.Ptr())
}

func (n NSUserActivity) ActivityType() string {
	result_ := C.C_NSUserActivity_ActivityType(n.Ptr())
	return MakeString(result_).String()
}

func (n NSUserActivity) Title() string {
	result_ := C.C_NSUserActivity_Title(n.Ptr())
	return MakeString(result_).String()
}

func (n NSUserActivity) SetTitle(value string) {
	C.C_NSUserActivity_SetTitle(n.Ptr(), NewString(value).Ptr())
}

func (n NSUserActivity) RequiredUserInfoKeys() Set {
	result_ := C.C_NSUserActivity_RequiredUserInfoKeys(n.Ptr())
	return MakeSet(result_)
}

func (n NSUserActivity) SetRequiredUserInfoKeys(value Set) {
	C.C_NSUserActivity_SetRequiredUserInfoKeys(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSUserActivity) TargetContentIdentifier() string {
	result_ := C.C_NSUserActivity_TargetContentIdentifier(n.Ptr())
	return MakeString(result_).String()
}

func (n NSUserActivity) SetTargetContentIdentifier(value string) {
	C.C_NSUserActivity_SetTargetContentIdentifier(n.Ptr(), NewString(value).Ptr())
}

func (n NSUserActivity) NeedsSave() bool {
	result_ := C.C_NSUserActivity_NeedsSave(n.Ptr())
	return bool(result_)
}

func (n NSUserActivity) SetNeedsSave(value bool) {
	C.C_NSUserActivity_SetNeedsSave(n.Ptr(), C.bool(value))
}

func (n NSUserActivity) Keywords() Set {
	result_ := C.C_NSUserActivity_Keywords(n.Ptr())
	return MakeSet(result_)
}

func (n NSUserActivity) SetKeywords(value Set) {
	C.C_NSUserActivity_SetKeywords(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSUserActivity) PersistentIdentifier() UserActivityPersistentIdentifier {
	result_ := C.C_NSUserActivity_PersistentIdentifier(n.Ptr())
	return UserActivityPersistentIdentifier(MakeString(result_).String())
}

func (n NSUserActivity) SetPersistentIdentifier(value UserActivityPersistentIdentifier) {
	C.C_NSUserActivity_SetPersistentIdentifier(n.Ptr(), NewString(string(value)).Ptr())
}

func (n NSUserActivity) IsEligibleForHandoff() bool {
	result_ := C.C_NSUserActivity_IsEligibleForHandoff(n.Ptr())
	return bool(result_)
}

func (n NSUserActivity) SetEligibleForHandoff(value bool) {
	C.C_NSUserActivity_SetEligibleForHandoff(n.Ptr(), C.bool(value))
}

func (n NSUserActivity) IsEligibleForSearch() bool {
	result_ := C.C_NSUserActivity_IsEligibleForSearch(n.Ptr())
	return bool(result_)
}

func (n NSUserActivity) SetEligibleForSearch(value bool) {
	C.C_NSUserActivity_SetEligibleForSearch(n.Ptr(), C.bool(value))
}

func (n NSUserActivity) IsEligibleForPublicIndexing() bool {
	result_ := C.C_NSUserActivity_IsEligibleForPublicIndexing(n.Ptr())
	return bool(result_)
}

func (n NSUserActivity) SetEligibleForPublicIndexing(value bool) {
	C.C_NSUserActivity_SetEligibleForPublicIndexing(n.Ptr(), C.bool(value))
}

func (n NSUserActivity) ExpirationDate() Date {
	result_ := C.C_NSUserActivity_ExpirationDate(n.Ptr())
	return MakeDate(result_)
}

func (n NSUserActivity) SetExpirationDate(value Date) {
	C.C_NSUserActivity_SetExpirationDate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSUserActivity) Delegate() objc.Object {
	result_ := C.C_NSUserActivity_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSUserActivity) SetDelegate(value objc.Object) {
	C.C_NSUserActivity_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSUserActivity) SupportsContinuationStreams() bool {
	result_ := C.C_NSUserActivity_SupportsContinuationStreams(n.Ptr())
	return bool(result_)
}

func (n NSUserActivity) SetSupportsContinuationStreams(value bool) {
	C.C_NSUserActivity_SetSupportsContinuationStreams(n.Ptr(), C.bool(value))
}

func (n NSUserActivity) WebpageURL() URL {
	result_ := C.C_NSUserActivity_WebpageURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSUserActivity) SetWebpageURL(value URL) {
	C.C_NSUserActivity_SetWebpageURL(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSUserActivity) ReferrerURL() URL {
	result_ := C.C_NSUserActivity_ReferrerURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSUserActivity) SetReferrerURL(value URL) {
	C.C_NSUserActivity_SetReferrerURL(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSUserActivity) ContextIdentifierPath() []string {
	result_ := C.C_NSUserActivity_ContextIdentifierPath(n.Ptr())
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
