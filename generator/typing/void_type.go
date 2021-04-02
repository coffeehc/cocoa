package typing

import "github.com/hsiafan/cocoa/generator/base"

var _ Type = (*VoidType)(nil)

// objc binary data type
type VoidType struct {
}

func (d *VoidType) GoImports() base.StringSet {
	return base.StringSet{}
}

func (d *VoidType) GoName(currentModule *Module) string {
	return ""
}

func (d *VoidType) CName() string {
	return "void"
}

func (d *VoidType) CgoName() string {
	return ""
}

func (d *VoidType) ObjcName() string {
	return "void"
}

func (d *VoidType) GoToCgoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	return nil, ""
}

func (d *VoidType) CToObjcCode(param string) (convertCodes []string, result string) {
	return nil, ""
}

func (d *VoidType) ObjcToCCode(param string) (convertCodes []string, result string) {
	return nil, ""
}

func (d *VoidType) CgoToGoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	return nil, ""
}
