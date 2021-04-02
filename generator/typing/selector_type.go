package typing

import (
	"fmt"

	"github.com/hsiafan/cocoa/generator/base"
)

var _ Type = (*SelectorType)(nil)

// objc selector type
type SelectorType struct {
}

func (s *SelectorType) GoImports() base.StringSet {
	return base.StringSet{"github.com/hsiafan/cocoa/" + ObjCRuntime.Package: struct{}{}}
}

func (s *SelectorType) GoName(currentModule *Module) string {
	return "*objc.Selector"
}

func (s *SelectorType) CName() string {
	return "void*"
}

func (s *SelectorType) CgoName() string {
	return "unsafe.Pointer"
}

func (s *SelectorType) ObjcName() string {
	return "SEL"
}

func (s *SelectorType) GoToCgoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("toPointer(%v)", param)
}

func (s *SelectorType) CToObjcCode(param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("(SEL)%v", param)
}

func (s *SelectorType) ObjcToCCode(param string) (convertCodes []string, result string) {
	return nil, param
}

func (s *SelectorType) CgoToGoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("objc.MakeSelector(%v)", param)
}
