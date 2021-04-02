package typing

import (
	"fmt"

	"github.com/hsiafan/cocoa/generator/base"
)

var _ Type = (*AliasType)(nil)

// type def types
type AliasType struct {
	Type   Type   // the inner type
	Alias  string // alias name for new go type
	Module Module // used when Alias is not empty
}

func (a *AliasType) GoImports() base.StringSet {
	return a.Type.GoImports()
}

func (a *AliasType) GoName(currentModule *Module) string {
	return fullGoName(a.Module, a.Alias, *currentModule)
}

func (a *AliasType) CName() string {
	return a.Type.CName()
}

func (a *AliasType) CgoName() string {
	return a.Type.CgoName()
}

func (a *AliasType) ObjcName() string {
	return a.Type.ObjcName()
}

func (a *AliasType) GoToCgoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	param = fmt.Sprintf(a.Type.GoName(currentModule) + "(" + param + ")")
	return a.Type.GoToCgoCode(currentModule, param)
}

func (a *AliasType) CToObjcCode(param string) (convertCodes []string, result string) {
	return a.Type.CToObjcCode(param)
}

func (a *AliasType) ObjcToCCode(param string) (convertCodes []string, result string) {
	return a.Type.ObjcToCCode(param)
}

func (a *AliasType) CgoToGoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	codes, result := a.Type.CgoToGoCode(currentModule, param)
	result = fmt.Sprintf("%v({%v})", a.Alias, result)
	return codes, result
}
