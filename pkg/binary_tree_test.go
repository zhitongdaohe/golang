package pkg

import (
	"reflect"
	"testing"
)

func buildBinaryTree() *BinaryTreeNode {
	node1 := &BinaryTreeNode{1, nil, nil}
	node2 := &BinaryTreeNode{2, nil, nil}
	node3 := &BinaryTreeNode{3, nil, nil}
	node4 := &BinaryTreeNode{4, nil, nil}
	node5 := &BinaryTreeNode{5, nil, nil}
	node6 := &BinaryTreeNode{6, nil, nil}
	node7 := &BinaryTreeNode{7, nil, nil}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	return node1
}

func TestPreorder(t *testing.T) {
	type args struct {
		root *BinaryTreeNode
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "test1_preorder",
			args: args{
				root: buildBinaryTree(),
			},
			wantResult: []int{1, 2, 4, 5, 3, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Preorder(tt.args.root); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Preorder() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestInorder(t *testing.T) {
	type args struct {
		root *BinaryTreeNode
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "test1_ineorder",
			args: args{
				root: buildBinaryTree(),
			},
			wantResult: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Inorder(tt.args.root); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Inorder() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestPostorder(t *testing.T) {
	type args struct {
		root *BinaryTreeNode
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "test1_posteorder",
			args: args{
				root: buildBinaryTree(),
			},
			wantResult: []int{4, 5, 2, 6, 7, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Postorder(tt.args.root); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Postorder() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
