package piscine

// import "fmt"

// type TreeNode struct {
// 	Left, Right *TreeNode
// 	Data        string
// }

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}
// 	if data < root.Data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 	} else {
// 		root.Right = BTreeInsertData(root.Right, data)
// 	}
// 	return root
// }

// func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
// 	if root != nil {
// 		if root.Left != nil {
// 			BTreeApplyInorder(root.Left, f)
// 			// f(root.Left.Data)
// 		}
// 		f(root.Data)

// 		if root.Right != nil {
// 			BTreeApplyInorder(root.Right, f)
// 			// f(root.Right.Data)
// 		}
// 	}
// }

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	// BTreeInsertData(root, "2")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	BTreeApplyInorder(root, fmt.Println)
// }
