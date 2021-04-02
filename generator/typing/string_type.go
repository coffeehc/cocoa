package typing

import (
	"fmt"

	"github.com/hsiafan/cocoa/generator/base"
)

var _ Type = (*StringType)(nil)

// string types
type StringType struct {
}

func (s *StringType) GoImports() base.StringSet {
	return base.StringSet{}
}

func (s *StringType) GoName(currentModule *Module) string {
	return "string"
}

func (s *StringType) CName() string {
	return "void*"
}

func (s *StringType) CgoName() string {
	return "unsafe.Pointer"
}

func (s *StringType) ObjcName() string {
	return "NSString*"
}

func (s *StringType) GoToCgoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	result = fmt.Sprintf("NewString(%v)", param)
	result = prependPackage(Foundation, result, *currentModule)
	return nil, result
}

func (s *StringType) CToObjcCode(param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("(NSString*)%v", param)
}

func (s *StringType) ObjcToCCode(param string) (convertCodes []string, result string) {
	return nil, param
}

func (s *StringType) CgoToGoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	result = fmt.Sprintf("MakeString(%v).String()", param)
	result = prependPackage(Foundation, result, *currentModule)
	return nil, result
}
