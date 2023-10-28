package structure

type Pair[A any, B any] struct {
	a A
	b B
}

func (pair Pair[A, B]) A() A {
	return pair.a
}

func (pair Pair[A, B]) B() B {
	return pair.b
}

func NewPair[A any, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{a: a, b: b}
}
