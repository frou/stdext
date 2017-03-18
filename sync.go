package stdext

import "sync"

// TODO(DH): Docs

// Clojure-esque construct: http://clojure.org/reference/atoms
type ConcAtom struct {
	mu  sync.Mutex
	val interface{}
}

func NewConcAtom(initVal interface{}) *ConcAtom {
	a := new(ConcAtom)
	a.Replace(initVal)
	return a
}

func (a *ConcAtom) Deref() interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.val
}

func (a *ConcAtom) Replace(val interface{}) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.val = val
}

func (a *ConcAtom) Progress(f func(interface{}) interface{}) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.val = f(a.val)
}
