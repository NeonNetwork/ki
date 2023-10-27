package structure

import (
	"github.com/heartbytenet/bblib/containers/optionals"
	"github.com/heartbytenet/bblib/objects"
)

type BinaryTreeNode[T any] struct {
	value     T
	nextLeft  *BinaryTreeNode[T]
	nextRight *BinaryTreeNode[T]
}

func (node *BinaryTreeNode[T]) Value() T {
	return node.value
}

func (node *BinaryTreeNode[T]) Left() optionals.Optional[*BinaryTreeNode[T]] {
	if node.nextLeft == nil {
		return optionals.None[*BinaryTreeNode[T]]()
	}

	return optionals.From[*BinaryTreeNode[T]](node.nextLeft)
}

func (node *BinaryTreeNode[T]) Right() optionals.Optional[*BinaryTreeNode[T]] {
	if node.nextRight == nil {
		return optionals.None[*BinaryTreeNode[T]]()
	}

	return optionals.From[*BinaryTreeNode[T]](node.nextRight)
}

func (node *BinaryTreeNode[T]) Init() *BinaryTreeNode[T] {
	return node
}

func NewBinaryTreeNode[T any](value T) (node *BinaryTreeNode[T]) {
	node = objects.Init[BinaryTreeNode[T]](&BinaryTreeNode[T]{
		value:     value,
		nextLeft:  nil,
		nextRight: nil,
	})

	return
}

func (node *BinaryTreeNode[T]) AddLeft(value T) (result *BinaryTreeNode[T]) {
	result = NewBinaryTreeNode(value)

	node.nextLeft = result

	return
}

func (node *BinaryTreeNode[T]) AddRight(value T) (result *BinaryTreeNode[T]) {
	result = NewBinaryTreeNode(value)

	node.nextRight = result

	return
}
