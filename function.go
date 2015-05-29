package gothic

type FunctionParam struct {
	name     string
	typename string
}

func NewFunctionParam(name, typename string) *FunctionParam {
	return &FunctionParam{name, typename}
}

type Function struct {
	name       string
	parameters []*FunctionParam
}

func NewFunction(name string) *Function {
	return &Function{name: name}
}

func (f Function) Name() string                     { return f.name }
func (f Function) Parameters() []*FunctionParam     { return f.parameters }
func (f Function) Parameter(idx int) *FunctionParam { return f.parameters[idx] }
func (f Function) NumParameters() int               { return len(f.parameters) }

func (f *Function) AppendParameter(param *FunctionParam) {
	f.parameters = append(f.parameters, param)
}
