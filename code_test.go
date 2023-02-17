package golang_interface_benchmark_test

import (
	code "github.com/deemson/golang-interface-benchmark"
	"testing"
)

func BenchmarkInterfaceWithStructImplementationMethodCall(b *testing.B) {
	i := code.NewInterfaceWithStructImplementation()
	for n := 0; n < b.N; n++ {
		i.Method()
	}
}

func BenchmarkInterfaceWithPointerImplementationMethodCall(b *testing.B) {
	i := code.NewInterfaceWithPointerImplementation()
	for n := 0; n < b.N; n++ {
		i.Method()
	}
}

func BenchmarkStructImplementationMethodCall(b *testing.B) {
	s := code.NewStructImplementation()
	for n := 0; n < b.N; n++ {
		s.Method()
	}
}

func BenchmarkPointerImplementationMethodCall(b *testing.B) {
	p := code.NewPointerImplementation()
	for n := 0; n < b.N; n++ {
		p.Method()
	}
}

func BenchmarkInterfaceWithStructImplementationInstantiationWithMethodCall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i := code.NewInterfaceWithStructImplementation()
		i.Method()
	}
}

func BenchmarkInterfaceWithPointerImplementationInstantiationWithMethodCall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i := code.NewInterfaceWithPointerImplementation()
		i.Method()
	}
}

func BenchmarkStructImplementationInstantiationWithMethodCall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := code.NewStructImplementation()
		s.Method()
	}
}

func BenchmarkPointerImplementationInstantiationWithMethodCall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		p := code.NewPointerImplementation()
		p.Method()
	}
}
