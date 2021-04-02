package gen

import (
	"strings"

	"github.com/hsiafan/cocoa/generator/base"
	"github.com/hsiafan/cocoa/generator/typing"
	"github.com/hsiafan/glow/stringx"
)

// Method is code generator for objective-c method
type Method struct {
	Name       string
	Params     []*Param
	ReturnType typing.Type
	Static     bool // true if is class method
	goFuncName string
}

// ToInitMethod return new init method.
func (m *Method) ToInitMethod(returnType *typing.ClassType) *Method {
	return &Method{
		Name:       m.Name,
		Params:     m.Params,
		ReturnType: returnType,
		Static:     m.Static,
		goFuncName: m.goFuncName,
	}
}

// WriteGoCallCode generate go method code to call c wrapper code
func (m *Method) WriteGoCallCode(currentModule *typing.Module, classType *typing.ClassType, goWriter *CodeWriter) {
	funcDeclare := m.GoFuncDeclare(currentModule)
	receiver := strings.ToLower(classType.Name[0:1])

	var args []string

	if m.Static {
		goWriter.WriteLine("func " + funcDeclare + "{")
	} else {
		goWriter.WriteLine("func (" + receiver + " " + classType.Name + ")" + funcDeclare + "{")
		args = append(args, receiver+".Ptr()")
	}

	goWriter.Indent()

	for _, p := range m.Params {
		pCodes, pResult := p.Type.GoToCgoCode(currentModule, p.Name)
		goWriter.WriteLines(pCodes)
		args = append(args, pResult)
	}

	callCode := "C." + m.CFuncName() + "(" + strings.Join(args, ", ") + ")"
	if _, ok := m.ReturnType.(*typing.VoidType); ok {
		goWriter.WriteLine(callCode)
	} else {
		var resultName = "result"
		goWriter.WriteLine(resultName + " := " + callCode)
		rCodes, rResult := m.ReturnType.CgoToGoCode(currentModule, resultName)
		goWriter.WriteLines(rCodes)
		goWriter.WriteLine("return " + rResult)
	}
	goWriter.UnIndent()
	goWriter.WriteLine("}")
}

// WriteGoInterfaceCode generate go interface method signature code
func (m *Method) WriteGoInterfaceCode(currentModule *typing.Module, classType *typing.ClassType, goWriter *CodeWriter) {
	if m.Static {
		return
	}
	funcDeclare := m.GoFuncDeclare(currentModule)
	goWriter.WriteLine(funcDeclare)
}

// GoFuncDeclare generate go function declaration
func (m *Method) GoFuncDeclare(currentModule *typing.Module) string {
	var paramStrs []string
	for _, p := range m.Params {
		paramStrs = append(paramStrs, p.GoDeclare(currentModule))
	}

	return "func " + m.GoFuncName() + "(" + strings.Join(paramStrs, ",") + ")" + " " + m.ReturnType.GoName(currentModule)
}

// GoFuncName return go func name
func (m *Method) GoFuncName() string {
	if m.goFuncName == "" {
		var sb strings.Builder
		sb.WriteString(m.Name)
		for idx, p := range m.Params {
			if idx == 0 {
				continue
			}
			sb.WriteByte('_')
			sb.WriteString(p.FieldName)
		}
		m.goFuncName = sb.String()
	}

	return m.goFuncName
}

// GoImports return all imports for go file
func (m *Method) GoImports() base.StringSet {
	var imports = base.StringSet{}
	for _, param := range m.Params {
		imports.AddSet(param.Type.GoImports())
	}
	imports.AddSet(m.ReturnType.GoImports())
	return imports
}

// WriteHeaderFunc generate c header signature
func (m *Method) WriteHeaderFunc(currentModule *typing.Module, classType *typing.ClassType, mWriter *CodeWriter) {
	mWriter.WriteLine(m.CFuncDeclare(currentModule) + ";")
}

// WriteCFuncCode generate c wrapper funtion
func (m *Method) WriteCFuncCode(currentModule *typing.Module, classType *typing.ClassType, mWriter *CodeWriter) {
	var paramStrs []string
	for _, p := range m.Params {
		paramStrs = append(paramStrs, p.CDeclare())
	}

	mWriter.WriteLine(m.CFuncDeclare(currentModule) + "{")
	mWriter.Indent()

	var callCodeBuilder strings.Builder
	callCodeBuilder.WriteString("[")
	if !m.Static {
		instanceName := stringx.DeCapitalize(classType.Name)
		mWriter.WriteLine(classType.ObjcName() + "* " + instanceName + " = (" + classType.ObjcName() + "*)ptr")
		callCodeBuilder.WriteString(instanceName)
	} else {
		callCodeBuilder.WriteString(classType.ObjcName())
	}
	for _, p := range m.Params {
		pCodes, pResult := p.Type.CToObjcCode(p.Name)
		mWriter.WriteLines(pCodes)
		callCodeBuilder.WriteRune(' ')
		callCodeBuilder.WriteString(p.Type.ObjcName())
		callCodeBuilder.WriteRune(':')
		callCodeBuilder.WriteString(pResult)
	}
	callCodeBuilder.WriteString("]")

	callCode := callCodeBuilder.String()
	if _, ok := m.ReturnType.(*typing.VoidType); ok {
		mWriter.WriteLine(callCode)
	} else {
		var resultName = "result"
		mWriter.WriteLine(m.ReturnType.ObjcName() + " " + resultName + " = " + callCode)
		rCodes, rResult := m.ReturnType.ObjcToCCode(resultName)
		mWriter.WriteLines(rCodes)
		mWriter.WriteLine("return " + rResult)
	}
	mWriter.UnIndent()
}

// CFuncDeclare generate c wrapper funtion signature
func (m *Method) CFuncDeclare(currentModule *typing.Module) string {
	var paramStrs []string
	if !m.Static {
		paramStrs = append(paramStrs, "void* ptr")
	}
	for _, p := range m.Params {
		paramStrs = append(paramStrs, p.CDeclare())
	}

	return m.ReturnType.CName() + " " + m.CFuncName() + "(" + strings.Join(paramStrs, ",") + ")"
}

// CFuncName return func name for c wrapper code
func (m *Method) CFuncName() string {
	return "C_" + m.GoFuncName()
}
