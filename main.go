package main

import "fmt"



func main() {
  fmt.Println("Welcome to Data Structures written in Go!")
}


type Node struct {
  Key int
  LeftChild *Node
  RightChild *Node
}

func isValidKey(key int, node *Node) bool {
  if (node == nil) {
    return True
  }


  if (node.LeftChild != nil) && (key > node.LeftChild.Key) {
    return False
  
  }

  if (node.RightChild != nil) && (key > node.RightChild.Key) {
        return False
    
    }

  return True
}

type Heap struct {
  head *Node
  Len int
}

func swapNodes(n1 *Node, n2 *Node) err {
if (n1 == nil ) && (n2 == nil) {
    return
  }

  if (n1 == nil) || (n2 == nil) {
    return fmt.Error("One of the nodes is empty")
  }

  n1New = Node{key=n2.Key, LeftChild=n2.LeftChild, RightChild=n2.RightChild}
  // n2New = Node(key=n1.Key, LeftChild=n1.LeftChild, RightChild=n1.RightChild)

  n2.Key = n1.Key
  n2.LeftChild = n1.LeftChild
  n2.RightChild = n1.RightChild

  n1.Key = n1New.Key
  n1.LeftChild = n1New.LeftChild
  n1.RightChild = n1New.RightChild  
}

func (h* Heap) heapAdd(key int) {
  todo
}

func heapifyNode(node *Node) {
  if (node == nil) {
    return
  }

  if node.LeftChild != nil && node.Key > node.LeftChild.Key {
    k = node.LeftChild.Key
    node.LeftChild.Key, node.key = node.Key, node.LeftChild.Key
    node.Key = k
  }
  
  if node.RightChild != nil && node.Key > node.RightChild.Key {
    k = node.RightChild.Key
    node.RightChild.Key, node.key = node.Key, node.RightChild.Key
    node.Key = k
  }

}

func (h *Heap) heapify() {
  if (h.head == nil ) || (h.Len) == 0 {
    return nil
  }

  for i := range(h.Len) {
    root = h.head

    while (root != nil) {
      if root.LeftChild != nil && root.Key > root.LeftChild.Key {
        k = root.LeftChild.Key
        root.LeftChild.Key, root.key = root.Key, root.LeftChild.Key
        root.Key = k
      }
      
      if root.RightChild != nil && root.Key > root.RightChild.Key {
        k = root.RightChild.Key
        root.RightChild.Key, root.key = root.Key, root.RightChild.Key
        root.Key = k
      }
    }

  }



}

func (h *Heap) heapPop() (int, err) {
  if (h.Len == 0) || (h.head == nil) {
    return nil, fmt.Error("Empty heap: nothing to pop")
  } 

  leftChild = h.LeftChild
  rightChild = h.RightChild

  root = h.head
  result = h.head.Key

  while (root != nil) {
    if (root.LeftChild == nil) {
      h.head.Key = root.Key
    }
    root = root.LeftChild
  }

  


  


}

