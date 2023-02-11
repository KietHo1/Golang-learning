package main

import (
	"errors"
	"fmt"
	"log"
)

type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}

// `Insert` inserts new data into the tree, at the position determined by the search value.
// Return values:
//
// * `true` if the data was successfully inserted,
// * `false` if the data value already exists in the tree.
func (n *Node) Insert(value, data string) error {

	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {
	// If the data is already in the tree, return.
	case value == n.Value:
		return nil
	// If the data value is less than the current node's value, and if the left child node is `nil`, insert a new left child node. Else call `Insert` on the left subtree.
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
			return nil
		}
		return n.Left.Insert(value, data)
	// If the data value is greater than the current node's value, do the same but for the right subtree.
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}

// `Find` searches for a string. It returns:
//
// * The data associated with the value and `true`, or
// * "" and `false` if the search string is not found in the tree.
func (n *Node) Find(s string) (string, bool) {

	if n == nil {
		return "", false
	}

	switch {
	// If the current node contains the value, return the node.
	case s == n.Value:
		return n.Data, true
	// If the data value is less than the current node's value, call `Find` for the left child node,
	case s < n.Value:
		return n.Left.Find(s)
		// else call `Find` for the right child node.
	default:
		return n.Right.Find(s)
	}
}

// `findMax` finds the maximum element in a (sub-)tree. Its value replaces the value of the
// to-be-deleted node.
// Return values: the node itself and its parent node.
func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

// `replaceNode` replaces the `parent`'s child pointer to `n` with a pointer to the `replacement` node.
// `parent` must not be `nil`.
func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

// `Delete` removes an element from the tree.
// It is an error to try deleting an element that does not exist.
// In order to remove an element properly, `Delete` needs to know the node's parent node.
// `parent` must not be `nil`.
func (n *Node) Delete(s string, parent *Node) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in the tree")
	}

	// Search the node to be deleted.
	switch {
	case s < n.Value:
		return n.Left.Delete(s, n)
	case s > n.Value:
		return n.Right.Delete(s, n)
	default:
		// We found the node to be deleted.
		// If the node has no children, simply remove it from its parent.
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}

		// If the node has one child: Replace the node with its child.
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}

		// If the node has two children:
		// Find the maximum element in the left subtree...
		replacement, replParent := n.Left.findMax(n)

		//...and replace the node's value and data with the replacement's value and data.
		n.Value = replacement.Value
		n.Data = replacement.Data

		// Then remove the replacement node.
		return replacement.Delete(replacement.Value, replParent)
	}
}

// A `Tree` basically consists of a root node.
type Tree struct {
	Root *Node
}

// `Insert` calls `Node.Insert` unless the root node is `nil`
func (t *Tree) Insert(value, data string) error {
	// If the tree is empty, create a new node,...
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return nil
	}
	// ...else call `Node.Insert`.
	return t.Root.Insert(value, data)
}

// `Find` calls `Node.Find` unless the root node is `nil`
func (t *Tree) Find(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.Find(s)
}

// `Delete` has one special case: the empty tree. (And deleting from an empty tree is an error.)
// In all other cases, it calls `Node.Delete`.
func (t *Tree) Delete(s string) error {

	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}

	// Call`Node.Delete`. Passing a "fake" parent node here *almost* avoids
	// having to treat the root node as a special case, with one exception.
	fakeParent := &Node{Right: t.Root}
	err := t.Root.Delete(s, fakeParent)
	if err != nil {
		return err
	}
	// If the root node is the only node in the tree, and if it is deleted,
	// then it *only* got removed from `fakeParent`. `t.Root` still points to the old node.
	// We rectify this by setting t.Root to nil.
	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
}

// `Traverse` is a simple method that traverses the tree in left-to-right order
// (which, *by pure incidence* ;-), is the same as traversing from smallest to
// largest value) and calls a custom function on each node.
func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

/* ## A Couple Of Tree Operations
Our `main` function does a quick sort by filling a tree and reading it out again. Then it searches for a particular node. No fancy output to see here; this is just the proof that the whole code above works as it should.
*/

// `main`
func main() {

	// Set up a slice of strings.
	values := []string{"d", "b", "c", "e", "a"}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}

	// Create a tree and fill it from the values.
	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}

	// Print the sorted values.
	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	// Find values.
	s := "d"
	fmt.Print("Find node '", s, "': ")
	d, found := tree.Find(s)
	if !found {
		log.Fatal("Cannot find '" + s + "'")
	}
	fmt.Println("Found " + s + ": '" + d + "'")

	// Delete a value.
	err := tree.Delete(s)
	if err != nil {
		log.Fatal("Error deleting "+s+": ", err)
	}
	fmt.Print("After deleting '" + s + "': ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	// Special case: A single-node tree. (See `Tree.Delete` about why this is a special case.)
	fmt.Println("Single-node tree")
	tree = &Tree{}

	tree.Insert("a", "alpha")
	fmt.Println("After insert:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	tree.Delete("a")
	fmt.Println("After delete:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

}
