package gen

import (
	"github.com/hsiafan/cocoa/generator/typing"
	"github.com/hsiafan/glow/stringx"
)

// Property is code generator for objective-c property
type Property struct {
	Name       string
	ReadOnly   bool
	GetterName string
	Type       typing.Type
	Static     bool
}

// WriteGoCallCode generate go method code to call c wrapper code
func (p *Property) WriteGoCallCode(currentModule *typing.Module, classType *typing.ClassType, goWriter *CodeWriter) {
	p.getter().WriteGoCallCode(currentModule, classType, goWriter)
	if !p.ReadOnly {
		p.setter().WriteGoCallCode(currentModule, classType, goWriter)
	}
}

// WriteGoInterfaceCode generate go interface method signature code
func (p *Property) WriteGoInterfaceCode(currentModule *typing.Module, classType *typing.ClassType, goWriter *CodeWriter) {
	p.getter().WriteGoInterfaceCode(currentModule, classType, goWriter)
	if !p.ReadOnly {
		p.setter().WriteGoInterfaceCode(currentModule, classType, goWriter)
	}
}

// WriteHeaderFunc generate c header signature
func (p *Property) WriteHeaderFunc(currentModule *typing.Module, classType *typing.ClassType, mWriter *CodeWriter) {
	p.getter().WriteHeaderFunc(currentModule, classType, mWriter)
	if !p.ReadOnly {
		p.setter().WriteHeaderFunc(currentModule, classType, mWriter)
	}
}

// WriteCFuncCode generate c wrapper funtion
func (p *Property) WriteCFuncCode(currentModule *typing.Module, classType *typing.ClassType, mWriter *CodeWriter) {
	p.getter().WriteCFuncCode(currentModule, classType, mWriter)
	if !p.ReadOnly {
		p.setter().WriteCFuncCode(currentModule, classType, mWriter)
	}
}

func (p *Property) getter() *Method {
	return &Method{
		Name:       p.GetterName,
		ReturnType: p.Type,
		Static:     p.Static,
	}
}

func (p *Property) setter() *Method {
	return &Method{
		Name:       "set" + stringx.Capitalize(p.Name),
		Params:     []*Param{{Name: "value", Type: p.Type}},
		ReturnType: &typing.VoidType{},
		Static:     p.Static,
	}
}
