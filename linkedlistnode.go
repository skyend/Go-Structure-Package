package structure

type LinkedListNode struct {
     data interface{}
     prev *LinkedListNode
     next *LinkedListNode
}

func NewLinkedListNode() *LinkedListNode {
     return new(LinkedListNode);
}

