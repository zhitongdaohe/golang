package list

import (
	"fmt"
	"golang/pkg"
)

func ReverseList(head *pkg.ListNode) *pkg.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cursor, next := head.Next, head.Next.Next

	head.Next = nil
	for cursor != nil {
		cursor.Next = head
		head = cursor
		cursor = next
		if cursor != nil {
			next = cursor.Next
		}
	}

	return head
}

func PrintList(head *pkg.ListNode) {
	fmt.Println("PrintList_begin")
	if head == nil {
		fmt.Println("empty list")
		return
	}

	cursor := head
	for cursor != nil {
		if cursor.Next != nil {
			fmt.Printf("%d ", cursor.Val)
		} else {
			fmt.Printf("%d\n", cursor.Val)
		}
		cursor = cursor.Next
	}

	fmt.Println("PrintList_end")
}
