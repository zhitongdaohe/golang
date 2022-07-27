package list

import (
	"fmt"
	"golang/pkg"
	"testing"
)

func TestPrintList(t *testing.T) {
	fmt.Println("TestPrintList")
	l1, l2, l3 := new(pkg.ListNode), new(pkg.ListNode), new(pkg.ListNode)
	l1.Val = 1
	l1.Next = l2
	l2.Val = 2
	l2.Next = l3
	l3.Val = 3
	l3.Next = nil

	type args struct {
		head *pkg.ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				head: l1,
			},
		},
	}
	for _, tt := range tests {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			PrintList(tt.args.head)
		})
	}
}
