package golang_interface_benchmark

type Interface interface {
	Method()
}

func NewInterfaceWithStructImplementation() Interface {
	return StructImplementation{}
}

func NewInterfaceWithPointerImplementation() Interface {
	return &PointerImplementation{}
}

// StructImplementation implements Interface with value receiver
type StructImplementation struct{}

func (s StructImplementation) Method() {}

func NewStructImplementation() StructImplementation {
	return StructImplementation{}
}

// PointerImplementation implements Interface with pointer receiver
type PointerImplementation struct{}

func (p *PointerImplementation) Method() {}

func NewPointerImplementation() *PointerImplementation {
	return &PointerImplementation{}
}
