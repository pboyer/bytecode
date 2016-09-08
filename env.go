package main

type symbol struct {
	pos        int
	symbolType string
	node       N
}

type env struct {
	parent *env
	data   map[string]*symbol
}

func newEnv(parent *env) *env {
	return &env{
		parent: parent,
		data:   make(map[string]*symbol),
	}
}

func (e *env) root() *env {
	if e.parent != nil {
		return e.parent.root()
	}
	return e
}

func (e *env) lookupRec(id string) (*symbol, bool) {
	s, ok := e.data[id]
	if ok {
		return s, true
	}

	if e.parent == nil {
		return nil, false
	}

	return e.parent.lookupRec(id)
}

func (e *env) lookup(id string) (*symbol, bool) {
	s, ok := e.data[id]
	return s, ok
}
