package piscine

// import "fmt"

// type TreeNode struct {
// 	Data                string
// 	Left, Right, Parent *TreeNode
// }

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}

// 	if data < root.Data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 		if root.Left != nil {
// 			root.Left.Parent = root
// 		}
// 	} else {
// 		root.Right = BTreeInsertData(root.Right, data)
// 		if root.Right != nil {
// 			root.Right.Parent = root
// 		}
// 	}
// 	return root
// }

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if elem > root.Data {
		if root.Right.Data == elem {
			return root.Right
		}
	} else {
		if root.Left.Data == elem {
			return root.Left
		}
	}
	return nil
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	selected := BTreeSearchItem(root, "7")
// 	fmt.Print("Item selected -> ")
// 	if selected != nil {
// 		fmt.Println(selected.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Parent of selected item -> ")
// 	if selected.Parent != nil {
// 		fmt.Println(selected.Parent.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Left child of selected item -> ")
// 	if selected.Left != nil {
// 		fmt.Println(selected.Left.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Right child of selected item -> ")
// 	if selected.Right != nil {
// 		fmt.Println(selected.Right.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}
// }
