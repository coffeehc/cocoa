package gen

import "github.com/hsiafan/cocoa/generator/typing"

// Param is code generator for objective-c method param
type Param struct {
	Name      string
	Type      typing.Type
	FieldName string // objc param field name(part of function name)
}

// GoDeclare return go param declare code
func (p *Param) GoDeclare(currentModule *typing.Module) string {
	return p.Name + " " + p.Type.GoName(currentModule)
}

func (p *Param) CDeclare() string {
	return p.Name + " " + p.Type.CName()
}

func (p *Param) ObjcDeclare() string {
	return p.Name + " " + p.Type.ObjcName()
}

func (p *Param) CgoDeclare() string {
	return p.Name + " " + p.Type.CgoName()
}
