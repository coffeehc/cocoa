package typing

import (
	"fmt"

	"github.com/hsiafan/cocoa/generator/base"
	"github.com/hsiafan/glow/stringx"
)

var _ Type = (*DataType)(nil)

// objc binary data type
type DataType struct {
}

func (d *DataType) GoImports() base.StringSet {
	return base.StringSet{}
}

func (d *DataType) GoName(currentModule *Module) string {
	return "[]byte"
}

func (d *DataType) CName() string {
	return "Array"
}

func (d *DataType) CgoName() string {
	return "C.Array"
}

func (d *DataType) ObjcName() string {
	return "NSData*"
}

func (d *DataType) GoToCgoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("C.Array{{data:unsafe.Pointer(&{%v}[0]), len:C.int(len({%v}))}}", param, param)
}

func (d *DataType) CToObjcCode(param string) (convertCodes []string, result string) {
	return nil, fmt.Sprintf("[[NSData alloc] initWithBytes:(Byte *){%v}.data length:{%v}.len]", param, param)
}

func (d *DataType) ObjcToCCode(param string) (convertCodes []string, result string) {
	arrayName := param + "array"
	codes := []string{
		fmt.Sprintf("Array %v;", arrayName),
		fmt.Sprintf("%v.data = [%v bytes];", arrayName, param),
		fmt.Sprintf("%v.len = %v.length;", arrayName, param),
	}
	return codes, arrayName
}

func (d *DataType) CgoToGoCode(currentModule *Module, param string) (convertCodes []string, result string) {
	bufferName := param + "Buffer"
	resultName := "go" + stringx.Capitalize(param)
	codes := []string{
		fmt.Sprintf("{%v} := (*[1 << 30]byte)(%v.data)[:C.int(%v.len)]", bufferName, param, param),
		fmt.Sprintf("%v := make([]byte, C.int(%v.len))", resultName, param),
		fmt.Sprintf("copy(%v, %v)", resultName, bufferName),
		fmt.Sprintf("C.free(%v.data)", param),
	}
	return codes, resultName
}
