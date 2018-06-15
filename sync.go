package stdext

import "sync"

// TODO(DH): Docs

// ConcAtom is intended to be a similar construct to: http://clojure.org/reference/atoms
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
		// It's possible that the == operator will panic when comparing two
		// values of interface type that have the same dynamic type, when that
		// type is not "comparable", e.g. a map
		//
		// https://golang.org/ref/spec#Comparison_operators
		_ = recover()
		// false will be returned from the outer function.
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
