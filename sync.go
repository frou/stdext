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

// TODO(DH): Rename to Swap?
func (a *ConcAtom) Replace(val interface{}) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.val = val
}

func (a *ConcAtom) CompareAndSwap(expected, val interface{}) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	defer func() {
		// It's possible that the == operator panicked with two interface value
		// operands: https://golang.org/ref/spec#Comparison_operators
		recover()
	}()

	if a.val == expected {
		a.val = val
		return true
	}
	return false
}

func (a *ConcAtom) Advance(old2new func(interface{}) interface{}) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.val = old2new(a.val)
}
