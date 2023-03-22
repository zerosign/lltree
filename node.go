package lltree

// Node repr for the list of the node
//
// A node can have a reference to it's own parent
// and it's siblings on in the same depth.
type node[T comparable] struct {
	val      T
	parent   *node[T]
	next     *node[T]
	prev     *node[T]
	children *node[T]
}

func (n *node[T]) Parent() (Node[T], bool) {

	if n.parent == nil {
		return nil, false
	}

	return n.parent, true
}

// impl current children for the node
func (n *node[T]) Children() (Cursor[T], bool) {

	if n.children == nil {

		return nil, false
	}

	return &cursor[T]{
		ref: n.children,
	}, true
}

// Parent is the root that contains tree of node
//
// Parent node doesn't store value by itself, but rather
// have a reference for the first child of the parents.
type parent[T comparable] struct {
	// this refs to the first node in the tree
	root node[T]
}
