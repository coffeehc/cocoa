package typing

import (
	"fmt"

	"github.com/hsiafan/cocoa/generator/base"
)

var _ Type = (*ClassType)(nil)

// Objective-c interface type
type ClassType struct {
	Name   string  // objc type name
	GName  string  // Go name, usually is objc type name without prefix 'NS'
	Module *Module // object-c module
}

func (c *ClassType) GoImports() base.StringSet {
	return base.StringSet{"github.com/hsiafan/cocoa/" + c.Module.Package: struct{}{}}
}

func (c *ClassType) GoName(currentModule *Module) string {
	return fullGoName(*c.Module, c.GName, *currentModule)
}

func (c *ClassType) CName() string {
	return "void*"
}

func (c *ClassType) CgoName() string {
	return "unsafe.Pointer"
}

func (c *ClassType) ObjcName() string {
	return c.Name
}

func (c *ClassType) GoToCgoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("toPointer(%v)", param)
}

func (c *ClassType) CToObjcCode(param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("(%v*)%v", c.GName, param)
}

func (c *ClassType) ObjcToCCode(param string) (convertCodes []string, result string) {
	return nil, param
}

func (c *ClassType) CgoToGoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	if *currentModule == *c.Module {
		return nil, fmt.Sprintf("Make%v(%v)", c.GName, param)
	}
	return nil, fmt.Sprintf("%v.Make%v(%v)", c.Module.Package, c.GName, param)
}
