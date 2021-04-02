package typing

func fullGoName(module Module, name string, currentModule Module) string {
	if module == currentModule {
		return name
	}
	return module.Package + "." + name
}

func prependPackage(module Module, s string, currentModule Module) string {
	if module == currentModule {
		return s
	}
	return module.Package + "." + s
}
