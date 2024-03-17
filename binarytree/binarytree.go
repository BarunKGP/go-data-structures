package "binarytree"


type BTreeNode struct {
  Key int
  LeftChild *BTreeNode
  RightChild *BTreeNode
}

func isEmpty(node *BTreeNode) bool {
  if node == nil || node.Key == nil {
    return True
  }

  return False
}

func isEmpty(tree *BTree) bool {
  if tree == nil || isEmpty(tree.Head) {
    return True
  }

  return False
}

type BTree struct {
  Head *BTreeNode

}

func (b *BTree)getHeight() int {
  if isEmpty(b) {
    return 0
  } 

  root := b.Head
  height := 1 + max(b.Head.LeftChild.getHeight(), b.Head.RightChild.getHeight())

  return height
}

func bfs(b *BTree) {
  if isEmpty(b) {
    return
  }

  var res []int

  root := b.Head
  fmt.Print(root.Key)

  bfs(root.LeftChild)
  bfs(root.RightChild)

}


func createBinaryTree(nums []int) BTree {
  btree := new(BTree)
  root := []*BTreeNode
  root = append(*btree.Head)

  for n := range(nums) {
    curNodePtr := pop(root)  
    curNodePtr.Key = n
    curNodePtr = append(&curNodePtr.RightChild, &curNodePtr.LeftChild)
  }

  return btree
}

func (b *BTree)totalElems() int {

}




