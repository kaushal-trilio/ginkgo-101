package ginkgo_101

type Person struct {
	Age int
}

func (p *Person) IsChild() bool {
	return p.Age < 18
}
