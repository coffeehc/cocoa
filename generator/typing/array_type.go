package typing

import (
	"fmt"

	"github.com/hsiafan/cocoa/generator/base"
	"github.com/hsiafan/glow/stringx"
)

var _ Type = (*ArrayType)(nil)

// array types, the element should be StringType or ClassType
type ArrayType struct {
	Type Type
}

func (a *ArrayType) GoImports() base.StringSet {
	return a.Type.GoImports()
}

func (a *ArrayType) GoName(currentModule *Module) string {
	return "[]" + a.Type.GoName(currentModule)
}

func (a *ArrayType) CName() string {
	return "Array"
}

func (a *ArrayType) CgoName() string {
	return "C.Array"
}

func (a *ArrayType) ObjcName() string {
	return "NSArray*"
}

func (a *ArrayType) GoToCgoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	cName := "c" + stringx.Capitalize(param)
	codes := []string{
		fmt.Sprintf("%vData := make([]unsafe.Pointer, len(%v))", cName, param),
		fmt.Sprintf("for idx, v := range %v {", param),
	}
	itemCodes, itemResult := a.Type.GoToCgoCode(currentModule, "v")
	for _, itemCode := range itemCodes {
		codes = append(codes, "\t"+itemCode)
	}

	codes = append(codes,
		fmt.Sprintf("\t%vData[idx] = %v", cName, itemResult),
		"}",
		fmt.Sprintf("%v := C.Array{data:unsafe.Pointer(&%vData[0]), len:C.int(len(%v))}", cName, cName, param),
	)
	return codes, cName
}

func (a *ArrayType) CToObjcCode(param string) (convertCodes []string, result string) {
	objcName := "objc" + stringx.Capitalize(param)
	codes := []string{
		fmt.Sprintf("NSMutableArray* %v = [[NSMutableArray alloc] init];", objcName),
		fmt.Sprintf("void** %vData = (void**)%v.data;", param, param),
		fmt.Sprintf("for (int i = 0; i < %v.len; i++) {", param),
		fmt.Sprint("\tvoid* p = Data[i];"),
	}
	itemCodes, itemResult := a.Type.CToObjcCode("p")
	for _, itemCode := range itemCodes {
		codes = append(codes, "\t"+itemCode)
	}

	codes = append(codes,
		fmt.Sprintf("\t[%v addObject:(%v)%v];", objcName, a.Type.ObjcName(), itemResult),
		"}",
	)
	return codes, objcName
}

func (a *ArrayType) ObjcToCCode(param string) (convertCodes []string, result string) {
	cDataName := param + "Data"
	cCountName := param + "count"
	cResult := param + "Array"
	codes := []string{
		fmt.Sprintf("int %v = [%v count];", cCountName, param),
		fmt.Sprintf("void** %v = malloc(%v * sizeof(void*));", cDataName, cCountName),
		fmt.Sprintf("for (int i = 0; i < %v; i++) {", cCountName),
		fmt.Sprintf("\t void* p = [%v objectAtIndex:i];", param),
	}
	itemCodes, itemResult := a.Type.ObjcToCCode("p")
	for _, itemCode := range itemCodes {
		codes = append(codes, "\t"+itemCode)
	}
	codes = append(codes,
		fmt.Sprintf("\t %v[i] = %v;", cDataName, itemResult),
		"}",
		fmt.Sprintf("Array %v;", cResult),
		fmt.Sprintf("%v.data = %v;", cResult, cDataName),
		fmt.Sprintf("%v.len = %v;", cResult, cCountName),
	)
	return codes, cResult
}

func (a *ArrayType) CgoToGoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	goTempSlice := param + "Slice"
	goName := "go" + stringx.Capitalize(param)
	codes := []string{
		fmt.Sprintf("defer C.free(%v.data)", param),
		fmt.Sprintf("%v := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(%v.data))[:%v.len:%v.len]", goTempSlice, param, param, param),
		fmt.Sprintf("var %v = make(%v, len(%v))", goName, a.GoName(currentModule), goTempSlice),
		fmt.Sprintf("for idx, r := range %v {", goTempSlice),
	}

	itemCodes, itemResult := a.Type.CgoToGoCode(currentModule, "r")
	for _, itemCode := range itemCodes {
		codes = append(codes, "\t"+itemCode)
	}

	codes = append(codes,
		fmt.Sprintf("\t%v[idx] = %v", goName, itemResult),
		"}",
	)
	return codes, goName
}
