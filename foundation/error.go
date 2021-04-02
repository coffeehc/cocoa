package foundation

import "github.com/hsiafan/cocoa/objc"

type Error interface {
	error
	objc.Object
	// Error code
	Code() int

	Domain() string
}

var _ Error = (*NSError)(nil)

type NSError struct {
	objc.NSObject
}

func (N *NSError) Error() string {
	panic("implement me")
}

func (N *NSError) Code() int {
	panic("implement me")
}

func (N *NSError) Domain() string {
	panic("implement me")
}
