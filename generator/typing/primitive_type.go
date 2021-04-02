package typing

import (
	"fmt"

	"github.com/hsiafan/cocoa/generator/base"
)

var _ Type = (*PrimitiveType)(nil)

var Bool = &PrimitiveType{Name: "bool", CgoName_: "bool", GoName_: "bool"}

var Int = &PrimitiveType{Name: "int", CgoName_: "int", GoName_: "int"}
var UInt = &PrimitiveType{Name: "unsigned int", CgoName_: "uint", GoName_: "uint"}

var Int64 = &PrimitiveType{Name: "long", CgoName_: "long", GoName_: "int64"}
var UInt64 = &PrimitiveType{Name: "unsigned long", CgoName_: "ulong", GoName_: "uint64"}

var Float32 = &PrimitiveType{Name: "float", CgoName_: "float", GoName_: "float32"}
var Float64 = &PrimitiveType{Name: "double", CgoName_: "double", GoName_: "float64"}

// primitive types
type PrimitiveType struct {
	Name     string // c type name, also can be used for objc
	CgoName_ string // c go type name
	GoName_  string // go type name
}

func (p *PrimitiveType) GoImports() base.StringSet {
	return base.StringSet{}
}

func (p *PrimitiveType) GoName(currentModule *Module) string {
	return p.GoName_
}

func (p *PrimitiveType) CName() string {
	return p.Name
}

func (p *PrimitiveType) CgoName() string {
	return p.CgoName_
}

func (p *PrimitiveType) ObjcName() string {
	return p.Name
}

func (p *PrimitiveType) GoToCgoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("C.%v(%v)", p.CgoName_, param)
}

func (p *PrimitiveType) CToObjcCode(param string) (convertCodes []string, result string) {
	return nil, param
}

func (p *PrimitiveType) ObjcToCCode(param string) (convertCodes []string, result string) {
	return nil, param
}

func (p *PrimitiveType) CgoToGoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("%v(%v)", p.GoName_, param)
}
