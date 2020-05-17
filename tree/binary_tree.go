package tree

type BinaryTree struct {
	Data                interface{}
	Left, Right, Parent *BinaryTree
}

func (b *BinaryTree) String() string {
	s, err := binaryTreeToString(b)
	if err != nil {
		panic(err)
	}
	return s
}

func (b BinaryTree) GetData() interface{} {
	return b.Data
}

func (b BinaryTree) GetLeft() TreeLike {
	return b.Left
}

func (b BinaryTree) GetRight() TreeLike {
	return b.Right
}
