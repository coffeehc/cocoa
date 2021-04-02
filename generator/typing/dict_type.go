package typing

import (
	"fmt"

	"github.com/hsiafan/cocoa/generator/base"
	"github.com/hsiafan/glow/stringx"
)

var _ Type = (*DictType)(nil)

// array types, the element should be StringType or ClassType
type DictType struct {
	KeyType   Type
	ValueType Type
}

func (a *DictType) GoImports() base.StringSet {
	var imports = base.StringSet{}
	imports.AddSet(a.KeyType.GoImports())
	imports.AddSet(a.ValueType.GoImports())
	return imports
}

func (a *DictType) GoName(currentModule *Module) string {
	return "map[" + a.KeyType.GoName(currentModule) + "]" + a.ValueType.GoName(currentModule)
}

func (a *DictType) CName() string {
	return "Dictionary"
}

func (a *DictType) CgoName() string {
	return "C.Dictionary"
}

func (a *DictType) ObjcName() string {
	return "NSDictionary*"
}

func (a *DictType) GoToCgoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	cName := "c" + stringx.Capitalize(param)
	codes := []string{
		fmt.Sprintf("%vKeyData := make([]unsafe.Pointer, len(%v))", cName, param),
		fmt.Sprintf("%vValueData := make([]unsafe.Pointer, len(%v))", cName, param),
		fmt.Sprintf("for k, v := range %v {", param),
	}
	keyItemCodes, keyItemResult := a.KeyType.GoToCgoCode(currentModule, "k")
	for _, itemCode := range keyItemCodes {
		codes = append(codes, "\t"+itemCode)
	}
	valueItemCodes, valueItemResult := a.ValueType.GoToCgoCode(currentModule, "v")
	for _, itemCode := range valueItemCodes {
		codes = append(codes, "\t"+itemCode)
	}

	codes = append(codes,
		fmt.Sprintf("\t%vKeyData[idx] = %v", cName, keyItemResult),
		fmt.Sprintf("\t%vValueData[idx] = %v", cName, valueItemResult),
		"}",
		fmt.Sprintf("%v := C.Dictionary{key_data:unsafe.Pointer(&%vKeyData[0]), va;ue_data:unsafe.Pointer(&%vValueData[0]), len:C.int(len(%v))}", cName, cName, cName, param),
	)
	return codes, cName
}

func (a *DictType) CToObjcCode(param string) (convertCodes []string, result string) {
	objcName := "objc" + stringx.Capitalize(param)
	codes := []string{
		fmt.Sprintf("NSMutableDictionary* %v = [[NSMutableDictionary alloc] initWithCapacity 10];", objcName),
		fmt.Sprintf("void** %vKeyData = (void**)%v.key_data;", param, param),
		fmt.Sprintf("void** %vValueData = (void**)%v.value_data;", param, param),
		fmt.Sprintf("for (int i = 0; i < %v.len; i++) {", param),
		fmt.Sprintf("\tvoid* kp = %vKeyData[i];", param),
		fmt.Sprintf("\tvoid* vp = %vValueData[i];", param),
	}
	keyItemCodes, keyItemResult := a.KeyType.CToObjcCode("kp")
	for _, itemCode := range keyItemCodes {
		codes = append(codes, "\t"+itemCode)
	}
	valueItemCodes, valueItemResult := a.KeyType.CToObjcCode("vp")
	for _, itemCode := range valueItemCodes {
		codes = append(codes, "\t"+itemCode)
	}

	codes = append(codes,
		fmt.Sprintf("\t[%v setObject:(%v)%v forKey:(%v)%v];", objcName, a.KeyType.ObjcName(), keyItemResult, a.ValueType.ObjcName(), valueItemResult),
		"}",
	)
	return codes, objcName
}

func (a *DictType) ObjcToCCode(param string) (convertCodes []string, result string) {
	keyDataName := param + "KeyData"
	valueDataName := param + "ValueData"
	cCountName := param + "Count"
	cResult := param + "Array"
	codes := []string{
		fmt.Sprintf("NSArray * %vKeys = [%v allKeys];", param, param),
		fmt.Sprintf("int %v = [%vKeys count];", cCountName, param),
		fmt.Sprintf("void** %v = malloc(%v * sizeof(void*));", keyDataName, cCountName),
		fmt.Sprintf("void** %v = malloc(%v * sizeof(void*));", valueDataName, cCountName),
		fmt.Sprintf("for (int i = 0; i < %v; i++) {", cCountName),
		fmt.Sprintf("\t %v kp = [%vKeys objectAtIndex:i];", a.KeyType.ObjcName(), param),
		fmt.Sprintf("\t %v vp = %v[kp];", a.ValueType.ObjcName(), param),
	}
	keyItemCodes, keyItemResult := a.KeyType.ObjcToCCode("kp")
	for _, itemCode := range keyItemCodes {
		codes = append(codes, "\t"+itemCode)
	}
	valueItemCodes, valueItemResult := a.KeyType.ObjcToCCode("vp")
	for _, itemCode := range valueItemCodes {
		codes = append(codes, "\t"+itemCode)
	}
	codes = append(codes,
		fmt.Sprintf("\t %v[i] = %v;", keyDataName, keyItemResult),
		fmt.Sprintf("\t %v[i] = %v;", valueDataName, valueItemResult),
		"}",
		fmt.Sprintf("Dictionary %v;", cResult),
		fmt.Sprintf("%v.key_data = %v;", cResult, keyDataName),
		fmt.Sprintf("%v.value_data = %v;", cResult, valueDataName),
		fmt.Sprintf("%v.len = %v;", cResult, cCountName),
	)
	return codes, cResult
}

func (a *DictType) CgoToGoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	goKeyTempSlice := param + "KeySlice"
	goValueTempSlice := param + "ValueSlice"
	goName := "go" + stringx.Capitalize(param)
	codes := []string{
		fmt.Sprintf("defer C.free(%v.key_data)", param),
		fmt.Sprintf("defer C.free(%v.value_data)", param),
		fmt.Sprintf("%v := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(%v.key_data))[:%v.len:%v.len]", goKeyTempSlice, param, param, param),
		fmt.Sprintf("%v := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(%v.value_data))[:%v.len:%v.len]", goValueTempSlice, param, param, param),
		fmt.Sprintf("var %v = make(%v)", goName, a.GoName(currentModule)),
		fmt.Sprintf("for idx, k := range %v {", goKeyTempSlice),
		fmt.Sprintf("\tv := %v[idx]", goValueTempSlice),
	}

	keyItemCodes, keyItemResult := a.KeyType.CgoToGoCode(currentModule, "k")
	for _, itemCode := range keyItemCodes {
		codes = append(codes, "\t"+itemCode)
	}
	valueItemCodes, valueItemResult := a.KeyType.CgoToGoCode(currentModule, "v")
	for _, itemCode := range valueItemCodes {
		codes = append(codes, "\t"+itemCode)
	}

	codes = append(codes,
		fmt.Sprintf("\t%v[%v] = %v", goName, keyItemResult, valueItemResult),
		"}",
	)
	return codes, goName
}
