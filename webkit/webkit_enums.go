package webkit

type ContentMode int

const (
	ContentModeRecommended ContentMode = iota
	ContentModeMobile
	ContentModeDesktop
)

type AudiovisualMediaTypes uint

const (
	AudiovisualMediaTypeNone  AudiovisualMediaTypes = 0
	AudiovisualMediaTypeAudio AudiovisualMediaTypes = 1 << 0
	AudiovisualMediaTypeVideo AudiovisualMediaTypes = 1 << 1
	AudiovisualMediaTypeAll   AudiovisualMediaTypes = 0xffffffffffffffff
)

type UserInterfaceDirectionPolicy int

const (
	UserInterfaceDirectionPolicyContent UserInterfaceDirectionPolicy = iota
	UserInterfaceDirectionPolicySystem
)

type UserScriptInjectionTime int

const (
	UserScriptInjectionTimeAtDocumentStart UserScriptInjectionTime = iota
	UserScriptInjectionTimeAtDocumentEnd
)

type NavigationType int

const (
	NavigationTypeOther           NavigationType = -1
	NavigationTypeLinkActivated   NavigationType = 0
	NavigationTypeFormSubmitted   NavigationType = 1
	NavigationTypeBackForward     NavigationType = 2
	NavigationTypeReload          NavigationType = 3
	NavigationTypeFormResubmitted NavigationType = 4
)
