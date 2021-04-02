package gen

import (
	"fmt"
	"os"

	"github.com/hsiafan/cocoa/generator/base"
	"github.com/hsiafan/cocoa/generator/typing"
	"github.com/hsiafan/glow/stringx"
)

// Class is code generator for objc class
type Class struct {
	Type        *typing.ClassType
	Parent      *Class
	Properties  []*Property
	InitMethods []*Method
	Methods     []*Method
}

func (c *Class) Init() {
	c.mergeParentInitMethods()
}

func (c *Class) mergeParentInitMethods() {
	if c.Parent != nil {
		c.InitMethods = append(c.InitMethods, c.Parent.InitMethods...)
	}
}

func (c *Class) fileName() string {
	return stringx.CamelToSnake(c.Type.GName)
}

func (c *Class) InterfaceName() string {
	return "I" + c.Type.GName
}

// WriteGoFile generate go source file
func (c *Class) WriteGoFile(fileDir string) {
	goFilePath := fileDir + "/" + c.fileName() + ".go"
	w, err := os.Open(goFilePath)
	if err != nil {
		panic(err)
	}
	cw := &CodeWriter{Writer: w, IndentStr: "\t"}
	c.writeGoPackageAndImports(cw)
	c.writeGoInterface(cw)
	c.writeGoStruct(cw)
}

func (c *Class) writeGoPackageAndImports(w *CodeWriter) {
	w.WriteLines([]string{
		"package" + c.Type.Module.Package,
		"// #cgo CFLAGS: -x objective-c",
		"// #cgo LDFLAGS: -framework " + c.Type.Module.Name,
		fmt.Sprintf("// #include \"%s.h\"", c.fileName()),
		"import \"C\"",
		"import (",
	})
	w.Indent()

	imports := &base.StringSet{}
	imports.Add("unsafe")
	for _, m := range c.Methods {
		imports.AddSet(m.GoImports())
	}
	imports.ForEach(func(value string) {
		w.WriteLine("\"" + value + "\"")
	})
	w.UnIndent()

	w.WriteLine(")")
}

func (c *Class) writeGoInterface(w *CodeWriter) {
	w.WriteLine("interface " + c.InterfaceName() + " {")
	w.Indent()

	if c.Parent != nil {
		w.WriteLine(c.Parent.InterfaceName())
	}

	for _, m := range c.Methods {
		if !m.Static {
			m.WriteGoInterfaceCode(c.Type.Module, c.Type, w)
		}
	}

	w.UnIndent()
	w.WriteLine("}")
}

func (c *Class) writeGoStruct(w *CodeWriter) {
	w.WriteLine("type " + c.Type.GName + " struct {")
	w.Indent()
	if c.Parent != nil {
		w.WriteLine(c.Parent.Type.GName)
	}
	w.UnIndent()
	w.WriteLine("}")

	// make func
	w.WriteLine(fmt.Sprintf("func Make%s(ptr unsafe.Pointer) *%s {", c.Type.GName, c.Type.GName))
	w.Indent()
	w.WriteLines([]string{
		"if ptr == nil {",
		"\treturn nil",
		"}",
		fmt.Sprintf("return &%s{", c.Type.GName),
	})
	if c.Parent != nil {
		pType := c.Parent.Type
		w.WriteLine(fmt.Sprintf("\t%s: *%s.Make%s(ptr),", pType.GName, pType.Module.Package, pType.GName))
	}
	w.WriteLine("}")
	w.UnIndent()
	w.WriteLine("}")

	// methods
	for _, m := range c.Methods {
		im := m.ToInitMethod(c.Type)
		im.WriteGoCallCode(c.Type.Module, c.Type, w)
	}

	for _, m := range c.Methods {
		m.WriteGoCallCode(c.Type.Module, c.Type, w)
	}

}

// WriteCHeaderFile generate c header source file
func (c *Class) WriteCHeaderFile(fileDir string) {
	goFilePath := fileDir + "/" + c.fileName() + ".h"
	w, err := os.Open(goFilePath)
	if err != nil {
		panic(err)
	}
	cw := &CodeWriter{Writer: w, IndentStr: "\t"}
	cw.WriteLines([]string{
		"#import <stdbool.h>",
		"#import <stdlib.h>",
		"#import <utils.h>",
	})

	// methods
	for _, m := range c.Methods {
		im := m.ToInitMethod(c.Type)
		im.WriteHeaderFunc(c.Type.Module, c.Type, cw)
	}

	for _, m := range c.Methods {
		m.WriteHeaderFunc(c.Type.Module, c.Type, cw)
	}
}

// WriteMFile generate objc wrapper source file
func (c *Class) WriteMFile(fileDir string) {
	goFilePath := fileDir + "/" + c.fileName() + ".m"
	w, err := os.Open(goFilePath)
	if err != nil {
		panic(err)
	}
	cw := &CodeWriter{Writer: w, IndentStr: "\t"}
	cw.WriteLines([]string{
		fmt.Sprintf("#import <%s>", c.Type.Module.Header),
		"#import \"url.h\"",
	})

	// methods
	for _, m := range c.Methods {
		im := m.ToInitMethod(c.Type)
		im.WriteCFuncCode(c.Type.Module, c.Type, cw)
	}

	for _, m := range c.Methods {
		m.WriteCFuncCode(c.Type.Module, c.Type, cw)
	}
}
