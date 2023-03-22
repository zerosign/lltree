package lltree

type cursor[T comparable] struct {
	// we don't need to checks whether current ref node is nil or not
	// since it's already checked on caller side
	ref *node[T]
}

func (c *cursor[T]) Previous() (Node[T], bool) {

	// this means we are already on the left most in the list
	if c.ref.prev == nil {
		return nil, false
	}

	return c.ref.prev, true
}

func (c *cursor[T]) Next() (Node[T], bool) {

	if c.ref.next == nil {
		return nil, false
	}

	return c.ref.next, false
}

func (c *cursor[T]) Parent() (Node[T], bool) {

	if c.ref.parent == nil {
		return nil, false
	}

	return c.ref.parent, true
}
