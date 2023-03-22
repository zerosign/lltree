package lltree

// Thjis stores current node reference before the node are being
// deleted from the tree.
//
// Node, refernece that being stored in here doesn't guarantee whether
// the pointer aren't being IGCed.
type Link[T comparable] struct {
	Parent Node[T]
	Next   Node[T]
	Prev   Node[T]
}

// interface for changing a value of a node
type EditValue[T comparable] interface {

	// change current node value
	Change(other T) bool
}

// any operation that can
type EditNode[T comparable] interface {

	// delete current node
	// if the node are being deleted, it will unref
	//
	// This will returns :
	// - current node
	// - any links that current node have
	// - flag whther the operatiion success or not
	//
	Delete() (Node[T], Link[T], bool)

	// replace current node with other node
	// other node links are intact, thus it means
	// current node links will be unreferenced#
	Replace(other Node[T]) bool
}

// CursorMut interface are interface that abstract mutable operation
// of current cursor.
type CursorMut[T comparable] interface{}

// Cursor are being used for iterating list of nodes
//
// cursor has reference to the parent of current list.
//
// Cursor are stateful iterator that can traverse a node.
type Cursor[T comparable] interface {

	// Get current parent of the current cursor list
	Parent() (Node[T], bool)

	// Get current next node in the current cursor of the list
	Next() (Node[T], bool)

	// Get previous node in the current cursor of the list
	Previous() (Node[T], bool)
}

type Parent[T comparable] interface {

	// Get current children repr as cursor
	Children() (Cursor[T], bool)
}

// repr of node interface
type Node[T comparable] interface {

	// node has children
	Parent[T]

	// node can be edited
	// EditNode(T)
}
