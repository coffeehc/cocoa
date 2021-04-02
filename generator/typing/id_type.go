package typing

import (
	"fmt"

	"github.com/hsiafan/cocoa/generator/base"
)

var _ Type = (*IDType)(nil)

// Objective-c interface type
type IDType struct {
}

func (c *IDType) GoImports() base.StringSet {
	return base.StringSet{"github.com/hsiafan/cocoa/" + ObjCRuntime.Package: struct{}{}}
}

func (c *IDType) GoName(currentModule *Module) string {
	if ObjCRuntime == *currentModule {
		return "Object"
	}
	return "objc.Object"
}

func (c *IDType) CName() string {
	return "void*"
}

func (c *IDType) CgoName() string {
	return "unsafe.Pointer"
}

func (c *IDType) ObjcName() string {
	return "id"
}

func (c *IDType) GoToCgoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("toPointer(%v)", param)
}

func (c *IDType) CToObjcCode(param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("(id)%v", param)
}

func (c *IDType) ObjcToCCode(param string) (convertCodes []string, result string) {
	return nil, param
}

func (c *IDType) CgoToGoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	if *currentModule == ObjCRuntime {
		return nil, fmt.Sprintf("MakeObject(%v)", param)
	}
	return nil, fmt.Sprintf("objc.MakeObject(%v)", param)
}
