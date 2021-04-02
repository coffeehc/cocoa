package typing

import "github.com/hsiafan/cocoa/generator/base"

var _ Type = (*InstanceType)(nil)

// InstanceType the objc instancetype.
// class method start with alloc or new, instance method start with autorelease，init，retain or self, return instancetype.
type InstanceType struct {
}

func (i *InstanceType) GoImports() base.StringSet {
	panic("implement me")
}

func (i *InstanceType) GoName(currentModule *Module) string {
	panic("implement me")
}

func (i *InstanceType) CName() string {
	panic("implement me")
}

func (i *InstanceType) CgoName() string {
	panic("implement me")
}

func (i *InstanceType) ObjcName() string {
	panic("implement me")
}

func (i *InstanceType) GoToCgoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	panic("implement me")
}

func (i *InstanceType) CToObjcCode(param string) (convertCodes []string, result string) {
	panic("implement me")
}

func (i *InstanceType) ObjcToCCode(param string) (convertCodes []string, result string) {
	panic("implement me")
}

func (i *InstanceType) CgoToGoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	panic("implement me")
}
