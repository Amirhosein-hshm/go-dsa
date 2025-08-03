package trees

import "iter"

type node[T int] struct {
	left_child  *node[T]
	right_child *node[T]
	v           T
}

type Tree[T int] struct {
	root *node[T]
}

func (t *Tree[T]) Insert(val T) {
	if t.root == nil {
		t.root = &node[T]{v: val}
		return
	}

	current := t.root
	for {
		if current.v > val {
			if current.left_child == nil {
				current.left_child = &node[T]{v: val}
				return
			}
			current = current.left_child
		} else {
			if current.right_child == nil {
				current.right_child = &node[T]{v: val}
				return
			}
			current = current.right_child
		}
	}
}

func recursion[T int](yield func(T) bool, n *node[T]) bool {
	if n.left_child != nil {
		if !recursion(yield, n.left_child) {
			return false
		}
	}
	if !yield(n.v) {
		return false
	}
	if n.right_child != nil {
		if !recursion(yield, n.right_child) {
			return false
		}
	}
	return true
}

func (t *Tree[T]) PreOrder() iter.Seq[T] {
	return func(yield func(T) bool) {
		recursion(yield, t.root)
	}
}
