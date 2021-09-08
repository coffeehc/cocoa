package delegates

import "github.com/hsiafan/cocoa/objc"

// GoDelegate is interface for go struct delegate impl
type GoDelegate interface {
	ToObjc() objc.Object
}

// HasDelegate is interface which can set delegate
type HasDelegate interface {
	SetDelegate(value objc.Object)
}

// Set is shortcut method for setting objc delegate using go protocol struct impl
func Set(obj HasDelegate, delegate GoDelegate) {
	obj.SetDelegate(delegate.ToObjc())
}
