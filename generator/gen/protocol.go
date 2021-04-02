package gen

// Protocol is code generator for objc protocol
type Protocol struct {
	Name        string
	Pkg         string
	Properties  []*Property
	InitMethods []*Method
	Methods     []*Method
}
