package typing

import (
	"fmt"

	"github.com/hsiafan/cocoa/generator/base"
)

var _ Type = (*StructType)(nil)

// struct type
type StructType struct {
	Name    string // c and objc type name
	GoName_ string // the go struct name
	Module  Module // used when Alias is not empty
}

func (s *StructType) GoImports() base.StringSet {
	return base.StringSet{"github.com/hsiafan/cocoa/" + s.Module.Package: struct{}{}}
}

func (s *StructType) GoName(currentModule *Module) string {
	return fullGoName(s.Module, s.GoName_, *currentModule)
}

func (s *StructType) CName() string {
	return s.Name
}

func (s *StructType) CgoName() string {
	return "C." + s.Name
}

func (s *StructType) ObjcName() string {
	return s.Name
}

func (s *StructType) GoToCgoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("to%v(%v)", s.Name, param)
}

func (s *StructType) CToObjcCode(param string) (convertCodes []string, result string) {
	return nil, param
}

func (s *StructType) ObjcToCCode(param string) (convertCodes []string, result string) {
	return nil, param
}

func (s *StructType) CgoToGoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("to%v(%v)", s.GoName_, param)
}
